package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"

	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
)

// CCTPRelayerDBReader is the interface for reading from the database.
type CCTPRelayerDBReader interface {
	// GetLastBlockNumber gets the last block number that had a message for the respective origin chain in the database.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetMessagesByState gets all messages with the given state.
	GetMessagesByState(ctx context.Context, states ...types.MessageState) ([]types.Message, error)
	// GetMessageByOriginHash gets a message by its origin hash.
	GetMessageByOriginHash(ctx context.Context, originHash common.Hash) (*types.Message, error)
	// GetMessageByRequestID gets a message by its request id.
	GetMessageByRequestID(ctx context.Context, requestID string) (*types.Message, error)
}

// CCTPRelayerDBWriter is the interface for writing to the database.
type CCTPRelayerDBWriter interface {
	// StoreMessage stores a message in the database.
	StoreMessage(ctx context.Context, message types.Message) error
}

// CCTPRelayerDB is the interface for the database service.
type CCTPRelayerDB interface {
	CCTPRelayerDBReader
	CCTPRelayerDBWriter
	SubmitterDB() submitterDB.Service
}
