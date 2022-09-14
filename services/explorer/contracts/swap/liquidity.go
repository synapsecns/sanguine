package swap

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
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

// GetFees gets the fees.
func (s SwapFlashLoanAddLiquidity) GetFees() []*big.Int {
	return s.Fees
}

// GetInvariant gets the invariant of the swap.
func (s SwapFlashLoanAddLiquidity) GetInvariant() *big.Int {
	return s.Invariant
}

// GetProvider gets the provider
func (s SwapFlashLoanAddLiquidity) GetProvider() common.Address {
	return s.Provider
}

// GetTokenAmounts gets the token amounts.
func (s SwapFlashLoanAddLiquidity) GetTokenAmounts() []*big.Int {
	return s.TokenAmounts
}

// GetLPTokenSupply gets the LP token supply.
func (s SwapFlashLoanAddLiquidity) GetLPTokenSupply() *big.Int {
	return s.LpTokenSupply
}

var _ swap.AddLiquidityLog = &SwapFlashLoanAddLiquidity{}

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
func (s SwapFlashLoanRemoveLiquidity) GetProvider() common.Address {
	return s.Provider
}

// GetTokenAmounts gets the invariant of the swap.
func (s SwapFlashLoanRemoveLiquidity) GetTokenAmounts() []*big.Int {
	return s.TokenAmounts
}

// GetLPTokenSupply gets the LP token supply.
func (s SwapFlashLoanRemoveLiquidity) GetLPTokenSupply() *big.Int {
	return s.LpTokenSupply
}

var _ swap.RemoveLiquidityLog = &SwapFlashLoanRemoveLiquidity{}

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
func (s SwapFlashLoanRemoveLiquidityOne) GetProvider() common.Address {
	return s.Provider
}

// GetLPTokenAmount gets the LP Token Amount.
func (s SwapFlashLoanRemoveLiquidityOne) GetLPTokenAmount() *big.Int {
	return s.LpTokenAmount
}

// GetLPTokenSupply gets the LP Token Supply.
func (s SwapFlashLoanRemoveLiquidityOne) GetLPTokenSupply() *big.Int {
	return s.LpTokenSupply
}

// GetBoughtId gets the bought id.
func (s SwapFlashLoanRemoveLiquidityOne) GetBoughtId() *big.Int {
	return s.BoughtId
}

// GetTokensBought gets the tokens bought.
func (s SwapFlashLoanRemoveLiquidityOne) GetTokensBought() *big.Int {
	return s.TokensBought
}

var _ swap.RemoveLiquidityOneLog = &SwapFlashLoanRemoveLiquidityOne{}

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
func (s SwapFlashLoanRemoveLiquidityImbalance) GetProvider() common.Address {
	return s.Provider
}

// GetTokenAmounts gets the invariant of the swap.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetTokenAmounts() []*big.Int {
	return s.TokenAmounts
}

// GetFees gets the gets fees.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetFees() []*big.Int {
	return s.Fees
}

// GetInvariant gets the invariant.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetInvariant() *big.Int {
	return s.Invariant
}

// GetLPTokenSupply gets the lp token supply.
func (s SwapFlashLoanRemoveLiquidityImbalance) GetLPTokenSupply() *big.Int {
	return s.LpTokenSupply
}

var _ swap.RemoveLiquidityImbalanceLog = &SwapFlashLoanRemoveLiquidityImbalance{}
