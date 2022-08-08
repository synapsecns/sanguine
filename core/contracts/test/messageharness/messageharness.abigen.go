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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c8edaddd1a73e02423359ded235b6369ad4b189b3082f964ab0af5f44ca5d88264736f6c634300080d0033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cf880a41369369c56329af38ab922b24e58c3510844966fc54314d843c59948764736f6c634300080d0033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"body\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_originDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"_updaterTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_relayerTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_proverTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_processorTip\",\"type\":\"uint96\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headerOffset\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_body\",\"type\":\"bytes\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"optimisticSeconds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"recipientAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"tips\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c97c703a": "body(bytes)",
		"c81aa9c8": "destination(bytes)",
		"81d030ef": "formatMessage(uint32,bytes32,uint32,uint32,bytes32,uint32,uint96,uint96,uint96,uint96,bytes)",
		"639c5570": "headerOffset()",
		"d7a7a72c": "leaf(bytes)",
		"46fad66e": "messageHash(uint32,bytes32,uint32,uint32,bytes32,uint32,bytes,bytes)",
		"52617f3c": "messageVersion()",
		"4e765004": "nonce(bytes)",
		"7c1cfff9": "optimisticSeconds(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"985a5c31": "recipient(bytes)",
		"f45387ba": "recipientAddress(bytes)",
		"6dc3c4f7": "sender(bytes)",
		"045c6c0b": "tips(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061193e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806381d030ef1161008c578063c97c703a11610066578063c97c703a146101de578063cb3eb0e1146101f1578063d7a7a72c14610204578063f45387ba1461021757600080fd5b806381d030ef146101a5578063985a5c31146101b8578063c81aa9c8146101cb57600080fd5b806352617f3c116100c857806352617f3c14610161578063639c5570146101775780636dc3c4f71461017f5780637c1cfff91461019257600080fd5b8063045c6c0b146100ef57806346fad66e146101185780634e76500414610139575b600080fd5b6101026100fd366004611367565b61024f565b60405161010f9190611416565b60405180910390f35b61012b610126366004611442565b61027c565b60405190815260200161010f565b61014c610147366004611367565b61032b565b60405163ffffffff909116815260200161010f565b60015b60405161ffff909116815260200161010f565b610164610352565b61012b61018d366004611367565b610368565b61014c6101a0366004611367565b610384565b6101026101b3366004611518565b6103a0565b61012b6101c6366004611367565b610477565b61014c6101d9366004611367565b610493565b6101026101ec366004611367565b6104af565b61014c6101ff366004611367565b6104cb565b61012b610212366004611367565b6104e7565b61022a610225366004611367565b610500565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010f565b606061027661026b6102608461051c565b62ffffff191661052a565b62ffffff191661056a565b92915050565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e08b811b82166022840152602683018b905289811b8216604684015288811b8216604a840152604e830188905286901b16606e82015281518082036052018152607290910190915260009061031d8185856105bd565b9a9950505050505050505050565b600061027661034761033c8461051c565b62ffffff19166105db565b62ffffff1916610612565b60006103606004600261164f565b60ff16905090565b600061027661037961033c8461051c565b62ffffff191661063c565b600061027661039561033c8461051c565b62ffffff191661065d565b606060006103b08787878761067e565b905060006104588e8e8e8e8e8e604080517e01000000000000000000000000000000000000000000000000000000000000602082015260e097881b7fffffffff000000000000000000000000000000000000000000000000000000009081166022830152602682019790975294871b8616604686015292861b8516604a850152604e84019190915290931b909116606e82015281516052818303018152607290910190915290565b9050610465818386610710565b9e9d5050505050505050505050505050565b600061027661048861033c8461051c565b62ffffff191661078a565b60006102766104a461033c8461051c565b62ffffff19166107ab565b606061027661026b6104c08461051c565b62ffffff19166107cc565b60006102766104dc61033c8461051c565b62ffffff191661080b565b60006102766104f58361051c565b62ffffff191661082c565b600061027661051161033c8461051c565b62ffffff1916610851565b600061027682610539610862565b60008161053f62ffffff198216610539610886565b506105618361054f8560026109aa565b61055a8660036109aa565b60026109dc565b91505b50919050565b60606000806105878460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060405191508192506105ac84836020016109fb565b508181016020016040529052919050565b60006105ca848484610710565b8051906020012090505b9392505050565b6000816105f062ffffff198216610539610886565b50610561836106008560016109aa565b61060b8660026109aa565b60016109dc565b60008161062860015b62ffffff19831690610886565b5061056162ffffff19841660266004610bd4565b600081610649600161061b565b5061056162ffffff19841660066020610c04565b60008161066a600161061b565b5061056162ffffff198416604e6004610bd4565b6040517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffffffffffff000000000000000000000000000000000000000060a086811b8216602284015285811b8216602e84015284811b8216603a84015283901b16604682015260609060520160405160208183030381529060405290505b949350505050565b82516060906000906107246004600261164f565b60ff166107319190611678565b905060008451826107429190611678565b905060016107526004600261164f565b60ff168383898989604051602001610770979695949392919061169e565b604051602081830303815290604052925050509392505050565b600081610797600161061b565b5061056162ffffff198416602e6020610c04565b6000816107b8600161061b565b5061056162ffffff198416602a6004610bd4565b6000816107e162ffffff198216610539610886565b50610561836107f18560036109aa565b601886901c6bffffffffffffffffffffffff1660036109dc565b600081610818600161061b565b5061056162ffffff19841660026004610bd4565b60008161084162ffffff198216610539610886565b5061056162ffffff198416610df6565b600061027661085f8361078a565b90565b81516000906020840161087d64ffffffffff85168284610e53565b95945050505050565b60006108928383610e9a565b6109a35760006108b16108a58560d81c90565b64ffffffffff16610ebd565b91505060006108c68464ffffffffff16610ebd565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099a9190611416565b60405180910390fd5b5090919050565b60006105d460028360048111156109c3576109c36115f1565b6109cd919061173c565b62ffffff198516906002610bd4565b600061087d846109ec8186611779565b62ffffff198816919085610fa7565b600062ffffff1980841603610a92576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f6572206465726566000000000000000000000000000000000000000000000000606482015260840161099a565b610a9b83611021565b610b27576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e746572206465726566000000000000000000000000000000000000000000606482015260840161099a565b6000610b418460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610b6b8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000604051905084811115610b905760206060fd5b8285848460045afa50610bca610ba68760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b6000610be1826020611790565b610bec90600861164f565b60ff16610bfa858585610c04565b901c949350505050565b60008160ff16600003610c19575060006105d4565b610c318460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610c4c60ff8416856117b3565b1115610cde57610cab610c6d8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610c938660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff1661105e565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099a9190611416565b60208260ff161115610d72576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e203332206279746573000000000000606482015260840161099a565b600882026000610d908660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b600080610e118360781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610e3b8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b600080610e6083856117b3565b9050604051811115610e70575060005b80600003610e855762ffffff199150506105d4565b5050606092831b9190911790911b1760181b90565b60008164ffffffffff16610eae8460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff161115610f30576000610edc82600861164f565b60ff1685901c9050610eed816110cc565b61ffff16841793508160ff16601014610f0857601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610ec3565b50600f5b60ff8160ff161015610fa1576000610f4d82600861164f565b60ff1685901c9050610f5e816110cc565b61ffff16831792508160ff16600014610f7957601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610f34565b50915091565b600080610fc28660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050610fdb866110fe565b84610fe687846117b3565b610ff091906117b3565b11156110035762ffffff19915050610708565b61100d85826117b3565b9050610bca8364ffffffffff168286610e53565b600061102d8260d81c90565b64ffffffffff1664ffffffffff0361104757506000919050565b6000611052836110fe565b60405110199392505050565b6060600061106b86610ebd565b915050600061107986610ebd565b915050600061108786610ebd565b915050600061109586610ebd565b915050838383836040516020016110af94939291906117cb565b604051602081830303815290604052945050505050949350505050565b60006110de60048360ff16901c611146565b60ff1661ffff919091161760081b6110f582611146565b60ff1617919050565b60006111188260181c6bffffffffffffffffffffffff1690565b6111308360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b600060f08083179060ff821690036111615750603092915050565b8060ff1660f1036111755750603192915050565b8060ff1660f2036111895750603292915050565b8060ff1660f30361119d5750603392915050565b8060ff1660f4036111b15750603492915050565b8060ff1660f5036111c55750603592915050565b8060ff1660f6036111d95750603692915050565b8060ff1660f7036111ed5750603792915050565b8060ff1660f8036112015750603892915050565b8060ff1660f9036112155750603992915050565b8060ff1660fa036112295750606192915050565b8060ff1660fb0361123d5750606292915050565b8060ff1660fc036112515750606392915050565b8060ff1660fd036112655750606492915050565b8060ff1660fe036112795750606592915050565b8060ff1660ff036105645750606692915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126112cd57600080fd5b813567ffffffffffffffff808211156112e8576112e861128d565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561132e5761132e61128d565b8160405283815286602085880101111561134757600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561137957600080fd5b813567ffffffffffffffff81111561139057600080fd5b610708848285016112bc565b60005b838110156113b757818101518382015260200161139f565b838111156113c6576000848401525b50505050565b600081518084526113e481602086016020860161139c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006105d460208301846113cc565b803563ffffffff8116811461143d57600080fd5b919050565b600080600080600080600080610100898b03121561145f57600080fd5b61146889611429565b97506020890135965061147d60408a01611429565b955061148b60608a01611429565b9450608089013593506114a060a08a01611429565b925060c089013567ffffffffffffffff808211156114bd57600080fd5b6114c98c838d016112bc565b935060e08b01359150808211156114df57600080fd5b506114ec8b828c016112bc565b9150509295985092959890939650565b80356bffffffffffffffffffffffff8116811461143d57600080fd5b60008060008060008060008060008060006101608c8e03121561153a57600080fd5b6115438c611429565b9a5060208c0135995061155860408d01611429565b985061156660608d01611429565b975060808c0135965061157b60a08d01611429565b955061158960c08d016114fc565b945061159760e08d016114fc565b93506115a66101008d016114fc565b92506115b56101208d016114fc565b91506101408c013567ffffffffffffffff8111156115d257600080fd5b6115de8e828f016112bc565b9150509295989b509295989b9093969950565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600060ff821660ff84168160ff048111821515161561167057611670611620565b029392505050565b600061ffff80831681851680830382111561169557611695611620565b01949350505050565b60007fffff000000000000000000000000000000000000000000000000000000000000808a60f01b168352808960f01b166002840152808860f01b166004840152808760f01b1660068401525084516116fe81600885016020890161139c565b84519083019061171581600884016020890161139c565b845191019061172b81600884016020880161139c565b016008019998505050505050505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561177457611774611620565b500290565b60008282101561178b5761178b611620565b500390565b600060ff821660ff8416808210156117aa576117aa611620565b90039392505050565b600082198211156117c6576117c6611620565b500190565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610bca56fea26469706673582212200d433cf58b1c411e28dea26e6cfce55626c077761ae2aefdef024731ec7fd18464736f6c634300080d0033",
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

// FormatMessage is a free data retrieval call binding the contract method 0x81d030ef.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, uint96 _updaterTip, uint96 _relayerTip, uint96 _proverTip, uint96 _processorTip, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCaller) FormatMessage(opts *bind.CallOpts, _originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _updaterTip *big.Int, _relayerTip *big.Int, _proverTip *big.Int, _processorTip *big.Int, _messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "formatMessage", _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _updaterTip, _relayerTip, _proverTip, _processorTip, _messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage is a free data retrieval call binding the contract method 0x81d030ef.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, uint96 _updaterTip, uint96 _relayerTip, uint96 _proverTip, uint96 _processorTip, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _updaterTip *big.Int, _relayerTip *big.Int, _proverTip *big.Int, _processorTip *big.Int, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _updaterTip, _relayerTip, _proverTip, _processorTip, _messageBody)
}

// FormatMessage is a free data retrieval call binding the contract method 0x81d030ef.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, uint96 _updaterTip, uint96 _relayerTip, uint96 _proverTip, uint96 _processorTip, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _updaterTip *big.Int, _relayerTip *big.Int, _proverTip *big.Int, _processorTip *big.Int, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _updaterTip, _relayerTip, _proverTip, _processorTip, _messageBody)
}

// HeaderOffset is a free data retrieval call binding the contract method 0x639c5570.
//
// Solidity: function headerOffset() pure returns(uint16)
func (_MessageHarness *MessageHarnessCaller) HeaderOffset(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "headerOffset")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// HeaderOffset is a free data retrieval call binding the contract method 0x639c5570.
//
// Solidity: function headerOffset() pure returns(uint16)
func (_MessageHarness *MessageHarnessSession) HeaderOffset() (uint16, error) {
	return _MessageHarness.Contract.HeaderOffset(&_MessageHarness.CallOpts)
}

// HeaderOffset is a free data retrieval call binding the contract method 0x639c5570.
//
// Solidity: function headerOffset() pure returns(uint16)
func (_MessageHarness *MessageHarnessCallerSession) HeaderOffset() (uint16, error) {
	return _MessageHarness.Contract.HeaderOffset(&_MessageHarness.CallOpts)
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

// MessageVersion is a free data retrieval call binding the contract method 0x52617f3c.
//
// Solidity: function messageVersion() pure returns(uint16)
func (_MessageHarness *MessageHarnessCaller) MessageVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "messageVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MessageVersion is a free data retrieval call binding the contract method 0x52617f3c.
//
// Solidity: function messageVersion() pure returns(uint16)
func (_MessageHarness *MessageHarnessSession) MessageVersion() (uint16, error) {
	return _MessageHarness.Contract.MessageVersion(&_MessageHarness.CallOpts)
}

// MessageVersion is a free data retrieval call binding the contract method 0x52617f3c.
//
// Solidity: function messageVersion() pure returns(uint16)
func (_MessageHarness *MessageHarnessCallerSession) MessageVersion() (uint16, error) {
	return _MessageHarness.Contract.MessageVersion(&_MessageHarness.CallOpts)
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

// TipsMetaData contains all meta data concerning the Tips contract.
var TipsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b0756d9546e7b2b4430f26d5c0341e164889948353dd82e994099eb78e3d3c6f64736f6c634300080d0033",
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c35e25f0946b278c613f5ebbafe60a8a7c609ce7200241743e59bf36c324e16f64736f6c634300080d0033",
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
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea2646970667358221220c22f1018f82777d7b0a5e7495d6d8c35c5962e2ace30e17414b6ee6723f7aaad64736f6c634300080d0033",
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
