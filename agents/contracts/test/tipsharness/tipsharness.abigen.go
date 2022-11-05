// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tipsharness

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

// SynapseTypesMetaData contains all meta data concerning the SynapseTypes contract.
var SynapseTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bd8a12c786b54003a475ba8cd5eb537ff6b83443c707f08a29ca81a9d6cb071d64736f6c63430008110033",
}

// SynapseTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseTypesMetaData.ABI instead.
var SynapseTypesABI = SynapseTypesMetaData.ABI

// SynapseTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseTypesMetaData.Bin instead.
var SynapseTypesBin = SynapseTypesMetaData.Bin

// DeploySynapseTypes deploys a new Ethereum contract, binding an instance of SynapseTypes to it.
func DeploySynapseTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseTypes, error) {
	parsed, err := SynapseTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseTypes{SynapseTypesCaller: SynapseTypesCaller{contract: contract}, SynapseTypesTransactor: SynapseTypesTransactor{contract: contract}, SynapseTypesFilterer: SynapseTypesFilterer{contract: contract}}, nil
}

// SynapseTypes is an auto generated Go binding around an Ethereum contract.
type SynapseTypes struct {
	SynapseTypesCaller     // Read-only binding to the contract
	SynapseTypesTransactor // Write-only binding to the contract
	SynapseTypesFilterer   // Log filterer for contract events
}

// SynapseTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseTypesSession struct {
	Contract     *SynapseTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseTypesCallerSession struct {
	Contract *SynapseTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SynapseTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseTypesTransactorSession struct {
	Contract     *SynapseTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SynapseTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseTypesRaw struct {
	Contract *SynapseTypes // Generic contract binding to access the raw methods on
}

// SynapseTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseTypesCallerRaw struct {
	Contract *SynapseTypesCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseTypesTransactorRaw struct {
	Contract *SynapseTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseTypes creates a new instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypes(address common.Address, backend bind.ContractBackend) (*SynapseTypes, error) {
	contract, err := bindSynapseTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseTypes{SynapseTypesCaller: SynapseTypesCaller{contract: contract}, SynapseTypesTransactor: SynapseTypesTransactor{contract: contract}, SynapseTypesFilterer: SynapseTypesFilterer{contract: contract}}, nil
}

// NewSynapseTypesCaller creates a new read-only instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesCaller(address common.Address, caller bind.ContractCaller) (*SynapseTypesCaller, error) {
	contract, err := bindSynapseTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesCaller{contract: contract}, nil
}

// NewSynapseTypesTransactor creates a new write-only instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseTypesTransactor, error) {
	contract, err := bindSynapseTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesTransactor{contract: contract}, nil
}

// NewSynapseTypesFilterer creates a new log filterer instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseTypesFilterer, error) {
	contract, err := bindSynapseTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesFilterer{contract: contract}, nil
}

// bindSynapseTypes binds a generic wrapper to an already deployed contract.
func bindSynapseTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseTypes *SynapseTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseTypes.Contract.SynapseTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseTypes *SynapseTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseTypes.Contract.SynapseTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseTypes *SynapseTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseTypes.Contract.SynapseTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseTypes *SynapseTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseTypes *SynapseTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseTypes *SynapseTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseTypes.Contract.contract.Transact(opts, method, params...)
}

// TipsMetaData contains all meta data concerning the Tips contract.
var TipsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a5bb92aa7f9a6c114be49400593121afa76aaa7950788d328abf525f67d9302964736f6c63430008110033",
}

// TipsABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsMetaData.ABI instead.
var TipsABI = TipsMetaData.ABI

// TipsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsMetaData.Bin instead.
var TipsBin = TipsMetaData.Bin

// DeployTips deploys a new Ethereum contract, binding an instance of Tips to it.
func DeployTips(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Tips, error) {
	parsed, err := TipsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Tips{TipsCaller: TipsCaller{contract: contract}, TipsTransactor: TipsTransactor{contract: contract}, TipsFilterer: TipsFilterer{contract: contract}}, nil
}

// Tips is an auto generated Go binding around an Ethereum contract.
type Tips struct {
	TipsCaller     // Read-only binding to the contract
	TipsTransactor // Write-only binding to the contract
	TipsFilterer   // Log filterer for contract events
}

// TipsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsSession struct {
	Contract     *Tips             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsCallerSession struct {
	Contract *TipsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TipsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsTransactorSession struct {
	Contract     *TipsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsRaw struct {
	Contract *Tips // Generic contract binding to access the raw methods on
}

// TipsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsCallerRaw struct {
	Contract *TipsCaller // Generic read-only contract binding to access the raw methods on
}

// TipsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsTransactorRaw struct {
	Contract *TipsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTips creates a new instance of Tips, bound to a specific deployed contract.
func NewTips(address common.Address, backend bind.ContractBackend) (*Tips, error) {
	contract, err := bindTips(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tips{TipsCaller: TipsCaller{contract: contract}, TipsTransactor: TipsTransactor{contract: contract}, TipsFilterer: TipsFilterer{contract: contract}}, nil
}

// NewTipsCaller creates a new read-only instance of Tips, bound to a specific deployed contract.
func NewTipsCaller(address common.Address, caller bind.ContractCaller) (*TipsCaller, error) {
	contract, err := bindTips(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsCaller{contract: contract}, nil
}

// NewTipsTransactor creates a new write-only instance of Tips, bound to a specific deployed contract.
func NewTipsTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsTransactor, error) {
	contract, err := bindTips(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsTransactor{contract: contract}, nil
}

// NewTipsFilterer creates a new log filterer instance of Tips, bound to a specific deployed contract.
func NewTipsFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsFilterer, error) {
	contract, err := bindTips(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsFilterer{contract: contract}, nil
}

// bindTips binds a generic wrapper to an already deployed contract.
func bindTips(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tips *TipsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tips.Contract.TipsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tips *TipsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tips.Contract.TipsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tips *TipsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tips.Contract.TipsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tips *TipsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tips.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tips *TipsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tips.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tips *TipsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tips.Contract.contract.Transact(opts, method, params...)
}

// TipsHarnessMetaData contains all meta data concerning the TipsHarness contract.
var TipsHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"broadcasterTip\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"castToTips\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emptyTips\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"executorTip\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_notaryTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_broadcasterTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_proverTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_executorTip\",\"type\":\"uint96\"}],\"name\":\"formatTips\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"isTips\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"notaryTip\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetBroadcaster\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetExecutor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetNotary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetProver\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"proverTip\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipsVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"tipsVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"totalTips\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"54fbddad": "broadcasterTip(uint40,bytes)",
		"8435d5ad": "castToTips(uint40,bytes)",
		"725bd463": "emptyTips()",
		"82739aa0": "executorTip(uint40,bytes)",
		"d024f867": "formatTips(uint96,uint96,uint96,uint96)",
		"993abc41": "isTips(bytes)",
		"fc39e482": "notaryTip(uint40,bytes)",
		"15bb7d2b": "offsetBroadcaster()",
		"51970c3f": "offsetExecutor()",
		"b4b4ccb2": "offsetNotary()",
		"98de8554": "offsetProver()",
		"0c096e8d": "offsetVersion()",
		"7b201de6": "proverTip(uint40,bytes)",
		"b440592e": "tipsLength()",
		"60fb5709": "tipsVersion()",
		"ecfa57cd": "tipsVersion(uint40,bytes)",
		"49adcc6a": "totalTips(uint40,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611341806100206000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c806382739aa0116100b2578063b440592e11610081578063d024f86711610066578063d024f86714610231578063ecfa57cd14610244578063fc39e4821461025757600080fd5b8063b440592e14610223578063b4b4ccb21461022a57600080fd5b806382739aa0146101c55780638435d5ad146101d857806398de8554146101f9578063993abc411461020057600080fd5b806354fbddad116100ee57806354fbddad1461017457806360fb570914610187578063725bd4631461019d5780637b201de6146101b257600080fd5b80630c096e8d1461012057806315bb7d2b1461013657806349adcc6a1461013d57806351970c3f1461016d575b600080fd5b60005b6040519081526020015b60405180910390f35b600e610123565b61015061014b366004610f88565b61026a565b6040516bffffffffffffffffffffffff909116815260200161012d565b6026610123565b610150610182366004610f88565b61028d565b60015b60405161ffff909116815260200161012d565b6101a56102a7565b60405161012d9190611046565b6101506101c0366004610f88565b61030b565b6101506101d3366004610f88565b610325565b6101eb6101e6366004610f88565b61033f565b60405161012d929190611059565b601a610123565b61021361020e366004611081565b610370565b604051901515815260200161012d565b6032610123565b6002610123565b6101a561023f3660046110d7565b610389565b61018a610252366004610f88565b61041b565b610150610265366004610f88565b610435565b6000610284610279838561044f565b62ffffff191661046a565b90505b92915050565b600061028461029c838561044f565b62ffffff19166104ae565b6060610306604080517e010000000000000000000000000000000000000000000000000000000000006020820152600060228201819052602e8201819052603a8201819052604682015281518082036032018152605290910190915290565b905090565b600061028461031a838561044f565b62ffffff19166104e3565b6000610284610334838561044f565b62ffffff191661050f565b60006060600061034e8461053b565b905060d881901c61036462ffffff19831661054c565b92509250509250929050565b600061028761037e8361053b565b62ffffff191661059f565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffffffffffff000000000000000000000000000000000000000060a087811b8216602284015286811b8216602e84015285811b8216603a84015284901b1660468201528151808203603201815260529091019091526060905b95945050505050565b600061028461042a838561044f565b62ffffff19166105e6565b6000610284610444838561044f565b62ffffff1916610612565b81516000906020840161041264ffffffffff8516828461063e565b60006104758261050f565b61047e836104e3565b610487846104ae565b61049085610612565b61049a919061115a565b6104a4919061115a565b610287919061115a565b6000816104c662ffffff198216640301020000610688565b506104da62ffffff198416600e600c6107ac565b91505b50919050565b6000816104fb62ffffff198216640301020000610688565b506104da62ffffff198416601a600c6107ac565b60008161052762ffffff198216640301020000610688565b506104da62ffffff1984166026600c6107ac565b60006102878264030102000061044f565b60606000806105698460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061058e84836020016107dc565b508181016020016040529052919050565b6000601882901c6bffffffffffffffffffffffff1660028110156105c65750600092915050565b60016105d1846105e6565b61ffff161480156104da575060321492915050565b6000816105fe62ffffff198216640301020000610688565b506104da62ffffff198416600060026107ac565b60008161062a62ffffff198216640301020000610688565b506104da62ffffff1984166002600c6107ac565b60008061064b8385611186565b905060405181111561065b575060005b806000036106705762ffffff19915050610681565b5050606083811b8317901b811760181b5b9392505050565b60006106948383610969565b6107a55760006106b36106a78560d81c90565b64ffffffffff1661098c565b91505060006106c88464ffffffffff1661098c565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c9190611046565b60405180910390fd5b5090919050565b60006107b9826020611199565b6107c49060086111b2565b60ff166107d2858585610a76565b901c949350505050565b600062ffffff198084160361084d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e746572206465726566000000000000604482015260640161079c565b61085683610c42565b6108bc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e746572206465726566000000604482015260640161079c565b60006108d68460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006109008560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006040519050848111156109255760206060fd5b8285848460045afa5061095f61093b8760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b60008164ffffffffff1661097d8460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff1611156109ff5760006109ab8260086111b2565b60ff1685901c90506109bc81610c7f565b61ffff16841793508160ff166010146109d757601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610992565b50600f5b60ff8160ff161015610a70576000610a1c8260086111b2565b60ff1685901c9050610a2d81610c7f565b61ffff16831792508160ff16600014610a4857601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610a03565b50915091565b60008160ff16600003610a8b57506000610681565b610aa38460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610abe60ff841685611186565b1115610b5057610b1d610adf8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610b058660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16610cb1565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161079c9190611046565b60208260ff161115610bbe576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e20333220627974657300000000000000604482015260640161079c565b600882026000610bdc8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000610c4e8260d81c90565b64ffffffffff1664ffffffffff03610c6857506000919050565b6000610c7383610d1f565b60405110199392505050565b6000610c9160048360ff16901c610d67565b60ff1661ffff919091161760081b610ca882610d67565b60ff1617919050565b60606000610cbe8661098c565b9150506000610ccc8661098c565b9150506000610cda8661098c565b9150506000610ce88661098c565b91505083838383604051602001610d0294939291906111ce565b604051602081830303815290604052945050505050949350505050565b6000610d398260181c6bffffffffffffffffffffffff1690565b610d518360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b600060f08083179060ff82169003610d825750603092915050565b8060ff1660f103610d965750603192915050565b8060ff1660f203610daa5750603292915050565b8060ff1660f303610dbe5750603392915050565b8060ff1660f403610dd25750603492915050565b8060ff1660f503610de65750603592915050565b8060ff1660f603610dfa5750603692915050565b8060ff1660f703610e0e5750603792915050565b8060ff1660f803610e225750603892915050565b8060ff1660f903610e365750603992915050565b8060ff1660fa03610e4a5750606192915050565b8060ff1660fb03610e5e5750606292915050565b8060ff1660fc03610e725750606392915050565b8060ff1660fd03610e865750606492915050565b8060ff1660fe03610e9a5750606592915050565b8060ff1660ff036104dd5750606692915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610eee57600080fd5b813567ffffffffffffffff80821115610f0957610f09610eae565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610f4f57610f4f610eae565b81604052838152866020858801011115610f6857600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610f9b57600080fd5b823564ffffffffff81168114610fb057600080fd5b9150602083013567ffffffffffffffff811115610fcc57600080fd5b610fd885828601610edd565b9150509250929050565b6000815180845260005b8181101561100857602081850181015186830182015201610fec565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006102846020830184610fe2565b64ffffffffff831681526040602082015260006110796040830184610fe2565b949350505050565b60006020828403121561109357600080fd5b813567ffffffffffffffff8111156110aa57600080fd5b61107984828501610edd565b80356bffffffffffffffffffffffff811681146110d257600080fd5b919050565b600080600080608085870312156110ed57600080fd5b6110f6856110b6565b9350611104602086016110b6565b9250611112604086016110b6565b9150611120606086016110b6565b905092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6bffffffffffffffffffffffff81811683821601908082111561117f5761117f61112b565b5092915050565b808201808211156102875761028761112b565b60ff82811682821603908111156102875761028761112b565b60ff818116838216029081169081811461117f5761117f61112b565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d820161095f56fea26469706673582212201984e8e31ca827280300f48db462ec4e3d9c5573a22618282d28b0f2a6c1174e64736f6c63430008110033",
}

// TipsHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsHarnessMetaData.ABI instead.
var TipsHarnessABI = TipsHarnessMetaData.ABI

// Deprecated: Use TipsHarnessMetaData.Sigs instead.
// TipsHarnessFuncSigs maps the 4-byte function signature to its string representation.
var TipsHarnessFuncSigs = TipsHarnessMetaData.Sigs

// TipsHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsHarnessMetaData.Bin instead.
var TipsHarnessBin = TipsHarnessMetaData.Bin

// DeployTipsHarness deploys a new Ethereum contract, binding an instance of TipsHarness to it.
func DeployTipsHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TipsHarness, error) {
	parsed, err := TipsHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TipsHarness{TipsHarnessCaller: TipsHarnessCaller{contract: contract}, TipsHarnessTransactor: TipsHarnessTransactor{contract: contract}, TipsHarnessFilterer: TipsHarnessFilterer{contract: contract}}, nil
}

// TipsHarness is an auto generated Go binding around an Ethereum contract.
type TipsHarness struct {
	TipsHarnessCaller     // Read-only binding to the contract
	TipsHarnessTransactor // Write-only binding to the contract
	TipsHarnessFilterer   // Log filterer for contract events
}

// TipsHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsHarnessSession struct {
	Contract     *TipsHarness      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsHarnessCallerSession struct {
	Contract *TipsHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TipsHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsHarnessTransactorSession struct {
	Contract     *TipsHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TipsHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsHarnessRaw struct {
	Contract *TipsHarness // Generic contract binding to access the raw methods on
}

// TipsHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsHarnessCallerRaw struct {
	Contract *TipsHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// TipsHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsHarnessTransactorRaw struct {
	Contract *TipsHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipsHarness creates a new instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarness(address common.Address, backend bind.ContractBackend) (*TipsHarness, error) {
	contract, err := bindTipsHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TipsHarness{TipsHarnessCaller: TipsHarnessCaller{contract: contract}, TipsHarnessTransactor: TipsHarnessTransactor{contract: contract}, TipsHarnessFilterer: TipsHarnessFilterer{contract: contract}}, nil
}

// NewTipsHarnessCaller creates a new read-only instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarnessCaller(address common.Address, caller bind.ContractCaller) (*TipsHarnessCaller, error) {
	contract, err := bindTipsHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsHarnessCaller{contract: contract}, nil
}

// NewTipsHarnessTransactor creates a new write-only instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsHarnessTransactor, error) {
	contract, err := bindTipsHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsHarnessTransactor{contract: contract}, nil
}

// NewTipsHarnessFilterer creates a new log filterer instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsHarnessFilterer, error) {
	contract, err := bindTipsHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsHarnessFilterer{contract: contract}, nil
}

// bindTipsHarness binds a generic wrapper to an already deployed contract.
func bindTipsHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsHarness *TipsHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsHarness.Contract.TipsHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsHarness *TipsHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsHarness.Contract.TipsHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsHarness *TipsHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsHarness.Contract.TipsHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsHarness *TipsHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsHarness *TipsHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsHarness *TipsHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsHarness.Contract.contract.Transact(opts, method, params...)
}

// BroadcasterTip is a free data retrieval call binding the contract method 0x54fbddad.
//
// Solidity: function broadcasterTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCaller) BroadcasterTip(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "broadcasterTip", _type, _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BroadcasterTip is a free data retrieval call binding the contract method 0x54fbddad.
//
// Solidity: function broadcasterTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessSession) BroadcasterTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.BroadcasterTip(&_TipsHarness.CallOpts, _type, _payload)
}

// BroadcasterTip is a free data retrieval call binding the contract method 0x54fbddad.
//
// Solidity: function broadcasterTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCallerSession) BroadcasterTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.BroadcasterTip(&_TipsHarness.CallOpts, _type, _payload)
}

// CastToTips is a free data retrieval call binding the contract method 0x8435d5ad.
//
// Solidity: function castToTips(uint40 , bytes _payload) view returns(uint40, bytes)
func (_TipsHarness *TipsHarnessCaller) CastToTips(opts *bind.CallOpts, arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "castToTips", arg0, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// CastToTips is a free data retrieval call binding the contract method 0x8435d5ad.
//
// Solidity: function castToTips(uint40 , bytes _payload) view returns(uint40, bytes)
func (_TipsHarness *TipsHarnessSession) CastToTips(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _TipsHarness.Contract.CastToTips(&_TipsHarness.CallOpts, arg0, _payload)
}

// CastToTips is a free data retrieval call binding the contract method 0x8435d5ad.
//
// Solidity: function castToTips(uint40 , bytes _payload) view returns(uint40, bytes)
func (_TipsHarness *TipsHarnessCallerSession) CastToTips(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _TipsHarness.Contract.CastToTips(&_TipsHarness.CallOpts, arg0, _payload)
}

// EmptyTips is a free data retrieval call binding the contract method 0x725bd463.
//
// Solidity: function emptyTips() pure returns(bytes)
func (_TipsHarness *TipsHarnessCaller) EmptyTips(opts *bind.CallOpts) ([]byte, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "emptyTips")

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EmptyTips is a free data retrieval call binding the contract method 0x725bd463.
//
// Solidity: function emptyTips() pure returns(bytes)
func (_TipsHarness *TipsHarnessSession) EmptyTips() ([]byte, error) {
	return _TipsHarness.Contract.EmptyTips(&_TipsHarness.CallOpts)
}

// EmptyTips is a free data retrieval call binding the contract method 0x725bd463.
//
// Solidity: function emptyTips() pure returns(bytes)
func (_TipsHarness *TipsHarnessCallerSession) EmptyTips() ([]byte, error) {
	return _TipsHarness.Contract.EmptyTips(&_TipsHarness.CallOpts)
}

// ExecutorTip is a free data retrieval call binding the contract method 0x82739aa0.
//
// Solidity: function executorTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCaller) ExecutorTip(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "executorTip", _type, _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutorTip is a free data retrieval call binding the contract method 0x82739aa0.
//
// Solidity: function executorTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessSession) ExecutorTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.ExecutorTip(&_TipsHarness.CallOpts, _type, _payload)
}

// ExecutorTip is a free data retrieval call binding the contract method 0x82739aa0.
//
// Solidity: function executorTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCallerSession) ExecutorTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.ExecutorTip(&_TipsHarness.CallOpts, _type, _payload)
}

// FormatTips is a free data retrieval call binding the contract method 0xd024f867.
//
// Solidity: function formatTips(uint96 _notaryTip, uint96 _broadcasterTip, uint96 _proverTip, uint96 _executorTip) pure returns(bytes)
func (_TipsHarness *TipsHarnessCaller) FormatTips(opts *bind.CallOpts, _notaryTip *big.Int, _broadcasterTip *big.Int, _proverTip *big.Int, _executorTip *big.Int) ([]byte, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "formatTips", _notaryTip, _broadcasterTip, _proverTip, _executorTip)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatTips is a free data retrieval call binding the contract method 0xd024f867.
//
// Solidity: function formatTips(uint96 _notaryTip, uint96 _broadcasterTip, uint96 _proverTip, uint96 _executorTip) pure returns(bytes)
func (_TipsHarness *TipsHarnessSession) FormatTips(_notaryTip *big.Int, _broadcasterTip *big.Int, _proverTip *big.Int, _executorTip *big.Int) ([]byte, error) {
	return _TipsHarness.Contract.FormatTips(&_TipsHarness.CallOpts, _notaryTip, _broadcasterTip, _proverTip, _executorTip)
}

// FormatTips is a free data retrieval call binding the contract method 0xd024f867.
//
// Solidity: function formatTips(uint96 _notaryTip, uint96 _broadcasterTip, uint96 _proverTip, uint96 _executorTip) pure returns(bytes)
func (_TipsHarness *TipsHarnessCallerSession) FormatTips(_notaryTip *big.Int, _broadcasterTip *big.Int, _proverTip *big.Int, _executorTip *big.Int) ([]byte, error) {
	return _TipsHarness.Contract.FormatTips(&_TipsHarness.CallOpts, _notaryTip, _broadcasterTip, _proverTip, _executorTip)
}

// IsTips is a free data retrieval call binding the contract method 0x993abc41.
//
// Solidity: function isTips(bytes _payload) pure returns(bool)
func (_TipsHarness *TipsHarnessCaller) IsTips(opts *bind.CallOpts, _payload []byte) (bool, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "isTips", _payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTips is a free data retrieval call binding the contract method 0x993abc41.
//
// Solidity: function isTips(bytes _payload) pure returns(bool)
func (_TipsHarness *TipsHarnessSession) IsTips(_payload []byte) (bool, error) {
	return _TipsHarness.Contract.IsTips(&_TipsHarness.CallOpts, _payload)
}

// IsTips is a free data retrieval call binding the contract method 0x993abc41.
//
// Solidity: function isTips(bytes _payload) pure returns(bool)
func (_TipsHarness *TipsHarnessCallerSession) IsTips(_payload []byte) (bool, error) {
	return _TipsHarness.Contract.IsTips(&_TipsHarness.CallOpts, _payload)
}

// NotaryTip is a free data retrieval call binding the contract method 0xfc39e482.
//
// Solidity: function notaryTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCaller) NotaryTip(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "notaryTip", _type, _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NotaryTip is a free data retrieval call binding the contract method 0xfc39e482.
//
// Solidity: function notaryTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessSession) NotaryTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.NotaryTip(&_TipsHarness.CallOpts, _type, _payload)
}

// NotaryTip is a free data retrieval call binding the contract method 0xfc39e482.
//
// Solidity: function notaryTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCallerSession) NotaryTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.NotaryTip(&_TipsHarness.CallOpts, _type, _payload)
}

// OffsetBroadcaster is a free data retrieval call binding the contract method 0x15bb7d2b.
//
// Solidity: function offsetBroadcaster() pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) OffsetBroadcaster(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "offsetBroadcaster")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetBroadcaster is a free data retrieval call binding the contract method 0x15bb7d2b.
//
// Solidity: function offsetBroadcaster() pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) OffsetBroadcaster() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetBroadcaster(&_TipsHarness.CallOpts)
}

// OffsetBroadcaster is a free data retrieval call binding the contract method 0x15bb7d2b.
//
// Solidity: function offsetBroadcaster() pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) OffsetBroadcaster() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetBroadcaster(&_TipsHarness.CallOpts)
}

// OffsetExecutor is a free data retrieval call binding the contract method 0x51970c3f.
//
// Solidity: function offsetExecutor() pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) OffsetExecutor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "offsetExecutor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetExecutor is a free data retrieval call binding the contract method 0x51970c3f.
//
// Solidity: function offsetExecutor() pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) OffsetExecutor() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetExecutor(&_TipsHarness.CallOpts)
}

// OffsetExecutor is a free data retrieval call binding the contract method 0x51970c3f.
//
// Solidity: function offsetExecutor() pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) OffsetExecutor() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetExecutor(&_TipsHarness.CallOpts)
}

// OffsetNotary is a free data retrieval call binding the contract method 0xb4b4ccb2.
//
// Solidity: function offsetNotary() pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) OffsetNotary(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "offsetNotary")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetNotary is a free data retrieval call binding the contract method 0xb4b4ccb2.
//
// Solidity: function offsetNotary() pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) OffsetNotary() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetNotary(&_TipsHarness.CallOpts)
}

// OffsetNotary is a free data retrieval call binding the contract method 0xb4b4ccb2.
//
// Solidity: function offsetNotary() pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) OffsetNotary() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetNotary(&_TipsHarness.CallOpts)
}

// OffsetProver is a free data retrieval call binding the contract method 0x98de8554.
//
// Solidity: function offsetProver() pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) OffsetProver(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "offsetProver")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetProver is a free data retrieval call binding the contract method 0x98de8554.
//
// Solidity: function offsetProver() pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) OffsetProver() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetProver(&_TipsHarness.CallOpts)
}

// OffsetProver is a free data retrieval call binding the contract method 0x98de8554.
//
// Solidity: function offsetProver() pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) OffsetProver() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetProver(&_TipsHarness.CallOpts)
}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) OffsetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "offsetVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) OffsetVersion() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetVersion(&_TipsHarness.CallOpts)
}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) OffsetVersion() (*big.Int, error) {
	return _TipsHarness.Contract.OffsetVersion(&_TipsHarness.CallOpts)
}

// ProverTip is a free data retrieval call binding the contract method 0x7b201de6.
//
// Solidity: function proverTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCaller) ProverTip(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "proverTip", _type, _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProverTip is a free data retrieval call binding the contract method 0x7b201de6.
//
// Solidity: function proverTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessSession) ProverTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.ProverTip(&_TipsHarness.CallOpts, _type, _payload)
}

// ProverTip is a free data retrieval call binding the contract method 0x7b201de6.
//
// Solidity: function proverTip(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCallerSession) ProverTip(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.ProverTip(&_TipsHarness.CallOpts, _type, _payload)
}

// TipsLength is a free data retrieval call binding the contract method 0xb440592e.
//
// Solidity: function tipsLength() pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) TipsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "tipsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TipsLength is a free data retrieval call binding the contract method 0xb440592e.
//
// Solidity: function tipsLength() pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) TipsLength() (*big.Int, error) {
	return _TipsHarness.Contract.TipsLength(&_TipsHarness.CallOpts)
}

// TipsLength is a free data retrieval call binding the contract method 0xb440592e.
//
// Solidity: function tipsLength() pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) TipsLength() (*big.Int, error) {
	return _TipsHarness.Contract.TipsLength(&_TipsHarness.CallOpts)
}

// TipsVersion is a free data retrieval call binding the contract method 0x60fb5709.
//
// Solidity: function tipsVersion() pure returns(uint16)
func (_TipsHarness *TipsHarnessCaller) TipsVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "tipsVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TipsVersion is a free data retrieval call binding the contract method 0x60fb5709.
//
// Solidity: function tipsVersion() pure returns(uint16)
func (_TipsHarness *TipsHarnessSession) TipsVersion() (uint16, error) {
	return _TipsHarness.Contract.TipsVersion(&_TipsHarness.CallOpts)
}

// TipsVersion is a free data retrieval call binding the contract method 0x60fb5709.
//
// Solidity: function tipsVersion() pure returns(uint16)
func (_TipsHarness *TipsHarnessCallerSession) TipsVersion() (uint16, error) {
	return _TipsHarness.Contract.TipsVersion(&_TipsHarness.CallOpts)
}

// TipsVersion0 is a free data retrieval call binding the contract method 0xecfa57cd.
//
// Solidity: function tipsVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_TipsHarness *TipsHarnessCaller) TipsVersion0(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint16, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "tipsVersion0", _type, _payload)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// TipsVersion0 is a free data retrieval call binding the contract method 0xecfa57cd.
//
// Solidity: function tipsVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_TipsHarness *TipsHarnessSession) TipsVersion0(_type *big.Int, _payload []byte) (uint16, error) {
	return _TipsHarness.Contract.TipsVersion0(&_TipsHarness.CallOpts, _type, _payload)
}

// TipsVersion0 is a free data retrieval call binding the contract method 0xecfa57cd.
//
// Solidity: function tipsVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_TipsHarness *TipsHarnessCallerSession) TipsVersion0(_type *big.Int, _payload []byte) (uint16, error) {
	return _TipsHarness.Contract.TipsVersion0(&_TipsHarness.CallOpts, _type, _payload)
}

// TotalTips is a free data retrieval call binding the contract method 0x49adcc6a.
//
// Solidity: function totalTips(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCaller) TotalTips(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "totalTips", _type, _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTips is a free data retrieval call binding the contract method 0x49adcc6a.
//
// Solidity: function totalTips(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessSession) TotalTips(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.TotalTips(&_TipsHarness.CallOpts, _type, _payload)
}

// TotalTips is a free data retrieval call binding the contract method 0x49adcc6a.
//
// Solidity: function totalTips(uint40 _type, bytes _payload) pure returns(uint96)
func (_TipsHarness *TipsHarnessCallerSession) TotalTips(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _TipsHarness.Contract.TotalTips(&_TipsHarness.CallOpts, _type, _payload)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122056790fc7bfbdaacd63f1e470337fe00f0425373ab3e7edcda0de7de956248dd564736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"LOW_12_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TWELVE_BYTES\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b286bae7": "LOW_12_MASK()",
		"f26be3fc": "NULL()",
		"406cba16": "TWELVE_BYTES()",
	},
	Bin: "0x61011561003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060475760003560e01c8063406cba1614604c578063b286bae714606a578063f26be3fc146089575b600080fd5b6053606081565b60405160ff90911681526020015b60405180910390f35b607c6bffffffffffffffffffffffff81565b6040519081526020016061565b60af7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000009091168152602001606156fea2646970667358221220a892da2a7920e7ef0b1769f557036757e866f714d2b40094d7108da2d338caad64736f6c63430008110033",
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

// LOW12MASK is a free data retrieval call binding the contract method 0xb286bae7.
//
// Solidity: function LOW_12_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) LOW12MASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "LOW_12_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOW12MASK is a free data retrieval call binding the contract method 0xb286bae7.
//
// Solidity: function LOW_12_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) LOW12MASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW12MASK(&_TypedMemView.CallOpts)
}

// LOW12MASK is a free data retrieval call binding the contract method 0xb286bae7.
//
// Solidity: function LOW_12_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) LOW12MASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW12MASK(&_TypedMemView.CallOpts)
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

// TWELVEBYTES is a free data retrieval call binding the contract method 0x406cba16.
//
// Solidity: function TWELVE_BYTES() view returns(uint8)
func (_TypedMemView *TypedMemViewCaller) TWELVEBYTES(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "TWELVE_BYTES")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TWELVEBYTES is a free data retrieval call binding the contract method 0x406cba16.
//
// Solidity: function TWELVE_BYTES() view returns(uint8)
func (_TypedMemView *TypedMemViewSession) TWELVEBYTES() (uint8, error) {
	return _TypedMemView.Contract.TWELVEBYTES(&_TypedMemView.CallOpts)
}

// TWELVEBYTES is a free data retrieval call binding the contract method 0x406cba16.
//
// Solidity: function TWELVE_BYTES() view returns(uint8)
func (_TypedMemView *TypedMemViewCallerSession) TWELVEBYTES() (uint8, error) {
	return _TypedMemView.Contract.TWELVEBYTES(&_TypedMemView.CallOpts)
}
