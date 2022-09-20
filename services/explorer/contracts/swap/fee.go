//nolint:revive,golint,stylecheck
package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"math/big"
)

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanNewAdminFee) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanNewAdminFee) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanNewAdminFee) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanNewAdminFee) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanNewAdminFee) GetEventType() swap.EventType {
	return swap.NewAdminFeeEvent
}

// GetNewAdminFee gets the admin fee.
func (s SwapFlashLoanNewAdminFee) GetNewAdminFee() *big.Int {
	return s.NewAdminFee
}

// GetSoldId gets the solid id.
func (s SwapFlashLoanNewAdminFee) GetSoldId() *big.Int {
	return nil
}

// GetTokenIndex gets the Token index.
func (s SwapFlashLoanNewAdminFee) GetTokenIndex() *uint8 {
	return nil
}

// GetAmount gets the amount.
func (s SwapFlashLoanNewAdminFee) GetAmount() *big.Int {
	return nil
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanNewAdminFee) GetAmountFee() *big.Int {
	return nil
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanNewAdminFee) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanNewAdminFee) GetBuyer() *common.Address {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanNewAdminFee) GetTokensSold() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanNewAdminFee) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanNewAdminFee) GetFutureTime() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanNewAdminFee) GetOldA() *big.Int {
	return nil
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanNewAdminFee) GetLPTokenAmount() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanNewAdminFee) GetNewA() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanNewAdminFee) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanNewAdminFee) GetTime() *big.Int {
	return nil
}

// GetBoughtId gets the bought id.
func (s SwapFlashLoanNewAdminFee) GetBoughtId() *big.Int {
	return nil
}

// GetTokensBought gets the amount of tokens bought.
func (s SwapFlashLoanNewAdminFee) GetTokensBought() *big.Int {
	return nil
}

// GetProvider gets the Provider of the swap.
func (s SwapFlashLoanNewAdminFee) GetProvider() *common.Address {
	return nil
}

// GetTokenAmounts gets the amount of tokens.
func (s SwapFlashLoanNewAdminFee) GetTokenAmounts() []*big.Int {
	return nil
}

// GetFees gets the gets fees.
func (s SwapFlashLoanNewAdminFee) GetFees() []*big.Int {
	return nil
}

// GetInvariant gets the invariant of the swap.
func (s SwapFlashLoanNewAdminFee) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets the LP token supply.
func (s SwapFlashLoanNewAdminFee) GetLPTokenSupply() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanNewAdminFee) GetNewSwapFee() *big.Int {
	return nil
}

var _ swap.EventLog = &SwapFlashLoanNewAdminFee{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanNewSwapFee) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanNewSwapFee) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanNewSwapFee) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanNewSwapFee) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanNewSwapFee) GetEventType() swap.EventType {
	return swap.NewAdminFeeEvent
}

// GetNewSwapFee gets the admin fee.
func (s SwapFlashLoanNewSwapFee) GetNewSwapFee() *big.Int {
	return s.NewSwapFee
}

// GetSoldId gets the solid id.
func (s SwapFlashLoanNewSwapFee) GetSoldId() *big.Int {
	return nil
}

// GetTokenIndex gets the Token index.
func (s SwapFlashLoanNewSwapFee) GetTokenIndex() *uint8 {
	return nil
}

// GetAmount gets the amount.
func (s SwapFlashLoanNewSwapFee) GetAmount() *big.Int {
	return nil
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanNewSwapFee) GetAmountFee() *big.Int {
	return nil
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanNewSwapFee) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanNewSwapFee) GetBuyer() *common.Address {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanNewSwapFee) GetNewAdminFee() *big.Int {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanNewSwapFee) GetTokensSold() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanNewSwapFee) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanNewSwapFee) GetFutureTime() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanNewSwapFee) GetOldA() *big.Int {
	return nil
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanNewSwapFee) GetLPTokenAmount() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanNewSwapFee) GetNewA() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanNewSwapFee) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanNewSwapFee) GetTime() *big.Int {
	return nil
}

// GetBoughtId gets the bought id.
func (s SwapFlashLoanNewSwapFee) GetBoughtId() *big.Int {
	return nil
}

// GetTokensBought gets the amount of tokens bought.
func (s SwapFlashLoanNewSwapFee) GetTokensBought() *big.Int {
	return nil
}

// GetProvider gets the Provider of the swap.
func (s SwapFlashLoanNewSwapFee) GetProvider() *common.Address {
	return nil
}

// GetTokenAmounts gets the amount of tokens.
func (s SwapFlashLoanNewSwapFee) GetTokenAmounts() []*big.Int {
	return nil
}

// GetFees gets the gets fees.
func (s SwapFlashLoanNewSwapFee) GetFees() []*big.Int {
	return nil
}

// GetInvariant gets the invariant of the swap.
func (s SwapFlashLoanNewSwapFee) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets the LP token supply.
func (s SwapFlashLoanNewSwapFee) GetLPTokenSupply() *big.Int {
	return nil
}

var _ swap.EventLog = &SwapFlashLoanNewSwapFee{}
