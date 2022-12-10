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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b3a5cda562bd3c850eea4a3eddea7cdd22ea7399cbe3c887fecb12c9b0e3b09a64736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"agentSignatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestationData\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"}],\"name\":\"attestationDomains\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"attestationKey\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedDestination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedDomains\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedKey\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedOrigin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"castToAttestation\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_guardSignatures\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_notarySignatures\",\"type\":\"bytes\"}],\"name\":\"formatAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"formatAttestationData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"guardSignature\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"guardSignatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"isAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"notarySignature\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"notarySignatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetAgentSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetDestination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetFirstSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"setIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f24caf61": "agentSignatures(uint40,bytes)",
		"9ffb971e": "attestationData(uint40,bytes)",
		"a104a5e5": "attestationDataLength()",
		"65dfb428": "attestationDomains(uint32,uint32)",
		"9f668e20": "attestationKey(uint32,uint32,uint32)",
		"3b830f3b": "attestedDestination(uint40,bytes)",
		"c231bd8f": "attestedDomains(uint40,bytes)",
		"79ce92a9": "attestedKey(uint40,bytes)",
		"8b445f51": "attestedNonce(uint40,bytes)",
		"badad7db": "attestedOrigin(uint40,bytes)",
		"91eedc1d": "attestedRoot(uint40,bytes)",
		"c2e19ed2": "castToAttestation(uint40,bytes)",
		"036227a2": "formatAttestation(bytes,bytes,bytes)",
		"2951eae3": "formatAttestationData(uint32,uint32,uint32,bytes32)",
		"4a0cfe0e": "guardSignature(uint40,bytes)",
		"dc1e976d": "guardSignatures(uint40,bytes)",
		"3ae7034d": "isAttestation(bytes)",
		"a43aa286": "notarySignature(uint40,bytes)",
		"cbe9d784": "notarySignatures(uint40,bytes)",
		"ce533592": "offsetAgentSignatures()",
		"d2c4428a": "offsetDestination()",
		"97d91f1a": "offsetFirstSignature()",
		"569e1eaf": "offsetNonce()",
		"320bfc44": "offsetOrigin()",
		"5b42242d": "offsetRoot()",
		"40a5737f": "setIndex(uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611d72806100206000396000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c806397d91f1a116100ee578063c231bd8f11610097578063ce53359211610071578063ce5335921461033f578063d2c4428a146103b7578063dc1e976d146103be578063f24caf61146103d157600080fd5b8063c231bd8f1461036c578063c2e19ed21461037f578063cbe9d7841461039257600080fd5b8063a104a5e5116100c8578063a104a5e51461033f578063a43aa28614610346578063badad7db1461035957600080fd5b806397d91f1a146103115780639f668e20146103195780639ffb971e1461032c57600080fd5b80634a0cfe0e1161015b57806365dfb4281161013557806365dfb4281461028f57806379ce92a9146102bb5780638b445f51146102eb57806391eedc1d146102fe57600080fd5b80634a0cfe0e14610260578063569e1eaf146102815780635b42242d1461028857600080fd5b80633ae7034d1161018c5780633ae7034d146102005780633b830f3b1461022357806340a5737f1461024b57600080fd5b8063036227a2146101b35780632951eae3146101dc578063320bfc44146101ef575b600080fd5b6101c66101c136600461196b565b6103fe565b6040516101d39190611a57565b60405180910390f35b6101c66101ea366004611a83565b610415565b60005b6040519081526020016101d3565b61021361020e366004611ace565b61047c565b60405190151581526020016101d3565b610236610231366004611b03565b61049b565b60405163ffffffff90911681526020016101d3565b61025e610259366004611b5d565b600055565b005b61027361026e366004611b03565b6104b5565b6040516101d3929190611b76565b60086101f2565b600c6101f2565b6102a261029d366004611b96565b61050c565b60405167ffffffffffffffff90911681526020016101d3565b6102ce6102c9366004611b03565b61052a565b6040516bffffffffffffffffffffffff90911681526020016101d3565b6102366102f9366004611b03565b610544565b6101f261030c366004611b03565b61055e565b6101f2610578565b6102ce610327366004611bc9565b61058b565b61027361033a366004611b03565b6105bd565b602c6101f2565b610273610354366004611b03565b6105da565b610236610367366004611b03565b610605565b6102a261037a366004611b03565b61061f565b61027361038d366004611b03565b610639565b6103a56103a0366004611b03565b610648565b60405160ff90911681526020016101d3565b60046101f2565b6103a56103cc366004611b03565b610662565b6103e46103df366004611b03565b61067c565b6040805160ff9384168152929091166020830152016101d3565b606061040b8484846106a2565b90505b9392505050565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e087811b8216602084015286811b8216602484015285901b166028820152602c80820184905282518083039091018152604c9091019091525b949350505050565b600061049561048a83610812565b62ffffff1916610823565b92915050565b600061040e6104aa83856108b9565b62ffffff19166108dd565b6000606060006104e06000546104d487876108b990919063ffffffff16565b62ffffff191690610908565b90506104f162ffffff1982166109df565b61050062ffffff198316610a03565b92509250509250929050565b600067ffffffff00000000602084901b1663ffffffff83161761040e565b600061040e61053983856108b9565b62ffffff1916610a56565b600061040e61055383856108b9565b62ffffff1916610a82565b600061040e61056d83856108b9565b62ffffff1916610aae565b6000610586602c6002611c3b565b905090565b600063ffffffff8216602084901b67ffffffff0000000016604086901b6bffffffff000000000000000016171761040b565b60006060816104e06105cf85876108b9565b62ffffff1916610ada565b6000606060006104e06000546105f987876108b990919063ffffffff16565b62ffffff191690610b0c565b600061040e61061483856108b9565b62ffffff1916610bf1565b600061040e61062e83856108b9565b62ffffff1916610c1d565b6000606060006104e084610812565b600061040e61065783856108b9565b62ffffff1916610c49565b600061040e61067183856108b9565b62ffffff1916610c6b565b60008061069761068c84866108b9565b62ffffff1916610c95565b915091509250929050565b606060006106af84610cc3565b905060006106bc84610cc3565b60408051600480825260a0820190925291925061ff00600885901b1660ff84161791600091602082016080803683370190505090506106fa88610db5565b8160008151811061070d5761070d611c4e565b62ffffff199092166020928302919091018201526040516107739161075f9185910160f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016815260020190565b604051602081830303815290604052610db5565b8160018151811061078657610786611c4e565b62ffffff19909216602092830291909101909101526107a487610db5565b816002815181106107b7576107b7611c4e565b62ffffff19909216602092830291909101909101526107d586610db5565b816003815181106107e8576107e8611c4e565b62ffffff199092166020928302919091019091015261080681610dc1565b98975050505050505050565b6000610495826401010000006108b9565b6000601882901c6bffffffffffffffffffffffff16610844602c6002611c3b565b8110156108545750600092915050565b60008061086085610e20565b60ff918216935016905060006108768284611c3b565b90508060000361088c5750600095945050505050565b610897604182611c7d565b6108a3602c6002611c3b565b6108ad9190611c3b565b90931495945050505050565b8151600090602084016108d464ffffffffff85168284610e4b565b95945050505050565b6000816108f562ffffff198216640101000000610e92565b5061040e62ffffff198416600480610fac565b60008261092062ffffff198216640101000000610e92565b50600061092c85610e20565b5090508060ff1684106109a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4f7574206f662072616e6765000000000000000000000000000000000000000060448201526064015b60405180910390fd5b6108d46109ae604186611c7d565b6109ba602c6002611c3b565b6109c49190611c3b565b62ffffff1987169060416301000000610fdc565b5092915050565b60008060606109ef816018611c3b565b6109f99190611c3b565b9290921c92915050565b6060600080610a208460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506040519150819250610a458483602001611047565b508181016020016040529052919050565b600081610a6e62ffffff198216640101000000610e92565b5061040e62ffffff1984166000600c610fac565b600081610a9a62ffffff198216640101000000610e92565b5061040e62ffffff19841660086004610fac565b600081610ac662ffffff198216640101000000610e92565b5061040e62ffffff198416600c602061122e565b600081610af262ffffff198216640101000000610e92565b5061040e62ffffff1984166000602c640101010000610fdc565b600082610b2462ffffff198216640101000000610e92565b50600080610b3186610e20565b915091508060ff168510610ba1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4f7574206f662072616e676500000000000000000000000000000000000000006044820152606401610997565b610be76041610bb360ff851688611c3b565b610bbd9190611c7d565b610bc9602c6002611c3b565b610bd39190611c3b565b62ffffff1988169060416301000000610fdc565b9695505050505050565b600081610c0962ffffff198216640101000000610e92565b5061040e62ffffff19841660006004610fac565b600081610c3562ffffff198216640101000000610e92565b5061040e62ffffff19841660006008610fac565b600081610c6162ffffff198216640101000000610e92565b5061047483610e20565b600081610c8362ffffff198216640101000000610e92565b50610c8d83610e20565b509392505050565b60008082610cae62ffffff198216640101000000610e92565b50610cb884610e20565b909590945092505050565b60008060418351610cd49190611c94565b8351909150610ce4604183611c7d565b14610d4b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f217369676e6174757265734c656e6774680000000000000000000000000000006044820152606401610997565b60ff8110610495576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f546f6f206d616e79207369676e617475726573000000000000000000000000006044820152606401610997565b600061049582826108b9565b6040516060906000610dd684602084016113dc565b90506000610df28260181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1690506000610e0d83611476565b9184525082016020016040525092915050565b60008080610e3762ffffff198516602c6002610fac565b60ff600882901c8116969116945092505050565b600080610e588385611c3b565b9050604051811115610e68575060005b80600003610e7d5762ffffff1991505061040e565b5050606092831b9190911790911b1760181b90565b6000610e9e838361148c565b610fa5576000610ebc610eb0856109df565b64ffffffffff166114ae565b9150506000610ed18464ffffffffff166114ae565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109979190611a57565b5090919050565b6000610fb9826020611ccf565b610fc4906008611ce8565b60ff16610fd285858561122e565b901c949350505050565b600080610fe886611598565b6bffffffffffffffffffffffff169050611001866115bf565b8461100c8784611c3b565b6110169190611c3b565b11156110295762ffffff19915050610474565b6110338582611c3b565b9050610be78364ffffffffff168286610e4b565b600062ffffff19808416036110b8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610997565b6110c1836115f8565b611127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610997565b60006111418460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061115c85611598565b6bffffffffffffffffffffffff1690506000806040519150858211156111825760206060fd5b8386858560045afa9050806111f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610997565b6112236111ff886109df565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b60008160ff166000036112435750600061040e565b61125b8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1661127660ff841685611c3b565b11156112f9576112c661128885611598565b6bffffffffffffffffffffffff166112ae8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16611634565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109979190611a57565b60208260ff161115611367576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610997565b60088202600061137686611598565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b6000604051828111156113ef5760206060fd5b506000805b845181101561146657600085828151811061141157611411611c4e565b6020026020010151905061142781848701611047565b506114408160181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff168301925050808061145e90611d04565b9150506113f4565b50606083901b811760181b610474565b6000611481826117c4565b610495906020611c7d565b60008164ffffffffff1661149f846109df565b64ffffffffff16149392505050565b600080601f5b600f8160ff1611156115215760006114cd826008611ce8565b60ff1685901c90506114de81611803565b61ffff16841793508160ff166010146114f957601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff016114b4565b50600f5b60ff8160ff16101561159257600061153e826008611ce8565b60ff1685901c905061154f81611803565b61ffff16831792508160ff1660001461156a57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611525565b50915091565b6000806115a760606018611c3b565b9290921c6bffffffffffffffffffffffff1692915050565b60006115d98260181c6bffffffffffffffffffffffff1690565b6115e283611598565b016bffffffffffffffffffffffff169050919050565b6000611603826109df565b64ffffffffff1664ffffffffff0361161d57506000919050565b6000611628836115bf565b60405110199392505050565b60606000611641866114ae565b915050600061164f866114ae565b915050600061165d866114ae565b915050600061166b866114ae565b604080517f54797065644d656d566965772f696e646578202d204f76657272616e2074686560208201527f20766965772e20536c6963652069732061742030780000000000000000000000818301527fffffffffffff000000000000000000000000000000000000000000000000000060d098891b811660558301527f2077697468206c656e6774682030780000000000000000000000000000000000605b830181905297891b8116606a8301527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060708301527f7800000000000000000000000000000000000000000000000000000000000000609083015295881b861660918201526097810196909652951b90921660a684015250507f2e0000000000000000000000000000000000000000000000000000000000000060ac8201528151808203608d01815260ad90910190915295945050505050565b600060206117e08360181c6bffffffffffffffffffffffff1690565b6117f9906bffffffffffffffffffffffff16601f611c3b565b6104959190611c94565b600061181560048360ff16901c611835565b60ff1661ffff919091161760081b61182c82611835565b60ff1617919050565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f8416918290811061188257611882611c4e565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126118d157600080fd5b813567ffffffffffffffff808211156118ec576118ec611891565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190828211818310171561193257611932611891565b8160405283815286602085880101111561194b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060006060848603121561198057600080fd5b833567ffffffffffffffff8082111561199857600080fd5b6119a4878388016118c0565b945060208601359150808211156119ba57600080fd5b6119c6878388016118c0565b935060408601359150808211156119dc57600080fd5b506119e9868287016118c0565b9150509250925092565b6000815180845260005b81811015611a19576020818501810151868301820152016119fd565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b60208152600061040e60208301846119f3565b803563ffffffff81168114611a7e57600080fd5b919050565b60008060008060808587031215611a9957600080fd5b611aa285611a6a565b9350611ab060208601611a6a565b9250611abe60408601611a6a565b9396929550929360600135925050565b600060208284031215611ae057600080fd5b813567ffffffffffffffff811115611af757600080fd5b610474848285016118c0565b60008060408385031215611b1657600080fd5b823564ffffffffff81168114611b2b57600080fd5b9150602083013567ffffffffffffffff811115611b4757600080fd5b611b53858286016118c0565b9150509250929050565b600060208284031215611b6f57600080fd5b5035919050565b64ffffffffff8316815260406020820152600061040b60408301846119f3565b60008060408385031215611ba957600080fd5b611bb283611a6a565b9150611bc060208401611a6a565b90509250929050565b600080600060608486031215611bde57600080fd5b611be784611a6a565b9250611bf560208501611a6a565b9150611c0360408501611a6a565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561049557610495611c0c565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b808202811582820484141761049557610495611c0c565b600082611cca577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60ff828116828216039081111561049557610495611c0c565b60ff81811683821602908116908181146109d8576109d8611c0c565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611d3557611d35611c0c565b506001019056fea2646970667358221220a4af5d7487275ed170b0735b512787d73f1cb74d41c65e9c414a31c0d03fae7d64736f6c63430008110033",
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

// AgentSignatures is a free data retrieval call binding the contract method 0xf24caf61.
//
// Solidity: function agentSignatures(uint40 _type, bytes _payload) pure returns(uint8, uint8)
func (_AttestationHarness *AttestationHarnessCaller) AgentSignatures(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint8, uint8, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "agentSignatures", _type, _payload)

	if err != nil {
		return *new(uint8), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return out0, out1, err

}

// AgentSignatures is a free data retrieval call binding the contract method 0xf24caf61.
//
// Solidity: function agentSignatures(uint40 _type, bytes _payload) pure returns(uint8, uint8)
func (_AttestationHarness *AttestationHarnessSession) AgentSignatures(_type *big.Int, _payload []byte) (uint8, uint8, error) {
	return _AttestationHarness.Contract.AgentSignatures(&_AttestationHarness.CallOpts, _type, _payload)
}

// AgentSignatures is a free data retrieval call binding the contract method 0xf24caf61.
//
// Solidity: function agentSignatures(uint40 _type, bytes _payload) pure returns(uint8, uint8)
func (_AttestationHarness *AttestationHarnessCallerSession) AgentSignatures(_type *big.Int, _payload []byte) (uint8, uint8, error) {
	return _AttestationHarness.Contract.AgentSignatures(&_AttestationHarness.CallOpts, _type, _payload)
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

// FormatAttestation is a free data retrieval call binding the contract method 0x036227a2.
//
// Solidity: function formatAttestation(bytes _data, bytes _guardSignatures, bytes _notarySignatures) view returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatAttestation(opts *bind.CallOpts, _data []byte, _guardSignatures []byte, _notarySignatures []byte) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatAttestation", _data, _guardSignatures, _notarySignatures)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatAttestation is a free data retrieval call binding the contract method 0x036227a2.
//
// Solidity: function formatAttestation(bytes _data, bytes _guardSignatures, bytes _notarySignatures) view returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatAttestation(_data []byte, _guardSignatures []byte, _notarySignatures []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, _data, _guardSignatures, _notarySignatures)
}

// FormatAttestation is a free data retrieval call binding the contract method 0x036227a2.
//
// Solidity: function formatAttestation(bytes _data, bytes _guardSignatures, bytes _notarySignatures) view returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatAttestation(_data []byte, _guardSignatures []byte, _notarySignatures []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestation(&_AttestationHarness.CallOpts, _data, _guardSignatures, _notarySignatures)
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

// GuardSignature is a free data retrieval call binding the contract method 0x4a0cfe0e.
//
// Solidity: function guardSignature(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCaller) GuardSignature(opts *bind.CallOpts, _type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "guardSignature", _type, _payload)

	if err != nil {
		return *new(*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return out0, out1, err

}

// GuardSignature is a free data retrieval call binding the contract method 0x4a0cfe0e.
//
// Solidity: function guardSignature(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessSession) GuardSignature(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.GuardSignature(&_AttestationHarness.CallOpts, _type, _payload)
}

// GuardSignature is a free data retrieval call binding the contract method 0x4a0cfe0e.
//
// Solidity: function guardSignature(uint40 _type, bytes _payload) view returns(uint40, bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) GuardSignature(_type *big.Int, _payload []byte) (*big.Int, []byte, error) {
	return _AttestationHarness.Contract.GuardSignature(&_AttestationHarness.CallOpts, _type, _payload)
}

// GuardSignatures is a free data retrieval call binding the contract method 0xdc1e976d.
//
// Solidity: function guardSignatures(uint40 _type, bytes _payload) pure returns(uint8)
func (_AttestationHarness *AttestationHarnessCaller) GuardSignatures(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint8, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "guardSignatures", _type, _payload)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GuardSignatures is a free data retrieval call binding the contract method 0xdc1e976d.
//
// Solidity: function guardSignatures(uint40 _type, bytes _payload) pure returns(uint8)
func (_AttestationHarness *AttestationHarnessSession) GuardSignatures(_type *big.Int, _payload []byte) (uint8, error) {
	return _AttestationHarness.Contract.GuardSignatures(&_AttestationHarness.CallOpts, _type, _payload)
}

// GuardSignatures is a free data retrieval call binding the contract method 0xdc1e976d.
//
// Solidity: function guardSignatures(uint40 _type, bytes _payload) pure returns(uint8)
func (_AttestationHarness *AttestationHarnessCallerSession) GuardSignatures(_type *big.Int, _payload []byte) (uint8, error) {
	return _AttestationHarness.Contract.GuardSignatures(&_AttestationHarness.CallOpts, _type, _payload)
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

// NotarySignatures is a free data retrieval call binding the contract method 0xcbe9d784.
//
// Solidity: function notarySignatures(uint40 _type, bytes _payload) pure returns(uint8)
func (_AttestationHarness *AttestationHarnessCaller) NotarySignatures(opts *bind.CallOpts, _type *big.Int, _payload []byte) (uint8, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "notarySignatures", _type, _payload)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// NotarySignatures is a free data retrieval call binding the contract method 0xcbe9d784.
//
// Solidity: function notarySignatures(uint40 _type, bytes _payload) pure returns(uint8)
func (_AttestationHarness *AttestationHarnessSession) NotarySignatures(_type *big.Int, _payload []byte) (uint8, error) {
	return _AttestationHarness.Contract.NotarySignatures(&_AttestationHarness.CallOpts, _type, _payload)
}

// NotarySignatures is a free data retrieval call binding the contract method 0xcbe9d784.
//
// Solidity: function notarySignatures(uint40 _type, bytes _payload) pure returns(uint8)
func (_AttestationHarness *AttestationHarnessCallerSession) NotarySignatures(_type *big.Int, _payload []byte) (uint8, error) {
	return _AttestationHarness.Contract.NotarySignatures(&_AttestationHarness.CallOpts, _type, _payload)
}

// OffsetAgentSignatures is a free data retrieval call binding the contract method 0xce533592.
//
// Solidity: function offsetAgentSignatures() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) OffsetAgentSignatures(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "offsetAgentSignatures")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetAgentSignatures is a free data retrieval call binding the contract method 0xce533592.
//
// Solidity: function offsetAgentSignatures() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) OffsetAgentSignatures() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetAgentSignatures(&_AttestationHarness.CallOpts)
}

// OffsetAgentSignatures is a free data retrieval call binding the contract method 0xce533592.
//
// Solidity: function offsetAgentSignatures() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) OffsetAgentSignatures() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetAgentSignatures(&_AttestationHarness.CallOpts)
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

// OffsetFirstSignature is a free data retrieval call binding the contract method 0x97d91f1a.
//
// Solidity: function offsetFirstSignature() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCaller) OffsetFirstSignature(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "offsetFirstSignature")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OffsetFirstSignature is a free data retrieval call binding the contract method 0x97d91f1a.
//
// Solidity: function offsetFirstSignature() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessSession) OffsetFirstSignature() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetFirstSignature(&_AttestationHarness.CallOpts)
}

// OffsetFirstSignature is a free data retrieval call binding the contract method 0x97d91f1a.
//
// Solidity: function offsetFirstSignature() pure returns(uint256)
func (_AttestationHarness *AttestationHarnessCallerSession) OffsetFirstSignature() (*big.Int, error) {
	return _AttestationHarness.Contract.OffsetFirstSignature(&_AttestationHarness.CallOpts)
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

// SetIndex is a paid mutator transaction binding the contract method 0x40a5737f.
//
// Solidity: function setIndex(uint256 index) returns()
func (_AttestationHarness *AttestationHarnessTransactor) SetIndex(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _AttestationHarness.contract.Transact(opts, "setIndex", index)
}

// SetIndex is a paid mutator transaction binding the contract method 0x40a5737f.
//
// Solidity: function setIndex(uint256 index) returns()
func (_AttestationHarness *AttestationHarnessSession) SetIndex(index *big.Int) (*types.Transaction, error) {
	return _AttestationHarness.Contract.SetIndex(&_AttestationHarness.TransactOpts, index)
}

// SetIndex is a paid mutator transaction binding the contract method 0x40a5737f.
//
// Solidity: function setIndex(uint256 index) returns()
func (_AttestationHarness *AttestationHarnessTransactorSession) SetIndex(index *big.Int) (*types.Transaction, error) {
	return _AttestationHarness.Contract.SetIndex(&_AttestationHarness.TransactOpts, index)
}

// ByteStringMetaData contains all meta data concerning the ByteString contract.
var ByteStringMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122020054a20c26a7fd65066415652aba36c70a6d3d42f215e6d034fb76af4c8ea7c64736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c7ed0cc2f4b8169de426526dd705ef97fb2a569c1222329285a320c37bc2f2c664736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220a25d33771efae5cc031dd1529c328e960bde285785c03b9f1423ae2f72f7926564736f6c63430008110033",
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
