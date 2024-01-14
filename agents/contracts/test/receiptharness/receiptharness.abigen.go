// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package receiptharness

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

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208c9ff5c8ae0945b0c310bf8ac4023f902435a7617ef72f632817973b5b0ac26664736f6c63430008110033",
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
	parsed, err := MemViewLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ReceiptHarnessMetaData contains all meta data concerning the ReceiptHarness contract.
var ReceiptHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedReceipt\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"attNotary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToReceipt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"destination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"equals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"finalExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"firstExecutor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destination_\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"snapshotRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"stateIndex_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"attNotary_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"firstExecutor_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"finalExecutor_\",\"type\":\"address\"}],\"name\":\"formatReceipt\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hashInvalid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hashValid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isReceipt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"messageHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"snapshotRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"stateIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e152a8cb": "attNotary(bytes)",
		"5cab5d3b": "castToReceipt(bytes)",
		"c81aa9c8": "destination(bytes)",
		"137e618a": "equals(bytes,bytes)",
		"f7e6a05b": "finalExecutor(bytes)",
		"27f2ee36": "firstExecutor(bytes)",
		"c3bfda6c": "formatReceipt(uint32,uint32,bytes32,bytes32,uint8,address,address,address)",
		"60cf3bf0": "hashInvalid(bytes)",
		"730dbf63": "hashValid(bytes)",
		"0bb3b580": "isReceipt(bytes)",
		"ed54c3b6": "messageHash(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"854dfcd7": "snapshotRoot(bytes)",
		"595271d1": "stateIndex(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610c09806100206000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063854dfcd71161008c578063cb3eb0e111610066578063cb3eb0e1146102f5578063e152a8cb14610308578063ed54c3b61461031b578063f7e6a05b1461032e57600080fd5b8063854dfcd7146101db578063c3bfda6c146101ee578063c81aa9c8146102cd57600080fd5b8063595271d1116100c8578063595271d1146101625780635cab5d3b1461018757806360cf3bf0146101a7578063730dbf63146101c857600080fd5b80630bb3b580146100ef578063137e618a1461011757806327f2ee361461012a575b600080fd5b6101026100fd3660046109c6565b610341565b60405190151581526020015b60405180910390f35b6101026101253660046109fb565b61036d565b61013d6101383660046109c6565b61039c565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161010e565b6101756101703660046109c6565b6103af565b60405160ff909116815260200161010e565b61019a6101953660046109c6565b6103c2565b60405161010e9190610a5f565b6101ba6101b53660046109c6565b6103da565b60405190815260200161010e565b6101ba6101d63660046109c6565b6103f0565b6101ba6101e93660046109c6565b610406565b61019a6101fc366004610b08565b6040805160e0998a1b7fffffffff0000000000000000000000000000000000000000000000000000000090811660208301529890991b90971660248901526028880195909552604887019390935260f89190911b7fff00000000000000000000000000000000000000000000000000000000000000166068860152606090811b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000908116606987015291811b8216607d8601529190911b1660918301528051608581840301815260a5909201905290565b6102e06102db3660046109c6565b610419565b60405163ffffffff909116815260200161010e565b6102e06103033660046109c6565b61042c565b61013d6103163660046109c6565b61043f565b6101ba6103293660046109c6565b610452565b61013d61033c3660046109c6565b610465565b600061036761034f83610478565b6fffffffffffffffffffffffffffffffff1660851490565b92915050565b600061039561038361037e84610478565b610493565b61038f61037e86610478565b906104e5565b9392505050565b60006103676103aa83610501565b61050f565b60006103676103bd83610501565b61051e565b606060006103cf83610501565b905061039581610530565b60006103676103eb61037e84610478565b61058d565b600061036761040161037e84610478565b6105bb565b600061036761041483610501565b6105e7565b600061036761042783610501565b6105f9565b600061036761043a83610501565b610607565b600061036761044d83610501565b610615565b600061036761046083610501565b610622565b600061036761047383610501565b610631565b80516000906020830161048b818361063e565b949350505050565b600060856fffffffffffffffffffffffffffffffff8316146104e1576040517f76b4e13c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b60006104f0826106a1565b6104f9846106a1565b149392505050565b600061036761037e83610478565b6000610367605d835b906106cc565b600061036760486001845b91906106d6565b6040518061054183602083016106f7565b506fffffffffffffffffffffffffffffffff83166000601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168301602001604052509052919050565b60006103677fdf42b2c0137811ba604f5c79e20c4d6b94770aa819cc524eca444056544f8ab7835b906107a6565b60006103677fb38669e8ca41a27fcd85729b868e8ab047d0f142073a017213e58f0a91e88ef3836105b5565b600061036760286020845b91906107e2565b600061036760048084610529565b600061036781600484610529565b6000610367604983610518565b600061036760086020846105f2565b6000610367607183610518565b60008061064b8385610b99565b905060405181111561065b575060005b80600003610695576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b831761048b565b6000806106ae8360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b6000610395838360145b6000806106e48585856107e2565b602084900360031b1c9150509392505050565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c9080851015610751576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa905080610794576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417979650505050505050565b6000816107b2846106a1565b60408051602081019390935282015260600160405160208183030381529060405280519060200120905092915050565b6000816000036107f457506000610395565b602082111561082f576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff841661084c8385610b99565b1115610884576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006108958660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261092c57600080fd5b813567ffffffffffffffff80821115610947576109476108ec565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561098d5761098d6108ec565b816040528381528660208588010111156109a657600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156109d857600080fd5b813567ffffffffffffffff8111156109ef57600080fd5b61048b8482850161091b565b60008060408385031215610a0e57600080fd5b823567ffffffffffffffff80821115610a2657600080fd5b610a328683870161091b565b93506020850135915080821115610a4857600080fd5b50610a558582860161091b565b9150509250929050565b600060208083528351808285015260005b81811015610a8c57858101830151858201604001528201610a70565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b803563ffffffff81168114610adf57600080fd5b919050565b803573ffffffffffffffffffffffffffffffffffffffff81168114610adf57600080fd5b600080600080600080600080610100898b031215610b2557600080fd5b610b2e89610acb565b9750610b3c60208a01610acb565b96506040890135955060608901359450608089013560ff81168114610b6057600080fd5b9350610b6e60a08a01610ae4565b9250610b7c60c08a01610ae4565b9150610b8a60e08a01610ae4565b90509295985092959890939650565b80820180821115610367577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea264697066735822122052b2b9231b76e4c4c31e614753009f7513a2ada3af2c039576adcd2563189f8664736f6c63430008110033",
}

// ReceiptHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiptHarnessMetaData.ABI instead.
var ReceiptHarnessABI = ReceiptHarnessMetaData.ABI

// Deprecated: Use ReceiptHarnessMetaData.Sigs instead.
// ReceiptHarnessFuncSigs maps the 4-byte function signature to its string representation.
var ReceiptHarnessFuncSigs = ReceiptHarnessMetaData.Sigs

// ReceiptHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiptHarnessMetaData.Bin instead.
var ReceiptHarnessBin = ReceiptHarnessMetaData.Bin

// DeployReceiptHarness deploys a new Ethereum contract, binding an instance of ReceiptHarness to it.
func DeployReceiptHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiptHarness, error) {
	parsed, err := ReceiptHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiptHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiptHarness{ReceiptHarnessCaller: ReceiptHarnessCaller{contract: contract}, ReceiptHarnessTransactor: ReceiptHarnessTransactor{contract: contract}, ReceiptHarnessFilterer: ReceiptHarnessFilterer{contract: contract}}, nil
}

// ReceiptHarness is an auto generated Go binding around an Ethereum contract.
type ReceiptHarness struct {
	ReceiptHarnessCaller     // Read-only binding to the contract
	ReceiptHarnessTransactor // Write-only binding to the contract
	ReceiptHarnessFilterer   // Log filterer for contract events
}

// ReceiptHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptHarnessSession struct {
	Contract     *ReceiptHarness   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiptHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptHarnessCallerSession struct {
	Contract *ReceiptHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ReceiptHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptHarnessTransactorSession struct {
	Contract     *ReceiptHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ReceiptHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptHarnessRaw struct {
	Contract *ReceiptHarness // Generic contract binding to access the raw methods on
}

// ReceiptHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptHarnessCallerRaw struct {
	Contract *ReceiptHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptHarnessTransactorRaw struct {
	Contract *ReceiptHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptHarness creates a new instance of ReceiptHarness, bound to a specific deployed contract.
func NewReceiptHarness(address common.Address, backend bind.ContractBackend) (*ReceiptHarness, error) {
	contract, err := bindReceiptHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptHarness{ReceiptHarnessCaller: ReceiptHarnessCaller{contract: contract}, ReceiptHarnessTransactor: ReceiptHarnessTransactor{contract: contract}, ReceiptHarnessFilterer: ReceiptHarnessFilterer{contract: contract}}, nil
}

// NewReceiptHarnessCaller creates a new read-only instance of ReceiptHarness, bound to a specific deployed contract.
func NewReceiptHarnessCaller(address common.Address, caller bind.ContractCaller) (*ReceiptHarnessCaller, error) {
	contract, err := bindReceiptHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptHarnessCaller{contract: contract}, nil
}

// NewReceiptHarnessTransactor creates a new write-only instance of ReceiptHarness, bound to a specific deployed contract.
func NewReceiptHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptHarnessTransactor, error) {
	contract, err := bindReceiptHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptHarnessTransactor{contract: contract}, nil
}

// NewReceiptHarnessFilterer creates a new log filterer instance of ReceiptHarness, bound to a specific deployed contract.
func NewReceiptHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptHarnessFilterer, error) {
	contract, err := bindReceiptHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptHarnessFilterer{contract: contract}, nil
}

// bindReceiptHarness binds a generic wrapper to an already deployed contract.
func bindReceiptHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiptHarnessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptHarness *ReceiptHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptHarness.Contract.ReceiptHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptHarness *ReceiptHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptHarness.Contract.ReceiptHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptHarness *ReceiptHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptHarness.Contract.ReceiptHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptHarness *ReceiptHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptHarness *ReceiptHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptHarness *ReceiptHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptHarness.Contract.contract.Transact(opts, method, params...)
}

// AttNotary is a free data retrieval call binding the contract method 0xe152a8cb.
//
// Solidity: function attNotary(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessCaller) AttNotary(opts *bind.CallOpts, payload []byte) (common.Address, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "attNotary", payload)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttNotary is a free data retrieval call binding the contract method 0xe152a8cb.
//
// Solidity: function attNotary(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessSession) AttNotary(payload []byte) (common.Address, error) {
	return _ReceiptHarness.Contract.AttNotary(&_ReceiptHarness.CallOpts, payload)
}

// AttNotary is a free data retrieval call binding the contract method 0xe152a8cb.
//
// Solidity: function attNotary(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessCallerSession) AttNotary(payload []byte) (common.Address, error) {
	return _ReceiptHarness.Contract.AttNotary(&_ReceiptHarness.CallOpts, payload)
}

// CastToReceipt is a free data retrieval call binding the contract method 0x5cab5d3b.
//
// Solidity: function castToReceipt(bytes payload) view returns(bytes)
func (_ReceiptHarness *ReceiptHarnessCaller) CastToReceipt(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "castToReceipt", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CastToReceipt is a free data retrieval call binding the contract method 0x5cab5d3b.
//
// Solidity: function castToReceipt(bytes payload) view returns(bytes)
func (_ReceiptHarness *ReceiptHarnessSession) CastToReceipt(payload []byte) ([]byte, error) {
	return _ReceiptHarness.Contract.CastToReceipt(&_ReceiptHarness.CallOpts, payload)
}

// CastToReceipt is a free data retrieval call binding the contract method 0x5cab5d3b.
//
// Solidity: function castToReceipt(bytes payload) view returns(bytes)
func (_ReceiptHarness *ReceiptHarnessCallerSession) CastToReceipt(payload []byte) ([]byte, error) {
	return _ReceiptHarness.Contract.CastToReceipt(&_ReceiptHarness.CallOpts, payload)
}

// Destination is a free data retrieval call binding the contract method 0xc81aa9c8.
//
// Solidity: function destination(bytes payload) pure returns(uint32)
func (_ReceiptHarness *ReceiptHarnessCaller) Destination(opts *bind.CallOpts, payload []byte) (uint32, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "destination", payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Destination is a free data retrieval call binding the contract method 0xc81aa9c8.
//
// Solidity: function destination(bytes payload) pure returns(uint32)
func (_ReceiptHarness *ReceiptHarnessSession) Destination(payload []byte) (uint32, error) {
	return _ReceiptHarness.Contract.Destination(&_ReceiptHarness.CallOpts, payload)
}

// Destination is a free data retrieval call binding the contract method 0xc81aa9c8.
//
// Solidity: function destination(bytes payload) pure returns(uint32)
func (_ReceiptHarness *ReceiptHarnessCallerSession) Destination(payload []byte) (uint32, error) {
	return _ReceiptHarness.Contract.Destination(&_ReceiptHarness.CallOpts, payload)
}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_ReceiptHarness *ReceiptHarnessCaller) Equals(opts *bind.CallOpts, a []byte, b []byte) (bool, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "equals", a, b)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_ReceiptHarness *ReceiptHarnessSession) Equals(a []byte, b []byte) (bool, error) {
	return _ReceiptHarness.Contract.Equals(&_ReceiptHarness.CallOpts, a, b)
}

// Equals is a free data retrieval call binding the contract method 0x137e618a.
//
// Solidity: function equals(bytes a, bytes b) pure returns(bool)
func (_ReceiptHarness *ReceiptHarnessCallerSession) Equals(a []byte, b []byte) (bool, error) {
	return _ReceiptHarness.Contract.Equals(&_ReceiptHarness.CallOpts, a, b)
}

// FinalExecutor is a free data retrieval call binding the contract method 0xf7e6a05b.
//
// Solidity: function finalExecutor(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessCaller) FinalExecutor(opts *bind.CallOpts, payload []byte) (common.Address, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "finalExecutor", payload)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FinalExecutor is a free data retrieval call binding the contract method 0xf7e6a05b.
//
// Solidity: function finalExecutor(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessSession) FinalExecutor(payload []byte) (common.Address, error) {
	return _ReceiptHarness.Contract.FinalExecutor(&_ReceiptHarness.CallOpts, payload)
}

// FinalExecutor is a free data retrieval call binding the contract method 0xf7e6a05b.
//
// Solidity: function finalExecutor(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessCallerSession) FinalExecutor(payload []byte) (common.Address, error) {
	return _ReceiptHarness.Contract.FinalExecutor(&_ReceiptHarness.CallOpts, payload)
}

// FirstExecutor is a free data retrieval call binding the contract method 0x27f2ee36.
//
// Solidity: function firstExecutor(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessCaller) FirstExecutor(opts *bind.CallOpts, payload []byte) (common.Address, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "firstExecutor", payload)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FirstExecutor is a free data retrieval call binding the contract method 0x27f2ee36.
//
// Solidity: function firstExecutor(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessSession) FirstExecutor(payload []byte) (common.Address, error) {
	return _ReceiptHarness.Contract.FirstExecutor(&_ReceiptHarness.CallOpts, payload)
}

// FirstExecutor is a free data retrieval call binding the contract method 0x27f2ee36.
//
// Solidity: function firstExecutor(bytes payload) pure returns(address)
func (_ReceiptHarness *ReceiptHarnessCallerSession) FirstExecutor(payload []byte) (common.Address, error) {
	return _ReceiptHarness.Contract.FirstExecutor(&_ReceiptHarness.CallOpts, payload)
}

// FormatReceipt is a free data retrieval call binding the contract method 0xc3bfda6c.
//
// Solidity: function formatReceipt(uint32 origin_, uint32 destination_, bytes32 messageHash_, bytes32 snapshotRoot_, uint8 stateIndex_, address attNotary_, address firstExecutor_, address finalExecutor_) pure returns(bytes)
func (_ReceiptHarness *ReceiptHarnessCaller) FormatReceipt(opts *bind.CallOpts, origin_ uint32, destination_ uint32, messageHash_ [32]byte, snapshotRoot_ [32]byte, stateIndex_ uint8, attNotary_ common.Address, firstExecutor_ common.Address, finalExecutor_ common.Address) ([]byte, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "formatReceipt", origin_, destination_, messageHash_, snapshotRoot_, stateIndex_, attNotary_, firstExecutor_, finalExecutor_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatReceipt is a free data retrieval call binding the contract method 0xc3bfda6c.
//
// Solidity: function formatReceipt(uint32 origin_, uint32 destination_, bytes32 messageHash_, bytes32 snapshotRoot_, uint8 stateIndex_, address attNotary_, address firstExecutor_, address finalExecutor_) pure returns(bytes)
func (_ReceiptHarness *ReceiptHarnessSession) FormatReceipt(origin_ uint32, destination_ uint32, messageHash_ [32]byte, snapshotRoot_ [32]byte, stateIndex_ uint8, attNotary_ common.Address, firstExecutor_ common.Address, finalExecutor_ common.Address) ([]byte, error) {
	return _ReceiptHarness.Contract.FormatReceipt(&_ReceiptHarness.CallOpts, origin_, destination_, messageHash_, snapshotRoot_, stateIndex_, attNotary_, firstExecutor_, finalExecutor_)
}

// FormatReceipt is a free data retrieval call binding the contract method 0xc3bfda6c.
//
// Solidity: function formatReceipt(uint32 origin_, uint32 destination_, bytes32 messageHash_, bytes32 snapshotRoot_, uint8 stateIndex_, address attNotary_, address firstExecutor_, address finalExecutor_) pure returns(bytes)
func (_ReceiptHarness *ReceiptHarnessCallerSession) FormatReceipt(origin_ uint32, destination_ uint32, messageHash_ [32]byte, snapshotRoot_ [32]byte, stateIndex_ uint8, attNotary_ common.Address, firstExecutor_ common.Address, finalExecutor_ common.Address) ([]byte, error) {
	return _ReceiptHarness.Contract.FormatReceipt(&_ReceiptHarness.CallOpts, origin_, destination_, messageHash_, snapshotRoot_, stateIndex_, attNotary_, firstExecutor_, finalExecutor_)
}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCaller) HashInvalid(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "hashInvalid", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessSession) HashInvalid(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.HashInvalid(&_ReceiptHarness.CallOpts, payload)
}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCallerSession) HashInvalid(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.HashInvalid(&_ReceiptHarness.CallOpts, payload)
}

// HashValid is a free data retrieval call binding the contract method 0x730dbf63.
//
// Solidity: function hashValid(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCaller) HashValid(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "hashValid", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashValid is a free data retrieval call binding the contract method 0x730dbf63.
//
// Solidity: function hashValid(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessSession) HashValid(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.HashValid(&_ReceiptHarness.CallOpts, payload)
}

// HashValid is a free data retrieval call binding the contract method 0x730dbf63.
//
// Solidity: function hashValid(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCallerSession) HashValid(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.HashValid(&_ReceiptHarness.CallOpts, payload)
}

// IsReceipt is a free data retrieval call binding the contract method 0x0bb3b580.
//
// Solidity: function isReceipt(bytes payload) pure returns(bool)
func (_ReceiptHarness *ReceiptHarnessCaller) IsReceipt(opts *bind.CallOpts, payload []byte) (bool, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "isReceipt", payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReceipt is a free data retrieval call binding the contract method 0x0bb3b580.
//
// Solidity: function isReceipt(bytes payload) pure returns(bool)
func (_ReceiptHarness *ReceiptHarnessSession) IsReceipt(payload []byte) (bool, error) {
	return _ReceiptHarness.Contract.IsReceipt(&_ReceiptHarness.CallOpts, payload)
}

// IsReceipt is a free data retrieval call binding the contract method 0x0bb3b580.
//
// Solidity: function isReceipt(bytes payload) pure returns(bool)
func (_ReceiptHarness *ReceiptHarnessCallerSession) IsReceipt(payload []byte) (bool, error) {
	return _ReceiptHarness.Contract.IsReceipt(&_ReceiptHarness.CallOpts, payload)
}

// MessageHash is a free data retrieval call binding the contract method 0xed54c3b6.
//
// Solidity: function messageHash(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCaller) MessageHash(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "messageHash", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MessageHash is a free data retrieval call binding the contract method 0xed54c3b6.
//
// Solidity: function messageHash(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessSession) MessageHash(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.MessageHash(&_ReceiptHarness.CallOpts, payload)
}

// MessageHash is a free data retrieval call binding the contract method 0xed54c3b6.
//
// Solidity: function messageHash(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCallerSession) MessageHash(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.MessageHash(&_ReceiptHarness.CallOpts, payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_ReceiptHarness *ReceiptHarnessCaller) Origin(opts *bind.CallOpts, payload []byte) (uint32, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "origin", payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_ReceiptHarness *ReceiptHarnessSession) Origin(payload []byte) (uint32, error) {
	return _ReceiptHarness.Contract.Origin(&_ReceiptHarness.CallOpts, payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes payload) pure returns(uint32)
func (_ReceiptHarness *ReceiptHarnessCallerSession) Origin(payload []byte) (uint32, error) {
	return _ReceiptHarness.Contract.Origin(&_ReceiptHarness.CallOpts, payload)
}

// SnapshotRoot is a free data retrieval call binding the contract method 0x854dfcd7.
//
// Solidity: function snapshotRoot(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCaller) SnapshotRoot(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "snapshotRoot", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SnapshotRoot is a free data retrieval call binding the contract method 0x854dfcd7.
//
// Solidity: function snapshotRoot(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessSession) SnapshotRoot(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.SnapshotRoot(&_ReceiptHarness.CallOpts, payload)
}

// SnapshotRoot is a free data retrieval call binding the contract method 0x854dfcd7.
//
// Solidity: function snapshotRoot(bytes payload) pure returns(bytes32)
func (_ReceiptHarness *ReceiptHarnessCallerSession) SnapshotRoot(payload []byte) ([32]byte, error) {
	return _ReceiptHarness.Contract.SnapshotRoot(&_ReceiptHarness.CallOpts, payload)
}

// StateIndex is a free data retrieval call binding the contract method 0x595271d1.
//
// Solidity: function stateIndex(bytes payload) pure returns(uint8)
func (_ReceiptHarness *ReceiptHarnessCaller) StateIndex(opts *bind.CallOpts, payload []byte) (uint8, error) {
	var out []interface{}
	err := _ReceiptHarness.contract.Call(opts, &out, "stateIndex", payload)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// StateIndex is a free data retrieval call binding the contract method 0x595271d1.
//
// Solidity: function stateIndex(bytes payload) pure returns(uint8)
func (_ReceiptHarness *ReceiptHarnessSession) StateIndex(payload []byte) (uint8, error) {
	return _ReceiptHarness.Contract.StateIndex(&_ReceiptHarness.CallOpts, payload)
}

// StateIndex is a free data retrieval call binding the contract method 0x595271d1.
//
// Solidity: function stateIndex(bytes payload) pure returns(uint8)
func (_ReceiptHarness *ReceiptHarnessCallerSession) StateIndex(payload []byte) (uint8, error) {
	return _ReceiptHarness.Contract.StateIndex(&_ReceiptHarness.CallOpts, payload)
}

// ReceiptLibMetaData contains all meta data concerning the ReceiptLib contract.
var ReceiptLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a50933beb11fb42dc6651c2cec100ea72eb83325f49b408f1b7033266650f01964736f6c63430008110033",
}

// ReceiptLibABI is the input ABI used to generate the binding from.
// Deprecated: Use ReceiptLibMetaData.ABI instead.
var ReceiptLibABI = ReceiptLibMetaData.ABI

// ReceiptLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReceiptLibMetaData.Bin instead.
var ReceiptLibBin = ReceiptLibMetaData.Bin

// DeployReceiptLib deploys a new Ethereum contract, binding an instance of ReceiptLib to it.
func DeployReceiptLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReceiptLib, error) {
	parsed, err := ReceiptLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReceiptLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReceiptLib{ReceiptLibCaller: ReceiptLibCaller{contract: contract}, ReceiptLibTransactor: ReceiptLibTransactor{contract: contract}, ReceiptLibFilterer: ReceiptLibFilterer{contract: contract}}, nil
}

// ReceiptLib is an auto generated Go binding around an Ethereum contract.
type ReceiptLib struct {
	ReceiptLibCaller     // Read-only binding to the contract
	ReceiptLibTransactor // Write-only binding to the contract
	ReceiptLibFilterer   // Log filterer for contract events
}

// ReceiptLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReceiptLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReceiptLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReceiptLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReceiptLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReceiptLibSession struct {
	Contract     *ReceiptLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReceiptLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReceiptLibCallerSession struct {
	Contract *ReceiptLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ReceiptLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReceiptLibTransactorSession struct {
	Contract     *ReceiptLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReceiptLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReceiptLibRaw struct {
	Contract *ReceiptLib // Generic contract binding to access the raw methods on
}

// ReceiptLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReceiptLibCallerRaw struct {
	Contract *ReceiptLibCaller // Generic read-only contract binding to access the raw methods on
}

// ReceiptLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReceiptLibTransactorRaw struct {
	Contract *ReceiptLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReceiptLib creates a new instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLib(address common.Address, backend bind.ContractBackend) (*ReceiptLib, error) {
	contract, err := bindReceiptLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReceiptLib{ReceiptLibCaller: ReceiptLibCaller{contract: contract}, ReceiptLibTransactor: ReceiptLibTransactor{contract: contract}, ReceiptLibFilterer: ReceiptLibFilterer{contract: contract}}, nil
}

// NewReceiptLibCaller creates a new read-only instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibCaller(address common.Address, caller bind.ContractCaller) (*ReceiptLibCaller, error) {
	contract, err := bindReceiptLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibCaller{contract: contract}, nil
}

// NewReceiptLibTransactor creates a new write-only instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibTransactor(address common.Address, transactor bind.ContractTransactor) (*ReceiptLibTransactor, error) {
	contract, err := bindReceiptLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibTransactor{contract: contract}, nil
}

// NewReceiptLibFilterer creates a new log filterer instance of ReceiptLib, bound to a specific deployed contract.
func NewReceiptLibFilterer(address common.Address, filterer bind.ContractFilterer) (*ReceiptLibFilterer, error) {
	contract, err := bindReceiptLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReceiptLibFilterer{contract: contract}, nil
}

// bindReceiptLib binds a generic wrapper to an already deployed contract.
func bindReceiptLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ReceiptLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptLib *ReceiptLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptLib.Contract.ReceiptLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptLib *ReceiptLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptLib.Contract.ReceiptLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptLib *ReceiptLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptLib.Contract.ReceiptLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReceiptLib *ReceiptLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReceiptLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReceiptLib *ReceiptLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReceiptLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReceiptLib *ReceiptLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReceiptLib.Contract.contract.Transact(opts, method, params...)
}
