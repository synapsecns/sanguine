package db

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/sinner/db/model"
	graphqlModel "github.com/synapsecns/sanguine/services/sinner/graphql/server/graph/model"
	"github.com/synapsecns/sanguine/services/sinner/types"
	"gorm.io/gorm"
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
	// StoreOrUpdateMessageStatus stores/updates the status of a message.
	StoreOrUpdateMessageStatus(ctx context.Context, txHash string, messageID string, messageType types.MessageType) error
}

// EventDBReader is an interface for reading events from a database.
//
//nolint:interfacebloat
type EventDBReader interface {
	// RetrieveOriginSent gets the Origin Sent event with a message hash.
	RetrieveOriginSent(ctx context.Context, messageHash string) (model.OriginSent, error)
	// RetrieveOriginSents gets the Origin Sent events tied to a tx hash.
	RetrieveOriginSents(ctx context.Context, chainID uint32, txHash string) ([]model.OriginSent, error)
	// RetrieveExecuted gets the Executed event with a message hash.
	RetrieveExecuted(ctx context.Context, messageHash string) (model.Executed, error)
	// RetrieveExecuteds gets the Executed events tied to a tx hash.
	RetrieveExecuteds(ctx context.Context, chainID uint32, txHash string) ([]model.Executed, error)
	// RetrieveMessageStatus gets status of a message.
	RetrieveMessageStatus(ctx context.Context, messageHash string) (graphqlModel.MessageStatus, error)
	// RetrieveLastStoredBlock gets the last block stored in sinner.
	RetrieveLastStoredBlock(ctx context.Context, chainID uint32, address common.Address) (uint64, error)
}

// EventDB stores events.
type EventDB interface {
	EventDBWriter
	EventDBReader
}

type TestEventDB interface {
	EventDB
	// UNSAFE_DB gets the underlying gorm db. This is not intended for use in production.
	//
	//nolint:golint
	UNSAFE_DB() *gorm.DB
}
