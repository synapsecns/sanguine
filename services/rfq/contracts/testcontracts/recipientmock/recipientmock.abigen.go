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

// IFastBridgeRecipientMetaData contains all meta data concerning the IFastBridgeRecipient contract.
var IFastBridgeRecipientMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"callParams\",\"type\":\"bytes\"}],\"name\":\"fastBridgeTransferReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"461e0c21": "fastBridgeTransferReceived(address,uint256,bytes)",
	},
}

// IFastBridgeRecipientABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastBridgeRecipientMetaData.ABI instead.
var IFastBridgeRecipientABI = IFastBridgeRecipientMetaData.ABI

// Deprecated: Use IFastBridgeRecipientMetaData.Sigs instead.
// IFastBridgeRecipientFuncSigs maps the 4-byte function signature to its string representation.
var IFastBridgeRecipientFuncSigs = IFastBridgeRecipientMetaData.Sigs

// IFastBridgeRecipient is an auto generated Go binding around an Ethereum contract.
type IFastBridgeRecipient struct {
	IFastBridgeRecipientCaller     // Read-only binding to the contract
	IFastBridgeRecipientTransactor // Write-only binding to the contract
	IFastBridgeRecipientFilterer   // Log filterer for contract events
}

// IFastBridgeRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFastBridgeRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastBridgeRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastBridgeRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastBridgeRecipientSession struct {
	Contract     *IFastBridgeRecipient // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IFastBridgeRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastBridgeRecipientCallerSession struct {
	Contract *IFastBridgeRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// IFastBridgeRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastBridgeRecipientTransactorSession struct {
	Contract     *IFastBridgeRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// IFastBridgeRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFastBridgeRecipientRaw struct {
	Contract *IFastBridgeRecipient // Generic contract binding to access the raw methods on
}

// IFastBridgeRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastBridgeRecipientCallerRaw struct {
	Contract *IFastBridgeRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// IFastBridgeRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastBridgeRecipientTransactorRaw struct {
	Contract *IFastBridgeRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastBridgeRecipient creates a new instance of IFastBridgeRecipient, bound to a specific deployed contract.
func NewIFastBridgeRecipient(address common.Address, backend bind.ContractBackend) (*IFastBridgeRecipient, error) {
	contract, err := bindIFastBridgeRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeRecipient{IFastBridgeRecipientCaller: IFastBridgeRecipientCaller{contract: contract}, IFastBridgeRecipientTransactor: IFastBridgeRecipientTransactor{contract: contract}, IFastBridgeRecipientFilterer: IFastBridgeRecipientFilterer{contract: contract}}, nil
}

// NewIFastBridgeRecipientCaller creates a new read-only instance of IFastBridgeRecipient, bound to a specific deployed contract.
func NewIFastBridgeRecipientCaller(address common.Address, caller bind.ContractCaller) (*IFastBridgeRecipientCaller, error) {
	contract, err := bindIFastBridgeRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeRecipientCaller{contract: contract}, nil
}

// NewIFastBridgeRecipientTransactor creates a new write-only instance of IFastBridgeRecipient, bound to a specific deployed contract.
func NewIFastBridgeRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*IFastBridgeRecipientTransactor, error) {
	contract, err := bindIFastBridgeRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeRecipientTransactor{contract: contract}, nil
}

// NewIFastBridgeRecipientFilterer creates a new log filterer instance of IFastBridgeRecipient, bound to a specific deployed contract.
func NewIFastBridgeRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*IFastBridgeRecipientFilterer, error) {
	contract, err := bindIFastBridgeRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeRecipientFilterer{contract: contract}, nil
}

// bindIFastBridgeRecipient binds a generic wrapper to an already deployed contract.
func bindIFastBridgeRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFastBridgeRecipientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeRecipient *IFastBridgeRecipientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeRecipient.Contract.IFastBridgeRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeRecipient *IFastBridgeRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeRecipient.Contract.IFastBridgeRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeRecipient *IFastBridgeRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeRecipient.Contract.IFastBridgeRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeRecipient *IFastBridgeRecipientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeRecipient *IFastBridgeRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeRecipient *IFastBridgeRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeRecipient.Contract.contract.Transact(opts, method, params...)
}

// FastBridgeTransferReceived is a paid mutator transaction binding the contract method 0x461e0c21.
//
// Solidity: function fastBridgeTransferReceived(address token, uint256 amount, bytes callParams) payable returns(bytes4)
func (_IFastBridgeRecipient *IFastBridgeRecipientTransactor) FastBridgeTransferReceived(opts *bind.TransactOpts, token common.Address, amount *big.Int, callParams []byte) (*types.Transaction, error) {
	return _IFastBridgeRecipient.contract.Transact(opts, "fastBridgeTransferReceived", token, amount, callParams)
}

// FastBridgeTransferReceived is a paid mutator transaction binding the contract method 0x461e0c21.
//
// Solidity: function fastBridgeTransferReceived(address token, uint256 amount, bytes callParams) payable returns(bytes4)
func (_IFastBridgeRecipient *IFastBridgeRecipientSession) FastBridgeTransferReceived(token common.Address, amount *big.Int, callParams []byte) (*types.Transaction, error) {
	return _IFastBridgeRecipient.Contract.FastBridgeTransferReceived(&_IFastBridgeRecipient.TransactOpts, token, amount, callParams)
}

// FastBridgeTransferReceived is a paid mutator transaction binding the contract method 0x461e0c21.
//
// Solidity: function fastBridgeTransferReceived(address token, uint256 amount, bytes callParams) payable returns(bytes4)
func (_IFastBridgeRecipient *IFastBridgeRecipientTransactorSession) FastBridgeTransferReceived(token common.Address, amount *big.Int, callParams []byte) (*types.Transaction, error) {
	return _IFastBridgeRecipient.Contract.FastBridgeTransferReceived(&_IFastBridgeRecipient.TransactOpts, token, amount, callParams)
}

// RecipientMockMetaData contains all meta data concerning the RecipientMock contract.
var RecipientMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"fastBridgeTransferReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"461e0c21": "fastBridgeTransferReceived(address,uint256,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610202806100206000396000f3fe6080604052600436106100225760003560e01c8063461e0c211461002e57600080fd5b3661002957005b600080fd5b61006461003c3660046100c8565b7f461e0c21000000000000000000000000000000000000000000000000000000009392505050565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200160405180910390f35b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000606084860312156100dd57600080fd5b833573ffffffffffffffffffffffffffffffffffffffff8116811461010157600080fd5b925060208401359150604084013567ffffffffffffffff8082111561012557600080fd5b818601915086601f83011261013957600080fd5b81358181111561014b5761014b610099565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561019157610191610099565b816040528281528960208487010111156101aa57600080fd5b826020860160208301376000602084830101528095505050505050925092509256fea2646970667358221220f1a74d56eefc9976c67d52f3119adaca862e91740662087e97a8e47c248fbb1d64736f6c63430008140033",
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

// FastBridgeTransferReceived is a paid mutator transaction binding the contract method 0x461e0c21.
//
// Solidity: function fastBridgeTransferReceived(address , uint256 , bytes ) payable returns(bytes4)
func (_RecipientMock *RecipientMockTransactor) FastBridgeTransferReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RecipientMock.contract.Transact(opts, "fastBridgeTransferReceived", arg0, arg1, arg2)
}

// FastBridgeTransferReceived is a paid mutator transaction binding the contract method 0x461e0c21.
//
// Solidity: function fastBridgeTransferReceived(address , uint256 , bytes ) payable returns(bytes4)
func (_RecipientMock *RecipientMockSession) FastBridgeTransferReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RecipientMock.Contract.FastBridgeTransferReceived(&_RecipientMock.TransactOpts, arg0, arg1, arg2)
}

// FastBridgeTransferReceived is a paid mutator transaction binding the contract method 0x461e0c21.
//
// Solidity: function fastBridgeTransferReceived(address , uint256 , bytes ) payable returns(bytes4)
func (_RecipientMock *RecipientMockTransactorSession) FastBridgeTransferReceived(arg0 common.Address, arg1 *big.Int, arg2 []byte) (*types.Transaction, error) {
	return _RecipientMock.Contract.FastBridgeTransferReceived(&_RecipientMock.TransactOpts, arg0, arg1, arg2)
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
