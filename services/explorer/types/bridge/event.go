package bridge

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all bridge events.
//
//nolint:interfacebloat
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetToken returns the token of the event.
	GetToken() common.Address
	// GetAmount returns the amount of the event.
	GetAmount() *big.Int
	// GetEventType returns the event type of the log.
	GetEventType() EventType
	// GetEventIndex returns the index of the log.
	GetEventIndex() uint64
	// GetRecipient returns the recipient of the event.
	GetRecipient() *common.Address
	// GetDestinationChainID returns the chain id of the token deposit.
	GetDestinationChainID() *big.Int
	// GetTokenIndexFrom returns the token index of the token deposit and swap.
	GetTokenIndexFrom() *uint8
	// GetTokenIndexTo returns the token index of the token deposit and swap.
	GetTokenIndexTo() *uint8
	// GetMinDy returns the min dy of the token deposit and swap.
	GetMinDy() *big.Int
	// GetDeadline returns the deadline of the token deposit and swap.
	GetDeadline() *big.Int
	// GetSwapTokenIndex returns the swap token index of the token redeem and remove.
	GetSwapTokenIndex() *uint8
	// GetSwapMinAmount returns the swap min amount of the token redeem and remove.
	GetSwapMinAmount() *big.Int
	// GetSwapDeadline returns the swap deadline of the token redeem and remove.
	GetSwapDeadline() *big.Int
	// GetRecipientBytes returns the recipient of the event.
	GetRecipientBytes() *[32]byte
	// GetFee returns the fee of the token withdraw.
	GetFee() *big.Int
	// GetKappa returns the kappa of the token withdraw.
	GetKappa() *[32]byte
	// GetSwapSuccess returns the swap success of the token withdraw and remove.
	GetSwapSuccess() *bool
}
