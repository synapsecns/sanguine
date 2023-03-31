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
func (s SwapFlashLoanFlashLoan) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) from the event.
func (s SwapFlashLoanFlashLoan) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanFlashLoan) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanFlashLoan) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanFlashLoan) GetEventType() swap.EventType {
	return swap.FlashLoanEvent
}

// GetAmount puts the amount in a map with it's associated token index.
func (s SwapFlashLoanFlashLoan) GetAmount() map[uint8]string {
	output := map[uint8]string{s.TokenIndex: core.CopyBigInt(s.Amount).String()}
	return output
}

// GetAmountFee gets the amount fee.
func (s SwapFlashLoanFlashLoan) GetAmountFee() map[uint8]string {
	output := map[uint8]string{s.TokenIndex: core.CopyBigInt(s.AmountFee).String()}
	return output
}

// GetProtocolFee gets the protocol fee.
func (s SwapFlashLoanFlashLoan) GetProtocolFee() *big.Int {
	return core.CopyBigInt(s.ProtocolFee)
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanFlashLoan) GetLPTokenAmount() *big.Int {
	return nil
}

// GetProvider gets the provider removing liquidity.
func (s SwapFlashLoanFlashLoan) GetProvider() *common.Address {
	return nil
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanFlashLoan) GetInvariant() *big.Int {
	return nil
}

// GetLPTokenSupply gets LP token supply.
func (s SwapFlashLoanFlashLoan) GetLPTokenSupply() *big.Int {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanFlashLoan) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanFlashLoan) GetNewSwapFee() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanFlashLoan) GetOldA() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanFlashLoan) GetNewA() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanFlashLoan) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanFlashLoan) GetFutureTime() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanFlashLoan) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanFlashLoan) GetTime() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanFlashLoan) GetBuyer() *common.Address {
	return nil
}

// GetTokensSold gets the token sold.
func (s SwapFlashLoanFlashLoan) GetTokensSold() *big.Int {
	return nil
}

// GetTokensBought gets the tokens bought.
func (s SwapFlashLoanFlashLoan) GetTokensBought() *big.Int {
	return nil
}

// GetSoldID gets the sold id.
func (s SwapFlashLoanFlashLoan) GetSoldID() *big.Int {
	return nil
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanFlashLoan) GetBoughtID() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanFlashLoan) GetReceiver() *common.Address {
	return &s.Receiver
}

func (s SwapFlashLoanFlashLoan) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanFlashLoan{}
