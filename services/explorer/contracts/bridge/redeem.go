package bridge

import (
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and swap event.
func (s SynapseBridgeTokenRedeemAndSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem and swap log.
func (s SynapseBridgeTokenRedeemAndSwap) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the redeem and swap log.
func (s SynapseBridgeTokenRedeemAndSwap) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the destination chain id from the redeem and swap log.
func (s SynapseBridgeTokenRedeemAndSwap) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the redeem and swap event.
func (s SynapseBridgeTokenRedeemAndSwap) GetEventType() types.EventType {
	return types.RedeemAndSwapEvent
}

// GetIdentifier gets the unique identifier (txhash) for the redeem and swap.
func (s SynapseBridgeTokenRedeemAndSwap) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeemAndSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeemAndSwap) GetContractAddress() string {
	return s.Raw.Address.String()
}

var _ types.CrossChainUserEventLog = &SynapseBridgeTokenRedeemAndSwap{}

// GetRaw gets the raw event logs from the redeem event.
func (s SynapseBridgeTokenRedeem) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeem) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeem) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenRedeem) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenRedeem) GetEventType() types.EventType {
	return types.RedeemEvent
}

// GetIdentifier gets the unique identifier (txhash) for the redeem.
func (s SynapseBridgeTokenRedeem) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeem) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeem) GetContractAddress() string {
	return s.Raw.Address.String()
}

var _ types.CrossChainUserEventLog = &SynapseBridgeTokenRedeem{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SynapseBridgeTokenRedeemAndRemove) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem and remove log.
func (s SynapseBridgeTokenRedeemAndRemove) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the redeem and remove log.
func (s SynapseBridgeTokenRedeemAndRemove) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenRedeemAndRemove) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenRedeemAndRemove) GetEventType() types.EventType {
	return types.RedeemAndRemoveEvent
}

// GetIdentifier gets the unique identifier (txhash) for the redeem.
func (s SynapseBridgeTokenRedeemAndRemove) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeemAndRemove) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeemAndRemove) GetContractAddress() string {
	return s.Raw.Address.String()
}

var _ types.CrossChainUserEventLog = &SynapseBridgeTokenRedeemAndRemove{}

// GetRaw gets the raw event logs from the redeem event.
func (s SynapseBridgeTokenRedeemV2) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetDestinationChainID gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeemV2) GetDestinationChainID() *big.Int {
	return s.ChainId
}

// GetToken gets the destination chain id from the redeem log.
func (s SynapseBridgeTokenRedeemV2) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenRedeemV2) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the redeem event type.
func (s SynapseBridgeTokenRedeemV2) GetEventType() types.EventType {
	return types.RedeemV2Event
}

// GetIdentifier gets the unique identifier (txhash) for the mint.
func (s SynapseBridgeTokenRedeemV2) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenRedeemV2) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenRedeemV2) GetContractAddress() string {
	return s.Raw.Address.String()
}

var _ types.CrossChainUserEventLog = &SynapseBridgeTokenRedeemV2{}
