package bridge

import (
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/agents/types"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SynapseBridgeTokenWithdraw) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetToken gets the destination chain id from the redeem and remove log.
func (s SynapseBridgeTokenWithdraw) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenWithdraw) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenWithdraw) GetEventType() types.EventType {
	return types.WithdrawEvent
}

// GetFee gets the fee for the token withdraw.
func (s SynapseBridgeTokenWithdraw) GetFee() *big.Int {
	return s.Fee
}

// GetKappa gets the gappa for the token withdraw.
func (s SynapseBridgeTokenWithdraw) GetKappa() [32]byte {
	return s.Kappa
}

// GetIdentifier gets the unique identifier (txhash) for the withdraw.
func (s SynapseBridgeTokenWithdraw) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenWithdraw) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenWithdraw) GetContractAddress() string {
	return s.Raw.Address.String()
}

var _ types.CrossChainBridgeEventLog = &SynapseBridgeTokenWithdraw{}

// GetToken gets the token for the withdraw and remove operation.
func (s SynapseBridgeTokenWithdrawAndRemove) GetToken() string {
	return s.Token.String()
}

// GetAmount gets the amount fo the withdraw.
func (s SynapseBridgeTokenWithdrawAndRemove) GetAmount() *big.Int {
	return s.Amount
}

// GetEventType gets the withdraw and remove event type.
func (s SynapseBridgeTokenWithdrawAndRemove) GetEventType() types.EventType {
	return types.WithdrawAndRemove
}

// GetIdentifier gets the unique identifier (txhash) for the withdraw and remove.
func (s SynapseBridgeTokenWithdrawAndRemove) GetIdentifier() string {
	return s.Raw.TxHash.String()
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenWithdrawAndRemove) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenWithdrawAndRemove) GetContractAddress() string {
	return s.Raw.Address.String()
}

// GetRaw gets the raw logs.
func (s SynapseBridgeTokenWithdrawAndRemove) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetFee gets the fee for the withdraw and remove.
func (s SynapseBridgeTokenWithdrawAndRemove) GetFee() *big.Int {
	return s.Fee
}

// GetKappa gets the kappa value for the withdraw and remove.
func (s SynapseBridgeTokenWithdrawAndRemove) GetKappa() [32]byte {
	return s.Kappa
}

var _ types.CrossChainBridgeEventLog = &SynapseBridgeTokenWithdrawAndRemove{}
