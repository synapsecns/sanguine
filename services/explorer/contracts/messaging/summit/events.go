package summit

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/summit"
	"math/big"
)

// GetEventType gets the execute event type.
func (s SummitReceiptAccepted) GetEventType() summit.EventType {
	return summit.ReceiptAcceptedEvent
}

// GetRaw gets the raw logs.
func (s SummitReceiptAccepted) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s SummitReceiptAccepted) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetEventIndex gets the event index.
func (s SummitReceiptAccepted) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

// GetBlockNumber gets the block number for the event.
func (s SummitReceiptAccepted) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SummitReceiptAccepted) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetDomain gets the domain of where the signed Notary/agent is active
func (s SummitReceiptAccepted) GetDomain() *uint32 {
	return &(s.Domain)
}

// GetAgent is the notary who signed the attestation
func (s SummitReceiptAccepted) GetAgent() *common.Address {
	return &(s.Notary)
}

// GetRcptPayload gets the raw payload with receipt data
func (s SummitReceiptAccepted) GetRcptPayload() []byte {
	return s.RcptPayload
}

// GetRcptSignature gets raw bytes of the message
func (s SummitReceiptAccepted) GetRcptSignature() []byte {
	return s.RcptSignature
}

// GetSnapshot gets raw payload with snapshot data
func (s SummitReceiptAccepted) GetSnapshot() []byte {
	return nil
}

// GetSnapSignature gets the agent signature for the snapshot
func (s SummitReceiptAccepted) GetSnapSignature() []byte {
	return nil
}

// GetTip gets the tip amount from the TipAwarded event.
func (s SummitReceiptAccepted) GetTip() *big.Int {
	return nil
}

var _ summit.EventLog = &SummitReceiptAccepted{}

// GetEventType gets the execute event type.
func (s SummitSnapshotAccepted) GetEventType() summit.EventType {
	return summit.SnapshotAcceptedEvent
}

// GetRaw gets the raw logs.
func (s SummitSnapshotAccepted) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s SummitSnapshotAccepted) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetEventIndex gets the event index.
func (s SummitSnapshotAccepted) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

// GetBlockNumber gets the block number for the event.
func (s SummitSnapshotAccepted) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SummitSnapshotAccepted) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetDomain gets the domain of where the signed Notary/agent is active
func (s SummitSnapshotAccepted) GetDomain() *uint32 {
	return &(s.Domain)
}

// GetAgent is the notary who signed the attestation
func (s SummitSnapshotAccepted) GetAgent() *common.Address {
	return &(s.Agent)
}

// GetRcptPayload gets the raw payload with receipt data
func (s SummitSnapshotAccepted) GetRcptPayload() []byte {
	return nil
}

// GetRcptSignature gets raw bytes of the message
func (s SummitSnapshotAccepted) GetRcptSignature() []byte {
	return nil
}

// GetSnapshot gets raw payload with snapshot data
func (s SummitSnapshotAccepted) GetSnapshot() []byte {
	return s.Snapshot
}

// GetSnapSignature gets the agent signature for the snapshot
func (s SummitSnapshotAccepted) GetSnapSignature() []byte {
	return s.SnapSignature
}

// GetTip gets the tip amount from the TipAwarded event.
func (s SummitSnapshotAccepted) GetTip() *big.Int {
	return nil
}

var _ summit.EventLog = &SummitSnapshotAccepted{}

// GetEventType gets the execute event type.
func (s SummitTipAwarded) GetEventType() summit.EventType {
	return summit.TipAwardedEvent
}

// GetRaw gets the raw logs.
func (s SummitTipAwarded) GetRaw() ethTypes.Log {
	return s.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (s SummitTipAwarded) GetTxHash() common.Hash {
	return s.Raw.TxHash
}

// GetEventIndex gets the event index.
func (s SummitTipAwarded) GetEventIndex() uint64 {
	return uint64(s.Raw.Index)
}

// GetBlockNumber gets the block number for the event.
func (s SummitTipAwarded) GetBlockNumber() uint64 {
	return s.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (s SummitTipAwarded) GetContractAddress() common.Address {
	return s.Raw.Address
}

// GetDomain gets the domain of where the signed Notary/agent is active
func (s SummitTipAwarded) GetDomain() *uint32 {
	return &(s.Domain)
}

// GetAgent is the notary who signed the attestation
func (s SummitTipAwarded) GetAgent() *common.Address {
	return &(s.Agent)
}

// GetRcptPayload gets the raw payload with receipt data
func (s SummitTipAwarded) GetRcptPayload() []byte {
	return nil
}

// GetRcptSignature gets raw bytes of the message
func (s SummitTipAwarded) GetRcptSignature() []byte {
	return nil
}

// GetSnapshot gets raw payload with snapshot data
func (s SummitTipAwarded) GetSnapshot() []byte {
	return s.Snapshot
}

// GetSnapSignature gets the agent signature for the snapshot
func (s SummitTipAwarded) GetSnapSignature() []byte {
	return s.SnapSignature
}

// GetTip gets the tip amount from the TipAwarded event.
func (s SummitTipAwarded) GetTip() *big.Int {
	return s.Tip
}

var _ summit.EventLog = &SummitTipAwarded{}
