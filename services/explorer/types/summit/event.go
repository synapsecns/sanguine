package summit

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// EventLog is the interface for all summit events.
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

	// GetDomain gets the domain of where the signed Notary/agent is active
	GetDomain() *uint32
	// GetAgent is the notary who signed the attestation
	GetAgent() *common.Address
	// GetRcptPayload gets the raw payload with receipt data
	GetRcptPayload() []byte
	// GetRcptSignature gets raw bytes of the message
	GetRcptSignature() []byte
	// GetSnapshot gets raw payload with snapshot data
	GetSnapshot() []byte
	// GetSnapSignature gets the agent signature for the snapshot
	GetSnapSignature() []byte
	// GetTip gets the tip amount for the TipAwarded event.
	GetTip() *big.Int
}
