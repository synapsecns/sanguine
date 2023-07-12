package indexer

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"
)

// GetLogs exports logs for testing.
func (x Indexer) GetLogs(ctx context.Context, startHeight, endHeight uint64) (<-chan types.Log, <-chan string) {
	return x.getLogs(ctx, startHeight, endHeight)
}

// IndexerConfig exports the indexers config for testing.
func (x Indexer) IndexerConfig() IndexerConfig {
	return x.indexerConfig
}
