package bridge

import (
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

// GetRaw gets the raw event logs from the deposit event.
func (s SynapseBridgeTokenDeposit) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the deposit log.
func (s SynapseBridgeTokenDeposit) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDeposit) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the destination chain id from the deposit log.
func (s SynapseBridgeTokenDeposit) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the tokend eposit event.
func (s SynapseBridgeTokenDeposit) GetEventType() types.EventType {
	return types.DepositEvent
}

// GetIdentifier gets the unique identifier (txhash) for the deposit.
func (s SynapseBridgeTokenDeposit) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenDeposit) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenDeposit) GetContractAddress() string {
	return s.Raw.Address.String()
}

// GetRecipient gets the recipient of a deposit.
func (s SynapseBridgeTokenDeposit) GetRecipient() string {
	return s.To.String()
}

var _ types.CrossChainDepositLog = &SynapseBridgeTokenDeposit{}

// GetRaw gets the raw event logs from the deposit and swap event.
func (s SynapseBridgeTokenDepositAndSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the destination chain id from the deposit and swap log.
func (s SynapseBridgeTokenDepositAndSwap) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the deposit and swap event.
func (s SynapseBridgeTokenDepositAndSwap) GetEventType() types.EventType {
	return types.DepositAndSwapEvent
}

// GetIdentifier gets the unique identifier (txhash) for the deposit.
func (s SynapseBridgeTokenDepositAndSwap) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenDepositAndSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenDepositAndSwap) GetContractAddress() string {
	return s.Raw.Address.String()
}

var _ types.CrossChainUserEventLog = &SynapseBridgeTokenDepositAndSwap{}
