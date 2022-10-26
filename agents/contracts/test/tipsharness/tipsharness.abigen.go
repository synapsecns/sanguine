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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bc29b13ed6c9408e777e330ccebb901a0a709c04e8b6686ee9da170d4eb8548164736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e8db5d77b0e75e92ca0fa3c6f1911eccc1fdddedb73dd2f06f79ad2e57243d4564736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"emptyTips\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_notaryTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_broadcasterTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_proverTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_executorTip\",\"type\":\"uint96\"}],\"name\":\"formatTips\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetBroadcaster\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetExecutor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetNotary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetProver\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tipsVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"725bd463": "emptyTips()",
		"d024f867": "formatTips(uint96,uint96,uint96,uint96)",
		"15bb7d2b": "offsetBroadcaster()",
		"51970c3f": "offsetExecutor()",
		"b4b4ccb2": "offsetNotary()",
		"98de8554": "offsetProver()",
		"0c096e8d": "offsetVersion()",
		"b440592e": "tipsLength()",
		"60fb5709": "tipsVersion()",
	},
	Bin: "0x608060405234801561001057600080fd5b50610318806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c8063725bd46311610076578063b440592e1161005b578063b440592e146100f7578063b4b4ccb2146100fe578063d024f8671461010557600080fd5b8063725bd463146100db57806398de8554146100f057600080fd5b80630c096e8d146100a857806315bb7d2b146100be57806351970c3f146100c557806360fb5709146100cc575b600080fd5b60005b6040519081526020015b60405180910390f35b600e6100ab565b60266100ab565b604051600181526020016100b5565b6100e361019c565b6040516100b59190610201565b601a6100ab565b60326100ab565b60026100ab565b6100e361011336600461028e565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffffffffffff000000000000000000000000000000000000000060a096871b8116602283015294861b8516602e82015292851b8416603a840152931b909116604682015281518082036032018152605290910190915290565b60606101fc604080517e010000000000000000000000000000000000000000000000000000000000006020820152600060228201819052602e8201819052603a820181905260468201528151808203603201815260529091019091526060905b905090565b600060208083528351808285015260005b8181101561022e57858101830151858201604001528201610212565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b80356bffffffffffffffffffffffff8116811461028957600080fd5b919050565b600080600080608085870312156102a457600080fd5b6102ad8561026d565b93506102bb6020860161026d565b92506102c96040860161026d565b91506102d76060860161026d565b90509295919450925056fea26469706673582212204f124516df2f4186d98a8bd976ca8a0b2a3c82e9ccd163dbc43b82c8f8dbe23e64736f6c63430008110033",
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122031b4b457fa538fd87b01e91bf686ad449b078e11fd204c674f5e17a88ed4dddd64736f6c63430008110033",
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
	Bin: "0x61011561003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060475760003560e01c8063406cba1614604c578063b286bae714606a578063f26be3fc146089575b600080fd5b6053606081565b60405160ff90911681526020015b60405180910390f35b607c6bffffffffffffffffffffffff81565b6040519081526020016061565b60af7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000009091168152602001606156fea26469706673582212201ae4b3f3b12d4f89a76be67e5d815ee50668ede18155fa0fc660572444bc56c764736f6c63430008110033",
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
