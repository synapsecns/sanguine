package db

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
)

// EventDBWriter is an interface for writing events to a database.
//
//nolint:interfacebloat
type EventDBWriter interface {
	// StoreOriginSent stores an OriginSent event.
	StoreOriginSent(ctx context.Context, originSent *model.OriginSent) error
	// StoreExecuted stores an Executed event.
	StoreExecuted(ctx context.Context, executed *model.Executed) error
	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error
}

// EventDBReader is an interface for reading events from a database.
//
//nolint:interfacebloat
type EventDBReader interface {
	// RetrieveMessageStatus gets status of a message.
	RetrieveMessageStatus(ctx context.Context, txhash string) (response string, err error)
	// GetLastStoredBlock gets the last block stored in sinner.
	GetLastStoredBlock(ctx context.Context, chainID uint32, address common.Address) (uint64, error)
}

// EventDB stores events.
//
//go:generate go run github.com/vektra/mockery/v2 --name EventDB --output ./mocks --case=underscore
type EventDB interface {
	EventDBWriter
	EventDBReader
}
