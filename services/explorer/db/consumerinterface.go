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
	StoreEvent(ctx context.Context, bridgeEvent *sql.BridgeEvent, swapEvent *sql.SwapEvent) error
	// UNSAFE_DB gets the underlying gorm db. This is not intended for use in production.
	//
	//nolint:golint
	UNSAFE_DB() *gorm.DB
}

// ConsumerDBReader is the interface for reading events from the ConsumerDB.
// nolint:interfacebloat
type ConsumerDBReader interface {
	// GetAllChainIDs gets all chain IDs that have been used in bridge events.
	GetAllChainIDs(ctx context.Context) ([]uint32, error)
	// GetSwapSuccess returns if an event had a successful swap.
	GetSwapSuccess(ctx context.Context, kappa string, chainID uint32) (*bool, error)
	// GetTxHashFromKappa returns the transaction hash for a given kappa.
	GetTxHashFromKappa(ctx context.Context, kappa string) (*string, error)
	// ReadBlockNumberByChainID reads an event from the database by chainID.
	ReadBlockNumberByChainID(ctx context.Context, eventType int8, chainID uint32) (*uint64, error)
	// BridgeEventCount returns the number of bridge events.
	BridgeEventCount(ctx context.Context, query string) (count uint64, err error)
	// GetTokenAddressesByChainID gets all token addresses that have been used in bridge events for a given chain ID.
	GetTokenAddressesByChainID(ctx context.Context, query string) ([]string, error)
	// PartialInfosFromIdentifiers returns events given identifiers.
	PartialInfosFromIdentifiers(ctx context.Context, query string) (partialInfos []*model.PartialInfo, err error)
	// GetKappaFromTxHash returns the kappa for a given transaction hash.
	GetKappaFromTxHash(ctx context.Context, query string) (*string, error)
	// GetTransactionCountForEveryAddress gets the count of transactions (origin) for each address.
	GetTransactionCountForEveryAddress(ctx context.Context, subQuery string) ([]*model.AddressRanking, error)
	// GetBridgeStatistic gets bridge statistics
	GetBridgeStatistic(ctx context.Context, subQuery string) (*string, error)
	// GetHistoricalData gets bridge historical data
	GetHistoricalData(ctx context.Context, subQuery string, typeArg *model.HistoricalResultType, filter string) (*model.HistoricalResult, error)
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
