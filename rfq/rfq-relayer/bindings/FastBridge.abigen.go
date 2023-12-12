// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// IFastBridgeBridgeParams is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeBridgeParams struct {
	DstChainId   uint32
	To           common.Address
	OriginToken  common.Address
	DestToken    common.Address
	OriginAmount *big.Int
	DestAmount   *big.Int
	Deadline     *big.Int
}

// IFastBridgeBridgeTransaction is an auto generated low-level Go binding around an user-defined struct.
type IFastBridgeBridgeTransaction struct {
	OriginChainId uint32
	DestChainId   uint32
	OriginSender  common.Address
	DestRecipient common.Address
	OriginToken   common.Address
	DestToken     common.Address
	OriginAmount  *big.Int
	DestAmount    *big.Int
	Deadline      *big.Int
	Nonce         *big.Int
}

// FastBridgeMetaData contains all meta data concerning the FastBridge contract.
var FastBridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ChainIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineNotExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeadlineTooShort\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodNotPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisputePeriodPassed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MsgValueIncorrect\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StatusIncorrect\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenNotETH\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TransactionRelayed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeDepositRefunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"BridgeProofDisputed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionHash\",\"type\":\"bytes32\"}],\"name\":\"BridgeProofProvided\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BridgeRelayed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"BridgeRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"guard\",\"type\":\"address\"}],\"name\":\"GuardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DISPUTE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GUARD_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_DEADLINE_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"addGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"addRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"dstChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeProofs\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"timestamp\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeRelays\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"bridgeStatuses\",\"outputs\":[{\"internalType\":\"enumFastBridge.BridgeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"transactionId\",\"type\":\"bytes32\"}],\"name\":\"dispute\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"getBridgeTransaction\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"originChainId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"destChainId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"originSender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destRecipient\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"originAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"internalType\":\"structIFastBridge.BridgeTransaction\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"destTxHash\",\"type\":\"bytes32\"}],\"name\":\"prove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"request\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_guard\",\"type\":\"address\"}],\"name\":\"removeGuard\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"name\":\"removeRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002203380380620022038339810160408190526200003491620001a8565b80806001600160a01b0381166200006557604051631e4fbdf760e01b81526000600482015260240160405180910390fd5b620000708162000087565b506200007e600082620000a5565b505050620001da565b600180546001600160a01b0319169055620000a28162000158565b50565b60008281526002602090815260408083206001600160a01b038516845290915281205460ff166200014e5760008381526002602090815260408083206001600160a01b03861684529091529020805460ff19166001179055620001053390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a450600162000152565b5060005b92915050565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600060208284031215620001bb57600080fd5b81516001600160a01b0381168114620001d357600080fd5b9392505050565b61201980620001ea6000396000f3fe6080604052600436106101d85760003560e01c80638da5cb5b11610102578063add98c7011610095578063dd39f00d11610064578063dd39f00d146105a2578063e30c3978146105c2578063ef059bd3146105e0578063f2fde38b1461060057600080fd5b8063add98c701461052c578063affed0e01461054c578063b623501614610562578063d547741f1461058257600080fd5b8063926d7d7f116100d1578063926d7d7f146104c8578063a217fddf146104ea578063a5bbe22b1461038e578063ac11fb1a146104ff57600080fd5b80638da5cb5b146103f45780638f0d6f171461042657806391ad50391461043957806391d14854146104a857600080fd5b806341fcb6121161017a57806379ba50971161014957806379ba509714610379578063820688d51461038e5780638379a24f146103a4578063886d36ff146103d457600080fd5b806341fcb6121461030457806360f0a5ac146103245780636913a63c14610344578063715018a61461036457600080fd5b80630728e3ad116101b65780630728e3ad1461027f578063248a9ca3146102945780632f2ff15d146102c457806336568abe146102e457600080fd5b806301ffc9a7146101dd57806303ed0ee514610212578063051287bc14610242575b600080fd5b3480156101e957600080fd5b506101fd6101f8366004611992565b610620565b60405190151581526020015b60405180910390f35b34801561021e57600080fd5b50610234600080516020611fa483398151915281565b604051908152602001610209565b34801561024e57600080fd5b5061027261025d3660046119bc565b60036020526000908152604090205460ff1681565b60405161020991906119eb565b61029261028d366004611a7a565b610657565b005b3480156102a057600080fd5b506102346102af3660046119bc565b60009081526002602052604090206001015490565b3480156102d057600080fd5b506102926102df366004611b22565b610865565b3480156102f057600080fd5b506102926102ff366004611b22565b610890565b34801561031057600080fd5b5061029261031f366004611bdf565b6108c8565b34801561033057600080fd5b5061029261033f366004611c26565b610a80565b34801561035057600080fd5b5061029261035f366004611c26565b610aed565b34801561037057600080fd5b50610292610b53565b34801561038557600080fd5b50610292610b67565b34801561039a57600080fd5b5061023461070881565b3480156103b057600080fd5b506101fd6103bf3660046119bc565b60056020526000908152604090205460ff1681565b3480156103e057600080fd5b506102926103ef366004611c43565b610bab565b34801561040057600080fd5b506000546001600160a01b03165b6040516001600160a01b039091168152602001610209565b610292610434366004611c88565b610cf7565b34801561044557600080fd5b506104816104543660046119bc565b6004602052600090815260409020546001600160601b03811690600160601b90046001600160a01b031682565b604080516001600160601b0390931683526001600160a01b03909116602083015201610209565b3480156104b457600080fd5b506101fd6104c3366004611b22565b610e37565b3480156104d457600080fd5b50610234600080516020611fc483398151915281565b3480156104f657600080fd5b50610234600081565b34801561050b57600080fd5b5061051f61051a366004611c88565b610e62565b6040516102099190611cc5565b34801561053857600080fd5b506102926105473660046119bc565b610ec5565b34801561055857600080fd5b5061023460065481565b34801561056e57600080fd5b5061029261057d366004611c26565b611028565b34801561058e57600080fd5b5061029261059d366004611b22565b61108e565b3480156105ae57600080fd5b506102926105bd366004611c26565b6110b3565b3480156105ce57600080fd5b506001546001600160a01b031661040e565b3480156105ec57600080fd5b506102926105fb366004611bdf565b611119565b34801561060c57600080fd5b5061029261061b366004611c26565b61124b565b60006001600160e01b03198216637965db0b60e01b148061065157506301ffc9a760e01b6001600160e01b03198316145b92915050565b46816000015163ffffffff160361068157604051637029fdf960e01b815260040160405180910390fd5b60808101511580610694575060a0810151155b156106b257604051631c71041960e31b815260040160405180910390fd5b60408101516001600160a01b031615806106d7575060608101516001600160a01b0316155b156106f55760405163d92e233d60e01b815260040160405180910390fd5b61070161070842611da0565b8160c001511015610724576040516296ff9960e31b815260040160405180910390fd5b600061073930836040015184608001516112bc565b905060006040518061014001604052804663ffffffff168152602001846000015163ffffffff168152602001336001600160a01b0316815260200184602001516001600160a01b0316815260200184604001516001600160a01b0316815260200184606001516001600160a01b031681526020018381526020018460a0015181526020018460c001518152602001600660008154809291906107da90611db3565b9091555090526040516107f09190602001611cc5565b60408051808303601f19018152828252805160208083019190912060008181526003909252929020805460ff191660011790559250907f2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf9061085790839033908690611df0565b60405180910390a150505050565b60008281526002602052604090206001015461088081611462565b61088a838361146c565b50505050565b6001600160a01b03811633146108b95760405163334bd91960e11b815260040160405180910390fd5b6108c38282611500565b505050565b6108e0600080516020611fc483398151915233610e37565b6109055760405162461bcd60e51b81526004016108fc90611e39565b60405180910390fd5b81516020830120600061091784610e62565b9050600260008381526003602052604090205460ff16600481111561093e5761093e6119d5565b1461095c576040516320a2c0b960e11b815260040160405180910390fd5b6000828152600460209081526040918290208251808401909352546001600160601b0381168352600160601b90046001600160a01b031690820181905233146109b8576040516304af43a960e41b815260040160405180910390fd5b80516107089042036001600160601b0316116109e757604051631992d0bd60e01b815260040160405180910390fd5b6000838152600360208190526040909120805460ff19169091179055608082015160c0830151610a216001600160a01b038316878361156d565b604080518681523360208201526001600160a01b0388811682840152841660608201526080810183905290517f582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac6789181900360a00190a150505050505050565b610a8b600033610e37565b610a9457600080fd5b610aac600080516020611fc483398151915282611500565b506040516001600160a01b03821681527f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b906020015b60405180910390a150565b610af8600033610e37565b610b0157600080fd5b610b19600080516020611fa48339815191528261146c565b506040516001600160a01b03821681527f93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f5990602001610ae2565b610b5b611654565b610b656000611681565b565b60015433906001600160a01b03168114610b9f5760405163118cdaa760e01b81526001600160a01b03821660048201526024016108fc565b610ba881611681565b50565b610bc3600080516020611fc483398151915233610e37565b610bdf5760405162461bcd60e51b81526004016108fc90611e39565b815160208301206000610bf184610e62565b9050806101000151421115610c195760405163559895a360e01b815260040160405180910390fd5b600160008381526003602052604090205460ff166004811115610c3e57610c3e6119d5565b14610c5c576040516320a2c0b960e11b815260040160405180910390fd5b6000828152600360209081526040808320805460ff19166002179055805180820182526001600160601b034281168252338285018181528887526004865295849020925195516001600160a01b0316600160601b02959091169490941790558051858152918201929092529081018490527f4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e490606001610857565b610d0f600080516020611fc483398151915233610e37565b610d2b5760405162461bcd60e51b81526004016108fc90611e39565b805160208201206000610d3d83610e62565b90504663ffffffff16816020015163ffffffff1614610d6f57604051637029fdf960e01b815260040160405180910390fd5b60008281526005602052604090205460ff1615610d9f5760405163bef7bb7d60e01b815260040160405180910390fd5b6000828152600560205260409020805460ff19166001179055606081015160a082015160e0830151610dd28383836112bc565b50604080518681523360208201526001600160a01b03808616928201929092529083166060820152608081018290527f7212c977e94cbafc6e4685aee721574624efaa1993ce2b48d1bb37c27ffa13f49060a0015b60405180910390a1505050505050565b60009182526002602090815260408084206001600160a01b0393909316845291905290205460ff1690565b604080516101408101825260008082526020808301829052928201819052606082018190526080820181905260a0820181905260c0820181905260e082018190526101008201819052610120820152825190916106519184018101908401611e8b565b610edd600080516020611fa483398151915233610e37565b610f215760405162461bcd60e51b815260206004820152601560248201527410d85b1b195c881a5cc81b9bdd08184819dd585c99605a1b60448201526064016108fc565b600260008281526003602052604090205460ff166004811115610f4657610f466119d5565b14610f64576040516320a2c0b960e11b815260040160405180910390fd5b6000818152600460209081526040918290208251808401909352546001600160601b03808216808552600160601b9092046001600160a01b03169390920192909252610708914203161115610fcc57604051630fa422ab60e21b815260040160405180910390fd5b6000818152600360209081526040808320805460ff191660011790556004825280832092909255815183815233918101919091527f0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad49101610ae2565b611033600033610e37565b61103c57600080fd5b611054600080516020611fa483398151915282611500565b506040516001600160a01b03821681527f59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d04890602001610ae2565b6000828152600260205260409020600101546110a981611462565b61088a8383611500565b6110be600033610e37565b6110c757600080fd5b6110df600080516020611fc48339815191528261146c565b506040516001600160a01b03821681527f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c590602001610ae2565b81516020830120600061112b84610e62565b60408101519091506001600160a01b0316331461115b576040516304af43a960e41b815260040160405180910390fd5b8061010001514211611180576040516370affcf560e11b815260040160405180910390fd5b600160008381526003602052604090205460ff1660048111156111a5576111a56119d5565b146111c3576040516320a2c0b960e11b815260040160405180910390fd5b6000828152600360205260409020805460ff19166004179055608081015160c08201516111fa6001600160a01b038316868361156d565b604080518581526001600160a01b038088166020830152841691810191909152606081018290527fb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc95890608001610e27565b611253611654565b600180546001600160a01b0383166001600160a01b031990911681179091556112846000546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6000346000036113d7576112d8836001600160a01b031661169a565b6040516370a0823160e01b81526001600160a01b0385811660048301528416906370a0823190602401602060405180830381865afa15801561131e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113429190611f39565b90506113596001600160a01b038416338685611703565b6040516370a0823160e01b81526001600160a01b0385811660048301528291908516906370a0823190602401602060405180830381865afa1580156113a2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c69190611f39565b6113d09190611f52565b905061145b565b6001600160a01b03831673eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee1461141457604051632eac7efb60e01b815260040160405180910390fd5b348214611434576040516381de0bf360e01b815260040160405180910390fd5b6001600160a01b0384163014611458576114586001600160a01b038416858461156d565b50345b9392505050565b610ba8813361176a565b60006114788383610e37565b6114f85760008381526002602090815260408083206001600160a01b03861684529091529020805460ff191660011790556114b03390565b6001600160a01b0316826001600160a01b0316847f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4506001610651565b506000610651565b600061150c8383610e37565b156114f85760008381526002602090815260408083206001600160a01b0386168085529252808320805460ff1916905551339286917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4506001610651565b306001600160a01b0383160361158257505050565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b03841601611640576000826001600160a01b03168260405160006040518083038185875af1925050503d80600081146115f4576040519150601f19603f3d011682016040523d82523d6000602084013e6115f9565b606091505b505090508061088a5760405162461bcd60e51b8152602060048201526013602482015272115512081d1c985b9cd9995c8819985a5b1959606a1b60448201526064016108fc565b6108c36001600160a01b03841683836117a7565b6000546001600160a01b03163314610b655760405163118cdaa760e01b81523360048201526024016108fc565b600180546001600160a01b0319169055610ba8816117d8565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeed196001600160a01b038216016116d857604051630fea47fd60e31b815260040160405180910390fd5b806001600160a01b03163b600003610ba857604051630fea47fd60e31b815260040160405180910390fd5b6040516001600160a01b03848116602483015283811660448301526064820183905261088a9186918216906323b872dd906084015b604051602081830303815290604052915060e01b6020820180516001600160e01b038381831617835250505050611828565b6117748282610e37565b6117a35760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016108fc565b5050565b6040516001600160a01b038381166024830152604482018390526108c391859182169063a9059cbb90606401611738565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b600061183d6001600160a01b0384168361188b565b905080516000141580156118625750808060200190518101906118609190611f65565b155b156108c357604051635274afe760e01b81526001600160a01b03841660048201526024016108fc565b606061145b8383600084600080856001600160a01b031684866040516118b19190611f87565b60006040518083038185875af1925050503d80600081146118ee576040519150601f19603f3d011682016040523d82523d6000602084013e6118f3565b606091505b509150915061190386838361190d565b9695505050505050565b6060826119225761191d82611969565b61145b565b815115801561193957506001600160a01b0384163b155b1561196257604051639996b31560e01b81526001600160a01b03851660048201526024016108fc565b508061145b565b8051156119795780518082602001fd5b604051630a12f52160e11b815260040160405180910390fd5b6000602082840312156119a457600080fd5b81356001600160e01b03198116811461145b57600080fd5b6000602082840312156119ce57600080fd5b5035919050565b634e487b7160e01b600052602160045260246000fd5b6020810160058310611a0d57634e487b7160e01b600052602160045260246000fd5b91905290565b634e487b7160e01b600052604160045260246000fd5b604051610140810167ffffffffffffffff81118282101715611a4d57611a4d611a13565b60405290565b63ffffffff81168114610ba857600080fd5b6001600160a01b0381168114610ba857600080fd5b600060e08284031215611a8c57600080fd5b60405160e0810181811067ffffffffffffffff82111715611aaf57611aaf611a13565b6040528235611abd81611a53565b81526020830135611acd81611a65565b60208201526040830135611ae081611a65565b60408201526060830135611af381611a65565b806060830152506080830135608082015260a083013560a082015260c083013560c08201528091505092915050565b60008060408385031215611b3557600080fd5b823591506020830135611b4781611a65565b809150509250929050565b600082601f830112611b6357600080fd5b813567ffffffffffffffff80821115611b7e57611b7e611a13565b604051601f8301601f19908116603f01168101908282118183101715611ba657611ba6611a13565b81604052838152866020858801011115611bbf57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215611bf257600080fd5b823567ffffffffffffffff811115611c0957600080fd5b611c1585828601611b52565b9250506020830135611b4781611a65565b600060208284031215611c3857600080fd5b813561145b81611a65565b60008060408385031215611c5657600080fd5b823567ffffffffffffffff811115611c6d57600080fd5b611c7985828601611b52565b95602094909401359450505050565b600060208284031215611c9a57600080fd5b813567ffffffffffffffff811115611cb157600080fd5b611cbd84828501611b52565b949350505050565b815163ffffffff16815261014081016020830151611ceb602084018263ffffffff169052565b506040830151611d0660408401826001600160a01b03169052565b506060830151611d2160608401826001600160a01b03169052565b506080830151611d3c60808401826001600160a01b03169052565b5060a0830151611d5760a08401826001600160a01b03169052565b5060c083015160c083015260e083015160e083015261010080840151818401525061012080840151818401525092915050565b634e487b7160e01b600052601160045260246000fd5b8082018082111561065157610651611d8a565b600060018201611dc557611dc5611d8a565b5060010190565b60005b83811015611de7578181015183820152602001611dcf565b50506000910152565b83815260018060a01b03831660208201526060604082015260008251806060840152611e23816080850160208701611dcc565b601f01601f191691909101608001949350505050565b60208082526017908201527f43616c6c6572206973206e6f7420612072656c61796572000000000000000000604082015260600190565b8051611e7b81611a53565b919050565b8051611e7b81611a65565b60006101408284031215611e9e57600080fd5b611ea6611a29565b611eaf83611e70565b8152611ebd60208401611e70565b6020820152611ece60408401611e80565b6040820152611edf60608401611e80565b6060820152611ef060808401611e80565b6080820152611f0160a08401611e80565b60a082015260c0838101519082015260e080840151908201526101008084015190820152610120928301519281019290925250919050565b600060208284031215611f4b57600080fd5b5051919050565b8181038181111561065157610651611d8a565b600060208284031215611f7757600080fd5b8151801515811461145b57600080fd5b60008251611f99818460208701611dcc565b919091019291505056fe043c983c49d46f0e102151eaf8085d4a2e6571d5df2d47b013f39bddfd4a639de2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4a2646970667358221220d833f950f5c5a94f30f986e95ad4e04bbd71de404b78dca3e783ec76dd968d0a64736f6c63430008140033",
}

// FastBridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use FastBridgeMetaData.ABI instead.
var FastBridgeABI = FastBridgeMetaData.ABI

// FastBridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FastBridgeMetaData.Bin instead.
var FastBridgeBin = FastBridgeMetaData.Bin

// DeployFastBridge deploys a new Ethereum contract, binding an instance of FastBridge to it.
func DeployFastBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address) (common.Address, *types.Transaction, *FastBridge, error) {
	parsed, err := FastBridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FastBridgeBin), backend, _owner)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FastBridge{FastBridgeCaller: FastBridgeCaller{contract: contract}, FastBridgeTransactor: FastBridgeTransactor{contract: contract}, FastBridgeFilterer: FastBridgeFilterer{contract: contract}}, nil
}

// FastBridge is an auto generated Go binding around an Ethereum contract.
type FastBridge struct {
	FastBridgeCaller     // Read-only binding to the contract
	FastBridgeTransactor // Write-only binding to the contract
	FastBridgeFilterer   // Log filterer for contract events
}

// FastBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type FastBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FastBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FastBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FastBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FastBridgeSession struct {
	Contract     *FastBridge       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FastBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FastBridgeCallerSession struct {
	Contract *FastBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FastBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FastBridgeTransactorSession struct {
	Contract     *FastBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FastBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type FastBridgeRaw struct {
	Contract *FastBridge // Generic contract binding to access the raw methods on
}

// FastBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FastBridgeCallerRaw struct {
	Contract *FastBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// FastBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FastBridgeTransactorRaw struct {
	Contract *FastBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFastBridge creates a new instance of FastBridge, bound to a specific deployed contract.
func NewFastBridge(address common.Address, backend bind.ContractBackend) (*FastBridge, error) {
	contract, err := bindFastBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FastBridge{FastBridgeCaller: FastBridgeCaller{contract: contract}, FastBridgeTransactor: FastBridgeTransactor{contract: contract}, FastBridgeFilterer: FastBridgeFilterer{contract: contract}}, nil
}

// NewFastBridgeCaller creates a new read-only instance of FastBridge, bound to a specific deployed contract.
func NewFastBridgeCaller(address common.Address, caller bind.ContractCaller) (*FastBridgeCaller, error) {
	contract, err := bindFastBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FastBridgeCaller{contract: contract}, nil
}

// NewFastBridgeTransactor creates a new write-only instance of FastBridge, bound to a specific deployed contract.
func NewFastBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*FastBridgeTransactor, error) {
	contract, err := bindFastBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FastBridgeTransactor{contract: contract}, nil
}

// NewFastBridgeFilterer creates a new log filterer instance of FastBridge, bound to a specific deployed contract.
func NewFastBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*FastBridgeFilterer, error) {
	contract, err := bindFastBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FastBridgeFilterer{contract: contract}, nil
}

// bindFastBridge binds a generic wrapper to an already deployed contract.
func bindFastBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FastBridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastBridge *FastBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastBridge.Contract.FastBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastBridge *FastBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridge.Contract.FastBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastBridge *FastBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastBridge.Contract.FastBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FastBridge *FastBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FastBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FastBridge *FastBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FastBridge *FastBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FastBridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FastBridge.Contract.DEFAULTADMINROLE(&_FastBridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _FastBridge.Contract.DEFAULTADMINROLE(&_FastBridge.CallOpts)
}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCaller) DISPUTEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "DISPUTE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeSession) DISPUTEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.DISPUTEPERIOD(&_FastBridge.CallOpts)
}

// DISPUTEPERIOD is a free data retrieval call binding the contract method 0xa5bbe22b.
//
// Solidity: function DISPUTE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) DISPUTEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.DISPUTEPERIOD(&_FastBridge.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCaller) GUARDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "GUARD_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeSession) GUARDROLE() ([32]byte, error) {
	return _FastBridge.Contract.GUARDROLE(&_FastBridge.CallOpts)
}

// GUARDROLE is a free data retrieval call binding the contract method 0x03ed0ee5.
//
// Solidity: function GUARD_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) GUARDROLE() ([32]byte, error) {
	return _FastBridge.Contract.GUARDROLE(&_FastBridge.CallOpts)
}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCaller) MINDEADLINEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "MIN_DEADLINE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeSession) MINDEADLINEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.MINDEADLINEPERIOD(&_FastBridge.CallOpts)
}

// MINDEADLINEPERIOD is a free data retrieval call binding the contract method 0x820688d5.
//
// Solidity: function MIN_DEADLINE_PERIOD() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) MINDEADLINEPERIOD() (*big.Int, error) {
	return _FastBridge.Contract.MINDEADLINEPERIOD(&_FastBridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeSession) RELAYERROLE() ([32]byte, error) {
	return _FastBridge.Contract.RELAYERROLE(&_FastBridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) RELAYERROLE() ([32]byte, error) {
	return _FastBridge.Contract.RELAYERROLE(&_FastBridge.CallOpts)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 ) view returns(uint96 timestamp, address relayer)
func (_FastBridge *FastBridgeCaller) BridgeProofs(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "bridgeProofs", arg0)

	outstruct := new(struct {
		Timestamp *big.Int
		Relayer   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Relayer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 ) view returns(uint96 timestamp, address relayer)
func (_FastBridge *FastBridgeSession) BridgeProofs(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _FastBridge.Contract.BridgeProofs(&_FastBridge.CallOpts, arg0)
}

// BridgeProofs is a free data retrieval call binding the contract method 0x91ad5039.
//
// Solidity: function bridgeProofs(bytes32 ) view returns(uint96 timestamp, address relayer)
func (_FastBridge *FastBridgeCallerSession) BridgeProofs(arg0 [32]byte) (struct {
	Timestamp *big.Int
	Relayer   common.Address
}, error) {
	return _FastBridge.Contract.BridgeProofs(&_FastBridge.CallOpts, arg0)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 ) view returns(bool)
func (_FastBridge *FastBridgeCaller) BridgeRelays(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "bridgeRelays", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 ) view returns(bool)
func (_FastBridge *FastBridgeSession) BridgeRelays(arg0 [32]byte) (bool, error) {
	return _FastBridge.Contract.BridgeRelays(&_FastBridge.CallOpts, arg0)
}

// BridgeRelays is a free data retrieval call binding the contract method 0x8379a24f.
//
// Solidity: function bridgeRelays(bytes32 ) view returns(bool)
func (_FastBridge *FastBridgeCallerSession) BridgeRelays(arg0 [32]byte) (bool, error) {
	return _FastBridge.Contract.BridgeRelays(&_FastBridge.CallOpts, arg0)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 ) view returns(uint8)
func (_FastBridge *FastBridgeCaller) BridgeStatuses(opts *bind.CallOpts, arg0 [32]byte) (uint8, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "bridgeStatuses", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 ) view returns(uint8)
func (_FastBridge *FastBridgeSession) BridgeStatuses(arg0 [32]byte) (uint8, error) {
	return _FastBridge.Contract.BridgeStatuses(&_FastBridge.CallOpts, arg0)
}

// BridgeStatuses is a free data retrieval call binding the contract method 0x051287bc.
//
// Solidity: function bridgeStatuses(bytes32 ) view returns(uint8)
func (_FastBridge *FastBridgeCallerSession) BridgeStatuses(arg0 [32]byte) (uint8, error) {
	return _FastBridge.Contract.BridgeStatuses(&_FastBridge.CallOpts, arg0)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256))
func (_FastBridge *FastBridgeCaller) GetBridgeTransaction(opts *bind.CallOpts, request []byte) (IFastBridgeBridgeTransaction, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "getBridgeTransaction", request)

	if err != nil {
		return *new(IFastBridgeBridgeTransaction), err
	}

	out0 := *abi.ConvertType(out[0], new(IFastBridgeBridgeTransaction)).(*IFastBridgeBridgeTransaction)

	return out0, err

}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256))
func (_FastBridge *FastBridgeSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _FastBridge.Contract.GetBridgeTransaction(&_FastBridge.CallOpts, request)
}

// GetBridgeTransaction is a free data retrieval call binding the contract method 0xac11fb1a.
//
// Solidity: function getBridgeTransaction(bytes request) pure returns((uint32,uint32,address,address,address,address,uint256,uint256,uint256,uint256))
func (_FastBridge *FastBridgeCallerSession) GetBridgeTransaction(request []byte) (IFastBridgeBridgeTransaction, error) {
	return _FastBridge.Contract.GetBridgeTransaction(&_FastBridge.CallOpts, request)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridge *FastBridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridge *FastBridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FastBridge.Contract.GetRoleAdmin(&_FastBridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_FastBridge *FastBridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _FastBridge.Contract.GetRoleAdmin(&_FastBridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridge *FastBridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridge *FastBridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FastBridge.Contract.HasRole(&_FastBridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_FastBridge *FastBridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _FastBridge.Contract.HasRole(&_FastBridge.CallOpts, role, account)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridge *FastBridgeCaller) Nonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "nonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridge *FastBridgeSession) Nonce() (*big.Int, error) {
	return _FastBridge.Contract.Nonce(&_FastBridge.CallOpts)
}

// Nonce is a free data retrieval call binding the contract method 0xaffed0e0.
//
// Solidity: function nonce() view returns(uint256)
func (_FastBridge *FastBridgeCallerSession) Nonce() (*big.Int, error) {
	return _FastBridge.Contract.Nonce(&_FastBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FastBridge *FastBridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FastBridge *FastBridgeSession) Owner() (common.Address, error) {
	return _FastBridge.Contract.Owner(&_FastBridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FastBridge *FastBridgeCallerSession) Owner() (common.Address, error) {
	return _FastBridge.Contract.Owner(&_FastBridge.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FastBridge *FastBridgeCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FastBridge *FastBridgeSession) PendingOwner() (common.Address, error) {
	return _FastBridge.Contract.PendingOwner(&_FastBridge.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FastBridge *FastBridgeCallerSession) PendingOwner() (common.Address, error) {
	return _FastBridge.Contract.PendingOwner(&_FastBridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridge *FastBridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FastBridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridge *FastBridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastBridge.Contract.SupportsInterface(&_FastBridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_FastBridge *FastBridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FastBridge.Contract.SupportsInterface(&_FastBridge.CallOpts, interfaceId)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FastBridge *FastBridgeTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FastBridge *FastBridgeSession) AcceptOwnership() (*types.Transaction, error) {
	return _FastBridge.Contract.AcceptOwnership(&_FastBridge.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FastBridge *FastBridgeTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FastBridge.Contract.AcceptOwnership(&_FastBridge.TransactOpts)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactor) AddGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "addGuard", _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_FastBridge *FastBridgeSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddGuard(&_FastBridge.TransactOpts, _guard)
}

// AddGuard is a paid mutator transaction binding the contract method 0x6913a63c.
//
// Solidity: function addGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactorSession) AddGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddGuard(&_FastBridge.TransactOpts, _guard)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactor) AddRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "addRelayer", _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddRelayer(&_FastBridge.TransactOpts, _relayer)
}

// AddRelayer is a paid mutator transaction binding the contract method 0xdd39f00d.
//
// Solidity: function addRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactorSession) AddRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.AddRelayer(&_FastBridge.TransactOpts, _relayer)
}

// Bridge is a paid mutator transaction binding the contract method 0x0728e3ad.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,uint256) params) payable returns()
func (_FastBridge *FastBridgeTransactor) Bridge(opts *bind.TransactOpts, params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "bridge", params)
}

// Bridge is a paid mutator transaction binding the contract method 0x0728e3ad.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,uint256) params) payable returns()
func (_FastBridge *FastBridgeSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridge.Contract.Bridge(&_FastBridge.TransactOpts, params)
}

// Bridge is a paid mutator transaction binding the contract method 0x0728e3ad.
//
// Solidity: function bridge((uint32,address,address,address,uint256,uint256,uint256) params) payable returns()
func (_FastBridge *FastBridgeTransactorSession) Bridge(params IFastBridgeBridgeParams) (*types.Transaction, error) {
	return _FastBridge.Contract.Bridge(&_FastBridge.TransactOpts, params)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactor) Claim(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "claim", request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridge *FastBridgeSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Claim(&_FastBridge.TransactOpts, request, to)
}

// Claim is a paid mutator transaction binding the contract method 0x41fcb612.
//
// Solidity: function claim(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactorSession) Claim(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Claim(&_FastBridge.TransactOpts, request, to)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridge *FastBridgeTransactor) Dispute(opts *bind.TransactOpts, transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "dispute", transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridge *FastBridgeSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Dispute(&_FastBridge.TransactOpts, transactionId)
}

// Dispute is a paid mutator transaction binding the contract method 0xadd98c70.
//
// Solidity: function dispute(bytes32 transactionId) returns()
func (_FastBridge *FastBridgeTransactorSession) Dispute(transactionId [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Dispute(&_FastBridge.TransactOpts, transactionId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.GrantRole(&_FastBridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.GrantRole(&_FastBridge.TransactOpts, role, account)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridge *FastBridgeTransactor) Prove(opts *bind.TransactOpts, request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "prove", request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridge *FastBridgeSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Prove(&_FastBridge.TransactOpts, request, destTxHash)
}

// Prove is a paid mutator transaction binding the contract method 0x886d36ff.
//
// Solidity: function prove(bytes request, bytes32 destTxHash) returns()
func (_FastBridge *FastBridgeTransactorSession) Prove(request []byte, destTxHash [32]byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Prove(&_FastBridge.TransactOpts, request, destTxHash)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactor) Refund(opts *bind.TransactOpts, request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "refund", request, to)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_FastBridge *FastBridgeSession) Refund(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Refund(&_FastBridge.TransactOpts, request, to)
}

// Refund is a paid mutator transaction binding the contract method 0xef059bd3.
//
// Solidity: function refund(bytes request, address to) returns()
func (_FastBridge *FastBridgeTransactorSession) Refund(request []byte, to common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.Refund(&_FastBridge.TransactOpts, request, to)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridge *FastBridgeTransactor) Relay(opts *bind.TransactOpts, request []byte) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "relay", request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridge *FastBridgeSession) Relay(request []byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Relay(&_FastBridge.TransactOpts, request)
}

// Relay is a paid mutator transaction binding the contract method 0x8f0d6f17.
//
// Solidity: function relay(bytes request) payable returns()
func (_FastBridge *FastBridgeTransactorSession) Relay(request []byte) (*types.Transaction, error) {
	return _FastBridge.Contract.Relay(&_FastBridge.TransactOpts, request)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactor) RemoveGuard(opts *bind.TransactOpts, _guard common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "removeGuard", _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_FastBridge *FastBridgeSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveGuard(&_FastBridge.TransactOpts, _guard)
}

// RemoveGuard is a paid mutator transaction binding the contract method 0xb6235016.
//
// Solidity: function removeGuard(address _guard) returns()
func (_FastBridge *FastBridgeTransactorSession) RemoveGuard(_guard common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveGuard(&_FastBridge.TransactOpts, _guard)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactor) RemoveRelayer(opts *bind.TransactOpts, _relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "removeRelayer", _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveRelayer(&_FastBridge.TransactOpts, _relayer)
}

// RemoveRelayer is a paid mutator transaction binding the contract method 0x60f0a5ac.
//
// Solidity: function removeRelayer(address _relayer) returns()
func (_FastBridge *FastBridgeTransactorSession) RemoveRelayer(_relayer common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RemoveRelayer(&_FastBridge.TransactOpts, _relayer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FastBridge *FastBridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FastBridge *FastBridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _FastBridge.Contract.RenounceOwnership(&_FastBridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FastBridge *FastBridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FastBridge.Contract.RenounceOwnership(&_FastBridge.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridge *FastBridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridge *FastBridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RenounceRole(&_FastBridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_FastBridge *FastBridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RenounceRole(&_FastBridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RevokeRole(&_FastBridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_FastBridge *FastBridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.RevokeRole(&_FastBridge.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FastBridge *FastBridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FastBridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FastBridge *FastBridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.TransferOwnership(&_FastBridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FastBridge *FastBridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FastBridge.Contract.TransferOwnership(&_FastBridge.TransactOpts, newOwner)
}

// FastBridgeBridgeDepositClaimedIterator is returned from FilterBridgeDepositClaimed and is used to iterate over the raw logs and unpacked data for BridgeDepositClaimed events raised by the FastBridge contract.
type FastBridgeBridgeDepositClaimedIterator struct {
	Event *FastBridgeBridgeDepositClaimed // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeDepositClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeDepositClaimed)
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
		it.Event = new(FastBridgeBridgeDepositClaimed)
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
func (it *FastBridgeBridgeDepositClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeDepositClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeDepositClaimed represents a BridgeDepositClaimed event raised by the FastBridge contract.
type FastBridgeBridgeDepositClaimed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositClaimed is a free log retrieval operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) FilterBridgeDepositClaimed(opts *bind.FilterOpts) (*FastBridgeBridgeDepositClaimedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeDepositClaimed")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeDepositClaimedIterator{contract: _FastBridge.contract, event: "BridgeDepositClaimed", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositClaimed is a free log subscription operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) WatchBridgeDepositClaimed(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeDepositClaimed) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeDepositClaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeDepositClaimed)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
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

// ParseBridgeDepositClaimed is a log parse operation binding the contract event 0x582211c35a2139ac3bbaac74663c6a1f56c6cbb658b41fe11fd45a82074ac678.
//
// Solidity: event BridgeDepositClaimed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) ParseBridgeDepositClaimed(log types.Log) (*FastBridgeBridgeDepositClaimed, error) {
	event := new(FastBridgeBridgeDepositClaimed)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeDepositRefundedIterator is returned from FilterBridgeDepositRefunded and is used to iterate over the raw logs and unpacked data for BridgeDepositRefunded events raised by the FastBridge contract.
type FastBridgeBridgeDepositRefundedIterator struct {
	Event *FastBridgeBridgeDepositRefunded // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeDepositRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeDepositRefunded)
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
		it.Event = new(FastBridgeBridgeDepositRefunded)
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
func (it *FastBridgeBridgeDepositRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeDepositRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeDepositRefunded represents a BridgeDepositRefunded event raised by the FastBridge contract.
type FastBridgeBridgeDepositRefunded struct {
	TransactionId [32]byte
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeDepositRefunded is a free log retrieval operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) FilterBridgeDepositRefunded(opts *bind.FilterOpts) (*FastBridgeBridgeDepositRefundedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeDepositRefunded")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeDepositRefundedIterator{contract: _FastBridge.contract, event: "BridgeDepositRefunded", logs: logs, sub: sub}, nil
}

// WatchBridgeDepositRefunded is a free log subscription operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) WatchBridgeDepositRefunded(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeDepositRefunded) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeDepositRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeDepositRefunded)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
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

// ParseBridgeDepositRefunded is a log parse operation binding the contract event 0xb4c55c0c9bc613519b920e88748090150b890a875d307f21bea7d4fb2e8bc958.
//
// Solidity: event BridgeDepositRefunded(bytes32 transactionId, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) ParseBridgeDepositRefunded(log types.Log) (*FastBridgeBridgeDepositRefunded, error) {
	event := new(FastBridgeBridgeDepositRefunded)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeDepositRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeProofDisputedIterator is returned from FilterBridgeProofDisputed and is used to iterate over the raw logs and unpacked data for BridgeProofDisputed events raised by the FastBridge contract.
type FastBridgeBridgeProofDisputedIterator struct {
	Event *FastBridgeBridgeProofDisputed // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeProofDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeProofDisputed)
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
		it.Event = new(FastBridgeBridgeProofDisputed)
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
func (it *FastBridgeBridgeProofDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeProofDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeProofDisputed represents a BridgeProofDisputed event raised by the FastBridge contract.
type FastBridgeBridgeProofDisputed struct {
	TransactionId [32]byte
	Relayer       common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofDisputed is a free log retrieval operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_FastBridge *FastBridgeFilterer) FilterBridgeProofDisputed(opts *bind.FilterOpts) (*FastBridgeBridgeProofDisputedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeProofDisputed")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeProofDisputedIterator{contract: _FastBridge.contract, event: "BridgeProofDisputed", logs: logs, sub: sub}, nil
}

// WatchBridgeProofDisputed is a free log subscription operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_FastBridge *FastBridgeFilterer) WatchBridgeProofDisputed(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeProofDisputed) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeProofDisputed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeProofDisputed)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
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

// ParseBridgeProofDisputed is a log parse operation binding the contract event 0x0695cf1d39b3055dcd0fe02d8b47eaf0d5a13e1996de925de59d0ef9b7f7fad4.
//
// Solidity: event BridgeProofDisputed(bytes32 transactionId, address relayer)
func (_FastBridge *FastBridgeFilterer) ParseBridgeProofDisputed(log types.Log) (*FastBridgeBridgeProofDisputed, error) {
	event := new(FastBridgeBridgeProofDisputed)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeProofDisputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeProofProvidedIterator is returned from FilterBridgeProofProvided and is used to iterate over the raw logs and unpacked data for BridgeProofProvided events raised by the FastBridge contract.
type FastBridgeBridgeProofProvidedIterator struct {
	Event *FastBridgeBridgeProofProvided // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeProofProvidedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeProofProvided)
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
		it.Event = new(FastBridgeBridgeProofProvided)
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
func (it *FastBridgeBridgeProofProvidedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeProofProvidedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeProofProvided represents a BridgeProofProvided event raised by the FastBridge contract.
type FastBridgeBridgeProofProvided struct {
	TransactionId   [32]byte
	Relayer         common.Address
	TransactionHash [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBridgeProofProvided is a free log retrieval operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_FastBridge *FastBridgeFilterer) FilterBridgeProofProvided(opts *bind.FilterOpts) (*FastBridgeBridgeProofProvidedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeProofProvided")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeProofProvidedIterator{contract: _FastBridge.contract, event: "BridgeProofProvided", logs: logs, sub: sub}, nil
}

// WatchBridgeProofProvided is a free log subscription operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_FastBridge *FastBridgeFilterer) WatchBridgeProofProvided(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeProofProvided) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeProofProvided")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeProofProvided)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
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

// ParseBridgeProofProvided is a log parse operation binding the contract event 0x4ac8af8a2cd87193d64dfc7a3b8d9923b714ec528b18725d080aa1299be0c5e4.
//
// Solidity: event BridgeProofProvided(bytes32 transactionId, address relayer, bytes32 transactionHash)
func (_FastBridge *FastBridgeFilterer) ParseBridgeProofProvided(log types.Log) (*FastBridgeBridgeProofProvided, error) {
	event := new(FastBridgeBridgeProofProvided)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeProofProvided", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeRelayedIterator is returned from FilterBridgeRelayed and is used to iterate over the raw logs and unpacked data for BridgeRelayed events raised by the FastBridge contract.
type FastBridgeBridgeRelayedIterator struct {
	Event *FastBridgeBridgeRelayed // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeRelayedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeRelayed)
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
		it.Event = new(FastBridgeBridgeRelayed)
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
func (it *FastBridgeBridgeRelayedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeRelayedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeRelayed represents a BridgeRelayed event raised by the FastBridge contract.
type FastBridgeBridgeRelayed struct {
	TransactionId [32]byte
	Relayer       common.Address
	To            common.Address
	Token         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRelayed is a free log retrieval operation binding the contract event 0x7212c977e94cbafc6e4685aee721574624efaa1993ce2b48d1bb37c27ffa13f4.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) FilterBridgeRelayed(opts *bind.FilterOpts) (*FastBridgeBridgeRelayedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeRelayed")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeRelayedIterator{contract: _FastBridge.contract, event: "BridgeRelayed", logs: logs, sub: sub}, nil
}

// WatchBridgeRelayed is a free log subscription operation binding the contract event 0x7212c977e94cbafc6e4685aee721574624efaa1993ce2b48d1bb37c27ffa13f4.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) WatchBridgeRelayed(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeRelayed) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeRelayed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeRelayed)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
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

// ParseBridgeRelayed is a log parse operation binding the contract event 0x7212c977e94cbafc6e4685aee721574624efaa1993ce2b48d1bb37c27ffa13f4.
//
// Solidity: event BridgeRelayed(bytes32 transactionId, address relayer, address to, address token, uint256 amount)
func (_FastBridge *FastBridgeFilterer) ParseBridgeRelayed(log types.Log) (*FastBridgeBridgeRelayed, error) {
	event := new(FastBridgeBridgeRelayed)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeRelayed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeBridgeRequestedIterator is returned from FilterBridgeRequested and is used to iterate over the raw logs and unpacked data for BridgeRequested events raised by the FastBridge contract.
type FastBridgeBridgeRequestedIterator struct {
	Event *FastBridgeBridgeRequested // Event containing the contract specifics and raw log

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
func (it *FastBridgeBridgeRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeBridgeRequested)
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
		it.Event = new(FastBridgeBridgeRequested)
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
func (it *FastBridgeBridgeRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeBridgeRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeBridgeRequested represents a BridgeRequested event raised by the FastBridge contract.
type FastBridgeBridgeRequested struct {
	TransactionId [32]byte
	Sender        common.Address
	Request       []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterBridgeRequested is a free log retrieval operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_FastBridge *FastBridgeFilterer) FilterBridgeRequested(opts *bind.FilterOpts) (*FastBridgeBridgeRequestedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "BridgeRequested")
	if err != nil {
		return nil, err
	}
	return &FastBridgeBridgeRequestedIterator{contract: _FastBridge.contract, event: "BridgeRequested", logs: logs, sub: sub}, nil
}

// WatchBridgeRequested is a free log subscription operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_FastBridge *FastBridgeFilterer) WatchBridgeRequested(opts *bind.WatchOpts, sink chan<- *FastBridgeBridgeRequested) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "BridgeRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeBridgeRequested)
				if err := _FastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
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

// ParseBridgeRequested is a log parse operation binding the contract event 0x2a8233b619c9d479346e133f609855c0a94d89fbcfa62f846a9f0cfdd1198ccf.
//
// Solidity: event BridgeRequested(bytes32 transactionId, address sender, bytes request)
func (_FastBridge *FastBridgeFilterer) ParseBridgeRequested(log types.Log) (*FastBridgeBridgeRequested, error) {
	event := new(FastBridgeBridgeRequested)
	if err := _FastBridge.contract.UnpackLog(event, "BridgeRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeGuardAddedIterator is returned from FilterGuardAdded and is used to iterate over the raw logs and unpacked data for GuardAdded events raised by the FastBridge contract.
type FastBridgeGuardAddedIterator struct {
	Event *FastBridgeGuardAdded // Event containing the contract specifics and raw log

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
func (it *FastBridgeGuardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeGuardAdded)
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
		it.Event = new(FastBridgeGuardAdded)
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
func (it *FastBridgeGuardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeGuardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeGuardAdded represents a GuardAdded event raised by the FastBridge contract.
type FastBridgeGuardAdded struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardAdded is a free log retrieval operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_FastBridge *FastBridgeFilterer) FilterGuardAdded(opts *bind.FilterOpts) (*FastBridgeGuardAddedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return &FastBridgeGuardAddedIterator{contract: _FastBridge.contract, event: "GuardAdded", logs: logs, sub: sub}, nil
}

// WatchGuardAdded is a free log subscription operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_FastBridge *FastBridgeFilterer) WatchGuardAdded(opts *bind.WatchOpts, sink chan<- *FastBridgeGuardAdded) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "GuardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeGuardAdded)
				if err := _FastBridge.contract.UnpackLog(event, "GuardAdded", log); err != nil {
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

// ParseGuardAdded is a log parse operation binding the contract event 0x93405f05cd04f0d1bd875f2de00f1f3890484ffd0589248953bdfd29ba7f2f59.
//
// Solidity: event GuardAdded(address guard)
func (_FastBridge *FastBridgeFilterer) ParseGuardAdded(log types.Log) (*FastBridgeGuardAdded, error) {
	event := new(FastBridgeGuardAdded)
	if err := _FastBridge.contract.UnpackLog(event, "GuardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeGuardRemovedIterator is returned from FilterGuardRemoved and is used to iterate over the raw logs and unpacked data for GuardRemoved events raised by the FastBridge contract.
type FastBridgeGuardRemovedIterator struct {
	Event *FastBridgeGuardRemoved // Event containing the contract specifics and raw log

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
func (it *FastBridgeGuardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeGuardRemoved)
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
		it.Event = new(FastBridgeGuardRemoved)
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
func (it *FastBridgeGuardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeGuardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeGuardRemoved represents a GuardRemoved event raised by the FastBridge contract.
type FastBridgeGuardRemoved struct {
	Guard common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterGuardRemoved is a free log retrieval operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_FastBridge *FastBridgeFilterer) FilterGuardRemoved(opts *bind.FilterOpts) (*FastBridgeGuardRemovedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return &FastBridgeGuardRemovedIterator{contract: _FastBridge.contract, event: "GuardRemoved", logs: logs, sub: sub}, nil
}

// WatchGuardRemoved is a free log subscription operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_FastBridge *FastBridgeFilterer) WatchGuardRemoved(opts *bind.WatchOpts, sink chan<- *FastBridgeGuardRemoved) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "GuardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeGuardRemoved)
				if err := _FastBridge.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
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

// ParseGuardRemoved is a log parse operation binding the contract event 0x59926e0a78d12238b668b31c8e3f6ece235a59a00ede111d883e255b68c4d048.
//
// Solidity: event GuardRemoved(address guard)
func (_FastBridge *FastBridgeFilterer) ParseGuardRemoved(log types.Log) (*FastBridgeGuardRemoved, error) {
	event := new(FastBridgeGuardRemoved)
	if err := _FastBridge.contract.UnpackLog(event, "GuardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FastBridge contract.
type FastBridgeOwnershipTransferStartedIterator struct {
	Event *FastBridgeOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FastBridgeOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeOwnershipTransferStarted)
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
		it.Event = new(FastBridgeOwnershipTransferStarted)
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
func (it *FastBridgeOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FastBridge contract.
type FastBridgeOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FastBridge *FastBridgeFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FastBridgeOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeOwnershipTransferStartedIterator{contract: _FastBridge.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FastBridge *FastBridgeFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FastBridgeOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeOwnershipTransferStarted)
				if err := _FastBridge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_FastBridge *FastBridgeFilterer) ParseOwnershipTransferStarted(log types.Log) (*FastBridgeOwnershipTransferStarted, error) {
	event := new(FastBridgeOwnershipTransferStarted)
	if err := _FastBridge.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FastBridge contract.
type FastBridgeOwnershipTransferredIterator struct {
	Event *FastBridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FastBridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeOwnershipTransferred)
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
		it.Event = new(FastBridgeOwnershipTransferred)
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
func (it *FastBridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeOwnershipTransferred represents a OwnershipTransferred event raised by the FastBridge contract.
type FastBridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FastBridge *FastBridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FastBridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeOwnershipTransferredIterator{contract: _FastBridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FastBridge *FastBridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FastBridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeOwnershipTransferred)
				if err := _FastBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FastBridge *FastBridgeFilterer) ParseOwnershipTransferred(log types.Log) (*FastBridgeOwnershipTransferred, error) {
	event := new(FastBridgeOwnershipTransferred)
	if err := _FastBridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the FastBridge contract.
type FastBridgeRelayerAddedIterator struct {
	Event *FastBridgeRelayerAdded // Event containing the contract specifics and raw log

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
func (it *FastBridgeRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRelayerAdded)
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
		it.Event = new(FastBridgeRelayerAdded)
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
func (it *FastBridgeRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRelayerAdded represents a RelayerAdded event raised by the FastBridge contract.
type FastBridgeRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_FastBridge *FastBridgeFilterer) FilterRelayerAdded(opts *bind.FilterOpts) (*FastBridgeRelayerAddedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return &FastBridgeRelayerAddedIterator{contract: _FastBridge.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_FastBridge *FastBridgeFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *FastBridgeRelayerAdded) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RelayerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRelayerAdded)
				if err := _FastBridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
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

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address relayer)
func (_FastBridge *FastBridgeFilterer) ParseRelayerAdded(log types.Log) (*FastBridgeRelayerAdded, error) {
	event := new(FastBridgeRelayerAdded)
	if err := _FastBridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the FastBridge contract.
type FastBridgeRelayerRemovedIterator struct {
	Event *FastBridgeRelayerRemoved // Event containing the contract specifics and raw log

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
func (it *FastBridgeRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRelayerRemoved)
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
		it.Event = new(FastBridgeRelayerRemoved)
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
func (it *FastBridgeRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRelayerRemoved represents a RelayerRemoved event raised by the FastBridge contract.
type FastBridgeRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_FastBridge *FastBridgeFilterer) FilterRelayerRemoved(opts *bind.FilterOpts) (*FastBridgeRelayerRemovedIterator, error) {

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return &FastBridgeRelayerRemovedIterator{contract: _FastBridge.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_FastBridge *FastBridgeFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *FastBridgeRelayerRemoved) (event.Subscription, error) {

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RelayerRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRelayerRemoved)
				if err := _FastBridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
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

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address relayer)
func (_FastBridge *FastBridgeFilterer) ParseRelayerRemoved(log types.Log) (*FastBridgeRelayerRemoved, error) {
	event := new(FastBridgeRelayerRemoved)
	if err := _FastBridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the FastBridge contract.
type FastBridgeRoleAdminChangedIterator struct {
	Event *FastBridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FastBridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRoleAdminChanged)
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
		it.Event = new(FastBridgeRoleAdminChanged)
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
func (it *FastBridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRoleAdminChanged represents a RoleAdminChanged event raised by the FastBridge contract.
type FastBridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FastBridge *FastBridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FastBridgeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeRoleAdminChangedIterator{contract: _FastBridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FastBridge *FastBridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FastBridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRoleAdminChanged)
				if err := _FastBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_FastBridge *FastBridgeFilterer) ParseRoleAdminChanged(log types.Log) (*FastBridgeRoleAdminChanged, error) {
	event := new(FastBridgeRoleAdminChanged)
	if err := _FastBridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the FastBridge contract.
type FastBridgeRoleGrantedIterator struct {
	Event *FastBridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *FastBridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRoleGranted)
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
		it.Event = new(FastBridgeRoleGranted)
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
func (it *FastBridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRoleGranted represents a RoleGranted event raised by the FastBridge contract.
type FastBridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FastBridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeRoleGrantedIterator{contract: _FastBridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FastBridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRoleGranted)
				if err := _FastBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) ParseRoleGranted(log types.Log) (*FastBridgeRoleGranted, error) {
	event := new(FastBridgeRoleGranted)
	if err := _FastBridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FastBridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the FastBridge contract.
type FastBridgeRoleRevokedIterator struct {
	Event *FastBridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FastBridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FastBridgeRoleRevoked)
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
		it.Event = new(FastBridgeRoleRevoked)
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
func (it *FastBridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FastBridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FastBridgeRoleRevoked represents a RoleRevoked event raised by the FastBridge contract.
type FastBridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FastBridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastBridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FastBridgeRoleRevokedIterator{contract: _FastBridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FastBridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _FastBridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FastBridgeRoleRevoked)
				if err := _FastBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_FastBridge *FastBridgeFilterer) ParseRoleRevoked(log types.Log) (*FastBridgeRoleRevoked, error) {
	event := new(FastBridgeRoleRevoked)
	if err := _FastBridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
