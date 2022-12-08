package db

import (
	"context"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
)

// ExecutorDBWriter is the interface for writing to the executor database.
type ExecutorDBWriter interface {
	// StoreMessage stores a message in the database.
	StoreMessage(ctx context.Context, message types.DBMessage) error
}

// ExecutorDBReader is the interface for reading from the executor database.
type ExecutorDBReader interface {
	// GetMessage gets a message from the database.
	GetMessage(ctx context.Context, messageMask types.DBMessage) (*types.DBMessage, error)
	// GetLastBlockNumber gets the last block number that had a message in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
}

// ExecutorDB is the interface for the executor database.
type ExecutorDB interface {
	ExecutorDBWriter
	ExecutorDBReader
}
