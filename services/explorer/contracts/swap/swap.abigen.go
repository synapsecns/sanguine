// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package swap

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212206d1ff01848a4a56d97f886756bad8b4fa638c4808ea712feaa090e069f5dbd6364736f6c634300060c0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220da3544b6a980726371ab0dc3244fd70790850284d4442daa7cfe73d080ad13de64736f6c634300060c0033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// AmplificationUtilsMetaData contains all meta data concerning the AmplificationUtils contract.
var AmplificationUtilsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"futureTime\",\"type\":\"uint256\"}],\"name\":\"RampA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StopRampA\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"A_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_A\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"d011f918": "A_PRECISION()",
		"39698415": "MAX_A()",
		"b0a14cfc": "getA(SwapUtils.Swap storage)",
		"c9b64dcb": "getAPrecise(SwapUtils.Swap storage)",
		"58fdd79b": "rampA(SwapUtils.Swap storage,uint256,uint256)",
		"f14e211e": "stopRampA(SwapUtils.Swap storage)",
	},
	Bin: "0x61081f610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061007c5760003560e01c8063c9b64dcb1161005a578063c9b64dcb146100f0578063d011f9181461010d578063f14e211e146101155761007c565b8063396984151461008157806358fdd79b1461009b578063b0a14cfc146100d3575b600080fd5b61008961013f565b60408051918252519081900360200190f35b8180156100a757600080fd5b506100d1600480360360608110156100be57600080fd5b5080359060208101359060400135610146565b005b610089600480360360208110156100e957600080fd5b5035610440565b6100896004803603602081101561010657600080fd5b503561045e565b610089610469565b81801561012157600080fd5b506100d16004803603602081101561013857600080fd5b503561046e565b620f424081565b60028301546101589062015180610545565b4210156101c657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f57616974203120646179206265666f7265207374617274696e672072616d7000604482015290519081900360640190fd5b6101d34262093a80610545565b81101561024157604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f496e73756666696369656e742072616d702074696d6500000000000000000000604482015290519081900360640190fd5b6000821180156102535750620f424082105b6102be57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f667574757265415f206d757374206265203e203020616e64203c204d41585f41604482015290519081900360640190fd5b60006102c9846105c2565b905060006102d884606461065d565b90508181101561036057816102ee82600261065d565b101561035b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f667574757265415f20697320746f6f20736d616c6c0000000000000000000000604482015290519081900360640190fd5b6103d9565b61036b82600261065d565b8111156103d957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f667574757265415f20697320746f6f206c617267650000000000000000000000604482015290519081900360640190fd5b8185556001850181905542600286018190556003860184905560408051848152602081018490528082019290925260608201859052517fa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c2549181900360800190a15050505050565b60006104566064610450846105c2565b906106d0565b90505b919050565b6000610456826105c2565b606481565b428160030154116104e057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f52616d7020697320616c72656164792073746f70706564000000000000000000604482015290519081900360640190fd5b60006104eb826105c2565b8083556001830181905542600284018190556003840181905560408051838152602081019290925280519293507f46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc20193892918290030190a15050565b6000828201838110156105b957604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b90505b92915050565b6003810154600182015460009190428211156106545760028401548454808311156106265761061b6106146105f78685610751565b6104506106044287610751565b61060e8887610751565b9061065d565b8290610545565b945050505050610459565b61061b61064d6106368685610751565b6104506106434287610751565b61060e8689610751565b8290610751565b91506104599050565b60008261066c575060006105bc565b8282028284828161067957fe5b04146105b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806107c96021913960400191505060405180910390fd5b600080821161074057604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b81838161074957fe5b049392505050565b6000828211156107c257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b5090039056fe536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a2646970667358221220c114667319e861dae68c50014502f4137a7fb296420876a9c950c35942f2b83764736f6c634300060c0033",
}

// AmplificationUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use AmplificationUtilsMetaData.ABI instead.
var AmplificationUtilsABI = AmplificationUtilsMetaData.ABI

// Deprecated: Use AmplificationUtilsMetaData.Sigs instead.
// AmplificationUtilsFuncSigs maps the 4-byte function signature to its string representation.
var AmplificationUtilsFuncSigs = AmplificationUtilsMetaData.Sigs

// AmplificationUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AmplificationUtilsMetaData.Bin instead.
var AmplificationUtilsBin = AmplificationUtilsMetaData.Bin

// DeployAmplificationUtils deploys a new Ethereum contract, binding an instance of AmplificationUtils to it.
func DeployAmplificationUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AmplificationUtils, error) {
	parsed, err := AmplificationUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AmplificationUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AmplificationUtils{AmplificationUtilsCaller: AmplificationUtilsCaller{contract: contract}, AmplificationUtilsTransactor: AmplificationUtilsTransactor{contract: contract}, AmplificationUtilsFilterer: AmplificationUtilsFilterer{contract: contract}}, nil
}

// AmplificationUtils is an auto generated Go binding around an Ethereum contract.
type AmplificationUtils struct {
	AmplificationUtilsCaller     // Read-only binding to the contract
	AmplificationUtilsTransactor // Write-only binding to the contract
	AmplificationUtilsFilterer   // Log filterer for contract events
}

// AmplificationUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AmplificationUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmplificationUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AmplificationUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmplificationUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AmplificationUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AmplificationUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AmplificationUtilsSession struct {
	Contract     *AmplificationUtils // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AmplificationUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AmplificationUtilsCallerSession struct {
	Contract *AmplificationUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AmplificationUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AmplificationUtilsTransactorSession struct {
	Contract     *AmplificationUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AmplificationUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AmplificationUtilsRaw struct {
	Contract *AmplificationUtils // Generic contract binding to access the raw methods on
}

// AmplificationUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AmplificationUtilsCallerRaw struct {
	Contract *AmplificationUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// AmplificationUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AmplificationUtilsTransactorRaw struct {
	Contract *AmplificationUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAmplificationUtils creates a new instance of AmplificationUtils, bound to a specific deployed contract.
func NewAmplificationUtils(address common.Address, backend bind.ContractBackend) (*AmplificationUtils, error) {
	contract, err := bindAmplificationUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AmplificationUtils{AmplificationUtilsCaller: AmplificationUtilsCaller{contract: contract}, AmplificationUtilsTransactor: AmplificationUtilsTransactor{contract: contract}, AmplificationUtilsFilterer: AmplificationUtilsFilterer{contract: contract}}, nil
}

// NewAmplificationUtilsCaller creates a new read-only instance of AmplificationUtils, bound to a specific deployed contract.
func NewAmplificationUtilsCaller(address common.Address, caller bind.ContractCaller) (*AmplificationUtilsCaller, error) {
	contract, err := bindAmplificationUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AmplificationUtilsCaller{contract: contract}, nil
}

// NewAmplificationUtilsTransactor creates a new write-only instance of AmplificationUtils, bound to a specific deployed contract.
func NewAmplificationUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*AmplificationUtilsTransactor, error) {
	contract, err := bindAmplificationUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AmplificationUtilsTransactor{contract: contract}, nil
}

// NewAmplificationUtilsFilterer creates a new log filterer instance of AmplificationUtils, bound to a specific deployed contract.
func NewAmplificationUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*AmplificationUtilsFilterer, error) {
	contract, err := bindAmplificationUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AmplificationUtilsFilterer{contract: contract}, nil
}

// bindAmplificationUtils binds a generic wrapper to an already deployed contract.
func bindAmplificationUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AmplificationUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmplificationUtils *AmplificationUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmplificationUtils.Contract.AmplificationUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmplificationUtils *AmplificationUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmplificationUtils.Contract.AmplificationUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmplificationUtils *AmplificationUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmplificationUtils.Contract.AmplificationUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AmplificationUtils *AmplificationUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AmplificationUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AmplificationUtils *AmplificationUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AmplificationUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AmplificationUtils *AmplificationUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AmplificationUtils.Contract.contract.Transact(opts, method, params...)
}

// APRECISION is a free data retrieval call binding the contract method 0xd011f918.
//
// Solidity: function A_PRECISION() view returns(uint256)
func (_AmplificationUtils *AmplificationUtilsCaller) APRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmplificationUtils.contract.Call(opts, &out, "A_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// APRECISION is a free data retrieval call binding the contract method 0xd011f918.
//
// Solidity: function A_PRECISION() view returns(uint256)
func (_AmplificationUtils *AmplificationUtilsSession) APRECISION() (*big.Int, error) {
	return _AmplificationUtils.Contract.APRECISION(&_AmplificationUtils.CallOpts)
}

// APRECISION is a free data retrieval call binding the contract method 0xd011f918.
//
// Solidity: function A_PRECISION() view returns(uint256)
func (_AmplificationUtils *AmplificationUtilsCallerSession) APRECISION() (*big.Int, error) {
	return _AmplificationUtils.Contract.APRECISION(&_AmplificationUtils.CallOpts)
}

// MAXA is a free data retrieval call binding the contract method 0x39698415.
//
// Solidity: function MAX_A() view returns(uint256)
func (_AmplificationUtils *AmplificationUtilsCaller) MAXA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AmplificationUtils.contract.Call(opts, &out, "MAX_A")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXA is a free data retrieval call binding the contract method 0x39698415.
//
// Solidity: function MAX_A() view returns(uint256)
func (_AmplificationUtils *AmplificationUtilsSession) MAXA() (*big.Int, error) {
	return _AmplificationUtils.Contract.MAXA(&_AmplificationUtils.CallOpts)
}

// MAXA is a free data retrieval call binding the contract method 0x39698415.
//
// Solidity: function MAX_A() view returns(uint256)
func (_AmplificationUtils *AmplificationUtilsCallerSession) MAXA() (*big.Int, error) {
	return _AmplificationUtils.Contract.MAXA(&_AmplificationUtils.CallOpts)
}

// AmplificationUtilsRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the AmplificationUtils contract.
type AmplificationUtilsRampAIterator struct {
	Event *AmplificationUtilsRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AmplificationUtilsRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmplificationUtilsRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AmplificationUtilsRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AmplificationUtilsRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmplificationUtilsRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmplificationUtilsRampA represents a RampA event raised by the AmplificationUtils contract.
type AmplificationUtilsRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_AmplificationUtils *AmplificationUtilsFilterer) FilterRampA(opts *bind.FilterOpts) (*AmplificationUtilsRampAIterator, error) {

	logs, sub, err := _AmplificationUtils.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &AmplificationUtilsRampAIterator{contract: _AmplificationUtils.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_AmplificationUtils *AmplificationUtilsFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *AmplificationUtilsRampA) (event.Subscription, error) {

	logs, sub, err := _AmplificationUtils.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmplificationUtilsRampA)
				if err := _AmplificationUtils.contract.UnpackLog(event, "RampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_AmplificationUtils *AmplificationUtilsFilterer) ParseRampA(log types.Log) (*AmplificationUtilsRampA, error) {
	event := new(AmplificationUtilsRampA)
	if err := _AmplificationUtils.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AmplificationUtilsStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the AmplificationUtils contract.
type AmplificationUtilsStopRampAIterator struct {
	Event *AmplificationUtilsStopRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AmplificationUtilsStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AmplificationUtilsStopRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AmplificationUtilsStopRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AmplificationUtilsStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AmplificationUtilsStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AmplificationUtilsStopRampA represents a StopRampA event raised by the AmplificationUtils contract.
type AmplificationUtilsStopRampA struct {
	CurrentA *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_AmplificationUtils *AmplificationUtilsFilterer) FilterStopRampA(opts *bind.FilterOpts) (*AmplificationUtilsStopRampAIterator, error) {

	logs, sub, err := _AmplificationUtils.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &AmplificationUtilsStopRampAIterator{contract: _AmplificationUtils.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_AmplificationUtils *AmplificationUtilsFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *AmplificationUtilsStopRampA) (event.Subscription, error) {

	logs, sub, err := _AmplificationUtils.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AmplificationUtilsStopRampA)
				if err := _AmplificationUtils.contract.UnpackLog(event, "StopRampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_AmplificationUtils *AmplificationUtilsFilterer) ParseStopRampA(log types.Log) (*AmplificationUtilsStopRampA, error) {
	event := new(AmplificationUtilsStopRampA)
	if err := _AmplificationUtils.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClonesMetaData contains all meta data concerning the Clones contract.
var ClonesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220605e0778555ecf862c6fa925f1fb4c3526175e69fa09aedca9251f2f192a904c64736f6c634300060c0033",
}

// ClonesABI is the input ABI used to generate the binding from.
// Deprecated: Use ClonesMetaData.ABI instead.
var ClonesABI = ClonesMetaData.ABI

// ClonesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClonesMetaData.Bin instead.
var ClonesBin = ClonesMetaData.Bin

// DeployClones deploys a new Ethereum contract, binding an instance of Clones to it.
func DeployClones(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Clones, error) {
	parsed, err := ClonesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClonesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// Clones is an auto generated Go binding around an Ethereum contract.
type Clones struct {
	ClonesCaller     // Read-only binding to the contract
	ClonesTransactor // Write-only binding to the contract
	ClonesFilterer   // Log filterer for contract events
}

// ClonesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClonesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClonesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClonesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClonesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClonesSession struct {
	Contract     *Clones           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClonesCallerSession struct {
	Contract *ClonesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClonesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClonesTransactorSession struct {
	Contract     *ClonesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClonesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClonesRaw struct {
	Contract *Clones // Generic contract binding to access the raw methods on
}

// ClonesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClonesCallerRaw struct {
	Contract *ClonesCaller // Generic read-only contract binding to access the raw methods on
}

// ClonesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClonesTransactorRaw struct {
	Contract *ClonesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClones creates a new instance of Clones, bound to a specific deployed contract.
func NewClones(address common.Address, backend bind.ContractBackend) (*Clones, error) {
	contract, err := bindClones(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Clones{ClonesCaller: ClonesCaller{contract: contract}, ClonesTransactor: ClonesTransactor{contract: contract}, ClonesFilterer: ClonesFilterer{contract: contract}}, nil
}

// NewClonesCaller creates a new read-only instance of Clones, bound to a specific deployed contract.
func NewClonesCaller(address common.Address, caller bind.ContractCaller) (*ClonesCaller, error) {
	contract, err := bindClones(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesCaller{contract: contract}, nil
}

// NewClonesTransactor creates a new write-only instance of Clones, bound to a specific deployed contract.
func NewClonesTransactor(address common.Address, transactor bind.ContractTransactor) (*ClonesTransactor, error) {
	contract, err := bindClones(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClonesTransactor{contract: contract}, nil
}

// NewClonesFilterer creates a new log filterer instance of Clones, bound to a specific deployed contract.
func NewClonesFilterer(address common.Address, filterer bind.ContractFilterer) (*ClonesFilterer, error) {
	contract, err := bindClones(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClonesFilterer{contract: contract}, nil
}

// bindClones binds a generic wrapper to an already deployed contract.
func bindClones(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClonesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.ClonesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.ClonesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Clones *ClonesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Clones.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Clones *ClonesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Clones *ClonesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Clones.Contract.contract.Transact(opts, method, params...)
}

// ContextMetaData contains all meta data concerning the Context contract.
var ContextMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextMetaData.ABI instead.
var ContextABI = ContextMetaData.ABI

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC20MetaData contains all meta data concerning the ERC20 contract.
var ERC20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e8738038062000e87833981810160405260408110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b5060405250508251620001b491506003906020850190620001e0565b508051620001ca906004906020840190620001e0565b50506005805460ff19166012179055506200027c565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022357805160ff191683800117855562000253565b8280016001018555821562000253579182015b828111156200025357825182559160200191906001019062000236565b506200026192915062000265565b5090565b5b8082111562000261576000815560010162000266565b610bfb806200028c6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610287578063a9059cbb146102c0578063dd62ed3e146102f9576100c9565b8063395093511461021357806370a082311461024c57806395d89b411461027f576100c9565b806318160ddd116100b257806318160ddd1461019857806323b872dd146101b2578063313ce567146101f5576100c9565b806306fdde03146100ce578063095ea7b31461014b575b600080fd5b6100d6610334565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101105781810151838201526020016100f8565b50505050905090810190601f16801561013d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101846004803603604081101561016157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356103e8565b604080519115158252519081900360200190f35b6101a0610405565b60408051918252519081900360200190f35b610184600480360360608110156101c857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135909116906040013561040b565b6101fd6104ac565b6040805160ff9092168252519081900360200190f35b6101846004803603604081101561022957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356104b5565b6101a06004803603602081101561026257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610510565b6100d6610538565b6101846004803603604081101561029d57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105b7565b610184600480360360408110156102d657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561062c565b6101a06004803603604081101561030f57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610640565b60038054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103de5780601f106103b3576101008083540402835291602001916103de565b820191906000526020600020905b8154815290600101906020018083116103c157829003601f168201915b5050505050905090565b60006103fc6103f5610678565b848461067c565b50600192915050565b60025490565b60006104188484846107c3565b6104a284610424610678565b61049d85604051806060016040528060288152602001610b306028913973ffffffffffffffffffffffffffffffffffffffff8a1660009081526001602052604081209061046f610678565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600020549190610993565b61067c565b5060019392505050565b60055460ff1690565b60006103fc6104c2610678565b8461049d85600160006104d3610678565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918c168152925290205490610a44565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b60048054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103de5780601f106103b3576101008083540402835291602001916103de565b60006103fc6105c4610678565b8461049d85604051806060016040528060258152602001610ba160259139600160006105ee610678565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918d16815292529020549190610993565b60006103fc610639610678565b84846107c3565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b3390565b73ffffffffffffffffffffffffffffffffffffffff83166106e8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610b7d6024913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610754576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610ae86022913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661082f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610b586025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661089b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180610ac56023913960400191505060405180910390fd5b6108a6838383610abf565b6108f081604051806060016040528060268152602001610b0a6026913973ffffffffffffffffffffffffffffffffffffffff86166000908152602081905260409020549190610993565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220939093559084168152205461092c9082610a44565b73ffffffffffffffffffffffffffffffffffffffff8084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610a3c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a015781810151838201526020016109e9565b50505050905090810190601f168015610a2e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610ab857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220231fb9c34a106dcbd4e72506478a058fb1814b86591c60e2e9911ad5d2d9da9a64736f6c634300060c0033",
}

// ERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MetaData.ABI instead.
var ERC20ABI = ERC20MetaData.ABI

// Deprecated: Use ERC20MetaData.Sigs instead.
// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = ERC20MetaData.Sigs

// ERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MetaData.Bin instead.
var ERC20Bin = ERC20MetaData.Bin

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := ERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20Bin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableUpgradeableMetaData contains all meta data concerning the ERC20BurnableUpgradeable contract.
var ERC20BurnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"42966c68": "burn(uint256)",
		"79cc6790": "burnFrom(address,uint256)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// ERC20BurnableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20BurnableUpgradeableMetaData.ABI instead.
var ERC20BurnableUpgradeableABI = ERC20BurnableUpgradeableMetaData.ABI

// Deprecated: Use ERC20BurnableUpgradeableMetaData.Sigs instead.
// ERC20BurnableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BurnableUpgradeableFuncSigs = ERC20BurnableUpgradeableMetaData.Sigs

// ERC20BurnableUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC20BurnableUpgradeable struct {
	ERC20BurnableUpgradeableCaller     // Read-only binding to the contract
	ERC20BurnableUpgradeableTransactor // Write-only binding to the contract
	ERC20BurnableUpgradeableFilterer   // Log filterer for contract events
}

// ERC20BurnableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BurnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BurnableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BurnableUpgradeableSession struct {
	Contract     *ERC20BurnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC20BurnableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BurnableUpgradeableCallerSession struct {
	Contract *ERC20BurnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ERC20BurnableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BurnableUpgradeableTransactorSession struct {
	Contract     *ERC20BurnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ERC20BurnableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BurnableUpgradeableRaw struct {
	Contract *ERC20BurnableUpgradeable // Generic contract binding to access the raw methods on
}

// ERC20BurnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeableCallerRaw struct {
	Contract *ERC20BurnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BurnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BurnableUpgradeableTransactorRaw struct {
	Contract *ERC20BurnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20BurnableUpgradeable creates a new instance of ERC20BurnableUpgradeable, bound to a specific deployed contract.
func NewERC20BurnableUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC20BurnableUpgradeable, error) {
	contract, err := bindERC20BurnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeable{ERC20BurnableUpgradeableCaller: ERC20BurnableUpgradeableCaller{contract: contract}, ERC20BurnableUpgradeableTransactor: ERC20BurnableUpgradeableTransactor{contract: contract}, ERC20BurnableUpgradeableFilterer: ERC20BurnableUpgradeableFilterer{contract: contract}}, nil
}

// NewERC20BurnableUpgradeableCaller creates a new read-only instance of ERC20BurnableUpgradeable, bound to a specific deployed contract.
func NewERC20BurnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC20BurnableUpgradeableCaller, error) {
	contract, err := bindERC20BurnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeableCaller{contract: contract}, nil
}

// NewERC20BurnableUpgradeableTransactor creates a new write-only instance of ERC20BurnableUpgradeable, bound to a specific deployed contract.
func NewERC20BurnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BurnableUpgradeableTransactor, error) {
	contract, err := bindERC20BurnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeableTransactor{contract: contract}, nil
}

// NewERC20BurnableUpgradeableFilterer creates a new log filterer instance of ERC20BurnableUpgradeable, bound to a specific deployed contract.
func NewERC20BurnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BurnableUpgradeableFilterer, error) {
	contract, err := bindERC20BurnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeableFilterer{contract: contract}, nil
}

// bindERC20BurnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC20BurnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BurnableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20BurnableUpgradeable.Contract.ERC20BurnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.ERC20BurnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.ERC20BurnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20BurnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeable.Contract.Allowance(&_ERC20BurnableUpgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeable.Contract.Allowance(&_ERC20BurnableUpgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeable.Contract.BalanceOf(&_ERC20BurnableUpgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20BurnableUpgradeable.Contract.BalanceOf(&_ERC20BurnableUpgradeable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) Decimals() (uint8, error) {
	return _ERC20BurnableUpgradeable.Contract.Decimals(&_ERC20BurnableUpgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCallerSession) Decimals() (uint8, error) {
	return _ERC20BurnableUpgradeable.Contract.Decimals(&_ERC20BurnableUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) Name() (string, error) {
	return _ERC20BurnableUpgradeable.Contract.Name(&_ERC20BurnableUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCallerSession) Name() (string, error) {
	return _ERC20BurnableUpgradeable.Contract.Name(&_ERC20BurnableUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) Symbol() (string, error) {
	return _ERC20BurnableUpgradeable.Contract.Symbol(&_ERC20BurnableUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCallerSession) Symbol() (string, error) {
	return _ERC20BurnableUpgradeable.Contract.Symbol(&_ERC20BurnableUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20BurnableUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _ERC20BurnableUpgradeable.Contract.TotalSupply(&_ERC20BurnableUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20BurnableUpgradeable.Contract.TotalSupply(&_ERC20BurnableUpgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.Approve(&_ERC20BurnableUpgradeable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.Approve(&_ERC20BurnableUpgradeable.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.Burn(&_ERC20BurnableUpgradeable.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.Burn(&_ERC20BurnableUpgradeable.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.BurnFrom(&_ERC20BurnableUpgradeable.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.BurnFrom(&_ERC20BurnableUpgradeable.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.DecreaseAllowance(&_ERC20BurnableUpgradeable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.DecreaseAllowance(&_ERC20BurnableUpgradeable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.IncreaseAllowance(&_ERC20BurnableUpgradeable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.IncreaseAllowance(&_ERC20BurnableUpgradeable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.Transfer(&_ERC20BurnableUpgradeable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.Transfer(&_ERC20BurnableUpgradeable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.TransferFrom(&_ERC20BurnableUpgradeable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20BurnableUpgradeable.Contract.TransferFrom(&_ERC20BurnableUpgradeable.TransactOpts, sender, recipient, amount)
}

// ERC20BurnableUpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20BurnableUpgradeable contract.
type ERC20BurnableUpgradeableApprovalIterator struct {
	Event *ERC20BurnableUpgradeableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableUpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableUpgradeableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableUpgradeableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableUpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableUpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableUpgradeableApproval represents a Approval event raised by the ERC20BurnableUpgradeable contract.
type ERC20BurnableUpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20BurnableUpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeableApprovalIterator{contract: _ERC20BurnableUpgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20BurnableUpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableUpgradeableApproval)
				if err := _ERC20BurnableUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableFilterer) ParseApproval(log types.Log) (*ERC20BurnableUpgradeableApproval, error) {
	event := new(ERC20BurnableUpgradeableApproval)
	if err := _ERC20BurnableUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20BurnableUpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20BurnableUpgradeable contract.
type ERC20BurnableUpgradeableTransferIterator struct {
	Event *ERC20BurnableUpgradeableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20BurnableUpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BurnableUpgradeableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20BurnableUpgradeableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20BurnableUpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BurnableUpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BurnableUpgradeableTransfer represents a Transfer event raised by the ERC20BurnableUpgradeable contract.
type ERC20BurnableUpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BurnableUpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BurnableUpgradeableTransferIterator{contract: _ERC20BurnableUpgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BurnableUpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20BurnableUpgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BurnableUpgradeableTransfer)
				if err := _ERC20BurnableUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20BurnableUpgradeable *ERC20BurnableUpgradeableFilterer) ParseTransfer(log types.Log) (*ERC20BurnableUpgradeableTransfer, error) {
	event := new(ERC20BurnableUpgradeableTransfer)
	if err := _ERC20BurnableUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UpgradeableMetaData contains all meta data concerning the ERC20Upgradeable contract.
var ERC20UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610bfd806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610287578063a9059cbb146102c0578063dd62ed3e146102f9576100c9565b8063395093511461021357806370a082311461024c57806395d89b411461027f576100c9565b806318160ddd116100b257806318160ddd1461019857806323b872dd146101b2578063313ce567146101f5576100c9565b806306fdde03146100ce578063095ea7b31461014b575b600080fd5b6100d6610334565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101105781810151838201526020016100f8565b50505050905090810190601f16801561013d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101846004803603604081101561016157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356103e8565b604080519115158252519081900360200190f35b6101a0610405565b60408051918252519081900360200190f35b610184600480360360608110156101c857600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020810135909116906040013561040b565b6101fd6104ac565b6040805160ff9092168252519081900360200190f35b6101846004803603604081101561022957600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356104b5565b6101a06004803603602081101561026257600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610510565b6100d6610538565b6101846004803603604081101561029d57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356105b7565b610184600480360360408110156102d657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561062c565b6101a06004803603604081101561030f57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610640565b60368054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103de5780601f106103b3576101008083540402835291602001916103de565b820191906000526020600020905b8154815290600101906020018083116103c157829003601f168201915b5050505050905090565b60006103fc6103f5610678565b848461067c565b50600192915050565b60355490565b60006104188484846107c3565b6104a284610424610678565b61049d85604051806060016040528060288152602001610b326028913973ffffffffffffffffffffffffffffffffffffffff8a1660009081526034602052604081209061046f610678565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600020549190610995565b61067c565b5060019392505050565b60385460ff1690565b60006103fc6104c2610678565b8461049d85603460006104d3610678565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918c168152925290205490610a46565b73ffffffffffffffffffffffffffffffffffffffff1660009081526033602052604090205490565b60378054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156103de5780601f106103b3576101008083540402835291602001916103de565b60006103fc6105c4610678565b8461049d85604051806060016040528060258152602001610ba360259139603460006105ee610678565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918d16815292529020549190610995565b60006103fc610639610678565b84846107c3565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260346020908152604080832093909416825291909152205490565b3390565b73ffffffffffffffffffffffffffffffffffffffff83166106e8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610b7f6024913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610754576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610aea6022913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff808416600081815260346020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b73ffffffffffffffffffffffffffffffffffffffff831661082f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526025815260200180610b5a6025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821661089b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180610ac76023913960400191505060405180910390fd5b6108a6838383610ac1565b6108f081604051806060016040528060268152602001610b0c6026913973ffffffffffffffffffffffffffffffffffffffff86166000908152603360205260409020549190610995565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260336020526040808220939093559084168152205461092c9082610a46565b73ffffffffffffffffffffffffffffffffffffffff80841660008181526033602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610a3e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610a035781810151838201526020016109eb565b50505050905090810190601f168015610a305780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610aba57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220988f03bd457a2310a250b390e92c42a6c260dba9692eb57db616b7791ad5d9bb64736f6c634300060c0033",
}

// ERC20UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20UpgradeableMetaData.ABI instead.
var ERC20UpgradeableABI = ERC20UpgradeableMetaData.ABI

// Deprecated: Use ERC20UpgradeableMetaData.Sigs instead.
// ERC20UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20UpgradeableFuncSigs = ERC20UpgradeableMetaData.Sigs

// ERC20UpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20UpgradeableMetaData.Bin instead.
var ERC20UpgradeableBin = ERC20UpgradeableMetaData.Bin

// DeployERC20Upgradeable deploys a new Ethereum contract, binding an instance of ERC20Upgradeable to it.
func DeployERC20Upgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Upgradeable, error) {
	parsed, err := ERC20UpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20UpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Upgradeable{ERC20UpgradeableCaller: ERC20UpgradeableCaller{contract: contract}, ERC20UpgradeableTransactor: ERC20UpgradeableTransactor{contract: contract}, ERC20UpgradeableFilterer: ERC20UpgradeableFilterer{contract: contract}}, nil
}

// ERC20Upgradeable is an auto generated Go binding around an Ethereum contract.
type ERC20Upgradeable struct {
	ERC20UpgradeableCaller     // Read-only binding to the contract
	ERC20UpgradeableTransactor // Write-only binding to the contract
	ERC20UpgradeableFilterer   // Log filterer for contract events
}

// ERC20UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20UpgradeableSession struct {
	Contract     *ERC20Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20UpgradeableCallerSession struct {
	Contract *ERC20UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20UpgradeableTransactorSession struct {
	Contract     *ERC20UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20UpgradeableRaw struct {
	Contract *ERC20Upgradeable // Generic contract binding to access the raw methods on
}

// ERC20UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20UpgradeableCallerRaw struct {
	Contract *ERC20UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20UpgradeableTransactorRaw struct {
	Contract *ERC20UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Upgradeable creates a new instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20Upgradeable(address common.Address, backend bind.ContractBackend) (*ERC20Upgradeable, error) {
	contract, err := bindERC20Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Upgradeable{ERC20UpgradeableCaller: ERC20UpgradeableCaller{contract: contract}, ERC20UpgradeableTransactor: ERC20UpgradeableTransactor{contract: contract}, ERC20UpgradeableFilterer: ERC20UpgradeableFilterer{contract: contract}}, nil
}

// NewERC20UpgradeableCaller creates a new read-only instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC20UpgradeableCaller, error) {
	contract, err := bindERC20Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableCaller{contract: contract}, nil
}

// NewERC20UpgradeableTransactor creates a new write-only instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20UpgradeableTransactor, error) {
	contract, err := bindERC20Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableTransactor{contract: contract}, nil
}

// NewERC20UpgradeableFilterer creates a new log filterer instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20UpgradeableFilterer, error) {
	contract, err := bindERC20Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableFilterer{contract: contract}, nil
}

// bindERC20Upgradeable binds a generic wrapper to an already deployed contract.
func bindERC20Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Upgradeable *ERC20UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Upgradeable.Contract.ERC20UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Upgradeable *ERC20UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.ERC20UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Upgradeable *ERC20UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.ERC20UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Upgradeable *ERC20UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Upgradeable *ERC20UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Upgradeable *ERC20UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.Allowance(&_ERC20Upgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.Allowance(&_ERC20Upgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.BalanceOf(&_ERC20Upgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.BalanceOf(&_ERC20Upgradeable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Decimals() (uint8, error) {
	return _ERC20Upgradeable.Contract.Decimals(&_ERC20Upgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Decimals() (uint8, error) {
	return _ERC20Upgradeable.Contract.Decimals(&_ERC20Upgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Name() (string, error) {
	return _ERC20Upgradeable.Contract.Name(&_ERC20Upgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Name() (string, error) {
	return _ERC20Upgradeable.Contract.Name(&_ERC20Upgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Symbol() (string, error) {
	return _ERC20Upgradeable.Contract.Symbol(&_ERC20Upgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Symbol() (string, error) {
	return _ERC20Upgradeable.Contract.Symbol(&_ERC20Upgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Upgradeable.Contract.TotalSupply(&_ERC20Upgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Upgradeable.Contract.TotalSupply(&_ERC20Upgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Approve(&_ERC20Upgradeable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Approve(&_ERC20Upgradeable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.DecreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.DecreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.IncreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.IncreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Transfer(&_ERC20Upgradeable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Transfer(&_ERC20Upgradeable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.TransferFrom(&_ERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.TransferFrom(&_ERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// ERC20UpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Upgradeable contract.
type ERC20UpgradeableApprovalIterator struct {
	Event *ERC20UpgradeableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20UpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UpgradeableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20UpgradeableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20UpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UpgradeableApproval represents a Approval event raised by the ERC20Upgradeable contract.
type ERC20UpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20UpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableApprovalIterator{contract: _ERC20Upgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20UpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UpgradeableApproval)
				if err := _ERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) ParseApproval(log types.Log) (*ERC20UpgradeableApproval, error) {
	event := new(ERC20UpgradeableApproval)
	if err := _ERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Upgradeable contract.
type ERC20UpgradeableTransferIterator struct {
	Event *ERC20UpgradeableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ERC20UpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UpgradeableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ERC20UpgradeableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ERC20UpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UpgradeableTransfer represents a Transfer event raised by the ERC20Upgradeable contract.
type ERC20UpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20UpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableTransferIterator{contract: _ERC20Upgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20UpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UpgradeableTransfer)
				if err := _ERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) ParseTransfer(log types.Log) (*ERC20UpgradeableTransfer, error) {
	event := new(ERC20UpgradeableTransfer)
	if err := _ERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20UpgradeableMetaData contains all meta data concerning the IERC20Upgradeable contract.
var IERC20UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20UpgradeableMetaData.ABI instead.
var IERC20UpgradeableABI = IERC20UpgradeableMetaData.ABI

// Deprecated: Use IERC20UpgradeableMetaData.Sigs instead.
// IERC20UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IERC20UpgradeableFuncSigs = IERC20UpgradeableMetaData.Sigs

// IERC20Upgradeable is an auto generated Go binding around an Ethereum contract.
type IERC20Upgradeable struct {
	IERC20UpgradeableCaller     // Read-only binding to the contract
	IERC20UpgradeableTransactor // Write-only binding to the contract
	IERC20UpgradeableFilterer   // Log filterer for contract events
}

// IERC20UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20UpgradeableSession struct {
	Contract     *IERC20Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC20UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20UpgradeableCallerSession struct {
	Contract *IERC20UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC20UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20UpgradeableTransactorSession struct {
	Contract     *IERC20UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC20UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20UpgradeableRaw struct {
	Contract *IERC20Upgradeable // Generic contract binding to access the raw methods on
}

// IERC20UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20UpgradeableCallerRaw struct {
	Contract *IERC20UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20UpgradeableTransactorRaw struct {
	Contract *IERC20UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Upgradeable creates a new instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20Upgradeable(address common.Address, backend bind.ContractBackend) (*IERC20Upgradeable, error) {
	contract, err := bindIERC20Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Upgradeable{IERC20UpgradeableCaller: IERC20UpgradeableCaller{contract: contract}, IERC20UpgradeableTransactor: IERC20UpgradeableTransactor{contract: contract}, IERC20UpgradeableFilterer: IERC20UpgradeableFilterer{contract: contract}}, nil
}

// NewIERC20UpgradeableCaller creates a new read-only instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC20UpgradeableCaller, error) {
	contract, err := bindIERC20Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableCaller{contract: contract}, nil
}

// NewIERC20UpgradeableTransactor creates a new write-only instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20UpgradeableTransactor, error) {
	contract, err := bindIERC20Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableTransactor{contract: contract}, nil
}

// NewIERC20UpgradeableFilterer creates a new log filterer instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20UpgradeableFilterer, error) {
	contract, err := bindIERC20Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableFilterer{contract: contract}, nil
}

// bindIERC20Upgradeable binds a generic wrapper to an already deployed contract.
func bindIERC20Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Upgradeable *IERC20UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Upgradeable *IERC20UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Upgradeable *IERC20UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.Allowance(&_IERC20Upgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.Allowance(&_IERC20Upgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.BalanceOf(&_IERC20Upgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.BalanceOf(&_IERC20Upgradeable.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) TotalSupply() (*big.Int, error) {
	return _IERC20Upgradeable.Contract.TotalSupply(&_IERC20Upgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Upgradeable.Contract.TotalSupply(&_IERC20Upgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Approve(&_IERC20Upgradeable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Approve(&_IERC20Upgradeable.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Transfer(&_IERC20Upgradeable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Transfer(&_IERC20Upgradeable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.TransferFrom(&_IERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.TransferFrom(&_IERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// IERC20UpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Upgradeable contract.
type IERC20UpgradeableApprovalIterator struct {
	Event *IERC20UpgradeableApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20UpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20UpgradeableApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20UpgradeableApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20UpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20UpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20UpgradeableApproval represents a Approval event raised by the IERC20Upgradeable contract.
type IERC20UpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20UpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableApprovalIterator{contract: _IERC20Upgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20UpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20UpgradeableApproval)
				if err := _IERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) ParseApproval(log types.Log) (*IERC20UpgradeableApproval, error) {
	event := new(IERC20UpgradeableApproval)
	if err := _IERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20UpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Upgradeable contract.
type IERC20UpgradeableTransferIterator struct {
	Event *IERC20UpgradeableTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20UpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20UpgradeableTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20UpgradeableTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20UpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20UpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20UpgradeableTransfer represents a Transfer event raised by the IERC20Upgradeable contract.
type IERC20UpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20UpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableTransferIterator{contract: _IERC20Upgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20UpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20UpgradeableTransfer)
				if err := _IERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) ParseTransfer(log types.Log) (*IERC20UpgradeableTransfer, error) {
	event := new(IERC20UpgradeableTransfer)
	if err := _IERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFlashLoanReceiverMetaData contains all meta data concerning the IFlashLoanReceiver contract.
var IFlashLoanReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swap\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"executeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"405b019d": "executeOperation(address,address,uint256,uint256,bytes)",
	},
}

// IFlashLoanReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IFlashLoanReceiverMetaData.ABI instead.
var IFlashLoanReceiverABI = IFlashLoanReceiverMetaData.ABI

// Deprecated: Use IFlashLoanReceiverMetaData.Sigs instead.
// IFlashLoanReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IFlashLoanReceiverFuncSigs = IFlashLoanReceiverMetaData.Sigs

// IFlashLoanReceiver is an auto generated Go binding around an Ethereum contract.
type IFlashLoanReceiver struct {
	IFlashLoanReceiverCaller     // Read-only binding to the contract
	IFlashLoanReceiverTransactor // Write-only binding to the contract
	IFlashLoanReceiverFilterer   // Log filterer for contract events
}

// IFlashLoanReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFlashLoanReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFlashLoanReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFlashLoanReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFlashLoanReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFlashLoanReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFlashLoanReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFlashLoanReceiverSession struct {
	Contract     *IFlashLoanReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFlashLoanReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFlashLoanReceiverCallerSession struct {
	Contract *IFlashLoanReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IFlashLoanReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFlashLoanReceiverTransactorSession struct {
	Contract     *IFlashLoanReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IFlashLoanReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFlashLoanReceiverRaw struct {
	Contract *IFlashLoanReceiver // Generic contract binding to access the raw methods on
}

// IFlashLoanReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFlashLoanReceiverCallerRaw struct {
	Contract *IFlashLoanReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IFlashLoanReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFlashLoanReceiverTransactorRaw struct {
	Contract *IFlashLoanReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFlashLoanReceiver creates a new instance of IFlashLoanReceiver, bound to a specific deployed contract.
func NewIFlashLoanReceiver(address common.Address, backend bind.ContractBackend) (*IFlashLoanReceiver, error) {
	contract, err := bindIFlashLoanReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFlashLoanReceiver{IFlashLoanReceiverCaller: IFlashLoanReceiverCaller{contract: contract}, IFlashLoanReceiverTransactor: IFlashLoanReceiverTransactor{contract: contract}, IFlashLoanReceiverFilterer: IFlashLoanReceiverFilterer{contract: contract}}, nil
}

// NewIFlashLoanReceiverCaller creates a new read-only instance of IFlashLoanReceiver, bound to a specific deployed contract.
func NewIFlashLoanReceiverCaller(address common.Address, caller bind.ContractCaller) (*IFlashLoanReceiverCaller, error) {
	contract, err := bindIFlashLoanReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFlashLoanReceiverCaller{contract: contract}, nil
}

// NewIFlashLoanReceiverTransactor creates a new write-only instance of IFlashLoanReceiver, bound to a specific deployed contract.
func NewIFlashLoanReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IFlashLoanReceiverTransactor, error) {
	contract, err := bindIFlashLoanReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFlashLoanReceiverTransactor{contract: contract}, nil
}

// NewIFlashLoanReceiverFilterer creates a new log filterer instance of IFlashLoanReceiver, bound to a specific deployed contract.
func NewIFlashLoanReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IFlashLoanReceiverFilterer, error) {
	contract, err := bindIFlashLoanReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFlashLoanReceiverFilterer{contract: contract}, nil
}

// bindIFlashLoanReceiver binds a generic wrapper to an already deployed contract.
func bindIFlashLoanReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFlashLoanReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFlashLoanReceiver *IFlashLoanReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFlashLoanReceiver.Contract.IFlashLoanReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFlashLoanReceiver *IFlashLoanReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFlashLoanReceiver.Contract.IFlashLoanReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFlashLoanReceiver *IFlashLoanReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFlashLoanReceiver.Contract.IFlashLoanReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFlashLoanReceiver *IFlashLoanReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFlashLoanReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFlashLoanReceiver *IFlashLoanReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFlashLoanReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFlashLoanReceiver *IFlashLoanReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFlashLoanReceiver.Contract.contract.Transact(opts, method, params...)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x405b019d.
//
// Solidity: function executeOperation(address swap, address token, uint256 amount, uint256 fee, bytes params) returns()
func (_IFlashLoanReceiver *IFlashLoanReceiverTransactor) ExecuteOperation(opts *bind.TransactOpts, pool common.Address, token common.Address, amount *big.Int, fee *big.Int, params []byte) (*types.Transaction, error) {
	return _IFlashLoanReceiver.contract.Transact(opts, "executeOperation", pool, token, amount, fee, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x405b019d.
//
// Solidity: function executeOperation(address swap, address token, uint256 amount, uint256 fee, bytes params) returns()
func (_IFlashLoanReceiver *IFlashLoanReceiverSession) ExecuteOperation(pool common.Address, token common.Address, amount *big.Int, fee *big.Int, params []byte) (*types.Transaction, error) {
	return _IFlashLoanReceiver.Contract.ExecuteOperation(&_IFlashLoanReceiver.TransactOpts, pool, token, amount, fee, params)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x405b019d.
//
// Solidity: function executeOperation(address swap, address token, uint256 amount, uint256 fee, bytes params) returns()
func (_IFlashLoanReceiver *IFlashLoanReceiverTransactorSession) ExecuteOperation(pool common.Address, token common.Address, amount *big.Int, fee *big.Int, params []byte) (*types.Transaction, error) {
	return _IFlashLoanReceiver.Contract.ExecuteOperation(&_IFlashLoanReceiver.TransactOpts, pool, token, amount, fee, params)
}

// ISwapMetaData contains all meta data concerning the ISwap contract.
var ISwapMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateRemoveLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"deposit\",\"type\":\"bool\"}],\"name\":\"calculateTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVirtualPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"pooledTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"decimals\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"lpTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lpTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adminFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpTokenTargetAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"maxBurnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4d49e87d": "addLiquidity(uint256[],uint256,uint256)",
		"f2fad2b6": "calculateRemoveLiquidity(uint256)",
		"342a87a1": "calculateRemoveLiquidityOneToken(uint256,uint8)",
		"a95b089f": "calculateSwap(uint8,uint8,uint256)",
		"e6ab2806": "calculateTokenAmount(uint256[],bool)",
		"d46300fd": "getA()",
		"82b86600": "getToken(uint8)",
		"91ceb3eb": "getTokenBalance(uint8)",
		"66c0bd24": "getTokenIndex(address)",
		"e25aa5fa": "getVirtualPrice()",
		"b28cb6dc": "initialize(address[],uint8[],string,string,uint256,uint256,uint256,address)",
		"31cd52b0": "removeLiquidity(uint256,uint256[],uint256)",
		"84cdd9bc": "removeLiquidityImbalance(uint256[],uint256,uint256)",
		"3e3a1560": "removeLiquidityOneToken(uint256,uint8,uint256,uint256)",
		"91695586": "swap(uint8,uint8,uint256,uint256,uint256)",
	},
}

// ISwapABI is the input ABI used to generate the binding from.
// Deprecated: Use ISwapMetaData.ABI instead.
var ISwapABI = ISwapMetaData.ABI

// Deprecated: Use ISwapMetaData.Sigs instead.
// ISwapFuncSigs maps the 4-byte function signature to its string representation.
var ISwapFuncSigs = ISwapMetaData.Sigs

// ISwap is an auto generated Go binding around an Ethereum contract.
type ISwap struct {
	ISwapCaller     // Read-only binding to the contract
	ISwapTransactor // Write-only binding to the contract
	ISwapFilterer   // Log filterer for contract events
}

// ISwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapSession struct {
	Contract     *ISwap            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapCallerSession struct {
	Contract *ISwapCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ISwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapTransactorSession struct {
	Contract     *ISwapTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapRaw struct {
	Contract *ISwap // Generic contract binding to access the raw methods on
}

// ISwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapCallerRaw struct {
	Contract *ISwapCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapTransactorRaw struct {
	Contract *ISwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwap creates a new instance of ISwap, bound to a specific deployed contract.
func NewISwap(address common.Address, backend bind.ContractBackend) (*ISwap, error) {
	contract, err := bindISwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwap{ISwapCaller: ISwapCaller{contract: contract}, ISwapTransactor: ISwapTransactor{contract: contract}, ISwapFilterer: ISwapFilterer{contract: contract}}, nil
}

// NewISwapCaller creates a new read-only instance of ISwap, bound to a specific deployed contract.
func NewISwapCaller(address common.Address, caller bind.ContractCaller) (*ISwapCaller, error) {
	contract, err := bindISwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapCaller{contract: contract}, nil
}

// NewISwapTransactor creates a new write-only instance of ISwap, bound to a specific deployed contract.
func NewISwapTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapTransactor, error) {
	contract, err := bindISwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapTransactor{contract: contract}, nil
}

// NewISwapFilterer creates a new log filterer instance of ISwap, bound to a specific deployed contract.
func NewISwapFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapFilterer, error) {
	contract, err := bindISwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapFilterer{contract: contract}, nil
}

// bindISwap binds a generic wrapper to an already deployed contract.
func bindISwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwap *ISwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwap.Contract.ISwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwap *ISwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwap.Contract.ISwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwap *ISwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwap.Contract.ISwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwap *ISwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwap *ISwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwap *ISwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwap.Contract.contract.Transact(opts, method, params...)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_ISwap *ISwapCaller) CalculateRemoveLiquidity(opts *bind.CallOpts, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateRemoveLiquidity", amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_ISwap *ISwapSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidity(&_ISwap.CallOpts, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_ISwap *ISwapCallerSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidity(&_ISwap.CallOpts, amount)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_ISwap *ISwapCaller) CalculateRemoveLiquidityOneToken(opts *bind.CallOpts, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateRemoveLiquidityOneToken", tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_ISwap *ISwapSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidityOneToken(&_ISwap.CallOpts, tokenAmount, tokenIndex)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_ISwap *ISwapCallerSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _ISwap.Contract.CalculateRemoveLiquidityOneToken(&_ISwap.CallOpts, tokenAmount, tokenIndex)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_ISwap *ISwapCaller) CalculateSwap(opts *bind.CallOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateSwap", tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_ISwap *ISwapSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwap.Contract.CalculateSwap(&_ISwap.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_ISwap *ISwapCallerSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _ISwap.Contract.CalculateSwap(&_ISwap.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_ISwap *ISwapCaller) CalculateTokenAmount(opts *bind.CallOpts, amounts []*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "calculateTokenAmount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_ISwap *ISwapSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _ISwap.Contract.CalculateTokenAmount(&_ISwap.CallOpts, amounts, deposit)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_ISwap *ISwapCallerSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _ISwap.Contract.CalculateTokenAmount(&_ISwap.CallOpts, amounts, deposit)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_ISwap *ISwapCaller) GetA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_ISwap *ISwapSession) GetA() (*big.Int, error) {
	return _ISwap.Contract.GetA(&_ISwap.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_ISwap *ISwapCallerSession) GetA() (*big.Int, error) {
	return _ISwap.Contract.GetA(&_ISwap.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_ISwap *ISwapCaller) GetToken(opts *bind.CallOpts, index uint8) (common.Address, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_ISwap *ISwapSession) GetToken(index uint8) (common.Address, error) {
	return _ISwap.Contract.GetToken(&_ISwap.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_ISwap *ISwapCallerSession) GetToken(index uint8) (common.Address, error) {
	return _ISwap.Contract.GetToken(&_ISwap.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_ISwap *ISwapCaller) GetTokenBalance(opts *bind.CallOpts, index uint8) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getTokenBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_ISwap *ISwapSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _ISwap.Contract.GetTokenBalance(&_ISwap.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_ISwap *ISwapCallerSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _ISwap.Contract.GetTokenBalance(&_ISwap.CallOpts, index)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_ISwap *ISwapCaller) GetTokenIndex(opts *bind.CallOpts, tokenAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getTokenIndex", tokenAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_ISwap *ISwapSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _ISwap.Contract.GetTokenIndex(&_ISwap.CallOpts, tokenAddress)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_ISwap *ISwapCallerSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _ISwap.Contract.GetTokenIndex(&_ISwap.CallOpts, tokenAddress)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_ISwap *ISwapCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ISwap.contract.Call(opts, &out, "getVirtualPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_ISwap *ISwapSession) GetVirtualPrice() (*big.Int, error) {
	return _ISwap.Contract.GetVirtualPrice(&_ISwap.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_ISwap *ISwapCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _ISwap.Contract.GetVirtualPrice(&_ISwap.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) AddLiquidity(opts *bind.TransactOpts, amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "addLiquidity", amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.AddLiquidity(&_ISwap.TransactOpts, amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.AddLiquidity(&_ISwap.TransactOpts, amounts, minToMint, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 a, uint256 fee, uint256 adminFee, address lpTokenTargetAddress) returns()
func (_ISwap *ISwapTransactor) Initialize(opts *bind.TransactOpts, pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, a *big.Int, fee *big.Int, adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "initialize", pooledTokens, decimals, lpTokenName, lpTokenSymbol, a, fee, adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 a, uint256 fee, uint256 adminFee, address lpTokenTargetAddress) returns()
func (_ISwap *ISwapSession) Initialize(pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, a *big.Int, fee *big.Int, adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _ISwap.Contract.Initialize(&_ISwap.TransactOpts, pooledTokens, decimals, lpTokenName, lpTokenSymbol, a, fee, adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 a, uint256 fee, uint256 adminFee, address lpTokenTargetAddress) returns()
func (_ISwap *ISwapTransactorSession) Initialize(pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, a *big.Int, fee *big.Int, adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _ISwap.Contract.Initialize(&_ISwap.TransactOpts, pooledTokens, decimals, lpTokenName, lpTokenSymbol, a, fee, adminFee, lpTokenTargetAddress)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_ISwap *ISwapTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "removeLiquidity", amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_ISwap *ISwapSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidity(&_ISwap.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_ISwap *ISwapTransactorSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidity(&_ISwap.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "removeLiquidityImbalance", amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityImbalance(&_ISwap.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityImbalance(&_ISwap.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) RemoveLiquidityOneToken(opts *bind.TransactOpts, tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "removeLiquidityOneToken", tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityOneToken(&_ISwap.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.RemoveLiquidityOneToken(&_ISwap.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactor) Swap(opts *bind.TransactOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.contract.Transact(opts, "swap", tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_ISwap *ISwapSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.Swap(&_ISwap.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_ISwap *ISwapTransactorSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _ISwap.Contract.Swap(&_ISwap.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// LPTokenMetaData contains all meta data concerning the LPToken contract.
var LPTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"42966c68": "burn(uint256)",
		"79cc6790": "burnFrom(address,uint256)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"4cd88b76": "initialize(string,string)",
		"40c10f19": "mint(address,uint256)",
		"06fdde03": "name()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611a49806100206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c806370a08231116100b257806395d89b4111610081578063a9059cbb11610066578063a9059cbb14610524578063dd62ed3e1461055d578063f2fde38b1461059857610136565b806395d89b41146104e3578063a457c2d7146104eb57610136565b806370a082311461043e578063715018a61461047157806379cc6790146104795780638da5cb5b146104b257610136565b8063313ce5671161010957806340c10f19116100ee57806340c10f19146102b957806342966c68146102f45780634cd88b761461031157610136565b8063313ce56714610262578063395093511461028057610136565b806306fdde031461013b578063095ea7b3146101b857806318160ddd1461020557806323b872dd1461021f575b600080fd5b6101436105cb565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561017d578181015183820152602001610165565b50505050905090810190601f1680156101aa5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101f1600480360360408110156101ce57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561067f565b604080519115158252519081900360200190f35b61020d61069c565b60408051918252519081900360200190f35b6101f16004803603606081101561023557600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356106a2565b61026a610743565b6040805160ff9092168252519081900360200190f35b6101f16004803603604081101561029657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561074c565b6102f2600480360360408110156102cf57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356107a7565b005b6102f26004803603602081101561030a57600080fd5b5035610895565b6101f16004803603604081101561032757600080fd5b81019060208101813564010000000081111561034257600080fd5b82018360208201111561035457600080fd5b8035906020019184600183028401116401000000008311171561037657600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156103c957600080fd5b8201836020820111156103db57600080fd5b803590602001918460018302840111640100000000831117156103fd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506108a9945050505050565b61020d6004803603602081101561045457600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661098c565b6102f26109b4565b6102f26004803603604081101561048f57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610ab1565b6104ba610b0b565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610143610b27565b6101f16004803603604081101561050157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610ba6565b6101f16004803603604081101561053a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610c1b565b61020d6004803603604081101561057357600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610c2f565b6102f2600480360360208110156105ae57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610c67565b60368054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106755780601f1061064a57610100808354040283529160200191610675565b820191906000526020600020905b81548152906001019060200180831161065857829003601f168201915b5050505050905090565b600061069361068c610dd5565b8484610dd9565b50600192915050565b60355490565b60006106af848484610eec565b610739846106bb610dd5565b610734856040518060600160405280602881526020016119396028913973ffffffffffffffffffffffffffffffffffffffff8a16600090815260346020526040812090610706610dd5565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002054919061108a565b610dd9565b5060019392505050565b60385460ff1690565b6000610693610759610dd5565b84610734856034600061076a610dd5565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918c168152925290205490611121565b6107af610dd5565b73ffffffffffffffffffffffffffffffffffffffff166107cd610b0b565b73ffffffffffffffffffffffffffffffffffffffff1614610835576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b80610887576040805162461bcd60e51b815260206004820152601660248201527f4c50546f6b656e3a2063616e6e6f74206d696e74203000000000000000000000604482015290519081900360640190fd5b6108918282611182565b5050565b6108a66108a0610dd5565b8261129b565b50565b60008054610100900460ff16806108c357506108c36113cb565b806108d1575060005460ff16155b61090c5760405162461bcd60e51b815260040180806020018281038252602e81526020018061190b602e913960400191505060405180910390fd5b600054610100900460ff1615801561095557600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0061ff0019909116610100171660011790555b61095d6113dc565b610967848461149b565b61096f6115af565b600191508015610985576000805461ff00191690555b5092915050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526033602052604090205490565b6109bc610dd5565b73ffffffffffffffffffffffffffffffffffffffff166109da610b0b565b73ffffffffffffffffffffffffffffffffffffffff1614610a42576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60975460405160009173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3609780547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b6000610ae88260405180606001604052806024815260200161196160249139610ae186610adc610dd5565b610c2f565b919061108a565b9050610afc83610af6610dd5565b83610dd9565b610b06838361129b565b505050565b60975473ffffffffffffffffffffffffffffffffffffffff1690565b60378054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156106755780601f1061064a57610100808354040283529160200191610675565b6000610693610bb3610dd5565b84610734856040518060600160405280602581526020016119ef6025913960346000610bdd610dd5565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918d1681529252902054919061108a565b6000610693610c28610dd5565b8484610eec565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260346020908152604080832093909416825291909152205490565b610c6f610dd5565b73ffffffffffffffffffffffffffffffffffffffff16610c8d610b0b565b73ffffffffffffffffffffffffffffffffffffffff1614610cf5576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8116610d475760405162461bcd60e51b815260040180806020018281038252602681526020018061189d6026913960400191505060405180910390fd5b60975460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b3390565b73ffffffffffffffffffffffffffffffffffffffff8316610e2b5760405162461bcd60e51b81526004018080602001828103825260248152602001806119cb6024913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610e7d5760405162461bcd60e51b81526004018080602001828103825260228152602001806118c36022913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff808416600081815260346020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316610f3e5760405162461bcd60e51b81526004018080602001828103825260258152602001806119a66025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610f905760405162461bcd60e51b81526004018080602001828103825260238152602001806118586023913960400191505060405180910390fd5b610f9b8383836116eb565b610fe5816040518060600160405280602681526020016118e56026913973ffffffffffffffffffffffffffffffffffffffff8616600090815260336020526040902054919061108a565b73ffffffffffffffffffffffffffffffffffffffff80851660009081526033602052604080822093909355908416815220546110219082611121565b73ffffffffffffffffffffffffffffffffffffffff80841660008181526033602090815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600081848411156111195760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156110de5781810151838201526020016110c6565b50505050905090810190601f16801561110b5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b60008282018381101561117b576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b73ffffffffffffffffffffffffffffffffffffffff82166111ea576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b6111f6600083836116eb565b6035546112039082611121565b60355573ffffffffffffffffffffffffffffffffffffffff82166000908152603360205260409020546112369082611121565b73ffffffffffffffffffffffffffffffffffffffff831660008181526033602090815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b73ffffffffffffffffffffffffffffffffffffffff82166112ed5760405162461bcd60e51b81526004018080602001828103825260218152602001806119856021913960400191505060405180910390fd5b6112f9826000836116eb565b6113438160405180606001604052806022815260200161187b6022913973ffffffffffffffffffffffffffffffffffffffff8516600090815260336020526040902054919061108a565b73ffffffffffffffffffffffffffffffffffffffff83166000908152603360205260409020556035546113769082611761565b60355560408051828152905160009173ffffffffffffffffffffffffffffffffffffffff8516917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35050565b60006113d6306117be565b15905090565b600054610100900460ff16806113f557506113f56113cb565b80611403575060005460ff16155b61143e5760405162461bcd60e51b815260040180806020018281038252602e81526020018061190b602e913960400191505060405180910390fd5b600054610100900460ff1615801561148757600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0061ff0019909116610100171660011790555b80156108a6576000805461ff001916905550565b600054610100900460ff16806114b457506114b46113cb565b806114c2575060005460ff16155b6114fd5760405162461bcd60e51b815260040180806020018281038252602e81526020018061190b602e913960400191505060405180910390fd5b600054610100900460ff1615801561154657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0061ff0019909116610100171660011790555b82516115599060369060208601906117c4565b50815161156d9060379060208501906117c4565b50603880547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660121790558015610b06576000805461ff0019169055505050565b600054610100900460ff16806115c857506115c86113cb565b806115d6575060005460ff16155b6116115760405162461bcd60e51b815260040180806020018281038252602e81526020018061190b602e913960400191505060405180910390fd5b600054610100900460ff1615801561165a57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0061ff0019909116610100171660011790555b6000611664610dd5565b609780547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a35080156108a6576000805461ff001916905550565b6116f6838383610b06565b73ffffffffffffffffffffffffffffffffffffffff8216301415610b06576040805162461bcd60e51b815260206004820152601e60248201527f4c50546f6b656e3a2063616e6e6f742073656e6420746f20697473656c660000604482015290519081900360640190fd5b6000828211156117b8576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b3b151590565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061180557805160ff1916838001178555611832565b82800160010185558215611832579182015b82811115611832578251825591602001919060010190611817565b5061183e929150611842565b5090565b5b8082111561183e576000815560010161184356fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a656445524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e20616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122010f05322837e1f19d8b1d0a51d3f4608c5e18cd5eb3988c66fe1d415b1d890a264736f6c634300060c0033",
}

// LPTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use LPTokenMetaData.ABI instead.
var LPTokenABI = LPTokenMetaData.ABI

// Deprecated: Use LPTokenMetaData.Sigs instead.
// LPTokenFuncSigs maps the 4-byte function signature to its string representation.
var LPTokenFuncSigs = LPTokenMetaData.Sigs

// LPTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LPTokenMetaData.Bin instead.
var LPTokenBin = LPTokenMetaData.Bin

// DeployLPToken deploys a new Ethereum contract, binding an instance of LPToken to it.
func DeployLPToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LPToken, error) {
	parsed, err := LPTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LPTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LPToken{LPTokenCaller: LPTokenCaller{contract: contract}, LPTokenTransactor: LPTokenTransactor{contract: contract}, LPTokenFilterer: LPTokenFilterer{contract: contract}}, nil
}

// LPToken is an auto generated Go binding around an Ethereum contract.
type LPToken struct {
	LPTokenCaller     // Read-only binding to the contract
	LPTokenTransactor // Write-only binding to the contract
	LPTokenFilterer   // Log filterer for contract events
}

// LPTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type LPTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LPTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LPTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LPTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LPTokenSession struct {
	Contract     *LPToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LPTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LPTokenCallerSession struct {
	Contract *LPTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LPTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LPTokenTransactorSession struct {
	Contract     *LPTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LPTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type LPTokenRaw struct {
	Contract *LPToken // Generic contract binding to access the raw methods on
}

// LPTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LPTokenCallerRaw struct {
	Contract *LPTokenCaller // Generic read-only contract binding to access the raw methods on
}

// LPTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LPTokenTransactorRaw struct {
	Contract *LPTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLPToken creates a new instance of LPToken, bound to a specific deployed contract.
func NewLPToken(address common.Address, backend bind.ContractBackend) (*LPToken, error) {
	contract, err := bindLPToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LPToken{LPTokenCaller: LPTokenCaller{contract: contract}, LPTokenTransactor: LPTokenTransactor{contract: contract}, LPTokenFilterer: LPTokenFilterer{contract: contract}}, nil
}

// NewLPTokenCaller creates a new read-only instance of LPToken, bound to a specific deployed contract.
func NewLPTokenCaller(address common.Address, caller bind.ContractCaller) (*LPTokenCaller, error) {
	contract, err := bindLPToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LPTokenCaller{contract: contract}, nil
}

// NewLPTokenTransactor creates a new write-only instance of LPToken, bound to a specific deployed contract.
func NewLPTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*LPTokenTransactor, error) {
	contract, err := bindLPToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LPTokenTransactor{contract: contract}, nil
}

// NewLPTokenFilterer creates a new log filterer instance of LPToken, bound to a specific deployed contract.
func NewLPTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*LPTokenFilterer, error) {
	contract, err := bindLPToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LPTokenFilterer{contract: contract}, nil
}

// bindLPToken binds a generic wrapper to an already deployed contract.
func bindLPToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LPTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LPToken *LPTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LPToken.Contract.LPTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LPToken *LPTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPToken.Contract.LPTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LPToken *LPTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LPToken.Contract.LPTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LPToken *LPTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LPToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LPToken *LPTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LPToken *LPTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LPToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LPToken *LPTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LPToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LPToken *LPTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LPToken.Contract.Allowance(&_LPToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LPToken *LPTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LPToken.Contract.Allowance(&_LPToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LPToken *LPTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LPToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LPToken *LPTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LPToken.Contract.BalanceOf(&_LPToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LPToken *LPTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LPToken.Contract.BalanceOf(&_LPToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LPToken *LPTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LPToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LPToken *LPTokenSession) Decimals() (uint8, error) {
	return _LPToken.Contract.Decimals(&_LPToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LPToken *LPTokenCallerSession) Decimals() (uint8, error) {
	return _LPToken.Contract.Decimals(&_LPToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LPToken *LPTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LPToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LPToken *LPTokenSession) Name() (string, error) {
	return _LPToken.Contract.Name(&_LPToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LPToken *LPTokenCallerSession) Name() (string, error) {
	return _LPToken.Contract.Name(&_LPToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPToken *LPTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LPToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPToken *LPTokenSession) Owner() (common.Address, error) {
	return _LPToken.Contract.Owner(&_LPToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LPToken *LPTokenCallerSession) Owner() (common.Address, error) {
	return _LPToken.Contract.Owner(&_LPToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LPToken *LPTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LPToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LPToken *LPTokenSession) Symbol() (string, error) {
	return _LPToken.Contract.Symbol(&_LPToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LPToken *LPTokenCallerSession) Symbol() (string, error) {
	return _LPToken.Contract.Symbol(&_LPToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LPToken *LPTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LPToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LPToken *LPTokenSession) TotalSupply() (*big.Int, error) {
	return _LPToken.Contract.TotalSupply(&_LPToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LPToken *LPTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _LPToken.Contract.TotalSupply(&_LPToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LPToken *LPTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LPToken *LPTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Approve(&_LPToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LPToken *LPTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Approve(&_LPToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_LPToken *LPTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_LPToken *LPTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Burn(&_LPToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_LPToken *LPTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Burn(&_LPToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_LPToken *LPTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_LPToken *LPTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.BurnFrom(&_LPToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_LPToken *LPTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.BurnFrom(&_LPToken.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LPToken *LPTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LPToken *LPTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.DecreaseAllowance(&_LPToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LPToken *LPTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.DecreaseAllowance(&_LPToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LPToken *LPTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LPToken *LPTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.IncreaseAllowance(&_LPToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LPToken *LPTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.IncreaseAllowance(&_LPToken.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns(bool)
func (_LPToken *LPTokenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "initialize", name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns(bool)
func (_LPToken *LPTokenSession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _LPToken.Contract.Initialize(&_LPToken.TransactOpts, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name, string symbol) returns(bool)
func (_LPToken *LPTokenTransactorSession) Initialize(name string, symbol string) (*types.Transaction, error) {
	return _LPToken.Contract.Initialize(&_LPToken.TransactOpts, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_LPToken *LPTokenTransactor) Mint(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "mint", recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_LPToken *LPTokenSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Mint(&_LPToken.TransactOpts, recipient, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address recipient, uint256 amount) returns()
func (_LPToken *LPTokenTransactorSession) Mint(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Mint(&_LPToken.TransactOpts, recipient, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPToken *LPTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPToken *LPTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _LPToken.Contract.RenounceOwnership(&_LPToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LPToken *LPTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LPToken.Contract.RenounceOwnership(&_LPToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LPToken *LPTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LPToken *LPTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Transfer(&_LPToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_LPToken *LPTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.Transfer(&_LPToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LPToken *LPTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LPToken *LPTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.TransferFrom(&_LPToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_LPToken *LPTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LPToken.Contract.TransferFrom(&_LPToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPToken *LPTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LPToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPToken *LPTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LPToken.Contract.TransferOwnership(&_LPToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LPToken *LPTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LPToken.Contract.TransferOwnership(&_LPToken.TransactOpts, newOwner)
}

// LPTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LPToken contract.
type LPTokenApprovalIterator struct {
	Event *LPTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LPTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LPTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LPTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPTokenApproval represents a Approval event raised by the LPToken contract.
type LPTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LPToken *LPTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LPTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LPToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LPTokenApprovalIterator{contract: _LPToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LPToken *LPTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LPTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LPToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPTokenApproval)
				if err := _LPToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LPToken *LPTokenFilterer) ParseApproval(log types.Log) (*LPTokenApproval, error) {
	event := new(LPTokenApproval)
	if err := _LPToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LPToken contract.
type LPTokenOwnershipTransferredIterator struct {
	Event *LPTokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LPTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPTokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LPTokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LPTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPTokenOwnershipTransferred represents a OwnershipTransferred event raised by the LPToken contract.
type LPTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPToken *LPTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LPTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LPTokenOwnershipTransferredIterator{contract: _LPToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPToken *LPTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LPTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LPToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPTokenOwnershipTransferred)
				if err := _LPToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LPToken *LPTokenFilterer) ParseOwnershipTransferred(log types.Log) (*LPTokenOwnershipTransferred, error) {
	event := new(LPTokenOwnershipTransferred)
	if err := _LPToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LPTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LPToken contract.
type LPTokenTransferIterator struct {
	Event *LPTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LPTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LPTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LPTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LPTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LPTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LPTokenTransfer represents a Transfer event raised by the LPToken contract.
type LPTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LPToken *LPTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LPTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LPToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LPTokenTransferIterator{contract: _LPToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LPToken *LPTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LPTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LPToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LPTokenTransfer)
				if err := _LPToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LPToken *LPTokenFilterer) ParseTransfer(log types.Log) (*LPTokenTransfer, error) {
	event := new(LPTokenTransfer)
	if err := _LPToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MathUtilsMetaData contains all meta data concerning the MathUtils contract.
var MathUtilsMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208fc54c301d7e54edbcc8fcebdd1df1120096cce97665d4b14d8773a7fa62873264736f6c634300060c0033",
}

// MathUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use MathUtilsMetaData.ABI instead.
var MathUtilsABI = MathUtilsMetaData.ABI

// MathUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MathUtilsMetaData.Bin instead.
var MathUtilsBin = MathUtilsMetaData.Bin

// DeployMathUtils deploys a new Ethereum contract, binding an instance of MathUtils to it.
func DeployMathUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MathUtils, error) {
	parsed, err := MathUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MathUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MathUtils{MathUtilsCaller: MathUtilsCaller{contract: contract}, MathUtilsTransactor: MathUtilsTransactor{contract: contract}, MathUtilsFilterer: MathUtilsFilterer{contract: contract}}, nil
}

// MathUtils is an auto generated Go binding around an Ethereum contract.
type MathUtils struct {
	MathUtilsCaller     // Read-only binding to the contract
	MathUtilsTransactor // Write-only binding to the contract
	MathUtilsFilterer   // Log filterer for contract events
}

// MathUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathUtilsSession struct {
	Contract     *MathUtils        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathUtilsCallerSession struct {
	Contract *MathUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MathUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathUtilsTransactorSession struct {
	Contract     *MathUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MathUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathUtilsRaw struct {
	Contract *MathUtils // Generic contract binding to access the raw methods on
}

// MathUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathUtilsCallerRaw struct {
	Contract *MathUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// MathUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathUtilsTransactorRaw struct {
	Contract *MathUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMathUtils creates a new instance of MathUtils, bound to a specific deployed contract.
func NewMathUtils(address common.Address, backend bind.ContractBackend) (*MathUtils, error) {
	contract, err := bindMathUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MathUtils{MathUtilsCaller: MathUtilsCaller{contract: contract}, MathUtilsTransactor: MathUtilsTransactor{contract: contract}, MathUtilsFilterer: MathUtilsFilterer{contract: contract}}, nil
}

// NewMathUtilsCaller creates a new read-only instance of MathUtils, bound to a specific deployed contract.
func NewMathUtilsCaller(address common.Address, caller bind.ContractCaller) (*MathUtilsCaller, error) {
	contract, err := bindMathUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathUtilsCaller{contract: contract}, nil
}

// NewMathUtilsTransactor creates a new write-only instance of MathUtils, bound to a specific deployed contract.
func NewMathUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*MathUtilsTransactor, error) {
	contract, err := bindMathUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathUtilsTransactor{contract: contract}, nil
}

// NewMathUtilsFilterer creates a new log filterer instance of MathUtils, bound to a specific deployed contract.
func NewMathUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*MathUtilsFilterer, error) {
	contract, err := bindMathUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathUtilsFilterer{contract: contract}, nil
}

// bindMathUtils binds a generic wrapper to an already deployed contract.
func bindMathUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathUtils *MathUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathUtils.Contract.MathUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathUtils *MathUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathUtils.Contract.MathUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathUtils *MathUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathUtils.Contract.MathUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MathUtils *MathUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MathUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MathUtils *MathUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MathUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MathUtils *MathUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MathUtils.Contract.contract.Transact(opts, method, params...)
}

// OwnableUpgradeableMetaData contains all meta data concerning the OwnableUpgradeable contract.
var OwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableUpgradeableMetaData.ABI instead.
var OwnableUpgradeableABI = OwnableUpgradeableMetaData.ABI

// Deprecated: Use OwnableUpgradeableMetaData.Sigs instead.
// OwnableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableUpgradeableFuncSigs = OwnableUpgradeableMetaData.Sigs

// OwnableUpgradeable is an auto generated Go binding around an Ethereum contract.
type OwnableUpgradeable struct {
	OwnableUpgradeableCaller     // Read-only binding to the contract
	OwnableUpgradeableTransactor // Write-only binding to the contract
	OwnableUpgradeableFilterer   // Log filterer for contract events
}

// OwnableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableUpgradeableSession struct {
	Contract     *OwnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OwnableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableUpgradeableCallerSession struct {
	Contract *OwnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OwnableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableUpgradeableTransactorSession struct {
	Contract     *OwnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableUpgradeableRaw struct {
	Contract *OwnableUpgradeable // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCallerRaw struct {
	Contract *OwnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactorRaw struct {
	Contract *OwnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeable creates a new instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeable(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeable, error) {
	contract, err := bindOwnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeable{OwnableUpgradeableCaller: OwnableUpgradeableCaller{contract: contract}, OwnableUpgradeableTransactor: OwnableUpgradeableTransactor{contract: contract}, OwnableUpgradeableFilterer: OwnableUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableCaller creates a new read-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableCaller, error) {
	contract, err := bindOwnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableCaller{contract: contract}, nil
}

// NewOwnableUpgradeableTransactor creates a new write-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableTransactor, error) {
	contract, err := bindOwnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableFilterer creates a new log filterer instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableFilterer, error) {
	contract, err := bindOwnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableFilterer{contract: contract}, nil
}

// bindOwnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// OwnableUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnableUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableOwnershipTransferredIterator{contract: _OwnableUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableOwnershipTransferred)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableOwnershipTransferred, error) {
	event := new(OwnableUpgradeableOwnershipTransferred)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerPausableUpgradeableMetaData contains all meta data concerning the OwnerPausableUpgradeable contract.
var OwnerPausableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
	},
}

// OwnerPausableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnerPausableUpgradeableMetaData.ABI instead.
var OwnerPausableUpgradeableABI = OwnerPausableUpgradeableMetaData.ABI

// Deprecated: Use OwnerPausableUpgradeableMetaData.Sigs instead.
// OwnerPausableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var OwnerPausableUpgradeableFuncSigs = OwnerPausableUpgradeableMetaData.Sigs

// OwnerPausableUpgradeable is an auto generated Go binding around an Ethereum contract.
type OwnerPausableUpgradeable struct {
	OwnerPausableUpgradeableCaller     // Read-only binding to the contract
	OwnerPausableUpgradeableTransactor // Write-only binding to the contract
	OwnerPausableUpgradeableFilterer   // Log filterer for contract events
}

// OwnerPausableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnerPausableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerPausableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnerPausableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerPausableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnerPausableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnerPausableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnerPausableUpgradeableSession struct {
	Contract     *OwnerPausableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OwnerPausableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnerPausableUpgradeableCallerSession struct {
	Contract *OwnerPausableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// OwnerPausableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnerPausableUpgradeableTransactorSession struct {
	Contract     *OwnerPausableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// OwnerPausableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnerPausableUpgradeableRaw struct {
	Contract *OwnerPausableUpgradeable // Generic contract binding to access the raw methods on
}

// OwnerPausableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnerPausableUpgradeableCallerRaw struct {
	Contract *OwnerPausableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnerPausableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnerPausableUpgradeableTransactorRaw struct {
	Contract *OwnerPausableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnerPausableUpgradeable creates a new instance of OwnerPausableUpgradeable, bound to a specific deployed contract.
func NewOwnerPausableUpgradeable(address common.Address, backend bind.ContractBackend) (*OwnerPausableUpgradeable, error) {
	contract, err := bindOwnerPausableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnerPausableUpgradeable{OwnerPausableUpgradeableCaller: OwnerPausableUpgradeableCaller{contract: contract}, OwnerPausableUpgradeableTransactor: OwnerPausableUpgradeableTransactor{contract: contract}, OwnerPausableUpgradeableFilterer: OwnerPausableUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnerPausableUpgradeableCaller creates a new read-only instance of OwnerPausableUpgradeable, bound to a specific deployed contract.
func NewOwnerPausableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*OwnerPausableUpgradeableCaller, error) {
	contract, err := bindOwnerPausableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerPausableUpgradeableCaller{contract: contract}, nil
}

// NewOwnerPausableUpgradeableTransactor creates a new write-only instance of OwnerPausableUpgradeable, bound to a specific deployed contract.
func NewOwnerPausableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnerPausableUpgradeableTransactor, error) {
	contract, err := bindOwnerPausableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnerPausableUpgradeableTransactor{contract: contract}, nil
}

// NewOwnerPausableUpgradeableFilterer creates a new log filterer instance of OwnerPausableUpgradeable, bound to a specific deployed contract.
func NewOwnerPausableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnerPausableUpgradeableFilterer, error) {
	contract, err := bindOwnerPausableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnerPausableUpgradeableFilterer{contract: contract}, nil
}

// bindOwnerPausableUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnerPausableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnerPausableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerPausableUpgradeable.Contract.OwnerPausableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.OwnerPausableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.OwnerPausableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnerPausableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnerPausableUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableSession) Owner() (common.Address, error) {
	return _OwnerPausableUpgradeable.Contract.Owner(&_OwnerPausableUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableCallerSession) Owner() (common.Address, error) {
	return _OwnerPausableUpgradeable.Contract.Owner(&_OwnerPausableUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OwnerPausableUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableSession) Paused() (bool, error) {
	return _OwnerPausableUpgradeable.Contract.Paused(&_OwnerPausableUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableCallerSession) Paused() (bool, error) {
	return _OwnerPausableUpgradeable.Contract.Paused(&_OwnerPausableUpgradeable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableSession) Pause() (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.Pause(&_OwnerPausableUpgradeable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactorSession) Pause() (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.Pause(&_OwnerPausableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.RenounceOwnership(&_OwnerPausableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.RenounceOwnership(&_OwnerPausableUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.TransferOwnership(&_OwnerPausableUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.TransferOwnership(&_OwnerPausableUpgradeable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableSession) Unpause() (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.Unpause(&_OwnerPausableUpgradeable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableTransactorSession) Unpause() (*types.Transaction, error) {
	return _OwnerPausableUpgradeable.Contract.Unpause(&_OwnerPausableUpgradeable.TransactOpts)
}

// OwnerPausableUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnerPausableUpgradeable contract.
type OwnerPausableUpgradeableOwnershipTransferredIterator struct {
	Event *OwnerPausableUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnerPausableUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerPausableUpgradeableOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnerPausableUpgradeableOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnerPausableUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerPausableUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerPausableUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the OwnerPausableUpgradeable contract.
type OwnerPausableUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnerPausableUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnerPausableUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnerPausableUpgradeableOwnershipTransferredIterator{contract: _OwnerPausableUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnerPausableUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnerPausableUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerPausableUpgradeableOwnershipTransferred)
				if err := _OwnerPausableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnerPausableUpgradeableOwnershipTransferred, error) {
	event := new(OwnerPausableUpgradeableOwnershipTransferred)
	if err := _OwnerPausableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerPausableUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OwnerPausableUpgradeable contract.
type OwnerPausableUpgradeablePausedIterator struct {
	Event *OwnerPausableUpgradeablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnerPausableUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerPausableUpgradeablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnerPausableUpgradeablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnerPausableUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerPausableUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerPausableUpgradeablePaused represents a Paused event raised by the OwnerPausableUpgradeable contract.
type OwnerPausableUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*OwnerPausableUpgradeablePausedIterator, error) {

	logs, sub, err := _OwnerPausableUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OwnerPausableUpgradeablePausedIterator{contract: _OwnerPausableUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OwnerPausableUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _OwnerPausableUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerPausableUpgradeablePaused)
				if err := _OwnerPausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) ParsePaused(log types.Log) (*OwnerPausableUpgradeablePaused, error) {
	event := new(OwnerPausableUpgradeablePaused)
	if err := _OwnerPausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnerPausableUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OwnerPausableUpgradeable contract.
type OwnerPausableUpgradeableUnpausedIterator struct {
	Event *OwnerPausableUpgradeableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *OwnerPausableUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnerPausableUpgradeableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(OwnerPausableUpgradeableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *OwnerPausableUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnerPausableUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnerPausableUpgradeableUnpaused represents a Unpaused event raised by the OwnerPausableUpgradeable contract.
type OwnerPausableUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OwnerPausableUpgradeableUnpausedIterator, error) {

	logs, sub, err := _OwnerPausableUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OwnerPausableUpgradeableUnpausedIterator{contract: _OwnerPausableUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OwnerPausableUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _OwnerPausableUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnerPausableUpgradeableUnpaused)
				if err := _OwnerPausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OwnerPausableUpgradeable *OwnerPausableUpgradeableFilterer) ParseUnpaused(log types.Log) (*OwnerPausableUpgradeableUnpaused, error) {
	event := new(OwnerPausableUpgradeableUnpaused)
	if err := _OwnerPausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUpgradeableMetaData contains all meta data concerning the PausableUpgradeable contract.
var PausableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c975abb": "paused()",
	},
}

// PausableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableUpgradeableMetaData.ABI instead.
var PausableUpgradeableABI = PausableUpgradeableMetaData.ABI

// Deprecated: Use PausableUpgradeableMetaData.Sigs instead.
// PausableUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var PausableUpgradeableFuncSigs = PausableUpgradeableMetaData.Sigs

// PausableUpgradeable is an auto generated Go binding around an Ethereum contract.
type PausableUpgradeable struct {
	PausableUpgradeableCaller     // Read-only binding to the contract
	PausableUpgradeableTransactor // Write-only binding to the contract
	PausableUpgradeableFilterer   // Log filterer for contract events
}

// PausableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableUpgradeableSession struct {
	Contract     *PausableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PausableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableUpgradeableCallerSession struct {
	Contract *PausableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// PausableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableUpgradeableTransactorSession struct {
	Contract     *PausableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// PausableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableUpgradeableRaw struct {
	Contract *PausableUpgradeable // Generic contract binding to access the raw methods on
}

// PausableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableUpgradeableCallerRaw struct {
	Contract *PausableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableUpgradeableTransactorRaw struct {
	Contract *PausableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausableUpgradeable creates a new instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeable(address common.Address, backend bind.ContractBackend) (*PausableUpgradeable, error) {
	contract, err := bindPausableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeable{PausableUpgradeableCaller: PausableUpgradeableCaller{contract: contract}, PausableUpgradeableTransactor: PausableUpgradeableTransactor{contract: contract}, PausableUpgradeableFilterer: PausableUpgradeableFilterer{contract: contract}}, nil
}

// NewPausableUpgradeableCaller creates a new read-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*PausableUpgradeableCaller, error) {
	contract, err := bindPausableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableCaller{contract: contract}, nil
}

// NewPausableUpgradeableTransactor creates a new write-only instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableUpgradeableTransactor, error) {
	contract, err := bindPausableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableTransactor{contract: contract}, nil
}

// NewPausableUpgradeableFilterer creates a new log filterer instance of PausableUpgradeable, bound to a specific deployed contract.
func NewPausableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableUpgradeableFilterer, error) {
	contract, err := bindPausableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableFilterer{contract: contract}, nil
}

// bindPausableUpgradeable binds a generic wrapper to an already deployed contract.
func bindPausableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.PausableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.PausableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PausableUpgradeable *PausableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PausableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PausableUpgradeable *PausableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PausableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PausableUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PausableUpgradeable *PausableUpgradeableCallerSession) Paused() (bool, error) {
	return _PausableUpgradeable.Contract.Paused(&_PausableUpgradeable.CallOpts)
}

// PausableUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PausableUpgradeable contract.
type PausableUpgradeablePausedIterator struct {
	Event *PausableUpgradeablePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausableUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeablePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausableUpgradeablePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausableUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeablePaused represents a Paused event raised by the PausableUpgradeable contract.
type PausableUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*PausableUpgradeablePausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeablePausedIterator{contract: _PausableUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeablePaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParsePaused(log types.Log) (*PausableUpgradeablePaused, error) {
	event := new(PausableUpgradeablePaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpausedIterator struct {
	Event *PausableUpgradeableUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PausableUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUpgradeableUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PausableUpgradeableUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PausableUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUpgradeableUnpaused represents a Unpaused event raised by the PausableUpgradeable contract.
type PausableUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PausableUpgradeableUnpausedIterator, error) {

	logs, sub, err := _PausableUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PausableUpgradeableUnpausedIterator{contract: _PausableUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PausableUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _PausableUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUpgradeableUnpaused)
				if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PausableUpgradeable *PausableUpgradeableFilterer) ParseUnpaused(log types.Log) (*PausableUpgradeableUnpaused, error) {
	event := new(PausableUpgradeableUnpaused)
	if err := _PausableUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReentrancyGuardUpgradeableMetaData contains all meta data concerning the ReentrancyGuardUpgradeable contract.
var ReentrancyGuardUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
}

// ReentrancyGuardUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ReentrancyGuardUpgradeableMetaData.ABI instead.
var ReentrancyGuardUpgradeableABI = ReentrancyGuardUpgradeableMetaData.ABI

// ReentrancyGuardUpgradeable is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeable struct {
	ReentrancyGuardUpgradeableCaller     // Read-only binding to the contract
	ReentrancyGuardUpgradeableTransactor // Write-only binding to the contract
	ReentrancyGuardUpgradeableFilterer   // Log filterer for contract events
}

// ReentrancyGuardUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardUpgradeableSession struct {
	Contract     *ReentrancyGuardUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardUpgradeableCallerSession struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// ReentrancyGuardUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardUpgradeableTransactorSession struct {
	Contract     *ReentrancyGuardUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// ReentrancyGuardUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableRaw struct {
	Contract *ReentrancyGuardUpgradeable // Generic contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableCallerRaw struct {
	Contract *ReentrancyGuardUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardUpgradeableTransactorRaw struct {
	Contract *ReentrancyGuardUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuardUpgradeable creates a new instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeable(address common.Address, backend bind.ContractBackend) (*ReentrancyGuardUpgradeable, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeable{ReentrancyGuardUpgradeableCaller: ReentrancyGuardUpgradeableCaller{contract: contract}, ReentrancyGuardUpgradeableTransactor: ReentrancyGuardUpgradeableTransactor{contract: contract}, ReentrancyGuardUpgradeableFilterer: ReentrancyGuardUpgradeableFilterer{contract: contract}}, nil
}

// NewReentrancyGuardUpgradeableCaller creates a new read-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardUpgradeableCaller, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableCaller{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableTransactor creates a new write-only instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardUpgradeableTransactor, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableTransactor{contract: contract}, nil
}

// NewReentrancyGuardUpgradeableFilterer creates a new log filterer instance of ReentrancyGuardUpgradeable, bound to a specific deployed contract.
func NewReentrancyGuardUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardUpgradeableFilterer, error) {
	contract, err := bindReentrancyGuardUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardUpgradeableFilterer{contract: contract}, nil
}

// bindReentrancyGuardUpgradeable binds a generic wrapper to an already deployed contract.
func bindReentrancyGuardUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.ReentrancyGuardUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuardUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuardUpgradeable *ReentrancyGuardUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuardUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122070735c410cb7c388bf2f5e577192a2cb6f39e581c6a16b3b51db7d2cf2a5fa2564736f6c634300060c0033",
}

// SafeERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20MetaData.ABI instead.
var SafeERC20ABI = SafeERC20MetaData.ABI

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20MetaData.Bin instead.
var SafeERC20Bin = SafeERC20MetaData.Bin

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fd964299fa0cb417ed0159037d108b77b9428ff98dcba2ab8a81daac352029a364736f6c634300060c0033",
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

// SafeMathUpgradeableMetaData contains all meta data concerning the SafeMathUpgradeable contract.
var SafeMathUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204413a1e9963614616b0998d22fff4dbf20b53757e8dd97aa71ecdefe35f1e07864736f6c634300060c0033",
}

// SafeMathUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathUpgradeableMetaData.ABI instead.
var SafeMathUpgradeableABI = SafeMathUpgradeableMetaData.ABI

// SafeMathUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathUpgradeableMetaData.Bin instead.
var SafeMathUpgradeableBin = SafeMathUpgradeableMetaData.Bin

// DeploySafeMathUpgradeable deploys a new Ethereum contract, binding an instance of SafeMathUpgradeable to it.
func DeploySafeMathUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMathUpgradeable, error) {
	parsed, err := SafeMathUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMathUpgradeable{SafeMathUpgradeableCaller: SafeMathUpgradeableCaller{contract: contract}, SafeMathUpgradeableTransactor: SafeMathUpgradeableTransactor{contract: contract}, SafeMathUpgradeableFilterer: SafeMathUpgradeableFilterer{contract: contract}}, nil
}

// SafeMathUpgradeable is an auto generated Go binding around an Ethereum contract.
type SafeMathUpgradeable struct {
	SafeMathUpgradeableCaller     // Read-only binding to the contract
	SafeMathUpgradeableTransactor // Write-only binding to the contract
	SafeMathUpgradeableFilterer   // Log filterer for contract events
}

// SafeMathUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathUpgradeableSession struct {
	Contract     *SafeMathUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeMathUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathUpgradeableCallerSession struct {
	Contract *SafeMathUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SafeMathUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathUpgradeableTransactorSession struct {
	Contract     *SafeMathUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SafeMathUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathUpgradeableRaw struct {
	Contract *SafeMathUpgradeable // Generic contract binding to access the raw methods on
}

// SafeMathUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathUpgradeableCallerRaw struct {
	Contract *SafeMathUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathUpgradeableTransactorRaw struct {
	Contract *SafeMathUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMathUpgradeable creates a new instance of SafeMathUpgradeable, bound to a specific deployed contract.
func NewSafeMathUpgradeable(address common.Address, backend bind.ContractBackend) (*SafeMathUpgradeable, error) {
	contract, err := bindSafeMathUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMathUpgradeable{SafeMathUpgradeableCaller: SafeMathUpgradeableCaller{contract: contract}, SafeMathUpgradeableTransactor: SafeMathUpgradeableTransactor{contract: contract}, SafeMathUpgradeableFilterer: SafeMathUpgradeableFilterer{contract: contract}}, nil
}

// NewSafeMathUpgradeableCaller creates a new read-only instance of SafeMathUpgradeable, bound to a specific deployed contract.
func NewSafeMathUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*SafeMathUpgradeableCaller, error) {
	contract, err := bindSafeMathUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathUpgradeableCaller{contract: contract}, nil
}

// NewSafeMathUpgradeableTransactor creates a new write-only instance of SafeMathUpgradeable, bound to a specific deployed contract.
func NewSafeMathUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathUpgradeableTransactor, error) {
	contract, err := bindSafeMathUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathUpgradeableTransactor{contract: contract}, nil
}

// NewSafeMathUpgradeableFilterer creates a new log filterer instance of SafeMathUpgradeable, bound to a specific deployed contract.
func NewSafeMathUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathUpgradeableFilterer, error) {
	contract, err := bindSafeMathUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathUpgradeableFilterer{contract: contract}, nil
}

// bindSafeMathUpgradeable binds a generic wrapper to an already deployed contract.
func bindSafeMathUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathUpgradeable *SafeMathUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathUpgradeable.Contract.SafeMathUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathUpgradeable *SafeMathUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathUpgradeable.Contract.SafeMathUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathUpgradeable *SafeMathUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathUpgradeable.Contract.SafeMathUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMathUpgradeable *SafeMathUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMathUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMathUpgradeable *SafeMathUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMathUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMathUpgradeable *SafeMathUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMathUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// SwapMetaData contains all meta data concerning the Swap contract.
var SwapMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invariant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAdminFee\",\"type\":\"uint256\"}],\"name\":\"NewAdminFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSwapFee\",\"type\":\"uint256\"}],\"name\":\"NewSwapFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"futureTime\",\"type\":\"uint256\"}],\"name\":\"RampA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invariant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityImbalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boughtId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityOne\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StopRampA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"soldId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"boughtId\",\"type\":\"uint128\"}],\"name\":\"TokenSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateRemoveLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"deposit\",\"type\":\"bool\"}],\"name\":\"calculateTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAPrecise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAdminBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVirtualPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"_pooledTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"decimals\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"lpTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lpTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_adminFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpTokenTargetAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"futureA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureTime\",\"type\":\"uint256\"}],\"name\":\"rampA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"maxBurnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAdminFee\",\"type\":\"uint256\"}],\"name\":\"setAdminFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSwapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopRampA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialATime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureATime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adminFee\",\"type\":\"uint256\"},{\"internalType\":\"contractLPToken\",\"name\":\"lpToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAdminFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"4d49e87d": "addLiquidity(uint256[],uint256,uint256)",
		"f2fad2b6": "calculateRemoveLiquidity(uint256)",
		"342a87a1": "calculateRemoveLiquidityOneToken(uint256,uint8)",
		"a95b089f": "calculateSwap(uint8,uint8,uint256)",
		"e6ab2806": "calculateTokenAmount(uint256[],bool)",
		"d46300fd": "getA()",
		"0ba81959": "getAPrecise()",
		"ef0a712f": "getAdminBalance(uint256)",
		"82b86600": "getToken(uint8)",
		"91ceb3eb": "getTokenBalance(uint8)",
		"66c0bd24": "getTokenIndex(address)",
		"e25aa5fa": "getVirtualPrice()",
		"b28cb6dc": "initialize(address[],uint8[],string,string,uint256,uint256,uint256,address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"593d132c": "rampA(uint256,uint256)",
		"31cd52b0": "removeLiquidity(uint256,uint256[],uint256)",
		"84cdd9bc": "removeLiquidityImbalance(uint256[],uint256,uint256)",
		"3e3a1560": "removeLiquidityOneToken(uint256,uint8,uint256,uint256)",
		"715018a6": "renounceOwnership()",
		"8beb60b6": "setAdminFee(uint256)",
		"34e19907": "setSwapFee(uint256)",
		"c4db7fa0": "stopRampA()",
		"91695586": "swap(uint8,uint8,uint256,uint256,uint256)",
		"5fd65f0f": "swapStorage()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"0419b45a": "withdrawAdminFees()",
	},
	Bin: "0x608060405234801561001057600080fd5b506133f6806100206000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80638456cb5911610104578063b28cb6dc116100a2578063e6ab280611610071578063e6ab280614610885578063ef0a712f146108f7578063f2fad2b614610914578063f2fde38b14610931576101cf565b8063b28cb6dc14610610578063c4db7fa01461086d578063d46300fd14610875578063e25aa5fa1461087d576101cf565b80638da5cb5b116100de5780638da5cb5b1461057c578063916955861461058457806391ceb3eb146105c0578063a95b089f146105e0576101cf565b80638456cb59146104e157806384cdd9bc146104e95780638beb60b61461055f576101cf565b80634d49e87d116101715780635fd65f0f1161014b5780635fd65f0f146103f157806366c0bd2414610447578063715018a61461049057806382b8660014610498576101cf565b80634d49e87d1461033c578063593d132c146103b25780635c975abb146103d5576101cf565b8063342a87a1116101ad578063342a87a1146102bf57806334e19907146102e55780633e3a1560146103025780633f4ba83a14610334576101cf565b80630419b45a146101d45780630ba81959146101de57806331cd52b0146101f8575b600080fd5b6101dc610964565b005b6101e6610a81565b60408051918252519081900360200190f35b61026f6004803603606081101561020e57600080fd5b8135919081019060408101602082013564010000000081111561023057600080fd5b82018360208201111561024257600080fd5b8035906020019184602083028401116401000000008311171561026457600080fd5b919350915035610b06565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156102ab578181015183820152602001610293565b505050509050019250505060405180910390f35b6101e6600480360360408110156102d557600080fd5b508035906020013560ff16610d43565b6101dc600480360360208110156102fb57600080fd5b5035610df7565b6101e66004803603608081101561031857600080fd5b5080359060ff6020820135169060408101359060600135610f12565b6101dc6110de565b6101e66004803603606081101561035257600080fd5b81019060208101813564010000000081111561036d57600080fd5b82018360208201111561037f57600080fd5b803590602001918460208302840111640100000000831117156103a157600080fd5b919350915080359060200135611176565b6101dc600480360360408110156103c857600080fd5b5080359060200135611316565b6103dd611439565b604080519115158252519081900360200190f35b6103f9611442565b604080519788526020880196909652868601949094526060860192909252608085015260a084015273ffffffffffffffffffffffffffffffffffffffff1660c0830152519081900360e00190f35b61047a6004803603602081101561045d57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16611470565b6040805160ff9092168252519081900360200190f35b6101dc611517565b6104b8600480360360208110156104ae57600080fd5b503560ff16611614565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101dc6116a9565b6101e6600480360360608110156104ff57600080fd5b81019060208101813564010000000081111561051a57600080fd5b82018360208201111561052c57600080fd5b8035906020019184602083028401116401000000008311171561054e57600080fd5b91935091508035906020013561173f565b6101dc6004803603602081101561057557600080fd5b50356118df565b6104b86119df565b6101e6600480360360a081101561059a57600080fd5b5060ff8135811691602081013590911690604081013590606081013590608001356119fb565b6101e6600480360360208110156105d657600080fd5b503560ff16611bd0565b6101e6600480360360608110156105f657600080fd5b5060ff813581169160208101359091169060400135611c4f565b6101dc600480360361010081101561062757600080fd5b81019060208101813564010000000081111561064257600080fd5b82018360208201111561065457600080fd5b8035906020019184602083028401116401000000008311171561067657600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156106c657600080fd5b8201836020820111156106d857600080fd5b803590602001918460208302840111640100000000831117156106fa57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561074a57600080fd5b82018360208201111561075c57600080fd5b8035906020019184600183028401116401000000008311171561077e57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156107d157600080fd5b8201836020820111156107e357600080fd5b8035906020019184600183028401116401000000008311171561080557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040810135906060013573ffffffffffffffffffffffffffffffffffffffff16611d0a565b6101dc61256e565b6101e6612667565b6101e66126bb565b6101e66004803603604081101561089b57600080fd5b8101906020810181356401000000008111156108b657600080fd5b8201836020820111156108c857600080fd5b803590602001918460208302840111640100000000831117156108ea57600080fd5b919350915035151561270f565b6101e66004803603602081101561090d57600080fd5b50356127a4565b61026f6004803603602081101561092a57600080fd5b5035612832565b6101dc6004803603602081101561094757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661297c565b61096c612aea565b73ffffffffffffffffffffffffffffffffffffffff1661098a6119df565b73ffffffffffffffffffffffffffffffffffffffff16146109f2576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73__$fb68600ab723a8a0a0b96b45bbac9939e0$__6324c5c75160c9610a166119df565b6040518363ffffffff1660e01b8152600401808381526020018273ffffffffffffffffffffffffffffffffffffffff1681526020019250505060006040518083038186803b158015610a6757600080fd5b505af4158015610a7b573d6000803e3d6000fd5b50505050565b600060c973__$34a89a4c622fc2e024417b7c254097dd7b$__63c9b64dcb90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610ad557600080fd5b505af4158015610ae9573d6000803e3d6000fd5b505050506040513d6020811015610aff57600080fd5b5051905090565b606060026097541415610b60576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026097558142811015610bbb576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b60c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6373fd6b3e90918888886040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060006040518083038186803b158015610c4c57600080fd5b505af4158015610c60573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526020811015610ca757600080fd5b8101908080516040519392919084640100000000821115610cc757600080fd5b908301906020820185811115610cdc57600080fd5b8251866020820283011164010000000082111715610cf957600080fd5b82525081516020918201928201910280838360005b83811015610d26578181015183820152602001610d0e565b505050509050016040525050509150506001609755949350505050565b604080517fe069274200000000000000000000000000000000000000000000000000000000815260c960048201526024810184905260ff83166044820152905160009173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163e069274291606480820192602092909190829003018186803b158015610dc257600080fd5b505af4158015610dd6573d6000803e3d6000fd5b505050506040513d6020811015610dec57600080fd5b505190505b92915050565b610dff612aea565b73ffffffffffffffffffffffffffffffffffffffff16610e1d6119df565b73ffffffffffffffffffffffffffffffffffffffff1614610e85576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517f467e186c00000000000000000000000000000000000000000000000000000000815260c9600482015260248101839052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163467e186c916044808301926000929190829003018186803b158015610ef757600080fd5b505af4158015610f0b573d6000803e3d6000fd5b5050505050565b600060026097541415610f6c576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002609755610f79611439565b15610fcb576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b8180421115611021576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b604080517fe7a4db8100000000000000000000000000000000000000000000000000000000815260c960048201526024810188905260ff8716604482015260648101869052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163e7a4db81916084808301926020929190829003018186803b1580156110a357600080fd5b505af41580156110b7573d6000803e3d6000fd5b505050506040513d60208110156110cd57600080fd5b505160016097559695505050505050565b6110e6612aea565b73ffffffffffffffffffffffffffffffffffffffff166111046119df565b73ffffffffffffffffffffffffffffffffffffffff161461116c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b611174612aee565b565b6000600260975414156111d0576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026097556111dd611439565b1561122f576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b8180421115611285576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b60c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6340370edf90918888886040518563ffffffff1660e01b815260040180858152602001806020018381526020018281038252858582818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060206040518083038186803b1580156110a357600080fd5b61131e612aea565b73ffffffffffffffffffffffffffffffffffffffff1661133c6119df565b73ffffffffffffffffffffffffffffffffffffffff16146113a4576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517f58fdd79b00000000000000000000000000000000000000000000000000000000815260c960048201526024810184905260448101839052905173__$34a89a4c622fc2e024417b7c254097dd7b$__916358fdd79b916064808301926000929190829003018186803b15801561141d57600080fd5b505af4158015611431573d6000803e3d6000fd5b505050505050565b60655460ff1690565b60c95460ca5460cb5460cc5460cd5460ce5460cf5473ffffffffffffffffffffffffffffffffffffffff1687565b73ffffffffffffffffffffffffffffffffffffffff8116600081815260d36020526040812054909160ff909116906114a782611614565b73ffffffffffffffffffffffffffffffffffffffff161461150f576040805162461bcd60e51b815260206004820152601460248201527f546f6b656e20646f6573206e6f74206578697374000000000000000000000000604482015290519081900360640190fd5b90505b919050565b61151f612aea565b73ffffffffffffffffffffffffffffffffffffffff1661153d6119df565b73ffffffffffffffffffffffffffffffffffffffff16146115a5576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60335460405160009173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60d05460009060ff831610611670576040805162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e67650000000000000000000000000000000000000000604482015290519081900360640190fd5b60d0805460ff841690811061168157fe5b60009182526020909120015473ffffffffffffffffffffffffffffffffffffffff1692915050565b6116b1612aea565b73ffffffffffffffffffffffffffffffffffffffff166116cf6119df565b73ffffffffffffffffffffffffffffffffffffffff1614611737576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b611174612ba4565b600060026097541415611799576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026097556117a6611439565b156117f8576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b818042111561184e576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b60c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6341b91c2690918888886040518563ffffffff1660e01b815260040180858152602001806020018381526020018281038252858582818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060206040518083038186803b1580156110a357600080fd5b6118e7612aea565b73ffffffffffffffffffffffffffffffffffffffff166119056119df565b73ffffffffffffffffffffffffffffffffffffffff161461196d576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517f7046727600000000000000000000000000000000000000000000000000000000815260c9600482015260248101839052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__916370467276916044808301926000929190829003018186803b158015610ef757600080fd5b60335473ffffffffffffffffffffffffffffffffffffffff1690565b600060026097541415611a55576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002609755611a62611439565b15611ab4576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b8180421115611b0a576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b604080517fa5397b2200000000000000000000000000000000000000000000000000000000815260c9600482015260ff808a166024830152881660448201526064810187905260848101869052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163a5397b229160a4808301926020929190829003018186803b158015611b9457600080fd5b505af4158015611ba8573d6000803e3d6000fd5b505050506040513d6020811015611bbe57600080fd5b50516001609755979650505050505050565b60d05460009060ff831610611c2c576040805162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e67650000000000000000000000000000604482015290519081900360640190fd5b60d2805460ff8416908110611c3d57fe5b90600052602060002001549050919050565b604080517f4b23603c00000000000000000000000000000000000000000000000000000000815260c9600482015260ff80861660248301528416604482015260648101839052905160009173__$fb68600ab723a8a0a0b96b45bbac9939e0$__91634b23603c91608480820192602092909190829003018186803b158015611cd657600080fd5b505af4158015611cea573d6000803e3d6000fd5b505050506040513d6020811015611d0057600080fd5b5051949350505050565b600054610100900460ff1680611d235750611d23612c34565b80611d31575060005460ff16155b611d6c5760405162461bcd60e51b815260040180806020018281038252602e815260200180613372602e913960400191505060405180910390fd5b600054610100900460ff16158015611d97576000805460ff1961ff0019909116610100171660011790555b611d9f612c45565b611da7612cff565b6001895111611dfd576040805162461bcd60e51b815260206004820152601960248201527f5f706f6f6c6564546f6b656e732e6c656e677468203c3d203100000000000000604482015290519081900360640190fd5b602089511115611e54576040805162461bcd60e51b815260206004820152601960248201527f5f706f6f6c6564546f6b656e732e6c656e677468203e20333200000000000000604482015290519081900360640190fd5b8751895114611eaa576040805162461bcd60e51b815260206004820152601f60248201527f5f706f6f6c6564546f6b656e7320646563696d616c73206d69736d6174636800604482015290519081900360640190fd5b6060885167ffffffffffffffff81118015611ec457600080fd5b50604051908082528060200260200182016040528015611eee578160200160208202803683370190505b50905060005b8a518160ff1610156121ba5760ff8116156120045760d360008c8360ff1681518110611f1c57fe5b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff1682528101919091526040016000205460ff16158015611fb357508a8160ff1681518110611f6757fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff168b600081518110611f9257fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1614155b612004576040805162461bcd60e51b815260206004820152601060248201527f4475706c696361746520746f6b656e7300000000000000000000000000000000604482015290519081900360640190fd5b600073ffffffffffffffffffffffffffffffffffffffff168b8260ff168151811061202b57fe5b602002602001015173ffffffffffffffffffffffffffffffffffffffff16141561209c576040805162461bcd60e51b815260206004820152601d60248201527f546865203020616464726573732069736e277420616e204552432d3230000000604482015290519081900360640190fd5b601260ff168a8260ff16815181106120b057fe5b602002602001015160ff16111561210e576040805162461bcd60e51b815260206004820152601a60248201527f546f6b656e20646563696d616c732065786365656473206d6178000000000000604482015290519081900360640190fd5b61213e8a8260ff168151811061212057fe5b602002602001015160ff16601260ff16612d9490919063ffffffff16565b600a0a828260ff168151811061215057fe5b6020026020010181815250508060d360008d8460ff168151811061217057fe5b60209081029190910181015173ffffffffffffffffffffffffffffffffffffffff168252810191909152604001600020805460ff191660ff92909216919091179055600101611ef4565b50620f42408610612212576040805162461bcd60e51b815260206004820152601260248201527f5f612065786365656473206d6178696d756d0000000000000000000000000000604482015290519081900360640190fd5b6305f5e100851061226a576040805162461bcd60e51b815260206004820152601460248201527f5f6665652065786365656473206d6178696d756d000000000000000000000000604482015290519081900360640190fd5b6402540be40084106122c3576040805162461bcd60e51b815260206004820152601960248201527f5f61646d696e4665652065786365656473206d6178696d756d00000000000000604482015290519081900360640190fd5b60006122ce84612df1565b90508073ffffffffffffffffffffffffffffffffffffffff16634cd88b768a8a6040518363ffffffff1660e01b8152600401808060200180602001838103835285818151815260200191508051906020019080838360005b8381101561233e578181015183820152602001612326565b50505050905090810190601f16801561236b5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561239e578181015183820152602001612386565b50505050905090810190601f1680156123cb5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b1580156123ec57600080fd5b505af1158015612400573d6000803e3d6000fd5b505050506040513d602081101561241657600080fd5b5051612469576040805162461bcd60e51b815260206004820152601c60248201527f636f756c64206e6f7420696e6974206c70546f6b656e20636c6f6e6500000000604482015290519081900360640190fd5b60cf80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83161790558a516124bc9060d09060208e019061322e565b5081516124d09060d19060208501906132b8565b508a5167ffffffffffffffff811180156124e957600080fd5b50604051908082528060200260200182016040528015612513578160200160208202803683370190505b5080516125289160d2916020909101906132b8565b50612534876064612eb9565b60c955612542876064612eb9565b60ca55505060cd84905560ce8390558015612563576000805461ff00191690555b505050505050505050565b612576612aea565b73ffffffffffffffffffffffffffffffffffffffff166125946119df565b73ffffffffffffffffffffffffffffffffffffffff16146125fc576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517ff14e211e00000000000000000000000000000000000000000000000000000000815260c96004820152905173__$34a89a4c622fc2e024417b7c254097dd7b$__9163f14e211e916024808301926000929190829003018186803b158015610a6757600080fd5b600060c973__$34a89a4c622fc2e024417b7c254097dd7b$__63b0a14cfc90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610ad557600080fd5b600060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6371906c2c90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610ad557600080fd5b600060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__63834b491090918686866040518563ffffffff1660e01b8152600401808581526020018060200183151581526020018281038252858582818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060206040518083038186803b158015611cd657600080fd5b600060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__637d0481609091846040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b15801561280057600080fd5b505af4158015612814573d6000803e3d6000fd5b505050506040513d602081101561282a57600080fd5b505192915050565b606060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6370703e4a9091846040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561288e57600080fd5b505af41580156128a2573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405260208110156128e957600080fd5b810190808051604051939291908464010000000082111561290957600080fd5b90830190602082018581111561291e57600080fd5b825186602082028301116401000000008211171561293b57600080fd5b82525081516020918201928201910280838360005b83811015612968578181015183820152602001612950565b505050509050016040525050509050919050565b612984612aea565b73ffffffffffffffffffffffffffffffffffffffff166129a26119df565b73ffffffffffffffffffffffffffffffffffffffff1614612a0a576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff8116612a5c5760405162461bcd60e51b815260040180806020018281038252602681526020018061334c6026913960400191505060405180910390fd5b60335460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b3390565b612af6611439565b612b47576040805162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015290519081900360640190fd5b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa612b7a612aea565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190a1565b612bac611439565b15612bfe576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258612b7a612aea565b6000612c3f30612f19565b15905090565b600054610100900460ff1680612c5e5750612c5e612c34565b80612c6c575060005460ff16155b612ca75760405162461bcd60e51b815260040180806020018281038252602e815260200180613372602e913960400191505060405180910390fd5b600054610100900460ff16158015612cd2576000805460ff1961ff0019909116610100171660011790555b612cda612f1f565b612ce2612fbf565b612cea6130dd565b8015612cfc576000805461ff00191690555b50565b600054610100900460ff1680612d185750612d18612c34565b80612d26575060005460ff16155b612d615760405162461bcd60e51b815260040180806020018281038252602e815260200180613372602e913960400191505060405180910390fd5b600054610100900460ff16158015612d8c576000805460ff1961ff0019909116610100171660011790555b612cea613188565b600082821115612deb576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b60006040517f3d602d80600a3d3981f3363d3d373d3d3d363d7300000000000000000000000081528260601b60148201527f5af43d82803e903d91602b57fd5bf3000000000000000000000000000000000060288201526037816000f091505073ffffffffffffffffffffffffffffffffffffffff8116611512576040805162461bcd60e51b815260206004820152601660248201527f455243313136373a20637265617465206661696c656400000000000000000000604482015290519081900360640190fd5b600082612ec857506000610df1565b82820282848281612ed557fe5b0414612f125760405162461bcd60e51b81526004018080602001828103825260218152602001806133a06021913960400191505060405180910390fd5b9392505050565b3b151590565b600054610100900460ff1680612f385750612f38612c34565b80612f46575060005460ff16155b612f815760405162461bcd60e51b815260040180806020018281038252602e815260200180613372602e913960400191505060405180910390fd5b600054610100900460ff16158015612cea576000805460ff1961ff0019909116610100171660011790558015612cfc576000805461ff001916905550565b600054610100900460ff1680612fd85750612fd8612c34565b80612fe6575060005460ff16155b6130215760405162461bcd60e51b815260040180806020018281038252602e815260200180613372602e913960400191505060405180910390fd5b600054610100900460ff1615801561304c576000805460ff1961ff0019909116610100171660011790555b6000613056612aea565b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a3508015612cfc576000805461ff001916905550565b600054610100900460ff16806130f657506130f6612c34565b80613104575060005460ff16155b61313f5760405162461bcd60e51b815260040180806020018281038252602e815260200180613372602e913960400191505060405180910390fd5b600054610100900460ff1615801561316a576000805460ff1961ff0019909116610100171660011790555b6065805460ff191690558015612cfc576000805461ff001916905550565b600054610100900460ff16806131a157506131a1612c34565b806131af575060005460ff16155b6131ea5760405162461bcd60e51b815260040180806020018281038252602e815260200180613372602e913960400191505060405180910390fd5b600054610100900460ff16158015613215576000805460ff1961ff0019909116610100171660011790555b60016097558015612cfc576000805461ff001916905550565b8280548282559060005260206000209081019282156132a8579160200282015b828111156132a857825182547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff90911617825560209092019160019091019061324e565b506132b49291506132ff565b5090565b8280548282559060005260206000209081019282156132f3579160200282015b828111156132f35782518255916020019190600101906132d8565b506132b4929150613336565b5b808211156132b45780547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155600101613300565b5b808211156132b4576000815560010161333756fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77a264697066735822122044a3c5676052ca688bd07a40762f95815819b51d89d675f1b3ce313b5389126564736f6c634300060c0033",
}

// SwapABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapMetaData.ABI instead.
var SwapABI = SwapMetaData.ABI

// Deprecated: Use SwapMetaData.Sigs instead.
// SwapFuncSigs maps the 4-byte function signature to its string representation.
var SwapFuncSigs = SwapMetaData.Sigs

// SwapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapMetaData.Bin instead.
var SwapBin = SwapMetaData.Bin

// DeploySwap deploys a new Ethereum contract, binding an instance of Swap to it.
func DeploySwap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Swap, error) {
	parsed, err := SwapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	amplificationUtilsAddr, _, _, _ := DeployAmplificationUtils(auth, backend)
	SwapBin = strings.ReplaceAll(SwapBin, "__$34a89a4c622fc2e024417b7c254097dd7b$__", amplificationUtilsAddr.String()[2:])

	swapUtilsAddr, _, _, _ := DeploySwapUtils(auth, backend)
	SwapBin = strings.ReplaceAll(SwapBin, "__$fb68600ab723a8a0a0b96b45bbac9939e0$__", swapUtilsAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// Swap is an auto generated Go binding around an Ethereum contract.
type Swap struct {
	SwapCaller     // Read-only binding to the contract
	SwapTransactor // Write-only binding to the contract
	SwapFilterer   // Log filterer for contract events
}

// SwapCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapSession struct {
	Contract     *Swap             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapCallerSession struct {
	Contract *SwapCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SwapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapTransactorSession struct {
	Contract     *SwapTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapRaw struct {
	Contract *Swap // Generic contract binding to access the raw methods on
}

// SwapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapCallerRaw struct {
	Contract *SwapCaller // Generic read-only contract binding to access the raw methods on
}

// SwapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapTransactorRaw struct {
	Contract *SwapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwap creates a new instance of Swap, bound to a specific deployed contract.
func NewSwap(address common.Address, backend bind.ContractBackend) (*Swap, error) {
	contract, err := bindSwap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Swap{SwapCaller: SwapCaller{contract: contract}, SwapTransactor: SwapTransactor{contract: contract}, SwapFilterer: SwapFilterer{contract: contract}}, nil
}

// NewSwapCaller creates a new read-only instance of Swap, bound to a specific deployed contract.
func NewSwapCaller(address common.Address, caller bind.ContractCaller) (*SwapCaller, error) {
	contract, err := bindSwap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapCaller{contract: contract}, nil
}

// NewSwapTransactor creates a new write-only instance of Swap, bound to a specific deployed contract.
func NewSwapTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapTransactor, error) {
	contract, err := bindSwap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapTransactor{contract: contract}, nil
}

// NewSwapFilterer creates a new log filterer instance of Swap, bound to a specific deployed contract.
func NewSwapFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFilterer, error) {
	contract, err := bindSwap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFilterer{contract: contract}, nil
}

// bindSwap binds a generic wrapper to an already deployed contract.
func bindSwap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.SwapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.SwapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Swap *SwapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Swap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Swap *SwapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Swap *SwapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Swap.Contract.contract.Transact(opts, method, params...)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_Swap *SwapCaller) CalculateRemoveLiquidity(opts *bind.CallOpts, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "calculateRemoveLiquidity", amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_Swap *SwapSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _Swap.Contract.CalculateRemoveLiquidity(&_Swap.CallOpts, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_Swap *SwapCallerSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _Swap.Contract.CalculateRemoveLiquidity(&_Swap.CallOpts, amount)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_Swap *SwapCaller) CalculateRemoveLiquidityOneToken(opts *bind.CallOpts, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "calculateRemoveLiquidityOneToken", tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_Swap *SwapSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _Swap.Contract.CalculateRemoveLiquidityOneToken(&_Swap.CallOpts, tokenAmount, tokenIndex)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_Swap *SwapCallerSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _Swap.Contract.CalculateRemoveLiquidityOneToken(&_Swap.CallOpts, tokenAmount, tokenIndex)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_Swap *SwapCaller) CalculateSwap(opts *bind.CallOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "calculateSwap", tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_Swap *SwapSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _Swap.Contract.CalculateSwap(&_Swap.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_Swap *SwapCallerSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _Swap.Contract.CalculateSwap(&_Swap.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_Swap *SwapCaller) CalculateTokenAmount(opts *bind.CallOpts, amounts []*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "calculateTokenAmount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_Swap *SwapSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _Swap.Contract.CalculateTokenAmount(&_Swap.CallOpts, amounts, deposit)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_Swap *SwapCallerSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _Swap.Contract.CalculateTokenAmount(&_Swap.CallOpts, amounts, deposit)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_Swap *SwapCaller) GetA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_Swap *SwapSession) GetA() (*big.Int, error) {
	return _Swap.Contract.GetA(&_Swap.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_Swap *SwapCallerSession) GetA() (*big.Int, error) {
	return _Swap.Contract.GetA(&_Swap.CallOpts)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_Swap *SwapCaller) GetAPrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getAPrecise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_Swap *SwapSession) GetAPrecise() (*big.Int, error) {
	return _Swap.Contract.GetAPrecise(&_Swap.CallOpts)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_Swap *SwapCallerSession) GetAPrecise() (*big.Int, error) {
	return _Swap.Contract.GetAPrecise(&_Swap.CallOpts)
}

// GetAdminBalance is a free data retrieval call binding the contract method 0xef0a712f.
//
// Solidity: function getAdminBalance(uint256 index) view returns(uint256)
func (_Swap *SwapCaller) GetAdminBalance(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getAdminBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAdminBalance is a free data retrieval call binding the contract method 0xef0a712f.
//
// Solidity: function getAdminBalance(uint256 index) view returns(uint256)
func (_Swap *SwapSession) GetAdminBalance(index *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetAdminBalance(&_Swap.CallOpts, index)
}

// GetAdminBalance is a free data retrieval call binding the contract method 0xef0a712f.
//
// Solidity: function getAdminBalance(uint256 index) view returns(uint256)
func (_Swap *SwapCallerSession) GetAdminBalance(index *big.Int) (*big.Int, error) {
	return _Swap.Contract.GetAdminBalance(&_Swap.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_Swap *SwapCaller) GetToken(opts *bind.CallOpts, index uint8) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_Swap *SwapSession) GetToken(index uint8) (common.Address, error) {
	return _Swap.Contract.GetToken(&_Swap.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_Swap *SwapCallerSession) GetToken(index uint8) (common.Address, error) {
	return _Swap.Contract.GetToken(&_Swap.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_Swap *SwapCaller) GetTokenBalance(opts *bind.CallOpts, index uint8) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getTokenBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_Swap *SwapSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _Swap.Contract.GetTokenBalance(&_Swap.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_Swap *SwapCallerSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _Swap.Contract.GetTokenBalance(&_Swap.CallOpts, index)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_Swap *SwapCaller) GetTokenIndex(opts *bind.CallOpts, tokenAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getTokenIndex", tokenAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_Swap *SwapSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _Swap.Contract.GetTokenIndex(&_Swap.CallOpts, tokenAddress)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_Swap *SwapCallerSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _Swap.Contract.GetTokenIndex(&_Swap.CallOpts, tokenAddress)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_Swap *SwapCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "getVirtualPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_Swap *SwapSession) GetVirtualPrice() (*big.Int, error) {
	return _Swap.Contract.GetVirtualPrice(&_Swap.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_Swap *SwapCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _Swap.Contract.GetVirtualPrice(&_Swap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Swap *SwapCallerSession) Owner() (common.Address, error) {
	return _Swap.Contract.Owner(&_Swap.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Swap *SwapCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Swap *SwapSession) Paused() (bool, error) {
	return _Swap.Contract.Paused(&_Swap.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Swap *SwapCallerSession) Paused() (bool, error) {
	return _Swap.Contract.Paused(&_Swap.CallOpts)
}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_Swap *SwapCaller) SwapStorage(opts *bind.CallOpts) (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	var out []interface{}
	err := _Swap.contract.Call(opts, &out, "swapStorage")

	outstruct := new(struct {
		InitialA     *big.Int
		FutureA      *big.Int
		InitialATime *big.Int
		FutureATime  *big.Int
		SwapFee      *big.Int
		AdminFee     *big.Int
		LpToken      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FutureA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitialATime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FutureATime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SwapFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AdminFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LpToken = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_Swap *SwapSession) SwapStorage() (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	return _Swap.Contract.SwapStorage(&_Swap.CallOpts)
}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_Swap *SwapCallerSession) SwapStorage() (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	return _Swap.Contract.SwapStorage(&_Swap.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactor) AddLiquidity(opts *bind.TransactOpts, amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "addLiquidity", amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_Swap *SwapSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AddLiquidity(&_Swap.TransactOpts, amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactorSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.AddLiquidity(&_Swap.TransactOpts, amounts, minToMint, deadline)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] _pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 _a, uint256 _fee, uint256 _adminFee, address lpTokenTargetAddress) returns()
func (_Swap *SwapTransactor) Initialize(opts *bind.TransactOpts, _pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, _a *big.Int, _fee *big.Int, _adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "initialize", _pooledTokens, decimals, lpTokenName, lpTokenSymbol, _a, _fee, _adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] _pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 _a, uint256 _fee, uint256 _adminFee, address lpTokenTargetAddress) returns()
func (_Swap *SwapSession) Initialize(_pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, _a *big.Int, _fee *big.Int, _adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _Swap.Contract.Initialize(&_Swap.TransactOpts, _pooledTokens, decimals, lpTokenName, lpTokenSymbol, _a, _fee, _adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] _pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 _a, uint256 _fee, uint256 _adminFee, address lpTokenTargetAddress) returns()
func (_Swap *SwapTransactorSession) Initialize(_pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, _a *big.Int, _fee *big.Int, _adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _Swap.Contract.Initialize(&_Swap.TransactOpts, _pooledTokens, decimals, lpTokenName, lpTokenSymbol, _a, _fee, _adminFee, lpTokenTargetAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Swap *SwapTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Swap *SwapSession) Pause() (*types.Transaction, error) {
	return _Swap.Contract.Pause(&_Swap.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Swap *SwapTransactorSession) Pause() (*types.Transaction, error) {
	return _Swap.Contract.Pause(&_Swap.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x593d132c.
//
// Solidity: function rampA(uint256 futureA, uint256 futureTime) returns()
func (_Swap *SwapTransactor) RampA(opts *bind.TransactOpts, futureA *big.Int, futureTime *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "rampA", futureA, futureTime)
}

// RampA is a paid mutator transaction binding the contract method 0x593d132c.
//
// Solidity: function rampA(uint256 futureA, uint256 futureTime) returns()
func (_Swap *SwapSession) RampA(futureA *big.Int, futureTime *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RampA(&_Swap.TransactOpts, futureA, futureTime)
}

// RampA is a paid mutator transaction binding the contract method 0x593d132c.
//
// Solidity: function rampA(uint256 futureA, uint256 futureTime) returns()
func (_Swap *SwapTransactorSession) RampA(futureA *big.Int, futureTime *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RampA(&_Swap.TransactOpts, futureA, futureTime)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_Swap *SwapTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "removeLiquidity", amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_Swap *SwapSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidity(&_Swap.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_Swap *SwapTransactorSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidity(&_Swap.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "removeLiquidityImbalance", amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_Swap *SwapSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidityImbalance(&_Swap.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactorSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidityImbalance(&_Swap.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactor) RemoveLiquidityOneToken(opts *bind.TransactOpts, tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "removeLiquidityOneToken", tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_Swap *SwapSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidityOneToken(&_Swap.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactorSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.RemoveLiquidityOneToken(&_Swap.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swap.Contract.RenounceOwnership(&_Swap.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Swap *SwapTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Swap.Contract.RenounceOwnership(&_Swap.TransactOpts)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x8beb60b6.
//
// Solidity: function setAdminFee(uint256 newAdminFee) returns()
func (_Swap *SwapTransactor) SetAdminFee(opts *bind.TransactOpts, newAdminFee *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "setAdminFee", newAdminFee)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x8beb60b6.
//
// Solidity: function setAdminFee(uint256 newAdminFee) returns()
func (_Swap *SwapSession) SetAdminFee(newAdminFee *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SetAdminFee(&_Swap.TransactOpts, newAdminFee)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x8beb60b6.
//
// Solidity: function setAdminFee(uint256 newAdminFee) returns()
func (_Swap *SwapTransactorSession) SetAdminFee(newAdminFee *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SetAdminFee(&_Swap.TransactOpts, newAdminFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 newSwapFee) returns()
func (_Swap *SwapTransactor) SetSwapFee(opts *bind.TransactOpts, newSwapFee *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "setSwapFee", newSwapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 newSwapFee) returns()
func (_Swap *SwapSession) SetSwapFee(newSwapFee *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SetSwapFee(&_Swap.TransactOpts, newSwapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 newSwapFee) returns()
func (_Swap *SwapTransactorSession) SetSwapFee(newSwapFee *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.SetSwapFee(&_Swap.TransactOpts, newSwapFee)
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_Swap *SwapTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "stopRampA")
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_Swap *SwapSession) StopRampA() (*types.Transaction, error) {
	return _Swap.Contract.StopRampA(&_Swap.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_Swap *SwapTransactorSession) StopRampA() (*types.Transaction, error) {
	return _Swap.Contract.StopRampA(&_Swap.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactor) Swap(opts *bind.TransactOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "swap", tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_Swap *SwapSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Swap(&_Swap.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_Swap *SwapTransactorSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Swap.Contract.Swap(&_Swap.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swap.Contract.TransferOwnership(&_Swap.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Swap *SwapTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Swap.Contract.TransferOwnership(&_Swap.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Swap *SwapTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Swap *SwapSession) Unpause() (*types.Transaction, error) {
	return _Swap.Contract.Unpause(&_Swap.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Swap *SwapTransactorSession) Unpause() (*types.Transaction, error) {
	return _Swap.Contract.Unpause(&_Swap.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x0419b45a.
//
// Solidity: function withdrawAdminFees() returns()
func (_Swap *SwapTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Swap.contract.Transact(opts, "withdrawAdminFees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x0419b45a.
//
// Solidity: function withdrawAdminFees() returns()
func (_Swap *SwapSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Swap.Contract.WithdrawAdminFees(&_Swap.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x0419b45a.
//
// Solidity: function withdrawAdminFees() returns()
func (_Swap *SwapTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _Swap.Contract.WithdrawAdminFees(&_Swap.TransactOpts)
}

// SwapAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Swap contract.
type SwapAddLiquidityIterator struct {
	Event *SwapAddLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapAddLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapAddLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapAddLiquidity represents a AddLiquidity event raised by the Swap contract.
type SwapAddLiquidity struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	Fees          []*big.Int
	Invariant     *big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SwapAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapAddLiquidityIterator{contract: _Swap.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *SwapAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapAddLiquidity)
				if err := _Swap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddLiquidity is a log parse operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) ParseAddLiquidity(log types.Log) (*SwapAddLiquidity, error) {
	event := new(SwapAddLiquidity)
	if err := _Swap.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapNewAdminFeeIterator is returned from FilterNewAdminFee and is used to iterate over the raw logs and unpacked data for NewAdminFee events raised by the Swap contract.
type SwapNewAdminFeeIterator struct {
	Event *SwapNewAdminFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapNewAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapNewAdminFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapNewAdminFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapNewAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapNewAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapNewAdminFee represents a NewAdminFee event raised by the Swap contract.
type SwapNewAdminFee struct {
	NewAdminFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewAdminFee is a free log retrieval operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_Swap *SwapFilterer) FilterNewAdminFee(opts *bind.FilterOpts) (*SwapNewAdminFeeIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "NewAdminFee")
	if err != nil {
		return nil, err
	}
	return &SwapNewAdminFeeIterator{contract: _Swap.contract, event: "NewAdminFee", logs: logs, sub: sub}, nil
}

// WatchNewAdminFee is a free log subscription operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_Swap *SwapFilterer) WatchNewAdminFee(opts *bind.WatchOpts, sink chan<- *SwapNewAdminFee) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "NewAdminFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapNewAdminFee)
				if err := _Swap.contract.UnpackLog(event, "NewAdminFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdminFee is a log parse operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_Swap *SwapFilterer) ParseNewAdminFee(log types.Log) (*SwapNewAdminFee, error) {
	event := new(SwapNewAdminFee)
	if err := _Swap.contract.UnpackLog(event, "NewAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapNewSwapFeeIterator is returned from FilterNewSwapFee and is used to iterate over the raw logs and unpacked data for NewSwapFee events raised by the Swap contract.
type SwapNewSwapFeeIterator struct {
	Event *SwapNewSwapFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapNewSwapFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapNewSwapFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapNewSwapFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapNewSwapFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapNewSwapFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapNewSwapFee represents a NewSwapFee event raised by the Swap contract.
type SwapNewSwapFee struct {
	NewSwapFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSwapFee is a free log retrieval operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_Swap *SwapFilterer) FilterNewSwapFee(opts *bind.FilterOpts) (*SwapNewSwapFeeIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "NewSwapFee")
	if err != nil {
		return nil, err
	}
	return &SwapNewSwapFeeIterator{contract: _Swap.contract, event: "NewSwapFee", logs: logs, sub: sub}, nil
}

// WatchNewSwapFee is a free log subscription operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_Swap *SwapFilterer) WatchNewSwapFee(opts *bind.WatchOpts, sink chan<- *SwapNewSwapFee) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "NewSwapFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapNewSwapFee)
				if err := _Swap.contract.UnpackLog(event, "NewSwapFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSwapFee is a log parse operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_Swap *SwapFilterer) ParseNewSwapFee(log types.Log) (*SwapNewSwapFee, error) {
	event := new(SwapNewSwapFee)
	if err := _Swap.contract.UnpackLog(event, "NewSwapFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Swap contract.
type SwapOwnershipTransferredIterator struct {
	Event *SwapOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapOwnershipTransferred represents a OwnershipTransferred event raised by the Swap contract.
type SwapOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapOwnershipTransferredIterator{contract: _Swap.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapOwnershipTransferred)
				if err := _Swap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Swap *SwapFilterer) ParseOwnershipTransferred(log types.Log) (*SwapOwnershipTransferred, error) {
	event := new(SwapOwnershipTransferred)
	if err := _Swap.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Swap contract.
type SwapPausedIterator struct {
	Event *SwapPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapPaused represents a Paused event raised by the Swap contract.
type SwapPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Swap *SwapFilterer) FilterPaused(opts *bind.FilterOpts) (*SwapPausedIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SwapPausedIterator{contract: _Swap.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Swap *SwapFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SwapPaused) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapPaused)
				if err := _Swap.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Swap *SwapFilterer) ParsePaused(log types.Log) (*SwapPaused, error) {
	event := new(SwapPaused)
	if err := _Swap.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the Swap contract.
type SwapRampAIterator struct {
	Event *SwapRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRampA represents a RampA event raised by the Swap contract.
type SwapRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_Swap *SwapFilterer) FilterRampA(opts *bind.FilterOpts) (*SwapRampAIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &SwapRampAIterator{contract: _Swap.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_Swap *SwapFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *SwapRampA) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRampA)
				if err := _Swap.contract.UnpackLog(event, "RampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_Swap *SwapFilterer) ParseRampA(log types.Log) (*SwapRampA, error) {
	event := new(SwapRampA)
	if err := _Swap.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Swap contract.
type SwapRemoveLiquidityIterator struct {
	Event *SwapRemoveLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRemoveLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapRemoveLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRemoveLiquidity represents a RemoveLiquidity event raised by the Swap contract.
type SwapRemoveLiquidity struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SwapRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapRemoveLiquidityIterator{contract: _Swap.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *SwapRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRemoveLiquidity)
				if err := _Swap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) ParseRemoveLiquidity(log types.Log) (*SwapRemoveLiquidity, error) {
	event := new(SwapRemoveLiquidity)
	if err := _Swap.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the Swap contract.
type SwapRemoveLiquidityImbalanceIterator struct {
	Event *SwapRemoveLiquidityImbalance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRemoveLiquidityImbalance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapRemoveLiquidityImbalance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the Swap contract.
type SwapRemoveLiquidityImbalance struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	Fees          []*big.Int
	Invariant     *big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*SwapRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapRemoveLiquidityImbalanceIterator{contract: _Swap.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *SwapRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRemoveLiquidityImbalance)
				if err := _Swap.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_Swap *SwapFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*SwapRemoveLiquidityImbalance, error) {
	event := new(SwapRemoveLiquidityImbalance)
	if err := _Swap.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the Swap contract.
type SwapRemoveLiquidityOneIterator struct {
	Event *SwapRemoveLiquidityOne // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapRemoveLiquidityOne)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapRemoveLiquidityOne)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the Swap contract.
type SwapRemoveLiquidityOne struct {
	Provider      common.Address
	LpTokenAmount *big.Int
	LpTokenSupply *big.Int
	BoughtId      *big.Int
	TokensBought  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_Swap *SwapFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*SwapRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapRemoveLiquidityOneIterator{contract: _Swap.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_Swap *SwapFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *SwapRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapRemoveLiquidityOne)
				if err := _Swap.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_Swap *SwapFilterer) ParseRemoveLiquidityOne(log types.Log) (*SwapRemoveLiquidityOne, error) {
	event := new(SwapRemoveLiquidityOne)
	if err := _Swap.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the Swap contract.
type SwapStopRampAIterator struct {
	Event *SwapStopRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapStopRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapStopRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapStopRampA represents a StopRampA event raised by the Swap contract.
type SwapStopRampA struct {
	CurrentA *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_Swap *SwapFilterer) FilterStopRampA(opts *bind.FilterOpts) (*SwapStopRampAIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &SwapStopRampAIterator{contract: _Swap.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_Swap *SwapFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *SwapStopRampA) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapStopRampA)
				if err := _Swap.contract.UnpackLog(event, "StopRampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_Swap *SwapFilterer) ParseStopRampA(log types.Log) (*SwapStopRampA, error) {
	event := new(SwapStopRampA)
	if err := _Swap.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapTokenSwapIterator is returned from FilterTokenSwap and is used to iterate over the raw logs and unpacked data for TokenSwap events raised by the Swap contract.
type SwapTokenSwapIterator struct {
	Event *SwapTokenSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapTokenSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapTokenSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapTokenSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapTokenSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapTokenSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapTokenSwap represents a TokenSwap event raised by the Swap contract.
type SwapTokenSwap struct {
	Buyer        common.Address
	TokensSold   *big.Int
	TokensBought *big.Int
	SoldId       *big.Int
	BoughtId     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenSwap is a free log retrieval operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_Swap *SwapFilterer) FilterTokenSwap(opts *bind.FilterOpts, buyer []common.Address) (*SwapTokenSwapIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Swap.contract.FilterLogs(opts, "TokenSwap", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SwapTokenSwapIterator{contract: _Swap.contract, event: "TokenSwap", logs: logs, sub: sub}, nil
}

// WatchTokenSwap is a free log subscription operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_Swap *SwapFilterer) WatchTokenSwap(opts *bind.WatchOpts, sink chan<- *SwapTokenSwap, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Swap.contract.WatchLogs(opts, "TokenSwap", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapTokenSwap)
				if err := _Swap.contract.UnpackLog(event, "TokenSwap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSwap is a log parse operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_Swap *SwapFilterer) ParseTokenSwap(log types.Log) (*SwapTokenSwap, error) {
	event := new(SwapTokenSwap)
	if err := _Swap.contract.UnpackLog(event, "TokenSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Swap contract.
type SwapUnpausedIterator struct {
	Event *SwapUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUnpaused represents a Unpaused event raised by the Swap contract.
type SwapUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Swap *SwapFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SwapUnpausedIterator, error) {

	logs, sub, err := _Swap.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SwapUnpausedIterator{contract: _Swap.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Swap *SwapFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SwapUnpaused) (event.Subscription, error) {

	logs, sub, err := _Swap.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUnpaused)
				if err := _Swap.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Swap *SwapFilterer) ParseUnpaused(log types.Log) (*SwapUnpaused, error) {
	event := new(SwapUnpaused)
	if err := _Swap.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanMetaData contains all meta data concerning the SwapFlashLoan contract.
var SwapFlashLoanMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invariant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAdminFee\",\"type\":\"uint256\"}],\"name\":\"NewAdminFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSwapFee\",\"type\":\"uint256\"}],\"name\":\"NewSwapFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"futureTime\",\"type\":\"uint256\"}],\"name\":\"RampA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invariant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityImbalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boughtId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityOne\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentA\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"StopRampA\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"soldId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"boughtId\",\"type\":\"uint128\"}],\"name\":\"TokenSwap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_BPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minToMint\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateRemoveLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableTokenAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool\",\"name\":\"deposit\",\"type\":\"bool\"}],\"name\":\"calculateTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"flashLoanFeeBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAPrecise\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAdminBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"index\",\"type\":\"uint8\"}],\"name\":\"getTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenIndex\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVirtualPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"_pooledTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint8[]\",\"name\":\"decimals\",\"type\":\"uint8[]\"},{\"internalType\":\"string\",\"name\":\"lpTokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"lpTokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_adminFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lpTokenTargetAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeShareBPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"futureA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureTime\",\"type\":\"uint256\"}],\"name\":\"rampA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"maxBurnAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityImbalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"removeLiquidityOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAdminFee\",\"type\":\"uint256\"}],\"name\":\"setAdminFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFlashLoanFeeBPS\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newProtocolFeeShareBPS\",\"type\":\"uint256\"}],\"name\":\"setFlashLoanFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSwapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopRampA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minDy\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapStorage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"initialA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialATime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"futureATime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adminFee\",\"type\":\"uint256\"},{\"internalType\":\"contractLPToken\",\"name\":\"lpToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawAdminFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"fd967f47": "MAX_BPS()",
		"4d49e87d": "addLiquidity(uint256[],uint256,uint256)",
		"f2fad2b6": "calculateRemoveLiquidity(uint256)",
		"342a87a1": "calculateRemoveLiquidityOneToken(uint256,uint8)",
		"a95b089f": "calculateSwap(uint8,uint8,uint256)",
		"e6ab2806": "calculateTokenAmount(uint256[],bool)",
		"5cffe9de": "flashLoan(address,address,uint256,bytes)",
		"7f1c825a": "flashLoanFeeBPS()",
		"d46300fd": "getA()",
		"0ba81959": "getAPrecise()",
		"ef0a712f": "getAdminBalance(uint256)",
		"82b86600": "getToken(uint8)",
		"91ceb3eb": "getTokenBalance(uint8)",
		"66c0bd24": "getTokenIndex(address)",
		"e25aa5fa": "getVirtualPrice()",
		"b28cb6dc": "initialize(address[],uint8[],string,string,uint256,uint256,uint256,address)",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"e8cc7fb6": "protocolFeeShareBPS()",
		"593d132c": "rampA(uint256,uint256)",
		"31cd52b0": "removeLiquidity(uint256,uint256[],uint256)",
		"84cdd9bc": "removeLiquidityImbalance(uint256[],uint256,uint256)",
		"3e3a1560": "removeLiquidityOneToken(uint256,uint8,uint256,uint256)",
		"715018a6": "renounceOwnership()",
		"8beb60b6": "setAdminFee(uint256)",
		"ef815967": "setFlashLoanFees(uint256,uint256)",
		"34e19907": "setSwapFee(uint256)",
		"c4db7fa0": "stopRampA()",
		"91695586": "swap(uint8,uint8,uint256,uint256,uint256)",
		"5fd65f0f": "swapStorage()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"0419b45a": "withdrawAdminFees()",
	},
	Bin: "0x608060405234801561001057600080fd5b50613dd3806100206000396000f3fe608060405234801561001057600080fd5b50600436106102265760003560e01c806384cdd9bc1161012a578063d46300fd116100bd578063ef0a712f1161008c578063f2fad2b611610071578063f2fad2b614610a30578063f2fde38b14610a4d578063fd967f4714610a7357610226565b8063ef0a712f146109f0578063ef81596714610a0d57610226565b8063d46300fd14610966578063e25aa5fa1461096e578063e6ab280614610976578063e8cc7fb6146109e857610226565b806391ceb3eb116100f957806391ceb3eb146106be578063a95b089f146106de578063b28cb6dc1461070e578063c4db7fa01461095e57610226565b806384cdd9bc146105e75780638beb60b61461065d5780638da5cb5b1461067a578063916955861461068257610226565b8063593d132c116101bd57806366c0bd241161018c5780637f1c825a116101715780637f1c825a1461059b57806382b86600146105a35780638456cb59146105df57610226565b806366c0bd2414610557578063715018a61461059357610226565b8063593d132c146104095780635c975abb1461042c5780635cffe9de146104485780635fd65f0f1461050e57610226565b806334e19907116101f957806334e199071461033c5780633e3a1560146103595780633f4ba83a1461038b5780634d49e87d1461039357610226565b80630419b45a1461022b5780630ba819591461023557806331cd52b01461024f578063342a87a114610316575b600080fd5b610233610a7b565b005b61023d610b71565b60408051918252519081900360200190f35b6102c66004803603606081101561026557600080fd5b8135919081019060408101602082013564010000000081111561028757600080fd5b82018360208201111561029957600080fd5b803590602001918460208302840111640100000000831117156102bb57600080fd5b919350915035610bf6565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103025781810151838201526020016102ea565b505050509050019250505060405180910390f35b61023d6004803603604081101561032c57600080fd5b508035906020013560ff16610e33565b6102336004803603602081101561035257600080fd5b5035610ee7565b61023d6004803603608081101561036f57600080fd5b5080359060ff6020820135169060408101359060600135610fe8565b6102336111b4565b61023d600480360360608110156103a957600080fd5b8101906020810181356401000000008111156103c457600080fd5b8201836020820111156103d657600080fd5b803590602001918460208302840111640100000000831117156103f857600080fd5b919350915080359060200135611232565b6102336004803603604081101561041f57600080fd5b50803590602001356113d2565b6104346114db565b604080519115158252519081900360200190f35b6102336004803603608081101561045e57600080fd5b6001600160a01b0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561049957600080fd5b8201836020820111156104ab57600080fd5b803590602001918460018302840111640100000000831117156104cd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506114e4945050505050565b610516611976565b604080519788526020880196909652868601949094526060860192909252608085015260a08401526001600160a01b031660c0830152519081900360e00190f35b61057d6004803603602081101561056d57600080fd5b50356001600160a01b0316611997565b6040805160ff9092168252519081900360200190f35b610233611a24565b61023d611afa565b6105c3600480360360208110156105b957600080fd5b503560ff16611b00565b604080516001600160a01b039092168252519081900360200190f35b610233611b88565b61023d600480360360608110156105fd57600080fd5b81019060208101813564010000000081111561061857600080fd5b82018360208201111561062a57600080fd5b8035906020019184602083028401116401000000008311171561064c57600080fd5b919350915080359060200135611c04565b6102336004803603602081101561067357600080fd5b5035611da4565b6105c3611e8a565b61023d600480360360a081101561069857600080fd5b5060ff813581169160208101359091169060408101359060608101359060800135611e99565b61023d600480360360208110156106d457600080fd5b503560ff1661206e565b61023d600480360360608110156106f457600080fd5b5060ff8135811691602081013590911690604001356120ed565b610233600480360361010081101561072557600080fd5b81019060208101813564010000000081111561074057600080fd5b82018360208201111561075257600080fd5b8035906020019184602083028401116401000000008311171561077457600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156107c457600080fd5b8201836020820111156107d657600080fd5b803590602001918460208302840111640100000000831117156107f857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561084857600080fd5b82018360208201111561085a57600080fd5b8035906020019184600183028401116401000000008311171561087c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092959493602081019350359150506401000000008111156108cf57600080fd5b8201836020820111156108e157600080fd5b8035906020019184600183028401116401000000008311171561090357600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602081013590604081013590606001356001600160a01b03166121aa565b61023361226e565b61023d61234d565b61023d6123a1565b61023d6004803603604081101561098c57600080fd5b8101906020810181356401000000008111156109a757600080fd5b8201836020820111156109b957600080fd5b803590602001918460208302840111640100000000831117156109db57600080fd5b91935091503515156123f5565b61023d61248a565b61023d60048036036020811015610a0657600080fd5b5035612490565b61023360048036036040811015610a2357600080fd5b508035906020013561251e565b6102c660048036036020811015610a4657600080fd5b503561260e565b61023360048036036020811015610a6357600080fd5b50356001600160a01b0316612758565b61023d612885565b610a8361288b565b6001600160a01b0316610a94611e8a565b6001600160a01b031614610aef576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b73__$fb68600ab723a8a0a0b96b45bbac9939e0$__6324c5c75160c9610b13611e8a565b6040518363ffffffff1660e01b815260040180838152602001826001600160a01b031681526020019250505060006040518083038186803b158015610b5757600080fd5b505af4158015610b6b573d6000803e3d6000fd5b50505050565b600060c973__$34a89a4c622fc2e024417b7c254097dd7b$__63c9b64dcb90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610bc557600080fd5b505af4158015610bd9573d6000803e3d6000fd5b505050506040513d6020811015610bef57600080fd5b5051905090565b606060026097541415610c50576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026097558142811015610cab576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b60c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6373fd6b3e90918888886040518563ffffffff1660e01b815260040180858152602001848152602001806020018281038252848482818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060006040518083038186803b158015610d3c57600080fd5b505af4158015610d50573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01682016040526020811015610d9757600080fd5b8101908080516040519392919084640100000000821115610db757600080fd5b908301906020820185811115610dcc57600080fd5b8251866020820283011164010000000082111715610de957600080fd5b82525081516020918201928201910280838360005b83811015610e16578181015183820152602001610dfe565b505050509050016040525050509150506001609755949350505050565b604080517fe069274200000000000000000000000000000000000000000000000000000000815260c960048201526024810184905260ff83166044820152905160009173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163e069274291606480820192602092909190829003018186803b158015610eb257600080fd5b505af4158015610ec6573d6000803e3d6000fd5b505050506040513d6020811015610edc57600080fd5b505190505b92915050565b610eef61288b565b6001600160a01b0316610f00611e8a565b6001600160a01b031614610f5b576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517f467e186c00000000000000000000000000000000000000000000000000000000815260c9600482015260248101839052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163467e186c916044808301926000929190829003018186803b158015610fcd57600080fd5b505af4158015610fe1573d6000803e3d6000fd5b5050505050565b600060026097541415611042576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b600260975561104f6114db565b156110a1576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b81804211156110f7576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b604080517fe7a4db8100000000000000000000000000000000000000000000000000000000815260c960048201526024810188905260ff8716604482015260648101869052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163e7a4db81916084808301926020929190829003018186803b15801561117957600080fd5b505af415801561118d573d6000803e3d6000fd5b505050506040513d60208110156111a357600080fd5b505160016097559695505050505050565b6111bc61288b565b6001600160a01b03166111cd611e8a565b6001600160a01b031614611228576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61123061288f565b565b60006002609754141561128c576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b60026097556112996114db565b156112eb576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b8180421115611341576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b60c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6340370edf90918888886040518563ffffffff1660e01b815260040180858152602001806020018381526020018281038252858582818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060206040518083038186803b15801561117957600080fd5b6113da61288b565b6001600160a01b03166113eb611e8a565b6001600160a01b031614611446576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517f58fdd79b00000000000000000000000000000000000000000000000000000000815260c960048201526024810184905260448101839052905173__$34a89a4c622fc2e024417b7c254097dd7b$__916358fdd79b916064808301926000929190829003018186803b1580156114bf57600080fd5b505af41580156114d3573d6000803e3d6000fd5b505050505050565b60655460ff1690565b6002609754141561153c576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002609755600061154c84611997565b90506000846001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561159d57600080fd5b505afa1580156115b1573d6000803e3d6000fd5b505050506040513d60208110156115c757600080fd5b505160d280549192506000916115ff919060ff86169081106115e557fe5b90600052602060002001548361293890919063ffffffff16565b90506000851180156116115750848210155b611662576040805162461bcd60e51b815260206004820152600e60248201527f696e76616c696420616d6f756e74000000000000000000000000000000000000604482015290519081900360640190fd5b600061168561271061167f60d4548961299590919063ffffffff16565b906129ee565b905060006116a461271061167f60d5548561299590919063ffffffff16565b9050600082116116fb576040805162461bcd60e51b815260206004820152601f60248201527f616d6f756e7420697320736d616c6c20666f72206120666c6173684c6f616e00604482015290519081900360640190fd5b61170f6001600160a01b0389168a89612a55565b886001600160a01b031663405b019d308a8a868b6040518663ffffffff1660e01b815260040180866001600160a01b03168152602001856001600160a01b0316815260200184815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611799578181015183820152602001611781565b50505050905090810190601f1680156117c65780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b1580156117e957600080fd5b505af11580156117fd573d6000803e3d6000fd5b505050506000886001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561185057600080fd5b505afa158015611864573d6000803e3d6000fd5b505050506040513d602081101561187a57600080fd5b505190506118888584612ada565b8110156118dc576040805162461bcd60e51b815260206004820152601860248201527f666c6173684c6f616e20666565206973206e6f74206d65740000000000000000604482015290519081900360640190fd5b6118f0826118ea8387612938565b90612938565b60d2805460ff891690811061190157fe5b600091825260209182902001919091556040805160ff891681529182018a905281810185905260608201849052516001600160a01b038c16917f7c186b2827b23e9024e7b29869cba58a97a4bac6567802a8ea6a8afa7b8c22f0919081900360800190a2505060016097555050505050505050565b60c95460ca5460cb5460cc5460cd5460ce5460cf546001600160a01b031687565b6001600160a01b038116600081815260d36020526040812054909160ff909116906119c182611b00565b6001600160a01b031614611a1c576040805162461bcd60e51b815260206004820152601460248201527f546f6b656e20646f6573206e6f74206578697374000000000000000000000000604482015290519081900360640190fd5b90505b919050565b611a2c61288b565b6001600160a01b0316611a3d611e8a565b6001600160a01b031614611a98576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6033546040516000916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60d45481565b60d05460009060ff831610611b5c576040805162461bcd60e51b815260206004820152600c60248201527f4f7574206f662072616e67650000000000000000000000000000000000000000604482015290519081900360640190fd5b60d0805460ff8416908110611b6d57fe5b6000918252602090912001546001600160a01b031692915050565b611b9061288b565b6001600160a01b0316611ba1611e8a565b6001600160a01b031614611bfc576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b611230612b34565b600060026097541415611c5e576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002609755611c6b6114db565b15611cbd576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b8180421115611d13576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b60c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6341b91c2690918888886040518563ffffffff1660e01b815260040180858152602001806020018381526020018281038252858582818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060206040518083038186803b15801561117957600080fd5b611dac61288b565b6001600160a01b0316611dbd611e8a565b6001600160a01b031614611e18576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517f7046727600000000000000000000000000000000000000000000000000000000815260c9600482015260248101839052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__916370467276916044808301926000929190829003018186803b158015610fcd57600080fd5b6033546001600160a01b031690565b600060026097541415611ef3576040805162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604482015290519081900360640190fd5b6002609755611f006114db565b15611f52576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b8180421115611fa8576040805162461bcd60e51b815260206004820152601060248201527f446561646c696e65206e6f74206d657400000000000000000000000000000000604482015290519081900360640190fd5b604080517fa5397b2200000000000000000000000000000000000000000000000000000000815260c9600482015260ff808a166024830152881660448201526064810187905260848101869052905173__$fb68600ab723a8a0a0b96b45bbac9939e0$__9163a5397b229160a4808301926020929190829003018186803b15801561203257600080fd5b505af4158015612046573d6000803e3d6000fd5b505050506040513d602081101561205c57600080fd5b50516001609755979650505050505050565b60d05460009060ff8316106120ca576040805162461bcd60e51b815260206004820152601260248201527f496e646578206f7574206f662072616e67650000000000000000000000000000604482015290519081900360640190fd5b60d2805460ff84169081106120db57fe5b90600052602060002001549050919050565b604080517f4b23603c00000000000000000000000000000000000000000000000000000000815260c9600482015260ff80861660248301528416604482015260648101839052905160009173__$fb68600ab723a8a0a0b96b45bbac9939e0$__91634b23603c91608480820192602092909190829003018186803b15801561217457600080fd5b505af4158015612188573d6000803e3d6000fd5b505050506040513d602081101561219e57600080fd5b505190505b9392505050565b600054610100900460ff16806121c357506121c3612bc4565b806121d1575060005460ff16155b61220c5760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff16158015612237576000805460ff1961ff0019909116610100171660011790555b6122478989898989898989612bd5565b600860d455600060d5558015612263576000805461ff00191690555b505050505050505050565b61227661288b565b6001600160a01b0316612287611e8a565b6001600160a01b0316146122e2576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b604080517ff14e211e00000000000000000000000000000000000000000000000000000000815260c96004820152905173__$34a89a4c622fc2e024417b7c254097dd7b$__9163f14e211e916024808301926000929190829003018186803b158015610b5757600080fd5b600060c973__$34a89a4c622fc2e024417b7c254097dd7b$__63b0a14cfc90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610bc557600080fd5b600060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6371906c2c90916040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610bc557600080fd5b600060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__63834b491090918686866040518563ffffffff1660e01b8152600401808581526020018060200183151581526020018281038252858582818152602001925060200280828437600081840152601f19601f8201169050808301925050509550505050505060206040518083038186803b15801561217457600080fd5b60d55481565b600060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__637d0481609091846040518363ffffffff1660e01b8152600401808381526020018281526020019250505060206040518083038186803b1580156124ec57600080fd5b505af4158015612500573d6000803e3d6000fd5b505050506040513d602081101561251657600080fd5b505192915050565b61252661288b565b6001600160a01b0316612537611e8a565b6001600160a01b031614612592576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6000821180156125a457506127108211155b80156125b257506127108111155b612603576040805162461bcd60e51b815260206004820152601b60248201527f6665657320617265206e6f7420696e2076616c69642072616e67650000000000604482015290519081900360640190fd5b60d49190915560d555565b606060c973__$fb68600ab723a8a0a0b96b45bbac9939e0$__6370703e4a9091846040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b15801561266a57600080fd5b505af415801561267e573d6000803e3d6000fd5b505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016820160405260208110156126c557600080fd5b81019080805160405193929190846401000000008211156126e557600080fd5b9083019060208201858111156126fa57600080fd5b825186602082028301116401000000008211171561271757600080fd5b82525081516020918201928201910280838360005b8381101561274457818101518382015260200161272c565b505050509050016040525050509050919050565b61276061288b565b6001600160a01b0316612771611e8a565b6001600160a01b0316146127cc576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6001600160a01b0381166128115760405162461bcd60e51b8152600401808060200182810382526026815260200180613cd96026913960400191505060405180910390fd5b6033546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b61271081565b3390565b6128976114db565b6128e8576040805162461bcd60e51b815260206004820152601460248201527f5061757361626c653a206e6f7420706175736564000000000000000000000000604482015290519081900360640190fd5b6065805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61291b61288b565b604080516001600160a01b039092168252519081900360200190a1565b60008282111561298f576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b6000826129a457506000610ee1565b828202828482816129b157fe5b04146121a35760405162461bcd60e51b8152600401808060200182810382526021815260200180613d536021913960400191505060405180910390fd5b6000808211612a44576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b818381612a4d57fe5b049392505050565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052612ad59084906133d0565b505050565b6000828201838110156121a3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b612b3c6114db565b15612b8e576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6065805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861291b61288b565b6000612bcf30613481565b15905090565b600054610100900460ff1680612bee5750612bee612bc4565b80612bfc575060005460ff16155b612c375760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff16158015612c62576000805460ff1961ff0019909116610100171660011790555b612c6a613487565b612c72613541565b6001895111612cc8576040805162461bcd60e51b815260206004820152601960248201527f5f706f6f6c6564546f6b656e732e6c656e677468203c3d203100000000000000604482015290519081900360640190fd5b602089511115612d1f576040805162461bcd60e51b815260206004820152601960248201527f5f706f6f6c6564546f6b656e732e6c656e677468203e20333200000000000000604482015290519081900360640190fd5b8751895114612d75576040805162461bcd60e51b815260206004820152601f60248201527f5f706f6f6c6564546f6b656e7320646563696d616c73206d69736d6174636800604482015290519081900360640190fd5b6060885167ffffffffffffffff81118015612d8f57600080fd5b50604051908082528060200260200182016040528015612db9578160200160208202803683370190505b50905060005b8a518160ff1610156130375760ff811615612ea85760d360008c8360ff1681518110612de757fe5b6020908102919091018101516001600160a01b031682528101919091526040016000205460ff16158015612e5757508a8160ff1681518110612e2557fe5b60200260200101516001600160a01b03168b600081518110612e4357fe5b60200260200101516001600160a01b031614155b612ea8576040805162461bcd60e51b815260206004820152601060248201527f4475706c696361746520746f6b656e7300000000000000000000000000000000604482015290519081900360640190fd5b60006001600160a01b03168b8260ff1681518110612ec257fe5b60200260200101516001600160a01b03161415612f26576040805162461bcd60e51b815260206004820152601d60248201527f546865203020616464726573732069736e277420616e204552432d3230000000604482015290519081900360640190fd5b601260ff168a8260ff1681518110612f3a57fe5b602002602001015160ff161115612f98576040805162461bcd60e51b815260206004820152601a60248201527f546f6b656e20646563696d616c732065786365656473206d6178000000000000604482015290519081900360640190fd5b612fc88a8260ff1681518110612faa57fe5b602002602001015160ff16601260ff1661293890919063ffffffff16565b600a0a828260ff1681518110612fda57fe5b6020026020010181815250508060d360008d8460ff1681518110612ffa57fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff191660ff92909216919091179055600101612dbf565b50620f4240861061308f576040805162461bcd60e51b815260206004820152601260248201527f5f612065786365656473206d6178696d756d0000000000000000000000000000604482015290519081900360640190fd5b6305f5e10085106130e7576040805162461bcd60e51b815260206004820152601460248201527f5f6665652065786365656473206d6178696d756d000000000000000000000000604482015290519081900360640190fd5b6402540be4008410613140576040805162461bcd60e51b815260206004820152601960248201527f5f61646d696e4665652065786365656473206d6178696d756d00000000000000604482015290519081900360640190fd5b600061314b846135d6565b9050806001600160a01b0316634cd88b768a8a6040518363ffffffff1660e01b8152600401808060200180602001838103835285818151815260200191508051906020019080838360005b838110156131ae578181015183820152602001613196565b50505050905090810190601f1680156131db5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b8381101561320e5781810151838201526020016131f6565b50505050905090810190601f16801561323b5780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561325c57600080fd5b505af1158015613270573d6000803e3d6000fd5b505050506040513d602081101561328657600080fd5b50516132d9576040805162461bcd60e51b815260206004820152601c60248201527f636f756c64206e6f7420696e6974206c70546f6b656e20636c6f6e6500000000604482015290519081900360640190fd5b60cf80547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383161790558a5161331f9060d09060208e0190613bc8565b5081516133339060d1906020850190613c45565b508a5167ffffffffffffffff8111801561334c57600080fd5b50604051908082528060200260200182016040528015613376578160200160208202803683370190505b50805161338b9160d291602090910190613c45565b50613397876064612995565b60c9556133a5876064612995565b60ca55505060cd84905560ce8390558015612263576000805461ff0019169055505050505050505050565b6060613425826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166136919092919063ffffffff16565b805190915015612ad55780806020019051602081101561344457600080fd5b5051612ad55760405162461bcd60e51b815260040180806020018281038252602a815260200180613d74602a913960400191505060405180910390fd5b3b151590565b600054610100900460ff16806134a057506134a0612bc4565b806134ae575060005460ff16155b6134e95760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff16158015613514576000805460ff1961ff0019909116610100171660011790555b61351c6136a8565b613524613748565b61352c613859565b801561353e576000805461ff00191690555b50565b600054610100900460ff168061355a575061355a612bc4565b80613568575060005460ff16155b6135a35760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff161580156135ce576000805460ff1961ff0019909116610100171660011790555b61352c613904565b60006040517f3d602d80600a3d3981f3363d3d373d3d3d363d7300000000000000000000000081528260601b60148201527f5af43d82803e903d91602b57fd5bf3000000000000000000000000000000000060288201526037816000f09150506001600160a01b038116611a1f576040805162461bcd60e51b815260206004820152601660248201527f455243313136373a20637265617465206661696c656400000000000000000000604482015290519081900360640190fd5b60606136a084846000856139aa565b949350505050565b600054610100900460ff16806136c157506136c1612bc4565b806136cf575060005460ff16155b61370a5760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff1615801561352c576000805460ff1961ff001990911661010017166001179055801561353e576000805461ff001916905550565b600054610100900460ff16806137615750613761612bc4565b8061376f575060005460ff16155b6137aa5760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff161580156137d5576000805460ff1961ff0019909116610100171660011790555b60006137df61288b565b603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038316908117909155604051919250906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350801561353e576000805461ff001916905550565b600054610100900460ff16806138725750613872612bc4565b80613880575060005460ff16155b6138bb5760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff161580156138e6576000805460ff1961ff0019909116610100171660011790555b6065805460ff19169055801561353e576000805461ff001916905550565b600054610100900460ff168061391d575061391d612bc4565b8061392b575060005460ff16155b6139665760405162461bcd60e51b815260040180806020018281038252602e815260200180613d25602e913960400191505060405180910390fd5b600054610100900460ff16158015613991576000805460ff1961ff0019909116610100171660011790555b6001609755801561353e576000805461ff001916905550565b6060824710156139eb5760405162461bcd60e51b8152600401808060200182810382526026815260200180613cff6026913960400191505060405180910390fd5b6139f485613481565b613a45576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b60208310613aa257805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101613a65565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114613b04576040519150601f19603f3d011682016040523d82523d6000602084013e613b09565b606091505b5091509150613b19828286613b24565b979650505050505050565b60608315613b335750816121a3565b825115613b435782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015613b8d578181015183820152602001613b75565b50505050905090810190601f168015613bba5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b828054828255906000526020600020908101928215613c35579160200282015b82811115613c3557825182547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b03909116178255602090920191600190910190613be8565b50613c41929150613c8c565b5090565b828054828255906000526020600020908101928215613c80579160200282015b82811115613c80578251825591602001919060010190613c65565b50613c41929150613cc3565b5b80821115613c415780547fffffffffffffffffffffffff0000000000000000000000000000000000000000168155600101613c8d565b5b80821115613c415760008155600101613cc456fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c496e697469616c697a61626c653a20636f6e747261637420697320616c726561647920696e697469616c697a6564536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220fd19277102b6b4cf17e276c36fef9dfb9c9b662bd1adb50fd544c520d37255d864736f6c634300060c0033",
}

// SwapFlashLoanABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapFlashLoanMetaData.ABI instead.
var SwapFlashLoanABI = SwapFlashLoanMetaData.ABI

// Deprecated: Use SwapFlashLoanMetaData.Sigs instead.
// SwapFlashLoanFuncSigs maps the 4-byte function signature to its string representation.
var SwapFlashLoanFuncSigs = SwapFlashLoanMetaData.Sigs

// SwapFlashLoanBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapFlashLoanMetaData.Bin instead.
var SwapFlashLoanBin = SwapFlashLoanMetaData.Bin

// DeploySwapFlashLoan deploys a new Ethereum contract, binding an instance of SwapFlashLoan to it.
func DeploySwapFlashLoan(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwapFlashLoan, error) {
	parsed, err := SwapFlashLoanMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	amplificationUtilsAddr, _, _, _ := DeployAmplificationUtils(auth, backend)
	SwapFlashLoanBin = strings.ReplaceAll(SwapFlashLoanBin, "__$34a89a4c622fc2e024417b7c254097dd7b$__", amplificationUtilsAddr.String()[2:])

	swapUtilsAddr, _, _, _ := DeploySwapUtils(auth, backend)
	SwapFlashLoanBin = strings.ReplaceAll(SwapFlashLoanBin, "__$fb68600ab723a8a0a0b96b45bbac9939e0$__", swapUtilsAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapFlashLoanBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapFlashLoan{SwapFlashLoanCaller: SwapFlashLoanCaller{contract: contract}, SwapFlashLoanTransactor: SwapFlashLoanTransactor{contract: contract}, SwapFlashLoanFilterer: SwapFlashLoanFilterer{contract: contract}}, nil
}

// SwapFlashLoan is an auto generated Go binding around an Ethereum contract.
type SwapFlashLoan struct {
	SwapFlashLoanCaller     // Read-only binding to the contract
	SwapFlashLoanTransactor // Write-only binding to the contract
	SwapFlashLoanFilterer   // Log filterer for contract events
}

// SwapFlashLoanCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapFlashLoanCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFlashLoanTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapFlashLoanTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFlashLoanFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapFlashLoanFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapFlashLoanSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapFlashLoanSession struct {
	Contract     *SwapFlashLoan    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapFlashLoanCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapFlashLoanCallerSession struct {
	Contract *SwapFlashLoanCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SwapFlashLoanTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapFlashLoanTransactorSession struct {
	Contract     *SwapFlashLoanTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SwapFlashLoanRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapFlashLoanRaw struct {
	Contract *SwapFlashLoan // Generic contract binding to access the raw methods on
}

// SwapFlashLoanCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapFlashLoanCallerRaw struct {
	Contract *SwapFlashLoanCaller // Generic read-only contract binding to access the raw methods on
}

// SwapFlashLoanTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapFlashLoanTransactorRaw struct {
	Contract *SwapFlashLoanTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapFlashLoan creates a new instance of SwapFlashLoan, bound to a specific deployed contract.
func NewSwapFlashLoan(address common.Address, backend bind.ContractBackend) (*SwapFlashLoan, error) {
	contract, err := bindSwapFlashLoan(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoan{SwapFlashLoanCaller: SwapFlashLoanCaller{contract: contract}, SwapFlashLoanTransactor: SwapFlashLoanTransactor{contract: contract}, SwapFlashLoanFilterer: SwapFlashLoanFilterer{contract: contract}}, nil
}

// NewSwapFlashLoanCaller creates a new read-only instance of SwapFlashLoan, bound to a specific deployed contract.
func NewSwapFlashLoanCaller(address common.Address, caller bind.ContractCaller) (*SwapFlashLoanCaller, error) {
	contract, err := bindSwapFlashLoan(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanCaller{contract: contract}, nil
}

// NewSwapFlashLoanTransactor creates a new write-only instance of SwapFlashLoan, bound to a specific deployed contract.
func NewSwapFlashLoanTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapFlashLoanTransactor, error) {
	contract, err := bindSwapFlashLoan(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanTransactor{contract: contract}, nil
}

// NewSwapFlashLoanFilterer creates a new log filterer instance of SwapFlashLoan, bound to a specific deployed contract.
func NewSwapFlashLoanFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapFlashLoanFilterer, error) {
	contract, err := bindSwapFlashLoan(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanFilterer{contract: contract}, nil
}

// bindSwapFlashLoan binds a generic wrapper to an already deployed contract.
func bindSwapFlashLoan(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapFlashLoanABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapFlashLoan *SwapFlashLoanRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapFlashLoan.Contract.SwapFlashLoanCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapFlashLoan *SwapFlashLoanRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SwapFlashLoanTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapFlashLoan *SwapFlashLoanRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SwapFlashLoanTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapFlashLoan *SwapFlashLoanCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapFlashLoan.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapFlashLoan *SwapFlashLoanTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapFlashLoan *SwapFlashLoanTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.contract.Transact(opts, method, params...)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) MAXBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "MAX_BPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) MAXBPS() (*big.Int, error) {
	return _SwapFlashLoan.Contract.MAXBPS(&_SwapFlashLoan.CallOpts)
}

// MAXBPS is a free data retrieval call binding the contract method 0xfd967f47.
//
// Solidity: function MAX_BPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) MAXBPS() (*big.Int, error) {
	return _SwapFlashLoan.Contract.MAXBPS(&_SwapFlashLoan.CallOpts)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_SwapFlashLoan *SwapFlashLoanCaller) CalculateRemoveLiquidity(opts *bind.CallOpts, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "calculateRemoveLiquidity", amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_SwapFlashLoan *SwapFlashLoanSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateRemoveLiquidity(&_SwapFlashLoan.CallOpts, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0xf2fad2b6.
//
// Solidity: function calculateRemoveLiquidity(uint256 amount) view returns(uint256[])
func (_SwapFlashLoan *SwapFlashLoanCallerSession) CalculateRemoveLiquidity(amount *big.Int) ([]*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateRemoveLiquidity(&_SwapFlashLoan.CallOpts, amount)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_SwapFlashLoan *SwapFlashLoanCaller) CalculateRemoveLiquidityOneToken(opts *bind.CallOpts, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "calculateRemoveLiquidityOneToken", tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_SwapFlashLoan *SwapFlashLoanSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateRemoveLiquidityOneToken(&_SwapFlashLoan.CallOpts, tokenAmount, tokenIndex)
}

// CalculateRemoveLiquidityOneToken is a free data retrieval call binding the contract method 0x342a87a1.
//
// Solidity: function calculateRemoveLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 availableTokenAmount)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) CalculateRemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateRemoveLiquidityOneToken(&_SwapFlashLoan.CallOpts, tokenAmount, tokenIndex)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) CalculateSwap(opts *bind.CallOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "calculateSwap", tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateSwap(&_SwapFlashLoan.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0xa95b089f.
//
// Solidity: function calculateSwap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) CalculateSwap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateSwap(&_SwapFlashLoan.CallOpts, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) CalculateTokenAmount(opts *bind.CallOpts, amounts []*big.Int, deposit bool) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "calculateTokenAmount", amounts, deposit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateTokenAmount(&_SwapFlashLoan.CallOpts, amounts, deposit)
}

// CalculateTokenAmount is a free data retrieval call binding the contract method 0xe6ab2806.
//
// Solidity: function calculateTokenAmount(uint256[] amounts, bool deposit) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) CalculateTokenAmount(amounts []*big.Int, deposit bool) (*big.Int, error) {
	return _SwapFlashLoan.Contract.CalculateTokenAmount(&_SwapFlashLoan.CallOpts, amounts, deposit)
}

// FlashLoanFeeBPS is a free data retrieval call binding the contract method 0x7f1c825a.
//
// Solidity: function flashLoanFeeBPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) FlashLoanFeeBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "flashLoanFeeBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FlashLoanFeeBPS is a free data retrieval call binding the contract method 0x7f1c825a.
//
// Solidity: function flashLoanFeeBPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) FlashLoanFeeBPS() (*big.Int, error) {
	return _SwapFlashLoan.Contract.FlashLoanFeeBPS(&_SwapFlashLoan.CallOpts)
}

// FlashLoanFeeBPS is a free data retrieval call binding the contract method 0x7f1c825a.
//
// Solidity: function flashLoanFeeBPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) FlashLoanFeeBPS() (*big.Int, error) {
	return _SwapFlashLoan.Contract.FlashLoanFeeBPS(&_SwapFlashLoan.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) GetA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "getA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) GetA() (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetA(&_SwapFlashLoan.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) GetA() (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetA(&_SwapFlashLoan.CallOpts)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) GetAPrecise(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "getAPrecise")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) GetAPrecise() (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetAPrecise(&_SwapFlashLoan.CallOpts)
}

// GetAPrecise is a free data retrieval call binding the contract method 0x0ba81959.
//
// Solidity: function getAPrecise() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) GetAPrecise() (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetAPrecise(&_SwapFlashLoan.CallOpts)
}

// GetAdminBalance is a free data retrieval call binding the contract method 0xef0a712f.
//
// Solidity: function getAdminBalance(uint256 index) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) GetAdminBalance(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "getAdminBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAdminBalance is a free data retrieval call binding the contract method 0xef0a712f.
//
// Solidity: function getAdminBalance(uint256 index) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) GetAdminBalance(index *big.Int) (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetAdminBalance(&_SwapFlashLoan.CallOpts, index)
}

// GetAdminBalance is a free data retrieval call binding the contract method 0xef0a712f.
//
// Solidity: function getAdminBalance(uint256 index) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) GetAdminBalance(index *big.Int) (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetAdminBalance(&_SwapFlashLoan.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_SwapFlashLoan *SwapFlashLoanCaller) GetToken(opts *bind.CallOpts, index uint8) (common.Address, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "getToken", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_SwapFlashLoan *SwapFlashLoanSession) GetToken(index uint8) (common.Address, error) {
	return _SwapFlashLoan.Contract.GetToken(&_SwapFlashLoan.CallOpts, index)
}

// GetToken is a free data retrieval call binding the contract method 0x82b86600.
//
// Solidity: function getToken(uint8 index) view returns(address)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) GetToken(index uint8) (common.Address, error) {
	return _SwapFlashLoan.Contract.GetToken(&_SwapFlashLoan.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) GetTokenBalance(opts *bind.CallOpts, index uint8) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "getTokenBalance", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetTokenBalance(&_SwapFlashLoan.CallOpts, index)
}

// GetTokenBalance is a free data retrieval call binding the contract method 0x91ceb3eb.
//
// Solidity: function getTokenBalance(uint8 index) view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) GetTokenBalance(index uint8) (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetTokenBalance(&_SwapFlashLoan.CallOpts, index)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_SwapFlashLoan *SwapFlashLoanCaller) GetTokenIndex(opts *bind.CallOpts, tokenAddress common.Address) (uint8, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "getTokenIndex", tokenAddress)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_SwapFlashLoan *SwapFlashLoanSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _SwapFlashLoan.Contract.GetTokenIndex(&_SwapFlashLoan.CallOpts, tokenAddress)
}

// GetTokenIndex is a free data retrieval call binding the contract method 0x66c0bd24.
//
// Solidity: function getTokenIndex(address tokenAddress) view returns(uint8)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) GetTokenIndex(tokenAddress common.Address) (uint8, error) {
	return _SwapFlashLoan.Contract.GetTokenIndex(&_SwapFlashLoan.CallOpts, tokenAddress)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) GetVirtualPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "getVirtualPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) GetVirtualPrice() (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetVirtualPrice(&_SwapFlashLoan.CallOpts)
}

// GetVirtualPrice is a free data retrieval call binding the contract method 0xe25aa5fa.
//
// Solidity: function getVirtualPrice() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) GetVirtualPrice() (*big.Int, error) {
	return _SwapFlashLoan.Contract.GetVirtualPrice(&_SwapFlashLoan.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapFlashLoan *SwapFlashLoanCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapFlashLoan *SwapFlashLoanSession) Owner() (common.Address, error) {
	return _SwapFlashLoan.Contract.Owner(&_SwapFlashLoan.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) Owner() (common.Address, error) {
	return _SwapFlashLoan.Contract.Owner(&_SwapFlashLoan.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SwapFlashLoan *SwapFlashLoanCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SwapFlashLoan *SwapFlashLoanSession) Paused() (bool, error) {
	return _SwapFlashLoan.Contract.Paused(&_SwapFlashLoan.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) Paused() (bool, error) {
	return _SwapFlashLoan.Contract.Paused(&_SwapFlashLoan.CallOpts)
}

// ProtocolFeeShareBPS is a free data retrieval call binding the contract method 0xe8cc7fb6.
//
// Solidity: function protocolFeeShareBPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCaller) ProtocolFeeShareBPS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "protocolFeeShareBPS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeeShareBPS is a free data retrieval call binding the contract method 0xe8cc7fb6.
//
// Solidity: function protocolFeeShareBPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) ProtocolFeeShareBPS() (*big.Int, error) {
	return _SwapFlashLoan.Contract.ProtocolFeeShareBPS(&_SwapFlashLoan.CallOpts)
}

// ProtocolFeeShareBPS is a free data retrieval call binding the contract method 0xe8cc7fb6.
//
// Solidity: function protocolFeeShareBPS() view returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) ProtocolFeeShareBPS() (*big.Int, error) {
	return _SwapFlashLoan.Contract.ProtocolFeeShareBPS(&_SwapFlashLoan.CallOpts)
}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_SwapFlashLoan *SwapFlashLoanCaller) SwapStorage(opts *bind.CallOpts) (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	var out []interface{}
	err := _SwapFlashLoan.contract.Call(opts, &out, "swapStorage")

	outstruct := new(struct {
		InitialA     *big.Int
		FutureA      *big.Int
		InitialATime *big.Int
		FutureATime  *big.Int
		SwapFee      *big.Int
		AdminFee     *big.Int
		LpToken      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.InitialA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FutureA = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitialATime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.FutureATime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SwapFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.AdminFee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LpToken = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_SwapFlashLoan *SwapFlashLoanSession) SwapStorage() (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	return _SwapFlashLoan.Contract.SwapStorage(&_SwapFlashLoan.CallOpts)
}

// SwapStorage is a free data retrieval call binding the contract method 0x5fd65f0f.
//
// Solidity: function swapStorage() view returns(uint256 initialA, uint256 futureA, uint256 initialATime, uint256 futureATime, uint256 swapFee, uint256 adminFee, address lpToken)
func (_SwapFlashLoan *SwapFlashLoanCallerSession) SwapStorage() (struct {
	InitialA     *big.Int
	FutureA      *big.Int
	InitialATime *big.Int
	FutureATime  *big.Int
	SwapFee      *big.Int
	AdminFee     *big.Int
	LpToken      common.Address
}, error) {
	return _SwapFlashLoan.Contract.SwapStorage(&_SwapFlashLoan.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactor) AddLiquidity(opts *bind.TransactOpts, amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "addLiquidity", amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.AddLiquidity(&_SwapFlashLoan.TransactOpts, amounts, minToMint, deadline)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4d49e87d.
//
// Solidity: function addLiquidity(uint256[] amounts, uint256 minToMint, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) AddLiquidity(amounts []*big.Int, minToMint *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.AddLiquidity(&_SwapFlashLoan.TransactOpts, amounts, minToMint, deadline)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes params) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "flashLoan", receiver, token, amount, params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes params) returns()
func (_SwapFlashLoan *SwapFlashLoanSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.FlashLoan(&_SwapFlashLoan.TransactOpts, receiver, token, amount, params)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x5cffe9de.
//
// Solidity: function flashLoan(address receiver, address token, uint256 amount, bytes params) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) FlashLoan(receiver common.Address, token common.Address, amount *big.Int, params []byte) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.FlashLoan(&_SwapFlashLoan.TransactOpts, receiver, token, amount, params)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] _pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 _a, uint256 _fee, uint256 _adminFee, address lpTokenTargetAddress) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) Initialize(opts *bind.TransactOpts, _pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, _a *big.Int, _fee *big.Int, _adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "initialize", _pooledTokens, decimals, lpTokenName, lpTokenSymbol, _a, _fee, _adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] _pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 _a, uint256 _fee, uint256 _adminFee, address lpTokenTargetAddress) returns()
func (_SwapFlashLoan *SwapFlashLoanSession) Initialize(_pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, _a *big.Int, _fee *big.Int, _adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Initialize(&_SwapFlashLoan.TransactOpts, _pooledTokens, decimals, lpTokenName, lpTokenSymbol, _a, _fee, _adminFee, lpTokenTargetAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xb28cb6dc.
//
// Solidity: function initialize(address[] _pooledTokens, uint8[] decimals, string lpTokenName, string lpTokenSymbol, uint256 _a, uint256 _fee, uint256 _adminFee, address lpTokenTargetAddress) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) Initialize(_pooledTokens []common.Address, decimals []uint8, lpTokenName string, lpTokenSymbol string, _a *big.Int, _fee *big.Int, _adminFee *big.Int, lpTokenTargetAddress common.Address) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Initialize(&_SwapFlashLoan.TransactOpts, _pooledTokens, decimals, lpTokenName, lpTokenSymbol, _a, _fee, _adminFee, lpTokenTargetAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SwapFlashLoan *SwapFlashLoanSession) Pause() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Pause(&_SwapFlashLoan.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) Pause() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Pause(&_SwapFlashLoan.TransactOpts)
}

// RampA is a paid mutator transaction binding the contract method 0x593d132c.
//
// Solidity: function rampA(uint256 futureA, uint256 futureTime) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) RampA(opts *bind.TransactOpts, futureA *big.Int, futureTime *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "rampA", futureA, futureTime)
}

// RampA is a paid mutator transaction binding the contract method 0x593d132c.
//
// Solidity: function rampA(uint256 futureA, uint256 futureTime) returns()
func (_SwapFlashLoan *SwapFlashLoanSession) RampA(futureA *big.Int, futureTime *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RampA(&_SwapFlashLoan.TransactOpts, futureA, futureTime)
}

// RampA is a paid mutator transaction binding the contract method 0x593d132c.
//
// Solidity: function rampA(uint256 futureA, uint256 futureTime) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) RampA(futureA *big.Int, futureTime *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RampA(&_SwapFlashLoan.TransactOpts, futureA, futureTime)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_SwapFlashLoan *SwapFlashLoanTransactor) RemoveLiquidity(opts *bind.TransactOpts, amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "removeLiquidity", amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_SwapFlashLoan *SwapFlashLoanSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RemoveLiquidity(&_SwapFlashLoan.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x31cd52b0.
//
// Solidity: function removeLiquidity(uint256 amount, uint256[] minAmounts, uint256 deadline) returns(uint256[])
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) RemoveLiquidity(amount *big.Int, minAmounts []*big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RemoveLiquidity(&_SwapFlashLoan.TransactOpts, amount, minAmounts, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactor) RemoveLiquidityImbalance(opts *bind.TransactOpts, amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "removeLiquidityImbalance", amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RemoveLiquidityImbalance(&_SwapFlashLoan.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityImbalance is a paid mutator transaction binding the contract method 0x84cdd9bc.
//
// Solidity: function removeLiquidityImbalance(uint256[] amounts, uint256 maxBurnAmount, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) RemoveLiquidityImbalance(amounts []*big.Int, maxBurnAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RemoveLiquidityImbalance(&_SwapFlashLoan.TransactOpts, amounts, maxBurnAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactor) RemoveLiquidityOneToken(opts *bind.TransactOpts, tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "removeLiquidityOneToken", tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RemoveLiquidityOneToken(&_SwapFlashLoan.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RemoveLiquidityOneToken is a paid mutator transaction binding the contract method 0x3e3a1560.
//
// Solidity: function removeLiquidityOneToken(uint256 tokenAmount, uint8 tokenIndex, uint256 minAmount, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) RemoveLiquidityOneToken(tokenAmount *big.Int, tokenIndex uint8, minAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RemoveLiquidityOneToken(&_SwapFlashLoan.TransactOpts, tokenAmount, tokenIndex, minAmount, deadline)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapFlashLoan *SwapFlashLoanSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RenounceOwnership(&_SwapFlashLoan.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.RenounceOwnership(&_SwapFlashLoan.TransactOpts)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x8beb60b6.
//
// Solidity: function setAdminFee(uint256 newAdminFee) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) SetAdminFee(opts *bind.TransactOpts, newAdminFee *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "setAdminFee", newAdminFee)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x8beb60b6.
//
// Solidity: function setAdminFee(uint256 newAdminFee) returns()
func (_SwapFlashLoan *SwapFlashLoanSession) SetAdminFee(newAdminFee *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SetAdminFee(&_SwapFlashLoan.TransactOpts, newAdminFee)
}

// SetAdminFee is a paid mutator transaction binding the contract method 0x8beb60b6.
//
// Solidity: function setAdminFee(uint256 newAdminFee) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) SetAdminFee(newAdminFee *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SetAdminFee(&_SwapFlashLoan.TransactOpts, newAdminFee)
}

// SetFlashLoanFees is a paid mutator transaction binding the contract method 0xef815967.
//
// Solidity: function setFlashLoanFees(uint256 newFlashLoanFeeBPS, uint256 newProtocolFeeShareBPS) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) SetFlashLoanFees(opts *bind.TransactOpts, newFlashLoanFeeBPS *big.Int, newProtocolFeeShareBPS *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "setFlashLoanFees", newFlashLoanFeeBPS, newProtocolFeeShareBPS)
}

// SetFlashLoanFees is a paid mutator transaction binding the contract method 0xef815967.
//
// Solidity: function setFlashLoanFees(uint256 newFlashLoanFeeBPS, uint256 newProtocolFeeShareBPS) returns()
func (_SwapFlashLoan *SwapFlashLoanSession) SetFlashLoanFees(newFlashLoanFeeBPS *big.Int, newProtocolFeeShareBPS *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SetFlashLoanFees(&_SwapFlashLoan.TransactOpts, newFlashLoanFeeBPS, newProtocolFeeShareBPS)
}

// SetFlashLoanFees is a paid mutator transaction binding the contract method 0xef815967.
//
// Solidity: function setFlashLoanFees(uint256 newFlashLoanFeeBPS, uint256 newProtocolFeeShareBPS) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) SetFlashLoanFees(newFlashLoanFeeBPS *big.Int, newProtocolFeeShareBPS *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SetFlashLoanFees(&_SwapFlashLoan.TransactOpts, newFlashLoanFeeBPS, newProtocolFeeShareBPS)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 newSwapFee) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) SetSwapFee(opts *bind.TransactOpts, newSwapFee *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "setSwapFee", newSwapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 newSwapFee) returns()
func (_SwapFlashLoan *SwapFlashLoanSession) SetSwapFee(newSwapFee *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SetSwapFee(&_SwapFlashLoan.TransactOpts, newSwapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 newSwapFee) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) SetSwapFee(newSwapFee *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.SetSwapFee(&_SwapFlashLoan.TransactOpts, newSwapFee)
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) StopRampA(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "stopRampA")
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_SwapFlashLoan *SwapFlashLoanSession) StopRampA() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.StopRampA(&_SwapFlashLoan.TransactOpts)
}

// StopRampA is a paid mutator transaction binding the contract method 0xc4db7fa0.
//
// Solidity: function stopRampA() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) StopRampA() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.StopRampA(&_SwapFlashLoan.TransactOpts)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactor) Swap(opts *bind.TransactOpts, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "swap", tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Swap(&_SwapFlashLoan.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// Swap is a paid mutator transaction binding the contract method 0x91695586.
//
// Solidity: function swap(uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx, uint256 minDy, uint256 deadline) returns(uint256)
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) Swap(tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int, minDy *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Swap(&_SwapFlashLoan.TransactOpts, tokenIndexFrom, tokenIndexTo, dx, minDy, deadline)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapFlashLoan *SwapFlashLoanSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.TransferOwnership(&_SwapFlashLoan.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.TransferOwnership(&_SwapFlashLoan.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SwapFlashLoan *SwapFlashLoanSession) Unpause() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Unpause(&_SwapFlashLoan.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) Unpause() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.Unpause(&_SwapFlashLoan.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x0419b45a.
//
// Solidity: function withdrawAdminFees() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactor) WithdrawAdminFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapFlashLoan.contract.Transact(opts, "withdrawAdminFees")
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x0419b45a.
//
// Solidity: function withdrawAdminFees() returns()
func (_SwapFlashLoan *SwapFlashLoanSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.WithdrawAdminFees(&_SwapFlashLoan.TransactOpts)
}

// WithdrawAdminFees is a paid mutator transaction binding the contract method 0x0419b45a.
//
// Solidity: function withdrawAdminFees() returns()
func (_SwapFlashLoan *SwapFlashLoanTransactorSession) WithdrawAdminFees() (*types.Transaction, error) {
	return _SwapFlashLoan.Contract.WithdrawAdminFees(&_SwapFlashLoan.TransactOpts)
}

// SwapFlashLoanAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the SwapFlashLoan contract.
type SwapFlashLoanAddLiquidityIterator struct {
	Event *SwapFlashLoanAddLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanAddLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanAddLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanAddLiquidity represents a AddLiquidity event raised by the SwapFlashLoan contract.
type SwapFlashLoanAddLiquidity struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	Fees          []*big.Int
	Invariant     *big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SwapFlashLoanAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanAddLiquidityIterator{contract: _SwapFlashLoan.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanAddLiquidity)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddLiquidity is a log parse operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseAddLiquidity(log types.Log) (*SwapFlashLoanAddLiquidity, error) {
	event := new(SwapFlashLoanAddLiquidity)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the SwapFlashLoan contract.
type SwapFlashLoanFlashLoanIterator struct {
	Event *SwapFlashLoanFlashLoan // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanFlashLoan)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanFlashLoan)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanFlashLoan represents a FlashLoan event raised by the SwapFlashLoan contract.
type SwapFlashLoanFlashLoan struct {
	Receiver    common.Address
	TokenIndex  uint8
	Amount      *big.Int
	AmountFee   *big.Int
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x7c186b2827b23e9024e7b29869cba58a97a4bac6567802a8ea6a8afa7b8c22f0.
//
// Solidity: event FlashLoan(address indexed receiver, uint8 tokenIndex, uint256 amount, uint256 amountFee, uint256 protocolFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterFlashLoan(opts *bind.FilterOpts, receiver []common.Address) (*SwapFlashLoanFlashLoanIterator, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "FlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanFlashLoanIterator{contract: _SwapFlashLoan.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x7c186b2827b23e9024e7b29869cba58a97a4bac6567802a8ea6a8afa7b8c22f0.
//
// Solidity: event FlashLoan(address indexed receiver, uint8 tokenIndex, uint256 amount, uint256 amountFee, uint256 protocolFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanFlashLoan, receiver []common.Address) (event.Subscription, error) {

	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "FlashLoan", receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanFlashLoan)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "FlashLoan", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFlashLoan is a log parse operation binding the contract event 0x7c186b2827b23e9024e7b29869cba58a97a4bac6567802a8ea6a8afa7b8c22f0.
//
// Solidity: event FlashLoan(address indexed receiver, uint8 tokenIndex, uint256 amount, uint256 amountFee, uint256 protocolFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseFlashLoan(log types.Log) (*SwapFlashLoanFlashLoan, error) {
	event := new(SwapFlashLoanFlashLoan)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanNewAdminFeeIterator is returned from FilterNewAdminFee and is used to iterate over the raw logs and unpacked data for NewAdminFee events raised by the SwapFlashLoan contract.
type SwapFlashLoanNewAdminFeeIterator struct {
	Event *SwapFlashLoanNewAdminFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanNewAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanNewAdminFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanNewAdminFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanNewAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanNewAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanNewAdminFee represents a NewAdminFee event raised by the SwapFlashLoan contract.
type SwapFlashLoanNewAdminFee struct {
	NewAdminFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewAdminFee is a free log retrieval operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterNewAdminFee(opts *bind.FilterOpts) (*SwapFlashLoanNewAdminFeeIterator, error) {

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "NewAdminFee")
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanNewAdminFeeIterator{contract: _SwapFlashLoan.contract, event: "NewAdminFee", logs: logs, sub: sub}, nil
}

// WatchNewAdminFee is a free log subscription operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchNewAdminFee(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanNewAdminFee) (event.Subscription, error) {

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "NewAdminFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanNewAdminFee)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "NewAdminFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdminFee is a log parse operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseNewAdminFee(log types.Log) (*SwapFlashLoanNewAdminFee, error) {
	event := new(SwapFlashLoanNewAdminFee)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "NewAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanNewSwapFeeIterator is returned from FilterNewSwapFee and is used to iterate over the raw logs and unpacked data for NewSwapFee events raised by the SwapFlashLoan contract.
type SwapFlashLoanNewSwapFeeIterator struct {
	Event *SwapFlashLoanNewSwapFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanNewSwapFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanNewSwapFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanNewSwapFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanNewSwapFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanNewSwapFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanNewSwapFee represents a NewSwapFee event raised by the SwapFlashLoan contract.
type SwapFlashLoanNewSwapFee struct {
	NewSwapFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSwapFee is a free log retrieval operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterNewSwapFee(opts *bind.FilterOpts) (*SwapFlashLoanNewSwapFeeIterator, error) {

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "NewSwapFee")
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanNewSwapFeeIterator{contract: _SwapFlashLoan.contract, event: "NewSwapFee", logs: logs, sub: sub}, nil
}

// WatchNewSwapFee is a free log subscription operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchNewSwapFee(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanNewSwapFee) (event.Subscription, error) {

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "NewSwapFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanNewSwapFee)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "NewSwapFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSwapFee is a log parse operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseNewSwapFee(log types.Log) (*SwapFlashLoanNewSwapFee, error) {
	event := new(SwapFlashLoanNewSwapFee)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "NewSwapFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SwapFlashLoan contract.
type SwapFlashLoanOwnershipTransferredIterator struct {
	Event *SwapFlashLoanOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanOwnershipTransferred represents a OwnershipTransferred event raised by the SwapFlashLoan contract.
type SwapFlashLoanOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SwapFlashLoanOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanOwnershipTransferredIterator{contract: _SwapFlashLoan.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanOwnershipTransferred)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseOwnershipTransferred(log types.Log) (*SwapFlashLoanOwnershipTransferred, error) {
	event := new(SwapFlashLoanOwnershipTransferred)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the SwapFlashLoan contract.
type SwapFlashLoanPausedIterator struct {
	Event *SwapFlashLoanPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanPaused represents a Paused event raised by the SwapFlashLoan contract.
type SwapFlashLoanPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterPaused(opts *bind.FilterOpts) (*SwapFlashLoanPausedIterator, error) {

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanPausedIterator{contract: _SwapFlashLoan.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanPaused) (event.Subscription, error) {

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanPaused)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParsePaused(log types.Log) (*SwapFlashLoanPaused, error) {
	event := new(SwapFlashLoanPaused)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanRampAIterator is returned from FilterRampA and is used to iterate over the raw logs and unpacked data for RampA events raised by the SwapFlashLoan contract.
type SwapFlashLoanRampAIterator struct {
	Event *SwapFlashLoanRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanRampA represents a RampA event raised by the SwapFlashLoan contract.
type SwapFlashLoanRampA struct {
	OldA        *big.Int
	NewA        *big.Int
	InitialTime *big.Int
	FutureTime  *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRampA is a free log retrieval operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterRampA(opts *bind.FilterOpts) (*SwapFlashLoanRampAIterator, error) {

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanRampAIterator{contract: _SwapFlashLoan.contract, event: "RampA", logs: logs, sub: sub}, nil
}

// WatchRampA is a free log subscription operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchRampA(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanRampA) (event.Subscription, error) {

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "RampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanRampA)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "RampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRampA is a log parse operation binding the contract event 0xa2b71ec6df949300b59aab36b55e189697b750119dd349fcfa8c0f779e83c254.
//
// Solidity: event RampA(uint256 oldA, uint256 newA, uint256 initialTime, uint256 futureTime)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseRampA(log types.Log) (*SwapFlashLoanRampA, error) {
	event := new(SwapFlashLoanRampA)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "RampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the SwapFlashLoan contract.
type SwapFlashLoanRemoveLiquidityIterator struct {
	Event *SwapFlashLoanRemoveLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanRemoveLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanRemoveLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanRemoveLiquidity represents a RemoveLiquidity event raised by the SwapFlashLoan contract.
type SwapFlashLoanRemoveLiquidity struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SwapFlashLoanRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanRemoveLiquidityIterator{contract: _SwapFlashLoan.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanRemoveLiquidity)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseRemoveLiquidity(log types.Log) (*SwapFlashLoanRemoveLiquidity, error) {
	event := new(SwapFlashLoanRemoveLiquidity)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the SwapFlashLoan contract.
type SwapFlashLoanRemoveLiquidityImbalanceIterator struct {
	Event *SwapFlashLoanRemoveLiquidityImbalance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanRemoveLiquidityImbalance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanRemoveLiquidityImbalance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the SwapFlashLoan contract.
type SwapFlashLoanRemoveLiquidityImbalance struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	Fees          []*big.Int
	Invariant     *big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*SwapFlashLoanRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanRemoveLiquidityImbalanceIterator{contract: _SwapFlashLoan.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanRemoveLiquidityImbalance)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*SwapFlashLoanRemoveLiquidityImbalance, error) {
	event := new(SwapFlashLoanRemoveLiquidityImbalance)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the SwapFlashLoan contract.
type SwapFlashLoanRemoveLiquidityOneIterator struct {
	Event *SwapFlashLoanRemoveLiquidityOne // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanRemoveLiquidityOne)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanRemoveLiquidityOne)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the SwapFlashLoan contract.
type SwapFlashLoanRemoveLiquidityOne struct {
	Provider      common.Address
	LpTokenAmount *big.Int
	LpTokenSupply *big.Int
	BoughtId      *big.Int
	TokensBought  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*SwapFlashLoanRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanRemoveLiquidityOneIterator{contract: _SwapFlashLoan.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanRemoveLiquidityOne)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseRemoveLiquidityOne(log types.Log) (*SwapFlashLoanRemoveLiquidityOne, error) {
	event := new(SwapFlashLoanRemoveLiquidityOne)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanStopRampAIterator is returned from FilterStopRampA and is used to iterate over the raw logs and unpacked data for StopRampA events raised by the SwapFlashLoan contract.
type SwapFlashLoanStopRampAIterator struct {
	Event *SwapFlashLoanStopRampA // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanStopRampAIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanStopRampA)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanStopRampA)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanStopRampAIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanStopRampAIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanStopRampA represents a StopRampA event raised by the SwapFlashLoan contract.
type SwapFlashLoanStopRampA struct {
	CurrentA *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterStopRampA is a free log retrieval operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterStopRampA(opts *bind.FilterOpts) (*SwapFlashLoanStopRampAIterator, error) {

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanStopRampAIterator{contract: _SwapFlashLoan.contract, event: "StopRampA", logs: logs, sub: sub}, nil
}

// WatchStopRampA is a free log subscription operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchStopRampA(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanStopRampA) (event.Subscription, error) {

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "StopRampA")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanStopRampA)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "StopRampA", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStopRampA is a log parse operation binding the contract event 0x46e22fb3709ad289f62ce63d469248536dbc78d82b84a3d7e74ad606dc201938.
//
// Solidity: event StopRampA(uint256 currentA, uint256 time)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseStopRampA(log types.Log) (*SwapFlashLoanStopRampA, error) {
	event := new(SwapFlashLoanStopRampA)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "StopRampA", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanTokenSwapIterator is returned from FilterTokenSwap and is used to iterate over the raw logs and unpacked data for TokenSwap events raised by the SwapFlashLoan contract.
type SwapFlashLoanTokenSwapIterator struct {
	Event *SwapFlashLoanTokenSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanTokenSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanTokenSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanTokenSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanTokenSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanTokenSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanTokenSwap represents a TokenSwap event raised by the SwapFlashLoan contract.
type SwapFlashLoanTokenSwap struct {
	Buyer        common.Address
	TokensSold   *big.Int
	TokensBought *big.Int
	SoldId       *big.Int
	BoughtId     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenSwap is a free log retrieval operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterTokenSwap(opts *bind.FilterOpts, buyer []common.Address) (*SwapFlashLoanTokenSwapIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "TokenSwap", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanTokenSwapIterator{contract: _SwapFlashLoan.contract, event: "TokenSwap", logs: logs, sub: sub}, nil
}

// WatchTokenSwap is a free log subscription operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchTokenSwap(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanTokenSwap, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "TokenSwap", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanTokenSwap)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "TokenSwap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSwap is a log parse operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseTokenSwap(log types.Log) (*SwapFlashLoanTokenSwap, error) {
	event := new(SwapFlashLoanTokenSwap)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "TokenSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapFlashLoanUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the SwapFlashLoan contract.
type SwapFlashLoanUnpausedIterator struct {
	Event *SwapFlashLoanUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapFlashLoanUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapFlashLoanUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapFlashLoanUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapFlashLoanUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapFlashLoanUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapFlashLoanUnpaused represents a Unpaused event raised by the SwapFlashLoan contract.
type SwapFlashLoanUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SwapFlashLoan *SwapFlashLoanFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SwapFlashLoanUnpausedIterator, error) {

	logs, sub, err := _SwapFlashLoan.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SwapFlashLoanUnpausedIterator{contract: _SwapFlashLoan.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SwapFlashLoan *SwapFlashLoanFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SwapFlashLoanUnpaused) (event.Subscription, error) {

	logs, sub, err := _SwapFlashLoan.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapFlashLoanUnpaused)
				if err := _SwapFlashLoan.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_SwapFlashLoan *SwapFlashLoanFilterer) ParseUnpaused(log types.Log) (*SwapFlashLoanUnpaused, error) {
	event := new(SwapFlashLoanUnpaused)
	if err := _SwapFlashLoan.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUtilsMetaData contains all meta data concerning the SwapUtils contract.
var SwapUtilsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invariant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAdminFee\",\"type\":\"uint256\"}],\"name\":\"NewAdminFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSwapFee\",\"type\":\"uint256\"}],\"name\":\"NewSwapFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"tokenAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invariant\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityImbalance\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lpTokenSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"boughtId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"}],\"name\":\"RemoveLiquidityOne\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensSold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokensBought\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"soldId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"boughtId\",\"type\":\"uint128\"}],\"name\":\"TokenSwap\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_ADMIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SWAP_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL_PRECISION_DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f3de0362": "MAX_ADMIN_FEE()",
		"ab3d8544": "MAX_SWAP_FEE()",
		"0296ab50": "POOL_PRECISION_DECIMALS()",
		"40370edf": "addLiquidity(SwapUtils.Swap storage,uint256[],uint256)",
		"70703e4a": "calculateRemoveLiquidity(SwapUtils.Swap storage,uint256)",
		"4b23603c": "calculateSwap(SwapUtils.Swap storage,uint8,uint8,uint256)",
		"834b4910": "calculateTokenAmount(SwapUtils.Swap storage,uint256[],bool)",
		"e0692742": "calculateWithdrawOneToken(SwapUtils.Swap storage,uint256,uint8)",
		"7d048160": "getAdminBalance(SwapUtils.Swap storage,uint256)",
		"71906c2c": "getVirtualPrice(SwapUtils.Swap storage)",
		"73fd6b3e": "removeLiquidity(SwapUtils.Swap storage,uint256,uint256[])",
		"41b91c26": "removeLiquidityImbalance(SwapUtils.Swap storage,uint256[],uint256)",
		"e7a4db81": "removeLiquidityOneToken(SwapUtils.Swap storage,uint256,uint8,uint256)",
		"70467276": "setAdminFee(SwapUtils.Swap storage,uint256)",
		"467e186c": "setSwapFee(SwapUtils.Swap storage,uint256)",
		"a5397b22": "swap(SwapUtils.Swap storage,uint8,uint8,uint256,uint256)",
		"24c5c751": "withdrawAdminFees(SwapUtils.Swap storage,address)",
	},
	Bin: "0x6140b5610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106101255760003560e01c806371906c2c116100bc578063a5397b221161008b578063e069274211610070578063e0692742146105a3578063e7a4db81146105cf578063f3de03621461060e57610125565b8063a5397b2214610552578063ab3d85441461059b57610125565b806371906c2c1461041057806373fd6b3e1461042d5780637d048160146104b6578063834b4910146104d957610125565b8063467e186c116100f8578063467e186c146103075780634b23603c14610337578063704672761461036d57806370703e4a1461039d57610125565b80630296ab501461012a57806324c5c7511461014857806340370edf1461018357806341b91c261461024e575b600080fd5b610132610616565b6040805160ff9092168252519081900360200190f35b81801561015457600080fd5b506101816004803603604081101561016b57600080fd5b50803590602001356001600160a01b031661061b565b005b81801561018f57600080fd5b5061023c600480360360608110156101a657600080fd5b813591908101906040810160208201356401000000008111156101c857600080fd5b8201836020820111156101da57600080fd5b803590602001918460208302840111640100000000831117156101fc57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550509135925061076a915050565b60408051918252519081900360200190f35b81801561025a57600080fd5b5061023c6004803603606081101561027157600080fd5b8135919081019060408101602082013564010000000081111561029357600080fd5b8201836020820111156102a557600080fd5b803590602001918460208302840111640100000000831117156102c757600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506110ea915050565b81801561031357600080fd5b506101816004803603604081101561032a57600080fd5b5080359060200135611886565b61023c6004803603608081101561034d57600080fd5b5080359060ff60208201358116916040810135909116906060013561191d565b81801561037957600080fd5b506101816004803603604081101561039057600080fd5b5080359060200135611988565b6103c0600480360360408110156103b357600080fd5b5080359060200135611a20565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156103fc5781810151838201526020016103e4565b505050509050019250505060405180910390f35b61023c6004803603602081101561042657600080fd5b5035611b03565b81801561043957600080fd5b506103c06004803603606081101561045057600080fd5b81359160208101359181019060608101604082013564010000000081111561047757600080fd5b82018360208201111561048957600080fd5b803590602001918460208302840111640100000000831117156104ab57600080fd5b509092509050611bda565b61023c600480360360408110156104cc57600080fd5b5080359060200135612034565b61023c600480360360608110156104ef57600080fd5b8135919081019060408101602082013564010000000081111561051157600080fd5b82018360208201111561052357600080fd5b8035906020019184602083028401116401000000008311171561054557600080fd5b9193509150351515612123565b81801561055e57600080fd5b5061023c600480360360a081101561057557600080fd5b5080359060ff602082013581169160408101359091169060608101359060800135612372565b61023c61275e565b61023c600480360360608110156105b957600080fd5b508035906020810135906040013560ff16612766565b8180156105db57600080fd5b5061023c600480360360808110156105f257600080fd5b5080359060208101359060ff60408201351690606001356127f0565b61023c612bb9565b601281565b60608260070180548060200260200160405190810160405280929190818152602001828054801561067557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610657575b5050505050905060005b815181101561076457600082828151811061069657fe5b60200260200101519050600061073e8660090184815481106106b457fe5b9060005260206000200154836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561070c57600080fd5b505afa158015610720573d6000803e3d6000fd5b505050506040513d602081101561073657600080fd5b505190612bc2565b9050801561075a5761075a6001600160a01b0383168683612c1f565b505060010161067f565b50505050565b60006060846007018054806020026020016040519081016040528092919081815260200182805480156107c657602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116107a8575b505050505090508051845114610823576040805162461bcd60e51b815260206004820181905260248201527f416d6f756e7473206d757374206d6174636820706f6f6c656420746f6b656e73604482015290519081900360640190fd5b61082b613f0e565b60405180610100016040528060008152602001600081526020016000815260200161085588612ca4565b81526020018760060160009054906101000a90046001600160a01b03166001600160a01b0316815260200160008152602001876009018054806020026020016040519081016040528092919081815260200182805480156108d557602002820191906000526020600020905b8154815260200190600101908083116108c1575b505050505081526020018760080180548060200260200160405190810160405280929190818152602001828054801561092d57602002820191906000526020600020905b815481526020019060010190808311610919575b5050505050815250905080608001516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561097457600080fd5b505afa158015610988573d6000803e3d6000fd5b505050506040513d602081101561099e57600080fd5b505160a08201819052156109ce576109cb6109c18260c001518360e00151612caf565b8260600151612da6565b81525b6060825167ffffffffffffffff811180156109e857600080fd5b50604051908082528060200260200182016040528015610a12578160200160208202803683370190505b50905060005b8351811015610c685760a0830151151580610a4657506000878281518110610a3c57fe5b6020026020010151115b610a97576040805162461bcd60e51b815260206004820152601e60248201527f4d75737420737570706c7920616c6c20746f6b656e7320696e20706f6f6c0000604482015290519081900360640190fd5b868181518110610aa357fe5b6020026020010151600014610c0c576000848281518110610ac057fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015610b1457600080fd5b505afa158015610b28573d6000803e3d6000fd5b505050506040513d6020811015610b3e57600080fd5b50518851909150610b8e90339030908b9086908110610b5957fe5b6020026020010151888681518110610b6d57fe5b60200260200101516001600160a01b0316612f12909392919063ffffffff16565b610bf281868481518110610b9e57fe5b60200260200101516001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561070c57600080fd5b888381518110610bfe57fe5b602002602001018181525050505b610c49878281518110610c1b57fe5b60200260200101518460c001518381518110610c3357fe5b6020026020010151612f9a90919063ffffffff16565b828281518110610c5557fe5b6020908102919091010152600101610a18565b50610c84610c7a828460e00151612caf565b8360600151612da6565b60208301819052825110610cdf576040805162461bcd60e51b815260206004820152601160248201527f442073686f756c6420696e637265617365000000000000000000000000000000604482015290519081900360640190fd5b60208201516040830152825160609067ffffffffffffffff81118015610d0457600080fd5b50604051908082528060200260200182016040528015610d2e578160200160208202803683370190505b5090508260a00151600014610ecc576000610d4e89600401548651612ff4565b905060005b8551811015610ea5576000610d9a8660000151610d948860c001518581518110610d7957fe5b6020026020010151896020015161300f90919063ffffffff16565b90613068565b9050610dd56402540be400610d94610dce888681518110610db757fe5b6020026020010151856130cf90919063ffffffff16565b869061300f565b848381518110610de157fe5b602002602001018181525050610e45610e236402540be400610d948e60050154888781518110610e0d57fe5b602002602001015161300f90919063ffffffff16565b868481518110610e2f57fe5b6020026020010151612bc290919063ffffffff16565b8b6009018381548110610e5457fe5b9060005260206000200181905550610e85848381518110610e7157fe5b6020026020010151868481518110610e2f57fe5b858381518110610e9157fe5b602090810291909101015250600101610d53565b50610ec1610eb7848660e00151612caf565b8560600151612da6565b604085015250610ee3565b8151610ee19060098a01906020850190613f5c565b505b60008360a0015160001415610efd57506020830151610f25565b835160a08501516040860151610f229291610d9491610f1c9084612bc2565b9061300f565b90505b86811015610f7a576040805162461bcd60e51b815260206004820152601b60248201527f436f756c646e2774206d696e74206d696e207265717565737465640000000000604482015290519081900360640190fd5b83608001516001600160a01b03166340c10f1933836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b158015610fd557600080fd5b505af1158015610fe9573d6000803e3d6000fd5b50505050336001600160a01b03167f189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a289848760200151611036868a60a00151612f9a90919063ffffffff16565b604051808060200180602001858152602001848152602001838103835287818151815260200191508051906020019060200280838360005b8381101561108657818101518382015260200161106e565b50505050905001838103825286818151815260200191508051906020019060200280838360005b838110156110c55781810151838201526020016110ad565b50505050905001965050505050505060405180910390a29450505050505b9392505050565b60006110f4613f0e565b60405180610100016040528060008152602001600081526020016000815260200161111e87612ca4565b81526020018660060160009054906101000a90046001600160a01b03166001600160a01b03168152602001600081526020018660090180548060200260200160405190810160405280929190818152602001828054801561119e57602002820191906000526020600020905b81548152602001906001019080831161118a575b50505050508152602001866008018054806020026020016040519081016040528092919081815260200182805480156111f657602002820191906000526020600020905b8154815260200190600101908083116111e2575b5050505050815250905080608001516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561123d57600080fd5b505afa158015611251573d6000803e3d6000fd5b505050506040513d602081101561126757600080fd5b505160a082015260078501805460408051602083810282018101909252828152606093909290918301828280156112c757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116112a9575b505050505090508051855114611324576040805162461bcd60e51b815260206004820181905260248201527f416d6f756e74732073686f756c64206d6174636820706f6f6c20746f6b656e73604482015290519081900360640190fd5b81608001516001600160a01b03166370a08231336040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561137557600080fd5b505afa158015611389573d6000803e3d6000fd5b505050506040513d602081101561139f57600080fd5b505184118015906113af57508315155b611400576040805162461bcd60e51b815260206004820152600d60248201527f3e4c502e62616c616e63654f6600000000000000000000000000000000000000604482015290519081900360640190fd5b600061141187600401548351612ff4565b90506060825167ffffffffffffffff8111801561142d57600080fd5b50604051908082528060200260200182016040528015611457578160200160208202803683370190505b5090506060835167ffffffffffffffff8111801561147457600080fd5b5060405190808252806020026020018201604052801561149e578160200160208202803683370190505b5090506114c06114b68660c001518760e00151612caf565b8660600151612da6565b855260005b8451811015611544576115258982815181106114dd57fe5b602002602001015160405180606001604052806023815260200161405d602391398860c00151848151811061150e57fe5b60200260200101516130e79092919063ffffffff16565b82828151811061153157fe5b60209081029190910101526001016114c5565b506115566114b6828760e00151612caf565b602086015260005b845181101561168357600061159f8760000151610d948960c00151858151811061158457fe5b60200260200101518a6020015161300f90919063ffffffff16565b905060006115c98484815181106115b257fe5b6020026020010151836130cf90919063ffffffff16565b90506115de6402540be400610d94888461300f565b8584815181106115ea57fe5b6020026020010181815250506116226116166402540be400610d948f60050154898881518110610e0d57fe5b858581518110610e2f57fe5b8c600901848154811061163157fe5b906000526020600020018190555061166285848151811061164e57fe5b6020026020010151858581518110610e2f57fe5b84848151811061166e57fe5b6020908102919091010152505060010161155e565b506116956114b6828760e00151612caf565b60408601819052855160a0870151600093506116bb92610d949190610f1c908490612bc2565b90508061170f576040805162461bcd60e51b815260206004820152601b60248201527f4275726e7420616d6f756e742063616e6e6f74206265207a65726f0000000000604482015290519081900360640190fd5b61171a816001612f9a565b905086811115611771576040805162461bcd60e51b815260206004820152601b60248201527f746f6b656e416d6f756e74203e206d61784275726e416d6f756e740000000000604482015290519081900360640190fd5b84608001516001600160a01b03166379cc679033836040518363ffffffff1660e01b815260040180836001600160a01b0316815260200182815260200192505050600060405180830381600087803b1580156117cc57600080fd5b505af11580156117e0573d6000803e3d6000fd5b5050505060005b845181101561183c57611834338a838151811061180057fe5b602002602001015187848151811061181457fe5b60200260200101516001600160a01b0316612c1f9092919063ffffffff16565b6001016117e7565b50336001600160a01b03167f3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af175589848860200151611036868b60a00151612bc290919063ffffffff16565b6305f5e1008111156118df576040805162461bcd60e51b815260206004820152600f60248201527f46656520697320746f6f20686967680000000000000000000000000000000000604482015290519081900360640190fd5b600482018190556040805182815290517fd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe59181900360200190a15050565b600061197e858585858960090180548060200260200160405190810160405280929190818152602001828054801561197457602002820191906000526020600020905b815481526020019060010190808311611960575b505050505061317e565b5095945050505050565b6402540be4008111156119e2576040805162461bcd60e51b815260206004820152600f60248201527f46656520697320746f6f20686967680000000000000000000000000000000000604482015290519081900360640190fd5b600582018190556040805182815290517fab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f389181900360200190a15050565b6060611afa83600901805480602002602001604051908101604052809291908181526020018280548015611a7357602002820191906000526020600020905b815481526020019060010190808311611a5f575b5050505050838560060160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015611ac957600080fd5b505afa158015611add573d6000803e3d6000fd5b505050506040513d6020811015611af357600080fd5b5051613326565b90505b92915050565b600080611b20611b1284613403565b611b1b85612ca4565b612da6565b905060008360060160009054906101000a90046001600160a01b031690506000816001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b7957600080fd5b505afa158015611b8d573d6000803e3d6000fd5b505050506040513d6020811015611ba357600080fd5b505190508015611bcd57611bc381610d9485670de0b6b3a764000061300f565b9350505050611bd5565b600093505050505b919050565b6006840154600785018054604080516020808402820181019092528281526060946001600160a01b03169385939192909190830182828015611c4557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311611c27575b50505050509050816001600160a01b03166370a08231336040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611c9957600080fd5b505afa158015611cad573d6000803e3d6000fd5b505050506040513d6020811015611cc357600080fd5b5051861115611d19576040805162461bcd60e51b815260206004820152600d60248201527f3e4c502e62616c616e63654f6600000000000000000000000000000000000000604482015290519081900360640190fd5b80518414611d6e576040805162461bcd60e51b815260206004820181905260248201527f6d696e416d6f756e7473206d757374206d6174636820706f6f6c546f6b656e73604482015290519081900360640190fd5b606087600901805480602002602001604051908101604052809291908181526020018280548015611dbe57602002820191906000526020600020905b815481526020019060010190808311611daa575b505050505090506000836001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b158015611e0057600080fd5b505afa158015611e14573d6000803e3d6000fd5b505050506040513d6020811015611e2a57600080fd5b505190506060611e3b838a84613326565b905060005b8151811015611f1b57888882818110611e5557fe5b90506020020135828281518110611e6857fe5b60200260200101511015611ec3576040805162461bcd60e51b815260206004820152601a60248201527f616d6f756e74735b695d203c206d696e416d6f756e74735b695d000000000000604482015290519081900360640190fd5b611ee6828281518110611ed257fe5b6020026020010151858381518110610e2f57fe5b8b6009018281548110611ef557fe5b9060005260206000200181905550611f133383838151811061180057fe5b600101611e40565b50604080517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018b905290516001600160a01b038716916379cc679091604480830192600092919082900301818387803b158015611f8357600080fd5b505af1158015611f97573d6000803e3d6000fd5b503392507f88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef9150839050611fcb858d612bc2565b6040518080602001838152602001828103825284818151815260200191508051906020019060200280838360005b83811015612011578181015183820152602001611ff9565b50505050905001935050505060405180910390a29450505050505b949350505050565b6007820154600090821061208f576040805162461bcd60e51b815260206004820152601860248201527f546f6b656e20696e646578206f7574206f662072616e67650000000000000000604482015290519081900360640190fd5b611afa8360090183815481106120a157fe5b90600052602060002001548460070184815481106120bb57fe5b60009182526020918290200154604080517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015290516001600160a01b03909216926370a0823192602480840193829003018186803b15801561070c57600080fd5b60008061212f86612ca4565b905060608660090180548060200260200160405190810160405280929190818152602001828054801561218157602002820191906000526020600020905b81548152602001906001019080831161216d575b505050505090506060876008018054806020026020016040519081016040528092919081815260200182805480156121d857602002820191906000526020600020905b8154815260200190600101908083116121c4575b5050505050905060006121f46121ee8484612caf565b85612da6565b905060005b83518110156122a35786156122475761222a89898381811061221757fe5b90506020020135858381518110610c3357fe5b84828151811061223657fe5b60200260200101818152505061229b565b61228289898381811061225657fe5b9050602002013560405180606001604052806023815260200161405d6023913986848151811061150e57fe5b84828151811061228e57fe5b6020026020010181815250505b6001016121f9565b5060006122b96122b38585612caf565b86612da6565b905060008a60060160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561230d57600080fd5b505afa158015612321573d6000803e3d6000fd5b505050506040513d602081101561233757600080fd5b5051905087156123605761235383610d9483610f1c8684612bc2565b965050505050505061202c565b61235383610d9483610f1c8387612bc2565b600080866007018660ff168154811061238757fe5b60009182526020918290200154604080517f70a0823100000000000000000000000000000000000000000000000000000000815233600482015290516001600160a01b03909216935083926370a0823192602480840193829003018186803b1580156123f257600080fd5b505afa158015612406573d6000803e3d6000fd5b505050506040513d602081101561241c57600080fd5b5051841115612472576040805162461bcd60e51b815260206004820152601d60248201527f43616e6e6f742073776170206d6f7265207468616e20796f75206f776e000000604482015290519081900360640190fd5b6000816001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156124c157600080fd5b505afa1580156124d5573d6000803e3d6000fd5b505050506040513d60208110156124eb57600080fd5b505190506125046001600160a01b038316333088612f12565b61255581836001600160a01b03166370a08231306040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561070c57600080fd5b945050506000806060886009018054806020026020016040519081016040528092919081815260200182805480156125ac57602002820191906000526020600020905b815481526020019060010190808311612598575b505050505090506125c0898989898561317e565b90935091508483101561261a576040805162461bcd60e51b815260206004820181905260248201527f53776170206469646e277420726573756c7420696e206d696e20746f6b656e73604482015290519081900360640190fd5b600061265c8a6008018960ff168154811061263157fe5b9060005260206000200154610d946402540be400610d948e600501548861300f90919063ffffffff16565b905061267187838b60ff1681518110610c3357fe5b8a6009018a60ff168154811061268357fe5b90600052602060002001819055506126ae816126a886858c60ff1681518110610e2f57fe5b90612bc2565b8a6009018960ff16815481106126c057fe5b906000526020600020018190555061270133858c6007018b60ff16815481106126e557fe5b6000918252602090912001546001600160a01b03169190612c1f565b604080518881526020810186905260ff808c16828401528a166060820152905133917fc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38919081900360800190a25091925050505b95945050505050565b6305f5e10081565b60008061197e8585858860060160009054906101000a90046001600160a01b03166001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156127bf57600080fd5b505afa1580156127d3573d6000803e3d6000fd5b505050506040513d60208110156127e957600080fd5b50516134b3565b6006840154600785018054604080516020808402820181019092528281526000946001600160a01b031693606093919290919083018282801561285c57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161283e575b50505050509050816001600160a01b03166370a08231336040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156128b057600080fd5b505afa1580156128c4573d6000803e3d6000fd5b505050506040513d60208110156128da57600080fd5b5051861115612930576040805162461bcd60e51b815260206004820152600d60248201527f3e4c502e62616c616e63654f6600000000000000000000000000000000000000604482015290519081900360640190fd5b80518560ff1610612988576040805162461bcd60e51b815260206004820152600f60248201527f546f6b656e206e6f7420666f756e640000000000000000000000000000000000604482015290519081900360640190fd5b6000826001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156129c357600080fd5b505afa1580156129d7573d6000803e3d6000fd5b505050506040513d60208110156129ed57600080fd5b50519050600080612a008a8a8a866134b3565b9150915086821015612a59576040805162461bcd60e51b815260206004820152600e60248201527f6479203c206d696e416d6f756e74000000000000000000000000000000000000604482015290519081900360640190fd5b612ab1612a86612a7f6402540be400610d948e600501548661300f90919063ffffffff16565b8490612f9a565b8b6009018a60ff1681548110612a9857fe5b9060005260206000200154612bc290919063ffffffff16565b8a6009018960ff1681548110612ac357fe5b6000918252602082200191909155604080517f79cc6790000000000000000000000000000000000000000000000000000000008152336004820152602481018c905290516001600160a01b038816926379cc6790926044808201939182900301818387803b158015612b3457600080fd5b505af1158015612b48573d6000803e3d6000fd5b50505050612b603383868b60ff168151811061181457fe5b604080518a81526020810185905260ff8a168183015260608101849052905133917f43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64919081900360800190a25098975050505050505050565b6402540be40081565b600082821115612c19576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb00000000000000000000000000000000000000000000000000000000179052612c9f908490613515565b505050565b6000611afd826135c6565b81518151606091908114612d0a576040805162461bcd60e51b815260206004820152601f60248201527f42616c616e636573206d757374206d61746368206d756c7469706c6965727300604482015290519081900360640190fd5b60608167ffffffffffffffff81118015612d2357600080fd5b50604051908082528060200260200182016040528015612d4d578160200160208202803683370190505b50905060005b82811015612d9d57612d7e858281518110612d6a57fe5b6020026020010151878381518110610e0d57fe5b828281518110612d8a57fe5b6020908102919091010152600101612d53565b50949350505050565b815160009081805b82811015612de657612ddc868281518110612dc557fe5b602002602001015183612f9a90919063ffffffff16565b9150600101612dae565b5080612df757600092505050611afd565b60008181612e05878661300f565b905060005b610100811015612ec4578260005b87811015612e4757612e3d612e33898d8481518110610e0d57fe5b610d94848861300f565b9150600101612e18565b509293508392612e9b612e7b612e6283610f1c8b6001612f9a565b612e756064610d9489610f1c8a84612bc2565b90612f9a565b610d9486610f1c612e8c868d61300f565b612e756064610d948b8f61300f565b9350612ea7848661365b565b15612ebb5783975050505050505050611afd565b50600101612e0a565b506040805162461bcd60e51b815260206004820152601360248201527f4420646f6573206e6f7420636f6e766572676500000000000000000000000000604482015290519081900360640190fd5b604080516001600160a01b0380861660248301528416604482015260648082018490528251808303909101815260849091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f23b872dd00000000000000000000000000000000000000000000000000000000179052610764908590613515565b600082820183811015611afa576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b6000611afa6130096004610f1c856001612bc2565b610d9485855b60008261301e57506000611afd565b8282028284828161302b57fe5b0414611afa5760405162461bcd60e51b81526004018080602001828103825260218152602001806140126021913960400191505060405180910390fd5b60008082116130be576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b8183816130c757fe5b049392505050565b6000818311156130e25750808203611afd565b500390565b600081848411156131765760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561313b578181015183820152602001613123565b50505050905090810190601f1680156131685780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000806060876008018054806020026020016040519081016040528092919081815260200182805480156131d157602002820191906000526020600020905b8154815260200190600101908083116131bd575b5050505050905060606131e48583612caf565b905080518860ff161080156131fc575080518760ff16105b61324d576040805162461bcd60e51b815260206004820152601860248201527f546f6b656e20696e646578206f7574206f662072616e67650000000000000000604482015290519081900360640190fd5b6000613292828a60ff168151811061326157fe5b6020026020010151612e75858c60ff168151811061327b57fe5b60200260200101518a61300f90919063ffffffff16565b905060006132ab6132a28c612ca4565b8b8b8587613672565b90506132c560016126a883868d60ff1681518110610e2f57fe5b95506132e76402540be400610d948d600401548961300f90919063ffffffff16565b9450613316848a60ff16815181106132fb57fe5b6020026020010151610d948789612bc290919063ffffffff16565b9550505050509550959350505050565b60608183111561337d576040805162461bcd60e51b815260206004820152601a60248201527f43616e6e6f742065786365656420746f74616c20737570706c79000000000000604482015290519081900360640190fd5b6060845167ffffffffffffffff8111801561339757600080fd5b506040519080825280602002602001820160405280156133c1578160200160208202803683370190505b50905060005b8551811015612d9d576133e484610d9487898581518110610e0d57fe5b8282815181106133f057fe5b60209081029190910101526001016133c7565b6060611afd8260090180548060200260200160405190810160405280929190818152602001828054801561345657602002820191906000526020600020905b815481526020019060010190808311613442575b5050505050836008018054806020026020016040519081016040528092919081815260200182805480156134a957602002820191906000526020600020905b815481526020019060010190808311613495575b5050505050612caf565b60008060008060006134c789888a896138b8565b8093508194508295505050506000613504846126a88c6008018b60ff16815481106134ee57fe5b600091825260209091200154610d948688612bc2565b939a93995092975050505050505050565b606061356a826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316613bb39092919063ffffffff16565b805190915015612c9f5780806020019051602081101561358957600080fd5b5051612c9f5760405162461bcd60e51b815260040180806020018281038252602a815260200180614033602a913960400191505060405180910390fd5b600381015460018201546000919042821115613652576002840154845480831115613624576136196136126135fb8685612bc2565b610d946136084287612bc2565b610f1c8887612bc2565b8290612f9a565b945050505050611bd5565b61361961364b6136348685612bc2565b610d946136414287612bc2565b610f1c8689612bc2565b8290612bc2565b9150611bd59050565b6000600161366984846130cf565b11159392505050565b805160009060ff86811690861614156136d2576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420636f6d7061726520746f6b656e20746f20697473656c66000000604482015290519081900360640190fd5b808660ff161080156136e65750808560ff16105b613737576040805162461bcd60e51b815260206004820152601660248201527f546f6b656e73206d75737420626520696e20706f6f6c00000000000000000000604482015290519081900360640190fd5b60006137438489612da6565b905080600080613753858c61300f565b90506000805b868110156137cc578b60ff168114156137745789915061379e565b8a60ff1681146137995788818151811061378a57fe5b6020026020010151915061379e565b6137c4565b6137a88483612f9a565b93506137c16137b7838961300f565b610d94878961300f565b94505b600101613759565b506137e96137da838861300f565b610d946064610f1c888a61300f565b935060006138066137ff84610d9489606461300f565b8590612f9a565b9050600086815b61010081101561386a5790915081906138406138328a6126a887612e7587600261300f565b610d948a612e75868061300f565b915061384c828461365b565b1561386257509850612755975050505050505050565b60010161380d565b506040805162461bcd60e51b815260206004820152601e60248201527f417070726f78696d6174696f6e20646964206e6f7420636f6e76657267650000604482015290519081900360640190fd5b600080600060606138c888613403565b905080518760ff1610613922576040805162461bcd60e51b815260206004820152601860248201527f546f6b656e20696e646578206f7574206f662072616e67650000000000000000604482015290519081900360640190fd5b61392a613fa7565b6040518060a00160405280600081526020016000815260200160008152602001600081526020016000815250905061396189612ca4565b60808201819052613973908390612da6565b8082526139939061398b908890610d94908b9061300f565b825190612bc2565b60208201528151829060ff8a169081106139a957fe5b6020026020010151871115613a05576040805162461bcd60e51b815260206004820152601a60248201527f5769746864726177206578636565647320617661696c61626c65000000000000604482015290519081900360640190fd5b613a19816080015189848460200151613bc2565b6040820152815160609067ffffffffffffffff81118015613a3957600080fd5b50604051908082528060200260200182016040528015613a63578160200160208202803683370190505b509050613a758a600401548451612ff4565b606083015260005b8351811015613b29576000848281518110613a9457fe5b60200260200101519050613b0961364b6402540be400610d9487606001518f60ff168714613ae257885160208a0151613add91613ad691610d94908a9061300f565b8790612bc2565b610f1c565b610f1c89604001516126a88b60000151610d948d602001518b61300f90919063ffffffff16565b838381518110613b1557fe5b602090810291909101015250600101613a7d565b506000613b52613b4384608001518c858760200151613bc2565b838c60ff1681518110610e2f57fe5b9050613b808b6008018b60ff1681548110613b6957fe5b600091825260209091200154610d94836001612bc2565b9050808360400151858c60ff1681518110613b9757fe5b6020026020010151965096509650505050509450945094915050565b606061202c8484600085613d28565b815160009060ff85168111613c1e576040805162461bcd60e51b815260206004820152600f60248201527f546f6b656e206e6f7420666f756e640000000000000000000000000000000000604482015290519081900360640190fd5b82600080613c2c898561300f565b905060005b84811015613c93578860ff168114613c8b57613c69888281518110613c5257fe5b602002602001015184612f9a90919063ffffffff16565b9250613c88613c7e868a8481518110610e0d57fe5b610d94868a61300f565b93505b600101613c31565b50613cb0613ca1828661300f565b610d946064610f1c878b61300f565b92506000613cc6612a7f83610d948a606461300f565b9050600087815b61010081101561386a579091508190613d00613cf28b6126a887612e7587600261300f565b610d9489612e75868061300f565b9150613d0c828461365b565b15613d205750965061202c95505050505050565b600101613ccd565b606082471015613d695760405162461bcd60e51b8152600401808060200182810382526026815260200180613fec6026913960400191505060405180910390fd5b613d7285613ea2565b613dc3576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b60006060866001600160a01b031685876040518082805190602001908083835b60208310613e2057805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101613de3565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114613e82576040519150601f19603f3d011682016040523d82523d6000602084013e613e87565b606091505b5091509150613e97828286613ea8565b979650505050505050565b3b151590565b60608315613eb75750816110e3565b825115613ec75782518084602001fd5b60405162461bcd60e51b815260206004820181815284516024840152845185939192839260440191908501908083836000831561313b578181015183820152602001613123565b6040518061010001604052806000815260200160008152602001600081526020016000815260200160006001600160a01b031681526020016000815260200160608152602001606081525090565b828054828255906000526020600020908101928215613f97579160200282015b82811115613f97578251825591602001919060010190613f7c565b50613fa3929150613fd6565b5090565b6040518060a0016040528060008152602001600081526020016000815260200160008152602001600081525090565b5b80821115613fa35760008155600101613fd756fe416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f775361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656443616e6e6f74207769746864726177206d6f7265207468616e20617661696c61626c65a2646970667358221220be351b008ac7c57dd94b1b14ca8208af4731752836c64efe950fa2f5d030372864736f6c634300060c0033",
}

// SwapUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapUtilsMetaData.ABI instead.
var SwapUtilsABI = SwapUtilsMetaData.ABI

// Deprecated: Use SwapUtilsMetaData.Sigs instead.
// SwapUtilsFuncSigs maps the 4-byte function signature to its string representation.
var SwapUtilsFuncSigs = SwapUtilsMetaData.Sigs

// SwapUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SwapUtilsMetaData.Bin instead.
var SwapUtilsBin = SwapUtilsMetaData.Bin

// DeploySwapUtils deploys a new Ethereum contract, binding an instance of SwapUtils to it.
func DeploySwapUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SwapUtils, error) {
	parsed, err := SwapUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SwapUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SwapUtils{SwapUtilsCaller: SwapUtilsCaller{contract: contract}, SwapUtilsTransactor: SwapUtilsTransactor{contract: contract}, SwapUtilsFilterer: SwapUtilsFilterer{contract: contract}}, nil
}

// SwapUtils is an auto generated Go binding around an Ethereum contract.
type SwapUtils struct {
	SwapUtilsCaller     // Read-only binding to the contract
	SwapUtilsTransactor // Write-only binding to the contract
	SwapUtilsFilterer   // Log filterer for contract events
}

// SwapUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapUtilsSession struct {
	Contract     *SwapUtils        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapUtilsCallerSession struct {
	Contract *SwapUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SwapUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapUtilsTransactorSession struct {
	Contract     *SwapUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SwapUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapUtilsRaw struct {
	Contract *SwapUtils // Generic contract binding to access the raw methods on
}

// SwapUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapUtilsCallerRaw struct {
	Contract *SwapUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// SwapUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapUtilsTransactorRaw struct {
	Contract *SwapUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapUtils creates a new instance of SwapUtils, bound to a specific deployed contract.
func NewSwapUtils(address common.Address, backend bind.ContractBackend) (*SwapUtils, error) {
	contract, err := bindSwapUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapUtils{SwapUtilsCaller: SwapUtilsCaller{contract: contract}, SwapUtilsTransactor: SwapUtilsTransactor{contract: contract}, SwapUtilsFilterer: SwapUtilsFilterer{contract: contract}}, nil
}

// NewSwapUtilsCaller creates a new read-only instance of SwapUtils, bound to a specific deployed contract.
func NewSwapUtilsCaller(address common.Address, caller bind.ContractCaller) (*SwapUtilsCaller, error) {
	contract, err := bindSwapUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsCaller{contract: contract}, nil
}

// NewSwapUtilsTransactor creates a new write-only instance of SwapUtils, bound to a specific deployed contract.
func NewSwapUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapUtilsTransactor, error) {
	contract, err := bindSwapUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsTransactor{contract: contract}, nil
}

// NewSwapUtilsFilterer creates a new log filterer instance of SwapUtils, bound to a specific deployed contract.
func NewSwapUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapUtilsFilterer, error) {
	contract, err := bindSwapUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsFilterer{contract: contract}, nil
}

// bindSwapUtils binds a generic wrapper to an already deployed contract.
func bindSwapUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwapUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapUtils *SwapUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapUtils.Contract.SwapUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapUtils *SwapUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapUtils.Contract.SwapUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapUtils *SwapUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapUtils.Contract.SwapUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapUtils *SwapUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapUtils *SwapUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapUtils *SwapUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapUtils.Contract.contract.Transact(opts, method, params...)
}

// MAXADMINFEE is a free data retrieval call binding the contract method 0xf3de0362.
//
// Solidity: function MAX_ADMIN_FEE() view returns(uint256)
func (_SwapUtils *SwapUtilsCaller) MAXADMINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapUtils.contract.Call(opts, &out, "MAX_ADMIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADMINFEE is a free data retrieval call binding the contract method 0xf3de0362.
//
// Solidity: function MAX_ADMIN_FEE() view returns(uint256)
func (_SwapUtils *SwapUtilsSession) MAXADMINFEE() (*big.Int, error) {
	return _SwapUtils.Contract.MAXADMINFEE(&_SwapUtils.CallOpts)
}

// MAXADMINFEE is a free data retrieval call binding the contract method 0xf3de0362.
//
// Solidity: function MAX_ADMIN_FEE() view returns(uint256)
func (_SwapUtils *SwapUtilsCallerSession) MAXADMINFEE() (*big.Int, error) {
	return _SwapUtils.Contract.MAXADMINFEE(&_SwapUtils.CallOpts)
}

// MAXSWAPFEE is a free data retrieval call binding the contract method 0xab3d8544.
//
// Solidity: function MAX_SWAP_FEE() view returns(uint256)
func (_SwapUtils *SwapUtilsCaller) MAXSWAPFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SwapUtils.contract.Call(opts, &out, "MAX_SWAP_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSWAPFEE is a free data retrieval call binding the contract method 0xab3d8544.
//
// Solidity: function MAX_SWAP_FEE() view returns(uint256)
func (_SwapUtils *SwapUtilsSession) MAXSWAPFEE() (*big.Int, error) {
	return _SwapUtils.Contract.MAXSWAPFEE(&_SwapUtils.CallOpts)
}

// MAXSWAPFEE is a free data retrieval call binding the contract method 0xab3d8544.
//
// Solidity: function MAX_SWAP_FEE() view returns(uint256)
func (_SwapUtils *SwapUtilsCallerSession) MAXSWAPFEE() (*big.Int, error) {
	return _SwapUtils.Contract.MAXSWAPFEE(&_SwapUtils.CallOpts)
}

// POOLPRECISIONDECIMALS is a free data retrieval call binding the contract method 0x0296ab50.
//
// Solidity: function POOL_PRECISION_DECIMALS() view returns(uint8)
func (_SwapUtils *SwapUtilsCaller) POOLPRECISIONDECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SwapUtils.contract.Call(opts, &out, "POOL_PRECISION_DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// POOLPRECISIONDECIMALS is a free data retrieval call binding the contract method 0x0296ab50.
//
// Solidity: function POOL_PRECISION_DECIMALS() view returns(uint8)
func (_SwapUtils *SwapUtilsSession) POOLPRECISIONDECIMALS() (uint8, error) {
	return _SwapUtils.Contract.POOLPRECISIONDECIMALS(&_SwapUtils.CallOpts)
}

// POOLPRECISIONDECIMALS is a free data retrieval call binding the contract method 0x0296ab50.
//
// Solidity: function POOL_PRECISION_DECIMALS() view returns(uint8)
func (_SwapUtils *SwapUtilsCallerSession) POOLPRECISIONDECIMALS() (uint8, error) {
	return _SwapUtils.Contract.POOLPRECISIONDECIMALS(&_SwapUtils.CallOpts)
}

// SwapUtilsAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the SwapUtils contract.
type SwapUtilsAddLiquidityIterator struct {
	Event *SwapUtilsAddLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUtilsAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUtilsAddLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUtilsAddLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUtilsAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUtilsAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUtilsAddLiquidity represents a AddLiquidity event raised by the SwapUtils contract.
type SwapUtilsAddLiquidity struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	Fees          []*big.Int
	Invariant     *big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) FilterAddLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SwapUtilsAddLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.FilterLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsAddLiquidityIterator{contract: _SwapUtils.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *SwapUtilsAddLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.WatchLogs(opts, "AddLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUtilsAddLiquidity)
				if err := _SwapUtils.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAddLiquidity is a log parse operation binding the contract event 0x189c623b666b1b45b83d7178f39b8c087cb09774317ca2f53c2d3c3726f222a2.
//
// Solidity: event AddLiquidity(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) ParseAddLiquidity(log types.Log) (*SwapUtilsAddLiquidity, error) {
	event := new(SwapUtilsAddLiquidity)
	if err := _SwapUtils.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUtilsNewAdminFeeIterator is returned from FilterNewAdminFee and is used to iterate over the raw logs and unpacked data for NewAdminFee events raised by the SwapUtils contract.
type SwapUtilsNewAdminFeeIterator struct {
	Event *SwapUtilsNewAdminFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUtilsNewAdminFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUtilsNewAdminFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUtilsNewAdminFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUtilsNewAdminFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUtilsNewAdminFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUtilsNewAdminFee represents a NewAdminFee event raised by the SwapUtils contract.
type SwapUtilsNewAdminFee struct {
	NewAdminFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewAdminFee is a free log retrieval operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_SwapUtils *SwapUtilsFilterer) FilterNewAdminFee(opts *bind.FilterOpts) (*SwapUtilsNewAdminFeeIterator, error) {

	logs, sub, err := _SwapUtils.contract.FilterLogs(opts, "NewAdminFee")
	if err != nil {
		return nil, err
	}
	return &SwapUtilsNewAdminFeeIterator{contract: _SwapUtils.contract, event: "NewAdminFee", logs: logs, sub: sub}, nil
}

// WatchNewAdminFee is a free log subscription operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_SwapUtils *SwapUtilsFilterer) WatchNewAdminFee(opts *bind.WatchOpts, sink chan<- *SwapUtilsNewAdminFee) (event.Subscription, error) {

	logs, sub, err := _SwapUtils.contract.WatchLogs(opts, "NewAdminFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUtilsNewAdminFee)
				if err := _SwapUtils.contract.UnpackLog(event, "NewAdminFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewAdminFee is a log parse operation binding the contract event 0xab599d640ca80cde2b09b128a4154a8dfe608cb80f4c9399c8b954b01fd35f38.
//
// Solidity: event NewAdminFee(uint256 newAdminFee)
func (_SwapUtils *SwapUtilsFilterer) ParseNewAdminFee(log types.Log) (*SwapUtilsNewAdminFee, error) {
	event := new(SwapUtilsNewAdminFee)
	if err := _SwapUtils.contract.UnpackLog(event, "NewAdminFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUtilsNewSwapFeeIterator is returned from FilterNewSwapFee and is used to iterate over the raw logs and unpacked data for NewSwapFee events raised by the SwapUtils contract.
type SwapUtilsNewSwapFeeIterator struct {
	Event *SwapUtilsNewSwapFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUtilsNewSwapFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUtilsNewSwapFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUtilsNewSwapFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUtilsNewSwapFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUtilsNewSwapFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUtilsNewSwapFee represents a NewSwapFee event raised by the SwapUtils contract.
type SwapUtilsNewSwapFee struct {
	NewSwapFee *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewSwapFee is a free log retrieval operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_SwapUtils *SwapUtilsFilterer) FilterNewSwapFee(opts *bind.FilterOpts) (*SwapUtilsNewSwapFeeIterator, error) {

	logs, sub, err := _SwapUtils.contract.FilterLogs(opts, "NewSwapFee")
	if err != nil {
		return nil, err
	}
	return &SwapUtilsNewSwapFeeIterator{contract: _SwapUtils.contract, event: "NewSwapFee", logs: logs, sub: sub}, nil
}

// WatchNewSwapFee is a free log subscription operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_SwapUtils *SwapUtilsFilterer) WatchNewSwapFee(opts *bind.WatchOpts, sink chan<- *SwapUtilsNewSwapFee) (event.Subscription, error) {

	logs, sub, err := _SwapUtils.contract.WatchLogs(opts, "NewSwapFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUtilsNewSwapFee)
				if err := _SwapUtils.contract.UnpackLog(event, "NewSwapFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewSwapFee is a log parse operation binding the contract event 0xd88ea5155021c6f8dafa1a741e173f595cdf77ce7c17d43342131d7f06afdfe5.
//
// Solidity: event NewSwapFee(uint256 newSwapFee)
func (_SwapUtils *SwapUtilsFilterer) ParseNewSwapFee(log types.Log) (*SwapUtilsNewSwapFee, error) {
	event := new(SwapUtilsNewSwapFee)
	if err := _SwapUtils.contract.UnpackLog(event, "NewSwapFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUtilsRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the SwapUtils contract.
type SwapUtilsRemoveLiquidityIterator struct {
	Event *SwapUtilsRemoveLiquidity // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUtilsRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUtilsRemoveLiquidity)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUtilsRemoveLiquidity)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUtilsRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUtilsRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUtilsRemoveLiquidity represents a RemoveLiquidity event raised by the SwapUtils contract.
type SwapUtilsRemoveLiquidity struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, provider []common.Address) (*SwapUtilsRemoveLiquidityIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.FilterLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsRemoveLiquidityIterator{contract: _SwapUtils.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *SwapUtilsRemoveLiquidity, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.WatchLogs(opts, "RemoveLiquidity", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUtilsRemoveLiquidity)
				if err := _SwapUtils.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x88d38ed598fdd809c2bf01ee49cd24b7fdabf379a83d29567952b60324d58cef.
//
// Solidity: event RemoveLiquidity(address indexed provider, uint256[] tokenAmounts, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) ParseRemoveLiquidity(log types.Log) (*SwapUtilsRemoveLiquidity, error) {
	event := new(SwapUtilsRemoveLiquidity)
	if err := _SwapUtils.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUtilsRemoveLiquidityImbalanceIterator is returned from FilterRemoveLiquidityImbalance and is used to iterate over the raw logs and unpacked data for RemoveLiquidityImbalance events raised by the SwapUtils contract.
type SwapUtilsRemoveLiquidityImbalanceIterator struct {
	Event *SwapUtilsRemoveLiquidityImbalance // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUtilsRemoveLiquidityImbalanceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUtilsRemoveLiquidityImbalance)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUtilsRemoveLiquidityImbalance)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUtilsRemoveLiquidityImbalanceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUtilsRemoveLiquidityImbalanceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUtilsRemoveLiquidityImbalance represents a RemoveLiquidityImbalance event raised by the SwapUtils contract.
type SwapUtilsRemoveLiquidityImbalance struct {
	Provider      common.Address
	TokenAmounts  []*big.Int
	Fees          []*big.Int
	Invariant     *big.Int
	LpTokenSupply *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityImbalance is a free log retrieval operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) FilterRemoveLiquidityImbalance(opts *bind.FilterOpts, provider []common.Address) (*SwapUtilsRemoveLiquidityImbalanceIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.FilterLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsRemoveLiquidityImbalanceIterator{contract: _SwapUtils.contract, event: "RemoveLiquidityImbalance", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityImbalance is a free log subscription operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) WatchRemoveLiquidityImbalance(opts *bind.WatchOpts, sink chan<- *SwapUtilsRemoveLiquidityImbalance, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.WatchLogs(opts, "RemoveLiquidityImbalance", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUtilsRemoveLiquidityImbalance)
				if err := _SwapUtils.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidityImbalance is a log parse operation binding the contract event 0x3631c28b1f9dd213e0319fb167b554d76b6c283a41143eb400a0d1adb1af1755.
//
// Solidity: event RemoveLiquidityImbalance(address indexed provider, uint256[] tokenAmounts, uint256[] fees, uint256 invariant, uint256 lpTokenSupply)
func (_SwapUtils *SwapUtilsFilterer) ParseRemoveLiquidityImbalance(log types.Log) (*SwapUtilsRemoveLiquidityImbalance, error) {
	event := new(SwapUtilsRemoveLiquidityImbalance)
	if err := _SwapUtils.contract.UnpackLog(event, "RemoveLiquidityImbalance", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUtilsRemoveLiquidityOneIterator is returned from FilterRemoveLiquidityOne and is used to iterate over the raw logs and unpacked data for RemoveLiquidityOne events raised by the SwapUtils contract.
type SwapUtilsRemoveLiquidityOneIterator struct {
	Event *SwapUtilsRemoveLiquidityOne // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUtilsRemoveLiquidityOneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUtilsRemoveLiquidityOne)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUtilsRemoveLiquidityOne)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUtilsRemoveLiquidityOneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUtilsRemoveLiquidityOneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUtilsRemoveLiquidityOne represents a RemoveLiquidityOne event raised by the SwapUtils contract.
type SwapUtilsRemoveLiquidityOne struct {
	Provider      common.Address
	LpTokenAmount *big.Int
	LpTokenSupply *big.Int
	BoughtId      *big.Int
	TokensBought  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidityOne is a free log retrieval operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_SwapUtils *SwapUtilsFilterer) FilterRemoveLiquidityOne(opts *bind.FilterOpts, provider []common.Address) (*SwapUtilsRemoveLiquidityOneIterator, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.FilterLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsRemoveLiquidityOneIterator{contract: _SwapUtils.contract, event: "RemoveLiquidityOne", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidityOne is a free log subscription operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_SwapUtils *SwapUtilsFilterer) WatchRemoveLiquidityOne(opts *bind.WatchOpts, sink chan<- *SwapUtilsRemoveLiquidityOne, provider []common.Address) (event.Subscription, error) {

	var providerRule []interface{}
	for _, providerItem := range provider {
		providerRule = append(providerRule, providerItem)
	}

	logs, sub, err := _SwapUtils.contract.WatchLogs(opts, "RemoveLiquidityOne", providerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUtilsRemoveLiquidityOne)
				if err := _SwapUtils.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRemoveLiquidityOne is a log parse operation binding the contract event 0x43fb02998f4e03da2e0e6fff53fdbf0c40a9f45f145dc377fc30615d7d7a8a64.
//
// Solidity: event RemoveLiquidityOne(address indexed provider, uint256 lpTokenAmount, uint256 lpTokenSupply, uint256 boughtId, uint256 tokensBought)
func (_SwapUtils *SwapUtilsFilterer) ParseRemoveLiquidityOne(log types.Log) (*SwapUtilsRemoveLiquidityOne, error) {
	event := new(SwapUtilsRemoveLiquidityOne)
	if err := _SwapUtils.contract.UnpackLog(event, "RemoveLiquidityOne", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwapUtilsTokenSwapIterator is returned from FilterTokenSwap and is used to iterate over the raw logs and unpacked data for TokenSwap events raised by the SwapUtils contract.
type SwapUtilsTokenSwapIterator struct {
	Event *SwapUtilsTokenSwap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SwapUtilsTokenSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwapUtilsTokenSwap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SwapUtilsTokenSwap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SwapUtilsTokenSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwapUtilsTokenSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwapUtilsTokenSwap represents a TokenSwap event raised by the SwapUtils contract.
type SwapUtilsTokenSwap struct {
	Buyer        common.Address
	TokensSold   *big.Int
	TokensBought *big.Int
	SoldId       *big.Int
	BoughtId     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokenSwap is a free log retrieval operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_SwapUtils *SwapUtilsFilterer) FilterTokenSwap(opts *bind.FilterOpts, buyer []common.Address) (*SwapUtilsTokenSwapIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _SwapUtils.contract.FilterLogs(opts, "TokenSwap", buyerRule)
	if err != nil {
		return nil, err
	}
	return &SwapUtilsTokenSwapIterator{contract: _SwapUtils.contract, event: "TokenSwap", logs: logs, sub: sub}, nil
}

// WatchTokenSwap is a free log subscription operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_SwapUtils *SwapUtilsFilterer) WatchTokenSwap(opts *bind.WatchOpts, sink chan<- *SwapUtilsTokenSwap, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _SwapUtils.contract.WatchLogs(opts, "TokenSwap", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwapUtilsTokenSwap)
				if err := _SwapUtils.contract.UnpackLog(event, "TokenSwap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenSwap is a log parse operation binding the contract event 0xc6c1e0630dbe9130cc068028486c0d118ddcea348550819defd5cb8c257f8a38.
//
// Solidity: event TokenSwap(address indexed buyer, uint256 tokensSold, uint256 tokensBought, uint128 soldId, uint128 boughtId)
func (_SwapUtils *SwapUtilsFilterer) ParseTokenSwap(log types.Log) (*SwapUtilsTokenSwap, error) {
	event := new(SwapUtilsTokenSwap)
	if err := _SwapUtils.contract.UnpackLog(event, "TokenSwap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
