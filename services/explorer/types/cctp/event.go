package cctp

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// EventLog is the interface for all cctp events.
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
	// GetRequestID returns the request id of the CCTP transfer.
	GetRequestID() [32]byte
	// GetOriginChainID returns the chain id of the CCTP transfer.
	GetOriginChainID() *big.Int
	// GetDestinationChainID returns the chain id of the CCTP transfer.
	GetDestinationChainID() *big.Int
	// GetSender returns the sender of the CCTP transfer.
	GetSender() *string
	// GetNonce returns the nonce of the CCTP transfer.
	GetNonce() *uint64
	// GetBurnToken returns the burn token of the CCTP transfer.
	GetBurnToken() *string
	// GetMintToken returns the mint token of the CCTP transfer.
	GetMintToken() *string
	// GetSentAmount returns the sent amount of the CCTP transfer.
	GetSentAmount() *big.Int
	// GetReceivedAmount returns the received amount of the CCTP transfer.
	GetReceivedAmount() *big.Int
	// GetRequestVersion returns the request version of the CCTP transfer.
	GetRequestVersion() *uint32
	// GetFormattedRequest returns the formatted request of the CCTP transfer.
	GetFormattedRequest() []byte
	// GetRecipient returns the receipient of the CCTP transfer.
	GetRecipient() *string
	// GetFee returns the fee of the CCTP transfer.
	GetFee() *big.Int
	// GetToken returns the address of the received token.
	GetToken() *string
}
