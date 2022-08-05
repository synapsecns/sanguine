// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package messageharness

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

// HeaderMetaData contains all meta data concerning the Header contract.
var HeaderMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220151de2853871dd0747b9d52f6ca623fb23b91958b265dd507ad724e447a11b3c64736f6c634300080d0033",
}

// HeaderABI is the input ABI used to generate the binding from.
// Deprecated: Use HeaderMetaData.ABI instead.
var HeaderABI = HeaderMetaData.ABI

// HeaderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HeaderMetaData.Bin instead.
var HeaderBin = HeaderMetaData.Bin

// DeployHeader deploys a new Ethereum contract, binding an instance of Header to it.
func DeployHeader(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Header, error) {
	parsed, err := HeaderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HeaderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Header{HeaderCaller: HeaderCaller{contract: contract}, HeaderTransactor: HeaderTransactor{contract: contract}, HeaderFilterer: HeaderFilterer{contract: contract}}, nil
}

// Header is an auto generated Go binding around an Ethereum contract.
type Header struct {
	HeaderCaller     // Read-only binding to the contract
	HeaderTransactor // Write-only binding to the contract
	HeaderFilterer   // Log filterer for contract events
}

// HeaderCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeaderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeaderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeaderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeaderSession struct {
	Contract     *Header           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeaderCallerSession struct {
	Contract *HeaderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HeaderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeaderTransactorSession struct {
	Contract     *HeaderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeaderRaw struct {
	Contract *Header // Generic contract binding to access the raw methods on
}

// HeaderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeaderCallerRaw struct {
	Contract *HeaderCaller // Generic read-only contract binding to access the raw methods on
}

// HeaderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeaderTransactorRaw struct {
	Contract *HeaderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeader creates a new instance of Header, bound to a specific deployed contract.
func NewHeader(address common.Address, backend bind.ContractBackend) (*Header, error) {
	contract, err := bindHeader(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Header{HeaderCaller: HeaderCaller{contract: contract}, HeaderTransactor: HeaderTransactor{contract: contract}, HeaderFilterer: HeaderFilterer{contract: contract}}, nil
}

// NewHeaderCaller creates a new read-only instance of Header, bound to a specific deployed contract.
func NewHeaderCaller(address common.Address, caller bind.ContractCaller) (*HeaderCaller, error) {
	contract, err := bindHeader(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderCaller{contract: contract}, nil
}

// NewHeaderTransactor creates a new write-only instance of Header, bound to a specific deployed contract.
func NewHeaderTransactor(address common.Address, transactor bind.ContractTransactor) (*HeaderTransactor, error) {
	contract, err := bindHeader(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderTransactor{contract: contract}, nil
}

// NewHeaderFilterer creates a new log filterer instance of Header, bound to a specific deployed contract.
func NewHeaderFilterer(address common.Address, filterer bind.ContractFilterer) (*HeaderFilterer, error) {
	contract, err := bindHeader(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeaderFilterer{contract: contract}, nil
}

// bindHeader binds a generic wrapper to an already deployed contract.
func bindHeader(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeaderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Header *HeaderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Header.Contract.HeaderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Header *HeaderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Header.Contract.HeaderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Header *HeaderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Header.Contract.HeaderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Header *HeaderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Header.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Header *HeaderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Header.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Header *HeaderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Header.Contract.contract.Transact(opts, method, params...)
}

// MessageMetaData contains all meta data concerning the Message contract.
var MessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208d58afd4704cdd3440a04f310b6727d2cf48c056ac3049266e870f046ba9116064736f6c634300080d0033",
}

// MessageABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageMetaData.ABI instead.
var MessageABI = MessageMetaData.ABI

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
	parsed, err := abi.JSON(strings.NewReader(MessageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// MessageHarnessMetaData contains all meta data concerning the MessageHarness contract.
var MessageHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"body\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_originDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_body\",\"type\":\"bytes\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"optimisticSeconds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"recipientAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"tips\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c97c703a": "body(bytes)",
		"c81aa9c8": "destination(bytes)",
		"7c517825": "formatMessage(uint32,bytes32,uint32,uint32,bytes32,uint32,bytes,bytes)",
		"d7a7a72c": "leaf(bytes)",
		"46fad66e": "messageHash(uint32,bytes32,uint32,uint32,bytes32,uint32,bytes,bytes)",
		"4e765004": "nonce(bytes)",
		"7c1cfff9": "optimisticSeconds(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"985a5c31": "recipient(bytes)",
		"f45387ba": "recipientAddress(bytes)",
		"6dc3c4f7": "sender(bytes)",
		"045c6c0b": "tips(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611742806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c8063985a5c3111610081578063cb3eb0e11161005b578063cb3eb0e1146101bd578063d7a7a72c146101d0578063f45387ba146101e357600080fd5b8063985a5c3114610184578063c81aa9c814610197578063c97c703a146101aa57600080fd5b80636dc3c4f7116100b25780636dc3c4f71461014b5780637c1cfff91461015e5780637c5178251461017157600080fd5b8063045c6c0b146100d957806346fad66e146101025780634e76500414610123575b600080fd5b6100ec6100e7366004611260565b61021b565b6040516100f9919061130f565b60405180910390f35b61011561011036600461133b565b610248565b6040519081526020016100f9565b610136610131366004611260565b6102f7565b60405163ffffffff90911681526020016100f9565b610115610159366004611260565b61031e565b61013661016c366004611260565b61033a565b6100ec61017f36600461133b565b610356565b610115610192366004611260565b6103f7565b6101366101a5366004611260565b610413565b6100ec6101b8366004611260565b61042f565b6101366101cb366004611260565b61044b565b6101156101de366004611260565b610467565b6101f66101f1366004611260565b610480565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100f9565b606061024261023761022c8461049c565b62ffffff19166104aa565b62ffffff19166104ea565b92915050565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e08b811b82166022840152602683018b905289811b8216604684015288811b8216604a840152604e830188905286901b16606e8201528151808203605201815260729091019091526000906102e981858561053d565b9a9950505050505050505050565b60006102426103136103088461049c565b62ffffff191661055b565b62ffffff1916610592565b600061024261032f6103088461049c565b62ffffff19166105bc565b600061024261034b6103088461049c565b62ffffff19166105dd565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e08b811b82166022840152602683018b905289811b8216604684015288811b8216604a840152604e830188905286901b16606e8201528151808203605201815260729091019091526060906102e98185856105fe565b60006102426104086103088461049c565b62ffffff1916610678565b60006102426104246103088461049c565b62ffffff1916610699565b60606102426102376104408461049c565b62ffffff19166106ba565b600061024261045c6103088461049c565b62ffffff19166106f9565b60006102426104758361049c565b62ffffff191661071a565b60006102426104916103088461049c565b62ffffff191661073f565b600061024282610539610750565b6000816104bf62ffffff198216610539610774565b506104e1836104cf856002610898565b6104da866003610898565b60026108ca565b91505b50919050565b60606000806105078460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061052c84836020016108f4565b508181016020016040529052919050565b600061054a8484846105fe565b8051906020012090505b9392505050565b60008161057062ffffff198216610539610774565b506104e183610580856001610898565b61058b866002610898565b60016108ca565b6000816105a860015b62ffffff19831690610774565b506104e162ffffff19841660266004610acd565b6000816105c9600161059b565b506104e162ffffff19841660066020610afd565b6000816105ea600161059b565b506104e162ffffff198416604e6004610acd565b825160609060009061061260046002611453565b60ff1661061f919061147c565b90506000845182610630919061147c565b9050600161064060046002611453565b60ff16838389898960405160200161065e97969594939291906114a2565b604051602081830303815290604052925050509392505050565b600081610685600161059b565b506104e162ffffff198416602e6020610afd565b6000816106a6600161059b565b506104e162ffffff198416602a6004610acd565b6000816106cf62ffffff198216610539610774565b506104e1836106df856003610898565b601886901c6bffffffffffffffffffffffff1660036108ca565b600081610706600161059b565b506104e162ffffff19841660026004610acd565b60008161072f62ffffff198216610539610774565b506104e162ffffff198416610cef565b600061024261074d83610678565b90565b81516000906020840161076b64ffffffffff85168284610d4c565b95945050505050565b60006107808383610d93565b61089157600061079f6107938560d81c90565b64ffffffffff16610db6565b91505060006107b48464ffffffffff16610db6565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610888919061130f565b60405180910390fd5b5090919050565b600061055460028360048111156108b1576108b16113f5565b6108bb9190611540565b62ffffff198516906002610acd565b60006108e9846108da818661157d565b62ffffff198816919085610ea0565b90505b949350505050565b600062ffffff198084160361098b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f65722064657265660000000000000000000000000000000000000000000000006064820152608401610888565b61099483610f1a565b610a20576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e7465722064657265660000000000000000000000000000000000000000006064820152608401610888565b6000610a3a8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610a648560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000604051905084811115610a895760206060fd5b8285848460045afa50610ac3610a9f8760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b6000610ada826020611594565b610ae5906008611453565b60ff16610af3858585610afd565b901c949350505050565b60008160ff16600003610b1257506000610554565b610b2a8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610b4560ff8416856115b7565b1115610bd757610ba4610b668560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610b8c8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16610f57565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610888919061130f565b60208260ff161115610c6b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e2033322062797465730000000000006064820152608401610888565b600882026000610c898660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b600080610d0a8360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610d348460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b600080610d5983856115b7565b9050604051811115610d69575060005b80600003610d7e5762ffffff19915050610554565b5050606092831b9190911790911b1760181b90565b60008164ffffffffff16610da78460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff161115610e29576000610dd5826008611453565b60ff1685901c9050610de681610fc5565b61ffff16841793508160ff16601014610e0157601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610dbc565b50600f5b60ff8160ff161015610e9a576000610e46826008611453565b60ff1685901c9050610e5781610fc5565b61ffff16831792508160ff16600014610e7257601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610e2d565b50915091565b600080610ebb8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050610ed486610ff7565b84610edf87846115b7565b610ee991906115b7565b1115610efc5762ffffff199150506108ec565b610f0685826115b7565b9050610ac38364ffffffffff168286610d4c565b6000610f268260d81c90565b64ffffffffff1664ffffffffff03610f4057506000919050565b6000610f4b83610ff7565b60405110199392505050565b60606000610f6486610db6565b9150506000610f7286610db6565b9150506000610f8086610db6565b9150506000610f8e86610db6565b91505083838383604051602001610fa894939291906115cf565b604051602081830303815290604052945050505050949350505050565b6000610fd760048360ff16901c61103f565b60ff1661ffff919091161760081b610fee8261103f565b60ff1617919050565b60006110118260181c6bffffffffffffffffffffffff1690565b6110298360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b600060f08083179060ff8216900361105a5750603092915050565b8060ff1660f10361106e5750603192915050565b8060ff1660f2036110825750603292915050565b8060ff1660f3036110965750603392915050565b8060ff1660f4036110aa5750603492915050565b8060ff1660f5036110be5750603592915050565b8060ff1660f6036110d25750603692915050565b8060ff1660f7036110e65750603792915050565b8060ff1660f8036110fa5750603892915050565b8060ff1660f90361110e5750603992915050565b8060ff1660fa036111225750606192915050565b8060ff1660fb036111365750606292915050565b8060ff1660fc0361114a5750606392915050565b8060ff1660fd0361115e5750606492915050565b8060ff1660fe036111725750606592915050565b8060ff1660ff036104e45750606692915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126111c657600080fd5b813567ffffffffffffffff808211156111e1576111e1611186565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561122757611227611186565b8160405283815286602085880101111561124057600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561127257600080fd5b813567ffffffffffffffff81111561128957600080fd5b6108ec848285016111b5565b60005b838110156112b0578181015183820152602001611298565b838111156112bf576000848401525b50505050565b600081518084526112dd816020860160208601611295565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061055460208301846112c5565b803563ffffffff8116811461133657600080fd5b919050565b600080600080600080600080610100898b03121561135857600080fd5b61136189611322565b97506020890135965061137660408a01611322565b955061138460608a01611322565b94506080890135935061139960a08a01611322565b925060c089013567ffffffffffffffff808211156113b657600080fd5b6113c28c838d016111b5565b935060e08b01359150808211156113d857600080fd5b506113e58b828c016111b5565b9150509295985092959890939650565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff821660ff84168160ff048111821515161561147457611474611424565b029392505050565b600061ffff80831681851680830382111561149957611499611424565b01949350505050565b60007fffff000000000000000000000000000000000000000000000000000000000000808a60f01b168352808960f01b166002840152808860f01b166004840152808760f01b166006840152508451611502816008850160208901611295565b845190830190611519816008840160208901611295565b845191019061152f816008840160208801611295565b016008019998505050505050505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561157857611578611424565b500290565b60008282101561158f5761158f611424565b500390565b600060ff821660ff8416808210156115ae576115ae611424565b90039392505050565b600082198211156115ca576115ca611424565b500190565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610ac356fea264697066735822122091a8f3707296d120f4c00eef1ddbd608a08744cf35154a101a76438fcd3c18b064736f6c634300080d0033",
}

// MessageHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageHarnessMetaData.ABI instead.
var MessageHarnessABI = MessageHarnessMetaData.ABI

// Deprecated: Use MessageHarnessMetaData.Sigs instead.
// MessageHarnessFuncSigs maps the 4-byte function signature to its string representation.
var MessageHarnessFuncSigs = MessageHarnessMetaData.Sigs

// MessageHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageHarnessMetaData.Bin instead.
var MessageHarnessBin = MessageHarnessMetaData.Bin

// DeployMessageHarness deploys a new Ethereum contract, binding an instance of MessageHarness to it.
func DeployMessageHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MessageHarness, error) {
	parsed, err := MessageHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageHarness{MessageHarnessCaller: MessageHarnessCaller{contract: contract}, MessageHarnessTransactor: MessageHarnessTransactor{contract: contract}, MessageHarnessFilterer: MessageHarnessFilterer{contract: contract}}, nil
}

// MessageHarness is an auto generated Go binding around an Ethereum contract.
type MessageHarness struct {
	MessageHarnessCaller     // Read-only binding to the contract
	MessageHarnessTransactor // Write-only binding to the contract
	MessageHarnessFilterer   // Log filterer for contract events
}

// MessageHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageHarnessSession struct {
	Contract     *MessageHarness   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MessageHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageHarnessCallerSession struct {
	Contract *MessageHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// MessageHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageHarnessTransactorSession struct {
	Contract     *MessageHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// MessageHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageHarnessRaw struct {
	Contract *MessageHarness // Generic contract binding to access the raw methods on
}

// MessageHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageHarnessCallerRaw struct {
	Contract *MessageHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// MessageHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageHarnessTransactorRaw struct {
	Contract *MessageHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageHarness creates a new instance of MessageHarness, bound to a specific deployed contract.
func NewMessageHarness(address common.Address, backend bind.ContractBackend) (*MessageHarness, error) {
	contract, err := bindMessageHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageHarness{MessageHarnessCaller: MessageHarnessCaller{contract: contract}, MessageHarnessTransactor: MessageHarnessTransactor{contract: contract}, MessageHarnessFilterer: MessageHarnessFilterer{contract: contract}}, nil
}

// NewMessageHarnessCaller creates a new read-only instance of MessageHarness, bound to a specific deployed contract.
func NewMessageHarnessCaller(address common.Address, caller bind.ContractCaller) (*MessageHarnessCaller, error) {
	contract, err := bindMessageHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHarnessCaller{contract: contract}, nil
}

// NewMessageHarnessTransactor creates a new write-only instance of MessageHarness, bound to a specific deployed contract.
func NewMessageHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageHarnessTransactor, error) {
	contract, err := bindMessageHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageHarnessTransactor{contract: contract}, nil
}

// NewMessageHarnessFilterer creates a new log filterer instance of MessageHarness, bound to a specific deployed contract.
func NewMessageHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageHarnessFilterer, error) {
	contract, err := bindMessageHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageHarnessFilterer{contract: contract}, nil
}

// bindMessageHarness binds a generic wrapper to an already deployed contract.
func bindMessageHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MessageHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageHarness *MessageHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHarness.Contract.MessageHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageHarness *MessageHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHarness.Contract.MessageHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageHarness *MessageHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHarness.Contract.MessageHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageHarness *MessageHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageHarness *MessageHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageHarness *MessageHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageHarness.Contract.contract.Transact(opts, method, params...)
}

// Body is a free data retrieval call binding the contract method 0xc97c703a.
//
// Solidity: function body(bytes _message) view returns(bytes)
func (_MessageHarness *MessageHarnessCaller) Body(opts *bind.CallOpts, _message []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "body", _message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Body is a free data retrieval call binding the contract method 0xc97c703a.
//
// Solidity: function body(bytes _message) view returns(bytes)
func (_MessageHarness *MessageHarnessSession) Body(_message []byte) ([]byte, error) {
	return _MessageHarness.Contract.Body(&_MessageHarness.CallOpts, _message)
}

// Body is a free data retrieval call binding the contract method 0xc97c703a.
//
// Solidity: function body(bytes _message) view returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) Body(_message []byte) ([]byte, error) {
	return _MessageHarness.Contract.Body(&_MessageHarness.CallOpts, _message)
}

// Destination is a free data retrieval call binding the contract method 0xc81aa9c8.
//
// Solidity: function destination(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCaller) Destination(opts *bind.CallOpts, _message []byte) (uint32, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "destination", _message)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xc81aa9c8.
//
// Solidity: function destination(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessSession) Destination(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.Destination(&_MessageHarness.CallOpts, _message)
}

// Destination is a free data retrieval call binding the contract method 0xc81aa9c8.
//
// Solidity: function destination(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCallerSession) Destination(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.Destination(&_MessageHarness.CallOpts, _message)
}

// FormatMessage is a free data retrieval call binding the contract method 0x7c517825.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCaller) FormatMessage(opts *bind.CallOpts, _originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "formatMessage", _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _tips, _messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage is a free data retrieval call binding the contract method 0x7c517825.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// FormatMessage is a free data retrieval call binding the contract method 0x7c517825.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) Leaf(opts *bind.CallOpts, _message []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "leaf", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) Leaf(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Leaf(&_MessageHarness.CallOpts, _message)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) Leaf(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Leaf(&_MessageHarness.CallOpts, _message)
}

// MessageHash is a free data retrieval call binding the contract method 0x46fad66e.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) MessageHash(opts *bind.CallOpts, _origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _body []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "messageHash", _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _tips, _body)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0x46fad66e.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) MessageHash(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _body []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _tips, _body)
}

// MessageHash is a free data retrieval call binding the contract method 0x46fad66e.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) MessageHash(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _body []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _tips, _body)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCaller) Nonce(opts *bind.CallOpts, _message []byte) (uint32, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "nonce", _message)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessSession) Nonce(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.Nonce(&_MessageHarness.CallOpts, _message)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCallerSession) Nonce(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.Nonce(&_MessageHarness.CallOpts, _message)
}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x7c1cfff9.
//
// Solidity: function optimisticSeconds(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCaller) OptimisticSeconds(opts *bind.CallOpts, _message []byte) (uint32, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "optimisticSeconds", _message)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x7c1cfff9.
//
// Solidity: function optimisticSeconds(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessSession) OptimisticSeconds(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.OptimisticSeconds(&_MessageHarness.CallOpts, _message)
}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x7c1cfff9.
//
// Solidity: function optimisticSeconds(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCallerSession) OptimisticSeconds(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.OptimisticSeconds(&_MessageHarness.CallOpts, _message)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCaller) Origin(opts *bind.CallOpts, _message []byte) (uint32, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "origin", _message)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessSession) Origin(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.Origin(&_MessageHarness.CallOpts, _message)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCallerSession) Origin(_message []byte) (uint32, error) {
	return _MessageHarness.Contract.Origin(&_MessageHarness.CallOpts, _message)
}

// Recipient is a free data retrieval call binding the contract method 0x985a5c31.
//
// Solidity: function recipient(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) Recipient(opts *bind.CallOpts, _message []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "recipient", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x985a5c31.
//
// Solidity: function recipient(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) Recipient(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Recipient(&_MessageHarness.CallOpts, _message)
}

// Recipient is a free data retrieval call binding the contract method 0x985a5c31.
//
// Solidity: function recipient(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) Recipient(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Recipient(&_MessageHarness.CallOpts, _message)
}

// RecipientAddress is a free data retrieval call binding the contract method 0xf45387ba.
//
// Solidity: function recipientAddress(bytes _message) pure returns(address)
func (_MessageHarness *MessageHarnessCaller) RecipientAddress(opts *bind.CallOpts, _message []byte) (common.Address, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "recipientAddress", _message)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientAddress is a free data retrieval call binding the contract method 0xf45387ba.
//
// Solidity: function recipientAddress(bytes _message) pure returns(address)
func (_MessageHarness *MessageHarnessSession) RecipientAddress(_message []byte) (common.Address, error) {
	return _MessageHarness.Contract.RecipientAddress(&_MessageHarness.CallOpts, _message)
}

// RecipientAddress is a free data retrieval call binding the contract method 0xf45387ba.
//
// Solidity: function recipientAddress(bytes _message) pure returns(address)
func (_MessageHarness *MessageHarnessCallerSession) RecipientAddress(_message []byte) (common.Address, error) {
	return _MessageHarness.Contract.RecipientAddress(&_MessageHarness.CallOpts, _message)
}

// Sender is a free data retrieval call binding the contract method 0x6dc3c4f7.
//
// Solidity: function sender(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) Sender(opts *bind.CallOpts, _message []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "sender", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0x6dc3c4f7.
//
// Solidity: function sender(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) Sender(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Sender(&_MessageHarness.CallOpts, _message)
}

// Sender is a free data retrieval call binding the contract method 0x6dc3c4f7.
//
// Solidity: function sender(bytes _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) Sender(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Sender(&_MessageHarness.CallOpts, _message)
}

// Tips is a free data retrieval call binding the contract method 0x045c6c0b.
//
// Solidity: function tips(bytes _message) view returns(bytes)
func (_MessageHarness *MessageHarnessCaller) Tips(opts *bind.CallOpts, _message []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "tips", _message)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Tips is a free data retrieval call binding the contract method 0x045c6c0b.
//
// Solidity: function tips(bytes _message) view returns(bytes)
func (_MessageHarness *MessageHarnessSession) Tips(_message []byte) ([]byte, error) {
	return _MessageHarness.Contract.Tips(&_MessageHarness.CallOpts, _message)
}

// Tips is a free data retrieval call binding the contract method 0x045c6c0b.
//
// Solidity: function tips(bytes _message) view returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) Tips(_message []byte) ([]byte, error) {
	return _MessageHarness.Contract.Tips(&_MessageHarness.CallOpts, _message)
}

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220272115c0d3c96e86b1f3a1b79c1d18812f629ad36943b92e2cb880d9130a014f64736f6c634300080d0033",
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
	ABI: "[{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f26be3fc": "NULL()",
	},
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea2646970667358221220d7745ca3a0dd1dcedb2899b14e00eea6ad2532032fb3c7ffb3294b0043fa0f0864736f6c634300080d0033",
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
