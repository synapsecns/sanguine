// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package messagetransmitter

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

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220692b67c87ddcd80ea0806a3c207ef65bcbcc1b11c3697fc6ea2f736cca0cda7964736f6c63430007060033",
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
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// AttestableMetaData contains all meta data concerning the Attestable contract.
var AttestableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAttesterManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"AttesterManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSignatureThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"SignatureThresholdUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attesterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"disableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttester\",\"type\":\"address\"}],\"name\":\"enableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getEnabledAttester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumEnabledAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"isEnabledAttester\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"setSignatureThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"updateAttesterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"9b0d94b7": "attesterManager()",
		"2d025080": "disableAttester(address)",
		"fae36879": "enableAttester(address)",
		"beb673d8": "getEnabledAttester(uint256)",
		"51079a53": "getNumEnabledAttesters()",
		"7af82f60": "isEnabledAttester(address)",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"bbde5374": "setSignatureThreshold(uint256)",
		"a82f2e26": "signatureThreshold()",
		"f2fde38b": "transferOwnership(address)",
		"de7769d4": "updateAttesterManager(address)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200116638038062001166833981810160405260208110156200003757600080fd5b50516200004d620000476200006f565b62000073565b62000058336200009d565b60016002556200006881620000bf565b50620002f6565b3390565b600180546001600160a01b03191690556200009a816200021f602090811b62000a2c17901c565b50565b600580546001600160a01b0319166001600160a01b0392909216919091179055565b6005546001600160a01b031633146200011f576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b6001600160a01b0381166200017b576040805162461bcd60e51b815260206004820152601c60248201527f4e6577206174746573746572206d757374206265206e6f6e7a65726f00000000604482015290519081900360640190fd5b620001968160036200026f60201b62000aa11790919060201c565b620001e8576040805162461bcd60e51b815260206004820152601860248201527f417474657374657220616c726561647920656e61626c65640000000000000000604482015290519081900360640190fd5b6040516001600160a01b038216907f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f90600090a250565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600062000286836001600160a01b0384166200028f565b90505b92915050565b60006200029d8383620002de565b620002d55750815460018181018455600084815260208082209093018490558454848252828601909352604090209190915562000289565b50600062000289565b60009081526001919091016020526040902054151590565b610e6080620003066000396000f3fe608060405234801561001057600080fd5b50600436106100df5760003560e01c8063a82f2e261161008c578063de7769d411610066578063de7769d4146101fd578063e30c397814610230578063f2fde38b14610238578063fae368791461026b576100df565b8063a82f2e26146101bb578063bbde5374146101c3578063beb673d8146101e0576100df565b80637af82f60116100bd5780637af82f601461013b5780638da5cb5b146101825780639b0d94b7146101b3576100df565b80632d025080146100e457806351079a531461011957806379ba509714610133575b600080fd5b610117600480360360208110156100fa57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661029e565b005b610121610462565b60408051918252519081900360200190f35b610117610473565b61016e6004803603602081101561015157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166104fc565b604080519115158252519081900360200190f35b61018a61050f565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b61018a61052b565b610121610547565b610117600480360360208110156101d957600080fd5b503561054d565b61018a600480360360208110156101f657600080fd5b5035610707565b6101176004803603602081101561021357600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610714565b61018a610804565b6101176004803603602081101561024e57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610820565b6101176004803603602081101561028157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166108b8565b60055473ffffffffffffffffffffffffffffffffffffffff16331461030a576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b6000610314610462565b90506001811161036b576040805162461bcd60e51b815260206004820152601960248201527f546f6f2066657720656e61626c65642061747465737465727300000000000000604482015290519081900360640190fd5b60025481116103c1576040805162461bcd60e51b815260206004820152601e60248201527f5369676e6174757265207468726573686f6c6420697320746f6f206c6f770000604482015290519081900360640190fd5b6103cc600383610aca565b61041d576040805162461bcd60e51b815260206004820152601960248201527f417474657374657220616c72656164792064697361626c656400000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8316907f78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc90600090a25050565b600061046e6003610aec565b905090565b600061047d610af7565b90508073ffffffffffffffffffffffffffffffffffffffff1661049e610804565b73ffffffffffffffffffffffffffffffffffffffff16146104f05760405162461bcd60e51b8152600401808060200182810382526029815260200180610e026029913960400191505060405180910390fd5b6104f981610afb565b50565b6000610509600383610b2c565b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60055473ffffffffffffffffffffffffffffffffffffffff1690565b60025481565b60055473ffffffffffffffffffffffffffffffffffffffff1633146105b9576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b8061060b576040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265207468726573686f6c640000000000604482015290519081900360640190fd5b6106156003610aec565b811115610669576040805162461bcd60e51b815260206004820181905260248201527f4e6577207369676e6174757265207468726573686f6c6420746f6f2068696768604482015290519081900360640190fd5b6002548114156106c0576040805162461bcd60e51b815260206004820152601f60248201527f5369676e6174757265207468726573686f6c6420616c72656164792073657400604482015290519081900360640190fd5b6002805490829055604080518281526020810184905281517f149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede929181900390910190a15050565b6000610509600383610b4e565b61071c610b5a565b73ffffffffffffffffffffffffffffffffffffffff8116610784576040805162461bcd60e51b815260206004820181905260248201527f496e76616c6964206174746573746572206d616e616765722061646472657373604482015290519081900360640190fd5b60055473ffffffffffffffffffffffffffffffffffffffff166107a682610bea565b8173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b60015473ffffffffffffffffffffffffffffffffffffffff1690565b610828610b5a565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831690811790915561087361050f565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60055473ffffffffffffffffffffffffffffffffffffffff163314610924576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b73ffffffffffffffffffffffffffffffffffffffff811661098c576040805162461bcd60e51b815260206004820152601c60248201527f4e6577206174746573746572206d757374206265206e6f6e7a65726f00000000604482015290519081900360640190fd5b610997600382610aa1565b6109e8576040805162461bcd60e51b815260206004820152601860248201527f417474657374657220616c726561647920656e61626c65640000000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f90600090a250565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000610ac38373ffffffffffffffffffffffffffffffffffffffff8416610c31565b9392505050565b6000610ac38373ffffffffffffffffffffffffffffffffffffffff8416610c7b565b600061050982610d5f565b3390565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556104f981610a2c565b6000610ac38373ffffffffffffffffffffffffffffffffffffffff8416610d63565b6000610ac38383610d7b565b610b62610af7565b73ffffffffffffffffffffffffffffffffffffffff16610b8061050f565b73ffffffffffffffffffffffffffffffffffffffff1614610be8576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b565b600580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b6000610c3d8383610d63565b610c7357508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610509565b506000610509565b60008181526001830160205260408120548015610d555783547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8083019190810190600090879083908110610ccc57fe5b9060005260206000200154905080876000018481548110610ce957fe5b600091825260208083209091019290925582815260018981019092526040902090840190558654879080610d1957fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050610509565b6000915050610509565b5490565b60009081526001919091016020526040902054151590565b81546000908210610dbd5760405162461bcd60e51b8152600401808060200182810382526022815260200180610de06022913960400191505060405180910390fd5b826000018281548110610dcc57fe5b906000526020600020015490509291505056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64734f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572a26469706673582212202a1b157f6886ccbf5d061a41222b10dee70e3f54a7b84409b3f7347251e260a864736f6c63430007060033",
}

// AttestableABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestableMetaData.ABI instead.
var AttestableABI = AttestableMetaData.ABI

// Deprecated: Use AttestableMetaData.Sigs instead.
// AttestableFuncSigs maps the 4-byte function signature to its string representation.
var AttestableFuncSigs = AttestableMetaData.Sigs

// AttestableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestableMetaData.Bin instead.
var AttestableBin = AttestableMetaData.Bin

// DeployAttestable deploys a new Ethereum contract, binding an instance of Attestable to it.
func DeployAttestable(auth *bind.TransactOpts, backend bind.ContractBackend, attester common.Address) (common.Address, *types.Transaction, *Attestable, error) {
	parsed, err := AttestableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestableBin), backend, attester)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attestable{AttestableCaller: AttestableCaller{contract: contract}, AttestableTransactor: AttestableTransactor{contract: contract}, AttestableFilterer: AttestableFilterer{contract: contract}}, nil
}

// Attestable is an auto generated Go binding around an Ethereum contract.
type Attestable struct {
	AttestableCaller     // Read-only binding to the contract
	AttestableTransactor // Write-only binding to the contract
	AttestableFilterer   // Log filterer for contract events
}

// AttestableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestableSession struct {
	Contract     *Attestable       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestableCallerSession struct {
	Contract *AttestableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AttestableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestableTransactorSession struct {
	Contract     *AttestableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AttestableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestableRaw struct {
	Contract *Attestable // Generic contract binding to access the raw methods on
}

// AttestableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestableCallerRaw struct {
	Contract *AttestableCaller // Generic read-only contract binding to access the raw methods on
}

// AttestableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestableTransactorRaw struct {
	Contract *AttestableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestable creates a new instance of Attestable, bound to a specific deployed contract.
func NewAttestable(address common.Address, backend bind.ContractBackend) (*Attestable, error) {
	contract, err := bindAttestable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attestable{AttestableCaller: AttestableCaller{contract: contract}, AttestableTransactor: AttestableTransactor{contract: contract}, AttestableFilterer: AttestableFilterer{contract: contract}}, nil
}

// NewAttestableCaller creates a new read-only instance of Attestable, bound to a specific deployed contract.
func NewAttestableCaller(address common.Address, caller bind.ContractCaller) (*AttestableCaller, error) {
	contract, err := bindAttestable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestableCaller{contract: contract}, nil
}

// NewAttestableTransactor creates a new write-only instance of Attestable, bound to a specific deployed contract.
func NewAttestableTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestableTransactor, error) {
	contract, err := bindAttestable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestableTransactor{contract: contract}, nil
}

// NewAttestableFilterer creates a new log filterer instance of Attestable, bound to a specific deployed contract.
func NewAttestableFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestableFilterer, error) {
	contract, err := bindAttestable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestableFilterer{contract: contract}, nil
}

// bindAttestable binds a generic wrapper to an already deployed contract.
func bindAttestable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttestableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestable *AttestableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestable.Contract.AttestableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestable *AttestableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestable.Contract.AttestableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestable *AttestableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestable.Contract.AttestableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestable *AttestableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestable *AttestableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestable *AttestableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestable.Contract.contract.Transact(opts, method, params...)
}

// AttesterManager is a free data retrieval call binding the contract method 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (_Attestable *AttestableCaller) AttesterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Attestable.contract.Call(opts, &out, "attesterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttesterManager is a free data retrieval call binding the contract method 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (_Attestable *AttestableSession) AttesterManager() (common.Address, error) {
	return _Attestable.Contract.AttesterManager(&_Attestable.CallOpts)
}

// AttesterManager is a free data retrieval call binding the contract method 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (_Attestable *AttestableCallerSession) AttesterManager() (common.Address, error) {
	return _Attestable.Contract.AttesterManager(&_Attestable.CallOpts)
}

// GetEnabledAttester is a free data retrieval call binding the contract method 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (_Attestable *AttestableCaller) GetEnabledAttester(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Attestable.contract.Call(opts, &out, "getEnabledAttester", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnabledAttester is a free data retrieval call binding the contract method 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (_Attestable *AttestableSession) GetEnabledAttester(index *big.Int) (common.Address, error) {
	return _Attestable.Contract.GetEnabledAttester(&_Attestable.CallOpts, index)
}

// GetEnabledAttester is a free data retrieval call binding the contract method 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (_Attestable *AttestableCallerSession) GetEnabledAttester(index *big.Int) (common.Address, error) {
	return _Attestable.Contract.GetEnabledAttester(&_Attestable.CallOpts, index)
}

// GetNumEnabledAttesters is a free data retrieval call binding the contract method 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (_Attestable *AttestableCaller) GetNumEnabledAttesters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestable.contract.Call(opts, &out, "getNumEnabledAttesters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumEnabledAttesters is a free data retrieval call binding the contract method 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (_Attestable *AttestableSession) GetNumEnabledAttesters() (*big.Int, error) {
	return _Attestable.Contract.GetNumEnabledAttesters(&_Attestable.CallOpts)
}

// GetNumEnabledAttesters is a free data retrieval call binding the contract method 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (_Attestable *AttestableCallerSession) GetNumEnabledAttesters() (*big.Int, error) {
	return _Attestable.Contract.GetNumEnabledAttesters(&_Attestable.CallOpts)
}

// IsEnabledAttester is a free data retrieval call binding the contract method 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (_Attestable *AttestableCaller) IsEnabledAttester(opts *bind.CallOpts, attester common.Address) (bool, error) {
	var out []interface{}
	err := _Attestable.contract.Call(opts, &out, "isEnabledAttester", attester)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnabledAttester is a free data retrieval call binding the contract method 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (_Attestable *AttestableSession) IsEnabledAttester(attester common.Address) (bool, error) {
	return _Attestable.Contract.IsEnabledAttester(&_Attestable.CallOpts, attester)
}

// IsEnabledAttester is a free data retrieval call binding the contract method 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (_Attestable *AttestableCallerSession) IsEnabledAttester(attester common.Address) (bool, error) {
	return _Attestable.Contract.IsEnabledAttester(&_Attestable.CallOpts, attester)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attestable *AttestableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Attestable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attestable *AttestableSession) Owner() (common.Address, error) {
	return _Attestable.Contract.Owner(&_Attestable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attestable *AttestableCallerSession) Owner() (common.Address, error) {
	return _Attestable.Contract.Owner(&_Attestable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Attestable *AttestableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Attestable.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Attestable *AttestableSession) PendingOwner() (common.Address, error) {
	return _Attestable.Contract.PendingOwner(&_Attestable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Attestable *AttestableCallerSession) PendingOwner() (common.Address, error) {
	return _Attestable.Contract.PendingOwner(&_Attestable.CallOpts)
}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (_Attestable *AttestableCaller) SignatureThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestable.contract.Call(opts, &out, "signatureThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (_Attestable *AttestableSession) SignatureThreshold() (*big.Int, error) {
	return _Attestable.Contract.SignatureThreshold(&_Attestable.CallOpts)
}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (_Attestable *AttestableCallerSession) SignatureThreshold() (*big.Int, error) {
	return _Attestable.Contract.SignatureThreshold(&_Attestable.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Attestable *AttestableTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestable.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Attestable *AttestableSession) AcceptOwnership() (*types.Transaction, error) {
	return _Attestable.Contract.AcceptOwnership(&_Attestable.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Attestable *AttestableTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Attestable.Contract.AcceptOwnership(&_Attestable.TransactOpts)
}

// DisableAttester is a paid mutator transaction binding the contract method 0x2d025080.
//
// Solidity: function disableAttester(address attester) returns()
func (_Attestable *AttestableTransactor) DisableAttester(opts *bind.TransactOpts, attester common.Address) (*types.Transaction, error) {
	return _Attestable.contract.Transact(opts, "disableAttester", attester)
}

// DisableAttester is a paid mutator transaction binding the contract method 0x2d025080.
//
// Solidity: function disableAttester(address attester) returns()
func (_Attestable *AttestableSession) DisableAttester(attester common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.DisableAttester(&_Attestable.TransactOpts, attester)
}

// DisableAttester is a paid mutator transaction binding the contract method 0x2d025080.
//
// Solidity: function disableAttester(address attester) returns()
func (_Attestable *AttestableTransactorSession) DisableAttester(attester common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.DisableAttester(&_Attestable.TransactOpts, attester)
}

// EnableAttester is a paid mutator transaction binding the contract method 0xfae36879.
//
// Solidity: function enableAttester(address newAttester) returns()
func (_Attestable *AttestableTransactor) EnableAttester(opts *bind.TransactOpts, newAttester common.Address) (*types.Transaction, error) {
	return _Attestable.contract.Transact(opts, "enableAttester", newAttester)
}

// EnableAttester is a paid mutator transaction binding the contract method 0xfae36879.
//
// Solidity: function enableAttester(address newAttester) returns()
func (_Attestable *AttestableSession) EnableAttester(newAttester common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.EnableAttester(&_Attestable.TransactOpts, newAttester)
}

// EnableAttester is a paid mutator transaction binding the contract method 0xfae36879.
//
// Solidity: function enableAttester(address newAttester) returns()
func (_Attestable *AttestableTransactorSession) EnableAttester(newAttester common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.EnableAttester(&_Attestable.TransactOpts, newAttester)
}

// SetSignatureThreshold is a paid mutator transaction binding the contract method 0xbbde5374.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (_Attestable *AttestableTransactor) SetSignatureThreshold(opts *bind.TransactOpts, newSignatureThreshold *big.Int) (*types.Transaction, error) {
	return _Attestable.contract.Transact(opts, "setSignatureThreshold", newSignatureThreshold)
}

// SetSignatureThreshold is a paid mutator transaction binding the contract method 0xbbde5374.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (_Attestable *AttestableSession) SetSignatureThreshold(newSignatureThreshold *big.Int) (*types.Transaction, error) {
	return _Attestable.Contract.SetSignatureThreshold(&_Attestable.TransactOpts, newSignatureThreshold)
}

// SetSignatureThreshold is a paid mutator transaction binding the contract method 0xbbde5374.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (_Attestable *AttestableTransactorSession) SetSignatureThreshold(newSignatureThreshold *big.Int) (*types.Transaction, error) {
	return _Attestable.Contract.SetSignatureThreshold(&_Attestable.TransactOpts, newSignatureThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Attestable *AttestableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Attestable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Attestable *AttestableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.TransferOwnership(&_Attestable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Attestable *AttestableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.TransferOwnership(&_Attestable.TransactOpts, newOwner)
}

// UpdateAttesterManager is a paid mutator transaction binding the contract method 0xde7769d4.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (_Attestable *AttestableTransactor) UpdateAttesterManager(opts *bind.TransactOpts, newAttesterManager common.Address) (*types.Transaction, error) {
	return _Attestable.contract.Transact(opts, "updateAttesterManager", newAttesterManager)
}

// UpdateAttesterManager is a paid mutator transaction binding the contract method 0xde7769d4.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (_Attestable *AttestableSession) UpdateAttesterManager(newAttesterManager common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.UpdateAttesterManager(&_Attestable.TransactOpts, newAttesterManager)
}

// UpdateAttesterManager is a paid mutator transaction binding the contract method 0xde7769d4.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (_Attestable *AttestableTransactorSession) UpdateAttesterManager(newAttesterManager common.Address) (*types.Transaction, error) {
	return _Attestable.Contract.UpdateAttesterManager(&_Attestable.TransactOpts, newAttesterManager)
}

// AttestableAttesterDisabledIterator is returned from FilterAttesterDisabled and is used to iterate over the raw logs and unpacked data for AttesterDisabled events raised by the Attestable contract.
type AttestableAttesterDisabledIterator struct {
	Event *AttestableAttesterDisabled // Event containing the contract specifics and raw log

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
func (it *AttestableAttesterDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestableAttesterDisabled)
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
		it.Event = new(AttestableAttesterDisabled)
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
func (it *AttestableAttesterDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestableAttesterDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestableAttesterDisabled represents a AttesterDisabled event raised by the Attestable contract.
type AttestableAttesterDisabled struct {
	Attester common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttesterDisabled is a free log retrieval operation binding the contract event 0x78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (_Attestable *AttestableFilterer) FilterAttesterDisabled(opts *bind.FilterOpts, attester []common.Address) (*AttestableAttesterDisabledIterator, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _Attestable.contract.FilterLogs(opts, "AttesterDisabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return &AttestableAttesterDisabledIterator{contract: _Attestable.contract, event: "AttesterDisabled", logs: logs, sub: sub}, nil
}

// WatchAttesterDisabled is a free log subscription operation binding the contract event 0x78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (_Attestable *AttestableFilterer) WatchAttesterDisabled(opts *bind.WatchOpts, sink chan<- *AttestableAttesterDisabled, attester []common.Address) (event.Subscription, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _Attestable.contract.WatchLogs(opts, "AttesterDisabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestableAttesterDisabled)
				if err := _Attestable.contract.UnpackLog(event, "AttesterDisabled", log); err != nil {
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

// ParseAttesterDisabled is a log parse operation binding the contract event 0x78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (_Attestable *AttestableFilterer) ParseAttesterDisabled(log types.Log) (*AttestableAttesterDisabled, error) {
	event := new(AttestableAttesterDisabled)
	if err := _Attestable.contract.UnpackLog(event, "AttesterDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestableAttesterEnabledIterator is returned from FilterAttesterEnabled and is used to iterate over the raw logs and unpacked data for AttesterEnabled events raised by the Attestable contract.
type AttestableAttesterEnabledIterator struct {
	Event *AttestableAttesterEnabled // Event containing the contract specifics and raw log

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
func (it *AttestableAttesterEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestableAttesterEnabled)
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
		it.Event = new(AttestableAttesterEnabled)
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
func (it *AttestableAttesterEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestableAttesterEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestableAttesterEnabled represents a AttesterEnabled event raised by the Attestable contract.
type AttestableAttesterEnabled struct {
	Attester common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttesterEnabled is a free log retrieval operation binding the contract event 0x5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (_Attestable *AttestableFilterer) FilterAttesterEnabled(opts *bind.FilterOpts, attester []common.Address) (*AttestableAttesterEnabledIterator, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _Attestable.contract.FilterLogs(opts, "AttesterEnabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return &AttestableAttesterEnabledIterator{contract: _Attestable.contract, event: "AttesterEnabled", logs: logs, sub: sub}, nil
}

// WatchAttesterEnabled is a free log subscription operation binding the contract event 0x5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (_Attestable *AttestableFilterer) WatchAttesterEnabled(opts *bind.WatchOpts, sink chan<- *AttestableAttesterEnabled, attester []common.Address) (event.Subscription, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _Attestable.contract.WatchLogs(opts, "AttesterEnabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestableAttesterEnabled)
				if err := _Attestable.contract.UnpackLog(event, "AttesterEnabled", log); err != nil {
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

// ParseAttesterEnabled is a log parse operation binding the contract event 0x5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (_Attestable *AttestableFilterer) ParseAttesterEnabled(log types.Log) (*AttestableAttesterEnabled, error) {
	event := new(AttestableAttesterEnabled)
	if err := _Attestable.contract.UnpackLog(event, "AttesterEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestableAttesterManagerUpdatedIterator is returned from FilterAttesterManagerUpdated and is used to iterate over the raw logs and unpacked data for AttesterManagerUpdated events raised by the Attestable contract.
type AttestableAttesterManagerUpdatedIterator struct {
	Event *AttestableAttesterManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AttestableAttesterManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestableAttesterManagerUpdated)
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
		it.Event = new(AttestableAttesterManagerUpdated)
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
func (it *AttestableAttesterManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestableAttesterManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestableAttesterManagerUpdated represents a AttesterManagerUpdated event raised by the Attestable contract.
type AttestableAttesterManagerUpdated struct {
	PreviousAttesterManager common.Address
	NewAttesterManager      common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAttesterManagerUpdated is a free log retrieval operation binding the contract event 0x0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d3460219.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (_Attestable *AttestableFilterer) FilterAttesterManagerUpdated(opts *bind.FilterOpts, previousAttesterManager []common.Address, newAttesterManager []common.Address) (*AttestableAttesterManagerUpdatedIterator, error) {

	var previousAttesterManagerRule []interface{}
	for _, previousAttesterManagerItem := range previousAttesterManager {
		previousAttesterManagerRule = append(previousAttesterManagerRule, previousAttesterManagerItem)
	}
	var newAttesterManagerRule []interface{}
	for _, newAttesterManagerItem := range newAttesterManager {
		newAttesterManagerRule = append(newAttesterManagerRule, newAttesterManagerItem)
	}

	logs, sub, err := _Attestable.contract.FilterLogs(opts, "AttesterManagerUpdated", previousAttesterManagerRule, newAttesterManagerRule)
	if err != nil {
		return nil, err
	}
	return &AttestableAttesterManagerUpdatedIterator{contract: _Attestable.contract, event: "AttesterManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchAttesterManagerUpdated is a free log subscription operation binding the contract event 0x0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d3460219.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (_Attestable *AttestableFilterer) WatchAttesterManagerUpdated(opts *bind.WatchOpts, sink chan<- *AttestableAttesterManagerUpdated, previousAttesterManager []common.Address, newAttesterManager []common.Address) (event.Subscription, error) {

	var previousAttesterManagerRule []interface{}
	for _, previousAttesterManagerItem := range previousAttesterManager {
		previousAttesterManagerRule = append(previousAttesterManagerRule, previousAttesterManagerItem)
	}
	var newAttesterManagerRule []interface{}
	for _, newAttesterManagerItem := range newAttesterManager {
		newAttesterManagerRule = append(newAttesterManagerRule, newAttesterManagerItem)
	}

	logs, sub, err := _Attestable.contract.WatchLogs(opts, "AttesterManagerUpdated", previousAttesterManagerRule, newAttesterManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestableAttesterManagerUpdated)
				if err := _Attestable.contract.UnpackLog(event, "AttesterManagerUpdated", log); err != nil {
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

// ParseAttesterManagerUpdated is a log parse operation binding the contract event 0x0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d3460219.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (_Attestable *AttestableFilterer) ParseAttesterManagerUpdated(log types.Log) (*AttestableAttesterManagerUpdated, error) {
	event := new(AttestableAttesterManagerUpdated)
	if err := _Attestable.contract.UnpackLog(event, "AttesterManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestableOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Attestable contract.
type AttestableOwnershipTransferStartedIterator struct {
	Event *AttestableOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *AttestableOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestableOwnershipTransferStarted)
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
		it.Event = new(AttestableOwnershipTransferStarted)
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
func (it *AttestableOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestableOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Attestable contract.
type AttestableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Attestable *AttestableFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AttestableOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Attestable.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AttestableOwnershipTransferStartedIterator{contract: _Attestable.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Attestable *AttestableFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *AttestableOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Attestable.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestableOwnershipTransferStarted)
				if err := _Attestable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Attestable *AttestableFilterer) ParseOwnershipTransferStarted(log types.Log) (*AttestableOwnershipTransferStarted, error) {
	event := new(AttestableOwnershipTransferStarted)
	if err := _Attestable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Attestable contract.
type AttestableOwnershipTransferredIterator struct {
	Event *AttestableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AttestableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestableOwnershipTransferred)
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
		it.Event = new(AttestableOwnershipTransferred)
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
func (it *AttestableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestableOwnershipTransferred represents a OwnershipTransferred event raised by the Attestable contract.
type AttestableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Attestable *AttestableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AttestableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Attestable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AttestableOwnershipTransferredIterator{contract: _Attestable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Attestable *AttestableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AttestableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Attestable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestableOwnershipTransferred)
				if err := _Attestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Attestable *AttestableFilterer) ParseOwnershipTransferred(log types.Log) (*AttestableOwnershipTransferred, error) {
	event := new(AttestableOwnershipTransferred)
	if err := _Attestable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestableSignatureThresholdUpdatedIterator is returned from FilterSignatureThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignatureThresholdUpdated events raised by the Attestable contract.
type AttestableSignatureThresholdUpdatedIterator struct {
	Event *AttestableSignatureThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *AttestableSignatureThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestableSignatureThresholdUpdated)
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
		it.Event = new(AttestableSignatureThresholdUpdated)
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
func (it *AttestableSignatureThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestableSignatureThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestableSignatureThresholdUpdated represents a SignatureThresholdUpdated event raised by the Attestable contract.
type AttestableSignatureThresholdUpdated struct {
	OldSignatureThreshold *big.Int
	NewSignatureThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSignatureThresholdUpdated is a free log retrieval operation binding the contract event 0x149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (_Attestable *AttestableFilterer) FilterSignatureThresholdUpdated(opts *bind.FilterOpts) (*AttestableSignatureThresholdUpdatedIterator, error) {

	logs, sub, err := _Attestable.contract.FilterLogs(opts, "SignatureThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &AttestableSignatureThresholdUpdatedIterator{contract: _Attestable.contract, event: "SignatureThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignatureThresholdUpdated is a free log subscription operation binding the contract event 0x149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (_Attestable *AttestableFilterer) WatchSignatureThresholdUpdated(opts *bind.WatchOpts, sink chan<- *AttestableSignatureThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Attestable.contract.WatchLogs(opts, "SignatureThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestableSignatureThresholdUpdated)
				if err := _Attestable.contract.UnpackLog(event, "SignatureThresholdUpdated", log); err != nil {
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

// ParseSignatureThresholdUpdated is a log parse operation binding the contract event 0x149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (_Attestable *AttestableFilterer) ParseSignatureThresholdUpdated(log types.Log) (*AttestableSignatureThresholdUpdated, error) {
	event := new(AttestableSignatureThresholdUpdated)
	if err := _Attestable.contract.UnpackLog(event, "SignatureThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// ECDSAMetaData contains all meta data concerning the ECDSA contract.
var ECDSAMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200a7b516bf5eddb7944fe18e5df7b8aa4bb3648fc7ef05cc98b893dd073c4af2264736f6c63430007060033",
}

// ECDSAABI is the input ABI used to generate the binding from.
// Deprecated: Use ECDSAMetaData.ABI instead.
var ECDSAABI = ECDSAMetaData.ABI

// ECDSABin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ECDSAMetaData.Bin instead.
var ECDSABin = ECDSAMetaData.Bin

// DeployECDSA deploys a new Ethereum contract, binding an instance of ECDSA to it.
func DeployECDSA(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ECDSA, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ECDSABin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// ECDSA is an auto generated Go binding around an Ethereum contract.
type ECDSA struct {
	ECDSACaller     // Read-only binding to the contract
	ECDSATransactor // Write-only binding to the contract
	ECDSAFilterer   // Log filterer for contract events
}

// ECDSACaller is an auto generated read-only Go binding around an Ethereum contract.
type ECDSACaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSATransactor is an auto generated write-only Go binding around an Ethereum contract.
type ECDSATransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSAFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ECDSAFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ECDSASession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ECDSASession struct {
	Contract     *ECDSA            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSACallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ECDSACallerSession struct {
	Contract *ECDSACaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ECDSATransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ECDSATransactorSession struct {
	Contract     *ECDSATransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ECDSARaw is an auto generated low-level Go binding around an Ethereum contract.
type ECDSARaw struct {
	Contract *ECDSA // Generic contract binding to access the raw methods on
}

// ECDSACallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ECDSACallerRaw struct {
	Contract *ECDSACaller // Generic read-only contract binding to access the raw methods on
}

// ECDSATransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ECDSATransactorRaw struct {
	Contract *ECDSATransactor // Generic write-only contract binding to access the raw methods on
}

// NewECDSA creates a new instance of ECDSA, bound to a specific deployed contract.
func NewECDSA(address common.Address, backend bind.ContractBackend) (*ECDSA, error) {
	contract, err := bindECDSA(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ECDSA{ECDSACaller: ECDSACaller{contract: contract}, ECDSATransactor: ECDSATransactor{contract: contract}, ECDSAFilterer: ECDSAFilterer{contract: contract}}, nil
}

// NewECDSACaller creates a new read-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSACaller(address common.Address, caller bind.ContractCaller) (*ECDSACaller, error) {
	contract, err := bindECDSA(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSACaller{contract: contract}, nil
}

// NewECDSATransactor creates a new write-only instance of ECDSA, bound to a specific deployed contract.
func NewECDSATransactor(address common.Address, transactor bind.ContractTransactor) (*ECDSATransactor, error) {
	contract, err := bindECDSA(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ECDSATransactor{contract: contract}, nil
}

// NewECDSAFilterer creates a new log filterer instance of ECDSA, bound to a specific deployed contract.
func NewECDSAFilterer(address common.Address, filterer bind.ContractFilterer) (*ECDSAFilterer, error) {
	contract, err := bindECDSA(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ECDSAFilterer{contract: contract}, nil
}

// bindECDSA binds a generic wrapper to an already deployed contract.
func bindECDSA(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ECDSAMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSARaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.ECDSACaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSARaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSARaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.ECDSATransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ECDSA *ECDSACallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ECDSA.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ECDSA *ECDSATransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ECDSA *ECDSATransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ECDSA.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122083242fa2568cfd593e8dd04016b3ebeb588aa88031a6347dada2680fdcc4a96264736f6c63430007060033",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
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
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// IMessageHandlerMetaData contains all meta data concerning the IMessageHandler contract.
var IMessageHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"handleReceiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"96abeb70": "handleReceiveMessage(uint32,bytes32,bytes)",
	},
}

// IMessageHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageHandlerMetaData.ABI instead.
var IMessageHandlerABI = IMessageHandlerMetaData.ABI

// Deprecated: Use IMessageHandlerMetaData.Sigs instead.
// IMessageHandlerFuncSigs maps the 4-byte function signature to its string representation.
var IMessageHandlerFuncSigs = IMessageHandlerMetaData.Sigs

// IMessageHandler is an auto generated Go binding around an Ethereum contract.
type IMessageHandler struct {
	IMessageHandlerCaller     // Read-only binding to the contract
	IMessageHandlerTransactor // Write-only binding to the contract
	IMessageHandlerFilterer   // Log filterer for contract events
}

// IMessageHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageHandlerSession struct {
	Contract     *IMessageHandler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageHandlerCallerSession struct {
	Contract *IMessageHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IMessageHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageHandlerTransactorSession struct {
	Contract     *IMessageHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IMessageHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageHandlerRaw struct {
	Contract *IMessageHandler // Generic contract binding to access the raw methods on
}

// IMessageHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageHandlerCallerRaw struct {
	Contract *IMessageHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageHandlerTransactorRaw struct {
	Contract *IMessageHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageHandler creates a new instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandler(address common.Address, backend bind.ContractBackend) (*IMessageHandler, error) {
	contract, err := bindIMessageHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageHandler{IMessageHandlerCaller: IMessageHandlerCaller{contract: contract}, IMessageHandlerTransactor: IMessageHandlerTransactor{contract: contract}, IMessageHandlerFilterer: IMessageHandlerFilterer{contract: contract}}, nil
}

// NewIMessageHandlerCaller creates a new read-only instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerCaller(address common.Address, caller bind.ContractCaller) (*IMessageHandlerCaller, error) {
	contract, err := bindIMessageHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerCaller{contract: contract}, nil
}

// NewIMessageHandlerTransactor creates a new write-only instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageHandlerTransactor, error) {
	contract, err := bindIMessageHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerTransactor{contract: contract}, nil
}

// NewIMessageHandlerFilterer creates a new log filterer instance of IMessageHandler, bound to a specific deployed contract.
func NewIMessageHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageHandlerFilterer, error) {
	contract, err := bindIMessageHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageHandlerFilterer{contract: contract}, nil
}

// bindIMessageHandler binds a generic wrapper to an already deployed contract.
func bindIMessageHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageHandlerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageHandler *IMessageHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageHandler.Contract.IMessageHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageHandler *IMessageHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageHandler.Contract.IMessageHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageHandler *IMessageHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageHandler.Contract.IMessageHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageHandler *IMessageHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageHandler *IMessageHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageHandler *IMessageHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageHandler.Contract.contract.Transact(opts, method, params...)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 sourceDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_IMessageHandler *IMessageHandlerTransactor) HandleReceiveMessage(opts *bind.TransactOpts, sourceDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageHandler.contract.Transact(opts, "handleReceiveMessage", sourceDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 sourceDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_IMessageHandler *IMessageHandlerSession) HandleReceiveMessage(sourceDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageHandler.Contract.HandleReceiveMessage(&_IMessageHandler.TransactOpts, sourceDomain, sender, messageBody)
}

// HandleReceiveMessage is a paid mutator transaction binding the contract method 0x96abeb70.
//
// Solidity: function handleReceiveMessage(uint32 sourceDomain, bytes32 sender, bytes messageBody) returns(bool)
func (_IMessageHandler *IMessageHandlerTransactorSession) HandleReceiveMessage(sourceDomain uint32, sender [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageHandler.Contract.HandleReceiveMessage(&_IMessageHandler.TransactOpts, sourceDomain, sender, messageBody)
}

// IMessageTransmitterMetaData contains all meta data concerning the IMessageTransmitter contract.
var IMessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"originalAttestation\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newMessageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newDestinationCaller\",\"type\":\"bytes32\"}],\"name\":\"replaceMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"b857b774": "replaceMessage(bytes,bytes,bytes,bytes32)",
		"0ba469bc": "sendMessage(uint32,bytes32,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
}

// IMessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageTransmitterMetaData.ABI instead.
var IMessageTransmitterABI = IMessageTransmitterMetaData.ABI

// Deprecated: Use IMessageTransmitterMetaData.Sigs instead.
// IMessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var IMessageTransmitterFuncSigs = IMessageTransmitterMetaData.Sigs

// IMessageTransmitter is an auto generated Go binding around an Ethereum contract.
type IMessageTransmitter struct {
	IMessageTransmitterCaller     // Read-only binding to the contract
	IMessageTransmitterTransactor // Write-only binding to the contract
	IMessageTransmitterFilterer   // Log filterer for contract events
}

// IMessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageTransmitterSession struct {
	Contract     *IMessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageTransmitterCallerSession struct {
	Contract *IMessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageTransmitterTransactorSession struct {
	Contract     *IMessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageTransmitterRaw struct {
	Contract *IMessageTransmitter // Generic contract binding to access the raw methods on
}

// IMessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageTransmitterCallerRaw struct {
	Contract *IMessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageTransmitterTransactorRaw struct {
	Contract *IMessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessageTransmitter creates a new instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitter(address common.Address, backend bind.ContractBackend) (*IMessageTransmitter, error) {
	contract, err := bindIMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitter{IMessageTransmitterCaller: IMessageTransmitterCaller{contract: contract}, IMessageTransmitterTransactor: IMessageTransmitterTransactor{contract: contract}, IMessageTransmitterFilterer: IMessageTransmitterFilterer{contract: contract}}, nil
}

// NewIMessageTransmitterCaller creates a new read-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*IMessageTransmitterCaller, error) {
	contract, err := bindIMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterCaller{contract: contract}, nil
}

// NewIMessageTransmitterTransactor creates a new write-only instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageTransmitterTransactor, error) {
	contract, err := bindIMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterTransactor{contract: contract}, nil
}

// NewIMessageTransmitterFilterer creates a new log filterer instance of IMessageTransmitter, bound to a specific deployed contract.
func NewIMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageTransmitterFilterer, error) {
	contract, err := bindIMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageTransmitterFilterer{contract: contract}, nil
}

// bindIMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindIMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageTransmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.IMessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.IMessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessageTransmitter *IMessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessageTransmitter *IMessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReceiveMessage(&_IMessageTransmitter.TransactOpts, message, signature)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IMessageTransmitter *IMessageTransmitterTransactor) ReplaceMessage(opts *bind.TransactOpts, originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "replaceMessage", originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IMessageTransmitter *IMessageTransmitterSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReplaceMessage(&_IMessageTransmitter.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.ReplaceMessage(&_IMessageTransmitter.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactor) SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "sendMessage", destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessage(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessage(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IMessageTransmitter *IMessageTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IMessageTransmitter.Contract.SendMessageWithCaller(&_IMessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// IReceiverMetaData contains all meta data concerning the IReceiver contract.
var IReceiverMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"57ecfd28": "receiveMessage(bytes,bytes)",
	},
}

// IReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use IReceiverMetaData.ABI instead.
var IReceiverABI = IReceiverMetaData.ABI

// Deprecated: Use IReceiverMetaData.Sigs instead.
// IReceiverFuncSigs maps the 4-byte function signature to its string representation.
var IReceiverFuncSigs = IReceiverMetaData.Sigs

// IReceiver is an auto generated Go binding around an Ethereum contract.
type IReceiver struct {
	IReceiverCaller     // Read-only binding to the contract
	IReceiverTransactor // Write-only binding to the contract
	IReceiverFilterer   // Log filterer for contract events
}

// IReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type IReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IReceiverSession struct {
	Contract     *IReceiver        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IReceiverCallerSession struct {
	Contract *IReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IReceiverTransactorSession struct {
	Contract     *IReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type IReceiverRaw struct {
	Contract *IReceiver // Generic contract binding to access the raw methods on
}

// IReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IReceiverCallerRaw struct {
	Contract *IReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// IReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IReceiverTransactorRaw struct {
	Contract *IReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIReceiver creates a new instance of IReceiver, bound to a specific deployed contract.
func NewIReceiver(address common.Address, backend bind.ContractBackend) (*IReceiver, error) {
	contract, err := bindIReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IReceiver{IReceiverCaller: IReceiverCaller{contract: contract}, IReceiverTransactor: IReceiverTransactor{contract: contract}, IReceiverFilterer: IReceiverFilterer{contract: contract}}, nil
}

// NewIReceiverCaller creates a new read-only instance of IReceiver, bound to a specific deployed contract.
func NewIReceiverCaller(address common.Address, caller bind.ContractCaller) (*IReceiverCaller, error) {
	contract, err := bindIReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiverCaller{contract: contract}, nil
}

// NewIReceiverTransactor creates a new write-only instance of IReceiver, bound to a specific deployed contract.
func NewIReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*IReceiverTransactor, error) {
	contract, err := bindIReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IReceiverTransactor{contract: contract}, nil
}

// NewIReceiverFilterer creates a new log filterer instance of IReceiver, bound to a specific deployed contract.
func NewIReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*IReceiverFilterer, error) {
	contract, err := bindIReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IReceiverFilterer{contract: contract}, nil
}

// bindIReceiver binds a generic wrapper to an already deployed contract.
func bindIReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiver *IReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiver.Contract.IReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiver *IReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiver.Contract.IReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiver *IReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiver.Contract.IReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IReceiver *IReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IReceiver *IReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IReceiver *IReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IReceiver.Contract.contract.Transact(opts, method, params...)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IReceiver *IReceiverTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, signature []byte) (*types.Transaction, error) {
	return _IReceiver.contract.Transact(opts, "receiveMessage", message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IReceiver *IReceiverSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IReceiver.Contract.ReceiveMessage(&_IReceiver.TransactOpts, message, signature)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes signature) returns(bool success)
func (_IReceiver *IReceiverTransactorSession) ReceiveMessage(message []byte, signature []byte) (*types.Transaction, error) {
	return _IReceiver.Contract.ReceiveMessage(&_IReceiver.TransactOpts, message, signature)
}

// IRelayerMetaData contains all meta data concerning the IRelayer contract.
var IRelayerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"originalAttestation\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newMessageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newDestinationCaller\",\"type\":\"bytes32\"}],\"name\":\"replaceMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b857b774": "replaceMessage(bytes,bytes,bytes,bytes32)",
		"0ba469bc": "sendMessage(uint32,bytes32,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
	},
}

// IRelayerABI is the input ABI used to generate the binding from.
// Deprecated: Use IRelayerMetaData.ABI instead.
var IRelayerABI = IRelayerMetaData.ABI

// Deprecated: Use IRelayerMetaData.Sigs instead.
// IRelayerFuncSigs maps the 4-byte function signature to its string representation.
var IRelayerFuncSigs = IRelayerMetaData.Sigs

// IRelayer is an auto generated Go binding around an Ethereum contract.
type IRelayer struct {
	IRelayerCaller     // Read-only binding to the contract
	IRelayerTransactor // Write-only binding to the contract
	IRelayerFilterer   // Log filterer for contract events
}

// IRelayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRelayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRelayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRelayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRelayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRelayerSession struct {
	Contract     *IRelayer         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRelayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRelayerCallerSession struct {
	Contract *IRelayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IRelayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRelayerTransactorSession struct {
	Contract     *IRelayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IRelayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRelayerRaw struct {
	Contract *IRelayer // Generic contract binding to access the raw methods on
}

// IRelayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRelayerCallerRaw struct {
	Contract *IRelayerCaller // Generic read-only contract binding to access the raw methods on
}

// IRelayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRelayerTransactorRaw struct {
	Contract *IRelayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRelayer creates a new instance of IRelayer, bound to a specific deployed contract.
func NewIRelayer(address common.Address, backend bind.ContractBackend) (*IRelayer, error) {
	contract, err := bindIRelayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRelayer{IRelayerCaller: IRelayerCaller{contract: contract}, IRelayerTransactor: IRelayerTransactor{contract: contract}, IRelayerFilterer: IRelayerFilterer{contract: contract}}, nil
}

// NewIRelayerCaller creates a new read-only instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerCaller(address common.Address, caller bind.ContractCaller) (*IRelayerCaller, error) {
	contract, err := bindIRelayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerCaller{contract: contract}, nil
}

// NewIRelayerTransactor creates a new write-only instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerTransactor(address common.Address, transactor bind.ContractTransactor) (*IRelayerTransactor, error) {
	contract, err := bindIRelayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRelayerTransactor{contract: contract}, nil
}

// NewIRelayerFilterer creates a new log filterer instance of IRelayer, bound to a specific deployed contract.
func NewIRelayerFilterer(address common.Address, filterer bind.ContractFilterer) (*IRelayerFilterer, error) {
	contract, err := bindIRelayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRelayerFilterer{contract: contract}, nil
}

// bindIRelayer binds a generic wrapper to an already deployed contract.
func bindIRelayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IRelayerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayer *IRelayerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayer.Contract.IRelayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayer *IRelayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayer.Contract.IRelayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayer *IRelayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayer.Contract.IRelayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRelayer *IRelayerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRelayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRelayer *IRelayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRelayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRelayer *IRelayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRelayer.Contract.contract.Transact(opts, method, params...)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IRelayer *IRelayerTransactor) ReplaceMessage(opts *bind.TransactOpts, originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "replaceMessage", originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IRelayer *IRelayerSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IRelayer.Contract.ReplaceMessage(&_IRelayer.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_IRelayer *IRelayerTransactorSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _IRelayer.Contract.ReplaceMessage(&_IRelayer.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactor) SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "sendMessage", destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessage(&_IRelayer.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactorSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessage(&_IRelayer.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessageWithCaller(&_IRelayer.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_IRelayer *IRelayerTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _IRelayer.Contract.SendMessageWithCaller(&_IRelayer.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// MessageMetaData contains all meta data concerning the Message contract.
var MessageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addressToBytes32\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_buf\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"82c947b7": "addressToBytes32(address)",
		"5ced058e": "bytes32ToAddress(bytes32)",
	},
	Bin: "0x610119610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80635ced058e14604257806382c947b7146085575b600080fd5b605c60048036036020811015605657600080fd5b503560c7565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60b560048036036020811015609957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1660ca565b60408051918252519081900360200190f35b90565b73ffffffffffffffffffffffffffffffffffffffff169056fea264697066735822122016096abc2c134d24a08f7ea8ec23179b7ba94fbed82239438af7360b92c1c2b564736f6c63430007060033",
}

// MessageABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageMetaData.ABI instead.
var MessageABI = MessageMetaData.ABI

// Deprecated: Use MessageMetaData.Sigs instead.
// MessageFuncSigs maps the 4-byte function signature to its string representation.
var MessageFuncSigs = MessageMetaData.Sigs

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
	parsed, err := MessageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address addr) pure returns(bytes32)
func (_Message *MessageCaller) AddressToBytes32(opts *bind.CallOpts, addr common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Message.contract.Call(opts, &out, "addressToBytes32", addr)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address addr) pure returns(bytes32)
func (_Message *MessageSession) AddressToBytes32(addr common.Address) ([32]byte, error) {
	return _Message.Contract.AddressToBytes32(&_Message.CallOpts, addr)
}

// AddressToBytes32 is a free data retrieval call binding the contract method 0x82c947b7.
//
// Solidity: function addressToBytes32(address addr) pure returns(bytes32)
func (_Message *MessageCallerSession) AddressToBytes32(addr common.Address) ([32]byte, error) {
	return _Message.Contract.AddressToBytes32(&_Message.CallOpts, addr)
}

// Bytes32ToAddress is a free data retrieval call binding the contract method 0x5ced058e.
//
// Solidity: function bytes32ToAddress(bytes32 _buf) pure returns(address)
func (_Message *MessageCaller) Bytes32ToAddress(opts *bind.CallOpts, _buf [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Message.contract.Call(opts, &out, "bytes32ToAddress", _buf)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bytes32ToAddress is a free data retrieval call binding the contract method 0x5ced058e.
//
// Solidity: function bytes32ToAddress(bytes32 _buf) pure returns(address)
func (_Message *MessageSession) Bytes32ToAddress(_buf [32]byte) (common.Address, error) {
	return _Message.Contract.Bytes32ToAddress(&_Message.CallOpts, _buf)
}

// Bytes32ToAddress is a free data retrieval call binding the contract method 0x5ced058e.
//
// Solidity: function bytes32ToAddress(bytes32 _buf) pure returns(address)
func (_Message *MessageCallerSession) Bytes32ToAddress(_buf [32]byte) (common.Address, error) {
	return _Message.Contract.Bytes32ToAddress(&_Message.CallOpts, _buf)
}

// MessageTransmitterMetaData contains all meta data concerning the MessageTransmitter contract.
var MessageTransmitterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_localDomain\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"_attester\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_maxMessageBodySize\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_version\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"AttesterEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousAttesterManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"AttesterManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxMessageBodySize\",\"type\":\"uint256\"}],\"name\":\"MaxMessageBodySizeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"sourceDomain\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"MessageReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"}],\"name\":\"MessageSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldSignatureThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"SignatureThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attesterManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"disableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttester\",\"type\":\"address\"}],\"name\":\"enableAttester\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getEnabledAttester\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumEnabledAttesters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"attester\",\"type\":\"address\"}],\"name\":\"isEnabledAttester\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"localDomain\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxMessageBodySize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextAvailableNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"message\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"attestation\",\"type\":\"bytes\"}],\"name\":\"receiveMessage\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"originalMessage\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"originalAttestation\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"newMessageBody\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"newDestinationCaller\",\"type\":\"bytes32\"}],\"name\":\"replaceMessage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessage\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"destinationDomain\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"destinationCaller\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"messageBody\",\"type\":\"bytes\"}],\"name\":\"sendMessageWithCaller\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMaxMessageBodySize\",\"type\":\"uint256\"}],\"name\":\"setMaxMessageBodySize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSignatureThreshold\",\"type\":\"uint256\"}],\"name\":\"setSignatureThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAttesterManager\",\"type\":\"address\"}],\"name\":\"updateAttesterManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"usedNonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"9b0d94b7": "attesterManager()",
		"2d025080": "disableAttester(address)",
		"fae36879": "enableAttester(address)",
		"beb673d8": "getEnabledAttester(uint256)",
		"51079a53": "getNumEnabledAttesters()",
		"7af82f60": "isEnabledAttester(address)",
		"8d3638f4": "localDomain()",
		"af47b9bb": "maxMessageBodySize()",
		"8371744e": "nextAvailableNonce()",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"9fd0506d": "pauser()",
		"e30c3978": "pendingOwner()",
		"57ecfd28": "receiveMessage(bytes,bytes)",
		"b857b774": "replaceMessage(bytes,bytes,bytes,bytes32)",
		"b2118a8d": "rescueERC20(address,address,uint256)",
		"38a63183": "rescuer()",
		"0ba469bc": "sendMessage(uint32,bytes32,bytes)",
		"f7259a75": "sendMessageWithCaller(uint32,bytes32,bytes32,bytes)",
		"92492c68": "setMaxMessageBodySize(uint256)",
		"bbde5374": "setSignatureThreshold(uint256)",
		"a82f2e26": "signatureThreshold()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"de7769d4": "updateAttesterManager(address)",
		"554bab3c": "updatePauser(address)",
		"2ab60045": "updateRescuer(address)",
		"feb61724": "usedNonces(bytes32)",
		"54fd4d50": "version()",
	},
	Bin: "0x60c06040526002805460ff60a01b191690553480156200001e57600080fd5b5060405162003cb738038062003cb7833981810160405260808110156200004457600080fd5b508051602082015160408301516060909301519192909182620000706200006a620000bb565b620000bf565b6200007b33620000e9565b60016004556200008b816200010b565b506001600160e01b031960e094851b811660805263ffffffff9290921660085590921b90911660a0525062000342565b3390565b600180546001600160a01b0319169055620000e6816200026b602090811b62001dbb17901c565b50565b600780546001600160a01b0319166001600160a01b0392909216919091179055565b6007546001600160a01b031633146200016b576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b6001600160a01b038116620001c7576040805162461bcd60e51b815260206004820152601c60248201527f4e6577206174746573746572206d757374206265206e6f6e7a65726f00000000604482015290519081900360640190fd5b620001e2816005620002bb60201b62001e231790919060201c565b62000234576040805162461bcd60e51b815260206004820152601860248201527f417474657374657220616c726561647920656e61626c65640000000000000000604482015290519081900360640190fd5b6040516001600160a01b038216907f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f90600090a250565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000620002d2836001600160a01b038416620002db565b90505b92915050565b6000620002e983836200032a565b6200032157508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155620002d5565b506000620002d5565b60009081526001919091016020526040902054151590565b60805160e01c60a05160e01c6139326200038560003980610bac5280610efb5280611f33525080610d605280611481528061173d5280611f5452506139326000f3fe608060405234801561001057600080fd5b50600436106101e55760003560e01c80638da5cb5b1161010f578063bbde5374116100a2578063f2fde38b11610071578063f2fde38b1461067f578063f7259a75146106a5578063fae368791461072e578063feb6172414610754576101e5565b8063bbde537414610617578063beb673d814610634578063de7769d414610651578063e30c397814610677576101e5565b8063a82f2e26116100de578063a82f2e26146104bd578063af47b9bb146104c5578063b2118a8d146104cd578063b857b77414610503576101e5565b80638da5cb5b1461048857806392492c68146104905780639b0d94b7146104ad5780639fd0506d146104b5576101e5565b8063554bab3c116101875780637af82f60116101565780637af82f601461044a5780638371744e146104705780638456cb59146104785780638d3638f414610480576101e5565b8063554bab3c1461033e57806357ecfd28146103645780635c975abb1461043a57806379ba509714610442576101e5565b806338a63183116101c357806338a63183146102d75780633f4ba83a146102fb57806351079a531461030357806354fd4d501461031d576101e5565b80630ba469bc146101ea5780632ab60045146102895780632d025080146102b1575b600080fd5b61026c6004803603606081101561020057600080fd5b63ffffffff8235169160208101359181019060608101604082013564010000000081111561022d57600080fd5b82018360208201111561023f57600080fd5b8035906020019184600183028401116401000000008311171561026157600080fd5b509092509050610771565b6040805167ffffffffffffffff9092168252519081900360200190f35b6102af6004803603602081101561029f57600080fd5b50356001600160a01b0316610895565b005b6102af600480360360208110156102c757600080fd5b50356001600160a01b0316610944565b6102df610aee565b604080516001600160a01b039092168252519081900360200190f35b6102af610afd565b61030b610b99565b60408051918252519081900360200190f35b610325610baa565b6040805163ffffffff9092168252519081900360200190f35b6102af6004803603602081101561035457600080fd5b50356001600160a01b0316610bce565b6104266004803603604081101561037a57600080fd5b81019060208101813564010000000081111561039557600080fd5b8201836020820111156103a757600080fd5b803590602001918460018302840111640100000000831117156103c957600080fd5b9193909290916020810190356401000000008111156103e757600080fd5b8201836020820111156103f957600080fd5b8035906020019184600183028401116401000000008311171561041b57600080fd5b509092509050610c83565b604080519115158252519081900360200190f35b610426611317565b6102af611338565b6104266004803603602081101561046057600080fd5b50356001600160a01b03166113a7565b61026c6113bc565b6102af6113cc565b61032561147f565b6102df6114a3565b6102af600480360360208110156104a657600080fd5b50356114b2565b6102df6114f5565b6102df611504565b61030b611513565b61030b611519565b6102af600480360360608110156104e357600080fd5b506001600160a01b0381358116916020810135909116906040013561151f565b6102af6004803603608081101561051957600080fd5b81019060208101813564010000000081111561053457600080fd5b82018360208201111561054657600080fd5b8035906020019184600183028401116401000000008311171561056857600080fd5b91939092909160208101903564010000000081111561058657600080fd5b82018360208201111561059857600080fd5b803590602001918460018302840111640100000000831117156105ba57600080fd5b9193909290916020810190356401000000008111156105d857600080fd5b8201836020820111156105ea57600080fd5b8035906020019184600183028401116401000000008311171561060c57600080fd5b919350915035611581565b6102af6004803603602081101561062d57600080fd5b50356117fc565b6102df6004803603602081101561064a57600080fd5b50356119a9565b6102af6004803603602081101561066757600080fd5b50356001600160a01b03166119b6565b6102df611a72565b6102af6004803603602081101561069557600080fd5b50356001600160a01b0316611a81565b61026c600480360360808110156106bb57600080fd5b63ffffffff82351691602081013591604082013591908101906080810160608201356401000000008111156106ef57600080fd5b82018360208201111561070157600080fd5b8035906020019184600183028401116401000000008311171561072357600080fd5b509092509050611aff565b6102af6004803603602081101561074457600080fd5b50356001600160a01b0316611c5c565b61030b6004803603602081101561076a57600080fd5b5035611da9565b60025460009074010000000000000000000000000000000000000000900460ff16156107e4576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6000806107ef611e41565b9050600073__$287656b8b950fc4c24b77afe782f2d94b1$__6382c947b7336040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561084b57600080fd5b505af415801561085f573d6000803e3d6000fd5b505050506040513d602081101561087557600080fd5b5051905061088888888584868b8b611e83565b509150505b949350505050565b61089d612059565b6001600160a01b0381166108e25760405162461bcd60e51b815260040180806020018281038252602a8152602001806136c9602a913960400191505060405180910390fd5b600380547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383169081179091556040517fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a90600090a250565b6007546001600160a01b031633146109a3576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b60006109ad610b99565b905060018111610a04576040805162461bcd60e51b815260206004820152601960248201527f546f6f2066657720656e61626c65642061747465737465727300000000000000604482015290519081900360640190fd5b6004548111610a5a576040805162461bcd60e51b815260206004820152601e60248201527f5369676e6174757265207468726573686f6c6420697320746f6f206c6f770000604482015290519081900360640190fd5b610a656005836120cf565b610ab6576040805162461bcd60e51b815260206004820152601960248201527f417474657374657220616c72656164792064697361626c656400000000000000604482015290519081900360640190fd5b6040516001600160a01b038316907f78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc90600090a25050565b6003546001600160a01b031690565b6002546001600160a01b03163314610b465760405162461bcd60e51b81526004018080602001828103825260228152602001806138296022913960400191505060405180910390fd5b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b6000610ba560056120e4565b905090565b7f000000000000000000000000000000000000000000000000000000000000000081565b610bd6612059565b6001600160a01b038116610c1b5760405162461bcd60e51b81526004018080602001828103825260288152602001806136566028913960400191505060405180910390fd5b600280547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0383811691909117918290556040519116907fb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a60490600090a250565b60025460009074010000000000000000000000000000000000000000900460ff1615610cf6576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b610d02858585856120ef565b6000610d48600087878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506122a49050565b9050610d5962ffffff1982166122c8565b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610d8f62ffffff19831661239a565b63ffffffff1614610de7576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c69642064657374696e6174696f6e20646f6d61696e000000000000604482015290519081900360640190fd5b6000610df862ffffff1983166123af565b14610ef457604080517f82c947b7000000000000000000000000000000000000000000000000000000008152336004820152905173__$287656b8b950fc4c24b77afe782f2d94b1$__916382c947b7916024808301926020929190829003018186803b158015610e6757600080fd5b505af4158015610e7b573d6000803e3d6000fd5b505050506040513d6020811015610e9157600080fd5b5051610ea262ffffff1983166123af565b14610ef4576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c69642063616c6c657220666f72206d657373616765000000000000604482015290519081900360640190fd5b63ffffffff7f000000000000000000000000000000000000000000000000000000000000000016610f2a62ffffff1983166123c4565b63ffffffff1614610f82576040805162461bcd60e51b815260206004820152601760248201527f496e76616c6964206d6573736167652076657273696f6e000000000000000000604482015290519081900360640190fd5b6000610f9362ffffff1983166123d8565b90506000610fa662ffffff1984166123ec565b90506000610fb48383612401565b6000818152600a602052604090205490915015611018576040805162461bcd60e51b815260206004820152601260248201527f4e6f6e636520616c726561647920757365640000000000000000000000000000604482015290519081900360640190fd5b6000818152600a602052604081206001905561103962ffffff19861661247c565b9050600061105a61104f62ffffff198816612491565b62ffffff19166124c8565b905073__$287656b8b950fc4c24b77afe782f2d94b1$__635ced058e61108562ffffff19891661250c565b6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156110b957600080fd5b505af41580156110cd573d6000803e3d6000fd5b505050506040513d60208110156110e357600080fd5b50516040517f96abeb7000000000000000000000000000000000000000000000000000000000815263ffffffff871660048201908152602482018590526060604483019081528451606484015284516001600160a01b03909416936396abeb70938a938893889391929091608490910190602085019080838360005b8381101561117757818101518382015260200161115f565b50505050905090810190601f1680156111a45780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b1580156111c557600080fd5b505af11580156111d9573d6000803e3d6000fd5b505050506040513d60208110156111ef57600080fd5b5051611242576040805162461bcd60e51b815260206004820152601d60248201527f68616e646c65526563656976654d6573736167652829206661696c6564000000604482015290519081900360640190fd5b8367ffffffffffffffff16336001600160a01b03167f58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d878585604051808463ffffffff16815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156112ca5781810151838201526020016112b2565b50505050905090810190601f1680156112f75780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a35060019a9950505050505050505050565b60025474010000000000000000000000000000000000000000900460ff1681565b6000611342612521565b9050806001600160a01b0316611356611a72565b6001600160a01b03161461139b5760405162461bcd60e51b81526004018080602001828103825260298152602001806136a06029913960400191505060405180910390fd5b6113a481612525565b50565b60006113b4600583612556565b90505b919050565b60095467ffffffffffffffff1681565b6002546001600160a01b031633146114155760405162461bcd60e51b81526004018080602001828103825260228152602001806138296022913960400191505060405180910390fd5b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b7f000000000000000000000000000000000000000000000000000000000000000081565b6000546001600160a01b031690565b6114ba612059565b60088190556040805182815290517fb13bf6bebed03d1b318e3ea32e4b2a3ad9f5e2312cdf340a2f4bbfaee39f928d9181900360200190a150565b6007546001600160a01b031690565b6002546001600160a01b031690565b60045481565b60085481565b6003546001600160a01b031633146115685760405162461bcd60e51b81526004018080602001828103825260248152602001806137cb6024913960400191505060405180910390fd5b61157c6001600160a01b038416838361256b565b505050565b60025474010000000000000000000000000000000000000000900460ff16156115f1576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b6115fd878787876120ef565b6000611643600089898080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092939250506122a49050565b905061165462ffffff1982166122c8565b600061166562ffffff19831661247c565b905073__$287656b8b950fc4c24b77afe782f2d94b1$__635ced058e826040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156116b657600080fd5b505af41580156116ca573d6000803e3d6000fd5b505050506040513d60208110156116e057600080fd5b50516001600160a01b031633146117285760405162461bcd60e51b81526004018080602001828103825260218152602001806137676021913960400191505060405180910390fd5b600061173962ffffff1984166123d8565b90507f000000000000000000000000000000000000000000000000000000000000000063ffffffff168163ffffffff16146117a55760405162461bcd60e51b815260040180806020018281038252602c81526020018061373b602c913960400191505060405180910390fd5b60006117b662ffffff19851661239a565b905060006117c962ffffff19861661250c565b905060006117dc62ffffff1987166123ec565b90506117ed83838988858e8e611e83565b50505050505050505050505050565b6007546001600160a01b0316331461185b576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b806118ad576040805162461bcd60e51b815260206004820152601b60248201527f496e76616c6964207369676e6174757265207468726573686f6c640000000000604482015290519081900360640190fd5b6118b760056120e4565b81111561190b576040805162461bcd60e51b815260206004820181905260248201527f4e6577207369676e6174757265207468726573686f6c6420746f6f2068696768604482015290519081900360640190fd5b600454811415611962576040805162461bcd60e51b815260206004820152601f60248201527f5369676e6174757265207468726573686f6c6420616c72656164792073657400604482015290519081900360640190fd5b6004805490829055604080518281526020810184905281517f149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede929181900390910190a15050565b60006113b46005836125eb565b6119be612059565b6001600160a01b038116611a19576040805162461bcd60e51b815260206004820181905260248201527f496e76616c6964206174746573746572206d616e616765722061646472657373604482015290519081900360640190fd5b6007546001600160a01b0316611a2e826125f7565b816001600160a01b0316816001600160a01b03167f0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d346021960405160405180910390a35050565b6001546001600160a01b031690565b611a89612059565b600180547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038316908117909155611ac76114a3565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b60025460009074010000000000000000000000000000000000000000900460ff1615611b72576040805162461bcd60e51b815260206004820152601060248201527f5061757361626c653a2070617573656400000000000000000000000000000000604482015290519081900360640190fd5b83611bae5760405162461bcd60e51b815260040180806020018281038252602281526020018061367e6022913960400191505060405180910390fd5b6000611bb8611e41565b9050600073__$287656b8b950fc4c24b77afe782f2d94b1$__6382c947b7336040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611c1457600080fd5b505af4158015611c28573d6000803e3d6000fd5b505050506040513d6020811015611c3e57600080fd5b50519050611c5188888884868a8a611e83565b509695505050505050565b6007546001600160a01b03163314611cbb576040805162461bcd60e51b815260206004820152601b60248201527f43616c6c6572206e6f74206174746573746572206d616e616765720000000000604482015290519081900360640190fd5b6001600160a01b038116611d16576040805162461bcd60e51b815260206004820152601c60248201527f4e6577206174746573746572206d757374206265206e6f6e7a65726f00000000604482015290519081900360640190fd5b611d21600582611e23565b611d72576040805162461bcd60e51b815260206004820152601860248201527f417474657374657220616c726561647920656e61626c65640000000000000000604482015290519081900360640190fd5b6040516001600160a01b038216907f5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f90600090a250565b600a6020526000908152604090205481565b600080546001600160a01b038381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000611e38836001600160a01b038416612631565b90505b92915050565b600980547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000008116600167ffffffffffffffff9283169081019092161790915590565b600854811115611eda576040805162461bcd60e51b815260206004820152601d60248201527f4d65737361676520626f64792065786365656473206d61782073697a65000000604482015290519081900360640190fd5b85611f2c576040805162461bcd60e51b815260206004820152601960248201527f526563697069656e74206d757374206265206e6f6e7a65726f00000000000000604482015290519081900360640190fd5b6000611fb47f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008a87898c8c8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061267b92505050565b90507f8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036816040518080602001828103825283818151815260200191508051906020019080838360005b83811015612015578181015183820152602001611ffd565b50505050905090810190601f1680156120425780820380516001836020036101000a031916815260200191505b509250505060405180910390a15050505050505050565b612061612521565b6001600160a01b03166120726114a3565b6001600160a01b0316146120cd576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b565b6000611e38836001600160a01b038416612770565b60006113b482612854565b6004546041028114612148576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c6964206174746573746174696f6e206c656e677468000000000000604482015290519081900360640190fd5b60008085856040518083838082843760405192018290039091209450600093505050505b60045481101561229b57600061218b604183810290810190878961360b565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201829052509394506121ce9250869150849050612858565b9050846001600160a01b0316816001600160a01b031611612236576040805162461bcd60e51b815260206004820152601f60248201527f496e76616c6964207369676e6174757265206f72646572206f72206475706500604482015290519081900360640190fd5b61223f816113a7565b612290576040805162461bcd60e51b815260206004820152601f60248201527f496e76616c6964207369676e61747572653a206e6f7420617474657374657200604482015290519081900360640190fd5b93505060010161216c565b50505050505050565b8151600090602084016122bf64ffffffffff85168284612864565b95945050505050565b6122d762ffffff1982166128a9565b612328576040805162461bcd60e51b815260206004820152601160248201527f4d616c666f726d6564206d657373616765000000000000000000000000000000604482015290519081900360640190fd5b607461233962ffffff1983166128e6565b6bffffffffffffffffffffffff1610156113a4576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c6964206d6573736167653a20746f6f2073686f7274000000000000604482015290519081900360640190fd5b60006113b462ffffff198316600860046128fa565b60006113b462ffffff1983166054602061291b565b60006113b462ffffff1983168260046128fa565b60006113b462ffffff1983166004806128fa565b60006113b462ffffff198316600c60086128fa565b6040805160e09390931b7fffffffff000000000000000000000000000000000000000000000000000000001660208085019190915260c09290921b7fffffffffffffffff0000000000000000000000000000000000000000000000001660248401528051808403600c018152602c9093019052815191012090565b60006113b462ffffff1983166014602061291b565b60006113b46074806124a862ffffff1986166128e6565b62ffffff19861692916bffffffffffffffffffffffff9103166000612a92565b60606000806124d6846128e6565b6bffffffffffffffffffffffff16905060405191508192506124fb8483602001612b06565b508181016020016040529052919050565b60006113b462ffffff1983166034602061291b565b3390565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556113a481611dbb565b6000611e38836001600160a01b038416612bfe565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb0000000000000000000000000000000000000000000000000000000017905261157c908490612c16565b6000611e388383612cc7565b600780547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b600061263d8383612bfe565b61267357508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155611e3b565b506000611e3b565b60608888888888888888604051602001808963ffffffff1660e01b81526004018863ffffffff1660e01b81526004018763ffffffff1660e01b81526004018667ffffffffffffffff1660c01b815260080185815260200184815260200183815260200182805190602001908083835b6020831061272757805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090920191602091820191016126ea565b6001836020036101000a03801982511681845116808217855250505050505090500198505050505050505050604051602081830303815290604052905098975050505050505050565b6000818152600183016020526040812054801561284a5783547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff80830191908101906000908790839081106127c157fe5b90600052602060002001549050808760000184815481106127de57fe5b60009182526020808320909101929092558281526001898101909252604090209084019055865487908061280e57fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050611e3b565b6000915050611e3b565b5490565b6000611e388383612d2b565b6000806128718484612da1565b9050604051811115612881575060005b806128935762ffffff199150506128a2565b61289e858585612dfb565b9150505b9392505050565b60006128b482612e0e565b64ffffffffff1664ffffffffff14156128cf575060006113b7565b60006128da83612e14565b60405110199392505050565b60181c6bffffffffffffffffffffffff1690565b60008160200360080260ff1661291185858561291b565b901c949350505050565b600060ff821661292d575060006128a2565b612936846128e6565b6bffffffffffffffffffffffff166129518460ff8516612da1565b1115612a165761299261296385612e3e565b6bffffffffffffffffffffffff1661297a866128e6565b6bffffffffffffffffffffffff16858560ff16612e52565b60405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156129db5781810151838201526020016129c3565b50505050905090810190601f168015612a085780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b60208260ff161115612a595760405162461bcd60e51b815260040180806020018281038252603a8152602001806137ef603a913960400191505060405180910390fd5b600882026000612a6886612e3e565b6bffffffffffffffffffffffff1690506000612a8383612fad565b91909501511695945050505050565b600080612a9e86612e3e565b6bffffffffffffffffffffffff169050612ab786612e14565b612acb85612ac58489612da1565b90612da1565b1115612ade5762ffffff1991505061088d565b612ae88186612da1565b9050612afc8364ffffffffff168286612864565b9695505050505050565b6000612b1183612ff6565b612b4c5760405162461bcd60e51b81526004018080602001828103825260288152602001806138756028913960400191505060405180910390fd5b612b55836128a9565b612b905760405162461bcd60e51b815260040180806020018281038252602b81526020018061389d602b913960400191505060405180910390fd5b6000612b9b846128e6565b6bffffffffffffffffffffffff1690506000612bb685612e3e565b6bffffffffffffffffffffffff1690506000604051905084811115612bdb5760206060fd5b8285848460045afa50612afc612bf087612e0e565b64ffffffffff168685612dfb565b60009081526001919091016020526040902054151590565b6000612c6b826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166130089092919063ffffffff16565b80519091501561157c57808060200190516020811015612c8a57600080fd5b505161157c5760405162461bcd60e51b815260040180806020018281038252602a81526020018061384b602a913960400191505060405180910390fd5b81546000908210612d095760405162461bcd60e51b81526004018080602001828103825260228152602001806136346022913960400191505060405180910390fd5b826000018281548110612d1857fe5b9060005260206000200154905092915050565b60008151604114612d83576040805162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015290519081900360640190fd5b60208201516040830151606084015160001a612afc86828585613017565b600082820183811015611e38576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b606092831b9190911790911b1760181b90565b60d81c90565b6000612e1f826128e6565b612e2883612e3e565b016bffffffffffffffffffffffff169050919050565b60781c6bffffffffffffffffffffffff1690565b60606000612e5f866131aa565b9150506000612e6d866131aa565b9150506000612e7b866131aa565b9150506000612e89866131aa565b9150508383838360405160200180806138c8603591397fffffffffffff000000000000000000000000000000000000000000000000000060d087811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b84015286901b16604a820152605001602161378882397fffffffffffff000000000000000000000000000000000000000000000000000060d094851b811660218301527f2077697468206c656e677468203078000000000000000000000000000000000060278301529290931b9091166036830152507f2e00000000000000000000000000000000000000000000000000000000000000603c82015260408051601d818403018152603d90920190529b9a5050505050505050505050565b7f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9091011d90565b60006130018261327e565b1592915050565b606061088d848460008561328a565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08211156130785760405162461bcd60e51b81526004018080602001828103825260228152602001806136f36022913960400191505060405180910390fd5b8360ff16601b148061308d57508360ff16601c145b6130c85760405162461bcd60e51b81526004018080602001828103825260228152602001806137a96022913960400191505060405180910390fd5b600060018686868660405160008152602001604052604051808581526020018460ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015613124573d6000803e3d6000fd5b50506040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe001519150506001600160a01b0381166122bf576040805162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015290519081900360640190fd5b600080601f5b600f8160ff1611156132125760ff600882021684901c6131cf81613403565b61ffff16841793508160ff166010146131ea57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016131b0565b50600f5b60ff8160ff1610156132785760ff600882021684901c61323581613403565b61ffff16831792508160ff1660001461325057601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01613216565b50915091565b62ffffff199081161490565b6060824710156132cb5760405162461bcd60e51b81526004018080602001828103825260268152602001806137156026913960400191505060405180910390fd5b6132d485613433565b613325576040805162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b600080866001600160a01b031685876040518082805190602001908083835b6020831061338157805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101613344565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146133e3576040519150601f19603f3d011682016040523d82523d6000602084013e6133e8565b606091505b50915091506133f8828286613439565b979650505050505050565b600061341560048360ff16901c61349f565b60ff161760081b62ffff001661342a8261349f565b60ff1617919050565b3b151590565b606083156134485750816128a2565b8251156134585782518084602001fd5b60405162461bcd60e51b81526020600482018181528451602484015284518593919283926044019190850190808383600083156129db5781810151838201526020016129c3565b600060f08083179060ff821614156134bb5760309150506113b7565b8060ff1660f114156134d15760319150506113b7565b8060ff1660f214156134e75760329150506113b7565b8060ff1660f314156134fd5760339150506113b7565b8060ff1660f414156135135760349150506113b7565b8060ff1660f514156135295760359150506113b7565b8060ff1660f6141561353f5760369150506113b7565b8060ff1660f714156135555760379150506113b7565b8060ff1660f8141561356b5760389150506113b7565b8060ff1660f914156135815760399150506113b7565b8060ff1660fa14156135975760619150506113b7565b8060ff1660fb14156135ad5760629150506113b7565b8060ff1660fc14156135c35760639150506113b7565b8060ff1660fd14156135d95760649150506113b7565b8060ff1660fe14156135ef5760659150506113b7565b8060ff1660ff14156136055760669150506113b7565b50919050565b6000808585111561361a578182fd5b83861115613626578182fd5b505082019391909203915056fe456e756d657261626c655365743a20696e646578206f7574206f6620626f756e64735061757361626c653a206e65772070617573657220697320746865207a65726f206164647265737344657374696e6174696f6e2063616c6c6572206d757374206265206e6f6e7a65726f4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f206164647265737345434453413a20696e76616c6964207369676e6174757265202773272076616c7565416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c4d657373616765206e6f74206f726967696e616c6c792073656e742066726f6d207468697320646f6d61696e53656e646572206e6f74207065726d697474656420746f20757365206e6f6e63652e20417474656d7074656420746f20696e646578206174206f666673657420307845434453413a20696e76616c6964207369676e6174757265202776272076616c7565526573637561626c653a2063616c6c6572206973206e6f7420746865207265736375657254797065644d656d566965772f696e646578202d20417474656d7074656420746f20696e646578206d6f7265207468616e2033322062797465735061757361626c653a2063616c6c6572206973206e6f7420746865207061757365725361666545524332303a204552433230206f7065726174696f6e20646964206e6f74207375636365656454797065644d656d566965772f636f7079546f202d204e756c6c20706f696e74657220646572656654797065644d656d566965772f636f7079546f202d20496e76616c696420706f696e74657220646572656654797065644d656d566965772f696e646578202d204f76657272616e2074686520766965772e20536c696365206973206174203078a264697066735822122080a00769da33d6f2913e3780432f8f71ea1b1e56d8fe2a3b7b18047055d2b51764736f6c63430007060033",
}

// MessageTransmitterABI is the input ABI used to generate the binding from.
// Deprecated: Use MessageTransmitterMetaData.ABI instead.
var MessageTransmitterABI = MessageTransmitterMetaData.ABI

// Deprecated: Use MessageTransmitterMetaData.Sigs instead.
// MessageTransmitterFuncSigs maps the 4-byte function signature to its string representation.
var MessageTransmitterFuncSigs = MessageTransmitterMetaData.Sigs

// MessageTransmitterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MessageTransmitterMetaData.Bin instead.
var MessageTransmitterBin = MessageTransmitterMetaData.Bin

// DeployMessageTransmitter deploys a new Ethereum contract, binding an instance of MessageTransmitter to it.
func DeployMessageTransmitter(auth *bind.TransactOpts, backend bind.ContractBackend, _localDomain uint32, _attester common.Address, _maxMessageBodySize uint32, _version uint32) (common.Address, *types.Transaction, *MessageTransmitter, error) {
	parsed, err := MessageTransmitterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	messageAddr, _, _, _ := DeployMessage(auth, backend)
	MessageTransmitterBin = strings.ReplaceAll(MessageTransmitterBin, "__$287656b8b950fc4c24b77afe782f2d94b1$__", messageAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MessageTransmitterBin), backend, _localDomain, _attester, _maxMessageBodySize, _version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MessageTransmitter{MessageTransmitterCaller: MessageTransmitterCaller{contract: contract}, MessageTransmitterTransactor: MessageTransmitterTransactor{contract: contract}, MessageTransmitterFilterer: MessageTransmitterFilterer{contract: contract}}, nil
}

// MessageTransmitter is an auto generated Go binding around an Ethereum contract.
type MessageTransmitter struct {
	MessageTransmitterCaller     // Read-only binding to the contract
	MessageTransmitterTransactor // Write-only binding to the contract
	MessageTransmitterFilterer   // Log filterer for contract events
}

// MessageTransmitterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MessageTransmitterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MessageTransmitterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MessageTransmitterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MessageTransmitterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MessageTransmitterSession struct {
	Contract     *MessageTransmitter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MessageTransmitterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MessageTransmitterCallerSession struct {
	Contract *MessageTransmitterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MessageTransmitterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MessageTransmitterTransactorSession struct {
	Contract     *MessageTransmitterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MessageTransmitterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MessageTransmitterRaw struct {
	Contract *MessageTransmitter // Generic contract binding to access the raw methods on
}

// MessageTransmitterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MessageTransmitterCallerRaw struct {
	Contract *MessageTransmitterCaller // Generic read-only contract binding to access the raw methods on
}

// MessageTransmitterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MessageTransmitterTransactorRaw struct {
	Contract *MessageTransmitterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMessageTransmitter creates a new instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitter(address common.Address, backend bind.ContractBackend) (*MessageTransmitter, error) {
	contract, err := bindMessageTransmitter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitter{MessageTransmitterCaller: MessageTransmitterCaller{contract: contract}, MessageTransmitterTransactor: MessageTransmitterTransactor{contract: contract}, MessageTransmitterFilterer: MessageTransmitterFilterer{contract: contract}}, nil
}

// NewMessageTransmitterCaller creates a new read-only instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitterCaller(address common.Address, caller bind.ContractCaller) (*MessageTransmitterCaller, error) {
	contract, err := bindMessageTransmitter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterCaller{contract: contract}, nil
}

// NewMessageTransmitterTransactor creates a new write-only instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitterTransactor(address common.Address, transactor bind.ContractTransactor) (*MessageTransmitterTransactor, error) {
	contract, err := bindMessageTransmitter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterTransactor{contract: contract}, nil
}

// NewMessageTransmitterFilterer creates a new log filterer instance of MessageTransmitter, bound to a specific deployed contract.
func NewMessageTransmitterFilterer(address common.Address, filterer bind.ContractFilterer) (*MessageTransmitterFilterer, error) {
	contract, err := bindMessageTransmitter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterFilterer{contract: contract}, nil
}

// bindMessageTransmitter binds a generic wrapper to an already deployed contract.
func bindMessageTransmitter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MessageTransmitterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTransmitter *MessageTransmitterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTransmitter.Contract.MessageTransmitterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTransmitter *MessageTransmitterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.MessageTransmitterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTransmitter *MessageTransmitterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.MessageTransmitterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MessageTransmitter *MessageTransmitterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MessageTransmitter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MessageTransmitter *MessageTransmitterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MessageTransmitter *MessageTransmitterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.contract.Transact(opts, method, params...)
}

// AttesterManager is a free data retrieval call binding the contract method 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (_MessageTransmitter *MessageTransmitterCaller) AttesterManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "attesterManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttesterManager is a free data retrieval call binding the contract method 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (_MessageTransmitter *MessageTransmitterSession) AttesterManager() (common.Address, error) {
	return _MessageTransmitter.Contract.AttesterManager(&_MessageTransmitter.CallOpts)
}

// AttesterManager is a free data retrieval call binding the contract method 0x9b0d94b7.
//
// Solidity: function attesterManager() view returns(address)
func (_MessageTransmitter *MessageTransmitterCallerSession) AttesterManager() (common.Address, error) {
	return _MessageTransmitter.Contract.AttesterManager(&_MessageTransmitter.CallOpts)
}

// GetEnabledAttester is a free data retrieval call binding the contract method 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (_MessageTransmitter *MessageTransmitterCaller) GetEnabledAttester(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "getEnabledAttester", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnabledAttester is a free data retrieval call binding the contract method 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (_MessageTransmitter *MessageTransmitterSession) GetEnabledAttester(index *big.Int) (common.Address, error) {
	return _MessageTransmitter.Contract.GetEnabledAttester(&_MessageTransmitter.CallOpts, index)
}

// GetEnabledAttester is a free data retrieval call binding the contract method 0xbeb673d8.
//
// Solidity: function getEnabledAttester(uint256 index) view returns(address)
func (_MessageTransmitter *MessageTransmitterCallerSession) GetEnabledAttester(index *big.Int) (common.Address, error) {
	return _MessageTransmitter.Contract.GetEnabledAttester(&_MessageTransmitter.CallOpts, index)
}

// GetNumEnabledAttesters is a free data retrieval call binding the contract method 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCaller) GetNumEnabledAttesters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "getNumEnabledAttesters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumEnabledAttesters is a free data retrieval call binding the contract method 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterSession) GetNumEnabledAttesters() (*big.Int, error) {
	return _MessageTransmitter.Contract.GetNumEnabledAttesters(&_MessageTransmitter.CallOpts)
}

// GetNumEnabledAttesters is a free data retrieval call binding the contract method 0x51079a53.
//
// Solidity: function getNumEnabledAttesters() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCallerSession) GetNumEnabledAttesters() (*big.Int, error) {
	return _MessageTransmitter.Contract.GetNumEnabledAttesters(&_MessageTransmitter.CallOpts)
}

// IsEnabledAttester is a free data retrieval call binding the contract method 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (_MessageTransmitter *MessageTransmitterCaller) IsEnabledAttester(opts *bind.CallOpts, attester common.Address) (bool, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "isEnabledAttester", attester)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnabledAttester is a free data retrieval call binding the contract method 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (_MessageTransmitter *MessageTransmitterSession) IsEnabledAttester(attester common.Address) (bool, error) {
	return _MessageTransmitter.Contract.IsEnabledAttester(&_MessageTransmitter.CallOpts, attester)
}

// IsEnabledAttester is a free data retrieval call binding the contract method 0x7af82f60.
//
// Solidity: function isEnabledAttester(address attester) view returns(bool)
func (_MessageTransmitter *MessageTransmitterCallerSession) IsEnabledAttester(attester common.Address) (bool, error) {
	return _MessageTransmitter.Contract.IsEnabledAttester(&_MessageTransmitter.CallOpts, attester)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterCaller) LocalDomain(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "localDomain")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterSession) LocalDomain() (uint32, error) {
	return _MessageTransmitter.Contract.LocalDomain(&_MessageTransmitter.CallOpts)
}

// LocalDomain is a free data retrieval call binding the contract method 0x8d3638f4.
//
// Solidity: function localDomain() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterCallerSession) LocalDomain() (uint32, error) {
	return _MessageTransmitter.Contract.LocalDomain(&_MessageTransmitter.CallOpts)
}

// MaxMessageBodySize is a free data retrieval call binding the contract method 0xaf47b9bb.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCaller) MaxMessageBodySize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "maxMessageBodySize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxMessageBodySize is a free data retrieval call binding the contract method 0xaf47b9bb.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterSession) MaxMessageBodySize() (*big.Int, error) {
	return _MessageTransmitter.Contract.MaxMessageBodySize(&_MessageTransmitter.CallOpts)
}

// MaxMessageBodySize is a free data retrieval call binding the contract method 0xaf47b9bb.
//
// Solidity: function maxMessageBodySize() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCallerSession) MaxMessageBodySize() (*big.Int, error) {
	return _MessageTransmitter.Contract.MaxMessageBodySize(&_MessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MessageTransmitter *MessageTransmitterCaller) NextAvailableNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "nextAvailableNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MessageTransmitter *MessageTransmitterSession) NextAvailableNonce() (uint64, error) {
	return _MessageTransmitter.Contract.NextAvailableNonce(&_MessageTransmitter.CallOpts)
}

// NextAvailableNonce is a free data retrieval call binding the contract method 0x8371744e.
//
// Solidity: function nextAvailableNonce() view returns(uint64)
func (_MessageTransmitter *MessageTransmitterCallerSession) NextAvailableNonce() (uint64, error) {
	return _MessageTransmitter.Contract.NextAvailableNonce(&_MessageTransmitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageTransmitter *MessageTransmitterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageTransmitter *MessageTransmitterSession) Owner() (common.Address, error) {
	return _MessageTransmitter.Contract.Owner(&_MessageTransmitter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MessageTransmitter *MessageTransmitterCallerSession) Owner() (common.Address, error) {
	return _MessageTransmitter.Contract.Owner(&_MessageTransmitter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageTransmitter *MessageTransmitterCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageTransmitter *MessageTransmitterSession) Paused() (bool, error) {
	return _MessageTransmitter.Contract.Paused(&_MessageTransmitter.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MessageTransmitter *MessageTransmitterCallerSession) Paused() (bool, error) {
	return _MessageTransmitter.Contract.Paused(&_MessageTransmitter.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_MessageTransmitter *MessageTransmitterCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_MessageTransmitter *MessageTransmitterSession) Pauser() (common.Address, error) {
	return _MessageTransmitter.Contract.Pauser(&_MessageTransmitter.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_MessageTransmitter *MessageTransmitterCallerSession) Pauser() (common.Address, error) {
	return _MessageTransmitter.Contract.Pauser(&_MessageTransmitter.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MessageTransmitter *MessageTransmitterCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MessageTransmitter *MessageTransmitterSession) PendingOwner() (common.Address, error) {
	return _MessageTransmitter.Contract.PendingOwner(&_MessageTransmitter.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MessageTransmitter *MessageTransmitterCallerSession) PendingOwner() (common.Address, error) {
	return _MessageTransmitter.Contract.PendingOwner(&_MessageTransmitter.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_MessageTransmitter *MessageTransmitterCaller) Rescuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "rescuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_MessageTransmitter *MessageTransmitterSession) Rescuer() (common.Address, error) {
	return _MessageTransmitter.Contract.Rescuer(&_MessageTransmitter.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_MessageTransmitter *MessageTransmitterCallerSession) Rescuer() (common.Address, error) {
	return _MessageTransmitter.Contract.Rescuer(&_MessageTransmitter.CallOpts)
}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCaller) SignatureThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "signatureThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterSession) SignatureThreshold() (*big.Int, error) {
	return _MessageTransmitter.Contract.SignatureThreshold(&_MessageTransmitter.CallOpts)
}

// SignatureThreshold is a free data retrieval call binding the contract method 0xa82f2e26.
//
// Solidity: function signatureThreshold() view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCallerSession) SignatureThreshold() (*big.Int, error) {
	return _MessageTransmitter.Contract.SignatureThreshold(&_MessageTransmitter.CallOpts)
}

// UsedNonces is a free data retrieval call binding the contract method 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCaller) UsedNonces(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "usedNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsedNonces is a free data retrieval call binding the contract method 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (_MessageTransmitter *MessageTransmitterSession) UsedNonces(arg0 [32]byte) (*big.Int, error) {
	return _MessageTransmitter.Contract.UsedNonces(&_MessageTransmitter.CallOpts, arg0)
}

// UsedNonces is a free data retrieval call binding the contract method 0xfeb61724.
//
// Solidity: function usedNonces(bytes32 ) view returns(uint256)
func (_MessageTransmitter *MessageTransmitterCallerSession) UsedNonces(arg0 [32]byte) (*big.Int, error) {
	return _MessageTransmitter.Contract.UsedNonces(&_MessageTransmitter.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterCaller) Version(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _MessageTransmitter.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterSession) Version() (uint32, error) {
	return _MessageTransmitter.Contract.Version(&_MessageTransmitter.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint32)
func (_MessageTransmitter *MessageTransmitterCallerSession) Version() (uint32, error) {
	return _MessageTransmitter.Contract.Version(&_MessageTransmitter.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MessageTransmitter *MessageTransmitterTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MessageTransmitter *MessageTransmitterSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessageTransmitter.Contract.AcceptOwnership(&_MessageTransmitter.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MessageTransmitter.Contract.AcceptOwnership(&_MessageTransmitter.TransactOpts)
}

// DisableAttester is a paid mutator transaction binding the contract method 0x2d025080.
//
// Solidity: function disableAttester(address attester) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) DisableAttester(opts *bind.TransactOpts, attester common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "disableAttester", attester)
}

// DisableAttester is a paid mutator transaction binding the contract method 0x2d025080.
//
// Solidity: function disableAttester(address attester) returns()
func (_MessageTransmitter *MessageTransmitterSession) DisableAttester(attester common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.DisableAttester(&_MessageTransmitter.TransactOpts, attester)
}

// DisableAttester is a paid mutator transaction binding the contract method 0x2d025080.
//
// Solidity: function disableAttester(address attester) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) DisableAttester(attester common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.DisableAttester(&_MessageTransmitter.TransactOpts, attester)
}

// EnableAttester is a paid mutator transaction binding the contract method 0xfae36879.
//
// Solidity: function enableAttester(address newAttester) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) EnableAttester(opts *bind.TransactOpts, newAttester common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "enableAttester", newAttester)
}

// EnableAttester is a paid mutator transaction binding the contract method 0xfae36879.
//
// Solidity: function enableAttester(address newAttester) returns()
func (_MessageTransmitter *MessageTransmitterSession) EnableAttester(newAttester common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.EnableAttester(&_MessageTransmitter.TransactOpts, newAttester)
}

// EnableAttester is a paid mutator transaction binding the contract method 0xfae36879.
//
// Solidity: function enableAttester(address newAttester) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) EnableAttester(newAttester common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.EnableAttester(&_MessageTransmitter.TransactOpts, newAttester)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageTransmitter *MessageTransmitterTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageTransmitter *MessageTransmitterSession) Pause() (*types.Transaction, error) {
	return _MessageTransmitter.Contract.Pause(&_MessageTransmitter.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) Pause() (*types.Transaction, error) {
	return _MessageTransmitter.Contract.Pause(&_MessageTransmitter.TransactOpts)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool success)
func (_MessageTransmitter *MessageTransmitterTransactor) ReceiveMessage(opts *bind.TransactOpts, message []byte, attestation []byte) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "receiveMessage", message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool success)
func (_MessageTransmitter *MessageTransmitterSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.ReceiveMessage(&_MessageTransmitter.TransactOpts, message, attestation)
}

// ReceiveMessage is a paid mutator transaction binding the contract method 0x57ecfd28.
//
// Solidity: function receiveMessage(bytes message, bytes attestation) returns(bool success)
func (_MessageTransmitter *MessageTransmitterTransactorSession) ReceiveMessage(message []byte, attestation []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.ReceiveMessage(&_MessageTransmitter.TransactOpts, message, attestation)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) ReplaceMessage(opts *bind.TransactOpts, originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "replaceMessage", originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_MessageTransmitter *MessageTransmitterSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.ReplaceMessage(&_MessageTransmitter.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// ReplaceMessage is a paid mutator transaction binding the contract method 0xb857b774.
//
// Solidity: function replaceMessage(bytes originalMessage, bytes originalAttestation, bytes newMessageBody, bytes32 newDestinationCaller) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) ReplaceMessage(originalMessage []byte, originalAttestation []byte, newMessageBody []byte, newDestinationCaller [32]byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.ReplaceMessage(&_MessageTransmitter.TransactOpts, originalMessage, originalAttestation, newMessageBody, newDestinationCaller)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) RescueERC20(opts *bind.TransactOpts, tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "rescueERC20", tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_MessageTransmitter *MessageTransmitterSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.RescueERC20(&_MessageTransmitter.TransactOpts, tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.RescueERC20(&_MessageTransmitter.TransactOpts, tokenContract, to, amount)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterTransactor) SendMessage(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "sendMessage", destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SendMessage(&_MessageTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessage is a paid mutator transaction binding the contract method 0x0ba469bc.
//
// Solidity: function sendMessage(uint32 destinationDomain, bytes32 recipient, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterTransactorSession) SendMessage(destinationDomain uint32, recipient [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SendMessage(&_MessageTransmitter.TransactOpts, destinationDomain, recipient, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterTransactor) SendMessageWithCaller(opts *bind.TransactOpts, destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "sendMessageWithCaller", destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SendMessageWithCaller(&_MessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SendMessageWithCaller is a paid mutator transaction binding the contract method 0xf7259a75.
//
// Solidity: function sendMessageWithCaller(uint32 destinationDomain, bytes32 recipient, bytes32 destinationCaller, bytes messageBody) returns(uint64)
func (_MessageTransmitter *MessageTransmitterTransactorSession) SendMessageWithCaller(destinationDomain uint32, recipient [32]byte, destinationCaller [32]byte, messageBody []byte) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SendMessageWithCaller(&_MessageTransmitter.TransactOpts, destinationDomain, recipient, destinationCaller, messageBody)
}

// SetMaxMessageBodySize is a paid mutator transaction binding the contract method 0x92492c68.
//
// Solidity: function setMaxMessageBodySize(uint256 newMaxMessageBodySize) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) SetMaxMessageBodySize(opts *bind.TransactOpts, newMaxMessageBodySize *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "setMaxMessageBodySize", newMaxMessageBodySize)
}

// SetMaxMessageBodySize is a paid mutator transaction binding the contract method 0x92492c68.
//
// Solidity: function setMaxMessageBodySize(uint256 newMaxMessageBodySize) returns()
func (_MessageTransmitter *MessageTransmitterSession) SetMaxMessageBodySize(newMaxMessageBodySize *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SetMaxMessageBodySize(&_MessageTransmitter.TransactOpts, newMaxMessageBodySize)
}

// SetMaxMessageBodySize is a paid mutator transaction binding the contract method 0x92492c68.
//
// Solidity: function setMaxMessageBodySize(uint256 newMaxMessageBodySize) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) SetMaxMessageBodySize(newMaxMessageBodySize *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SetMaxMessageBodySize(&_MessageTransmitter.TransactOpts, newMaxMessageBodySize)
}

// SetSignatureThreshold is a paid mutator transaction binding the contract method 0xbbde5374.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) SetSignatureThreshold(opts *bind.TransactOpts, newSignatureThreshold *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "setSignatureThreshold", newSignatureThreshold)
}

// SetSignatureThreshold is a paid mutator transaction binding the contract method 0xbbde5374.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (_MessageTransmitter *MessageTransmitterSession) SetSignatureThreshold(newSignatureThreshold *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SetSignatureThreshold(&_MessageTransmitter.TransactOpts, newSignatureThreshold)
}

// SetSignatureThreshold is a paid mutator transaction binding the contract method 0xbbde5374.
//
// Solidity: function setSignatureThreshold(uint256 newSignatureThreshold) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) SetSignatureThreshold(newSignatureThreshold *big.Int) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.SetSignatureThreshold(&_MessageTransmitter.TransactOpts, newSignatureThreshold)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageTransmitter *MessageTransmitterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.TransferOwnership(&_MessageTransmitter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.TransferOwnership(&_MessageTransmitter.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageTransmitter *MessageTransmitterTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageTransmitter *MessageTransmitterSession) Unpause() (*types.Transaction, error) {
	return _MessageTransmitter.Contract.Unpause(&_MessageTransmitter.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) Unpause() (*types.Transaction, error) {
	return _MessageTransmitter.Contract.Unpause(&_MessageTransmitter.TransactOpts)
}

// UpdateAttesterManager is a paid mutator transaction binding the contract method 0xde7769d4.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) UpdateAttesterManager(opts *bind.TransactOpts, newAttesterManager common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "updateAttesterManager", newAttesterManager)
}

// UpdateAttesterManager is a paid mutator transaction binding the contract method 0xde7769d4.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (_MessageTransmitter *MessageTransmitterSession) UpdateAttesterManager(newAttesterManager common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.UpdateAttesterManager(&_MessageTransmitter.TransactOpts, newAttesterManager)
}

// UpdateAttesterManager is a paid mutator transaction binding the contract method 0xde7769d4.
//
// Solidity: function updateAttesterManager(address newAttesterManager) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) UpdateAttesterManager(newAttesterManager common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.UpdateAttesterManager(&_MessageTransmitter.TransactOpts, newAttesterManager)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) UpdatePauser(opts *bind.TransactOpts, _newPauser common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "updatePauser", _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_MessageTransmitter *MessageTransmitterSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.UpdatePauser(&_MessageTransmitter.TransactOpts, _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.UpdatePauser(&_MessageTransmitter.TransactOpts, _newPauser)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_MessageTransmitter *MessageTransmitterTransactor) UpdateRescuer(opts *bind.TransactOpts, newRescuer common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.contract.Transact(opts, "updateRescuer", newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_MessageTransmitter *MessageTransmitterSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.UpdateRescuer(&_MessageTransmitter.TransactOpts, newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_MessageTransmitter *MessageTransmitterTransactorSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _MessageTransmitter.Contract.UpdateRescuer(&_MessageTransmitter.TransactOpts, newRescuer)
}

// MessageTransmitterAttesterDisabledIterator is returned from FilterAttesterDisabled and is used to iterate over the raw logs and unpacked data for AttesterDisabled events raised by the MessageTransmitter contract.
type MessageTransmitterAttesterDisabledIterator struct {
	Event *MessageTransmitterAttesterDisabled // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterAttesterDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterAttesterDisabled)
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
		it.Event = new(MessageTransmitterAttesterDisabled)
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
func (it *MessageTransmitterAttesterDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterAttesterDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterAttesterDisabled represents a AttesterDisabled event raised by the MessageTransmitter contract.
type MessageTransmitterAttesterDisabled struct {
	Attester common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttesterDisabled is a free log retrieval operation binding the contract event 0x78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterAttesterDisabled(opts *bind.FilterOpts, attester []common.Address) (*MessageTransmitterAttesterDisabledIterator, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "AttesterDisabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterAttesterDisabledIterator{contract: _MessageTransmitter.contract, event: "AttesterDisabled", logs: logs, sub: sub}, nil
}

// WatchAttesterDisabled is a free log subscription operation binding the contract event 0x78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchAttesterDisabled(opts *bind.WatchOpts, sink chan<- *MessageTransmitterAttesterDisabled, attester []common.Address) (event.Subscription, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "AttesterDisabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterAttesterDisabled)
				if err := _MessageTransmitter.contract.UnpackLog(event, "AttesterDisabled", log); err != nil {
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

// ParseAttesterDisabled is a log parse operation binding the contract event 0x78e573a18c75957b7cadaab01511aa1c19a659f06ecf53e01de37ed92d3261fc.
//
// Solidity: event AttesterDisabled(address indexed attester)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseAttesterDisabled(log types.Log) (*MessageTransmitterAttesterDisabled, error) {
	event := new(MessageTransmitterAttesterDisabled)
	if err := _MessageTransmitter.contract.UnpackLog(event, "AttesterDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterAttesterEnabledIterator is returned from FilterAttesterEnabled and is used to iterate over the raw logs and unpacked data for AttesterEnabled events raised by the MessageTransmitter contract.
type MessageTransmitterAttesterEnabledIterator struct {
	Event *MessageTransmitterAttesterEnabled // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterAttesterEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterAttesterEnabled)
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
		it.Event = new(MessageTransmitterAttesterEnabled)
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
func (it *MessageTransmitterAttesterEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterAttesterEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterAttesterEnabled represents a AttesterEnabled event raised by the MessageTransmitter contract.
type MessageTransmitterAttesterEnabled struct {
	Attester common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAttesterEnabled is a free log retrieval operation binding the contract event 0x5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterAttesterEnabled(opts *bind.FilterOpts, attester []common.Address) (*MessageTransmitterAttesterEnabledIterator, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "AttesterEnabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterAttesterEnabledIterator{contract: _MessageTransmitter.contract, event: "AttesterEnabled", logs: logs, sub: sub}, nil
}

// WatchAttesterEnabled is a free log subscription operation binding the contract event 0x5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchAttesterEnabled(opts *bind.WatchOpts, sink chan<- *MessageTransmitterAttesterEnabled, attester []common.Address) (event.Subscription, error) {

	var attesterRule []interface{}
	for _, attesterItem := range attester {
		attesterRule = append(attesterRule, attesterItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "AttesterEnabled", attesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterAttesterEnabled)
				if err := _MessageTransmitter.contract.UnpackLog(event, "AttesterEnabled", log); err != nil {
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

// ParseAttesterEnabled is a log parse operation binding the contract event 0x5b99bab45c72ce67e89466dbc47480b9c1fde1400e7268bbf463b8354ee4653f.
//
// Solidity: event AttesterEnabled(address indexed attester)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseAttesterEnabled(log types.Log) (*MessageTransmitterAttesterEnabled, error) {
	event := new(MessageTransmitterAttesterEnabled)
	if err := _MessageTransmitter.contract.UnpackLog(event, "AttesterEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterAttesterManagerUpdatedIterator is returned from FilterAttesterManagerUpdated and is used to iterate over the raw logs and unpacked data for AttesterManagerUpdated events raised by the MessageTransmitter contract.
type MessageTransmitterAttesterManagerUpdatedIterator struct {
	Event *MessageTransmitterAttesterManagerUpdated // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterAttesterManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterAttesterManagerUpdated)
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
		it.Event = new(MessageTransmitterAttesterManagerUpdated)
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
func (it *MessageTransmitterAttesterManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterAttesterManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterAttesterManagerUpdated represents a AttesterManagerUpdated event raised by the MessageTransmitter contract.
type MessageTransmitterAttesterManagerUpdated struct {
	PreviousAttesterManager common.Address
	NewAttesterManager      common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterAttesterManagerUpdated is a free log retrieval operation binding the contract event 0x0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d3460219.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterAttesterManagerUpdated(opts *bind.FilterOpts, previousAttesterManager []common.Address, newAttesterManager []common.Address) (*MessageTransmitterAttesterManagerUpdatedIterator, error) {

	var previousAttesterManagerRule []interface{}
	for _, previousAttesterManagerItem := range previousAttesterManager {
		previousAttesterManagerRule = append(previousAttesterManagerRule, previousAttesterManagerItem)
	}
	var newAttesterManagerRule []interface{}
	for _, newAttesterManagerItem := range newAttesterManager {
		newAttesterManagerRule = append(newAttesterManagerRule, newAttesterManagerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "AttesterManagerUpdated", previousAttesterManagerRule, newAttesterManagerRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterAttesterManagerUpdatedIterator{contract: _MessageTransmitter.contract, event: "AttesterManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchAttesterManagerUpdated is a free log subscription operation binding the contract event 0x0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d3460219.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchAttesterManagerUpdated(opts *bind.WatchOpts, sink chan<- *MessageTransmitterAttesterManagerUpdated, previousAttesterManager []common.Address, newAttesterManager []common.Address) (event.Subscription, error) {

	var previousAttesterManagerRule []interface{}
	for _, previousAttesterManagerItem := range previousAttesterManager {
		previousAttesterManagerRule = append(previousAttesterManagerRule, previousAttesterManagerItem)
	}
	var newAttesterManagerRule []interface{}
	for _, newAttesterManagerItem := range newAttesterManager {
		newAttesterManagerRule = append(newAttesterManagerRule, newAttesterManagerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "AttesterManagerUpdated", previousAttesterManagerRule, newAttesterManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterAttesterManagerUpdated)
				if err := _MessageTransmitter.contract.UnpackLog(event, "AttesterManagerUpdated", log); err != nil {
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

// ParseAttesterManagerUpdated is a log parse operation binding the contract event 0x0cee1b7ae04f3c788dd3a46c6fa677eb95b913611ef7ab59524fdc09d3460219.
//
// Solidity: event AttesterManagerUpdated(address indexed previousAttesterManager, address indexed newAttesterManager)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseAttesterManagerUpdated(log types.Log) (*MessageTransmitterAttesterManagerUpdated, error) {
	event := new(MessageTransmitterAttesterManagerUpdated)
	if err := _MessageTransmitter.contract.UnpackLog(event, "AttesterManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterMaxMessageBodySizeUpdatedIterator is returned from FilterMaxMessageBodySizeUpdated and is used to iterate over the raw logs and unpacked data for MaxMessageBodySizeUpdated events raised by the MessageTransmitter contract.
type MessageTransmitterMaxMessageBodySizeUpdatedIterator struct {
	Event *MessageTransmitterMaxMessageBodySizeUpdated // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterMaxMessageBodySizeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterMaxMessageBodySizeUpdated)
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
		it.Event = new(MessageTransmitterMaxMessageBodySizeUpdated)
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
func (it *MessageTransmitterMaxMessageBodySizeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterMaxMessageBodySizeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterMaxMessageBodySizeUpdated represents a MaxMessageBodySizeUpdated event raised by the MessageTransmitter contract.
type MessageTransmitterMaxMessageBodySizeUpdated struct {
	NewMaxMessageBodySize *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterMaxMessageBodySizeUpdated is a free log retrieval operation binding the contract event 0xb13bf6bebed03d1b318e3ea32e4b2a3ad9f5e2312cdf340a2f4bbfaee39f928d.
//
// Solidity: event MaxMessageBodySizeUpdated(uint256 newMaxMessageBodySize)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterMaxMessageBodySizeUpdated(opts *bind.FilterOpts) (*MessageTransmitterMaxMessageBodySizeUpdatedIterator, error) {

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "MaxMessageBodySizeUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterMaxMessageBodySizeUpdatedIterator{contract: _MessageTransmitter.contract, event: "MaxMessageBodySizeUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxMessageBodySizeUpdated is a free log subscription operation binding the contract event 0xb13bf6bebed03d1b318e3ea32e4b2a3ad9f5e2312cdf340a2f4bbfaee39f928d.
//
// Solidity: event MaxMessageBodySizeUpdated(uint256 newMaxMessageBodySize)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchMaxMessageBodySizeUpdated(opts *bind.WatchOpts, sink chan<- *MessageTransmitterMaxMessageBodySizeUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "MaxMessageBodySizeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterMaxMessageBodySizeUpdated)
				if err := _MessageTransmitter.contract.UnpackLog(event, "MaxMessageBodySizeUpdated", log); err != nil {
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

// ParseMaxMessageBodySizeUpdated is a log parse operation binding the contract event 0xb13bf6bebed03d1b318e3ea32e4b2a3ad9f5e2312cdf340a2f4bbfaee39f928d.
//
// Solidity: event MaxMessageBodySizeUpdated(uint256 newMaxMessageBodySize)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseMaxMessageBodySizeUpdated(log types.Log) (*MessageTransmitterMaxMessageBodySizeUpdated, error) {
	event := new(MessageTransmitterMaxMessageBodySizeUpdated)
	if err := _MessageTransmitter.contract.UnpackLog(event, "MaxMessageBodySizeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the MessageTransmitter contract.
type MessageTransmitterMessageReceivedIterator struct {
	Event *MessageTransmitterMessageReceived // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterMessageReceived)
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
		it.Event = new(MessageTransmitterMessageReceived)
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
func (it *MessageTransmitterMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterMessageReceived represents a MessageReceived event raised by the MessageTransmitter contract.
type MessageTransmitterMessageReceived struct {
	Caller       common.Address
	SourceDomain uint32
	Nonce        uint64
	Sender       [32]byte
	MessageBody  []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterMessageReceived(opts *bind.FilterOpts, caller []common.Address, nonce []uint64) (*MessageTransmitterMessageReceivedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "MessageReceived", callerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterMessageReceivedIterator{contract: _MessageTransmitter.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *MessageTransmitterMessageReceived, caller []common.Address, nonce []uint64) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	var nonceRule []interface{}
	for _, nonceItem := range nonce {
		nonceRule = append(nonceRule, nonceItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "MessageReceived", callerRule, nonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterMessageReceived)
				if err := _MessageTransmitter.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x58200b4c34ae05ee816d710053fff3fb75af4395915d3d2a771b24aa10e3cc5d.
//
// Solidity: event MessageReceived(address indexed caller, uint32 sourceDomain, uint64 indexed nonce, bytes32 sender, bytes messageBody)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseMessageReceived(log types.Log) (*MessageTransmitterMessageReceived, error) {
	event := new(MessageTransmitterMessageReceived)
	if err := _MessageTransmitter.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterMessageSentIterator is returned from FilterMessageSent and is used to iterate over the raw logs and unpacked data for MessageSent events raised by the MessageTransmitter contract.
type MessageTransmitterMessageSentIterator struct {
	Event *MessageTransmitterMessageSent // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterMessageSent)
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
		it.Event = new(MessageTransmitterMessageSent)
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
func (it *MessageTransmitterMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterMessageSent represents a MessageSent event raised by the MessageTransmitter contract.
type MessageTransmitterMessageSent struct {
	Message []byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMessageSent is a free log retrieval operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterMessageSent(opts *bind.FilterOpts) (*MessageTransmitterMessageSentIterator, error) {

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterMessageSentIterator{contract: _MessageTransmitter.contract, event: "MessageSent", logs: logs, sub: sub}, nil
}

// WatchMessageSent is a free log subscription operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchMessageSent(opts *bind.WatchOpts, sink chan<- *MessageTransmitterMessageSent) (event.Subscription, error) {

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "MessageSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterMessageSent)
				if err := _MessageTransmitter.contract.UnpackLog(event, "MessageSent", log); err != nil {
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

// ParseMessageSent is a log parse operation binding the contract event 0x8c5261668696ce22758910d05bab8f186d6eb247ceac2af2e82c7dc17669b036.
//
// Solidity: event MessageSent(bytes message)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseMessageSent(log types.Log) (*MessageTransmitterMessageSent, error) {
	event := new(MessageTransmitterMessageSent)
	if err := _MessageTransmitter.contract.UnpackLog(event, "MessageSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the MessageTransmitter contract.
type MessageTransmitterOwnershipTransferStartedIterator struct {
	Event *MessageTransmitterOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterOwnershipTransferStarted)
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
		it.Event = new(MessageTransmitterOwnershipTransferStarted)
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
func (it *MessageTransmitterOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MessageTransmitter contract.
type MessageTransmitterOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageTransmitterOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterOwnershipTransferStartedIterator{contract: _MessageTransmitter.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MessageTransmitterOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterOwnershipTransferStarted)
				if err := _MessageTransmitter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseOwnershipTransferStarted(log types.Log) (*MessageTransmitterOwnershipTransferStarted, error) {
	event := new(MessageTransmitterOwnershipTransferStarted)
	if err := _MessageTransmitter.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MessageTransmitter contract.
type MessageTransmitterOwnershipTransferredIterator struct {
	Event *MessageTransmitterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterOwnershipTransferred)
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
		it.Event = new(MessageTransmitterOwnershipTransferred)
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
func (it *MessageTransmitterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterOwnershipTransferred represents a OwnershipTransferred event raised by the MessageTransmitter contract.
type MessageTransmitterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MessageTransmitterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterOwnershipTransferredIterator{contract: _MessageTransmitter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MessageTransmitterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterOwnershipTransferred)
				if err := _MessageTransmitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MessageTransmitter *MessageTransmitterFilterer) ParseOwnershipTransferred(log types.Log) (*MessageTransmitterOwnershipTransferred, error) {
	event := new(MessageTransmitterOwnershipTransferred)
	if err := _MessageTransmitter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the MessageTransmitter contract.
type MessageTransmitterPauseIterator struct {
	Event *MessageTransmitterPause // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterPause)
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
		it.Event = new(MessageTransmitterPause)
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
func (it *MessageTransmitterPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterPause represents a Pause event raised by the MessageTransmitter contract.
type MessageTransmitterPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_MessageTransmitter *MessageTransmitterFilterer) FilterPause(opts *bind.FilterOpts) (*MessageTransmitterPauseIterator, error) {

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterPauseIterator{contract: _MessageTransmitter.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_MessageTransmitter *MessageTransmitterFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *MessageTransmitterPause) (event.Subscription, error) {

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterPause)
				if err := _MessageTransmitter.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_MessageTransmitter *MessageTransmitterFilterer) ParsePause(log types.Log) (*MessageTransmitterPause, error) {
	event := new(MessageTransmitterPause)
	if err := _MessageTransmitter.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterPauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the MessageTransmitter contract.
type MessageTransmitterPauserChangedIterator struct {
	Event *MessageTransmitterPauserChanged // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterPauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterPauserChanged)
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
		it.Event = new(MessageTransmitterPauserChanged)
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
func (it *MessageTransmitterPauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterPauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterPauserChanged represents a PauserChanged event raised by the MessageTransmitter contract.
type MessageTransmitterPauserChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterPauserChanged(opts *bind.FilterOpts, newAddress []common.Address) (*MessageTransmitterPauserChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterPauserChangedIterator{contract: _MessageTransmitter.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *MessageTransmitterPauserChanged, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterPauserChanged)
				if err := _MessageTransmitter.contract.UnpackLog(event, "PauserChanged", log); err != nil {
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

// ParsePauserChanged is a log parse operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_MessageTransmitter *MessageTransmitterFilterer) ParsePauserChanged(log types.Log) (*MessageTransmitterPauserChanged, error) {
	event := new(MessageTransmitterPauserChanged)
	if err := _MessageTransmitter.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterRescuerChangedIterator is returned from FilterRescuerChanged and is used to iterate over the raw logs and unpacked data for RescuerChanged events raised by the MessageTransmitter contract.
type MessageTransmitterRescuerChangedIterator struct {
	Event *MessageTransmitterRescuerChanged // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterRescuerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterRescuerChanged)
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
		it.Event = new(MessageTransmitterRescuerChanged)
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
func (it *MessageTransmitterRescuerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterRescuerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterRescuerChanged represents a RescuerChanged event raised by the MessageTransmitter contract.
type MessageTransmitterRescuerChanged struct {
	NewRescuer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRescuerChanged is a free log retrieval operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterRescuerChanged(opts *bind.FilterOpts, newRescuer []common.Address) (*MessageTransmitterRescuerChangedIterator, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterRescuerChangedIterator{contract: _MessageTransmitter.contract, event: "RescuerChanged", logs: logs, sub: sub}, nil
}

// WatchRescuerChanged is a free log subscription operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchRescuerChanged(opts *bind.WatchOpts, sink chan<- *MessageTransmitterRescuerChanged, newRescuer []common.Address) (event.Subscription, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterRescuerChanged)
				if err := _MessageTransmitter.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
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

// ParseRescuerChanged is a log parse operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseRescuerChanged(log types.Log) (*MessageTransmitterRescuerChanged, error) {
	event := new(MessageTransmitterRescuerChanged)
	if err := _MessageTransmitter.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterSignatureThresholdUpdatedIterator is returned from FilterSignatureThresholdUpdated and is used to iterate over the raw logs and unpacked data for SignatureThresholdUpdated events raised by the MessageTransmitter contract.
type MessageTransmitterSignatureThresholdUpdatedIterator struct {
	Event *MessageTransmitterSignatureThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterSignatureThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterSignatureThresholdUpdated)
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
		it.Event = new(MessageTransmitterSignatureThresholdUpdated)
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
func (it *MessageTransmitterSignatureThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterSignatureThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterSignatureThresholdUpdated represents a SignatureThresholdUpdated event raised by the MessageTransmitter contract.
type MessageTransmitterSignatureThresholdUpdated struct {
	OldSignatureThreshold *big.Int
	NewSignatureThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSignatureThresholdUpdated is a free log retrieval operation binding the contract event 0x149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (_MessageTransmitter *MessageTransmitterFilterer) FilterSignatureThresholdUpdated(opts *bind.FilterOpts) (*MessageTransmitterSignatureThresholdUpdatedIterator, error) {

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "SignatureThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterSignatureThresholdUpdatedIterator{contract: _MessageTransmitter.contract, event: "SignatureThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchSignatureThresholdUpdated is a free log subscription operation binding the contract event 0x149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (_MessageTransmitter *MessageTransmitterFilterer) WatchSignatureThresholdUpdated(opts *bind.WatchOpts, sink chan<- *MessageTransmitterSignatureThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "SignatureThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterSignatureThresholdUpdated)
				if err := _MessageTransmitter.contract.UnpackLog(event, "SignatureThresholdUpdated", log); err != nil {
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

// ParseSignatureThresholdUpdated is a log parse operation binding the contract event 0x149153f58b4da003a8cfd4523709a202402182cb5aa335046911277a1be6eede.
//
// Solidity: event SignatureThresholdUpdated(uint256 oldSignatureThreshold, uint256 newSignatureThreshold)
func (_MessageTransmitter *MessageTransmitterFilterer) ParseSignatureThresholdUpdated(log types.Log) (*MessageTransmitterSignatureThresholdUpdated, error) {
	event := new(MessageTransmitterSignatureThresholdUpdated)
	if err := _MessageTransmitter.contract.UnpackLog(event, "SignatureThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MessageTransmitterUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the MessageTransmitter contract.
type MessageTransmitterUnpauseIterator struct {
	Event *MessageTransmitterUnpause // Event containing the contract specifics and raw log

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
func (it *MessageTransmitterUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MessageTransmitterUnpause)
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
		it.Event = new(MessageTransmitterUnpause)
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
func (it *MessageTransmitterUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MessageTransmitterUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MessageTransmitterUnpause represents a Unpause event raised by the MessageTransmitter contract.
type MessageTransmitterUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_MessageTransmitter *MessageTransmitterFilterer) FilterUnpause(opts *bind.FilterOpts) (*MessageTransmitterUnpauseIterator, error) {

	logs, sub, err := _MessageTransmitter.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &MessageTransmitterUnpauseIterator{contract: _MessageTransmitter.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_MessageTransmitter *MessageTransmitterFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *MessageTransmitterUnpause) (event.Subscription, error) {

	logs, sub, err := _MessageTransmitter.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MessageTransmitterUnpause)
				if err := _MessageTransmitter.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_MessageTransmitter *MessageTransmitterFilterer) ParseUnpause(log types.Log) (*MessageTransmitterUnpause, error) {
	event := new(MessageTransmitterUnpause)
	if err := _MessageTransmitter.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableMetaData contains all meta data concerning the Ownable contract.
var OwnableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8da5cb5b": "owner()",
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

// Ownable2StepMetaData contains all meta data concerning the Ownable2Step contract.
var Ownable2StepMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"f2fde38b": "transferOwnership(address)",
	},
}

// Ownable2StepABI is the input ABI used to generate the binding from.
// Deprecated: Use Ownable2StepMetaData.ABI instead.
var Ownable2StepABI = Ownable2StepMetaData.ABI

// Deprecated: Use Ownable2StepMetaData.Sigs instead.
// Ownable2StepFuncSigs maps the 4-byte function signature to its string representation.
var Ownable2StepFuncSigs = Ownable2StepMetaData.Sigs

// Ownable2Step is an auto generated Go binding around an Ethereum contract.
type Ownable2Step struct {
	Ownable2StepCaller     // Read-only binding to the contract
	Ownable2StepTransactor // Write-only binding to the contract
	Ownable2StepFilterer   // Log filterer for contract events
}

// Ownable2StepCaller is an auto generated read-only Go binding around an Ethereum contract.
type Ownable2StepCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepTransactor is an auto generated write-only Go binding around an Ethereum contract.
type Ownable2StepTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Ownable2StepFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Ownable2StepSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Ownable2StepSession struct {
	Contract     *Ownable2Step     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Ownable2StepCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Ownable2StepCallerSession struct {
	Contract *Ownable2StepCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// Ownable2StepTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Ownable2StepTransactorSession struct {
	Contract     *Ownable2StepTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// Ownable2StepRaw is an auto generated low-level Go binding around an Ethereum contract.
type Ownable2StepRaw struct {
	Contract *Ownable2Step // Generic contract binding to access the raw methods on
}

// Ownable2StepCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Ownable2StepCallerRaw struct {
	Contract *Ownable2StepCaller // Generic read-only contract binding to access the raw methods on
}

// Ownable2StepTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Ownable2StepTransactorRaw struct {
	Contract *Ownable2StepTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable2Step creates a new instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2Step(address common.Address, backend bind.ContractBackend) (*Ownable2Step, error) {
	contract, err := bindOwnable2Step(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable2Step{Ownable2StepCaller: Ownable2StepCaller{contract: contract}, Ownable2StepTransactor: Ownable2StepTransactor{contract: contract}, Ownable2StepFilterer: Ownable2StepFilterer{contract: contract}}, nil
}

// NewOwnable2StepCaller creates a new read-only instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepCaller(address common.Address, caller bind.ContractCaller) (*Ownable2StepCaller, error) {
	contract, err := bindOwnable2Step(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepCaller{contract: contract}, nil
}

// NewOwnable2StepTransactor creates a new write-only instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepTransactor(address common.Address, transactor bind.ContractTransactor) (*Ownable2StepTransactor, error) {
	contract, err := bindOwnable2Step(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepTransactor{contract: contract}, nil
}

// NewOwnable2StepFilterer creates a new log filterer instance of Ownable2Step, bound to a specific deployed contract.
func NewOwnable2StepFilterer(address common.Address, filterer bind.ContractFilterer) (*Ownable2StepFilterer, error) {
	contract, err := bindOwnable2Step(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepFilterer{contract: contract}, nil
}

// bindOwnable2Step binds a generic wrapper to an already deployed contract.
func bindOwnable2Step(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Ownable2StepMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2Step *Ownable2StepRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2Step.Contract.Ownable2StepCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2Step *Ownable2StepRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.Contract.Ownable2StepTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2Step *Ownable2StepRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2Step.Contract.Ownable2StepTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable2Step *Ownable2StepCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable2Step.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable2Step *Ownable2StepTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable2Step *Ownable2StepTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable2Step.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2Step.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepSession) Owner() (common.Address, error) {
	return _Ownable2Step.Contract.Owner(&_Ownable2Step.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable2Step *Ownable2StepCallerSession) Owner() (common.Address, error) {
	return _Ownable2Step.Contract.Owner(&_Ownable2Step.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable2Step.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepSession) PendingOwner() (common.Address, error) {
	return _Ownable2Step.Contract.PendingOwner(&_Ownable2Step.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Ownable2Step *Ownable2StepCallerSession) PendingOwner() (common.Address, error) {
	return _Ownable2Step.Contract.PendingOwner(&_Ownable2Step.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable2Step.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.AcceptOwnership(&_Ownable2Step.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Ownable2Step *Ownable2StepTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Ownable2Step.Contract.AcceptOwnership(&_Ownable2Step.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.Contract.TransferOwnership(&_Ownable2Step.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable2Step *Ownable2StepTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable2Step.Contract.TransferOwnership(&_Ownable2Step.TransactOpts, newOwner)
}

// Ownable2StepOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferStartedIterator struct {
	Event *Ownable2StepOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *Ownable2StepOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepOwnershipTransferStarted)
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
		it.Event = new(Ownable2StepOwnershipTransferStarted)
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
func (it *Ownable2StepOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepOwnershipTransferStartedIterator{contract: _Ownable2Step.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *Ownable2StepOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepOwnershipTransferStarted)
				if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) ParseOwnershipTransferStarted(log types.Log) (*Ownable2StepOwnershipTransferStarted, error) {
	event := new(Ownable2StepOwnershipTransferStarted)
	if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Ownable2StepOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferredIterator struct {
	Event *Ownable2StepOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Ownable2StepOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Ownable2StepOwnershipTransferred)
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
		it.Event = new(Ownable2StepOwnershipTransferred)
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
func (it *Ownable2StepOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Ownable2StepOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Ownable2StepOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable2Step contract.
type Ownable2StepOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Ownable2StepOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Ownable2StepOwnershipTransferredIterator{contract: _Ownable2Step.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable2Step *Ownable2StepFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Ownable2StepOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable2Step.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Ownable2StepOwnershipTransferred)
				if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable2Step *Ownable2StepFilterer) ParseOwnershipTransferred(log types.Log) (*Ownable2StepOwnershipTransferred, error) {
	event := new(Ownable2StepOwnershipTransferred)
	if err := _Ownable2Step.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableMetaData contains all meta data concerning the Pausable contract.
var PausableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PauserChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauser\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newPauser\",\"type\":\"address\"}],\"name\":\"updatePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"8456cb59": "pause()",
		"5c975abb": "paused()",
		"9fd0506d": "pauser()",
		"e30c3978": "pendingOwner()",
		"f2fde38b": "transferOwnership(address)",
		"3f4ba83a": "unpause()",
		"554bab3c": "updatePauser(address)",
	},
	Bin: "0x60806040526002805460ff60a01b1916905534801561001d57600080fd5b5061002e610029610033565b610037565b6100ae565b3390565b600180546001600160a01b031916905561005b8161005e602090811b6105bb17901c565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6107b8806100bd6000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c80638456cb59116100765780639fd0506d1161005b5780639fd0506d14610142578063e30c39781461014a578063f2fde38b14610152576100a3565b80638456cb59146101095780638da5cb5b14610111576100a3565b80633f4ba83a146100a8578063554bab3c146100b25780635c975abb146100e557806379ba509714610101575b600080fd5b6100b0610185565b005b6100b0600480360360208110156100c857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610248565b6100ed610331565b604080519115158252519081900360200190f35b6100b0610352565b6100b06103f5565b6101196104cf565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6101196104eb565b610119610507565b6100b06004803603602081101561016857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610523565b60025473ffffffffffffffffffffffffffffffffffffffff1633146101f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806107616022913960400191505060405180910390fd5b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff1690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b610250610630565b73ffffffffffffffffffffffffffffffffffffffff81166102bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260288152602001806107106028913960400191505060405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691909117918290556040519116907fb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a60490600090a250565b60025474010000000000000000000000000000000000000000900460ff1681565b600061035c6106da565b90508073ffffffffffffffffffffffffffffffffffffffff1661037d610507565b73ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806107386029913960400191505060405180910390fd5b6103f2816106de565b50565b60025473ffffffffffffffffffffffffffffffffffffffff163314610465576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806107616022913960400191505060405180910390fd5b600280547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60025473ffffffffffffffffffffffffffffffffffffffff1690565b60015473ffffffffffffffffffffffffffffffffffffffff1690565b61052b610630565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556105766104cf565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6106386106da565b73ffffffffffffffffffffffffffffffffffffffff166106566104cf565b73ffffffffffffffffffffffffffffffffffffffff16146106d857604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b565b3390565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556103f2816105bb56fe5061757361626c653a206e65772070617573657220697320746865207a65726f20616464726573734f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e65725061757361626c653a2063616c6c6572206973206e6f742074686520706175736572a264697066735822122032a3fbfc8c028f60210e0b9ce54aac646ede4563efc487a16a16868056543a2e64736f6c63430007060033",
}

// PausableABI is the input ABI used to generate the binding from.
// Deprecated: Use PausableMetaData.ABI instead.
var PausableABI = PausableMetaData.ABI

// Deprecated: Use PausableMetaData.Sigs instead.
// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = PausableMetaData.Sigs

// PausableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PausableMetaData.Bin instead.
var PausableBin = PausableMetaData.Bin

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := PausableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PausableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Pausable *PausableCaller) Pauser(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "pauser")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Pausable *PausableSession) Pauser() (common.Address, error) {
	return _Pausable.Contract.Pauser(&_Pausable.CallOpts)
}

// Pauser is a free data retrieval call binding the contract method 0x9fd0506d.
//
// Solidity: function pauser() view returns(address)
func (_Pausable *PausableCallerSession) Pauser() (common.Address, error) {
	return _Pausable.Contract.Pauser(&_Pausable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Pausable *PausableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Pausable *PausableSession) PendingOwner() (common.Address, error) {
	return _Pausable.Contract.PendingOwner(&_Pausable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Pausable *PausableCallerSession) PendingOwner() (common.Address, error) {
	return _Pausable.Contract.PendingOwner(&_Pausable.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Pausable *PausableTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Pausable *PausableSession) AcceptOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.AcceptOwnership(&_Pausable.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Pausable *PausableTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Pausable.Contract.AcceptOwnership(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Pausable *PausableTransactor) UpdatePauser(opts *bind.TransactOpts, _newPauser common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "updatePauser", _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Pausable *PausableSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.UpdatePauser(&_Pausable.TransactOpts, _newPauser)
}

// UpdatePauser is a paid mutator transaction binding the contract method 0x554bab3c.
//
// Solidity: function updatePauser(address _newPauser) returns()
func (_Pausable *PausableTransactorSession) UpdatePauser(_newPauser common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.UpdatePauser(&_Pausable.TransactOpts, _newPauser)
}

// PausableOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Pausable contract.
type PausableOwnershipTransferStartedIterator struct {
	Event *PausableOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferStarted)
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
		it.Event = new(PausableOwnershipTransferStarted)
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
func (it *PausableOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Pausable contract.
type PausableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferStartedIterator{contract: _Pausable.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferStarted)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) ParseOwnershipTransferStarted(log types.Log) (*PausableOwnershipTransferStarted, error) {
	event := new(PausableOwnershipTransferStarted)
	if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pausable contract.
type PausableOwnershipTransferredIterator struct {
	Event *PausableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PausableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableOwnershipTransferred)
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
		it.Event = new(PausableOwnershipTransferred)
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
func (it *PausableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableOwnershipTransferred represents a OwnershipTransferred event raised by the Pausable contract.
type PausableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PausableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PausableOwnershipTransferredIterator{contract: _Pausable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pausable *PausableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PausableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableOwnershipTransferred)
				if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pausable *PausableFilterer) ParseOwnershipTransferred(log types.Log) (*PausableOwnershipTransferred, error) {
	event := new(PausableOwnershipTransferred)
	if err := _Pausable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) ParsePause(log types.Log) (*PausablePause, error) {
	event := new(PausablePause)
	if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausablePauserChangedIterator is returned from FilterPauserChanged and is used to iterate over the raw logs and unpacked data for PauserChanged events raised by the Pausable contract.
type PausablePauserChangedIterator struct {
	Event *PausablePauserChanged // Event containing the contract specifics and raw log

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
func (it *PausablePauserChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePauserChanged)
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
		it.Event = new(PausablePauserChanged)
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
func (it *PausablePauserChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauserChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePauserChanged represents a PauserChanged event raised by the Pausable contract.
type PausablePauserChanged struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPauserChanged is a free log retrieval operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Pausable *PausableFilterer) FilterPauserChanged(opts *bind.FilterOpts, newAddress []common.Address) (*PausablePauserChangedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &PausablePauserChangedIterator{contract: _Pausable.contract, event: "PauserChanged", logs: logs, sub: sub}, nil
}

// WatchPauserChanged is a free log subscription operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Pausable *PausableFilterer) WatchPauserChanged(opts *bind.WatchOpts, sink chan<- *PausablePauserChanged, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "PauserChanged", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePauserChanged)
				if err := _Pausable.contract.UnpackLog(event, "PauserChanged", log); err != nil {
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

// ParsePauserChanged is a log parse operation binding the contract event 0xb80482a293ca2e013eda8683c9bd7fc8347cfdaeea5ede58cba46df502c2a604.
//
// Solidity: event PauserChanged(address indexed newAddress)
func (_Pausable *PausableFilterer) ParsePauserChanged(log types.Log) (*PausablePauserChanged, error) {
	event := new(PausablePauserChanged)
	if err := _Pausable.contract.UnpackLog(event, "PauserChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) ParseUnpause(log types.Log) (*PausableUnpause, error) {
	event := new(PausableUnpause)
	if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RescuableMetaData contains all meta data concerning the Rescuable contract.
var RescuableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"RescuerChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rescueERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rescuer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRescuer\",\"type\":\"address\"}],\"name\":\"updateRescuer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"79ba5097": "acceptOwnership()",
		"8da5cb5b": "owner()",
		"e30c3978": "pendingOwner()",
		"b2118a8d": "rescueERC20(address,address,uint256)",
		"38a63183": "rescuer()",
		"f2fde38b": "transferOwnership(address)",
		"2ab60045": "updateRescuer(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061002161001c610026565b61002a565b6100a1565b3390565b600180546001600160a01b031916905561004e81610051602090811b61047e17901c565b50565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610ac6806100b06000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100f0578063b2118a8d146100f8578063e30c39781461013b578063f2fde38b146101435761007d565b80632ab600451461008257806338a63183146100b757806379ba5097146100e8575b600080fd5b6100b56004803603602081101561009857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610176565b005b6100bf610259565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6100b5610275565b6100bf610318565b6100b56004803603606081101561010e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101359091169060400135610334565b6100bf6103ca565b6100b56004803603602081101561015957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166103e6565b61017e6104f3565b73ffffffffffffffffffffffffffffffffffffffff81166101ea576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001806109f3602a913960400191505060405180910390fd5b600280547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a90600090a250565b60025473ffffffffffffffffffffffffffffffffffffffff1690565b600061027f61059d565b90508073ffffffffffffffffffffffffffffffffffffffff166102a06103ca565b73ffffffffffffffffffffffffffffffffffffffff161461030c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001806109ca6029913960400191505060405180910390fd5b610315816105a1565b50565b60005473ffffffffffffffffffffffffffffffffffffffff1690565b60025473ffffffffffffffffffffffffffffffffffffffff1633146103a4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180610a436024913960400191505060405180910390fd5b6103c573ffffffffffffffffffffffffffffffffffffffff841683836105d2565b505050565b60015473ffffffffffffffffffffffffffffffffffffffff1690565b6103ee6104f3565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8316908117909155610439610318565b73ffffffffffffffffffffffffffffffffffffffff167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6104fb61059d565b73ffffffffffffffffffffffffffffffffffffffff16610519610318565b73ffffffffffffffffffffffffffffffffffffffff161461059b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b565b3390565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001690556103158161047e565b6040805173ffffffffffffffffffffffffffffffffffffffff8416602482015260448082018490528251808303909101815260649091019091526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fa9059cbb000000000000000000000000000000000000000000000000000000001790526103c590849060006106bc826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff166107329092919063ffffffff16565b8051909150156103c5578080602001905160208110156106db57600080fd5b50516103c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a815260200180610a67602a913960400191505060405180910390fd5b6060610741848460008561074b565b90505b9392505050565b6060824710156107a6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180610a1d6026913960400191505060405180910390fd5b6107af85610905565b61081a57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015290519081900360640190fd5b6000808673ffffffffffffffffffffffffffffffffffffffff1685876040518082805190602001908083835b6020831061088357805182527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe09092019160209182019101610846565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d80600081146108e5576040519150601f19603f3d011682016040523d82523d6000602084013e6108ea565b606091505b50915091506108fa82828661090b565b979650505050505050565b3b151590565b6060831561091a575081610744565b82511561092a5782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561098e578181015183820152602001610976565b50505050905090810190601f1680156109bb5780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe4f776e61626c6532537465703a2063616c6c6572206973206e6f7420746865206e6577206f776e6572526573637561626c653a206e6577207265736375657220697320746865207a65726f2061646472657373416464726573733a20696e73756666696369656e742062616c616e636520666f722063616c6c526573637561626c653a2063616c6c6572206973206e6f742074686520726573637565725361666545524332303a204552433230206f7065726174696f6e20646964206e6f742073756363656564a2646970667358221220e941ed2d04087529217b5000e5a7eea32ea0b22465a1f390fea977c3bb6c7b9d64736f6c63430007060033",
}

// RescuableABI is the input ABI used to generate the binding from.
// Deprecated: Use RescuableMetaData.ABI instead.
var RescuableABI = RescuableMetaData.ABI

// Deprecated: Use RescuableMetaData.Sigs instead.
// RescuableFuncSigs maps the 4-byte function signature to its string representation.
var RescuableFuncSigs = RescuableMetaData.Sigs

// RescuableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RescuableMetaData.Bin instead.
var RescuableBin = RescuableMetaData.Bin

// DeployRescuable deploys a new Ethereum contract, binding an instance of Rescuable to it.
func DeployRescuable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Rescuable, error) {
	parsed, err := RescuableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RescuableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rescuable{RescuableCaller: RescuableCaller{contract: contract}, RescuableTransactor: RescuableTransactor{contract: contract}, RescuableFilterer: RescuableFilterer{contract: contract}}, nil
}

// Rescuable is an auto generated Go binding around an Ethereum contract.
type Rescuable struct {
	RescuableCaller     // Read-only binding to the contract
	RescuableTransactor // Write-only binding to the contract
	RescuableFilterer   // Log filterer for contract events
}

// RescuableCaller is an auto generated read-only Go binding around an Ethereum contract.
type RescuableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RescuableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RescuableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RescuableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RescuableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RescuableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RescuableSession struct {
	Contract     *Rescuable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RescuableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RescuableCallerSession struct {
	Contract *RescuableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// RescuableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RescuableTransactorSession struct {
	Contract     *RescuableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// RescuableRaw is an auto generated low-level Go binding around an Ethereum contract.
type RescuableRaw struct {
	Contract *Rescuable // Generic contract binding to access the raw methods on
}

// RescuableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RescuableCallerRaw struct {
	Contract *RescuableCaller // Generic read-only contract binding to access the raw methods on
}

// RescuableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RescuableTransactorRaw struct {
	Contract *RescuableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRescuable creates a new instance of Rescuable, bound to a specific deployed contract.
func NewRescuable(address common.Address, backend bind.ContractBackend) (*Rescuable, error) {
	contract, err := bindRescuable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rescuable{RescuableCaller: RescuableCaller{contract: contract}, RescuableTransactor: RescuableTransactor{contract: contract}, RescuableFilterer: RescuableFilterer{contract: contract}}, nil
}

// NewRescuableCaller creates a new read-only instance of Rescuable, bound to a specific deployed contract.
func NewRescuableCaller(address common.Address, caller bind.ContractCaller) (*RescuableCaller, error) {
	contract, err := bindRescuable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RescuableCaller{contract: contract}, nil
}

// NewRescuableTransactor creates a new write-only instance of Rescuable, bound to a specific deployed contract.
func NewRescuableTransactor(address common.Address, transactor bind.ContractTransactor) (*RescuableTransactor, error) {
	contract, err := bindRescuable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RescuableTransactor{contract: contract}, nil
}

// NewRescuableFilterer creates a new log filterer instance of Rescuable, bound to a specific deployed contract.
func NewRescuableFilterer(address common.Address, filterer bind.ContractFilterer) (*RescuableFilterer, error) {
	contract, err := bindRescuable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RescuableFilterer{contract: contract}, nil
}

// bindRescuable binds a generic wrapper to an already deployed contract.
func bindRescuable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RescuableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rescuable *RescuableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rescuable.Contract.RescuableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rescuable *RescuableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rescuable.Contract.RescuableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rescuable *RescuableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rescuable.Contract.RescuableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rescuable *RescuableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rescuable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rescuable *RescuableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rescuable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rescuable *RescuableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rescuable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rescuable *RescuableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rescuable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rescuable *RescuableSession) Owner() (common.Address, error) {
	return _Rescuable.Contract.Owner(&_Rescuable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rescuable *RescuableCallerSession) Owner() (common.Address, error) {
	return _Rescuable.Contract.Owner(&_Rescuable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Rescuable *RescuableCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rescuable.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Rescuable *RescuableSession) PendingOwner() (common.Address, error) {
	return _Rescuable.Contract.PendingOwner(&_Rescuable.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Rescuable *RescuableCallerSession) PendingOwner() (common.Address, error) {
	return _Rescuable.Contract.PendingOwner(&_Rescuable.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Rescuable *RescuableCaller) Rescuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rescuable.contract.Call(opts, &out, "rescuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Rescuable *RescuableSession) Rescuer() (common.Address, error) {
	return _Rescuable.Contract.Rescuer(&_Rescuable.CallOpts)
}

// Rescuer is a free data retrieval call binding the contract method 0x38a63183.
//
// Solidity: function rescuer() view returns(address)
func (_Rescuable *RescuableCallerSession) Rescuer() (common.Address, error) {
	return _Rescuable.Contract.Rescuer(&_Rescuable.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Rescuable *RescuableTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Rescuable *RescuableSession) AcceptOwnership() (*types.Transaction, error) {
	return _Rescuable.Contract.AcceptOwnership(&_Rescuable.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Rescuable *RescuableTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Rescuable.Contract.AcceptOwnership(&_Rescuable.TransactOpts)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Rescuable *RescuableTransactor) RescueERC20(opts *bind.TransactOpts, tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "rescueERC20", tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Rescuable *RescuableSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rescuable.Contract.RescueERC20(&_Rescuable.TransactOpts, tokenContract, to, amount)
}

// RescueERC20 is a paid mutator transaction binding the contract method 0xb2118a8d.
//
// Solidity: function rescueERC20(address tokenContract, address to, uint256 amount) returns()
func (_Rescuable *RescuableTransactorSession) RescueERC20(tokenContract common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rescuable.Contract.RescueERC20(&_Rescuable.TransactOpts, tokenContract, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rescuable *RescuableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rescuable *RescuableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.TransferOwnership(&_Rescuable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rescuable *RescuableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.TransferOwnership(&_Rescuable.TransactOpts, newOwner)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Rescuable *RescuableTransactor) UpdateRescuer(opts *bind.TransactOpts, newRescuer common.Address) (*types.Transaction, error) {
	return _Rescuable.contract.Transact(opts, "updateRescuer", newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Rescuable *RescuableSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.UpdateRescuer(&_Rescuable.TransactOpts, newRescuer)
}

// UpdateRescuer is a paid mutator transaction binding the contract method 0x2ab60045.
//
// Solidity: function updateRescuer(address newRescuer) returns()
func (_Rescuable *RescuableTransactorSession) UpdateRescuer(newRescuer common.Address) (*types.Transaction, error) {
	return _Rescuable.Contract.UpdateRescuer(&_Rescuable.TransactOpts, newRescuer)
}

// RescuableOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the Rescuable contract.
type RescuableOwnershipTransferStartedIterator struct {
	Event *RescuableOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *RescuableOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RescuableOwnershipTransferStarted)
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
		it.Event = new(RescuableOwnershipTransferStarted)
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
func (it *RescuableOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RescuableOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RescuableOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the Rescuable contract.
type RescuableOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RescuableOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RescuableOwnershipTransferStartedIterator{contract: _Rescuable.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *RescuableOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RescuableOwnershipTransferStarted)
				if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) ParseOwnershipTransferStarted(log types.Log) (*RescuableOwnershipTransferStarted, error) {
	event := new(RescuableOwnershipTransferStarted)
	if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RescuableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rescuable contract.
type RescuableOwnershipTransferredIterator struct {
	Event *RescuableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RescuableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RescuableOwnershipTransferred)
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
		it.Event = new(RescuableOwnershipTransferred)
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
func (it *RescuableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RescuableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RescuableOwnershipTransferred represents a OwnershipTransferred event raised by the Rescuable contract.
type RescuableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RescuableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RescuableOwnershipTransferredIterator{contract: _Rescuable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rescuable *RescuableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RescuableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rescuable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RescuableOwnershipTransferred)
				if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rescuable *RescuableFilterer) ParseOwnershipTransferred(log types.Log) (*RescuableOwnershipTransferred, error) {
	event := new(RescuableOwnershipTransferred)
	if err := _Rescuable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RescuableRescuerChangedIterator is returned from FilterRescuerChanged and is used to iterate over the raw logs and unpacked data for RescuerChanged events raised by the Rescuable contract.
type RescuableRescuerChangedIterator struct {
	Event *RescuableRescuerChanged // Event containing the contract specifics and raw log

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
func (it *RescuableRescuerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RescuableRescuerChanged)
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
		it.Event = new(RescuableRescuerChanged)
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
func (it *RescuableRescuerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RescuableRescuerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RescuableRescuerChanged represents a RescuerChanged event raised by the Rescuable contract.
type RescuableRescuerChanged struct {
	NewRescuer common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRescuerChanged is a free log retrieval operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Rescuable *RescuableFilterer) FilterRescuerChanged(opts *bind.FilterOpts, newRescuer []common.Address) (*RescuableRescuerChangedIterator, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _Rescuable.contract.FilterLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return &RescuableRescuerChangedIterator{contract: _Rescuable.contract, event: "RescuerChanged", logs: logs, sub: sub}, nil
}

// WatchRescuerChanged is a free log subscription operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Rescuable *RescuableFilterer) WatchRescuerChanged(opts *bind.WatchOpts, sink chan<- *RescuableRescuerChanged, newRescuer []common.Address) (event.Subscription, error) {

	var newRescuerRule []interface{}
	for _, newRescuerItem := range newRescuer {
		newRescuerRule = append(newRescuerRule, newRescuerItem)
	}

	logs, sub, err := _Rescuable.contract.WatchLogs(opts, "RescuerChanged", newRescuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RescuableRescuerChanged)
				if err := _Rescuable.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
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

// ParseRescuerChanged is a log parse operation binding the contract event 0xe475e580d85111348e40d8ca33cfdd74c30fe1655c2d8537a13abc10065ffa5a.
//
// Solidity: event RescuerChanged(address indexed newRescuer)
func (_Rescuable *RescuableFilterer) ParseRescuerChanged(log types.Log) (*RescuableRescuerChanged, error) {
	event := new(RescuableRescuerChanged)
	if err := _Rescuable.contract.UnpackLog(event, "RescuerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeERC20MetaData contains all meta data concerning the SafeERC20 contract.
var SafeERC20MetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122049d650fa5b28cfa0c973a0a960917819b7d0e4fa81f4a4ec4cde43af45791f1d64736f6c63430007060033",
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
	parsed, err := SafeERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220af9492fa74450900982792140a6f65f4b785f62d74325ae05a4f942a997ef82064736f6c63430007060033",
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
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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

// TypedMemViewMetaData contains all meta data concerning the TypedMemView contract.
var TypedMemViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f26be3fc": "NULL()",
	},
	Bin: "0x60cd610025600b82828239805160001a60731461001857fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c8063f26be3fc146038575b600080fd5b603e6073565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000009092168252519081900360200190f35b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000008156fea264697066735822122078d7da5a74bf463c97c4a8e3f4a12c2bd32aaf813b57e7960ca2f727f847ab5f64736f6c63430007060033",
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
	parsed, err := TypedMemViewMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
