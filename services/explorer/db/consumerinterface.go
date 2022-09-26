package db

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"gorm.io/gorm"
)

// ConsumerDBWriter is the interface for writing to the ConsumerDB.
type ConsumerDBWriter interface {
	// StoreEvent stores an event.
	StoreEvent(ctx context.Context, bridgeEvent bridge.EventLog, swapEvent swap.EventLog, chainID uint32, tokenID *string) error
	// StoreFailedLog stores a log that was failed to be parsed and stored.
	StoreFailedLog(ctx context.Context, log types.Log, chainID uint32) error
	// StoreLastLoggedBlock stores the last confirmed block for a chain.
	StoreLastLoggedBlock(ctx context.Context, chainID uint32, blockNumber uint64) error
	// DeleteFailedLog deletes a failed log from the table.
	DeleteFailedLog(ctx context.Context, log types.Log, chainID uint32) error
}

// ConsumerDBReader is the interface for reading events from the ConsumerDB.
type ConsumerDBReader interface {
	// ReadBlockNumberByChainID reads an event from the database by chainID.
	ReadBlockNumberByChainID(ctx context.Context, eventType int8, chainID uint32) (*uint64, error)
	// RetrieveFailedLogs gets all failed logs.
	RetrieveFailedLogs(ctx context.Context, chainID uint32) (logs []types.Log, err error)
	// RetrieveLastLoggedBlock retrieves the last block number that has been confirmed.
	RetrieveLastLoggedBlock(ctx context.Context, chainID uint32) (uint64, error)
	// DB gets the underlying gorm db.
	DB() *gorm.DB
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
