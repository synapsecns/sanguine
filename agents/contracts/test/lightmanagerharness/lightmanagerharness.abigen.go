// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lightmanagerharness

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205d9fc98cb0389b4946c85a9cbcc83e8c30c457c364625b1f79455117dfcfc00064736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"b269681d": "destination()",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"ddeffa66": "getStoredSignature(uint256)",
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

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_AgentManager *AgentManagerCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _AgentManager.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_AgentManager *AgentManagerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _AgentManager.Contract.GetStoredSignature(&_AgentManager.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_AgentManager *AgentManagerCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _AgentManager.Contract.GetStoredSignature(&_AgentManager.CallOpts, index)
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

// AgentManagerAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the AgentManager contract.
type AgentManagerAttestationAcceptedIterator struct {
	Event *AgentManagerAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerAttestationAccepted represents a AttestationAccepted event raised by the AgentManager contract.
type AgentManagerAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_AgentManager *AgentManagerFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*AgentManagerAttestationAcceptedIterator, error) {

	logs, sub, err := _AgentManager.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &AgentManagerAttestationAcceptedIterator{contract: _AgentManager.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_AgentManager *AgentManagerFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *AgentManagerAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _AgentManager.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerAttestationAccepted)
				if err := _AgentManager.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManager *AgentManagerFilterer) ParseAttestationAccepted(log types.Log) (*AgentManagerAttestationAccepted, error) {
	event := new(AgentManagerAttestationAccepted)
	if err := _AgentManager.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"}]",
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

// AgentManagerEventsAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the AgentManagerEvents contract.
type AgentManagerEventsAttestationAcceptedIterator struct {
	Event *AgentManagerEventsAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerEventsAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerEventsAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerEventsAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerEventsAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerEventsAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerEventsAttestationAccepted represents a AttestationAccepted event raised by the AgentManagerEvents contract.
type AgentManagerEventsAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*AgentManagerEventsAttestationAcceptedIterator, error) {

	logs, sub, err := _AgentManagerEvents.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &AgentManagerEventsAttestationAcceptedIterator{contract: _AgentManagerEvents.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_AgentManagerEvents *AgentManagerEventsFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *AgentManagerEventsAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _AgentManagerEvents.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerEventsAttestationAccepted)
				if err := _AgentManagerEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerEvents *AgentManagerEventsFilterer) ParseAttestationAccepted(log types.Log) (*AgentManagerEventsAttestationAccepted, error) {
	event := new(AgentManagerEventsAttestationAccepted)
	if err := _AgentManagerEvents.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// AgentManagerHarnessMetaData contains all meta data concerning the AgentManagerHarness contract.
var AgentManagerHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"openDisputeExposed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"remoteMockFunc\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"sensitiveMockFunc\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"sensitiveMockFuncOver32Bytes\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sensitiveMockFuncVoid\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"slashAgentExposed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"b269681d": "destination()",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"ddeffa66": "getStoredSignature(uint256)",
		"8d3638f4": "localDomain()",
		"c166ca67": "openDisputeExposed(address,address)",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"a149352c": "remoteMockFunc(uint32,uint256,bytes32)",
		"715018a6": "renounceOwnership()",
		"127a2c9d": "sensitiveMockFunc(address,uint8,bytes32)",
		"0e6bfcd5": "sensitiveMockFuncOver32Bytes(uint16,bytes4,bytes32)",
		"c9f1a03f": "sensitiveMockFuncVoid(uint16,bytes4,bytes32)",
		"69978b0d": "slashAgentExposed(uint32,address,address)",
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

// AgentManagerHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use AgentManagerHarnessMetaData.ABI instead.
var AgentManagerHarnessABI = AgentManagerHarnessMetaData.ABI

// Deprecated: Use AgentManagerHarnessMetaData.Sigs instead.
// AgentManagerHarnessFuncSigs maps the 4-byte function signature to its string representation.
var AgentManagerHarnessFuncSigs = AgentManagerHarnessMetaData.Sigs

// AgentManagerHarness is an auto generated Go binding around an Ethereum contract.
type AgentManagerHarness struct {
	AgentManagerHarnessCaller     // Read-only binding to the contract
	AgentManagerHarnessTransactor // Write-only binding to the contract
	AgentManagerHarnessFilterer   // Log filterer for contract events
}

// AgentManagerHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type AgentManagerHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AgentManagerHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AgentManagerHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AgentManagerHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AgentManagerHarnessSession struct {
	Contract     *AgentManagerHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AgentManagerHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AgentManagerHarnessCallerSession struct {
	Contract *AgentManagerHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// AgentManagerHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AgentManagerHarnessTransactorSession struct {
	Contract     *AgentManagerHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// AgentManagerHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type AgentManagerHarnessRaw struct {
	Contract *AgentManagerHarness // Generic contract binding to access the raw methods on
}

// AgentManagerHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AgentManagerHarnessCallerRaw struct {
	Contract *AgentManagerHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// AgentManagerHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AgentManagerHarnessTransactorRaw struct {
	Contract *AgentManagerHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAgentManagerHarness creates a new instance of AgentManagerHarness, bound to a specific deployed contract.
func NewAgentManagerHarness(address common.Address, backend bind.ContractBackend) (*AgentManagerHarness, error) {
	contract, err := bindAgentManagerHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarness{AgentManagerHarnessCaller: AgentManagerHarnessCaller{contract: contract}, AgentManagerHarnessTransactor: AgentManagerHarnessTransactor{contract: contract}, AgentManagerHarnessFilterer: AgentManagerHarnessFilterer{contract: contract}}, nil
}

// NewAgentManagerHarnessCaller creates a new read-only instance of AgentManagerHarness, bound to a specific deployed contract.
func NewAgentManagerHarnessCaller(address common.Address, caller bind.ContractCaller) (*AgentManagerHarnessCaller, error) {
	contract, err := bindAgentManagerHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessCaller{contract: contract}, nil
}

// NewAgentManagerHarnessTransactor creates a new write-only instance of AgentManagerHarness, bound to a specific deployed contract.
func NewAgentManagerHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*AgentManagerHarnessTransactor, error) {
	contract, err := bindAgentManagerHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessTransactor{contract: contract}, nil
}

// NewAgentManagerHarnessFilterer creates a new log filterer instance of AgentManagerHarness, bound to a specific deployed contract.
func NewAgentManagerHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*AgentManagerHarnessFilterer, error) {
	contract, err := bindAgentManagerHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessFilterer{contract: contract}, nil
}

// bindAgentManagerHarness binds a generic wrapper to an already deployed contract.
func bindAgentManagerHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AgentManagerHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManagerHarness *AgentManagerHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManagerHarness.Contract.AgentManagerHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManagerHarness *AgentManagerHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.AgentManagerHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManagerHarness *AgentManagerHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.AgentManagerHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AgentManagerHarness *AgentManagerHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AgentManagerHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AgentManagerHarness *AgentManagerHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AgentManagerHarness *AgentManagerHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.contract.Transact(opts, method, params...)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_AgentManagerHarness *AgentManagerHarnessCaller) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "agentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_AgentManagerHarness *AgentManagerHarnessSession) AgentRoot() ([32]byte, error) {
	return _AgentManagerHarness.Contract.AgentRoot(&_AgentManagerHarness.CallOpts)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) AgentRoot() ([32]byte, error) {
	return _AgentManagerHarness.Contract.AgentRoot(&_AgentManagerHarness.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_AgentManagerHarness *AgentManagerHarnessCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_AgentManagerHarness *AgentManagerHarnessSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _AgentManagerHarness.Contract.AgentStatus(&_AgentManagerHarness.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _AgentManagerHarness.Contract.AgentStatus(&_AgentManagerHarness.CallOpts, agent)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessSession) Destination() (common.Address, error) {
	return _AgentManagerHarness.Contract.Destination(&_AgentManagerHarness.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) Destination() (common.Address, error) {
	return _AgentManagerHarness.Contract.Destination(&_AgentManagerHarness.CallOpts)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_AgentManagerHarness *AgentManagerHarnessCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (Dispute, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(Dispute), err
	}

	out0 := *abi.ConvertType(out[0], new(Dispute)).(*Dispute)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_AgentManagerHarness *AgentManagerHarnessSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _AgentManagerHarness.Contract.DisputeStatus(&_AgentManagerHarness.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _AgentManagerHarness.Contract.DisputeStatus(&_AgentManagerHarness.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_AgentManagerHarness *AgentManagerHarnessCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "getAgent", index)

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
func (_AgentManagerHarness *AgentManagerHarnessSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _AgentManagerHarness.Contract.GetAgent(&_AgentManagerHarness.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _AgentManagerHarness.Contract.GetAgent(&_AgentManagerHarness.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_AgentManagerHarness *AgentManagerHarnessCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_AgentManagerHarness *AgentManagerHarnessSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _AgentManagerHarness.Contract.GetStoredSignature(&_AgentManagerHarness.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _AgentManagerHarness.Contract.GetStoredSignature(&_AgentManagerHarness.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentManagerHarness *AgentManagerHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentManagerHarness *AgentManagerHarnessSession) LocalDomain() (uint32, error) {
	return _AgentManagerHarness.Contract.LocalDomain(&_AgentManagerHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) LocalDomain() (uint32, error) {
	return _AgentManagerHarness.Contract.LocalDomain(&_AgentManagerHarness.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessSession) Origin() (common.Address, error) {
	return _AgentManagerHarness.Contract.Origin(&_AgentManagerHarness.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) Origin() (common.Address, error) {
	return _AgentManagerHarness.Contract.Origin(&_AgentManagerHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessSession) Owner() (common.Address, error) {
	return _AgentManagerHarness.Contract.Owner(&_AgentManagerHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) Owner() (common.Address, error) {
	return _AgentManagerHarness.Contract.Owner(&_AgentManagerHarness.CallOpts)
}

// RemoteMockFunc is a free data retrieval call binding the contract method 0xa149352c.
//
// Solidity: function remoteMockFunc(uint32 , uint256 , bytes32 ) view returns(bytes4)
func (_AgentManagerHarness *AgentManagerHarnessCaller) RemoteMockFunc(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int, arg2 [32]byte) ([4]byte, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "remoteMockFunc", arg0, arg1, arg2)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// RemoteMockFunc is a free data retrieval call binding the contract method 0xa149352c.
//
// Solidity: function remoteMockFunc(uint32 , uint256 , bytes32 ) view returns(bytes4)
func (_AgentManagerHarness *AgentManagerHarnessSession) RemoteMockFunc(arg0 uint32, arg1 *big.Int, arg2 [32]byte) ([4]byte, error) {
	return _AgentManagerHarness.Contract.RemoteMockFunc(&_AgentManagerHarness.CallOpts, arg0, arg1, arg2)
}

// RemoteMockFunc is a free data retrieval call binding the contract method 0xa149352c.
//
// Solidity: function remoteMockFunc(uint32 , uint256 , bytes32 ) view returns(bytes4)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) RemoteMockFunc(arg0 uint32, arg1 *big.Int, arg2 [32]byte) ([4]byte, error) {
	return _AgentManagerHarness.Contract.RemoteMockFunc(&_AgentManagerHarness.CallOpts, arg0, arg1, arg2)
}

// SensitiveMockFunc is a free data retrieval call binding the contract method 0x127a2c9d.
//
// Solidity: function sensitiveMockFunc(address , uint8 , bytes32 data) view returns(bytes32)
func (_AgentManagerHarness *AgentManagerHarnessCaller) SensitiveMockFunc(opts *bind.CallOpts, arg0 common.Address, arg1 uint8, data [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "sensitiveMockFunc", arg0, arg1, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SensitiveMockFunc is a free data retrieval call binding the contract method 0x127a2c9d.
//
// Solidity: function sensitiveMockFunc(address , uint8 , bytes32 data) view returns(bytes32)
func (_AgentManagerHarness *AgentManagerHarnessSession) SensitiveMockFunc(arg0 common.Address, arg1 uint8, data [32]byte) ([32]byte, error) {
	return _AgentManagerHarness.Contract.SensitiveMockFunc(&_AgentManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFunc is a free data retrieval call binding the contract method 0x127a2c9d.
//
// Solidity: function sensitiveMockFunc(address , uint8 , bytes32 data) view returns(bytes32)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) SensitiveMockFunc(arg0 common.Address, arg1 uint8, data [32]byte) ([32]byte, error) {
	return _AgentManagerHarness.Contract.SensitiveMockFunc(&_AgentManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFuncOver32Bytes is a free data retrieval call binding the contract method 0x0e6bfcd5.
//
// Solidity: function sensitiveMockFuncOver32Bytes(uint16 , bytes4 , bytes32 data) view returns(bytes4, bytes32)
func (_AgentManagerHarness *AgentManagerHarnessCaller) SensitiveMockFuncOver32Bytes(opts *bind.CallOpts, arg0 uint16, arg1 [4]byte, data [32]byte) ([4]byte, [32]byte, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "sensitiveMockFuncOver32Bytes", arg0, arg1, data)

	if err != nil {
		return *new([4]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// SensitiveMockFuncOver32Bytes is a free data retrieval call binding the contract method 0x0e6bfcd5.
//
// Solidity: function sensitiveMockFuncOver32Bytes(uint16 , bytes4 , bytes32 data) view returns(bytes4, bytes32)
func (_AgentManagerHarness *AgentManagerHarnessSession) SensitiveMockFuncOver32Bytes(arg0 uint16, arg1 [4]byte, data [32]byte) ([4]byte, [32]byte, error) {
	return _AgentManagerHarness.Contract.SensitiveMockFuncOver32Bytes(&_AgentManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFuncOver32Bytes is a free data retrieval call binding the contract method 0x0e6bfcd5.
//
// Solidity: function sensitiveMockFuncOver32Bytes(uint16 , bytes4 , bytes32 data) view returns(bytes4, bytes32)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) SensitiveMockFuncOver32Bytes(arg0 uint16, arg1 [4]byte, data [32]byte) ([4]byte, [32]byte, error) {
	return _AgentManagerHarness.Contract.SensitiveMockFuncOver32Bytes(&_AgentManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFuncVoid is a free data retrieval call binding the contract method 0xc9f1a03f.
//
// Solidity: function sensitiveMockFuncVoid(uint16 , bytes4 , bytes32 ) view returns()
func (_AgentManagerHarness *AgentManagerHarnessCaller) SensitiveMockFuncVoid(opts *bind.CallOpts, arg0 uint16, arg1 [4]byte, arg2 [32]byte) error {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "sensitiveMockFuncVoid", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// SensitiveMockFuncVoid is a free data retrieval call binding the contract method 0xc9f1a03f.
//
// Solidity: function sensitiveMockFuncVoid(uint16 , bytes4 , bytes32 ) view returns()
func (_AgentManagerHarness *AgentManagerHarnessSession) SensitiveMockFuncVoid(arg0 uint16, arg1 [4]byte, arg2 [32]byte) error {
	return _AgentManagerHarness.Contract.SensitiveMockFuncVoid(&_AgentManagerHarness.CallOpts, arg0, arg1, arg2)
}

// SensitiveMockFuncVoid is a free data retrieval call binding the contract method 0xc9f1a03f.
//
// Solidity: function sensitiveMockFuncVoid(uint16 , bytes4 , bytes32 ) view returns()
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) SensitiveMockFuncVoid(arg0 uint16, arg1 [4]byte, arg2 [32]byte) error {
	return _AgentManagerHarness.Contract.SensitiveMockFuncVoid(&_AgentManagerHarness.CallOpts, arg0, arg1, arg2)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentManagerHarness *AgentManagerHarnessCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AgentManagerHarness.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentManagerHarness *AgentManagerHarnessSession) Version() (string, error) {
	return _AgentManagerHarness.Contract.Version(&_AgentManagerHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_AgentManagerHarness *AgentManagerHarnessCallerSession) Version() (string, error) {
	return _AgentManagerHarness.Contract.Version(&_AgentManagerHarness.CallOpts)
}

// OpenDisputeExposed is a paid mutator transaction binding the contract method 0xc166ca67.
//
// Solidity: function openDisputeExposed(address guard, address notary) returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactor) OpenDisputeExposed(opts *bind.TransactOpts, guard common.Address, notary common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "openDisputeExposed", guard, notary)
}

// OpenDisputeExposed is a paid mutator transaction binding the contract method 0xc166ca67.
//
// Solidity: function openDisputeExposed(address guard, address notary) returns()
func (_AgentManagerHarness *AgentManagerHarnessSession) OpenDisputeExposed(guard common.Address, notary common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.OpenDisputeExposed(&_AgentManagerHarness.TransactOpts, guard, notary)
}

// OpenDisputeExposed is a paid mutator transaction binding the contract method 0xc166ca67.
//
// Solidity: function openDisputeExposed(address guard, address notary) returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) OpenDisputeExposed(guard common.Address, notary common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.OpenDisputeExposed(&_AgentManagerHarness.TransactOpts, guard, notary)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentManagerHarness *AgentManagerHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.RenounceOwnership(&_AgentManagerHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.RenounceOwnership(&_AgentManagerHarness.TransactOpts)
}

// SlashAgentExposed is a paid mutator transaction binding the contract method 0x69978b0d.
//
// Solidity: function slashAgentExposed(uint32 domain, address agent, address prover) returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactor) SlashAgentExposed(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "slashAgentExposed", domain, agent, prover)
}

// SlashAgentExposed is a paid mutator transaction binding the contract method 0x69978b0d.
//
// Solidity: function slashAgentExposed(uint32 domain, address agent, address prover) returns()
func (_AgentManagerHarness *AgentManagerHarnessSession) SlashAgentExposed(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SlashAgentExposed(&_AgentManagerHarness.TransactOpts, domain, agent, prover)
}

// SlashAgentExposed is a paid mutator transaction binding the contract method 0x69978b0d.
//
// Solidity: function slashAgentExposed(uint32 domain, address agent, address prover) returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) SlashAgentExposed(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SlashAgentExposed(&_AgentManagerHarness.TransactOpts, domain, agent, prover)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SubmitStateReportWithAttestation(&_AgentManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SubmitStateReportWithAttestation(&_AgentManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SubmitStateReportWithSnapshot(&_AgentManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SubmitStateReportWithSnapshot(&_AgentManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SubmitStateReportWithSnapshotProof(&_AgentManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.SubmitStateReportWithSnapshotProof(&_AgentManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentManagerHarness *AgentManagerHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.TransferOwnership(&_AgentManagerHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.TransferOwnership(&_AgentManagerHarness.TransactOpts, newOwner)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_AgentManagerHarness *AgentManagerHarnessSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyReceipt(&_AgentManagerHarness.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyReceipt(&_AgentManagerHarness.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) VerifyStateReport(opts *bind.TransactOpts, srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "verifyStateReport", srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_AgentManagerHarness *AgentManagerHarnessSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateReport(&_AgentManagerHarness.TransactOpts, srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateReport(&_AgentManagerHarness.TransactOpts, srPayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateWithAttestation(&_AgentManagerHarness.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateWithAttestation(&_AgentManagerHarness.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateWithSnapshot(&_AgentManagerHarness.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateWithSnapshot(&_AgentManagerHarness.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateWithSnapshotProof(&_AgentManagerHarness.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_AgentManagerHarness *AgentManagerHarnessTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _AgentManagerHarness.Contract.VerifyStateWithSnapshotProof(&_AgentManagerHarness.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// AgentManagerHarnessAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the AgentManagerHarness contract.
type AgentManagerHarnessAttestationAcceptedIterator struct {
	Event *AgentManagerHarnessAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessAttestationAccepted represents a AttestationAccepted event raised by the AgentManagerHarness contract.
type AgentManagerHarnessAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*AgentManagerHarnessAttestationAcceptedIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessAttestationAcceptedIterator{contract: _AgentManagerHarness.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessAttestationAccepted)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseAttestationAccepted(log types.Log) (*AgentManagerHarnessAttestationAccepted, error) {
	event := new(AgentManagerHarnessAttestationAccepted)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessDisputeUpdatedIterator is returned from FilterDisputeUpdated and is used to iterate over the raw logs and unpacked data for DisputeUpdated events raised by the AgentManagerHarness contract.
type AgentManagerHarnessDisputeUpdatedIterator struct {
	Event *AgentManagerHarnessDisputeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessDisputeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessDisputeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessDisputeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessDisputeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessDisputeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessDisputeUpdated represents a DisputeUpdated event raised by the AgentManagerHarness contract.
type AgentManagerHarnessDisputeUpdated struct {
	Agent   common.Address
	Dispute Dispute
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeUpdated is a free log retrieval operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterDisputeUpdated(opts *bind.FilterOpts) (*AgentManagerHarnessDisputeUpdatedIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessDisputeUpdatedIterator{contract: _AgentManagerHarness.contract, event: "DisputeUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeUpdated is a free log subscription operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchDisputeUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessDisputeUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessDisputeUpdated)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseDisputeUpdated(log types.Log) (*AgentManagerHarnessDisputeUpdated, error) {
	event := new(AgentManagerHarnessDisputeUpdated)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AgentManagerHarness contract.
type AgentManagerHarnessInitializedIterator struct {
	Event *AgentManagerHarnessInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessInitialized represents a Initialized event raised by the AgentManagerHarness contract.
type AgentManagerHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*AgentManagerHarnessInitializedIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessInitializedIterator{contract: _AgentManagerHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessInitialized)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseInitialized(log types.Log) (*AgentManagerHarnessInitialized, error) {
	event := new(AgentManagerHarnessInitialized)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidReceiptIterator struct {
	Event *AgentManagerHarnessInvalidReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessInvalidReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessInvalidReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessInvalidReceipt represents a InvalidReceipt event raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*AgentManagerHarnessInvalidReceiptIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessInvalidReceiptIterator{contract: _AgentManagerHarness.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessInvalidReceipt)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseInvalidReceipt(log types.Log) (*AgentManagerHarnessInvalidReceipt, error) {
	event := new(AgentManagerHarnessInvalidReceipt)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidStateReportIterator struct {
	Event *AgentManagerHarnessInvalidStateReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessInvalidStateReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessInvalidStateReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessInvalidStateReport represents a InvalidStateReport event raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*AgentManagerHarnessInvalidStateReportIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessInvalidStateReportIterator{contract: _AgentManagerHarness.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessInvalidStateReport)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseInvalidStateReport(log types.Log) (*AgentManagerHarnessInvalidStateReport, error) {
	event := new(AgentManagerHarnessInvalidStateReport)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidStateWithAttestationIterator struct {
	Event *AgentManagerHarnessInvalidStateWithAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessInvalidStateWithAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessInvalidStateWithAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*AgentManagerHarnessInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessInvalidStateWithAttestationIterator{contract: _AgentManagerHarness.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessInvalidStateWithAttestation)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseInvalidStateWithAttestation(log types.Log) (*AgentManagerHarnessInvalidStateWithAttestation, error) {
	event := new(AgentManagerHarnessInvalidStateWithAttestation)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidStateWithSnapshotIterator struct {
	Event *AgentManagerHarnessInvalidStateWithSnapshot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessInvalidStateWithSnapshot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessInvalidStateWithSnapshot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the AgentManagerHarness contract.
type AgentManagerHarnessInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*AgentManagerHarnessInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessInvalidStateWithSnapshotIterator{contract: _AgentManagerHarness.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessInvalidStateWithSnapshot)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*AgentManagerHarnessInvalidStateWithSnapshot, error) {
	event := new(AgentManagerHarnessInvalidStateWithSnapshot)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AgentManagerHarness contract.
type AgentManagerHarnessOwnershipTransferredIterator struct {
	Event *AgentManagerHarnessOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the AgentManagerHarness contract.
type AgentManagerHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AgentManagerHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessOwnershipTransferredIterator{contract: _AgentManagerHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessOwnershipTransferred)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*AgentManagerHarnessOwnershipTransferred, error) {
	event := new(AgentManagerHarnessOwnershipTransferred)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the AgentManagerHarness contract.
type AgentManagerHarnessRootUpdatedIterator struct {
	Event *AgentManagerHarnessRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessRootUpdated represents a RootUpdated event raised by the AgentManagerHarness contract.
type AgentManagerHarnessRootUpdated struct {
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*AgentManagerHarnessRootUpdatedIterator, error) {

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessRootUpdatedIterator{contract: _AgentManagerHarness.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessRootUpdated) (event.Subscription, error) {

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessRootUpdated)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "RootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseRootUpdated(log types.Log) (*AgentManagerHarnessRootUpdated, error) {
	event := new(AgentManagerHarnessRootUpdated)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AgentManagerHarnessStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the AgentManagerHarness contract.
type AgentManagerHarnessStatusUpdatedIterator struct {
	Event *AgentManagerHarnessStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AgentManagerHarnessStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AgentManagerHarnessStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AgentManagerHarnessStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AgentManagerHarnessStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AgentManagerHarnessStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AgentManagerHarnessStatusUpdated represents a StatusUpdated event raised by the AgentManagerHarness contract.
type AgentManagerHarnessStatusUpdated struct {
	Flag   uint8
	Domain uint32
	Agent  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) FilterStatusUpdated(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*AgentManagerHarnessStatusUpdatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManagerHarness.contract.FilterLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &AgentManagerHarnessStatusUpdatedIterator{contract: _AgentManagerHarness.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_AgentManagerHarness *AgentManagerHarnessFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *AgentManagerHarnessStatusUpdated, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _AgentManagerHarness.contract.WatchLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AgentManagerHarnessStatusUpdated)
				if err := _AgentManagerHarness.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_AgentManagerHarness *AgentManagerHarnessFilterer) ParseStatusUpdated(log types.Log) (*AgentManagerHarnessStatusUpdated, error) {
	event := new(AgentManagerHarnessStatusUpdated)
	if err := _AgentManagerHarness.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220418882ba19eeda21dc232e9732391c98ca6be47ad2062cd77c7bf44db68eeeec64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204fb965d4f389633228c39c6e714640559280ca2641083e09d5b48e0be8bffeb464736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220095c76a00d2fff7ef7d7651e0d78391a66292212503031a3376b9453bc3e048e64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bc29deb98906fce6cca59a73cfd5dcf8cc29374db21ea738027c9cef7e32207664736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"ddeffa66": "getStoredSignature(uint256)",
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

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IAgentManager *IAgentManagerCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _IAgentManager.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IAgentManager *IAgentManagerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _IAgentManager.Contract.GetStoredSignature(&_IAgentManager.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_IAgentManager *IAgentManagerCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _IAgentManager.Contract.GetStoredSignature(&_IAgentManager.CallOpts, index)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"msgPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"originProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"gasLimit\",\"type\":\"uint64\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"snapRoot\",\"type\":\"bytes32\"}],\"name\":\"getAttestationNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"}],\"name\":\"isValidReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"messageStatus\",\"outputs\":[{\"internalType\":\"enumMessageStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"}],\"name\":\"receiptBody\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"32ff14d2": "execute(bytes,bytes32[],bytes32[],uint256,uint64)",
		"4f127567": "getAttestationNonce(bytes32)",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"addAgent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"completeUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getActiveAgents\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"agents\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"indexFrom\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"getLeafs\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"leafs\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"getProof\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"initiateUnstaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leafsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"attNotaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"attNonce\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rcptBodyPayload\",\"type\":\"bytes\"}],\"name\":\"passReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"remoteSlashAgent\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"submitReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidAttestation\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"}],\"name\":\"verifyAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTips\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
		"6b47b3bc": "passReceipt(uint32,uint32,uint256,bytes)",
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

// PassReceipt is a paid mutator transaction binding the contract method 0x6b47b3bc.
//
// Solidity: function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes rcptBodyPayload) returns(bool wasAccepted)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) PassReceipt(opts *bind.TransactOpts, attNotaryIndex uint32, attNonce uint32, paddedTips *big.Int, rcptBodyPayload []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "passReceipt", attNotaryIndex, attNonce, paddedTips, rcptBodyPayload)
}

// PassReceipt is a paid mutator transaction binding the contract method 0x6b47b3bc.
//
// Solidity: function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes rcptBodyPayload) returns(bool wasAccepted)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) PassReceipt(attNotaryIndex uint32, attNonce uint32, paddedTips *big.Int, rcptBodyPayload []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.PassReceipt(&_InterfaceBondingManager.TransactOpts, attNotaryIndex, attNonce, paddedTips, rcptBodyPayload)
}

// PassReceipt is a paid mutator transaction binding the contract method 0x6b47b3bc.
//
// Solidity: function passReceipt(uint32 attNotaryIndex, uint32 attNonce, uint256 paddedTips, bytes rcptBodyPayload) returns(bool wasAccepted)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactorSession) PassReceipt(attNotaryIndex uint32, attNonce uint32, paddedTips *big.Int, rcptBodyPayload []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.PassReceipt(&_InterfaceBondingManager.TransactOpts, attNotaryIndex, attNonce, paddedTips, rcptBodyPayload)
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
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_InterfaceBondingManager *InterfaceBondingManagerTransactor) SubmitSnapshot(opts *bind.TransactOpts, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.contract.Transact(opts, "submitSnapshot", snapPayload, snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
func (_InterfaceBondingManager *InterfaceBondingManagerSession) SubmitSnapshot(snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _InterfaceBondingManager.Contract.SubmitSnapshot(&_InterfaceBondingManager.TransactOpts, snapPayload, snapSignature)
}

// SubmitSnapshot is a paid mutator transaction binding the contract method 0x4bb73ea5.
//
// Solidity: function submitSnapshot(bytes snapPayload, bytes snapSignature) returns(bytes attPayload, bytes32 agentRoot, uint256[] snapGas)
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

// InterfaceDestinationMetaData contains all meta data concerning the InterfaceDestination contract.
var InterfaceDestinationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"sigIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"ChainGas[]\",\"name\":\"snapGas\",\"type\":\"uint128[]\"}],\"name\":\"acceptAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destStatus\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"snapRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"agentRootTime\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"notaryIndex\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"getGasData\",\"outputs\":[{\"internalType\":\"GasData\",\"name\":\"gasData\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"dataMaturity\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAgentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"passAgentRoot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"rootPassed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"rootPending\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"39fe2736": "acceptAttestation(uint32,uint256,bytes,bytes32,uint128[])",
		"3cf7b120": "attestationsAmount()",
		"40989152": "destStatus()",
		"29be4db2": "getAttestation(uint256)",
		"d0dd0675": "getGasData(uint32)",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"remoteWithdrawTips\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"}],\"name\":\"setAgentRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas\",\"type\":\"uint256[]\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"updateAgentStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1fa07138": "remoteWithdrawTips(uint32,uint256,address,uint256)",
		"58668176": "setAgentRoot(bytes32)",
		"6c38f723": "submitAttestation(bytes,bytes,bytes32,uint256[])",
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

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightManager.contract.Transact(opts, "submitAttestation", attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SubmitAttestation(&_InterfaceLightManager.TransactOpts, attPayload, attSignature, agentRoot, snapGas)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot, uint256[] snapGas) returns(bool wasAccepted)
func (_InterfaceLightManager *InterfaceLightManagerTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot [32]byte, snapGas []*big.Int) (*types.Transaction, error) {
	return _InterfaceLightManager.Contract.SubmitAttestation(&_InterfaceLightManager.TransactOpts, attPayload, attSignature, agentRoot, snapGas)
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

// LightManagerMetaData contains all meta data concerning the LightManager contract.
var LightManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"remoteWithdrawTips\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"}],\"name\":\"setAgentRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas_\",\"type\":\"uint256[]\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"updateAgentStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"b269681d": "destination()",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"ddeffa66": "getStoredSignature(uint256)",
		"485cc955": "initialize(address,address)",
		"8d3638f4": "localDomain()",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"1fa07138": "remoteWithdrawTips(uint32,uint256,address,uint256)",
		"715018a6": "renounceOwnership()",
		"58668176": "setAgentRoot(bytes32)",
		"6c38f723": "submitAttestation(bytes,bytes,bytes32,uint256[])",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
		"235d51b1": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes,bytes)",
		"708cdc82": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"cbd05965": "updateAgentStatus(address,(uint8,uint32,uint32),bytes32[])",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b5060405162004c1838038062004c188339810160408190526200003491620000dc565b60408051808201909152600580825264302e302e3360d81b6020830152608052818162000065565b60405180910390fd5b62000070816200010b565b60a0525063ffffffff90811660c0528216600919019050620000d55760405162461bcd60e51b815260206004820152601d60248201527f43616e2774206265206465706c6f796564206f6e2053796e436861696e00000060448201526064016200005c565b5062000133565b600060208284031215620000ef57600080fd5b815163ffffffff811681146200010457600080fd5b9392505050565b805160208083015191908110156200012d576000198160200360031b1b821691505b50919050565b60805160a05160c051614aae6200016a600039600081816103b20152610d7101526000610314015260006102f10152614aae6000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c8063708cdc82116100ee578063938b5f3211610097578063cbd0596511610071578063cbd0596514610447578063ddeffa661461045a578063dfe396751461046d578063f2fde38b1461048057600080fd5b8063938b5f321461040e578063b269681d14610421578063c25aa5851461043457600080fd5b80637be8e738116100c85780637be8e7381461039a5780638d3638f4146103ad5780638da5cb5b146103e957600080fd5b8063708cdc821461036c578063715018a61461037f57806377ec5c101461038757600080fd5b80632de5aaf71161015b578063485cc95511610135578063485cc955146102d057806354fd4d50146102e557806358668176146103465780636c38f7231461035957600080fd5b80632de5aaf7146102785780633463d1b11461029957806336cba43c146102b957600080fd5b8063213a6ddb1161018c578063213a6ddb14610232578063235d51b11461024557806328f3fac91461025857600080fd5b80630db27e77146101b35780631fa07138146101db578063200f6b661461021f575b600080fd5b6101c66101c1366004613d59565b610493565b60405190151581526020015b60405180910390f35b6101ee6101e9366004613e5c565b61051d565b6040517fffffffff0000000000000000000000000000000000000000000000000000000090911681526020016101d2565b6101c661022d366004613ea2565b6106d3565b6101c6610240366004613f34565b61086c565b6101c6610253366004613fa1565b61096d565b61026b610266366004614029565b610ab8565b6040516101d291906140b3565b61028b6102863660046140c1565b610b3e565b6040516101d29291906140da565b6102ac6102a7366004614029565b610b8b565b6040516101d2919061412f565b6102c260fb5481565b6040519081526020016101d2565b6102e36102de36600461413d565b610c26565b005b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516101d291906141b6565b6102e36103543660046140c1565b610ce1565b6101c66103673660046141c9565b610d47565b6101c661037a3660046142af565b610fcc565b6102e36110fd565b6101c6610395366004614366565b611159565b6101c66103a83660046143b5565b6111c8565b6103d47f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff90911681526020016101d2565b6033546001600160a01b03165b6040516001600160a01b0390911681526020016101d2565b60c9546103f6906001600160a01b031681565b60ca546103f6906001600160a01b031681565b6101c661044236600461441a565b6112f3565b6102e361045536600461447e565b6113fd565b6103396104683660046140c1565b6116b7565b6101c661047b36600461441a565b611766565b6102e361048e366004614029565b611844565b60008061049f87611923565b90506000806104ae838961193c565b915091506104bb826119cf565b60006104c687611a8c565b90506000806104d58389611a9f565b915091506104e282611b27565b6104f6838e6104f089611b95565b8d611ba8565b61050a848660400151838560400151611c96565b5060019c9b505050505050505050505050565b60ca546000906001600160a01b0316331461057f5760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e000000000000000000000000000000000000000060448201526064015b60405180910390fd5b63ffffffff8516600a146105d55760405162461bcd60e51b815260206004820152600e60248201527f2173796e61707365446f6d61696e0000000000000000000000000000000000006044820152606401610576565b620151808410156106285760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610576565b60c9546040517f4e04e7a70000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526024820185905290911690634e04e7a790604401600060405180830381600087803b15801561068f57600080fd5b505af11580156106a3573d6000803e3d6000fd5b507f1fa071380000000000000000000000000000000000000000000000000000000093505050505b949350505050565b6000806106df84611a8c565b90506000806106ee8386611a9f565b915091506106fb82611b27565b600061070688611e14565b905061071184611e27565b61071a82611e38565b146107675760405162461bcd60e51b815260206004820181905260248201527f4174746573746174696f6e206e6f74206d61746368657320736e617073686f746044820152606401610576565b600061078161077c610779848d611f11565b90565b611fb2565b60c9546040517fa9dcf22d0000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063a9dcf22d906107cb9084906004016141b6565b602060405180830381865afa1580156107e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080c919061451c565b95508561085f577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a828a8a604051610848949392919061453e565b60405180910390a161085f84602001518433611ff1565b5050505050949350505050565b60008061087884611e14565b905060008061088783866120fb565b9150915061089482611b27565b60c9546001600160a01b031663a9dcf22d6108b561077c610779878c611f11565b6040518263ffffffff1660e01b81526004016108d191906141b6565b602060405180830381865afa1580156108ee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610912919061451c565b935083610963577f8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd187878760405161094c9392919061457d565b60405180910390a161096382602001518233611ff1565b5050509392505050565b60008061097987611923565b9050600080610988838961193c565b91509150600061099788611e14565b90506109b56109a585611b95565b6109af838e611f11565b9061212f565b610a015760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610576565b610a0a836119cf565b6000610a1588611a8c565b9050600080610a24838a611a9f565b91509150610a3182611b27565b610a3a83611e27565b610a4385611e38565b14610a905760405162461bcd60e51b815260206004820181905260248201527f4174746573746174696f6e206e6f74206d61746368657320736e617073686f746044820152606401610576565b610aa4858760400151838560400151611c96565b5060019d9c50505050505050505050505050565b6040805160608101825260008082526020820181905291810191909152610ade82612150565b905060026001600160a01b038316600090815260cb602052604090205460ff166002811115610b0f57610b0f614044565b148015610b2f5750600581516005811115610b2c57610b2c614044565b14155b15610b3957600481525b919050565b60408051606081018252600080825260208201819052918101829052600083815260fd60205260409020546001600160a01b031691508115610b8657610b8382610ab8565b90505b915091565b60408051606081018252600080825260208201819052918101919091526001600160a01b038216600090815260cb6020526040908190208151606081019092528054829060ff166002811115610be357610be3614044565b6002811115610bf457610bf4614044565b81529054610100810463ffffffff1660208301526501000000000090046001600160a01b031660409091015292915050565b6000610c3260016121ec565b90508015610c6757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610c71838361233e565b610c79612401565b8015610cdc57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60ca546001600160a01b03163314610d3b5760405162461bcd60e51b815260206004820181905260248201527f4f6e6c792044657374696e6174696f6e2073657473206167656e7420726f6f746044820152606401610576565b610d4481612486565b50565b600080610d5386611a8c565b9050600080610d628388611a9f565b91509150610d6f826119cf565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff16826020015163ffffffff1614610dee5760405162461bcd60e51b815260206004820152601360248201527f57726f6e67204e6f7461727920646f6d61696e000000000000000000000000006044820152606401610576565b6001600160a01b038116600090815260cb602052604081205460ff166002811115610e1b57610e1b614044565b14610e685760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610576565b84610e7b87610e76836124ca565b6124fa565b610e8485612530565b14610ed15760405162461bcd60e51b815260206004820152600f60248201527f496e76616c696420736e617047617300000000000000000000000000000000006044820152606401610576565b6000610edc8961253e565b60ca5460408087015190517f39fe27360000000000000000000000000000000000000000000000000000000081529293506001600160a01b03909116916339fe273691610f339185908f908e9089906004016145b2565b6020604051808303816000875af1158015610f52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f76919061451c565b9550851561085f577f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea00658460200151848c8c604051610fb79493929190614634565b60405180910390a15050505050949350505050565b600080610fd886611923565b9050600080610fe7838861193c565b91509150610ff4826119cf565b6000610fff87611e14565b905060008061100e83896120fb565b91509150816020015163ffffffff1660000361106c5760405162461bcd60e51b815260206004820152601f60248201527f536e617073686f74207369676e6572206973206e6f742061204e6f74617279006044820152606401610576565b61107582611b27565b61108b61108187611b95565b6109af858f611f11565b6110d75760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610576565b6110eb848660400151838560400151611c96565b5060019b9a5050505050505050505050565b6033546001600160a01b031633146111575760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610576565b565b60008061116585612581565b90506000806111748387612594565b91509150611181826119cf565b600080611196611190866125bd565b88611a9f565b915091506111a382611b27565b6111b7838560400151838560400151611c96565b6001955050505050505b9392505050565b6000806111d484611a8c565b90506000806111e38386611a9f565b915091506111f082611b27565b60006111fb896125cb565b9050611209848b838b611ba8565b60c9546040517fa9dcf22d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a9dcf22d90611252908c906004016141b6565b602060405180830381865afa15801561126f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611293919061451c565b9450846112e6577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a8a89896040516112cf949392919061453e565b60405180910390a16112e683602001518333611ff1565b5050505095945050505050565b6000806112ff846125d9565b905060008061130e83866125ec565b9150915061131b82611b27565b60ca546040517fe2f006f70000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e2f006f7906113649089906004016141b6565b602060405180830381865afa158015611381573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a5919061451c565b9350836113f4577f4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d86866040516113dd929190614674565b60405180910390a16113f482602001518233611ff1565b50505092915050565b60408083015163ffffffff16600090815260fd60205220546001600160a01b031680158061143c5750836001600160a01b0316816001600160a01b0316145b6114885760405162461bcd60e51b815260206004820152601360248201527f496e76616c6964206167656e7420696e646578000000000000000000000000006044820152606401610576565b600061149d8460000151856020015187612615565b9050600060fb549050806114be866040015163ffffffff168487602061264b565b1461150b5760405162461bcd60e51b815260206004820152600d60248201527f496e76616c69642070726f6f66000000000000000000000000000000000000006044820152606401610576565b6001600160a01b0383166115625760408581015163ffffffff16600090815260fd6020522080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0388161790555b600081815260fc602090815260408083206001600160a01b038a1684529091529020855181548792919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018360058111156115c4576115c4614044565b021790555060208281015182546040948501517fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff90911661010063ffffffff938416027fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff1617650100000000009183169190910217909255870151875192516001600160a01b038a169391909216917f8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e9161167e91614699565b60405180910390a360058551600581111561169b5761169b614044565b036116af576116af86866040015133612709565b505050505050565b606060cc82815481106116cc576116cc6146a7565b9060005260206000200180546116e1906146d6565b80601f016020809104026020016040519081016040528092919081815260200182805461170d906146d6565b801561175a5780601f1061172f5761010080835404028352916020019161175a565b820191906000526020600020905b81548152906001019060200180831161173d57829003601f168201915b50505050509050919050565b60008061177284611923565b9050600080611781838661193c565b9150915061178e82611b27565b60c9546001600160a01b031663a9dcf22d6117ae61077c61077987611b95565b6040518263ffffffff1660e01b81526004016117ca91906141b6565b602060405180830381865afa1580156117e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061180b919061451c565b159350836113f4577f9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d86866040516113dd929190614674565b6033546001600160a01b0316331461189e5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610576565b6001600160a01b03811661191a5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610576565b610d448161286f565b6000611936611931836128d9565b6128ec565b92915050565b604080516060810182526000808252602082018190529181018290529061196b61196585612947565b84612975565b6020820151919350915063ffffffff16156119c85760405162461bcd60e51b815260206004820152601560248201527f5369676e6572206973206e6f74206120477561726400000000000000000000006044820152606401610576565b9250929050565b60015b815160058111156119e5576119e5614044565b14816020015163ffffffff16600014611a33576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611a6a565b6040518060400160405280601381526020017f4e6f7420616e20616374697665206775617264000000000000000000000000008152505b90611a885760405162461bcd60e51b815260040161057691906141b6565b5050565b6000611936611a9a836128d9565b612a1f565b6040805160608101825260008082526020820181905291810182905290611ac861196585612a76565b6020820151919350915063ffffffff166000036119c85760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610576565b600181516005811115611b3c57611b3c614044565b1480611b4a575060026119d2565b602082015163ffffffff1615611a33576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611a6a565b6000611936611ba383612aa2565b612aaf565b6000611bb383612b06565b9150508082600081518110611bca57611bca6146a7565b602002602001015114611c1f5760405162461bcd60e51b815260206004820152601260248201527f496e636f72726563742070726f6f665b305d00000000000000000000000000006044820152606401610576565b6000611c3d611c2d85611e27565b611c3686612b30565b8588612b3f565b905080611c4987611e27565b146116af5760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f740000000000000000006044820152606401610576565b6001600160a01b038416600090815260cb602052604081205460ff166002811115611cc357611cc3614044565b14611d105760405162461bcd60e51b815260206004820152601860248201527f477561726420616c726561647920696e206469737075746500000000000000006044820152606401610576565b6001600160a01b038216600090815260cb602052604081205460ff166002811115611d3d57611d3d614044565b14611d8a5760405162461bcd60e51b815260206004820152601960248201527f4e6f7461727920616c726561647920696e2064697370757465000000000000006044820152606401610576565b611dc784604051806060016040528060016002811115611dac57611dac614044565b815263ffffffff851660208201526000604090910152612bbc565b611e0482604051806060016040528060016002811115611de957611de9614044565b815263ffffffff871660208201526000604090910152612bbc565b611e0e8382612ccc565b50505050565b6000611936611e22836128d9565b612d4d565b6000611936816020845b9190612da4565b600080611e4483612eae565b905060008167ffffffffffffffff811115611e6157611e61613bd1565b604051908082528060200260200182016040528015611e8a578160200160208202803683370190505b50905060005b82811015611ed757611eaa611ea58683611f11565b612ed8565b828281518110611ebc57611ebc6146a7565b6020908102919091010152611ed081614752565b9050611e90565b50611eed81611ee86001600661478a565b612f17565b80600081518110611f0057611f006146a7565b602002602001015192505050919050565b60008281611f21600c603261479d565b611f2b90856147b0565b90506fffffffffffffffffffffffffffffffff82168110611f8e5760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610576565b611fa9611ba382611fa1600c603261479d565b85919061301a565b95945050505050565b60405180611fc3836020830161308b565b506fffffffffffffffffffffffffffffffff83166000601f8201601f19168301602001604052509052919050565b6000611ffc83612150565b905060018151600581111561201357612013614044565b1480612031575060028151600581111561202f5761202f614044565b145b801561204c57508363ffffffff16816020015163ffffffff16145b6120985760405162461bcd60e51b815260206004820152601f60248201527f536c617368696e6720636f756c64206e6f7420626520696e69746961746564006044820152606401610576565b826001600160a01b03168463ffffffff167f8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e60046040516120d99190614699565b60405180910390a36120f083826040015184612709565b611e0e848484613134565b60408051606081018252600080825260208201819052918101829052906121246119658561324a565b909590945092505050565b600061213f82613276565b613276565b61214884613276565b149392505050565b60408051606080820183526000808352602080840182905283850182905260fb54825260fc81528482206001600160a01b038716835290528390208351918201909352825491929091829060ff1660058111156121af576121af614044565b60058111156121c0576121c0614044565b8152905463ffffffff610100820481166020840152650100000000009091041660409091015292915050565b60008054610100900460ff1615612289578160ff16600114801561220f5750303b155b6122815760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610576565b506000919050565b60005460ff8084169116106123065760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610576565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff166123bb5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610576565b60c980546001600160a01b039384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560ca8054929093169116179055565b600054610100900460ff1661247e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610576565b6111576132a1565b8060fb5414610d445760fb8190556040518181527f2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa49060200160405180910390a150565b6000816040516020016124dd91906147c7565b604051602081830303815290604052805190602001209050919050565b60408051602081018490529081018290526000906060015b60405160208183030381529060405280519060200120905092915050565b600061193660208084611e31565b60cc80546001810182556000919091527f47197230e1e4b29fc0bd84d7d78966c0925452aff72a2a121538b102457e9ebe810161257b8382614855565b50919050565b600061193661258f836128d9565b613327565b604080516060810182526000808252602082018190529181018290529061196b6119658561337e565b6000611936611a9a83612aa2565b6000611936611ba3836128d9565b60006119366125e7836128d9565b6133aa565b6040805160608101825260008082526020820181905291810182905290611ac861196585613401565b600083838360405160200161262c93929190614951565b6040516020818303038152906040528051906020012090509392505050565b8151600090828111156126a05760405162461bcd60e51b815260206004820152600e60248201527f50726f6f6620746f6f206c6f6e670000000000000000000000000000000000006044820152606401610576565b84915060005b818110156126dd576126d3838683815181106126c4576126c46146a7565b6020026020010151898461342d565b92506001016126a6565b50805b838110156126ff576126f5836000898461342d565b92506001016126e0565b5050949350505050565b6001600160a01b038316600090815260cb60205260408082208151606081019092528054829060ff16600281111561274357612743614044565b600281111561275457612754614044565b81529054610100810463ffffffff1660208301526501000000000090046001600160a01b0316604090910152905060028151600281111561279757612797614044565b036127e45760405162461bcd60e51b815260206004820152601860248201527f4469737075746520616c7265616479207265736f6c76656400000000000000006044820152606401610576565b6001600160a01b0382166040820152600281526128018482612bbc565b602081015163ffffffff16156128615761286161283d826020015163ffffffff16600090815260fd60205260409020546001600160a01b031690565b60408051606081019091528060008152600060208201819052604090910152612bbc565b611e0e838260200151613456565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b8051600090602083016106cb81836134ad565b60006128f782613510565b6129435760405162461bcd60e51b815260206004820152601260248201527f4e6f742061207374617465207265706f727400000000000000000000000000006044820152606401610576565b5090565b60006119367f7919c62641a21cff2eb6e116b4dc34ce70919052c470953e4621535c155ccbc8835b90613562565b60408051606081018252600080825260208201819052918101829052908061299c85613585565b90506129a881856135c0565b91506129b382610ab8565b92506000835160058111156129ca576129ca614044565b03612a175760405162461bcd60e51b815260206004820152600d60248201527f556e6b6e6f776e206167656e74000000000000000000000000000000000000006044820152606401610576565b509250929050565b6000612a2a826135e4565b6129435760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610576565b60006119367f569efb4f951664b562fe9283d8f1a49928bec7335bab838210b64c85e11be59e8361296f565b6000611936826001613603565b6000612aba82613669565b6129435760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610576565b60008082612b1861213a82602461368f565b9250612b2861213a826024613603565b915050915091565b6000611936826020600461369c565b6000600182901b60408110612b965760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610576565b6000612ba287876136bd565b9050612bb1828287600661264b565b979650505050505050565b6001600160a01b038216600090815260cb60205260409020815181548392919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836002811115612c1457612c14614044565b0217905550602082015181546040938401516001600160a01b031665010000000000027fffffffffffffff0000000000000000000000000000000000000000ffffffffff63ffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff90911617179055517f4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f790612cc090849084906149c9565b60405180910390a15050565b60ca546040517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff8085166004830152831660248201526001600160a01b039091169063a2155c34906044015b600060405180830381600087803b158015612d3957600080fd5b505af11580156116af573d6000803e3d6000fd5b6000612d5882613700565b6129435760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610576565b600081600003612db6575060006111c1565b6020821115612df1576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff8416612e0e838561479d565b1115612e46576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b6000612e578660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b6000612ebc600c603261479d565b611936906fffffffffffffffffffffffffffffffff84166149e6565b6000806000612ee684612b06565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b81516001821b811115612f6c5760405162461bcd60e51b815260206004820152600e60248201527f48656967687420746f6f206c6f770000000000000000000000000000000000006044820152606401610576565b60005b82811015611e0e5760005b8281101561300b5760008160010190506000868381518110612f9e57612f9e6146a7565b602002602001015190506000858310612fb8576000612fd3565b878381518110612fca57612fca6146a7565b60200260200101515b9050612fdf8282613756565b88600186901c81518110612ff557612ff56146a7565b6020908102919091010152505050600201612f7a565b506001918201821c9101612f6f565b6000806130278560801c90565b9050613032856137a2565b8361303d868461479d565b613047919061479d565b111561307f576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611fa9848201846134ad565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c90808510156130e5576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa905080613128576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417612bb1565b60c9546040805163ffffffff861660248201526001600160a01b03858116604483015284811660648084019190915283518084039091018152608490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9d228a510000000000000000000000000000000000000000000000000000000017905291517fa1c702a7000000000000000000000000000000000000000000000000000000008152919092169163a1c702a79161320191600a91620151809190600401614a21565b60408051808303816000875af115801561321f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132439190614a4a565b5050505050565b60006119367fdfe02260445526f7b137cb9caf995dcdead56fff547ac8de4b3e3305217231488361296f565b6000806132838360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b600054610100900460ff1661331e5760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610576565b6111573361286f565b6000613332826137c8565b6129435760405162461bcd60e51b815260206004820152601960248201527f4e6f7420616e206174746573746174696f6e207265706f7274000000000000006044820152606401610576565b60006119367fbf180edbd986dd1b6d6de1afe33dbc4c91ee49032bd1af9001bf3a96c95e6fb08361296f565b60006133b58261381a565b6129435760405162461bcd60e51b815260206004820152600d60248201527f4e6f7420612072656365697074000000000000000000000000000000000000006044820152606401610576565b60006119367f293501048791dbdbd4a6187fddcc1046f21c1173ad2502f4b7275f89714771d48361296f565b6000600183831c16810361344c576134458585613756565b90506106cb565b6134458486613756565b60ca546040517f6116921800000000000000000000000000000000000000000000000000000000815263ffffffff8085166004830152831660248201526001600160a01b0390911690636116921890604401612d1f565b6000806134ba838561479d565b90506040518111156134ca575060005b80600003613504576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b83176106cb565b600060016fffffffffffffffffffffffffffffffff8316101561353557506000919050565b600061354083613859565b60ff16111561355157506000919050565b61193661355d83612aa2565b613669565b60008161356e84613276565b604080516020810193909352820152606001612512565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c016124dd565b60008060006135cf8585613867565b915091506135dc816138a9565b509392505050565b6000604e6fffffffffffffffffffffffffffffffff83165b1492915050565b60006fffffffffffffffffffffffffffffffff831680831115613652576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6106cb836136608660801c90565b018483036134ad565b6000613677600c603261479d565b6fffffffffffffffffffffffffffffffff83166135fc565b60006111c183828461301a565b6000806136aa858585612da4565b602084900360031b1c9150509392505050565b6000828260405160200161251292919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60006fffffffffffffffffffffffffffffffff821681613722600c603261479d565b61372c90836149e6565b90508161373b600c603261479d565b61374590836147b0565b1480156106cb57506106cb81613a95565b600082158015613764575081155b1561377157506000611936565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050611936565b60006fffffffffffffffffffffffffffffffff82166137c18360801c90565b0192915050565b600060016fffffffffffffffffffffffffffffffff831610156137ed57506000919050565b60006137f883613859565b60ff16111561380957506000919050565b61193661381583612aa2565b6135e4565b60006138286020608561479d565b6fffffffffffffffffffffffffffffffff83161461384857506000919050565b61193661385483613aba565b613ac8565b60006119368282600161369c565b600080825160410361389d5760208301516040840151606085015160001a61389187828585613ae4565b945094505050506119c8565b506000905060026119c8565b60008160048111156138bd576138bd614044565b036138c55750565b60018160048111156138d9576138d9614044565b036139265760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610576565b600281600481111561393a5761393a614044565b036139875760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610576565b600381600481111561399b5761399b614044565b03613a0e5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610576565b6004816004811115613a2257613a22614044565b03610d445760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610576565b600081158015906119365750613aad6001600661478a565b6001901b82111592915050565b60006119368282608561301a565b600060856fffffffffffffffffffffffffffffffff83166135fc565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613b1b5750600090506003613bc8565b8460ff16601b14158015613b3357508460ff16601c14155b15613b445750600090506004613bc8565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613b98573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613bc157600060019250925050613bc8565b9150600090505b94509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715613c2357613c23613bd1565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715613c5257613c52613bd1565b604052919050565b600082601f830112613c6b57600080fd5b813567ffffffffffffffff811115613c8557613c85613bd1565b613c986020601f19601f84011601613c29565b818152846020838601011115613cad57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff821115613ce457613ce4613bd1565b5060051b60200190565b600082601f830112613cff57600080fd5b81356020613d14613d0f83613cca565b613c29565b82815260059290921b84018101918181019086841115613d3357600080fd5b8286015b84811015613d4e5780358352918301918301613d37565b509695505050505050565b60008060008060008060c08789031215613d7257600080fd5b86359550602087013567ffffffffffffffff80821115613d9157600080fd5b613d9d8a838b01613c5a565b96506040890135915080821115613db357600080fd5b613dbf8a838b01613c5a565b95506060890135915080821115613dd557600080fd5b613de18a838b01613cee565b94506080890135915080821115613df757600080fd5b613e038a838b01613c5a565b935060a0890135915080821115613e1957600080fd5b50613e2689828a01613c5a565b9150509295509295509295565b63ffffffff81168114610d4457600080fd5b80356001600160a01b0381168114610b3957600080fd5b60008060008060808587031215613e7257600080fd5b8435613e7d81613e33565b935060208501359250613e9260408601613e45565b9396929550929360600135925050565b60008060008060808587031215613eb857600080fd5b84359350602085013567ffffffffffffffff80821115613ed757600080fd5b613ee388838901613c5a565b94506040870135915080821115613ef957600080fd5b613f0588838901613c5a565b93506060870135915080821115613f1b57600080fd5b50613f2887828801613c5a565b91505092959194509250565b600080600060608486031215613f4957600080fd5b83359250602084013567ffffffffffffffff80821115613f6857600080fd5b613f7487838801613c5a565b93506040860135915080821115613f8a57600080fd5b50613f9786828701613c5a565b9150509250925092565b60008060008060008060c08789031215613fba57600080fd5b86359550602087013567ffffffffffffffff80821115613fd957600080fd5b613fe58a838b01613c5a565b96506040890135915080821115613ffb57600080fd5b6140078a838b01613c5a565b9550606089013591508082111561401d57600080fd5b613de18a838b01613c5a565b60006020828403121561403b57600080fd5b6111c182613e45565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6006811061408357614083614044565b9052565b614092828251614073565b60208181015163ffffffff9081169184019190915260409182015116910152565b606081016119368284614087565b6000602082840312156140d357600080fd5b5035919050565b6001600160a01b0383168152608081016111c16020830184614087565b80516003811061410957614109614044565b825260208181015163ffffffff16908301526040908101516001600160a01b0316910152565b6060810161193682846140f7565b6000806040838503121561415057600080fd5b61415983613e45565b915061416760208401613e45565b90509250929050565b6000815180845260005b818110156141965760208185018101518683018201520161417a565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006111c16020830184614170565b600080600080608085870312156141df57600080fd5b843567ffffffffffffffff808211156141f757600080fd5b61420388838901613c5a565b955060209150818701358181111561421a57600080fd5b61422689828a01613c5a565b9550506040870135935060608701358181111561424257600080fd5b87019050601f8101881361425557600080fd5b8035614263613d0f82613cca565b81815260059190911b8201830190838101908a83111561428257600080fd5b928401925b828410156142a057833582529284019290840190614287565b979a9699509497505050505050565b600080600080600060a086880312156142c757600080fd5b85359450602086013567ffffffffffffffff808211156142e657600080fd5b6142f289838a01613c5a565b9550604088013591508082111561430857600080fd5b61431489838a01613c5a565b9450606088013591508082111561432a57600080fd5b61433689838a01613c5a565b9350608088013591508082111561434c57600080fd5b5061435988828901613c5a565b9150509295509295909350565b60008060006060848603121561437b57600080fd5b833567ffffffffffffffff8082111561439357600080fd5b61439f87838801613c5a565b94506020860135915080821115613f6857600080fd5b600080600080600060a086880312156143cd57600080fd5b85359450602086013567ffffffffffffffff808211156143ec57600080fd5b6143f889838a01613c5a565b9550604088013591508082111561440e57600080fd5b61431489838a01613cee565b6000806040838503121561442d57600080fd5b823567ffffffffffffffff8082111561444557600080fd5b61445186838701613c5a565b9350602085013591508082111561446757600080fd5b5061447485828601613c5a565b9150509250929050565b600080600083850360a081121561449457600080fd5b61449d85613e45565b93506060601f19820112156144b157600080fd5b506144ba613c00565b6020850135600681106144cc57600080fd5b815260408501356144dc81613e33565b602082015260608501356144ef81613e33565b60408201529150608084013567ffffffffffffffff81111561451057600080fd5b613f9786828701613cee565b60006020828403121561452e57600080fd5b815180151581146111c157600080fd5b8481526080602082015260006145576080830186614170565b82810360408401526145698186614170565b90508281036060840152612bb18185614170565b8381526060602082015260006145966060830185614170565b82810360408401526145a88185614170565b9695505050505050565b63ffffffff8616815260006020868184015260a060408401526145d860a0840187614170565b60608401869052838103608085015284518082528286019183019060005b818110156146245783516fffffffffffffffffffffffffffffffff16835292840192918401916001016145f6565b50909a9950505050505050505050565b63ffffffff851681526001600160a01b03841660208201526080604082015260006146626080830185614170565b8281036060840152612bb18185614170565b6040815260006146876040830185614170565b8281036020840152611fa98185614170565b602081016119368284614073565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c908216806146ea57607f821691505b60208210810361257b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361478357614783614723565b5060010190565b8181038181111561193657611936614723565b8082018082111561193657611936614723565b808202811582820484141761193657611936614723565b815160009082906020808601845b838110156148035781516fffffffffffffffffffffffffffffffff16855293820193908201906001016147d5565b50929695505050505050565b601f821115610cdc57600081815260208120601f850160051c810160208610156148365750805b601f850160051c820191505b818110156116af57828155600101614842565b815167ffffffffffffffff81111561486f5761486f613bd1565b6148838161487d84546146d6565b8461480f565b602080601f8311600181146148d657600084156148a05750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556116af565b600085815260208120601f198616915b82811015614905578886015182559484019460019091019084016148e6565b508582101561494157878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b60006006851061496357614963614044565b5060f89390931b835260e09190911b7fffffffff0000000000000000000000000000000000000000000000000000000016600183015260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600582015260190190565b6001600160a01b0383168152608081016111c160208301846140f7565b600082614a1c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600063ffffffff808616835280851660208401525060606040830152611fa96060830184614170565b60008060408385031215614a5d57600080fd5b8251614a6881613e33565b602093909301519294929350505056fea2646970667358221220ae149cfe94001e8fdaad0800853031a0d32d2397334d1431bc718e864cbaa68464736f6c63430008110033",
}

// LightManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use LightManagerMetaData.ABI instead.
var LightManagerABI = LightManagerMetaData.ABI

// Deprecated: Use LightManagerMetaData.Sigs instead.
// LightManagerFuncSigs maps the 4-byte function signature to its string representation.
var LightManagerFuncSigs = LightManagerMetaData.Sigs

// LightManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LightManagerMetaData.Bin instead.
var LightManagerBin = LightManagerMetaData.Bin

// DeployLightManager deploys a new Ethereum contract, binding an instance of LightManager to it.
func DeployLightManager(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32) (common.Address, *types.Transaction, *LightManager, error) {
	parsed, err := LightManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LightManagerBin), backend, domain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LightManager{LightManagerCaller: LightManagerCaller{contract: contract}, LightManagerTransactor: LightManagerTransactor{contract: contract}, LightManagerFilterer: LightManagerFilterer{contract: contract}}, nil
}

// LightManager is an auto generated Go binding around an Ethereum contract.
type LightManager struct {
	LightManagerCaller     // Read-only binding to the contract
	LightManagerTransactor // Write-only binding to the contract
	LightManagerFilterer   // Log filterer for contract events
}

// LightManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightManagerSession struct {
	Contract     *LightManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LightManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightManagerCallerSession struct {
	Contract *LightManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LightManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightManagerTransactorSession struct {
	Contract     *LightManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LightManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightManagerRaw struct {
	Contract *LightManager // Generic contract binding to access the raw methods on
}

// LightManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightManagerCallerRaw struct {
	Contract *LightManagerCaller // Generic read-only contract binding to access the raw methods on
}

// LightManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightManagerTransactorRaw struct {
	Contract *LightManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightManager creates a new instance of LightManager, bound to a specific deployed contract.
func NewLightManager(address common.Address, backend bind.ContractBackend) (*LightManager, error) {
	contract, err := bindLightManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightManager{LightManagerCaller: LightManagerCaller{contract: contract}, LightManagerTransactor: LightManagerTransactor{contract: contract}, LightManagerFilterer: LightManagerFilterer{contract: contract}}, nil
}

// NewLightManagerCaller creates a new read-only instance of LightManager, bound to a specific deployed contract.
func NewLightManagerCaller(address common.Address, caller bind.ContractCaller) (*LightManagerCaller, error) {
	contract, err := bindLightManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightManagerCaller{contract: contract}, nil
}

// NewLightManagerTransactor creates a new write-only instance of LightManager, bound to a specific deployed contract.
func NewLightManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*LightManagerTransactor, error) {
	contract, err := bindLightManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightManagerTransactor{contract: contract}, nil
}

// NewLightManagerFilterer creates a new log filterer instance of LightManager, bound to a specific deployed contract.
func NewLightManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*LightManagerFilterer, error) {
	contract, err := bindLightManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightManagerFilterer{contract: contract}, nil
}

// bindLightManager binds a generic wrapper to an already deployed contract.
func bindLightManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LightManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightManager *LightManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightManager.Contract.LightManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightManager *LightManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightManager.Contract.LightManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightManager *LightManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightManager.Contract.LightManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightManager *LightManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightManager *LightManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightManager *LightManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightManager.Contract.contract.Transact(opts, method, params...)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_LightManager *LightManagerCaller) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "agentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_LightManager *LightManagerSession) AgentRoot() ([32]byte, error) {
	return _LightManager.Contract.AgentRoot(&_LightManager.CallOpts)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_LightManager *LightManagerCallerSession) AgentRoot() ([32]byte, error) {
	return _LightManager.Contract.AgentRoot(&_LightManager.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_LightManager *LightManagerCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_LightManager *LightManagerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _LightManager.Contract.AgentStatus(&_LightManager.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_LightManager *LightManagerCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _LightManager.Contract.AgentStatus(&_LightManager.CallOpts, agent)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightManager *LightManagerCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightManager *LightManagerSession) Destination() (common.Address, error) {
	return _LightManager.Contract.Destination(&_LightManager.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightManager *LightManagerCallerSession) Destination() (common.Address, error) {
	return _LightManager.Contract.Destination(&_LightManager.CallOpts)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_LightManager *LightManagerCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (Dispute, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(Dispute), err
	}

	out0 := *abi.ConvertType(out[0], new(Dispute)).(*Dispute)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_LightManager *LightManagerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _LightManager.Contract.DisputeStatus(&_LightManager.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_LightManager *LightManagerCallerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _LightManager.Contract.DisputeStatus(&_LightManager.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_LightManager *LightManagerCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "getAgent", index)

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
func (_LightManager *LightManagerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _LightManager.Contract.GetAgent(&_LightManager.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_LightManager *LightManagerCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _LightManager.Contract.GetAgent(&_LightManager.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightManager *LightManagerCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightManager *LightManagerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightManager.Contract.GetStoredSignature(&_LightManager.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightManager *LightManagerCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightManager.Contract.GetStoredSignature(&_LightManager.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightManager *LightManagerCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightManager *LightManagerSession) LocalDomain() (uint32, error) {
	return _LightManager.Contract.LocalDomain(&_LightManager.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightManager *LightManagerCallerSession) LocalDomain() (uint32, error) {
	return _LightManager.Contract.LocalDomain(&_LightManager.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightManager *LightManagerCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightManager *LightManagerSession) Origin() (common.Address, error) {
	return _LightManager.Contract.Origin(&_LightManager.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightManager *LightManagerCallerSession) Origin() (common.Address, error) {
	return _LightManager.Contract.Origin(&_LightManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightManager *LightManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightManager *LightManagerSession) Owner() (common.Address, error) {
	return _LightManager.Contract.Owner(&_LightManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightManager *LightManagerCallerSession) Owner() (common.Address, error) {
	return _LightManager.Contract.Owner(&_LightManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightManager *LightManagerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightManager.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightManager *LightManagerSession) Version() (string, error) {
	return _LightManager.Contract.Version(&_LightManager.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightManager *LightManagerCallerSession) Version() (string, error) {
	return _LightManager.Contract.Version(&_LightManager.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_LightManager *LightManagerTransactor) Initialize(opts *bind.TransactOpts, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "initialize", origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_LightManager *LightManagerSession) Initialize(origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightManager.Contract.Initialize(&_LightManager.TransactOpts, origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_LightManager *LightManagerTransactorSession) Initialize(origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightManager.Contract.Initialize(&_LightManager.TransactOpts, origin_, destination_)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_LightManager *LightManagerTransactor) RemoteWithdrawTips(opts *bind.TransactOpts, msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "remoteWithdrawTips", msgOrigin, proofMaturity, recipient, amount)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_LightManager *LightManagerSession) RemoteWithdrawTips(msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LightManager.Contract.RemoteWithdrawTips(&_LightManager.TransactOpts, msgOrigin, proofMaturity, recipient, amount)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_LightManager *LightManagerTransactorSession) RemoteWithdrawTips(msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LightManager.Contract.RemoteWithdrawTips(&_LightManager.TransactOpts, msgOrigin, proofMaturity, recipient, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightManager *LightManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightManager *LightManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightManager.Contract.RenounceOwnership(&_LightManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightManager *LightManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightManager.Contract.RenounceOwnership(&_LightManager.TransactOpts)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot_) returns()
func (_LightManager *LightManagerTransactor) SetAgentRoot(opts *bind.TransactOpts, agentRoot_ [32]byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "setAgentRoot", agentRoot_)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot_) returns()
func (_LightManager *LightManagerSession) SetAgentRoot(agentRoot_ [32]byte) (*types.Transaction, error) {
	return _LightManager.Contract.SetAgentRoot(&_LightManager.TransactOpts, agentRoot_)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot_) returns()
func (_LightManager *LightManagerTransactorSession) SetAgentRoot(agentRoot_ [32]byte) (*types.Transaction, error) {
	return _LightManager.Contract.SetAgentRoot(&_LightManager.TransactOpts, agentRoot_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "submitAttestation", attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightManager *LightManagerSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitAttestation(&_LightManager.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitAttestation(&_LightManager.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactor) SubmitAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "submitAttestationReport", arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitAttestationReport(&_LightManager.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactorSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitAttestationReport(&_LightManager.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitStateReportWithAttestation(&_LightManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitStateReportWithAttestation(&_LightManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitStateReportWithSnapshot(&_LightManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitStateReportWithSnapshot(&_LightManager.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitStateReportWithSnapshotProof(&_LightManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManager *LightManagerTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.SubmitStateReportWithSnapshotProof(&_LightManager.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightManager *LightManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightManager *LightManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightManager.Contract.TransferOwnership(&_LightManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightManager *LightManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightManager.Contract.TransferOwnership(&_LightManager.TransactOpts, newOwner)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_LightManager *LightManagerTransactor) UpdateAgentStatus(opts *bind.TransactOpts, agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "updateAgentStatus", agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_LightManager *LightManagerSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _LightManager.Contract.UpdateAgentStatus(&_LightManager.TransactOpts, agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_LightManager *LightManagerTransactorSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _LightManager.Contract.UpdateAgentStatus(&_LightManager.TransactOpts, agent, status, proof)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightManager *LightManagerTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightManager *LightManagerSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyReceipt(&_LightManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightManager *LightManagerTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyReceipt(&_LightManager.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_LightManager *LightManagerTransactor) VerifyStateReport(opts *bind.TransactOpts, srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "verifyStateReport", srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_LightManager *LightManagerSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateReport(&_LightManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_LightManager *LightManagerTransactorSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateReport(&_LightManager.TransactOpts, srPayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManager *LightManagerTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManager *LightManagerSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateWithAttestation(&_LightManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManager *LightManagerTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateWithAttestation(&_LightManager.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightManager *LightManagerTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightManager *LightManagerSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateWithSnapshot(&_LightManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightManager *LightManagerTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateWithSnapshot(&_LightManager.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManager *LightManagerTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManager *LightManagerSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateWithSnapshotProof(&_LightManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManager *LightManagerTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManager.Contract.VerifyStateWithSnapshotProof(&_LightManager.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// LightManagerAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the LightManager contract.
type LightManagerAttestationAcceptedIterator struct {
	Event *LightManagerAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerAttestationAccepted represents a AttestationAccepted event raised by the LightManager contract.
type LightManagerAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightManager *LightManagerFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*LightManagerAttestationAcceptedIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &LightManagerAttestationAcceptedIterator{contract: _LightManager.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightManager *LightManagerFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *LightManagerAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerAttestationAccepted)
				if err := _LightManager.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseAttestationAccepted(log types.Log) (*LightManagerAttestationAccepted, error) {
	event := new(LightManagerAttestationAccepted)
	if err := _LightManager.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerDisputeUpdatedIterator is returned from FilterDisputeUpdated and is used to iterate over the raw logs and unpacked data for DisputeUpdated events raised by the LightManager contract.
type LightManagerDisputeUpdatedIterator struct {
	Event *LightManagerDisputeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerDisputeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerDisputeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerDisputeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerDisputeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerDisputeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerDisputeUpdated represents a DisputeUpdated event raised by the LightManager contract.
type LightManagerDisputeUpdated struct {
	Agent   common.Address
	Dispute Dispute
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeUpdated is a free log retrieval operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_LightManager *LightManagerFilterer) FilterDisputeUpdated(opts *bind.FilterOpts) (*LightManagerDisputeUpdatedIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return &LightManagerDisputeUpdatedIterator{contract: _LightManager.contract, event: "DisputeUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeUpdated is a free log subscription operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_LightManager *LightManagerFilterer) WatchDisputeUpdated(opts *bind.WatchOpts, sink chan<- *LightManagerDisputeUpdated) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerDisputeUpdated)
				if err := _LightManager.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseDisputeUpdated(log types.Log) (*LightManagerDisputeUpdated, error) {
	event := new(LightManagerDisputeUpdated)
	if err := _LightManager.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightManager contract.
type LightManagerInitializedIterator struct {
	Event *LightManagerInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerInitialized represents a Initialized event raised by the LightManager contract.
type LightManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightManager *LightManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightManagerInitializedIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightManagerInitializedIterator{contract: _LightManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightManager *LightManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerInitialized)
				if err := _LightManager.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseInitialized(log types.Log) (*LightManagerInitialized, error) {
	event := new(LightManagerInitialized)
	if err := _LightManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the LightManager contract.
type LightManagerInvalidReceiptIterator struct {
	Event *LightManagerInvalidReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerInvalidReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerInvalidReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerInvalidReceipt represents a InvalidReceipt event raised by the LightManager contract.
type LightManagerInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightManager *LightManagerFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*LightManagerInvalidReceiptIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &LightManagerInvalidReceiptIterator{contract: _LightManager.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightManager *LightManagerFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *LightManagerInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerInvalidReceipt)
				if err := _LightManager.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseInvalidReceipt(log types.Log) (*LightManagerInvalidReceipt, error) {
	event := new(LightManagerInvalidReceipt)
	if err := _LightManager.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the LightManager contract.
type LightManagerInvalidStateReportIterator struct {
	Event *LightManagerInvalidStateReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerInvalidStateReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerInvalidStateReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerInvalidStateReport represents a InvalidStateReport event raised by the LightManager contract.
type LightManagerInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightManager *LightManagerFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*LightManagerInvalidStateReportIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &LightManagerInvalidStateReportIterator{contract: _LightManager.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightManager *LightManagerFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *LightManagerInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerInvalidStateReport)
				if err := _LightManager.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseInvalidStateReport(log types.Log) (*LightManagerInvalidStateReport, error) {
	event := new(LightManagerInvalidStateReport)
	if err := _LightManager.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the LightManager contract.
type LightManagerInvalidStateWithAttestationIterator struct {
	Event *LightManagerInvalidStateWithAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerInvalidStateWithAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerInvalidStateWithAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the LightManager contract.
type LightManagerInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightManager *LightManagerFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*LightManagerInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &LightManagerInvalidStateWithAttestationIterator{contract: _LightManager.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightManager *LightManagerFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *LightManagerInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerInvalidStateWithAttestation)
				if err := _LightManager.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseInvalidStateWithAttestation(log types.Log) (*LightManagerInvalidStateWithAttestation, error) {
	event := new(LightManagerInvalidStateWithAttestation)
	if err := _LightManager.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the LightManager contract.
type LightManagerInvalidStateWithSnapshotIterator struct {
	Event *LightManagerInvalidStateWithSnapshot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerInvalidStateWithSnapshot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerInvalidStateWithSnapshot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the LightManager contract.
type LightManagerInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightManager *LightManagerFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*LightManagerInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &LightManagerInvalidStateWithSnapshotIterator{contract: _LightManager.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightManager *LightManagerFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *LightManagerInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerInvalidStateWithSnapshot)
				if err := _LightManager.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*LightManagerInvalidStateWithSnapshot, error) {
	event := new(LightManagerInvalidStateWithSnapshot)
	if err := _LightManager.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LightManager contract.
type LightManagerOwnershipTransferredIterator struct {
	Event *LightManagerOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerOwnershipTransferred represents a OwnershipTransferred event raised by the LightManager contract.
type LightManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightManager *LightManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LightManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LightManagerOwnershipTransferredIterator{contract: _LightManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightManager *LightManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LightManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerOwnershipTransferred)
				if err := _LightManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseOwnershipTransferred(log types.Log) (*LightManagerOwnershipTransferred, error) {
	event := new(LightManagerOwnershipTransferred)
	if err := _LightManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the LightManager contract.
type LightManagerRootUpdatedIterator struct {
	Event *LightManagerRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerRootUpdated represents a RootUpdated event raised by the LightManager contract.
type LightManagerRootUpdated struct {
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_LightManager *LightManagerFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*LightManagerRootUpdatedIterator, error) {

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &LightManagerRootUpdatedIterator{contract: _LightManager.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_LightManager *LightManagerFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *LightManagerRootUpdated) (event.Subscription, error) {

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerRootUpdated)
				if err := _LightManager.contract.UnpackLog(event, "RootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseRootUpdated(log types.Log) (*LightManagerRootUpdated, error) {
	event := new(LightManagerRootUpdated)
	if err := _LightManager.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the LightManager contract.
type LightManagerStatusUpdatedIterator struct {
	Event *LightManagerStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerStatusUpdated represents a StatusUpdated event raised by the LightManager contract.
type LightManagerStatusUpdated struct {
	Flag   uint8
	Domain uint32
	Agent  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_LightManager *LightManagerFilterer) FilterStatusUpdated(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*LightManagerStatusUpdatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _LightManager.contract.FilterLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &LightManagerStatusUpdatedIterator{contract: _LightManager.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_LightManager *LightManagerFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *LightManagerStatusUpdated, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _LightManager.contract.WatchLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerStatusUpdated)
				if err := _LightManager.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManager *LightManagerFilterer) ParseStatusUpdated(log types.Log) (*LightManagerStatusUpdated, error) {
	event := new(LightManagerStatusUpdated)
	if err := _LightManager.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessMetaData contains all meta data concerning the LightManagerHarness contract.
var LightManagerHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"AttestationAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structDispute\",\"name\":\"dispute\",\"type\":\"tuple\"}],\"name\":\"DisputeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidReceipt\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithAttestation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"InvalidStateWithSnapshot\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"newRoot\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"StatusUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"agentStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"}],\"name\":\"disputeStatus\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDisputeFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"rivalIndex\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"fraudProver\",\"type\":\"address\"}],\"internalType\":\"structDispute\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAgent\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getStoredSignature\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"origin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"openDisputeExposed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"remoteMockFunc\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"msgOrigin\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"proofMaturity\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"remoteWithdrawTips\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"magicValue\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"sensitiveMockFunc\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"data\",\"type\":\"bytes32\"}],\"name\":\"sensitiveMockFuncOver32Bytes\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sensitiveMockFuncVoid\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"}],\"name\":\"setAgentRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prover\",\"type\":\"address\"}],\"name\":\"slashAgentExposed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"snapGas_\",\"type\":\"uint256[]\"}],\"name\":\"submitAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"arPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"arSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitAttestationReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"submitStateReportWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAccepted\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"agent\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumAgentFlag\",\"name\":\"flag\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"index\",\"type\":\"uint32\"}],\"internalType\":\"structAgentStatus\",\"name\":\"status\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"updateAgentStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"rcptPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rcptSignature\",\"type\":\"bytes\"}],\"name\":\"verifyReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReceipt\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"srPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"srSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateReport\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidReport\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"snapPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"snapSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"statePayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes32[]\",\"name\":\"snapProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes\",\"name\":\"attPayload\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attSignature\",\"type\":\"bytes\"}],\"name\":\"verifyStateWithSnapshotProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isValidState\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"versionString\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"36cba43c": "agentRoot()",
		"28f3fac9": "agentStatus(address)",
		"b269681d": "destination()",
		"3463d1b1": "disputeStatus(address)",
		"2de5aaf7": "getAgent(uint256)",
		"ddeffa66": "getStoredSignature(uint256)",
		"485cc955": "initialize(address,address)",
		"8d3638f4": "localDomain()",
		"c166ca67": "openDisputeExposed(address,address)",
		"938b5f32": "origin()",
		"8da5cb5b": "owner()",
		"a149352c": "remoteMockFunc(uint32,uint256,bytes32)",
		"1fa07138": "remoteWithdrawTips(uint32,uint256,address,uint256)",
		"715018a6": "renounceOwnership()",
		"127a2c9d": "sensitiveMockFunc(address,uint8,bytes32)",
		"0e6bfcd5": "sensitiveMockFuncOver32Bytes(uint16,bytes4,bytes32)",
		"c9f1a03f": "sensitiveMockFuncVoid(uint16,bytes4,bytes32)",
		"58668176": "setAgentRoot(bytes32)",
		"69978b0d": "slashAgentExposed(uint32,address,address)",
		"6c38f723": "submitAttestation(bytes,bytes,bytes32,uint256[])",
		"77ec5c10": "submitAttestationReport(bytes,bytes,bytes)",
		"235d51b1": "submitStateReportWithAttestation(uint256,bytes,bytes,bytes,bytes,bytes)",
		"708cdc82": "submitStateReportWithSnapshot(uint256,bytes,bytes,bytes,bytes)",
		"0db27e77": "submitStateReportWithSnapshotProof(uint256,bytes,bytes,bytes32[],bytes,bytes)",
		"f2fde38b": "transferOwnership(address)",
		"cbd05965": "updateAgentStatus(address,(uint8,uint32,uint32),bytes32[])",
		"c25aa585": "verifyReceipt(bytes,bytes)",
		"dfe39675": "verifyStateReport(bytes,bytes)",
		"200f6b66": "verifyStateWithAttestation(uint256,bytes,bytes,bytes)",
		"213a6ddb": "verifyStateWithSnapshot(uint256,bytes,bytes)",
		"7be8e738": "verifyStateWithSnapshotProof(uint256,bytes,bytes32[],bytes,bytes)",
		"54fd4d50": "version()",
	},
	Bin: "0x60e06040523480156200001157600080fd5b5060405162005077380380620050778339810160408190526200003491620000df565b60408051808201909152600580825264302e302e3360d81b60208301526080528190818162000067565b60405180910390fd5b62000072816200010e565b60a0525063ffffffff90811660c0528216600919019050620000d75760405162461bcd60e51b815260206004820152601d60248201527f43616e2774206265206465706c6f796564206f6e2053796e436861696e00000060448201526064016200005e565b505062000136565b600060208284031215620000f257600080fd5b815163ffffffff811681146200010757600080fd5b9392505050565b8051602080830151919081101562000130576000198160200360031b1b821691505b50919050565b60805160a05160c051614f0a6200016d600039600081816104650152610fa5015260006103b4015260006103910152614f0a6000f3fe608060405234801561001057600080fd5b50600436106101f05760003560e01c8063708cdc821161010f578063b269681d116100a2578063cbd0596511610071578063cbd0596514610533578063ddeffa6614610546578063dfe3967514610559578063f2fde38b1461056c57600080fd5b8063b269681d146104e7578063c166ca67146104fa578063c25aa5851461050d578063c9f1a03f1461052057600080fd5b80638d3638f4116100de5780638d3638f4146104605780638da5cb5b1461049c578063938b5f32146104c1578063a149352c146104d457600080fd5b8063708cdc821461041f578063715018a61461043257806377ec5c101461043a5780637be8e7381461044d57600080fd5b80632de5aaf71161018757806354fd4d501161015657806354fd4d501461038557806358668176146103e657806369978b0d146103f95780636c38f7231461040c57600080fd5b80632de5aaf7146103265780633463d1b11461034757806336cba43c14610367578063485cc9551461037057600080fd5b8063200f6b66116101c3578063200f6b66146102cd578063213a6ddb146102e0578063235d51b1146102f357806328f3fac91461030657600080fd5b80630db27e77146101f55780630e6bfcd51461021d578063127a2c9d146102685780631fa0713814610289575b600080fd5b61020861020336600461409a565b61057f565b60405190151581526020015b60405180910390f35b61023061022b366004614174565b610609565b604080517fffffffff000000000000000000000000000000000000000000000000000000009093168352602083019190915201610214565b61027b6102763660046141f8565b610699565b604051908152602001610214565b61029c61029736600461423e565b61074b565b6040517fffffffff000000000000000000000000000000000000000000000000000000009091168152602001610214565b6102086102db366004614284565b6108fc565b6102086102ee366004614316565b610a95565b610208610301366004614383565b610b96565b61031961031436600461440b565b610ce1565b6040516102149190614495565b6103396103343660046144a3565b610d67565b6040516102149291906144bc565b61035a61035536600461440b565b610db4565b6040516102149190614511565b61027b60fb5481565b61038361037e36600461451f565b610e4f565b005b604080518082019091527f000000000000000000000000000000000000000000000000000000000000000081527f000000000000000000000000000000000000000000000000000000000000000060208201525b6040516102149190614598565b6103836103f43660046144a3565b610f0a565b6103836104073660046145ab565b610f70565b61020861041a3660046145f0565b610f7b565b61020861042d3660046146d6565b611200565b610383611331565b61020861044836600461478d565b61138d565b61020861045b3660046147dc565b6113fa565b6104877f000000000000000000000000000000000000000000000000000000000000000081565b60405163ffffffff9091168152602001610214565b6033546001600160a01b03165b6040516001600160a01b039091168152602001610214565b60c9546104a9906001600160a01b031681565b61029c6104e2366004614841565b611525565b60ca546104a9906001600160a01b031681565b61038361050836600461451f565b6115ab565b61020861051b366004614876565b6115d3565b61038361052e366004614174565b6116dd565b6103836105413660046148da565b611737565b6103d96105543660046144a3565b6119f1565b610208610567366004614876565b611aa0565b61038361057a36600461440b565b611b7e565b60008061058b87611c5d565b905060008061059a8389611c76565b915091506105a782611d09565b60006105b287611dc2565b90506000806105c18389611dd5565b915091506105ce82611e5d565b6105e2838e6105dc89611ecb565b8d611ede565b6105f6848660400151838560400151611fcc565b5060019c9b505050505050505050505050565b60ca5460009081906001600160a01b0316331461066d5760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e000000000000000000000000000000000000000060448201526064015b60405180910390fd5b507f0e6bfcd5000000000000000000000000000000000000000000000000000000009491935090915050565b60ca546000906001600160a01b031633146106f65760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610664565b8161071f577f474d00000000000000000000000000000000000000000000000000000000000091505b507f127a2c9d0000000000000000000000000000000000000000000000000000000081185b9392505050565b60ca546000906001600160a01b031633146107a85760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610664565b63ffffffff8516600a146107fe5760405162461bcd60e51b815260206004820152600e60248201527f2173796e61707365446f6d61696e0000000000000000000000000000000000006044820152606401610664565b620151808410156108515760405162461bcd60e51b815260206004820152601160248201527f216f7074696d6973746963506572696f640000000000000000000000000000006044820152606401610664565b60c9546040517f4e04e7a70000000000000000000000000000000000000000000000000000000081526001600160a01b0385811660048301526024820185905290911690634e04e7a790604401600060405180830381600087803b1580156108b857600080fd5b505af11580156108cc573d6000803e3d6000fd5b507f1fa071380000000000000000000000000000000000000000000000000000000093505050505b949350505050565b60008061090884611dc2565b90506000806109178386611dd5565b9150915061092482611e5d565b600061092f8861214a565b905061093a8461215d565b6109438261216e565b146109905760405162461bcd60e51b815260206004820181905260248201527f4174746573746174696f6e206e6f74206d61746368657320736e617073686f746044820152606401610664565b60006109aa6109a56109a2848d612247565b90565b6122e8565b60c9546040517fa9dcf22d0000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063a9dcf22d906109f4908490600401614598565b602060405180830381865afa158015610a11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a359190614978565b955085610a88577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a828a8a604051610a71949392919061499a565b60405180910390a1610a8884602001518433612327565b5050505050949350505050565b600080610aa18461214a565b9050600080610ab08386612431565b91509150610abd82611e5d565b60c9546001600160a01b031663a9dcf22d610ade6109a56109a2878c612247565b6040518263ffffffff1660e01b8152600401610afa9190614598565b602060405180830381865afa158015610b17573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b3b9190614978565b935083610b8c577f8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1878787604051610b75939291906149d9565b60405180910390a1610b8c82602001518233612327565b5050509392505050565b600080610ba287611c5d565b9050600080610bb18389611c76565b915091506000610bc08861214a565b9050610bde610bce85611ecb565b610bd8838e612247565b90612465565b610c2a5760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610664565b610c3383611d09565b6000610c3e88611dc2565b9050600080610c4d838a611dd5565b91509150610c5a82611e5d565b610c638361215d565b610c6c8561216e565b14610cb95760405162461bcd60e51b815260206004820181905260248201527f4174746573746174696f6e206e6f74206d61746368657320736e617073686f746044820152606401610664565b610ccd858760400151838560400151611fcc565b5060019d9c50505050505050505050505050565b6040805160608101825260008082526020820181905291810191909152610d0782612486565b905060026001600160a01b038316600090815260cb602052604090205460ff166002811115610d3857610d38614426565b148015610d585750600581516005811115610d5557610d55614426565b14155b15610d6257600481525b919050565b60408051606081018252600080825260208201819052918101829052600083815260fd60205260409020546001600160a01b031691508115610daf57610dac82610ce1565b90505b915091565b60408051606081018252600080825260208201819052918101919091526001600160a01b038216600090815260cb6020526040908190208151606081019092528054829060ff166002811115610e0c57610e0c614426565b6002811115610e1d57610e1d614426565b81529054610100810463ffffffff1660208301526501000000000090046001600160a01b031660409091015292915050565b6000610e5b6001612522565b90508015610e9057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b610e9a8383612674565b610ea2612737565b8015610f0557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b505050565b60ca546001600160a01b03163314610f645760405162461bcd60e51b815260206004820181905260248201527f4f6e6c792044657374696e6174696f6e2073657473206167656e7420726f6f746044820152606401610664565b610f6d816127bc565b50565b610f05838383612327565b600080610f8786611dc2565b9050600080610f968388611dd5565b91509150610fa382611d09565b7f000000000000000000000000000000000000000000000000000000000000000063ffffffff16826020015163ffffffff16146110225760405162461bcd60e51b815260206004820152601360248201527f57726f6e67204e6f7461727920646f6d61696e000000000000000000000000006044820152606401610664565b6001600160a01b038116600090815260cb602052604081205460ff16600281111561104f5761104f614426565b1461109c5760405162461bcd60e51b815260206004820152601460248201527f4e6f7461727920697320696e20646973707574650000000000000000000000006044820152606401610664565b846110af876110aa83612800565b612830565b6110b885612866565b146111055760405162461bcd60e51b815260206004820152600f60248201527f496e76616c696420736e617047617300000000000000000000000000000000006044820152606401610664565b600061111089612874565b60ca5460408087015190517f39fe27360000000000000000000000000000000000000000000000000000000081529293506001600160a01b03909116916339fe2736916111679185908f908e908990600401614a0e565b6020604051808303816000875af1158015611186573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111aa9190614978565b95508515610a88577f5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea00658460200151848c8c6040516111eb9493929190614a90565b60405180910390a15050505050949350505050565b60008061120c86611c5d565b905060008061121b8388611c76565b9150915061122882611d09565b60006112338761214a565b90506000806112428389612431565b91509150816020015163ffffffff166000036112a05760405162461bcd60e51b815260206004820152601f60248201527f536e617073686f74207369676e6572206973206e6f742061204e6f74617279006044820152606401610664565b6112a982611e5d565b6112bf6112b587611ecb565b610bd8858f612247565b61130b5760405162461bcd60e51b815260206004820152601260248201527f53746174657320646f6e2774206d6174636800000000000000000000000000006044820152606401610664565b61131f848660400151838560400151611fcc565b5060019b9a5050505050505050505050565b6033546001600160a01b0316331461138b5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610664565b565b600080611399856128b7565b90506000806113a883876128ca565b915091506113b582611d09565b6000806113ca6113c4866128f3565b88611dd5565b915091506113d782611e5d565b6113eb838560400151838560400151611fcc565b50600198975050505050505050565b60008061140684611dc2565b90506000806114158386611dd5565b9150915061142282611e5d565b600061142d89612901565b905061143b848b838b611ede565b60c9546040517fa9dcf22d0000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063a9dcf22d90611484908c90600401614598565b602060405180830381865afa1580156114a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114c59190614978565b945084611518577f541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a4928a8a8989604051611501949392919061499a565b60405180910390a161151883602001518333612327565b5050505095945050505050565b60ca546000906001600160a01b031633146115825760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610664565b507fa149352c000000000000000000000000000000000000000000000000000000009392505050565b6115cf826115b884612486565b60400151836115c685612486565b60400151611fcc565b5050565b6000806115df8461290f565b90506000806115ee8386612922565b915091506115fb82611e5d565b60ca546040517fe2f006f70000000000000000000000000000000000000000000000000000000081526001600160a01b039091169063e2f006f790611644908990600401614598565b602060405180830381865afa158015611661573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116859190614978565b9350836116d4577f4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d86866040516116bd929190614ad0565b60405180910390a16116d482602001518233612327565b50505092915050565b60ca546001600160a01b03163314610f055760405162461bcd60e51b815260206004820152600c60248201527f2164657374696e6174696f6e00000000000000000000000000000000000000006044820152606401610664565b60408083015163ffffffff16600090815260fd60205220546001600160a01b03168015806117765750836001600160a01b0316816001600160a01b0316145b6117c25760405162461bcd60e51b815260206004820152601360248201527f496e76616c6964206167656e7420696e646578000000000000000000000000006044820152606401610664565b60006117d7846000015185602001518761294b565b9050600060fb549050806117f8866040015163ffffffff1684876020612981565b146118455760405162461bcd60e51b815260206004820152600d60248201527f496e76616c69642070726f6f66000000000000000000000000000000000000006044820152606401610664565b6001600160a01b03831661189c5760408581015163ffffffff16600090815260fd6020522080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0388161790555b600081815260fc602090815260408083206001600160a01b038a1684529091529020855181548792919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660018360058111156118fe576118fe614426565b021790555060208281015182546040948501517fffffffffffffffffffffffffffffffffffffffffffffff0000000000000000ff90911661010063ffffffff938416027fffffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffff1617650100000000009183169190910217909255870151875192516001600160a01b038a169391909216917f8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e916119b891614af5565b60405180910390a36005855160058111156119d5576119d5614426565b036119e9576119e986866040015133612a3f565b505050505050565b606060cc8281548110611a0657611a06614b03565b906000526020600020018054611a1b90614b32565b80601f0160208091040260200160405190810160405280929190818152602001828054611a4790614b32565b8015611a945780601f10611a6957610100808354040283529160200191611a94565b820191906000526020600020905b815481529060010190602001808311611a7757829003601f168201915b50505050509050919050565b600080611aac84611c5d565b9050600080611abb8386611c76565b91509150611ac882611e5d565b60c9546001600160a01b031663a9dcf22d611ae86109a56109a287611ecb565b6040518263ffffffff1660e01b8152600401611b049190614598565b602060405180830381865afa158015611b21573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b459190614978565b159350836116d4577f9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d86866040516116bd929190614ad0565b6033546001600160a01b03163314611bd85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610664565b6001600160a01b038116611c545760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610664565b610f6d81612ba5565b6000611c70611c6b83612c0f565b612c22565b92915050565b6040805160608101825260008082526020820181905291810182905290611ca5611c9f85612c7d565b84612cab565b6020820151919350915063ffffffff1615611d025760405162461bcd60e51b815260206004820152601560248201527f5369676e6572206973206e6f74206120477561726400000000000000000000006044820152606401610664565b9250929050565b60015b81516005811115611d1f57611d1f614426565b14816020015163ffffffff16600014611d6d576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611da4565b6040518060400160405280601381526020017f4e6f7420616e20616374697665206775617264000000000000000000000000008152505b906115cf5760405162461bcd60e51b81526004016106649190614598565b6000611c70611dd083612c0f565b612d55565b6040805160608101825260008082526020820181905291810182905290611dfe611c9f85612dac565b6020820151919350915063ffffffff16600003611d025760405162461bcd60e51b815260206004820152601660248201527f5369676e6572206973206e6f742061204e6f74617279000000000000000000006044820152606401610664565b600181516005811115611e7257611e72614426565b1480611e8057506002611d0c565b602082015163ffffffff1615611d6d576040518060400160405280601481526020017f4e6f7420616e20616374697665206e6f74617279000000000000000000000000815250611da4565b6000611c70611ed983612dd8565b612de5565b6000611ee983612e3c565b9150508082600081518110611f0057611f00614b03565b602002602001015114611f555760405162461bcd60e51b815260206004820152601260248201527f496e636f72726563742070726f6f665b305d00000000000000000000000000006044820152606401610664565b6000611f73611f638561215d565b611f6c86612e66565b8588612e75565b905080611f7f8761215d565b146119e95760405162461bcd60e51b815260206004820152601760248201527f496e636f727265637420736e617073686f7420726f6f740000000000000000006044820152606401610664565b6001600160a01b038416600090815260cb602052604081205460ff166002811115611ff957611ff9614426565b146120465760405162461bcd60e51b815260206004820152601860248201527f477561726420616c726561647920696e206469737075746500000000000000006044820152606401610664565b6001600160a01b038216600090815260cb602052604081205460ff16600281111561207357612073614426565b146120c05760405162461bcd60e51b815260206004820152601960248201527f4e6f7461727920616c726561647920696e2064697370757465000000000000006044820152606401610664565b6120fd846040518060600160405280600160028111156120e2576120e2614426565b815263ffffffff851660208201526000604090910152612ef2565b61213a8260405180606001604052806001600281111561211f5761211f614426565b815263ffffffff871660208201526000604090910152612ef2565b6121448382613002565b50505050565b6000611c7061215883612c0f565b613083565b6000611c70816020845b91906130da565b60008061217a836131e4565b905060008167ffffffffffffffff81111561219757612197613f12565b6040519080825280602002602001820160405280156121c0578160200160208202803683370190505b50905060005b8281101561220d576121e06121db8683612247565b61320e565b8282815181106121f2576121f2614b03565b602090810291909101015261220681614bae565b90506121c6565b506122238161221e60016006614be6565b61324d565b8060008151811061223657612236614b03565b602002602001015192505050919050565b60008281612257600c6032614bf9565b6122619085614c0c565b90506fffffffffffffffffffffffffffffffff821681106122c45760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610664565b6122df611ed9826122d7600c6032614bf9565b859190613350565b95945050505050565b604051806122f983602083016133c1565b506fffffffffffffffffffffffffffffffff83166000601f8201601f19168301602001604052509052919050565b600061233283612486565b905060018151600581111561234957612349614426565b1480612367575060028151600581111561236557612365614426565b145b801561238257508363ffffffff16816020015163ffffffff16145b6123ce5760405162461bcd60e51b815260206004820152601f60248201527f536c617368696e6720636f756c64206e6f7420626520696e69746961746564006044820152606401610664565b826001600160a01b03168463ffffffff167f8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e600460405161240f9190614af5565b60405180910390a361242683826040015184612a3f565b61214484848461346a565b604080516060810182526000808252602082018190529181018290529061245a611c9f85613475565b909590945092505050565b6000612475826134a1565b6134a1565b61247e846134a1565b149392505050565b60408051606080820183526000808352602080840182905283850182905260fb54825260fc81528482206001600160a01b038716835290528390208351918201909352825491929091829060ff1660058111156124e5576124e5614426565b60058111156124f6576124f6614426565b8152905463ffffffff610100820481166020840152650100000000009091041660409091015292915050565b60008054610100900460ff16156125bf578160ff1660011480156125455750303b155b6125b75760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610664565b506000919050565b60005460ff80841691161061263c5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610664565b50600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660ff92909216919091179055600190565b600054610100900460ff166126f15760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610664565b60c980546001600160a01b039384167fffffffffffffffffffffffff00000000000000000000000000000000000000009182161790915560ca8054929093169116179055565b600054610100900460ff166127b45760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610664565b61138b6134cc565b8060fb5414610f6d5760fb8190556040518181527f2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa49060200160405180910390a150565b6000816040516020016128139190614c23565b604051602081830303815290604052805190602001209050919050565b60408051602081018490529081018290526000906060015b60405160208183030381529060405280519060200120905092915050565b6000611c7060208084612167565b60cc80546001810182556000919091527f47197230e1e4b29fc0bd84d7d78966c0925452aff72a2a121538b102457e9ebe81016128b18382614cb1565b50919050565b6000611c706128c583612c0f565b613552565b6040805160608101825260008082526020820181905291810182905290611ca5611c9f856135a9565b6000611c70611dd083612dd8565b6000611c70611ed983612c0f565b6000611c7061291d83612c0f565b6135d5565b6040805160608101825260008082526020820181905291810182905290611dfe611c9f8561362c565b600083838360405160200161296293929190614dad565b6040516020818303038152906040528051906020012090509392505050565b8151600090828111156129d65760405162461bcd60e51b815260206004820152600e60248201527f50726f6f6620746f6f206c6f6e670000000000000000000000000000000000006044820152606401610664565b84915060005b81811015612a1357612a09838683815181106129fa576129fa614b03565b60200260200101518984613658565b92506001016129dc565b50805b83811015612a3557612a2b8360008984613658565b9250600101612a16565b5050949350505050565b6001600160a01b038316600090815260cb60205260408082208151606081019092528054829060ff166002811115612a7957612a79614426565b6002811115612a8a57612a8a614426565b81529054610100810463ffffffff1660208301526501000000000090046001600160a01b03166040909101529050600281516002811115612acd57612acd614426565b03612b1a5760405162461bcd60e51b815260206004820152601860248201527f4469737075746520616c7265616479207265736f6c76656400000000000000006044820152606401610664565b6001600160a01b038216604082015260028152612b378482612ef2565b602081015163ffffffff1615612b9757612b97612b73826020015163ffffffff16600090815260fd60205260409020546001600160a01b031690565b60408051606081019091528060008152600060208201819052604090910152612ef2565b612144838260200151613681565b603380546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b8051600090602083016108f481836136d8565b6000612c2d8261373b565b612c795760405162461bcd60e51b815260206004820152601260248201527f4e6f742061207374617465207265706f727400000000000000000000000000006044820152606401610664565b5090565b6000611c707f7919c62641a21cff2eb6e116b4dc34ce70919052c470953e4621535c155ccbc8835b9061378d565b604080516060810182526000808252602082018190529181018290529080612cd2856137b0565b9050612cde81856137eb565b9150612ce982610ce1565b9250600083516005811115612d0057612d00614426565b03612d4d5760405162461bcd60e51b815260206004820152600d60248201527f556e6b6e6f776e206167656e74000000000000000000000000000000000000006044820152606401610664565b509250929050565b6000612d608261380f565b612c795760405162461bcd60e51b815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e00000000000000000000000000006044820152606401610664565b6000611c707f569efb4f951664b562fe9283d8f1a49928bec7335bab838210b64c85e11be59e83612ca5565b6000611c7082600161382e565b6000612df082613894565b612c795760405162461bcd60e51b815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610664565b60008082612e4e6124708260246138ba565b9250612e5e61247082602461382e565b915050915091565b6000611c7082602060046138c7565b6000600182901b60408110612ecc5760405162461bcd60e51b815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e676500000000000000006044820152606401610664565b6000612ed887876138e8565b9050612ee78282876006612981565b979650505050505050565b6001600160a01b038216600090815260cb60205260409020815181548392919082907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836002811115612f4a57612f4a614426565b0217905550602082015181546040938401516001600160a01b031665010000000000027fffffffffffffff0000000000000000000000000000000000000000ffffffffff63ffffffff90931661010002929092167fffffffffffffff000000000000000000000000000000000000000000000000ff90911617179055517f4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f790612ff69084908490614e25565b60405180910390a15050565b60ca546040517fa2155c3400000000000000000000000000000000000000000000000000000000815263ffffffff8085166004830152831660248201526001600160a01b039091169063a2155c34906044015b600060405180830381600087803b15801561306f57600080fd5b505af11580156119e9573d6000803e3d6000fd5b600061308e8261392b565b612c795760405162461bcd60e51b815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610664565b6000816000036130ec57506000610744565b6020821115613127576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166131448385614bf9565b111561317c576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b600061318d8660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b60006131f2600c6032614bf9565b611c70906fffffffffffffffffffffffffffffffff8416614e42565b600080600061321c84612e3c565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b81516001821b8111156132a25760405162461bcd60e51b815260206004820152600e60248201527f48656967687420746f6f206c6f770000000000000000000000000000000000006044820152606401610664565b60005b828110156121445760005b8281101561334157600081600101905060008683815181106132d4576132d4614b03565b6020026020010151905060008583106132ee576000613309565b87838151811061330057613300614b03565b60200260200101515b90506133158282613981565b88600186901c8151811061332b5761332b614b03565b60209081029190910101525050506002016132b0565b506001918201821c91016132a5565b60008061335d8560801c90565b9050613368856139cd565b836133738684614bf9565b61337d9190614bf9565b11156133b5576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6122df848201846136d8565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c908085101561341b576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa90508061345e576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417612ee7565b610f058383836139f3565b6000611c707fdfe02260445526f7b137cb9caf995dcdead56fff547ac8de4b3e33052172314883612ca5565b6000806134ae8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b600054610100900460ff166135495760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610664565b61138b33612ba5565b600061355d82613b09565b612c795760405162461bcd60e51b815260206004820152601960248201527f4e6f7420616e206174746573746174696f6e207265706f7274000000000000006044820152606401610664565b6000611c707fbf180edbd986dd1b6d6de1afe33dbc4c91ee49032bd1af9001bf3a96c95e6fb083612ca5565b60006135e082613b5b565b612c795760405162461bcd60e51b815260206004820152600d60248201527f4e6f7420612072656365697074000000000000000000000000000000000000006044820152606401610664565b6000611c707f293501048791dbdbd4a6187fddcc1046f21c1173ad2502f4b7275f89714771d483612ca5565b6000600183831c168103613677576136708585613981565b90506108f4565b6136708486613981565b60ca546040517f6116921800000000000000000000000000000000000000000000000000000000815263ffffffff8085166004830152831660248201526001600160a01b0390911690636116921890604401613055565b6000806136e58385614bf9565b90506040518111156136f5575060005b8060000361372f576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b83176108f4565b600060016fffffffffffffffffffffffffffffffff8316101561376057506000919050565b600061376b83613b9a565b60ff16111561377c57506000919050565b611c7061378883612dd8565b613894565b600081613799846134a1565b604080516020810193909352820152606001612848565b6040517f19457468657265756d205369676e6564204d6573736167653a0a3332000000006020820152603c8101829052600090605c01612813565b60008060006137fa8585613ba8565b9150915061380781613bea565b509392505050565b6000604e6fffffffffffffffffffffffffffffffff83165b1492915050565b60006fffffffffffffffffffffffffffffffff83168083111561387d576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6108f48361388b8660801c90565b018483036136d8565b60006138a2600c6032614bf9565b6fffffffffffffffffffffffffffffffff8316613827565b6000610744838284613350565b6000806138d58585856130da565b602084900360031b1c9150509392505050565b6000828260405160200161284892919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60006fffffffffffffffffffffffffffffffff82168161394d600c6032614bf9565b6139579083614e42565b905081613966600c6032614bf9565b6139709083614c0c565b1480156108f457506108f481613dd6565b60008215801561398f575081155b1561399c57506000611c70565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050611c70565b60006fffffffffffffffffffffffffffffffff82166139ec8360801c90565b0192915050565b60c9546040805163ffffffff861660248201526001600160a01b03858116604483015284811660648084019190915283518084039091018152608490920183526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f9d228a510000000000000000000000000000000000000000000000000000000017905291517fa1c702a7000000000000000000000000000000000000000000000000000000008152919092169163a1c702a791613ac091600a91620151809190600401614e7d565b60408051808303816000875af1158015613ade573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b029190614ea6565b5050505050565b600060016fffffffffffffffffffffffffffffffff83161015613b2e57506000919050565b6000613b3983613b9a565b60ff161115613b4a57506000919050565b611c70613b5683612dd8565b61380f565b6000613b6960206085614bf9565b6fffffffffffffffffffffffffffffffff831614613b8957506000919050565b611c70613b9583613dfb565b613e09565b6000611c70828260016138c7565b6000808251604103613bde5760208301516040840151606085015160001a613bd287828585613e25565b94509450505050611d02565b50600090506002611d02565b6000816004811115613bfe57613bfe614426565b03613c065750565b6001816004811115613c1a57613c1a614426565b03613c675760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610664565b6002816004811115613c7b57613c7b614426565b03613cc85760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610664565b6003816004811115613cdc57613cdc614426565b03613d4f5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610664565b6004816004811115613d6357613d63614426565b03610f6d5760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202776272076616c60448201527f75650000000000000000000000000000000000000000000000000000000000006064820152608401610664565b60008115801590611c705750613dee60016006614be6565b6001901b82111592915050565b6000611c7082826085613350565b600060856fffffffffffffffffffffffffffffffff8316613827565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115613e5c5750600090506003613f09565b8460ff16601b14158015613e7457508460ff16601c14155b15613e855750600090506004613f09565b6040805160008082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015613ed9573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613f0257600060019250925050613f09565b9150600090505b94509492505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040516060810167ffffffffffffffff81118282101715613f6457613f64613f12565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715613f9357613f93613f12565b604052919050565b600082601f830112613fac57600080fd5b813567ffffffffffffffff811115613fc657613fc6613f12565b613fd96020601f19601f84011601613f6a565b818152846020838601011115613fee57600080fd5b816020850160208301376000918101602001919091529392505050565b600067ffffffffffffffff82111561402557614025613f12565b5060051b60200190565b600082601f83011261404057600080fd5b813560206140556140508361400b565b613f6a565b82815260059290921b8401810191818101908684111561407457600080fd5b8286015b8481101561408f5780358352918301918301614078565b509695505050505050565b60008060008060008060c087890312156140b357600080fd5b86359550602087013567ffffffffffffffff808211156140d257600080fd5b6140de8a838b01613f9b565b965060408901359150808211156140f457600080fd5b6141008a838b01613f9b565b9550606089013591508082111561411657600080fd5b6141228a838b0161402f565b9450608089013591508082111561413857600080fd5b6141448a838b01613f9b565b935060a089013591508082111561415a57600080fd5b5061416789828a01613f9b565b9150509295509295509295565b60008060006060848603121561418957600080fd5b833561ffff8116811461419b57600080fd5b925060208401357fffffffff00000000000000000000000000000000000000000000000000000000811681146141d057600080fd5b929592945050506040919091013590565b80356001600160a01b0381168114610d6257600080fd5b60008060006060848603121561420d57600080fd5b614216846141e1565b9250602084013560ff811681146141d057600080fd5b63ffffffff81168114610f6d57600080fd5b6000806000806080858703121561425457600080fd5b843561425f8161422c565b935060208501359250614274604086016141e1565b9396929550929360600135925050565b6000806000806080858703121561429a57600080fd5b84359350602085013567ffffffffffffffff808211156142b957600080fd5b6142c588838901613f9b565b945060408701359150808211156142db57600080fd5b6142e788838901613f9b565b935060608701359150808211156142fd57600080fd5b5061430a87828801613f9b565b91505092959194509250565b60008060006060848603121561432b57600080fd5b83359250602084013567ffffffffffffffff8082111561434a57600080fd5b61435687838801613f9b565b9350604086013591508082111561436c57600080fd5b5061437986828701613f9b565b9150509250925092565b60008060008060008060c0878903121561439c57600080fd5b86359550602087013567ffffffffffffffff808211156143bb57600080fd5b6143c78a838b01613f9b565b965060408901359150808211156143dd57600080fd5b6143e98a838b01613f9b565b955060608901359150808211156143ff57600080fd5b6141228a838b01613f9b565b60006020828403121561441d57600080fd5b610744826141e1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6006811061446557614465614426565b9052565b614474828251614455565b60208181015163ffffffff9081169184019190915260409182015116910152565b60608101611c708284614469565b6000602082840312156144b557600080fd5b5035919050565b6001600160a01b0383168152608081016107446020830184614469565b8051600381106144eb576144eb614426565b825260208181015163ffffffff16908301526040908101516001600160a01b0316910152565b60608101611c7082846144d9565b6000806040838503121561453257600080fd5b61453b836141e1565b9150614549602084016141e1565b90509250929050565b6000815180845260005b818110156145785760208185018101518683018201520161455c565b506000602082860101526020601f19601f83011685010191505092915050565b6020815260006107446020830184614552565b6000806000606084860312156145c057600080fd5b83356145cb8161422c565b92506145d9602085016141e1565b91506145e7604085016141e1565b90509250925092565b6000806000806080858703121561460657600080fd5b843567ffffffffffffffff8082111561461e57600080fd5b61462a88838901613f9b565b955060209150818701358181111561464157600080fd5b61464d89828a01613f9b565b9550506040870135935060608701358181111561466957600080fd5b87019050601f8101881361467c57600080fd5b803561468a6140508261400b565b81815260059190911b8201830190838101908a8311156146a957600080fd5b928401925b828410156146c7578335825292840192908401906146ae565b979a9699509497505050505050565b600080600080600060a086880312156146ee57600080fd5b85359450602086013567ffffffffffffffff8082111561470d57600080fd5b61471989838a01613f9b565b9550604088013591508082111561472f57600080fd5b61473b89838a01613f9b565b9450606088013591508082111561475157600080fd5b61475d89838a01613f9b565b9350608088013591508082111561477357600080fd5b5061478088828901613f9b565b9150509295509295909350565b6000806000606084860312156147a257600080fd5b833567ffffffffffffffff808211156147ba57600080fd5b6147c687838801613f9b565b9450602086013591508082111561434a57600080fd5b600080600080600060a086880312156147f457600080fd5b85359450602086013567ffffffffffffffff8082111561481357600080fd5b61481f89838a01613f9b565b9550604088013591508082111561483557600080fd5b61473b89838a0161402f565b60008060006060848603121561485657600080fd5b83356148618161422c565b95602085013595506040909401359392505050565b6000806040838503121561488957600080fd5b823567ffffffffffffffff808211156148a157600080fd5b6148ad86838701613f9b565b935060208501359150808211156148c357600080fd5b506148d085828601613f9b565b9150509250929050565b600080600083850360a08112156148f057600080fd5b6148f9856141e1565b93506060601f198201121561490d57600080fd5b50614916613f41565b60208501356006811061492857600080fd5b815260408501356149388161422c565b6020820152606085013561494b8161422c565b60408201529150608084013567ffffffffffffffff81111561496c57600080fd5b6143798682870161402f565b60006020828403121561498a57600080fd5b8151801515811461074457600080fd5b8481526080602082015260006149b36080830186614552565b82810360408401526149c58186614552565b90508281036060840152612ee78185614552565b8381526060602082015260006149f26060830185614552565b8281036040840152614a048185614552565b9695505050505050565b63ffffffff8616815260006020868184015260a06040840152614a3460a0840187614552565b60608401869052838103608085015284518082528286019183019060005b81811015614a805783516fffffffffffffffffffffffffffffffff1683529284019291840191600101614a52565b50909a9950505050505050505050565b63ffffffff851681526001600160a01b0384166020820152608060408201526000614abe6080830185614552565b8281036060840152612ee78185614552565b604081526000614ae36040830185614552565b82810360208401526122df8185614552565b60208101611c708284614455565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600181811c90821680614b4657607f821691505b6020821081036128b1577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203614bdf57614bdf614b7f565b5060010190565b81810381811115611c7057611c70614b7f565b80820180821115611c7057611c70614b7f565b8082028115828204841417611c7057611c70614b7f565b815160009082906020808601845b83811015614c5f5781516fffffffffffffffffffffffffffffffff1685529382019390820190600101614c31565b50929695505050505050565b601f821115610f0557600081815260208120601f850160051c81016020861015614c925750805b601f850160051c820191505b818110156119e957828155600101614c9e565b815167ffffffffffffffff811115614ccb57614ccb613f12565b614cdf81614cd98454614b32565b84614c6b565b602080601f831160018114614d325760008415614cfc5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b1785556119e9565b600085815260208120601f198616915b82811015614d6157888601518255948401946001909101908401614d42565b5085821015614d9d57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b600060068510614dbf57614dbf614426565b5060f89390931b835260e09190911b7fffffffff0000000000000000000000000000000000000000000000000000000016600183015260601b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016600582015260190190565b6001600160a01b03831681526080810161074460208301846144d9565b600082614e78577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600063ffffffff8086168352808516602084015250606060408301526122df6060830184614552565b60008060408385031215614eb957600080fd5b8251614ec48161422c565b602093909301519294929350505056fea264697066735822122065e0f9fd7c4342d8ec2851c07f5a63df77683b0844ec7f766e914bf8d32cc1dc64736f6c63430008110033",
}

// LightManagerHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use LightManagerHarnessMetaData.ABI instead.
var LightManagerHarnessABI = LightManagerHarnessMetaData.ABI

// Deprecated: Use LightManagerHarnessMetaData.Sigs instead.
// LightManagerHarnessFuncSigs maps the 4-byte function signature to its string representation.
var LightManagerHarnessFuncSigs = LightManagerHarnessMetaData.Sigs

// LightManagerHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LightManagerHarnessMetaData.Bin instead.
var LightManagerHarnessBin = LightManagerHarnessMetaData.Bin

// DeployLightManagerHarness deploys a new Ethereum contract, binding an instance of LightManagerHarness to it.
func DeployLightManagerHarness(auth *bind.TransactOpts, backend bind.ContractBackend, domain uint32) (common.Address, *types.Transaction, *LightManagerHarness, error) {
	parsed, err := LightManagerHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LightManagerHarnessBin), backend, domain)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LightManagerHarness{LightManagerHarnessCaller: LightManagerHarnessCaller{contract: contract}, LightManagerHarnessTransactor: LightManagerHarnessTransactor{contract: contract}, LightManagerHarnessFilterer: LightManagerHarnessFilterer{contract: contract}}, nil
}

// LightManagerHarness is an auto generated Go binding around an Ethereum contract.
type LightManagerHarness struct {
	LightManagerHarnessCaller     // Read-only binding to the contract
	LightManagerHarnessTransactor // Write-only binding to the contract
	LightManagerHarnessFilterer   // Log filterer for contract events
}

// LightManagerHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type LightManagerHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightManagerHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LightManagerHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightManagerHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LightManagerHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LightManagerHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LightManagerHarnessSession struct {
	Contract     *LightManagerHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// LightManagerHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LightManagerHarnessCallerSession struct {
	Contract *LightManagerHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// LightManagerHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LightManagerHarnessTransactorSession struct {
	Contract     *LightManagerHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// LightManagerHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type LightManagerHarnessRaw struct {
	Contract *LightManagerHarness // Generic contract binding to access the raw methods on
}

// LightManagerHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LightManagerHarnessCallerRaw struct {
	Contract *LightManagerHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// LightManagerHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LightManagerHarnessTransactorRaw struct {
	Contract *LightManagerHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLightManagerHarness creates a new instance of LightManagerHarness, bound to a specific deployed contract.
func NewLightManagerHarness(address common.Address, backend bind.ContractBackend) (*LightManagerHarness, error) {
	contract, err := bindLightManagerHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LightManagerHarness{LightManagerHarnessCaller: LightManagerHarnessCaller{contract: contract}, LightManagerHarnessTransactor: LightManagerHarnessTransactor{contract: contract}, LightManagerHarnessFilterer: LightManagerHarnessFilterer{contract: contract}}, nil
}

// NewLightManagerHarnessCaller creates a new read-only instance of LightManagerHarness, bound to a specific deployed contract.
func NewLightManagerHarnessCaller(address common.Address, caller bind.ContractCaller) (*LightManagerHarnessCaller, error) {
	contract, err := bindLightManagerHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessCaller{contract: contract}, nil
}

// NewLightManagerHarnessTransactor creates a new write-only instance of LightManagerHarness, bound to a specific deployed contract.
func NewLightManagerHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*LightManagerHarnessTransactor, error) {
	contract, err := bindLightManagerHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessTransactor{contract: contract}, nil
}

// NewLightManagerHarnessFilterer creates a new log filterer instance of LightManagerHarness, bound to a specific deployed contract.
func NewLightManagerHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*LightManagerHarnessFilterer, error) {
	contract, err := bindLightManagerHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessFilterer{contract: contract}, nil
}

// bindLightManagerHarness binds a generic wrapper to an already deployed contract.
func bindLightManagerHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LightManagerHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightManagerHarness *LightManagerHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightManagerHarness.Contract.LightManagerHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightManagerHarness *LightManagerHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.LightManagerHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightManagerHarness *LightManagerHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.LightManagerHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LightManagerHarness *LightManagerHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LightManagerHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LightManagerHarness *LightManagerHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LightManagerHarness *LightManagerHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.contract.Transact(opts, method, params...)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_LightManagerHarness *LightManagerHarnessCaller) AgentRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "agentRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_LightManagerHarness *LightManagerHarnessSession) AgentRoot() ([32]byte, error) {
	return _LightManagerHarness.Contract.AgentRoot(&_LightManagerHarness.CallOpts)
}

// AgentRoot is a free data retrieval call binding the contract method 0x36cba43c.
//
// Solidity: function agentRoot() view returns(bytes32)
func (_LightManagerHarness *LightManagerHarnessCallerSession) AgentRoot() ([32]byte, error) {
	return _LightManagerHarness.Contract.AgentRoot(&_LightManagerHarness.CallOpts)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_LightManagerHarness *LightManagerHarnessCaller) AgentStatus(opts *bind.CallOpts, agent common.Address) (AgentStatus, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "agentStatus", agent)

	if err != nil {
		return *new(AgentStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(AgentStatus)).(*AgentStatus)

	return out0, err

}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_LightManagerHarness *LightManagerHarnessSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _LightManagerHarness.Contract.AgentStatus(&_LightManagerHarness.CallOpts, agent)
}

// AgentStatus is a free data retrieval call binding the contract method 0x28f3fac9.
//
// Solidity: function agentStatus(address agent) view returns((uint8,uint32,uint32) status)
func (_LightManagerHarness *LightManagerHarnessCallerSession) AgentStatus(agent common.Address) (AgentStatus, error) {
	return _LightManagerHarness.Contract.AgentStatus(&_LightManagerHarness.CallOpts, agent)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightManagerHarness *LightManagerHarnessCaller) Destination(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "destination")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightManagerHarness *LightManagerHarnessSession) Destination() (common.Address, error) {
	return _LightManagerHarness.Contract.Destination(&_LightManagerHarness.CallOpts)
}

// Destination is a free data retrieval call binding the contract method 0xb269681d.
//
// Solidity: function destination() view returns(address)
func (_LightManagerHarness *LightManagerHarnessCallerSession) Destination() (common.Address, error) {
	return _LightManagerHarness.Contract.Destination(&_LightManagerHarness.CallOpts)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_LightManagerHarness *LightManagerHarnessCaller) DisputeStatus(opts *bind.CallOpts, agent common.Address) (Dispute, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "disputeStatus", agent)

	if err != nil {
		return *new(Dispute), err
	}

	out0 := *abi.ConvertType(out[0], new(Dispute)).(*Dispute)

	return out0, err

}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_LightManagerHarness *LightManagerHarnessSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _LightManagerHarness.Contract.DisputeStatus(&_LightManagerHarness.CallOpts, agent)
}

// DisputeStatus is a free data retrieval call binding the contract method 0x3463d1b1.
//
// Solidity: function disputeStatus(address agent) view returns((uint8,uint32,address))
func (_LightManagerHarness *LightManagerHarnessCallerSession) DisputeStatus(agent common.Address) (Dispute, error) {
	return _LightManagerHarness.Contract.DisputeStatus(&_LightManagerHarness.CallOpts, agent)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_LightManagerHarness *LightManagerHarnessCaller) GetAgent(opts *bind.CallOpts, index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "getAgent", index)

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
func (_LightManagerHarness *LightManagerHarnessSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _LightManagerHarness.Contract.GetAgent(&_LightManagerHarness.CallOpts, index)
}

// GetAgent is a free data retrieval call binding the contract method 0x2de5aaf7.
//
// Solidity: function getAgent(uint256 index) view returns(address agent, (uint8,uint32,uint32) status)
func (_LightManagerHarness *LightManagerHarnessCallerSession) GetAgent(index *big.Int) (struct {
	Agent  common.Address
	Status AgentStatus
}, error) {
	return _LightManagerHarness.Contract.GetAgent(&_LightManagerHarness.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightManagerHarness *LightManagerHarnessCaller) GetStoredSignature(opts *bind.CallOpts, index *big.Int) ([]byte, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "getStoredSignature", index)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightManagerHarness *LightManagerHarnessSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightManagerHarness.Contract.GetStoredSignature(&_LightManagerHarness.CallOpts, index)
}

// GetStoredSignature is a free data retrieval call binding the contract method 0xddeffa66.
//
// Solidity: function getStoredSignature(uint256 index) view returns(bytes)
func (_LightManagerHarness *LightManagerHarnessCallerSession) GetStoredSignature(index *big.Int) ([]byte, error) {
	return _LightManagerHarness.Contract.GetStoredSignature(&_LightManagerHarness.CallOpts, index)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightManagerHarness *LightManagerHarnessCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightManagerHarness *LightManagerHarnessSession) LocalDomain() (uint32, error) {
	return _LightManagerHarness.Contract.LocalDomain(&_LightManagerHarness.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_LightManagerHarness *LightManagerHarnessCallerSession) LocalDomain() (uint32, error) {
	return _LightManagerHarness.Contract.LocalDomain(&_LightManagerHarness.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightManagerHarness *LightManagerHarnessCaller) Origin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "origin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightManagerHarness *LightManagerHarnessSession) Origin() (common.Address, error) {
	return _LightManagerHarness.Contract.Origin(&_LightManagerHarness.CallOpts)
}

// Origin is a free data retrieval call binding the contract method 0x938b5f32.
//
// Solidity: function origin() view returns(address)
func (_LightManagerHarness *LightManagerHarnessCallerSession) Origin() (common.Address, error) {
	return _LightManagerHarness.Contract.Origin(&_LightManagerHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightManagerHarness *LightManagerHarnessCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightManagerHarness *LightManagerHarnessSession) Owner() (common.Address, error) {
	return _LightManagerHarness.Contract.Owner(&_LightManagerHarness.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LightManagerHarness *LightManagerHarnessCallerSession) Owner() (common.Address, error) {
	return _LightManagerHarness.Contract.Owner(&_LightManagerHarness.CallOpts)
}

// RemoteMockFunc is a free data retrieval call binding the contract method 0xa149352c.
//
// Solidity: function remoteMockFunc(uint32 , uint256 , bytes32 ) view returns(bytes4)
func (_LightManagerHarness *LightManagerHarnessCaller) RemoteMockFunc(opts *bind.CallOpts, arg0 uint32, arg1 *big.Int, arg2 [32]byte) ([4]byte, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "remoteMockFunc", arg0, arg1, arg2)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// RemoteMockFunc is a free data retrieval call binding the contract method 0xa149352c.
//
// Solidity: function remoteMockFunc(uint32 , uint256 , bytes32 ) view returns(bytes4)
func (_LightManagerHarness *LightManagerHarnessSession) RemoteMockFunc(arg0 uint32, arg1 *big.Int, arg2 [32]byte) ([4]byte, error) {
	return _LightManagerHarness.Contract.RemoteMockFunc(&_LightManagerHarness.CallOpts, arg0, arg1, arg2)
}

// RemoteMockFunc is a free data retrieval call binding the contract method 0xa149352c.
//
// Solidity: function remoteMockFunc(uint32 , uint256 , bytes32 ) view returns(bytes4)
func (_LightManagerHarness *LightManagerHarnessCallerSession) RemoteMockFunc(arg0 uint32, arg1 *big.Int, arg2 [32]byte) ([4]byte, error) {
	return _LightManagerHarness.Contract.RemoteMockFunc(&_LightManagerHarness.CallOpts, arg0, arg1, arg2)
}

// SensitiveMockFunc is a free data retrieval call binding the contract method 0x127a2c9d.
//
// Solidity: function sensitiveMockFunc(address , uint8 , bytes32 data) view returns(bytes32)
func (_LightManagerHarness *LightManagerHarnessCaller) SensitiveMockFunc(opts *bind.CallOpts, arg0 common.Address, arg1 uint8, data [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "sensitiveMockFunc", arg0, arg1, data)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SensitiveMockFunc is a free data retrieval call binding the contract method 0x127a2c9d.
//
// Solidity: function sensitiveMockFunc(address , uint8 , bytes32 data) view returns(bytes32)
func (_LightManagerHarness *LightManagerHarnessSession) SensitiveMockFunc(arg0 common.Address, arg1 uint8, data [32]byte) ([32]byte, error) {
	return _LightManagerHarness.Contract.SensitiveMockFunc(&_LightManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFunc is a free data retrieval call binding the contract method 0x127a2c9d.
//
// Solidity: function sensitiveMockFunc(address , uint8 , bytes32 data) view returns(bytes32)
func (_LightManagerHarness *LightManagerHarnessCallerSession) SensitiveMockFunc(arg0 common.Address, arg1 uint8, data [32]byte) ([32]byte, error) {
	return _LightManagerHarness.Contract.SensitiveMockFunc(&_LightManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFuncOver32Bytes is a free data retrieval call binding the contract method 0x0e6bfcd5.
//
// Solidity: function sensitiveMockFuncOver32Bytes(uint16 , bytes4 , bytes32 data) view returns(bytes4, bytes32)
func (_LightManagerHarness *LightManagerHarnessCaller) SensitiveMockFuncOver32Bytes(opts *bind.CallOpts, arg0 uint16, arg1 [4]byte, data [32]byte) ([4]byte, [32]byte, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "sensitiveMockFuncOver32Bytes", arg0, arg1, data)

	if err != nil {
		return *new([4]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// SensitiveMockFuncOver32Bytes is a free data retrieval call binding the contract method 0x0e6bfcd5.
//
// Solidity: function sensitiveMockFuncOver32Bytes(uint16 , bytes4 , bytes32 data) view returns(bytes4, bytes32)
func (_LightManagerHarness *LightManagerHarnessSession) SensitiveMockFuncOver32Bytes(arg0 uint16, arg1 [4]byte, data [32]byte) ([4]byte, [32]byte, error) {
	return _LightManagerHarness.Contract.SensitiveMockFuncOver32Bytes(&_LightManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFuncOver32Bytes is a free data retrieval call binding the contract method 0x0e6bfcd5.
//
// Solidity: function sensitiveMockFuncOver32Bytes(uint16 , bytes4 , bytes32 data) view returns(bytes4, bytes32)
func (_LightManagerHarness *LightManagerHarnessCallerSession) SensitiveMockFuncOver32Bytes(arg0 uint16, arg1 [4]byte, data [32]byte) ([4]byte, [32]byte, error) {
	return _LightManagerHarness.Contract.SensitiveMockFuncOver32Bytes(&_LightManagerHarness.CallOpts, arg0, arg1, data)
}

// SensitiveMockFuncVoid is a free data retrieval call binding the contract method 0xc9f1a03f.
//
// Solidity: function sensitiveMockFuncVoid(uint16 , bytes4 , bytes32 ) view returns()
func (_LightManagerHarness *LightManagerHarnessCaller) SensitiveMockFuncVoid(opts *bind.CallOpts, arg0 uint16, arg1 [4]byte, arg2 [32]byte) error {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "sensitiveMockFuncVoid", arg0, arg1, arg2)

	if err != nil {
		return err
	}

	return err

}

// SensitiveMockFuncVoid is a free data retrieval call binding the contract method 0xc9f1a03f.
//
// Solidity: function sensitiveMockFuncVoid(uint16 , bytes4 , bytes32 ) view returns()
func (_LightManagerHarness *LightManagerHarnessSession) SensitiveMockFuncVoid(arg0 uint16, arg1 [4]byte, arg2 [32]byte) error {
	return _LightManagerHarness.Contract.SensitiveMockFuncVoid(&_LightManagerHarness.CallOpts, arg0, arg1, arg2)
}

// SensitiveMockFuncVoid is a free data retrieval call binding the contract method 0xc9f1a03f.
//
// Solidity: function sensitiveMockFuncVoid(uint16 , bytes4 , bytes32 ) view returns()
func (_LightManagerHarness *LightManagerHarnessCallerSession) SensitiveMockFuncVoid(arg0 uint16, arg1 [4]byte, arg2 [32]byte) error {
	return _LightManagerHarness.Contract.SensitiveMockFuncVoid(&_LightManagerHarness.CallOpts, arg0, arg1, arg2)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightManagerHarness *LightManagerHarnessCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LightManagerHarness.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightManagerHarness *LightManagerHarnessSession) Version() (string, error) {
	return _LightManagerHarness.Contract.Version(&_LightManagerHarness.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string versionString)
func (_LightManagerHarness *LightManagerHarnessCallerSession) Version() (string, error) {
	return _LightManagerHarness.Contract.Version(&_LightManagerHarness.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_LightManagerHarness *LightManagerHarnessTransactor) Initialize(opts *bind.TransactOpts, origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "initialize", origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_LightManagerHarness *LightManagerHarnessSession) Initialize(origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.Initialize(&_LightManagerHarness.TransactOpts, origin_, destination_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address origin_, address destination_) returns()
func (_LightManagerHarness *LightManagerHarnessTransactorSession) Initialize(origin_ common.Address, destination_ common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.Initialize(&_LightManagerHarness.TransactOpts, origin_, destination_)
}

// OpenDisputeExposed is a paid mutator transaction binding the contract method 0xc166ca67.
//
// Solidity: function openDisputeExposed(address guard, address notary) returns()
func (_LightManagerHarness *LightManagerHarnessTransactor) OpenDisputeExposed(opts *bind.TransactOpts, guard common.Address, notary common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "openDisputeExposed", guard, notary)
}

// OpenDisputeExposed is a paid mutator transaction binding the contract method 0xc166ca67.
//
// Solidity: function openDisputeExposed(address guard, address notary) returns()
func (_LightManagerHarness *LightManagerHarnessSession) OpenDisputeExposed(guard common.Address, notary common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.OpenDisputeExposed(&_LightManagerHarness.TransactOpts, guard, notary)
}

// OpenDisputeExposed is a paid mutator transaction binding the contract method 0xc166ca67.
//
// Solidity: function openDisputeExposed(address guard, address notary) returns()
func (_LightManagerHarness *LightManagerHarnessTransactorSession) OpenDisputeExposed(guard common.Address, notary common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.OpenDisputeExposed(&_LightManagerHarness.TransactOpts, guard, notary)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_LightManagerHarness *LightManagerHarnessTransactor) RemoteWithdrawTips(opts *bind.TransactOpts, msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "remoteWithdrawTips", msgOrigin, proofMaturity, recipient, amount)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_LightManagerHarness *LightManagerHarnessSession) RemoteWithdrawTips(msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.RemoteWithdrawTips(&_LightManagerHarness.TransactOpts, msgOrigin, proofMaturity, recipient, amount)
}

// RemoteWithdrawTips is a paid mutator transaction binding the contract method 0x1fa07138.
//
// Solidity: function remoteWithdrawTips(uint32 msgOrigin, uint256 proofMaturity, address recipient, uint256 amount) returns(bytes4 magicValue)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) RemoteWithdrawTips(msgOrigin uint32, proofMaturity *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.RemoteWithdrawTips(&_LightManagerHarness.TransactOpts, msgOrigin, proofMaturity, recipient, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightManagerHarness *LightManagerHarnessTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightManagerHarness *LightManagerHarnessSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightManagerHarness.Contract.RenounceOwnership(&_LightManagerHarness.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LightManagerHarness *LightManagerHarnessTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LightManagerHarness.Contract.RenounceOwnership(&_LightManagerHarness.TransactOpts)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot_) returns()
func (_LightManagerHarness *LightManagerHarnessTransactor) SetAgentRoot(opts *bind.TransactOpts, agentRoot_ [32]byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "setAgentRoot", agentRoot_)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot_) returns()
func (_LightManagerHarness *LightManagerHarnessSession) SetAgentRoot(agentRoot_ [32]byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SetAgentRoot(&_LightManagerHarness.TransactOpts, agentRoot_)
}

// SetAgentRoot is a paid mutator transaction binding the contract method 0x58668176.
//
// Solidity: function setAgentRoot(bytes32 agentRoot_) returns()
func (_LightManagerHarness *LightManagerHarnessTransactorSession) SetAgentRoot(agentRoot_ [32]byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SetAgentRoot(&_LightManagerHarness.TransactOpts, agentRoot_)
}

// SlashAgentExposed is a paid mutator transaction binding the contract method 0x69978b0d.
//
// Solidity: function slashAgentExposed(uint32 domain, address agent, address prover) returns()
func (_LightManagerHarness *LightManagerHarnessTransactor) SlashAgentExposed(opts *bind.TransactOpts, domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "slashAgentExposed", domain, agent, prover)
}

// SlashAgentExposed is a paid mutator transaction binding the contract method 0x69978b0d.
//
// Solidity: function slashAgentExposed(uint32 domain, address agent, address prover) returns()
func (_LightManagerHarness *LightManagerHarnessSession) SlashAgentExposed(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SlashAgentExposed(&_LightManagerHarness.TransactOpts, domain, agent, prover)
}

// SlashAgentExposed is a paid mutator transaction binding the contract method 0x69978b0d.
//
// Solidity: function slashAgentExposed(uint32 domain, address agent, address prover) returns()
func (_LightManagerHarness *LightManagerHarnessTransactorSession) SlashAgentExposed(domain uint32, agent common.Address, prover common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SlashAgentExposed(&_LightManagerHarness.TransactOpts, domain, agent, prover)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactor) SubmitAttestation(opts *bind.TransactOpts, attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "submitAttestation", attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitAttestation(&_LightManagerHarness.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestation is a paid mutator transaction binding the contract method 0x6c38f723.
//
// Solidity: function submitAttestation(bytes attPayload, bytes attSignature, bytes32 agentRoot_, uint256[] snapGas_) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) SubmitAttestation(attPayload []byte, attSignature []byte, agentRoot_ [32]byte, snapGas_ []*big.Int) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitAttestation(&_LightManagerHarness.TransactOpts, attPayload, attSignature, agentRoot_, snapGas_)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactor) SubmitAttestationReport(opts *bind.TransactOpts, arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "submitAttestationReport", arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitAttestationReport(&_LightManagerHarness.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitAttestationReport is a paid mutator transaction binding the contract method 0x77ec5c10.
//
// Solidity: function submitAttestationReport(bytes arPayload, bytes arSignature, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) SubmitAttestationReport(arPayload []byte, arSignature []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitAttestationReport(&_LightManagerHarness.TransactOpts, arPayload, arSignature, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactor) SubmitStateReportWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "submitStateReportWithAttestation", stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitStateReportWithAttestation(&_LightManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithAttestation is a paid mutator transaction binding the contract method 0x235d51b1.
//
// Solidity: function submitStateReportWithAttestation(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) SubmitStateReportWithAttestation(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitStateReportWithAttestation(&_LightManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, attPayload, attSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactor) SubmitStateReportWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "submitStateReportWithSnapshot", stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitStateReportWithSnapshot(&_LightManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshot is a paid mutator transaction binding the contract method 0x708cdc82.
//
// Solidity: function submitStateReportWithSnapshot(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes snapPayload, bytes snapSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) SubmitStateReportWithSnapshot(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitStateReportWithSnapshot(&_LightManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapPayload, snapSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactor) SubmitStateReportWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "submitStateReportWithSnapshotProof", stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitStateReportWithSnapshotProof(&_LightManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// SubmitStateReportWithSnapshotProof is a paid mutator transaction binding the contract method 0x0db27e77.
//
// Solidity: function submitStateReportWithSnapshotProof(uint256 stateIndex, bytes srPayload, bytes srSignature, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool wasAccepted)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) SubmitStateReportWithSnapshotProof(stateIndex *big.Int, srPayload []byte, srSignature []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.SubmitStateReportWithSnapshotProof(&_LightManagerHarness.TransactOpts, stateIndex, srPayload, srSignature, snapProof, attPayload, attSignature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightManagerHarness *LightManagerHarnessTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightManagerHarness *LightManagerHarnessSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.TransferOwnership(&_LightManagerHarness.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LightManagerHarness *LightManagerHarnessTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.TransferOwnership(&_LightManagerHarness.TransactOpts, newOwner)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_LightManagerHarness *LightManagerHarnessTransactor) UpdateAgentStatus(opts *bind.TransactOpts, agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "updateAgentStatus", agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_LightManagerHarness *LightManagerHarnessSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.UpdateAgentStatus(&_LightManagerHarness.TransactOpts, agent, status, proof)
}

// UpdateAgentStatus is a paid mutator transaction binding the contract method 0xcbd05965.
//
// Solidity: function updateAgentStatus(address agent, (uint8,uint32,uint32) status, bytes32[] proof) returns()
func (_LightManagerHarness *LightManagerHarnessTransactorSession) UpdateAgentStatus(agent common.Address, status AgentStatus, proof [][32]byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.UpdateAgentStatus(&_LightManagerHarness.TransactOpts, agent, status, proof)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightManagerHarness *LightManagerHarnessTransactor) VerifyReceipt(opts *bind.TransactOpts, rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "verifyReceipt", rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightManagerHarness *LightManagerHarnessSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyReceipt(&_LightManagerHarness.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyReceipt is a paid mutator transaction binding the contract method 0xc25aa585.
//
// Solidity: function verifyReceipt(bytes rcptPayload, bytes rcptSignature) returns(bool isValidReceipt)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) VerifyReceipt(rcptPayload []byte, rcptSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyReceipt(&_LightManagerHarness.TransactOpts, rcptPayload, rcptSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_LightManagerHarness *LightManagerHarnessTransactor) VerifyStateReport(opts *bind.TransactOpts, srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "verifyStateReport", srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_LightManagerHarness *LightManagerHarnessSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateReport(&_LightManagerHarness.TransactOpts, srPayload, srSignature)
}

// VerifyStateReport is a paid mutator transaction binding the contract method 0xdfe39675.
//
// Solidity: function verifyStateReport(bytes srPayload, bytes srSignature) returns(bool isValidReport)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) VerifyStateReport(srPayload []byte, srSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateReport(&_LightManagerHarness.TransactOpts, srPayload, srSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessTransactor) VerifyStateWithAttestation(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "verifyStateWithAttestation", stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateWithAttestation(&_LightManagerHarness.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithAttestation is a paid mutator transaction binding the contract method 0x200f6b66.
//
// Solidity: function verifyStateWithAttestation(uint256 stateIndex, bytes snapPayload, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) VerifyStateWithAttestation(stateIndex *big.Int, snapPayload []byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateWithAttestation(&_LightManagerHarness.TransactOpts, stateIndex, snapPayload, attPayload, attSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessTransactor) VerifyStateWithSnapshot(opts *bind.TransactOpts, stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "verifyStateWithSnapshot", stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateWithSnapshot(&_LightManagerHarness.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshot is a paid mutator transaction binding the contract method 0x213a6ddb.
//
// Solidity: function verifyStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) VerifyStateWithSnapshot(stateIndex *big.Int, snapPayload []byte, snapSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateWithSnapshot(&_LightManagerHarness.TransactOpts, stateIndex, snapPayload, snapSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessTransactor) VerifyStateWithSnapshotProof(opts *bind.TransactOpts, stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.contract.Transact(opts, "verifyStateWithSnapshotProof", stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateWithSnapshotProof(&_LightManagerHarness.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// VerifyStateWithSnapshotProof is a paid mutator transaction binding the contract method 0x7be8e738.
//
// Solidity: function verifyStateWithSnapshotProof(uint256 stateIndex, bytes statePayload, bytes32[] snapProof, bytes attPayload, bytes attSignature) returns(bool isValidState)
func (_LightManagerHarness *LightManagerHarnessTransactorSession) VerifyStateWithSnapshotProof(stateIndex *big.Int, statePayload []byte, snapProof [][32]byte, attPayload []byte, attSignature []byte) (*types.Transaction, error) {
	return _LightManagerHarness.Contract.VerifyStateWithSnapshotProof(&_LightManagerHarness.TransactOpts, stateIndex, statePayload, snapProof, attPayload, attSignature)
}

// LightManagerHarnessAttestationAcceptedIterator is returned from FilterAttestationAccepted and is used to iterate over the raw logs and unpacked data for AttestationAccepted events raised by the LightManagerHarness contract.
type LightManagerHarnessAttestationAcceptedIterator struct {
	Event *LightManagerHarnessAttestationAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessAttestationAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessAttestationAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessAttestationAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessAttestationAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessAttestationAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessAttestationAccepted represents a AttestationAccepted event raised by the LightManagerHarness contract.
type LightManagerHarnessAttestationAccepted struct {
	Domain       uint32
	Notary       common.Address
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAttestationAccepted is a free log retrieval operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterAttestationAccepted(opts *bind.FilterOpts) (*LightManagerHarnessAttestationAcceptedIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessAttestationAcceptedIterator{contract: _LightManagerHarness.contract, event: "AttestationAccepted", logs: logs, sub: sub}, nil
}

// WatchAttestationAccepted is a free log subscription operation binding the contract event 0x5fb28b72a4ff089027990125e187d936f30d65013d66fac1e54e0625f7ea0065.
//
// Solidity: event AttestationAccepted(uint32 domain, address notary, bytes attPayload, bytes attSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchAttestationAccepted(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessAttestationAccepted) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "AttestationAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessAttestationAccepted)
				if err := _LightManagerHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseAttestationAccepted(log types.Log) (*LightManagerHarnessAttestationAccepted, error) {
	event := new(LightManagerHarnessAttestationAccepted)
	if err := _LightManagerHarness.contract.UnpackLog(event, "AttestationAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessDisputeUpdatedIterator is returned from FilterDisputeUpdated and is used to iterate over the raw logs and unpacked data for DisputeUpdated events raised by the LightManagerHarness contract.
type LightManagerHarnessDisputeUpdatedIterator struct {
	Event *LightManagerHarnessDisputeUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessDisputeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessDisputeUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessDisputeUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessDisputeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessDisputeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessDisputeUpdated represents a DisputeUpdated event raised by the LightManagerHarness contract.
type LightManagerHarnessDisputeUpdated struct {
	Agent   common.Address
	Dispute Dispute
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDisputeUpdated is a free log retrieval operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterDisputeUpdated(opts *bind.FilterOpts) (*LightManagerHarnessDisputeUpdatedIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessDisputeUpdatedIterator{contract: _LightManagerHarness.contract, event: "DisputeUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeUpdated is a free log subscription operation binding the contract event 0x4a4274ddec5680bd48ca531f9f0fb4cbcf1598a0e55deb8171fd95e99673e0f7.
//
// Solidity: event DisputeUpdated(address agent, (uint8,uint32,address) dispute)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchDisputeUpdated(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessDisputeUpdated) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "DisputeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessDisputeUpdated)
				if err := _LightManagerHarness.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseDisputeUpdated(log types.Log) (*LightManagerHarnessDisputeUpdated, error) {
	event := new(LightManagerHarnessDisputeUpdated)
	if err := _LightManagerHarness.contract.UnpackLog(event, "DisputeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the LightManagerHarness contract.
type LightManagerHarnessInitializedIterator struct {
	Event *LightManagerHarnessInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessInitialized represents a Initialized event raised by the LightManagerHarness contract.
type LightManagerHarnessInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterInitialized(opts *bind.FilterOpts) (*LightManagerHarnessInitializedIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessInitializedIterator{contract: _LightManagerHarness.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessInitialized) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessInitialized)
				if err := _LightManagerHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseInitialized(log types.Log) (*LightManagerHarnessInitialized, error) {
	event := new(LightManagerHarnessInitialized)
	if err := _LightManagerHarness.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessInvalidReceiptIterator is returned from FilterInvalidReceipt and is used to iterate over the raw logs and unpacked data for InvalidReceipt events raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidReceiptIterator struct {
	Event *LightManagerHarnessInvalidReceipt // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessInvalidReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessInvalidReceipt)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessInvalidReceipt)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessInvalidReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessInvalidReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessInvalidReceipt represents a InvalidReceipt event raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidReceipt struct {
	RcptPayload   []byte
	RcptSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidReceipt is a free log retrieval operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterInvalidReceipt(opts *bind.FilterOpts) (*LightManagerHarnessInvalidReceiptIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessInvalidReceiptIterator{contract: _LightManagerHarness.contract, event: "InvalidReceipt", logs: logs, sub: sub}, nil
}

// WatchInvalidReceipt is a free log subscription operation binding the contract event 0x4d4c3a87f0d5fbcea3c51d5baa727fceedb200dd7c9287f7ef85b60b794d6a8d.
//
// Solidity: event InvalidReceipt(bytes rcptPayload, bytes rcptSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchInvalidReceipt(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessInvalidReceipt) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "InvalidReceipt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessInvalidReceipt)
				if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseInvalidReceipt(log types.Log) (*LightManagerHarnessInvalidReceipt, error) {
	event := new(LightManagerHarnessInvalidReceipt)
	if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessInvalidStateReportIterator is returned from FilterInvalidStateReport and is used to iterate over the raw logs and unpacked data for InvalidStateReport events raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidStateReportIterator struct {
	Event *LightManagerHarnessInvalidStateReport // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessInvalidStateReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessInvalidStateReport)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessInvalidStateReport)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessInvalidStateReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessInvalidStateReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessInvalidStateReport represents a InvalidStateReport event raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidStateReport struct {
	SrPayload   []byte
	SrSignature []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateReport is a free log retrieval operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterInvalidStateReport(opts *bind.FilterOpts) (*LightManagerHarnessInvalidStateReportIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessInvalidStateReportIterator{contract: _LightManagerHarness.contract, event: "InvalidStateReport", logs: logs, sub: sub}, nil
}

// WatchInvalidStateReport is a free log subscription operation binding the contract event 0x9b0db5e74572fe0188dcef5afafe498161864c5706c3003c98ee506ae5c0282d.
//
// Solidity: event InvalidStateReport(bytes srPayload, bytes srSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchInvalidStateReport(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessInvalidStateReport) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "InvalidStateReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessInvalidStateReport)
				if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseInvalidStateReport(log types.Log) (*LightManagerHarnessInvalidStateReport, error) {
	event := new(LightManagerHarnessInvalidStateReport)
	if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidStateReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessInvalidStateWithAttestationIterator is returned from FilterInvalidStateWithAttestation and is used to iterate over the raw logs and unpacked data for InvalidStateWithAttestation events raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidStateWithAttestationIterator struct {
	Event *LightManagerHarnessInvalidStateWithAttestation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessInvalidStateWithAttestationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessInvalidStateWithAttestation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessInvalidStateWithAttestation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessInvalidStateWithAttestationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessInvalidStateWithAttestationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessInvalidStateWithAttestation represents a InvalidStateWithAttestation event raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidStateWithAttestation struct {
	StateIndex   *big.Int
	StatePayload []byte
	AttPayload   []byte
	AttSignature []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithAttestation is a free log retrieval operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterInvalidStateWithAttestation(opts *bind.FilterOpts) (*LightManagerHarnessInvalidStateWithAttestationIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessInvalidStateWithAttestationIterator{contract: _LightManagerHarness.contract, event: "InvalidStateWithAttestation", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithAttestation is a free log subscription operation binding the contract event 0x541491c63a99c21d0612ba7b3c4d90f7662f54a123363e65fae5c51e34e8a492.
//
// Solidity: event InvalidStateWithAttestation(uint256 stateIndex, bytes statePayload, bytes attPayload, bytes attSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchInvalidStateWithAttestation(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessInvalidStateWithAttestation) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "InvalidStateWithAttestation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessInvalidStateWithAttestation)
				if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseInvalidStateWithAttestation(log types.Log) (*LightManagerHarnessInvalidStateWithAttestation, error) {
	event := new(LightManagerHarnessInvalidStateWithAttestation)
	if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidStateWithAttestation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessInvalidStateWithSnapshotIterator is returned from FilterInvalidStateWithSnapshot and is used to iterate over the raw logs and unpacked data for InvalidStateWithSnapshot events raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidStateWithSnapshotIterator struct {
	Event *LightManagerHarnessInvalidStateWithSnapshot // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessInvalidStateWithSnapshotIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessInvalidStateWithSnapshot)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessInvalidStateWithSnapshot)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessInvalidStateWithSnapshotIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessInvalidStateWithSnapshotIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessInvalidStateWithSnapshot represents a InvalidStateWithSnapshot event raised by the LightManagerHarness contract.
type LightManagerHarnessInvalidStateWithSnapshot struct {
	StateIndex    *big.Int
	SnapPayload   []byte
	SnapSignature []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterInvalidStateWithSnapshot is a free log retrieval operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterInvalidStateWithSnapshot(opts *bind.FilterOpts) (*LightManagerHarnessInvalidStateWithSnapshotIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessInvalidStateWithSnapshotIterator{contract: _LightManagerHarness.contract, event: "InvalidStateWithSnapshot", logs: logs, sub: sub}, nil
}

// WatchInvalidStateWithSnapshot is a free log subscription operation binding the contract event 0x8ec8481d2e38a1ffe2c6ec35459332dc9e6248b1219d287dfa4143a68a75bbd1.
//
// Solidity: event InvalidStateWithSnapshot(uint256 stateIndex, bytes snapPayload, bytes snapSignature)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchInvalidStateWithSnapshot(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessInvalidStateWithSnapshot) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "InvalidStateWithSnapshot")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessInvalidStateWithSnapshot)
				if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseInvalidStateWithSnapshot(log types.Log) (*LightManagerHarnessInvalidStateWithSnapshot, error) {
	event := new(LightManagerHarnessInvalidStateWithSnapshot)
	if err := _LightManagerHarness.contract.UnpackLog(event, "InvalidStateWithSnapshot", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LightManagerHarness contract.
type LightManagerHarnessOwnershipTransferredIterator struct {
	Event *LightManagerHarnessOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessOwnershipTransferred represents a OwnershipTransferred event raised by the LightManagerHarness contract.
type LightManagerHarnessOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LightManagerHarnessOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessOwnershipTransferredIterator{contract: _LightManagerHarness.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessOwnershipTransferred)
				if err := _LightManagerHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseOwnershipTransferred(log types.Log) (*LightManagerHarnessOwnershipTransferred, error) {
	event := new(LightManagerHarnessOwnershipTransferred)
	if err := _LightManagerHarness.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the LightManagerHarness contract.
type LightManagerHarnessRootUpdatedIterator struct {
	Event *LightManagerHarnessRootUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessRootUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessRootUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessRootUpdated represents a RootUpdated event raised by the LightManagerHarness contract.
type LightManagerHarnessRootUpdated struct {
	NewRoot [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterRootUpdated(opts *bind.FilterOpts) (*LightManagerHarnessRootUpdatedIterator, error) {

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessRootUpdatedIterator{contract: _LightManagerHarness.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 newRoot)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessRootUpdated) (event.Subscription, error) {

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "RootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessRootUpdated)
				if err := _LightManagerHarness.contract.UnpackLog(event, "RootUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseRootUpdated(log types.Log) (*LightManagerHarnessRootUpdated, error) {
	event := new(LightManagerHarnessRootUpdated)
	if err := _LightManagerHarness.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LightManagerHarnessStatusUpdatedIterator is returned from FilterStatusUpdated and is used to iterate over the raw logs and unpacked data for StatusUpdated events raised by the LightManagerHarness contract.
type LightManagerHarnessStatusUpdatedIterator struct {
	Event *LightManagerHarnessStatusUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LightManagerHarnessStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LightManagerHarnessStatusUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LightManagerHarnessStatusUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LightManagerHarnessStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LightManagerHarnessStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LightManagerHarnessStatusUpdated represents a StatusUpdated event raised by the LightManagerHarness contract.
type LightManagerHarnessStatusUpdated struct {
	Flag   uint8
	Domain uint32
	Agent  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStatusUpdated is a free log retrieval operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_LightManagerHarness *LightManagerHarnessFilterer) FilterStatusUpdated(opts *bind.FilterOpts, domain []uint32, agent []common.Address) (*LightManagerHarnessStatusUpdatedIterator, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _LightManagerHarness.contract.FilterLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return &LightManagerHarnessStatusUpdatedIterator{contract: _LightManagerHarness.contract, event: "StatusUpdated", logs: logs, sub: sub}, nil
}

// WatchStatusUpdated is a free log subscription operation binding the contract event 0x8f9b8b0f4f062833bec85ea9a8465e4a1207b4be6eb565bbd0ae8b913588d04e.
//
// Solidity: event StatusUpdated(uint8 flag, uint32 indexed domain, address indexed agent)
func (_LightManagerHarness *LightManagerHarnessFilterer) WatchStatusUpdated(opts *bind.WatchOpts, sink chan<- *LightManagerHarnessStatusUpdated, domain []uint32, agent []common.Address) (event.Subscription, error) {

	var domainRule []interface{}
	for _, domainItem := range domain {
		domainRule = append(domainRule, domainItem)
	}
	var agentRule []interface{}
	for _, agentItem := range agent {
		agentRule = append(agentRule, agentItem)
	}

	logs, sub, err := _LightManagerHarness.contract.WatchLogs(opts, "StatusUpdated", domainRule, agentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LightManagerHarnessStatusUpdated)
				if err := _LightManagerHarness.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_LightManagerHarness *LightManagerHarnessFilterer) ParseStatusUpdated(log types.Log) (*LightManagerHarnessStatusUpdated, error) {
	event := new(LightManagerHarnessStatusUpdated)
	if err := _LightManagerHarness.contract.UnpackLog(event, "StatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122070eaa50dfba157234d05a4072a11f9baf061011f59cc0faa29f3254544ef2fc364736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c99fd1e097f4afbaf8ad182202b12c1664b15f62bd43ca2b11e7f67a2e537e8f64736f6c63430008110033",
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

// NumberLibMetaData contains all meta data concerning the NumberLib contract.
var NumberLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a72e000a9e8edfeba65de3c78483ab6bebe2c50926820a21b023108a8a461a6764736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220553583fd5782568502a975588eb3a27afc69f324d16f72dd7f31347e257100fa64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122092960e6632f76a41647d743746b748955b7e92466a9a726ecbc9a2383b5b60fc64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220db5810c2cd78f02afdefd5c1ef3145dbee5a0b8a1fab0e07b6faf1e8bd541caf64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122017c1fac4811ddd919b66ce0ab1a752e040a5de0488b40b71f227eb3aae8ec1d664736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209f0f41b4563b4ee5fd6ff5596c4a7d8d27b9821ac5317c3402fa6de243998f7864736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220695e16c15ae80b645b86acd2ef1428562ffb04f949d0b0eb0594b2ac32ec021864736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a05a3bced092145dca75ae2af97acc128bbdeb2b3596ac9885b76ade62e8dc2164736f6c63430008110033",
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
