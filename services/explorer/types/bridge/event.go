package bridge

import "github.com/ethereum/go-ethereum/common"

// EventLog is the interface for all bridge events.
type EventLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetChainID returns the chain id of the log.
	GetChainID() uint32
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetEventType returns the event type of the log.
	GetEventType() EventType
}
