// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dfkpet

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

// Pet is an auto generated low-level Go binding around an user-defined struct.
type Pet struct {
	Id                *big.Int
	OriginId          uint8
	Name              string
	Season            uint8
	EggType           uint8
	Rarity            uint8
	Element           uint8
	BonusCount        uint8
	ProfBonus         uint8
	ProfBonusScalar   uint8
	CraftBonus        uint8
	CraftBonusScalar  uint8
	CombatBonus       uint8
	CombatBonusScalar uint8
	Appearance        uint16
	Background        uint8
	Shiny             uint8
	HungryAt          uint64
	EquippableAt      uint64
	EquippedTo        *big.Int
}

// PetBridgeUpgradeableMessageFormat is an auto generated low-level Go binding around an user-defined struct.
type PetBridgeUpgradeableMessageFormat struct {
	DstPet   Pet
	DstUser  common.Address
	DstPetId *big.Int
}

// PetOptions is an auto generated low-level Go binding around an user-defined struct.
type PetOptions struct {
	OriginId          uint8
	Name              string
	Season            uint8
	EggType           uint8
	Rarity            uint8
	Element           uint8
	BonusCount        uint8
	ProfBonus         uint8
	ProfBonusScalar   uint8
	CraftBonus        uint8
	CraftBonusScalar  uint8
	CombatBonus       uint8
	CombatBonusScalar uint8
	Appearance        uint16
	Background        uint8
	Shiny             uint8
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e523e02595067863309d7f7eca1e1e795e71e44f7a6c74c2f793a40e421f01ce64736f6c634300080d0033",
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

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
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

// IMessageBusMetaData contains all meta data concerning the IMessageBus contract.
var IMessageBusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"estimateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_srcAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_dstAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"_messageId\",\"type\":\"bytes32\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_receiver\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_options\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5da6d2c4": "estimateFee(uint256,bytes)",
		"21730efc": "executeMessage(uint256,bytes,address,uint256,uint256,bytes,bytes32)",
		"ac8a4c1b": "sendMessage(bytes32,uint256,bytes,bytes)",
		"1ac3ddeb": "withdrawFee(address)",
	},
}

// IMessageBusABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageBusMetaData.ABI instead.
var IMessageBusABI = IMessageBusMetaData.ABI

// Deprecated: Use IMessageBusMetaData.Sigs instead.
// IMessageBusFuncSigs maps the 4-byte function signature to its string representation.
var IMessageBusFuncSigs = IMessageBusMetaData.Sigs

// IMessageBus is an auto generated Go binding around an Ethereum contract.
type IMessageBus struct {
	IMessageBusCaller     // Read-only binding to the contract
	IMessageBusTransactor // Write-only binding to the contract
	IMessageBusFilterer   // Log filterer for contract events
}

// IMessageBusCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageBusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageBusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageBusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageBusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageBusSession struct {
	Contract     *IMessageBus      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageBusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageBusCallerSession struct {
	Contract *IMessageBusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IMessageBusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageBusTransactorSession struct {
	Contract     *IMessageBusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IMessageBusRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageBusRaw struct {
	Contract *IMessageBus // Generic contract binding to access the raw methods on
}

// IMessageBusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageBusCallerRaw struct {
	Contract *IMessageBusCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageBusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageBusTransactorRaw struct {
	Contract *IMessageBusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageBus creates a new instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBus(address common.Address, backend bind.ContractBackend) (*IMessageBus, error) {
	contract, err := bindIMessageBus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageBus{IMessageBusCaller: IMessageBusCaller{contract: contract}, IMessageBusTransactor: IMessageBusTransactor{contract: contract}, IMessageBusFilterer: IMessageBusFilterer{contract: contract}}, nil
}

// NewIMessageBusCaller creates a new read-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusCaller(address common.Address, caller bind.ContractCaller) (*IMessageBusCaller, error) {
	contract, err := bindIMessageBus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusCaller{contract: contract}, nil
}

// NewIMessageBusTransactor creates a new write-only instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageBusTransactor, error) {
	contract, err := bindIMessageBus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageBusTransactor{contract: contract}, nil
}

// NewIMessageBusFilterer creates a new log filterer instance of IMessageBus, bound to a specific deployed contract.
func NewIMessageBusFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageBusFilterer, error) {
	contract, err := bindIMessageBus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageBusFilterer{contract: contract}, nil
}

// bindIMessageBus binds a generic wrapper to an already deployed contract.
func bindIMessageBus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageBusMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.IMessageBusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.IMessageBusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageBus *IMessageBusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageBus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageBus *IMessageBusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageBus *IMessageBusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageBus.Contract.contract.Transact(opts, method, params...)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IMessageBus *IMessageBusTransactor) EstimateFee(opts *bind.TransactOpts, _dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "estimateFee", _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IMessageBus *IMessageBusSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.EstimateFee(&_IMessageBus.TransactOpts, _dstChainId, _options)
}

// EstimateFee is a paid mutator transaction binding the contract method 0x5da6d2c4.
//
// Solidity: function estimateFee(uint256 _dstChainId, bytes _options) returns(uint256)
func (_IMessageBus *IMessageBusTransactorSession) EstimateFee(_dstChainId *big.Int, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.EstimateFee(&_IMessageBus.TransactOpts, _dstChainId, _options)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x21730efc.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_IMessageBus *IMessageBusTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress []byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "executeMessage", _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x21730efc.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_IMessageBus *IMessageBusSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress []byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0x21730efc.
//
// Solidity: function executeMessage(uint256 _srcChainId, bytes _srcAddress, address _dstAddress, uint256 _gasLimit, uint256 _nonce, bytes _message, bytes32 _messageId) returns()
func (_IMessageBus *IMessageBusTransactorSession) ExecuteMessage(_srcChainId *big.Int, _srcAddress []byte, _dstAddress common.Address, _gasLimit *big.Int, _nonce *big.Int, _message []byte, _messageId [32]byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.ExecuteMessage(&_IMessageBus.TransactOpts, _srcChainId, _srcAddress, _dstAddress, _gasLimit, _nonce, _message, _messageId)
}

// SendMessage is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_IMessageBus *IMessageBusTransactor) SendMessage(opts *bind.TransactOpts, _receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "sendMessage", _receiver, _dstChainId, _message, _options)
}

// SendMessage is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_IMessageBus *IMessageBusSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// SendMessage is a paid mutator transaction binding the contract method 0xac8a4c1b.
//
// Solidity: function sendMessage(bytes32 _receiver, uint256 _dstChainId, bytes _message, bytes _options) payable returns()
func (_IMessageBus *IMessageBusTransactorSession) SendMessage(_receiver [32]byte, _dstChainId *big.Int, _message []byte, _options []byte) (*types.Transaction, error) {
	return _IMessageBus.Contract.SendMessage(&_IMessageBus.TransactOpts, _receiver, _dstChainId, _message, _options)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1ac3ddeb.
//
// Solidity: function withdrawFee(address _account) returns()
func (_IMessageBus *IMessageBusTransactor) WithdrawFee(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _IMessageBus.contract.Transact(opts, "withdrawFee", _account)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1ac3ddeb.
//
// Solidity: function withdrawFee(address _account) returns()
func (_IMessageBus *IMessageBusSession) WithdrawFee(_account common.Address) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0x1ac3ddeb.
//
// Solidity: function withdrawFee(address _account) returns()
func (_IMessageBus *IMessageBusTransactorSession) WithdrawFee(_account common.Address) (*types.Transaction, error) {
	return _IMessageBus.Contract.WithdrawFee(&_IMessageBus.TransactOpts, _account)
}

// IPetCoreUpgradeableMetaData contains all meta data concerning the IPetCoreUpgradeable contract.
var IPetCoreUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"bridgeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getPet\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"originId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"season\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"eggType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"element\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"bonusCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"appearance\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"background\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shiny\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"hungryAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"equippableAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"equippedTo\",\"type\":\"uint256\"}],\"internalType\":\"structPet\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUserPets\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"originId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"season\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"eggType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"element\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"bonusCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"appearance\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"background\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shiny\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"hungryAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"equippableAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"equippedTo\",\"type\":\"uint256\"}],\"internalType\":\"structPet[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"originId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"season\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"eggType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"element\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"bonusCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"appearance\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"background\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shiny\",\"type\":\"uint8\"}],\"internalType\":\"structPetOptions\",\"name\":\"_petOptions\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"hatchPet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"originId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"season\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"eggType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"element\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"bonusCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"appearance\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"background\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shiny\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"hungryAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"equippableAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"equippedTo\",\"type\":\"uint256\"}],\"internalType\":\"structPet\",\"name\":\"_pet\",\"type\":\"tuple\"}],\"name\":\"updatePet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"095ea7b3": "approve(address,uint256)",
		"1b827671": "bridgeMint(uint256,address)",
		"59d55194": "getPet(uint256)",
		"ba8cd532": "getUserPets(address)",
		"2a4138d2": "hatchPet((uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8),address)",
		"6352211e": "ownerOf(uint256)",
		"42842e0e": "safeTransferFrom(address,address,uint256)",
		"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"f5882186": "updatePet((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256))",
	},
}

// IPetCoreUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IPetCoreUpgradeableMetaData.ABI instead.
var IPetCoreUpgradeableABI = IPetCoreUpgradeableMetaData.ABI

// Deprecated: Use IPetCoreUpgradeableMetaData.Sigs instead.
// IPetCoreUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IPetCoreUpgradeableFuncSigs = IPetCoreUpgradeableMetaData.Sigs

// IPetCoreUpgradeable is an auto generated Go binding around an Ethereum contract.
type IPetCoreUpgradeable struct {
	IPetCoreUpgradeableCaller     // Read-only binding to the contract
	IPetCoreUpgradeableTransactor // Write-only binding to the contract
	IPetCoreUpgradeableFilterer   // Log filterer for contract events
}

// IPetCoreUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPetCoreUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPetCoreUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPetCoreUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPetCoreUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPetCoreUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPetCoreUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPetCoreUpgradeableSession struct {
	Contract     *IPetCoreUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IPetCoreUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPetCoreUpgradeableCallerSession struct {
	Contract *IPetCoreUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IPetCoreUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPetCoreUpgradeableTransactorSession struct {
	Contract     *IPetCoreUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IPetCoreUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPetCoreUpgradeableRaw struct {
	Contract *IPetCoreUpgradeable // Generic contract binding to access the raw methods on
}

// IPetCoreUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPetCoreUpgradeableCallerRaw struct {
	Contract *IPetCoreUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IPetCoreUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPetCoreUpgradeableTransactorRaw struct {
	Contract *IPetCoreUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPetCoreUpgradeable creates a new instance of IPetCoreUpgradeable, bound to a specific deployed contract.
func NewIPetCoreUpgradeable(address common.Address, backend bind.ContractBackend) (*IPetCoreUpgradeable, error) {
	contract, err := bindIPetCoreUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPetCoreUpgradeable{IPetCoreUpgradeableCaller: IPetCoreUpgradeableCaller{contract: contract}, IPetCoreUpgradeableTransactor: IPetCoreUpgradeableTransactor{contract: contract}, IPetCoreUpgradeableFilterer: IPetCoreUpgradeableFilterer{contract: contract}}, nil
}

// NewIPetCoreUpgradeableCaller creates a new read-only instance of IPetCoreUpgradeable, bound to a specific deployed contract.
func NewIPetCoreUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IPetCoreUpgradeableCaller, error) {
	contract, err := bindIPetCoreUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPetCoreUpgradeableCaller{contract: contract}, nil
}

// NewIPetCoreUpgradeableTransactor creates a new write-only instance of IPetCoreUpgradeable, bound to a specific deployed contract.
func NewIPetCoreUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IPetCoreUpgradeableTransactor, error) {
	contract, err := bindIPetCoreUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPetCoreUpgradeableTransactor{contract: contract}, nil
}

// NewIPetCoreUpgradeableFilterer creates a new log filterer instance of IPetCoreUpgradeable, bound to a specific deployed contract.
func NewIPetCoreUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IPetCoreUpgradeableFilterer, error) {
	contract, err := bindIPetCoreUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPetCoreUpgradeableFilterer{contract: contract}, nil
}

// bindIPetCoreUpgradeable binds a generic wrapper to an already deployed contract.
func bindIPetCoreUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IPetCoreUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPetCoreUpgradeable *IPetCoreUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPetCoreUpgradeable.Contract.IPetCoreUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPetCoreUpgradeable *IPetCoreUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.IPetCoreUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPetCoreUpgradeable *IPetCoreUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.IPetCoreUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPetCoreUpgradeable *IPetCoreUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPetCoreUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// GetPet is a free data retrieval call binding the contract method 0x59d55194.
//
// Solidity: function getPet(uint256 _id) view returns((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256))
func (_IPetCoreUpgradeable *IPetCoreUpgradeableCaller) GetPet(opts *bind.CallOpts, _id *big.Int) (Pet, error) {
	var out []interface{}
	err := _IPetCoreUpgradeable.contract.Call(opts, &out, "getPet", _id)

	if err != nil {
		return *new(Pet), err
	}

	out0 := *abi.ConvertType(out[0], new(Pet)).(*Pet)

	return out0, err

}

// GetPet is a free data retrieval call binding the contract method 0x59d55194.
//
// Solidity: function getPet(uint256 _id) view returns((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256))
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) GetPet(_id *big.Int) (Pet, error) {
	return _IPetCoreUpgradeable.Contract.GetPet(&_IPetCoreUpgradeable.CallOpts, _id)
}

// GetPet is a free data retrieval call binding the contract method 0x59d55194.
//
// Solidity: function getPet(uint256 _id) view returns((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256))
func (_IPetCoreUpgradeable *IPetCoreUpgradeableCallerSession) GetPet(_id *big.Int) (Pet, error) {
	return _IPetCoreUpgradeable.Contract.GetPet(&_IPetCoreUpgradeable.CallOpts, _id)
}

// GetUserPets is a free data retrieval call binding the contract method 0xba8cd532.
//
// Solidity: function getUserPets(address _address) view returns((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256)[])
func (_IPetCoreUpgradeable *IPetCoreUpgradeableCaller) GetUserPets(opts *bind.CallOpts, _address common.Address) ([]Pet, error) {
	var out []interface{}
	err := _IPetCoreUpgradeable.contract.Call(opts, &out, "getUserPets", _address)

	if err != nil {
		return *new([]Pet), err
	}

	out0 := *abi.ConvertType(out[0], new([]Pet)).(*[]Pet)

	return out0, err

}

// GetUserPets is a free data retrieval call binding the contract method 0xba8cd532.
//
// Solidity: function getUserPets(address _address) view returns((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256)[])
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) GetUserPets(_address common.Address) ([]Pet, error) {
	return _IPetCoreUpgradeable.Contract.GetUserPets(&_IPetCoreUpgradeable.CallOpts, _address)
}

// GetUserPets is a free data retrieval call binding the contract method 0xba8cd532.
//
// Solidity: function getUserPets(address _address) view returns((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256)[])
func (_IPetCoreUpgradeable *IPetCoreUpgradeableCallerSession) GetUserPets(_address common.Address) ([]Pet, error) {
	return _IPetCoreUpgradeable.Contract.GetUserPets(&_IPetCoreUpgradeable.CallOpts, _address)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IPetCoreUpgradeable *IPetCoreUpgradeableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IPetCoreUpgradeable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IPetCoreUpgradeable.Contract.OwnerOf(&_IPetCoreUpgradeable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IPetCoreUpgradeable *IPetCoreUpgradeableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IPetCoreUpgradeable.Contract.OwnerOf(&_IPetCoreUpgradeable.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.Approve(&_IPetCoreUpgradeable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.Approve(&_IPetCoreUpgradeable.TransactOpts, to, tokenId)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x1b827671.
//
// Solidity: function bridgeMint(uint256 _id, address _to) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactor) BridgeMint(opts *bind.TransactOpts, _id *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.contract.Transact(opts, "bridgeMint", _id, _to)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x1b827671.
//
// Solidity: function bridgeMint(uint256 _id, address _to) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) BridgeMint(_id *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.BridgeMint(&_IPetCoreUpgradeable.TransactOpts, _id, _to)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x1b827671.
//
// Solidity: function bridgeMint(uint256 _id, address _to) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorSession) BridgeMint(_id *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.BridgeMint(&_IPetCoreUpgradeable.TransactOpts, _id, _to)
}

// HatchPet is a paid mutator transaction binding the contract method 0x2a4138d2.
//
// Solidity: function hatchPet((uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8) _petOptions, address owner) returns(uint256)
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactor) HatchPet(opts *bind.TransactOpts, _petOptions PetOptions, owner common.Address) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.contract.Transact(opts, "hatchPet", _petOptions, owner)
}

// HatchPet is a paid mutator transaction binding the contract method 0x2a4138d2.
//
// Solidity: function hatchPet((uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8) _petOptions, address owner) returns(uint256)
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) HatchPet(_petOptions PetOptions, owner common.Address) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.HatchPet(&_IPetCoreUpgradeable.TransactOpts, _petOptions, owner)
}

// HatchPet is a paid mutator transaction binding the contract method 0x2a4138d2.
//
// Solidity: function hatchPet((uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8) _petOptions, address owner) returns(uint256)
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorSession) HatchPet(_petOptions PetOptions, owner common.Address) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.HatchPet(&_IPetCoreUpgradeable.TransactOpts, _petOptions, owner)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.SafeTransferFrom(&_IPetCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.SafeTransferFrom(&_IPetCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.SafeTransferFrom0(&_IPetCoreUpgradeable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.SafeTransferFrom0(&_IPetCoreUpgradeable.TransactOpts, from, to, tokenId, _data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.TransferFrom(&_IPetCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.TransferFrom(&_IPetCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// UpdatePet is a paid mutator transaction binding the contract method 0xf5882186.
//
// Solidity: function updatePet((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256) _pet) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactor) UpdatePet(opts *bind.TransactOpts, _pet Pet) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.contract.Transact(opts, "updatePet", _pet)
}

// UpdatePet is a paid mutator transaction binding the contract method 0xf5882186.
//
// Solidity: function updatePet((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256) _pet) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableSession) UpdatePet(_pet Pet) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.UpdatePet(&_IPetCoreUpgradeable.TransactOpts, _pet)
}

// UpdatePet is a paid mutator transaction binding the contract method 0xf5882186.
//
// Solidity: function updatePet((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256) _pet) returns()
func (_IPetCoreUpgradeable *IPetCoreUpgradeableTransactorSession) UpdatePet(_pet Pet) (*types.Transaction, error) {
	return _IPetCoreUpgradeable.Contract.UpdatePet(&_IPetCoreUpgradeable.TransactOpts, _pet)
}

// ISynMessagingReceiverMetaData contains all meta data concerning the ISynMessagingReceiver contract.
var ISynMessagingReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
	},
}

// ISynMessagingReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use ISynMessagingReceiverMetaData.ABI instead.
var ISynMessagingReceiverABI = ISynMessagingReceiverMetaData.ABI

// Deprecated: Use ISynMessagingReceiverMetaData.Sigs instead.
// ISynMessagingReceiverFuncSigs maps the 4-byte function signature to its string representation.
var ISynMessagingReceiverFuncSigs = ISynMessagingReceiverMetaData.Sigs

// ISynMessagingReceiver is an auto generated Go binding around an Ethereum contract.
type ISynMessagingReceiver struct {
	ISynMessagingReceiverCaller     // Read-only binding to the contract
	ISynMessagingReceiverTransactor // Write-only binding to the contract
	ISynMessagingReceiverFilterer   // Log filterer for contract events
}

// ISynMessagingReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISynMessagingReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISynMessagingReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISynMessagingReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISynMessagingReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISynMessagingReceiverSession struct {
	Contract     *ISynMessagingReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ISynMessagingReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISynMessagingReceiverCallerSession struct {
	Contract *ISynMessagingReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// ISynMessagingReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISynMessagingReceiverTransactorSession struct {
	Contract     *ISynMessagingReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// ISynMessagingReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISynMessagingReceiverRaw struct {
	Contract *ISynMessagingReceiver // Generic contract binding to access the raw methods on
}

// ISynMessagingReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISynMessagingReceiverCallerRaw struct {
	Contract *ISynMessagingReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ISynMessagingReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISynMessagingReceiverTransactorRaw struct {
	Contract *ISynMessagingReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISynMessagingReceiver creates a new instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiver(address common.Address, backend bind.ContractBackend) (*ISynMessagingReceiver, error) {
	contract, err := bindISynMessagingReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiver{ISynMessagingReceiverCaller: ISynMessagingReceiverCaller{contract: contract}, ISynMessagingReceiverTransactor: ISynMessagingReceiverTransactor{contract: contract}, ISynMessagingReceiverFilterer: ISynMessagingReceiverFilterer{contract: contract}}, nil
}

// NewISynMessagingReceiverCaller creates a new read-only instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverCaller(address common.Address, caller bind.ContractCaller) (*ISynMessagingReceiverCaller, error) {
	contract, err := bindISynMessagingReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverCaller{contract: contract}, nil
}

// NewISynMessagingReceiverTransactor creates a new write-only instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ISynMessagingReceiverTransactor, error) {
	contract, err := bindISynMessagingReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverTransactor{contract: contract}, nil
}

// NewISynMessagingReceiverFilterer creates a new log filterer instance of ISynMessagingReceiver, bound to a specific deployed contract.
func NewISynMessagingReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ISynMessagingReceiverFilterer, error) {
	contract, err := bindISynMessagingReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISynMessagingReceiverFilterer{contract: contract}, nil
}

// bindISynMessagingReceiver binds a generic wrapper to an already deployed contract.
func bindISynMessagingReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ISynMessagingReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynMessagingReceiver *ISynMessagingReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ISynMessagingReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISynMessagingReceiver *ISynMessagingReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISynMessagingReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ExecuteMessage(&_ISynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_ISynMessagingReceiver *ISynMessagingReceiverTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _ISynMessagingReceiver.Contract.ExecuteMessage(&_ISynMessagingReceiver.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
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

// OwnableUpgradeableMetaData contains all meta data concerning the OwnableUpgradeable contract.
var OwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// PetBridgeUpgradeableMetaData contains all meta data concerning the PetBridgeUpgradeable contract.
var PetBridgeUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"petId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arrivalChainId\",\"type\":\"uint256\"}],\"name\":\"PetArrived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"petId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arrivalChainId\",\"type\":\"uint256\"}],\"name\":\"PetSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"decodeMessage\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"originId\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"season\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"eggType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"element\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"bonusCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"profBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"craftBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonus\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"combatBonusScalar\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"appearance\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"background\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shiny\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"hungryAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"equippableAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"equippedTo\",\"type\":\"uint256\"}],\"internalType\":\"structPet\",\"name\":\"dstPet\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"dstUser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstPetId\",\"type\":\"uint256\"}],\"internalType\":\"structPetBridgeUpgradeable.MessageFormat\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getTrustedRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"trustedRemote\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_pets\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_petId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"}],\"name\":\"sendPet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_msgGasLimit\",\"type\":\"uint256\"}],\"name\":\"setMsgGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"634d45b2": "decodeMessage(bytes)",
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
		"84a12b0f": "getTrustedRemote(uint256)",
		"485cc955": "initialize(address,address)",
		"a1a227fa": "messageBus()",
		"c0e07f28": "msgGasLimit()",
		"8da5cb5b": "owner()",
		"a1227b68": "pets()",
		"715018a6": "renounceOwnership()",
		"bb5e613b": "sendPet(uint256,uint256)",
		"547cad12": "setMessageBus(address)",
		"f9ecc6f5": "setMsgGasLimit(uint256)",
		"bd3583ae": "setTrustedRemote(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611ba6806100206000396000f3fe6080604052600436106100dd5760003560e01c8063a1a227fa1161007f578063bd3583ae11610059578063bd3583ae14610283578063c0e07f28146102a3578063f2fde38b146102b9578063f9ecc6f5146102d957600080fd5b8063a1a227fa14610223578063a606087114610250578063bb5e613b1461027057600080fd5b8063715018a6116100bb578063715018a61461015a57806384a12b0f1461016f5780638da5cb5b146101aa578063a1227b68146101f657600080fd5b8063485cc955146100e2578063547cad1214610104578063634d45b214610124575b600080fd5b3480156100ee57600080fd5b506101026100fd366004611350565b6102f9565b005b34801561011057600080fd5b5061010261011f366004611389565b61045f565b34801561013057600080fd5b5061014461013f366004611494565b61050d565b60405161015191906116e8565b60405180910390f35b34801561016657600080fd5b506101026105d6565b34801561017b57600080fd5b5061019c61018a36600461173a565b60009081526066602052604090205490565b604051908152602001610151565b3480156101b657600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610151565b34801561020257600080fd5b506097546101d19073ffffffffffffffffffffffffffffffffffffffff1681565b34801561022f57600080fd5b506065546101d19073ffffffffffffffffffffffffffffffffffffffff1681565b34801561025c57600080fd5b5061010261026b366004611753565b610649565b61010261027e3660046117ee565b610757565b34801561028f57600080fd5b5061010261029e3660046117ee565b610aaa565b3480156102af57600080fd5b5061019c60985481565b3480156102c557600080fd5b506101026102d4366004611389565b610b60565b3480156102e557600080fd5b506101026102f436600461173a565b610c5c565b600054610100900460ff166103145760005460ff1615610318565b303b155b61038f5760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600054610100900460ff161580156103ce57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011790555b6103d6610cc8565b6065805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556097805492851692909116919091179055801561045a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146104c65760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610386565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b604080516102e081018252600060608083018281526080840183905260a084019190915260c0830182905260e08301829052610100830182905261012083018290526101408301829052610160830182905261018083018290526101a083018290526101c083018290526101e08301829052610200830182905261022083018290526102408301829052610260830182905261028083018290526102a083018290526102c08301829052825260208201819052918101919091526105d082610d4e565b92915050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461063d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610386565b6106476000610e25565b565b60655473ffffffffffffffffffffffffffffffffffffffff1633146106b05760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610386565b600084815260666020526040902054851461070d5760405162461bcd60e51b815260206004820152601a60248201527f496e76616c696420736f757263652073656e64696e67206170700000000000006044820152606401610386565b610750858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250610e9c915050565b5050505050565b6097546040517f59d55194000000000000000000000000000000000000000000000000000000008152600481018490528391839160009173ffffffffffffffffffffffffffffffffffffffff16906359d5519490602401600060405180830381865afa1580156107cb573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526108119190810190611a29565b90508061026001516000146108685760405162461bcd60e51b815260206004820152600f60248201527f70657420697320657175697070656400000000000000000000000000000000006044820152606401610386565b60008281526066602052604081205490610883853385611165565b905060006108d1609854604080517e01000000000000000000000000000000000000000000000000000000000000602082015260228082019390935281518082039093018352604201905290565b6097546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810189905291925073ffffffffffffffffffffffffffffffffffffffff16906323b872dd90606401600060405180830381600087803b15801561094a57600080fd5b505af115801561095e573d6000803e3d6000fd5b50506097546040517f6352211e000000000000000000000000000000000000000000000000000000008152600481018a905230935073ffffffffffffffffffffffffffffffffffffffff9091169150636352211e90602401602060405180830381865afa1580156109d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109f79190611a5e565b73ffffffffffffffffffffffffffffffffffffffff1614610a5a5760405162461bcd60e51b815260206004820152601260248201527f4661696c656420746f206c6f636b2050657400000000000000000000000000006044820152606401610386565b610a66838684846111c6565b857f0158ec556cf37866a866d820878380a6a27533e8e17b9d692f800dd8de5ae84f86604051610a9891815260200190565b60405180910390a25050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610b115760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610386565b60008281526066602090815260409182902083905581518481529081018390527f642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03910160405180910390a15050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610bc75760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610386565b73ffffffffffffffffffffffffffffffffffffffff8116610c505760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610386565b610c5981610e25565b50565b60335473ffffffffffffffffffffffffffffffffffffffff163314610cc35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610386565b609855565b600054610100900460ff16610d455760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610386565b61064733610e25565b604080516102e081018252600060608083018281526080840183905260a084019190915260c0830182905260e08301829052610100830182905261012083018290526101408301829052610160830182905261018083018290526101a083018290526101c083018290526101e08301829052610200830182905261022083018290526102408301829052610260830182905261028083018290526102a083018290526102c0830182905282526020820181905291810191909152600082806020019051810190610e1e9190611a7b565b9392505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b6000610ea783610d4e565b8051602082015160408084015160975491517f6352211e00000000000000000000000000000000000000000000000000000000815260048101829052949550929391929173ffffffffffffffffffffffffffffffffffffffff90911690636352211e90602401602060405180830381865afa925050508015610f64575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252610f6191810190611a5e565b60015b610ff9576097546040517f1b8276710000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff848116602483015290911690631b82767190604401600060405180830381600087803b158015610fdc57600080fd5b505af1158015610ff0573d6000803e3d6000fd5b505050506110ab565b3073ffffffffffffffffffffffffffffffffffffffff8216036110a9576097546040517f42842e0e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff858116602483015260448201859052909116906342842e0e90606401600060405180830381600087803b15801561109057600080fd5b505af11580156110a4573d6000803e3d6000fd5b505050505b505b6097546040517ff588218600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063f588218690611101908690600401611b21565b600060405180830381600087803b15801561111b57600080fd5b505af115801561112f573d6000803e3d6000fd5b50505050807f6bd58bb696941d116b542481b456cfc0a45988f5a1fca53066073374992d2f3746604051610a9891815260200190565b6060600060405180606001604052808481526020018573ffffffffffffffffffffffffffffffffffffffff168152602001868152509050806040516020016111ad91906116e8565b6040516020818303038152906040529150509392505050565b600083815260666020526040902054806112225760405162461bcd60e51b815260206004820152601f60248201527f4e6f2072656d6f7465206170702073657420666f722064737420636861696e006044820152606401610386565b8481146112975760405162461bcd60e51b815260206004820152602660248201527f5265636569766572206973206e6f7420696e20747275737465642072656d6f7460448201527f65206170707300000000000000000000000000000000000000000000000000006064820152608401610386565b6065546040517fac8a4c1b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063ac8a4c1b9034906112f5908990899089908990600401611b34565b6000604051808303818588803b15801561130e57600080fd5b505af1158015611322573d6000803e3d6000fd5b50505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff81168114610c5957600080fd5b6000806040838503121561136357600080fd5b823561136e8161132e565b9150602083013561137e8161132e565b809150509250929050565b60006020828403121561139b57600080fd5b8135610e1e8161132e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610280810167ffffffffffffffff811182821017156113f9576113f96113a6565b60405290565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611446576114466113a6565b604052919050565b600067ffffffffffffffff821115611468576114686113a6565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b6000602082840312156114a657600080fd5b813567ffffffffffffffff8111156114bd57600080fd5b8201601f810184136114ce57600080fd5b80356114e16114dc8261144e565b6113ff565b8181528560208385010111156114f657600080fd5b81602084016020830137600091810160200191909152949350505050565b60005b8381101561152f578181015183820152602001611517565b8381111561153e576000848401525b50505050565b6000815180845261155c816020860160208601611514565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60006102808251845260208301516115ab602086018260ff169052565b5060408301518160408601526115c382860182611544565b91505060608301516115da606086018260ff169052565b5060808301516115ef608086018260ff169052565b5060a083015161160460a086018260ff169052565b5060c083015161161960c086018260ff169052565b5060e083015161162e60e086018260ff169052565b506101008381015160ff90811691860191909152610120808501518216908601526101408085015182169086015261016080850151821690860152610180808501518216908601526101a0808501518216908601526101c08085015161ffff16908601526101e08085015182169086015261020080850151909116908501526102208084015167ffffffffffffffff9081169186019190915261024080850151909116908501526102609283015192909301919091525090565b602081526000825160606020840152611704608084018261158e565b905073ffffffffffffffffffffffffffffffffffffffff6020850151166040840152604084015160608401528091505092915050565b60006020828403121561174c57600080fd5b5035919050565b60008060008060006080868803121561176b57600080fd5b8535945060208601359350604086013567ffffffffffffffff8082111561179157600080fd5b818801915088601f8301126117a557600080fd5b8135818111156117b457600080fd5b8960208285010111156117c657600080fd5b60208301955080945050505060608601356117e08161132e565b809150509295509295909350565b6000806040838503121561180157600080fd5b50508035926020909101359150565b805160ff8116811461182157600080fd5b919050565b600082601f83011261183757600080fd5b81516118456114dc8261144e565b81815284602083860101111561185a57600080fd5b61186b826020830160208701611514565b949350505050565b805161ffff8116811461182157600080fd5b805167ffffffffffffffff8116811461182157600080fd5b600061028082840312156118b057600080fd5b6118b86113d5565b9050815181526118ca60208301611810565b6020820152604082015167ffffffffffffffff8111156118e957600080fd5b6118f584828501611826565b60408301525061190760608301611810565b606082015261191860808301611810565b608082015261192960a08301611810565b60a082015261193a60c08301611810565b60c082015261194b60e08301611810565b60e082015261010061195e818401611810565b90820152610120611970838201611810565b90820152610140611982838201611810565b90820152610160611994838201611810565b908201526101806119a6838201611810565b908201526101a06119b8838201611810565b908201526101c06119ca838201611873565b908201526101e06119dc838201611810565b908201526102006119ee838201611810565b90820152610220611a00838201611885565b90820152610240611a12838201611885565b818301525061026080830151818301525092915050565b600060208284031215611a3b57600080fd5b815167ffffffffffffffff811115611a5257600080fd5b61186b8482850161189d565b600060208284031215611a7057600080fd5b8151610e1e8161132e565b600060208284031215611a8d57600080fd5b815167ffffffffffffffff80821115611aa557600080fd5b9083019060608286031215611ab957600080fd5b604051606081018181108382111715611ad457611ad46113a6565b604052825182811115611ae657600080fd5b611af28782860161189d565b82525060208301519150611b058261132e565b8160208201526040830151604082015280935050505092915050565b602081526000610e1e602083018461158e565b848152836020820152608060408201526000611b536080830185611544565b8281036060840152611b658185611544565b97965050505050505056fea26469706673582212207b7b9dbea5688d4c71a262b1fac7f585af4bdba4c0af18bd263923d8080dd2e264736f6c634300080d0033",
}

// PetBridgeUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use PetBridgeUpgradeableMetaData.ABI instead.
var PetBridgeUpgradeableABI = PetBridgeUpgradeableMetaData.ABI

// Deprecated: Use PetBridgeUpgradeableMetaData.Sigs instead.
// PetBridgeUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var PetBridgeUpgradeableFuncSigs = PetBridgeUpgradeableMetaData.Sigs

// PetBridgeUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PetBridgeUpgradeableMetaData.Bin instead.
var PetBridgeUpgradeableBin = PetBridgeUpgradeableMetaData.Bin

// DeployPetBridgeUpgradeable deploys a new Ethereum contract, binding an instance of PetBridgeUpgradeable to it.
func DeployPetBridgeUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PetBridgeUpgradeable, error) {
	parsed, err := PetBridgeUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PetBridgeUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PetBridgeUpgradeable{PetBridgeUpgradeableCaller: PetBridgeUpgradeableCaller{contract: contract}, PetBridgeUpgradeableTransactor: PetBridgeUpgradeableTransactor{contract: contract}, PetBridgeUpgradeableFilterer: PetBridgeUpgradeableFilterer{contract: contract}}, nil
}

// PetBridgeUpgradeable is an auto generated Go binding around an Ethereum contract.
type PetBridgeUpgradeable struct {
	PetBridgeUpgradeableCaller     // Read-only binding to the contract
	PetBridgeUpgradeableTransactor // Write-only binding to the contract
	PetBridgeUpgradeableFilterer   // Log filterer for contract events
}

// PetBridgeUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PetBridgeUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PetBridgeUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PetBridgeUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PetBridgeUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PetBridgeUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PetBridgeUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PetBridgeUpgradeableSession struct {
	Contract     *PetBridgeUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PetBridgeUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PetBridgeUpgradeableCallerSession struct {
	Contract *PetBridgeUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// PetBridgeUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PetBridgeUpgradeableTransactorSession struct {
	Contract     *PetBridgeUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// PetBridgeUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PetBridgeUpgradeableRaw struct {
	Contract *PetBridgeUpgradeable // Generic contract binding to access the raw methods on
}

// PetBridgeUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PetBridgeUpgradeableCallerRaw struct {
	Contract *PetBridgeUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// PetBridgeUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PetBridgeUpgradeableTransactorRaw struct {
	Contract *PetBridgeUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPetBridgeUpgradeable creates a new instance of PetBridgeUpgradeable, bound to a specific deployed contract.
func NewPetBridgeUpgradeable(address common.Address, backend bind.ContractBackend) (*PetBridgeUpgradeable, error) {
	contract, err := bindPetBridgeUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeable{PetBridgeUpgradeableCaller: PetBridgeUpgradeableCaller{contract: contract}, PetBridgeUpgradeableTransactor: PetBridgeUpgradeableTransactor{contract: contract}, PetBridgeUpgradeableFilterer: PetBridgeUpgradeableFilterer{contract: contract}}, nil
}

// NewPetBridgeUpgradeableCaller creates a new read-only instance of PetBridgeUpgradeable, bound to a specific deployed contract.
func NewPetBridgeUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*PetBridgeUpgradeableCaller, error) {
	contract, err := bindPetBridgeUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeableCaller{contract: contract}, nil
}

// NewPetBridgeUpgradeableTransactor creates a new write-only instance of PetBridgeUpgradeable, bound to a specific deployed contract.
func NewPetBridgeUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*PetBridgeUpgradeableTransactor, error) {
	contract, err := bindPetBridgeUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeableTransactor{contract: contract}, nil
}

// NewPetBridgeUpgradeableFilterer creates a new log filterer instance of PetBridgeUpgradeable, bound to a specific deployed contract.
func NewPetBridgeUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*PetBridgeUpgradeableFilterer, error) {
	contract, err := bindPetBridgeUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeableFilterer{contract: contract}, nil
}

// bindPetBridgeUpgradeable binds a generic wrapper to an already deployed contract.
func bindPetBridgeUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PetBridgeUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PetBridgeUpgradeable *PetBridgeUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PetBridgeUpgradeable.Contract.PetBridgeUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PetBridgeUpgradeable *PetBridgeUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.PetBridgeUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PetBridgeUpgradeable *PetBridgeUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.PetBridgeUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PetBridgeUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns(((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256),address,uint256))
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCaller) DecodeMessage(opts *bind.CallOpts, _message []byte) (PetBridgeUpgradeableMessageFormat, error) {
	var out []interface{}
	err := _PetBridgeUpgradeable.contract.Call(opts, &out, "decodeMessage", _message)

	if err != nil {
		return *new(PetBridgeUpgradeableMessageFormat), err
	}

	out0 := *abi.ConvertType(out[0], new(PetBridgeUpgradeableMessageFormat)).(*PetBridgeUpgradeableMessageFormat)

	return out0, err

}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns(((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256),address,uint256))
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) DecodeMessage(_message []byte) (PetBridgeUpgradeableMessageFormat, error) {
	return _PetBridgeUpgradeable.Contract.DecodeMessage(&_PetBridgeUpgradeable.CallOpts, _message)
}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns(((uint256,uint8,string,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint8,uint16,uint8,uint8,uint64,uint64,uint256),address,uint256))
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCallerSession) DecodeMessage(_message []byte) (PetBridgeUpgradeableMessageFormat, error) {
	return _PetBridgeUpgradeable.Contract.DecodeMessage(&_PetBridgeUpgradeable.CallOpts, _message)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCaller) GetTrustedRemote(opts *bind.CallOpts, _chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PetBridgeUpgradeable.contract.Call(opts, &out, "getTrustedRemote", _chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _PetBridgeUpgradeable.Contract.GetTrustedRemote(&_PetBridgeUpgradeable.CallOpts, _chainId)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCallerSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _PetBridgeUpgradeable.Contract.GetTrustedRemote(&_PetBridgeUpgradeable.CallOpts, _chainId)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PetBridgeUpgradeable.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) MessageBus() (common.Address, error) {
	return _PetBridgeUpgradeable.Contract.MessageBus(&_PetBridgeUpgradeable.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCallerSession) MessageBus() (common.Address, error) {
	return _PetBridgeUpgradeable.Contract.MessageBus(&_PetBridgeUpgradeable.CallOpts)
}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCaller) MsgGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PetBridgeUpgradeable.contract.Call(opts, &out, "msgGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) MsgGasLimit() (*big.Int, error) {
	return _PetBridgeUpgradeable.Contract.MsgGasLimit(&_PetBridgeUpgradeable.CallOpts)
}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCallerSession) MsgGasLimit() (*big.Int, error) {
	return _PetBridgeUpgradeable.Contract.MsgGasLimit(&_PetBridgeUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PetBridgeUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) Owner() (common.Address, error) {
	return _PetBridgeUpgradeable.Contract.Owner(&_PetBridgeUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCallerSession) Owner() (common.Address, error) {
	return _PetBridgeUpgradeable.Contract.Owner(&_PetBridgeUpgradeable.CallOpts)
}

// Pets is a free data retrieval call binding the contract method 0xa1227b68.
//
// Solidity: function pets() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCaller) Pets(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PetBridgeUpgradeable.contract.Call(opts, &out, "pets")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pets is a free data retrieval call binding the contract method 0xa1227b68.
//
// Solidity: function pets() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) Pets() (common.Address, error) {
	return _PetBridgeUpgradeable.Contract.Pets(&_PetBridgeUpgradeable.CallOpts)
}

// Pets is a free data retrieval call binding the contract method 0xa1227b68.
//
// Solidity: function pets() view returns(address)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableCallerSession) Pets() (common.Address, error) {
	return _PetBridgeUpgradeable.Contract.Pets(&_PetBridgeUpgradeable.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.ExecuteMessage(&_PetBridgeUpgradeable.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.ExecuteMessage(&_PetBridgeUpgradeable.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messageBus, address _pets) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) Initialize(opts *bind.TransactOpts, _messageBus common.Address, _pets common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "initialize", _messageBus, _pets)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messageBus, address _pets) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) Initialize(_messageBus common.Address, _pets common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.Initialize(&_PetBridgeUpgradeable.TransactOpts, _messageBus, _pets)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _messageBus, address _pets) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) Initialize(_messageBus common.Address, _pets common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.Initialize(&_PetBridgeUpgradeable.TransactOpts, _messageBus, _pets)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.RenounceOwnership(&_PetBridgeUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.RenounceOwnership(&_PetBridgeUpgradeable.TransactOpts)
}

// SendPet is a paid mutator transaction binding the contract method 0xbb5e613b.
//
// Solidity: function sendPet(uint256 _petId, uint256 _dstChainId) payable returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) SendPet(opts *bind.TransactOpts, _petId *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "sendPet", _petId, _dstChainId)
}

// SendPet is a paid mutator transaction binding the contract method 0xbb5e613b.
//
// Solidity: function sendPet(uint256 _petId, uint256 _dstChainId) payable returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) SendPet(_petId *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SendPet(&_PetBridgeUpgradeable.TransactOpts, _petId, _dstChainId)
}

// SendPet is a paid mutator transaction binding the contract method 0xbb5e613b.
//
// Solidity: function sendPet(uint256 _petId, uint256 _dstChainId) payable returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) SendPet(_petId *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SendPet(&_PetBridgeUpgradeable.TransactOpts, _petId, _dstChainId)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SetMessageBus(&_PetBridgeUpgradeable.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SetMessageBus(&_PetBridgeUpgradeable.TransactOpts, _messageBus)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) SetMsgGasLimit(opts *bind.TransactOpts, _msgGasLimit *big.Int) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "setMsgGasLimit", _msgGasLimit)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) SetMsgGasLimit(_msgGasLimit *big.Int) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SetMsgGasLimit(&_PetBridgeUpgradeable.TransactOpts, _msgGasLimit)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) SetMsgGasLimit(_msgGasLimit *big.Int) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SetMsgGasLimit(&_PetBridgeUpgradeable.TransactOpts, _msgGasLimit)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "setTrustedRemote", _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SetTrustedRemote(&_PetBridgeUpgradeable.TransactOpts, _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.SetTrustedRemote(&_PetBridgeUpgradeable.TransactOpts, _srcChainId, _srcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.TransferOwnership(&_PetBridgeUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PetBridgeUpgradeable *PetBridgeUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PetBridgeUpgradeable.Contract.TransferOwnership(&_PetBridgeUpgradeable.TransactOpts, newOwner)
}

// PetBridgeUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeableOwnershipTransferredIterator struct {
	Event *PetBridgeUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PetBridgeUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PetBridgeUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PetBridgeUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PetBridgeUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PetBridgeUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PetBridgeUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PetBridgeUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PetBridgeUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeableOwnershipTransferredIterator{contract: _PetBridgeUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PetBridgeUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PetBridgeUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PetBridgeUpgradeableOwnershipTransferred)
				if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*PetBridgeUpgradeableOwnershipTransferred, error) {
	event := new(PetBridgeUpgradeableOwnershipTransferred)
	if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PetBridgeUpgradeablePetArrivedIterator is returned from FilterPetArrived and is used to iterate over the raw logs and unpacked data for PetArrived events raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeablePetArrivedIterator struct {
	Event *PetBridgeUpgradeablePetArrived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PetBridgeUpgradeablePetArrivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PetBridgeUpgradeablePetArrived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PetBridgeUpgradeablePetArrived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PetBridgeUpgradeablePetArrivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PetBridgeUpgradeablePetArrivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PetBridgeUpgradeablePetArrived represents a PetArrived event raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeablePetArrived struct {
	PetId          *big.Int
	ArrivalChainId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPetArrived is a free log retrieval operation binding the contract event 0x6bd58bb696941d116b542481b456cfc0a45988f5a1fca53066073374992d2f37.
//
// Solidity: event PetArrived(uint256 indexed petId, uint256 arrivalChainId)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) FilterPetArrived(opts *bind.FilterOpts, petId []*big.Int) (*PetBridgeUpgradeablePetArrivedIterator, error) {

	var petIdRule []interface{}
	for _, petIdItem := range petId {
		petIdRule = append(petIdRule, petIdItem)
	}

	logs, sub, err := _PetBridgeUpgradeable.contract.FilterLogs(opts, "PetArrived", petIdRule)
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeablePetArrivedIterator{contract: _PetBridgeUpgradeable.contract, event: "PetArrived", logs: logs, sub: sub}, nil
}

// WatchPetArrived is a free log subscription operation binding the contract event 0x6bd58bb696941d116b542481b456cfc0a45988f5a1fca53066073374992d2f37.
//
// Solidity: event PetArrived(uint256 indexed petId, uint256 arrivalChainId)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) WatchPetArrived(opts *bind.WatchOpts, sink chan<- *PetBridgeUpgradeablePetArrived, petId []*big.Int) (event.Subscription, error) {

	var petIdRule []interface{}
	for _, petIdItem := range petId {
		petIdRule = append(petIdRule, petIdItem)
	}

	logs, sub, err := _PetBridgeUpgradeable.contract.WatchLogs(opts, "PetArrived", petIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PetBridgeUpgradeablePetArrived)
				if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "PetArrived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePetArrived is a log parse operation binding the contract event 0x6bd58bb696941d116b542481b456cfc0a45988f5a1fca53066073374992d2f37.
//
// Solidity: event PetArrived(uint256 indexed petId, uint256 arrivalChainId)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) ParsePetArrived(log types.Log) (*PetBridgeUpgradeablePetArrived, error) {
	event := new(PetBridgeUpgradeablePetArrived)
	if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "PetArrived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PetBridgeUpgradeablePetSentIterator is returned from FilterPetSent and is used to iterate over the raw logs and unpacked data for PetSent events raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeablePetSentIterator struct {
	Event *PetBridgeUpgradeablePetSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PetBridgeUpgradeablePetSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PetBridgeUpgradeablePetSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PetBridgeUpgradeablePetSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PetBridgeUpgradeablePetSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PetBridgeUpgradeablePetSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PetBridgeUpgradeablePetSent represents a PetSent event raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeablePetSent struct {
	PetId          *big.Int
	ArrivalChainId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPetSent is a free log retrieval operation binding the contract event 0x0158ec556cf37866a866d820878380a6a27533e8e17b9d692f800dd8de5ae84f.
//
// Solidity: event PetSent(uint256 indexed petId, uint256 arrivalChainId)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) FilterPetSent(opts *bind.FilterOpts, petId []*big.Int) (*PetBridgeUpgradeablePetSentIterator, error) {

	var petIdRule []interface{}
	for _, petIdItem := range petId {
		petIdRule = append(petIdRule, petIdItem)
	}

	logs, sub, err := _PetBridgeUpgradeable.contract.FilterLogs(opts, "PetSent", petIdRule)
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeablePetSentIterator{contract: _PetBridgeUpgradeable.contract, event: "PetSent", logs: logs, sub: sub}, nil
}

// WatchPetSent is a free log subscription operation binding the contract event 0x0158ec556cf37866a866d820878380a6a27533e8e17b9d692f800dd8de5ae84f.
//
// Solidity: event PetSent(uint256 indexed petId, uint256 arrivalChainId)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) WatchPetSent(opts *bind.WatchOpts, sink chan<- *PetBridgeUpgradeablePetSent, petId []*big.Int) (event.Subscription, error) {

	var petIdRule []interface{}
	for _, petIdItem := range petId {
		petIdRule = append(petIdRule, petIdItem)
	}

	logs, sub, err := _PetBridgeUpgradeable.contract.WatchLogs(opts, "PetSent", petIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PetBridgeUpgradeablePetSent)
				if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "PetSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePetSent is a log parse operation binding the contract event 0x0158ec556cf37866a866d820878380a6a27533e8e17b9d692f800dd8de5ae84f.
//
// Solidity: event PetSent(uint256 indexed petId, uint256 arrivalChainId)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) ParsePetSent(log types.Log) (*PetBridgeUpgradeablePetSent, error) {
	event := new(PetBridgeUpgradeablePetSent)
	if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "PetSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PetBridgeUpgradeableSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeableSetTrustedRemoteIterator struct {
	Event *PetBridgeUpgradeableSetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PetBridgeUpgradeableSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PetBridgeUpgradeableSetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PetBridgeUpgradeableSetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PetBridgeUpgradeableSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PetBridgeUpgradeableSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PetBridgeUpgradeableSetTrustedRemote represents a SetTrustedRemote event raised by the PetBridgeUpgradeable contract.
type PetBridgeUpgradeableSetTrustedRemote struct {
	SrcChainId *big.Int
	SrcAddress [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*PetBridgeUpgradeableSetTrustedRemoteIterator, error) {

	logs, sub, err := _PetBridgeUpgradeable.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &PetBridgeUpgradeableSetTrustedRemoteIterator{contract: _PetBridgeUpgradeable.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *PetBridgeUpgradeableSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _PetBridgeUpgradeable.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PetBridgeUpgradeableSetTrustedRemote)
				if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemote is a log parse operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_PetBridgeUpgradeable *PetBridgeUpgradeableFilterer) ParseSetTrustedRemote(log types.Log) (*PetBridgeUpgradeableSetTrustedRemote, error) {
	event := new(PetBridgeUpgradeableSetTrustedRemote)
	if err := _PetBridgeUpgradeable.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynMessagingReceiverUpgradeableMetaData contains all meta data concerning the SynMessagingReceiverUpgradeable contract.
var SynMessagingReceiverUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getTrustedRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"trustedRemote\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
		"84a12b0f": "getTrustedRemote(uint256)",
		"a1a227fa": "messageBus()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"547cad12": "setMessageBus(address)",
		"bd3583ae": "setTrustedRemote(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
	},
}

// SynMessagingReceiverUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use SynMessagingReceiverUpgradeableMetaData.ABI instead.
var SynMessagingReceiverUpgradeableABI = SynMessagingReceiverUpgradeableMetaData.ABI

// Deprecated: Use SynMessagingReceiverUpgradeableMetaData.Sigs instead.
// SynMessagingReceiverUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var SynMessagingReceiverUpgradeableFuncSigs = SynMessagingReceiverUpgradeableMetaData.Sigs

// SynMessagingReceiverUpgradeable is an auto generated Go binding around an Ethereum contract.
type SynMessagingReceiverUpgradeable struct {
	SynMessagingReceiverUpgradeableCaller     // Read-only binding to the contract
	SynMessagingReceiverUpgradeableTransactor // Write-only binding to the contract
	SynMessagingReceiverUpgradeableFilterer   // Log filterer for contract events
}

// SynMessagingReceiverUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynMessagingReceiverUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynMessagingReceiverUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynMessagingReceiverUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynMessagingReceiverUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynMessagingReceiverUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynMessagingReceiverUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynMessagingReceiverUpgradeableSession struct {
	Contract     *SynMessagingReceiverUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                    // Call options to use throughout this session
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// SynMessagingReceiverUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynMessagingReceiverUpgradeableCallerSession struct {
	Contract *SynMessagingReceiverUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                          // Call options to use throughout this session
}

// SynMessagingReceiverUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynMessagingReceiverUpgradeableTransactorSession struct {
	Contract     *SynMessagingReceiverUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                          // Transaction auth options to use throughout this session
}

// SynMessagingReceiverUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynMessagingReceiverUpgradeableRaw struct {
	Contract *SynMessagingReceiverUpgradeable // Generic contract binding to access the raw methods on
}

// SynMessagingReceiverUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynMessagingReceiverUpgradeableCallerRaw struct {
	Contract *SynMessagingReceiverUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// SynMessagingReceiverUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynMessagingReceiverUpgradeableTransactorRaw struct {
	Contract *SynMessagingReceiverUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynMessagingReceiverUpgradeable creates a new instance of SynMessagingReceiverUpgradeable, bound to a specific deployed contract.
func NewSynMessagingReceiverUpgradeable(address common.Address, backend bind.ContractBackend) (*SynMessagingReceiverUpgradeable, error) {
	contract, err := bindSynMessagingReceiverUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverUpgradeable{SynMessagingReceiverUpgradeableCaller: SynMessagingReceiverUpgradeableCaller{contract: contract}, SynMessagingReceiverUpgradeableTransactor: SynMessagingReceiverUpgradeableTransactor{contract: contract}, SynMessagingReceiverUpgradeableFilterer: SynMessagingReceiverUpgradeableFilterer{contract: contract}}, nil
}

// NewSynMessagingReceiverUpgradeableCaller creates a new read-only instance of SynMessagingReceiverUpgradeable, bound to a specific deployed contract.
func NewSynMessagingReceiverUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*SynMessagingReceiverUpgradeableCaller, error) {
	contract, err := bindSynMessagingReceiverUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverUpgradeableCaller{contract: contract}, nil
}

// NewSynMessagingReceiverUpgradeableTransactor creates a new write-only instance of SynMessagingReceiverUpgradeable, bound to a specific deployed contract.
func NewSynMessagingReceiverUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*SynMessagingReceiverUpgradeableTransactor, error) {
	contract, err := bindSynMessagingReceiverUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverUpgradeableTransactor{contract: contract}, nil
}

// NewSynMessagingReceiverUpgradeableFilterer creates a new log filterer instance of SynMessagingReceiverUpgradeable, bound to a specific deployed contract.
func NewSynMessagingReceiverUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*SynMessagingReceiverUpgradeableFilterer, error) {
	contract, err := bindSynMessagingReceiverUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverUpgradeableFilterer{contract: contract}, nil
}

// bindSynMessagingReceiverUpgradeable binds a generic wrapper to an already deployed contract.
func bindSynMessagingReceiverUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SynMessagingReceiverUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynMessagingReceiverUpgradeable.Contract.SynMessagingReceiverUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.SynMessagingReceiverUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.SynMessagingReceiverUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynMessagingReceiverUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableCaller) GetTrustedRemote(opts *bind.CallOpts, _chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _SynMessagingReceiverUpgradeable.contract.Call(opts, &out, "getTrustedRemote", _chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _SynMessagingReceiverUpgradeable.Contract.GetTrustedRemote(&_SynMessagingReceiverUpgradeable.CallOpts, _chainId)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableCallerSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _SynMessagingReceiverUpgradeable.Contract.GetTrustedRemote(&_SynMessagingReceiverUpgradeable.CallOpts, _chainId)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynMessagingReceiverUpgradeable.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) MessageBus() (common.Address, error) {
	return _SynMessagingReceiverUpgradeable.Contract.MessageBus(&_SynMessagingReceiverUpgradeable.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableCallerSession) MessageBus() (common.Address, error) {
	return _SynMessagingReceiverUpgradeable.Contract.MessageBus(&_SynMessagingReceiverUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SynMessagingReceiverUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) Owner() (common.Address, error) {
	return _SynMessagingReceiverUpgradeable.Contract.Owner(&_SynMessagingReceiverUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableCallerSession) Owner() (common.Address, error) {
	return _SynMessagingReceiverUpgradeable.Contract.Owner(&_SynMessagingReceiverUpgradeable.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.ExecuteMessage(&_SynMessagingReceiverUpgradeable.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.ExecuteMessage(&_SynMessagingReceiverUpgradeable.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.RenounceOwnership(&_SynMessagingReceiverUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.RenounceOwnership(&_SynMessagingReceiverUpgradeable.TransactOpts)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.SetMessageBus(&_SynMessagingReceiverUpgradeable.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.SetMessageBus(&_SynMessagingReceiverUpgradeable.TransactOpts, _messageBus)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.contract.Transact(opts, "setTrustedRemote", _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.SetTrustedRemote(&_SynMessagingReceiverUpgradeable.TransactOpts, _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactorSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.SetTrustedRemote(&_SynMessagingReceiverUpgradeable.TransactOpts, _srcChainId, _srcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.TransferOwnership(&_SynMessagingReceiverUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SynMessagingReceiverUpgradeable.Contract.TransferOwnership(&_SynMessagingReceiverUpgradeable.TransactOpts, newOwner)
}

// SynMessagingReceiverUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SynMessagingReceiverUpgradeable contract.
type SynMessagingReceiverUpgradeableOwnershipTransferredIterator struct {
	Event *SynMessagingReceiverUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynMessagingReceiverUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynMessagingReceiverUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynMessagingReceiverUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynMessagingReceiverUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynMessagingReceiverUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynMessagingReceiverUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the SynMessagingReceiverUpgradeable contract.
type SynMessagingReceiverUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SynMessagingReceiverUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynMessagingReceiverUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverUpgradeableOwnershipTransferredIterator{contract: _SynMessagingReceiverUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SynMessagingReceiverUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SynMessagingReceiverUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynMessagingReceiverUpgradeableOwnershipTransferred)
				if err := _SynMessagingReceiverUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*SynMessagingReceiverUpgradeableOwnershipTransferred, error) {
	event := new(SynMessagingReceiverUpgradeableOwnershipTransferred)
	if err := _SynMessagingReceiverUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SynMessagingReceiverUpgradeableSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the SynMessagingReceiverUpgradeable contract.
type SynMessagingReceiverUpgradeableSetTrustedRemoteIterator struct {
	Event *SynMessagingReceiverUpgradeableSetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SynMessagingReceiverUpgradeableSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SynMessagingReceiverUpgradeableSetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SynMessagingReceiverUpgradeableSetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SynMessagingReceiverUpgradeableSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SynMessagingReceiverUpgradeableSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SynMessagingReceiverUpgradeableSetTrustedRemote represents a SetTrustedRemote event raised by the SynMessagingReceiverUpgradeable contract.
type SynMessagingReceiverUpgradeableSetTrustedRemote struct {
	SrcChainId *big.Int
	SrcAddress [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*SynMessagingReceiverUpgradeableSetTrustedRemoteIterator, error) {

	logs, sub, err := _SynMessagingReceiverUpgradeable.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &SynMessagingReceiverUpgradeableSetTrustedRemoteIterator{contract: _SynMessagingReceiverUpgradeable.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *SynMessagingReceiverUpgradeableSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _SynMessagingReceiverUpgradeable.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SynMessagingReceiverUpgradeableSetTrustedRemote)
				if err := _SynMessagingReceiverUpgradeable.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetTrustedRemote is a log parse operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_SynMessagingReceiverUpgradeable *SynMessagingReceiverUpgradeableFilterer) ParseSetTrustedRemote(log types.Log) (*SynMessagingReceiverUpgradeableSetTrustedRemote, error) {
	event := new(SynMessagingReceiverUpgradeableSetTrustedRemote)
	if err := _SynMessagingReceiverUpgradeable.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
