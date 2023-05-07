package origin

import (
	"github.com/ethereum/go-ethereum/common"
)

// EventLog is the interface for all message events.
//
//nolint:interfacebloat
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetEventType returns the event type of the log.
	GetEventType() EventType
	// GetEventIndex returns the index of the log.
	GetEventIndex() uint64
	// GetMessageHash gets the hash of message; the leaf inserted to the Merkle tree for the message
	GetMessageHash() [32]byte
	// GetNonce gets the nonce of sent message (starts from 1)
	GetNonce() uint32
	// GetDestination gets the destination domain
	GetDestination() uint32
	// GetMessage gets raw bytes of the message
	GetMessage() []byte
}
