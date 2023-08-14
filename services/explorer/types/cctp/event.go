package cctp

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
	// GetRequestID returns the request id of the CCTP transfer.
	GetRequestID() [32]byte
	// GetToken returns the address of the received token.
	GetToken() string
	// GetAmount returns the amount of the CCTP transfer.
	GetAmount() *big.Int
	// GetOriginChainID returns the chain id of the CCTP transfer.
	GetOriginChainID() *big.Int
	// GetDestinationChainID returns the chain id of the CCTP transfer.
	GetDestinationChainID() *big.Int
	// GetSender returns the sender of the CCTP transfer.
	GetSender() *string
	// GetNonce returns the nonce of the CCTP transfer.
	GetNonce() *uint64
	// GetMintToken returns the mint token of the CCTP transfer.
	GetMintToken() *string

	// GetRequestVersion returns the request version of the CCTP transfer.
	GetRequestVersion() *uint32
	// GetFormattedRequest returns the formatted request of the CCTP transfer.
	GetFormattedRequest() *[]byte
	// GetRecipient returns the receipient of the CCTP transfer.
	GetRecipient() *string
	// GetFee returns the fee of the CCTP transfer.
	GetFee() *big.Int
}
