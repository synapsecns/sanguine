// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package snapshotharness

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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fef610aaae7ff3bedc9071a08a621a860ce8e4792583b12c9ca1c0cb21e4499c64736f6c63430008110033",
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

// MerkleMathMetaData contains all meta data concerning the MerkleMath contract.
var MerkleMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122023db8815f4f7dda42fcf7e8806774098c334c5466c28d056d3ae9b01458b076f64736f6c63430008110033",
}

// MerkleMathABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleMathMetaData.ABI instead.
var MerkleMathABI = MerkleMathMetaData.ABI

// MerkleMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleMathMetaData.Bin instead.
var MerkleMathBin = MerkleMathMetaData.Bin

// DeployMerkleMath deploys a new Ethereum contract, binding an instance of MerkleMath to it.
func DeployMerkleMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleMath, error) {
	parsed, err := MerkleMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleMath{MerkleMathCaller: MerkleMathCaller{contract: contract}, MerkleMathTransactor: MerkleMathTransactor{contract: contract}, MerkleMathFilterer: MerkleMathFilterer{contract: contract}}, nil
}

// MerkleMath is an auto generated Go binding around an Ethereum contract.
type MerkleMath struct {
	MerkleMathCaller     // Read-only binding to the contract
	MerkleMathTransactor // Write-only binding to the contract
	MerkleMathFilterer   // Log filterer for contract events
}

// MerkleMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleMathSession struct {
	Contract     *MerkleMath       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleMathCallerSession struct {
	Contract *MerkleMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleMathTransactorSession struct {
	Contract     *MerkleMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleMathRaw struct {
	Contract *MerkleMath // Generic contract binding to access the raw methods on
}

// MerkleMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleMathCallerRaw struct {
	Contract *MerkleMathCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleMathTransactorRaw struct {
	Contract *MerkleMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleMath creates a new instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMath(address common.Address, backend bind.ContractBackend) (*MerkleMath, error) {
	contract, err := bindMerkleMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleMath{MerkleMathCaller: MerkleMathCaller{contract: contract}, MerkleMathTransactor: MerkleMathTransactor{contract: contract}, MerkleMathFilterer: MerkleMathFilterer{contract: contract}}, nil
}

// NewMerkleMathCaller creates a new read-only instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathCaller(address common.Address, caller bind.ContractCaller) (*MerkleMathCaller, error) {
	contract, err := bindMerkleMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleMathCaller{contract: contract}, nil
}

// NewMerkleMathTransactor creates a new write-only instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleMathTransactor, error) {
	contract, err := bindMerkleMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleMathTransactor{contract: contract}, nil
}

// NewMerkleMathFilterer creates a new log filterer instance of MerkleMath, bound to a specific deployed contract.
func NewMerkleMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleMathFilterer, error) {
	contract, err := bindMerkleMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleMathFilterer{contract: contract}, nil
}

// bindMerkleMath binds a generic wrapper to an already deployed contract.
func bindMerkleMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleMath *MerkleMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleMath.Contract.MerkleMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleMath *MerkleMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleMath.Contract.MerkleMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleMath *MerkleMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleMath.Contract.MerkleMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleMath *MerkleMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleMath *MerkleMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleMath *MerkleMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleMath.Contract.contract.Transact(opts, method, params...)
}

// SnapshotHarnessMetaData contains all meta data concerning the SnapshotHarness contract.
var SnapshotHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"calculateRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"statePayloads\",\"type\":\"bytes[]\"}],\"name\":\"formatSnapshot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isSnapshot\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"stateIndex\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"statesAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c9b2c4b4": "calculateRoot(bytes)",
		"925ea687": "castToSnapshot(bytes)",
		"a641fa33": "formatSnapshot(bytes[])",
		"aa1e84de": "hash(bytes)",
		"8aae3c34": "isSnapshot(bytes)",
		"1406cde1": "state(bytes,uint256)",
		"450701c5": "statesAmount(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b506110bb806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063925ea6871161005b578063925ea687146100ef578063a641fa3314610102578063aa1e84de14610115578063c9b2c4b41461012857600080fd5b80631406cde114610082578063450701c5146100ab5780638aae3c34146100cc575b600080fd5b610095610090366004610dce565b61013b565b6040516100a29190610e13565b60405180910390f35b6100be6100b9366004610e7f565b610167565b6040519081526020016100a2565b6100df6100da366004610e7f565b61017a565b60405190151581526020016100a2565b6100956100fd366004610e7f565b61018d565b610095610110366004610eb4565b6101ac565b6100be610123366004610e7f565b610264565b6100be610136366004610e7f565b610277565b606061015e610159610156846101508761028a565b9061029d565b90565b61034b565b90505b92915050565b60006101616101758361028a565b6103a8565b6000610161610188836103c7565b6103da565b6060600061019a8361028a565b90506101a58161034b565b9392505050565b805160609060008167ffffffffffffffff8111156101cc576101cc610cc2565b6040519080825280602002602001820160405280156101f5578160200160208202803683370190505b50905060005b828110156102525761022585828151811061021857610218610f77565b602002602001015161041a565b82828151811061023757610237610f77565b602090810291909101015261024b81610fd5565b90506101fb565b5061025c81610428565b949350505050565b60006101616102728361028a565b610545565b60006101616102858361028a565b610571565b6000610161610298836103c7565b61064a565b600082816102ac60328561100d565b90506fffffffffffffffffffffffffffffffff8216811061032e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f537461746520696e646578206f7574206f662072616e6765000000000000000060448201526064015b60405180910390fd5b61034261033d838360326106bf565b610730565b95945050505050565b6040518061035c83602083016107ae565b506fffffffffffffffffffffffffffffffff83166000601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168301602001604052509052919050565b600061016160326fffffffffffffffffffffffffffffffff8416611024565b80516000906020830161025c818361085d565b60006fffffffffffffffffffffffffffffffff8216816103fb603283611024565b90508161040960328361100d565b14801561025c575061025c816108c0565b600061016161033d836103c7565b606061043482516108c0565b61049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601560248201527f496e76616c69642073746174657320616d6f756e7400000000000000000000006044820152606401610325565b815160008167ffffffffffffffff8111156104b7576104b7610cc2565b6040519080825280602002602001820160405280156104e0578160200160208202803683370190505b50905060005b8281101561053b5761050e85828151811061050357610503610f77565b602002602001015190565b82828151811061052057610520610f77565b602090810291909101015261053481610fd5565b90506104e6565b5061025c816108e5565b6000610161827fdfe02260445526f7b137cb9caf995dcdead56fff547ac8de4b3e330521723148610941565b60008061057d836103a8565b905060008167ffffffffffffffff81111561059a5761059a610cc2565b6040519080825280602002602001820160405280156105c3578160200160208202803683370190505b50905060005b82811015610610576105e36105de868361029d565b61097d565b8282815181106105f5576105f5610f77565b602090810291909101015261060981610fd5565b90506105c9565b50610626816106216001600661105f565b6109bc565b8060008151811061063957610639610f77565b602002602001015192505050919050565b6000610655826103da565b6106bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f4e6f74206120736e617073686f740000000000000000000000000000000000006044820152606401610325565b5090565b6000806106cc8560801c90565b90506106d785610adf565b836106e28684611072565b6106ec9190611072565b1115610724576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6103428482018461085d565b600060326fffffffffffffffffffffffffffffffff8316146106bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f7420612073746174650000000000000000000000000000000000000000006044820152606401610325565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c9080851015610808576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa90508061084b576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417979650505050505050565b60008061086a8385611072565b905060405181111561087a575060005b806000036108b4576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b831761025c565b6000811580159061016157506108d86001600661105f565b6001901b82111592915050565b6040518060006108f88460208401610b05565b6fffffffffffffffffffffffffffffffff16601f81017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016830160200160405290915250919050565b60008161094d84610ba9565b60408051602081019390935282015260600160405160208183030381529060405280519060200120905092915050565b600080600061098b84610bd4565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b81516001821b811115610a2b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600e60248201527f48656967687420746f6f206c6f770000000000000000000000000000000000006044820152606401610325565b60005b82811015610ad95760005b82811015610aca5760008160010190506000868381518110610a5d57610a5d610f77565b602002602001015190506000858310610a77576000610a92565b878381518110610a8957610a89610f77565b60200260200101515b9050610a9e8282610c03565b88600186901c81518110610ab457610ab4610f77565b6020908102919091010152505050600201610a39565b506001918201821c9101610a2e565b50505050565b60006fffffffffffffffffffffffffffffffff8216610afe8360801c90565b0192915050565b60405160009080831015610b45576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b8551811015610b9c576000868281518110610b6657610b66610f77565b60200260200101519050610b7c818488016107ae565b506fffffffffffffffffffffffffffffffff169190910190600101610b49565b50608084901b8117610342565b600080610bb68360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b60008082610beb610be6826024610c4f565b610ba9565b9250610bfb610be6826024610c5c565b915050915091565b600082158015610c11575081155b15610c1e57506000610161565b6040805160208101859052908101839052606001604051602081830303815290604052805190602001209050610161565b600061015e8382846106bf565b60006fffffffffffffffffffffffffffffffff831680831115610cab576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61025c83610cb98660801c90565b0184830361085d565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715610d3857610d38610cc2565b604052919050565b600082601f830112610d5157600080fd5b813567ffffffffffffffff811115610d6b57610d6b610cc2565b610d9c60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601610cf1565b818152846020838601011115610db157600080fd5b816020850160208301376000918101602001919091529392505050565b60008060408385031215610de157600080fd5b823567ffffffffffffffff811115610df857600080fd5b610e0485828601610d40565b95602094909401359450505050565b600060208083528351808285015260005b81811015610e4057858101830151858201604001528201610e24565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b600060208284031215610e9157600080fd5b813567ffffffffffffffff811115610ea857600080fd5b61025c84828501610d40565b60006020808385031215610ec757600080fd5b823567ffffffffffffffff80821115610edf57600080fd5b818501915085601f830112610ef357600080fd5b813581811115610f0557610f05610cc2565b8060051b610f14858201610cf1565b9182528381018501918581019089841115610f2e57600080fd5b86860192505b83831015610f6a57823585811115610f4c5760008081fd5b610f5a8b89838a0101610d40565b8352509186019190860190610f34565b9998505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361100657611006610fa6565b5060010190565b808202811582820484141761016157610161610fa6565b60008261105a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b8181038181111561016157610161610fa6565b8082018082111561016157610161610fa656fea2646970667358221220f27c4181a1c31c7dfbec0128aeca21e81150c5ea3bb8e1add91f04cf57eedcc364736f6c63430008110033",
}

// SnapshotHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotHarnessMetaData.ABI instead.
var SnapshotHarnessABI = SnapshotHarnessMetaData.ABI

// Deprecated: Use SnapshotHarnessMetaData.Sigs instead.
// SnapshotHarnessFuncSigs maps the 4-byte function signature to its string representation.
var SnapshotHarnessFuncSigs = SnapshotHarnessMetaData.Sigs

// SnapshotHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SnapshotHarnessMetaData.Bin instead.
var SnapshotHarnessBin = SnapshotHarnessMetaData.Bin

// DeploySnapshotHarness deploys a new Ethereum contract, binding an instance of SnapshotHarness to it.
func DeploySnapshotHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SnapshotHarness, error) {
	parsed, err := SnapshotHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SnapshotHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SnapshotHarness{SnapshotHarnessCaller: SnapshotHarnessCaller{contract: contract}, SnapshotHarnessTransactor: SnapshotHarnessTransactor{contract: contract}, SnapshotHarnessFilterer: SnapshotHarnessFilterer{contract: contract}}, nil
}

// SnapshotHarness is an auto generated Go binding around an Ethereum contract.
type SnapshotHarness struct {
	SnapshotHarnessCaller     // Read-only binding to the contract
	SnapshotHarnessTransactor // Write-only binding to the contract
	SnapshotHarnessFilterer   // Log filterer for contract events
}

// SnapshotHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotHarnessSession struct {
	Contract     *SnapshotHarness  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotHarnessCallerSession struct {
	Contract *SnapshotHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SnapshotHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotHarnessTransactorSession struct {
	Contract     *SnapshotHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SnapshotHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotHarnessRaw struct {
	Contract *SnapshotHarness // Generic contract binding to access the raw methods on
}

// SnapshotHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotHarnessCallerRaw struct {
	Contract *SnapshotHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotHarnessTransactorRaw struct {
	Contract *SnapshotHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshotHarness creates a new instance of SnapshotHarness, bound to a specific deployed contract.
func NewSnapshotHarness(address common.Address, backend bind.ContractBackend) (*SnapshotHarness, error) {
	contract, err := bindSnapshotHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SnapshotHarness{SnapshotHarnessCaller: SnapshotHarnessCaller{contract: contract}, SnapshotHarnessTransactor: SnapshotHarnessTransactor{contract: contract}, SnapshotHarnessFilterer: SnapshotHarnessFilterer{contract: contract}}, nil
}

// NewSnapshotHarnessCaller creates a new read-only instance of SnapshotHarness, bound to a specific deployed contract.
func NewSnapshotHarnessCaller(address common.Address, caller bind.ContractCaller) (*SnapshotHarnessCaller, error) {
	contract, err := bindSnapshotHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotHarnessCaller{contract: contract}, nil
}

// NewSnapshotHarnessTransactor creates a new write-only instance of SnapshotHarness, bound to a specific deployed contract.
func NewSnapshotHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotHarnessTransactor, error) {
	contract, err := bindSnapshotHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotHarnessTransactor{contract: contract}, nil
}

// NewSnapshotHarnessFilterer creates a new log filterer instance of SnapshotHarness, bound to a specific deployed contract.
func NewSnapshotHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotHarnessFilterer, error) {
	contract, err := bindSnapshotHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotHarnessFilterer{contract: contract}, nil
}

// bindSnapshotHarness binds a generic wrapper to an already deployed contract.
func bindSnapshotHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotHarness *SnapshotHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotHarness.Contract.SnapshotHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotHarness *SnapshotHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHarness.Contract.SnapshotHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotHarness *SnapshotHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotHarness.Contract.SnapshotHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotHarness *SnapshotHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotHarness *SnapshotHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotHarness *SnapshotHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotHarness.Contract.contract.Transact(opts, method, params...)
}

// CalculateRoot is a free data retrieval call binding the contract method 0xc9b2c4b4.
//
// Solidity: function calculateRoot(bytes payload) pure returns(bytes32)
func (_SnapshotHarness *SnapshotHarnessCaller) CalculateRoot(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _SnapshotHarness.contract.Call(opts, &out, "calculateRoot", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CalculateRoot is a free data retrieval call binding the contract method 0xc9b2c4b4.
//
// Solidity: function calculateRoot(bytes payload) pure returns(bytes32)
func (_SnapshotHarness *SnapshotHarnessSession) CalculateRoot(payload []byte) ([32]byte, error) {
	return _SnapshotHarness.Contract.CalculateRoot(&_SnapshotHarness.CallOpts, payload)
}

// CalculateRoot is a free data retrieval call binding the contract method 0xc9b2c4b4.
//
// Solidity: function calculateRoot(bytes payload) pure returns(bytes32)
func (_SnapshotHarness *SnapshotHarnessCallerSession) CalculateRoot(payload []byte) ([32]byte, error) {
	return _SnapshotHarness.Contract.CalculateRoot(&_SnapshotHarness.CallOpts, payload)
}

// CastToSnapshot is a free data retrieval call binding the contract method 0x925ea687.
//
// Solidity: function castToSnapshot(bytes payload) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessCaller) CastToSnapshot(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHarness.contract.Call(opts, &out, "castToSnapshot", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CastToSnapshot is a free data retrieval call binding the contract method 0x925ea687.
//
// Solidity: function castToSnapshot(bytes payload) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessSession) CastToSnapshot(payload []byte) ([]byte, error) {
	return _SnapshotHarness.Contract.CastToSnapshot(&_SnapshotHarness.CallOpts, payload)
}

// CastToSnapshot is a free data retrieval call binding the contract method 0x925ea687.
//
// Solidity: function castToSnapshot(bytes payload) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessCallerSession) CastToSnapshot(payload []byte) ([]byte, error) {
	return _SnapshotHarness.Contract.CastToSnapshot(&_SnapshotHarness.CallOpts, payload)
}

// FormatSnapshot is a free data retrieval call binding the contract method 0xa641fa33.
//
// Solidity: function formatSnapshot(bytes[] statePayloads) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessCaller) FormatSnapshot(opts *bind.CallOpts, statePayloads [][]byte) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHarness.contract.Call(opts, &out, "formatSnapshot", statePayloads)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatSnapshot is a free data retrieval call binding the contract method 0xa641fa33.
//
// Solidity: function formatSnapshot(bytes[] statePayloads) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessSession) FormatSnapshot(statePayloads [][]byte) ([]byte, error) {
	return _SnapshotHarness.Contract.FormatSnapshot(&_SnapshotHarness.CallOpts, statePayloads)
}

// FormatSnapshot is a free data retrieval call binding the contract method 0xa641fa33.
//
// Solidity: function formatSnapshot(bytes[] statePayloads) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessCallerSession) FormatSnapshot(statePayloads [][]byte) ([]byte, error) {
	return _SnapshotHarness.Contract.FormatSnapshot(&_SnapshotHarness.CallOpts, statePayloads)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes payload) pure returns(bytes32)
func (_SnapshotHarness *SnapshotHarnessCaller) Hash(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _SnapshotHarness.contract.Call(opts, &out, "hash", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes payload) pure returns(bytes32)
func (_SnapshotHarness *SnapshotHarnessSession) Hash(payload []byte) ([32]byte, error) {
	return _SnapshotHarness.Contract.Hash(&_SnapshotHarness.CallOpts, payload)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes payload) pure returns(bytes32)
func (_SnapshotHarness *SnapshotHarnessCallerSession) Hash(payload []byte) ([32]byte, error) {
	return _SnapshotHarness.Contract.Hash(&_SnapshotHarness.CallOpts, payload)
}

// IsSnapshot is a free data retrieval call binding the contract method 0x8aae3c34.
//
// Solidity: function isSnapshot(bytes payload) pure returns(bool)
func (_SnapshotHarness *SnapshotHarnessCaller) IsSnapshot(opts *bind.CallOpts, payload []byte) (bool, error) {
	var out []interface{}
	err := _SnapshotHarness.contract.Call(opts, &out, "isSnapshot", payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSnapshot is a free data retrieval call binding the contract method 0x8aae3c34.
//
// Solidity: function isSnapshot(bytes payload) pure returns(bool)
func (_SnapshotHarness *SnapshotHarnessSession) IsSnapshot(payload []byte) (bool, error) {
	return _SnapshotHarness.Contract.IsSnapshot(&_SnapshotHarness.CallOpts, payload)
}

// IsSnapshot is a free data retrieval call binding the contract method 0x8aae3c34.
//
// Solidity: function isSnapshot(bytes payload) pure returns(bool)
func (_SnapshotHarness *SnapshotHarnessCallerSession) IsSnapshot(payload []byte) (bool, error) {
	return _SnapshotHarness.Contract.IsSnapshot(&_SnapshotHarness.CallOpts, payload)
}

// State is a free data retrieval call binding the contract method 0x1406cde1.
//
// Solidity: function state(bytes payload, uint256 stateIndex) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessCaller) State(opts *bind.CallOpts, payload []byte, stateIndex *big.Int) ([]byte, error) {
	var out []interface{}
	err := _SnapshotHarness.contract.Call(opts, &out, "state", payload, stateIndex)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x1406cde1.
//
// Solidity: function state(bytes payload, uint256 stateIndex) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessSession) State(payload []byte, stateIndex *big.Int) ([]byte, error) {
	return _SnapshotHarness.Contract.State(&_SnapshotHarness.CallOpts, payload, stateIndex)
}

// State is a free data retrieval call binding the contract method 0x1406cde1.
//
// Solidity: function state(bytes payload, uint256 stateIndex) view returns(bytes)
func (_SnapshotHarness *SnapshotHarnessCallerSession) State(payload []byte, stateIndex *big.Int) ([]byte, error) {
	return _SnapshotHarness.Contract.State(&_SnapshotHarness.CallOpts, payload, stateIndex)
}

// StatesAmount is a free data retrieval call binding the contract method 0x450701c5.
//
// Solidity: function statesAmount(bytes payload) pure returns(uint256)
func (_SnapshotHarness *SnapshotHarnessCaller) StatesAmount(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _SnapshotHarness.contract.Call(opts, &out, "statesAmount", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StatesAmount is a free data retrieval call binding the contract method 0x450701c5.
//
// Solidity: function statesAmount(bytes payload) pure returns(uint256)
func (_SnapshotHarness *SnapshotHarnessSession) StatesAmount(payload []byte) (*big.Int, error) {
	return _SnapshotHarness.Contract.StatesAmount(&_SnapshotHarness.CallOpts, payload)
}

// StatesAmount is a free data retrieval call binding the contract method 0x450701c5.
//
// Solidity: function statesAmount(bytes payload) pure returns(uint256)
func (_SnapshotHarness *SnapshotHarnessCallerSession) StatesAmount(payload []byte) (*big.Int, error) {
	return _SnapshotHarness.Contract.StatesAmount(&_SnapshotHarness.CallOpts, payload)
}

// SnapshotLibMetaData contains all meta data concerning the SnapshotLib contract.
var SnapshotLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a508bb05ac779e71dcbe418917d8f5def57f44508e5f5a5b31df72fae75a282064736f6c63430008110033",
}

// SnapshotLibABI is the input ABI used to generate the binding from.
// Deprecated: Use SnapshotLibMetaData.ABI instead.
var SnapshotLibABI = SnapshotLibMetaData.ABI

// SnapshotLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SnapshotLibMetaData.Bin instead.
var SnapshotLibBin = SnapshotLibMetaData.Bin

// DeploySnapshotLib deploys a new Ethereum contract, binding an instance of SnapshotLib to it.
func DeploySnapshotLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SnapshotLib, error) {
	parsed, err := SnapshotLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SnapshotLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SnapshotLib{SnapshotLibCaller: SnapshotLibCaller{contract: contract}, SnapshotLibTransactor: SnapshotLibTransactor{contract: contract}, SnapshotLibFilterer: SnapshotLibFilterer{contract: contract}}, nil
}

// SnapshotLib is an auto generated Go binding around an Ethereum contract.
type SnapshotLib struct {
	SnapshotLibCaller     // Read-only binding to the contract
	SnapshotLibTransactor // Write-only binding to the contract
	SnapshotLibFilterer   // Log filterer for contract events
}

// SnapshotLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnapshotLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnapshotLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnapshotLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnapshotLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnapshotLibSession struct {
	Contract     *SnapshotLib      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnapshotLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnapshotLibCallerSession struct {
	Contract *SnapshotLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SnapshotLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnapshotLibTransactorSession struct {
	Contract     *SnapshotLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SnapshotLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnapshotLibRaw struct {
	Contract *SnapshotLib // Generic contract binding to access the raw methods on
}

// SnapshotLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnapshotLibCallerRaw struct {
	Contract *SnapshotLibCaller // Generic read-only contract binding to access the raw methods on
}

// SnapshotLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnapshotLibTransactorRaw struct {
	Contract *SnapshotLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSnapshotLib creates a new instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLib(address common.Address, backend bind.ContractBackend) (*SnapshotLib, error) {
	contract, err := bindSnapshotLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SnapshotLib{SnapshotLibCaller: SnapshotLibCaller{contract: contract}, SnapshotLibTransactor: SnapshotLibTransactor{contract: contract}, SnapshotLibFilterer: SnapshotLibFilterer{contract: contract}}, nil
}

// NewSnapshotLibCaller creates a new read-only instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibCaller(address common.Address, caller bind.ContractCaller) (*SnapshotLibCaller, error) {
	contract, err := bindSnapshotLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibCaller{contract: contract}, nil
}

// NewSnapshotLibTransactor creates a new write-only instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibTransactor(address common.Address, transactor bind.ContractTransactor) (*SnapshotLibTransactor, error) {
	contract, err := bindSnapshotLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibTransactor{contract: contract}, nil
}

// NewSnapshotLibFilterer creates a new log filterer instance of SnapshotLib, bound to a specific deployed contract.
func NewSnapshotLibFilterer(address common.Address, filterer bind.ContractFilterer) (*SnapshotLibFilterer, error) {
	contract, err := bindSnapshotLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnapshotLibFilterer{contract: contract}, nil
}

// bindSnapshotLib binds a generic wrapper to an already deployed contract.
func bindSnapshotLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnapshotLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotLib *SnapshotLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotLib.Contract.SnapshotLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotLib *SnapshotLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotLib.Contract.SnapshotLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotLib *SnapshotLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotLib.Contract.SnapshotLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SnapshotLib *SnapshotLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SnapshotLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SnapshotLib *SnapshotLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SnapshotLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SnapshotLib *SnapshotLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SnapshotLib.Contract.contract.Transact(opts, method, params...)
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220df50f440737edf4061e7f8af2dfe7b54f7cdf99ab54b17398398786e4c16e0bd64736f6c63430008110033",
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
