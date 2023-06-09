package db

import "context"

// CCTPRelayerDBReader is the interface for reading from the database.
// TODO(dwasse): impl db interactions.
type CCTPRelayerDBReader interface {
	// GetLastBlockNumber gets the last block number that had a message in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
}

// Service is the interface for the database service.
type Service interface {
	CCTPRelayerDBReader
}
