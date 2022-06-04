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

// MessageMetaData contains all meta data concerning the Message contract.
var MessageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208646c6017fddc2d31041252c0c6093f8ea36f61cb11b6bce878eae37636b3fbb64736f6c634300080d0033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"body\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_originDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_body\",\"type\":\"bytes\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"optimisticSeconds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"recipientAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_message\",\"type\":\"bytes\"}],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c97c703a": "body(bytes)",
		"c81aa9c8": "destination(bytes)",
		"9a1b300a": "formatMessage(uint32,bytes32,uint32,uint32,bytes32,uint32,bytes)",
		"d7a7a72c": "leaf(bytes)",
		"311f82e3": "messageHash(uint32,bytes32,uint32,uint32,bytes32,uint32,bytes)",
		"4e765004": "nonce(bytes)",
		"7c1cfff9": "optimisticSeconds(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"985a5c31": "recipient(bytes)",
		"f45387ba": "recipientAddress(bytes)",
		"6dc3c4f7": "sender(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611223806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80639a1b300a11610081578063cb3eb0e11161005b578063cb3eb0e11461019b578063d7a7a72c146101ae578063f45387ba146101c157600080fd5b80639a1b300a14610155578063c81aa9c814610175578063c97c703a1461018857600080fd5b80636dc3c4f7116100b25780636dc3c4f71461011c5780637c1cfff91461012f578063985a5c311461014257600080fd5b8063311f82e3146100ce5780634e765004146100f4575b600080fd5b6100e16100dc366004610e32565b6101f9565b6040519081526020015b60405180910390f35b610107610102366004610ec5565b610216565b60405163ffffffff90911681526020016100eb565b6100e161012a366004610ec5565b610236565b61010761013d366004610ec5565b610250565b6100e1610150366004610ec5565b61026a565b610168610163366004610e32565b610284565b6040516100eb9190610f74565b610107610183366004610ec5565b610295565b610168610196366004610ec5565b6102af565b6101076101a9366004610ec5565b6102d8565b6100e16101bc366004610ec5565b6102f2565b6101d46101cf366004610ec5565b61030c565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100eb565b600061020a88888888888888610326565b98975050505050505050565b6000610230610225838361034b565b62ffffff191661036f565b92915050565b6000610230610245838361034b565b62ffffff1916610384565b600061023061025f838361034b565b62ffffff1916610399565b6000610230610279838361034b565b62ffffff19166103ae565b606061020a888888888888886103c3565b60006102306102a4838361034b565b62ffffff19166103fe565b60606102306102cd6102c284600061034b565b62ffffff1916610413565b62ffffff1916610446565b60006102306102e7838361034b565b62ffffff1916610499565b6000610230610301838361034b565b62ffffff19166104ad565b600061023061031b838361034b565b62ffffff19166104fe565b6000610337888888888888886103c3565b805190602001209050979650505050505050565b81516000906020840161036664ffffffffff8516828461050f565b95945050505050565b600061023062ffffff19831660246004610559565b600061023062ffffff19831660046020610589565b600061023062ffffff198316604c6004610559565b600061023062ffffff198316602c6020610589565b6060878787878787876040516020016103e29796959493929190610f87565b6040516020818303038152906040529050979650505050505050565b600061023062ffffff19831660286004610559565b6000610230605061043681601886901c6bffffffffffffffffffffffff16611035565b62ffffff19851691906000610784565b60606000806104638460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506040519150819250610488848360200161080a565b508181016020016040529052919050565b600061023062ffffff198316826004610559565b60006102306104bb83610499565b6104c484610384565b6104cd8561036f565b6104d6866103fe565b6104df876103ae565b6104e888610399565b6104f96104f48a610413565b610446565b610326565b600061023061050c836103ae565b90565b60008061051c838561104c565b905060405181111561052c575060005b806000036105415762ffffff19915050610552565b5050606083811b8317901b811760181b5b9392505050565b6000610566826020611064565b610571906008611087565b60ff1661057f858585610589565b901c949350505050565b60008160ff1660000361059e57506000610552565b6105b68460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166105d160ff84168561104c565b111561066c576106306105f28560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166106188660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff166109e3565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106639190610f74565b60405180910390fd5b60208260ff161115610700576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e2033322062797465730000000000006064820152608401610663565b60088202600061071e8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b60008061079f8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506107b886610a51565b846107c3878461104c565b6107cd919061104c565b11156107e05762ffffff19915050610802565b6107ea858261104c565b90506107fe8364ffffffffff16828661050f565b9150505b949350505050565b600062ffffff19808416036108a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f65722064657265660000000000000000000000000000000000000000000000006064820152608401610663565b6108aa83610a99565b610936576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e7465722064657265660000000000000000000000000000000000000000006064820152608401610663565b60006109508460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061097a8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600060405190508481111561099f5760206060fd5b8285848460045afa506109d96109b58760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b606060006109f086610ad6565b91505060006109fe86610ad6565b9150506000610a0c86610ad6565b9150506000610a1a86610ad6565b91505083838383604051602001610a3494939291906110b0565b604051602081830303815290604052945050505050949350505050565b6000610a6b8260181c6bffffffffffffffffffffffff1690565b610a838360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b6000610aa58260d81c90565b64ffffffffff1664ffffffffff03610abf57506000919050565b6000610aca83610a51565b60405110199392505050565b600080601f5b600f8160ff161115610b49576000610af5826008611087565b60ff1685901c9050610b0681610bc0565b61ffff16841793508160ff16601014610b2157601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610adc565b50600f5b60ff8160ff161015610bba576000610b66826008611087565b60ff1685901c9050610b7781610bc0565b61ffff16831792508160ff16600014610b9257601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610b4d565b50915091565b6000610bd260048360ff16901c610bf2565b60ff1661ffff919091161760081b610be982610bf2565b60ff1617919050565b600060f08083179060ff82169003610c0d5750603092915050565b8060ff1660f103610c215750603192915050565b8060ff1660f203610c355750603292915050565b8060ff1660f303610c495750603392915050565b8060ff1660f403610c5d5750603492915050565b8060ff1660f503610c715750603592915050565b8060ff1660f603610c855750603692915050565b8060ff1660f703610c995750603792915050565b8060ff1660f803610cad5750603892915050565b8060ff1660f903610cc15750603992915050565b8060ff1660fa03610cd55750606192915050565b8060ff1660fb03610ce95750606292915050565b8060ff1660fc03610cfd5750606392915050565b8060ff1660fd03610d115750606492915050565b8060ff1660fe03610d255750606592915050565b8060ff1660ff03610d395750606692915050565b50919050565b803563ffffffff81168114610d5357600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610d9857600080fd5b813567ffffffffffffffff80821115610db357610db3610d58565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610df957610df9610d58565b81604052838152866020858801011115610e1257600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600080600060e0888a031215610e4d57600080fd5b610e5688610d3f565b965060208801359550610e6b60408901610d3f565b9450610e7960608901610d3f565b935060808801359250610e8e60a08901610d3f565b915060c088013567ffffffffffffffff811115610eaa57600080fd5b610eb68a828b01610d87565b91505092959891949750929550565b600060208284031215610ed757600080fd5b813567ffffffffffffffff811115610eee57600080fd5b61080284828501610d87565b60005b83811015610f15578181015183820152602001610efd565b83811115610f24576000848401525b50505050565b60008151808452610f42816020860160208601610efa565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006105526020830184610f2a565b60007fffffffff00000000000000000000000000000000000000000000000000000000808a60e01b168352886004840152808860e01b166024840152808760e01b16602884015285602c840152808560e01b16604c840152508251610ff3816050850160208701610efa565b9190910160500198975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008282101561104757611047611006565b500390565b6000821982111561105f5761105f611006565b500190565b600060ff821660ff84168082101561107e5761107e611006565b90039392505050565b600060ff821660ff84168160ff04811182151516156110a8576110a8611006565b029392505050565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d82016109d956fea2646970667358221220d6f14ac7903e5fe192d3954445afcd0a84175f2fb356db10b29a79a89fb43ece64736f6c634300080d0033",
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

// FormatMessage is a free data retrieval call binding the contract method 0x9a1b300a.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCaller) FormatMessage(opts *bind.CallOpts, _originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "formatMessage", _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage is a free data retrieval call binding the contract method 0x9a1b300a.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _messageBody)
}

// FormatMessage is a free data retrieval call binding the contract method 0x9a1b300a.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, uint32 _optimisticSeconds, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _optimisticSeconds uint32, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _optimisticSeconds, _messageBody)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes _message) view returns(bytes32)
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
// Solidity: function leaf(bytes _message) view returns(bytes32)
func (_MessageHarness *MessageHarnessSession) Leaf(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Leaf(&_MessageHarness.CallOpts, _message)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes _message) view returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) Leaf(_message []byte) ([32]byte, error) {
	return _MessageHarness.Contract.Leaf(&_MessageHarness.CallOpts, _message)
}

// MessageHash is a free data retrieval call binding the contract method 0x311f82e3.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) MessageHash(opts *bind.CallOpts, _origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _body []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "messageHash", _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _body)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0x311f82e3.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) MessageHash(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _body []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _body)
}

// MessageHash is a free data retrieval call binding the contract method 0x311f82e3.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) MessageHash(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _body []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _body)
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d448e8387a2513d63cbeeac1ceb7053391478ba4b5a8413e1df0284d792537e464736f6c634300080d0033",
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
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea26469706673582212204fa828d35af921efe5d4e6f0a9cdd65ffd9f86461db607d57acc306abc0f4b5664736f6c634300080d0033",
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
