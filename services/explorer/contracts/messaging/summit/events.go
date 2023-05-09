package summit

import (
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/origin"
)

// GetEventType gets the execute event type.
func (o OriginSent) GetEventType() origin.EventType {
	return origin.SentEvent
}

// GetRaw gets the raw logs.
func (o OriginSent) GetRaw() ethTypes.Log {
	return o.Raw
}

// GetTxHash gets the unique identifier (txhash) for the event.
func (o OriginSent) GetTxHash() common.Hash {
	return o.Raw.TxHash
}

// GetEventIndex gets the event index.
func (o OriginSent) GetEventIndex() uint64 {
	return uint64(o.Raw.Index)
}

// GetBlockNumber gets the block number for the event.
func (o OriginSent) GetBlockNumber() uint64 {
	return o.Raw.BlockNumber
}

// GetContractAddress gets the contract address the event occurred on.
func (o OriginSent) GetContractAddress() common.Address {
	return o.Raw.Address
}

// GetMessageHash gets the message hash.
func (o OriginSent) GetMessageHash() [32]byte {
	return o.MessageHash
}

// GetDestination gets the destination.
func (o OriginSent) GetDestination() uint32 {
	return o.Destination
}

// GetMessage gets the Message.
func (o OriginSent) GetMessage() []byte {
	return o.Message
}

// GetNonce gets the Nonce.
func (o OriginSent) GetNonce() uint32 {
	return o.Nonce
}

var _ origin.EventLog = &OriginSent{}
