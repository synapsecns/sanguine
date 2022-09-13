package bridge

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all bridge events.
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetToken returns the token of the event.
	GetToken() common.Address
	// GetRecipient returns the recipient of the event.
	GetRecipient() common.Address
	// GetAmount returns the amount of the event.
	GetAmount() *big.Int
	// GetEventType returns the event type of the log.
	GetEventType() EventType
}

// DepositLog is the interface for the token deposit event.
type DepositLog interface {
	EventLog
	// GetDestinationChainID returns the chain id of the token deposit.
	GetDestinationChainID() *big.Int
}

// DepositAndSwapLog is the interface for the token deposit and swap event.
type DepositAndSwapLog interface {
	EventLog
	// GetDestinationChainID returns the chain id of the token deposit and swap.
	GetDestinationChainID() *big.Int
	// GetTokenIndexFrom returns the token index of the token deposit and swap.
	GetTokenIndexFrom() uint8
	// GetTokenIndexTo returns the token index of the token deposit and swap.
	GetTokenIndexTo() uint8
	// GetMinDy returns the min dy of the token deposit and swap.
	GetMinDy() *big.Int
	// GetDeadline returns the deadline of the token deposit and swap.
	GetDeadline() *big.Int
}

// RedeemLog is the interface for the token redeem event.
type RedeemLog interface {
	EventLog
	// GetDestinationChainID returns the chain id of the token redeem.
	GetDestinationChainID() *big.Int
}

// RedeemAndSwapLog is the interface for the token redeem and swap event.
type RedeemAndSwapLog interface {
	EventLog
	// GetDestinationChainID returns the chain id of the token redeem and swap.
	GetDestinationChainID() *big.Int
	// GetTokenIndexFrom returns the token index of the token redeem and swap.
	GetTokenIndexFrom() uint8
	// GetTokenIndexTo returns the token index of the token redeem and swap.
	GetTokenIndexTo() uint8
	// GetMinDy returns the min dy of the token redeem and swap.
	GetMinDy() *big.Int
	// GetDeadline returns the deadline of the token redeem and swap.
	GetDeadline() *big.Int
}

// RedeemAndRemoveLog is the interface for the token redeem and remove event.
type RedeemAndRemoveLog interface {
	EventLog
	// GetDestinationChainID returns the chain id of the token redeem and remove.
	GetDestinationChainID() *big.Int
	// GetSwapTokenIndex returns the swap token index of the token redeem and remove.
	GetSwapTokenIndex() uint8
	// GetSwapMinAmount returns the swap min amount of the token redeem and remove.
	GetSwapMinAmount() *big.Int
	// GetSwapDeadline returns the swap deadline of the token redeem and remove.
	GetSwapDeadline() *big.Int
}

// RedeemV2Log is the interface for the token redeem v2 event.
type RedeemV2Log interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetToken returns the token of the event.
	GetToken() common.Address
	// GetRecipient returns the recipient of the event.
	GetRecipient() [32]byte
	// GetAmount returns the amount of the event.
	GetAmount() *big.Int
	// GetEventType returns the event type of the log.
	GetEventType() EventType
	// GetDestinationChainID returns the chain id of the token redeem v2.
	GetDestinationChainID() *big.Int
}

// WithdrawLog is the interface for the token withdraw event.
type WithdrawLog interface {
	EventLog
	// GetFee returns the fee of the token withdraw.
	GetFee() *big.Int
	// GetKappa returns the kappa of the token withdraw.
	GetKappa() [32]byte
}

// WithdrawAndRemoveLog is the interface for the token withdraw and remove event.
type WithdrawAndRemoveLog interface {
	EventLog
	// GetFee returns the fee of the token withdraw and remove.
	GetFee() *big.Int
	// GetSwapTokenIndex returns the swap token index of the token withdraw and remove.
	GetSwapTokenIndex() uint8
	// GetSwapMinAmount returns the swap min amount of the token withdraw and remove.
	GetSwapMinAmount() *big.Int
	// GetSwapDeadline returns the swap deadline of the token withdraw and remove.
	GetSwapDeadline() *big.Int
	// GetSwapSuccess returns the swap success of the token withdraw and remove.
	GetSwapSuccess() bool
	// GetKappa returns the kappa of the token withdraw and remove.
	GetKappa() [32]byte
}

// MintLog is the interface for the token mint event.
type MintLog interface {
	EventLog
	// GetFee returns the fee of the token mint.
	GetFee() *big.Int
	// GetKappa returns the kappa of the token mint.
	GetKappa() [32]byte
}

// MintAndSwapLog is the interface for the token mint and swap event.
type MintAndSwapLog interface {
	EventLog
	// GetFee returns the fee of the token mint and swap.
	GetFee() *big.Int
	// GetTokenIndexFrom returns the token index of the token mint and swap.
	GetTokenIndexFrom() uint8
	// GetTokenIndexTo returns the token index of the token mint and swap.
	GetTokenIndexTo() uint8
	// GetMinDy returns the min dy of the token mint and swap.
	GetMinDy() *big.Int
	// GetDeadline returns the deadline of the token mint and swap.
	GetDeadline() *big.Int
	// GetSwapSuccess returns the swap success of the token mint and swap.
	GetSwapSuccess() bool
	// GetKappa returns the kappa of the token mint and swap.
	GetKappa() [32]byte
}
