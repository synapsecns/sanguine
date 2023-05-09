package destination

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/destination"
)

// GetEventType gets the execute event type.
func (d DestinationAttestationAccepted) GetEventType() destination.EventType {
	return destination.AttestationAcceptedEvent
}

// GetRaw gets the raw logs.
func (d DestinationAttestationAccepted) GetRaw() ethTypes.Log {
	return d.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (d DestinationAttestationAccepted) GetTxHash() common.Hash {
	return d.Raw.TxHash
}

// GetEventIndex gets the event index.
func (d DestinationAttestationAccepted) GetEventIndex() uint64 {
	return uint64(d.Raw.Index)
}

// GetBlockNumber gets the block number for the event.
func (d DestinationAttestationAccepted) GetBlockNumber() uint64 {
	return d.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (d DestinationAttestationAccepted) GetContractAddress() common.Address {
	return d.Raw.Address
}

// GetDomain gets domain where the agent is active
func (d DestinationAttestationAccepted) GetDomain() *uint32 {
	return &(d.Domain)
}

// GetAgent gets the address of the agent
func (d DestinationAttestationAccepted) GetAgent() *common.Address {
	return &(d.Notary)
}

// GetAttestation gets raw payload with attestation data
func (d DestinationAttestationAccepted) GetAttestation() []byte {
	return d.Attestation
}

// GetAttSignature gets the notary signature for the attestation
func (d DestinationAttestationAccepted) GetAttSignature() []byte {
	return d.AttSignature
}

// GetAgentRoot gets the root
func (d DestinationAttestationAccepted) GetAgentRoot() *[32]byte {
	return nil
}

var _ destination.EventLog = &DestinationAttestationAccepted{}

// GetEventType gets the execute event type.
func (d DestinationAgentRootAccepted) GetEventType() destination.EventType {
	return destination.AgentRootAcceptedEvent
}

// GetRaw gets the raw logs.
func (d DestinationAgentRootAccepted) GetRaw() ethTypes.Log {
	return d.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (d DestinationAgentRootAccepted) GetTxHash() common.Hash {
	return d.Raw.TxHash
}

// GetEventIndex gets the event index.
func (d DestinationAgentRootAccepted) GetEventIndex() uint64 {
	return uint64(d.Raw.Index)
}

// GetBlockNumber gets the block number for the event.
func (d DestinationAgentRootAccepted) GetBlockNumber() uint64 {
	return d.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (d DestinationAgentRootAccepted) GetContractAddress() common.Address {
	return d.Raw.Address
}

// GetDomain gets domain where the agent is active
func (d DestinationAgentRootAccepted) GetDomain() *uint32 {
	return nil
}

// GetAgent gets the address of the agent
func (d DestinationAgentRootAccepted) GetAgent() *common.Address {
	return nil
}

// GetAttestation gets raw payload with attestation data
func (d DestinationAgentRootAccepted) GetAttestation() []byte {
	return nil
}

// GetAttSignature gets the notary signature for the attestation
func (d DestinationAgentRootAccepted) GetAttSignature() []byte {
	return nil
}

// GetAgentRoot gets the root
func (d DestinationAgentRootAccepted) GetAgentRoot() *[32]byte {
	return &(d.AgentRoot)
}

var _ destination.EventLog = &DestinationAgentRootAccepted{}
