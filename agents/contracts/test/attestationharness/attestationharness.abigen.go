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

// AttestationMetaData contains all meta data concerning the Attestation contract.
var AttestationMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122055925333591b582ca8e63cfc04ec0ff2e4124b0a9e60439cf73e45a944e5bee764736f6c63430008110033",
}

// AttestationABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationMetaData.ABI instead.
var AttestationABI = AttestationMetaData.ABI

// AttestationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttestationMetaData.Bin instead.
var AttestationBin = AttestationMetaData.Bin

// DeployAttestation deploys a new Ethereum contract, binding an instance of Attestation to it.
func DeployAttestation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Attestation, error) {
	parsed, err := AttestationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttestationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attestation{AttestationCaller: AttestationCaller{contract: contract}, AttestationTransactor: AttestationTransactor{contract: contract}, AttestationFilterer: AttestationFilterer{contract: contract}}, nil
}

// Attestation is an auto generated Go binding around an Ethereum contract.
type Attestation struct {
	AttestationCaller     // Read-only binding to the contract
	AttestationTransactor // Write-only binding to the contract
	AttestationFilterer   // Log filterer for contract events
}

// AttestationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationSession struct {
	Contract     *Attestation      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationCallerSession struct {
	Contract *AttestationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AttestationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationTransactorSession struct {
	Contract     *AttestationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AttestationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationRaw struct {
	Contract *Attestation // Generic contract binding to access the raw methods on
}

// AttestationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationCallerRaw struct {
	Contract *AttestationCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationTransactorRaw struct {
	Contract *AttestationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestation creates a new instance of Attestation, bound to a specific deployed contract.
func NewAttestation(address common.Address, backend bind.ContractBackend) (*Attestation, error) {
	contract, err := bindAttestation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attestation{AttestationCaller: AttestationCaller{contract: contract}, AttestationTransactor: AttestationTransactor{contract: contract}, AttestationFilterer: AttestationFilterer{contract: contract}}, nil
}

// NewAttestationCaller creates a new read-only instance of Attestation, bound to a specific deployed contract.
func NewAttestationCaller(address common.Address, caller bind.ContractCaller) (*AttestationCaller, error) {
	contract, err := bindAttestation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationCaller{contract: contract}, nil
}

// NewAttestationTransactor creates a new write-only instance of Attestation, bound to a specific deployed contract.
func NewAttestationTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationTransactor, error) {
	contract, err := bindAttestation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationTransactor{contract: contract}, nil
}

// NewAttestationFilterer creates a new log filterer instance of Attestation, bound to a specific deployed contract.
func NewAttestationFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationFilterer, error) {
	contract, err := bindAttestation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationFilterer{contract: contract}, nil
}

// bindAttestation binds a generic wrapper to an already deployed contract.
func bindAttestation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestation *AttestationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestation.Contract.AttestationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestation *AttestationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestation.Contract.AttestationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestation *AttestationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestation.Contract.AttestationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestation *AttestationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestation *AttestationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestation *AttestationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestation.Contract.contract.Transact(opts, method, params...)
}

// AttestationHarnessMetaData contains all meta data concerning the AttestationHarness contract.
var AttestationHarnessMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestationData\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"}],\"name\":\"attestationDomains\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"attestationKey\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedDestination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedDomains\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedKey\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedOrigin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"castToAttestation\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"formatAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"formatAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"formatAttestationData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"isAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"notarySignature\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetDestination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9ffb971e": "attestationData(uint40,bytes)",
		"a104a5e5": "attestationDataLength()",
		"65dfb428": "attestationDomains(uint32,uint32)",
		"9f668e20": "attestationKey(uint32,uint32,uint32)",
		"e152990b": "attestationLength()",
		"3b830f3b": "attestedDestination(uint40,bytes)",
		"c231bd8f": "attestedDomains(uint40,bytes)",
		"79ce92a9": "attestedKey(uint40,bytes)",
		"8b445f51": "attestedNonce(uint40,bytes)",
		"badad7db": "attestedOrigin(uint40,bytes)",
		"91eedc1d": "attestedRoot(uint40,bytes)",
		"c2e19ed2": "castToAttestation(uint40,bytes)",
		"81646674": "formatAttestation(bytes,bytes)",
		"01a74313": "formatAttestation(uint32,uint32,uint32,bytes32,bytes)",
		"2951eae3": "formatAttestationData(uint32,uint32,uint32,bytes32)",
		"3ae7034d": "isAttestation(bytes)",
		"a43aa286": "notarySignature(uint40,bytes)",
		"d2c4428a": "offsetDestination()",
		"569e1eaf": "offsetNonce()",
		"320bfc44": "offsetOrigin()",
		"5b42242d": "offsetRoot()",
		"740261a6": "offsetSignature()",
	},
	Bin: "0x608060405234801561001057600080fd5b5061168b806100206000396000f3fe608060405234801561001057600080fd5b50600436106101825760003560e01c80638b445f51116100d8578063a43aa2861161008c578063c2e19ed211610066578063c2e19ed214610366578063d2c4428a14610379578063e152990b1461038057600080fd5b8063a43aa2861461032d578063badad7db14610340578063c231bd8f1461035357600080fd5b80639f668e20116100bd5780639f668e20146102c95780639ffb971e1461030c578063a104a5e51461025957600080fd5b80638b445f51146102a357806391eedc1d146102b657600080fd5b8063569e1eaf1161013a578063740261a611610114578063740261a61461025957806379ce92a914610260578063816466741461029057600080fd5b8063569e1eaf1461021f5780635b42242d1461022657806365dfb4281461022d57600080fd5b8063320bfc441161016b578063320bfc44146101c35780633ae7034d146101d45780633b830f3b146101f757600080fd5b806301a74313146101875780632951eae3146101b0575b600080fd5b61019a6101953660046111a7565b610388565b6040516101a7919061128e565b60405180910390f35b61019a6101be3660046112a1565b6103a9565b60005b6040519081526020016101a7565b6101e76101e23660046112ec565b610410565b60405190151581526020016101a7565b61020a610205366004611321565b61042f565b60405163ffffffff90911681526020016101a7565b60086101c6565b600c6101c6565b61024061023b36600461137b565b610450565b60405167ffffffffffffffff90911681526020016101a7565b602c6101c6565b61027361026e366004611321565b61046e565b6040516bffffffffffffffffffffffff90911681526020016101a7565b61019a61029e3660046113ae565b610488565b61020a6102b1366004611321565b610494565b6101c66102c4366004611321565b6104ae565b6102736102d7366004611408565b600063ffffffff8216602084901b67ffffffff0000000016604086901b6bffffffff0000000000000000161717949350505050565b61031f61031a366004611321565b6104c8565b6040516101a792919061144b565b61031f61033b366004611321565b610507565b61020a61034e366004611321565b610524565b610240610361366004611321565b61053e565b61031f610374366004611321565b610558565b60046101c6565b6101c6610567565b606061039f610399878787876103a9565b83610488565b9695505050505050565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e087811b8216602084015286811b8216602484015285901b166028820152602c80820184905282518083039091018152604c9091019091525b949350505050565b600061042961041e8361057a565b62ffffff191661058b565b92915050565b600061044961043e83856105c1565b62ffffff19166105e5565b9392505050565b600067ffffffff00000000602084901b1663ffffffff831617610449565b600061044961047d83856105c1565b62ffffff1916610619565b60606104498383610645565b60006104496104a383856105c1565b62ffffff1916610671565b60006104496104bd83856105c1565b62ffffff191661069d565b60006060816104e56104da85876105c1565b62ffffff19166106c9565b905060d881901c6104fb62ffffff1983166106fb565b92509250509250929050565b60006060816104e561051985876105c1565b62ffffff191661074e565b600061044961053383856105c1565b62ffffff191661077f565b600061044961054d83856105c1565b62ffffff19166107ab565b6000606060006104e58461057a565b6000610575602c604161149a565b905090565b6000610429826401010000006105c1565b6000610599602c604161149a565b6bffffffffffffffffffffffff601884901c166bffffffffffffffffffffffff161492915050565b8151600090602084016105dc64ffffffffff851682846107d7565b95945050505050565b6000816105fd62ffffff19821664010100000061081e565b5061061062ffffff198416600480610942565b91505b50919050565b60008161063162ffffff19821664010100000061081e565b5061061062ffffff1984166000600c610942565b6060828260405160200161065a9291906114ad565b604051602081830303815290604052905092915050565b60008161068962ffffff19821664010100000061081e565b5061061062ffffff19841660086004610942565b6000816106b562ffffff19821664010100000061081e565b5061061062ffffff198416600c6020610972565b6000816106e162ffffff19821664010100000061081e565b5061061062ffffff1984166000602c640101010000610b3e565b60606000806107188460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050604051915081925061073d8483602001610bb8565b508181016020016040529052919050565b60008161076662ffffff19821664010100000061081e565b5061061062ffffff198416602c60416301000000610b3e565b60008161079762ffffff19821664010100000061081e565b5061061062ffffff19841660006004610942565b6000816107c362ffffff19821664010100000061081e565b5061061062ffffff19841660006008610942565b6000806107e4838561149a565b90506040518111156107f4575060005b806000036108095762ffffff19915050610449565b5050606092831b9190911790911b1760181b90565b600061082a8383610d3b565b61093b57600061084961083d8560d81c90565b64ffffffffff16610d5e565b915050600061085e8464ffffffffff16610d5e565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610932919061128e565b60405180910390fd5b5090919050565b600061094f8260206114dc565b61095a9060086114f5565b60ff16610968858585610972565b901c949350505050565b60008160ff1660000361098757506000610449565b61099f8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff166109ba60ff84168561149a565b1115610a4c57610a196109db8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16610a018660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16610e48565b6040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610932919061128e565b60208260ff161115610aba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610932565b600882026000610ad88660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b600080610b598660781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050610b7286610eb6565b84610b7d878461149a565b610b87919061149a565b1115610b9a5762ffffff19915050610408565b610ba4858261149a565b905061039f8364ffffffffff1682866107d7565b600062ffffff1980841603610c29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610932565b610c3283610efe565b610c98576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610932565b6000610cb28460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610cdc8560781c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000604051905084811115610d015760206060fd5b8285848460045afa5061039f610d178760d81c90565b70ffffffffff000000000000000000000000606091821b168717901b841760181b90565b60008164ffffffffff16610d4f8460d81c90565b64ffffffffff16149392505050565b600080601f5b600f8160ff161115610dd1576000610d7d8260086114f5565b60ff1685901c9050610d8e81610f3b565b61ffff16841793508160ff16601014610da957601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610d64565b50600f5b60ff8160ff161015610e42576000610dee8260086114f5565b60ff1685901c9050610dff81610f3b565b61ffff16831792508160ff16600014610e1a57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01610dd5565b50915091565b60606000610e5586610d5e565b9150506000610e6386610d5e565b9150506000610e7186610d5e565b9150506000610e7f86610d5e565b91505083838383604051602001610e999493929190611518565b604051602081830303815290604052945050505050949350505050565b6000610ed08260181c6bffffffffffffffffffffffff1690565b610ee88360781c6bffffffffffffffffffffffff1690565b016bffffffffffffffffffffffff169050919050565b6000610f0a8260d81c90565b64ffffffffff1664ffffffffff03610f2457506000919050565b6000610f2f83610eb6565b60405110199392505050565b6000610f4d60048360ff16901c610f6d565b60ff1661ffff919091161760081b610f6482610f6d565b60ff1617919050565b600060f08083179060ff82169003610f885750603092915050565b8060ff1660f103610f9c5750603192915050565b8060ff1660f203610fb05750603292915050565b8060ff1660f303610fc45750603392915050565b8060ff1660f403610fd85750603492915050565b8060ff1660f503610fec5750603592915050565b8060ff1660f6036110005750603692915050565b8060ff1660f7036110145750603792915050565b8060ff1660f8036110285750603892915050565b8060ff1660f90361103c5750603992915050565b8060ff1660fa036110505750606192915050565b8060ff1660fb036110645750606292915050565b8060ff1660fc036110785750606392915050565b8060ff1660fd0361108c5750606492915050565b8060ff1660fe036110a05750606592915050565b8060ff1660ff036106135750606692915050565b803563ffffffff811681146110c857600080fd5b919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f83011261110d57600080fd5b813567ffffffffffffffff80821115611128576111286110cd565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561116e5761116e6110cd565b8160405283815286602085880101111561118757600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600080600060a086880312156111bf57600080fd5b6111c8866110b4565b94506111d6602087016110b4565b93506111e4604087016110b4565b925060608601359150608086013567ffffffffffffffff81111561120757600080fd5b611213888289016110fc565b9150509295509295909350565b60005b8381101561123b578181015183820152602001611223565b50506000910152565b6000815180845261125c816020860160208601611220565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006104496020830184611244565b600080600080608085870312156112b757600080fd5b6112c0856110b4565b93506112ce602086016110b4565b92506112dc604086016110b4565b9396929550929360600135925050565b6000602082840312156112fe57600080fd5b813567ffffffffffffffff81111561131557600080fd5b610408848285016110fc565b6000806040838503121561133457600080fd5b823564ffffffffff8116811461134957600080fd5b9150602083013567ffffffffffffffff81111561136557600080fd5b611371858286016110fc565b9150509250929050565b6000806040838503121561138e57600080fd5b611397836110b4565b91506113a5602084016110b4565b90509250929050565b600080604083850312156113c157600080fd5b823567ffffffffffffffff808211156113d957600080fd5b6113e5868387016110fc565b935060208501359150808211156113fb57600080fd5b50611371858286016110fc565b60008060006060848603121561141d57600080fd5b611426846110b4565b9250611434602085016110b4565b9150611442604085016110b4565b90509250925092565b64ffffffffff831681526040602082015260006104086040830184611244565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b808201808211156104295761042961146b565b600083516114bf818460208801611220565b8351908301906114d3818360208801611220565b01949350505050565b60ff82811682821603908111156104295761042961146b565b60ff81811683821602908116908181146115115761151161146b565b5092915050565b7f54797065644d656d566965772f696e646578202d204f76657272616e2074686581527f20766965772e20536c696365206973206174203078000000000000000000000060208201527fffffffffffff000000000000000000000000000000000000000000000000000060d086811b821660358401527f2077697468206c656e6774682030780000000000000000000000000000000000603b840181905286821b8316604a8501527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060508501527f7800000000000000000000000000000000000000000000000000000000000000607085015285821b83166071850152607784015283901b1660868201527f2e00000000000000000000000000000000000000000000000000000000000000608c8201526000608d820161039f56fea26469706673582212205c438101af78835ccf3351be0e1488a13c6f4dc3ff3df68c8a462558c5c4263964736f6c63430008110033",
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

// AttestationData is a free data retrieval call binding the contract method 0x9ffb971e.
//
// Solidity: function attestationData(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCaller) AttestationData(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestationData", _type, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// AttestationData is a free data retrieval call binding the contract method 0x9ffb971e.
//
// Solidity: function attestationData(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessSession) AttestationData(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.AttestationData(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestationData is a free data retrieval call binding the contract method 0x9ffb971e.
//
// Solidity: function attestationData(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestationData(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.AttestationData(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestationDataLength is a free data retrieval call binding the contract method 0xa104a5e5.
//
// Solidity: function attestationDataLength() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) AttestationDataLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestationDataLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationDataLength is a free data retrieval call binding the contract method 0xa104a5e5.
//
// Solidity: function attestationDataLength() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) AttestationDataLength() (*big.Int, error) {
	return _AttestationHarness.Contract.AttestationDataLength(&_AttestationHarness.CallOpts)
}

// AttestationDataLength is a free data retrieval call binding the contract method 0xa104a5e5.
//
// Solidity: function attestationDataLength() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestationDataLength() (*big.Int, error) {
	return _AttestationHarness.Contract.AttestationDataLength(&_AttestationHarness.CallOpts)
}

// AttestationDomains is a free data retrieval call binding the contract method 0x65dfb428.
//
// Solidity: function attestationDomains(uint32 _origin, uint32 _destination) pure returns(uint64)
func (_AttestationHarness *AttestationHarnessCaller) AttestationDomains(opts *bind.CallOpts, _origin uint32, _destination uint32) (uint64, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestationDomains", _origin, _destination)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AttestationDomains is a free data retrieval call binding the contract method 0x65dfb428.
//
// Solidity: function attestationDomains(uint32 _origin, uint32 _destination) pure returns(uint64)
func (_AttestationHarness *AttestationHarnessSession) AttestationDomains(_origin uint32, _destination uint32) (uint64, error) {
	return _AttestationHarness.Contract.AttestationDomains(&_AttestationHarness.CallOpts, _origin, _destination)
}

// AttestationDomains is a free data retrieval call binding the contract method 0x65dfb428.
//
// Solidity: function attestationDomains(uint32 _origin, uint32 _destination) pure returns(uint64)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestationDomains(_origin uint32, _destination uint32) (uint64, error) {
	return _AttestationHarness.Contract.AttestationDomains(&_AttestationHarness.CallOpts, _origin, _destination)
}

// AttestationKey is a free data retrieval call binding the contract method 0x9f668e20.
//
// Solidity: function attestationKey(uint32 _origin, uint32 _destination, uint32 _nonce) pure returns(uint96)
func (_AttestationHarness *AttestationHarnessCaller) AttestationKey(opts *bind.CallOpts, _origin uint32, _destination uint32, _nonce uint32) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestationKey", _origin, _destination, _nonce)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationKey is a free data retrieval call binding the contract method 0x9f668e20.
//
// Solidity: function attestationKey(uint32 _origin, uint32 _destination, uint32 _nonce) pure returns(uint96)
func (_AttestationHarness *AttestationHarnessSession) AttestationKey(_origin uint32, _destination uint32, _nonce uint32) (*big.Int, error) {
	return _AttestationHarness.Contract.AttestationKey(&_AttestationHarness.CallOpts, _origin, _destination, _nonce)
}

// AttestationKey is a free data retrieval call binding the contract method 0x9f668e20.
//
// Solidity: function attestationKey(uint32 _origin, uint32 _destination, uint32 _nonce) pure returns(uint96)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestationKey(_origin uint32, _destination uint32, _nonce uint32) (*big.Int, error) {
	return _AttestationHarness.Contract.AttestationKey(&_AttestationHarness.CallOpts, _origin, _destination, _nonce)
}

// AttestationLength is a free data retrieval call binding the contract method 0xe152990b.
//
// Solidity: function attestationLength() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) AttestationLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestationLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationLength is a free data retrieval call binding the contract method 0xe152990b.
//
// Solidity: function attestationLength() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) AttestationLength() (*big.Int, error) {
	return _AttestationHarness.Contract.AttestationLength(&_AttestationHarness.CallOpts)
}

// AttestationLength is a free data retrieval call binding the contract method 0xe152990b.
//
// Solidity: function attestationLength() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestationLength() (*big.Int, error) {
	return _AttestationHarness.Contract.AttestationLength(&_AttestationHarness.CallOpts)
}

// AttestedDestination is a free data retrieval call binding the contract method 0x3b830f3b.
//
// Solidity: function attestedDestination(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCaller) AttestedDestination(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestedDestination", _type, _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AttestedDestination is a free data retrieval call binding the contract method 0x3b830f3b.
//
// Solidity: function attestedDestination(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessSession) AttestedDestination(_type *big.Int, _payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.AttestedDestination(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedDestination is a free data retrieval call binding the contract method 0x3b830f3b.
//
// Solidity: function attestedDestination(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestedDestination(_type *big.Int, _payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.AttestedDestination(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedDomains is a free data retrieval call binding the contract method 0xc231bd8f.
//
// Solidity: function attestedDomains(uint40 _type, bytes _payload) pure returns(uint64)
func (_AttestationHarness *AttestationHarnessCaller) AttestedDomains(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint64, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestedDomains", _type, _payload)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// AttestedDomains is a free data retrieval call binding the contract method 0xc231bd8f.
//
// Solidity: function attestedDomains(uint40 _type, bytes _payload) pure returns(uint64)
func (_AttestationHarness *AttestationHarnessSession) AttestedDomains(_type *big.Int, _payload []byte) (uint64, error) {
	return _AttestationHarness.Contract.AttestedDomains(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedDomains is a free data retrieval call binding the contract method 0xc231bd8f.
//
// Solidity: function attestedDomains(uint40 _type, bytes _payload) pure returns(uint64)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestedDomains(_type *big.Int, _payload []byte) (uint64, error) {
	return _AttestationHarness.Contract.AttestedDomains(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedKey is a free data retrieval call binding the contract method 0x79ce92a9.
//
// Solidity: function attestedKey(uint40 _type, bytes _payload) pure returns(uint96)
func (_AttestationHarness *AttestationHarnessCaller) AttestedKey(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestedKey", _type, _payload)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestedKey is a free data retrieval call binding the contract method 0x79ce92a9.
//
// Solidity: function attestedKey(uint40 _type, bytes _payload) pure returns(uint96)
func (_AttestationHarness *AttestationHarnessSession) AttestedKey(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _AttestationHarness.Contract.AttestedKey(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedKey is a free data retrieval call binding the contract method 0x79ce92a9.
//
// Solidity: function attestedKey(uint40 _type, bytes _payload) pure returns(uint96)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestedKey(_type *big.Int, _payload []byte) (*big.Int, error) {
	return _AttestationHarness.Contract.AttestedKey(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedNonce is a free data retrieval call binding the contract method 0x8b445f51.
//
// Solidity: function attestedNonce(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCaller) AttestedNonce(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestedNonce", _type, _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AttestedNonce is a free data retrieval call binding the contract method 0x8b445f51.
//
// Solidity: function attestedNonce(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessSession) AttestedNonce(_type *big.Int, _payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.AttestedNonce(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedNonce is a free data retrieval call binding the contract method 0x8b445f51.
//
// Solidity: function attestedNonce(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestedNonce(_type *big.Int, _payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.AttestedNonce(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedOrigin is a free data retrieval call binding the contract method 0xbadad7db.
//
// Solidity: function attestedOrigin(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCaller) AttestedOrigin(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint32, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestedOrigin", _type, _payload)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// AttestedOrigin is a free data retrieval call binding the contract method 0xbadad7db.
//
// Solidity: function attestedOrigin(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessSession) AttestedOrigin(_type *big.Int, _payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.AttestedOrigin(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedOrigin is a free data retrieval call binding the contract method 0xbadad7db.
//
// Solidity: function attestedOrigin(uint40 _type, bytes _payload) pure returns(uint32)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestedOrigin(_type *big.Int, _payload []byte) (uint32, error) {
	return _AttestationHarness.Contract.AttestedOrigin(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedRoot is a free data retrieval call binding the contract method 0x91eedc1d.
//
// Solidity: function attestedRoot(uint40 _type, bytes _payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCaller) AttestedRoot(opts *bind.CallOpts, _type *big.Int, _payload []byte) ([32]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "attestedRoot", _type, _payload)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AttestedRoot is a free data retrieval call binding the contract method 0x91eedc1d.
//
// Solidity: function attestedRoot(uint40 _type, bytes _payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessSession) AttestedRoot(_type *big.Int, _payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.AttestedRoot(&_AttestationHarness.CallOpts, _type, _payload)
}

// AttestedRoot is a free data retrieval call binding the contract method 0x91eedc1d.
//
// Solidity: function attestedRoot(uint40 _type, bytes _payload) pure returns(bytes32)
func (_AttestationHarness *AttestationHarnessCallerSession) AttestedRoot(_type *big.Int, _payload []byte) ([32]byte, error) {
	return _AttestationHarness.Contract.AttestedRoot(&_AttestationHarness.CallOpts, _type, _payload)
}

// CastToAttestation is a free data retrieval call binding the contract method 0xc2e19ed2.
//
// Solidity: function castToAttestation(uint40 , bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCaller) CastToAttestation(opts *bind.CallOpts, arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "castToAttestation", arg0, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// CastToAttestation is a free data retrieval call binding the contract method 0xc2e19ed2.
//
// Solidity: function castToAttestation(uint40 , bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessSession) CastToAttestation(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.CastToAttestation(&_AttestationHarness.CallOpts, arg0, _payload)
}

// CastToAttestation is a free data retrieval call binding the contract method 0xc2e19ed2.
//
// Solidity: function castToAttestation(uint40 , bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) CastToAttestation(arg0 *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.CastToAttestation(&_AttestationHarness.CallOpts, arg0, _payload)
}

// FormatAttestation is a free data retrieval call binding the contract method 0x01a74313.
//
// Solidity: function formatAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root, bytes _signature) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatAttestation(opts *bind.CallOpts, _origin uint32, _destination uint32, _nonce uint32, _root [32]byte, _signature []byte) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatAttestation", _origin, _destination, _nonce, _root, _signature)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatAttestation is a free data retrieval call binding the contract method 0x01a74313.
//
// Solidity: function formatAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root, bytes _signature) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatAttestation(_origin uint32, _destination uint32, _nonce uint32, _root [32]byte, _signature []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, _origin, _destination, _nonce, _root, _signature)
}

// FormatAttestation is a free data retrieval call binding the contract method 0x01a74313.
//
// Solidity: function formatAttestation(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root, bytes _signature) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatAttestation(_origin uint32, _destination uint32, _nonce uint32, _root [32]byte, _signature []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, _origin, _destination, _nonce, _root, _signature)
}

// FormatAttestation0 is a free data retrieval call binding the contract method 0x81646674.
//
// Solidity: function formatAttestation(bytes _data, bytes _signature) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatAttestation0(opts *bind.CallOpts, _data []byte, _signature []byte) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatAttestation0", _data, _signature)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatAttestation0 is a free data retrieval call binding the contract method 0x81646674.
//
// Solidity: function formatAttestation(bytes _data, bytes _signature) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatAttestation0(_data []byte, _signature []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation0(&_AttestationHarness.CallOpts, _data, _signature)
}

// FormatAttestation0 is a free data retrieval call binding the contract method 0x81646674.
//
// Solidity: function formatAttestation(bytes _data, bytes _signature) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatAttestation0(_data []byte, _signature []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation0(&_AttestationHarness.CallOpts, _data, _signature)
}

// FormatAttestationData is a free data retrieval call binding the contract method 0x2951eae3.
//
// Solidity: function formatAttestationData(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatAttestationData(opts *bind.CallOpts, _origin uint32, _destination uint32, _nonce uint32, _root [32]byte) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatAttestationData", _origin, _destination, _nonce, _root)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatAttestationData is a free data retrieval call binding the contract method 0x2951eae3.
//
// Solidity: function formatAttestationData(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatAttestationData(_origin uint32, _destination uint32, _nonce uint32, _root [32]byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestationData(&_AttestationHarness.CallOpts, _origin, _destination, _nonce, _root)
}

// FormatAttestationData is a free data retrieval call binding the contract method 0x2951eae3.
//
// Solidity: function formatAttestationData(uint32 _origin, uint32 _destination, uint32 _nonce, bytes32 _root) pure returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatAttestationData(_origin uint32, _destination uint32, _nonce uint32, _root [32]byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestationData(&_AttestationHarness.CallOpts, _origin, _destination, _nonce, _root)
}

// IsAttestation is a free data retrieval call binding the contract method 0x3ae7034d.
//
// Solidity: function isAttestation(bytes _payload) pure returns(bool)
func (_AttestationHarness *AttestationHarnessCaller) IsAttestation(opts *bind.CallOpts, _payload []byte) (bool, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "isAttestation", _payload)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAttestation is a free data retrieval call binding the contract method 0x3ae7034d.
//
// Solidity: function isAttestation(bytes _payload) pure returns(bool)
func (_AttestationHarness *AttestationHarnessSession) IsAttestation(_payload []byte) (bool, error) {
	return _AttestationHarness.Contract.IsAttestation(&_AttestationHarness.CallOpts, _payload)
}

// IsAttestation is a free data retrieval call binding the contract method 0x3ae7034d.
//
// Solidity: function isAttestation(bytes _payload) pure returns(bool)
func (_AttestationHarness *AttestationHarnessCallerSession) IsAttestation(_payload []byte) (bool, error) {
	return _AttestationHarness.Contract.IsAttestation(&_AttestationHarness.CallOpts, _payload)
}

// NotarySignature is a free data retrieval call binding the contract method 0xa43aa286.
//
// Solidity: function notarySignature(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCaller) NotarySignature(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "notarySignature", _type, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// NotarySignature is a free data retrieval call binding the contract method 0xa43aa286.
//
// Solidity: function notarySignature(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessSession) NotarySignature(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.NotarySignature(&_AttestationHarness.CallOpts, _type, _payload)
}

// NotarySignature is a free data retrieval call binding the contract method 0xa43aa286.
//
// Solidity: function notarySignature(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) NotarySignature(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.NotarySignature(&_AttestationHarness.CallOpts, _type, _payload)
}

// OffsetDestination is a free data retrieval call binding the contract method 0xd2c4428a.
//
// Solidity: function offsetDestination() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) OffsetDestination(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "offsetDestination")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetDestination is a free data retrieval call binding the contract method 0xd2c4428a.
//
// Solidity: function offsetDestination() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) OffsetDestination() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetDestination(&_AttestationHarness.CallOpts)
}

// OffsetDestination is a free data retrieval call binding the contract method 0xd2c4428a.
//
// Solidity: function offsetDestination() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) OffsetDestination() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetDestination(&_AttestationHarness.CallOpts)
}

// OffsetNonce is a free data retrieval call binding the contract method 0x569e1eaf.
//
// Solidity: function offsetNonce() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) OffsetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "offsetNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetNonce is a free data retrieval call binding the contract method 0x569e1eaf.
//
// Solidity: function offsetNonce() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) OffsetNonce() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetNonce(&_AttestationHarness.CallOpts)
}

// OffsetNonce is a free data retrieval call binding the contract method 0x569e1eaf.
//
// Solidity: function offsetNonce() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) OffsetNonce() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetNonce(&_AttestationHarness.CallOpts)
}

// OffsetOrigin is a free data retrieval call binding the contract method 0x320bfc44.
//
// Solidity: function offsetOrigin() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) OffsetOrigin(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "offsetOrigin")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetOrigin is a free data retrieval call binding the contract method 0x320bfc44.
//
// Solidity: function offsetOrigin() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) OffsetOrigin() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetOrigin(&_AttestationHarness.CallOpts)
}

// OffsetOrigin is a free data retrieval call binding the contract method 0x320bfc44.
//
// Solidity: function offsetOrigin() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) OffsetOrigin() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetOrigin(&_AttestationHarness.CallOpts)
}

// OffsetRoot is a free data retrieval call binding the contract method 0x5b42242d.
//
// Solidity: function offsetRoot() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) OffsetRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "offsetRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetRoot is a free data retrieval call binding the contract method 0x5b42242d.
//
// Solidity: function offsetRoot() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) OffsetRoot() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetRoot(&_AttestationHarness.CallOpts)
}

// OffsetRoot is a free data retrieval call binding the contract method 0x5b42242d.
//
// Solidity: function offsetRoot() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) OffsetRoot() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetRoot(&_AttestationHarness.CallOpts)
}

// OffsetSignature is a free data retrieval call binding the contract method 0x740261a6.
//
// Solidity: function offsetSignature() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) OffsetSignature(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "offsetSignature")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetSignature is a free data retrieval call binding the contract method 0x740261a6.
//
// Solidity: function offsetSignature() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) OffsetSignature() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetSignature(&_AttestationHarness.CallOpts)
}

// OffsetSignature is a free data retrieval call binding the contract method 0x740261a6.
//
// Solidity: function offsetSignature() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) OffsetSignature() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetSignature(&_AttestationHarness.CallOpts)
}

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122095cb27a2580dd37e71bf433a6c11f729b0c1c7a11be0d43af5fda064494dceaf64736f6c63430008110033",
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

// SynapseTypesMetaData contains all meta data concerning the SynapseTypes contract.
var SynapseTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212201279501b008bf08993ce56d023466326263dbae453f78986b7c4b70e26d0950a64736f6c63430008110033",
}

// SynapseTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use SynapseTypesMetaData.ABI instead.
var SynapseTypesABI = SynapseTypesMetaData.ABI

// SynapseTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SynapseTypesMetaData.Bin instead.
var SynapseTypesBin = SynapseTypesMetaData.Bin

// DeploySynapseTypes deploys a new Ethereum contract, binding an instance of SynapseTypes to it.
func DeploySynapseTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SynapseTypes, error) {
	parsed, err := SynapseTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SynapseTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SynapseTypes{SynapseTypesCaller: SynapseTypesCaller{contract: contract}, SynapseTypesTransactor: SynapseTypesTransactor{contract: contract}, SynapseTypesFilterer: SynapseTypesFilterer{contract: contract}}, nil
}

// SynapseTypes is an auto generated Go binding around an Ethereum contract.
type SynapseTypes struct {
	SynapseTypesCaller     // Read-only binding to the contract
	SynapseTypesTransactor // Write-only binding to the contract
	SynapseTypesFilterer   // Log filterer for contract events
}

// SynapseTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type SynapseTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SynapseTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SynapseTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SynapseTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SynapseTypesSession struct {
	Contract     *SynapseTypes     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SynapseTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SynapseTypesCallerSession struct {
	Contract *SynapseTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// SynapseTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SynapseTypesTransactorSession struct {
	Contract     *SynapseTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// SynapseTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type SynapseTypesRaw struct {
	Contract *SynapseTypes // Generic contract binding to access the raw methods on
}

// SynapseTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SynapseTypesCallerRaw struct {
	Contract *SynapseTypesCaller // Generic read-only contract binding to access the raw methods on
}

// SynapseTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SynapseTypesTransactorRaw struct {
	Contract *SynapseTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSynapseTypes creates a new instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypes(address common.Address, backend bind.ContractBackend) (*SynapseTypes, error) {
	contract, err := bindSynapseTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SynapseTypes{SynapseTypesCaller: SynapseTypesCaller{contract: contract}, SynapseTypesTransactor: SynapseTypesTransactor{contract: contract}, SynapseTypesFilterer: SynapseTypesFilterer{contract: contract}}, nil
}

// NewSynapseTypesCaller creates a new read-only instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesCaller(address common.Address, caller bind.ContractCaller) (*SynapseTypesCaller, error) {
	contract, err := bindSynapseTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesCaller{contract: contract}, nil
}

// NewSynapseTypesTransactor creates a new write-only instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*SynapseTypesTransactor, error) {
	contract, err := bindSynapseTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesTransactor{contract: contract}, nil
}

// NewSynapseTypesFilterer creates a new log filterer instance of SynapseTypes, bound to a specific deployed contract.
func NewSynapseTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*SynapseTypesFilterer, error) {
	contract, err := bindSynapseTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SynapseTypesFilterer{contract: contract}, nil
}

// bindSynapseTypes binds a generic wrapper to an already deployed contract.
func bindSynapseTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SynapseTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseTypes *SynapseTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseTypes.Contract.SynapseTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseTypes *SynapseTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseTypes.Contract.SynapseTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseTypes *SynapseTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseTypes.Contract.SynapseTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SynapseTypes *SynapseTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SynapseTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SynapseTypes *SynapseTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SynapseTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SynapseTypes *SynapseTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SynapseTypes.Contract.contract.Transact(opts, method, params...)
}

// TypedMemViewMetaData contains all meta data concerning the TypedMemView contract.
var TypedMemViewMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"LOW_12_MASK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NULL\",\"outputs\":[{\"internalType\":\"bytes29\",\"name\":\"\",\"type\":\"bytes29\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TWELVE_BYTES\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b286bae7": "LOW_12_MASK()",
		"f26be3fc": "NULL()",
		"406cba16": "TWELVE_BYTES()",
	},
	Bin: "0x61011561003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060475760003560e01c8063406cba1614604c578063b286bae714606a578063f26be3fc146089575b600080fd5b6053606081565b60405160ff90911681526020015b60405180910390f35b607c6bffffffffffffffffffffffff81565b6040519081526020016061565b60af7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0000009091168152602001606156fea2646970667358221220555793bcb7701a0ae6cc57382f042c46031b637a19b2f3d8d8e97aa348a9448364736f6c63430008110033",
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

// LOW12MASK is a free data retrieval call binding the contract method 0xb286bae7.
//
// Solidity: function LOW_12_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCaller) LOW12MASK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "LOW_12_MASK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOW12MASK is a free data retrieval call binding the contract method 0xb286bae7.
//
// Solidity: function LOW_12_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewSession) LOW12MASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW12MASK(&_TypedMemView.CallOpts)
}

// LOW12MASK is a free data retrieval call binding the contract method 0xb286bae7.
//
// Solidity: function LOW_12_MASK() view returns(uint256)
func (_TypedMemView *TypedMemViewCallerSession) LOW12MASK() (*big.Int, error) {
	return _TypedMemView.Contract.LOW12MASK(&_TypedMemView.CallOpts)
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

// TWELVEBYTES is a free data retrieval call binding the contract method 0x406cba16.
//
// Solidity: function TWELVE_BYTES() view returns(uint8)
func (_TypedMemView *TypedMemViewCaller) TWELVEBYTES(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TypedMemView.contract.Call(opts, &out, "TWELVE_BYTES")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TWELVEBYTES is a free data retrieval call binding the contract method 0x406cba16.
//
// Solidity: function TWELVE_BYTES() view returns(uint8)
func (_TypedMemView *TypedMemViewSession) TWELVEBYTES() (uint8, error) {
	return _TypedMemView.Contract.TWELVEBYTES(&_TypedMemView.CallOpts)
}

// TWELVEBYTES is a free data retrieval call binding the contract method 0x406cba16.
//
// Solidity: function TWELVE_BYTES() view returns(uint8)
func (_TypedMemView *TypedMemViewCallerSession) TWELVEBYTES() (uint8, error) {
	return _TypedMemView.Contract.TWELVEBYTES(&_TypedMemView.CallOpts)
}
