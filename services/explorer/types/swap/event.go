package swap

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all swap events.
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.'
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetEventType returns the event type of the log.
	GetEventType() EventType
}

// FlashLoanLog is the interface for the flash loan event.
type FlashLoanLog interface {
	EventLog
	// GetTokenIndex returns the token index of the flash loan.
	GetTokenIndex() uint8
	// GetAmount returns the amount of the flash loan.
	GetAmount() *big.Int
	// GetAmountFee returns the amount fee of the flash loan.
	GetAmountFee() *big.Int
	// GetProtocolFee returns the protocol fee of the flash loan.
	GetProtocolFee() *big.Int
}

// TokenSwapLog is the interface for the token swap event.
type TokenSwapLog interface {
	EventLog
	// GetBuyer returns the buyer of the token swap.
	GetBuyer() common.Address
	// GetTokensSold returns the amount of tokens sold.
	GetTokensSold() *big.Int
	// GetTokensBought returns the amount of tokens bought.
	GetTokensBought() *big.Int
	// GetSoldId returns the token id of the sold token.
	GetSoldId() *big.Int
	// GetBoughtId returns the token id of the bought token.
	GetBoughtId() *big.Int
}

// AddLiquidityLog is the interface for the add liquidity event.
type AddLiquidityLog interface {
	EventLog
	// GetProvider returns the provider adding liquidity.
	GetProvider() common.Address
	// GetTokenAmounts returns the amount of tokens added.
	GetTokenAmounts() []*big.Int
	// GetFees returns the fees for each token.
	GetFees() []*big.Int
	// GetInvariant returns the invariant.
	GetInvariant() *big.Int
	// GetLPTokenSupply returns the LP token supply.
	GetLPTokenSupply() *big.Int
}

// RemoveLiquidityLog is the interface for the remove liquidity event.
type RemoveLiquidityLog interface {
	EventLog
	// GetProvider returns the provider removing liquidity.
	GetProvider() common.Address
	// GetTokenAmounts returns the amount of tokens removed.
	GetTokenAmounts() []*big.Int
	// GetLPTokenSupply returns the LP token supply.
	GetLPTokenSupply() *big.Int
}

// RemoveLiquidityOneLog is the interface for the remove liquidity one event.
type RemoveLiquidityOneLog interface {
	EventLog
	// GetProvider returns the provider removing liquidity.
	GetProvider() common.Address
	// GetLPTokenAmount returns the LP token amount.
	GetLPTokenAmount() *big.Int
	// GetLPTokenSupply returns the LP token supply.
	GetLPTokenSupply() *big.Int
	// GetBoughtId returns the token id of the bought token.
	GetBoughtId() *big.Int
	// GetTokensBought returns the amount of tokens bought.
	GetTokensBought() *big.Int
}

// RemoveLiquidityImbalanceLog is the interface for the remove liquidity imbalance event.
type RemoveLiquidityImbalanceLog interface {
	EventLog
	// GetProvider returns the provider removing liquidity.
	GetProvider() common.Address
	// GetTokenAmounts returns the amount of tokens removed.
	GetTokenAmounts() []*big.Int
	// GetFees returns the fees for each token.
	GetFees() []*big.Int
	// GetInvariant returns the invariant.
	GetInvariant() *big.Int
	// GetLPTokenSupply returns the LP token supply.
	GetLPTokenSupply() *big.Int
}

// NewAdminFeeLog is the interface for the new admin fee event.
type NewAdminFeeLog interface {
	EventLog
	// GetNewAdminFee returns the new admin fee.
	GetNewAdminFee() *big.Int
}

// NewSwapFeeLog is the interface for the new swap fee event.
type NewSwapFeeLog interface {
	EventLog
	// GetNewSwapFee returns the new swap fee.
	GetNewSwapFee() *big.Int
}

// RampALog is the interface for the ramp A event.
type RampALog interface {
	EventLog
	// GetOldA returns the old A.
	GetOldA() *big.Int
	// GetNewA returns the new A.
	GetNewA() *big.Int
	// GetInitialTime returns the initial time.
	GetInitialTime() *big.Int
	// GetFutureTime returns the future time.
	GetFutureTime() *big.Int
}

// StopRampALog is the interface for the stop ramp A event.
type StopRampALog interface {
	EventLog
	// GetCurrentA returns the current A.
	GetCurrentA() *big.Int
	// GetTime returns the time.
	GetTime() *big.Int
}
