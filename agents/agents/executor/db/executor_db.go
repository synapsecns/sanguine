package db

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/agents/agents/executor/types"
	agentsTypes "github.com/synapsecns/sanguine/agents/types"
)

// ExecutorDBWriter is the interface for writing to the executor database.
type ExecutorDBWriter interface {
	// StoreMessage stores a message in the database.
	StoreMessage(ctx context.Context, message agentsTypes.Message, root common.Hash, blockNumber uint64) error
}

// ExecutorDBReader is the interface for reading from the executor database.
type ExecutorDBReader interface {
	// GetMessage gets a message from the database.
	GetMessage(ctx context.Context, messageMask types.DBMessage) (*agentsTypes.Message, error)
	// GetMessages gets messages from the database, paginated and ordered in ascending order by nonce.
	GetMessages(ctx context.Context, messageMask types.DBMessage, page int) ([]agentsTypes.Message, error)
	// GetRoot gets the root of a message from the database.
	GetRoot(ctx context.Context, messageMask types.DBMessage) (common.Hash, error)
	// GetLastBlockNumber gets the last block number that had a message in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
}

// ExecutorDB is the interface for the executor database.
type ExecutorDB interface {
	ExecutorDBWriter
	ExecutorDBReader
}
