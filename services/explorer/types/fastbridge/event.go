package fastbridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// EventLog is the interface for all cctp events.
//
//nolint:interfacebloat
type EventLog interface {
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.
	GetBlockNumber() uint64
	// GetEventType returns the event type of the log.
	GetEventType() EventType
	// GetEventIndex returns the index of the log.
	GetEventIndex() uint64
	// GetTransactionID returns the transaction id of the RFQ transfer.
	GetTransactionID() [32]byte
	// GetRelayer returns the address of the RFQ relayer.
	GetRelayer() *string
	// GetRecipient returns the Dest address of the RFQ transfer.
	GetTo() *string
	// GetSender returns the sender of the RFQ transfer.
	GetSender() *string
	// GetRequest returns the request info of the RFQ transfer.
	GetRequest() *[]byte
	// GetOriginChainID returns the chain id of the RFQ transfer.
	GetOriginChainID() *uint32
	// GetDestChainID returns the chain id of the RFQ transfer.
	GetDestChainID() *uint32
	// GetOriginToken returns the origin token of the RFQ transfer.
	GetOriginToken() common.Address
	// GetDestToken returns the Dest token of the RFQ transfer.
	GetDestToken() common.Address
	// GetOriginAmount returns the origin amount of the RFQ transfer.
	GetOriginAmount() *big.Int
	// GetDestAmount returns the Dest amount of the RFQ transfer.
	GetDestAmount() *big.Int
	// GetChainGasAmount returns the chain gas amount of the RFQ transfer.
	GetChainGasAmount() *big.Int
	// GetSendChainGas returns if the RFQ transfer will send gas to the recipient.
	GetSendChainGas() *bool
}
