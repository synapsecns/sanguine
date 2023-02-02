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
func (s SwapFlashLoanAddLiquidity) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanAddLiquidity) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanAddLiquidity) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanAddLiquidity) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanAddLiquidity) GetEventType() swap.EventType {
	return swap.AddLiquidityEvent
}

// GetInvariant gets the invariant of the swap.
func (s SwapFlashLoanAddLiquidity) GetInvariant() *big.Int {
	return core.CopyBigInt(s.Invariant)
}

// GetProvider gets the provider.
func (s SwapFlashLoanAddLiquidity) GetProvider() *common.Address {
	return &s.Provider
}

// GetAmount gets the token amounts.
func (s SwapFlashLoanAddLiquidity) GetAmount() map[uint8]string {
	tokenAmounts := make(map[uint8]string)
	zero := big.NewInt(0)
	for i, tokenAmount := range s.TokenAmounts {
		if tokenAmount.Cmp(zero) == 1 {
			tokenAmounts[uint8(i)] = core.CopyBigInt(tokenAmount).String()
		}
	}
	return tokenAmounts
}

// GetAmountFee gets the fees.
func (s SwapFlashLoanAddLiquidity) GetAmountFee() map[uint8]string {
	feeAmounts := make(map[uint8]string)
	zero := big.NewInt(0)
	for i, feeAmount := range s.Fees {
		if feeAmount.Cmp(zero) == 1 {
			feeAmounts[uint8(i)] = core.CopyBigInt(feeAmount).String()
		}
	}
	return feeAmounts
}

// GetLPTokenSupply gets the LP token supply.
func (s SwapFlashLoanAddLiquidity) GetLPTokenSupply() *big.Int {
	return core.CopyBigInt(s.LpTokenSupply)
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanAddLiquidity) GetBoughtID() *big.Int {
	return nil
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanAddLiquidity) GetSoldID() *big.Int {
	return nil
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanAddLiquidity) GetLPTokenAmount() *big.Int {
	return nil
}

// GetTokensBought gets the amount of tokens bought.
func (s SwapFlashLoanAddLiquidity) GetTokensBought() *big.Int {
	return nil
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanAddLiquidity) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanAddLiquidity) GetBuyer() *common.Address {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanAddLiquidity) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanAddLiquidity) GetNewSwapFee() *big.Int {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanAddLiquidity) GetTokensSold() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanAddLiquidity) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanAddLiquidity) GetFutureTime() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanAddLiquidity) GetOldA() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanAddLiquidity) GetNewA() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanAddLiquidity) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanAddLiquidity) GetTime() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanAddLiquidity) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanAddLiquidity) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanAddLiquidity{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanRemoveLiquidity) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanRemoveLiquidity) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanRemoveLiquidity) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanRemoveLiquidity) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanRemoveLiquidity) GetEventType() swap.EventType {
	return swap.RemoveLiquidityEvent
}

// GetProvider gets the Provider of the swap.
func (s SwapFlashLoanRemoveLiquidity) GetProvider() *common.Address {
	return &s.Provider
}

// GetAmount gets the token amounts.
func (s SwapFlashLoanRemoveLiquidity) GetAmount() map[uint8]string {
	tokenAmounts := make(map[uint8]string)
	zero := big.NewInt(0)
	for i, tokenAmount := range s.TokenAmounts {
		if tokenAmount.Cmp(zero) == 1 {
			tokenAmounts[uint8(i)] = core.CopyBigInt(tokenAmount).String()
		}
	}
	return tokenAmounts
}

// GetAmountFee gets the fees.
func (s SwapFlashLoanRemoveLiquidity) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
}

// GetLPTokenSupply gets the LP token supply.
func (s SwapFlashLoanRemoveLiquidity) GetLPTokenSupply() *big.Int {
	return s.LpTokenSupply
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanRemoveLiquidity) GetBoughtID() *big.Int {
	return nil
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanRemoveLiquidity) GetSoldID() *big.Int {
	return nil
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanRemoveLiquidity) GetLPTokenAmount() *big.Int {
	return nil
}

// GetTokensBought gets the amount of tokens bought.
func (s SwapFlashLoanRemoveLiquidity) GetTokensBought() *big.Int {
	return nil
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanRemoveLiquidity) GetInvariant() *big.Int {
	return nil
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanRemoveLiquidity) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanRemoveLiquidity) GetBuyer() *common.Address {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanRemoveLiquidity) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanRemoveLiquidity) GetNewSwapFee() *big.Int {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanRemoveLiquidity) GetTokensSold() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanRemoveLiquidity) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanRemoveLiquidity) GetFutureTime() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanRemoveLiquidity) GetOldA() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanRemoveLiquidity) GetNewA() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanRemoveLiquidity) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanRemoveLiquidity) GetTime() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanRemoveLiquidity) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanRemoveLiquidity) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanRemoveLiquidity{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanRemoveLiquidityOne) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanRemoveLiquidityOne) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanRemoveLiquidityOne) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanRemoveLiquidityOne) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanRemoveLiquidityOne) GetEventType() swap.EventType {
	return swap.RemoveLiquidityOneEvent
}

// GetProvider gets the Provider of the swap.
func (s SwapFlashLoanRemoveLiquidityOne) GetProvider() *common.Address {
	return &s.Provider
}

// GetLPTokenAmount gets the LP Token Amount.
func (s SwapFlashLoanRemoveLiquidityOne) GetLPTokenAmount() *big.Int {
	return core.CopyBigInt(s.LpTokenAmount)
}

// GetLPTokenSupply gets the LP Token Supply.
func (s SwapFlashLoanRemoveLiquidityOne) GetLPTokenSupply() *big.Int {
	return core.CopyBigInt(s.LpTokenSupply)
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanRemoveLiquidityOne) GetBoughtID() *big.Int {
	return core.CopyBigInt(s.BoughtId)
}

// GetTokensBought gets the tokens bought.
func (s SwapFlashLoanRemoveLiquidityOne) GetTokensBought() *big.Int {
	return core.CopyBigInt(s.TokensBought)
}

// GetAmount puts the amount in a map with its associated token index.
func (s SwapFlashLoanRemoveLiquidityOne) GetAmount() map[uint8]string {
	output := map[uint8]string{uint8(s.BoughtId.Int64()): core.CopyBigInt(s.TokensBought).String()}
	return output
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanRemoveLiquidityOne) GetSoldID() *big.Int {
	return nil
}

// GetAmountFee gets the amount.
func (s SwapFlashLoanRemoveLiquidityOne) GetAmountFee() map[uint8]string {
	return map[uint8]string{}
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanRemoveLiquidityOne) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanRemoveLiquidityOne) GetBuyer() *common.Address {
	return nil
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanRemoveLiquidityOne) GetInvariant() *big.Int {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanRemoveLiquidityOne) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanRemoveLiquidityOne) GetNewSwapFee() *big.Int {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanRemoveLiquidityOne) GetTokensSold() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanRemoveLiquidityOne) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanRemoveLiquidityOne) GetFutureTime() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanRemoveLiquidityOne) GetOldA() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanRemoveLiquidityOne) GetNewA() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanRemoveLiquidityOne) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanRemoveLiquidityOne) GetTime() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanRemoveLiquidityOne) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanRemoveLiquidityOne) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanRemoveLiquidityOne{}

// GetRaw gets the raw event logs from the redeem and remove event.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the withdraw.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetBlockNumber gets the block number for the event.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetEventType gets the type of the redeem event.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetEventType() swap.EventType {
	return swap.RemoveLiquidityImbalanceEvent
}

// GetProvider gets the Provider of the swap.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetProvider() *common.Address {
	return &s.Provider
}

// GetAmount gets the token amounts.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetAmount() map[uint8]string {
	tokenAmounts := make(map[uint8]string)
	zero := big.NewInt(0)
	for i, tokenAmount := range s.TokenAmounts {
		if tokenAmount.Cmp(zero) == 1 {
			tokenAmounts[uint8(i)] = core.CopyBigInt(tokenAmount).String()
		}
	}
	return tokenAmounts
}

// GetAmountFee gets the fees.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetAmountFee() map[uint8]string {
	feeAmounts := make(map[uint8]string)
	zero := big.NewInt(0)
	for i, feeAmount := range s.Fees {
		if feeAmount.Cmp(zero) == 1 {
			feeAmounts[uint8(i)] = core.CopyBigInt(feeAmount).String()
		}
	}
	return feeAmounts
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetInvariant() *big.Int {
	return core.CopyBigInt(s.Invariant)
}

// GetLPTokenSupply gets the lp token supply.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetLPTokenSupply() *big.Int {
	return core.CopyBigInt(s.LpTokenSupply)
}

// GetSoldID gets the solid id.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetSoldID() *big.Int {
	return nil
}

// GetProtocolFee gets the protocol fee of the tx.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetProtocolFee() *big.Int {
	return nil
}

// GetBuyer gets the buyer.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetBuyer() *common.Address {
	return nil
}

// GetNewAdminFee gets the new admin fee.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetNewAdminFee() *big.Int {
	return nil
}

// GetNewSwapFee gets the new swap fee.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetNewSwapFee() *big.Int {
	return nil
}

// GetTokensSold gets the tokens sold.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetTokensSold() *big.Int {
	return nil
}

// GetInitialTime gets the initial time.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetInitialTime() *big.Int {
	return nil
}

// GetFutureTime gets the future time.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetFutureTime() *big.Int {
	return nil
}

// GetOldA gets the old A.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetOldA() *big.Int {
	return nil
}

// GetLPTokenAmount gets the LP token supply.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetLPTokenAmount() *big.Int {
	return nil
}

// GetNewA gets the new A.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetNewA() *big.Int {
	return nil
}

// GetCurrentA gets the current A.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetCurrentA() *big.Int {
	return nil
}

// GetTime gets the current time.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetTime() *big.Int {
	return nil
}

// GetBoughtID gets the bought id.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetBoughtID() *big.Int {
	return nil
}

// GetTokensBought gets the amount of tokens bought.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetTokensBought() *big.Int {
	return nil
}

// GetReceiver gets the receiver.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetReceiver() *common.Address {
	return nil
}

func (s SwapFlashLoanRemoveLiquidityImbalance) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

var _ swap.EventLog = &SwapFlashLoanRemoveLiquidityImbalance{}
