// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridgetransactionv2

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

// IFastBridgeBridgeParams is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeBridgeParams struct {
	DstChainId   uint32
	Sender       common.Address
	To           common.Address
	OriginToken  common.Address
	DestToken    common.Address
	OriginAmount *big.Int
	DestAmount   *big.Int
	SendChainGas bool
	Deadline     *big.Int
}

// IFastBridgeBridgeTransaction is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeBridgeTransaction struct {
	OriginChainId   uint32
	DestChainId     uint32
	OriginSender    common.Address
	DestRecipient   common.Address
	OriginToken     common.Address
	DestToken       common.Address
	OriginAmount    *big.Int
	DestAmount      *big.Int
	OriginFeeAmount *big.Int
	SendChainGas    bool
	Deadline        *big.Int
	Nonce           *big.Int
}

// IFastBridgeV2BridgeParamsV2 is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeV2BridgeParamsV2 struct {
	QuoteRelayer            common.Address
	QuoteExclusivitySeconds *big.Int
	QuoteId                 []byte
	ZapNative               *big.Int
	ZapData                 []byte
}

// IFastBridgeV2BridgeTransactionV2 is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeV2BridgeTransactionV2 struct {
	OriginChainId      uint32
	DestChainId        uint32
	OriginSender       common.Address
	DestRecipient      common.Address
	OriginToken        common.Address
	DestToken          common.Address
	OriginAmount       *big.Int
	DestAmount         *big.Int
	OriginFeeAmount    *big.Int
	Deadline           *big.Int
	Nonce              *big.Int
	ExclusivityRelayer common.Address
	ExclusivityEndTime *big.Int
	ZapNative          *big.Int
	ZapData            []byte
}

// BridgeTransactionV2LibMetaData contains all meta data concerning the BridgeTransactionV2Lib contract.
var BridgeTransactionV2LibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BridgeTransactionV2__InvalidEncodedTx\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"version\",\"type\":\"uint16\"}],\"name\":\"BridgeTransactionV2__UnsupportedVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122015e035caa40406d51f5c7b0d5428e946b6e76cdb0ace342f3f3b4702af80f9a564736f6c63430008180033",
}

// BridgeTransactionV2LibABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeTransactionV2LibMetaData.ABI instead.
var BridgeTransactionV2LibABI = BridgeTransactionV2LibMetaData.ABI

// BridgeTransactionV2LibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeTransactionV2LibMetaData.Bin instead.
var BridgeTransactionV2LibBin = BridgeTransactionV2LibMetaData.Bin

// DeployBridgeTransactionV2Lib deploys a new Ethereum contract, binding an instance of BridgeTransactionV2Lib to it.
func DeployBridgeTransactionV2Lib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BridgeTransactionV2Lib, error) {
	parsed, err := BridgeTransactionV2LibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeTransactionV2LibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeTransactionV2Lib{BridgeTransactionV2LibCaller: BridgeTransactionV2LibCaller{contract: contract}, BridgeTransactionV2LibTransactor: BridgeTransactionV2LibTransactor{contract: contract}, BridgeTransactionV2LibFilterer: BridgeTransactionV2LibFilterer{contract: contract}}, nil
}

// BridgeTransactionV2Lib is an auto generated Go binding around an Ethereum contract.
type BridgeTransactionV2Lib struct {
	BridgeTransactionV2LibCaller     // Read-only binding to the contract
	BridgeTransactionV2LibTransactor // Write-only binding to the contract
	BridgeTransactionV2LibFilterer   // Log filterer for contract events
}

// BridgeTransactionV2LibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactionV2LibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactionV2LibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeTransactionV2LibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactionV2LibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTransactionV2LibSession struct {
	Contract     *BridgeTransactionV2Lib // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeTransactionV2LibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTransactionV2LibCallerSession struct {
	Contract *BridgeTransactionV2LibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// BridgeTransactionV2LibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactionV2LibTransactorSession struct {
	Contract     *BridgeTransactionV2LibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// BridgeTransactionV2LibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTransactionV2LibRaw struct {
	Contract *BridgeTransactionV2Lib // Generic contract binding to access the raw methods on
}

// BridgeTransactionV2LibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibCallerRaw struct {
	Contract *BridgeTransactionV2LibCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactionV2LibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactionV2LibTransactorRaw struct {
	Contract *BridgeTransactionV2LibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeTransactionV2Lib creates a new instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2Lib(address common.Address, backend bind.ContractBackend) (*BridgeTransactionV2Lib, error) {
	contract, err := bindBridgeTransactionV2Lib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2Lib{BridgeTransactionV2LibCaller: BridgeTransactionV2LibCaller{contract: contract}, BridgeTransactionV2LibTransactor: BridgeTransactionV2LibTransactor{contract: contract}, BridgeTransactionV2LibFilterer: BridgeTransactionV2LibFilterer{contract: contract}}, nil
}

// NewBridgeTransactionV2LibCaller creates a new read-only instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2LibCaller(address common.Address, caller bind.ContractCaller) (*BridgeTransactionV2LibCaller, error) {
	contract, err := bindBridgeTransactionV2Lib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2LibCaller{contract: contract}, nil
}

// NewBridgeTransactionV2LibTransactor creates a new write-only instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2LibTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactionV2LibTransactor, error) {
	contract, err := bindBridgeTransactionV2Lib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2LibTransactor{contract: contract}, nil
}

// NewBridgeTransactionV2LibFilterer creates a new log filterer instance of BridgeTransactionV2Lib, bound to a specific deployed contract.
func NewBridgeTransactionV2LibFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeTransactionV2LibFilterer, error) {
	contract, err := bindBridgeTransactionV2Lib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactionV2LibFilterer{contract: contract}, nil
}

// bindBridgeTransactionV2Lib binds a generic wrapper to an already deployed contract.
func bindBridgeTransactionV2Lib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeTransactionV2LibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTransactionV2Lib.Contract.BridgeTransactionV2LibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.BridgeTransactionV2LibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.BridgeTransactionV2LibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeTransactionV2Lib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeTransactionV2Lib *BridgeTransactionV2LibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeTransactionV2Lib.Contract.contract.Transact(opts, method, params...)
}

// IFastBridgeMetaData contains all meta data concerning the IFastBridge contract.
var IFastBridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"45851694": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256))",
		"aa9641ab": "canClaim(bytes32,address)",
		"41fcb612": "claim(bytes,address)",
		"add98c70": "dispute(bytes32)",
		"ac11fb1a": "getBridgeTransaction(bytes)",
		"886d36ff": "prove(bytes,bytes32)",
		"5eb7d946": "refund(bytes)",
		"8f0d6f17": "relay(bytes)",
	},
}

// IFastBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastBridgeMetaData.ABI instead.
var IFastBridgeABI = IFastBridgeMetaData.ABI

// Deprecated: Use IFastBridgeMetaData.Sigs instead.
// IFastBridgeFuncSigs maps the 4-byte function signature to its string representation.
var IFastBridgeFuncSigs = IFastBridgeMetaData.Sigs

// IFastBridge is an auto generated Go binding around an Ethereum contract.
type IFastBridge struct {
	IFastBridgeCaller     // Read-only binding to the contract
	IFastBridgeTransactor // Write-only binding to the contract
	IFastBridgeFilterer   // Log filterer for contract events
}

// IFastBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFastBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastBridgeSession struct {
	Contract     *IFastBridge      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFastBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastBridgeCallerSession struct {
	Contract *IFastBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IFastBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastBridgeTransactorSession struct {
	Contract     *IFastBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IFastBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFastBridgeRaw struct {
	Contract *IFastBridge // Generic contract binding to access the raw methods on
}

// IFastBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastBridgeCallerRaw struct {
	Contract *IFastBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// IFastBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastBridgeTransactorRaw struct {
	Contract *IFastBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastBridge creates a new instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridge(address common.Address, backend bind.ContractBackend) (*IFastBridge, error) {
	contract, err := bindIFastBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastBridge{IFastBridgeCaller: IFastBridgeCaller{contract: contract}, IFastBridgeTransactor: IFastBridgeTransactor{contract: contract}, IFastBridgeFilterer: IFastBridgeFilterer{contract: contract}}, nil
}

// NewIFastBridgeCaller creates a new read-only instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridgeCaller(address common.Address, caller bind.ContractCaller) (*IFastBridgeCaller, error) {
	contract, err := bindIFastBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeCaller{contract: contract}, nil
}

// NewIFastBridgeTransactor creates a new write-only instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*IFastBridgeTransactor, error) {
	contract, err := bindIFastBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeTransactor{contract: contract}, nil
}

// NewIFastBridgeFilterer creates a new log filterer instance of IFastBridge, bound to a specific deployed contract.
func NewIFastBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*IFastBridgeFilterer, error) {
	contract, err := bindIFastBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeFilterer{contract: contract}, nil
}

// bindIFastBridge binds a generic wrapper to an already deployed contract.
func bindIFastBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFastBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridge *IFastBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridge.Contract.IFastBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridge *IFastBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridge.Contract.IFastBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridge *IFastBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridge.Contract.IFastBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridge *IFastBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridge *IFastBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridge *IFastBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridge.Contract.contract.Transact(opts, method, params...)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridge *IFastBridgeCaller) CanClaim(opts *bind.CallOpts, transactionId [32]byte, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _IFastBridge.contract.Call(opts, &out, "canClaim", transactionId, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridge *IFastBridgeSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridge.Contract.CanClaim(&_IFastBridge.CallOpts, transactionId, relayer)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridge *IFastBridgeCallerSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridge.Contract.CanClaim(&_IFastBridge.CallOpts, transactionId, relayer)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeCaller) GetBridgeTransaction(opts *bind.CallOpts, request []byte) (IFastBridgeBridgeTransaction, error) {
	var out []interface{}
	err := _IFastBridge.contract.Call(opts, &out, "getBridgeTransaction", request)

	if err != nil {
		return *new(IFastBridgeBridgeTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeBridgeTransaction)).(*IFastBridgeBridgeTransaction)

	return out0, err

}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridge.Contract.GetBridgeTransaction(&_IFastBridge.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridge *IFastBridgeCallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridge.Contract.GetBridgeTransaction(&_IFastBridge.CallOpts, request)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeTransactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.Contract.Bridge(&_IFastBridge.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridge *IFastBridgeTransactorSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridge.Contract.Bridge(&_IFastBridge.TransactOpts, params)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeTransactor) Claim(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "claim", request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.Contract.Claim(&_IFastBridge.TransactOpts, request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridge.Contract.Claim(&_IFastBridge.TransactOpts, request, to)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridge *IFastBridgeTransactor) Dispute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "dispute", transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridge *IFastBridgeSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Dispute(&_IFastBridge.TransactOpts, transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Dispute(&_IFastBridge.TransactOpts, transactionId)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridge *IFastBridgeTransactor) Prove(opts *bind.TransactOpts, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "prove", request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridge *IFastBridgeSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Prove(&_IFastBridge.TransactOpts, request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Prove(&_IFastBridge.TransactOpts, request, destTxHash)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridge *IFastBridgeTransactor) Refund(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "refund", request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridge *IFastBridgeSession) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Refund(&_IFastBridge.TransactOpts, request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridge *IFastBridgeTransactorSession) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Refund(&_IFastBridge.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridge *IFastBridgeTransactor) Relay(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridge.contract.Transact(opts, "relay", request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridge *IFastBridgeSession) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Relay(&_IFastBridge.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridge *IFastBridgeTransactorSession) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridge.Contract.Relay(&_IFastBridge.TransactOpts, request)
}

// IFastBridgeBridgeDepositClaimedIterator is returned from FilterBridgeDepositClaimed and is used to iterate over the raw logs and unpacked data for BridgeDepositClaimed events raised by the IFastBridge contract.
type IFastBridgeBridgeDepositClaimedIterator struct {
	Event *IFastBridgeBridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeDepositClaimed)
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
		it.Event = new(IFastBridgeBridgeDepositClaimed)
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
func (it *IFastBridgeBridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeDepositClaimed represents a BridgeDepositClaimed event raised by the IFastBridge contract.
type IFastBridgeBridgeDepositClaimed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositClaimed is a free log retrieval operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeBridgeDepositClaimedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeDepositClaimedIterator{contract: _IFastBridge.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeDepositClaimed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeDepositClaimed)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
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

// ParseBridgeDepositClaimed is a log parse operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeDepositClaimed(log types.Log) (*IFastBridgeBridgeDepositClaimed, error) {
	event := new(IFastBridgeBridgeDepositClaimed)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeDepositRefundedIterator is returned from FilterBridgeDepositRefunded and is used to iterate over the raw logs and unpacked data for BridgeDepositRefunded events raised by the IFastBridge contract.
type IFastBridgeBridgeDepositRefundedIterator struct {
	Event *IFastBridgeBridgeDepositRefunded // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeDepositRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeDepositRefunded)
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
		it.Event = new(IFastBridgeBridgeDepositRefunded)
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
func (it *IFastBridgeBridgeDepositRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeDepositRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeDepositRefunded represents a BridgeDepositRefunded event raised by the IFastBridge contract.
type IFastBridgeBridgeDepositRefunded struct {
	TransactionId [32]byte
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositRefunded is a free log retrieval operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts, transactionId [][32]byte, to []common.Address) (*IFastBridgeBridgeDepositRefundedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeDepositRefundedIterator{contract: _IFastBridge.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeDepositRefunded, transactionId [][32]byte, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeDepositRefunded)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
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

// ParseBridgeDepositRefunded is a log parse operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeDepositRefunded(log types.Log) (*IFastBridgeBridgeDepositRefunded, error) {
	event := new(IFastBridgeBridgeDepositRefunded)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeProofDisputedIterator is returned from FilterBridgeProofDisputed and is used to iterate over the raw logs and unpacked data for BridgeProofDisputed events raised by the IFastBridge contract.
type IFastBridgeBridgeProofDisputedIterator struct {
	Event *IFastBridgeBridgeProofDisputed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeProofDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeProofDisputed)
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
		it.Event = new(IFastBridgeBridgeProofDisputed)
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
func (it *IFastBridgeBridgeProofDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeProofDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeProofDisputed represents a BridgeProofDisputed event raised by the IFastBridge contract.
type IFastBridgeBridgeProofDisputed struct {
	TransactionId [32]byte
	Relayer       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofDisputed is a free log retrieval operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeProofDisputed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeBridgeProofDisputedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeProofDisputedIterator{contract: _IFastBridge.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeProofDisputed, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeProofDisputed)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
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

// ParseBridgeProofDisputed is a log parse operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeProofDisputed(log types.Log) (*IFastBridgeBridgeProofDisputed, error) {
	event := new(IFastBridgeBridgeProofDisputed)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeProofProvidedIterator is returned from FilterBridgeProofProvided and is used to iterate over the raw logs and unpacked data for BridgeProofProvided events raised by the IFastBridge contract.
type IFastBridgeBridgeProofProvidedIterator struct {
	Event *IFastBridgeBridgeProofProvided // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeProofProvided)
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
		it.Event = new(IFastBridgeBridgeProofProvided)
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
func (it *IFastBridgeBridgeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeProofProvided represents a BridgeProofProvided event raised by the IFastBridge contract.
type IFastBridgeBridgeProofProvided struct {
	TransactionId   [32]byte
	Relayer         common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofProvided is a free log retrieval operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeProofProvided(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeBridgeProofProvidedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeProofProvidedIterator{contract: _IFastBridge.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeProofProvided, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeProofProvided)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
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

// ParseBridgeProofProvided is a log parse operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeProofProvided(log types.Log) (*IFastBridgeBridgeProofProvided, error) {
	event := new(IFastBridgeBridgeProofProvided)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeRelayedIterator is returned from FilterBridgeRelayed and is used to iterate over the raw logs and unpacked data for BridgeRelayed events raised by the IFastBridge contract.
type IFastBridgeBridgeRelayedIterator struct {
	Event *IFastBridgeBridgeRelayed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeRelayed)
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
		it.Event = new(IFastBridgeBridgeRelayed)
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
func (it *IFastBridgeBridgeRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeRelayed represents a BridgeRelayed event raised by the IFastBridge contract.
type IFastBridgeBridgeRelayed struct {
	TransactionId  [32]byte
	Relayer        common.Address
	To             common.Address
	OriginChainId  uint32
	OriginToken    common.Address
	DestToken      common.Address
	OriginAmount   *big.Int
	DestAmount     *big.Int
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeRelayed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeBridgeRelayedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeRelayedIterator{contract: _IFastBridge.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeRelayed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeRelayed)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeRelayed(log types.Log) (*IFastBridgeBridgeRelayed, error) {
	event := new(IFastBridgeBridgeRelayed)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeBridgeRequestedIterator is returned from FilterBridgeRequested and is used to iterate over the raw logs and unpacked data for BridgeRequested events raised by the IFastBridge contract.
type IFastBridgeBridgeRequestedIterator struct {
	Event *IFastBridgeBridgeRequested // Event containing the contract specifics and raw log

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
func (it *IFastBridgeBridgeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeBridgeRequested)
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
		it.Event = new(IFastBridgeBridgeRequested)
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
func (it *IFastBridgeBridgeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeBridgeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeBridgeRequested represents a BridgeRequested event raised by the IFastBridge contract.
type IFastBridgeBridgeRequested struct {
	TransactionId [32]byte
	Sender        common.Address
	Request       []byte
	DestChainId   uint32
	OriginToken   common.Address
	DestToken     common.Address
	OriginAmount  *big.Int
	DestAmount    *big.Int
	SendChainGas  bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridge *IFastBridgeFilterer) FilterBridgeRequested(opts *bind.FilterOpts, transactionId [][32]byte, sender []common.Address) (*IFastBridgeBridgeRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridge.contract.FilterLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeBridgeRequestedIterator{contract: _IFastBridge.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridge *IFastBridgeFilterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *IFastBridgeBridgeRequested, transactionId [][32]byte, sender []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridge.contract.WatchLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeBridgeRequested)
				if err := _IFastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridge *IFastBridgeFilterer) ParseBridgeRequested(log types.Log) (*IFastBridgeBridgeRequested, error) {
	event := new(IFastBridgeBridgeRequested)
	if err := _IFastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2MetaData contains all meta data concerning the IFastBridgeV2 contract.
var IFastBridgeV2MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"}],\"name\":\"BridgeQuoteDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainGasAmount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"quoteRelayer\",\"type\":\"address\"},{\"internalType\":\"int256\",\"name\":\"quoteExclusivitySeconds\",\"type\":\"int256\"},{\"internalType\":\"bytes\",\"name\":\"quoteId\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeParamsV2\",\"name\":\"paramsV2\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeProofs\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"timestamp\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"bridgeStatuses\",\"outputs\":[{\"internalType\":\"enumIFastBridgeV2.BridgeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"canClaim\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sendChainGas\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransactionV2\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"originFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"exclusivityRelayer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exclusivityEndTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"zapNative\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"zapData\",\"type\":\"bytes\"}],\"internalType\":\"structIFastBridgeV2.BridgeTransactionV2\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"45851694": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256))",
		"bfc7c607": "bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256),(address,int256,bytes,uint256,bytes))",
		"91ad5039": "bridgeProofs(bytes32)",
		"8379a24f": "bridgeRelays(bytes32)",
		"051287bc": "bridgeStatuses(bytes32)",
		"aa9641ab": "canClaim(bytes32,address)",
		"c63ff8dd": "claim(bytes)",
		"41fcb612": "claim(bytes,address)",
		"add98c70": "dispute(bytes32)",
		"ac11fb1a": "getBridgeTransaction(bytes)",
		"5aa6ccba": "getBridgeTransactionV2(bytes)",
		"886d36ff": "prove(bytes,bytes32)",
		"18e4357d": "prove(bytes32,bytes32,address)",
		"5eb7d946": "refund(bytes)",
		"8f0d6f17": "relay(bytes)",
		"9c9545f0": "relay(bytes,address)",
	},
}

// IFastBridgeV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use IFastBridgeV2MetaData.ABI instead.
var IFastBridgeV2ABI = IFastBridgeV2MetaData.ABI

// Deprecated: Use IFastBridgeV2MetaData.Sigs instead.
// IFastBridgeV2FuncSigs maps the 4-byte function signature to its string representation.
var IFastBridgeV2FuncSigs = IFastBridgeV2MetaData.Sigs

// IFastBridgeV2 is an auto generated Go binding around an Ethereum contract.
type IFastBridgeV2 struct {
	IFastBridgeV2Caller     // Read-only binding to the contract
	IFastBridgeV2Transactor // Write-only binding to the contract
	IFastBridgeV2Filterer   // Log filterer for contract events
}

// IFastBridgeV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type IFastBridgeV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IFastBridgeV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFastBridgeV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFastBridgeV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFastBridgeV2Session struct {
	Contract     *IFastBridgeV2    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFastBridgeV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFastBridgeV2CallerSession struct {
	Contract *IFastBridgeV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IFastBridgeV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFastBridgeV2TransactorSession struct {
	Contract     *IFastBridgeV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IFastBridgeV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type IFastBridgeV2Raw struct {
	Contract *IFastBridgeV2 // Generic contract binding to access the raw methods on
}

// IFastBridgeV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFastBridgeV2CallerRaw struct {
	Contract *IFastBridgeV2Caller // Generic read-only contract binding to access the raw methods on
}

// IFastBridgeV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFastBridgeV2TransactorRaw struct {
	Contract *IFastBridgeV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIFastBridgeV2 creates a new instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2(address common.Address, backend bind.ContractBackend) (*IFastBridgeV2, error) {
	contract, err := bindIFastBridgeV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2{IFastBridgeV2Caller: IFastBridgeV2Caller{contract: contract}, IFastBridgeV2Transactor: IFastBridgeV2Transactor{contract: contract}, IFastBridgeV2Filterer: IFastBridgeV2Filterer{contract: contract}}, nil
}

// NewIFastBridgeV2Caller creates a new read-only instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2Caller(address common.Address, caller bind.ContractCaller) (*IFastBridgeV2Caller, error) {
	contract, err := bindIFastBridgeV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2Caller{contract: contract}, nil
}

// NewIFastBridgeV2Transactor creates a new write-only instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2Transactor(address common.Address, transactor bind.ContractTransactor) (*IFastBridgeV2Transactor, error) {
	contract, err := bindIFastBridgeV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2Transactor{contract: contract}, nil
}

// NewIFastBridgeV2Filterer creates a new log filterer instance of IFastBridgeV2, bound to a specific deployed contract.
func NewIFastBridgeV2Filterer(address common.Address, filterer bind.ContractFilterer) (*IFastBridgeV2Filterer, error) {
	contract, err := bindIFastBridgeV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2Filterer{contract: contract}, nil
}

// bindIFastBridgeV2 binds a generic wrapper to an already deployed contract.
func bindIFastBridgeV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFastBridgeV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeV2 *IFastBridgeV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeV2.Contract.IFastBridgeV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeV2 *IFastBridgeV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.IFastBridgeV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeV2 *IFastBridgeV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.IFastBridgeV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFastBridgeV2 *IFastBridgeV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFastBridgeV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFastBridgeV2 *IFastBridgeV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFastBridgeV2 *IFastBridgeV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.contract.Transact(opts, method, params...)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_IFastBridgeV2 *IFastBridgeV2Caller) BridgeProofs(opts *bind.CallOpts, transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "bridgeProofs", transactionId)

	outstruct := new(struct {
		Timestamp *big.Int
		Relayer   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Relayer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_IFastBridgeV2 *IFastBridgeV2Session) BridgeProofs(transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _IFastBridgeV2.Contract.BridgeProofs(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 transactionId) view returns(uint96 timestamp, address relayer)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) BridgeProofs(transactionId [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _IFastBridgeV2.Contract.BridgeProofs(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Caller) BridgeRelays(opts *bind.CallOpts, transactionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "bridgeRelays", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Session) BridgeRelays(transactionId [32]byte) (bool, error) {
	return _IFastBridgeV2.Contract.BridgeRelays(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 transactionId) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) BridgeRelays(transactionId [32]byte) (bool, error) {
	return _IFastBridgeV2.Contract.BridgeRelays(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8)
func (_IFastBridgeV2 *IFastBridgeV2Caller) BridgeStatuses(opts *bind.CallOpts, transactionId [32]byte) (uint8, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "bridgeStatuses", transactionId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8)
func (_IFastBridgeV2 *IFastBridgeV2Session) BridgeStatuses(transactionId [32]byte) (uint8, error) {
	return _IFastBridgeV2.Contract.BridgeStatuses(&_IFastBridgeV2.CallOpts, transactionId)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 transactionId) view returns(uint8)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) BridgeStatuses(transactionId [32]byte) (uint8, error) {
	return _IFastBridgeV2.Contract.BridgeStatuses(&_IFastBridgeV2.CallOpts, transactionId)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Caller) CanClaim(opts *bind.CallOpts, transactionId [32]byte, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "canClaim", transactionId, relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2Session) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridgeV2.Contract.CanClaim(&_IFastBridgeV2.CallOpts, transactionId, relayer)
}

// CanClaim is a free data retrieval call binding the contract method 0xaa9641ab.
//
// Solidity: function canClaim(bytes32 transactionId, address relayer) view returns(bool)
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) CanClaim(transactionId [32]byte, relayer common.Address) (bool, error) {
	return _IFastBridgeV2.Contract.CanClaim(&_IFastBridgeV2.CallOpts, transactionId, relayer)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridgeV2 *IFastBridgeV2Caller) GetBridgeTransaction(opts *bind.CallOpts, request []byte) (IFastBridgeBridgeTransaction, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "getBridgeTransaction", request)

	if err != nil {
		return *new(IFastBridgeBridgeTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeBridgeTransaction)).(*IFastBridgeBridgeTransaction)

	return out0, err

}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridgeV2 *IFastBridgeV2Session) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransaction(&_IFastBridgeV2.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,bool,uint256,uint256))
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransaction(&_IFastBridgeV2.CallOpts, request)
}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_IFastBridgeV2 *IFastBridgeV2Caller) GetBridgeTransactionV2(opts *bind.CallOpts, request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	var out []interface{}
	err := _IFastBridgeV2.contract.Call(opts, &out, "getBridgeTransactionV2", request)

	if err != nil {
		return *new(IFastBridgeV2BridgeTransactionV2), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeV2BridgeTransactionV2)).(*IFastBridgeV2BridgeTransactionV2)

	return out0, err

}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_IFastBridgeV2 *IFastBridgeV2Session) GetBridgeTransactionV2(request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransactionV2(&_IFastBridgeV2.CallOpts, request)
}

// GetBridgeTransactionV2 is a free data retrieval call binding the contract method 0x5aa6ccba.
//
// Solidity: function getBridgeTransactionV2(bytes request) view returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256,uint256,address,uint256,uint256,bytes))
func (_IFastBridgeV2 *IFastBridgeV2CallerSession) GetBridgeTransactionV2(request []byte) (IFastBridgeV2BridgeTransactionV2, error) {
	return _IFastBridgeV2.Contract.GetBridgeTransactionV2(&_IFastBridgeV2.CallOpts, request)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge(&_IFastBridgeV2.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x45851694.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge(&_IFastBridgeV2.TransactOpts, params)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Bridge0(opts *bind.TransactOpts, params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "bridge0", params, paramsV2)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Bridge0(params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge0(&_IFastBridgeV2.TransactOpts, params, paramsV2)
}

// Bridge0 is a paid mutator transaction binding the contract method 0xbfc7c607.
//
// Solidity: function bridge((uint32,address,address,address,address,uint256,uint256,bool,uint256) params, (address,int256,bytes,uint256,bytes) paramsV2) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Bridge0(params IFastBridgeBridgeParams, paramsV2 IFastBridgeV2BridgeParamsV2) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Bridge0(&_IFastBridgeV2.TransactOpts, params, paramsV2)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Claim(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "claim", request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim(&_IFastBridgeV2.TransactOpts, request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim(&_IFastBridgeV2.TransactOpts, request, to)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Claim0(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "claim0", request)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Claim0(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim0(&_IFastBridgeV2.TransactOpts, request)
}

// Claim0 is a paid mutator transaction binding the contract method 0xc63ff8dd.
//
// Solidity: function claim(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Claim0(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Claim0(&_IFastBridgeV2.TransactOpts, request)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Dispute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "dispute", transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Dispute(&_IFastBridgeV2.TransactOpts, transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Dispute(&_IFastBridgeV2.TransactOpts, transactionId)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Prove(opts *bind.TransactOpts, transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "prove", transactionId, destTxHash, relayer)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Prove(transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove(&_IFastBridgeV2.TransactOpts, transactionId, destTxHash, relayer)
}

// Prove is a paid mutator transaction binding the contract method 0x18e4357d.
//
// Solidity: function prove(bytes32 transactionId, bytes32 destTxHash, address relayer) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Prove(transactionId [32]byte, destTxHash [32]byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove(&_IFastBridgeV2.TransactOpts, transactionId, destTxHash, relayer)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Prove0(opts *bind.TransactOpts, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "prove0", request, destTxHash)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Prove0(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove0(&_IFastBridgeV2.TransactOpts, request, destTxHash)
}

// Prove0 is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Prove0(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Prove0(&_IFastBridgeV2.TransactOpts, request, destTxHash)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Refund(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "refund", request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Refund(&_IFastBridgeV2.TransactOpts, request)
}

// Refund is a paid mutator transaction binding the contract method 0x5eb7d946.
//
// Solidity: function refund(bytes request) returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Refund(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Refund(&_IFastBridgeV2.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Relay(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "relay", request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay(&_IFastBridgeV2.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Relay(request []byte) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay(&_IFastBridgeV2.TransactOpts, request)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Transactor) Relay0(opts *bind.TransactOpts, request []byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.contract.Transact(opts, "relay0", request, relayer)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2Session) Relay0(request []byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay0(&_IFastBridgeV2.TransactOpts, request, relayer)
}

// Relay0 is a paid mutator transaction binding the contract method 0x9c9545f0.
//
// Solidity: function relay(bytes request, address relayer) payable returns()
func (_IFastBridgeV2 *IFastBridgeV2TransactorSession) Relay0(request []byte, relayer common.Address) (*types.Transaction, error) {
	return _IFastBridgeV2.Contract.Relay0(&_IFastBridgeV2.TransactOpts, request, relayer)
}

// IFastBridgeV2BridgeDepositClaimedIterator is returned from FilterBridgeDepositClaimed and is used to iterate over the raw logs and unpacked data for BridgeDepositClaimed events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositClaimedIterator struct {
	Event *IFastBridgeV2BridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeDepositClaimed)
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
		it.Event = new(IFastBridgeV2BridgeDepositClaimed)
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
func (it *IFastBridgeV2BridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeDepositClaimed represents a BridgeDepositClaimed event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositClaimed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositClaimed is a free log retrieval operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeV2BridgeDepositClaimedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeDepositClaimedIterator{contract: _IFastBridgeV2.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeDepositClaimed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeDepositClaimed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeDepositClaimed)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
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

// ParseBridgeDepositClaimed is a log parse operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 indexed transactionId, address indexed relayer, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeDepositClaimed(log types.Log) (*IFastBridgeV2BridgeDepositClaimed, error) {
	event := new(IFastBridgeV2BridgeDepositClaimed)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeDepositRefundedIterator is returned from FilterBridgeDepositRefunded and is used to iterate over the raw logs and unpacked data for BridgeDepositRefunded events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositRefundedIterator struct {
	Event *IFastBridgeV2BridgeDepositRefunded // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeDepositRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeDepositRefunded)
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
		it.Event = new(IFastBridgeV2BridgeDepositRefunded)
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
func (it *IFastBridgeV2BridgeDepositRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeDepositRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeDepositRefunded represents a BridgeDepositRefunded event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeDepositRefunded struct {
	TransactionId [32]byte
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositRefunded is a free log retrieval operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts, transactionId [][32]byte, to []common.Address) (*IFastBridgeV2BridgeDepositRefundedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeDepositRefundedIterator{contract: _IFastBridgeV2.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeDepositRefunded, transactionId [][32]byte, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeDepositRefunded", transactionIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeDepositRefunded)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
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

// ParseBridgeDepositRefunded is a log parse operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 indexed transactionId, address indexed to, address token, uint256 amount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeDepositRefunded(log types.Log) (*IFastBridgeV2BridgeDepositRefunded, error) {
	event := new(IFastBridgeV2BridgeDepositRefunded)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeProofDisputedIterator is returned from FilterBridgeProofDisputed and is used to iterate over the raw logs and unpacked data for BridgeProofDisputed events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofDisputedIterator struct {
	Event *IFastBridgeV2BridgeProofDisputed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeProofDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeProofDisputed)
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
		it.Event = new(IFastBridgeV2BridgeProofDisputed)
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
func (it *IFastBridgeV2BridgeProofDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeProofDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeProofDisputed represents a BridgeProofDisputed event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofDisputed struct {
	TransactionId [32]byte
	Relayer       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofDisputed is a free log retrieval operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeProofDisputed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeV2BridgeProofDisputedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeProofDisputedIterator{contract: _IFastBridgeV2.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeProofDisputed, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeProofDisputed", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeProofDisputed)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
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

// ParseBridgeProofDisputed is a log parse operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 indexed transactionId, address indexed relayer)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeProofDisputed(log types.Log) (*IFastBridgeV2BridgeProofDisputed, error) {
	event := new(IFastBridgeV2BridgeProofDisputed)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeProofProvidedIterator is returned from FilterBridgeProofProvided and is used to iterate over the raw logs and unpacked data for BridgeProofProvided events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofProvidedIterator struct {
	Event *IFastBridgeV2BridgeProofProvided // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeProofProvided)
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
		it.Event = new(IFastBridgeV2BridgeProofProvided)
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
func (it *IFastBridgeV2BridgeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeProofProvided represents a BridgeProofProvided event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeProofProvided struct {
	TransactionId   [32]byte
	Relayer         common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofProvided is a free log retrieval operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeProofProvided(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address) (*IFastBridgeV2BridgeProofProvidedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeProofProvidedIterator{contract: _IFastBridgeV2.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeProofProvided, transactionId [][32]byte, relayer []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeProofProvided", transactionIdRule, relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeProofProvided)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
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

// ParseBridgeProofProvided is a log parse operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 indexed transactionId, address indexed relayer, bytes32 transactionHash)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeProofProvided(log types.Log) (*IFastBridgeV2BridgeProofProvided, error) {
	event := new(IFastBridgeV2BridgeProofProvided)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeQuoteDetailsIterator is returned from FilterBridgeQuoteDetails and is used to iterate over the raw logs and unpacked data for BridgeQuoteDetails events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeQuoteDetailsIterator struct {
	Event *IFastBridgeV2BridgeQuoteDetails // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeQuoteDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeQuoteDetails)
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
		it.Event = new(IFastBridgeV2BridgeQuoteDetails)
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
func (it *IFastBridgeV2BridgeQuoteDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeQuoteDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeQuoteDetails represents a BridgeQuoteDetails event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeQuoteDetails struct {
	TransactionId [32]byte
	QuoteId       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeQuoteDetails is a free log retrieval operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeQuoteDetails(opts *bind.FilterOpts, transactionId [][32]byte) (*IFastBridgeV2BridgeQuoteDetailsIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeQuoteDetails", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeQuoteDetailsIterator{contract: _IFastBridgeV2.contract, event: "BridgeQuoteDetails", logs: logs, sub: sub}, nil
}

// WatchBridgeQuoteDetails is a free log subscription operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeQuoteDetails(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeQuoteDetails, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeQuoteDetails", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeQuoteDetails)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeQuoteDetails", log); err != nil {
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

// ParseBridgeQuoteDetails is a log parse operation binding the contract event 0x3120e2bb59c86aca6890191a589a96af3662838efa374fbdcdf4c95bfe4a6c0e.
//
// Solidity: event BridgeQuoteDetails(bytes32 indexed transactionId, bytes quoteId)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeQuoteDetails(log types.Log) (*IFastBridgeV2BridgeQuoteDetails, error) {
	event := new(IFastBridgeV2BridgeQuoteDetails)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeQuoteDetails", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeRelayedIterator is returned from FilterBridgeRelayed and is used to iterate over the raw logs and unpacked data for BridgeRelayed events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRelayedIterator struct {
	Event *IFastBridgeV2BridgeRelayed // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeRelayed)
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
		it.Event = new(IFastBridgeV2BridgeRelayed)
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
func (it *IFastBridgeV2BridgeRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeRelayed represents a BridgeRelayed event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRelayed struct {
	TransactionId  [32]byte
	Relayer        common.Address
	To             common.Address
	OriginChainId  uint32
	OriginToken    common.Address
	DestToken      common.Address
	OriginAmount   *big.Int
	DestAmount     *big.Int
	ChainGasAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeRelayed(opts *bind.FilterOpts, transactionId [][32]byte, relayer []common.Address, to []common.Address) (*IFastBridgeV2BridgeRelayedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeRelayedIterator{contract: _IFastBridgeV2.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeRelayed, transactionId [][32]byte, relayer []common.Address, to []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeRelayed", transactionIdRule, relayerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeRelayed)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0xf8ae392d784b1ea5e8881bfa586d81abf07ef4f1e2fc75f7fe51c90f05199a5c.
//
// Solidity: event BridgeRelayed(bytes32 indexed transactionId, address indexed relayer, address indexed to, uint32 originChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, uint256 chainGasAmount)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeRelayed(log types.Log) (*IFastBridgeV2BridgeRelayed, error) {
	event := new(IFastBridgeV2BridgeRelayed)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFastBridgeV2BridgeRequestedIterator is returned from FilterBridgeRequested and is used to iterate over the raw logs and unpacked data for BridgeRequested events raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRequestedIterator struct {
	Event *IFastBridgeV2BridgeRequested // Event containing the contract specifics and raw log

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
func (it *IFastBridgeV2BridgeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFastBridgeV2BridgeRequested)
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
		it.Event = new(IFastBridgeV2BridgeRequested)
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
func (it *IFastBridgeV2BridgeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFastBridgeV2BridgeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFastBridgeV2BridgeRequested represents a BridgeRequested event raised by the IFastBridgeV2 contract.
type IFastBridgeV2BridgeRequested struct {
	TransactionId [32]byte
	Sender        common.Address
	Request       []byte
	DestChainId   uint32
	OriginToken   common.Address
	DestToken     common.Address
	OriginAmount  *big.Int
	DestAmount    *big.Int
	SendChainGas  bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) FilterBridgeRequested(opts *bind.FilterOpts, transactionId [][32]byte, sender []common.Address) (*IFastBridgeV2BridgeRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.FilterLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &IFastBridgeV2BridgeRequestedIterator{contract: _IFastBridgeV2.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *IFastBridgeV2BridgeRequested, transactionId [][32]byte, sender []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _IFastBridgeV2.contract.WatchLogs(opts, "BridgeRequested", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFastBridgeV2BridgeRequested)
				if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x120ea0364f36cdac7983bcfdd55270ca09d7f9b314a2ebc425a3b01ab1d6403a.
//
// Solidity: event BridgeRequested(bytes32 indexed transactionId, address indexed sender, bytes request, uint32 destChainId, address originToken, address destToken, uint256 originAmount, uint256 destAmount, bool sendChainGas)
func (_IFastBridgeV2 *IFastBridgeV2Filterer) ParseBridgeRequested(log types.Log) (*IFastBridgeV2BridgeRequested, error) {
	event := new(IFastBridgeV2BridgeRequested)
	if err := _IFastBridgeV2.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
