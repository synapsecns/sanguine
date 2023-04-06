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

// OriginState is an auto generated low-level Go binding around an user-defined struct.
type OriginState struct {
	BlockNumber *big.Int
	Timestamp   *big.Int
}

// SummitState is an auto generated low-level Go binding around an user-defined struct.
type SummitState struct {
	Root        [32]byte
	Origin      uint32
	Nonce       uint32
	BlockNumber *big.Int
	Timestamp   *big.Int
}

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220edda4b2ed6a72f1889b70697425cbddfc15ffcdc5f233391a83321e043c9d3ed64736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structOriginState\",\"name\":\"originState_\",\"type\":\"tuple\"}],\"name\":\"equalToOrigin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"a\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"equals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structOriginState\",\"name\":\"originState_\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"}],\"name\":\"formatOriginState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"formatState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structSummitState\",\"name\":\"summitState\",\"type\":\"tuple\"}],\"name\":\"formatSummitState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"leaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"root_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin_\",\"type\":\"uint32\"}],\"name\":\"leftLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"originState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structOriginState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"rightLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"subLeafs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"toSummitState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structSummitState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e948e600": "blockNumber(bytes)",
		"5fed0261": "castToState(bytes)",
		"431f6acc": "equalToOrigin(bytes,(uint40,uint40))",
		"137e618a": "equals(bytes,bytes)",
		"3af50def": "formatOriginState((uint40,uint40),bytes32,uint32,uint32)",
		"365b4b67": "formatState(bytes32,uint32,uint32,uint40,uint40)",
		"461e44d3": "formatSummitState((bytes32,uint32,uint32,uint40,uint40))",
		"aae6d884": "isState(bytes)",
		"d7a7a72c": "leaf(bytes)",
		"edaa471d": "leftLeaf(bytes32,uint32)",
		"4e765004": "nonce(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"bcb513e5": "originState()",
		"503d0bed": "rightLeaf(uint32,uint40,uint40)",
		"c2e9e208": "root(bytes)",
		"9aaa1826": "subLeafs(bytes)",
		"1c9aa222": "timestamp(bytes)",
		"9677fe8d": "toSummitState(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b506119bd806100206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80639677fe8d116100b2578063c2e9e20811610081578063d7a7a72c11610066578063d7a7a72c1461033d578063e948e60014610350578063edaa471d1461036357600080fd5b8063c2e9e20814610317578063cb3eb0e11461032a57600080fd5b80639677fe8d146102415780639aaa1826146102af578063aae6d884146102d7578063bcb513e5146102ea57600080fd5b8063431f6acc116101095780634e765004116100ee5780634e765004146101e5578063503d0bed1461020d5780635fed02611461022e57600080fd5b8063431f6acc146101bf578063461e44d3146101d257600080fd5b8063137e618a1461013b5780631c9aa22214610163578063365b4b671461018c5780633af50def146101ac575b600080fd5b61014e610149366004611406565b610376565b60405190151581526020015b60405180910390f35b61017661017136600461146a565b6103a2565b60405164ffffffffff909116815260200161015a565b61019f61019a3660046114cd565b6103bb565b60405161015a919061158f565b61019f6101ba366004611600565b610451565b61014e6101cd36600461164e565b61046a565b61019f6101e036600461169d565b610485565b6101f86101f336600461146a565b610490565b60405163ffffffff909116815260200161015a565b61022061021b366004611724565b6104a9565b60405190815260200161015a565b61019f61023c36600461146a565b610538565b61025461024f36600461146a565b610556565b60405161015a9190600060a08201905082518252602083015163ffffffff80821660208501528060408601511660408501525050606083015164ffffffffff8082166060850152806080860151166080850152505092915050565b6102c26102bd36600461146a565b610598565b6040805192835260208301919091520161015a565b61014e6102e536600461146a565b6105bb565b6102f26105e7565b60408051825164ffffffffff908116825260209384015116928101929092520161015a565b61022061032536600461146a565b610623565b6101f861033836600461146a565b61063c565b61022061034b36600461146a565b610655565b61017661035e36600461146a565b61066e565b610220610371366004611767565b610687565b600061039961038483610693565b61038d85610693565b62ffffff1916906106a6565b90505b92915050565b600061039c6103b083610693565b62ffffff19166106d5565b60408051602081018790527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168385015286901b1660448201527fffffffffff00000000000000000000000000000000000000000000000000000060d885811b8216604884015284901b16604d8201528151808203603201815260529091019091526060905b9695505050505050565b606061045f858585856106eb565b90505b949350505050565b60006103998261047985610693565b62ffffff191690610783565b606061039c826107d7565b600061039c61049e83610693565b62ffffff1916610885565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b166020808301919091527fffffffffff00000000000000000000000000000000000000000000000000000060d886811b8216602485015285901b1660298301528251808303600e018152602e90920190925280519101206000905b90505b9392505050565b6060600061054583610693565b905061053162ffffff19821661089b565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261039c61058d83610693565b62ffffff19166108ee565b6000806105b26105a784610693565b62ffffff1916610999565b91509150915091565b600061039c6105ca83836109d2565b62ffffff191660181c6bffffffffffffffffffffffff1660321490565b604080518082019091526000808252602082015261061e604080518082019091524364ffffffffff90811682524216602082015290565b905090565b600061039c61063183610693565b62ffffff19166109f6565b600061039c61064a83610693565b62ffffff1916610a0b565b600061039c61066383610693565b62ffffff1916610a21565b600061039c61067c83610693565b62ffffff1916610a65565b60006103998383610a7b565b600061039c6106a183610adc565b610ae8565b60006106be62ffffff1983165b62ffffff1916610b74565b6106cd62ffffff1985166106b3565b149392505050565b600062ffffff19821661053181602d6005610bc2565b8351602080860151604080519283018790527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168584015286901b1660448401527fffffffffff00000000000000000000000000000000000000000000000000000060d894851b811660488501529190931b16604d82015281518082036032018152605290910190915260609061045f565b805160009064ffffffffff1661079e62ffffff198516610a65565b64ffffffffff161480156103995750602082015164ffffffffff166107c862ffffff1985166106d5565b64ffffffffff16149392505050565b606061039c8260000151836020015184604001518560600151866080015160408051602081019690965260e094851b7fffffffff00000000000000000000000000000000000000000000000000000000908116878301529390941b909216604485015260d890811b7fffffffffff000000000000000000000000000000000000000000000000000000908116604886015291901b16604d830152805160328184030181526052909201905290565b600062ffffff1982166105318160246004610bc2565b60606000806108b88460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060405191508192506108dd8483602001610bf2565b508181016020016040529052919050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261092862ffffff1983166109f6565b815261093962ffffff198316610a0b565b63ffffffff16602082015261095362ffffff198316610885565b63ffffffff16604082015261096d62ffffff198316610a65565b64ffffffffff16606082015261098862ffffff1983166106d5565b64ffffffffff166080820152919050565b60008062ffffff1983166109b26106b382602485610dd9565b92506109ca6106b362ffffff19831660246000610de8565b915050915091565b8151600090602084016109ed64ffffffffff85168284610e26565b95945050505050565b600062ffffff19821661053181836020610e6d565b600062ffffff1982166105318160206004610bc2565b60008080610a3462ffffff198516610999565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b600062ffffff1982166105318160286005610bc2565b60008282604051602001610abe92919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b600061039c82826109d2565b6000610b058260181c6bffffffffffffffffffffffff1660321490565b610b70576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f74206120737461746500000000000000000000000000000000000000000060448201526064015b60405180910390fd5b5090565b600080610b808361101b565b6bffffffffffffffffffffffff1690506000610baa8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6000610bcf8260206117b9565b610bda9060086117d2565b60ff16610be8858585610e6d565b901c949350505050565b600062ffffff1980841603610c63576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610b67565b610c6c83611042565b610cd2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610b67565b6000610cec8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610d078561101b565b6bffffffffffffffffffffffff169050600080604051915085821115610d2d5760206060fd5b8386858560045afa905080610d9e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610b67565b610dce610daa8861107e565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b600061052e84600085856110a2565b600061052e848485610e088860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610e2091906117f5565b856110a2565b600080610e338385611808565b9050604051811115610e43575060005b80600003610e585762ffffff19915050610531565b5050606092831b9190911790911b1760181b90565b60008160ff16600003610e8257506000610531565b610e9a8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610eb560ff841685611808565b1115610f3857610f05610ec78561101b565b6bffffffffffffffffffffffff16610eed8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff1661110d565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b67919061158f565b60208260ff161115610fa6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610b67565b600882026000610fb58661101b565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b60008061102a60606018611808565b9290921c6bffffffffffffffffffffffff1692915050565b600061104d8261107e565b64ffffffffff1664ffffffffff0361106757506000919050565b60006110728361117b565b60405110199392505050565b600080606061108e816018611808565b6110989190611808565b9290921c92915050565b6000806110ae8661101b565b6bffffffffffffffffffffffff1690506110c78661117b565b846110d28784611808565b6110dc9190611808565b11156110ef5762ffffff19915050610462565b6110f98582611808565b90506104478364ffffffffff168286610e26565b6060600061111a866111b4565b9150506000611128866111b4565b9150506000611136866111b4565b9150506000611144866111b4565b9150508383838360405160200161115e949392919061181b565b604051602081830303815290604052945050505050949350505050565b60006111958260181c6bffffffffffffffffffffffff1690565b61119e8361101b565b016bffffffffffffffffffffffff169050919050565b600080601f5b600f8160ff1611156112275760006111d38260086117d2565b60ff1685901c90506111e48161129e565b61ffff16841793508160ff166010146111ff57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016111ba565b50600f5b60ff8160ff1610156112985760006112448260086117d2565b60ff1685901c90506112558161129e565b61ffff16831792508160ff1660001461127057601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0161122b565b50915091565b60006112b060048360ff16901c6112d0565b60ff1661ffff919091161760081b6112c7826112d0565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f8416918290811061131d5761131d611958565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261136c57600080fd5b813567ffffffffffffffff808211156113875761138761132c565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019082821181831017156113cd576113cd61132c565b816040528381528660208588010111156113e657600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561141957600080fd5b823567ffffffffffffffff8082111561143157600080fd5b61143d8683870161135b565b9350602085013591508082111561145357600080fd5b506114608582860161135b565b9150509250929050565b60006020828403121561147c57600080fd5b813567ffffffffffffffff81111561149357600080fd5b6104628482850161135b565b803563ffffffff811681146114b357600080fd5b919050565b803564ffffffffff811681146114b357600080fd5b600080600080600060a086880312156114e557600080fd5b853594506114f56020870161149f565b93506115036040870161149f565b9250611511606087016114b8565b915061151f608087016114b8565b90509295509295909350565b6000815180845260005b8181101561155157602081850181015186830182015201611535565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000610399602083018461152b565b6000604082840312156115b457600080fd5b6040516040810181811067ffffffffffffffff821117156115d7576115d761132c565b6040529050806115e6836114b8565b81526115f4602084016114b8565b60208201525092915050565b60008060008060a0858703121561161657600080fd5b61162086866115a2565b9350604085013592506116356060860161149f565b91506116436080860161149f565b905092959194509250565b6000806060838503121561166157600080fd5b823567ffffffffffffffff81111561167857600080fd5b6116848582860161135b565b92505061169484602085016115a2565b90509250929050565b600060a082840312156116af57600080fd5b60405160a0810181811067ffffffffffffffff821117156116d2576116d261132c565b604052823581526116e56020840161149f565b60208201526116f66040840161149f565b6040820152611707606084016114b8565b6060820152611718608084016114b8565b60808201529392505050565b60008060006060848603121561173957600080fd5b6117428461149f565b9250611750602085016114b8565b915061175e604085016114b8565b90509250925092565b6000806040838503121561177a57600080fd5b823591506116946020840161149f565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff828116828216039081111561039c5761039c61178a565b60ff81811683821602908116908181146117ee576117ee61178a565b5092915050565b8181038181111561039c5761039c61178a565b8082018082111561039c5761039c61178a565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610447565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea264697066735822122029058ca778825a6fa89f5c9887c188238ffc66b84c8cfce08d592f378ba41f5f64736f6c63430008110033",
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

// EqualToOrigin is a free data retrieval call binding the contract method 0x431f6acc.
//
// Solidity: function equalToOrigin(bytes payload, (uint40,uint40) originState_) pure returns(bool)
func (_StateHarness *StateHarnessCaller) EqualToOrigin(opts *bind.CallOpts, payload []byte, originState_ OriginState) (bool, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "equalToOrigin", payload, originState_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EqualToOrigin is a free data retrieval call binding the contract method 0x431f6acc.
//
// Solidity: function equalToOrigin(bytes payload, (uint40,uint40) originState_) pure returns(bool)
func (_StateHarness *StateHarnessSession) EqualToOrigin(payload []byte, originState_ OriginState) (bool, error) {
	return _StateHarness.Contract.EqualToOrigin(&_StateHarness.CallOpts, payload, originState_)
}

// EqualToOrigin is a free data retrieval call binding the contract method 0x431f6acc.
//
// Solidity: function equalToOrigin(bytes payload, (uint40,uint40) originState_) pure returns(bool)
func (_StateHarness *StateHarnessCallerSession) EqualToOrigin(payload []byte, originState_ OriginState) (bool, error) {
	return _StateHarness.Contract.EqualToOrigin(&_StateHarness.CallOpts, payload, originState_)
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

// FormatOriginState is a free data retrieval call binding the contract method 0x3af50def.
//
// Solidity: function formatOriginState((uint40,uint40) originState_, bytes32 root_, uint32 origin_, uint32 nonce_) pure returns(bytes)
func (_StateHarness *StateHarnessCaller) FormatOriginState(opts *bind.CallOpts, originState_ OriginState, root_ [32]byte, origin_ uint32, nonce_ uint32) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "formatOriginState", originState_, root_, origin_, nonce_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatOriginState is a free data retrieval call binding the contract method 0x3af50def.
//
// Solidity: function formatOriginState((uint40,uint40) originState_, bytes32 root_, uint32 origin_, uint32 nonce_) pure returns(bytes)
func (_StateHarness *StateHarnessSession) FormatOriginState(originState_ OriginState, root_ [32]byte, origin_ uint32, nonce_ uint32) ([]byte, error) {
	return _StateHarness.Contract.FormatOriginState(&_StateHarness.CallOpts, originState_, root_, origin_, nonce_)
}

// FormatOriginState is a free data retrieval call binding the contract method 0x3af50def.
//
// Solidity: function formatOriginState((uint40,uint40) originState_, bytes32 root_, uint32 origin_, uint32 nonce_) pure returns(bytes)
func (_StateHarness *StateHarnessCallerSession) FormatOriginState(originState_ OriginState, root_ [32]byte, origin_ uint32, nonce_ uint32) ([]byte, error) {
	return _StateHarness.Contract.FormatOriginState(&_StateHarness.CallOpts, originState_, root_, origin_, nonce_)
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

// FormatSummitState is a free data retrieval call binding the contract method 0x461e44d3.
//
// Solidity: function formatSummitState((bytes32,uint32,uint32,uint40,uint40) summitState) pure returns(bytes)
func (_StateHarness *StateHarnessCaller) FormatSummitState(opts *bind.CallOpts, summitState SummitState) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "formatSummitState", summitState)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatSummitState is a free data retrieval call binding the contract method 0x461e44d3.
//
// Solidity: function formatSummitState((bytes32,uint32,uint32,uint40,uint40) summitState) pure returns(bytes)
func (_StateHarness *StateHarnessSession) FormatSummitState(summitState SummitState) ([]byte, error) {
	return _StateHarness.Contract.FormatSummitState(&_StateHarness.CallOpts, summitState)
}

// FormatSummitState is a free data retrieval call binding the contract method 0x461e44d3.
//
// Solidity: function formatSummitState((bytes32,uint32,uint32,uint40,uint40) summitState) pure returns(bytes)
func (_StateHarness *StateHarnessCallerSession) FormatSummitState(summitState SummitState) ([]byte, error) {
	return _StateHarness.Contract.FormatSummitState(&_StateHarness.CallOpts, summitState)
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

// OriginState is a free data retrieval call binding the contract method 0xbcb513e5.
//
// Solidity: function originState() view returns((uint40,uint40) state)
func (_StateHarness *StateHarnessCaller) OriginState(opts *bind.CallOpts) (OriginState, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "originState")

	if err != nil {
		return *new(OriginState), err
	}

	out0 := *abi.ConvertType(out[0], new(OriginState)).(*OriginState)

	return out0, err

}

// OriginState is a free data retrieval call binding the contract method 0xbcb513e5.
//
// Solidity: function originState() view returns((uint40,uint40) state)
func (_StateHarness *StateHarnessSession) OriginState() (OriginState, error) {
	return _StateHarness.Contract.OriginState(&_StateHarness.CallOpts)
}

// OriginState is a free data retrieval call binding the contract method 0xbcb513e5.
//
// Solidity: function originState() view returns((uint40,uint40) state)
func (_StateHarness *StateHarnessCallerSession) OriginState() (OriginState, error) {
	return _StateHarness.Contract.OriginState(&_StateHarness.CallOpts)
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

// ToSummitState is a free data retrieval call binding the contract method 0x9677fe8d.
//
// Solidity: function toSummitState(bytes payload) pure returns((bytes32,uint32,uint32,uint40,uint40) state)
func (_StateHarness *StateHarnessCaller) ToSummitState(opts *bind.CallOpts, payload []byte) (SummitState, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "toSummitState", payload)

	if err != nil {
		return *new(SummitState), err
	}

	out0 := *abi.ConvertType(out[0], new(SummitState)).(*SummitState)

	return out0, err

}

// ToSummitState is a free data retrieval call binding the contract method 0x9677fe8d.
//
// Solidity: function toSummitState(bytes payload) pure returns((bytes32,uint32,uint32,uint40,uint40) state)
func (_StateHarness *StateHarnessSession) ToSummitState(payload []byte) (SummitState, error) {
	return _StateHarness.Contract.ToSummitState(&_StateHarness.CallOpts, payload)
}

// ToSummitState is a free data retrieval call binding the contract method 0x9677fe8d.
//
// Solidity: function toSummitState(bytes payload) pure returns((bytes32,uint32,uint32,uint40,uint40) state)
func (_StateHarness *StateHarnessCallerSession) ToSummitState(payload []byte) (SummitState, error) {
	return _StateHarness.Contract.ToSummitState(&_StateHarness.CallOpts, payload)
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220752711dd99295a34893c6f499317df6aa79c45b312bc275f71a1f157670cbf5f64736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220f5bbbd9fe28e8a26f1d6fb434a39001acb77a346c94a3ea9731863276d25e68564736f6c63430008110033",
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
