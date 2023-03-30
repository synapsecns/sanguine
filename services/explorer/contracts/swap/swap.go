//nolint:revive,golint,stylecheck
package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs.
func (s SwapFlashLoanTokenSwap) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s SwapFlashLoanTokenSwap) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanTokenSwap) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanTokenSwap) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the swap event.
func (s SwapFlashLoanTokenSwap) GetEventType() swap.EventType {
	return swap.TokenSwapEvent
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanTokenSwap) GetBuyer() *common.Address {
	return &s.Buyer
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanTokenSwap) GetTokensSold() *big.Int {
	return core.CopyBigInt(s.TokensSold)
}

// GetTokensBought gets the tokens bought.
func (s SwapFlashLoanTokenSwap) GetTokensBought() *big.Int {
	return core.CopyBigInt(s.TokensBought)
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanTokenSwap) GetSoldID() *big.Int {
	return core.CopyBigInt(s.SoldId)
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanTokenSwap) GetBoughtID() *big.Int {
	return core.CopyBigInt(s.BoughtId)
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanTokenSwap) GetLPTokenAmount() *big.Int {
	return nil
}

// GetAmount puts the amount in a map with it's associated token index.
func (s SwapFlashLoanTokenSwap) GetAmount() map[uint8]string {
	output := map[uint8]string{uint8(s.SoldId.Int64()): core.CopyBigInt(s.TokensSold).String(), uint8(s.BoughtId.Int64()): core.CopyBigInt(s.TokensBought).String()}
	return output
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanTokenSwap) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanTokenSwap) GetProtocolFee() *big.Int {
	return nil
}

// GetProvider gets the provider removing liquidity.
func (s SwapFlashLoanTokenSwap) GetProvider() *common.Address {
	return nil
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanTokenSwap) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets LP token supply.
func (s SwapFlashLoanTokenSwap) GetLPTokenSupply() *big.Int {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanTokenSwap) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanTokenSwap) GetNewSwapFee() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanTokenSwap) GetOldA() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanTokenSwap) GetNewA() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanTokenSwap) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanTokenSwap) GetFutureTime() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanTokenSwap) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanTokenSwap) GetTime() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanTokenSwap) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanTokenSwap) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanTokenSwap{}
