// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package recipientmock

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

// IZapRecipientMetaData contains all meta data concerning the IZapRecipient contract.
var IZapRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"name\":\"zap\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e85e13dd": "zap(address,uint256,bytes)",
	},
}

// IZapRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use IZapRecipientMetaData.ABI instead.
var IZapRecipientABI = IZapRecipientMetaData.ABI

// Deprecated: Use IZapRecipientMetaData.Sigs instead.
// IZapRecipientFuncSigs maps the 4-byte function signature to its string representation.
var IZapRecipientFuncSigs = IZapRecipientMetaData.Sigs

// IZapRecipient is an auto generated Go binding around an Ethereum contract.
type IZapRecipient struct {
	IZapRecipientCaller     // Read-only binding to the contract
	IZapRecipientTransactor // Write-only binding to the contract
	IZapRecipientFilterer   // Log filterer for contract events
}

// IZapRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IZapRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZapRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IZapRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZapRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IZapRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IZapRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IZapRecipientSession struct {
	Contract     *IZapRecipient    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IZapRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IZapRecipientCallerSession struct {
	Contract *IZapRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IZapRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IZapRecipientTransactorSession struct {
	Contract     *IZapRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IZapRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IZapRecipientRaw struct {
	Contract *IZapRecipient // Generic contract binding to access the raw methods on
}

// IZapRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IZapRecipientCallerRaw struct {
	Contract *IZapRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IZapRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IZapRecipientTransactorRaw struct {
	Contract *IZapRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIZapRecipient creates a new instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipient(address common.Address, backend bind.ContractBackend) (*IZapRecipient, error) {
	contract, err := bindIZapRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IZapRecipient{IZapRecipientCaller: IZapRecipientCaller{contract: contract}, IZapRecipientTransactor: IZapRecipientTransactor{contract: contract}, IZapRecipientFilterer: IZapRecipientFilterer{contract: contract}}, nil
}

// NewIZapRecipientCaller creates a new read-only instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipientCaller(address common.Address, caller bind.ContractCaller) (*IZapRecipientCaller, error) {
	contract, err := bindIZapRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IZapRecipientCaller{contract: contract}, nil
}

// NewIZapRecipientTransactor creates a new write-only instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IZapRecipientTransactor, error) {
	contract, err := bindIZapRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IZapRecipientTransactor{contract: contract}, nil
}

// NewIZapRecipientFilterer creates a new log filterer instance of IZapRecipient, bound to a specific deployed contract.
func NewIZapRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IZapRecipientFilterer, error) {
	contract, err := bindIZapRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IZapRecipientFilterer{contract: contract}, nil
}

// bindIZapRecipient binds a generic wrapper to an already deployed contract.
func bindIZapRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IZapRecipientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZapRecipient *IZapRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZapRecipient.Contract.IZapRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZapRecipient *IZapRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZapRecipient.Contract.IZapRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZapRecipient *IZapRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZapRecipient.Contract.IZapRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IZapRecipient *IZapRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IZapRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IZapRecipient *IZapRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IZapRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IZapRecipient *IZapRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IZapRecipient.Contract.contract.Transact(opts, method, params...)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address token, uint256 amount, bytes zapData) payable returns(bytes4)
func (_IZapRecipient *IZapRecipientTransactor) Zap(opts *bind.TransactOpts, token common.Address, amount *big.Int, zapData []byte) (*types.Transaction, error) {
	return _IZapRecipient.contract.Transact(opts, "zap", token, amount, zapData)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address token, uint256 amount, bytes zapData) payable returns(bytes4)
func (_IZapRecipient *IZapRecipientSession) Zap(token common.Address, amount *big.Int, zapData []byte) (*types.Transaction, error) {
	return _IZapRecipient.Contract.Zap(&_IZapRecipient.TransactOpts, token, amount, zapData)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address token, uint256 amount, bytes zapData) payable returns(bytes4)
func (_IZapRecipient *IZapRecipientTransactorSession) Zap(token common.Address, amount *big.Int, zapData []byte) (*types.Transaction, error) {
	return _IZapRecipient.Contract.Zap(&_IZapRecipient.TransactOpts, token, amount, zapData)
}

// RecipientMockMetaData contains all meta data concerning the RecipientMock contract.
var RecipientMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"testRecipientMock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"zap\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"b239c091": "testRecipientMock()",
		"e85e13dd": "zap(address,uint256,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061021b806100206000396000f3fe60806040526004361061002d5760003560e01c8063b239c09114610039578063e85e13dd1461004757600080fd5b3661003457005b600080fd5b34801561004557600080fd5b005b61007d6100553660046100e1565b7fe85e13dd000000000000000000000000000000000000000000000000000000009392505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200160405180910390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000606084860312156100f657600080fd5b833573ffffffffffffffffffffffffffffffffffffffff8116811461011a57600080fd5b925060208401359150604084013567ffffffffffffffff8082111561013e57600080fd5b818601915086601f83011261015257600080fd5b813581811115610164576101646100b2565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156101aa576101aa6100b2565b816040528281528960208487010111156101c357600080fd5b826020860160208301376000602084830101528095505050505050925092509256fea26469706673582212206a5bc8b6ce97f6638e1238777c2558656268131efa652b2ea633dc4c60fe400a64736f6c63430008140033",
}

// RecipientMockABI is the input ABI used to generate the binding from.
// Deprecated: Use RecipientMockMetaData.ABI instead.
var RecipientMockABI = RecipientMockMetaData.ABI

// Deprecated: Use RecipientMockMetaData.Sigs instead.
// RecipientMockFuncSigs maps the 4-byte function signature to its string representation.
var RecipientMockFuncSigs = RecipientMockMetaData.Sigs

// RecipientMockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RecipientMockMetaData.Bin instead.
var RecipientMockBin = RecipientMockMetaData.Bin

// DeployRecipientMock deploys a new Ethereum contract, binding an instance of RecipientMock to it.
func DeployRecipientMock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RecipientMock, error) {
	parsed, err := RecipientMockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RecipientMockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RecipientMock{RecipientMockCaller: RecipientMockCaller{contract: contract}, RecipientMockTransactor: RecipientMockTransactor{contract: contract}, RecipientMockFilterer: RecipientMockFilterer{contract: contract}}, nil
}

// RecipientMock is an auto generated Go binding around an Ethereum contract.
type RecipientMock struct {
	RecipientMockCaller     // Read-only binding to the contract
	RecipientMockTransactor // Write-only binding to the contract
	RecipientMockFilterer   // Log filterer for contract events
}

// RecipientMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecipientMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecipientMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecipientMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecipientMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecipientMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecipientMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecipientMockSession struct {
	Contract     *RecipientMock    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecipientMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecipientMockCallerSession struct {
	Contract *RecipientMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RecipientMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecipientMockTransactorSession struct {
	Contract     *RecipientMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RecipientMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecipientMockRaw struct {
	Contract *RecipientMock // Generic contract binding to access the raw methods on
}

// RecipientMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecipientMockCallerRaw struct {
	Contract *RecipientMockCaller // Generic read-only contract binding to access the raw methods on
}

// RecipientMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecipientMockTransactorRaw struct {
	Contract *RecipientMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecipientMock creates a new instance of RecipientMock, bound to a specific deployed contract.
func NewRecipientMock(address common.Address, backend bind.ContractBackend) (*RecipientMock, error) {
	contract, err := bindRecipientMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RecipientMock{RecipientMockCaller: RecipientMockCaller{contract: contract}, RecipientMockTransactor: RecipientMockTransactor{contract: contract}, RecipientMockFilterer: RecipientMockFilterer{contract: contract}}, nil
}

// NewRecipientMockCaller creates a new read-only instance of RecipientMock, bound to a specific deployed contract.
func NewRecipientMockCaller(address common.Address, caller bind.ContractCaller) (*RecipientMockCaller, error) {
	contract, err := bindRecipientMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecipientMockCaller{contract: contract}, nil
}

// NewRecipientMockTransactor creates a new write-only instance of RecipientMock, bound to a specific deployed contract.
func NewRecipientMockTransactor(address common.Address, transactor bind.ContractTransactor) (*RecipientMockTransactor, error) {
	contract, err := bindRecipientMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecipientMockTransactor{contract: contract}, nil
}

// NewRecipientMockFilterer creates a new log filterer instance of RecipientMock, bound to a specific deployed contract.
func NewRecipientMockFilterer(address common.Address, filterer bind.ContractFilterer) (*RecipientMockFilterer, error) {
	contract, err := bindRecipientMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecipientMockFilterer{contract: contract}, nil
}

// bindRecipientMock binds a generic wrapper to an already deployed contract.
func bindRecipientMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecipientMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecipientMock *RecipientMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecipientMock.Contract.RecipientMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecipientMock *RecipientMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecipientMock.Contract.RecipientMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecipientMock *RecipientMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecipientMock.Contract.RecipientMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RecipientMock *RecipientMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RecipientMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RecipientMock *RecipientMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecipientMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RecipientMock *RecipientMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RecipientMock.Contract.contract.Transact(opts, method, params...)
}

// TestRecipientMock is a paid mutator transaction binding the contract method 0xb239c091.
//
// Solidity: function testRecipientMock() returns()
func (_RecipientMock *RecipientMockTransactor) TestRecipientMock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecipientMock.contract.Transact(opts, "testRecipientMock")
}

// TestRecipientMock is a paid mutator transaction binding the contract method 0xb239c091.
//
// Solidity: function testRecipientMock() returns()
func (_RecipientMock *RecipientMockSession) TestRecipientMock() (*types.Transaction, error) {
	return _RecipientMock.Contract.TestRecipientMock(&_RecipientMock.TransactOpts)
}

// TestRecipientMock is a paid mutator transaction binding the contract method 0xb239c091.
//
// Solidity: function testRecipientMock() returns()
func (_RecipientMock *RecipientMockTransactorSession) TestRecipientMock() (*types.Transaction, error) {
	return _RecipientMock.Contract.TestRecipientMock(&_RecipientMock.TransactOpts)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address , uint256 , bytes ) payable returns(bytes4)
func (_RecipientMock *RecipientMockTransactor) Zap(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RecipientMock.contract.Transact(opts, "zap", arg0, arg1, arg2)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address , uint256 , bytes ) payable returns(bytes4)
func (_RecipientMock *RecipientMockSession) Zap(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RecipientMock.Contract.Zap(&_RecipientMock.TransactOpts, arg0, arg1, arg2)
}

// Zap is a paid mutator transaction binding the contract method 0xe85e13dd.
//
// Solidity: function zap(address , uint256 , bytes ) payable returns(bytes4)
func (_RecipientMock *RecipientMockTransactorSession) Zap(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RecipientMock.Contract.Zap(&_RecipientMock.TransactOpts, arg0, arg1, arg2)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RecipientMock *RecipientMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RecipientMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RecipientMock *RecipientMockSession) Receive() (*types.Transaction, error) {
	return _RecipientMock.Contract.Receive(&_RecipientMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_RecipientMock *RecipientMockTransactorSession) Receive() (*types.Transaction, error) {
	return _RecipientMock.Contract.Receive(&_RecipientMock.TransactOpts)
}
