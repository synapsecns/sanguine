// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package executionservice

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
	parsed, err := ContextMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// ExecutionServiceMetaData contains all meta data concerning the ExecutionService contract.
var ExecutionServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOAUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"interchainClient\",\"type\":\"address\"}],\"name\":\"InterchainClientUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gasOracle\",\"outputs\":[{\"internalType\":\"contractIGasOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"interchainClient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_executorEOA\",\"type\":\"address\"}],\"name\":\"setExecutorEOA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gasOracle\",\"type\":\"address\"}],\"name\":\"setGasOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_interchainClient\",\"type\":\"address\"}],\"name\":\"setInterchainClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"5d62a8dd": "gasOracle()",
		"c473e7e8": "getExecutionFee(uint256,uint256,bytes)",
		"c9a64e91": "interchainClient()",
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"e4e06522": "requestExecution(uint256,uint256,bytes32,uint256,bytes)",
		"2d54566c": "setExecutorEOA(address)",
		"a87b8152": "setGasOracle(address)",
		"27efcbb7": "setInterchainClient(address)",
		"f2fde38b": "transferOwnership(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051610d9c380380610d9c83398101604081905261002f916100be565b806001600160a01b03811661005e57604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b6100678161006e565b50506100ee565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100d057600080fd5b81516001600160a01b03811681146100e757600080fd5b9392505050565b610c9f806100fd6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80638da5cb5b11610081578063c9a64e911161005b578063c9a64e91146101ba578063e4e06522146101da578063f2fde38b146101ed57600080fd5b80638da5cb5b14610168578063a87b815214610186578063c473e7e81461019957600080fd5b80635d62a8dd116100b25780635d62a8dd146100f657806362014bad14610140578063715018a61461016057600080fd5b806327efcbb7146100ce5780632d54566c146100e3575b600080fd5b6100e16100dc366004610913565b610200565b005b6100e16100f1366004610913565b610277565b6003546101169073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b6002546101169073ffffffffffffffffffffffffffffffffffffffff1681565b6100e16102ee565b60005473ffffffffffffffffffffffffffffffffffffffff16610116565b6100e1610194366004610913565b610302565b6101ac6101a7366004610a63565b610379565b604051908152602001610137565b6001546101169073ffffffffffffffffffffffffffffffffffffffff1681565b6100e16101e8366004610ab3565b6105c4565b6100e16101fb366004610913565b610741565b6102086107a5565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f33e3a51bf9fedee6e205cd10237a9a8dd3fffed45ac3dee88d2eba92ef8b5d6290600090a250565b61027f6107a5565b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f8a101f72b53416c657555bd234f732510711b6f59f28b161db43892c89b6b3b490600090a250565b6102f66107a5565b61030060006107f8565b565b61030a6107a5565b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517f0b987d971d57fe66ec2d0cff095b1b547c3076cdd07fdf5b0863f320700dbb9090600090a250565b60008060006103878461086d565b90925090507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60ff8316016105305760006103c18561088f565b60035481516040517f5cbd3c48000000000000000000000000000000000000000000000000000000008152600481018b905260248101919091526044810189905291925060009173ffffffffffffffffffffffffffffffffffffffff90911690635cbd3c4890606401602060405180830381865afa158015610447573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061046b9190610b17565b6020830151909150156105255760035460208301516040517f1e7b9287000000000000000000000000000000000000000000000000000000008152600481018b9052602481019190915273ffffffffffffffffffffffffffffffffffffffff90911690631e7b928790604401602060405180830381865afa1580156104f4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105189190610b17565b6105229082610b30565b90505b93506105bd92505050565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f556e737570706f72746564206f7074696f6e732076657273696f6e3a2076657260448201527f73696f6e206d757374206265204f5054494f4e535f563100000000000000000060648201526084015b60405180910390fd5b9392505050565b60015473ffffffffffffffffffffffffffffffffffffffff16331461066b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f457865637574696f6e536572766963653a2063616c6c6572206973206e6f742060448201527f74686520496e746572636861696e436c69656e7400000000000000000000000060648201526084016105b4565b610676858583610379565b821015610705576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f457865637574696f6e536572766963653a20657865637574696f6e206665652060448201527f6973206e6f74206869676820656e6f756768000000000000000000000000000060648201526084016105b4565b60405133815283907f507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd13829060200160405180910390a25050505050565b6107496107a5565b73ffffffffffffffffffffffffffffffffffffffff8116610799576040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600060048201526024016105b4565b6107a2816107f8565b50565b60005473ffffffffffffffffffffffffffffffffffffffff163314610300576040517f118cdaa70000000000000000000000000000000000000000000000000000000081523360048201526024016105b4565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60006060828060200190518101906108859190610b70565b9094909350915050565b60408051808201909152600080825260208201526000806108af8461086d565b9092509050600160ff831610156108f7576040517fbd91a21500000000000000000000000000000000000000000000000000000000815260ff831660048201526024016105b4565b8080602001905181019061090b9190610c1a565b949350505050565b60006020828403121561092557600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146105bd57600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156109bf576109bf610949565b604052919050565b600067ffffffffffffffff8211156109e1576109e1610949565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600082601f830112610a1e57600080fd5b8135610a31610a2c826109c7565b610978565b818152846020838601011115610a4657600080fd5b816020850160208301376000918101602001919091529392505050565b600080600060608486031215610a7857600080fd5b8335925060208401359150604084013567ffffffffffffffff811115610a9d57600080fd5b610aa986828701610a0d565b9150509250925092565b600080600080600060a08688031215610acb57600080fd5b85359450602086013593506040860135925060608601359150608086013567ffffffffffffffff811115610afe57600080fd5b610b0a88828901610a0d565b9150509295509295909350565b600060208284031215610b2957600080fd5b5051919050565b80820180821115610b6a577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b92915050565b60008060408385031215610b8357600080fd5b825160ff81168114610b9457600080fd5b8092505060208084015167ffffffffffffffff811115610bb357600080fd5b8401601f81018613610bc457600080fd5b8051610bd2610a2c826109c7565b8181528784838501011115610be657600080fd5b60005b82811015610c04578381018501518282018601528401610be9565b5060009181019093015250919491935090915050565b600060408284031215610c2c57600080fd5b6040516040810181811067ffffffffffffffff82111715610c4f57610c4f610949565b60405282518152602092830151928101929092525091905056fea26469706673582212200641b0b7a9fd32751747fc147907f58c5f85135d24f10373e9d8067e5651333064736f6c63430008140033",
}

// ExecutionServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionServiceMetaData.ABI instead.
var ExecutionServiceABI = ExecutionServiceMetaData.ABI

// Deprecated: Use ExecutionServiceMetaData.Sigs instead.
// ExecutionServiceFuncSigs maps the 4-byte function signature to its string representation.
var ExecutionServiceFuncSigs = ExecutionServiceMetaData.Sigs

// ExecutionServiceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExecutionServiceMetaData.Bin instead.
var ExecutionServiceBin = ExecutionServiceMetaData.Bin

// DeployExecutionService deploys a new Ethereum contract, binding an instance of ExecutionService to it.
func DeployExecutionService(auth *bind.TransactOpts, backend bind.ContractBackend, owner_ common.Address) (common.Address, *types.Transaction, *ExecutionService, error) {
	parsed, err := ExecutionServiceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExecutionServiceBin), backend, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExecutionService{ExecutionServiceCaller: ExecutionServiceCaller{contract: contract}, ExecutionServiceTransactor: ExecutionServiceTransactor{contract: contract}, ExecutionServiceFilterer: ExecutionServiceFilterer{contract: contract}}, nil
}

// ExecutionService is an auto generated Go binding around an Ethereum contract.
type ExecutionService struct {
	ExecutionServiceCaller     // Read-only binding to the contract
	ExecutionServiceTransactor // Write-only binding to the contract
	ExecutionServiceFilterer   // Log filterer for contract events
}

// ExecutionServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionServiceSession struct {
	Contract     *ExecutionService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutionServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionServiceCallerSession struct {
	Contract *ExecutionServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ExecutionServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionServiceTransactorSession struct {
	Contract     *ExecutionServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ExecutionServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionServiceRaw struct {
	Contract *ExecutionService // Generic contract binding to access the raw methods on
}

// ExecutionServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionServiceCallerRaw struct {
	Contract *ExecutionServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionServiceTransactorRaw struct {
	Contract *ExecutionServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionService creates a new instance of ExecutionService, bound to a specific deployed contract.
func NewExecutionService(address common.Address, backend bind.ContractBackend) (*ExecutionService, error) {
	contract, err := bindExecutionService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionService{ExecutionServiceCaller: ExecutionServiceCaller{contract: contract}, ExecutionServiceTransactor: ExecutionServiceTransactor{contract: contract}, ExecutionServiceFilterer: ExecutionServiceFilterer{contract: contract}}, nil
}

// NewExecutionServiceCaller creates a new read-only instance of ExecutionService, bound to a specific deployed contract.
func NewExecutionServiceCaller(address common.Address, caller bind.ContractCaller) (*ExecutionServiceCaller, error) {
	contract, err := bindExecutionService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceCaller{contract: contract}, nil
}

// NewExecutionServiceTransactor creates a new write-only instance of ExecutionService, bound to a specific deployed contract.
func NewExecutionServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionServiceTransactor, error) {
	contract, err := bindExecutionService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceTransactor{contract: contract}, nil
}

// NewExecutionServiceFilterer creates a new log filterer instance of ExecutionService, bound to a specific deployed contract.
func NewExecutionServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionServiceFilterer, error) {
	contract, err := bindExecutionService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceFilterer{contract: contract}, nil
}

// bindExecutionService binds a generic wrapper to an already deployed contract.
func bindExecutionService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutionServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionService *ExecutionServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionService.Contract.ExecutionServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionService *ExecutionServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionService.Contract.ExecutionServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionService *ExecutionServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionService.Contract.ExecutionServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionService *ExecutionServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionService *ExecutionServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionService *ExecutionServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionService.Contract.contract.Transact(opts, method, params...)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ExecutionService *ExecutionServiceCaller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionService.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ExecutionService *ExecutionServiceSession) ExecutorEOA() (common.Address, error) {
	return _ExecutionService.Contract.ExecutorEOA(&_ExecutionService.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_ExecutionService *ExecutionServiceCallerSession) ExecutorEOA() (common.Address, error) {
	return _ExecutionService.Contract.ExecutorEOA(&_ExecutionService.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ExecutionService *ExecutionServiceCaller) GasOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionService.contract.Call(opts, &out, "gasOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ExecutionService *ExecutionServiceSession) GasOracle() (common.Address, error) {
	return _ExecutionService.Contract.GasOracle(&_ExecutionService.CallOpts)
}

// GasOracle is a free data retrieval call binding the contract method 0x5d62a8dd.
//
// Solidity: function gasOracle() view returns(address)
func (_ExecutionService *ExecutionServiceCallerSession) GasOracle() (common.Address, error) {
	return _ExecutionService.Contract.GasOracle(&_ExecutionService.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ExecutionService *ExecutionServiceCaller) GetExecutionFee(opts *bind.CallOpts, dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _ExecutionService.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ExecutionService *ExecutionServiceSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _ExecutionService.Contract.GetExecutionFee(&_ExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_ExecutionService *ExecutionServiceCallerSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _ExecutionService.Contract.GetExecutionFee(&_ExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// InterchainClient is a free data retrieval call binding the contract method 0xc9a64e91.
//
// Solidity: function interchainClient() view returns(address)
func (_ExecutionService *ExecutionServiceCaller) InterchainClient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionService.contract.Call(opts, &out, "interchainClient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterchainClient is a free data retrieval call binding the contract method 0xc9a64e91.
//
// Solidity: function interchainClient() view returns(address)
func (_ExecutionService *ExecutionServiceSession) InterchainClient() (common.Address, error) {
	return _ExecutionService.Contract.InterchainClient(&_ExecutionService.CallOpts)
}

// InterchainClient is a free data retrieval call binding the contract method 0xc9a64e91.
//
// Solidity: function interchainClient() view returns(address)
func (_ExecutionService *ExecutionServiceCallerSession) InterchainClient() (common.Address, error) {
	return _ExecutionService.Contract.InterchainClient(&_ExecutionService.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionService *ExecutionServiceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExecutionService.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionService *ExecutionServiceSession) Owner() (common.Address, error) {
	return _ExecutionService.Contract.Owner(&_ExecutionService.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExecutionService *ExecutionServiceCallerSession) Owner() (common.Address, error) {
	return _ExecutionService.Contract.Owner(&_ExecutionService.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionService *ExecutionServiceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionService.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionService *ExecutionServiceSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExecutionService.Contract.RenounceOwnership(&_ExecutionService.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExecutionService *ExecutionServiceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExecutionService.Contract.RenounceOwnership(&_ExecutionService.TransactOpts)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ExecutionService *ExecutionServiceTransactor) RequestExecution(opts *bind.TransactOpts, dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ExecutionService.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ExecutionService *ExecutionServiceSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ExecutionService.Contract.RequestExecution(&_ExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_ExecutionService *ExecutionServiceTransactorSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _ExecutionService.Contract.RequestExecution(&_ExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address _executorEOA) returns()
func (_ExecutionService *ExecutionServiceTransactor) SetExecutorEOA(opts *bind.TransactOpts, _executorEOA common.Address) (*types.Transaction, error) {
	return _ExecutionService.contract.Transact(opts, "setExecutorEOA", _executorEOA)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address _executorEOA) returns()
func (_ExecutionService *ExecutionServiceSession) SetExecutorEOA(_executorEOA common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.SetExecutorEOA(&_ExecutionService.TransactOpts, _executorEOA)
}

// SetExecutorEOA is a paid mutator transaction binding the contract method 0x2d54566c.
//
// Solidity: function setExecutorEOA(address _executorEOA) returns()
func (_ExecutionService *ExecutionServiceTransactorSession) SetExecutorEOA(_executorEOA common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.SetExecutorEOA(&_ExecutionService.TransactOpts, _executorEOA)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address _gasOracle) returns()
func (_ExecutionService *ExecutionServiceTransactor) SetGasOracle(opts *bind.TransactOpts, _gasOracle common.Address) (*types.Transaction, error) {
	return _ExecutionService.contract.Transact(opts, "setGasOracle", _gasOracle)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address _gasOracle) returns()
func (_ExecutionService *ExecutionServiceSession) SetGasOracle(_gasOracle common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.SetGasOracle(&_ExecutionService.TransactOpts, _gasOracle)
}

// SetGasOracle is a paid mutator transaction binding the contract method 0xa87b8152.
//
// Solidity: function setGasOracle(address _gasOracle) returns()
func (_ExecutionService *ExecutionServiceTransactorSession) SetGasOracle(_gasOracle common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.SetGasOracle(&_ExecutionService.TransactOpts, _gasOracle)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address _interchainClient) returns()
func (_ExecutionService *ExecutionServiceTransactor) SetInterchainClient(opts *bind.TransactOpts, _interchainClient common.Address) (*types.Transaction, error) {
	return _ExecutionService.contract.Transact(opts, "setInterchainClient", _interchainClient)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address _interchainClient) returns()
func (_ExecutionService *ExecutionServiceSession) SetInterchainClient(_interchainClient common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.SetInterchainClient(&_ExecutionService.TransactOpts, _interchainClient)
}

// SetInterchainClient is a paid mutator transaction binding the contract method 0x27efcbb7.
//
// Solidity: function setInterchainClient(address _interchainClient) returns()
func (_ExecutionService *ExecutionServiceTransactorSession) SetInterchainClient(_interchainClient common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.SetInterchainClient(&_ExecutionService.TransactOpts, _interchainClient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionService *ExecutionServiceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionService.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionService *ExecutionServiceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.TransferOwnership(&_ExecutionService.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExecutionService *ExecutionServiceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExecutionService.Contract.TransferOwnership(&_ExecutionService.TransactOpts, newOwner)
}

// ExecutionServiceExecutionRequestedIterator is returned from FilterExecutionRequested and is used to iterate over the raw logs and unpacked data for ExecutionRequested events raised by the ExecutionService contract.
type ExecutionServiceExecutionRequestedIterator struct {
	Event *ExecutionServiceExecutionRequested // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceExecutionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceExecutionRequested)
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
		it.Event = new(ExecutionServiceExecutionRequested)
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
func (it *ExecutionServiceExecutionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceExecutionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceExecutionRequested represents a ExecutionRequested event raised by the ExecutionService contract.
type ExecutionServiceExecutionRequested struct {
	TransactionId [32]byte
	Client        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_ExecutionService *ExecutionServiceFilterer) FilterExecutionRequested(opts *bind.FilterOpts, transactionId [][32]byte) (*ExecutionServiceExecutionRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ExecutionService.contract.FilterLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceExecutionRequestedIterator{contract: _ExecutionService.contract, event: "ExecutionRequested", logs: logs, sub: sub}, nil
}

// WatchExecutionRequested is a free log subscription operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_ExecutionService *ExecutionServiceFilterer) WatchExecutionRequested(opts *bind.WatchOpts, sink chan<- *ExecutionServiceExecutionRequested, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ExecutionService.contract.WatchLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceExecutionRequested)
				if err := _ExecutionService.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
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

// ParseExecutionRequested is a log parse operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_ExecutionService *ExecutionServiceFilterer) ParseExecutionRequested(log types.Log) (*ExecutionServiceExecutionRequested, error) {
	event := new(ExecutionServiceExecutionRequested)
	if err := _ExecutionService.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceExecutorEOAUpdatedIterator is returned from FilterExecutorEOAUpdated and is used to iterate over the raw logs and unpacked data for ExecutorEOAUpdated events raised by the ExecutionService contract.
type ExecutionServiceExecutorEOAUpdatedIterator struct {
	Event *ExecutionServiceExecutorEOAUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceExecutorEOAUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceExecutorEOAUpdated)
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
		it.Event = new(ExecutionServiceExecutorEOAUpdated)
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
func (it *ExecutionServiceExecutorEOAUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceExecutorEOAUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceExecutorEOAUpdated represents a ExecutorEOAUpdated event raised by the ExecutionService contract.
type ExecutionServiceExecutorEOAUpdated struct {
	ExecutorEOA common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutorEOAUpdated is a free log retrieval operation binding the contract event 0x8a101f72b53416c657555bd234f732510711b6f59f28b161db43892c89b6b3b4.
//
// Solidity: event ExecutorEOAUpdated(address indexed executorEOA)
func (_ExecutionService *ExecutionServiceFilterer) FilterExecutorEOAUpdated(opts *bind.FilterOpts, executorEOA []common.Address) (*ExecutionServiceExecutorEOAUpdatedIterator, error) {

	var executorEOARule []interface{}
	for _, executorEOAItem := range executorEOA {
		executorEOARule = append(executorEOARule, executorEOAItem)
	}

	logs, sub, err := _ExecutionService.contract.FilterLogs(opts, "ExecutorEOAUpdated", executorEOARule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceExecutorEOAUpdatedIterator{contract: _ExecutionService.contract, event: "ExecutorEOAUpdated", logs: logs, sub: sub}, nil
}

// WatchExecutorEOAUpdated is a free log subscription operation binding the contract event 0x8a101f72b53416c657555bd234f732510711b6f59f28b161db43892c89b6b3b4.
//
// Solidity: event ExecutorEOAUpdated(address indexed executorEOA)
func (_ExecutionService *ExecutionServiceFilterer) WatchExecutorEOAUpdated(opts *bind.WatchOpts, sink chan<- *ExecutionServiceExecutorEOAUpdated, executorEOA []common.Address) (event.Subscription, error) {

	var executorEOARule []interface{}
	for _, executorEOAItem := range executorEOA {
		executorEOARule = append(executorEOARule, executorEOAItem)
	}

	logs, sub, err := _ExecutionService.contract.WatchLogs(opts, "ExecutorEOAUpdated", executorEOARule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceExecutorEOAUpdated)
				if err := _ExecutionService.contract.UnpackLog(event, "ExecutorEOAUpdated", log); err != nil {
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

// ParseExecutorEOAUpdated is a log parse operation binding the contract event 0x8a101f72b53416c657555bd234f732510711b6f59f28b161db43892c89b6b3b4.
//
// Solidity: event ExecutorEOAUpdated(address indexed executorEOA)
func (_ExecutionService *ExecutionServiceFilterer) ParseExecutorEOAUpdated(log types.Log) (*ExecutionServiceExecutorEOAUpdated, error) {
	event := new(ExecutionServiceExecutorEOAUpdated)
	if err := _ExecutionService.contract.UnpackLog(event, "ExecutorEOAUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceGasOracleUpdatedIterator is returned from FilterGasOracleUpdated and is used to iterate over the raw logs and unpacked data for GasOracleUpdated events raised by the ExecutionService contract.
type ExecutionServiceGasOracleUpdatedIterator struct {
	Event *ExecutionServiceGasOracleUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceGasOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceGasOracleUpdated)
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
		it.Event = new(ExecutionServiceGasOracleUpdated)
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
func (it *ExecutionServiceGasOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceGasOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceGasOracleUpdated represents a GasOracleUpdated event raised by the ExecutionService contract.
type ExecutionServiceGasOracleUpdated struct {
	GasOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOracleUpdated is a free log retrieval operation binding the contract event 0x0b987d971d57fe66ec2d0cff095b1b547c3076cdd07fdf5b0863f320700dbb90.
//
// Solidity: event GasOracleUpdated(address indexed gasOracle)
func (_ExecutionService *ExecutionServiceFilterer) FilterGasOracleUpdated(opts *bind.FilterOpts, gasOracle []common.Address) (*ExecutionServiceGasOracleUpdatedIterator, error) {

	var gasOracleRule []interface{}
	for _, gasOracleItem := range gasOracle {
		gasOracleRule = append(gasOracleRule, gasOracleItem)
	}

	logs, sub, err := _ExecutionService.contract.FilterLogs(opts, "GasOracleUpdated", gasOracleRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceGasOracleUpdatedIterator{contract: _ExecutionService.contract, event: "GasOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchGasOracleUpdated is a free log subscription operation binding the contract event 0x0b987d971d57fe66ec2d0cff095b1b547c3076cdd07fdf5b0863f320700dbb90.
//
// Solidity: event GasOracleUpdated(address indexed gasOracle)
func (_ExecutionService *ExecutionServiceFilterer) WatchGasOracleUpdated(opts *bind.WatchOpts, sink chan<- *ExecutionServiceGasOracleUpdated, gasOracle []common.Address) (event.Subscription, error) {

	var gasOracleRule []interface{}
	for _, gasOracleItem := range gasOracle {
		gasOracleRule = append(gasOracleRule, gasOracleItem)
	}

	logs, sub, err := _ExecutionService.contract.WatchLogs(opts, "GasOracleUpdated", gasOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceGasOracleUpdated)
				if err := _ExecutionService.contract.UnpackLog(event, "GasOracleUpdated", log); err != nil {
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

// ParseGasOracleUpdated is a log parse operation binding the contract event 0x0b987d971d57fe66ec2d0cff095b1b547c3076cdd07fdf5b0863f320700dbb90.
//
// Solidity: event GasOracleUpdated(address indexed gasOracle)
func (_ExecutionService *ExecutionServiceFilterer) ParseGasOracleUpdated(log types.Log) (*ExecutionServiceGasOracleUpdated, error) {
	event := new(ExecutionServiceGasOracleUpdated)
	if err := _ExecutionService.contract.UnpackLog(event, "GasOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceInterchainClientUpdatedIterator is returned from FilterInterchainClientUpdated and is used to iterate over the raw logs and unpacked data for InterchainClientUpdated events raised by the ExecutionService contract.
type ExecutionServiceInterchainClientUpdatedIterator struct {
	Event *ExecutionServiceInterchainClientUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceInterchainClientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceInterchainClientUpdated)
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
		it.Event = new(ExecutionServiceInterchainClientUpdated)
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
func (it *ExecutionServiceInterchainClientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceInterchainClientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceInterchainClientUpdated represents a InterchainClientUpdated event raised by the ExecutionService contract.
type ExecutionServiceInterchainClientUpdated struct {
	InterchainClient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientUpdated is a free log retrieval operation binding the contract event 0x33e3a51bf9fedee6e205cd10237a9a8dd3fffed45ac3dee88d2eba92ef8b5d62.
//
// Solidity: event InterchainClientUpdated(address indexed interchainClient)
func (_ExecutionService *ExecutionServiceFilterer) FilterInterchainClientUpdated(opts *bind.FilterOpts, interchainClient []common.Address) (*ExecutionServiceInterchainClientUpdatedIterator, error) {

	var interchainClientRule []interface{}
	for _, interchainClientItem := range interchainClient {
		interchainClientRule = append(interchainClientRule, interchainClientItem)
	}

	logs, sub, err := _ExecutionService.contract.FilterLogs(opts, "InterchainClientUpdated", interchainClientRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceInterchainClientUpdatedIterator{contract: _ExecutionService.contract, event: "InterchainClientUpdated", logs: logs, sub: sub}, nil
}

// WatchInterchainClientUpdated is a free log subscription operation binding the contract event 0x33e3a51bf9fedee6e205cd10237a9a8dd3fffed45ac3dee88d2eba92ef8b5d62.
//
// Solidity: event InterchainClientUpdated(address indexed interchainClient)
func (_ExecutionService *ExecutionServiceFilterer) WatchInterchainClientUpdated(opts *bind.WatchOpts, sink chan<- *ExecutionServiceInterchainClientUpdated, interchainClient []common.Address) (event.Subscription, error) {

	var interchainClientRule []interface{}
	for _, interchainClientItem := range interchainClient {
		interchainClientRule = append(interchainClientRule, interchainClientItem)
	}

	logs, sub, err := _ExecutionService.contract.WatchLogs(opts, "InterchainClientUpdated", interchainClientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceInterchainClientUpdated)
				if err := _ExecutionService.contract.UnpackLog(event, "InterchainClientUpdated", log); err != nil {
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

// ParseInterchainClientUpdated is a log parse operation binding the contract event 0x33e3a51bf9fedee6e205cd10237a9a8dd3fffed45ac3dee88d2eba92ef8b5d62.
//
// Solidity: event InterchainClientUpdated(address indexed interchainClient)
func (_ExecutionService *ExecutionServiceFilterer) ParseInterchainClientUpdated(log types.Log) (*ExecutionServiceInterchainClientUpdated, error) {
	event := new(ExecutionServiceInterchainClientUpdated)
	if err := _ExecutionService.contract.UnpackLog(event, "InterchainClientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExecutionService contract.
type ExecutionServiceOwnershipTransferredIterator struct {
	Event *ExecutionServiceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceOwnershipTransferred)
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
		it.Event = new(ExecutionServiceOwnershipTransferred)
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
func (it *ExecutionServiceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceOwnershipTransferred represents a OwnershipTransferred event raised by the ExecutionService contract.
type ExecutionServiceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExecutionService *ExecutionServiceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExecutionServiceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExecutionService.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceOwnershipTransferredIterator{contract: _ExecutionService.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExecutionService *ExecutionServiceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExecutionServiceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExecutionService.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceOwnershipTransferred)
				if err := _ExecutionService.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ExecutionService *ExecutionServiceFilterer) ParseOwnershipTransferred(log types.Log) (*ExecutionServiceOwnershipTransferred, error) {
	event := new(ExecutionServiceOwnershipTransferred)
	if err := _ExecutionService.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceEventsMetaData contains all meta data concerning the ExecutionServiceEvents contract.
var ExecutionServiceEventsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"client\",\"type\":\"address\"}],\"name\":\"ExecutionRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executorEOA\",\"type\":\"address\"}],\"name\":\"ExecutorEOAUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gasOracle\",\"type\":\"address\"}],\"name\":\"GasOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"interchainClient\",\"type\":\"address\"}],\"name\":\"InterchainClientUpdated\",\"type\":\"event\"}]",
}

// ExecutionServiceEventsABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutionServiceEventsMetaData.ABI instead.
var ExecutionServiceEventsABI = ExecutionServiceEventsMetaData.ABI

// ExecutionServiceEvents is an auto generated Go binding around an Ethereum contract.
type ExecutionServiceEvents struct {
	ExecutionServiceEventsCaller     // Read-only binding to the contract
	ExecutionServiceEventsTransactor // Write-only binding to the contract
	ExecutionServiceEventsFilterer   // Log filterer for contract events
}

// ExecutionServiceEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutionServiceEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutionServiceEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceEventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutionServiceEventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutionServiceEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutionServiceEventsSession struct {
	Contract     *ExecutionServiceEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExecutionServiceEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutionServiceEventsCallerSession struct {
	Contract *ExecutionServiceEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ExecutionServiceEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutionServiceEventsTransactorSession struct {
	Contract     *ExecutionServiceEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ExecutionServiceEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutionServiceEventsRaw struct {
	Contract *ExecutionServiceEvents // Generic contract binding to access the raw methods on
}

// ExecutionServiceEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutionServiceEventsCallerRaw struct {
	Contract *ExecutionServiceEventsCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutionServiceEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutionServiceEventsTransactorRaw struct {
	Contract *ExecutionServiceEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutionServiceEvents creates a new instance of ExecutionServiceEvents, bound to a specific deployed contract.
func NewExecutionServiceEvents(address common.Address, backend bind.ContractBackend) (*ExecutionServiceEvents, error) {
	contract, err := bindExecutionServiceEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEvents{ExecutionServiceEventsCaller: ExecutionServiceEventsCaller{contract: contract}, ExecutionServiceEventsTransactor: ExecutionServiceEventsTransactor{contract: contract}, ExecutionServiceEventsFilterer: ExecutionServiceEventsFilterer{contract: contract}}, nil
}

// NewExecutionServiceEventsCaller creates a new read-only instance of ExecutionServiceEvents, bound to a specific deployed contract.
func NewExecutionServiceEventsCaller(address common.Address, caller bind.ContractCaller) (*ExecutionServiceEventsCaller, error) {
	contract, err := bindExecutionServiceEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEventsCaller{contract: contract}, nil
}

// NewExecutionServiceEventsTransactor creates a new write-only instance of ExecutionServiceEvents, bound to a specific deployed contract.
func NewExecutionServiceEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutionServiceEventsTransactor, error) {
	contract, err := bindExecutionServiceEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEventsTransactor{contract: contract}, nil
}

// NewExecutionServiceEventsFilterer creates a new log filterer instance of ExecutionServiceEvents, bound to a specific deployed contract.
func NewExecutionServiceEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutionServiceEventsFilterer, error) {
	contract, err := bindExecutionServiceEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEventsFilterer{contract: contract}, nil
}

// bindExecutionServiceEvents binds a generic wrapper to an already deployed contract.
func bindExecutionServiceEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExecutionServiceEventsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionServiceEvents *ExecutionServiceEventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionServiceEvents.Contract.ExecutionServiceEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionServiceEvents *ExecutionServiceEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionServiceEvents.Contract.ExecutionServiceEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionServiceEvents *ExecutionServiceEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionServiceEvents.Contract.ExecutionServiceEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExecutionServiceEvents *ExecutionServiceEventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExecutionServiceEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExecutionServiceEvents *ExecutionServiceEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExecutionServiceEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExecutionServiceEvents *ExecutionServiceEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExecutionServiceEvents.Contract.contract.Transact(opts, method, params...)
}

// ExecutionServiceEventsExecutionRequestedIterator is returned from FilterExecutionRequested and is used to iterate over the raw logs and unpacked data for ExecutionRequested events raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsExecutionRequestedIterator struct {
	Event *ExecutionServiceEventsExecutionRequested // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceEventsExecutionRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceEventsExecutionRequested)
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
		it.Event = new(ExecutionServiceEventsExecutionRequested)
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
func (it *ExecutionServiceEventsExecutionRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceEventsExecutionRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceEventsExecutionRequested represents a ExecutionRequested event raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsExecutionRequested struct {
	TransactionId [32]byte
	Client        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionRequested is a free log retrieval operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) FilterExecutionRequested(opts *bind.FilterOpts, transactionId [][32]byte) (*ExecutionServiceEventsExecutionRequestedIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.FilterLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEventsExecutionRequestedIterator{contract: _ExecutionServiceEvents.contract, event: "ExecutionRequested", logs: logs, sub: sub}, nil
}

// WatchExecutionRequested is a free log subscription operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) WatchExecutionRequested(opts *bind.WatchOpts, sink chan<- *ExecutionServiceEventsExecutionRequested, transactionId [][32]byte) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.WatchLogs(opts, "ExecutionRequested", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceEventsExecutionRequested)
				if err := _ExecutionServiceEvents.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
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

// ParseExecutionRequested is a log parse operation binding the contract event 0x507e9bdb8950e371753b8570704cf098fb0d67ca247f09a0629971ac80bd1382.
//
// Solidity: event ExecutionRequested(bytes32 indexed transactionId, address client)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) ParseExecutionRequested(log types.Log) (*ExecutionServiceEventsExecutionRequested, error) {
	event := new(ExecutionServiceEventsExecutionRequested)
	if err := _ExecutionServiceEvents.contract.UnpackLog(event, "ExecutionRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceEventsExecutorEOAUpdatedIterator is returned from FilterExecutorEOAUpdated and is used to iterate over the raw logs and unpacked data for ExecutorEOAUpdated events raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsExecutorEOAUpdatedIterator struct {
	Event *ExecutionServiceEventsExecutorEOAUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceEventsExecutorEOAUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceEventsExecutorEOAUpdated)
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
		it.Event = new(ExecutionServiceEventsExecutorEOAUpdated)
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
func (it *ExecutionServiceEventsExecutorEOAUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceEventsExecutorEOAUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceEventsExecutorEOAUpdated represents a ExecutorEOAUpdated event raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsExecutorEOAUpdated struct {
	ExecutorEOA common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecutorEOAUpdated is a free log retrieval operation binding the contract event 0x8a101f72b53416c657555bd234f732510711b6f59f28b161db43892c89b6b3b4.
//
// Solidity: event ExecutorEOAUpdated(address indexed executorEOA)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) FilterExecutorEOAUpdated(opts *bind.FilterOpts, executorEOA []common.Address) (*ExecutionServiceEventsExecutorEOAUpdatedIterator, error) {

	var executorEOARule []interface{}
	for _, executorEOAItem := range executorEOA {
		executorEOARule = append(executorEOARule, executorEOAItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.FilterLogs(opts, "ExecutorEOAUpdated", executorEOARule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEventsExecutorEOAUpdatedIterator{contract: _ExecutionServiceEvents.contract, event: "ExecutorEOAUpdated", logs: logs, sub: sub}, nil
}

// WatchExecutorEOAUpdated is a free log subscription operation binding the contract event 0x8a101f72b53416c657555bd234f732510711b6f59f28b161db43892c89b6b3b4.
//
// Solidity: event ExecutorEOAUpdated(address indexed executorEOA)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) WatchExecutorEOAUpdated(opts *bind.WatchOpts, sink chan<- *ExecutionServiceEventsExecutorEOAUpdated, executorEOA []common.Address) (event.Subscription, error) {

	var executorEOARule []interface{}
	for _, executorEOAItem := range executorEOA {
		executorEOARule = append(executorEOARule, executorEOAItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.WatchLogs(opts, "ExecutorEOAUpdated", executorEOARule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceEventsExecutorEOAUpdated)
				if err := _ExecutionServiceEvents.contract.UnpackLog(event, "ExecutorEOAUpdated", log); err != nil {
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

// ParseExecutorEOAUpdated is a log parse operation binding the contract event 0x8a101f72b53416c657555bd234f732510711b6f59f28b161db43892c89b6b3b4.
//
// Solidity: event ExecutorEOAUpdated(address indexed executorEOA)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) ParseExecutorEOAUpdated(log types.Log) (*ExecutionServiceEventsExecutorEOAUpdated, error) {
	event := new(ExecutionServiceEventsExecutorEOAUpdated)
	if err := _ExecutionServiceEvents.contract.UnpackLog(event, "ExecutorEOAUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceEventsGasOracleUpdatedIterator is returned from FilterGasOracleUpdated and is used to iterate over the raw logs and unpacked data for GasOracleUpdated events raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsGasOracleUpdatedIterator struct {
	Event *ExecutionServiceEventsGasOracleUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceEventsGasOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceEventsGasOracleUpdated)
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
		it.Event = new(ExecutionServiceEventsGasOracleUpdated)
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
func (it *ExecutionServiceEventsGasOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceEventsGasOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceEventsGasOracleUpdated represents a GasOracleUpdated event raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsGasOracleUpdated struct {
	GasOracle common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGasOracleUpdated is a free log retrieval operation binding the contract event 0x0b987d971d57fe66ec2d0cff095b1b547c3076cdd07fdf5b0863f320700dbb90.
//
// Solidity: event GasOracleUpdated(address indexed gasOracle)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) FilterGasOracleUpdated(opts *bind.FilterOpts, gasOracle []common.Address) (*ExecutionServiceEventsGasOracleUpdatedIterator, error) {

	var gasOracleRule []interface{}
	for _, gasOracleItem := range gasOracle {
		gasOracleRule = append(gasOracleRule, gasOracleItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.FilterLogs(opts, "GasOracleUpdated", gasOracleRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEventsGasOracleUpdatedIterator{contract: _ExecutionServiceEvents.contract, event: "GasOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchGasOracleUpdated is a free log subscription operation binding the contract event 0x0b987d971d57fe66ec2d0cff095b1b547c3076cdd07fdf5b0863f320700dbb90.
//
// Solidity: event GasOracleUpdated(address indexed gasOracle)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) WatchGasOracleUpdated(opts *bind.WatchOpts, sink chan<- *ExecutionServiceEventsGasOracleUpdated, gasOracle []common.Address) (event.Subscription, error) {

	var gasOracleRule []interface{}
	for _, gasOracleItem := range gasOracle {
		gasOracleRule = append(gasOracleRule, gasOracleItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.WatchLogs(opts, "GasOracleUpdated", gasOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceEventsGasOracleUpdated)
				if err := _ExecutionServiceEvents.contract.UnpackLog(event, "GasOracleUpdated", log); err != nil {
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

// ParseGasOracleUpdated is a log parse operation binding the contract event 0x0b987d971d57fe66ec2d0cff095b1b547c3076cdd07fdf5b0863f320700dbb90.
//
// Solidity: event GasOracleUpdated(address indexed gasOracle)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) ParseGasOracleUpdated(log types.Log) (*ExecutionServiceEventsGasOracleUpdated, error) {
	event := new(ExecutionServiceEventsGasOracleUpdated)
	if err := _ExecutionServiceEvents.contract.UnpackLog(event, "GasOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExecutionServiceEventsInterchainClientUpdatedIterator is returned from FilterInterchainClientUpdated and is used to iterate over the raw logs and unpacked data for InterchainClientUpdated events raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsInterchainClientUpdatedIterator struct {
	Event *ExecutionServiceEventsInterchainClientUpdated // Event containing the contract specifics and raw log

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
func (it *ExecutionServiceEventsInterchainClientUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExecutionServiceEventsInterchainClientUpdated)
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
		it.Event = new(ExecutionServiceEventsInterchainClientUpdated)
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
func (it *ExecutionServiceEventsInterchainClientUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExecutionServiceEventsInterchainClientUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExecutionServiceEventsInterchainClientUpdated represents a InterchainClientUpdated event raised by the ExecutionServiceEvents contract.
type ExecutionServiceEventsInterchainClientUpdated struct {
	InterchainClient common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInterchainClientUpdated is a free log retrieval operation binding the contract event 0x33e3a51bf9fedee6e205cd10237a9a8dd3fffed45ac3dee88d2eba92ef8b5d62.
//
// Solidity: event InterchainClientUpdated(address indexed interchainClient)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) FilterInterchainClientUpdated(opts *bind.FilterOpts, interchainClient []common.Address) (*ExecutionServiceEventsInterchainClientUpdatedIterator, error) {

	var interchainClientRule []interface{}
	for _, interchainClientItem := range interchainClient {
		interchainClientRule = append(interchainClientRule, interchainClientItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.FilterLogs(opts, "InterchainClientUpdated", interchainClientRule)
	if err != nil {
		return nil, err
	}
	return &ExecutionServiceEventsInterchainClientUpdatedIterator{contract: _ExecutionServiceEvents.contract, event: "InterchainClientUpdated", logs: logs, sub: sub}, nil
}

// WatchInterchainClientUpdated is a free log subscription operation binding the contract event 0x33e3a51bf9fedee6e205cd10237a9a8dd3fffed45ac3dee88d2eba92ef8b5d62.
//
// Solidity: event InterchainClientUpdated(address indexed interchainClient)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) WatchInterchainClientUpdated(opts *bind.WatchOpts, sink chan<- *ExecutionServiceEventsInterchainClientUpdated, interchainClient []common.Address) (event.Subscription, error) {

	var interchainClientRule []interface{}
	for _, interchainClientItem := range interchainClient {
		interchainClientRule = append(interchainClientRule, interchainClientItem)
	}

	logs, sub, err := _ExecutionServiceEvents.contract.WatchLogs(opts, "InterchainClientUpdated", interchainClientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExecutionServiceEventsInterchainClientUpdated)
				if err := _ExecutionServiceEvents.contract.UnpackLog(event, "InterchainClientUpdated", log); err != nil {
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

// ParseInterchainClientUpdated is a log parse operation binding the contract event 0x33e3a51bf9fedee6e205cd10237a9a8dd3fffed45ac3dee88d2eba92ef8b5d62.
//
// Solidity: event InterchainClientUpdated(address indexed interchainClient)
func (_ExecutionServiceEvents *ExecutionServiceEventsFilterer) ParseInterchainClientUpdated(log types.Log) (*ExecutionServiceEventsInterchainClientUpdated, error) {
	event := new(ExecutionServiceEventsInterchainClientUpdated)
	if err := _ExecutionServiceEvents.contract.UnpackLog(event, "InterchainClientUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IExecutionServiceMetaData contains all meta data concerning the IExecutionService contract.
var IExecutionServiceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"executorEOA\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"getExecutionFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"dstChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txPayloadSize\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"executionFee\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"options\",\"type\":\"bytes\"}],\"name\":\"requestExecution\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62014bad": "executorEOA()",
		"c473e7e8": "getExecutionFee(uint256,uint256,bytes)",
		"e4e06522": "requestExecution(uint256,uint256,bytes32,uint256,bytes)",
	},
}

// IExecutionServiceABI is the input ABI used to generate the binding from.
// Deprecated: Use IExecutionServiceMetaData.ABI instead.
var IExecutionServiceABI = IExecutionServiceMetaData.ABI

// Deprecated: Use IExecutionServiceMetaData.Sigs instead.
// IExecutionServiceFuncSigs maps the 4-byte function signature to its string representation.
var IExecutionServiceFuncSigs = IExecutionServiceMetaData.Sigs

// IExecutionService is an auto generated Go binding around an Ethereum contract.
type IExecutionService struct {
	IExecutionServiceCaller     // Read-only binding to the contract
	IExecutionServiceTransactor // Write-only binding to the contract
	IExecutionServiceFilterer   // Log filterer for contract events
}

// IExecutionServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IExecutionServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IExecutionServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IExecutionServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IExecutionServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IExecutionServiceSession struct {
	Contract     *IExecutionService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IExecutionServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IExecutionServiceCallerSession struct {
	Contract *IExecutionServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IExecutionServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IExecutionServiceTransactorSession struct {
	Contract     *IExecutionServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IExecutionServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IExecutionServiceRaw struct {
	Contract *IExecutionService // Generic contract binding to access the raw methods on
}

// IExecutionServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IExecutionServiceCallerRaw struct {
	Contract *IExecutionServiceCaller // Generic read-only contract binding to access the raw methods on
}

// IExecutionServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IExecutionServiceTransactorRaw struct {
	Contract *IExecutionServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIExecutionService creates a new instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionService(address common.Address, backend bind.ContractBackend) (*IExecutionService, error) {
	contract, err := bindIExecutionService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IExecutionService{IExecutionServiceCaller: IExecutionServiceCaller{contract: contract}, IExecutionServiceTransactor: IExecutionServiceTransactor{contract: contract}, IExecutionServiceFilterer: IExecutionServiceFilterer{contract: contract}}, nil
}

// NewIExecutionServiceCaller creates a new read-only instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceCaller(address common.Address, caller bind.ContractCaller) (*IExecutionServiceCaller, error) {
	contract, err := bindIExecutionService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceCaller{contract: contract}, nil
}

// NewIExecutionServiceTransactor creates a new write-only instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*IExecutionServiceTransactor, error) {
	contract, err := bindIExecutionService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceTransactor{contract: contract}, nil
}

// NewIExecutionServiceFilterer creates a new log filterer instance of IExecutionService, bound to a specific deployed contract.
func NewIExecutionServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*IExecutionServiceFilterer, error) {
	contract, err := bindIExecutionService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IExecutionServiceFilterer{contract: contract}, nil
}

// bindIExecutionService binds a generic wrapper to an already deployed contract.
func bindIExecutionService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IExecutionServiceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionService *IExecutionServiceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionService.Contract.IExecutionServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionService *IExecutionServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionService.Contract.IExecutionServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionService *IExecutionServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionService.Contract.IExecutionServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IExecutionService *IExecutionServiceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IExecutionService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IExecutionService *IExecutionServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IExecutionService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IExecutionService *IExecutionServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IExecutionService.Contract.contract.Transact(opts, method, params...)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceCaller) ExecutorEOA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IExecutionService.contract.Call(opts, &out, "executorEOA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceSession) ExecutorEOA() (common.Address, error) {
	return _IExecutionService.Contract.ExecutorEOA(&_IExecutionService.CallOpts)
}

// ExecutorEOA is a free data retrieval call binding the contract method 0x62014bad.
//
// Solidity: function executorEOA() view returns(address)
func (_IExecutionService *IExecutionServiceCallerSession) ExecutorEOA() (common.Address, error) {
	return _IExecutionService.Contract.ExecutorEOA(&_IExecutionService.CallOpts)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceCaller) GetExecutionFee(opts *bind.CallOpts, dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	var out []interface{}
	err := _IExecutionService.contract.Call(opts, &out, "getExecutionFee", dstChainId, txPayloadSize, options)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _IExecutionService.Contract.GetExecutionFee(&_IExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// GetExecutionFee is a free data retrieval call binding the contract method 0xc473e7e8.
//
// Solidity: function getExecutionFee(uint256 dstChainId, uint256 txPayloadSize, bytes options) view returns(uint256)
func (_IExecutionService *IExecutionServiceCallerSession) GetExecutionFee(dstChainId *big.Int, txPayloadSize *big.Int, options []byte) (*big.Int, error) {
	return _IExecutionService.Contract.GetExecutionFee(&_IExecutionService.CallOpts, dstChainId, txPayloadSize, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceTransactor) RequestExecution(opts *bind.TransactOpts, dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.contract.Transact(opts, "requestExecution", dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// RequestExecution is a paid mutator transaction binding the contract method 0xe4e06522.
//
// Solidity: function requestExecution(uint256 dstChainId, uint256 txPayloadSize, bytes32 transactionId, uint256 executionFee, bytes options) returns()
func (_IExecutionService *IExecutionServiceTransactorSession) RequestExecution(dstChainId *big.Int, txPayloadSize *big.Int, transactionId [32]byte, executionFee *big.Int, options []byte) (*types.Transaction, error) {
	return _IExecutionService.Contract.RequestExecution(&_IExecutionService.TransactOpts, dstChainId, txPayloadSize, transactionId, executionFee, options)
}

// IGasOracleMetaData contains all meta data concerning the IGasOracle contract.
var IGasOracleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"convertRemoteValueToLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInLocalUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"remoteChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldataSize\",\"type\":\"uint256\"}],\"name\":\"estimateTxCostInRemoteUnits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1e7b9287": "convertRemoteValueToLocalUnits(uint256,uint256)",
		"5cbd3c48": "estimateTxCostInLocalUnits(uint256,uint256,uint256)",
		"fd6a7167": "estimateTxCostInRemoteUnits(uint256,uint256,uint256)",
	},
}

// IGasOracleABI is the input ABI used to generate the binding from.
// Deprecated: Use IGasOracleMetaData.ABI instead.
var IGasOracleABI = IGasOracleMetaData.ABI

// Deprecated: Use IGasOracleMetaData.Sigs instead.
// IGasOracleFuncSigs maps the 4-byte function signature to its string representation.
var IGasOracleFuncSigs = IGasOracleMetaData.Sigs

// IGasOracle is an auto generated Go binding around an Ethereum contract.
type IGasOracle struct {
	IGasOracleCaller     // Read-only binding to the contract
	IGasOracleTransactor // Write-only binding to the contract
	IGasOracleFilterer   // Log filterer for contract events
}

// IGasOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGasOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGasOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGasOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGasOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGasOracleSession struct {
	Contract     *IGasOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGasOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGasOracleCallerSession struct {
	Contract *IGasOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IGasOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGasOracleTransactorSession struct {
	Contract     *IGasOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IGasOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGasOracleRaw struct {
	Contract *IGasOracle // Generic contract binding to access the raw methods on
}

// IGasOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGasOracleCallerRaw struct {
	Contract *IGasOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IGasOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGasOracleTransactorRaw struct {
	Contract *IGasOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGasOracle creates a new instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracle(address common.Address, backend bind.ContractBackend) (*IGasOracle, error) {
	contract, err := bindIGasOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGasOracle{IGasOracleCaller: IGasOracleCaller{contract: contract}, IGasOracleTransactor: IGasOracleTransactor{contract: contract}, IGasOracleFilterer: IGasOracleFilterer{contract: contract}}, nil
}

// NewIGasOracleCaller creates a new read-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleCaller(address common.Address, caller bind.ContractCaller) (*IGasOracleCaller, error) {
	contract, err := bindIGasOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleCaller{contract: contract}, nil
}

// NewIGasOracleTransactor creates a new write-only instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IGasOracleTransactor, error) {
	contract, err := bindIGasOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGasOracleTransactor{contract: contract}, nil
}

// NewIGasOracleFilterer creates a new log filterer instance of IGasOracle, bound to a specific deployed contract.
func NewIGasOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IGasOracleFilterer, error) {
	contract, err := bindIGasOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGasOracleFilterer{contract: contract}, nil
}

// bindIGasOracle binds a generic wrapper to an already deployed contract.
func bindIGasOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGasOracleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.IGasOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.IGasOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGasOracle *IGasOracleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGasOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGasOracle *IGasOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGasOracle *IGasOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGasOracle.Contract.contract.Transact(opts, method, params...)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) ConvertRemoteValueToLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "convertRemoteValueToLocalUnits", remoteChainId, value)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// ConvertRemoteValueToLocalUnits is a free data retrieval call binding the contract method 0x1e7b9287.
//
// Solidity: function convertRemoteValueToLocalUnits(uint256 remoteChainId, uint256 value) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) ConvertRemoteValueToLocalUnits(remoteChainId *big.Int, value *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.ConvertRemoteValueToLocalUnits(&_IGasOracle.CallOpts, remoteChainId, value)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInLocalUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInLocalUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInLocalUnits is a free data retrieval call binding the contract method 0x5cbd3c48.
//
// Solidity: function estimateTxCostInLocalUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInLocalUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInLocalUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCaller) EstimateTxCostInRemoteUnits(opts *bind.CallOpts, remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IGasOracle.contract.Call(opts, &out, "estimateTxCostInRemoteUnits", remoteChainId, gasLimit, calldataSize)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// EstimateTxCostInRemoteUnits is a free data retrieval call binding the contract method 0xfd6a7167.
//
// Solidity: function estimateTxCostInRemoteUnits(uint256 remoteChainId, uint256 gasLimit, uint256 calldataSize) view returns(uint256)
func (_IGasOracle *IGasOracleCallerSession) EstimateTxCostInRemoteUnits(remoteChainId *big.Int, gasLimit *big.Int, calldataSize *big.Int) (*big.Int, error) {
	return _IGasOracle.Contract.EstimateTxCostInRemoteUnits(&_IGasOracle.CallOpts, remoteChainId, gasLimit, calldataSize)
}

// OptionsLibMetaData contains all meta data concerning the OptionsLib contract.
var OptionsLibMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"OptionsLib__IncorrectVersion\",\"type\":\"error\"}]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122092579146fe89b4651b63b94a9c6c15c0f5814d6d2f920060c2750740d224962264736f6c63430008140033",
}

// OptionsLibABI is the input ABI used to generate the binding from.
// Deprecated: Use OptionsLibMetaData.ABI instead.
var OptionsLibABI = OptionsLibMetaData.ABI

// OptionsLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OptionsLibMetaData.Bin instead.
var OptionsLibBin = OptionsLibMetaData.Bin

// DeployOptionsLib deploys a new Ethereum contract, binding an instance of OptionsLib to it.
func DeployOptionsLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OptionsLib, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OptionsLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// OptionsLib is an auto generated Go binding around an Ethereum contract.
type OptionsLib struct {
	OptionsLibCaller     // Read-only binding to the contract
	OptionsLibTransactor // Write-only binding to the contract
	OptionsLibFilterer   // Log filterer for contract events
}

// OptionsLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type OptionsLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OptionsLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OptionsLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OptionsLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OptionsLibSession struct {
	Contract     *OptionsLib       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OptionsLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OptionsLibCallerSession struct {
	Contract *OptionsLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OptionsLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OptionsLibTransactorSession struct {
	Contract     *OptionsLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OptionsLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type OptionsLibRaw struct {
	Contract *OptionsLib // Generic contract binding to access the raw methods on
}

// OptionsLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OptionsLibCallerRaw struct {
	Contract *OptionsLibCaller // Generic read-only contract binding to access the raw methods on
}

// OptionsLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OptionsLibTransactorRaw struct {
	Contract *OptionsLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOptionsLib creates a new instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLib(address common.Address, backend bind.ContractBackend) (*OptionsLib, error) {
	contract, err := bindOptionsLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OptionsLib{OptionsLibCaller: OptionsLibCaller{contract: contract}, OptionsLibTransactor: OptionsLibTransactor{contract: contract}, OptionsLibFilterer: OptionsLibFilterer{contract: contract}}, nil
}

// NewOptionsLibCaller creates a new read-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibCaller(address common.Address, caller bind.ContractCaller) (*OptionsLibCaller, error) {
	contract, err := bindOptionsLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibCaller{contract: contract}, nil
}

// NewOptionsLibTransactor creates a new write-only instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibTransactor(address common.Address, transactor bind.ContractTransactor) (*OptionsLibTransactor, error) {
	contract, err := bindOptionsLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OptionsLibTransactor{contract: contract}, nil
}

// NewOptionsLibFilterer creates a new log filterer instance of OptionsLib, bound to a specific deployed contract.
func NewOptionsLibFilterer(address common.Address, filterer bind.ContractFilterer) (*OptionsLibFilterer, error) {
	contract, err := bindOptionsLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OptionsLibFilterer{contract: contract}, nil
}

// bindOptionsLib binds a generic wrapper to an already deployed contract.
func bindOptionsLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OptionsLibMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.OptionsLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.OptionsLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OptionsLib *OptionsLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OptionsLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OptionsLib *OptionsLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OptionsLib *OptionsLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OptionsLib.Contract.contract.Transact(opts, method, params...)
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
		"715018a6": "renounceOwnership()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// OwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableMetaData.ABI instead.
var OwnableABI = OwnableMetaData.ABI

// Deprecated: Use OwnableMetaData.Sigs instead.
// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = OwnableMetaData.Sigs

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
