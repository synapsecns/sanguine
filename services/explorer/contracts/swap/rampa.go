//nolint:revive,golint,stylecheck
package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs from the Ramp A Event.
func (s SwapFlashLoanRampA) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s SwapFlashLoanRampA) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanRampA) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanRampA) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the ramp A event.
func (s SwapFlashLoanRampA) GetEventType() swap.EventType {
	return swap.RampAEvent
}

// GetOldA gets the old A.
func (s SwapFlashLoanRampA) GetOldA() *big.Int {
	return core.CopyBigInt(s.OldA)
}

// GetNewA gets the new A.
func (s SwapFlashLoanRampA) GetNewA() *big.Int {
	return core.CopyBigInt(s.NewA)
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanRampA) GetInitialTime() *big.Int {
	return core.CopyBigInt(s.InitialTime)
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanRampA) GetFutureTime() *big.Int {
	return core.CopyBigInt(s.FutureTime)
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanRampA) GetBoughtID() *big.Int {
	return nil
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanRampA) GetSoldID() *big.Int {
	return nil
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanRampA) GetLPTokenAmount() *big.Int {
	return nil
}

// GetTokensBought gets the amount of tokens bought.
func (s SwapFlashLoanRampA) GetTokensBought() *big.Int {
	return nil
}

// GetAmount gets the amount.
func (s SwapFlashLoanRampA) GetAmount() map[uint8]string {
	return nil
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanRampA) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanRampA) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanRampA) GetBuyer() *common.Address {
	return nil
}

// GetProvider gets the provider removing liquidity.
func (s SwapFlashLoanRampA) GetProvider() *common.Address {
	return nil
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanRampA) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets LP token supply.
func (s SwapFlashLoanRampA) GetLPTokenSupply() *big.Int {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanRampA) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanRampA) GetNewSwapFee() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanRampA) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanRampA) GetTime() *big.Int {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanRampA) GetTokensSold() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanRampA) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanRampA) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanRampA{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanStopRampA) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanStopRampA) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanStopRampA) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanStopRampA) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanStopRampA) GetEventType() swap.EventType {
	return swap.StopRampAEvent
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanStopRampA) GetCurrentA() *big.Int {
	return core.CopyBigInt(s.CurrentA)
}

// GetTime gets the time.
func (s SwapFlashLoanStopRampA) GetTime() *big.Int {
	return core.CopyBigInt(s.Time)
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanStopRampA) GetBoughtID() *big.Int {
	return nil
}

// GetSoldID gets the sold id.
func (s SwapFlashLoanStopRampA) GetSoldID() *big.Int {
	return nil
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanStopRampA) GetLPTokenAmount() *big.Int {
	return nil
}

// GetTokensBought gets the amount of tokens bought.
func (s SwapFlashLoanStopRampA) GetTokensBought() *big.Int {
	return nil
}

// GetAmount gets the amount.
func (s SwapFlashLoanStopRampA) GetAmount() map[uint8]string {
	return nil
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanStopRampA) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanStopRampA) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanStopRampA) GetBuyer() *common.Address {
	return nil
}

// GetProvider gets the provider removing liquidity.
func (s SwapFlashLoanStopRampA) GetProvider() *common.Address {
	return nil
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanStopRampA) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets LP token supply.
func (s SwapFlashLoanStopRampA) GetLPTokenSupply() *big.Int {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanStopRampA) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanStopRampA) GetNewSwapFee() *big.Int {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanStopRampA) GetTokensSold() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanStopRampA) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanStopRampA) GetFutureTime() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanStopRampA) GetOldA() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanStopRampA) GetNewA() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanStopRampA) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanStopRampA) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanStopRampA{}
