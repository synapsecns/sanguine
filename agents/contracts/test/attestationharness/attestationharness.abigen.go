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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bcc85e029dd4d1bca2480d411ba280a68c7e88807a6773a6ce976182da1b0c3564736f6c63430008110033",
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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"agentSignatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestationData\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"}],\"name\":\"attestationDomains\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"}],\"name\":\"attestationKey\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedDestination\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedDomains\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedKey\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedNonce\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedOrigin\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"attestedRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"castToAttestation\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_guardSignatures\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_notarySignatures\",\"type\":\"bytes\"}],\"name\":\"formatAttestation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_origin\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_destination\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_nonce\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"_root\",\"type\":\"bytes32\"}],\"name\":\"formatAttestationData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_guardSignatures\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_notarySignatures\",\"type\":\"bytes\"}],\"name\":\"formatAttestationFromViews\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"guardSignature\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"guardSignatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"isAttestation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"notarySignature\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_type\",\"type\":\"uint40\"},{\"internalType\":\"bytes\",\"name\":\"_payload\",\"type\":\"bytes\"}],\"name\":\"notarySignatures\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetAgentSignatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetDestination\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetFirstSignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetOrigin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"offsetRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"setIndex\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"_attestationDomains\",\"type\":\"uint64\"}],\"name\":\"unpackDomains\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint96\",\"name\":\"_attestationKey\",\"type\":\"uint96\"}],\"name\":\"unpackKey\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"domains\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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
		"794dd214": "formatAttestationFromViews(bytes,bytes,bytes)",
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
		"7214101b": "unpackDomains(uint64)",
		"308514be": "unpackKey(uint96)",
	},
	Bin: "0x608060405234801561001057600080fd5b50611eeb806100206000396000f3fe608060405234801561001057600080fd5b50600436106101cf5760003560e01c80638b445f5111610104578063badad7db116100a2578063ce53359211610071578063ce533592146103db578063d2c4428a14610453578063dc1e976d1461045a578063f24caf611461046d57600080fd5b8063badad7db146103f5578063c231bd8f14610408578063c2e19ed21461041b578063cbe9d7841461042e57600080fd5b80639f668e20116100de5780639f668e20146103b55780639ffb971e146103c8578063a104a5e5146103db578063a43aa286146103e257600080fd5b80638b445f511461038757806391eedc1d1461039a57806397d91f1a146103ad57600080fd5b80634a0cfe0e1161017157806365dfb4281161014b57806365dfb428146102e85780637214101b14610314578063794dd2141461034457806379ce92a91461035757600080fd5b80634a0cfe0e146102b9578063569e1eaf146102da5780635b42242d146102e157600080fd5b8063320bfc44116101ad578063320bfc44146102485780633ae7034d146102595780633b830f3b1461027c57806340a5737f146102a457600080fd5b8063036227a2146101d45780632951eae3146101fd578063308514be14610210575b600080fd5b6101e76101e2366004611a8c565b61049a565b6040516101f49190611b78565b60405180910390f35b6101e761020b366004611ba4565b6104b1565b61022361021e366004611bef565b610518565b6040805167ffffffffffffffff909316835263ffffffff9091166020830152016101f4565b60005b6040519081526020016101f4565b61026c610267366004611c1d565b61053b565b60405190151581526020016101f4565b61028f61028a366004611c52565b61055a565b60405163ffffffff90911681526020016101f4565b6102b76102b2366004611cac565b600055565b005b6102cc6102c7366004611c52565b610574565b6040516101f4929190611cc5565b600861024b565b600c61024b565b6102fb6102f6366004611ce5565b6105cb565b60405167ffffffffffffffff90911681526020016101f4565b610327610322366004611d18565b6105e9565b6040805163ffffffff9384168152929091166020830152016101f4565b6101e7610352366004611a8c565b610600565b61036a610365366004611c52565b61062b565b6040516bffffffffffffffffffffffff90911681526020016101f4565b61028f610395366004611c52565b610645565b61024b6103a8366004611c52565b61065f565b61024b610679565b61036a6103c3366004611d42565b61068c565b6102cc6103d6366004611c52565b6106be565b602c61024b565b6102cc6103f0366004611c52565b6106db565b61028f610403366004611c52565b610706565b6102fb610416366004611c52565b610720565b6102cc610429366004611c52565b61073a565b61044161043c366004611c52565b610749565b60405160ff90911681526020016101f4565b600461024b565b610441610468366004611c52565b610763565b61048061047b366004611c52565b61077d565b6040805160ff9384168152929091166020830152016101f4565b60606104a78484846107a3565b90505b9392505050565b604080517fffffffff0000000000000000000000000000000000000000000000000000000060e087811b8216602084015286811b8216602484015285901b166028820152602c80820184905282518083039091018152604c9091019091525b949350505050565b600080602083901c67ffffffffffffffff1663ffffffff84165b91509150915091565b6000610554610549836107c3565b62ffffff19166107d4565b92915050565b60006104aa610569838561086a565b62ffffff191661088e565b60006060600061059f600054610593878761086a90919063ffffffff16565b62ffffff1916906108b9565b90506105b062ffffff198216610990565b6105bf62ffffff1983166109b4565b92509250509250929050565b600067ffffffff00000000602084901b1663ffffffff8316176104aa565b60008063ffffffff602084901c8116908416610532565b60606104a761061085600061086a565b61061b85600061086a565b61062685600061086a565b610a07565b60006104aa61063a838561086a565b62ffffff1916610b69565b60006104aa610654838561086a565b62ffffff1916610b95565b60006104aa61066e838561086a565b62ffffff1916610bc1565b6000610687602c6002611db4565b905090565b600063ffffffff8216602084901b67ffffffff0000000016604086901b6bffffffff00000000000000001617176104a7565b600060608161059f6106d0858761086a565b62ffffff1916610bed565b60006060600061059f6000546106fa878761086a90919063ffffffff16565b62ffffff191690610c1f565b60006104aa610715838561086a565b62ffffff1916610d04565b60006104aa61072f838561086a565b62ffffff1916610d30565b60006060600061059f846107c3565b60006104aa610758838561086a565b62ffffff1916610d5c565b60006104aa610772838561086a565b62ffffff1916610d7e565b60008061079861078d848661086a565b62ffffff1916610da8565b915091509250929050565b60606104a76107b185610dd6565b6107ba85610dd6565b61062685610dd6565b60006105548264010100000061086a565b6000601882901c6bffffffffffffffffffffffff166107f5602c6002611db4565b8110156108055750600092915050565b60008061081185610de2565b60ff918216935016905060006108278284611db4565b90508060000361083d5750600095945050505050565b610848604182611dc7565b610854602c6002611db4565b61085e9190611db4565b90931495945050505050565b81516000906020840161088564ffffffffff85168284610e0d565b95945050505050565b6000816108a662ffffff198216640101000000610e54565b506104aa62ffffff198416600480610f6e565b6000826108d162ffffff198216640101000000610e54565b5060006108dd85610de2565b5090508060ff168410610951576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4f7574206f662072616e6765000000000000000000000000000000000000000060448201526064015b60405180910390fd5b61088561095f604186611dc7565b61096b602c6002611db4565b6109759190611db4565b62ffffff1987169060416301000000610f9e565b5092915050565b60008060606109a0816018611db4565b6109aa9190611db4565b9290921c92915050565b60606000806109d18460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16905060405191508192506109f68483602001611009565b508181016020016040529052919050565b60606000610a14846111f0565b90506000610a21846111f0565b60408051600480825260a0820190925291925061ff00600885901b1660ff84161791600091602082016080803683370190505090508781600081518110610a6a57610a6a611dde565b62ffffff19909216602092830291909101820152604051610ad091610abc9185910160f09190911b7fffff00000000000000000000000000000000000000000000000000000000000016815260020190565b604051602081830303815290604052610dd6565b81600181518110610ae357610ae3611dde565b602002602001019062ffffff1916908162ffffff1916815250508681600281518110610b1157610b11611dde565b602002602001019062ffffff1916908162ffffff1916815250508581600381518110610b3f57610b3f611dde565b62ffffff1990921660209283029190910190910152610b5d816112f0565b98975050505050505050565b600081610b8162ffffff198216640101000000610e54565b506104aa62ffffff1984166000600c610f6e565b600081610bad62ffffff198216640101000000610e54565b506104aa62ffffff19841660086004610f6e565b600081610bd962ffffff198216640101000000610e54565b506104aa62ffffff198416600c602061134f565b600081610c0562ffffff198216640101000000610e54565b506104aa62ffffff1984166000602c640101010000610f9e565b600082610c3762ffffff198216640101000000610e54565b50600080610c4486610de2565b915091508060ff168510610cb4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f4f7574206f662072616e676500000000000000000000000000000000000000006044820152606401610948565b610cfa6041610cc660ff851688611db4565b610cd09190611dc7565b610cdc602c6002611db4565b610ce69190611db4565b62ffffff1988169060416301000000610f9e565b9695505050505050565b600081610d1c62ffffff198216640101000000610e54565b506104aa62ffffff19841660006004610f6e565b600081610d4862ffffff198216640101000000610e54565b506104aa62ffffff19841660006008610f6e565b600081610d7462ffffff198216640101000000610e54565b5061051083610de2565b600081610d9662ffffff198216640101000000610e54565b50610da083610de2565b509392505050565b60008082610dc162ffffff198216640101000000610e54565b50610dcb84610de2565b909590945092505050565b6000610554828261086a565b60008080610df962ffffff198516602c6002610f6e565b60ff600882901c8116969116945092505050565b600080610e1a8385611db4565b9050604051811115610e2a575060005b80600003610e3f5762ffffff199150506104aa565b5050606092831b9190911790911b1760181b90565b6000610e6083836114fd565b610f67576000610e7e610e7285610990565b64ffffffffff1661151f565b9150506000610e938464ffffffffff1661151f565b6040517f5479706520617373657274696f6e206661696c65642e20476f7420307800000060208201527fffffffffffffffffffff0000000000000000000000000000000000000000000060b086811b8216603d8401527f2e20457870656374656420307800000000000000000000000000000000000000604784015283901b16605482015290925060009150605e016040516020818303038152906040529050806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109489190611b78565b5090919050565b6000610f7b826020611e0d565b610f86906008611e26565b60ff16610f9485858561134f565b901c949350505050565b600080610faa86611609565b6bffffffffffffffffffffffff169050610fc386611630565b84610fce8784611db4565b610fd89190611db4565b1115610feb5762ffffff19915050610510565b610ff58582611db4565b9050610cfa8364ffffffffff168286610e0d565b600062ffffff198084160361107a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f636f7079546f3a204e756c6c20706f696e7465722064657265660000000000006044820152606401610948565b61108383611669565b6110e9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f636f7079546f3a20496e76616c696420706f696e7465722064657265660000006044820152606401610948565b60006111038460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061111e85611609565b6bffffffffffffffffffffffff1690506000806040519150858211156111445760206060fd5b8386858560045afa9050806111b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f6964656e746974793a206f7574206f66206761730000000000000000000000006044820152606401610948565b6111e56111c188610990565b70ffffffffff000000000000000000000000606091821b168817901b851760181b90565b979650505050505050565b6000601882901c6bffffffffffffffffffffffff1681611211604183611e42565b90508161121f604183611dc7565b14611286576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601160248201527f217369676e6174757265734c656e6774680000000000000000000000000000006044820152606401610948565b60ff81106104aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f546f6f206d616e79207369676e617475726573000000000000000000000000006044820152606401610948565b604051606090600061130584602084016116a5565b905060006113218260181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff169050600061133c8361173f565b9184525082016020016040525092915050565b60008160ff16600003611364575060006104aa565b61137c8460181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff1661139760ff841685611db4565b111561141a576113e76113a985611609565b6bffffffffffffffffffffffff166113cf8660181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff16858560ff16611755565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109489190611b78565b60208260ff161115611488576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f496e6465783a206d6f7265207468616e203332206279746573000000000000006044820152606401610948565b60088202600061149786611609565b6bffffffffffffffffffffffff16905060007f80000000000000000000000000000000000000000000000000000000000000007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84011d91909501511695945050505050565b60008164ffffffffff1661151084610990565b64ffffffffff16149392505050565b600080601f5b600f8160ff16111561159257600061153e826008611e26565b60ff1685901c905061154f816118e5565b61ffff16841793508160ff1660101461156a57601084901b93505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611525565b50600f5b60ff8160ff1610156116035760006115af826008611e26565b60ff1685901c90506115c0816118e5565b61ffff16831792508160ff166000146115db57601083901b92505b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff01611596565b50915091565b60008061161860606018611db4565b9290921c6bffffffffffffffffffffffff1692915050565b600061164a8260181c6bffffffffffffffffffffffff1690565b61165383611609565b016bffffffffffffffffffffffff169050919050565b600061167482610990565b64ffffffffff1664ffffffffff0361168e57506000919050565b600061169983611630565b60405110199392505050565b6000604051828111156116b85760206060fd5b506000805b845181101561172f5760008582815181106116da576116da611dde565b602002602001015190506116f081848701611009565b506117098160181c6bffffffffffffffffffffffff1690565b6bffffffffffffffffffffffff168301925050808061172790611e7d565b9150506116bd565b50606083901b811760181b610510565b600061174a82611917565b610554906020611dc7565b606060006117628661151f565b91505060006117708661151f565b915050600061177e8661151f565b915050600061178c8661151f565b604080517f54797065644d656d566965772f696e646578202d204f76657272616e2074686560208201527f20766965772e20536c6963652069732061742030780000000000000000000000818301527fffffffffffff000000000000000000000000000000000000000000000000000060d098891b811660558301527f2077697468206c656e6774682030780000000000000000000000000000000000605b830181905297891b8116606a8301527f2e20417474656d7074656420746f20696e646578206174206f6666736574203060708301527f7800000000000000000000000000000000000000000000000000000000000000609083015295881b861660918201526097810196909652951b90921660a684015250507f2e0000000000000000000000000000000000000000000000000000000000000060ac8201528151808203608d01815260ad90910190915295945050505050565b60006118f760048360ff16901c611956565b60ff1661ffff919091161760081b61190e82611956565b60ff1617919050565b600060206119338360181c6bffffffffffffffffffffffff1690565b61194c906bffffffffffffffffffffffff16601f611db4565b6105549190611e42565b6040805180820190915260108082527f30313233343536373839616263646566000000000000000000000000000000006020830152600091600f841691829081106119a3576119a3611dde565b016020015160f81c9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126119f257600080fd5b813567ffffffffffffffff80821115611a0d57611a0d6119b2565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611a5357611a536119b2565b81604052838152866020858801011115611a6c57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215611aa157600080fd5b833567ffffffffffffffff80821115611ab957600080fd5b611ac5878388016119e1565b94506020860135915080821115611adb57600080fd5b611ae7878388016119e1565b93506040860135915080821115611afd57600080fd5b50611b0a868287016119e1565b9150509250925092565b6000815180845260005b81811015611b3a57602081850181015186830182015201611b1e565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b6020815260006104aa6020830184611b14565b803563ffffffff81168114611b9f57600080fd5b919050565b60008060008060808587031215611bba57600080fd5b611bc385611b8b565b9350611bd160208601611b8b565b9250611bdf60408601611b8b565b9396929550929360600135925050565b600060208284031215611c0157600080fd5b81356bffffffffffffffffffffffff811681146104aa57600080fd5b600060208284031215611c2f57600080fd5b813567ffffffffffffffff811115611c4657600080fd5b610510848285016119e1565b60008060408385031215611c6557600080fd5b823564ffffffffff81168114611c7a57600080fd5b9150602083013567ffffffffffffffff811115611c9657600080fd5b611ca2858286016119e1565b9150509250929050565b600060208284031215611cbe57600080fd5b5035919050565b64ffffffffff831681526040602082015260006104a76040830184611b14565b60008060408385031215611cf857600080fd5b611d0183611b8b565b9150611d0f60208401611b8b565b90509250929050565b600060208284031215611d2a57600080fd5b813567ffffffffffffffff811681146104aa57600080fd5b600080600060608486031215611d5757600080fd5b611d6084611b8b565b9250611d6e60208501611b8b565b9150611d7c60408501611b8b565b90509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b8082018082111561055457610554611d85565b808202811582820484141761055457610554611d85565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60ff828116828216039081111561055457610554611d85565b60ff818116838216029081169081811461098957610989611d85565b600082611e78577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611eae57611eae611d85565b506001019056fea26469706673582212204485c0e451e93a24acd24b92b89a1caca671b441f66f5f18f590f74141ba331364736f6c63430008110033",
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

// FormatAttestationFromViews is a free data retrieval call binding the contract method 0x794dd214.
//
// Solidity: function formatAttestationFromViews(bytes _data, bytes _guardSignatures, bytes _notarySignatures) view returns(bytes)
func (_AttestationHarness *AttestationHarnessCaller) FormatAttestationFromViews(opts *bind.CallOpts, _data []byte, _guardSignatures []byte, _notarySignatures []byte) ([]byte, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "formatAttestationFromViews", _data, _guardSignatures, _notarySignatures)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// FormatAttestationFromViews is a free data retrieval call binding the contract method 0x794dd214.
//
// Solidity: function formatAttestationFromViews(bytes _data, bytes _guardSignatures, bytes _notarySignatures) view returns(bytes)
func (_AttestationHarness *AttestationHarnessSession) FormatAttestationFromViews(_data []byte, _guardSignatures []byte, _notarySignatures []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestationFromViews(&_AttestationHarness.CallOpts, _data, _guardSignatures, _notarySignatures)
}

// FormatAttestationFromViews is a free data retrieval call binding the contract method 0x794dd214.
//
// Solidity: function formatAttestationFromViews(bytes _data, bytes _guardSignatures, bytes _notarySignatures) view returns(bytes)
func (_AttestationHarness *AttestationHarnessCallerSession) FormatAttestationFromViews(_data []byte, _guardSignatures []byte, _notarySignatures []byte) ([]byte, error) {
	return _AttestationHarness.Contract.FormatAttestationFromViews(&_AttestationHarness.CallOpts, _data, _guardSignatures, _notarySignatures)
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

// UnpackDomains is a free data retrieval call binding the contract method 0x7214101b.
//
// Solidity: function unpackDomains(uint64 _attestationDomains) pure returns(uint32, uint32)
func (_AttestationHarness *AttestationHarnessCaller) UnpackDomains(opts *bind.CallOpts, _attestationDomains uint64) (uint32, uint32, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "unpackDomains", _attestationDomains)

	if err != nil {
		return *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// UnpackDomains is a free data retrieval call binding the contract method 0x7214101b.
//
// Solidity: function unpackDomains(uint64 _attestationDomains) pure returns(uint32, uint32)
func (_AttestationHarness *AttestationHarnessSession) UnpackDomains(_attestationDomains uint64) (uint32, uint32, error) {
	return _AttestationHarness.Contract.UnpackDomains(&_AttestationHarness.CallOpts, _attestationDomains)
}

// UnpackDomains is a free data retrieval call binding the contract method 0x7214101b.
//
// Solidity: function unpackDomains(uint64 _attestationDomains) pure returns(uint32, uint32)
func (_AttestationHarness *AttestationHarnessCallerSession) UnpackDomains(_attestationDomains uint64) (uint32, uint32, error) {
	return _AttestationHarness.Contract.UnpackDomains(&_AttestationHarness.CallOpts, _attestationDomains)
}

// UnpackKey is a free data retrieval call binding the contract method 0x308514be.
//
// Solidity: function unpackKey(uint96 _attestationKey) pure returns(uint64 domains, uint32 nonce)
func (_AttestationHarness *AttestationHarnessCaller) UnpackKey(opts *bind.CallOpts, _attestationKey *big.Int) (struct {
	Domains uint64
	Nonce   uint32
}, error) {
	var out []interface{}
	err := _AttestationHarness.contract.Call(opts, &out, "unpackKey", _attestationKey)

	outstruct := new(struct {
		Domains uint64
		Nonce   uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Domains = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.Nonce = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// UnpackKey is a free data retrieval call binding the contract method 0x308514be.
//
// Solidity: function unpackKey(uint96 _attestationKey) pure returns(uint64 domains, uint32 nonce)
func (_AttestationHarness *AttestationHarnessSession) UnpackKey(_attestationKey *big.Int) (struct {
	Domains uint64
	Nonce   uint32
}, error) {
	return _AttestationHarness.Contract.UnpackKey(&_AttestationHarness.CallOpts, _attestationKey)
}

// UnpackKey is a free data retrieval call binding the contract method 0x308514be.
//
// Solidity: function unpackKey(uint96 _attestationKey) pure returns(uint64 domains, uint32 nonce)
func (_AttestationHarness *AttestationHarnessCallerSession) UnpackKey(_attestationKey *big.Int) (struct {
	Domains uint64
	Nonce   uint32
}, error) {
	return _AttestationHarness.Contract.UnpackKey(&_AttestationHarness.CallOpts, _attestationKey)
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b208ccdc5e9d75f1a63da44c7a83a614f682c82490eaba5ec38855f8900a280564736f6c63430008110033",
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
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204aee4a6b37c6af87f9819463d669c53693a74aff75cd5b48d242170a851fd03c64736f6c63430008110033",
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
	Bin: "0x6101f061003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100ad5760003560e01c806397b8ad4a11610080578063eb74062811610065578063eb740628146100f8578063f26be3fc14610100578063fb734584146100f857600080fd5b806397b8ad4a146100cd578063b602d173146100e557600080fd5b806310153fce146100b25780631136e7ea146100cd57806313090c5a146100d55780631bfe17ce146100dd575b600080fd5b6100ba602881565b6040519081526020015b60405180910390f35b6100ba601881565b6100ba610158565b6100ba610172565b6100ba6bffffffffffffffffffffffff81565b6100ba606081565b6101277fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000081565b6040517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00000090911681526020016100c4565b606061016581601861017a565b61016f919061017a565b81565b61016f606060185b808201808211156101b4577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b9291505056fea2646970667358221220a581a65672c6517980835fe9a38549634cc9d743d493792bfac43c59aa7ad80264736f6c63430008110033",
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
