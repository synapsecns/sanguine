//nolint:revive,golint
package bridgev1

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SynapseBridgeTokenWithdraw) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetToken gets the destination chain id from the redeem and remove log.
func (s SynapseBridgeTokenWithdraw) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the token amount.
func (s SynapseBridgeTokenWithdraw) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the type of the redeem event.
func (s SynapseBridgeTokenWithdraw) GetEventType() bridge.EventType {
	return bridge.WithdrawEvent
}

// GetFee gets the fee for the token withdraw.
func (s SynapseBridgeTokenWithdraw) GetFee() *big.Int {
	return core.CopyBigInt(s.Fee)
}

// GetKappa gets the gappa for the token withdraw.
func (s SynapseBridgeTokenWithdraw) GetKappa() *[32]byte {
	return &s.Kappa
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SynapseBridgeTokenWithdraw) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenWithdraw) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenWithdraw) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRecipient gets the recipient of the withdrawal.
func (s SynapseBridgeTokenWithdraw) GetRecipient() *common.Address {
	return &s.To
}

func (s SynapseBridgeTokenWithdraw) GetDestinationChainID() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetTokenIndexFrom() *uint8 {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetTokenIndexTo() *uint8 {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetMinDy() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetSwapTokenIndex() *uint8 {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetSwapMinAmount() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetSwapDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetSwapSuccess() *bool {
	return nil
}

func (s SynapseBridgeTokenWithdraw) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenWithdraw{}

// GetToken gets the token for the withdraw and remove operation.
func (s SynapseBridgeTokenWithdrawAndRemove) GetToken() common.Address {
	return s.Token
}

// GetAmount gets the amount fo the withdraw.
func (s SynapseBridgeTokenWithdrawAndRemove) GetAmount() *big.Int {
	return core.CopyBigInt(s.Amount)
}

// GetEventType gets the withdraw and remove event type.
func (s SynapseBridgeTokenWithdrawAndRemove) GetEventType() bridge.EventType {
	return bridge.WithdrawAndRemoveEvent
}

// GetTxHash gets the unique identifier (txhash) for the withdraw and remove.
func (s SynapseBridgeTokenWithdrawAndRemove) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SynapseBridgeTokenWithdrawAndRemove) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SynapseBridgeTokenWithdrawAndRemove) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetRaw gets the raw logs.
func (s SynapseBridgeTokenWithdrawAndRemove) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetFee gets the fee for the withdraw and remove.
func (s SynapseBridgeTokenWithdrawAndRemove) GetFee() *big.Int {
	return core.CopyBigInt(s.Fee)
}

// GetKappa gets the kappa value for the withdraw and remove.
func (s SynapseBridgeTokenWithdrawAndRemove) GetKappa() *[32]byte {
	return &s.Kappa
}

// GetRecipient gets the recipient of the withdraw and remove.
func (s SynapseBridgeTokenWithdrawAndRemove) GetRecipient() *common.Address {
	return &s.To
}

// GetSwapTokenIndex gets the index of the token to swap.
func (s SynapseBridgeTokenWithdrawAndRemove) GetSwapTokenIndex() *uint8 {
	return &s.SwapTokenIndex
}

// GetSwapMinAmount gets the minimum amount to swap.
func (s SynapseBridgeTokenWithdrawAndRemove) GetSwapMinAmount() *big.Int {
	return core.CopyBigInt(s.SwapMinAmount)
}

// GetSwapDeadline gets the swap deadline.
func (s SynapseBridgeTokenWithdrawAndRemove) GetSwapDeadline() *big.Int {
	return core.CopyBigInt(s.SwapDeadline)
}

// GetSwapSuccess gets the swap success.
func (s SynapseBridgeTokenWithdrawAndRemove) GetSwapSuccess() *bool {
	return &s.SwapSuccess
}

func (s SynapseBridgeTokenWithdrawAndRemove) GetDestinationChainID() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdrawAndRemove) GetTokenIndexFrom() *uint8 {
	return nil
}

func (s SynapseBridgeTokenWithdrawAndRemove) GetTokenIndexTo() *uint8 {
	return nil
}

func (s SynapseBridgeTokenWithdrawAndRemove) GetMinDy() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdrawAndRemove) GetDeadline() *big.Int {
	return nil
}

func (s SynapseBridgeTokenWithdrawAndRemove) GetRecipientBytes() *[32]byte {
	return nil
}

func (s SynapseBridgeTokenWithdrawAndRemove) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ bridge.EventLog = &SynapseBridgeTokenWithdrawAndRemove{}
