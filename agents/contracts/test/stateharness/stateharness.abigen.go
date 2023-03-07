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
	Root        [32]byte
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a37f2e149b03db46d5109283c9460d513c3cc0ee2f48313ca09d6757284793cd64736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"castToState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structOriginState\",\"name\":\"_originState\",\"type\":\"tuple\"}],\"name\":\"equalToOrigin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structOriginState\",\"name\":\"_originState\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"formatOriginState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"_blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_timestamp\",\"type\":\"uint40\"}],\"name\":\"formatState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structSummitState\",\"name\":\"_summitState\",\"type\":\"tuple\"}],\"name\":\"formatSummitState\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"isState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"}],\"name\":\"leftLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"origin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"originState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structOriginState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"_blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"_timestamp\",\"type\":\"uint40\"}],\"name\":\"rightLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"root\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"subLeafs\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"toSummitState\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structSummitState\",\"name\":\"state\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e948e600": "blockNumber(bytes)",
		"5fed0261": "castToState(bytes)",
		"3047b5ea": "equalToOrigin(bytes,(bytes32,uint40,uint40))",
		"e8643e25": "formatOriginState((bytes32,uint40,uint40),uint32,uint32)",
		"365b4b67": "formatState(bytes32,uint32,uint32,uint40,uint40)",
		"461e44d3": "formatSummitState((bytes32,uint32,uint32,uint40,uint40))",
		"aa1e84de": "hash(bytes)",
		"aae6d884": "isState(bytes)",
		"edaa471d": "leftLeaf(bytes32,uint32)",
		"4e765004": "nonce(bytes)",
		"cb3eb0e1": "origin(bytes)",
		"1efcb714": "originState(bytes32)",
		"503d0bed": "rightLeaf(uint32,uint40,uint40)",
		"c2e9e208": "root(bytes)",
		"9aaa1826": "subLeafs(bytes)",
		"1c9aa222": "timestamp(bytes)",
		"9677fe8d": "toSummitState(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061192f806100206000396000f3fe608060405234801561001057600080fd5b506004361061011b5760003560e01c80639677fe8d116100b2578063c2e9e20811610081578063e8643e2511610066578063e8643e2514610326578063e948e60014610339578063edaa471d1461034c57600080fd5b8063c2e9e20814610300578063cb3eb0e11461031357600080fd5b80639677fe8d146102445780639aaa1826146102b2578063aa1e84de146102da578063aae6d884146102ed57600080fd5b8063461e44d3116100ee578063461e44d3146101d55780634e765004146101e8578063503d0bed146102105780635fed02611461023157600080fd5b80631c9aa222146101205780631efcb7141461014e5780633047b5ea14610192578063365b4b67146101b5575b600080fd5b61013361012e3660046113cc565b61035f565b60405164ffffffffff90911681526020015b60405180910390f35b61016161015c366004611401565b61037e565b604080518251815260208084015164ffffffffff908116918301919091529282015190921690820152606001610145565b6101a56101a036600461149c565b6103c4565b6040519015158152602001610145565b6101c86101c33660046114ff565b6103e6565b60405161014591906115c1565b6101c86101e33660046115d4565b61047c565b6101fb6101f63660046113cc565b610487565b60405163ffffffff9091168152602001610145565b61022361021e36600461165b565b6104a0565b604051908152602001610145565b6101c861023f3660046113cc565b61052d565b6102576102523660046113cc565b61054b565b6040516101459190600060a08201905082518252602083015163ffffffff80821660208501528060408601511660408501525050606083015164ffffffffff8082166060850152806080860151166080850152505092915050565b6102c56102c03660046113cc565b61058d565b60408051928352602083019190915201610145565b6102236102e83660046113cc565b6105b0565b6101a56102fb3660046113cc565b6105c9565b61022361030e3660046113cc565b6105f5565b6101fb6103213660046113cc565b61060e565b6101c861033436600461169e565b610627565b6101336103473660046113cc565b610634565b61022361035a3660046116d9565b61064d565b600061037861036d83610659565b62ffffff191661066c565b92915050565b604080516060810182526000808252602082018190529181019190915261037882604080516060810182529182524364ffffffffff908116602084015242169082015290565b60006103df826103d385610659565b62ffffff191690610682565b9392505050565b60408051602081018790527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168385015286901b1660448201527fffffffffff00000000000000000000000000000000000000000000000000000060d885811b8216604884015284901b16604d8201528151808203603201815260529091019091526060905b9695505050505050565b6060610378826106f3565b600061037861049583610659565b62ffffff19166107a1565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b166020808301919091527fffffffffff00000000000000000000000000000000000000000000000000000060d886811b8216602485015285901b1660298301528251808303600e018152602e90920190925280519101206000905b949350505050565b6060600061053a83610659565b90506103df62ffffff1982166107b7565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261037861058283610659565b62ffffff191661080a565b6000806105a761059c84610659565b62ffffff19166108b5565b91509150915091565b60006103786105be83610659565b62ffffff19166108f9565b60006103786105d8838361093d565b62ffffff191660181c6bffffffffffffffffffffffff1660321490565b600061037861060383610659565b62ffffff1916610961565b600061037861061c83610659565b62ffffff1916610976565b606061052584848461098c565b600061037861064283610659565b62ffffff1916610a2b565b60006103df8383610a41565b600061037861066783610aa2565b610aae565b600062ffffff1982166103df81602d6005610b3a565b805160009061069662ffffff198516610961565b1480156106c25750602082015164ffffffffff166106b962ffffff198516610a2b565b64ffffffffff16145b80156103df5750604082015164ffffffffff166106e462ffffff19851661066c565b64ffffffffff16149392505050565b60606103788260000151836020015184604001518560600151866080015160408051602081019690965260e094851b7fffffffff00000000000000000000000000000000000000000000000000000000908116878301529390941b909216604485015260d890811b7fffffffffff000000000000000000000000000000000000000000000000000000908116604886015291901b16604d830152805160328184030181526052909201905290565b600062ffffff1982166103df8160246004610b3a565b60606000806107d48460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060405191508192506107f98483602001610b6a565b508181016020016040529052919050565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915261084462ffffff198316610961565b815261085562ffffff198316610976565b63ffffffff16602082015261086f62ffffff1983166107a1565b63ffffffff16604082015261088962ffffff198316610a2b565b64ffffffffff1660608201526108a462ffffff19831661066c565b64ffffffffff166080820152919050565b60008062ffffff1983166108d96108ce82602485610d51565b62ffffff1916610d60565b92506108f16108ce62ffffff19831660246000610dae565b915050915091565b6000808061090c62ffffff1985166108b5565b6040805160208082019490945280820192909252805180830382018152606090920190528051910120949350505050565b81516000906020840161095864ffffffffff85168284610dec565b95945050505050565b600062ffffff1982166103df81836020610e33565b600062ffffff1982166103df8160206004610b3a565b82516020808501516040808701518151938401949094527fffffffff0000000000000000000000000000000000000000000000000000000060e087811b82168584015286901b1660448401527fffffffffff00000000000000000000000000000000000000000000000000000060d892831b811660488501529390911b909216604d820152815180820360320181526052909101909152606090610525565b600062ffffff1982166103df8160286005610b3a565b60008282604051602001610a8492919091825260e01b7fffffffff0000000000000000000000000000000000000000000000000000000016602082015260240190565b60405160208183030381529060405280519060200120905092915050565b6000610378828261093d565b6000610acb8260181c6bffffffffffffffffffffffff1660321490565b610b36576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600b60248201527f4e6f74206120737461746500000000000000000000000000000000000000000060448201526064015b60405180910390fd5b5090565b6000610b4782602061172b565b610b52906008611744565b60ff16610b60858585610e33565b901c949350505050565b600062ffffff1980841603610bdb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610b2d565b610be483610fe1565b610c4a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610b2d565b6000610c648460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610c7f8561101d565b6bffffffffffffffffffffffff169050600080604051915085821115610ca55760206060fd5b8386858560045afa905080610d16576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610b2d565b610d46610d2288611044565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b60006105258460008585611068565b600080610d6c8361101d565b6bffffffffffffffffffffffff1690506000610d968460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6000610525848485610dce8860181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610de69190611767565b85611068565b600080610df9838561177a565b9050604051811115610e09575060005b80600003610e1e5762ffffff199150506103df565b5050606092831b9190911790911b1760181b90565b60008160ff16600003610e48575060006103df565b610e608460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610e7b60ff84168561177a565b1115610efe57610ecb610e8d8561101d565b6bffffffffffffffffffffffff16610eb38660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff166110d3565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b2d91906115c1565b60208260ff161115610f6c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610b2d565b600882026000610f7b8661101d565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000610fec82611044565b64ffffffffff1664ffffffffff0361100657506000919050565b600061101183611141565b60405110199392505050565b60008061102c6060601861177a565b9290921c6bffffffffffffffffffffffff1692915050565b600080606061105481601861177a565b61105e919061177a565b9290921c92915050565b6000806110748661101d565b6bffffffffffffffffffffffff16905061108d86611141565b84611098878461177a565b6110a2919061177a565b11156110b55762ffffff19915050610525565b6110bf858261177a565b90506104728364ffffffffff168286610dec565b606060006110e08661117a565b91505060006110ee8661117a565b91505060006110fc8661117a565b915050600061110a8661117a565b91505083838383604051602001611124949392919061178d565b604051602081830303815290604052945050505050949350505050565b600061115b8260181c6bffffffffffffffffffffffff1690565b6111648361101d565b016bffffffffffffffffffffffff169050919050565b600080601f5b600f8160ff1611156111ed576000611199826008611744565b60ff1685901c90506111aa81611264565b61ffff16841793508160ff166010146111c557601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611180565b50600f5b60ff8160ff16101561125e57600061120a826008611744565b60ff1685901c905061121b81611264565b61ffff16831792508160ff1660001461123657601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016111f1565b50915091565b600061127660048360ff16901c611296565b60ff1661ffff919091161760081b61128d82611296565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f841691829081106112e3576112e36118ca565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261133257600080fd5b813567ffffffffffffffff8082111561134d5761134d6112f2565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611393576113936112f2565b816040528381528660208588010111156113ac57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000602082840312156113de57600080fd5b813567ffffffffffffffff8111156113f557600080fd5b61052584828501611321565b60006020828403121561141357600080fd5b5035919050565b803564ffffffffff8116811461142f57600080fd5b919050565b60006060828403121561144657600080fd5b6040516060810181811067ffffffffffffffff82111715611469576114696112f2565b6040528235815290508061147f6020840161141a565b60208201526114906040840161141a565b60408201525092915050565b600080608083850312156114af57600080fd5b823567ffffffffffffffff8111156114c657600080fd5b6114d285828601611321565b9250506114e28460208501611434565b90509250929050565b803563ffffffff8116811461142f57600080fd5b600080600080600060a0868803121561151757600080fd5b85359450611527602087016114eb565b9350611535604087016114eb565b92506115436060870161141a565b91506115516080870161141a565b90509295509295909350565b6000815180845260005b8181101561158357602081850181015186830182015201611567565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006103df602083018461155d565b600060a082840312156115e657600080fd5b60405160a0810181811067ffffffffffffffff82111715611609576116096112f2565b6040528235815261161c602084016114eb565b602082015261162d604084016114eb565b604082015261163e6060840161141a565b606082015261164f6080840161141a565b60808201529392505050565b60008060006060848603121561167057600080fd5b611679846114eb565b92506116876020850161141a565b91506116956040850161141a565b90509250925092565b600080600060a084860312156116b357600080fd5b6116bd8585611434565b92506116cb606085016114eb565b9150611695608085016114eb565b600080604083850312156116ec57600080fd5b823591506114e2602084016114eb565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff8281168282160390811115610378576103786116fc565b60ff8181168382160290811690818114611760576117606116fc565b5092915050565b81810381811115610378576103786116fc565b80820180821115610378576103786116fc565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d8201610472565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220c3947a969705f42ee5ec3708119538de1d76fec38e223a9c06abc746ecc8de0f64736f6c63430008110033",
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
// Solidity: function blockNumber(bytes _payload) pure returns(uint40)
func (_StateHarness *StateHarnessCaller) BlockNumber(opts *bind.CallOpts, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "blockNumber", _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes _payload) pure returns(uint40)
func (_StateHarness *StateHarnessSession) BlockNumber(_payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.BlockNumber(&_StateHarness.CallOpts, _payload)
}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes _payload) pure returns(uint40)
func (_StateHarness *StateHarnessCallerSession) BlockNumber(_payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.BlockNumber(&_StateHarness.CallOpts, _payload)
}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes _payload) view returns(bytes)
func (_StateHarness *StateHarnessCaller) CastToState(opts *bind.CallOpts, _payload []byte) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "castToState", _payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes _payload) view returns(bytes)
func (_StateHarness *StateHarnessSession) CastToState(_payload []byte) ([]byte, error) {
	return _StateHarness.Contract.CastToState(&_StateHarness.CallOpts, _payload)
}

// CastToState is a free data retrieval call binding the contract method 0x5fed0261.
//
// Solidity: function castToState(bytes _payload) view returns(bytes)
func (_StateHarness *StateHarnessCallerSession) CastToState(_payload []byte) ([]byte, error) {
	return _StateHarness.Contract.CastToState(&_StateHarness.CallOpts, _payload)
}

// EqualToOrigin is a free data retrieval call binding the contract method 0x3047b5ea.
//
// Solidity: function equalToOrigin(bytes _payload, (bytes32,uint40,uint40) _originState) pure returns(bool)
func (_StateHarness *StateHarnessCaller) EqualToOrigin(opts *bind.CallOpts, _payload []byte, _originState OriginState) (bool, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "equalToOrigin", _payload, _originState)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EqualToOrigin is a free data retrieval call binding the contract method 0x3047b5ea.
//
// Solidity: function equalToOrigin(bytes _payload, (bytes32,uint40,uint40) _originState) pure returns(bool)
func (_StateHarness *StateHarnessSession) EqualToOrigin(_payload []byte, _originState OriginState) (bool, error) {
	return _StateHarness.Contract.EqualToOrigin(&_StateHarness.CallOpts, _payload, _originState)
}

// EqualToOrigin is a free data retrieval call binding the contract method 0x3047b5ea.
//
// Solidity: function equalToOrigin(bytes _payload, (bytes32,uint40,uint40) _originState) pure returns(bool)
func (_StateHarness *StateHarnessCallerSession) EqualToOrigin(_payload []byte, _originState OriginState) (bool, error) {
	return _StateHarness.Contract.EqualToOrigin(&_StateHarness.CallOpts, _payload, _originState)
}

// FormatOriginState is a free data retrieval call binding the contract method 0xe8643e25.
//
// Solidity: function formatOriginState((bytes32,uint40,uint40) _originState, uint32 _origin, uint32 _nonce) pure returns(bytes)
func (_StateHarness *StateHarnessCaller) FormatOriginState(opts *bind.CallOpts, _originState OriginState, _origin uint32, _nonce uint32) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "formatOriginState", _originState, _origin, _nonce)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatOriginState is a free data retrieval call binding the contract method 0xe8643e25.
//
// Solidity: function formatOriginState((bytes32,uint40,uint40) _originState, uint32 _origin, uint32 _nonce) pure returns(bytes)
func (_StateHarness *StateHarnessSession) FormatOriginState(_originState OriginState, _origin uint32, _nonce uint32) ([]byte, error) {
	return _StateHarness.Contract.FormatOriginState(&_StateHarness.CallOpts, _originState, _origin, _nonce)
}

// FormatOriginState is a free data retrieval call binding the contract method 0xe8643e25.
//
// Solidity: function formatOriginState((bytes32,uint40,uint40) _originState, uint32 _origin, uint32 _nonce) pure returns(bytes)
func (_StateHarness *StateHarnessCallerSession) FormatOriginState(_originState OriginState, _origin uint32, _nonce uint32) ([]byte, error) {
	return _StateHarness.Contract.FormatOriginState(&_StateHarness.CallOpts, _originState, _origin, _nonce)
}

// FormatState is a free data retrieval call binding the contract method 0x365b4b67.
//
// Solidity: function formatState(bytes32 _root, uint32 _origin, uint32 _nonce, uint40 _blockNumber, uint40 _timestamp) pure returns(bytes)
func (_StateHarness *StateHarnessCaller) FormatState(opts *bind.CallOpts, _root [32]byte, _origin uint32, _nonce uint32, _blockNumber *big.Int, _timestamp *big.Int) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "formatState", _root, _origin, _nonce, _blockNumber, _timestamp)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatState is a free data retrieval call binding the contract method 0x365b4b67.
//
// Solidity: function formatState(bytes32 _root, uint32 _origin, uint32 _nonce, uint40 _blockNumber, uint40 _timestamp) pure returns(bytes)
func (_StateHarness *StateHarnessSession) FormatState(_root [32]byte, _origin uint32, _nonce uint32, _blockNumber *big.Int, _timestamp *big.Int) ([]byte, error) {
	return _StateHarness.Contract.FormatState(&_StateHarness.CallOpts, _root, _origin, _nonce, _blockNumber, _timestamp)
}

// FormatState is a free data retrieval call binding the contract method 0x365b4b67.
//
// Solidity: function formatState(bytes32 _root, uint32 _origin, uint32 _nonce, uint40 _blockNumber, uint40 _timestamp) pure returns(bytes)
func (_StateHarness *StateHarnessCallerSession) FormatState(_root [32]byte, _origin uint32, _nonce uint32, _blockNumber *big.Int, _timestamp *big.Int) ([]byte, error) {
	return _StateHarness.Contract.FormatState(&_StateHarness.CallOpts, _root, _origin, _nonce, _blockNumber, _timestamp)
}

// FormatSummitState is a free data retrieval call binding the contract method 0x461e44d3.
//
// Solidity: function formatSummitState((bytes32,uint32,uint32,uint40,uint40) _summitState) pure returns(bytes)
func (_StateHarness *StateHarnessCaller) FormatSummitState(opts *bind.CallOpts, _summitState SummitState) ([]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "formatSummitState", _summitState)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatSummitState is a free data retrieval call binding the contract method 0x461e44d3.
//
// Solidity: function formatSummitState((bytes32,uint32,uint32,uint40,uint40) _summitState) pure returns(bytes)
func (_StateHarness *StateHarnessSession) FormatSummitState(_summitState SummitState) ([]byte, error) {
	return _StateHarness.Contract.FormatSummitState(&_StateHarness.CallOpts, _summitState)
}

// FormatSummitState is a free data retrieval call binding the contract method 0x461e44d3.
//
// Solidity: function formatSummitState((bytes32,uint32,uint32,uint40,uint40) _summitState) pure returns(bytes)
func (_StateHarness *StateHarnessCallerSession) FormatSummitState(_summitState SummitState) ([]byte, error) {
	return _StateHarness.Contract.FormatSummitState(&_StateHarness.CallOpts, _summitState)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) Hash(opts *bind.CallOpts, _payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "hash", _payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _payload) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) Hash(_payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Hash(&_StateHarness.CallOpts, _payload)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) Hash(_payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Hash(&_StateHarness.CallOpts, _payload)
}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes _payload) pure returns(bool)
func (_StateHarness *StateHarnessCaller) IsState(opts *bind.CallOpts, _payload []byte) (bool, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "isState", _payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes _payload) pure returns(bool)
func (_StateHarness *StateHarnessSession) IsState(_payload []byte) (bool, error) {
	return _StateHarness.Contract.IsState(&_StateHarness.CallOpts, _payload)
}

// IsState is a free data retrieval call binding the contract method 0xaae6d884.
//
// Solidity: function isState(bytes _payload) pure returns(bool)
func (_StateHarness *StateHarnessCallerSession) IsState(_payload []byte) (bool, error) {
	return _StateHarness.Contract.IsState(&_StateHarness.CallOpts, _payload)
}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 _root, uint32 _origin) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) LeftLeaf(opts *bind.CallOpts, _root [32]byte, _origin uint32) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "leftLeaf", _root, _origin)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 _root, uint32 _origin) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) LeftLeaf(_root [32]byte, _origin uint32) ([32]byte, error) {
	return _StateHarness.Contract.LeftLeaf(&_StateHarness.CallOpts, _root, _origin)
}

// LeftLeaf is a free data retrieval call binding the contract method 0xedaa471d.
//
// Solidity: function leftLeaf(bytes32 _root, uint32 _origin) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) LeftLeaf(_root [32]byte, _origin uint32) ([32]byte, error) {
	return _StateHarness.Contract.LeftLeaf(&_StateHarness.CallOpts, _root, _origin)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes _payload) pure returns(uint32)
func (_StateHarness *StateHarnessCaller) Nonce(opts *bind.CallOpts, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "nonce", _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes _payload) pure returns(uint32)
func (_StateHarness *StateHarnessSession) Nonce(_payload []byte) (uint32, error) {
	return _StateHarness.Contract.Nonce(&_StateHarness.CallOpts, _payload)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes _payload) pure returns(uint32)
func (_StateHarness *StateHarnessCallerSession) Nonce(_payload []byte) (uint32, error) {
	return _StateHarness.Contract.Nonce(&_StateHarness.CallOpts, _payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes _payload) pure returns(uint32)
func (_StateHarness *StateHarnessCaller) Origin(opts *bind.CallOpts, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "origin", _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes _payload) pure returns(uint32)
func (_StateHarness *StateHarnessSession) Origin(_payload []byte) (uint32, error) {
	return _StateHarness.Contract.Origin(&_StateHarness.CallOpts, _payload)
}

// Origin is a free data retrieval call binding the contract method 0xcb3eb0e1.
//
// Solidity: function origin(bytes _payload) pure returns(uint32)
func (_StateHarness *StateHarnessCallerSession) Origin(_payload []byte) (uint32, error) {
	return _StateHarness.Contract.Origin(&_StateHarness.CallOpts, _payload)
}

// OriginState is a free data retrieval call binding the contract method 0x1efcb714.
//
// Solidity: function originState(bytes32 _root) view returns((bytes32,uint40,uint40) state)
func (_StateHarness *StateHarnessCaller) OriginState(opts *bind.CallOpts, _root [32]byte) (OriginState, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "originState", _root)

	if err != nil {
		return *new(OriginState), err
	}

	out0 := *abi.ConvertType(out[0], new(OriginState)).(*OriginState)

	return out0, err

}

// OriginState is a free data retrieval call binding the contract method 0x1efcb714.
//
// Solidity: function originState(bytes32 _root) view returns((bytes32,uint40,uint40) state)
func (_StateHarness *StateHarnessSession) OriginState(_root [32]byte) (OriginState, error) {
	return _StateHarness.Contract.OriginState(&_StateHarness.CallOpts, _root)
}

// OriginState is a free data retrieval call binding the contract method 0x1efcb714.
//
// Solidity: function originState(bytes32 _root) view returns((bytes32,uint40,uint40) state)
func (_StateHarness *StateHarnessCallerSession) OriginState(_root [32]byte) (OriginState, error) {
	return _StateHarness.Contract.OriginState(&_StateHarness.CallOpts, _root)
}

// RightLeaf is a free data retrieval call binding the contract method 0x503d0bed.
//
// Solidity: function rightLeaf(uint32 _nonce, uint40 _blockNumber, uint40 _timestamp) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) RightLeaf(opts *bind.CallOpts, _nonce uint32, _blockNumber *big.Int, _timestamp *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "rightLeaf", _nonce, _blockNumber, _timestamp)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RightLeaf is a free data retrieval call binding the contract method 0x503d0bed.
//
// Solidity: function rightLeaf(uint32 _nonce, uint40 _blockNumber, uint40 _timestamp) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) RightLeaf(_nonce uint32, _blockNumber *big.Int, _timestamp *big.Int) ([32]byte, error) {
	return _StateHarness.Contract.RightLeaf(&_StateHarness.CallOpts, _nonce, _blockNumber, _timestamp)
}

// RightLeaf is a free data retrieval call binding the contract method 0x503d0bed.
//
// Solidity: function rightLeaf(uint32 _nonce, uint40 _blockNumber, uint40 _timestamp) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) RightLeaf(_nonce uint32, _blockNumber *big.Int, _timestamp *big.Int) ([32]byte, error) {
	return _StateHarness.Contract.RightLeaf(&_StateHarness.CallOpts, _nonce, _blockNumber, _timestamp)
}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes _payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCaller) Root(opts *bind.CallOpts, _payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "root", _payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes _payload) pure returns(bytes32)
func (_StateHarness *StateHarnessSession) Root(_payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Root(&_StateHarness.CallOpts, _payload)
}

// Root is a free data retrieval call binding the contract method 0xc2e9e208.
//
// Solidity: function root(bytes _payload) pure returns(bytes32)
func (_StateHarness *StateHarnessCallerSession) Root(_payload []byte) ([32]byte, error) {
	return _StateHarness.Contract.Root(&_StateHarness.CallOpts, _payload)
}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes _payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessCaller) SubLeafs(opts *bind.CallOpts, _payload []byte) ([32]byte, [32]byte, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "subLeafs", _payload)

	if err != nil {
		return *new([32]byte), *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	out1 := *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return out0, out1, err

}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes _payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessSession) SubLeafs(_payload []byte) ([32]byte, [32]byte, error) {
	return _StateHarness.Contract.SubLeafs(&_StateHarness.CallOpts, _payload)
}

// SubLeafs is a free data retrieval call binding the contract method 0x9aaa1826.
//
// Solidity: function subLeafs(bytes _payload) pure returns(bytes32, bytes32)
func (_StateHarness *StateHarnessCallerSession) SubLeafs(_payload []byte) ([32]byte, [32]byte, error) {
	return _StateHarness.Contract.SubLeafs(&_StateHarness.CallOpts, _payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes _payload) pure returns(uint40)
func (_StateHarness *StateHarnessCaller) Timestamp(opts *bind.CallOpts, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "timestamp", _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes _payload) pure returns(uint40)
func (_StateHarness *StateHarnessSession) Timestamp(_payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.Timestamp(&_StateHarness.CallOpts, _payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes _payload) pure returns(uint40)
func (_StateHarness *StateHarnessCallerSession) Timestamp(_payload []byte) (*big.Int, error) {
	return _StateHarness.Contract.Timestamp(&_StateHarness.CallOpts, _payload)
}

// ToSummitState is a free data retrieval call binding the contract method 0x9677fe8d.
//
// Solidity: function toSummitState(bytes _payload) pure returns((bytes32,uint32,uint32,uint40,uint40) state)
func (_StateHarness *StateHarnessCaller) ToSummitState(opts *bind.CallOpts, _payload []byte) (SummitState, error) {
	var out []interface{}
	err := _StateHarness.contract.Call(opts, &out, "toSummitState", _payload)

	if err != nil {
		return *new(SummitState), err
	}

	out0 := *abi.ConvertType(out[0], new(SummitState)).(*SummitState)

	return out0, err

}

// ToSummitState is a free data retrieval call binding the contract method 0x9677fe8d.
//
// Solidity: function toSummitState(bytes _payload) pure returns((bytes32,uint32,uint32,uint40,uint40) state)
func (_StateHarness *StateHarnessSession) ToSummitState(_payload []byte) (SummitState, error) {
	return _StateHarness.Contract.ToSummitState(&_StateHarness.CallOpts, _payload)
}

// ToSummitState is a free data retrieval call binding the contract method 0x9677fe8d.
//
// Solidity: function toSummitState(bytes _payload) pure returns((bytes32,uint32,uint32,uint40,uint40) state)
func (_StateHarness *StateHarnessCallerSession) ToSummitState(_payload []byte) (SummitState, error) {
	return _StateHarness.Contract.ToSummitState(&_StateHarness.CallOpts, _payload)
}

// StateLibMetaData contains all meta data concerning the StateLib contract.
var StateLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220904c255ffaf38ed0dbb4595487ccec533224f338bbc9695356ce26a6755af3a964736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220125b438d5bdca38436074c73afdad82949d3678adcbc8b10c84aacc4dc89209d64736f6c63430008110033",
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
