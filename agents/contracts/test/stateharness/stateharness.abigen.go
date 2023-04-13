// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stateharness

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

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122056892eb3eed8af38c911c7c780bb32db581e4c7f95e584148caee9317750d77764736f6c63430008110033",
}

// ByteStringABI is the input ABI used to generate the binding from.
// Deprecated: Use ByteStringMetaData.ABI instead.
var ByteStringABI = ByteStringMetaData.ABI

// ByteStringBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ByteStringMetaData.Bin instead.
var ByteStringBin = ByteStringMetaData.Bin

// DeployByteString deploys a new Ethereum contract, binding an instance of ByteString to it.
func DeployByteString(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ByteString, error) {
	parsed, err := ByteStringMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ByteStringBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// ByteString is an auto generated Go binding around an Ethereum contract.
type ByteString struct {
	ByteStringCaller     // Read-only binding to the contract
	ByteStringTransactor // Write-only binding to the contract
	ByteStringFilterer   // Log filterer for contract events
}

// ByteStringCaller is an auto generated read-only Go binding around an Ethereum contract.
type ByteStringCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ByteStringTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ByteStringFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ByteStringSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ByteStringSession struct {
	Contract     *ByteString       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ByteStringCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ByteStringCallerSession struct {
	Contract *ByteStringCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ByteStringTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ByteStringTransactorSession struct {
	Contract     *ByteStringTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ByteStringRaw is an auto generated low-level Go binding around an Ethereum contract.
type ByteStringRaw struct {
	Contract *ByteString // Generic contract binding to access the raw methods on
}

// ByteStringCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ByteStringCallerRaw struct {
	Contract *ByteStringCaller // Generic read-only contract binding to access the raw methods on
}

// ByteStringTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ByteStringTransactorRaw struct {
	Contract *ByteStringTransactor // Generic write-only contract binding to access the raw methods on
}

// NewByteString creates a new instance of ByteString, bound to a specific deployed contract.
func NewByteString(address common.Address, backend bind.ContractBackend) (*ByteString, error) {
	contract, err := bindByteString(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ByteString{ByteStringCaller: ByteStringCaller{contract: contract}, ByteStringTransactor: ByteStringTransactor{contract: contract}, ByteStringFilterer: ByteStringFilterer{contract: contract}}, nil
}

// NewByteStringCaller creates a new read-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringCaller(address common.Address, caller bind.ContractCaller) (*ByteStringCaller, error) {
	contract, err := bindByteString(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringCaller{contract: contract}, nil
}

// NewByteStringTransactor creates a new write-only instance of ByteString, bound to a specific deployed contract.
func NewByteStringTransactor(address common.Address, transactor bind.ContractTransactor) (*ByteStringTransactor, error) {
	contract, err := bindByteString(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ByteStringTransactor{contract: contract}, nil
}

// NewByteStringFilterer creates a new log filterer instance of ByteString, bound to a specific deployed contract.
func NewByteStringFilterer(address common.Address, filterer bind.ContractFilterer) (*ByteStringFilterer, error) {
	contract, err := bindByteString(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ByteStringFilterer{contract: contract}, nil
}

// bindByteString binds a generic wrapper to an already deployed contract.
func bindByteString(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ByteStringABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.ByteStringCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.ByteStringTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ByteString *ByteStringCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ByteString.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ByteString *ByteStringTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ByteString *ByteStringTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ByteString.Contract.contract.Transact(opts, method, params...)
}

// StateHarnessMetaData contains all meta data concerning the StateHarness contract.
var StateHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"equals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"formatState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"}],\"name\":\"leftLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"rightLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"subLeafs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e948e600": "blockNumber(bytes)",
		"5fed0261": "castToState(bytes)",
		"137e618a": "equals(bytes,bytes)",
		"365b4b67": "formatState(bytes32,uint32,uint32,uint40,uint40)",
		"aae6d884": "isState(bytes)",
		"d7a7a72c": "leaf(bytes)",
		"edaa471d": "leftLeaf(bytes32,uint32)",
		"4e765004": "nonce(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"503d0bed": "rightLeaf(uint32,uint40,uint40)",
		"c2e9e208": "root(bytes)",
		"9aaa1826": "subLeafs(bytes)",
		"1c9aa222": "timestamp(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611423806100206000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80639aaa18261161008c578063cb3eb0e111610066578063cb3eb0e1146101ff578063d7a7a72c14610212578063e948e60014610225578063edaa471d1461023857600080fd5b80639aaa1826146101b1578063aae6d884146101d9578063c2e9e208146101ec57600080fd5b80634e765004116100bd5780634e76500414610155578063503d0bed1461017d5780635fed02611461019e57600080fd5b8063137e618a146100e45780631c9aa2221461010c578063365b4b6714610135575b600080fd5b6100f76100f2366004610fe5565b61024b565b60405190151581526020015b60405180910390f35b61011f61011a366004611049565b610277565b60405164ffffffffff9091168152602001610103565b6101486101433660046110ac565b610290565b604051610103919061116e565b610168610163366004611049565b610326565b60405163ffffffff9091168152602001610103565b61019061018b366004611181565b61033f565b604051908152602001610103565b6101486101ac366004611049565b6103ce565b6101c46101bf366004611049565b6103ec565b60408051928352602083019190915201610103565b6100f76101e7366004611049565b61040f565b6101906101fa366004611049565b61043b565b61016861020d366004611049565b610454565b610190610220366004611049565b61046d565b61011f610233366004611049565b610486565b6101906102463660046111c4565b61049f565b600061026e610259836104ab565b610262856104ab565b62ffffff1916906104be565b90505b92915050565b6000610271610285836104ab565b62ffffff19166104ed565b60408051602081018790527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168385015286901b1660448201527fffffffffff00000000000000000000000000000000000000000000000000000060d885811b8216604884015284901b16604d8201528151808203603201815260529091019091526060905b9695505050505050565b6000610271610334836104ab565b62ffffff1916610503565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b166020808301919091527fffffffffff00000000000000000000000000000000000000000000000000000060d886811b8216602485015285901b1660298301528251808303600e018152602e90920190925280519101206000905b90505b9392505050565b606060006103db836104ab565b90506103c762ffffff198216610519565b6000806104066103fb846104ab565b62ffffff191661056c565b91509150915091565b600061027161041e83836105a5565b62ffffff191660181c6bffffffffffffffffffffffff1660321490565b6000610271610449836104ab565b62ffffff19166105c9565b6000610271610462836104ab565b62ffffff19166105de565b600061027161047b836104ab565b62ffffff19166105f4565b6000610271610494836104ab565b62ffffff1916610638565b600061026e838361064e565b60006102716104b9836106af565b6106bb565b60006104d662ffffff1983165b62ffffff1916610747565b6104e562ffffff1985166104cb565b149392505050565b600062ffffff1982166103c781602d6005610795565b600062ffffff1982166103c78160246004610795565b60606000806105368460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061055b84836020016107c5565b508181016020016040529052919050565b60008062ffffff1983166105856104cb826024856109ac565b925061059d6104cb62ffffff198316602460006109bb565b915050915091565b8151600090602084016105c064ffffffffff851682846109f9565b95945050505050565b600062ffffff1982166103c781836020610a40565b600062ffffff1982166103c78160206004610795565b6000808061060762ffffff19851661056c565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b600062ffffff1982166103c78160286005610795565b6000828260405160200161069192919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b600061027182826105a5565b60006106d88260181c6bffffffffffffffffffffffff1660321490565b610743576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f74206120737461746500000000000000000000000000000000000000000060448201526064015b60405180910390fd5b5090565b60008061075383610bee565b6bffffffffffffffffffffffff169050600061077d8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b60006107a282602061121f565b6107ad906008611238565b60ff166107bb858585610a40565b901c949350505050565b600062ffffff1980841603610836576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e746572206465726566000000000000604482015260640161073a565b61083f83610c15565b6108a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e746572206465726566000000604482015260640161073a565b60006108bf8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006108da85610bee565b6bffffffffffffffffffffffff1690506000806040519150858211156109005760206060fd5b8386858560045afa905080610971576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f6620676173000000000000000000000000604482015260640161073a565b6109a161097d88610c51565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b60006103c48460008585610c75565b60006103c48484856109db8860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166109f3919061125b565b85610c75565b600080610a06838561126e565b9050604051811115610a16575060005b80600003610a2b5762ffffff199150506103c7565b5050606092831b9190911790911b1760181b90565b60008160ff16600003610a55575060006103c7565b610a6d8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610a8860ff84168561126e565b1115610b0b57610ad8610a9a85610bee565b6bffffffffffffffffffffffff16610ac08660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16610cec565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073a919061116e565b60208260ff161115610b79576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e20333220627974657300000000000000604482015260640161073a565b600882026000610b8886610bee565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b600080610bfd6060601861126e565b9290921c6bffffffffffffffffffffffff1692915050565b6000610c2082610c51565b64ffffffffff1664ffffffffff03610c3a57506000919050565b6000610c4583610d5a565b60405110199392505050565b6000806060610c6181601861126e565b610c6b919061126e565b9290921c92915050565b600080610c8186610bee565b6bffffffffffffffffffffffff169050610c9a86610d5a565b84610ca5878461126e565b610caf919061126e565b1115610cc25762ffffff19915050610ce4565b610ccc858261126e565b9050610ce08364ffffffffff1682866109f9565b9150505b949350505050565b60606000610cf986610d93565b9150506000610d0786610d93565b9150506000610d1586610d93565b9150506000610d2386610d93565b91505083838383604051602001610d3d9493929190611281565b604051602081830303815290604052945050505050949350505050565b6000610d748260181c6bffffffffffffffffffffffff1690565b610d7d83610bee565b016bffffffffffffffffffffffff169050919050565b600080601f5b600f8160ff161115610e06576000610db2826008611238565b60ff1685901c9050610dc381610e7d565b61ffff16841793508160ff16601014610dde57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610d99565b50600f5b60ff8160ff161015610e77576000610e23826008611238565b60ff1685901c9050610e3481610e7d565b61ffff16831792508160ff16600014610e4f57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610e0a565b50915091565b6000610e8f60048360ff16901c610eaf565b60ff1661ffff919091161760081b610ea682610eaf565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110610efc57610efc6113be565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610f4b57600080fd5b813567ffffffffffffffff80821115610f6657610f66610f0b565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610fac57610fac610f0b565b81604052838152866020858801011115610fc557600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610ff857600080fd5b823567ffffffffffffffff8082111561101057600080fd5b61101c86838701610f3a565b9350602085013591508082111561103257600080fd5b5061103f85828601610f3a565b9150509250929050565b60006020828403121561105b57600080fd5b813567ffffffffffffffff81111561107257600080fd5b610ce484828501610f3a565b803563ffffffff8116811461109257600080fd5b919050565b803564ffffffffff8116811461109257600080fd5b600080600080600060a086880312156110c457600080fd5b853594506110d46020870161107e565b93506110e26040870161107e565b92506110f060608701611097565b91506110fe60808701611097565b90509295509295909350565b6000815180845260005b8181101561113057602081850181015186830182015201611114565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061026e602083018461110a565b60008060006060848603121561119657600080fd5b61119f8461107e565b92506111ad60208501611097565b91506111bb60408501611097565b90509250925092565b600080604083850312156111d757600080fd5b823591506111e76020840161107e565b90509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff8281168282160390811115610271576102716111f0565b60ff8181168382160290811690818114611254576112546111f0565b5092915050565b81810381811115610271576102716111f0565b80820180821115610271576102716111f0565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d820161031c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220b0dcdfb1cc872b2477bc8127678bd592ec57c2420fbcf6391467aa77b15a642e64736f6c63430008110033",
}

// StateHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use StateHarnessMetaData.ABI instead.
var StateHarnessABI = StateHarnessMetaData.ABI

// Deprecated: Use StateHarnessMetaData.Sigs instead.
// StateHarnessFuncSigs maps the 4-byte function signature to its string representation.
var StateHarnessFuncSigs = StateHarnessMetaData.Sigs

// StateHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateHarnessMetaData.Bin instead.
var StateHarnessBin = StateHarnessMetaData.Bin

// DeployStateHarness deploys a new Ethereum contract, binding an instance of StateHarness to it.
func DeployStateHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateHarness, error) {
	parsed, err := StateHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateHarness{StateHarnessCaller: StateHarnessCaller{contract: contract}, StateHarnessTransactor: StateHarnessTransactor{contract: contract}, StateHarnessFilterer: StateHarnessFilterer{contract: contract}}, nil
}

// StateHarness is an auto generated Go binding around an Ethereum contract.
type StateHarness struct {
	StateHarnessCaller     // Read-only binding to the contract
	StateHarnessTransactor // Write-only binding to the contract
	StateHarnessFilterer   // Log filterer for contract events
}

// StateHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateHarnessSession struct {
	Contract     *StateHarness     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateHarnessCallerSession struct {
	Contract *StateHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StateHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateHarnessTransactorSession struct {
	Contract     *StateHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StateHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateHarnessRaw struct {
	Contract *StateHarness // Generic contract binding to access the raw methods on
}

// StateHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateHarnessCallerRaw struct {
	Contract *StateHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// StateHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateHarnessTransactorRaw struct {
	Contract *StateHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateHarness creates a new instance of StateHarness, bound to a specific deployed contract.
func NewStateHarness(address common.Address, backend bind.ContractBackend) (*StateHarness, error) {
	contract, err := bindStateHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateHarness{StateHarnessCaller: StateHarnessCaller{contract: contract}, StateHarnessTransactor: StateHarnessTransactor{contract: contract}, StateHarnessFilterer: StateHarnessFilterer{contract: contract}}, nil
}

// NewStateHarnessCaller creates a new read-only instance of StateHarness, bound to a specific deployed contract.
func NewStateHarnessCaller(address common.Address, caller bind.ContractCaller) (*StateHarnessCaller, error) {
	contract, err := bindStateHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateHarnessCaller{contract: contract}, nil
}

// NewStateHarnessTransactor creates a new write-only instance of StateHarness, bound to a specific deployed contract.
func NewStateHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*StateHarnessTransactor, error) {
	contract, err := bindStateHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateHarnessTransactor{contract: contract}, nil
}

// NewStateHarnessFilterer creates a new log filterer instance of StateHarness, bound to a specific deployed contract.
func NewStateHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*StateHarnessFilterer, error) {
	contract, err := bindStateHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateHarnessFilterer{contract: contract}, nil
}

// bindStateHarness binds a generic wrapper to an already deployed contract.
func bindStateHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHarness *StateHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHarness.Contract.StateHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHarness *StateHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHarness.Contract.StateHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHarness *StateHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHarness.Contract.StateHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateHarness *StateHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateHarness *StateHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateHarness *StateHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateHarness.Contract.contract.Transact(opts, method, params...)
}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCaller) BlockNumber(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "blockNumber", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessSession) BlockNumber(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.BlockNumber(&_StateHarness.CallOpts, payload)
}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCallerSession) BlockNumber(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.BlockNumber(&_StateHarness.CallOpts, payload)
}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes payload) view returns(bytes)
func (_StateHarness *StateHarnessCaller) CastToState(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "castToState", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes payload) view returns(bytes)
func (_StateHarness *StateHarnessSession) CastToState(payload []byte) ([]byte, error) {
	return _StateHarness.Contract.CastToState(&_StateHarness.CallOpts, payload)
}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes payload) view returns(bytes)
func (_StateHarness *StateHarnessCallerSession) CastToState(payload []byte) ([]byte, error) {
	return _StateHarness.Contract.CastToState(&_StateHarness.CallOpts, payload)
}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_StateHarness *StateHarnessCaller) Equals(opts *bind.CallOpts, a []byte, b []byte) (bool, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "equals", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_StateHarness *StateHarnessSession) Equals(a []byte, b []byte) (bool, error) {
	return _StateHarness.Contract.Equals(&_StateHarness.CallOpts, a, b)
}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_StateHarness *StateHarnessCallerSession) Equals(a []byte, b []byte) (bool, error) {
	return _StateHarness.Contract.Equals(&_StateHarness.CallOpts, a, b)
}

// FormatState is a free data retrieval call binding the contract method 0x365b4b67.
//
// Solidity: function formatState(bytes32 root_, uint32 origin_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_StateHarness *StateHarnessCaller) FormatState(opts *bind.CallOpts, root_ [32]byte, origin_ uint32, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "formatState", root_, origin_, nonce_, blockNumber_, timestamp_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatState is a free data retrieval call binding the contract method 0x365b4b67.
//
// Solidity: function formatState(bytes32 root_, uint32 origin_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_StateHarness *StateHarnessSession) FormatState(root_ [32]byte, origin_ uint32, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	return _StateHarness.Contract.FormatState(&_StateHarness.CallOpts, root_, origin_, nonce_, blockNumber_, timestamp_)
}

// FormatState is a free data retrieval call binding the contract method 0x365b4b67.
//
// Solidity: function formatState(bytes32 root_, uint32 origin_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_StateHarness *StateHarnessCallerSession) FormatState(root_ [32]byte, origin_ uint32, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	return _StateHarness.Contract.FormatState(&_StateHarness.CallOpts, root_, origin_, nonce_, blockNumber_, timestamp_)
}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes payload) pure returns(bool)
func (_StateHarness *StateHarnessCaller) IsState(opts *bind.CallOpts, payload []byte) (bool, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "isState", payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes payload) pure returns(bool)
func (_StateHarness *StateHarnessSession) IsState(payload []byte) (bool, error) {
	return _StateHarness.Contract.IsState(&_StateHarness.CallOpts, payload)
}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes payload) pure returns(bool)
func (_StateHarness *StateHarnessCallerSession) IsState(payload []byte) (bool, error) {
	return _StateHarness.Contract.IsState(&_StateHarness.CallOpts, payload)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) Leaf(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "leaf", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) Leaf(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Leaf(&_StateHarness.CallOpts, payload)
}

// Leaf is a free data retrieval call binding the contract method 0xd7a7a72c.
//
// Solidity: function leaf(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) Leaf(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Leaf(&_StateHarness.CallOpts, payload)
}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 root_, uint32 origin_) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) LeftLeaf(opts *bind.CallOpts, root_ [32]byte, origin_ uint32) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "leftLeaf", root_, origin_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 root_, uint32 origin_) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) LeftLeaf(root_ [32]byte, origin_ uint32) ([32]byte, error) {
	return _StateHarness.Contract.LeftLeaf(&_StateHarness.CallOpts, root_, origin_)
}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 root_, uint32 origin_) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) LeftLeaf(root_ [32]byte, origin_ uint32) ([32]byte, error) {
	return _StateHarness.Contract.LeftLeaf(&_StateHarness.CallOpts, root_, origin_)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCaller) Nonce(opts *bind.CallOpts, payload []byte) (uint32, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "nonce", payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessSession) Nonce(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Nonce(&_StateHarness.CallOpts, payload)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCallerSession) Nonce(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Nonce(&_StateHarness.CallOpts, payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCaller) Origin(opts *bind.CallOpts, payload []byte) (uint32, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "origin", payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessSession) Origin(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Origin(&_StateHarness.CallOpts, payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_StateHarness *StateHarnessCallerSession) Origin(payload []byte) (uint32, error) {
	return _StateHarness.Contract.Origin(&_StateHarness.CallOpts, payload)
}

// RightLeaf is a free data retrieval call binding the contract method 0x503d0bed.
//
// Solidity: function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) RightLeaf(opts *bind.CallOpts, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "rightLeaf", nonce_, blockNumber_, timestamp_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RightLeaf is a free data retrieval call binding the contract method 0x503d0bed.
//
// Solidity: function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) RightLeaf(nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([32]byte, error) {
	return _StateHarness.Contract.RightLeaf(&_StateHarness.CallOpts, nonce_, blockNumber_, timestamp_)
}

// RightLeaf is a free data retrieval call binding the contract method 0x503d0bed.
//
// Solidity: function rightLeaf(uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) RightLeaf(nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([32]byte, error) {
	return _StateHarness.Contract.RightLeaf(&_StateHarness.CallOpts, nonce_, blockNumber_, timestamp_)
}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) Root(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "root", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) Root(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Root(&_StateHarness.CallOpts, payload)
}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) Root(payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Root(&_StateHarness.CallOpts, payload)
}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessCaller) SubLeafs(opts *bind.CallOpts, payload []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "subLeafs", payload)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessSession) SubLeafs(payload []byte) ([32]byte, [32]byte, error) {
	return _StateHarness.Contract.SubLeafs(&_StateHarness.CallOpts, payload)
}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessCallerSession) SubLeafs(payload []byte) ([32]byte, [32]byte, error) {
	return _StateHarness.Contract.SubLeafs(&_StateHarness.CallOpts, payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCaller) Timestamp(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "timestamp", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessSession) Timestamp(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.Timestamp(&_StateHarness.CallOpts, payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_StateHarness *StateHarnessCallerSession) Timestamp(payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.Timestamp(&_StateHarness.CallOpts, payload)
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d3c34b61cad8a0e6b6136e211b453ff5cab0a4c778ef7249099b5ce6ec0a225c64736f6c63430008110033",
}

// StateLibABI is the input ABI used to generate the binding from.
// Deprecated: Use StateLibMetaData.ABI instead.
var StateLibABI = StateLibMetaData.ABI

// StateLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateLibMetaData.Bin instead.
var StateLibBin = StateLibMetaData.Bin

// DeployStateLib deploys a new Ethereum contract, binding an instance of StateLib to it.
func DeployStateLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StateLib, error) {
	parsed, err := StateLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StateLib{StateLibCaller: StateLibCaller{contract: contract}, StateLibTransactor: StateLibTransactor{contract: contract}, StateLibFilterer: StateLibFilterer{contract: contract}}, nil
}

// StateLib is an auto generated Go binding around an Ethereum contract.
type StateLib struct {
	StateLibCaller     // Read-only binding to the contract
	StateLibTransactor // Write-only binding to the contract
	StateLibFilterer   // Log filterer for contract events
}

// StateLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateLibSession struct {
	Contract     *StateLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateLibCallerSession struct {
	Contract *StateLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StateLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateLibTransactorSession struct {
	Contract     *StateLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StateLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateLibRaw struct {
	Contract *StateLib // Generic contract binding to access the raw methods on
}

// StateLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateLibCallerRaw struct {
	Contract *StateLibCaller // Generic read-only contract binding to access the raw methods on
}

// StateLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateLibTransactorRaw struct {
	Contract *StateLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStateLib creates a new instance of StateLib, bound to a specific deployed contract.
func NewStateLib(address common.Address, backend bind.ContractBackend) (*StateLib, error) {
	contract, err := bindStateLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StateLib{StateLibCaller: StateLibCaller{contract: contract}, StateLibTransactor: StateLibTransactor{contract: contract}, StateLibFilterer: StateLibFilterer{contract: contract}}, nil
}

// NewStateLibCaller creates a new read-only instance of StateLib, bound to a specific deployed contract.
func NewStateLibCaller(address common.Address, caller bind.ContractCaller) (*StateLibCaller, error) {
	contract, err := bindStateLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateLibCaller{contract: contract}, nil
}

// NewStateLibTransactor creates a new write-only instance of StateLib, bound to a specific deployed contract.
func NewStateLibTransactor(address common.Address, transactor bind.ContractTransactor) (*StateLibTransactor, error) {
	contract, err := bindStateLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateLibTransactor{contract: contract}, nil
}

// NewStateLibFilterer creates a new log filterer instance of StateLib, bound to a specific deployed contract.
func NewStateLibFilterer(address common.Address, filterer bind.ContractFilterer) (*StateLibFilterer, error) {
	contract, err := bindStateLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateLibFilterer{contract: contract}, nil
}

// bindStateLib binds a generic wrapper to an already deployed contract.
func bindStateLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateLib *StateLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateLib.Contract.StateLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateLib *StateLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateLib.Contract.StateLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateLib *StateLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateLib.Contract.StateLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StateLib *StateLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StateLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StateLib *StateLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StateLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StateLib *StateLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StateLib.Contract.contract.Transact(opts, method, params...)
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220d26cf239b558065d9b59692fe9194eefa862444eb2cac8124915cd7ca3863a7564736f6c63430008110033",
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
