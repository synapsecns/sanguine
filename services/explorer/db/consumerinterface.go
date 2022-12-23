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
	// StoreEvents stores a list of events.
	StoreEvents(ctx context.Context, events []interface{}) error
	// StoreLastBlock stores the last block number that has been backfilled for a given chain.
	StoreLastBlock(ctx context.Context, chainID uint32, blockNumber uint64, contractAddress string) error
	// StoreTokenIndex stores the token index data.
	StoreTokenIndex(ctx context.Context, chainID uint32, tokenIndex uint8, tokenAddress string, contractAddress string) error
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
	// GetTxCounts gets the counts for each of tx_hash from a given query.
	GetTxCounts(ctx context.Context, query string) ([]*model.TransactionCountResult, error)
	// GetTokenCounts gets the counts for each of token address from a given query.
	GetTokenCounts(ctx context.Context, query string) ([]*model.TokenCountResult, error)
	// GetBridgeEvent returns a bridge event.
	GetBridgeEvent(ctx context.Context, query string) (*sql.BridgeEvent, error)
	// GetBridgeEvents returns a bridge event.
	GetBridgeEvents(ctx context.Context, query string) ([]sql.BridgeEvent, error)
	// GetAllBridgeEvents returns a bridge event.
	GetAllBridgeEvents(ctx context.Context, query string) ([]sql.HybridBridgeEvent, error)
	// GetDateResults gets day by day data for a given query.
	GetDateResults(ctx context.Context, query string) ([]*model.DateResult, error)
	// GetAddressRanking gets AddressRanking for a given query.
	GetAddressRanking(ctx context.Context, query string) ([]*model.AddressRanking, error)

	// PartialInfosFromIdentifiers returns events given identifiers.
	PartialInfosFromIdentifiers(ctx context.Context, query string) ([]*model.PartialInfo, error)
	// PartialInfosFromIdentifiersByChain returns events given identifiers.
	PartialInfosFromIdentifiersByChain(ctx context.Context, query string) (map[int]*model.PartialInfo, error)
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
