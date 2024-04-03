package rfq

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
	// GetAddress returns the address of the RFQ relayer.
	GetAddress() string
	// GetRecipient returns the destination address of the RFQ transfer.
	GetRecipient() string
	// GetSender returns the sender of the RFQ transfer.
	GetSender() string
	// GetRequest returns the request info of the RFQ transfer.
	GetRequest() *[]byte
	// GetOriginChainID returns the chain id of the RFQ transfer.
	GetOriginChainID() *big.Int
	// GetDestinationChainID returns the chain id of the RFQ transfer.
	GetDestinationChainID() *big.Int
	// GetOriginToken returns the origin token of the RFQ transfer.
	GetOriginToken() *string
	// GetDestinationToken returns the destination token of the RFQ transfer.
	GetDestinationToken() *string
	// GetOriginAmount returns the origin amount of the RFQ transfer.
	GetOriginAmount() *big.Int
	// GetDestinationAmount returns the destination amount of the RFQ transfer.
	GetDestinationAmount() *big.Int
	// GetChainGasAmount returns the chain gas amount of the RFQ transfer.
	GetChainGasAmount() *big.Int
	// GetSendChainGas returns if the RFQ transfer will send gas to the recipient.
	GetSendChainGas() bool
}
