package destination

import (
	"github.com/ethereum/go-ethereum/common"
)

// EventLog is the interface for all destination events.
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
	// GetAttestation gets raw payload with attestation data
	GetAttestation() []byte
	// GetAttSignature gets the notary signature for the attestation
	GetAttSignature() []byte
	// GetAgentRoot gets the root
	GetAgentRoot() *[32]byte
}
