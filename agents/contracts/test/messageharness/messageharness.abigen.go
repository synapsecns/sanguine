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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200ce3cc281e49de8f86cc7c70511c07fbf6b6ca7dbd085f4e2f90d8f1fcf312a564736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c661fed1aaa7fa26551d9c3952c12415e326161b084816b06d16f756222b651f64736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"body\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"castToMessage\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint96\",\"name\":\"_notaryTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_broadcasterTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_proverTip\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"_executorTip\",\"type\":\"uint96\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"formatMessage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"header\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"isMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_tips\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_messageBody\",\"type\":\"bytes\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"messageVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"messageVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"tips\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"27bcefdf": "body(uint40,bytes)",
		"81038946": "castToMessage(uint40,bytes)",
		"0b9255e9": "formatMessage(bytes,bytes,bytes)",
		"7c517825": "formatMessage(uint32,bytes32,uint32,uint32,bytes32,uint32,bytes,bytes)",
		"81d030ef": "formatMessage(uint32,bytes32,uint32,uint32,bytes32,uint32,uint96,uint96,uint96,uint96,bytes)",
		"40e0cd87": "header(uint40,bytes)",
		"f9893ddd": "isMessage(bytes)",
		"02e5f230": "messageHash(bytes,bytes,bytes)",
		"52617f3c": "messageVersion()",
		"36861181": "messageVersion(uint40,bytes)",
		"1515f90c": "offsetHeader()",
		"0c096e8d": "offsetVersion()",
		"cd993304": "tips(uint40,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061192b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c806340e0cd871161008c578063810389461161006657806381038946146101ad57806381d030ef146101c0578063cd993304146101d3578063f9893ddd146101e657600080fd5b806340e0cd871461018057806352617f3c146101935780637c5178251461019a57600080fd5b80631515f90c116100bd5780631515f90c1461013157806327bcefdf14610139578063368611811461015a57600080fd5b806302e5f230146100e45780630b9255e91461010a5780630c096e8d1461012a575b600080fd5b6100f76100f23660046113dd565b610209565b6040519081526020015b60405180910390f35b61011d6101183660046113dd565b610220565b60405161010191906114d3565b60006100f7565b6100f761022d565b61014c6101473660046114e6565b610243565b604051610101929190611540565b61016d6101683660046114e6565b61028c565b60405161ffff9091168152602001610101565b61014c61018e3660046114e6565b6102af565b600161016d565b61011d6101a8366004611579565b6102cc565b61014c6101bb3660046114e6565b6102eb565b61011d6101ce36600461164f565b6102fa565b61014c6101e13660046114e6565b6103d1565b6101f96101f4366004611728565b6103ee565b6040519015158152602001610101565b6000610216848484610407565b90505b9392505050565b6060610216848484610424565b600061023b600360026117bb565b60ff16919050565b6000606081610260610255858761045b565b62ffffff191661047f565b905061027162ffffff1982166104e3565b61028062ffffff198316610507565b92509250509250929050565b60006102a661029b838561045b565b62ffffff191661055a565b90505b92915050565b60006060816102606102c1858761045b565b62ffffff1916610586565b60606102de89898989898989896105d0565b9998505050505050505050565b60006060600061026084610671565b6060600061030a87878787610682565b905060006103b28e8e8e8e8e8e604080517e01000000000000000000000000000000000000000000000000000000000000602082015260e097881b7fffffffff000000000000000000000000000000000000000000000000000000009081166022830152602682019790975294871b8616604686015292861b8516604a850152604e84019190915290931b909116606e82015281516052818303018152607290910190915290565b90506103bf818386610220565b9e9d5050505050505050505050505050565b60006060816102606103e3858761045b565b62ffffff1916610714565b60006102a96103fc83610671565b62ffffff1916610773565b6000610414848484610424565b8051906020012090509392505050565b82518251604051606092610443926001928890889088906020016117de565b60405160208183030381529060405290509392505050565b81516000906020840161047664ffffffffff8516828461086b565b95945050505050565b60008161049762ffffff1982166403010000006108b2565b506102196104a68460026109d5565b6104b18560016109d5565b6104bd600360026117bb565b6104ca919060ff16611870565b6104d49190611870565b62ffffff198516906000610a07565b60008060606104f3816018611870565b6104fd9190611870565b9290921c92915050565b60606000806105248460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060405191508192506105498483602001610a45565b508181016020016040529052919050565b60008161057262ffffff1982166403010000006108b2565b5061021962ffffff19841660006002610c2c565b60008161059e62ffffff1982166403010000006108b2565b506102196105ae600360026117bb565b60ff166105bc8560016109d5565b62ffffff1986169190640301010000610c5c565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e08b811b82166022840152602683018b905289811b8216604684015288811b8216604a840152604e830188905286901b16606e8201528151808203605201815260729091019091526060906102de908484610424565b60006102a98264030100000061045b565b6040517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffffffffffffffffffff000000000000000000000000000000000000000060a086811b8216602284015285811b8216602e84015284811b8216603a84015283901b16604682015260609060520160405160208183030381529060405290505b949350505050565b60008161072c62ffffff1982166403010000006108b2565b5061021961073b8460016109d5565b610747600360026117bb565b610754919060ff16611870565b61075f8560026109d5565b62ffffff1986169190640301020000610c5c565b6000601882901c6bffffffffffffffffffffffff16610794600360026117bb565b60ff168110156107a75750600092915050565b60016107b28461055a565b61ffff16146107c45750600092915050565b60006107d18460016109d5565b905060006107e08560026109d5565b90508281836107f1600360026117bb565b6107fe919060ff16611870565b6108089190611870565b111561081957506000949350505050565b61083061082586610586565b62ffffff1916610cd1565b1580610850575061084e61084386610714565b62ffffff1916610d18565b155b1561086057506000949350505050565b506001949350505050565b6000806108788385611870565b9050604051811115610888575060005b8060000361089d5762ffffff19915050610219565b5050606092831b9190911790911b1760181b90565b60006108be8383610d5f565b6109ce5760006108dc6108d0856104e3565b64ffffffffff16610d81565b91505060006108f18464ffffffffff16610d81565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c591906114d3565b60405180910390fd5b5090919050565b60006102a660028360038111156109ee576109ee61175d565b6109f89190611883565b62ffffff198516906002610c2c565b6000610216848485610a278860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610a3f919061189a565b85610c5c565b600062ffffff1980841603610ab6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e74657220646572656600000000000060448201526064016109c5565b610abf83610e6b565b610b25576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e74657220646572656600000060448201526064016109c5565b6000610b3f8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610b5a85610ea7565b6bffffffffffffffffffffffff169050600080604051915085821115610b805760206060fd5b8386858560045afa905080610bf1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f662067617300000000000000000000000060448201526064016109c5565b610c21610bfd886104e3565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b6000610c398260206118ad565b610c449060086117bb565b60ff16610c52858585610ece565b901c949350505050565b600080610c6886610ea7565b6bffffffffffffffffffffffff169050610c818661107c565b84610c8c8784611870565b610c969190611870565b1115610ca95762ffffff1991505061070c565b610cb38582611870565b9050610cc78364ffffffffff16828661086b565b9695505050505050565b6000601882901c6bffffffffffffffffffffffff166002811015610cf85750600092915050565b6001610d03846110b5565b61ffff16148015610219575060521492915050565b6000601882901c6bffffffffffffffffffffffff166002811015610d3f5750600092915050565b6001610d4a846110cd565b61ffff16148015610219575060321492915050565b60008164ffffffffff16610d72846104e3565b64ffffffffff16149392505050565b600080601f5b600f8160ff161115610df4576000610da08260086117bb565b60ff1685901c9050610db1816110e5565b61ffff16841793508160ff16601014610dcc57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610d87565b50600f5b60ff8160ff161015610e65576000610e118260086117bb565b60ff1685901c9050610e22816110e5565b61ffff16831792508160ff16600014610e3d57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610df8565b50915091565b6000610e76826104e3565b64ffffffffff1664ffffffffff03610e9057506000919050565b6000610e9b8361107c565b60405110199392505050565b600080610eb660606018611870565b9290921c6bffffffffffffffffffffffff1692915050565b60008160ff16600003610ee357506000610219565b610efb8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610f1660ff841685611870565b1115610f9957610f66610f2885610ea7565b6bffffffffffffffffffffffff16610f4e8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16611117565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109c591906114d3565b60208260ff161115611007576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e2033322062797465730000000000000060448201526064016109c5565b60088202600061101686610ea7565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b60006110968260181c6bffffffffffffffffffffffff1690565b61109f83610ea7565b016bffffffffffffffffffffffff169050919050565b60008161057262ffffff1982166403010100006108b2565b60008161057262ffffff1982166403010200006108b2565b60006110f760048360ff16901c6112a7565b60ff1661ffff919091161760081b61110e826112a7565b60ff1617919050565b6060600061112486610d81565b915050600061113286610d81565b915050600061114086610d81565b915050600061114e86610d81565b604080517f54797065644d656d566965772f696e646578202d204f76657272616e2074686560208201527f20766965772e20536c6963652069732061742030780000000000000000000000818301527fffffffffffff000000000000000000000000000000000000000000000000000060d098891b811660558301527f2077697468206c656e6774682030780000000000000000000000000000000000605b830181905297891b8116606a8301527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060708301527f7800000000000000000000000000000000000000000000000000000000000000609083015295881b861660918201526097810196909652951b90921660a684015250507f2e0000000000000000000000000000000000000000000000000000000000000060ac8201528151808203608d01815260ad90910190915295945050505050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f841691829081106112f4576112f46118c6565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261134357600080fd5b813567ffffffffffffffff8082111561135e5761135e611303565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156113a4576113a4611303565b816040528381528660208588010111156113bd57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156113f257600080fd5b833567ffffffffffffffff8082111561140a57600080fd5b61141687838801611332565b9450602086013591508082111561142c57600080fd5b61143887838801611332565b9350604086013591508082111561144e57600080fd5b5061145b86828701611332565b9150509250925092565b60005b83811015611480578181015183820152602001611468565b50506000910152565b600081518084526114a1816020860160208601611465565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006102a66020830184611489565b600080604083850312156114f957600080fd5b823564ffffffffff8116811461150e57600080fd5b9150602083013567ffffffffffffffff81111561152a57600080fd5b61153685828601611332565b9150509250929050565b64ffffffffff831681526040602082015260006102166040830184611489565b803563ffffffff8116811461157457600080fd5b919050565b600080600080600080600080610100898b03121561159657600080fd5b61159f89611560565b9750602089013596506115b460408a01611560565b95506115c260608a01611560565b9450608089013593506115d760a08a01611560565b925060c089013567ffffffffffffffff808211156115f457600080fd5b6116008c838d01611332565b935060e08b013591508082111561161657600080fd5b506116238b828c01611332565b9150509295985092959890939650565b80356bffffffffffffffffffffffff8116811461157457600080fd5b60008060008060008060008060008060006101608c8e03121561167157600080fd5b61167a8c611560565b9a5060208c0135995061168f60408d01611560565b985061169d60608d01611560565b975060808c013596506116b260a08d01611560565b95506116c060c08d01611633565b94506116ce60e08d01611633565b93506116dd6101008d01611633565b92506116ec6101208d01611633565b91506101408c013567ffffffffffffffff81111561170957600080fd5b6117158e828f01611332565b9150509295989b509295989b9093969950565b60006020828403121561173a57600080fd5b813567ffffffffffffffff81111561175157600080fd5b61070c84828501611332565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff81811683821602908116908181146117d7576117d761178c565b5092915050565b60007fffff000000000000000000000000000000000000000000000000000000000000808960f01b168352808860f01b166002840152808760f01b166004840152508451611833816006850160208901611465565b84519083019061184a816006840160208901611465565b8451910190611860816006840160208801611465565b0160060198975050505050505050565b808201808211156102a9576102a961178c565b80820281158282048414176102a9576102a961178c565b818103818111156102a9576102a961178c565b60ff82811682821603908111156102a9576102a961178c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212205fb98c6eb0036dc7932040bf9d3171ea0c1527c7d1a18ef4801848a362554fd764736f6c63430008110033",
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

// Body is a free data retrieval call binding the contract method 0x27bcefdf.
//
// Solidity: function body(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCaller) Body(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "body", _type, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// Body is a free data retrieval call binding the contract method 0x27bcefdf.
//
// Solidity: function body(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessSession) Body(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.Body(&_MessageHarness.CallOpts, _type, _payload)
}

// Body is a free data retrieval call binding the contract method 0x27bcefdf.
//
// Solidity: function body(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCallerSession) Body(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.Body(&_MessageHarness.CallOpts, _type, _payload)
}

// CastToMessage is a free data retrieval call binding the contract method 0x81038946.
//
// Solidity: function castToMessage(uint40 , bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCaller) CastToMessage(opts *bind.CallOpts, arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "castToMessage", arg0, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// CastToMessage is a free data retrieval call binding the contract method 0x81038946.
//
// Solidity: function castToMessage(uint40 , bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessSession) CastToMessage(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.CastToMessage(&_MessageHarness.CallOpts, arg0, _payload)
}

// CastToMessage is a free data retrieval call binding the contract method 0x81038946.
//
// Solidity: function castToMessage(uint40 , bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCallerSession) CastToMessage(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.CastToMessage(&_MessageHarness.CallOpts, arg0, _payload)
}

// FormatMessage is a free data retrieval call binding the contract method 0x0b9255e9.
//
// Solidity: function formatMessage(bytes _header, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCaller) FormatMessage(opts *bind.CallOpts, _header []byte, _tips []byte, _messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "formatMessage", _header, _tips, _messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage is a free data retrieval call binding the contract method 0x0b9255e9.
//
// Solidity: function formatMessage(bytes _header, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessSession) FormatMessage(_header []byte, _tips []byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _header, _tips, _messageBody)
}

// FormatMessage is a free data retrieval call binding the contract method 0x0b9255e9.
//
// Solidity: function formatMessage(bytes _header, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) FormatMessage(_header []byte, _tips []byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage(&_MessageHarness.CallOpts, _header, _tips, _messageBody)
}

// FormatMessage0 is a free data retrieval call binding the contract method 0x7c517825.
//
// Solidity: function formatMessage(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCaller) FormatMessage0(opts *bind.CallOpts, _origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "formatMessage0", _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage0 is a free data retrieval call binding the contract method 0x7c517825.
//
// Solidity: function formatMessage(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessSession) FormatMessage0(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage0(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// FormatMessage0 is a free data retrieval call binding the contract method 0x7c517825.
//
// Solidity: function formatMessage(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, bytes _tips, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) FormatMessage0(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _tips []byte, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage0(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _tips, _messageBody)
}

// FormatMessage1 is a free data retrieval call binding the contract method 0x81d030ef.
//
// Solidity: function formatMessage(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, uint96 _notaryTip, uint96 _broadcasterTip, uint96 _proverTip, uint96 _executorTip, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCaller) FormatMessage1(opts *bind.CallOpts, _origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _notaryTip *big.Int, _broadcasterTip *big.Int, _proverTip *big.Int, _executorTip *big.Int, _messageBody []byte) ([]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "formatMessage1", _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _notaryTip, _broadcasterTip, _proverTip, _executorTip, _messageBody)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatMessage1 is a free data retrieval call binding the contract method 0x81d030ef.
//
// Solidity: function formatMessage(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, uint96 _notaryTip, uint96 _broadcasterTip, uint96 _proverTip, uint96 _executorTip, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessSession) FormatMessage1(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _notaryTip *big.Int, _broadcasterTip *big.Int, _proverTip *big.Int, _executorTip *big.Int, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage1(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _notaryTip, _broadcasterTip, _proverTip, _executorTip, _messageBody)
}

// FormatMessage1 is a free data retrieval call binding the contract method 0x81d030ef.
//
// Solidity: function formatMessage(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds, uint96 _notaryTip, uint96 _broadcasterTip, uint96 _proverTip, uint96 _executorTip, bytes _messageBody) pure returns(bytes)
func (_MessageHarness *MessageHarnessCallerSession) FormatMessage1(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32, _notaryTip *big.Int, _broadcasterTip *big.Int, _proverTip *big.Int, _executorTip *big.Int, _messageBody []byte) ([]byte, error) {
	return _MessageHarness.Contract.FormatMessage1(&_MessageHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds, _notaryTip, _broadcasterTip, _proverTip, _executorTip, _messageBody)
}

// Header is a free data retrieval call binding the contract method 0x40e0cd87.
//
// Solidity: function header(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCaller) Header(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "header", _type, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// Header is a free data retrieval call binding the contract method 0x40e0cd87.
//
// Solidity: function header(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessSession) Header(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.Header(&_MessageHarness.CallOpts, _type, _payload)
}

// Header is a free data retrieval call binding the contract method 0x40e0cd87.
//
// Solidity: function header(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCallerSession) Header(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.Header(&_MessageHarness.CallOpts, _type, _payload)
}

// IsMessage is a free data retrieval call binding the contract method 0xf9893ddd.
//
// Solidity: function isMessage(bytes _payload) pure returns(bool)
func (_MessageHarness *MessageHarnessCaller) IsMessage(opts *bind.CallOpts, _payload []byte) (bool, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "isMessage", _payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMessage is a free data retrieval call binding the contract method 0xf9893ddd.
//
// Solidity: function isMessage(bytes _payload) pure returns(bool)
func (_MessageHarness *MessageHarnessSession) IsMessage(_payload []byte) (bool, error) {
	return _MessageHarness.Contract.IsMessage(&_MessageHarness.CallOpts, _payload)
}

// IsMessage is a free data retrieval call binding the contract method 0xf9893ddd.
//
// Solidity: function isMessage(bytes _payload) pure returns(bool)
func (_MessageHarness *MessageHarnessCallerSession) IsMessage(_payload []byte) (bool, error) {
	return _MessageHarness.Contract.IsMessage(&_MessageHarness.CallOpts, _payload)
}

// MessageHash is a free data retrieval call binding the contract method 0x02e5f230.
//
// Solidity: function messageHash(bytes _header, bytes _tips, bytes _messageBody) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCaller) MessageHash(opts *bind.CallOpts, _header []byte, _tips []byte, _messageBody []byte) ([32]byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "messageHash", _header, _tips, _messageBody)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0x02e5f230.
//
// Solidity: function messageHash(bytes _header, bytes _tips, bytes _messageBody) pure returns(bytes32)
func (_MessageHarness *MessageHarnessSession) MessageHash(_header []byte, _tips []byte, _messageBody []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _header, _tips, _messageBody)
}

// MessageHash is a free data retrieval call binding the contract method 0x02e5f230.
//
// Solidity: function messageHash(bytes _header, bytes _tips, bytes _messageBody) pure returns(bytes32)
func (_MessageHarness *MessageHarnessCallerSession) MessageHash(_header []byte, _tips []byte, _messageBody []byte) ([32]byte, error) {
	return _MessageHarness.Contract.MessageHash(&_MessageHarness.CallOpts, _header, _tips, _messageBody)
}

// MessageVersion is a free data retrieval call binding the contract method 0x36861181.
//
// Solidity: function messageVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_MessageHarness *MessageHarnessCaller) MessageVersion(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint16, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "messageVersion", _type, _payload)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MessageVersion is a free data retrieval call binding the contract method 0x36861181.
//
// Solidity: function messageVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_MessageHarness *MessageHarnessSession) MessageVersion(_type *big.Int, _payload []byte) (uint16, error) {
	return _MessageHarness.Contract.MessageVersion(&_MessageHarness.CallOpts, _type, _payload)
}

// MessageVersion is a free data retrieval call binding the contract method 0x36861181.
//
// Solidity: function messageVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_MessageHarness *MessageHarnessCallerSession) MessageVersion(_type *big.Int, _payload []byte) (uint16, error) {
	return _MessageHarness.Contract.MessageVersion(&_MessageHarness.CallOpts, _type, _payload)
}

// MessageVersion0 is a free data retrieval call binding the contract method 0x52617f3c.
//
// Solidity: function messageVersion() pure returns(uint16)
func (_MessageHarness *MessageHarnessCaller) MessageVersion0(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "messageVersion0")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MessageVersion0 is a free data retrieval call binding the contract method 0x52617f3c.
//
// Solidity: function messageVersion() pure returns(uint16)
func (_MessageHarness *MessageHarnessSession) MessageVersion0() (uint16, error) {
	return _MessageHarness.Contract.MessageVersion0(&_MessageHarness.CallOpts)
}

// MessageVersion0 is a free data retrieval call binding the contract method 0x52617f3c.
//
// Solidity: function messageVersion() pure returns(uint16)
func (_MessageHarness *MessageHarnessCallerSession) MessageVersion0() (uint16, error) {
	return _MessageHarness.Contract.MessageVersion0(&_MessageHarness.CallOpts)
}

// OffsetHeader is a free data retrieval call binding the contract method 0x1515f90c.
//
// Solidity: function offsetHeader() pure returns(uint256)
func (_MessageHarness *MessageHarnessCaller) OffsetHeader(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "offsetHeader")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetHeader is a free data retrieval call binding the contract method 0x1515f90c.
//
// Solidity: function offsetHeader() pure returns(uint256)
func (_MessageHarness *MessageHarnessSession) OffsetHeader() (*big.Int, error) {
	return _MessageHarness.Contract.OffsetHeader(&_MessageHarness.CallOpts)
}

// OffsetHeader is a free data retrieval call binding the contract method 0x1515f90c.
//
// Solidity: function offsetHeader() pure returns(uint256)
func (_MessageHarness *MessageHarnessCallerSession) OffsetHeader() (*big.Int, error) {
	return _MessageHarness.Contract.OffsetHeader(&_MessageHarness.CallOpts)
}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_MessageHarness *MessageHarnessCaller) OffsetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "offsetVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_MessageHarness *MessageHarnessSession) OffsetVersion() (*big.Int, error) {
	return _MessageHarness.Contract.OffsetVersion(&_MessageHarness.CallOpts)
}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_MessageHarness *MessageHarnessCallerSession) OffsetVersion() (*big.Int, error) {
	return _MessageHarness.Contract.OffsetVersion(&_MessageHarness.CallOpts)
}

// Tips is a free data retrieval call binding the contract method 0xcd993304.
//
// Solidity: function tips(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCaller) Tips(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _MessageHarness.contract.Call(opts, &out, "tips", _type, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// Tips is a free data retrieval call binding the contract method 0xcd993304.
//
// Solidity: function tips(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessSession) Tips(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.Tips(&_MessageHarness.CallOpts, _type, _payload)
}

// Tips is a free data retrieval call binding the contract method 0xcd993304.
//
// Solidity: function tips(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_MessageHarness *MessageHarnessCallerSession) Tips(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _MessageHarness.Contract.Tips(&_MessageHarness.CallOpts, _type, _payload)
}

// SynapseTypesMetaData contains all meta data concerning the SynapseTypes contract.
var SynapseTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bc97f2d0f4db4d01afcc612a21dcb4e3b76163651164aa7ef07c20d9b199b5b164736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208b3a879112ca33be589153ad9d33290d52c59a4edc0818b7a043c858d92aa89864736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e677e00d8056f3cff61c7681fe15a66385f925305df3b72135b74bd7d09ca5c464736f6c63430008110033",
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
	ABI: "[{\"inputs\":[],\"name\":\"BITS_EMPTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BITS_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LOW_96_BITS_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_LOC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SHIFT_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"97b8ad4a": "BITS_EMPTY()",
		"eb740628": "BITS_LEN()",
		"fb734584": "BITS_LOC()",
		"10153fce": "BITS_TYPE()",
		"b602d173": "LOW_96_BITS_MASK()",
		"f26be3fc": "NULL()",
		"1136e7ea": "SHIFT_LEN()",
		"1bfe17ce": "SHIFT_LOC()",
		"13090c5a": "SHIFT_TYPE()",
	},
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea26469706673582212205a371e560fa172f173a3be10938201a9dc81497059ea8ad62c2b9a8e423c65fd64736f6c63430008110033",
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

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSEMPTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_EMPTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSEMPTY is a free data retrieval call binding the contract method 0x97b8ad4a.
//
// Solidity: function BITS_EMPTY() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSEMPTY() (*big.Int, error) {
	return _TypedMemView.Contract.BITSEMPTY(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLEN is a free data retrieval call binding the contract method 0xeb740628.
//
// Solidity: function BITS_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLEN() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLEN(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSLOC is a free data retrieval call binding the contract method 0xfb734584.
//
// Solidity: function BITS_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSLOC() (*big.Int, error) {
	return _TypedMemView.Contract.BITSLOC(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) BITSTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "BITS_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// BITSTYPE is a free data retrieval call binding the contract method 0x10153fce.
//
// Solidity: function BITS_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) BITSTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.BITSTYPE(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) LOW96BITSMASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "LOW_96_BITS_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
}

// LOW96BITSMASK is a free data retrieval call binding the contract method 0xb602d173.
//
// Solidity: function LOW_96_BITS_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) LOW96BITSMASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW96BITSMASK(&_TypedMemView.CallOpts)
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

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLEN is a free data retrieval call binding the contract method 0x1136e7ea.
//
// Solidity: function SHIFT_LEN() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLEN() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLEN(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTLOC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_LOC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTLOC is a free data retrieval call binding the contract method 0x1bfe17ce.
//
// Solidity: function SHIFT_LOC() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTLOC() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTLOC(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) SHIFTTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "SHIFT_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}

// SHIFTTYPE is a free data retrieval call binding the contract method 0x13090c5a.
//
// Solidity: function SHIFT_TYPE() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) SHIFTTYPE() (*big.Int, error) {
	return _TypedMemView.Contract.SHIFTTYPE(&_TypedMemView.CallOpts)
}
