package message

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// MessageLog is the interface for all message events.
//
//nolint:interfacebloat
type MessageLog interface {
	// GetContractAddress returns the contract address of the log.
	GetContractAddress() common.Address
	// GetBlockNumber returns the block number of the log.'
	GetBlockNumber() uint64
	// GetTxHash returns the transaction hash of the log.
	GetTxHash() common.Hash
	// GetEventType returns the event type of the log.
	GetEventType() EventType
	// GetEventIndex returns the index of the log.
	GetEventIndex() uint64

	// GetMessageId returns the message id of the event.
	GetMessageId() *[32]byte
	// GetStatus returns the status of the event.
	GetStatus() *string
	// GetSourceAddress gets the address that the message will be passed from.
	GetSourceAddress() *common.Address
	// GetDestinationAddress gets the address that the message will be passed to.
	GetDestinationAddress() *common.Address
	// GetSourceChainID returns the chain id of the message's source chain.
	GetSourceChainID() *big.Int
	// GetDestinationChainID returns the chain id of the message's destination chain.
	GetDestinationChainID() *big.Int
	// GetGasLimit returns the gas limit to be passed alongside the message, depending on the fee paid on the source chain.
	GetGasLimit() *big.Int
	// GetSourceNonce returns the source nonce of the message.
	GetSourceNonce() *big.Int
	// GetNonce returns the nonce of the message.
	GetNonce() *big.Int
	// GetMessage gets the message.
	GetMessage() *[]byte
	// GetReceiver returns the receiver of the event.
	GetReceiver() *[32]byte
	// GetOptions gets the message.
	GetOptions() *[]byte
	// GetFee returns the fee of the message.
	GetFee() *big.Int
}
