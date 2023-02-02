package swap

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all swap events.
//
//nolint:interfacebloat
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.'
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetEventType returns the event type of the log.
	GetEventType() EventType
	// GetEventIndex returns the index of the log.
	GetEventIndex() uint64
	// GetAmount returns the token amount with its token index.
	GetAmount() map[uint8]string
	// GetAmountFee returns the amount fee of the tx.
	GetAmountFee() map[uint8]string
	// GetProtocolFee returns the protocol fee from the tx.
	GetProtocolFee() *big.Int
	// GetBuyer returns the buyer of the token swap.
	GetBuyer() *common.Address
	// GetTokensSold returns the amount of tokens sold.
	GetTokensSold() *big.Int
	// GetSoldID returns the token id of the sold token.
	GetSoldID() *big.Int
	// GetLPTokenAmount returns the LP token amount.
	GetLPTokenAmount() *big.Int
	// GetBoughtID returns the token id of the bought token.
	GetBoughtID() *big.Int
	// GetTokensBought returns the amount of tokens bought.
	GetTokensBought() *big.Int
	// GetProvider returns the provider removing liquidity.
	GetProvider() *common.Address
	// GetInvariant returns the invariant.
	GetInvariant() *big.Int
	// GetLPTokenSupply returns the LP token supply.
	GetLPTokenSupply() *big.Int
	// GetNewAdminFee returns the new admin fee.
	GetNewAdminFee() *big.Int
	// GetNewSwapFee returns the new swap fee.
	GetNewSwapFee() *big.Int
	// GetOldA returns the old A.
	GetOldA() *big.Int
	// GetNewA returns the new A.
	GetNewA() *big.Int
	// GetInitialTime returns the initial time.
	GetInitialTime() *big.Int
	// GetFutureTime returns the future time.
	GetFutureTime() *big.Int
	// GetCurrentA returns the current A.
	GetCurrentA() *big.Int
	// GetTime returns the time.
	GetTime() *big.Int
	// GetReceiver returns the receiver.
	GetReceiver() *common.Address
}
