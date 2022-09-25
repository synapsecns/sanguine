//nolint:revive,golint,stylecheck
package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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
	return s.TokensSold
}

// GetTokensBought gets the tokens bought.
func (s SwapFlashLoanTokenSwap) GetTokensBought() *big.Int {
	return s.TokensBought
}

// GetSoldId gets the solid id.
func (s SwapFlashLoanTokenSwap) GetSoldId() *big.Int {
	return s.SoldId
}

// GetBoughtId gets the bought id.
func (s SwapFlashLoanTokenSwap) GetBoughtId() *big.Int {
	return s.BoughtId
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanTokenSwap) GetLPTokenAmount() *big.Int {
	return nil
}

// GetTokenIndex gets the Token index.
func (s SwapFlashLoanTokenSwap) GetTokenIndex() *uint8 {
	return nil
}

// GetAmount gets the amount.
func (s SwapFlashLoanTokenSwap) GetAmount() *big.Int {
	return nil
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanTokenSwap) GetAmountFee() *big.Int {
	return nil
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanTokenSwap) GetProtocolFee() *big.Int {
	return nil
}

// GetProvider gets the provider removing liquidity.
func (s SwapFlashLoanTokenSwap) GetProvider() *common.Address {
	return nil
}

// GetTokenAmounts gets the amount of tokens.
func (s SwapFlashLoanTokenSwap) GetTokenAmounts() []*big.Int {
	return nil
}

// GetFees gets the fees for each token.
func (s SwapFlashLoanTokenSwap) GetFees() []*big.Int {
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

var _ swap.EventLog = &SwapFlashLoanTokenSwap{}
