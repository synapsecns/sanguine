package db

import (
	"context"

	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
)

// CCTPRelayerDBReader is the interface for reading from the database.
type CCTPRelayerDBReader interface {
	// GetLastBlockNumber gets the last block number that had a message in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
}

// CCTPRelayerDBWriter is the interface for writing to the database.
type CCTPRelayerDBWriter interface {
	// StoreMessage stores a message in the database.
	StoreMessage(ctx context.Context, message types.Message)
}

// CCTPRelayerDB is the interface for the database service.
type CCTPRelayerDB interface {
	CCTPRelayerDBReader
}
