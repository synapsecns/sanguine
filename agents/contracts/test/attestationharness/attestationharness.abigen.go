// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package attestationharness

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

// AttestationHarnessMetaData contains all meta data concerning the AttestationHarness contract.
var AttestationHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"IndexedTooMuch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OccupiedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PrecompileOutOfGas\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnallocatedMemory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnformattedAttestation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ViewOverrun\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"snapGasHash_\",\"type\":\"bytes32\"}],\"name\":\"dataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"dataHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"snapRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"formatAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hashInvalid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hashValid\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"snapRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"e948e600": "blockNumber(bytes)",
		"11f6389c": "castToAttestation(bytes)",
		"c4e1fa9c": "dataHash(bytes)",
		"717b6ca9": "dataHash(bytes32,bytes32)",
		"86bf2253": "formatAttestation(bytes32,bytes32,uint32,uint40,uint40)",
		"60cf3bf0": "hashInvalid(bytes)",
		"730dbf63": "hashValid(bytes)",
		"3ae7034d": "isAttestation(bytes)",
		"4e765004": "nonce(bytes)",
		"91bacf60": "snapRoot(bytes)",
		"1c9aa222": "timestamp(bytes)",
	},
	Bin: "0x608060405234801561001057600080fd5b50610a1f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c8063717b6ca91161008157806391bacf601161005b57806391bacf601461024b578063c4e1fa9c1461025e578063e948e6001461027157600080fd5b8063717b6ca91461018c578063730dbf631461019f57806386bf2253146101b257600080fd5b80633ae7034d116100b25780633ae7034d146101205780634e7650041461014357806360cf3bf01461016b57600080fd5b806311f6389c146100ce5780631c9aa222146100f7575b600080fd5b6100e16100dc3660046107d6565b610284565b6040516100ee91906108a5565b60405180910390f35b61010a6101053660046107d6565b6102a3565b60405164ffffffffff90911681526020016100ee565b61013361012e3660046107d6565b6102bc565b60405190151581526020016100ee565b6101566101513660046107d6565b6102e2565b60405163ffffffff90911681526020016100ee565b61017e6101793660046107d6565b6102f5565b6040519081526020016100ee565b61017e61019a366004610911565b610308565b61017e6101ad3660046107d6565b610314565b6100e16101c036600461094d565b6040805160208101969096528581019490945260e09290921b7fffffffff0000000000000000000000000000000000000000000000000000000016606085015260d890811b7fffffffffff000000000000000000000000000000000000000000000000000000908116606486015291901b1660698301528051604e818403018152606e909201905290565b61017e6102593660046107d6565b610327565b61017e61026c3660046107d6565b61033a565b61010a61027f3660046107d6565b61034d565b6060600061029183610360565b905061029c81610373565b9392505050565b60006102b66102b183610360565b6103d0565b92915050565b60006102b66102ca836103e2565b6fffffffffffffffffffffffffffffffff16604e1490565b60006102b66102f083610360565b6103fd565b60006102b661030383610360565b61040c565b600061029c838361043a565b60006102b661032283610360565b610470565b60006102b661033583610360565b61049c565b60006102b661034883610360565b6104ad565b60006102b661035b83610360565b6104bb565b60006102b661036e836103e2565b6104ca565b60405180610384836020830161051c565b506fffffffffffffffffffffffffffffffff83166000601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168301602001604052509052919050565b60006102b660496005845b91906105cb565b8051600090602083016103f581836105ec565b949350505050565b60006102b660406004846103db565b60006102b67fccfadb9c399e4e4257b6d0c3f92e1f9a9c00b1802b55a2f7d511702faa769090835b9061064f565b60408051602081018490529081018290526000906060015b60405160208183030381529060405280519060200120905092915050565b60006102b67f3464bf887f210604c594030208052a323ac6628785466262d75241769120164183610434565b60006102b6816020845b9190610672565b60006102b6602080846104a6565b60006102b660446005846103db565b6000604e6fffffffffffffffffffffffffffffffff831614610518576040517feb92662c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5090565b6040516000906fffffffffffffffffffffffffffffffff841690608085901c9080851015610576576040517f4b2a158c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008386858560045afa9050806105b9576040517f7c7d772f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608086901b8417979650505050505050565b6000806105d9858585610672565b602084900360031b1c9150509392505050565b6000806105f983856109af565b9050604051811115610609575060005b80600003610643576040517f10bef38600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b608084901b83176103f5565b60008161065b8461077c565b604080516020810193909352820152606001610452565b6000816000036106845750600061029c565b60208211156106bf576040517f31d784a800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6fffffffffffffffffffffffffffffffff84166106dc83856109af565b1115610714576040517fa3b99ded00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600382901b60006107258660801c90565b909401517f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff929092019190911d16949350505050565b6000806107898360801c90565b6fffffffffffffffffffffffffffffffff9390931690922092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156107e857600080fd5b813567ffffffffffffffff8082111561080057600080fd5b818401915084601f83011261081457600080fd5b813581811115610826576108266107a7565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561086c5761086c6107a7565b8160405282815287602084870101111561088557600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156108d2578581018301518582016040015282016108b6565b5060006040828601015260407fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8301168501019250505092915050565b6000806040838503121561092457600080fd5b50508035926020909101359150565b803564ffffffffff8116811461094857600080fd5b919050565b600080600080600060a0868803121561096557600080fd5b8535945060208601359350604086013563ffffffff8116811461098757600080fd5b925061099560608701610933565b91506109a360808701610933565b90509295509295909350565b808201808211156102b6577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fdfea2646970667358221220046372e6e1f70e205a39db92fa1357c72822d5893d9ec1277bdf30790af27c0364736f6c63430008110033",
}

// AttestationHarnessABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationHarnessMetaData.ABI instead.
var AttestationHarnessABI = AttestationHarnessMetaData.ABI

// Deprecated: Use AttestationHarnessMetaData.Sigs instead.
// AttestationHarnessFuncSigs maps the 4-byte function signature to its string representation.
var AttestationHarnessFuncSigs = AttestationHarnessMetaData.Sigs

// AttestationHarnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationHarnessMetaData.Bin instead.
var AttestationHarnessBin = AttestationHarnessMetaData.Bin

// DeployAttestationHarness deploys a new Ethereum contract, binding an instance of AttestationHarness to it.
func DeployAttestationHarness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttestationHarness, error) {
	parsed, err := AttestationHarnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationHarnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttestationHarness{AttestationHarnessCaller: AttestationHarnessCaller{contract: contract}, AttestationHarnessTransactor: AttestationHarnessTransactor{contract: contract}, AttestationHarnessFilterer: AttestationHarnessFilterer{contract: contract}}, nil
}

// AttestationHarness is an auto generated Go binding around an Ethereum contract.
type AttestationHarness struct {
	AttestationHarnessCaller     // Read-only binding to the contract
	AttestationHarnessTransactor // Write-only binding to the contract
	AttestationHarnessFilterer   // Log filterer for contract events
}

// AttestationHarnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationHarnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHarnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationHarnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHarnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationHarnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationHarnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationHarnessSession struct {
	Contract     *AttestationHarness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AttestationHarnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationHarnessCallerSession struct {
	Contract *AttestationHarnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AttestationHarnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationHarnessTransactorSession struct {
	Contract     *AttestationHarnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AttestationHarnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationHarnessRaw struct {
	Contract *AttestationHarness // Generic contract binding to access the raw methods on
}

// AttestationHarnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationHarnessCallerRaw struct {
	Contract *AttestationHarnessCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationHarnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationHarnessTransactorRaw struct {
	Contract *AttestationHarnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationHarness creates a new instance of AttestationHarness, bound to a specific deployed contract.
func NewAttestationHarness(address common.Address, backend bind.ContractBackend) (*AttestationHarness, error) {
	contract, err := bindAttestationHarness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationHarness{AttestationHarnessCaller: AttestationHarnessCaller{contract: contract}, AttestationHarnessTransactor: AttestationHarnessTransactor{contract: contract}, AttestationHarnessFilterer: AttestationHarnessFilterer{contract: contract}}, nil
}

// NewAttestationHarnessCaller creates a new read-only instance of AttestationHarness, bound to a specific deployed contract.
func NewAttestationHarnessCaller(address common.Address, caller bind.ContractCaller) (*AttestationHarnessCaller, error) {
	contract, err := bindAttestationHarness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationHarnessCaller{contract: contract}, nil
}

// NewAttestationHarnessTransactor creates a new write-only instance of AttestationHarness, bound to a specific deployed contract.
func NewAttestationHarnessTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationHarnessTransactor, error) {
	contract, err := bindAttestationHarness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationHarnessTransactor{contract: contract}, nil
}

// NewAttestationHarnessFilterer creates a new log filterer instance of AttestationHarness, bound to a specific deployed contract.
func NewAttestationHarnessFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationHarnessFilterer, error) {
	contract, err := bindAttestationHarness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationHarnessFilterer{contract: contract}, nil
}

// bindAttestationHarness binds a generic wrapper to an already deployed contract.
func bindAttestationHarness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationHarnessABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationHarness *AttestationHarnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationHarness.Contract.AttestationHarnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationHarness *AttestationHarnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationHarness.Contract.AttestationHarnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationHarness *AttestationHarnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationHarness.Contract.AttestationHarnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationHarness *AttestationHarnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationHarness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationHarness *AttestationHarnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationHarness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationHarness *AttestationHarnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationHarness.Contract.contract.Transact(opts, method, params...)
}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_AttestationHarness *AttestationHarnessCaller) BlockNumber(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "blockNumber", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_AttestationHarness *AttestationHarnessSession) BlockNumber(payload []byte) (*big.Int, error) {
	return _AttestationHarness.Contract.BlockNumber(&_AttestationHarness.CallOpts, payload)
}

// BlockNumber is a free data retrieval call binding the contract method 0xe948e600.
//
// Solidity: function blockNumber(bytes payload) pure returns(uint40)
func (_AttestationHarness *AttestationHarnessCallerSession) BlockNumber(payload []byte) (*big.Int, error) {
	return _AttestationHarness.Contract.BlockNumber(&_AttestationHarness.CallOpts, payload)
}

// CastToAttestation is a free data retrieval call binding the contract method 0x11f6389c.
//
// Solidity: function castToAttestation(bytes payload) view returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) CastToAttestation(opts *bind.CallOpts, payload []byte) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "castToAttestation", payload)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// CastToAttestation is a free data retrieval call binding the contract method 0x11f6389c.
//
// Solidity: function castToAttestation(bytes payload) view returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) CastToAttestation(payload []byte) ([]byte, error) {
	return _AttestationHarness.Contract.CastToAttestation(&_AttestationHarness.CallOpts, payload)
}

// CastToAttestation is a free data retrieval call binding the contract method 0x11f6389c.
//
// Solidity: function castToAttestation(bytes payload) view returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) CastToAttestation(payload []byte) ([]byte, error) {
	return _AttestationHarness.Contract.CastToAttestation(&_AttestationHarness.CallOpts, payload)
}

// DataHash is a free data retrieval call binding the contract method 0x717b6ca9.
//
// Solidity: function dataHash(bytes32 agentRoot_, bytes32 snapGasHash_) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) DataHash(opts *bind.CallOpts, agentRoot_ [32]byte, snapGasHash_ [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "dataHash", agentRoot_, snapGasHash_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataHash is a free data retrieval call binding the contract method 0x717b6ca9.
//
// Solidity: function dataHash(bytes32 agentRoot_, bytes32 snapGasHash_) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) DataHash(agentRoot_ [32]byte, snapGasHash_ [32]byte) ([32]byte, error) {
	return _AttestationHarness.Contract.DataHash(&_AttestationHarness.CallOpts, agentRoot_, snapGasHash_)
}

// DataHash is a free data retrieval call binding the contract method 0x717b6ca9.
//
// Solidity: function dataHash(bytes32 agentRoot_, bytes32 snapGasHash_) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) DataHash(agentRoot_ [32]byte, snapGasHash_ [32]byte) ([32]byte, error) {
	return _AttestationHarness.Contract.DataHash(&_AttestationHarness.CallOpts, agentRoot_, snapGasHash_)
}

// DataHash0 is a free data retrieval call binding the contract method 0xc4e1fa9c.
//
// Solidity: function dataHash(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) DataHash0(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "dataHash0", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DataHash0 is a free data retrieval call binding the contract method 0xc4e1fa9c.
//
// Solidity: function dataHash(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) DataHash0(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.DataHash0(&_AttestationHarness.CallOpts, payload)
}

// DataHash0 is a free data retrieval call binding the contract method 0xc4e1fa9c.
//
// Solidity: function dataHash(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) DataHash0(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.DataHash0(&_AttestationHarness.CallOpts, payload)
}

// FormatAttestation is a free data retrieval call binding the contract method 0x86bf2253.
//
// Solidity: function formatAttestation(bytes32 snapRoot_, bytes32 dataHash_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatAttestation(opts *bind.CallOpts, snapRoot_ [32]byte, dataHash_ [32]byte, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatAttestation", snapRoot_, dataHash_, nonce_, blockNumber_, timestamp_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatAttestation is a free data retrieval call binding the contract method 0x86bf2253.
//
// Solidity: function formatAttestation(bytes32 snapRoot_, bytes32 dataHash_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatAttestation(snapRoot_ [32]byte, dataHash_ [32]byte, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, snapRoot_, dataHash_, nonce_, blockNumber_, timestamp_)
}

// FormatAttestation is a free data retrieval call binding the contract method 0x86bf2253.
//
// Solidity: function formatAttestation(bytes32 snapRoot_, bytes32 dataHash_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatAttestation(snapRoot_ [32]byte, dataHash_ [32]byte, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, snapRoot_, dataHash_, nonce_, blockNumber_, timestamp_)
}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) HashInvalid(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "hashInvalid", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) HashInvalid(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.HashInvalid(&_AttestationHarness.CallOpts, payload)
}

// HashInvalid is a free data retrieval call binding the contract method 0x60cf3bf0.
//
// Solidity: function hashInvalid(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) HashInvalid(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.HashInvalid(&_AttestationHarness.CallOpts, payload)
}

// HashValid is a free data retrieval call binding the contract method 0x730dbf63.
//
// Solidity: function hashValid(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) HashValid(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "hashValid", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashValid is a free data retrieval call binding the contract method 0x730dbf63.
//
// Solidity: function hashValid(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) HashValid(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.HashValid(&_AttestationHarness.CallOpts, payload)
}

// HashValid is a free data retrieval call binding the contract method 0x730dbf63.
//
// Solidity: function hashValid(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) HashValid(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.HashValid(&_AttestationHarness.CallOpts, payload)
}

// IsAttestation is a free data retrieval call binding the contract method 0x3ae7034d.
//
// Solidity: function isAttestation(bytes payload) pure returns(bool)
func (_AttestationHarness *AttestationHarnessCaller) IsAttestation(opts *bind.CallOpts, payload []byte) (bool, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "isAttestation", payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAttestation is a free data retrieval call binding the contract method 0x3ae7034d.
//
// Solidity: function isAttestation(bytes payload) pure returns(bool)
func (_AttestationHarness *AttestationHarnessSession) IsAttestation(payload []byte) (bool, error) {
	return _AttestationHarness.Contract.IsAttestation(&_AttestationHarness.CallOpts, payload)
}

// IsAttestation is a free data retrieval call binding the contract method 0x3ae7034d.
//
// Solidity: function isAttestation(bytes payload) pure returns(bool)
func (_AttestationHarness *AttestationHarnessCallerSession) IsAttestation(payload []byte) (bool, error) {
	return _AttestationHarness.Contract.IsAttestation(&_AttestationHarness.CallOpts, payload)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCaller) Nonce(opts *bind.CallOpts, payload []byte) (uint32, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "nonce", payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessSession) Nonce(payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.Nonce(&_AttestationHarness.CallOpts, payload)
}

// Nonce is a free data retrieval call binding the contract method 0x4e765004.
//
// Solidity: function nonce(bytes payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCallerSession) Nonce(payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.Nonce(&_AttestationHarness.CallOpts, payload)
}

// SnapRoot is a free data retrieval call binding the contract method 0x91bacf60.
//
// Solidity: function snapRoot(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) SnapRoot(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "snapRoot", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SnapRoot is a free data retrieval call binding the contract method 0x91bacf60.
//
// Solidity: function snapRoot(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) SnapRoot(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.SnapRoot(&_AttestationHarness.CallOpts, payload)
}

// SnapRoot is a free data retrieval call binding the contract method 0x91bacf60.
//
// Solidity: function snapRoot(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) SnapRoot(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.SnapRoot(&_AttestationHarness.CallOpts, payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_AttestationHarness *AttestationHarnessCaller) Timestamp(opts *bind.CallOpts, payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "timestamp", payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_AttestationHarness *AttestationHarnessSession) Timestamp(payload []byte) (*big.Int, error) {
	return _AttestationHarness.Contract.Timestamp(&_AttestationHarness.CallOpts, payload)
}

// Timestamp is a free data retrieval call binding the contract method 0x1c9aa222.
//
// Solidity: function timestamp(bytes payload) pure returns(uint40)
func (_AttestationHarness *AttestationHarnessCallerSession) Timestamp(payload []byte) (*big.Int, error) {
	return _AttestationHarness.Contract.Timestamp(&_AttestationHarness.CallOpts, payload)
}

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220f6f5d82dd1114b7af0b473f769ad1e79d84a001d2f1f19e4aebe17aa5b27156e64736f6c63430008110033",
}

// AttestationLibABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationLibMetaData.ABI instead.
var AttestationLibABI = AttestationLibMetaData.ABI

// AttestationLibBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationLibMetaData.Bin instead.
var AttestationLibBin = AttestationLibMetaData.Bin

// DeployAttestationLib deploys a new Ethereum contract, binding an instance of AttestationLib to it.
func DeployAttestationLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AttestationLib, error) {
	parsed, err := AttestationLibMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AttestationLib{AttestationLibCaller: AttestationLibCaller{contract: contract}, AttestationLibTransactor: AttestationLibTransactor{contract: contract}, AttestationLibFilterer: AttestationLibFilterer{contract: contract}}, nil
}

// AttestationLib is an auto generated Go binding around an Ethereum contract.
type AttestationLib struct {
	AttestationLibCaller     // Read-only binding to the contract
	AttestationLibTransactor // Write-only binding to the contract
	AttestationLibFilterer   // Log filterer for contract events
}

// AttestationLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationLibSession struct {
	Contract     *AttestationLib   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestationLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationLibCallerSession struct {
	Contract *AttestationLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// AttestationLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationLibTransactorSession struct {
	Contract     *AttestationLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AttestationLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationLibRaw struct {
	Contract *AttestationLib // Generic contract binding to access the raw methods on
}

// AttestationLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationLibCallerRaw struct {
	Contract *AttestationLibCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationLibTransactorRaw struct {
	Contract *AttestationLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestationLib creates a new instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLib(address common.Address, backend bind.ContractBackend) (*AttestationLib, error) {
	contract, err := bindAttestationLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AttestationLib{AttestationLibCaller: AttestationLibCaller{contract: contract}, AttestationLibTransactor: AttestationLibTransactor{contract: contract}, AttestationLibFilterer: AttestationLibFilterer{contract: contract}}, nil
}

// NewAttestationLibCaller creates a new read-only instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibCaller(address common.Address, caller bind.ContractCaller) (*AttestationLibCaller, error) {
	contract, err := bindAttestationLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationLibCaller{contract: contract}, nil
}

// NewAttestationLibTransactor creates a new write-only instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationLibTransactor, error) {
	contract, err := bindAttestationLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationLibTransactor{contract: contract}, nil
}

// NewAttestationLibFilterer creates a new log filterer instance of AttestationLib, bound to a specific deployed contract.
func NewAttestationLibFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationLibFilterer, error) {
	contract, err := bindAttestationLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationLibFilterer{contract: contract}, nil
}

// bindAttestationLib binds a generic wrapper to an already deployed contract.
func bindAttestationLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationLib *AttestationLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationLib.Contract.AttestationLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationLib *AttestationLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationLib.Contract.AttestationLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationLib *AttestationLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationLib.Contract.AttestationLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AttestationLib *AttestationLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AttestationLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AttestationLib *AttestationLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AttestationLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AttestationLib *AttestationLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AttestationLib.Contract.contract.Transact(opts, method, params...)
}

// MemViewLibMetaData contains all meta data concerning the MemViewLib contract.
var MemViewLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122032e1deb551f0f843bc8574a7310f40a5ed56f751272fc35e6a7e713c8e125e0864736f6c63430008110033",
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
