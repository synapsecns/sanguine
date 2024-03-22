// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenmessenger

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c15c51f6a775ad1c746a4f31af5246f1b4b8f127284e651d413674677162991664736f6c63430007060033",
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
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// BurnMessageMetaData contains all meta data concerning the BurnMessage contract.
var BurnMessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122065ac1d04477ca7cb9a700a2c2608f3ab4089475856c0b14cd9f685022d4e59db64736f6c63430007060033",
}

// BurnMessageABI is the input ABI used to generate the binding from.
// Deprecated: Use BurnMessageMetaData.ABI instead.
var BurnMessageABI = BurnMessageMetaData.ABI

// BurnMessageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BurnMessageMetaData.Bin instead.
var BurnMessageBin = BurnMessageMetaData.Bin

// DeployBurnMessage deploys a new Ethereum contract, binding an instance of BurnMessage to it.
func DeployBurnMessage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BurnMessage, error) {
	parsed, err := BurnMessageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BurnMessageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BurnMessage{BurnMessageCaller: BurnMessageCaller{contract: contract}, BurnMessageTransactor: BurnMessageTransactor{contract: contract}, BurnMessageFilterer: BurnMessageFilterer{contract: contract}}, nil
}

// BurnMessage is an auto generated Go binding around an Ethereum contract.
type BurnMessage struct {
	BurnMessageCaller     // Read-only binding to the contract
	BurnMessageTransactor // Write-only binding to the contract
	BurnMessageFilterer   // Log filterer for contract events
}

// BurnMessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BurnMessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnMessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BurnMessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnMessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BurnMessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BurnMessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BurnMessageSession struct {
	Contract     *BurnMessage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BurnMessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BurnMessageCallerSession struct {
	Contract *BurnMessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BurnMessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BurnMessageTransactorSession struct {
	Contract     *BurnMessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BurnMessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BurnMessageRaw struct {
	Contract *BurnMessage // Generic contract binding to access the raw methods on
}

// BurnMessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BurnMessageCallerRaw struct {
	Contract *BurnMessageCaller // Generic read-only contract binding to access the raw methods on
}

// BurnMessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BurnMessageTransactorRaw struct {
	Contract *BurnMessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBurnMessage creates a new instance of BurnMessage, bound to a specific deployed contract.
func NewBurnMessage(address common.Address, backend bind.ContractBackend) (*BurnMessage, error) {
	contract, err := bindBurnMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BurnMessage{BurnMessageCaller: BurnMessageCaller{contract: contract}, BurnMessageTransactor: BurnMessageTransactor{contract: contract}, BurnMessageFilterer: BurnMessageFilterer{contract: contract}}, nil
}

// NewBurnMessageCaller creates a new read-only instance of BurnMessage, bound to a specific deployed contract.
func NewBurnMessageCaller(address common.Address, caller bind.ContractCaller) (*BurnMessageCaller, error) {
	contract, err := bindBurnMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMessageCaller{contract: contract}, nil
}

// NewBurnMessageTransactor creates a new write-only instance of BurnMessage, bound to a specific deployed contract.
func NewBurnMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*BurnMessageTransactor, error) {
	contract, err := bindBurnMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BurnMessageTransactor{contract: contract}, nil
}

// NewBurnMessageFilterer creates a new log filterer instance of BurnMessage, bound to a specific deployed contract.
func NewBurnMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*BurnMessageFilterer, error) {
	contract, err := bindBurnMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BurnMessageFilterer{contract: contract}, nil
}

// bindBurnMessage binds a generic wrapper to an already deployed contract.
func bindBurnMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BurnMessageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnMessage *BurnMessageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMessage.Contract.BurnMessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnMessage *BurnMessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMessage.Contract.BurnMessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnMessage *BurnMessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMessage.Contract.BurnMessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BurnMessage *BurnMessageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BurnMessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BurnMessage *BurnMessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BurnMessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BurnMessage *BurnMessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BurnMessage.Contract.contract.Transact(opts, method, params...)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMessageHandlerMetaData contains all meta data concerning the IMessageHandler contract.
var IMessageHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"96abeb70": "handleReceiveMessage(uint32,bytes32,bytes)",
	},
}

// IMessageHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageHandlerMetaData.ABI instead.
var IMessageHandlerABI = IMessageHandlerMetaData.ABI

// Deprecated: Use IMessageHandlerMetaData.Sigs instead.
// IMessageHandlerFuncSigs maps the 4-byte function signature to its string representation.
var IMessageHandlerFuncSigs = IMessageHandlerMetaData.Sigs

// IMessageHandler is an auto generated Go binding around an Ethereum contract.
type IMessageHandler struct {
	IMessageHandlerCaller     // Read-only binding to the contract
	IMessageHandlerTransactor // Write-only binding to the contract
	IMessageHandlerFilterer   // Log filterer for contract events
}

// IMessageHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageHandlerSession struct {
	Contract     *IMessageHandler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageHandlerCallerSession struct {
	Contract *IMessageHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IMessageHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageHandlerTransactorSession struct {
	Contract     *IMessageHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IMessageHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageHandlerRaw struct {
	Contract *IMessageHandler // Generic contract binding to access the raw methods on
}

// IMessageHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageHandlerCallerRaw struct {
	Contract *IMessageHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageHandlerTransactorRaw struct {
	Contract *IMessageHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageHandler creates a new instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandler(address common.Address, backend bind.ContractBackend) (*IMessageHandler, error) {
	contract, err := bindIMessageHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageHandler{IMessageHandlerCaller: IMessageHandlerCaller{contract: contract}, IMessageHandlerTransactor: IMessageHandlerTransactor{contract: contract}, IMessageHandlerFilterer: IMessageHandlerFilterer{contract: contract}}, nil
}

// NewIMessageHandlerCaller creates a new read-only instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerCaller(address common.Address, caller bind.ContractCaller) (*IMessageHandlerCaller, error) {
	contract, err := bindIMessageHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerCaller{contract: contract}, nil
}

// NewIMessageHandlerTransactor creates a new write-only instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageHandlerTransactor, error) {
	contract, err := bindIMessageHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerTransactor{contract: contract}, nil
}

// NewIMessageHandlerFilterer creates a new log filterer instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageHandlerFilterer, error) {
	contract, err := bindIMessageHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerFilterer{contract: contract}, nil
}

// bindIMessageHandler binds a generic wrapper to an already deployed contract.
func bindIMessageHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageHandler *IMessageHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageHandler.Contract.IMessageHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageHandler *IMessageHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageHandler.Contract.IMessageHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageHandler *IMessageHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageHandler.Contract.IMessageHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageHandler *IMessageHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageHandler *IMessageHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageHandler *IMessageHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageHandler.Contract.contract.Transact(opts, method, params...)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 sourceDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_IMessageHandler *IMessageHandlerTransactor) HandleReceiveMessage(opts *bind.TransactOpts, sourceDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageHandler.contract.Transact(opts, "handleReceiveMessage", sourceDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 sourceDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_IMessageHandler *IMessageHandlerSession) HandleReceiveMessage(sourceDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageHandler.Contract.HandleReceiveMessage(&_IMessageHandler.TransactOpts, sourceDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 sourceDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_IMessageHandler *IMessageHandlerTransactorSession) HandleReceiveMessage(sourceDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageHandler.Contract.HandleReceiveMessage(&_IMessageHandler.TransactOpts, sourceDomain, sender, messageBody)
}

// IMessageTransmitterMetaData contains all meta data concerning the IMessageTransmitter contract.
var IMessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"originalAttestation\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newMessageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newDestinationCaller\",\"type\":\"bytes32\"}],\"name\":\"replaceMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"b857b774": "replaceMessage(bytes,bytes,bytes,bytes32)",
		"0ba469bc": "sendMessage(uint32,bytes32,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
}

// IMessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageTransmitterMetaData.ABI instead.
var IMessageTransmitterABI = IMessageTransmitterMetaData.ABI

// Deprecated: Use IMessageTransmitterMetaData.Sigs instead.
// IMessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var IMessageTransmitterFuncSigs = IMessageTransmitterMetaData.Sigs

// IMessageTransmitter is an auto generated Go binding around an Ethereum contract.
type IMessageTransmitter struct {
	IMessageTransmitterCaller     // Read-only binding to the contract
	IMessageTransmitterTransactor // Write-only binding to the contract
	IMessageTransmitterFilterer   // Log filterer for contract events
}

// IMessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageTransmitterSession struct {
	Contract     *IMessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageTransmitterCallerSession struct {
	Contract *IMessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageTransmitterTransactorSession struct {
	Contract     *IMessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageTransmitterRaw struct {
	Contract *IMessageTransmitter // Generic contract binding to access the raw methods on
}

// IMessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageTransmitterCallerRaw struct {
	Contract *IMessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactorRaw struct {
	Contract *IMessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageTransmitter creates a new instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitter(address common.Address, backend bind.ContractBackend) (*IMessageTransmitter, error) {
	contract, err := bindIMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitter{IMessageTransmitterCaller: IMessageTransmitterCaller{contract: contract}, IMessageTransmitterTransactor: IMessageTransmitterTransactor{contract: contract}, IMessageTransmitterFilterer: IMessageTransmitterFilterer{contract: contract}}, nil
}

// NewIMessageTransmitterCaller creates a new read-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*IMessageTransmitterCaller, error) {
	contract, err := bindIMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterCaller{contract: contract}, nil
}

// NewIMessageTransmitterTransactor creates a new write-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageTransmitterTransactor, error) {
	contract, err := bindIMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterTransactor{contract: contract}, nil
}

// NewIMessageTransmitterFilterer creates a new log filterer instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageTransmitterFilterer, error) {
	contract, err := bindIMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterFilterer{contract: contract}, nil
}

// bindIMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindIMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageTransmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.IMessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IMessageTransmitter *IMessageTransmitterTransactor) ReplaceMessage(opts *bind.TransactOpts, originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "replaceMessage", originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IMessageTransmitter *IMessageTransmitterSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReplaceMessage(&_IMessageTransmitter.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReplaceMessage(&_IMessageTransmitter.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactor) SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "sendMessage", destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessage(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessage(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// IMintBurnTokenMetaData contains all meta data concerning the IMintBurnToken contract.
var IMintBurnTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"42966c68": "burn(uint256)",
		"40c10f19": "mint(address,uint256)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IMintBurnTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IMintBurnTokenMetaData.ABI instead.
var IMintBurnTokenABI = IMintBurnTokenMetaData.ABI

// Deprecated: Use IMintBurnTokenMetaData.Sigs instead.
// IMintBurnTokenFuncSigs maps the 4-byte function signature to its string representation.
var IMintBurnTokenFuncSigs = IMintBurnTokenMetaData.Sigs

// IMintBurnToken is an auto generated Go binding around an Ethereum contract.
type IMintBurnToken struct {
	IMintBurnTokenCaller     // Read-only binding to the contract
	IMintBurnTokenTransactor // Write-only binding to the contract
	IMintBurnTokenFilterer   // Log filterer for contract events
}

// IMintBurnTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMintBurnTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintBurnTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMintBurnTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintBurnTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMintBurnTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMintBurnTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMintBurnTokenSession struct {
	Contract     *IMintBurnToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMintBurnTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMintBurnTokenCallerSession struct {
	Contract *IMintBurnTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IMintBurnTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMintBurnTokenTransactorSession struct {
	Contract     *IMintBurnTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IMintBurnTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMintBurnTokenRaw struct {
	Contract *IMintBurnToken // Generic contract binding to access the raw methods on
}

// IMintBurnTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMintBurnTokenCallerRaw struct {
	Contract *IMintBurnTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IMintBurnTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMintBurnTokenTransactorRaw struct {
	Contract *IMintBurnTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMintBurnToken creates a new instance of IMintBurnToken, bound to a specific deployed contract.
func NewIMintBurnToken(address common.Address, backend bind.ContractBackend) (*IMintBurnToken, error) {
	contract, err := bindIMintBurnToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMintBurnToken{IMintBurnTokenCaller: IMintBurnTokenCaller{contract: contract}, IMintBurnTokenTransactor: IMintBurnTokenTransactor{contract: contract}, IMintBurnTokenFilterer: IMintBurnTokenFilterer{contract: contract}}, nil
}

// NewIMintBurnTokenCaller creates a new read-only instance of IMintBurnToken, bound to a specific deployed contract.
func NewIMintBurnTokenCaller(address common.Address, caller bind.ContractCaller) (*IMintBurnTokenCaller, error) {
	contract, err := bindIMintBurnToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMintBurnTokenCaller{contract: contract}, nil
}

// NewIMintBurnTokenTransactor creates a new write-only instance of IMintBurnToken, bound to a specific deployed contract.
func NewIMintBurnTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IMintBurnTokenTransactor, error) {
	contract, err := bindIMintBurnToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMintBurnTokenTransactor{contract: contract}, nil
}

// NewIMintBurnTokenFilterer creates a new log filterer instance of IMintBurnToken, bound to a specific deployed contract.
func NewIMintBurnTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IMintBurnTokenFilterer, error) {
	contract, err := bindIMintBurnToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMintBurnTokenFilterer{contract: contract}, nil
}

// bindIMintBurnToken binds a generic wrapper to an already deployed contract.
func bindIMintBurnToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMintBurnTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintBurnToken *IMintBurnTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintBurnToken.Contract.IMintBurnTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintBurnToken *IMintBurnTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.IMintBurnTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintBurnToken *IMintBurnTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.IMintBurnTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMintBurnToken *IMintBurnTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMintBurnToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMintBurnToken *IMintBurnTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMintBurnToken *IMintBurnTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMintBurnToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMintBurnToken.Contract.Allowance(&_IMintBurnToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IMintBurnToken.Contract.Allowance(&_IMintBurnToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IMintBurnToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IMintBurnToken.Contract.BalanceOf(&_IMintBurnToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IMintBurnToken.Contract.BalanceOf(&_IMintBurnToken.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IMintBurnToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenSession) TotalSupply() (*big.Int, error) {
	return _IMintBurnToken.Contract.TotalSupply(&_IMintBurnToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IMintBurnToken *IMintBurnTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IMintBurnToken.Contract.TotalSupply(&_IMintBurnToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Approve(&_IMintBurnToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Approve(&_IMintBurnToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IMintBurnToken *IMintBurnTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IMintBurnToken *IMintBurnTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Burn(&_IMintBurnToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_IMintBurnToken *IMintBurnTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Burn(&_IMintBurnToken.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Mint(&_IMintBurnToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Mint(&_IMintBurnToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Transfer(&_IMintBurnToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.Transfer(&_IMintBurnToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.TransferFrom(&_IMintBurnToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IMintBurnToken *IMintBurnTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMintBurnToken.Contract.TransferFrom(&_IMintBurnToken.TransactOpts, sender, recipient, amount)
}

// IMintBurnTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IMintBurnToken contract.
type IMintBurnTokenApprovalIterator struct {
	Event *IMintBurnTokenApproval // Event containing the contract specifics and raw log

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
func (it *IMintBurnTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMintBurnTokenApproval)
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
		it.Event = new(IMintBurnTokenApproval)
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
func (it *IMintBurnTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMintBurnTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMintBurnTokenApproval represents a Approval event raised by the IMintBurnToken contract.
type IMintBurnTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMintBurnToken *IMintBurnTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IMintBurnTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMintBurnToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IMintBurnTokenApprovalIterator{contract: _IMintBurnToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMintBurnToken *IMintBurnTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IMintBurnTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IMintBurnToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMintBurnTokenApproval)
				if err := _IMintBurnToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IMintBurnToken *IMintBurnTokenFilterer) ParseApproval(log types.Log) (*IMintBurnTokenApproval, error) {
	event := new(IMintBurnTokenApproval)
	if err := _IMintBurnToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMintBurnTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IMintBurnToken contract.
type IMintBurnTokenTransferIterator struct {
	Event *IMintBurnTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IMintBurnTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMintBurnTokenTransfer)
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
		it.Event = new(IMintBurnTokenTransfer)
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
func (it *IMintBurnTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMintBurnTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMintBurnTokenTransfer represents a Transfer event raised by the IMintBurnToken contract.
type IMintBurnTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMintBurnToken *IMintBurnTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IMintBurnTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMintBurnToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IMintBurnTokenTransferIterator{contract: _IMintBurnToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMintBurnToken *IMintBurnTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IMintBurnTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IMintBurnToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMintBurnTokenTransfer)
				if err := _IMintBurnToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IMintBurnToken *IMintBurnTokenFilterer) ParseTransfer(log types.Log) (*IMintBurnTokenTransfer, error) {
	event := new(IMintBurnTokenTransfer)
	if err := _IMintBurnToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IReceiverMetaData contains all meta data concerning the IReceiver contract.
var IReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"57ecfd28": "receiveMessage(bytes,bytes)",
	},
}

// IReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IReceiverMetaData.ABI instead.
var IReceiverABI = IReceiverMetaData.ABI

// Deprecated: Use IReceiverMetaData.Sigs instead.
// IReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IReceiverFuncSigs = IReceiverMetaData.Sigs

// IReceiver is an auto generated Go binding around an Ethereum contract.
type IReceiver struct {
	IReceiverCaller     // Read-only binding to the contract
	IReceiverTransactor // Write-only binding to the contract
	IReceiverFilterer   // Log filterer for contract events
}

// IReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReceiverSession struct {
	Contract     *IReceiver        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReceiverCallerSession struct {
	Contract *IReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReceiverTransactorSession struct {
	Contract     *IReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReceiverRaw struct {
	Contract *IReceiver // Generic contract binding to access the raw methods on
}

// IReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReceiverCallerRaw struct {
	Contract *IReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReceiverTransactorRaw struct {
	Contract *IReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReceiver creates a new instance of IReceiver, bound to a specific deployed contract.
func NewIReceiver(address common.Address, backend bind.ContractBackend) (*IReceiver, error) {
	contract, err := bindIReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReceiver{IReceiverCaller: IReceiverCaller{contract: contract}, IReceiverTransactor: IReceiverTransactor{contract: contract}, IReceiverFilterer: IReceiverFilterer{contract: contract}}, nil
}

// NewIReceiverCaller creates a new read-only instance of IReceiver, bound to a specific deployed contract.
func NewIReceiverCaller(address common.Address, caller bind.ContractCaller) (*IReceiverCaller, error) {
	contract, err := bindIReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiverCaller{contract: contract}, nil
}

// NewIReceiverTransactor creates a new write-only instance of IReceiver, bound to a specific deployed contract.
func NewIReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IReceiverTransactor, error) {
	contract, err := bindIReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiverTransactor{contract: contract}, nil
}

// NewIReceiverFilterer creates a new log filterer instance of IReceiver, bound to a specific deployed contract.
func NewIReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IReceiverFilterer, error) {
	contract, err := bindIReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReceiverFilterer{contract: contract}, nil
}

// bindIReceiver binds a generic wrapper to an already deployed contract.
func bindIReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiver *IReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiver.Contract.IReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiver *IReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiver.Contract.IReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiver *IReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiver.Contract.IReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiver *IReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiver *IReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiver *IReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiver.Contract.contract.Transact(opts, method, params...)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IReceiver *IReceiverTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _IReceiver.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IReceiver *IReceiverSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IReceiver.Contract.ReceiveMessage(&_IReceiver.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IReceiver *IReceiverTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IReceiver.Contract.ReceiveMessage(&_IReceiver.TransactOpts, message, signature)
}

// IRelayerMetaData contains all meta data concerning the IRelayer contract.
var IRelayerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"originalAttestation\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newMessageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newDestinationCaller\",\"type\":\"bytes32\"}],\"name\":\"replaceMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b857b774": "replaceMessage(bytes,bytes,bytes,bytes32)",
		"0ba469bc": "sendMessage(uint32,bytes32,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
}

// IRelayerABI is the input ABI used to generate the binding from.
// Deprecated: Use IRelayerMetaData.ABI instead.
var IRelayerABI = IRelayerMetaData.ABI

// Deprecated: Use IRelayerMetaData.Sigs instead.
// IRelayerFuncSigs maps the 4-byte function signature to its string representation.
var IRelayerFuncSigs = IRelayerMetaData.Sigs

// IRelayer is an auto generated Go binding around an Ethereum contract.
type IRelayer struct {
	IRelayerCaller     // Read-only binding to the contract
	IRelayerTransactor // Write-only binding to the contract
	IRelayerFilterer   // Log filterer for contract events
}

// IRelayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayerSession struct {
	Contract     *IRelayer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRelayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayerCallerSession struct {
	Contract *IRelayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IRelayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayerTransactorSession struct {
	Contract     *IRelayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IRelayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayerRaw struct {
	Contract *IRelayer // Generic contract binding to access the raw methods on
}

// IRelayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayerCallerRaw struct {
	Contract *IRelayerCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayerTransactorRaw struct {
	Contract *IRelayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayer creates a new instance of IRelayer, bound to a specific deployed contract.
func NewIRelayer(address common.Address, backend bind.ContractBackend) (*IRelayer, error) {
	contract, err := bindIRelayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayer{IRelayerCaller: IRelayerCaller{contract: contract}, IRelayerTransactor: IRelayerTransactor{contract: contract}, IRelayerFilterer: IRelayerFilterer{contract: contract}}, nil
}

// NewIRelayerCaller creates a new read-only instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerCaller(address common.Address, caller bind.ContractCaller) (*IRelayerCaller, error) {
	contract, err := bindIRelayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerCaller{contract: contract}, nil
}

// NewIRelayerTransactor creates a new write-only instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayerTransactor, error) {
	contract, err := bindIRelayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerTransactor{contract: contract}, nil
}

// NewIRelayerFilterer creates a new log filterer instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayerFilterer, error) {
	contract, err := bindIRelayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayerFilterer{contract: contract}, nil
}

// bindIRelayer binds a generic wrapper to an already deployed contract.
func bindIRelayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRelayerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayer *IRelayerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayer.Contract.IRelayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayer *IRelayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayer.Contract.IRelayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayer *IRelayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayer.Contract.IRelayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayer *IRelayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayer *IRelayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayer *IRelayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayer.Contract.contract.Transact(opts, method, params...)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IRelayer *IRelayerTransactor) ReplaceMessage(opts *bind.TransactOpts, originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "replaceMessage", originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IRelayer *IRelayerSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IRelayer.Contract.ReplaceMessage(&_IRelayer.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IRelayer *IRelayerTransactorSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IRelayer.Contract.ReplaceMessage(&_IRelayer.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactor) SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "sendMessage", destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessage(&_IRelayer.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactorSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessage(&_IRelayer.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessageWithCaller(&_IRelayer.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessageWithCaller(&_IRelayer.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// ITokenMinterMetaData contains all meta data concerning the ITokenMinter contract.
var ITokenMinterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"remoteToken\",\"type\":\"bytes32\"}],\"name\":\"getLocalToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"burnToken\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTokenController\",\"type\":\"address\"}],\"name\":\"setTokenController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9dc29fac": "burn(address,uint256)",
		"78a0565e": "getLocalToken(uint32,bytes32)",
		"d54de06f": "mint(uint32,bytes32,address,uint256)",
		"e102baab": "setTokenController(address)",
	},
}

// ITokenMinterABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMinterMetaData.ABI instead.
var ITokenMinterABI = ITokenMinterMetaData.ABI

// Deprecated: Use ITokenMinterMetaData.Sigs instead.
// ITokenMinterFuncSigs maps the 4-byte function signature to its string representation.
var ITokenMinterFuncSigs = ITokenMinterMetaData.Sigs

// ITokenMinter is an auto generated Go binding around an Ethereum contract.
type ITokenMinter struct {
	ITokenMinterCaller     // Read-only binding to the contract
	ITokenMinterTransactor // Write-only binding to the contract
	ITokenMinterFilterer   // Log filterer for contract events
}

// ITokenMinterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenMinterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenMinterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenMinterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenMinterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenMinterSession struct {
	Contract     *ITokenMinter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenMinterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenMinterCallerSession struct {
	Contract *ITokenMinterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ITokenMinterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenMinterTransactorSession struct {
	Contract     *ITokenMinterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ITokenMinterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenMinterRaw struct {
	Contract *ITokenMinter // Generic contract binding to access the raw methods on
}

// ITokenMinterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenMinterCallerRaw struct {
	Contract *ITokenMinterCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenMinterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenMinterTransactorRaw struct {
	Contract *ITokenMinterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITokenMinter creates a new instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinter(address common.Address, backend bind.ContractBackend) (*ITokenMinter, error) {
	contract, err := bindITokenMinter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITokenMinter{ITokenMinterCaller: ITokenMinterCaller{contract: contract}, ITokenMinterTransactor: ITokenMinterTransactor{contract: contract}, ITokenMinterFilterer: ITokenMinterFilterer{contract: contract}}, nil
}

// NewITokenMinterCaller creates a new read-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterCaller(address common.Address, caller bind.ContractCaller) (*ITokenMinterCaller, error) {
	contract, err := bindITokenMinter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterCaller{contract: contract}, nil
}

// NewITokenMinterTransactor creates a new write-only instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenMinterTransactor, error) {
	contract, err := bindITokenMinter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterTransactor{contract: contract}, nil
}

// NewITokenMinterFilterer creates a new log filterer instance of ITokenMinter, bound to a specific deployed contract.
func NewITokenMinterFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenMinterFilterer, error) {
	contract, err := bindITokenMinter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenMinterFilterer{contract: contract}, nil
}

// bindITokenMinter binds a generic wrapper to an already deployed contract.
func bindITokenMinter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ITokenMinterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.ITokenMinterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.ITokenMinterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITokenMinter *ITokenMinterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITokenMinter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITokenMinter *ITokenMinterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITokenMinter.Contract.contract.Transact(opts, method, params...)
}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterCaller) GetLocalToken(opts *bind.CallOpts, remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ITokenMinter.contract.Call(opts, &out, "getLocalToken", remoteDomain, remoteToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _ITokenMinter.Contract.GetLocalToken(&_ITokenMinter.CallOpts, remoteDomain, remoteToken)
}

// GetLocalToken is a free data retrieval call binding the contract method 0x78a0565e.
//
// Solidity: function getLocalToken(uint32 remoteDomain, bytes32 remoteToken) view returns(address)
func (_ITokenMinter *ITokenMinterCallerSession) GetLocalToken(remoteDomain uint32, remoteToken [32]byte) (common.Address, error) {
	return _ITokenMinter.Contract.GetLocalToken(&_ITokenMinter.CallOpts, remoteDomain, remoteToken)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterTransactor) Burn(opts *bind.TransactOpts, burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "burn", burnToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterSession) Burn(burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, burnToken, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address burnToken, uint256 amount) returns()
func (_ITokenMinter *ITokenMinterTransactorSession) Burn(burnToken common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Burn(&_ITokenMinter.TransactOpts, burnToken, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterTransactor) Mint(opts *bind.TransactOpts, sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "mint", sourceDomain, burnToken, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterSession) Mint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, sourceDomain, burnToken, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xd54de06f.
//
// Solidity: function mint(uint32 sourceDomain, bytes32 burnToken, address to, uint256 amount) returns(address mintToken)
func (_ITokenMinter *ITokenMinterTransactorSession) Mint(sourceDomain uint32, burnToken [32]byte, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITokenMinter.Contract.Mint(&_ITokenMinter.TransactOpts, sourceDomain, burnToken, to, amount)
}

// SetTokenController is a paid mutator transaction binding the contract method 0xe102baab.
//
// Solidity: function setTokenController(address newTokenController) returns()
func (_ITokenMinter *ITokenMinterTransactor) SetTokenController(opts *bind.TransactOpts, newTokenController common.Address) (*types.Transaction, error) {
	return _ITokenMinter.contract.Transact(opts, "setTokenController", newTokenController)
}

// SetTokenController is a paid mutator transaction binding the contract method 0xe102baab.
//
// Solidity: function setTokenController(address newTokenController) returns()
func (_ITokenMinter *ITokenMinterSession) SetTokenController(newTokenController common.Address) (*types.Transaction, error) {
	return _ITokenMinter.Contract.SetTokenController(&_ITokenMinter.TransactOpts, newTokenController)
}

// SetTokenController is a paid mutator transaction binding the contract method 0xe102baab.
//
// Solidity: function setTokenController(address newTokenController) returns()
func (_ITokenMinter *ITokenMinterTransactorSession) SetTokenController(newTokenController common.Address) (*types.Transaction, error) {
	return _ITokenMinter.Contract.SetTokenController(&_ITokenMinter.TransactOpts, newTokenController)
}

// MessageMetaData contains all meta data concerning the Message contract.
var MessageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_buf\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"82c947b7": "addressToBytes32(address)",
		"5ced058e": "bytes32ToAddress(bytes32)",
	},
	Bin: "0x610119610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80635ced058e14604257806382c947b7146085575b600080fd5b605c60048036036020811015605657600080fd5b503560c7565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60b560048036036020811015609957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1660ca565b60408051918252519081900360200190f35b90565b73ffffffffffffffffffffffffffffffffffffffff169056fea2646970667358221220f604d6ff0f07e23e9aa7a0cc513d199a6e023d0baa4a78c87be130028d6477c064736f6c63430007060033",
}

// MessageABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageMetaData.ABI instead.
var MessageABI = MessageMetaData.ABI

// Deprecated: Use MessageMetaData.Sigs instead.
// MessageFuncSigs maps the 4-byte function signature to its string representation.
var MessageFuncSigs = MessageMetaData.Sigs

// MessageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageMetaData.Bin instead.
var MessageBin = MessageMetaData.Bin

// DeployMessage deploys a new Ethereum contract, binding an instance of Message to it.
func DeployMessage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Message, error) {
	parsed, err := MessageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Message{MessageCaller: MessageCaller{contract: contract}, MessageTransactor: MessageTransactor{contract: contract}, MessageFilterer: MessageFilterer{contract: contract}}, nil
}

// Message is an auto generated Go binding around an Ethereum contract.
type Message struct {
	MessageCaller     // Read-only binding to the contract
	MessageTransactor // Write-only binding to the contract
	MessageFilterer   // Log filterer for contract events
}

// MessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageSession struct {
	Contract     *Message          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageCallerSession struct {
	Contract *MessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTransactorSession struct {
	Contract     *MessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageRaw struct {
	Contract *Message // Generic contract binding to access the raw methods on
}

// MessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageCallerRaw struct {
	Contract *MessageCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTransactorRaw struct {
	Contract *MessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessage creates a new instance of Message, bound to a specific deployed contract.
func NewMessage(address common.Address, backend bind.ContractBackend) (*Message, error) {
	contract, err := bindMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Message{MessageCaller: MessageCaller{contract: contract}, MessageTransactor: MessageTransactor{contract: contract}, MessageFilterer: MessageFilterer{contract: contract}}, nil
}

// NewMessageCaller creates a new read-only instance of Message, bound to a specific deployed contract.
func NewMessageCaller(address common.Address, caller bind.ContractCaller) (*MessageCaller, error) {
	contract, err := bindMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageCaller{contract: contract}, nil
}

// NewMessageTransactor creates a new write-only instance of Message, bound to a specific deployed contract.
func NewMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTransactor, error) {
	contract, err := bindMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransactor{contract: contract}, nil
}

// NewMessageFilterer creates a new log filterer instance of Message, bound to a specific deployed contract.
func NewMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageFilterer, error) {
	contract, err := bindMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageFilterer{contract: contract}, nil
}

// bindMessage binds a generic wrapper to an already deployed contract.
func bindMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Message *MessageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Message.Contract.MessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Message *MessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.Contract.MessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Message *MessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Message.Contract.MessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Message *MessageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Message.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Message *MessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Message.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Message *MessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Message.Contract.contract.Transact(opts, method, params...)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address addr) pure returns(bytes32)
func (_Message *MessageCaller) AddressToBytes32(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Message.contract.Call(opts, &out, "addressToBytes32", addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address addr) pure returns(bytes32)
func (_Message *MessageSession) AddressToBytes32(addr common.Address) ([32]byte, error) {
	return _Message.Contract.AddressToBytes32(&_Message.CallOpts, addr)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address addr) pure returns(bytes32)
func (_Message *MessageCallerSession) AddressToBytes32(addr common.Address) ([32]byte, error) {
	return _Message.Contract.AddressToBytes32(&_Message.CallOpts, addr)
}

// Bytes32ToAddress is a free data retrieval call binding the contract method 0x5ced058e.
//
// Solidity: function bytes32ToAddress(bytes32 _buf) pure returns(address)
func (_Message *MessageCaller) Bytes32ToAddress(opts *bind.CallOpts, _buf [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Message.contract.Call(opts, &out, "bytes32ToAddress", _buf)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bytes32ToAddress is a free data retrieval call binding the contract method 0x5ced058e.
//
// Solidity: function bytes32ToAddress(bytes32 _buf) pure returns(address)
func (_Message *MessageSession) Bytes32ToAddress(_buf [32]byte) (common.Address, error) {
	return _Message.Contract.Bytes32ToAddress(&_Message.CallOpts, _buf)
}

// Bytes32ToAddress is a free data retrieval call binding the contract method 0x5ced058e.
//
// Solidity: function bytes32ToAddress(bytes32 _buf) pure returns(address)
func (_Message *MessageCallerSession) Bytes32ToAddress(_buf [32]byte) (common.Address, error) {
	return _Message.Contract.Bytes32ToAddress(&_Message.CallOpts, _buf)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ownable2StepMetaData contains all meta data concerning the Ownable2Step contract.
var Ownable2StepMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// Ownable2StepABI is the input ABI used to generate the binding from.
// Deprecated: Use Ownable2StepMetaData.ABI instead.
var Ownable2StepABI = Ownable2StepMetaData.ABI

// Deprecated: Use Ownable2StepMetaData.Sigs instead.
// Ownable2StepFuncSigs maps the 4-byte function signature to its string representation.
var Ownable2StepFuncSigs = Ownable2StepMetaData.Sigs

// Ownable2Step is an auto generated Go binding around an Ethereum contract.
type Ownable2Step struct {
	Ownable2StepCaller     // Read-only binding to the contract
	Ownable2StepTransactor // Write-only binding to the contract
	Ownable2StepFilterer   // Log filterer for contract events
}

// Ownable2StepCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ownable2StepCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ownable2StepTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ownable2StepFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ownable2StepSession struct {
	Contract     *Ownable2Step     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ownable2StepCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ownable2StepCallerSession struct {
	Contract *Ownable2StepCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Ownable2StepTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ownable2StepTransactorSession struct {
	Contract     *Ownable2StepTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Ownable2StepRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ownable2StepRaw struct {
	Contract *Ownable2Step // Generic contract binding to access the raw methods on
}

// Ownable2StepCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ownable2StepCallerRaw struct {
	Contract *Ownable2StepCaller // Generic read-only contract binding to access the raw methods on
}

// Ownable2StepTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ownable2StepTransactorRaw struct {
	Contract *Ownable2StepTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable2Step creates a new instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2Step(address common.Address, backend bind.ContractBackend) (*Ownable2Step, error) {
	contract, err := bindOwnable2Step(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable2Step{Ownable2StepCaller: Ownable2StepCaller{contract: contract}, Ownable2StepTransactor: Ownable2StepTransactor{contract: contract}, Ownable2StepFilterer: Ownable2StepFilterer{contract: contract}}, nil
}

// NewOwnable2StepCaller creates a new read-only instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepCaller(address common.Address, caller bind.ContractCaller) (*Ownable2StepCaller, error) {
	contract, err := bindOwnable2Step(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepCaller{contract: contract}, nil
}

// NewOwnable2StepTransactor creates a new write-only instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepTransactor(address common.Address, transactor bind.ContractTransactor) (*Ownable2StepTransactor, error) {
	contract, err := bindOwnable2Step(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepTransactor{contract: contract}, nil
}

// NewOwnable2StepFilterer creates a new log filterer instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepFilterer(address common.Address, filterer bind.ContractFilterer) (*Ownable2StepFilterer, error) {
	contract, err := bindOwnable2Step(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepFilterer{contract: contract}, nil
}

// bindOwnable2Step binds a generic wrapper to an already deployed contract.
func bindOwnable2Step(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ownable2StepMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2Step *Ownable2StepRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2Step.Contract.Ownable2StepCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2Step *Ownable2StepRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.Contract.Ownable2StepTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2Step *Ownable2StepRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2Step.Contract.Ownable2StepTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2Step *Ownable2StepCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2Step.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2Step *Ownable2StepTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2Step *Ownable2StepTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2Step.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2Step.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepSession) Owner() (common.Address, error) {
	return _Ownable2Step.Contract.Owner(&_Ownable2Step.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepCallerSession) Owner() (common.Address, error) {
	return _Ownable2Step.Contract.Owner(&_Ownable2Step.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2Step.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepSession) PendingOwner() (common.Address, error) {
	return _Ownable2Step.Contract.PendingOwner(&_Ownable2Step.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepCallerSession) PendingOwner() (common.Address, error) {
	return _Ownable2Step.Contract.PendingOwner(&_Ownable2Step.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.AcceptOwnership(&_Ownable2Step.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.AcceptOwnership(&_Ownable2Step.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.Contract.TransferOwnership(&_Ownable2Step.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.Contract.TransferOwnership(&_Ownable2Step.TransactOpts, newOwner)
}

// Ownable2StepOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferStartedIterator struct {
	Event *Ownable2StepOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *Ownable2StepOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepOwnershipTransferStarted)
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
		it.Event = new(Ownable2StepOwnershipTransferStarted)
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
func (it *Ownable2StepOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepOwnershipTransferStartedIterator{contract: _Ownable2Step.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *Ownable2StepOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepOwnershipTransferStarted)
				if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Ownable2Step *Ownable2StepFilterer) ParseOwnershipTransferStarted(log types.Log) (*Ownable2StepOwnershipTransferStarted, error) {
	event := new(Ownable2StepOwnershipTransferStarted)
	if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ownable2StepOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferredIterator struct {
	Event *Ownable2StepOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Ownable2StepOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepOwnershipTransferred)
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
		it.Event = new(Ownable2StepOwnershipTransferred)
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
func (it *Ownable2StepOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepOwnershipTransferredIterator{contract: _Ownable2Step.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Ownable2StepOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepOwnershipTransferred)
				if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable2Step *Ownable2StepFilterer) ParseOwnershipTransferred(log types.Log) (*Ownable2StepOwnershipTransferred, error) {
	event := new(Ownable2StepOwnershipTransferred)
	if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RescuableMetaData contains all meta data concerning the Rescuable contract.
var RescuableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"b2118a8d": "rescueERC20(address,address,uint256)",
		"38a63183": "rescuer()",
		"f2fde38b": "transferOwnership(address)",
		"2ab60045": "updateRescuer(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061002161001c610026565b61002a565b6100a1565b3390565b600180546001600160a01b031916905561004e81610051602090811b61047e17901c565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610ac6806100b06000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100f0578063b2118a8d146100f8578063e30c39781461013b578063f2fde38b146101435761007d565b80632ab600451461008257806338a63183146100b757806379ba5097146100e8575b600080fd5b6100b56004803603602081101561009857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610176565b005b6100bf610259565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6100b5610275565b6100bf610318565b6100b56004803603606081101561010e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101359091169060400135610334565b6100bf6103ca565b6100b56004803603602081101561015957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166103e6565b61017e6104f3565b73ffffffffffffffffffffffffffffffffffffffff81166101ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806109f3602a913960400191505060405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a90600090a250565b60025473ffffffffffffffffffffffffffffffffffffffff1690565b600061027f61059d565b90508073ffffffffffffffffffffffffffffffffffffffff166102a06103ca565b73ffffffffffffffffffffffffffffffffffffffff161461030c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806109ca6029913960400191505060405180910390fd5b610315816105a1565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60025473ffffffffffffffffffffffffffffffffffffffff1633146103a4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610a436024913960400191505060405180910390fd5b6103c573ffffffffffffffffffffffffffffffffffffffff841683836105d2565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff1690565b6103ee6104f3565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155610439610318565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6104fb61059d565b73ffffffffffffffffffffffffffffffffffffffff16610519610318565b73ffffffffffffffffffffffffffffffffffffffff161461059b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b565b3390565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556103158161047e565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526103c590849060006106bc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107329092919063ffffffff16565b8051909150156103c5578080602001905160208110156106db57600080fd5b50516103c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610a67602a913960400191505060405180910390fd5b6060610741848460008561074b565b90505b9392505050565b6060824710156107a6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610a1d6026913960400191505060405180910390fd5b6107af85610905565b61081a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b6020831061088357805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610846565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146108e5576040519150601f19603f3d011682016040523d82523d6000602084013e6108ea565b606091505b50915091506108fa82828661090b565b979650505050505050565b3b151590565b6060831561091a575081610744565b82511561092a5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561098e578181015183820152602001610976565b50505050905090810190601f1680156109bb5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c526573637561626c653a2063616c6c6572206973206e6f742074686520726573637565725361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220111c14accabfad16a2deb2939dae03d625c7d5f7274599139a4e49d7182bb26564736f6c63430007060033",
}

// RescuableABI is the input ABI used to generate the binding from.
// Deprecated: Use RescuableMetaData.ABI instead.
var RescuableABI = RescuableMetaData.ABI

// Deprecated: Use RescuableMetaData.Sigs instead.
// RescuableFuncSigs maps the 4-byte function signature to its string representation.
var RescuableFuncSigs = RescuableMetaData.Sigs

// RescuableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RescuableMetaData.Bin instead.
var RescuableBin = RescuableMetaData.Bin

// DeployRescuable deploys a new Ethereum contract, binding an instance of Rescuable to it.
func DeployRescuable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Rescuable, error) {
	parsed, err := RescuableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RescuableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rescuable{RescuableCaller: RescuableCaller{contract: contract}, RescuableTransactor: RescuableTransactor{contract: contract}, RescuableFilterer: RescuableFilterer{contract: contract}}, nil
}

// Rescuable is an auto generated Go binding around an Ethereum contract.
type Rescuable struct {
	RescuableCaller     // Read-only binding to the contract
	RescuableTransactor // Write-only binding to the contract
	RescuableFilterer   // Log filterer for contract events
}

// RescuableCaller is an auto generated read-only Go binding around an Ethereum contract.
type RescuableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RescuableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RescuableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RescuableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RescuableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RescuableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RescuableSession struct {
	Contract     *Rescuable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RescuableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RescuableCallerSession struct {
	Contract *RescuableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RescuableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RescuableTransactorSession struct {
	Contract     *RescuableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RescuableRaw is an auto generated low-level Go binding around an Ethereum contract.
type RescuableRaw struct {
	Contract *Rescuable // Generic contract binding to access the raw methods on
}

// RescuableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RescuableCallerRaw struct {
	Contract *RescuableCaller // Generic read-only contract binding to access the raw methods on
}

// RescuableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RescuableTransactorRaw struct {
	Contract *RescuableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRescuable creates a new instance of Rescuable, bound to a specific deployed contract.
func NewRescuable(address common.Address, backend bind.ContractBackend) (*Rescuable, error) {
	contract, err := bindRescuable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rescuable{RescuableCaller: RescuableCaller{contract: contract}, RescuableTransactor: RescuableTransactor{contract: contract}, RescuableFilterer: RescuableFilterer{contract: contract}}, nil
}

// NewRescuableCaller creates a new read-only instance of Rescuable, bound to a specific deployed contract.
func NewRescuableCaller(address common.Address, caller bind.ContractCaller) (*RescuableCaller, error) {
	contract, err := bindRescuable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RescuableCaller{contract: contract}, nil
}

// NewRescuableTransactor creates a new write-only instance of Rescuable, bound to a specific deployed contract.
func NewRescuableTransactor(address common.Address, transactor bind.ContractTransactor) (*RescuableTransactor, error) {
	contract, err := bindRescuable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RescuableTransactor{contract: contract}, nil
}

// NewRescuableFilterer creates a new log filterer instance of Rescuable, bound to a specific deployed contract.
func NewRescuableFilterer(address common.Address, filterer bind.ContractFilterer) (*RescuableFilterer, error) {
	contract, err := bindRescuable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RescuableFilterer{contract: contract}, nil
}

// bindRescuable binds a generic wrapper to an already deployed contract.
func bindRescuable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RescuableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rescuable *RescuableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rescuable.Contract.RescuableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rescuable *RescuableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rescuable.Contract.RescuableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rescuable *RescuableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rescuable.Contract.RescuableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rescuable *RescuableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rescuable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rescuable *RescuableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rescuable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rescuable *RescuableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rescuable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rescuable *RescuableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rescuable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rescuable *RescuableSession) Owner() (common.Address, error) {
	return _Rescuable.Contract.Owner(&_Rescuable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rescuable *RescuableCallerSession) Owner() (common.Address, error) {
	return _Rescuable.Contract.Owner(&_Rescuable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Rescuable *RescuableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rescuable.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Rescuable *RescuableSession) PendingOwner() (common.Address, error) {
	return _Rescuable.Contract.PendingOwner(&_Rescuable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Rescuable *RescuableCallerSession) PendingOwner() (common.Address, error) {
	return _Rescuable.Contract.PendingOwner(&_Rescuable.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Rescuable *RescuableCaller) Rescuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rescuable.contract.Call(opts, &out, "rescuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Rescuable *RescuableSession) Rescuer() (common.Address, error) {
	return _Rescuable.Contract.Rescuer(&_Rescuable.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Rescuable *RescuableCallerSession) Rescuer() (common.Address, error) {
	return _Rescuable.Contract.Rescuer(&_Rescuable.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Rescuable *RescuableTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Rescuable *RescuableSession) AcceptOwnership() (*types.Transaction, error) {
	return _Rescuable.Contract.AcceptOwnership(&_Rescuable.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Rescuable *RescuableTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Rescuable.Contract.AcceptOwnership(&_Rescuable.TransactOpts)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Rescuable *RescuableTransactor) RescueERC20(opts *bind.TransactOpts, tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "rescueERC20", tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Rescuable *RescuableSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rescuable.Contract.RescueERC20(&_Rescuable.TransactOpts, tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Rescuable *RescuableTransactorSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rescuable.Contract.RescueERC20(&_Rescuable.TransactOpts, tokenContract, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rescuable *RescuableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rescuable *RescuableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.TransferOwnership(&_Rescuable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rescuable *RescuableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.TransferOwnership(&_Rescuable.TransactOpts, newOwner)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Rescuable *RescuableTransactor) UpdateRescuer(opts *bind.TransactOpts, newRescuer common.Address) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "updateRescuer", newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Rescuable *RescuableSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.UpdateRescuer(&_Rescuable.TransactOpts, newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Rescuable *RescuableTransactorSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.UpdateRescuer(&_Rescuable.TransactOpts, newRescuer)
}

// RescuableOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Rescuable contract.
type RescuableOwnershipTransferStartedIterator struct {
	Event *RescuableOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *RescuableOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RescuableOwnershipTransferStarted)
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
		it.Event = new(RescuableOwnershipTransferStarted)
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
func (it *RescuableOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RescuableOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RescuableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Rescuable contract.
type RescuableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RescuableOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RescuableOwnershipTransferStartedIterator{contract: _Rescuable.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *RescuableOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RescuableOwnershipTransferStarted)
				if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_Rescuable *RescuableFilterer) ParseOwnershipTransferStarted(log types.Log) (*RescuableOwnershipTransferStarted, error) {
	event := new(RescuableOwnershipTransferStarted)
	if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RescuableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rescuable contract.
type RescuableOwnershipTransferredIterator struct {
	Event *RescuableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RescuableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RescuableOwnershipTransferred)
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
		it.Event = new(RescuableOwnershipTransferred)
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
func (it *RescuableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RescuableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RescuableOwnershipTransferred represents a OwnershipTransferred event raised by the Rescuable contract.
type RescuableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RescuableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RescuableOwnershipTransferredIterator{contract: _Rescuable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RescuableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RescuableOwnershipTransferred)
				if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rescuable *RescuableFilterer) ParseOwnershipTransferred(log types.Log) (*RescuableOwnershipTransferred, error) {
	event := new(RescuableOwnershipTransferred)
	if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RescuableRescuerChangedIterator is returned from FilterRescuerChanged and is used to iterate over the raw logs and unpacked data for RescuerChanged events raised by the Rescuable contract.
type RescuableRescuerChangedIterator struct {
	Event *RescuableRescuerChanged // Event containing the contract specifics and raw log

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
func (it *RescuableRescuerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RescuableRescuerChanged)
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
		it.Event = new(RescuableRescuerChanged)
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
func (it *RescuableRescuerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RescuableRescuerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RescuableRescuerChanged represents a RescuerChanged event raised by the Rescuable contract.
type RescuableRescuerChanged struct {
	NewRescuer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRescuerChanged is a free log retrieval operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Rescuable *RescuableFilterer) FilterRescuerChanged(opts *bind.FilterOpts, newRescuer []common.Address) (*RescuableRescuerChangedIterator, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _Rescuable.contract.FilterLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return &RescuableRescuerChangedIterator{contract: _Rescuable.contract, event: "RescuerChanged", logs: logs, sub: sub}, nil
}

// WatchRescuerChanged is a free log subscription operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Rescuable *RescuableFilterer) WatchRescuerChanged(opts *bind.WatchOpts, sink chan<- *RescuableRescuerChanged, newRescuer []common.Address) (event.Subscription, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _Rescuable.contract.WatchLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RescuableRescuerChanged)
				if err := _Rescuable.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
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

// ParseRescuerChanged is a log parse operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Rescuable *RescuableFilterer) ParseRescuerChanged(log types.Log) (*RescuableRescuerChanged, error) {
	event := new(RescuableRescuerChanged)
	if err := _Rescuable.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220633ba2adca837bb8edb640b403b7d7cbeb6bedac9e09f0a5e9c31c97ed986bec64736f6c63430007060033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206f0cfac63ce46d7c0a7224280b8de479bad9d100e41a2c5c44189ec264d9690364736f6c63430007060033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TokenMessengerMetaData contains all meta data concerning the TokenMessenger contract.
var TokenMessengerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_messageTransmitter\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_messageBodyVersion\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"depositor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationTokenMessenger\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"DepositForBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localMinter\",\"type\":\"address\"}],\"name\":\"LocalMinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"localMinter\",\"type\":\"address\"}],\"name\":\"LocalMinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mintToken\",\"type\":\"address\"}],\"name\":\"MintAndWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"RemoteTokenMessengerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"RemoteTokenMessengerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLocalMinter\",\"type\":\"address\"}],\"name\":\"addLocalMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"tokenMessenger\",\"type\":\"bytes32\"}],\"name\":\"addRemoteTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"}],\"name\":\"depositForBurn\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"mintRecipient\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"burnToken\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"}],\"name\":\"depositForBurnWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"remoteDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMessageTransmitter\",\"outputs\":[{\"internalType\":\"contractIMessageTransmitter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localMinter\",\"outputs\":[{\"internalType\":\"contractITokenMinter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageBodyVersion\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"remoteTokenMessengers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeLocalMinter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"domain\",\"type\":\"uint32\"}],\"name\":\"removeRemoteTokenMessenger\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"originalAttestation\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newDestinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newMintRecipient\",\"type\":\"bytes32\"}],\"name\":\"replaceDepositForBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8197beb9": "addLocalMinter(address)",
		"da87e448": "addRemoteTokenMessenger(uint32,bytes32)",
		"6fd3504e": "depositForBurn(uint256,uint32,bytes32,address)",
		"f856ddb6": "depositForBurnWithCaller(uint256,uint32,bytes32,address,bytes32)",
		"96abeb70": "handleReceiveMessage(uint32,bytes32,bytes)",
		"2c121921": "localMessageTransmitter()",
		"cb75c11c": "localMinter()",
		"9cdbb181": "messageBodyVersion()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"82a5e665": "remoteTokenMessengers(uint32)",
		"91f17888": "removeLocalMinter()",
		"f79fd08e": "removeRemoteTokenMessenger(uint32)",
		"29a78e33": "replaceDepositForBurn(bytes,bytes,bytes32,bytes32)",
		"b2118a8d": "rescueERC20(address,address,uint256)",
		"38a63183": "rescuer()",
		"f2fde38b": "transferOwnership(address)",
		"2ab60045": "updateRescuer(address)",
	},
	Bin: "0x60c06040523480156200001157600080fd5b5060405162002d6e38038062002d6e833981810160405260408110156200003757600080fd5b508051602090910151620000546200004e620000d9565b620000dd565b6001600160a01b038216620000b0576040805162461bcd60e51b815260206004820152601a60248201527f4d6573736167655472616e736d6974746572206e6f7420736574000000000000604482015290519081900360640190fd5b60609190911b6001600160601b03191660805260e01b6001600160e01b03191660a05262000157565b3390565b600180546001600160a01b0319169055620001048162000107602090811b6200130217901c565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60805160601c60a05160e01c612bc5620001a9600039806106d55280610e2e5280610f9552806118f05250806107015280610a595280611af85280611b3452806120f252806121c95250612bc56000f3fe608060405234801561001057600080fd5b50600436106101515760003560e01c806391f17888116100cd578063da87e44811610081578063f2fde38b11610066578063f2fde38b14610466578063f79fd08e1461048c578063f856ddb6146104af57610151565b8063da87e44814610435578063e30c39781461045e57610151565b80639cdbb181116100b25780639cdbb181146103d6578063b2118a8d146103f7578063cb75c11c1461042d57610151565b806391f178881461033857806396abeb701461034057610151565b80636fd3504e116101245780638197beb9116101095780638197beb9146102d557806382a5e665146102fb5780638da5cb5b1461033057610151565b80636fd3504e1461027257806379ba5097146102cd57610151565b806329a78e33146101565780632ab60045146102205780632c1219211461024657806338a631831461026a575b600080fd5b61021e6004803603608081101561016c57600080fd5b81019060208101813564010000000081111561018757600080fd5b82018360208201111561019957600080fd5b803590602001918460018302840111640100000000831117156101bb57600080fd5b9193909290916020810190356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b9193509150803590602001356104f3565b005b61021e6004803603602081101561023657600080fd5b50356001600160a01b03166109a8565b61024e610a57565b604080516001600160a01b039092168252519081900360200190f35b61024e610a7b565b6102b06004803603608081101561028857600080fd5b50803590602081013563ffffffff1690604081013590606001356001600160a01b0316610a8a565b6040805167ffffffffffffffff9092168252519081900360200190f35b61021e610aa4565b61021e600480360360208110156102eb57600080fd5b50356001600160a01b0316610b13565b61031e6004803603602081101561031157600080fd5b503563ffffffff16610c40565b60408051918252519081900360200190f35b61024e610c52565b61021e610c61565b6103c26004803603606081101561035657600080fd5b63ffffffff8235169160208101359181019060608101604082013564010000000081111561038357600080fd5b82018360208201111561039557600080fd5b803590602001918460018302840111640100000000831117156103b757600080fd5b509092509050610d2e565b604080519115158252519081900360200190f35b6103de610f93565b6040805163ffffffff9092168252519081900360200190f35b61021e6004803603606081101561040d57600080fd5b506001600160a01b03813581169160208101359091169060400135610fb7565b61024e611019565b61021e6004803603604081101561044b57600080fd5b5063ffffffff8135169060200135611028565b61024e611140565b61021e6004803603602081101561047c57600080fd5b50356001600160a01b031661114f565b61021e600480360360208110156104a257600080fd5b503563ffffffff166111cd565b6102b0600480360360a08110156104c557600080fd5b5080359063ffffffff602082013516906040810135906001600160a01b036060820135169060800135611297565b6000610539600088888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505061136a9050565b905061054a62ffffff19821661138e565b600061055b62ffffff198316611460565b905061056c62ffffff19821661149f565b600061057d62ffffff198316611570565b905073__$b72828ef9544669316767c66e430328604$__635ced058e826040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156105ce57600080fd5b505af41580156105e2573d6000803e3d6000fd5b505050506040513d60208110156105f857600080fd5b50516001600160a01b03163314610656576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c69642073656e64657220666f72206d657373616765000000000000604482015290519081900360640190fd5b836106a8576040805162461bcd60e51b815260206004820152601e60248201527f4d696e7420726563697069656e74206d757374206265206e6f6e7a65726f0000604482015290519081900360640190fd5b60006106b962ffffff198416611585565b905060006106cc62ffffff19851661159a565b905060006106fd7f0000000000000000000000000000000000000000000000000000000000000000848985886115af565b90507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663b857b7748d8d8d8d868e6040518763ffffffff1660e01b81526004018080602001806020018060200185815260200184810384528a8a82818152602001925080828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016909101858103845288815260200190508888808284376000838201819052601f9091017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169092018681038452885181528851602091820193918a019250908190849084905b838110156108165781810151838201526020016107fe565b50505050905090810190601f1680156108435780820380516001836020036101000a031916815260200191505b509950505050505050505050600060405180830381600087803b15801561086957600080fd5b505af115801561087d573d6000803e3d6000fd5b50505050336001600160a01b031673__$b72828ef9544669316767c66e430328604$__635ced058e856040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156108da57600080fd5b505af41580156108ee573d6000803e3d6000fd5b505050506040513d602081101561090457600080fd5b50516001600160a01b031661091e62ffffff198916611612565b67ffffffffffffffff167f2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c0858b61095a62ffffff198d16611627565b61096962ffffff198e1661163c565b60408051948552602085019390935263ffffffff909116838301526060830152608082018e9052519081900360a00190a4505050505050505050505050565b6109b0611651565b6001600160a01b0381166109f55760405162461bcd60e51b815260040180806020018281038252602a815260200180612a41602a913960400191505060405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040517fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a90600090a250565b7f000000000000000000000000000000000000000000000000000000000000000081565b6002546001600160a01b031690565b6000610a9985858585856116c7565b90505b949350505050565b6000610aae611abf565b9050806001600160a01b0316610ac2611140565b6001600160a01b031614610b075760405162461bcd60e51b8152600401808060200182810382526029815260200180612a186029913960400191505060405180910390fd5b610b1081611ac3565b50565b610b1b611651565b6001600160a01b038116610b76576040805162461bcd60e51b815260206004820152601860248201527f5a65726f2061646472657373206e6f7420616c6c6f7765640000000000000000604482015290519081900360640190fd5b6003546001600160a01b031615610bd4576040805162461bcd60e51b815260206004820152601c60248201527f4c6f63616c206d696e74657220697320616c7265616479207365742e00000000604482015290519081900360640190fd5b600380546001600160a01b0383167fffffffffffffffffffffffff0000000000000000000000000000000000000000909116811790915560408051918252517f109bb3e70cbf1931e295b49e75c67013b85ff80d64e6f1d321f37157b90c38309181900360200190a150565b60046020526000908152604090205481565b6000546001600160a01b031690565b610c69611651565b6003546001600160a01b031680610cc7576040805162461bcd60e51b815260206004820152601760248201527f4e6f206c6f63616c206d696e746572206973207365742e000000000000000000604482015290519081900360640190fd5b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055604080516001600160a01b038316815290517f2db49fbf671271826a27b02ebc496209c85fffffb4bccc67430d2a0f22b4d1ac9181900360200190a150565b6000610d38611af4565b610d89576040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964206d657373616765207472616e736d69747465720000000000604482015290519081900360640190fd5b8484610d958282611b5b565b610dd05760405162461bcd60e51b8152600401808060200182810382526021815260200180612a6b6021913960400191505060405180910390fd5b6000610e16600087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505061136a9050565b9050610e2762ffffff19821661149f565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610e5d62ffffff198316611b87565b63ffffffff1614610eb5576040805162461bcd60e51b815260206004820152601c60248201527f496e76616c6964206d65737361676520626f64792076657273696f6e00000000604482015290519081900360640190fd5b6000610ec662ffffff198316611b9b565b90506000610ed962ffffff198416611585565b90506000610eec62ffffff19851661159a565b90506000610ef8611bb0565b9050610f81818d8573__$b72828ef9544669316767c66e430328604$__635ced058e896040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610f4f57600080fd5b505af4158015610f63573d6000803e3d6000fd5b505050506040513d6020811015610f7957600080fd5b505186611c20565b5060019b9a5050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b6002546001600160a01b031633146110005760405162461bcd60e51b8152600401808060200182810382526024815260200180612ad36024913960400191505060405180910390fd5b6110146001600160a01b0384168383611d1f565b505050565b6003546001600160a01b031681565b611030611651565b80611082576040805162461bcd60e51b815260206004820152601660248201527f62797465733332283029206e6f7420616c6c6f77656400000000000000000000604482015290519081900360640190fd5b63ffffffff8216600090815260046020526040902054156110ea576040805162461bcd60e51b815260206004820152601a60248201527f546f6b656e4d657373656e67657220616c726561647920736574000000000000604482015290519081900360640190fd5b63ffffffff82166000818152600460209081526040918290208490558151928352820183905280517f4bba2b08298cf59661b4895e384cc2ac3962ce2d71f1b7c11bca52e1169f95999281900390910190a15050565b6001546001600160a01b031690565b611157611651565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038316908117909155611195610c52565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6111d5611651565b63ffffffff811660009081526004602052604090205461123c576040805162461bcd60e51b815260206004820152601560248201527f4e6f20546f6b656e4d657373656e676572207365740000000000000000000000604482015290519081900360640190fd5b63ffffffff8116600081815260046020908152604080832080549390558051938452908301829052805191927f3dcea012093dbca2bb8ed7fd2b2ff90305ab70bddda8bbb94d4152735a98f0b1929081900390910190a15050565b6000816112eb576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c69642064657374696e6174696f6e2063616c6c6572000000000000604482015290519081900360640190fd5b6112f886868686866116c7565b9695505050505050565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b81516000906020840161138564ffffffffff85168284611d9f565b95945050505050565b61139d62ffffff198216611dd9565b6113ee576040805162461bcd60e51b815260206004820152601160248201527f4d616c666f726d6564206d657373616765000000000000000000000000000000604482015290519081900360640190fd5b60746113ff62ffffff198316611e16565b6bffffffffffffffffffffffff161015610b10576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c6964206d6573736167653a20746f6f2073686f7274000000000000604482015290519081900360640190fd5b600061149760748061147762ffffff198616611e16565b62ffffff19861692916bffffffffffffffffffffffff9103166000611e2a565b90505b919050565b6114ae62ffffff198216611dd9565b6114ff576040805162461bcd60e51b815260206004820152601160248201527f4d616c666f726d6564206d657373616765000000000000000000000000000000604482015290519081900360640190fd5b608461151062ffffff198316611e16565b6bffffffffffffffffffffffff1614610b10576040805162461bcd60e51b815260206004820152601660248201527f496e76616c6964206d657373616765206c656e67746800000000000000000000604482015290519081900360640190fd5b600061149762ffffff19831660646020611e94565b600061149762ffffff19831660046020611e94565b600061149762ffffff1983166044602061200b565b6040805160e09690961b7fffffffff000000000000000000000000000000000000000000000000000000001660208701526024860194909452604485019290925260648401526084808401919091528151808403909101815260a4909201905290565b600061149762ffffff198316600c600861200b565b600061149762ffffff1983166008600461200b565b600061149762ffffff19831660346020611e94565b611659611abf565b6001600160a01b031661166a610c52565b6001600160a01b0316146116c5576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b565b600080861161171d576040805162461bcd60e51b815260206004820152601660248201527f416d6f756e74206d757374206265206e6f6e7a65726f00000000000000000000604482015290519081900360640190fd5b8361176f576040805162461bcd60e51b815260206004820152601e60248201527f4d696e7420726563697069656e74206d757374206265206e6f6e7a65726f0000604482015290519081900360640190fd5b600061177a8661202c565b90506000611786611bb0565b604080517f23b872dd0000000000000000000000000000000000000000000000000000000081523360048201526001600160a01b038084166024830152604482018c905291519293508792918316916323b872dd916064808201926020929091908290030181600087803b1580156117fd57600080fd5b505af1158015611811573d6000803e3d6000fd5b505050506040513d602081101561182757600080fd5b505161187a576040805162461bcd60e51b815260206004820152601960248201527f5472616e73666572206f7065726174696f6e206661696c656400000000000000604482015290519081900360640190fd5b816001600160a01b0316639dc29fac878b6040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156118d157600080fd5b505af11580156118e5573d6000803e3d6000fd5b505050506000611a327f000000000000000000000000000000000000000000000000000000000000000073__$b72828ef9544669316767c66e430328604$__6382c947b78a6040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561196757600080fd5b505af415801561197b573d6000803e3d6000fd5b505050506040513d602081101561199157600080fd5b5051604080517f82c947b700000000000000000000000000000000000000000000000000000000815233600482015290518c918f9173__$b72828ef9544669316767c66e430328604$__916382c947b7916024808301926020929190829003018186803b158015611a0157600080fd5b505af4158015611a15573d6000803e3d6000fd5b505050506040513d6020811015611a2b57600080fd5b50516115af565b90506000611a428a868985612094565b604080518d8152602081018c905263ffffffff8d168183015260608101889052608081018a9052905191925033916001600160a01b038b169167ffffffffffffffff8516917f2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c09181900360a00190a49a9950505050505050505050565b3390565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055610b1081611302565b60007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031615801590611b565750336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016145b905090565b60008115801590611b80575063ffffffff831660009081526004602052604090205482145b9392505050565b600061149762ffffff19831682600461200b565b600061149762ffffff19831660246020611e94565b6003546000906001600160a01b0316611c10576040805162461bcd60e51b815260206004820152601760248201527f4c6f63616c206d696e746572206973206e6f7420736574000000000000000000604482015290519081900360640190fd5b506003546001600160a01b031690565b604080517fd54de06f00000000000000000000000000000000000000000000000000000000815263ffffffff86166004820152602481018590526001600160a01b03848116604483015260648201849052915187926000929084169163d54de06f9160848082019260209290919082900301818787803b158015611ca357600080fd5b505af1158015611cb7573d6000803e3d6000fd5b505050506040513d6020811015611ccd57600080fd5b50516040805185815290519192506001600160a01b0380841692908716917f1b2a7ff080b8cb6ff436ce0372e399692bbfb6d4ae5766fd8d58a7b8cc6142e6919081900360200190a350505050505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526110149084906122ad565b600080611dac848461235e565b9050604051811115611dbc575060005b80611dce5762ffffff19915050611b80565b6113858585856123b8565b6000611de4826123cb565b64ffffffffff1664ffffffffff1415611dff5750600061149a565b6000611e0a836123d1565b60405110199392505050565b60181c6bffffffffffffffffffffffff1690565b600080611e36866123fb565b6bffffffffffffffffffffffff169050611e4f866123d1565b611e6385611e5d848961235e565b9061235e565b1115611e765762ffffff19915050610a9c565b611e80818661235e565b90506112f88364ffffffffff168286611d9f565b600060ff8216611ea657506000611b80565b611eaf84611e16565b6bffffffffffffffffffffffff16611eca8460ff851661235e565b1115611f8f57611f0b611edc856123fb565b6bffffffffffffffffffffffff16611ef386611e16565b6bffffffffffffffffffffffff16858560ff1661240f565b60405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611f54578181015183820152602001611f3c565b50505050905090810190601f168015611f815780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b60208260ff161115611fd25760405162461bcd60e51b815260040180806020018281038252603a815260200180612af7603a913960400191505060405180910390fd5b600882026000611fe1866123fb565b6bffffffffffffffffffffffff1690506000611ffc8361256a565b91909501511695945050505050565b60008160200360080260ff16612022858585611e94565b901c949350505050565b63ffffffff811660009081526004602052604081205480611497576040805162461bcd60e51b815260206004820152601c60248201527f4e6f20546f6b656e4d657373656e67657220666f7220646f6d61696e00000000604482015290519081900360640190fd5b6000826121c7576040517f0ba469bc00000000000000000000000000000000000000000000000000000000815263ffffffff861660048201908152602482018690526060604483019081528451606484015284516001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001693630ba469bc938a938a93899360840190602085019080838360005b8381101561214657818101518382015260200161212e565b50505050905090810190601f1680156121735780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561219457600080fd5b505af11580156121a8573d6000803e3d6000fd5b505050506040513d60208110156121be57600080fd5b50519050610a9c565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663f7259a75868686866040518563ffffffff1660e01b8152600401808563ffffffff16815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561225e578181015183820152602001612246565b50505050905090810190601f16801561228b5780820380516001836020036101000a031916815260200191505b5095505050505050602060405180830381600087803b15801561219457600080fd5b6000612302826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166125b39092919063ffffffff16565b8051909150156110145780806020019051602081101561232157600080fd5b50516110145760405162461bcd60e51b815260040180806020018281038252602a815260200180612b31602a913960400191505060405180910390fd5b600082820183811015611b80576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b606092831b9190911790911b1760181b90565b60d81c90565b60006123dc82611e16565b6123e5836123fb565b016bffffffffffffffffffffffff169050919050565b60781c6bffffffffffffffffffffffff1690565b6060600061241c866125c2565b915050600061242a866125c2565b9150506000612438866125c2565b9150506000612446866125c2565b915050838383836040516020018080612b5b603591397fffffffffffff000000000000000000000000000000000000000000000000000060d087811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b84015286901b16604a8201526050016021612ab282397fffffffffffff000000000000000000000000000000000000000000000000000060d094851b811660218301527f2077697468206c656e677468203078000000000000000000000000000000000060278301529290931b9091166036830152507f2e00000000000000000000000000000000000000000000000000000000000000603c82015260408051601d818403018152603d90920190529b9a5050505050505050505050565b7f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9091011d90565b6060610a9c8484600085612696565b600080601f5b600f8160ff16111561262a5760ff600882021684901c6125e78161280f565b61ffff16841793508160ff1660101461260257601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016125c8565b50600f5b60ff8160ff1610156126905760ff600882021684901c61264d8161280f565b61ffff16831792508160ff1660001461266857601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161262e565b50915091565b6060824710156126d75760405162461bcd60e51b8152600401808060200182810382526026815260200180612a8c6026913960400191505060405180910390fd5b6126e08561283f565b612731576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b6020831061278d57805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101612750565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146127ef576040519150601f19603f3d011682016040523d82523d6000602084013e6127f4565b606091505b5091509150612804828286612845565b979650505050505050565b600061282160048360ff16901c6128ab565b60ff161760081b62ffff0016612836826128ab565b60ff1617919050565b3b151590565b60608315612854575081611b80565b8251156128645782518084602001fd5b60405162461bcd60e51b8152602060048201818152845160248401528451859391928392604401919085019080838360008315611f54578181015183820152602001611f3c565b600060f08083179060ff821614156128c757603091505061149a565b8060ff1660f114156128dd57603191505061149a565b8060ff1660f214156128f357603291505061149a565b8060ff1660f3141561290957603391505061149a565b8060ff1660f4141561291f57603491505061149a565b8060ff1660f5141561293557603591505061149a565b8060ff1660f6141561294b57603691505061149a565b8060ff1660f7141561296157603791505061149a565b8060ff1660f8141561297757603891505061149a565b8060ff1660f9141561298d57603991505061149a565b8060ff1660fa14156129a357606191505061149a565b8060ff1660fb14156129b957606291505061149a565b8060ff1660fc14156129cf57606391505061149a565b8060ff1660fd14156129e557606491505061149a565b8060ff1660fe14156129fb57606591505061149a565b8060ff1660ff1415612a1157606691505061149a565b5091905056fe4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f206164647265737352656d6f746520546f6b656e4d657373656e67657220756e737570706f72746564416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c2e20417474656d7074656420746f20696e646578206174206f6666736574203078526573637561626c653a2063616c6c6572206973206e6f7420746865207265736375657254797065644d656d566965772f696e646578202d20417474656d7074656420746f20696e646578206d6f7265207468616e2033322062797465735361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656454797065644d656d566965772f696e646578202d204f76657272616e2074686520766965772e20536c696365206973206174203078a26469706673582212202c33f5cd5a19bb15b4d28742277e10a8d7140893dd90aaecef8a44d80cbf8fd264736f6c63430007060033",
}

// TokenMessengerABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenMessengerMetaData.ABI instead.
var TokenMessengerABI = TokenMessengerMetaData.ABI

// Deprecated: Use TokenMessengerMetaData.Sigs instead.
// TokenMessengerFuncSigs maps the 4-byte function signature to its string representation.
var TokenMessengerFuncSigs = TokenMessengerMetaData.Sigs

// TokenMessengerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenMessengerMetaData.Bin instead.
var TokenMessengerBin = TokenMessengerMetaData.Bin

// DeployTokenMessenger deploys a new Ethereum contract, binding an instance of TokenMessenger to it.
func DeployTokenMessenger(auth *bind.TransactOpts, backend bind.ContractBackend, _messageTransmitter common.Address, _messageBodyVersion uint32) (common.Address, *types.Transaction, *TokenMessenger, error) {
	parsed, err := TokenMessengerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	messageAddr, _, _, _ := DeployMessage(auth, backend)
	TokenMessengerBin = strings.ReplaceAll(TokenMessengerBin, "__$b72828ef9544669316767c66e430328604$__", messageAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenMessengerBin), backend, _messageTransmitter, _messageBodyVersion)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenMessenger{TokenMessengerCaller: TokenMessengerCaller{contract: contract}, TokenMessengerTransactor: TokenMessengerTransactor{contract: contract}, TokenMessengerFilterer: TokenMessengerFilterer{contract: contract}}, nil
}

// TokenMessenger is an auto generated Go binding around an Ethereum contract.
type TokenMessenger struct {
	TokenMessengerCaller     // Read-only binding to the contract
	TokenMessengerTransactor // Write-only binding to the contract
	TokenMessengerFilterer   // Log filterer for contract events
}

// TokenMessengerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenMessengerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMessengerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenMessengerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMessengerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenMessengerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenMessengerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenMessengerSession struct {
	Contract     *TokenMessenger   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenMessengerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenMessengerCallerSession struct {
	Contract *TokenMessengerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenMessengerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenMessengerTransactorSession struct {
	Contract     *TokenMessengerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenMessengerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenMessengerRaw struct {
	Contract *TokenMessenger // Generic contract binding to access the raw methods on
}

// TokenMessengerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenMessengerCallerRaw struct {
	Contract *TokenMessengerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenMessengerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenMessengerTransactorRaw struct {
	Contract *TokenMessengerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenMessenger creates a new instance of TokenMessenger, bound to a specific deployed contract.
func NewTokenMessenger(address common.Address, backend bind.ContractBackend) (*TokenMessenger, error) {
	contract, err := bindTokenMessenger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenMessenger{TokenMessengerCaller: TokenMessengerCaller{contract: contract}, TokenMessengerTransactor: TokenMessengerTransactor{contract: contract}, TokenMessengerFilterer: TokenMessengerFilterer{contract: contract}}, nil
}

// NewTokenMessengerCaller creates a new read-only instance of TokenMessenger, bound to a specific deployed contract.
func NewTokenMessengerCaller(address common.Address, caller bind.ContractCaller) (*TokenMessengerCaller, error) {
	contract, err := bindTokenMessenger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerCaller{contract: contract}, nil
}

// NewTokenMessengerTransactor creates a new write-only instance of TokenMessenger, bound to a specific deployed contract.
func NewTokenMessengerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenMessengerTransactor, error) {
	contract, err := bindTokenMessenger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerTransactor{contract: contract}, nil
}

// NewTokenMessengerFilterer creates a new log filterer instance of TokenMessenger, bound to a specific deployed contract.
func NewTokenMessengerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenMessengerFilterer, error) {
	contract, err := bindTokenMessenger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerFilterer{contract: contract}, nil
}

// bindTokenMessenger binds a generic wrapper to an already deployed contract.
func bindTokenMessenger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenMessengerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenMessenger *TokenMessengerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenMessenger.Contract.TokenMessengerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenMessenger *TokenMessengerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMessenger.Contract.TokenMessengerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenMessenger *TokenMessengerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenMessenger.Contract.TokenMessengerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenMessenger *TokenMessengerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenMessenger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenMessenger *TokenMessengerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMessenger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenMessenger *TokenMessengerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenMessenger.Contract.contract.Transact(opts, method, params...)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_TokenMessenger *TokenMessengerCaller) LocalMessageTransmitter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenMessenger.contract.Call(opts, &out, "localMessageTransmitter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_TokenMessenger *TokenMessengerSession) LocalMessageTransmitter() (common.Address, error) {
	return _TokenMessenger.Contract.LocalMessageTransmitter(&_TokenMessenger.CallOpts)
}

// LocalMessageTransmitter is a free data retrieval call binding the contract method 0x2c121921.
//
// Solidity: function localMessageTransmitter() view returns(address)
func (_TokenMessenger *TokenMessengerCallerSession) LocalMessageTransmitter() (common.Address, error) {
	return _TokenMessenger.Contract.LocalMessageTransmitter(&_TokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_TokenMessenger *TokenMessengerCaller) LocalMinter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenMessenger.contract.Call(opts, &out, "localMinter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_TokenMessenger *TokenMessengerSession) LocalMinter() (common.Address, error) {
	return _TokenMessenger.Contract.LocalMinter(&_TokenMessenger.CallOpts)
}

// LocalMinter is a free data retrieval call binding the contract method 0xcb75c11c.
//
// Solidity: function localMinter() view returns(address)
func (_TokenMessenger *TokenMessengerCallerSession) LocalMinter() (common.Address, error) {
	return _TokenMessenger.Contract.LocalMinter(&_TokenMessenger.CallOpts)
}

// MessageBodyVersion is a free data retrieval call binding the contract method 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (_TokenMessenger *TokenMessengerCaller) MessageBodyVersion(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _TokenMessenger.contract.Call(opts, &out, "messageBodyVersion")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MessageBodyVersion is a free data retrieval call binding the contract method 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (_TokenMessenger *TokenMessengerSession) MessageBodyVersion() (uint32, error) {
	return _TokenMessenger.Contract.MessageBodyVersion(&_TokenMessenger.CallOpts)
}

// MessageBodyVersion is a free data retrieval call binding the contract method 0x9cdbb181.
//
// Solidity: function messageBodyVersion() view returns(uint32)
func (_TokenMessenger *TokenMessengerCallerSession) MessageBodyVersion() (uint32, error) {
	return _TokenMessenger.Contract.MessageBodyVersion(&_TokenMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenMessenger *TokenMessengerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenMessenger.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenMessenger *TokenMessengerSession) Owner() (common.Address, error) {
	return _TokenMessenger.Contract.Owner(&_TokenMessenger.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TokenMessenger *TokenMessengerCallerSession) Owner() (common.Address, error) {
	return _TokenMessenger.Contract.Owner(&_TokenMessenger.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenMessenger *TokenMessengerCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenMessenger.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenMessenger *TokenMessengerSession) PendingOwner() (common.Address, error) {
	return _TokenMessenger.Contract.PendingOwner(&_TokenMessenger.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_TokenMessenger *TokenMessengerCallerSession) PendingOwner() (common.Address, error) {
	return _TokenMessenger.Contract.PendingOwner(&_TokenMessenger.CallOpts)
}

// RemoteTokenMessengers is a free data retrieval call binding the contract method 0x82a5e665.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (_TokenMessenger *TokenMessengerCaller) RemoteTokenMessengers(opts *bind.CallOpts, arg0 uint32) ([32]byte, error) {
	var out []interface{}
	err := _TokenMessenger.contract.Call(opts, &out, "remoteTokenMessengers", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RemoteTokenMessengers is a free data retrieval call binding the contract method 0x82a5e665.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (_TokenMessenger *TokenMessengerSession) RemoteTokenMessengers(arg0 uint32) ([32]byte, error) {
	return _TokenMessenger.Contract.RemoteTokenMessengers(&_TokenMessenger.CallOpts, arg0)
}

// RemoteTokenMessengers is a free data retrieval call binding the contract method 0x82a5e665.
//
// Solidity: function remoteTokenMessengers(uint32 ) view returns(bytes32)
func (_TokenMessenger *TokenMessengerCallerSession) RemoteTokenMessengers(arg0 uint32) ([32]byte, error) {
	return _TokenMessenger.Contract.RemoteTokenMessengers(&_TokenMessenger.CallOpts, arg0)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_TokenMessenger *TokenMessengerCaller) Rescuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TokenMessenger.contract.Call(opts, &out, "rescuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_TokenMessenger *TokenMessengerSession) Rescuer() (common.Address, error) {
	return _TokenMessenger.Contract.Rescuer(&_TokenMessenger.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_TokenMessenger *TokenMessengerCallerSession) Rescuer() (common.Address, error) {
	return _TokenMessenger.Contract.Rescuer(&_TokenMessenger.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenMessenger *TokenMessengerTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenMessenger *TokenMessengerSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenMessenger.Contract.AcceptOwnership(&_TokenMessenger.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_TokenMessenger *TokenMessengerTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _TokenMessenger.Contract.AcceptOwnership(&_TokenMessenger.TransactOpts)
}

// AddLocalMinter is a paid mutator transaction binding the contract method 0x8197beb9.
//
// Solidity: function addLocalMinter(address newLocalMinter) returns()
func (_TokenMessenger *TokenMessengerTransactor) AddLocalMinter(opts *bind.TransactOpts, newLocalMinter common.Address) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "addLocalMinter", newLocalMinter)
}

// AddLocalMinter is a paid mutator transaction binding the contract method 0x8197beb9.
//
// Solidity: function addLocalMinter(address newLocalMinter) returns()
func (_TokenMessenger *TokenMessengerSession) AddLocalMinter(newLocalMinter common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.AddLocalMinter(&_TokenMessenger.TransactOpts, newLocalMinter)
}

// AddLocalMinter is a paid mutator transaction binding the contract method 0x8197beb9.
//
// Solidity: function addLocalMinter(address newLocalMinter) returns()
func (_TokenMessenger *TokenMessengerTransactorSession) AddLocalMinter(newLocalMinter common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.AddLocalMinter(&_TokenMessenger.TransactOpts, newLocalMinter)
}

// AddRemoteTokenMessenger is a paid mutator transaction binding the contract method 0xda87e448.
//
// Solidity: function addRemoteTokenMessenger(uint32 domain, bytes32 tokenMessenger) returns()
func (_TokenMessenger *TokenMessengerTransactor) AddRemoteTokenMessenger(opts *bind.TransactOpts, domain uint32, tokenMessenger [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "addRemoteTokenMessenger", domain, tokenMessenger)
}

// AddRemoteTokenMessenger is a paid mutator transaction binding the contract method 0xda87e448.
//
// Solidity: function addRemoteTokenMessenger(uint32 domain, bytes32 tokenMessenger) returns()
func (_TokenMessenger *TokenMessengerSession) AddRemoteTokenMessenger(domain uint32, tokenMessenger [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.AddRemoteTokenMessenger(&_TokenMessenger.TransactOpts, domain, tokenMessenger)
}

// AddRemoteTokenMessenger is a paid mutator transaction binding the contract method 0xda87e448.
//
// Solidity: function addRemoteTokenMessenger(uint32 domain, bytes32 tokenMessenger) returns()
func (_TokenMessenger *TokenMessengerTransactorSession) AddRemoteTokenMessenger(domain uint32, tokenMessenger [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.AddRemoteTokenMessenger(&_TokenMessenger.TransactOpts, domain, tokenMessenger)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken) returns(uint64 _nonce)
func (_TokenMessenger *TokenMessengerTransactor) DepositForBurn(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "depositForBurn", amount, destinationDomain, mintRecipient, burnToken)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken) returns(uint64 _nonce)
func (_TokenMessenger *TokenMessengerSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.DepositForBurn(&_TokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken)
}

// DepositForBurn is a paid mutator transaction binding the contract method 0x6fd3504e.
//
// Solidity: function depositForBurn(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken) returns(uint64 _nonce)
func (_TokenMessenger *TokenMessengerTransactorSession) DepositForBurn(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.DepositForBurn(&_TokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_TokenMessenger *TokenMessengerTransactor) DepositForBurnWithCaller(opts *bind.TransactOpts, amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "depositForBurnWithCaller", amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_TokenMessenger *TokenMessengerSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.DepositForBurnWithCaller(&_TokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// DepositForBurnWithCaller is a paid mutator transaction binding the contract method 0xf856ddb6.
//
// Solidity: function depositForBurnWithCaller(uint256 amount, uint32 destinationDomain, bytes32 mintRecipient, address burnToken, bytes32 destinationCaller) returns(uint64 nonce)
func (_TokenMessenger *TokenMessengerTransactorSession) DepositForBurnWithCaller(amount *big.Int, destinationDomain uint32, mintRecipient [32]byte, burnToken common.Address, destinationCaller [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.DepositForBurnWithCaller(&_TokenMessenger.TransactOpts, amount, destinationDomain, mintRecipient, burnToken, destinationCaller)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_TokenMessenger *TokenMessengerTransactor) HandleReceiveMessage(opts *bind.TransactOpts, remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "handleReceiveMessage", remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_TokenMessenger *TokenMessengerSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.HandleReceiveMessage(&_TokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 remoteDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_TokenMessenger *TokenMessengerTransactorSession) HandleReceiveMessage(remoteDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.HandleReceiveMessage(&_TokenMessenger.TransactOpts, remoteDomain, sender, messageBody)
}

// RemoveLocalMinter is a paid mutator transaction binding the contract method 0x91f17888.
//
// Solidity: function removeLocalMinter() returns()
func (_TokenMessenger *TokenMessengerTransactor) RemoveLocalMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "removeLocalMinter")
}

// RemoveLocalMinter is a paid mutator transaction binding the contract method 0x91f17888.
//
// Solidity: function removeLocalMinter() returns()
func (_TokenMessenger *TokenMessengerSession) RemoveLocalMinter() (*types.Transaction, error) {
	return _TokenMessenger.Contract.RemoveLocalMinter(&_TokenMessenger.TransactOpts)
}

// RemoveLocalMinter is a paid mutator transaction binding the contract method 0x91f17888.
//
// Solidity: function removeLocalMinter() returns()
func (_TokenMessenger *TokenMessengerTransactorSession) RemoveLocalMinter() (*types.Transaction, error) {
	return _TokenMessenger.Contract.RemoveLocalMinter(&_TokenMessenger.TransactOpts)
}

// RemoveRemoteTokenMessenger is a paid mutator transaction binding the contract method 0xf79fd08e.
//
// Solidity: function removeRemoteTokenMessenger(uint32 domain) returns()
func (_TokenMessenger *TokenMessengerTransactor) RemoveRemoteTokenMessenger(opts *bind.TransactOpts, domain uint32) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "removeRemoteTokenMessenger", domain)
}

// RemoveRemoteTokenMessenger is a paid mutator transaction binding the contract method 0xf79fd08e.
//
// Solidity: function removeRemoteTokenMessenger(uint32 domain) returns()
func (_TokenMessenger *TokenMessengerSession) RemoveRemoteTokenMessenger(domain uint32) (*types.Transaction, error) {
	return _TokenMessenger.Contract.RemoveRemoteTokenMessenger(&_TokenMessenger.TransactOpts, domain)
}

// RemoveRemoteTokenMessenger is a paid mutator transaction binding the contract method 0xf79fd08e.
//
// Solidity: function removeRemoteTokenMessenger(uint32 domain) returns()
func (_TokenMessenger *TokenMessengerTransactorSession) RemoveRemoteTokenMessenger(domain uint32) (*types.Transaction, error) {
	return _TokenMessenger.Contract.RemoveRemoteTokenMessenger(&_TokenMessenger.TransactOpts, domain)
}

// ReplaceDepositForBurn is a paid mutator transaction binding the contract method 0x29a78e33.
//
// Solidity: function replaceDepositForBurn(bytes originalMessage, bytes originalAttestation, bytes32 newDestinationCaller, bytes32 newMintRecipient) returns()
func (_TokenMessenger *TokenMessengerTransactor) ReplaceDepositForBurn(opts *bind.TransactOpts, originalMessage []byte, originalAttestation []byte, newDestinationCaller [32]byte, newMintRecipient [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "replaceDepositForBurn", originalMessage, originalAttestation, newDestinationCaller, newMintRecipient)
}

// ReplaceDepositForBurn is a paid mutator transaction binding the contract method 0x29a78e33.
//
// Solidity: function replaceDepositForBurn(bytes originalMessage, bytes originalAttestation, bytes32 newDestinationCaller, bytes32 newMintRecipient) returns()
func (_TokenMessenger *TokenMessengerSession) ReplaceDepositForBurn(originalMessage []byte, originalAttestation []byte, newDestinationCaller [32]byte, newMintRecipient [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.ReplaceDepositForBurn(&_TokenMessenger.TransactOpts, originalMessage, originalAttestation, newDestinationCaller, newMintRecipient)
}

// ReplaceDepositForBurn is a paid mutator transaction binding the contract method 0x29a78e33.
//
// Solidity: function replaceDepositForBurn(bytes originalMessage, bytes originalAttestation, bytes32 newDestinationCaller, bytes32 newMintRecipient) returns()
func (_TokenMessenger *TokenMessengerTransactorSession) ReplaceDepositForBurn(originalMessage []byte, originalAttestation []byte, newDestinationCaller [32]byte, newMintRecipient [32]byte) (*types.Transaction, error) {
	return _TokenMessenger.Contract.ReplaceDepositForBurn(&_TokenMessenger.TransactOpts, originalMessage, originalAttestation, newDestinationCaller, newMintRecipient)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_TokenMessenger *TokenMessengerTransactor) RescueERC20(opts *bind.TransactOpts, tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "rescueERC20", tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_TokenMessenger *TokenMessengerSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMessenger.Contract.RescueERC20(&_TokenMessenger.TransactOpts, tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_TokenMessenger *TokenMessengerTransactorSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenMessenger.Contract.RescueERC20(&_TokenMessenger.TransactOpts, tokenContract, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenMessenger *TokenMessengerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenMessenger *TokenMessengerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.TransferOwnership(&_TokenMessenger.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TokenMessenger *TokenMessengerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.TransferOwnership(&_TokenMessenger.TransactOpts, newOwner)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_TokenMessenger *TokenMessengerTransactor) UpdateRescuer(opts *bind.TransactOpts, newRescuer common.Address) (*types.Transaction, error) {
	return _TokenMessenger.contract.Transact(opts, "updateRescuer", newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_TokenMessenger *TokenMessengerSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.UpdateRescuer(&_TokenMessenger.TransactOpts, newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_TokenMessenger *TokenMessengerTransactorSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _TokenMessenger.Contract.UpdateRescuer(&_TokenMessenger.TransactOpts, newRescuer)
}

// TokenMessengerDepositForBurnIterator is returned from FilterDepositForBurn and is used to iterate over the raw logs and unpacked data for DepositForBurn events raised by the TokenMessenger contract.
type TokenMessengerDepositForBurnIterator struct {
	Event *TokenMessengerDepositForBurn // Event containing the contract specifics and raw log

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
func (it *TokenMessengerDepositForBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerDepositForBurn)
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
		it.Event = new(TokenMessengerDepositForBurn)
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
func (it *TokenMessengerDepositForBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerDepositForBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerDepositForBurn represents a DepositForBurn event raised by the TokenMessenger contract.
type TokenMessengerDepositForBurn struct {
	Nonce                     uint64
	BurnToken                 common.Address
	Amount                    *big.Int
	Depositor                 common.Address
	MintRecipient             [32]byte
	DestinationDomain         uint32
	DestinationTokenMessenger [32]byte
	DestinationCaller         [32]byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterDepositForBurn is a free log retrieval operation binding the contract event 0x2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c0.
//
// Solidity: event DepositForBurn(uint64 indexed nonce, address indexed burnToken, uint256 amount, address indexed depositor, bytes32 mintRecipient, uint32 destinationDomain, bytes32 destinationTokenMessenger, bytes32 destinationCaller)
func (_TokenMessenger *TokenMessengerFilterer) FilterDepositForBurn(opts *bind.FilterOpts, nonce []uint64, burnToken []common.Address, depositor []common.Address) (*TokenMessengerDepositForBurnIterator, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerDepositForBurnIterator{contract: _TokenMessenger.contract, event: "DepositForBurn", logs: logs, sub: sub}, nil
}

// WatchDepositForBurn is a free log subscription operation binding the contract event 0x2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c0.
//
// Solidity: event DepositForBurn(uint64 indexed nonce, address indexed burnToken, uint256 amount, address indexed depositor, bytes32 mintRecipient, uint32 destinationDomain, bytes32 destinationTokenMessenger, bytes32 destinationCaller)
func (_TokenMessenger *TokenMessengerFilterer) WatchDepositForBurn(opts *bind.WatchOpts, sink chan<- *TokenMessengerDepositForBurn, nonce []uint64, burnToken []common.Address, depositor []common.Address) (event.Subscription, error) {

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}
	var burnTokenRule []interface{}
	for _, burnTokenItem := range burnToken {
		burnTokenRule = append(burnTokenRule, burnTokenItem)
	}

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "DepositForBurn", nonceRule, burnTokenRule, depositorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerDepositForBurn)
				if err := _TokenMessenger.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
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

// ParseDepositForBurn is a log parse operation binding the contract event 0x2fa9ca894982930190727e75500a97d8dc500233a5065e0f3126c48fbe0343c0.
//
// Solidity: event DepositForBurn(uint64 indexed nonce, address indexed burnToken, uint256 amount, address indexed depositor, bytes32 mintRecipient, uint32 destinationDomain, bytes32 destinationTokenMessenger, bytes32 destinationCaller)
func (_TokenMessenger *TokenMessengerFilterer) ParseDepositForBurn(log types.Log) (*TokenMessengerDepositForBurn, error) {
	event := new(TokenMessengerDepositForBurn)
	if err := _TokenMessenger.contract.UnpackLog(event, "DepositForBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerLocalMinterAddedIterator is returned from FilterLocalMinterAdded and is used to iterate over the raw logs and unpacked data for LocalMinterAdded events raised by the TokenMessenger contract.
type TokenMessengerLocalMinterAddedIterator struct {
	Event *TokenMessengerLocalMinterAdded // Event containing the contract specifics and raw log

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
func (it *TokenMessengerLocalMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerLocalMinterAdded)
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
		it.Event = new(TokenMessengerLocalMinterAdded)
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
func (it *TokenMessengerLocalMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerLocalMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerLocalMinterAdded represents a LocalMinterAdded event raised by the TokenMessenger contract.
type TokenMessengerLocalMinterAdded struct {
	LocalMinter common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLocalMinterAdded is a free log retrieval operation binding the contract event 0x109bb3e70cbf1931e295b49e75c67013b85ff80d64e6f1d321f37157b90c3830.
//
// Solidity: event LocalMinterAdded(address localMinter)
func (_TokenMessenger *TokenMessengerFilterer) FilterLocalMinterAdded(opts *bind.FilterOpts) (*TokenMessengerLocalMinterAddedIterator, error) {

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "LocalMinterAdded")
	if err != nil {
		return nil, err
	}
	return &TokenMessengerLocalMinterAddedIterator{contract: _TokenMessenger.contract, event: "LocalMinterAdded", logs: logs, sub: sub}, nil
}

// WatchLocalMinterAdded is a free log subscription operation binding the contract event 0x109bb3e70cbf1931e295b49e75c67013b85ff80d64e6f1d321f37157b90c3830.
//
// Solidity: event LocalMinterAdded(address localMinter)
func (_TokenMessenger *TokenMessengerFilterer) WatchLocalMinterAdded(opts *bind.WatchOpts, sink chan<- *TokenMessengerLocalMinterAdded) (event.Subscription, error) {

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "LocalMinterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerLocalMinterAdded)
				if err := _TokenMessenger.contract.UnpackLog(event, "LocalMinterAdded", log); err != nil {
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

// ParseLocalMinterAdded is a log parse operation binding the contract event 0x109bb3e70cbf1931e295b49e75c67013b85ff80d64e6f1d321f37157b90c3830.
//
// Solidity: event LocalMinterAdded(address localMinter)
func (_TokenMessenger *TokenMessengerFilterer) ParseLocalMinterAdded(log types.Log) (*TokenMessengerLocalMinterAdded, error) {
	event := new(TokenMessengerLocalMinterAdded)
	if err := _TokenMessenger.contract.UnpackLog(event, "LocalMinterAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerLocalMinterRemovedIterator is returned from FilterLocalMinterRemoved and is used to iterate over the raw logs and unpacked data for LocalMinterRemoved events raised by the TokenMessenger contract.
type TokenMessengerLocalMinterRemovedIterator struct {
	Event *TokenMessengerLocalMinterRemoved // Event containing the contract specifics and raw log

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
func (it *TokenMessengerLocalMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerLocalMinterRemoved)
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
		it.Event = new(TokenMessengerLocalMinterRemoved)
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
func (it *TokenMessengerLocalMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerLocalMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerLocalMinterRemoved represents a LocalMinterRemoved event raised by the TokenMessenger contract.
type TokenMessengerLocalMinterRemoved struct {
	LocalMinter common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLocalMinterRemoved is a free log retrieval operation binding the contract event 0x2db49fbf671271826a27b02ebc496209c85fffffb4bccc67430d2a0f22b4d1ac.
//
// Solidity: event LocalMinterRemoved(address localMinter)
func (_TokenMessenger *TokenMessengerFilterer) FilterLocalMinterRemoved(opts *bind.FilterOpts) (*TokenMessengerLocalMinterRemovedIterator, error) {

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "LocalMinterRemoved")
	if err != nil {
		return nil, err
	}
	return &TokenMessengerLocalMinterRemovedIterator{contract: _TokenMessenger.contract, event: "LocalMinterRemoved", logs: logs, sub: sub}, nil
}

// WatchLocalMinterRemoved is a free log subscription operation binding the contract event 0x2db49fbf671271826a27b02ebc496209c85fffffb4bccc67430d2a0f22b4d1ac.
//
// Solidity: event LocalMinterRemoved(address localMinter)
func (_TokenMessenger *TokenMessengerFilterer) WatchLocalMinterRemoved(opts *bind.WatchOpts, sink chan<- *TokenMessengerLocalMinterRemoved) (event.Subscription, error) {

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "LocalMinterRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerLocalMinterRemoved)
				if err := _TokenMessenger.contract.UnpackLog(event, "LocalMinterRemoved", log); err != nil {
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

// ParseLocalMinterRemoved is a log parse operation binding the contract event 0x2db49fbf671271826a27b02ebc496209c85fffffb4bccc67430d2a0f22b4d1ac.
//
// Solidity: event LocalMinterRemoved(address localMinter)
func (_TokenMessenger *TokenMessengerFilterer) ParseLocalMinterRemoved(log types.Log) (*TokenMessengerLocalMinterRemoved, error) {
	event := new(TokenMessengerLocalMinterRemoved)
	if err := _TokenMessenger.contract.UnpackLog(event, "LocalMinterRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerMintAndWithdrawIterator is returned from FilterMintAndWithdraw and is used to iterate over the raw logs and unpacked data for MintAndWithdraw events raised by the TokenMessenger contract.
type TokenMessengerMintAndWithdrawIterator struct {
	Event *TokenMessengerMintAndWithdraw // Event containing the contract specifics and raw log

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
func (it *TokenMessengerMintAndWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerMintAndWithdraw)
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
		it.Event = new(TokenMessengerMintAndWithdraw)
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
func (it *TokenMessengerMintAndWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerMintAndWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerMintAndWithdraw represents a MintAndWithdraw event raised by the TokenMessenger contract.
type TokenMessengerMintAndWithdraw struct {
	MintRecipient common.Address
	Amount        *big.Int
	MintToken     common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMintAndWithdraw is a free log retrieval operation binding the contract event 0x1b2a7ff080b8cb6ff436ce0372e399692bbfb6d4ae5766fd8d58a7b8cc6142e6.
//
// Solidity: event MintAndWithdraw(address indexed mintRecipient, uint256 amount, address indexed mintToken)
func (_TokenMessenger *TokenMessengerFilterer) FilterMintAndWithdraw(opts *bind.FilterOpts, mintRecipient []common.Address, mintToken []common.Address) (*TokenMessengerMintAndWithdrawIterator, error) {

	var mintRecipientRule []interface{}
	for _, mintRecipientItem := range mintRecipient {
		mintRecipientRule = append(mintRecipientRule, mintRecipientItem)
	}

	var mintTokenRule []interface{}
	for _, mintTokenItem := range mintToken {
		mintTokenRule = append(mintTokenRule, mintTokenItem)
	}

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "MintAndWithdraw", mintRecipientRule, mintTokenRule)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerMintAndWithdrawIterator{contract: _TokenMessenger.contract, event: "MintAndWithdraw", logs: logs, sub: sub}, nil
}

// WatchMintAndWithdraw is a free log subscription operation binding the contract event 0x1b2a7ff080b8cb6ff436ce0372e399692bbfb6d4ae5766fd8d58a7b8cc6142e6.
//
// Solidity: event MintAndWithdraw(address indexed mintRecipient, uint256 amount, address indexed mintToken)
func (_TokenMessenger *TokenMessengerFilterer) WatchMintAndWithdraw(opts *bind.WatchOpts, sink chan<- *TokenMessengerMintAndWithdraw, mintRecipient []common.Address, mintToken []common.Address) (event.Subscription, error) {

	var mintRecipientRule []interface{}
	for _, mintRecipientItem := range mintRecipient {
		mintRecipientRule = append(mintRecipientRule, mintRecipientItem)
	}

	var mintTokenRule []interface{}
	for _, mintTokenItem := range mintToken {
		mintTokenRule = append(mintTokenRule, mintTokenItem)
	}

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "MintAndWithdraw", mintRecipientRule, mintTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerMintAndWithdraw)
				if err := _TokenMessenger.contract.UnpackLog(event, "MintAndWithdraw", log); err != nil {
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

// ParseMintAndWithdraw is a log parse operation binding the contract event 0x1b2a7ff080b8cb6ff436ce0372e399692bbfb6d4ae5766fd8d58a7b8cc6142e6.
//
// Solidity: event MintAndWithdraw(address indexed mintRecipient, uint256 amount, address indexed mintToken)
func (_TokenMessenger *TokenMessengerFilterer) ParseMintAndWithdraw(log types.Log) (*TokenMessengerMintAndWithdraw, error) {
	event := new(TokenMessengerMintAndWithdraw)
	if err := _TokenMessenger.contract.UnpackLog(event, "MintAndWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the TokenMessenger contract.
type TokenMessengerOwnershipTransferStartedIterator struct {
	Event *TokenMessengerOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *TokenMessengerOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerOwnershipTransferStarted)
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
		it.Event = new(TokenMessengerOwnershipTransferStarted)
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
func (it *TokenMessengerOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the TokenMessenger contract.
type TokenMessengerOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TokenMessenger *TokenMessengerFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenMessengerOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerOwnershipTransferStartedIterator{contract: _TokenMessenger.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_TokenMessenger *TokenMessengerFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *TokenMessengerOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerOwnershipTransferStarted)
				if err := _TokenMessenger.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_TokenMessenger *TokenMessengerFilterer) ParseOwnershipTransferStarted(log types.Log) (*TokenMessengerOwnershipTransferStarted, error) {
	event := new(TokenMessengerOwnershipTransferStarted)
	if err := _TokenMessenger.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenMessenger contract.
type TokenMessengerOwnershipTransferredIterator struct {
	Event *TokenMessengerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenMessengerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerOwnershipTransferred)
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
		it.Event = new(TokenMessengerOwnershipTransferred)
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
func (it *TokenMessengerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerOwnershipTransferred represents a OwnershipTransferred event raised by the TokenMessenger contract.
type TokenMessengerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenMessenger *TokenMessengerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenMessengerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerOwnershipTransferredIterator{contract: _TokenMessenger.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TokenMessenger *TokenMessengerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenMessengerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerOwnershipTransferred)
				if err := _TokenMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TokenMessenger *TokenMessengerFilterer) ParseOwnershipTransferred(log types.Log) (*TokenMessengerOwnershipTransferred, error) {
	event := new(TokenMessengerOwnershipTransferred)
	if err := _TokenMessenger.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerRemoteTokenMessengerAddedIterator is returned from FilterRemoteTokenMessengerAdded and is used to iterate over the raw logs and unpacked data for RemoteTokenMessengerAdded events raised by the TokenMessenger contract.
type TokenMessengerRemoteTokenMessengerAddedIterator struct {
	Event *TokenMessengerRemoteTokenMessengerAdded // Event containing the contract specifics and raw log

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
func (it *TokenMessengerRemoteTokenMessengerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerRemoteTokenMessengerAdded)
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
		it.Event = new(TokenMessengerRemoteTokenMessengerAdded)
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
func (it *TokenMessengerRemoteTokenMessengerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerRemoteTokenMessengerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerRemoteTokenMessengerAdded represents a RemoteTokenMessengerAdded event raised by the TokenMessenger contract.
type TokenMessengerRemoteTokenMessengerAdded struct {
	Domain         uint32
	TokenMessenger [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemoteTokenMessengerAdded is a free log retrieval operation binding the contract event 0x4bba2b08298cf59661b4895e384cc2ac3962ce2d71f1b7c11bca52e1169f9599.
//
// Solidity: event RemoteTokenMessengerAdded(uint32 domain, bytes32 tokenMessenger)
func (_TokenMessenger *TokenMessengerFilterer) FilterRemoteTokenMessengerAdded(opts *bind.FilterOpts) (*TokenMessengerRemoteTokenMessengerAddedIterator, error) {

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "RemoteTokenMessengerAdded")
	if err != nil {
		return nil, err
	}
	return &TokenMessengerRemoteTokenMessengerAddedIterator{contract: _TokenMessenger.contract, event: "RemoteTokenMessengerAdded", logs: logs, sub: sub}, nil
}

// WatchRemoteTokenMessengerAdded is a free log subscription operation binding the contract event 0x4bba2b08298cf59661b4895e384cc2ac3962ce2d71f1b7c11bca52e1169f9599.
//
// Solidity: event RemoteTokenMessengerAdded(uint32 domain, bytes32 tokenMessenger)
func (_TokenMessenger *TokenMessengerFilterer) WatchRemoteTokenMessengerAdded(opts *bind.WatchOpts, sink chan<- *TokenMessengerRemoteTokenMessengerAdded) (event.Subscription, error) {

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "RemoteTokenMessengerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerRemoteTokenMessengerAdded)
				if err := _TokenMessenger.contract.UnpackLog(event, "RemoteTokenMessengerAdded", log); err != nil {
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

// ParseRemoteTokenMessengerAdded is a log parse operation binding the contract event 0x4bba2b08298cf59661b4895e384cc2ac3962ce2d71f1b7c11bca52e1169f9599.
//
// Solidity: event RemoteTokenMessengerAdded(uint32 domain, bytes32 tokenMessenger)
func (_TokenMessenger *TokenMessengerFilterer) ParseRemoteTokenMessengerAdded(log types.Log) (*TokenMessengerRemoteTokenMessengerAdded, error) {
	event := new(TokenMessengerRemoteTokenMessengerAdded)
	if err := _TokenMessenger.contract.UnpackLog(event, "RemoteTokenMessengerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerRemoteTokenMessengerRemovedIterator is returned from FilterRemoteTokenMessengerRemoved and is used to iterate over the raw logs and unpacked data for RemoteTokenMessengerRemoved events raised by the TokenMessenger contract.
type TokenMessengerRemoteTokenMessengerRemovedIterator struct {
	Event *TokenMessengerRemoteTokenMessengerRemoved // Event containing the contract specifics and raw log

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
func (it *TokenMessengerRemoteTokenMessengerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerRemoteTokenMessengerRemoved)
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
		it.Event = new(TokenMessengerRemoteTokenMessengerRemoved)
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
func (it *TokenMessengerRemoteTokenMessengerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerRemoteTokenMessengerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerRemoteTokenMessengerRemoved represents a RemoteTokenMessengerRemoved event raised by the TokenMessenger contract.
type TokenMessengerRemoteTokenMessengerRemoved struct {
	Domain         uint32
	TokenMessenger [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRemoteTokenMessengerRemoved is a free log retrieval operation binding the contract event 0x3dcea012093dbca2bb8ed7fd2b2ff90305ab70bddda8bbb94d4152735a98f0b1.
//
// Solidity: event RemoteTokenMessengerRemoved(uint32 domain, bytes32 tokenMessenger)
func (_TokenMessenger *TokenMessengerFilterer) FilterRemoteTokenMessengerRemoved(opts *bind.FilterOpts) (*TokenMessengerRemoteTokenMessengerRemovedIterator, error) {

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "RemoteTokenMessengerRemoved")
	if err != nil {
		return nil, err
	}
	return &TokenMessengerRemoteTokenMessengerRemovedIterator{contract: _TokenMessenger.contract, event: "RemoteTokenMessengerRemoved", logs: logs, sub: sub}, nil
}

// WatchRemoteTokenMessengerRemoved is a free log subscription operation binding the contract event 0x3dcea012093dbca2bb8ed7fd2b2ff90305ab70bddda8bbb94d4152735a98f0b1.
//
// Solidity: event RemoteTokenMessengerRemoved(uint32 domain, bytes32 tokenMessenger)
func (_TokenMessenger *TokenMessengerFilterer) WatchRemoteTokenMessengerRemoved(opts *bind.WatchOpts, sink chan<- *TokenMessengerRemoteTokenMessengerRemoved) (event.Subscription, error) {

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "RemoteTokenMessengerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerRemoteTokenMessengerRemoved)
				if err := _TokenMessenger.contract.UnpackLog(event, "RemoteTokenMessengerRemoved", log); err != nil {
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

// ParseRemoteTokenMessengerRemoved is a log parse operation binding the contract event 0x3dcea012093dbca2bb8ed7fd2b2ff90305ab70bddda8bbb94d4152735a98f0b1.
//
// Solidity: event RemoteTokenMessengerRemoved(uint32 domain, bytes32 tokenMessenger)
func (_TokenMessenger *TokenMessengerFilterer) ParseRemoteTokenMessengerRemoved(log types.Log) (*TokenMessengerRemoteTokenMessengerRemoved, error) {
	event := new(TokenMessengerRemoteTokenMessengerRemoved)
	if err := _TokenMessenger.contract.UnpackLog(event, "RemoteTokenMessengerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenMessengerRescuerChangedIterator is returned from FilterRescuerChanged and is used to iterate over the raw logs and unpacked data for RescuerChanged events raised by the TokenMessenger contract.
type TokenMessengerRescuerChangedIterator struct {
	Event *TokenMessengerRescuerChanged // Event containing the contract specifics and raw log

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
func (it *TokenMessengerRescuerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMessengerRescuerChanged)
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
		it.Event = new(TokenMessengerRescuerChanged)
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
func (it *TokenMessengerRescuerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMessengerRescuerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMessengerRescuerChanged represents a RescuerChanged event raised by the TokenMessenger contract.
type TokenMessengerRescuerChanged struct {
	NewRescuer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRescuerChanged is a free log retrieval operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_TokenMessenger *TokenMessengerFilterer) FilterRescuerChanged(opts *bind.FilterOpts, newRescuer []common.Address) (*TokenMessengerRescuerChangedIterator, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _TokenMessenger.contract.FilterLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return &TokenMessengerRescuerChangedIterator{contract: _TokenMessenger.contract, event: "RescuerChanged", logs: logs, sub: sub}, nil
}

// WatchRescuerChanged is a free log subscription operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_TokenMessenger *TokenMessengerFilterer) WatchRescuerChanged(opts *bind.WatchOpts, sink chan<- *TokenMessengerRescuerChanged, newRescuer []common.Address) (event.Subscription, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _TokenMessenger.contract.WatchLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMessengerRescuerChanged)
				if err := _TokenMessenger.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
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

// ParseRescuerChanged is a log parse operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_TokenMessenger *TokenMessengerFilterer) ParseRescuerChanged(log types.Log) (*TokenMessengerRescuerChanged, error) {
	event := new(TokenMessengerRescuerChanged)
	if err := _TokenMessenger.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TypedMemViewMetaData contains all meta data concerning the TypedMemView contract.
var TypedMemViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f26be3fc": "NULL()",
	},
	Bin: "0x60cd610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b603e6073565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000009092168252519081900360200190f35b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008156fea2646970667358221220a21227c20433f6a4b6380811afd62868553688f85286db2e3a286c277076b7ff64736f6c63430007060033",
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
	parsed, err := TypedMemViewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
