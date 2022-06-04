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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c6c72b981e8a15eed6f2f4fe808a8b74e3166264e4d012d1edcea5d41bcd125464736f6c634300080d0033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"body\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_originDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_body\",\"type\":\"bytes\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"recipientAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes29\",\"name\":\"_message\",\"type\":\"bytes29\"}],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"70cb2c85": "body(bytes29)",
		"11cd4add": "destination(bytes29)",
		"2651fb15": "formatMessage(uint32,bytes32,uint32,uint32,bytes32,bytes)",
		"db681644": "leaf(bytes29)",
		"3cd63012": "messageHash(uint32,bytes32,uint32,uint32,bytes32,bytes)",
		"58c8a98f": "nonce(bytes29)",
		"a9f877ad": "origin(bytes29)",
		"418fad6d": "recipient(bytes29)",
		"6eddb7f3": "recipientAddress(bytes29)",
		"a0d41e58": "sender(bytes29)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611143806100206000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c80636eddb7f311610076578063a0d41e581161005b578063a0d41e58146101b7578063a9f877ad146101ca578063db681644146101dd57600080fd5b80636eddb7f31461015757806370cb2c851461018f57600080fd5b80633cd63012116100a75780633cd6301214610110578063418fad6d1461013157806358c8a98f1461014457600080fd5b806311cd4add146100c35780632651fb15146100f0575b600080fd5b6100d66100d1366004610c75565b6101f0565b60405163ffffffff90911681526020015b60405180910390f35b6101036100fe366004610ce3565b610207565b6040516100e79190610e70565b61012361011e366004610ce3565b610222565b6040519081526020016100e7565b61012361013f366004610c75565b610232565b6100d6610152366004610c75565b610243565b61016a610165366004610c75565b610254565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e7565b6101a261019d366004610c75565b610265565b60405162ffffff1990911681526020016100e7565b6101236101c5366004610c75565b610276565b6100d66101d8366004610c75565b610287565b6101236101eb366004610c75565b610298565b600061020162ffffff1983166102a9565b92915050565b60606102178787878787876102be565b979650505050505050565b60006102178787878787876102f6565b600061020162ffffff198316610319565b600061020162ffffff19831661032e565b600061020162ffffff198316610343565b600061020162ffffff198316610354565b600061020162ffffff198316610387565b600061020162ffffff19831661039c565b600061020162ffffff1983166103b0565b600061020162ffffff198316602860046103f8565b60608686868686866040516020016102db96959493929190610e83565b60405160208183030381529060405290509695505050505050565b60006103068787878787876102be565b8051906020012090509695505050505050565b600061020162ffffff198316602c602061042a565b600061020162ffffff198316602460046103f8565b600061020161035183610319565b90565b6000610201604c61037781601886901c6bffffffffffffffffffffffff16610f25565b62ffffff19851691906000610630565b600061020162ffffff1983166004602061042a565b600061020162ffffff1983168260046103f8565b60006102016103be8361039c565b6103c784610387565b6103d08561032e565b6103d9866102a9565b6103e287610319565b6103f36103ee89610354565b6106b5565b6102f6565b6000610405826020610f3c565b610410906008610f5f565b60ff1661041e85858561042a565b901c90505b9392505050565b60008160ff1660000361043f57506000610423565b6104578460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166104728460ff851661070f565b111561050d576104d16104938560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166104b98660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff1661071b565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105049190610e70565b60405180910390fd5b60208260ff1611156105a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f54797065644d656d566965772f696e646578202d20417474656d70746564207460448201527f6f20696e646578206d6f7265207468616e2033322062797465730000000000006064820152608401610504565b60006105ae836008610f5f565b905060006105ca8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b60008061064b8660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905061066486610789565b61067885610672848961070f565b9061070f565b111561068b5762ffffff199150506106ad565b610695818661070f565b90506106a98364ffffffffff1682866107d9565b9150505b949350505050565b60606000806106d28460181c6bffffffffffffffffffffffff1690565b60405193508392506bffffffffffffffffffffffff1690506106fe846106f9846020610f88565b610820565b508181016020016040529052919050565b60006104238284610f88565b60606000610728866109f9565b9150506000610736866109f9565b9150506000610744866109f9565b9150506000610752866109f9565b9150508383838360405160200161076c9493929190610fa0565b604051602081830303815290604052945050505050949350505050565b60006107a38260181c6bffffffffffffffffffffffff1690565b6107bb8360781c6bffffffffffffffffffffffff1690565b6107c591906110dd565b6bffffffffffffffffffffffff1692915050565b6000806107e6848461070f565b90506040518111156107f6575060005b8060000361080b5762ffffff19915050610423565b5050606092831b9190911790911b1760181b90565b600062ffffff19808416036108b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f54797065644d656d566965772f636f7079546f202d204e756c6c20706f696e7460448201527f65722064657265660000000000000000000000000000000000000000000000006064820152608401610504565b6108c083610ab9565b61094c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f54797065644d656d566965772f636f7079546f202d20496e76616c696420706f60448201527f696e7465722064657265660000000000000000000000000000000000000000006064820152608401610504565b60006109668460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006109908560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006040519050848111156109b55760206060fd5b8285848460045afa506109ef6109cb8760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b9695505050505050565b600080601f5b600f8160ff161115610a57576000610a18826008610f5f565b60ff1685901c9050610a2981610af6565b61ffff16841793508160ff16601014610a4457601084901b93505b50610a50600182610f3c565b90506109ff565b50600f5b60ff8160ff161015610ab3576000610a74826008610f5f565b60ff1685901c9050610a8581610af6565b61ffff16831792508160ff16600014610aa057601083901b92505b50610aac600182610f3c565b9050610a5b565b50915091565b6000610ac58260d81c90565b64ffffffffff1664ffffffffff03610adf57506000919050565b6000610aea83610789565b60405110199392505050565b6000610b0860048360ff16901c610b28565b60ff1661ffff919091161760081b610b1f82610b28565b60ff1617919050565b600060f08083179060ff82169003610b435750603092915050565b8060ff1660f103610b575750603192915050565b8060ff1660f203610b6b5750603292915050565b8060ff1660f303610b7f5750603392915050565b8060ff1660f403610b935750603492915050565b8060ff1660f503610ba75750603592915050565b8060ff1660f603610bbb5750603692915050565b8060ff1660f703610bcf5750603792915050565b8060ff1660f803610be35750603892915050565b8060ff1660f903610bf75750603992915050565b8060ff1660fa03610c0b5750606192915050565b8060ff1660fb03610c1f5750606292915050565b8060ff1660fc03610c335750606392915050565b8060ff1660fd03610c475750606492915050565b8060ff1660fe03610c5b5750606592915050565b8060ff1660ff03610c6f5750606692915050565b50919050565b600060208284031215610c8757600080fd5b813562ffffff198116811461042357600080fd5b803563ffffffff81168114610caf57600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008060008060008060c08789031215610cfc57600080fd5b610d0587610c9b565b955060208701359450610d1a60408801610c9b565b9350610d2860608801610c9b565b92506080870135915060a087013567ffffffffffffffff80821115610d4c57600080fd5b818901915089601f830112610d6057600080fd5b813581811115610d7257610d72610cb4565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610db857610db8610cb4565b816040528281528c6020848701011115610dd157600080fd5b8260208601602083013760006020848301015280955050505050509295509295509295565b60005b83811015610e11578181015183820152602001610df9565b83811115610e20576000848401525b50505050565b60008151808452610e3e816020860160208601610df6565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006104236020830184610e26565b60007fffffffff00000000000000000000000000000000000000000000000000000000808960e01b168352876004840152808760e01b166024840152808660e01b1660288401525083602c8301528251610ee481604c850160208701610df6565b91909101604c01979650505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015610f3757610f37610ef6565b500390565b600060ff821660ff841680821015610f5657610f56610ef6565b90039392505050565b600060ff821660ff84168160ff0481118215151615610f8057610f80610ef6565b029392505050565b60008219821115610f9b57610f9b610ef6565b500190565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d82016109ef565b60006bffffffffffffffffffffffff80831681851680830382111561110457611104610ef6565b0194935050505056fea26469706673582212206de0e1f4b2f228ebeab769a6f80b2230f559880b1a622d17b01f0cf69e67058c64736f6c634300080d0033",
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

// Body is a free data retrieval call binding the contract method 0x70cb2c85.
//
// Solidity: function body(bytes29 _message) pure returns(bytes29)
func (_MessageHarness *MessageHarnessCaller) Body(opts *bind.CallOpts, _message [29]byte) ([29]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "body", _message)

	if err != nil {
		return *new([29]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([29]byte)).(*[29]byte)

	return out0, err

}

// Body is a free data retrieval call binding the contract method 0x70cb2c85.
//
// Solidity: function body(bytes29 _message) pure returns(bytes29)
func (_MessageHarness *MessageHarnessSession) Body(_message [29]byte) ([29]byte, error) {
	return _MessageHarness.Contract.Body(&_MessageHarness.CallOpts, _message)
}

// Body is a free data retrieval call binding the contract method 0x70cb2c85.
//
// Solidity: function body(bytes29 _message) pure returns(bytes29)
func (_MessageHarness *MessageHarnessCallerSession) Body(_message [29]byte) ([29]byte, error) {
	return _MessageHarness.Contract.Body(&_MessageHarness.CallOpts, _message)
}

// Destination is a free data retrieval call binding the contract method 0x11cd4add.
//
// Solidity: function destination(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCaller) Destination(opts *bind.CallOpts, _message [29]byte) (uint32, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "destination", _message)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0x11cd4add.
//
// Solidity: function destination(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessSession) Destination(_message [29]byte) (uint32, error) {
	return _MessageHarness.Contract.Destination(&_MessageHarness.CallOpts, _message)
}

// Destination is a free data retrieval call binding the contract method 0x11cd4add.
//
// Solidity: function destination(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCallerSession) Destination(_message [29]byte) (uint32, error) {
	return _MessageHarness.Contract.Destination(&_MessageHarness.CallOpts, _message)
}

// FormatMessage is a free data retrieval call binding the contract method 0x2651fb15.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCaller) FormatMessage(opts *bind.CallOpts, _originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "formatMessage", _originDomain, _sender, _nonce, _destinationDomain, _recipient, _messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage is a free data retrieval call binding the contract method 0x2651fb15.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _messageBody)
}

// FormatMessage is a free data retrieval call binding the contract method 0x2651fb15.
//
// Solidity: function formatMessage(uint32 _originDomain, bytes32 _sender, uint32 _nonce, uint32 _destinationDomain, bytes32 _recipient, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) FormatMessage(_originDomain uint32, _sender [32]byte, _nonce uint32, _destinationDomain uint32, _recipient [32]byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _originDomain, _sender, _nonce, _destinationDomain, _recipient, _messageBody)
}

// Leaf is a free data retrieval call binding the contract method 0xdb681644.
//
// Solidity: function leaf(bytes29 _message) view returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) Leaf(opts *bind.CallOpts, _message [29]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "leaf", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Leaf is a free data retrieval call binding the contract method 0xdb681644.
//
// Solidity: function leaf(bytes29 _message) view returns(bytes32)
func (_MessageHarness *MessageHarnessSession) Leaf(_message [29]byte) ([32]byte, error) {
	return _MessageHarness.Contract.Leaf(&_MessageHarness.CallOpts, _message)
}

// Leaf is a free data retrieval call binding the contract method 0xdb681644.
//
// Solidity: function leaf(bytes29 _message) view returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) Leaf(_message [29]byte) ([32]byte, error) {
	return _MessageHarness.Contract.Leaf(&_MessageHarness.CallOpts, _message)
}

// MessageHash is a free data retrieval call binding the contract method 0x3cd63012.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) MessageHash(opts *bind.CallOpts, _origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _body []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "messageHash", _origin, _sender, _nonce, _destination, _recipient, _body)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0x3cd63012.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) MessageHash(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _body []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _body)
}

// MessageHash is a free data retrieval call binding the contract method 0x3cd63012.
//
// Solidity: function messageHash(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, bytes _body) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) MessageHash(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _body []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _body)
}

// Nonce is a free data retrieval call binding the contract method 0x58c8a98f.
//
// Solidity: function nonce(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCaller) Nonce(opts *bind.CallOpts, _message [29]byte) (uint32, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "nonce", _message)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x58c8a98f.
//
// Solidity: function nonce(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessSession) Nonce(_message [29]byte) (uint32, error) {
	return _MessageHarness.Contract.Nonce(&_MessageHarness.CallOpts, _message)
}

// Nonce is a free data retrieval call binding the contract method 0x58c8a98f.
//
// Solidity: function nonce(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCallerSession) Nonce(_message [29]byte) (uint32, error) {
	return _MessageHarness.Contract.Nonce(&_MessageHarness.CallOpts, _message)
}

// Origin is a free data retrieval call binding the contract method 0xa9f877ad.
//
// Solidity: function origin(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCaller) Origin(opts *bind.CallOpts, _message [29]byte) (uint32, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "origin", _message)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0xa9f877ad.
//
// Solidity: function origin(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessSession) Origin(_message [29]byte) (uint32, error) {
	return _MessageHarness.Contract.Origin(&_MessageHarness.CallOpts, _message)
}

// Origin is a free data retrieval call binding the contract method 0xa9f877ad.
//
// Solidity: function origin(bytes29 _message) pure returns(uint32)
func (_MessageHarness *MessageHarnessCallerSession) Origin(_message [29]byte) (uint32, error) {
	return _MessageHarness.Contract.Origin(&_MessageHarness.CallOpts, _message)
}

// Recipient is a free data retrieval call binding the contract method 0x418fad6d.
//
// Solidity: function recipient(bytes29 _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) Recipient(opts *bind.CallOpts, _message [29]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "recipient", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x418fad6d.
//
// Solidity: function recipient(bytes29 _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) Recipient(_message [29]byte) ([32]byte, error) {
	return _MessageHarness.Contract.Recipient(&_MessageHarness.CallOpts, _message)
}

// Recipient is a free data retrieval call binding the contract method 0x418fad6d.
//
// Solidity: function recipient(bytes29 _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) Recipient(_message [29]byte) ([32]byte, error) {
	return _MessageHarness.Contract.Recipient(&_MessageHarness.CallOpts, _message)
}

// RecipientAddress is a free data retrieval call binding the contract method 0x6eddb7f3.
//
// Solidity: function recipientAddress(bytes29 _message) pure returns(address)
func (_MessageHarness *MessageHarnessCaller) RecipientAddress(opts *bind.CallOpts, _message [29]byte) (common.Address, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "recipientAddress", _message)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientAddress is a free data retrieval call binding the contract method 0x6eddb7f3.
//
// Solidity: function recipientAddress(bytes29 _message) pure returns(address)
func (_MessageHarness *MessageHarnessSession) RecipientAddress(_message [29]byte) (common.Address, error) {
	return _MessageHarness.Contract.RecipientAddress(&_MessageHarness.CallOpts, _message)
}

// RecipientAddress is a free data retrieval call binding the contract method 0x6eddb7f3.
//
// Solidity: function recipientAddress(bytes29 _message) pure returns(address)
func (_MessageHarness *MessageHarnessCallerSession) RecipientAddress(_message [29]byte) (common.Address, error) {
	return _MessageHarness.Contract.RecipientAddress(&_MessageHarness.CallOpts, _message)
}

// Sender is a free data retrieval call binding the contract method 0xa0d41e58.
//
// Solidity: function sender(bytes29 _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) Sender(opts *bind.CallOpts, _message [29]byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "sender", _message)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0xa0d41e58.
//
// Solidity: function sender(bytes29 _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) Sender(_message [29]byte) ([32]byte, error) {
	return _MessageHarness.Contract.Sender(&_MessageHarness.CallOpts, _message)
}

// Sender is a free data retrieval call binding the contract method 0xa0d41e58.
//
// Solidity: function sender(bytes29 _message) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) Sender(_message [29]byte) ([32]byte, error) {
	return _MessageHarness.Contract.Sender(&_MessageHarness.CallOpts, _message)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b435ee8eb901cd08ee0eef07ea2ce68002a284d5cab762a1fe3f63f3197a2a3c64736f6c634300080d0033",
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
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201d9374fbaa14e68c88093df0b9080adf40d9d941bf9fcd500700c609d7b7d80a64736f6c634300080d0033",
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
	Bin: "0x60c9610038600b82828239805160001a607314602b57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b605e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff000000909116815260200160405180910390f3fea26469706673582212207b1e41d39790ec9465242265ca8fa171afe4b4ca83fec3d4ba343a161b5707c064736f6c634300080d0033",
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
