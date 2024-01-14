// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dfkhero

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

// Hero is an auto generated low-level Go binding around an user-defined struct.
type Hero struct {
	Id                  *big.Int
	SummoningInfo       SummoningInfo
	Info                HeroInfo
	State               HeroState
	Stats               HeroStats
	PrimaryStatGrowth   HeroStatGrowth
	SecondaryStatGrowth HeroStatGrowth
	Professions         HeroProfessions
}

// HeroBridgeUpgradeableMessageFormat is an auto generated low-level Go binding around an user-defined struct.
type HeroBridgeUpgradeableMessageFormat struct {
	DstHero   Hero
	DstUser   common.Address
	DstHeroId *big.Int
}

// HeroCrystal is an auto generated low-level Go binding around an user-defined struct.
type HeroCrystal struct {
	Owner            common.Address
	SummonerId       *big.Int
	AssistantId      *big.Int
	Generation       uint16
	CreatedBlock     *big.Int
	HeroId           *big.Int
	SummonerTears    uint8
	AssistantTears   uint8
	EnhancementStone common.Address
	MaxSummons       uint32
	FirstName        uint32
	LastName         uint32
	ShinyStyle       uint8
}

// HeroInfo is an auto generated low-level Go binding around an user-defined struct.
type HeroInfo struct {
	StatGenes   *big.Int
	VisualGenes *big.Int
	Rarity      uint8
	Shiny       bool
	Generation  uint16
	FirstName   uint32
	LastName    uint32
	ShinyStyle  uint8
	Class       uint8
	SubClass    uint8
}

// HeroProfessions is an auto generated low-level Go binding around an user-defined struct.
type HeroProfessions struct {
	Mining    uint16
	Gardening uint16
	Foraging  uint16
	Fishing   uint16
}

// HeroStatGrowth is an auto generated low-level Go binding around an user-defined struct.
type HeroStatGrowth struct {
	Strength     uint16
	Agility      uint16
	Intelligence uint16
	Wisdom       uint16
	Luck         uint16
	Vitality     uint16
	Endurance    uint16
	Dexterity    uint16
	HpSm         uint16
	HpRg         uint16
	HpLg         uint16
	MpSm         uint16
	MpRg         uint16
	MpLg         uint16
}

// HeroState is an auto generated low-level Go binding around an user-defined struct.
type HeroState struct {
	StaminaFullAt *big.Int
	HpFullAt      *big.Int
	MpFullAt      *big.Int
	Level         uint16
	Xp            uint64
	CurrentQuest  common.Address
	Sp            uint8
	Status        uint8
}

// HeroStats is an auto generated low-level Go binding around an user-defined struct.
type HeroStats struct {
	Strength     uint16
	Agility      uint16
	Intelligence uint16
	Wisdom       uint16
	Luck         uint16
	Vitality     uint16
	Endurance    uint16
	Dexterity    uint16
	Hp           uint16
	Mp           uint16
	Stamina      uint16
}

// SummoningInfo is an auto generated low-level Go binding around an user-defined struct.
type SummoningInfo struct {
	SummonedTime   *big.Int
	NextSummonTime *big.Int
	SummonerId     *big.Int
	AssistantId    *big.Int
	Summons        uint32
	MaxSummons     uint32
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122091fbd8e0f5c6fd0b2077ea63bad344da23c029f00b9345777ec49da0d994805164736f6c634300080d0033",
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

// HeroBridgeUpgradeableMetaData contains all meta data concerning the HeroBridgeUpgradeable contract.
var HeroBridgeUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arrivalChainId\",\"type\":\"uint256\"}],\"name\":\"HeroArrived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"arrivalChainId\",\"type\":\"uint256\"}],\"name\":\"HeroSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"SetTrustedRemote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"assistingAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"decodeMessage\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"summonedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSummonTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"summons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"}],\"internalType\":\"structSummoningInfo\",\"name\":\"summoningInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumRarity\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"shiny\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"class\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"subClass\",\"type\":\"uint8\"}],\"internalType\":\"structHeroInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"staminaFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"xp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"currentQuest\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sp\",\"type\":\"uint8\"},{\"internalType\":\"enumHeroStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structHeroState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stamina\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStats\",\"name\":\"stats\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStatGrowth\",\"name\":\"primaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStatGrowth\",\"name\":\"secondaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"mining\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gardening\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"foraging\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fishing\",\"type\":\"uint16\"}],\"internalType\":\"structHeroProfessions\",\"name\":\"professions\",\"type\":\"tuple\"}],\"internalType\":\"structHero\",\"name\":\"dstHero\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"dstUser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dstHeroId\",\"type\":\"uint256\"}],\"internalType\":\"structHeroBridgeUpgradeable.MessageFormat\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"executeMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getTrustedRemote\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"trustedRemote\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"heroes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_heroes\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_assistingAuction\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBus\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"msgGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dstChainId\",\"type\":\"uint256\"}],\"name\":\"sendHero\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_assistingAuction\",\"type\":\"address\"}],\"name\":\"setAssistingAuctionAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageBus\",\"type\":\"address\"}],\"name\":\"setMessageBus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_msgGasLimit\",\"type\":\"uint256\"}],\"name\":\"setMsgGasLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_srcChainId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_srcAddress\",\"type\":\"bytes32\"}],\"name\":\"setTrustedRemote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"7246a948": "assistingAuction()",
		"634d45b2": "decodeMessage(bytes)",
		"a6060871": "executeMessage(bytes32,uint256,bytes,address)",
		"84a12b0f": "getTrustedRemote(uint256)",
		"230bb9f6": "heroes()",
		"c0c53b8b": "initialize(address,address,address)",
		"a1a227fa": "messageBus()",
		"c0e07f28": "msgGasLimit()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"1efedbe5": "sendHero(uint256,uint256)",
		"5c9c7c73": "setAssistingAuctionAddress(address)",
		"547cad12": "setMessageBus(address)",
		"f9ecc6f5": "setMsgGasLimit(uint256)",
		"bd3583ae": "setTrustedRemote(uint256,bytes32)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506124e4806100206000396000f3fe6080604052600436106100f35760003560e01c80638da5cb5b1161008a578063c0c53b8b11610059578063c0c53b8b146102e6578063c0e07f2814610306578063f2fde38b1461031c578063f9ecc6f51461033c57600080fd5b80638da5cb5b1461024e578063a1a227fa14610279578063a6060871146102a6578063bd3583ae146102c657600080fd5b8063634d45b2116100c6578063634d45b2146101a4578063715018a6146101d15780637246a948146101e657806384a12b0f1461021357600080fd5b80631efedbe5146100f8578063230bb9f61461010d578063547cad12146101645780635c9c7c7314610184575b600080fd5b61010b61010636600461166a565b61035c565b005b34801561011957600080fd5b5060975461013a9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b34801561017057600080fd5b5061010b61017f3660046116ae565b61078d565b34801561019057600080fd5b5061010b61019f3660046116ae565b61083b565b3480156101b057600080fd5b506101c46101bf366004611790565b6108e9565b60405161015b9190611ca5565b3480156101dd57600080fd5b5061010b610900565b3480156101f257600080fd5b5060985461013a9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561021f57600080fd5b5061024061022e366004611cea565b60009081526066602052604090205490565b60405190815260200161015b565b34801561025a57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff1661013a565b34801561028557600080fd5b5060655461013a9073ffffffffffffffffffffffffffffffffffffffff1681565b3480156102b257600080fd5b5061010b6102c1366004611d03565b610973565b3480156102d257600080fd5b5061010b6102e136600461166a565b610a81565b3480156102f257600080fd5b5061010b610301366004611d9e565b610b37565b34801561031257600080fd5b5061024060995481565b34801561032857600080fd5b5061010b6103373660046116ae565b610ca6565b34801561034857600080fd5b5061010b610357366004611cea565b610da2565b6097546040517f21d80111000000000000000000000000000000000000000000000000000000008152600481018490528391839160009173ffffffffffffffffffffffffffffffffffffffff16906321d801119060240161088060405180830381865afa1580156103d1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103f59190612337565b606081015160a0015190915073ffffffffffffffffffffffffffffffffffffffff16156104695760405162461bcd60e51b815260206004820152601060248201527f6865726f206973207175657374696e670000000000000000000000000000000060448201526064015b60405180910390fd5b6098546040517f37e246ad0000000000000000000000000000000000000000000000000000000081526004810185905273ffffffffffffffffffffffffffffffffffffffff909116906337e246ad906024016020604051808303816000875af11580156104da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fe9190612354565b1561054b5760405162461bcd60e51b815260206004820152601160248201527f617373697374696e672061756374696f6e0000000000000000000000000000006044820152606401610460565b60008281526066602052604081205490610566853385610e0e565b905060006105b4609954604080517e01000000000000000000000000000000000000000000000000000000000000602082015260228082019390935281518082039093018352604201905290565b6097546040517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201523060248201526044810189905291925073ffffffffffffffffffffffffffffffffffffffff16906323b872dd90606401600060405180830381600087803b15801561062d57600080fd5b505af1158015610641573d6000803e3d6000fd5b50506097546040517f6352211e000000000000000000000000000000000000000000000000000000008152600481018a905230935073ffffffffffffffffffffffffffffffffffffffff9091169150636352211e90602401602060405180830381865afa1580156106b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106da919061236f565b73ffffffffffffffffffffffffffffffffffffffff161461073d5760405162461bcd60e51b815260206004820152601360248201527f4661696c656420746f206c6f636b204865726f000000000000000000000000006044820152606401610460565b61074983868484610e6f565b857fff49e1183195e20d542761218d85b1f570951619d00481ca84c2195144092bb28660405161077b91815260200190565b60405180910390a25050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146107f45760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b606580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60335473ffffffffffffffffffffffffffffffffffffffff1633146108a25760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b609880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6108f16113c2565b6108fa82610fd7565b92915050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146109675760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b6109716000610ffc565b565b60655473ffffffffffffffffffffffffffffffffffffffff1633146109da5760405162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f74206d65737361676520627573000000000000006044820152606401610460565b6000848152606660205260409020548514610a375760405162461bcd60e51b815260206004820152601a60248201527f496e76616c696420736f757263652073656e64696e67206170700000000000006044820152606401610460565b610a7a858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250879250611073915050565b5050505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610ae85760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b60008281526066602090815260409182902083905581518481529081018390527f642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03910160405180910390a15050565b600054610100900460ff16610b525760005460ff1615610b56565b303b155b610bc85760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610460565b600054610100900460ff16158015610c0757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000166101011790555b610c0f61133c565b6065805473ffffffffffffffffffffffffffffffffffffffff8087167fffffffffffffffffffffffff0000000000000000000000000000000000000000928316179092556097805486841690831617905560988054928516929091169190911790558015610ca057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610d0d5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b73ffffffffffffffffffffffffffffffffffffffff8116610d965760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610460565b610d9f81610ffc565b50565b60335473ffffffffffffffffffffffffffffffffffffffff163314610e095760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610460565b609955565b6060600060405180606001604052808481526020018573ffffffffffffffffffffffffffffffffffffffff16815260200186815250905080604051602001610e569190611ca5565b6040516020818303038152906040529150509392505050565b60008381526066602052604090205480610ecb5760405162461bcd60e51b815260206004820152601f60248201527f4e6f2072656d6f7465206170702073657420666f722064737420636861696e006044820152606401610460565b848114610f405760405162461bcd60e51b815260206004820152602660248201527f5265636569766572206973206e6f7420696e20747275737465642072656d6f7460448201527f65206170707300000000000000000000000000000000000000000000000000006064820152608401610460565b6065546040517fac8a4c1b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063ac8a4c1b903490610f9e9089908990899089906004016123f7565b6000604051808303818588803b158015610fb757600080fd5b505af1158015610fcb573d6000803e3d6000fd5b50505050505050505050565b610fdf6113c2565b600082806020019051810190610ff59190612433565b9392505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600061107e83610fd7565b8051602082015160408084015160975491517f6352211e00000000000000000000000000000000000000000000000000000000815260048101829052949550929391929173ffffffffffffffffffffffffffffffffffffffff90911690636352211e90602401602060405180830381865afa92505050801561113b575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682019092526111389181019061236f565b60015b6111d0576097546040517f1b8276710000000000000000000000000000000000000000000000000000000081526004810183905273ffffffffffffffffffffffffffffffffffffffff848116602483015290911690631b82767190604401600060405180830381600087803b1580156111b357600080fd5b505af11580156111c7573d6000803e3d6000fd5b50505050611282565b3073ffffffffffffffffffffffffffffffffffffffff821603611280576097546040517f42842e0e00000000000000000000000000000000000000000000000000000000815230600482015273ffffffffffffffffffffffffffffffffffffffff858116602483015260448201859052909116906342842e0e90606401600060405180830381600087803b15801561126757600080fd5b505af115801561127b573d6000803e3d6000fd5b505050505b505b6097546040517fb006410300000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9091169063b0064103906112d890869060040161249f565b600060405180830381600087803b1580156112f257600080fd5b505af1158015611306573d6000803e3d6000fd5b50505050807f1f3a4fd9d309a82e1d37743b6e0a35dfa60738cceb88aa6117021faad35957764660405161077b91815260200190565b600054610100900460ff166113b95760405162461bcd60e51b815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610460565b61097133610ffc565b60405180606001604052806113d56113e9565b815260006020820181905260409091015290565b604051806101000160405280600081526020016114416040518060c0016040528060008152602001600081526020016000815260200160008152602001600063ffffffff168152602001600063ffffffff1681525090565b815260200161144e611606565b81526020016114986040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081018290529060e082015290565b8152604080516101608101825260008082526020828101829052928201819052606082018190526080820181905260a0820181905260c0820181905260e08201819052610100820181905261012082018190526101408201529101908152604080516101c08101825260008082526020828101829052928201819052606082018190526080820181905260a0820181905260c0820181905260e08201819052610100820181905261012082018190526101408201819052610160820181905261018082018190526101a08201529101908152604080516101c08101825260008082526020828101829052928201819052606082018190526080820181905260a0820181905260c0820181905260e08201819052610100820181905261012082018190526101408201819052610160820181905261018082018190526101a082015291019081526040805160808101825260008082526020828101829052928201819052606082015291015290565b6040518061014001604052806000815260200160008152602001600060048111156116335761163361185f565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c0820181905260e09091015290565b6000806040838503121561167d57600080fd5b50508035926020909101359150565b73ffffffffffffffffffffffffffffffffffffffff81168114610d9f57600080fd5b6000602082840312156116c057600080fd5b8135610ff58161168c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff8111828210171561171e5761171e6116cb565b60405290565b604051610100810167ffffffffffffffff8111828210171561171e5761171e6116cb565b604051610160810167ffffffffffffffff8111828210171561171e5761171e6116cb565b6040516101c0810167ffffffffffffffff8111828210171561171e5761171e6116cb565b6000602082840312156117a257600080fd5b813567ffffffffffffffff808211156117ba57600080fd5b818401915084601f8301126117ce57600080fd5b8135818111156117e0576117e06116cb565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715611826576118266116cb565b8160405282815287602084870101111561183f57600080fd5b826020860160208301376000928101602001929092525095945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6005811061189e5761189e61185f565b9052565b805182526020810151602083015260408101516118c2604084018261188e565b5060608101516118d6606084018215159052565b5060808101516118ec608084018261ffff169052565b5060a081015161190460a084018263ffffffff169052565b5060c081015161191c60c084018263ffffffff169052565b5060e081015161193160e084018260ff169052565b506101008181015160ff81168483015250506101208181015160ff811684830152610ca0565b80518252602081015160208301526040810151604083015261ffff606082015116606083015267ffffffffffffffff608082015116608083015273ffffffffffffffffffffffffffffffffffffffff60a08201511660a083015260ff60c08201511660c083015260e0810151600281106119d3576119d361185f565b8060e0840152505050565b805161ffff16825260208101516119fb602084018261ffff169052565b506040810151611a11604084018261ffff169052565b506060810151611a27606084018261ffff169052565b506080810151611a3d608084018261ffff169052565b5060a0810151611a5360a084018261ffff169052565b5060c0810151611a6960c084018261ffff169052565b5060e0810151611a7f60e084018261ffff169052565b506101008181015161ffff81168483015250506101208181015161ffff81168483015250506101408181015161ffff811684830152610ca0565b805161ffff1682526020810151611ad6602084018261ffff169052565b506040810151611aec604084018261ffff169052565b506060810151611b02606084018261ffff169052565b506080810151611b18608084018261ffff169052565b5060a0810151611b2e60a084018261ffff169052565b5060c0810151611b4460c084018261ffff169052565b5060e0810151611b5a60e084018261ffff169052565b506101008181015161ffff90811691840191909152610120808301518216908401526101408083015182169084015261016080830151821690840152610180808301518216908401526101a0808301519182168185015290610ca0565b80518252602081015180516020840152602081015160408401526040810151606084015260608101516080840152608081015163ffffffff80821660a08601528060a08401511660c08601525050506040810151611c1860e08401826118a2565b506060810151611c2c610220840182611957565b506080810151611c406103208401826119de565b5060a0810151611c54610480840182611ab9565b5060c0810151611c68610640840182611ab9565b5060e00151805161ffff90811661080084015260208201518116610820840152604082015181166108408401526060909101511661086090910152565b60006108c082019050611cb9828451611bb7565b602083015173ffffffffffffffffffffffffffffffffffffffff166108808301526040909201516108a09091015290565b600060208284031215611cfc57600080fd5b5035919050565b600080600080600060808688031215611d1b57600080fd5b8535945060208601359350604086013567ffffffffffffffff80821115611d4157600080fd5b818801915088601f830112611d5557600080fd5b813581811115611d6457600080fd5b896020828501011115611d7657600080fd5b6020830195508094505050506060860135611d908161168c565b809150509295509295909350565b600080600060608486031215611db357600080fd5b8335611dbe8161168c565b92506020840135611dce8161168c565b91506040840135611dde8161168c565b809150509250925092565b805163ffffffff81168114611dfd57600080fd5b919050565b600060c08284031215611e1457600080fd5b60405160c0810181811067ffffffffffffffff82111715611e3757611e376116cb565b806040525080915082518152602083015160208201526040830151604082015260608301516060820152611e6d60808401611de9565b6080820152611e7e60a08401611de9565b60a08201525092915050565b805160058110611dfd57600080fd5b80518015158114611dfd57600080fd5b805161ffff81168114611dfd57600080fd5b805160ff81168114611dfd57600080fd5b60006101408284031215611edf57600080fd5b611ee76116fa565b90508151815260208201516020820152611f0360408301611e8a565b6040820152611f1460608301611e99565b6060820152611f2560808301611ea9565b6080820152611f3660a08301611de9565b60a0820152611f4760c08301611de9565b60c0820152611f5860e08301611ebb565b60e0820152610100611f6b818401611ebb565b90820152610120611f7d838201611ebb565b9082015292915050565b8051611dfd8161168c565b805160028110611dfd57600080fd5b60006101008284031215611fb457600080fd5b611fbc611724565b9050815181526020820151602082015260408201516040820152611fe260608301611ea9565b6060820152608082015167ffffffffffffffff8116811461200257600080fd5b608082015261201360a08301611f87565b60a082015261202460c08301611ebb565b60c082015261203560e08301611f92565b60e082015292915050565b6000610160828403121561205357600080fd5b61205b611748565b905061206682611ea9565b815261207460208301611ea9565b602082015261208560408301611ea9565b604082015261209660608301611ea9565b60608201526120a760808301611ea9565b60808201526120b860a08301611ea9565b60a08201526120c960c08301611ea9565b60c08201526120da60e08301611ea9565b60e08201526101006120ed818401611ea9565b908201526101206120ff838201611ea9565b90820152610140611f7d838201611ea9565b60006101c0828403121561212457600080fd5b61212c61176c565b905061213782611ea9565b815261214560208301611ea9565b602082015261215660408301611ea9565b604082015261216760608301611ea9565b606082015261217860808301611ea9565b608082015261218960a08301611ea9565b60a082015261219a60c08301611ea9565b60c08201526121ab60e08301611ea9565b60e08201526101006121be818401611ea9565b908201526101206121d0838201611ea9565b908201526101406121e2838201611ea9565b908201526101606121f4838201611ea9565b90820152610180612206838201611ea9565b908201526101a0611f7d838201611ea9565b60006080828403121561222a57600080fd5b6040516080810181811067ffffffffffffffff8211171561224d5761224d6116cb565b60405290508061225c83611ea9565b815261226a60208401611ea9565b602082015261227b60408401611ea9565b604082015261228c60608401611ea9565b60608201525092915050565b600061088082840312156122ab57600080fd5b6122b3611724565b9050815181526122c68360208401611e02565b60208201526122d88360e08401611ecc565b60408201526122eb836102208401611fa1565b60608201526122fe836103208401612040565b6080820152612311836104808401612111565b60a0820152612324836106408401612111565b60c0820152612035836108008401612218565b6000610880828403121561234a57600080fd5b610ff58383612298565b60006020828403121561236657600080fd5b610ff582611e99565b60006020828403121561238157600080fd5b8151610ff58161168c565b6000815180845260005b818110156123b257602081850181015186830182015201612396565b818111156123c4576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b848152836020820152608060408201526000612416608083018561238c565b8281036060840152612428818561238c565b979650505050505050565b60006108c0828403121561244657600080fd5b6040516060810181811067ffffffffffffffff82111715612469576124696116cb565b6040526124768484612298565b81526108808301516124878161168c565b60208201526108a09290920151604083015250919050565b61088081016108fa8284611bb756fea2646970667358221220a492c15d1336b7418bba663bf0175c1140180dad27b80895d6ec6762a05ca9c064736f6c634300080d0033",
}

// HeroBridgeUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use HeroBridgeUpgradeableMetaData.ABI instead.
var HeroBridgeUpgradeableABI = HeroBridgeUpgradeableMetaData.ABI

// Deprecated: Use HeroBridgeUpgradeableMetaData.Sigs instead.
// HeroBridgeUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var HeroBridgeUpgradeableFuncSigs = HeroBridgeUpgradeableMetaData.Sigs

// HeroBridgeUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HeroBridgeUpgradeableMetaData.Bin instead.
var HeroBridgeUpgradeableBin = HeroBridgeUpgradeableMetaData.Bin

// DeployHeroBridgeUpgradeable deploys a new Ethereum contract, binding an instance of HeroBridgeUpgradeable to it.
func DeployHeroBridgeUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HeroBridgeUpgradeable, error) {
	parsed, err := HeroBridgeUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HeroBridgeUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HeroBridgeUpgradeable{HeroBridgeUpgradeableCaller: HeroBridgeUpgradeableCaller{contract: contract}, HeroBridgeUpgradeableTransactor: HeroBridgeUpgradeableTransactor{contract: contract}, HeroBridgeUpgradeableFilterer: HeroBridgeUpgradeableFilterer{contract: contract}}, nil
}

// HeroBridgeUpgradeable is an auto generated Go binding around an Ethereum contract.
type HeroBridgeUpgradeable struct {
	HeroBridgeUpgradeableCaller     // Read-only binding to the contract
	HeroBridgeUpgradeableTransactor // Write-only binding to the contract
	HeroBridgeUpgradeableFilterer   // Log filterer for contract events
}

// HeroBridgeUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeroBridgeUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeroBridgeUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeroBridgeUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeroBridgeUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeroBridgeUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeroBridgeUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeroBridgeUpgradeableSession struct {
	Contract     *HeroBridgeUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// HeroBridgeUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeroBridgeUpgradeableCallerSession struct {
	Contract *HeroBridgeUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// HeroBridgeUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeroBridgeUpgradeableTransactorSession struct {
	Contract     *HeroBridgeUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// HeroBridgeUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeroBridgeUpgradeableRaw struct {
	Contract *HeroBridgeUpgradeable // Generic contract binding to access the raw methods on
}

// HeroBridgeUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeroBridgeUpgradeableCallerRaw struct {
	Contract *HeroBridgeUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// HeroBridgeUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeroBridgeUpgradeableTransactorRaw struct {
	Contract *HeroBridgeUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeroBridgeUpgradeable creates a new instance of HeroBridgeUpgradeable, bound to a specific deployed contract.
func NewHeroBridgeUpgradeable(address common.Address, backend bind.ContractBackend) (*HeroBridgeUpgradeable, error) {
	contract, err := bindHeroBridgeUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeable{HeroBridgeUpgradeableCaller: HeroBridgeUpgradeableCaller{contract: contract}, HeroBridgeUpgradeableTransactor: HeroBridgeUpgradeableTransactor{contract: contract}, HeroBridgeUpgradeableFilterer: HeroBridgeUpgradeableFilterer{contract: contract}}, nil
}

// NewHeroBridgeUpgradeableCaller creates a new read-only instance of HeroBridgeUpgradeable, bound to a specific deployed contract.
func NewHeroBridgeUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*HeroBridgeUpgradeableCaller, error) {
	contract, err := bindHeroBridgeUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeableCaller{contract: contract}, nil
}

// NewHeroBridgeUpgradeableTransactor creates a new write-only instance of HeroBridgeUpgradeable, bound to a specific deployed contract.
func NewHeroBridgeUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*HeroBridgeUpgradeableTransactor, error) {
	contract, err := bindHeroBridgeUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeableTransactor{contract: contract}, nil
}

// NewHeroBridgeUpgradeableFilterer creates a new log filterer instance of HeroBridgeUpgradeable, bound to a specific deployed contract.
func NewHeroBridgeUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*HeroBridgeUpgradeableFilterer, error) {
	contract, err := bindHeroBridgeUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeableFilterer{contract: contract}, nil
}

// bindHeroBridgeUpgradeable binds a generic wrapper to an already deployed contract.
func bindHeroBridgeUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := HeroBridgeUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeroBridgeUpgradeable.Contract.HeroBridgeUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.HeroBridgeUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.HeroBridgeUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeroBridgeUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// AssistingAuction is a free data retrieval call binding the contract method 0x7246a948.
//
// Solidity: function assistingAuction() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCaller) AssistingAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HeroBridgeUpgradeable.contract.Call(opts, &out, "assistingAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssistingAuction is a free data retrieval call binding the contract method 0x7246a948.
//
// Solidity: function assistingAuction() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) AssistingAuction() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.AssistingAuction(&_HeroBridgeUpgradeable.CallOpts)
}

// AssistingAuction is a free data retrieval call binding the contract method 0x7246a948.
//
// Solidity: function assistingAuction() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerSession) AssistingAuction() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.AssistingAuction(&_HeroBridgeUpgradeable.CallOpts)
}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns(((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)),address,uint256))
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCaller) DecodeMessage(opts *bind.CallOpts, _message []byte) (HeroBridgeUpgradeableMessageFormat, error) {
	var out []interface{}
	err := _HeroBridgeUpgradeable.contract.Call(opts, &out, "decodeMessage", _message)

	if err != nil {
		return *new(HeroBridgeUpgradeableMessageFormat), err
	}

	out0 := *abi.ConvertType(out[0], new(HeroBridgeUpgradeableMessageFormat)).(*HeroBridgeUpgradeableMessageFormat)

	return out0, err

}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns(((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)),address,uint256))
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) DecodeMessage(_message []byte) (HeroBridgeUpgradeableMessageFormat, error) {
	return _HeroBridgeUpgradeable.Contract.DecodeMessage(&_HeroBridgeUpgradeable.CallOpts, _message)
}

// DecodeMessage is a free data retrieval call binding the contract method 0x634d45b2.
//
// Solidity: function decodeMessage(bytes _message) pure returns(((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)),address,uint256))
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerSession) DecodeMessage(_message []byte) (HeroBridgeUpgradeableMessageFormat, error) {
	return _HeroBridgeUpgradeable.Contract.DecodeMessage(&_HeroBridgeUpgradeable.CallOpts, _message)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCaller) GetTrustedRemote(opts *bind.CallOpts, _chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _HeroBridgeUpgradeable.contract.Call(opts, &out, "getTrustedRemote", _chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _HeroBridgeUpgradeable.Contract.GetTrustedRemote(&_HeroBridgeUpgradeable.CallOpts, _chainId)
}

// GetTrustedRemote is a free data retrieval call binding the contract method 0x84a12b0f.
//
// Solidity: function getTrustedRemote(uint256 _chainId) view returns(bytes32 trustedRemote)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerSession) GetTrustedRemote(_chainId *big.Int) ([32]byte, error) {
	return _HeroBridgeUpgradeable.Contract.GetTrustedRemote(&_HeroBridgeUpgradeable.CallOpts, _chainId)
}

// Heroes is a free data retrieval call binding the contract method 0x230bb9f6.
//
// Solidity: function heroes() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCaller) Heroes(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HeroBridgeUpgradeable.contract.Call(opts, &out, "heroes")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Heroes is a free data retrieval call binding the contract method 0x230bb9f6.
//
// Solidity: function heroes() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) Heroes() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.Heroes(&_HeroBridgeUpgradeable.CallOpts)
}

// Heroes is a free data retrieval call binding the contract method 0x230bb9f6.
//
// Solidity: function heroes() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerSession) Heroes() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.Heroes(&_HeroBridgeUpgradeable.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCaller) MessageBus(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HeroBridgeUpgradeable.contract.Call(opts, &out, "messageBus")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) MessageBus() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.MessageBus(&_HeroBridgeUpgradeable.CallOpts)
}

// MessageBus is a free data retrieval call binding the contract method 0xa1a227fa.
//
// Solidity: function messageBus() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerSession) MessageBus() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.MessageBus(&_HeroBridgeUpgradeable.CallOpts)
}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCaller) MsgGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeroBridgeUpgradeable.contract.Call(opts, &out, "msgGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) MsgGasLimit() (*big.Int, error) {
	return _HeroBridgeUpgradeable.Contract.MsgGasLimit(&_HeroBridgeUpgradeable.CallOpts)
}

// MsgGasLimit is a free data retrieval call binding the contract method 0xc0e07f28.
//
// Solidity: function msgGasLimit() view returns(uint256)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerSession) MsgGasLimit() (*big.Int, error) {
	return _HeroBridgeUpgradeable.Contract.MsgGasLimit(&_HeroBridgeUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _HeroBridgeUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) Owner() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.Owner(&_HeroBridgeUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableCallerSession) Owner() (common.Address, error) {
	return _HeroBridgeUpgradeable.Contract.Owner(&_HeroBridgeUpgradeable.CallOpts)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) ExecuteMessage(opts *bind.TransactOpts, _srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "executeMessage", _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.ExecuteMessage(&_HeroBridgeUpgradeable.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// ExecuteMessage is a paid mutator transaction binding the contract method 0xa6060871.
//
// Solidity: function executeMessage(bytes32 _srcAddress, uint256 _srcChainId, bytes _message, address _executor) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) ExecuteMessage(_srcAddress [32]byte, _srcChainId *big.Int, _message []byte, _executor common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.ExecuteMessage(&_HeroBridgeUpgradeable.TransactOpts, _srcAddress, _srcChainId, _message, _executor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _messageBus, address _heroes, address _assistingAuction) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) Initialize(opts *bind.TransactOpts, _messageBus common.Address, _heroes common.Address, _assistingAuction common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "initialize", _messageBus, _heroes, _assistingAuction)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _messageBus, address _heroes, address _assistingAuction) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) Initialize(_messageBus common.Address, _heroes common.Address, _assistingAuction common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.Initialize(&_HeroBridgeUpgradeable.TransactOpts, _messageBus, _heroes, _assistingAuction)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _messageBus, address _heroes, address _assistingAuction) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) Initialize(_messageBus common.Address, _heroes common.Address, _assistingAuction common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.Initialize(&_HeroBridgeUpgradeable.TransactOpts, _messageBus, _heroes, _assistingAuction)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.RenounceOwnership(&_HeroBridgeUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.RenounceOwnership(&_HeroBridgeUpgradeable.TransactOpts)
}

// SendHero is a paid mutator transaction binding the contract method 0x1efedbe5.
//
// Solidity: function sendHero(uint256 _heroId, uint256 _dstChainId) payable returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) SendHero(opts *bind.TransactOpts, _heroId *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "sendHero", _heroId, _dstChainId)
}

// SendHero is a paid mutator transaction binding the contract method 0x1efedbe5.
//
// Solidity: function sendHero(uint256 _heroId, uint256 _dstChainId) payable returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) SendHero(_heroId *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SendHero(&_HeroBridgeUpgradeable.TransactOpts, _heroId, _dstChainId)
}

// SendHero is a paid mutator transaction binding the contract method 0x1efedbe5.
//
// Solidity: function sendHero(uint256 _heroId, uint256 _dstChainId) payable returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) SendHero(_heroId *big.Int, _dstChainId *big.Int) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SendHero(&_HeroBridgeUpgradeable.TransactOpts, _heroId, _dstChainId)
}

// SetAssistingAuctionAddress is a paid mutator transaction binding the contract method 0x5c9c7c73.
//
// Solidity: function setAssistingAuctionAddress(address _assistingAuction) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) SetAssistingAuctionAddress(opts *bind.TransactOpts, _assistingAuction common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "setAssistingAuctionAddress", _assistingAuction)
}

// SetAssistingAuctionAddress is a paid mutator transaction binding the contract method 0x5c9c7c73.
//
// Solidity: function setAssistingAuctionAddress(address _assistingAuction) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) SetAssistingAuctionAddress(_assistingAuction common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetAssistingAuctionAddress(&_HeroBridgeUpgradeable.TransactOpts, _assistingAuction)
}

// SetAssistingAuctionAddress is a paid mutator transaction binding the contract method 0x5c9c7c73.
//
// Solidity: function setAssistingAuctionAddress(address _assistingAuction) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) SetAssistingAuctionAddress(_assistingAuction common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetAssistingAuctionAddress(&_HeroBridgeUpgradeable.TransactOpts, _assistingAuction)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) SetMessageBus(opts *bind.TransactOpts, _messageBus common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "setMessageBus", _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetMessageBus(&_HeroBridgeUpgradeable.TransactOpts, _messageBus)
}

// SetMessageBus is a paid mutator transaction binding the contract method 0x547cad12.
//
// Solidity: function setMessageBus(address _messageBus) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) SetMessageBus(_messageBus common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetMessageBus(&_HeroBridgeUpgradeable.TransactOpts, _messageBus)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) SetMsgGasLimit(opts *bind.TransactOpts, _msgGasLimit *big.Int) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "setMsgGasLimit", _msgGasLimit)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) SetMsgGasLimit(_msgGasLimit *big.Int) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetMsgGasLimit(&_HeroBridgeUpgradeable.TransactOpts, _msgGasLimit)
}

// SetMsgGasLimit is a paid mutator transaction binding the contract method 0xf9ecc6f5.
//
// Solidity: function setMsgGasLimit(uint256 _msgGasLimit) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) SetMsgGasLimit(_msgGasLimit *big.Int) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetMsgGasLimit(&_HeroBridgeUpgradeable.TransactOpts, _msgGasLimit)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) SetTrustedRemote(opts *bind.TransactOpts, _srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "setTrustedRemote", _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetTrustedRemote(&_HeroBridgeUpgradeable.TransactOpts, _srcChainId, _srcAddress)
}

// SetTrustedRemote is a paid mutator transaction binding the contract method 0xbd3583ae.
//
// Solidity: function setTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) SetTrustedRemote(_srcChainId *big.Int, _srcAddress [32]byte) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.SetTrustedRemote(&_HeroBridgeUpgradeable.TransactOpts, _srcChainId, _srcAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.TransferOwnership(&_HeroBridgeUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _HeroBridgeUpgradeable.Contract.TransferOwnership(&_HeroBridgeUpgradeable.TransactOpts, newOwner)
}

// HeroBridgeUpgradeableHeroArrivedIterator is returned from FilterHeroArrived and is used to iterate over the raw logs and unpacked data for HeroArrived events raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableHeroArrivedIterator struct {
	Event *HeroBridgeUpgradeableHeroArrived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroBridgeUpgradeableHeroArrivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroBridgeUpgradeableHeroArrived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroBridgeUpgradeableHeroArrived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroBridgeUpgradeableHeroArrivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroBridgeUpgradeableHeroArrivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroBridgeUpgradeableHeroArrived represents a HeroArrived event raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableHeroArrived struct {
	HeroId         *big.Int
	ArrivalChainId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterHeroArrived is a free log retrieval operation binding the contract event 0x1f3a4fd9d309a82e1d37743b6e0a35dfa60738cceb88aa6117021faad3595776.
//
// Solidity: event HeroArrived(uint256 indexed heroId, uint256 arrivalChainId)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) FilterHeroArrived(opts *bind.FilterOpts, heroId []*big.Int) (*HeroBridgeUpgradeableHeroArrivedIterator, error) {

	var heroIdRule []interface{}
	for _, heroIdItem := range heroId {
		heroIdRule = append(heroIdRule, heroIdItem)
	}

	logs, sub, err := _HeroBridgeUpgradeable.contract.FilterLogs(opts, "HeroArrived", heroIdRule)
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeableHeroArrivedIterator{contract: _HeroBridgeUpgradeable.contract, event: "HeroArrived", logs: logs, sub: sub}, nil
}

// WatchHeroArrived is a free log subscription operation binding the contract event 0x1f3a4fd9d309a82e1d37743b6e0a35dfa60738cceb88aa6117021faad3595776.
//
// Solidity: event HeroArrived(uint256 indexed heroId, uint256 arrivalChainId)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) WatchHeroArrived(opts *bind.WatchOpts, sink chan<- *HeroBridgeUpgradeableHeroArrived, heroId []*big.Int) (event.Subscription, error) {

	var heroIdRule []interface{}
	for _, heroIdItem := range heroId {
		heroIdRule = append(heroIdRule, heroIdItem)
	}

	logs, sub, err := _HeroBridgeUpgradeable.contract.WatchLogs(opts, "HeroArrived", heroIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroBridgeUpgradeableHeroArrived)
				if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "HeroArrived", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHeroArrived is a log parse operation binding the contract event 0x1f3a4fd9d309a82e1d37743b6e0a35dfa60738cceb88aa6117021faad3595776.
//
// Solidity: event HeroArrived(uint256 indexed heroId, uint256 arrivalChainId)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) ParseHeroArrived(log types.Log) (*HeroBridgeUpgradeableHeroArrived, error) {
	event := new(HeroBridgeUpgradeableHeroArrived)
	if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "HeroArrived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroBridgeUpgradeableHeroSentIterator is returned from FilterHeroSent and is used to iterate over the raw logs and unpacked data for HeroSent events raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableHeroSentIterator struct {
	Event *HeroBridgeUpgradeableHeroSent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroBridgeUpgradeableHeroSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroBridgeUpgradeableHeroSent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroBridgeUpgradeableHeroSent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroBridgeUpgradeableHeroSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroBridgeUpgradeableHeroSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroBridgeUpgradeableHeroSent represents a HeroSent event raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableHeroSent struct {
	HeroId         *big.Int
	ArrivalChainId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterHeroSent is a free log retrieval operation binding the contract event 0xff49e1183195e20d542761218d85b1f570951619d00481ca84c2195144092bb2.
//
// Solidity: event HeroSent(uint256 indexed heroId, uint256 arrivalChainId)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) FilterHeroSent(opts *bind.FilterOpts, heroId []*big.Int) (*HeroBridgeUpgradeableHeroSentIterator, error) {

	var heroIdRule []interface{}
	for _, heroIdItem := range heroId {
		heroIdRule = append(heroIdRule, heroIdItem)
	}

	logs, sub, err := _HeroBridgeUpgradeable.contract.FilterLogs(opts, "HeroSent", heroIdRule)
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeableHeroSentIterator{contract: _HeroBridgeUpgradeable.contract, event: "HeroSent", logs: logs, sub: sub}, nil
}

// WatchHeroSent is a free log subscription operation binding the contract event 0xff49e1183195e20d542761218d85b1f570951619d00481ca84c2195144092bb2.
//
// Solidity: event HeroSent(uint256 indexed heroId, uint256 arrivalChainId)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) WatchHeroSent(opts *bind.WatchOpts, sink chan<- *HeroBridgeUpgradeableHeroSent, heroId []*big.Int) (event.Subscription, error) {

	var heroIdRule []interface{}
	for _, heroIdItem := range heroId {
		heroIdRule = append(heroIdRule, heroIdItem)
	}

	logs, sub, err := _HeroBridgeUpgradeable.contract.WatchLogs(opts, "HeroSent", heroIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroBridgeUpgradeableHeroSent)
				if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "HeroSent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseHeroSent is a log parse operation binding the contract event 0xff49e1183195e20d542761218d85b1f570951619d00481ca84c2195144092bb2.
//
// Solidity: event HeroSent(uint256 indexed heroId, uint256 arrivalChainId)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) ParseHeroSent(log types.Log) (*HeroBridgeUpgradeableHeroSent, error) {
	event := new(HeroBridgeUpgradeableHeroSent)
	if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "HeroSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroBridgeUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableOwnershipTransferredIterator struct {
	Event *HeroBridgeUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroBridgeUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroBridgeUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroBridgeUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroBridgeUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroBridgeUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroBridgeUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HeroBridgeUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HeroBridgeUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeableOwnershipTransferredIterator{contract: _HeroBridgeUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HeroBridgeUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _HeroBridgeUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroBridgeUpgradeableOwnershipTransferred)
				if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*HeroBridgeUpgradeableOwnershipTransferred, error) {
	event := new(HeroBridgeUpgradeableOwnershipTransferred)
	if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HeroBridgeUpgradeableSetTrustedRemoteIterator is returned from FilterSetTrustedRemote and is used to iterate over the raw logs and unpacked data for SetTrustedRemote events raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableSetTrustedRemoteIterator struct {
	Event *HeroBridgeUpgradeableSetTrustedRemote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HeroBridgeUpgradeableSetTrustedRemoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HeroBridgeUpgradeableSetTrustedRemote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HeroBridgeUpgradeableSetTrustedRemote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HeroBridgeUpgradeableSetTrustedRemoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HeroBridgeUpgradeableSetTrustedRemoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HeroBridgeUpgradeableSetTrustedRemote represents a SetTrustedRemote event raised by the HeroBridgeUpgradeable contract.
type HeroBridgeUpgradeableSetTrustedRemote struct {
	SrcChainId *big.Int
	SrcAddress [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetTrustedRemote is a free log retrieval operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) FilterSetTrustedRemote(opts *bind.FilterOpts) (*HeroBridgeUpgradeableSetTrustedRemoteIterator, error) {

	logs, sub, err := _HeroBridgeUpgradeable.contract.FilterLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return &HeroBridgeUpgradeableSetTrustedRemoteIterator{contract: _HeroBridgeUpgradeable.contract, event: "SetTrustedRemote", logs: logs, sub: sub}, nil
}

// WatchSetTrustedRemote is a free log subscription operation binding the contract event 0x642e74356c0610a9f944fb1a2d88d2fb82c6b74921566eee8bc0f9bb30f74f03.
//
// Solidity: event SetTrustedRemote(uint256 _srcChainId, bytes32 _srcAddress)
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) WatchSetTrustedRemote(opts *bind.WatchOpts, sink chan<- *HeroBridgeUpgradeableSetTrustedRemote) (event.Subscription, error) {

	logs, sub, err := _HeroBridgeUpgradeable.contract.WatchLogs(opts, "SetTrustedRemote")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HeroBridgeUpgradeableSetTrustedRemote)
				if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
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
func (_HeroBridgeUpgradeable *HeroBridgeUpgradeableFilterer) ParseSetTrustedRemote(log types.Log) (*HeroBridgeUpgradeableSetTrustedRemote, error) {
	event := new(HeroBridgeUpgradeableSetTrustedRemote)
	if err := _HeroBridgeUpgradeable.contract.UnpackLog(event, "SetTrustedRemote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAssistingAuctionMetaData contains all meta data concerning the IAssistingAuction contract.
var IAssistingAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"}],\"name\":\"bidFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"heroCore\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isOnAuction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"jewelToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_feeAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_feePercents\",\"type\":\"uint256[]\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"598647f8": "bid(uint256,uint256)",
		"402ca9e0": "bidFor(address,uint256,uint256)",
		"96b5a755": "cancelAuction(uint256)",
		"878eb368": "cancelAuctionWhenPaused(uint256)",
		"431f21da": "createAuction(uint256,uint256,uint256,uint256)",
		"78bd7935": "getAuction(uint256)",
		"c55d0f56": "getCurrentPrice(uint256)",
		"e7bbda11": "heroCore()",
		"37e246ad": "isOnAuction(uint256)",
		"237cc032": "jewelToken()",
		"8da5cb5b": "owner()",
		"83b5ff8b": "ownerCut()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"16f81524": "setFees(address[],uint256[])",
		"f2fde38b": "transferOwnership(address)",
	},
}

// IAssistingAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use IAssistingAuctionMetaData.ABI instead.
var IAssistingAuctionABI = IAssistingAuctionMetaData.ABI

// Deprecated: Use IAssistingAuctionMetaData.Sigs instead.
// IAssistingAuctionFuncSigs maps the 4-byte function signature to its string representation.
var IAssistingAuctionFuncSigs = IAssistingAuctionMetaData.Sigs

// IAssistingAuction is an auto generated Go binding around an Ethereum contract.
type IAssistingAuction struct {
	IAssistingAuctionCaller     // Read-only binding to the contract
	IAssistingAuctionTransactor // Write-only binding to the contract
	IAssistingAuctionFilterer   // Log filterer for contract events
}

// IAssistingAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAssistingAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssistingAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAssistingAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssistingAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAssistingAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAssistingAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAssistingAuctionSession struct {
	Contract     *IAssistingAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IAssistingAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAssistingAuctionCallerSession struct {
	Contract *IAssistingAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IAssistingAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAssistingAuctionTransactorSession struct {
	Contract     *IAssistingAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IAssistingAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAssistingAuctionRaw struct {
	Contract *IAssistingAuction // Generic contract binding to access the raw methods on
}

// IAssistingAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAssistingAuctionCallerRaw struct {
	Contract *IAssistingAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// IAssistingAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAssistingAuctionTransactorRaw struct {
	Contract *IAssistingAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAssistingAuction creates a new instance of IAssistingAuction, bound to a specific deployed contract.
func NewIAssistingAuction(address common.Address, backend bind.ContractBackend) (*IAssistingAuction, error) {
	contract, err := bindIAssistingAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAssistingAuction{IAssistingAuctionCaller: IAssistingAuctionCaller{contract: contract}, IAssistingAuctionTransactor: IAssistingAuctionTransactor{contract: contract}, IAssistingAuctionFilterer: IAssistingAuctionFilterer{contract: contract}}, nil
}

// NewIAssistingAuctionCaller creates a new read-only instance of IAssistingAuction, bound to a specific deployed contract.
func NewIAssistingAuctionCaller(address common.Address, caller bind.ContractCaller) (*IAssistingAuctionCaller, error) {
	contract, err := bindIAssistingAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAssistingAuctionCaller{contract: contract}, nil
}

// NewIAssistingAuctionTransactor creates a new write-only instance of IAssistingAuction, bound to a specific deployed contract.
func NewIAssistingAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*IAssistingAuctionTransactor, error) {
	contract, err := bindIAssistingAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAssistingAuctionTransactor{contract: contract}, nil
}

// NewIAssistingAuctionFilterer creates a new log filterer instance of IAssistingAuction, bound to a specific deployed contract.
func NewIAssistingAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*IAssistingAuctionFilterer, error) {
	contract, err := bindIAssistingAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAssistingAuctionFilterer{contract: contract}, nil
}

// bindIAssistingAuction binds a generic wrapper to an already deployed contract.
func bindIAssistingAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAssistingAuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssistingAuction *IAssistingAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssistingAuction.Contract.IAssistingAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssistingAuction *IAssistingAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.IAssistingAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssistingAuction *IAssistingAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.IAssistingAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAssistingAuction *IAssistingAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAssistingAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAssistingAuction *IAssistingAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAssistingAuction *IAssistingAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_IAssistingAuction *IAssistingAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	var out []interface{}
	err := _IAssistingAuction.contract.Call(opts, &out, "getAuction", _tokenId)

	outstruct := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StartingPrice = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.EndingPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_IAssistingAuction *IAssistingAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _IAssistingAuction.Contract.GetAuction(&_IAssistingAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_IAssistingAuction *IAssistingAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _IAssistingAuction.Contract.GetAuction(&_IAssistingAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_IAssistingAuction *IAssistingAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IAssistingAuction.contract.Call(opts, &out, "getCurrentPrice", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_IAssistingAuction *IAssistingAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _IAssistingAuction.Contract.GetCurrentPrice(&_IAssistingAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_IAssistingAuction *IAssistingAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _IAssistingAuction.Contract.GetCurrentPrice(&_IAssistingAuction.CallOpts, _tokenId)
}

// HeroCore is a free data retrieval call binding the contract method 0xe7bbda11.
//
// Solidity: function heroCore() view returns(address)
func (_IAssistingAuction *IAssistingAuctionCaller) HeroCore(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAssistingAuction.contract.Call(opts, &out, "heroCore")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HeroCore is a free data retrieval call binding the contract method 0xe7bbda11.
//
// Solidity: function heroCore() view returns(address)
func (_IAssistingAuction *IAssistingAuctionSession) HeroCore() (common.Address, error) {
	return _IAssistingAuction.Contract.HeroCore(&_IAssistingAuction.CallOpts)
}

// HeroCore is a free data retrieval call binding the contract method 0xe7bbda11.
//
// Solidity: function heroCore() view returns(address)
func (_IAssistingAuction *IAssistingAuctionCallerSession) HeroCore() (common.Address, error) {
	return _IAssistingAuction.Contract.HeroCore(&_IAssistingAuction.CallOpts)
}

// JewelToken is a free data retrieval call binding the contract method 0x237cc032.
//
// Solidity: function jewelToken() view returns(address)
func (_IAssistingAuction *IAssistingAuctionCaller) JewelToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAssistingAuction.contract.Call(opts, &out, "jewelToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// JewelToken is a free data retrieval call binding the contract method 0x237cc032.
//
// Solidity: function jewelToken() view returns(address)
func (_IAssistingAuction *IAssistingAuctionSession) JewelToken() (common.Address, error) {
	return _IAssistingAuction.Contract.JewelToken(&_IAssistingAuction.CallOpts)
}

// JewelToken is a free data retrieval call binding the contract method 0x237cc032.
//
// Solidity: function jewelToken() view returns(address)
func (_IAssistingAuction *IAssistingAuctionCallerSession) JewelToken() (common.Address, error) {
	return _IAssistingAuction.Contract.JewelToken(&_IAssistingAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IAssistingAuction *IAssistingAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IAssistingAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IAssistingAuction *IAssistingAuctionSession) Owner() (common.Address, error) {
	return _IAssistingAuction.Contract.Owner(&_IAssistingAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IAssistingAuction *IAssistingAuctionCallerSession) Owner() (common.Address, error) {
	return _IAssistingAuction.Contract.Owner(&_IAssistingAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_IAssistingAuction *IAssistingAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAssistingAuction.contract.Call(opts, &out, "ownerCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_IAssistingAuction *IAssistingAuctionSession) OwnerCut() (*big.Int, error) {
	return _IAssistingAuction.Contract.OwnerCut(&_IAssistingAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_IAssistingAuction *IAssistingAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _IAssistingAuction.Contract.OwnerCut(&_IAssistingAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IAssistingAuction *IAssistingAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IAssistingAuction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IAssistingAuction *IAssistingAuctionSession) Paused() (bool, error) {
	return _IAssistingAuction.Contract.Paused(&_IAssistingAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IAssistingAuction *IAssistingAuctionCallerSession) Paused() (bool, error) {
	return _IAssistingAuction.Contract.Paused(&_IAssistingAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _tokenId, uint256 _bidAmount) returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "bid", _tokenId, _bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _tokenId, uint256 _bidAmount) returns()
func (_IAssistingAuction *IAssistingAuctionSession) Bid(_tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.Bid(&_IAssistingAuction.TransactOpts, _tokenId, _bidAmount)
}

// Bid is a paid mutator transaction binding the contract method 0x598647f8.
//
// Solidity: function bid(uint256 _tokenId, uint256 _bidAmount) returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) Bid(_tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.Bid(&_IAssistingAuction.TransactOpts, _tokenId, _bidAmount)
}

// BidFor is a paid mutator transaction binding the contract method 0x402ca9e0.
//
// Solidity: function bidFor(address _bidder, uint256 _tokenId, uint256 _bidAmount) returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) BidFor(opts *bind.TransactOpts, _bidder common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "bidFor", _bidder, _tokenId, _bidAmount)
}

// BidFor is a paid mutator transaction binding the contract method 0x402ca9e0.
//
// Solidity: function bidFor(address _bidder, uint256 _tokenId, uint256 _bidAmount) returns()
func (_IAssistingAuction *IAssistingAuctionSession) BidFor(_bidder common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.BidFor(&_IAssistingAuction.TransactOpts, _bidder, _tokenId, _bidAmount)
}

// BidFor is a paid mutator transaction binding the contract method 0x402ca9e0.
//
// Solidity: function bidFor(address _bidder, uint256 _tokenId, uint256 _bidAmount) returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) BidFor(_bidder common.Address, _tokenId *big.Int, _bidAmount *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.BidFor(&_IAssistingAuction.TransactOpts, _bidder, _tokenId, _bidAmount)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_IAssistingAuction *IAssistingAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.CancelAuction(&_IAssistingAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.CancelAuction(&_IAssistingAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_IAssistingAuction *IAssistingAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.CancelAuctionWhenPaused(&_IAssistingAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.CancelAuctionWhenPaused(&_IAssistingAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x431f21da.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x431f21da.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IAssistingAuction *IAssistingAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.CreateAuction(&_IAssistingAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x431f21da.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.CreateAuction(&_IAssistingAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration)
}

// IsOnAuction is a paid mutator transaction binding the contract method 0x37e246ad.
//
// Solidity: function isOnAuction(uint256 _tokenId) returns(bool)
func (_IAssistingAuction *IAssistingAuctionTransactor) IsOnAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "isOnAuction", _tokenId)
}

// IsOnAuction is a paid mutator transaction binding the contract method 0x37e246ad.
//
// Solidity: function isOnAuction(uint256 _tokenId) returns(bool)
func (_IAssistingAuction *IAssistingAuctionSession) IsOnAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.IsOnAuction(&_IAssistingAuction.TransactOpts, _tokenId)
}

// IsOnAuction is a paid mutator transaction binding the contract method 0x37e246ad.
//
// Solidity: function isOnAuction(uint256 _tokenId) returns(bool)
func (_IAssistingAuction *IAssistingAuctionTransactorSession) IsOnAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.IsOnAuction(&_IAssistingAuction.TransactOpts, _tokenId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IAssistingAuction *IAssistingAuctionSession) RenounceOwnership() (*types.Transaction, error) {
	return _IAssistingAuction.Contract.RenounceOwnership(&_IAssistingAuction.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _IAssistingAuction.Contract.RenounceOwnership(&_IAssistingAuction.TransactOpts)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) SetFees(opts *bind.TransactOpts, _feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "setFees", _feeAddresses, _feePercents)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_IAssistingAuction *IAssistingAuctionSession) SetFees(_feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.SetFees(&_IAssistingAuction.TransactOpts, _feeAddresses, _feePercents)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) SetFees(_feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.SetFees(&_IAssistingAuction.TransactOpts, _feeAddresses, _feePercents)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IAssistingAuction *IAssistingAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IAssistingAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IAssistingAuction *IAssistingAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.TransferOwnership(&_IAssistingAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IAssistingAuction *IAssistingAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IAssistingAuction.Contract.TransferOwnership(&_IAssistingAuction.TransactOpts, newOwner)
}

// IHeroCoreUpgradeableMetaData contains all meta data concerning the IHeroCoreUpgradeable contract.
var IHeroCoreUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"HERO_MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAMINA_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assistingAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseCooldown\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseSummonFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"bridgeMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"}],\"name\":\"calculateSummoningCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownPerGen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cooldownPerSummon\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createAssistingAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumRarity\",\"name\":\"_rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_shiny\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint256\",\"name\":\"createdBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"summonerTears\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"assistantTears\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"enhancementStone\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"}],\"internalType\":\"structHeroCrystal\",\"name\":\"_crystal\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_crystalId\",\"type\":\"uint256\"}],\"name\":\"createHero\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"crystalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_staminaDeduction\",\"type\":\"uint256\"}],\"name\":\"deductStamina\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"digits\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"extractNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"}],\"name\":\"getCurrentStamina\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getHero\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"summonedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSummonTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"summons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"}],\"internalType\":\"structSummoningInfo\",\"name\":\"summoningInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumRarity\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"shiny\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"class\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"subClass\",\"type\":\"uint8\"}],\"internalType\":\"structHeroInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"staminaFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"xp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"currentQuest\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sp\",\"type\":\"uint8\"},{\"internalType\":\"enumHeroStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structHeroState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stamina\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStats\",\"name\":\"stats\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStatGrowth\",\"name\":\"primaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStatGrowth\",\"name\":\"secondaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"mining\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gardening\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"foraging\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fishing\",\"type\":\"uint16\"}],\"internalType\":\"structHeroProfessions\",\"name\":\"professions\",\"type\":\"tuple\"}],\"internalType\":\"structHero\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increasePerGen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increasePerSummon\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_crystalAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_heroId\",\"type\":\"uint256\"}],\"name\":\"isReadyToSummon\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_crystalId\",\"type\":\"uint256\"}],\"name\":\"openCrystal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setAssistingAuctionAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_feeAddresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_feePercents\",\"type\":\"uint256[]\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseCooldown\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cooldownPerSummon\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_cooldownPerGen\",\"type\":\"uint256\"}],\"name\":\"setSummonCooldowns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_baseSummonFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_increasePerSummon\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_increasePerGen\",\"type\":\"uint256\"}],\"name\":\"setSummonFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_timePerStamina\",\"type\":\"uint256\"}],\"name\":\"setTimePerStamina\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_summonerTears\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_assistantTears\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_enhancementStone\",\"type\":\"address\"}],\"name\":\"summonCrystal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timePerStamina\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"summonedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextSummonTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"summonerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assistantId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"summons\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"maxSummons\",\"type\":\"uint32\"}],\"internalType\":\"structSummoningInfo\",\"name\":\"summoningInfo\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"statGenes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"visualGenes\",\"type\":\"uint256\"},{\"internalType\":\"enumRarity\",\"name\":\"rarity\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"shiny\",\"type\":\"bool\"},{\"internalType\":\"uint16\",\"name\":\"generation\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"firstName\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"lastName\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"shinyStyle\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"class\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"subClass\",\"type\":\"uint8\"}],\"internalType\":\"structHeroInfo\",\"name\":\"info\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"staminaFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mpFullAt\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"level\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"xp\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"currentQuest\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"sp\",\"type\":\"uint8\"},{\"internalType\":\"enumHeroStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"internalType\":\"structHeroState\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mp\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stamina\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStats\",\"name\":\"stats\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStatGrowth\",\"name\":\"primaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"strength\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"agility\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"intelligence\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"wisdom\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"luck\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"vitality\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"endurance\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"dexterity\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"hpLg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpSm\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpRg\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"mpLg\",\"type\":\"uint16\"}],\"internalType\":\"structHeroStatGrowth\",\"name\":\"secondaryStatGrowth\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"mining\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"gardening\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"foraging\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"fishing\",\"type\":\"uint16\"}],\"internalType\":\"structHeroProfessions\",\"name\":\"professions\",\"type\":\"tuple\"}],\"internalType\":\"structHero\",\"name\":\"_hero\",\"type\":\"tuple\"}],\"name\":\"updateHero\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"vrf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"39ab52d5": "HERO_MODERATOR_ROLE()",
		"d5391393": "MINTER_ROLE()",
		"797669c9": "MODERATOR_ROLE()",
		"e63ab1e9": "PAUSER_ROLE()",
		"630a0179": "STAMINA_ROLE()",
		"095ea7b3": "approve(address,uint256)",
		"7246a948": "assistingAuction()",
		"70a08231": "balanceOf(address)",
		"93deec27": "baseCooldown()",
		"45d31394": "baseSummonFee()",
		"1b827671": "bridgeMint(uint256,address)",
		"42966c68": "burn(uint256)",
		"a66f2039": "calculateSummoningCost(uint256)",
		"a8029920": "cooldownPerGen()",
		"0a595eb5": "cooldownPerSummon()",
		"9d6fac6f": "cooldowns(uint256)",
		"32847f0e": "createAssistingAuction(uint256,uint256,uint256,uint256)",
		"306e56f3": "createHero(uint256,uint256,uint8,bool,(address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8),uint256)",
		"3d7d3f5a": "createSaleAuction(uint256,uint256,uint256,uint256)",
		"dbf4025b": "crystalToken()",
		"f2e58229": "deductStamina(uint256,uint256)",
		"b00a7b9e": "extractNumber(uint256,uint256,uint256)",
		"f2b47d52": "geneScience()",
		"081812fc": "getApproved(uint256)",
		"df52458a": "getCurrentStamina(uint256)",
		"21d80111": "getHero(uint256)",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"9010d07c": "getRoleMember(bytes32,uint256)",
		"ca15c873": "getRoleMemberCount(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"20c20a07": "increasePerGen()",
		"e9dea449": "increasePerSummon()",
		"c4d66de8": "initialize(address)",
		"a6487c53": "initialize(string,string,string)",
		"e985e9c5": "isApprovedForAll(address,address)",
		"45ef4ecb": "isReadyToSummon(uint256)",
		"6a627842": "mint(address)",
		"06fdde03": "name()",
		"8ff41141": "openCrystal(uint256)",
		"6352211e": "ownerOf(uint256)",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"42842e0e": "safeTransferFrom(address,address,uint256)",
		"b88d4fde": "safeTransferFrom(address,address,uint256,bytes)",
		"e6cbe351": "saleAuction()",
		"a22cb465": "setApprovalForAll(address,bool)",
		"5c9c7c73": "setAssistingAuctionAddress(address)",
		"16f81524": "setFees(address[],uint256[])",
		"6fbde40d": "setSaleAuctionAddress(address)",
		"4e970324": "setSummonCooldowns(uint256,uint256,uint256)",
		"03466cfd": "setSummonFees(uint256,uint256,uint256)",
		"7fd73850": "setTimePerStamina(uint256)",
		"5880d8e6": "summonCrystal(uint256,uint256,uint8,uint8,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"95d89b41": "symbol()",
		"96072223": "timePerStamina()",
		"4f6ccce7": "tokenByIndex(uint256)",
		"2f745c59": "tokenOfOwnerByIndex(address,uint256)",
		"c87b56dd": "tokenURI(uint256)",
		"18160ddd": "totalSupply()",
		"23b872dd": "transferFrom(address,address,uint256)",
		"3f4ba83a": "unpause()",
		"b0064103": "updateHero((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)))",
		"4b757f98": "vrf(uint256)",
	},
}

// IHeroCoreUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IHeroCoreUpgradeableMetaData.ABI instead.
var IHeroCoreUpgradeableABI = IHeroCoreUpgradeableMetaData.ABI

// Deprecated: Use IHeroCoreUpgradeableMetaData.Sigs instead.
// IHeroCoreUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IHeroCoreUpgradeableFuncSigs = IHeroCoreUpgradeableMetaData.Sigs

// IHeroCoreUpgradeable is an auto generated Go binding around an Ethereum contract.
type IHeroCoreUpgradeable struct {
	IHeroCoreUpgradeableCaller     // Read-only binding to the contract
	IHeroCoreUpgradeableTransactor // Write-only binding to the contract
	IHeroCoreUpgradeableFilterer   // Log filterer for contract events
}

// IHeroCoreUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IHeroCoreUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHeroCoreUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IHeroCoreUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHeroCoreUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IHeroCoreUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IHeroCoreUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IHeroCoreUpgradeableSession struct {
	Contract     *IHeroCoreUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IHeroCoreUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IHeroCoreUpgradeableCallerSession struct {
	Contract *IHeroCoreUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IHeroCoreUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IHeroCoreUpgradeableTransactorSession struct {
	Contract     *IHeroCoreUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IHeroCoreUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IHeroCoreUpgradeableRaw struct {
	Contract *IHeroCoreUpgradeable // Generic contract binding to access the raw methods on
}

// IHeroCoreUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IHeroCoreUpgradeableCallerRaw struct {
	Contract *IHeroCoreUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IHeroCoreUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IHeroCoreUpgradeableTransactorRaw struct {
	Contract *IHeroCoreUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIHeroCoreUpgradeable creates a new instance of IHeroCoreUpgradeable, bound to a specific deployed contract.
func NewIHeroCoreUpgradeable(address common.Address, backend bind.ContractBackend) (*IHeroCoreUpgradeable, error) {
	contract, err := bindIHeroCoreUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IHeroCoreUpgradeable{IHeroCoreUpgradeableCaller: IHeroCoreUpgradeableCaller{contract: contract}, IHeroCoreUpgradeableTransactor: IHeroCoreUpgradeableTransactor{contract: contract}, IHeroCoreUpgradeableFilterer: IHeroCoreUpgradeableFilterer{contract: contract}}, nil
}

// NewIHeroCoreUpgradeableCaller creates a new read-only instance of IHeroCoreUpgradeable, bound to a specific deployed contract.
func NewIHeroCoreUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IHeroCoreUpgradeableCaller, error) {
	contract, err := bindIHeroCoreUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IHeroCoreUpgradeableCaller{contract: contract}, nil
}

// NewIHeroCoreUpgradeableTransactor creates a new write-only instance of IHeroCoreUpgradeable, bound to a specific deployed contract.
func NewIHeroCoreUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IHeroCoreUpgradeableTransactor, error) {
	contract, err := bindIHeroCoreUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IHeroCoreUpgradeableTransactor{contract: contract}, nil
}

// NewIHeroCoreUpgradeableFilterer creates a new log filterer instance of IHeroCoreUpgradeable, bound to a specific deployed contract.
func NewIHeroCoreUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IHeroCoreUpgradeableFilterer, error) {
	contract, err := bindIHeroCoreUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IHeroCoreUpgradeableFilterer{contract: contract}, nil
}

// bindIHeroCoreUpgradeable binds a generic wrapper to an already deployed contract.
func bindIHeroCoreUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IHeroCoreUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHeroCoreUpgradeable.Contract.IHeroCoreUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.IHeroCoreUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.IHeroCoreUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IHeroCoreUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.DEFAULTADMINROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.DEFAULTADMINROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// HEROMODERATORROLE is a free data retrieval call binding the contract method 0x39ab52d5.
//
// Solidity: function HERO_MODERATOR_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) HEROMODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "HERO_MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HEROMODERATORROLE is a free data retrieval call binding the contract method 0x39ab52d5.
//
// Solidity: function HERO_MODERATOR_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) HEROMODERATORROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.HEROMODERATORROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// HEROMODERATORROLE is a free data retrieval call binding the contract method 0x39ab52d5.
//
// Solidity: function HERO_MODERATOR_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) HEROMODERATORROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.HEROMODERATORROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) MINTERROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.MINTERROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) MINTERROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.MINTERROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) MODERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "MODERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) MODERATORROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.MODERATORROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// MODERATORROLE is a free data retrieval call binding the contract method 0x797669c9.
//
// Solidity: function MODERATOR_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) MODERATORROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.MODERATORROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) PAUSERROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.PAUSERROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) PAUSERROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.PAUSERROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// STAMINAROLE is a free data retrieval call binding the contract method 0x630a0179.
//
// Solidity: function STAMINA_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) STAMINAROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "STAMINA_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAMINAROLE is a free data retrieval call binding the contract method 0x630a0179.
//
// Solidity: function STAMINA_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) STAMINAROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.STAMINAROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// STAMINAROLE is a free data retrieval call binding the contract method 0x630a0179.
//
// Solidity: function STAMINA_ROLE() view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) STAMINAROLE() ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.STAMINAROLE(&_IHeroCoreUpgradeable.CallOpts)
}

// AssistingAuction is a free data retrieval call binding the contract method 0x7246a948.
//
// Solidity: function assistingAuction() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) AssistingAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "assistingAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AssistingAuction is a free data retrieval call binding the contract method 0x7246a948.
//
// Solidity: function assistingAuction() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) AssistingAuction() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.AssistingAuction(&_IHeroCoreUpgradeable.CallOpts)
}

// AssistingAuction is a free data retrieval call binding the contract method 0x7246a948.
//
// Solidity: function assistingAuction() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) AssistingAuction() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.AssistingAuction(&_IHeroCoreUpgradeable.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.BalanceOf(&_IHeroCoreUpgradeable.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.BalanceOf(&_IHeroCoreUpgradeable.CallOpts, owner)
}

// BaseCooldown is a free data retrieval call binding the contract method 0x93deec27.
//
// Solidity: function baseCooldown() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) BaseCooldown(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "baseCooldown")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseCooldown is a free data retrieval call binding the contract method 0x93deec27.
//
// Solidity: function baseCooldown() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) BaseCooldown() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.BaseCooldown(&_IHeroCoreUpgradeable.CallOpts)
}

// BaseCooldown is a free data retrieval call binding the contract method 0x93deec27.
//
// Solidity: function baseCooldown() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) BaseCooldown() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.BaseCooldown(&_IHeroCoreUpgradeable.CallOpts)
}

// BaseSummonFee is a free data retrieval call binding the contract method 0x45d31394.
//
// Solidity: function baseSummonFee() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) BaseSummonFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "baseSummonFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseSummonFee is a free data retrieval call binding the contract method 0x45d31394.
//
// Solidity: function baseSummonFee() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) BaseSummonFee() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.BaseSummonFee(&_IHeroCoreUpgradeable.CallOpts)
}

// BaseSummonFee is a free data retrieval call binding the contract method 0x45d31394.
//
// Solidity: function baseSummonFee() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) BaseSummonFee() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.BaseSummonFee(&_IHeroCoreUpgradeable.CallOpts)
}

// CalculateSummoningCost is a free data retrieval call binding the contract method 0xa66f2039.
//
// Solidity: function calculateSummoningCost(uint256 _heroId) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) CalculateSummoningCost(opts *bind.CallOpts, _heroId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "calculateSummoningCost", _heroId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSummoningCost is a free data retrieval call binding the contract method 0xa66f2039.
//
// Solidity: function calculateSummoningCost(uint256 _heroId) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) CalculateSummoningCost(_heroId *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.CalculateSummoningCost(&_IHeroCoreUpgradeable.CallOpts, _heroId)
}

// CalculateSummoningCost is a free data retrieval call binding the contract method 0xa66f2039.
//
// Solidity: function calculateSummoningCost(uint256 _heroId) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) CalculateSummoningCost(_heroId *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.CalculateSummoningCost(&_IHeroCoreUpgradeable.CallOpts, _heroId)
}

// CooldownPerGen is a free data retrieval call binding the contract method 0xa8029920.
//
// Solidity: function cooldownPerGen() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) CooldownPerGen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "cooldownPerGen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownPerGen is a free data retrieval call binding the contract method 0xa8029920.
//
// Solidity: function cooldownPerGen() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) CooldownPerGen() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.CooldownPerGen(&_IHeroCoreUpgradeable.CallOpts)
}

// CooldownPerGen is a free data retrieval call binding the contract method 0xa8029920.
//
// Solidity: function cooldownPerGen() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) CooldownPerGen() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.CooldownPerGen(&_IHeroCoreUpgradeable.CallOpts)
}

// CooldownPerSummon is a free data retrieval call binding the contract method 0x0a595eb5.
//
// Solidity: function cooldownPerSummon() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) CooldownPerSummon(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "cooldownPerSummon")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CooldownPerSummon is a free data retrieval call binding the contract method 0x0a595eb5.
//
// Solidity: function cooldownPerSummon() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) CooldownPerSummon() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.CooldownPerSummon(&_IHeroCoreUpgradeable.CallOpts)
}

// CooldownPerSummon is a free data retrieval call binding the contract method 0x0a595eb5.
//
// Solidity: function cooldownPerSummon() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) CooldownPerSummon() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.CooldownPerSummon(&_IHeroCoreUpgradeable.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _IHeroCoreUpgradeable.Contract.Cooldowns(&_IHeroCoreUpgradeable.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _IHeroCoreUpgradeable.Contract.Cooldowns(&_IHeroCoreUpgradeable.CallOpts, arg0)
}

// CrystalToken is a free data retrieval call binding the contract method 0xdbf4025b.
//
// Solidity: function crystalToken() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) CrystalToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "crystalToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CrystalToken is a free data retrieval call binding the contract method 0xdbf4025b.
//
// Solidity: function crystalToken() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) CrystalToken() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.CrystalToken(&_IHeroCoreUpgradeable.CallOpts)
}

// CrystalToken is a free data retrieval call binding the contract method 0xdbf4025b.
//
// Solidity: function crystalToken() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) CrystalToken() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.CrystalToken(&_IHeroCoreUpgradeable.CallOpts)
}

// ExtractNumber is a free data retrieval call binding the contract method 0xb00a7b9e.
//
// Solidity: function extractNumber(uint256 randomNumber, uint256 digits, uint256 offset) pure returns(uint256 result)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) ExtractNumber(opts *bind.CallOpts, randomNumber *big.Int, digits *big.Int, offset *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "extractNumber", randomNumber, digits, offset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtractNumber is a free data retrieval call binding the contract method 0xb00a7b9e.
//
// Solidity: function extractNumber(uint256 randomNumber, uint256 digits, uint256 offset) pure returns(uint256 result)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) ExtractNumber(randomNumber *big.Int, digits *big.Int, offset *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.ExtractNumber(&_IHeroCoreUpgradeable.CallOpts, randomNumber, digits, offset)
}

// ExtractNumber is a free data retrieval call binding the contract method 0xb00a7b9e.
//
// Solidity: function extractNumber(uint256 randomNumber, uint256 digits, uint256 offset) pure returns(uint256 result)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) ExtractNumber(randomNumber *big.Int, digits *big.Int, offset *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.ExtractNumber(&_IHeroCoreUpgradeable.CallOpts, randomNumber, digits, offset)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "geneScience")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GeneScience() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.GeneScience(&_IHeroCoreUpgradeable.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) GeneScience() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.GeneScience(&_IHeroCoreUpgradeable.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.GetApproved(&_IHeroCoreUpgradeable.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.GetApproved(&_IHeroCoreUpgradeable.CallOpts, tokenId)
}

// GetCurrentStamina is a free data retrieval call binding the contract method 0xdf52458a.
//
// Solidity: function getCurrentStamina(uint256 _heroId) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) GetCurrentStamina(opts *bind.CallOpts, _heroId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "getCurrentStamina", _heroId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentStamina is a free data retrieval call binding the contract method 0xdf52458a.
//
// Solidity: function getCurrentStamina(uint256 _heroId) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GetCurrentStamina(_heroId *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.GetCurrentStamina(&_IHeroCoreUpgradeable.CallOpts, _heroId)
}

// GetCurrentStamina is a free data retrieval call binding the contract method 0xdf52458a.
//
// Solidity: function getCurrentStamina(uint256 _heroId) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) GetCurrentStamina(_heroId *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.GetCurrentStamina(&_IHeroCoreUpgradeable.CallOpts, _heroId)
}

// GetHero is a free data retrieval call binding the contract method 0x21d80111.
//
// Solidity: function getHero(uint256 _id) view returns((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)))
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) GetHero(opts *bind.CallOpts, _id *big.Int) (Hero, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "getHero", _id)

	if err != nil {
		return *new(Hero), err
	}

	out0 := *abi.ConvertType(out[0], new(Hero)).(*Hero)

	return out0, err

}

// GetHero is a free data retrieval call binding the contract method 0x21d80111.
//
// Solidity: function getHero(uint256 _id) view returns((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)))
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GetHero(_id *big.Int) (Hero, error) {
	return _IHeroCoreUpgradeable.Contract.GetHero(&_IHeroCoreUpgradeable.CallOpts, _id)
}

// GetHero is a free data retrieval call binding the contract method 0x21d80111.
//
// Solidity: function getHero(uint256 _id) view returns((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)))
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) GetHero(_id *big.Int) (Hero, error) {
	return _IHeroCoreUpgradeable.Contract.GetHero(&_IHeroCoreUpgradeable.CallOpts, _id)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.GetRoleAdmin(&_IHeroCoreUpgradeable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.GetRoleAdmin(&_IHeroCoreUpgradeable.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.GetRoleMember(&_IHeroCoreUpgradeable.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.GetRoleMember(&_IHeroCoreUpgradeable.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.GetRoleMemberCount(&_IHeroCoreUpgradeable.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.GetRoleMemberCount(&_IHeroCoreUpgradeable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.HasRole(&_IHeroCoreUpgradeable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.HasRole(&_IHeroCoreUpgradeable.CallOpts, role, account)
}

// IncreasePerGen is a free data retrieval call binding the contract method 0x20c20a07.
//
// Solidity: function increasePerGen() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) IncreasePerGen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "increasePerGen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePerGen is a free data retrieval call binding the contract method 0x20c20a07.
//
// Solidity: function increasePerGen() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) IncreasePerGen() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.IncreasePerGen(&_IHeroCoreUpgradeable.CallOpts)
}

// IncreasePerGen is a free data retrieval call binding the contract method 0x20c20a07.
//
// Solidity: function increasePerGen() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) IncreasePerGen() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.IncreasePerGen(&_IHeroCoreUpgradeable.CallOpts)
}

// IncreasePerSummon is a free data retrieval call binding the contract method 0xe9dea449.
//
// Solidity: function increasePerSummon() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) IncreasePerSummon(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "increasePerSummon")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePerSummon is a free data retrieval call binding the contract method 0xe9dea449.
//
// Solidity: function increasePerSummon() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) IncreasePerSummon() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.IncreasePerSummon(&_IHeroCoreUpgradeable.CallOpts)
}

// IncreasePerSummon is a free data retrieval call binding the contract method 0xe9dea449.
//
// Solidity: function increasePerSummon() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) IncreasePerSummon() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.IncreasePerSummon(&_IHeroCoreUpgradeable.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.IsApprovedForAll(&_IHeroCoreUpgradeable.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.IsApprovedForAll(&_IHeroCoreUpgradeable.CallOpts, owner, operator)
}

// IsReadyToSummon is a free data retrieval call binding the contract method 0x45ef4ecb.
//
// Solidity: function isReadyToSummon(uint256 _heroId) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) IsReadyToSummon(opts *bind.CallOpts, _heroId *big.Int) (bool, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "isReadyToSummon", _heroId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyToSummon is a free data retrieval call binding the contract method 0x45ef4ecb.
//
// Solidity: function isReadyToSummon(uint256 _heroId) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) IsReadyToSummon(_heroId *big.Int) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.IsReadyToSummon(&_IHeroCoreUpgradeable.CallOpts, _heroId)
}

// IsReadyToSummon is a free data retrieval call binding the contract method 0x45ef4ecb.
//
// Solidity: function isReadyToSummon(uint256 _heroId) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) IsReadyToSummon(_heroId *big.Int) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.IsReadyToSummon(&_IHeroCoreUpgradeable.CallOpts, _heroId)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Name() (string, error) {
	return _IHeroCoreUpgradeable.Contract.Name(&_IHeroCoreUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) Name() (string, error) {
	return _IHeroCoreUpgradeable.Contract.Name(&_IHeroCoreUpgradeable.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.OwnerOf(&_IHeroCoreUpgradeable.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.OwnerOf(&_IHeroCoreUpgradeable.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Paused() (bool, error) {
	return _IHeroCoreUpgradeable.Contract.Paused(&_IHeroCoreUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) Paused() (bool, error) {
	return _IHeroCoreUpgradeable.Contract.Paused(&_IHeroCoreUpgradeable.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SaleAuction() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.SaleAuction(&_IHeroCoreUpgradeable.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) SaleAuction() (common.Address, error) {
	return _IHeroCoreUpgradeable.Contract.SaleAuction(&_IHeroCoreUpgradeable.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.SupportsInterface(&_IHeroCoreUpgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IHeroCoreUpgradeable.Contract.SupportsInterface(&_IHeroCoreUpgradeable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Symbol() (string, error) {
	return _IHeroCoreUpgradeable.Contract.Symbol(&_IHeroCoreUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) Symbol() (string, error) {
	return _IHeroCoreUpgradeable.Contract.Symbol(&_IHeroCoreUpgradeable.CallOpts)
}

// TimePerStamina is a free data retrieval call binding the contract method 0x96072223.
//
// Solidity: function timePerStamina() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) TimePerStamina(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "timePerStamina")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimePerStamina is a free data retrieval call binding the contract method 0x96072223.
//
// Solidity: function timePerStamina() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) TimePerStamina() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TimePerStamina(&_IHeroCoreUpgradeable.CallOpts)
}

// TimePerStamina is a free data retrieval call binding the contract method 0x96072223.
//
// Solidity: function timePerStamina() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) TimePerStamina() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TimePerStamina(&_IHeroCoreUpgradeable.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TokenByIndex(&_IHeroCoreUpgradeable.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TokenByIndex(&_IHeroCoreUpgradeable.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TokenOfOwnerByIndex(&_IHeroCoreUpgradeable.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TokenOfOwnerByIndex(&_IHeroCoreUpgradeable.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IHeroCoreUpgradeable.Contract.TokenURI(&_IHeroCoreUpgradeable.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _IHeroCoreUpgradeable.Contract.TokenURI(&_IHeroCoreUpgradeable.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TotalSupply(&_IHeroCoreUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _IHeroCoreUpgradeable.Contract.TotalSupply(&_IHeroCoreUpgradeable.CallOpts)
}

// Vrf is a free data retrieval call binding the contract method 0x4b757f98.
//
// Solidity: function vrf(uint256 blockNumber) view returns(bytes32 result)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCaller) Vrf(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _IHeroCoreUpgradeable.contract.Call(opts, &out, "vrf", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Vrf is a free data retrieval call binding the contract method 0x4b757f98.
//
// Solidity: function vrf(uint256 blockNumber) view returns(bytes32 result)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Vrf(blockNumber *big.Int) ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.Vrf(&_IHeroCoreUpgradeable.CallOpts, blockNumber)
}

// Vrf is a free data retrieval call binding the contract method 0x4b757f98.
//
// Solidity: function vrf(uint256 blockNumber) view returns(bytes32 result)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableCallerSession) Vrf(blockNumber *big.Int) ([32]byte, error) {
	return _IHeroCoreUpgradeable.Contract.Vrf(&_IHeroCoreUpgradeable.CallOpts, blockNumber)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Approve(&_IHeroCoreUpgradeable.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Approve(&_IHeroCoreUpgradeable.TransactOpts, to, tokenId)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x1b827671.
//
// Solidity: function bridgeMint(uint256 _id, address _to) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) BridgeMint(opts *bind.TransactOpts, _id *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "bridgeMint", _id, _to)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x1b827671.
//
// Solidity: function bridgeMint(uint256 _id, address _to) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) BridgeMint(_id *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.BridgeMint(&_IHeroCoreUpgradeable.TransactOpts, _id, _to)
}

// BridgeMint is a paid mutator transaction binding the contract method 0x1b827671.
//
// Solidity: function bridgeMint(uint256 _id, address _to) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) BridgeMint(_id *big.Int, _to common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.BridgeMint(&_IHeroCoreUpgradeable.TransactOpts, _id, _to)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Burn(&_IHeroCoreUpgradeable.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Burn(&_IHeroCoreUpgradeable.TransactOpts, tokenId)
}

// CreateAssistingAuction is a paid mutator transaction binding the contract method 0x32847f0e.
//
// Solidity: function createAssistingAuction(uint256 _heroId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) CreateAssistingAuction(opts *bind.TransactOpts, _heroId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "createAssistingAuction", _heroId, _startingPrice, _endingPrice, _duration)
}

// CreateAssistingAuction is a paid mutator transaction binding the contract method 0x32847f0e.
//
// Solidity: function createAssistingAuction(uint256 _heroId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) CreateAssistingAuction(_heroId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.CreateAssistingAuction(&_IHeroCoreUpgradeable.TransactOpts, _heroId, _startingPrice, _endingPrice, _duration)
}

// CreateAssistingAuction is a paid mutator transaction binding the contract method 0x32847f0e.
//
// Solidity: function createAssistingAuction(uint256 _heroId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) CreateAssistingAuction(_heroId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.CreateAssistingAuction(&_IHeroCoreUpgradeable.TransactOpts, _heroId, _startingPrice, _endingPrice, _duration)
}

// CreateHero is a paid mutator transaction binding the contract method 0x306e56f3.
//
// Solidity: function createHero(uint256 _statGenes, uint256 _visualGenes, uint8 _rarity, bool _shiny, (address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8) _crystal, uint256 _crystalId) returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) CreateHero(opts *bind.TransactOpts, _statGenes *big.Int, _visualGenes *big.Int, _rarity uint8, _shiny bool, _crystal HeroCrystal, _crystalId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "createHero", _statGenes, _visualGenes, _rarity, _shiny, _crystal, _crystalId)
}

// CreateHero is a paid mutator transaction binding the contract method 0x306e56f3.
//
// Solidity: function createHero(uint256 _statGenes, uint256 _visualGenes, uint8 _rarity, bool _shiny, (address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8) _crystal, uint256 _crystalId) returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) CreateHero(_statGenes *big.Int, _visualGenes *big.Int, _rarity uint8, _shiny bool, _crystal HeroCrystal, _crystalId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.CreateHero(&_IHeroCoreUpgradeable.TransactOpts, _statGenes, _visualGenes, _rarity, _shiny, _crystal, _crystalId)
}

// CreateHero is a paid mutator transaction binding the contract method 0x306e56f3.
//
// Solidity: function createHero(uint256 _statGenes, uint256 _visualGenes, uint8 _rarity, bool _shiny, (address,uint256,uint256,uint16,uint256,uint256,uint8,uint8,address,uint32,uint32,uint32,uint8) _crystal, uint256 _crystalId) returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) CreateHero(_statGenes *big.Int, _visualGenes *big.Int, _rarity uint8, _shiny bool, _crystal HeroCrystal, _crystalId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.CreateHero(&_IHeroCoreUpgradeable.TransactOpts, _statGenes, _visualGenes, _rarity, _shiny, _crystal, _crystalId)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _heroId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) CreateSaleAuction(opts *bind.TransactOpts, _heroId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "createSaleAuction", _heroId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _heroId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) CreateSaleAuction(_heroId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.CreateSaleAuction(&_IHeroCoreUpgradeable.TransactOpts, _heroId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _heroId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) CreateSaleAuction(_heroId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.CreateSaleAuction(&_IHeroCoreUpgradeable.TransactOpts, _heroId, _startingPrice, _endingPrice, _duration)
}

// DeductStamina is a paid mutator transaction binding the contract method 0xf2e58229.
//
// Solidity: function deductStamina(uint256 _heroId, uint256 _staminaDeduction) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) DeductStamina(opts *bind.TransactOpts, _heroId *big.Int, _staminaDeduction *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "deductStamina", _heroId, _staminaDeduction)
}

// DeductStamina is a paid mutator transaction binding the contract method 0xf2e58229.
//
// Solidity: function deductStamina(uint256 _heroId, uint256 _staminaDeduction) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) DeductStamina(_heroId *big.Int, _staminaDeduction *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.DeductStamina(&_IHeroCoreUpgradeable.TransactOpts, _heroId, _staminaDeduction)
}

// DeductStamina is a paid mutator transaction binding the contract method 0xf2e58229.
//
// Solidity: function deductStamina(uint256 _heroId, uint256 _staminaDeduction) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) DeductStamina(_heroId *big.Int, _staminaDeduction *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.DeductStamina(&_IHeroCoreUpgradeable.TransactOpts, _heroId, _staminaDeduction)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.GrantRole(&_IHeroCoreUpgradeable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.GrantRole(&_IHeroCoreUpgradeable.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string baseTokenURI) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, baseTokenURI string) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "initialize", name, symbol, baseTokenURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string baseTokenURI) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Initialize(name string, symbol string, baseTokenURI string) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Initialize(&_IHeroCoreUpgradeable.TransactOpts, name, symbol, baseTokenURI)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6487c53.
//
// Solidity: function initialize(string name, string symbol, string baseTokenURI) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) Initialize(name string, symbol string, baseTokenURI string) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Initialize(&_IHeroCoreUpgradeable.TransactOpts, name, symbol, baseTokenURI)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _crystalAddress) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) Initialize0(opts *bind.TransactOpts, _crystalAddress common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "initialize0", _crystalAddress)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _crystalAddress) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Initialize0(_crystalAddress common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Initialize0(&_IHeroCoreUpgradeable.TransactOpts, _crystalAddress)
}

// Initialize0 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _crystalAddress) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) Initialize0(_crystalAddress common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Initialize0(&_IHeroCoreUpgradeable.TransactOpts, _crystalAddress)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Mint(&_IHeroCoreUpgradeable.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Mint(&_IHeroCoreUpgradeable.TransactOpts, to)
}

// OpenCrystal is a paid mutator transaction binding the contract method 0x8ff41141.
//
// Solidity: function openCrystal(uint256 _crystalId) returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) OpenCrystal(opts *bind.TransactOpts, _crystalId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "openCrystal", _crystalId)
}

// OpenCrystal is a paid mutator transaction binding the contract method 0x8ff41141.
//
// Solidity: function openCrystal(uint256 _crystalId) returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) OpenCrystal(_crystalId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.OpenCrystal(&_IHeroCoreUpgradeable.TransactOpts, _crystalId)
}

// OpenCrystal is a paid mutator transaction binding the contract method 0x8ff41141.
//
// Solidity: function openCrystal(uint256 _crystalId) returns(uint256)
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) OpenCrystal(_crystalId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.OpenCrystal(&_IHeroCoreUpgradeable.TransactOpts, _crystalId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Pause() (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Pause(&_IHeroCoreUpgradeable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) Pause() (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Pause(&_IHeroCoreUpgradeable.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.RenounceRole(&_IHeroCoreUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.RenounceRole(&_IHeroCoreUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.RevokeRole(&_IHeroCoreUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.RevokeRole(&_IHeroCoreUpgradeable.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SafeTransferFrom(&_IHeroCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SafeTransferFrom(&_IHeroCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SafeTransferFrom0(&_IHeroCoreUpgradeable.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SafeTransferFrom0(&_IHeroCoreUpgradeable.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetApprovalForAll(&_IHeroCoreUpgradeable.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetApprovalForAll(&_IHeroCoreUpgradeable.TransactOpts, operator, approved)
}

// SetAssistingAuctionAddress is a paid mutator transaction binding the contract method 0x5c9c7c73.
//
// Solidity: function setAssistingAuctionAddress(address _address) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SetAssistingAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "setAssistingAuctionAddress", _address)
}

// SetAssistingAuctionAddress is a paid mutator transaction binding the contract method 0x5c9c7c73.
//
// Solidity: function setAssistingAuctionAddress(address _address) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SetAssistingAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetAssistingAuctionAddress(&_IHeroCoreUpgradeable.TransactOpts, _address)
}

// SetAssistingAuctionAddress is a paid mutator transaction binding the contract method 0x5c9c7c73.
//
// Solidity: function setAssistingAuctionAddress(address _address) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SetAssistingAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetAssistingAuctionAddress(&_IHeroCoreUpgradeable.TransactOpts, _address)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SetFees(opts *bind.TransactOpts, _feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "setFees", _feeAddresses, _feePercents)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SetFees(_feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetFees(&_IHeroCoreUpgradeable.TransactOpts, _feeAddresses, _feePercents)
}

// SetFees is a paid mutator transaction binding the contract method 0x16f81524.
//
// Solidity: function setFees(address[] _feeAddresses, uint256[] _feePercents) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SetFees(_feeAddresses []common.Address, _feePercents []*big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetFees(&_IHeroCoreUpgradeable.TransactOpts, _feeAddresses, _feePercents)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetSaleAuctionAddress(&_IHeroCoreUpgradeable.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetSaleAuctionAddress(&_IHeroCoreUpgradeable.TransactOpts, _address)
}

// SetSummonCooldowns is a paid mutator transaction binding the contract method 0x4e970324.
//
// Solidity: function setSummonCooldowns(uint256 _baseCooldown, uint256 _cooldownPerSummon, uint256 _cooldownPerGen) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SetSummonCooldowns(opts *bind.TransactOpts, _baseCooldown *big.Int, _cooldownPerSummon *big.Int, _cooldownPerGen *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "setSummonCooldowns", _baseCooldown, _cooldownPerSummon, _cooldownPerGen)
}

// SetSummonCooldowns is a paid mutator transaction binding the contract method 0x4e970324.
//
// Solidity: function setSummonCooldowns(uint256 _baseCooldown, uint256 _cooldownPerSummon, uint256 _cooldownPerGen) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SetSummonCooldowns(_baseCooldown *big.Int, _cooldownPerSummon *big.Int, _cooldownPerGen *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetSummonCooldowns(&_IHeroCoreUpgradeable.TransactOpts, _baseCooldown, _cooldownPerSummon, _cooldownPerGen)
}

// SetSummonCooldowns is a paid mutator transaction binding the contract method 0x4e970324.
//
// Solidity: function setSummonCooldowns(uint256 _baseCooldown, uint256 _cooldownPerSummon, uint256 _cooldownPerGen) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SetSummonCooldowns(_baseCooldown *big.Int, _cooldownPerSummon *big.Int, _cooldownPerGen *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetSummonCooldowns(&_IHeroCoreUpgradeable.TransactOpts, _baseCooldown, _cooldownPerSummon, _cooldownPerGen)
}

// SetSummonFees is a paid mutator transaction binding the contract method 0x03466cfd.
//
// Solidity: function setSummonFees(uint256 _baseSummonFee, uint256 _increasePerSummon, uint256 _increasePerGen) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SetSummonFees(opts *bind.TransactOpts, _baseSummonFee *big.Int, _increasePerSummon *big.Int, _increasePerGen *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "setSummonFees", _baseSummonFee, _increasePerSummon, _increasePerGen)
}

// SetSummonFees is a paid mutator transaction binding the contract method 0x03466cfd.
//
// Solidity: function setSummonFees(uint256 _baseSummonFee, uint256 _increasePerSummon, uint256 _increasePerGen) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SetSummonFees(_baseSummonFee *big.Int, _increasePerSummon *big.Int, _increasePerGen *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetSummonFees(&_IHeroCoreUpgradeable.TransactOpts, _baseSummonFee, _increasePerSummon, _increasePerGen)
}

// SetSummonFees is a paid mutator transaction binding the contract method 0x03466cfd.
//
// Solidity: function setSummonFees(uint256 _baseSummonFee, uint256 _increasePerSummon, uint256 _increasePerGen) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SetSummonFees(_baseSummonFee *big.Int, _increasePerSummon *big.Int, _increasePerGen *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetSummonFees(&_IHeroCoreUpgradeable.TransactOpts, _baseSummonFee, _increasePerSummon, _increasePerGen)
}

// SetTimePerStamina is a paid mutator transaction binding the contract method 0x7fd73850.
//
// Solidity: function setTimePerStamina(uint256 _timePerStamina) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SetTimePerStamina(opts *bind.TransactOpts, _timePerStamina *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "setTimePerStamina", _timePerStamina)
}

// SetTimePerStamina is a paid mutator transaction binding the contract method 0x7fd73850.
//
// Solidity: function setTimePerStamina(uint256 _timePerStamina) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SetTimePerStamina(_timePerStamina *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetTimePerStamina(&_IHeroCoreUpgradeable.TransactOpts, _timePerStamina)
}

// SetTimePerStamina is a paid mutator transaction binding the contract method 0x7fd73850.
//
// Solidity: function setTimePerStamina(uint256 _timePerStamina) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SetTimePerStamina(_timePerStamina *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SetTimePerStamina(&_IHeroCoreUpgradeable.TransactOpts, _timePerStamina)
}

// SummonCrystal is a paid mutator transaction binding the contract method 0x5880d8e6.
//
// Solidity: function summonCrystal(uint256 _summonerId, uint256 _assistantId, uint8 _summonerTears, uint8 _assistantTears, address _enhancementStone) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) SummonCrystal(opts *bind.TransactOpts, _summonerId *big.Int, _assistantId *big.Int, _summonerTears uint8, _assistantTears uint8, _enhancementStone common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "summonCrystal", _summonerId, _assistantId, _summonerTears, _assistantTears, _enhancementStone)
}

// SummonCrystal is a paid mutator transaction binding the contract method 0x5880d8e6.
//
// Solidity: function summonCrystal(uint256 _summonerId, uint256 _assistantId, uint8 _summonerTears, uint8 _assistantTears, address _enhancementStone) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) SummonCrystal(_summonerId *big.Int, _assistantId *big.Int, _summonerTears uint8, _assistantTears uint8, _enhancementStone common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SummonCrystal(&_IHeroCoreUpgradeable.TransactOpts, _summonerId, _assistantId, _summonerTears, _assistantTears, _enhancementStone)
}

// SummonCrystal is a paid mutator transaction binding the contract method 0x5880d8e6.
//
// Solidity: function summonCrystal(uint256 _summonerId, uint256 _assistantId, uint8 _summonerTears, uint8 _assistantTears, address _enhancementStone) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) SummonCrystal(_summonerId *big.Int, _assistantId *big.Int, _summonerTears uint8, _assistantTears uint8, _enhancementStone common.Address) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.SummonCrystal(&_IHeroCoreUpgradeable.TransactOpts, _summonerId, _assistantId, _summonerTears, _assistantTears, _enhancementStone)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.TransferFrom(&_IHeroCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.TransferFrom(&_IHeroCoreUpgradeable.TransactOpts, from, to, tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) Unpause() (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Unpause(&_IHeroCoreUpgradeable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) Unpause() (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.Unpause(&_IHeroCoreUpgradeable.TransactOpts)
}

// UpdateHero is a paid mutator transaction binding the contract method 0xb0064103.
//
// Solidity: function updateHero((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactor) UpdateHero(opts *bind.TransactOpts, _hero Hero) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.contract.Transact(opts, "updateHero", _hero)
}

// UpdateHero is a paid mutator transaction binding the contract method 0xb0064103.
//
// Solidity: function updateHero((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableSession) UpdateHero(_hero Hero) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.UpdateHero(&_IHeroCoreUpgradeable.TransactOpts, _hero)
}

// UpdateHero is a paid mutator transaction binding the contract method 0xb0064103.
//
// Solidity: function updateHero((uint256,(uint256,uint256,uint256,uint256,uint32,uint32),(uint256,uint256,uint8,bool,uint16,uint32,uint32,uint8,uint8,uint8),(uint256,uint256,uint256,uint16,uint64,address,uint8,uint8),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16,uint16),(uint16,uint16,uint16,uint16)) _hero) returns()
func (_IHeroCoreUpgradeable *IHeroCoreUpgradeableTransactorSession) UpdateHero(_hero Hero) (*types.Transaction, error) {
	return _IHeroCoreUpgradeable.Contract.UpdateHero(&_IHeroCoreUpgradeable.TransactOpts, _hero)
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
