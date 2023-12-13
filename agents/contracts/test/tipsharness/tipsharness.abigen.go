// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tipsharness

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

// TipsHarnessMetaData contains all meta data concerning the TipsHarness contract.
var TipsHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"TipsOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TipsValueTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"attestationTip\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"deliveryTip\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emptyTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"summitTip_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"attestationTip_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"executionTip_\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"deliveryTip_\",\"type\":\"uint64\"}],\"name\":\"encodeTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"summitTip_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attestationTip_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionTip_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deliveryTip_\",\"type\":\"uint256\"}],\"name\":\"encodeTips256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"executionTip\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Tips\",\"name\":\"tips\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"matchValue\",\"outputs\":[{\"internalType\":\"Tips\",\"name\":\"newTips\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"summitTip\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"value\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"paddedTips\",\"type\":\"uint256\"}],\"name\":\"wrapPadded\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0453e80e": "attestationTip(uint256)",
		"ecbf034e": "deliveryTip(uint256)",
		"725bd463": "emptyTips()",
		"4f2a6f9e": "encodeTips(uint64,uint64,uint64,uint64)",
		"72d3f3cd": "encodeTips256(uint256,uint256,uint256,uint256)",
		"4c63c701": "executionTip(uint256)",
		"f472a58a": "leaf(uint256)",
		"86450b88": "matchValue(uint256,uint256)",
		"b284b609": "summitTip(uint256)",
		"c5a46ee6": "value(uint256)",
		"138ac42f": "wrapPadded(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610579806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c806372d3f3cd11610081578063c5a46ee61161005b578063c5a46ee614610187578063ecbf034e1461019a578063f472a58a146101a857600080fd5b806372d3f3cd1461014e57806386450b8814610161578063b284b6091461017457600080fd5b80634c63c701116100b25780634c63c701146101205780634f2a6f9e14610133578063725bd4631461014657600080fd5b80630453e80e146100ce578063138ac42f146100ff575b600080fd5b6100e16100dc36600461042b565b6101bb565b60405167ffffffffffffffff90911681526020015b60405180910390f35b61011261010d36600461042b565b6101cd565b6040519081526020016100f6565b6100e161012e36600461042b565b6101d5565b610112610141366004610461565b6101e1565b61011261025a565b61011261015c3660046104b5565b610263565b61011261016f3660046104e7565b6102dc565b6100e161018236600461042b565b6102ef565b61011261019536600461042b565b6102fb565b6100e161010d36600461042b565b6101126101b636600461042b565b610306565b60006101c78260801c90565b92915050565b6000816101c7565b60006101c78260401c90565b60008067ffffffffffffffff8316604085901b6fffffffffffffffff000000000000000016608087901b77ffffffffffffffff000000000000000000000000000000001660c089901b7fffffffffffffffff000000000000000000000000000000000000000000000000161717175b9695505050505050565b600080806101c7565b60008067ffffffffffffffff602084811c919091169085901b6fffffffffffffffff000000000000000016606087901b77ffffffffffffffff000000000000000000000000000000001660a089901b7fffffffffffffffff00000000000000000000000000000000000000000000000016171717610250565b60006102e88383610314565b9392505050565b60006101c78260c01c90565b60006101c7826103bb565b6000818152602081206101c7565b600080610320846103bb565b90508083101561035c576040517f429726c100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b80830360201c67ffffffffffffffff8567ffffffffffffffff16820111156103b0576040517f1b438b3300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b939093019392505050565b60008167ffffffffffffffff166103d28360401c90565b67ffffffffffffffff166103e68460801c90565b67ffffffffffffffff166103fa8560c01c90565b67ffffffffffffffff1661040e9190610509565b6104189190610509565b6104229190610509565b60201b92915050565b60006020828403121561043d57600080fd5b5035919050565b803567ffffffffffffffff8116811461045c57600080fd5b919050565b6000806000806080858703121561047757600080fd5b61048085610444565b935061048e60208601610444565b925061049c60408601610444565b91506104aa60608601610444565b905092959194509250565b600080600080608085870312156104cb57600080fd5b5050823594602084013594506040840135936060013592509050565b600080604083850312156104fa57600080fd5b50508035926020909101359150565b808201808211156101c7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea26469706673582212209f0aa7bdc0739bafacea5e25b90d085026bc2fc02d70e140f11c1acdb137392b64736f6c63430008110033",
}

// TipsHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsHarnessMetaData.ABI instead.
var TipsHarnessABI = TipsHarnessMetaData.ABI

// Deprecated: Use TipsHarnessMetaData.Sigs instead.
// TipsHarnessFuncSigs maps the 4-byte function signature to its string representation.
var TipsHarnessFuncSigs = TipsHarnessMetaData.Sigs

// TipsHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsHarnessMetaData.Bin instead.
var TipsHarnessBin = TipsHarnessMetaData.Bin

// DeployTipsHarness deploys a new Ethereum contract, binding an instance of TipsHarness to it.
func DeployTipsHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TipsHarness, error) {
	parsed, err := TipsHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TipsHarness{TipsHarnessCaller: TipsHarnessCaller{contract: contract}, TipsHarnessTransactor: TipsHarnessTransactor{contract: contract}, TipsHarnessFilterer: TipsHarnessFilterer{contract: contract}}, nil
}

// TipsHarness is an auto generated Go binding around an Ethereum contract.
type TipsHarness struct {
	TipsHarnessCaller     // Read-only binding to the contract
	TipsHarnessTransactor // Write-only binding to the contract
	TipsHarnessFilterer   // Log filterer for contract events
}

// TipsHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsHarnessSession struct {
	Contract     *TipsHarness      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsHarnessCallerSession struct {
	Contract *TipsHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TipsHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsHarnessTransactorSession struct {
	Contract     *TipsHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TipsHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsHarnessRaw struct {
	Contract *TipsHarness // Generic contract binding to access the raw methods on
}

// TipsHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsHarnessCallerRaw struct {
	Contract *TipsHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// TipsHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsHarnessTransactorRaw struct {
	Contract *TipsHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipsHarness creates a new instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarness(address common.Address, backend bind.ContractBackend) (*TipsHarness, error) {
	contract, err := bindTipsHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TipsHarness{TipsHarnessCaller: TipsHarnessCaller{contract: contract}, TipsHarnessTransactor: TipsHarnessTransactor{contract: contract}, TipsHarnessFilterer: TipsHarnessFilterer{contract: contract}}, nil
}

// NewTipsHarnessCaller creates a new read-only instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarnessCaller(address common.Address, caller bind.ContractCaller) (*TipsHarnessCaller, error) {
	contract, err := bindTipsHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsHarnessCaller{contract: contract}, nil
}

// NewTipsHarnessTransactor creates a new write-only instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsHarnessTransactor, error) {
	contract, err := bindTipsHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsHarnessTransactor{contract: contract}, nil
}

// NewTipsHarnessFilterer creates a new log filterer instance of TipsHarness, bound to a specific deployed contract.
func NewTipsHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsHarnessFilterer, error) {
	contract, err := bindTipsHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsHarnessFilterer{contract: contract}, nil
}

// bindTipsHarness binds a generic wrapper to an already deployed contract.
func bindTipsHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsHarness *TipsHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsHarness.Contract.TipsHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsHarness *TipsHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsHarness.Contract.TipsHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsHarness *TipsHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsHarness.Contract.TipsHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsHarness *TipsHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsHarness *TipsHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsHarness *TipsHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsHarness.Contract.contract.Transact(opts, method, params...)
}

// AttestationTip is a free data retrieval call binding the contract method 0x0453e80e.
//
// Solidity: function attestationTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCaller) AttestationTip(opts *bind.CallOpts, paddedTips *big.Int) (uint64, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "attestationTip", paddedTips)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AttestationTip is a free data retrieval call binding the contract method 0x0453e80e.
//
// Solidity: function attestationTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessSession) AttestationTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.AttestationTip(&_TipsHarness.CallOpts, paddedTips)
}

// AttestationTip is a free data retrieval call binding the contract method 0x0453e80e.
//
// Solidity: function attestationTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCallerSession) AttestationTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.AttestationTip(&_TipsHarness.CallOpts, paddedTips)
}

// DeliveryTip is a free data retrieval call binding the contract method 0xecbf034e.
//
// Solidity: function deliveryTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCaller) DeliveryTip(opts *bind.CallOpts, paddedTips *big.Int) (uint64, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "deliveryTip", paddedTips)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DeliveryTip is a free data retrieval call binding the contract method 0xecbf034e.
//
// Solidity: function deliveryTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessSession) DeliveryTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.DeliveryTip(&_TipsHarness.CallOpts, paddedTips)
}

// DeliveryTip is a free data retrieval call binding the contract method 0xecbf034e.
//
// Solidity: function deliveryTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCallerSession) DeliveryTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.DeliveryTip(&_TipsHarness.CallOpts, paddedTips)
}

// EmptyTips is a free data retrieval call binding the contract method 0x725bd463.
//
// Solidity: function emptyTips() pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) EmptyTips(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "emptyTips")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmptyTips is a free data retrieval call binding the contract method 0x725bd463.
//
// Solidity: function emptyTips() pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) EmptyTips() (*big.Int, error) {
	return _TipsHarness.Contract.EmptyTips(&_TipsHarness.CallOpts)
}

// EmptyTips is a free data retrieval call binding the contract method 0x725bd463.
//
// Solidity: function emptyTips() pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) EmptyTips() (*big.Int, error) {
	return _TipsHarness.Contract.EmptyTips(&_TipsHarness.CallOpts)
}

// EncodeTips is a free data retrieval call binding the contract method 0x4f2a6f9e.
//
// Solidity: function encodeTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_) pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) EncodeTips(opts *bind.CallOpts, summitTip_ uint64, attestationTip_ uint64, executionTip_ uint64, deliveryTip_ uint64) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "encodeTips", summitTip_, attestationTip_, executionTip_, deliveryTip_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeTips is a free data retrieval call binding the contract method 0x4f2a6f9e.
//
// Solidity: function encodeTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_) pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) EncodeTips(summitTip_ uint64, attestationTip_ uint64, executionTip_ uint64, deliveryTip_ uint64) (*big.Int, error) {
	return _TipsHarness.Contract.EncodeTips(&_TipsHarness.CallOpts, summitTip_, attestationTip_, executionTip_, deliveryTip_)
}

// EncodeTips is a free data retrieval call binding the contract method 0x4f2a6f9e.
//
// Solidity: function encodeTips(uint64 summitTip_, uint64 attestationTip_, uint64 executionTip_, uint64 deliveryTip_) pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) EncodeTips(summitTip_ uint64, attestationTip_ uint64, executionTip_ uint64, deliveryTip_ uint64) (*big.Int, error) {
	return _TipsHarness.Contract.EncodeTips(&_TipsHarness.CallOpts, summitTip_, attestationTip_, executionTip_, deliveryTip_)
}

// EncodeTips256 is a free data retrieval call binding the contract method 0x72d3f3cd.
//
// Solidity: function encodeTips256(uint256 summitTip_, uint256 attestationTip_, uint256 executionTip_, uint256 deliveryTip_) pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) EncodeTips256(opts *bind.CallOpts, summitTip_ *big.Int, attestationTip_ *big.Int, executionTip_ *big.Int, deliveryTip_ *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "encodeTips256", summitTip_, attestationTip_, executionTip_, deliveryTip_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EncodeTips256 is a free data retrieval call binding the contract method 0x72d3f3cd.
//
// Solidity: function encodeTips256(uint256 summitTip_, uint256 attestationTip_, uint256 executionTip_, uint256 deliveryTip_) pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) EncodeTips256(summitTip_ *big.Int, attestationTip_ *big.Int, executionTip_ *big.Int, deliveryTip_ *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.EncodeTips256(&_TipsHarness.CallOpts, summitTip_, attestationTip_, executionTip_, deliveryTip_)
}

// EncodeTips256 is a free data retrieval call binding the contract method 0x72d3f3cd.
//
// Solidity: function encodeTips256(uint256 summitTip_, uint256 attestationTip_, uint256 executionTip_, uint256 deliveryTip_) pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) EncodeTips256(summitTip_ *big.Int, attestationTip_ *big.Int, executionTip_ *big.Int, deliveryTip_ *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.EncodeTips256(&_TipsHarness.CallOpts, summitTip_, attestationTip_, executionTip_, deliveryTip_)
}

// ExecutionTip is a free data retrieval call binding the contract method 0x4c63c701.
//
// Solidity: function executionTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCaller) ExecutionTip(opts *bind.CallOpts, paddedTips *big.Int) (uint64, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "executionTip", paddedTips)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExecutionTip is a free data retrieval call binding the contract method 0x4c63c701.
//
// Solidity: function executionTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessSession) ExecutionTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.ExecutionTip(&_TipsHarness.CallOpts, paddedTips)
}

// ExecutionTip is a free data retrieval call binding the contract method 0x4c63c701.
//
// Solidity: function executionTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCallerSession) ExecutionTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.ExecutionTip(&_TipsHarness.CallOpts, paddedTips)
}

// Leaf is a free data retrieval call binding the contract method 0xf472a58a.
//
// Solidity: function leaf(uint256 paddedTips) pure returns(bytes32)
func (_TipsHarness *TipsHarnessCaller) Leaf(opts *bind.CallOpts, paddedTips *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "leaf", paddedTips)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Leaf is a free data retrieval call binding the contract method 0xf472a58a.
//
// Solidity: function leaf(uint256 paddedTips) pure returns(bytes32)
func (_TipsHarness *TipsHarnessSession) Leaf(paddedTips *big.Int) ([32]byte, error) {
	return _TipsHarness.Contract.Leaf(&_TipsHarness.CallOpts, paddedTips)
}

// Leaf is a free data retrieval call binding the contract method 0xf472a58a.
//
// Solidity: function leaf(uint256 paddedTips) pure returns(bytes32)
func (_TipsHarness *TipsHarnessCallerSession) Leaf(paddedTips *big.Int) ([32]byte, error) {
	return _TipsHarness.Contract.Leaf(&_TipsHarness.CallOpts, paddedTips)
}

// MatchValue is a free data retrieval call binding the contract method 0x86450b88.
//
// Solidity: function matchValue(uint256 tips, uint256 newValue) pure returns(uint256 newTips)
func (_TipsHarness *TipsHarnessCaller) MatchValue(opts *bind.CallOpts, tips *big.Int, newValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "matchValue", tips, newValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MatchValue is a free data retrieval call binding the contract method 0x86450b88.
//
// Solidity: function matchValue(uint256 tips, uint256 newValue) pure returns(uint256 newTips)
func (_TipsHarness *TipsHarnessSession) MatchValue(tips *big.Int, newValue *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.MatchValue(&_TipsHarness.CallOpts, tips, newValue)
}

// MatchValue is a free data retrieval call binding the contract method 0x86450b88.
//
// Solidity: function matchValue(uint256 tips, uint256 newValue) pure returns(uint256 newTips)
func (_TipsHarness *TipsHarnessCallerSession) MatchValue(tips *big.Int, newValue *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.MatchValue(&_TipsHarness.CallOpts, tips, newValue)
}

// SummitTip is a free data retrieval call binding the contract method 0xb284b609.
//
// Solidity: function summitTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCaller) SummitTip(opts *bind.CallOpts, paddedTips *big.Int) (uint64, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "summitTip", paddedTips)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SummitTip is a free data retrieval call binding the contract method 0xb284b609.
//
// Solidity: function summitTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessSession) SummitTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.SummitTip(&_TipsHarness.CallOpts, paddedTips)
}

// SummitTip is a free data retrieval call binding the contract method 0xb284b609.
//
// Solidity: function summitTip(uint256 paddedTips) pure returns(uint64)
func (_TipsHarness *TipsHarnessCallerSession) SummitTip(paddedTips *big.Int) (uint64, error) {
	return _TipsHarness.Contract.SummitTip(&_TipsHarness.CallOpts, paddedTips)
}

// Value is a free data retrieval call binding the contract method 0xc5a46ee6.
//
// Solidity: function value(uint256 paddedTips) pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) Value(opts *bind.CallOpts, paddedTips *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "value", paddedTips)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0xc5a46ee6.
//
// Solidity: function value(uint256 paddedTips) pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) Value(paddedTips *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.Value(&_TipsHarness.CallOpts, paddedTips)
}

// Value is a free data retrieval call binding the contract method 0xc5a46ee6.
//
// Solidity: function value(uint256 paddedTips) pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) Value(paddedTips *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.Value(&_TipsHarness.CallOpts, paddedTips)
}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedTips) pure returns(uint256)
func (_TipsHarness *TipsHarnessCaller) WrapPadded(opts *bind.CallOpts, paddedTips *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TipsHarness.contract.Call(opts, &out, "wrapPadded", paddedTips)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedTips) pure returns(uint256)
func (_TipsHarness *TipsHarnessSession) WrapPadded(paddedTips *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.WrapPadded(&_TipsHarness.CallOpts, paddedTips)
}

// WrapPadded is a free data retrieval call binding the contract method 0x138ac42f.
//
// Solidity: function wrapPadded(uint256 paddedTips) pure returns(uint256)
func (_TipsHarness *TipsHarnessCallerSession) WrapPadded(paddedTips *big.Int) (*big.Int, error) {
	return _TipsHarness.Contract.WrapPadded(&_TipsHarness.CallOpts, paddedTips)
}

// TipsLibMetaData contains all meta data concerning the TipsLib contract.
var TipsLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d5578b5b4ad0066276edaecff5b95123aa98707d57007b8ed7c18e91e133678364736f6c63430008110033",
}

// TipsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use TipsLibMetaData.ABI instead.
var TipsLibABI = TipsLibMetaData.ABI

// TipsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TipsLibMetaData.Bin instead.
var TipsLibBin = TipsLibMetaData.Bin

// DeployTipsLib deploys a new Ethereum contract, binding an instance of TipsLib to it.
func DeployTipsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TipsLib, error) {
	parsed, err := TipsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TipsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// TipsLib is an auto generated Go binding around an Ethereum contract.
type TipsLib struct {
	TipsLibCaller     // Read-only binding to the contract
	TipsLibTransactor // Write-only binding to the contract
	TipsLibFilterer   // Log filterer for contract events
}

// TipsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type TipsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TipsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TipsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TipsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TipsLibSession struct {
	Contract     *TipsLib          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TipsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TipsLibCallerSession struct {
	Contract *TipsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TipsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TipsLibTransactorSession struct {
	Contract     *TipsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TipsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type TipsLibRaw struct {
	Contract *TipsLib // Generic contract binding to access the raw methods on
}

// TipsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TipsLibCallerRaw struct {
	Contract *TipsLibCaller // Generic read-only contract binding to access the raw methods on
}

// TipsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TipsLibTransactorRaw struct {
	Contract *TipsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTipsLib creates a new instance of TipsLib, bound to a specific deployed contract.
func NewTipsLib(address common.Address, backend bind.ContractBackend) (*TipsLib, error) {
	contract, err := bindTipsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TipsLib{TipsLibCaller: TipsLibCaller{contract: contract}, TipsLibTransactor: TipsLibTransactor{contract: contract}, TipsLibFilterer: TipsLibFilterer{contract: contract}}, nil
}

// NewTipsLibCaller creates a new read-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibCaller(address common.Address, caller bind.ContractCaller) (*TipsLibCaller, error) {
	contract, err := bindTipsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibCaller{contract: contract}, nil
}

// NewTipsLibTransactor creates a new write-only instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*TipsLibTransactor, error) {
	contract, err := bindTipsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TipsLibTransactor{contract: contract}, nil
}

// NewTipsLibFilterer creates a new log filterer instance of TipsLib, bound to a specific deployed contract.
func NewTipsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*TipsLibFilterer, error) {
	contract, err := bindTipsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TipsLibFilterer{contract: contract}, nil
}

// bindTipsLib binds a generic wrapper to an already deployed contract.
func bindTipsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TipsLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.TipsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.TipsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TipsLib *TipsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TipsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TipsLib *TipsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TipsLib *TipsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TipsLib.Contract.contract.Transact(opts, method, params...)
}
