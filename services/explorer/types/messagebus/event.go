package messagebus

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all message events.
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
	// GetMessageID returns the message id of the event.
	GetMessageID() *string
	// GetSourceChainID returns the chain id of the message's source chain.
	GetSourceChainID() *big.Int

	// GetSourceAddress gets the address that the message will be passed from.
	GetSourceAddress() *string
	// GetStatus returns the status of the event.
	GetStatus() *string
	// GetDestinationAddress gets the address that the message will be passed to.
	GetDestinationAddress() *string
	// GetDestinationChainID returns the chain id of the message's destination chain.
	GetDestinationChainID() *big.Int

	// GetNonce returns the nonce of the message.
	GetNonce() *big.Int
	// GetMessage gets the message.
	GetMessage() *string
	// GetReceiver returns the receiver of the event.
	GetReceiver() *string
	// GetOptions gets the message.
	GetOptions() *string
	// GetFee returns the fee of the message.
	GetFee() *big.Int
	// GetRevertReason returns the reason why the event was reverted.
	GetRevertReason() *string
}
