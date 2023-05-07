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

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e977210b43116a9b7d5cef0cce72d6aa126fa0559678c54764d4f845823bb7af64736f6c63430008110033",
}

// MemViewLibABI is the input ABI used to generate the binding from.
// Deprecated: Use MemViewLibMetaData.ABI instead.
var MemViewLibABI = MemViewLibMetaData.ABI

// MemViewLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MemViewLibMetaData.Bin instead.
var MemViewLibBin = MemViewLibMetaData.Bin

// DeployMemViewLib deploys a new Ethereum contract, binding an instance of MemViewLib to it.
func DeployMemViewLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MemViewLib, error) {
	parsed, err := MemViewLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MemViewLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MemViewLib{MemViewLibCaller: MemViewLibCaller{contract: contract}, MemViewLibTransactor: MemViewLibTransactor{contract: contract}, MemViewLibFilterer: MemViewLibFilterer{contract: contract}}, nil
}

// MemViewLib is an auto generated Go binding around an Ethereum contract.
type MemViewLib struct {
	MemViewLibCaller     // Read-only binding to the contract
	MemViewLibTransactor // Write-only binding to the contract
	MemViewLibFilterer   // Log filterer for contract events
}

// MemViewLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MemViewLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MemViewLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MemViewLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MemViewLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MemViewLibSession struct {
	Contract     *MemViewLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MemViewLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MemViewLibCallerSession struct {
	Contract *MemViewLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MemViewLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MemViewLibTransactorSession struct {
	Contract     *MemViewLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MemViewLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MemViewLibRaw struct {
	Contract *MemViewLib // Generic contract binding to access the raw methods on
}

// MemViewLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MemViewLibCallerRaw struct {
	Contract *MemViewLibCaller // Generic read-only contract binding to access the raw methods on
}

// MemViewLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MemViewLibTransactorRaw struct {
	Contract *MemViewLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMemViewLib creates a new instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLib(address common.Address, backend bind.ContractBackend) (*MemViewLib, error) {
	contract, err := bindMemViewLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MemViewLib{MemViewLibCaller: MemViewLibCaller{contract: contract}, MemViewLibTransactor: MemViewLibTransactor{contract: contract}, MemViewLibFilterer: MemViewLibFilterer{contract: contract}}, nil
}

// NewMemViewLibCaller creates a new read-only instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibCaller(address common.Address, caller bind.ContractCaller) (*MemViewLibCaller, error) {
	contract, err := bindMemViewLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MemViewLibCaller{contract: contract}, nil
}

// NewMemViewLibTransactor creates a new write-only instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MemViewLibTransactor, error) {
	contract, err := bindMemViewLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MemViewLibTransactor{contract: contract}, nil
}

// NewMemViewLibFilterer creates a new log filterer instance of MemViewLib, bound to a specific deployed contract.
func NewMemViewLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MemViewLibFilterer, error) {
	contract, err := bindMemViewLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MemViewLibFilterer{contract: contract}, nil
}

// bindMemViewLib binds a generic wrapper to an already deployed contract.
func bindMemViewLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MemViewLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemViewLib *MemViewLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemViewLib.Contract.MemViewLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemViewLib *MemViewLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemViewLib.Contract.MemViewLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemViewLib *MemViewLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemViewLib.Contract.MemViewLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MemViewLib *MemViewLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MemViewLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MemViewLib *MemViewLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MemViewLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MemViewLib *MemViewLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MemViewLib.Contract.contract.Transact(opts, method, params...)
}

// StateHarnessMetaData contains all meta data concerning the StateHarness contract.
var StateHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"equals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"formatState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"}],\"name\":\"leftLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"rightLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"subLeafs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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
	Bin: "0x608060405234801561001057600080fd5b50610d66806100206000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c80639aaa18261161008c578063cb3eb0e111610066578063cb3eb0e114610288578063d7a7a72c1461029b578063e948e600146102ae578063edaa471d146102c157600080fd5b80639aaa18261461023a578063aae6d88414610262578063c2e9e2081461027557600080fd5b80634e765004116100bd5780634e765004146101de578063503d0bed146102065780635fed02611461022757600080fd5b8063137e618a146100e45780631c9aa2221461010c578063365b4b6714610135575b600080fd5b6100f76100f2366004610af6565b6102d4565b60405190151581526020015b60405180910390f35b61011f61011a366004610b5a565b6102fa565b60405164ffffffffff9091168152602001610103565b6101d1610143366004610bbd565b6040805160208101969096527fffffffff0000000000000000000000000000000000000000000000000000000060e095861b8116878301529390941b90921660448501527fffffffffff00000000000000000000000000000000000000000000000000000060d891821b8116604886015291901b16604d830152805180830360320181526052909201905290565b6040516101039190610c1b565b6101f16101ec366004610b5a565b61030d565b60405163ffffffff9091168152602001610103565b610219610214366004610c87565b610320565b604051908152602001610103565b6101d1610235366004610b5a565b6103a9565b61024d610248366004610b5a565b6103c1565b60408051928352602083019190915201610103565b6100f7610270366004610b5a565b6103de565b610219610283366004610b5a565b610404565b6101f1610296366004610b5a565b610417565b6102196102a9366004610b5a565b61042a565b61011f6102bc366004610b5a565b61043d565b6102196102cf366004610cca565b610450565b60006102f16102e28361045c565b6102eb8561045c565b9061046f565b90505b92915050565b60006102f46103088361045c565b610490565b60006102f461031b8361045c565b6104a2565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b166020808301919091527fffffffffff00000000000000000000000000000000000000000000000000000060d886811b8216602485015285901b1660298301528251808303600e018152602e90920190925280519101205b9392505050565b606060006103b68361045c565b90506103a2816104b1565b6000806103d56103d08461045c565b61050e565b91509150915091565b60006102f46103ec83610538565b6fffffffffffffffffffffffffffffffff1660321490565b60006102f46104128361045c565b610553565b60006102f46104258361045c565b610561565b60006102f46104388361045c565b610570565b60006102f461044b8361045c565b6105af565b60006102f183836105be565b60006102f461046a83610538565b61061f565b600061047f826106a5565b6106a5565b610488846106a5565b149392505050565b60006102f4602d6005845b91906106d0565b60006102f4602460048461049b565b604051806104c283602083016106f1565b506fffffffffffffffffffffffffffffffff83166000601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168301602001604052509052919050565b6000808261052061047a8260246107a0565b925061053061047a8260246107ad565b915050915091565b80516000906020830161054b818361080f565b949350505050565b60006102f482826020610872565b60006102f4602060048461049b565b600080600061057e8461050e565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b60006102f4602860058461049b565b6000828260405160200161060192919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b600060326fffffffffffffffffffffffffffffffff8316146106a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f742061207374617465000000000000000000000000000000000000000000604482015260640160405180910390fd5b5090565b6000806106b28360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b6000806106de858585610872565b602084900360031b1c9150509392505050565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c908085101561074b576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa90508061078e576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417979650505050505050565b60006102f183828461097c565b60006fffffffffffffffffffffffffffffffff8316808311156107fc576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61054b8361080a8660801c90565b018483035b60008061081c8385610cf6565b905060405181111561082c575060005b80600003610866576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b831761054b565b600081600003610884575060006103a2565b60208211156108bf576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166108dc8385610cf6565b1115610914576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006109258660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b6000806109898560801c90565b9050610994856109f6565b8361099f8684610cf6565b6109a99190610cf6565b11156109e1576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109ed8482018461080f565b95945050505050565b60006fffffffffffffffffffffffffffffffff8216610a158360801c90565b0192915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610a5c57600080fd5b813567ffffffffffffffff80821115610a7757610a77610a1c565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610abd57610abd610a1c565b81604052838152866020858801011115610ad657600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610b0957600080fd5b823567ffffffffffffffff80821115610b2157600080fd5b610b2d86838701610a4b565b93506020850135915080821115610b4357600080fd5b50610b5085828601610a4b565b9150509250929050565b600060208284031215610b6c57600080fd5b813567ffffffffffffffff811115610b8357600080fd5b61054b84828501610a4b565b803563ffffffff81168114610ba357600080fd5b919050565b803564ffffffffff81168114610ba357600080fd5b600080600080600060a08688031215610bd557600080fd5b85359450610be560208701610b8f565b9350610bf360408701610b8f565b9250610c0160608701610ba8565b9150610c0f60808701610ba8565b90509295509295909350565b600060208083528351808285015260005b81811015610c4857858101830151858201604001528201610c2c565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b600080600060608486031215610c9c57600080fd5b610ca584610b8f565b9250610cb360208501610ba8565b9150610cc160408501610ba8565b90509250925092565b60008060408385031215610cdd57600080fd5b82359150610ced60208401610b8f565b90509250929050565b808201808211156102f4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea26469706673582212204af8b155d8943eb4d02773987ad55d1ad2f3b9529bede77204a6ac67556c632c64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122044bff0d319625920ed6ba5fe468215c81ae9e670c102a8ed2ccf14dab848b77f64736f6c63430008110033",
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
