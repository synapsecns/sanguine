package db

import (
	"context"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"gorm.io/gorm"
)

// ConsumerDBWriter is the interface for writing to the ConsumerDB.
type ConsumerDBWriter interface {
	// StoreEvent stores an event.
	StoreEvent(ctx context.Context, event interface{}) error
	// StoreLastBlock stores the last block number that has been backfilled for a given chain.
	StoreLastBlock(ctx context.Context, chainID uint32, blockNumber uint64) error
	// UNSAFE_DB gets the underlying gorm db. This is not intended for use in production.
	//
	//nolint:golint
	UNSAFE_DB() *gorm.DB
}

// ConsumerDBReader is the interface for reading events from the ConsumerDB.
// nolint:interfacebloat
type ConsumerDBReader interface {
	// GetUint64 gets a uint64 for a given query.
	GetUint64(ctx context.Context, query string) (uint64, error)
	// GetFloat64 gets a float64 out of the database
	GetFloat64(ctx context.Context, query string) (float64, error)
	// GetStringArray gets an array of strings from a given query.
	GetStringArray(ctx context.Context, query string) ([]string, error)
	// GetBridgeEvent returns a bridge event.
	GetBridgeEvent(ctx context.Context, query string) (*sql.BridgeEvent, error)
	// GetDateResults gets day by day data for a given query.
	GetDateResults(ctx context.Context, query string) ([]*model.DateResult, error)
	// GetAddressRanking gets AddressRanking for a given query.
	GetAddressRanking(ctx context.Context, query string) ([]*model.AddressRanking, error)

	// GetAllChainIDs gets all chain IDs that have been used in bridge events.
	GetAllChainIDs(ctx context.Context) ([]int, error)
	// PartialInfosFromIdentifiers returns events given identifiers.
	PartialInfosFromIdentifiers(ctx context.Context, query string) ([]*model.PartialInfo, error)
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
