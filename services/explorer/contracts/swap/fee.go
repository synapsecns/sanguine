//nolint:revive,golint,stylecheck
package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core"
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
	return core.CopyBigInt(s.NewAdminFee)
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanNewAdminFee) GetSoldID() *big.Int {
	return nil
}

// GetAmount gets the amount(s).
func (s SwapFlashLoanNewAdminFee) GetAmount() map[uint8]string {
	return nil
}

// GetAmountFee gets the fee amount(s).
func (s SwapFlashLoanNewAdminFee) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
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

// GetBoughtID gets the bought id.
func (s SwapFlashLoanNewAdminFee) GetBoughtID() *big.Int {
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

// GetReceiver gets the receiver.
func (s SwapFlashLoanNewAdminFee) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanNewAdminFee) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
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
	return swap.NewSwapFeeEvent
}

// GetNewSwapFee gets the admin fee.
func (s SwapFlashLoanNewSwapFee) GetNewSwapFee() *big.Int {
	return core.CopyBigInt(s.NewSwapFee)
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanNewSwapFee) GetSoldID() *big.Int {
	return nil
}

// GetAmount gets the amount.
func (s SwapFlashLoanNewSwapFee) GetAmount() map[uint8]string {
	return nil
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanNewSwapFee) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
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

// GetBoughtID gets the bought id.
func (s SwapFlashLoanNewSwapFee) GetBoughtID() *big.Int {
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

// GetInvariant gets the invariant of the swap.
func (s SwapFlashLoanNewSwapFee) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets the LP token supply.
func (s SwapFlashLoanNewSwapFee) GetLPTokenSupply() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanNewSwapFee) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanNewSwapFee) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanNewSwapFee{}
