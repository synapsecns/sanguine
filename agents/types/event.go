package types

import (
	"math/big"
)

// CrossChainEventLog is the common cross chain event log for user and bridge events.
//go:generate go run github.com/vektra/mockery/v2 --name CrossChainEventLog --output ./mocks --case=underscore
type CrossChainEventLog interface {
	// GetToken gets the token of the log
	GetToken() string
	// GetAmount gets the amount from fee
	GetAmount() *big.Int
	// GetEventType gets the event type
	GetEventType() EventType
	// GetIdentifier gets the unique identifier for a transaction.
	//
	// For example: In the case of eth, this is a txhash. For terra, it's a deposit identifier generated at runtime
	GetIdentifier() string

	// GetBlockNumber gets the block number an event occurred
	GetBlockNumber() uint64
	// GetContractAddress gets the contract address an event occurred at
	GetContractAddress() string
}

// CrossChainUserEventLog is a common interface for getting data from cross chain event logs initiated by the user
//go:generate go run github.com/vektra/mockery/v2 --name CrossChainUserEventLog --output ./mocks --case=underscore
type CrossChainUserEventLog interface {
	CrossChainEventLog
	// GetDestinationChainID gets the destination chain id from the event log
	GetDestinationChainID() *big.Int
}

// CrossChainBridgeEventLog is a common interface for getting data from cross chain event logs initiated by
// the bridge.
//go:generate go run github.com/vektra/mockery/v2 --name CrossChainBridgeEventLog --output ./mocks --case=underscore
type CrossChainBridgeEventLog interface {
	CrossChainEventLog
	// GetFee gets the fee assessed via CalculateSwapFee()
	GetFee() *big.Int
	// GetEventType gets the event type
	GetEventType() EventType
	// GetKappa gets the kappa value for the event
	GetKappa() [32]byte
}

// CrossChainDepositLog is a cross chain deposit log.
type CrossChainDepositLog interface {
	CrossChainUserEventLog
	// GetRecipient gets the recipient of the deposit
	GetRecipient() string
}
