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

// ExecutionAttestation is an auto generated low-level Go binding around an user-defined struct.
type ExecutionAttestation struct {
	Notary      common.Address
	Nonce       uint32
	SubmittedAt *big.Int
}

// SummitAttestation is an auto generated low-level Go binding around an user-defined struct.
type SummitAttestation struct {
	SnapRoot    [32]byte
	AgentRoot   [32]byte
	BlockNumber *big.Int
	Timestamp   *big.Int
}

// AttestationHarnessMetaData contains all meta data concerning the AttestationHarness contract.
var AttestationHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"agentRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"blockNumber\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"castToAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"snapRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot_\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber_\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp_\",\"type\":\"uint40\"}],\"name\":\"formatAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"snapRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"agentRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint40\",\"name\":\"blockNumber\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"timestamp\",\"type\":\"uint40\"}],\"internalType\":\"structSummitAttestation\",\"name\":\"summitAtt\",\"type\":\"tuple\"},{\"internalType\":\"uint32\",\"name\":\"nonce_\",\"type\":\"uint32\"}],\"name\":\"formatSummitAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"isAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submittedAt\",\"type\":\"uint40\"}],\"internalType\":\"structExecutionAttestation\",\"name\":\"execAtt\",\"type\":\"tuple\"}],\"name\":\"isEmpty\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"snapRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"name\":\"timestamp\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"}],\"name\":\"toExecutionAttestation\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"notary\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"uint40\",\"name\":\"submittedAt\",\"type\":\"uint40\"}],\"internalType\":\"structExecutionAttestation\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3c4d0a2d": "agentRoot(bytes)",
		"e948e600": "blockNumber(bytes)",
		"11f6389c": "castToAttestation(bytes)",
		"86bf2253": "formatAttestation(bytes32,bytes32,uint32,uint40,uint40)",
		"ca3634c2": "formatSummitAttestation((bytes32,bytes32,uint40,uint40),uint32)",
		"aa1e84de": "hash(bytes)",
		"3ae7034d": "isAttestation(bytes)",
		"26721a76": "isEmpty((address,uint32,uint40))",
		"4e765004": "nonce(bytes)",
		"91bacf60": "snapRoot(bytes)",
		"1c9aa222": "timestamp(bytes)",
		"d9569d9c": "toExecutionAttestation(bytes,address)",
	},
	Bin: "0x608060405234801561001057600080fd5b506113ce806100206000396000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c806386bf225311610081578063ca3634c21161005b578063ca3634c2146101e3578063d9569d9c146101f6578063e948e6001461024e57600080fd5b806386bf2253146101aa57806391bacf60146101bd578063aa1e84de146101d057600080fd5b80633ae7034d116100b25780633ae7034d1461014e5780633c4d0a2d146101615780634e7650041461018257600080fd5b806311f6389c146100d95780631c9aa2221461010257806326721a761461012b575b600080fd5b6100ec6100e7366004610f0e565b610261565b6040516100f99190610faf565b60405180910390f35b610115610110366004610f0e565b610286565b60405164ffffffffff90911681526020016100f9565b61013e610139366004611014565b6102a5565b60405190151581526020016100f9565b61013e61015c366004610f0e565b6102c6565b61017461016f366004610f0e565b6102f2565b6040519081526020016100f9565b610195610190366004610f0e565b61030b565b60405163ffffffff90911681526020016100f9565b6100ec6101b8366004611080565b610324565b6101746101cb366004610f0e565b6103b8565b6101746101de366004610f0e565b6103d1565b6100ec6101f13660046110d7565b6103ea565b610209610204366004611169565b6103f6565b60408051825173ffffffffffffffffffffffffffffffffffffffff16815260208084015163ffffffff16908201529181015164ffffffffff16908201526060016100f9565b61011561025c366004610f0e565b61042c565b6060600061026e83610445565b905061027f62ffffff198216610458565b9392505050565b600061029f61029483610445565b62ffffff19166104ab565b92915050565b600061029f825173ffffffffffffffffffffffffffffffffffffffff161590565b600061029f6102d583836104c1565b62ffffff191660181c6bffffffffffffffffffffffff16604e1490565b600061029f61030083610445565b62ffffff19166104e5565b600061029f61031983610445565b62ffffff19166104fa565b60408051602081018790528082018690527fffffffff0000000000000000000000000000000000000000000000000000000060e086901b166060828101919091527fffffffffff00000000000000000000000000000000000000000000000000000060d886811b8216606485015285901b1660698301528251808303604e018152606e9092019092525b9695505050505050565b600061029f6103c683610445565b62ffffff1916610510565b600061029f6103df83610445565b62ffffff1916610525565b606061027f8383610588565b604080516060810182526000808252602082018190529181019190915261027f8261042085610445565b62ffffff191690610626565b600061029f61043a83610445565b62ffffff1916610686565b600061029f6104538361069c565b6106a8565b60606000806104758460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061049a8483602001610734565b508181016020016040529052919050565b600062ffffff19821661027f816049600561091b565b8151600090602084016104dc64ffffffffff8516828461094b565b95945050505050565b600062ffffff19821661027f81602080610992565b600062ffffff19821661027f816040600461091b565b600062ffffff19821661027f81836020610992565b600062ffffff1982167f569efb4f951664b562fe9283d8f1a49928bec7335bab838210b64c85e11be59e61055882610b40565b60408051602081019390935282015260600160405160208183030381529060405280519060200120915050919050565b8151602080840151604080860151606087810151835195860196909652848301939093527fffffffff0000000000000000000000000000000000000000000000000000000060e087901b16848401527fffffffffff00000000000000000000000000000000000000000000000000000060d891821b8116606486015294901b90931660698301528251808303604e018152606e90920190925261027f565b604080516060810182526000602082018190529181019190915273ffffffffffffffffffffffffffffffffffffffff8216815261066862ffffff1984166104fa565b63ffffffff16602082015264ffffffffff4216604082015292915050565b600062ffffff19821661027f816044600561091b565b600061029f82826104c1565b60006106c58260181c6bffffffffffffffffffffffff16604e1490565b610730576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601260248201527f4e6f7420616e206174746573746174696f6e000000000000000000000000000060448201526064015b60405180910390fd5b5090565b600062ffffff19808416036107a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610727565b6107ae83610b8e565b610814576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610727565b600061082e8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061084985610bca565b6bffffffffffffffffffffffff16905060008060405191508582111561086f5760206060fd5b8386858560045afa9050806108e0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610727565b6109106108ec88610bf1565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b60006109288260206111dd565b6109339060086111f6565b60ff16610941858585610992565b901c949350505050565b6000806109588385611219565b9050604051811115610968575060005b8060000361097d5762ffffff1991505061027f565b5050606092831b9190911790911b1760181b90565b60008160ff166000036109a75750600061027f565b6109bf8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166109da60ff841685611219565b1115610a5d57610a2a6109ec85610bca565b6bffffffffffffffffffffffff16610a128660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16610c15565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107279190610faf565b60208260ff161115610acb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610727565b600882026000610ada86610bca565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b600080610b4c83610bca565b6bffffffffffffffffffffffff1690506000610b768460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169091209392505050565b6000610b9982610bf1565b64ffffffffff1664ffffffffff03610bb357506000919050565b6000610bbe83610c83565b60405110199392505050565b600080610bd960606018611219565b9290921c6bffffffffffffffffffffffff1692915050565b6000806060610c01816018611219565b610c0b9190611219565b9290921c92915050565b60606000610c2286610cbc565b9150506000610c3086610cbc565b9150506000610c3e86610cbc565b9150506000610c4c86610cbc565b91505083838383604051602001610c66949392919061122c565b604051602081830303815290604052945050505050949350505050565b6000610c9d8260181c6bffffffffffffffffffffffff1690565b610ca683610bca565b016bffffffffffffffffffffffff169050919050565b600080601f5b600f8160ff161115610d2f576000610cdb8260086111f6565b60ff1685901c9050610cec81610da6565b61ffff16841793508160ff16601014610d0757601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610cc2565b50600f5b60ff8160ff161015610da0576000610d4c8260086111f6565b60ff1685901c9050610d5d81610da6565b61ffff16831792508160ff16600014610d7857601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610d33565b50915091565b6000610db860048360ff16901c610dd8565b60ff1661ffff919091161760081b610dcf82610dd8565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f84169182908110610e2557610e25611369565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f830112610e7457600080fd5b813567ffffffffffffffff80821115610e8f57610e8f610e34565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610ed557610ed5610e34565b81604052838152866020858801011115610eee57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215610f2057600080fd5b813567ffffffffffffffff811115610f3757600080fd5b610f4384828501610e63565b949350505050565b6000815180845260005b81811015610f7157602081850181015186830182015201610f55565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061027f6020830184610f4b565b803573ffffffffffffffffffffffffffffffffffffffff81168114610fe657600080fd5b919050565b803563ffffffff81168114610fe657600080fd5b803564ffffffffff81168114610fe657600080fd5b60006060828403121561102657600080fd5b6040516060810181811067ffffffffffffffff8211171561104957611049610e34565b60405261105583610fc2565b815261106360208401610feb565b602082015261107460408401610fff565b60408201529392505050565b600080600080600060a0868803121561109857600080fd5b85359450602086013593506110af60408701610feb565b92506110bd60608701610fff565b91506110cb60808701610fff565b90509295509295909350565b60008082840360a08112156110eb57600080fd5b60808112156110f957600080fd5b506040516080810181811067ffffffffffffffff8211171561111d5761111d610e34565b8060405250833581526020840135602082015261113c60408501610fff565b604082015261114d60608501610fff565b6060820152915061116060808401610feb565b90509250929050565b6000806040838503121561117c57600080fd5b823567ffffffffffffffff81111561119357600080fd5b61119f85828601610e63565b92505061116060208401610fc2565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60ff828116828216039081111561029f5761029f6111ae565b60ff8181168382160290811690818114611212576112126111ae565b5092915050565b8082018082111561029f5761029f6111ae565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d82016103ae565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220f452a816619f401a6a91774b3004ebf9f398601993643c8e6dc8d55df212f97764736f6c63430008110033",
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

// AgentRoot is a free data retrieval call binding the contract method 0x3c4d0a2d.
//
// Solidity: function agentRoot(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) AgentRoot(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "agentRoot", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AgentRoot is a free data retrieval call binding the contract method 0x3c4d0a2d.
//
// Solidity: function agentRoot(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) AgentRoot(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.AgentRoot(&_AttestationHarness.CallOpts, payload)
}

// AgentRoot is a free data retrieval call binding the contract method 0x3c4d0a2d.
//
// Solidity: function agentRoot(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) AgentRoot(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.AgentRoot(&_AttestationHarness.CallOpts, payload)
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

// FormatAttestation is a free data retrieval call binding the contract method 0x86bf2253.
//
// Solidity: function formatAttestation(bytes32 snapRoot_, bytes32 agentRoot_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatAttestation(opts *bind.CallOpts, snapRoot_ [32]byte, agentRoot_ [32]byte, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatAttestation", snapRoot_, agentRoot_, nonce_, blockNumber_, timestamp_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatAttestation is a free data retrieval call binding the contract method 0x86bf2253.
//
// Solidity: function formatAttestation(bytes32 snapRoot_, bytes32 agentRoot_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatAttestation(snapRoot_ [32]byte, agentRoot_ [32]byte, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, snapRoot_, agentRoot_, nonce_, blockNumber_, timestamp_)
}

// FormatAttestation is a free data retrieval call binding the contract method 0x86bf2253.
//
// Solidity: function formatAttestation(bytes32 snapRoot_, bytes32 agentRoot_, uint32 nonce_, uint40 blockNumber_, uint40 timestamp_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatAttestation(snapRoot_ [32]byte, agentRoot_ [32]byte, nonce_ uint32, blockNumber_ *big.Int, timestamp_ *big.Int) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, snapRoot_, agentRoot_, nonce_, blockNumber_, timestamp_)
}

// FormatSummitAttestation is a free data retrieval call binding the contract method 0xca3634c2.
//
// Solidity: function formatSummitAttestation((bytes32,bytes32,uint40,uint40) summitAtt, uint32 nonce_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatSummitAttestation(opts *bind.CallOpts, summitAtt SummitAttestation, nonce_ uint32) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatSummitAttestation", summitAtt, nonce_)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatSummitAttestation is a free data retrieval call binding the contract method 0xca3634c2.
//
// Solidity: function formatSummitAttestation((bytes32,bytes32,uint40,uint40) summitAtt, uint32 nonce_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatSummitAttestation(summitAtt SummitAttestation, nonce_ uint32) ([]byte, error) {
	return _AttestationHarness.Contract.FormatSummitAttestation(&_AttestationHarness.CallOpts, summitAtt, nonce_)
}

// FormatSummitAttestation is a free data retrieval call binding the contract method 0xca3634c2.
//
// Solidity: function formatSummitAttestation((bytes32,bytes32,uint40,uint40) summitAtt, uint32 nonce_) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatSummitAttestation(summitAtt SummitAttestation, nonce_ uint32) ([]byte, error) {
	return _AttestationHarness.Contract.FormatSummitAttestation(&_AttestationHarness.CallOpts, summitAtt, nonce_)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) Hash(opts *bind.CallOpts, payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "hash", payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) Hash(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.Hash(&_AttestationHarness.CallOpts, payload)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) Hash(payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.Hash(&_AttestationHarness.CallOpts, payload)
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

// IsEmpty is a free data retrieval call binding the contract method 0x26721a76.
//
// Solidity: function isEmpty((address,uint32,uint40) execAtt) pure returns(bool)
func (_AttestationHarness *AttestationHarnessCaller) IsEmpty(opts *bind.CallOpts, execAtt ExecutionAttestation) (bool, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "isEmpty", execAtt)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEmpty is a free data retrieval call binding the contract method 0x26721a76.
//
// Solidity: function isEmpty((address,uint32,uint40) execAtt) pure returns(bool)
func (_AttestationHarness *AttestationHarnessSession) IsEmpty(execAtt ExecutionAttestation) (bool, error) {
	return _AttestationHarness.Contract.IsEmpty(&_AttestationHarness.CallOpts, execAtt)
}

// IsEmpty is a free data retrieval call binding the contract method 0x26721a76.
//
// Solidity: function isEmpty((address,uint32,uint40) execAtt) pure returns(bool)
func (_AttestationHarness *AttestationHarnessCallerSession) IsEmpty(execAtt ExecutionAttestation) (bool, error) {
	return _AttestationHarness.Contract.IsEmpty(&_AttestationHarness.CallOpts, execAtt)
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

// ToExecutionAttestation is a free data retrieval call binding the contract method 0xd9569d9c.
//
// Solidity: function toExecutionAttestation(bytes payload, address notary) view returns((address,uint32,uint40))
func (_AttestationHarness *AttestationHarnessCaller) ToExecutionAttestation(opts *bind.CallOpts, payload []byte, notary common.Address) (ExecutionAttestation, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "toExecutionAttestation", payload, notary)

	if err != nil {
		return *new(ExecutionAttestation), err
	}

	out0 := *abi.ConvertType(out[0], new(ExecutionAttestation)).(*ExecutionAttestation)

	return out0, err

}

// ToExecutionAttestation is a free data retrieval call binding the contract method 0xd9569d9c.
//
// Solidity: function toExecutionAttestation(bytes payload, address notary) view returns((address,uint32,uint40))
func (_AttestationHarness *AttestationHarnessSession) ToExecutionAttestation(payload []byte, notary common.Address) (ExecutionAttestation, error) {
	return _AttestationHarness.Contract.ToExecutionAttestation(&_AttestationHarness.CallOpts, payload, notary)
}

// ToExecutionAttestation is a free data retrieval call binding the contract method 0xd9569d9c.
//
// Solidity: function toExecutionAttestation(bytes payload, address notary) view returns((address,uint32,uint40))
func (_AttestationHarness *AttestationHarnessCallerSession) ToExecutionAttestation(payload []byte, notary common.Address) (ExecutionAttestation, error) {
	return _AttestationHarness.Contract.ToExecutionAttestation(&_AttestationHarness.CallOpts, payload, notary)
}

// AttestationLibMetaData contains all meta data concerning the AttestationLib contract.
var AttestationLibMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205f04fb9b3bbbe648585cf8ce8eb33ce85bae55cd49dee7b804784c431511c1d364736f6c63430008110033",
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

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207770bb92e6032144968d0aac3c1fcc4c647d6602fc88a5181a88b5d29b57f3c464736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea26469706673582212206216ef17d8a7f443bcfbf6003d946c804cb34daa8db803202f3870075ef303b964736f6c63430008110033",
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
