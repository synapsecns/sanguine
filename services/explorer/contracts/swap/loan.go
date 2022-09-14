package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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

// GetTokenIndex gets the token index.
func (s SwapFlashLoanFlashLoan) GetTokenIndex() *uint8 {
	return &s.TokenIndex
}

// GetAmount gets the amount.
func (s SwapFlashLoanFlashLoan) GetAmount() *big.Int {
	return s.Amount
}

// GetAmountFee gets the amount fee.
func (s SwapFlashLoanFlashLoan) GetAmountFee() *big.Int {
	return s.AmountFee
}

// GetProtocolFee gets the protocol fee.
func (s SwapFlashLoanFlashLoan) GetProtocolFee() *big.Int {
	return s.ProtocolFee
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanFlashLoan) GetLPTokenAmount() *big.Int {
	return nil
}

// GetProvider gets the provider removing liquidity.
func (s SwapFlashLoanFlashLoan) GetProvider() *common.Address {
	return nil
}

// GetTokenAmounts gets the amount of tokens.
func (s SwapFlashLoanFlashLoan) GetTokenAmounts() []*big.Int {
	return nil
}

// GetFees gets the fees for each token.
func (s SwapFlashLoanFlashLoan) GetFees() []*big.Int {
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

// GetSoldId gets the sold id.
func (s SwapFlashLoanFlashLoan) GetSoldId() *big.Int {
	return nil
}

// GetBoughtId gets the bought id.
func (s SwapFlashLoanFlashLoan) GetBoughtId() *big.Int {
	return nil
}

var _ swap.EventLog = &SwapFlashLoanFlashLoan{}
