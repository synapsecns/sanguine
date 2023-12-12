package db

import (
	"context"

	submitterDB "github.com/synapsecns/sanguine/ethergo/submitter/db"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db/model"

	"github.com/ethereum/go-ethereum/core/types"

	"gorm.io/gorm"
)

// DBWriter is an interface for writing events to a database.
type DBWriter interface {
	StoreOriginBridgeEvent(ctx context.Context, chainID uint32, log *types.Log, event *bindings.FastBridgeBridgeRequested) error
	StoreDestinationBridgeEvent(ctx context.Context, log *types.Log, originEvent *model.OriginBridgeEvent) error
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error
	StoreToken(ctx context.Context, token *model.Token) error
	StoreDeadlineQueueEvent(ctx context.Context, entry *model.DeadlineQueue) error
	RemoveDeadlineQueueEvent(ctx context.Context, eventID string) error
}

// DBReader is an interface for reading events from a database.
type DBReader interface {
	GetOriginBridgeEvent(ctx context.Context, transactionID string) (*model.OriginBridgeEvent, error)
	GetDestinationBridgeEvent(ctx context.Context, transactionID string) (*model.DestinationBridgeEvent, error)
	GetLastIndexed(ctx context.Context, chainID uint32, address common.Address) (uint64, error)
	GetToken(ctx context.Context, tokenID string) (*model.Token, error)
	GetDeadlineQueueEvents(ctx context.Context) ([]*model.DeadlineQueue, error)
}

// DB stores events.
type DB interface {
	DBWriter
	DBReader
	SubmitterDB() submitterDB.Service
}

type TestDB interface {
	DB
	// UNSAFE_DB gets the underlying gorm db. This is not intended for use in production.
	//
	//nolint:golint
	UNSAFE_DB() *gorm.DB
}
