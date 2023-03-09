// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package headerharness

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220955c01215ff254b5b1d51dd650b5d04070d88bdebdd7e00010f9481b0c1afa9e64736f6c63430008110033",
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

// HeaderHarnessMetaData contains all meta data concerning the HeaderHarness contract.
var HeaderHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"castToHeader\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_sender\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_optimisticSeconds\",\"type\":\"uint32\"}],\"name\":\"formatHeader\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headerLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"headerVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"headerVersion\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"isHeader\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetDestination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetOptimisticSeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetRecipient\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetSender\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"optimisticSeconds\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"recipient\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"recipientAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"sender\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"72a2f4a8": "castToHeader(uint40,bytes)",
		"06dc2d1a": "destination(uint40,bytes)",
		"ac124002": "formatHeader(uint32,bytes32,uint32,uint32,bytes32,uint32)",
		"80bfe8a3": "headerLength()",
		"5cf682c6": "headerVersion()",
		"9a3c57ad": "headerVersion(uint40,bytes)",
		"aabd9956": "isHeader(bytes)",
		"170cd79b": "nonce(uint40,bytes)",
		"d2c4428a": "offsetDestination()",
		"569e1eaf": "offsetNonce()",
		"4155c3d5": "offsetOptimisticSeconds()",
		"320bfc44": "offsetOrigin()",
		"a2ce1f35": "offsetRecipient()",
		"07fd670d": "offsetSender()",
		"0c096e8d": "offsetVersion()",
		"1900888f": "optimisticSeconds(uint40,bytes)",
		"d455d504": "origin(uint40,bytes)",
		"9b011d88": "recipient(uint40,bytes)",
		"11dca44d": "recipientAddress(uint40,bytes)",
		"ce429474": "sender(uint40,bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611366806100206000396000f3fe608060405234801561001057600080fd5b506004361061016c5760003560e01c806372a2f4a8116100cd578063aabd995611610081578063ce42947411610066578063ce429474146102d7578063d2c4428a146102ea578063d455d504146102f157600080fd5b8063aabd995614610294578063ac124002146102b757600080fd5b80639a3c57ad116100b25780639a3c57ad146102675780639b011d881461027a578063a2ce1f351461028d57600080fd5b806372a2f4a81461023f57806380bfe8a31461026057600080fd5b80631900888f116101245780634155c3d5116101095780634155c3d51461021b578063569e1eaf146102225780635cf682c61461022957600080fd5b80631900888f14610201578063320bfc441461021457600080fd5b80630c096e8d116101555780630c096e8d146101af57806311dca44d146101b6578063170cd79b146101ee57600080fd5b806306dc2d1a1461017157806307fd670d1461019e575b600080fd5b61018461017f3660046110d6565b610304565b60405163ffffffff90911681526020015b60405180910390f35b60065b604051908152602001610195565b60006101a1565b6101c96101c43660046110d6565b610327565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610195565b6101846101fc3660046110d6565b610341565b61018461020f3660046110d6565b61035b565b60026101a1565b604e6101a1565b60266101a1565b60015b60405161ffff9091168152602001610195565b61025261024d3660046110d6565b610375565b604051610195929190611194565b60526101a1565b61022c6102753660046110d6565b6103b0565b6101a16102883660046110d6565b6103ca565b602e6101a1565b6102a76102a23660046111bc565b6103e4565b6040519015158152602001610195565b6102ca6102c536600461120a565b6103fd565b6040516101959190611270565b6101a16102e53660046110d6565b61049f565b602a6101a1565b6101846102ff3660046110d6565b6104b9565b600061031e61031383856104d3565b62ffffff19166104f7565b90505b92915050565b600061031e61033683856104d3565b62ffffff191661052a565b600061031e61035083856104d3565b62ffffff191661053b565b600061031e61036a83856104d3565b62ffffff1916610567565b60006060600061038484610593565b905061039562ffffff1982166105a4565b6103a462ffffff1983166105c8565b92509250509250929050565b600061031e6103bf83856104d3565b62ffffff191661061b565b600061031e6103d983856104d3565b62ffffff1916610647565b60006103216103f283610593565b62ffffff1916610673565b604080517e0100000000000000000000000000000000000000000000000000000000000060208201527fffffffff0000000000000000000000000000000000000000000000000000000060e089811b821660228401526026830189905287811b8216604684015286811b8216604a840152604e830186905284901b16606e8201528151808203605201815260729091019091526060905b979650505050505050565b600061031e6104ae83856104d3565b62ffffff19166106ba565b600061031e6104c883856104d3565b62ffffff19166106e6565b8151600090602084016104ee64ffffffffff85168284610712565b95945050505050565b60008161050f62ffffff198216640301010000610759565b5061052362ffffff198416602a600461087c565b9392505050565b600061032161053883610647565b90565b60008161055362ffffff198216640301010000610759565b5061052362ffffff1984166026600461087c565b60008161057f62ffffff198216640301010000610759565b5061052362ffffff198416604e600461087c565b6000610321826403010100006104d3565b60008060606105b48160186112b2565b6105be91906112b2565b9290921c92915050565b60606000806105e58460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061060a84836020016108ac565b508181016020016040529052919050565b60008161063362ffffff198216640301010000610759565b5061052362ffffff1984166000600261087c565b60008161065f62ffffff198216640301010000610759565b5061052362ffffff198416602e6020610a88565b6000601882901c6bffffffffffffffffffffffff16600281101561069a5750600092915050565b60016106a58461061b565b61ffff16148015610523575060521492915050565b6000816106d262ffffff198216640301010000610759565b5061052362ffffff19841660066020610a88565b6000816106fe62ffffff198216640301010000610759565b5061052362ffffff1984166002600461087c565b60008061071f83856112b2565b905060405181111561072f575060005b806000036107445762ffffff19915050610523565b5050606092831b9190911790911b1760181b90565b60006107658383610c36565b610875576000610783610777856105a4565b64ffffffffff16610c58565b91505060006107988464ffffffffff16610c58565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086c9190611270565b60405180910390fd5b5090919050565b60006108898260206112c5565b6108949060086112de565b60ff166108a2858585610a88565b901c949350505050565b600062ffffff198084160361091d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e746572206465726566000000000000604482015260640161086c565b61092683610d42565b61098c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e746572206465726566000000604482015260640161086c565b60006109a68460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060006109c185610d7e565b6bffffffffffffffffffffffff1690506000806040519150858211156109e75760206060fd5b8386858560045afa905080610a58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f6620676173000000000000000000000000604482015260640161086c565b610494610a64886105a4565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b60008160ff16600003610a9d57506000610523565b610ab58460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610ad060ff8416856112b2565b1115610b5357610b20610ae285610d7e565b6bffffffffffffffffffffffff16610b088660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16610da5565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086c9190611270565b60208260ff161115610bc1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e20333220627974657300000000000000604482015260640161086c565b600882026000610bd086610d7e565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b60008164ffffffffff16610c49846105a4565b64ffffffffff16149392505050565b600080601f5b600f8160ff161115610ccb576000610c778260086112de565b60ff1685901c9050610c8881610f35565b61ffff16841793508160ff16601014610ca357601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610c5e565b50600f5b60ff8160ff161015610d3c576000610ce88260086112de565b60ff1685901c9050610cf981610f35565b61ffff16831792508160ff16600014610d1457601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610ccf565b50915091565b6000610d4d826105a4565b64ffffffffff1664ffffffffff03610d6757506000919050565b6000610d7283610f67565b60405110199392505050565b600080610d8d606060186112b2565b9290921c6bffffffffffffffffffffffff1692915050565b60606000610db286610c58565b9150506000610dc086610c58565b9150506000610dce86610c58565b9150506000610ddc86610c58565b604080517f54797065644d656d566965772f696e646578202d204f76657272616e2074686560208201527f20766965772e20536c6963652069732061742030780000000000000000000000818301527fffffffffffff000000000000000000000000000000000000000000000000000060d098891b811660558301527f2077697468206c656e6774682030780000000000000000000000000000000000605b830181905297891b8116606a8301527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060708301527f7800000000000000000000000000000000000000000000000000000000000000609083015295881b861660918201526097810196909652951b90921660a684015250507f2e0000000000000000000000000000000000000000000000000000000000000060ac8201528151808203608d01815260ad90910190915295945050505050565b6000610f4760048360ff16901c610fa0565b60ff1661ffff919091161760081b610f5e82610fa0565b60ff1617919050565b6000610f818260181c6bffffffffffffffffffffffff1690565b610f8a83610d7e565b016bffffffffffffffffffffffff169050919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110610fed57610fed611301565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261103c57600080fd5b813567ffffffffffffffff8082111561105757611057610ffc565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561109d5761109d610ffc565b816040528381528660208588010111156110b657600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156110e957600080fd5b823564ffffffffff811681146110fe57600080fd5b9150602083013567ffffffffffffffff81111561111a57600080fd5b6111268582860161102b565b9150509250929050565b6000815180845260005b818110156111565760208185018101518683018201520161113a565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b64ffffffffff831681526040602082015260006111b46040830184611130565b949350505050565b6000602082840312156111ce57600080fd5b813567ffffffffffffffff8111156111e557600080fd5b6111b48482850161102b565b803563ffffffff8116811461120557600080fd5b919050565b60008060008060008060c0878903121561122357600080fd5b61122c876111f1565b955060208701359450611241604088016111f1565b935061124f606088016111f1565b92506080870135915061126460a088016111f1565b90509295509295509295565b60208152600061031e6020830184611130565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561032157610321611283565b60ff828116828216039081111561032157610321611283565b60ff81811683821602908116908181146112fa576112fa611283565b5092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea26469706673582212202dd73610e73a922114a97f7b37b9ab865573f8def2e711967fbfb6a918d1ed3764736f6c63430008110033",
}

// HeaderHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use HeaderHarnessMetaData.ABI instead.
var HeaderHarnessABI = HeaderHarnessMetaData.ABI

// Deprecated: Use HeaderHarnessMetaData.Sigs instead.
// HeaderHarnessFuncSigs maps the 4-byte function signature to its string representation.
var HeaderHarnessFuncSigs = HeaderHarnessMetaData.Sigs

// HeaderHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HeaderHarnessMetaData.Bin instead.
var HeaderHarnessBin = HeaderHarnessMetaData.Bin

// DeployHeaderHarness deploys a new Ethereum contract, binding an instance of HeaderHarness to it.
func DeployHeaderHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HeaderHarness, error) {
	parsed, err := HeaderHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HeaderHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HeaderHarness{HeaderHarnessCaller: HeaderHarnessCaller{contract: contract}, HeaderHarnessTransactor: HeaderHarnessTransactor{contract: contract}, HeaderHarnessFilterer: HeaderHarnessFilterer{contract: contract}}, nil
}

// HeaderHarness is an auto generated Go binding around an Ethereum contract.
type HeaderHarness struct {
	HeaderHarnessCaller     // Read-only binding to the contract
	HeaderHarnessTransactor // Write-only binding to the contract
	HeaderHarnessFilterer   // Log filterer for contract events
}

// HeaderHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type HeaderHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HeaderHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HeaderHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HeaderHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HeaderHarnessSession struct {
	Contract     *HeaderHarness    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HeaderHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HeaderHarnessCallerSession struct {
	Contract *HeaderHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HeaderHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HeaderHarnessTransactorSession struct {
	Contract     *HeaderHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HeaderHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type HeaderHarnessRaw struct {
	Contract *HeaderHarness // Generic contract binding to access the raw methods on
}

// HeaderHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HeaderHarnessCallerRaw struct {
	Contract *HeaderHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// HeaderHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HeaderHarnessTransactorRaw struct {
	Contract *HeaderHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHeaderHarness creates a new instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarness(address common.Address, backend bind.ContractBackend) (*HeaderHarness, error) {
	contract, err := bindHeaderHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HeaderHarness{HeaderHarnessCaller: HeaderHarnessCaller{contract: contract}, HeaderHarnessTransactor: HeaderHarnessTransactor{contract: contract}, HeaderHarnessFilterer: HeaderHarnessFilterer{contract: contract}}, nil
}

// NewHeaderHarnessCaller creates a new read-only instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarnessCaller(address common.Address, caller bind.ContractCaller) (*HeaderHarnessCaller, error) {
	contract, err := bindHeaderHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderHarnessCaller{contract: contract}, nil
}

// NewHeaderHarnessTransactor creates a new write-only instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*HeaderHarnessTransactor, error) {
	contract, err := bindHeaderHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HeaderHarnessTransactor{contract: contract}, nil
}

// NewHeaderHarnessFilterer creates a new log filterer instance of HeaderHarness, bound to a specific deployed contract.
func NewHeaderHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*HeaderHarnessFilterer, error) {
	contract, err := bindHeaderHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HeaderHarnessFilterer{contract: contract}, nil
}

// bindHeaderHarness binds a generic wrapper to an already deployed contract.
func bindHeaderHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HeaderHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderHarness *HeaderHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderHarness.Contract.HeaderHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderHarness *HeaderHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderHarness.Contract.HeaderHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderHarness *HeaderHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderHarness.Contract.HeaderHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HeaderHarness *HeaderHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HeaderHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HeaderHarness *HeaderHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HeaderHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HeaderHarness *HeaderHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HeaderHarness.Contract.contract.Transact(opts, method, params...)
}

// CastToHeader is a free data retrieval call binding the contract method 0x72a2f4a8.
//
// Solidity: function castToHeader(uint40 , bytes _payload) view returns(uint40, bytes)
func (_HeaderHarness *HeaderHarnessCaller) CastToHeader(opts *bind.CallOpts, arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "castToHeader", arg0, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// CastToHeader is a free data retrieval call binding the contract method 0x72a2f4a8.
//
// Solidity: function castToHeader(uint40 , bytes _payload) view returns(uint40, bytes)
func (_HeaderHarness *HeaderHarnessSession) CastToHeader(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _HeaderHarness.Contract.CastToHeader(&_HeaderHarness.CallOpts, arg0, _payload)
}

// CastToHeader is a free data retrieval call binding the contract method 0x72a2f4a8.
//
// Solidity: function castToHeader(uint40 , bytes _payload) view returns(uint40, bytes)
func (_HeaderHarness *HeaderHarnessCallerSession) CastToHeader(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _HeaderHarness.Contract.CastToHeader(&_HeaderHarness.CallOpts, arg0, _payload)
}

// Destination is a free data retrieval call binding the contract method 0x06dc2d1a.
//
// Solidity: function destination(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) Destination(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "destination", _type, _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0x06dc2d1a.
//
// Solidity: function destination(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) Destination(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.Destination(&_HeaderHarness.CallOpts, _type, _payload)
}

// Destination is a free data retrieval call binding the contract method 0x06dc2d1a.
//
// Solidity: function destination(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) Destination(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.Destination(&_HeaderHarness.CallOpts, _type, _payload)
}

// FormatHeader is a free data retrieval call binding the contract method 0xac124002.
//
// Solidity: function formatHeader(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds) pure returns(bytes)
func (_HeaderHarness *HeaderHarnessCaller) FormatHeader(opts *bind.CallOpts, _origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32) ([]byte, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "formatHeader", _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatHeader is a free data retrieval call binding the contract method 0xac124002.
//
// Solidity: function formatHeader(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds) pure returns(bytes)
func (_HeaderHarness *HeaderHarnessSession) FormatHeader(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32) ([]byte, error) {
	return _HeaderHarness.Contract.FormatHeader(&_HeaderHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds)
}

// FormatHeader is a free data retrieval call binding the contract method 0xac124002.
//
// Solidity: function formatHeader(uint32 _origin, bytes32 _sender, uint32 _nonce, uint32 _destination, bytes32 _recipient, uint32 _optimisticSeconds) pure returns(bytes)
func (_HeaderHarness *HeaderHarnessCallerSession) FormatHeader(_origin uint32, _sender [32]byte, _nonce uint32, _destination uint32, _recipient [32]byte, _optimisticSeconds uint32) ([]byte, error) {
	return _HeaderHarness.Contract.FormatHeader(&_HeaderHarness.CallOpts, _origin, _sender, _nonce, _destination, _recipient, _optimisticSeconds)
}

// HeaderLength is a free data retrieval call binding the contract method 0x80bfe8a3.
//
// Solidity: function headerLength() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) HeaderLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "headerLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HeaderLength is a free data retrieval call binding the contract method 0x80bfe8a3.
//
// Solidity: function headerLength() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) HeaderLength() (*big.Int, error) {
	return _HeaderHarness.Contract.HeaderLength(&_HeaderHarness.CallOpts)
}

// HeaderLength is a free data retrieval call binding the contract method 0x80bfe8a3.
//
// Solidity: function headerLength() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) HeaderLength() (*big.Int, error) {
	return _HeaderHarness.Contract.HeaderLength(&_HeaderHarness.CallOpts)
}

// HeaderVersion is a free data retrieval call binding the contract method 0x5cf682c6.
//
// Solidity: function headerVersion() pure returns(uint16)
func (_HeaderHarness *HeaderHarnessCaller) HeaderVersion(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "headerVersion")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// HeaderVersion is a free data retrieval call binding the contract method 0x5cf682c6.
//
// Solidity: function headerVersion() pure returns(uint16)
func (_HeaderHarness *HeaderHarnessSession) HeaderVersion() (uint16, error) {
	return _HeaderHarness.Contract.HeaderVersion(&_HeaderHarness.CallOpts)
}

// HeaderVersion is a free data retrieval call binding the contract method 0x5cf682c6.
//
// Solidity: function headerVersion() pure returns(uint16)
func (_HeaderHarness *HeaderHarnessCallerSession) HeaderVersion() (uint16, error) {
	return _HeaderHarness.Contract.HeaderVersion(&_HeaderHarness.CallOpts)
}

// HeaderVersion0 is a free data retrieval call binding the contract method 0x9a3c57ad.
//
// Solidity: function headerVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_HeaderHarness *HeaderHarnessCaller) HeaderVersion0(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint16, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "headerVersion0", _type, _payload)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// HeaderVersion0 is a free data retrieval call binding the contract method 0x9a3c57ad.
//
// Solidity: function headerVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_HeaderHarness *HeaderHarnessSession) HeaderVersion0(_type *big.Int, _payload []byte) (uint16, error) {
	return _HeaderHarness.Contract.HeaderVersion0(&_HeaderHarness.CallOpts, _type, _payload)
}

// HeaderVersion0 is a free data retrieval call binding the contract method 0x9a3c57ad.
//
// Solidity: function headerVersion(uint40 _type, bytes _payload) pure returns(uint16)
func (_HeaderHarness *HeaderHarnessCallerSession) HeaderVersion0(_type *big.Int, _payload []byte) (uint16, error) {
	return _HeaderHarness.Contract.HeaderVersion0(&_HeaderHarness.CallOpts, _type, _payload)
}

// IsHeader is a free data retrieval call binding the contract method 0xaabd9956.
//
// Solidity: function isHeader(bytes _payload) pure returns(bool)
func (_HeaderHarness *HeaderHarnessCaller) IsHeader(opts *bind.CallOpts, _payload []byte) (bool, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "isHeader", _payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHeader is a free data retrieval call binding the contract method 0xaabd9956.
//
// Solidity: function isHeader(bytes _payload) pure returns(bool)
func (_HeaderHarness *HeaderHarnessSession) IsHeader(_payload []byte) (bool, error) {
	return _HeaderHarness.Contract.IsHeader(&_HeaderHarness.CallOpts, _payload)
}

// IsHeader is a free data retrieval call binding the contract method 0xaabd9956.
//
// Solidity: function isHeader(bytes _payload) pure returns(bool)
func (_HeaderHarness *HeaderHarnessCallerSession) IsHeader(_payload []byte) (bool, error) {
	return _HeaderHarness.Contract.IsHeader(&_HeaderHarness.CallOpts, _payload)
}

// Nonce is a free data retrieval call binding the contract method 0x170cd79b.
//
// Solidity: function nonce(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) Nonce(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "nonce", _type, _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x170cd79b.
//
// Solidity: function nonce(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) Nonce(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.Nonce(&_HeaderHarness.CallOpts, _type, _payload)
}

// Nonce is a free data retrieval call binding the contract method 0x170cd79b.
//
// Solidity: function nonce(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) Nonce(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.Nonce(&_HeaderHarness.CallOpts, _type, _payload)
}

// OffsetDestination is a free data retrieval call binding the contract method 0xd2c4428a.
//
// Solidity: function offsetDestination() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) OffsetDestination(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "offsetDestination")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetDestination is a free data retrieval call binding the contract method 0xd2c4428a.
//
// Solidity: function offsetDestination() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) OffsetDestination() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetDestination(&_HeaderHarness.CallOpts)
}

// OffsetDestination is a free data retrieval call binding the contract method 0xd2c4428a.
//
// Solidity: function offsetDestination() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) OffsetDestination() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetDestination(&_HeaderHarness.CallOpts)
}

// OffsetNonce is a free data retrieval call binding the contract method 0x569e1eaf.
//
// Solidity: function offsetNonce() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) OffsetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "offsetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetNonce is a free data retrieval call binding the contract method 0x569e1eaf.
//
// Solidity: function offsetNonce() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) OffsetNonce() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetNonce(&_HeaderHarness.CallOpts)
}

// OffsetNonce is a free data retrieval call binding the contract method 0x569e1eaf.
//
// Solidity: function offsetNonce() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) OffsetNonce() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetNonce(&_HeaderHarness.CallOpts)
}

// OffsetOptimisticSeconds is a free data retrieval call binding the contract method 0x4155c3d5.
//
// Solidity: function offsetOptimisticSeconds() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) OffsetOptimisticSeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "offsetOptimisticSeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetOptimisticSeconds is a free data retrieval call binding the contract method 0x4155c3d5.
//
// Solidity: function offsetOptimisticSeconds() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) OffsetOptimisticSeconds() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetOptimisticSeconds(&_HeaderHarness.CallOpts)
}

// OffsetOptimisticSeconds is a free data retrieval call binding the contract method 0x4155c3d5.
//
// Solidity: function offsetOptimisticSeconds() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) OffsetOptimisticSeconds() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetOptimisticSeconds(&_HeaderHarness.CallOpts)
}

// OffsetOrigin is a free data retrieval call binding the contract method 0x320bfc44.
//
// Solidity: function offsetOrigin() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) OffsetOrigin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "offsetOrigin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetOrigin is a free data retrieval call binding the contract method 0x320bfc44.
//
// Solidity: function offsetOrigin() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) OffsetOrigin() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetOrigin(&_HeaderHarness.CallOpts)
}

// OffsetOrigin is a free data retrieval call binding the contract method 0x320bfc44.
//
// Solidity: function offsetOrigin() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) OffsetOrigin() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetOrigin(&_HeaderHarness.CallOpts)
}

// OffsetRecipient is a free data retrieval call binding the contract method 0xa2ce1f35.
//
// Solidity: function offsetRecipient() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) OffsetRecipient(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "offsetRecipient")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetRecipient is a free data retrieval call binding the contract method 0xa2ce1f35.
//
// Solidity: function offsetRecipient() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) OffsetRecipient() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetRecipient(&_HeaderHarness.CallOpts)
}

// OffsetRecipient is a free data retrieval call binding the contract method 0xa2ce1f35.
//
// Solidity: function offsetRecipient() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) OffsetRecipient() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetRecipient(&_HeaderHarness.CallOpts)
}

// OffsetSender is a free data retrieval call binding the contract method 0x07fd670d.
//
// Solidity: function offsetSender() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) OffsetSender(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "offsetSender")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetSender is a free data retrieval call binding the contract method 0x07fd670d.
//
// Solidity: function offsetSender() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) OffsetSender() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetSender(&_HeaderHarness.CallOpts)
}

// OffsetSender is a free data retrieval call binding the contract method 0x07fd670d.
//
// Solidity: function offsetSender() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) OffsetSender() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetSender(&_HeaderHarness.CallOpts)
}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCaller) OffsetVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "offsetVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessSession) OffsetVersion() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetVersion(&_HeaderHarness.CallOpts)
}

// OffsetVersion is a free data retrieval call binding the contract method 0x0c096e8d.
//
// Solidity: function offsetVersion() pure returns(uint256)
func (_HeaderHarness *HeaderHarnessCallerSession) OffsetVersion() (*big.Int, error) {
	return _HeaderHarness.Contract.OffsetVersion(&_HeaderHarness.CallOpts)
}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x1900888f.
//
// Solidity: function optimisticSeconds(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) OptimisticSeconds(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "optimisticSeconds", _type, _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x1900888f.
//
// Solidity: function optimisticSeconds(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) OptimisticSeconds(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.OptimisticSeconds(&_HeaderHarness.CallOpts, _type, _payload)
}

// OptimisticSeconds is a free data retrieval call binding the contract method 0x1900888f.
//
// Solidity: function optimisticSeconds(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) OptimisticSeconds(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.OptimisticSeconds(&_HeaderHarness.CallOpts, _type, _payload)
}

// Origin is a free data retrieval call binding the contract method 0xd455d504.
//
// Solidity: function origin(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCaller) Origin(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "origin", _type, _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0xd455d504.
//
// Solidity: function origin(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessSession) Origin(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.Origin(&_HeaderHarness.CallOpts, _type, _payload)
}

// Origin is a free data retrieval call binding the contract method 0xd455d504.
//
// Solidity: function origin(uint40 _type, bytes _payload) pure returns(uint32)
func (_HeaderHarness *HeaderHarnessCallerSession) Origin(_type *big.Int, _payload []byte) (uint32, error) {
	return _HeaderHarness.Contract.Origin(&_HeaderHarness.CallOpts, _type, _payload)
}

// Recipient is a free data retrieval call binding the contract method 0x9b011d88.
//
// Solidity: function recipient(uint40 _type, bytes _payload) pure returns(bytes32)
func (_HeaderHarness *HeaderHarnessCaller) Recipient(opts *bind.CallOpts, _type *big.Int, _payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "recipient", _type, _payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Recipient is a free data retrieval call binding the contract method 0x9b011d88.
//
// Solidity: function recipient(uint40 _type, bytes _payload) pure returns(bytes32)
func (_HeaderHarness *HeaderHarnessSession) Recipient(_type *big.Int, _payload []byte) ([32]byte, error) {
	return _HeaderHarness.Contract.Recipient(&_HeaderHarness.CallOpts, _type, _payload)
}

// Recipient is a free data retrieval call binding the contract method 0x9b011d88.
//
// Solidity: function recipient(uint40 _type, bytes _payload) pure returns(bytes32)
func (_HeaderHarness *HeaderHarnessCallerSession) Recipient(_type *big.Int, _payload []byte) ([32]byte, error) {
	return _HeaderHarness.Contract.Recipient(&_HeaderHarness.CallOpts, _type, _payload)
}

// RecipientAddress is a free data retrieval call binding the contract method 0x11dca44d.
//
// Solidity: function recipientAddress(uint40 _type, bytes _payload) pure returns(address)
func (_HeaderHarness *HeaderHarnessCaller) RecipientAddress(opts *bind.CallOpts, _type *big.Int, _payload []byte) (common.Address, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "recipientAddress", _type, _payload)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecipientAddress is a free data retrieval call binding the contract method 0x11dca44d.
//
// Solidity: function recipientAddress(uint40 _type, bytes _payload) pure returns(address)
func (_HeaderHarness *HeaderHarnessSession) RecipientAddress(_type *big.Int, _payload []byte) (common.Address, error) {
	return _HeaderHarness.Contract.RecipientAddress(&_HeaderHarness.CallOpts, _type, _payload)
}

// RecipientAddress is a free data retrieval call binding the contract method 0x11dca44d.
//
// Solidity: function recipientAddress(uint40 _type, bytes _payload) pure returns(address)
func (_HeaderHarness *HeaderHarnessCallerSession) RecipientAddress(_type *big.Int, _payload []byte) (common.Address, error) {
	return _HeaderHarness.Contract.RecipientAddress(&_HeaderHarness.CallOpts, _type, _payload)
}

// Sender is a free data retrieval call binding the contract method 0xce429474.
//
// Solidity: function sender(uint40 _type, bytes _payload) pure returns(bytes32)
func (_HeaderHarness *HeaderHarnessCaller) Sender(opts *bind.CallOpts, _type *big.Int, _payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _HeaderHarness.contract.Call(opts, &out, "sender", _type, _payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Sender is a free data retrieval call binding the contract method 0xce429474.
//
// Solidity: function sender(uint40 _type, bytes _payload) pure returns(bytes32)
func (_HeaderHarness *HeaderHarnessSession) Sender(_type *big.Int, _payload []byte) ([32]byte, error) {
	return _HeaderHarness.Contract.Sender(&_HeaderHarness.CallOpts, _type, _payload)
}

// Sender is a free data retrieval call binding the contract method 0xce429474.
//
// Solidity: function sender(uint40 _type, bytes _payload) pure returns(bytes32)
func (_HeaderHarness *HeaderHarnessCallerSession) Sender(_type *big.Int, _payload []byte) ([32]byte, error) {
	return _HeaderHarness.Contract.Sender(&_HeaderHarness.CallOpts, _type, _payload)
}

// SynapseTypesMetaData contains all meta data concerning the SynapseTypes contract.
var SynapseTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203e3bcd05141bd0e6718dd149f8813ee0acc5812f69bc423cea42c7ae3eb7dd3c64736f6c63430008110033",
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

// TypeCastsMetaData contains all meta data concerning the TypeCasts contract.
var TypeCastsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200b2b11c0ad038f76e18987278659539d2eab6d7c7d88acc0e915dde290d0dea764736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea264697066735822122035e6a80023e3e894a57d0d576756dd4bc2dc7b8e58568ca7ba988f478106fdd864736f6c63430008110033",
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
