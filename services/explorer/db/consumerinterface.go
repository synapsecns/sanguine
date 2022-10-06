package db

import (
	"context"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
	"github.com/synapsecns/sanguine/services/explorer/types/swap"
	"gorm.io/gorm"
)

// ConsumerDBWriter is the interface for writing to the ConsumerDB.
type ConsumerDBWriter interface {
	// StoreEvent stores an event.
	StoreEvent(ctx context.Context, bridgeEvent bridge.EventLog, swapEvent swap.EventLog, chainID uint32, tokenID *string, sender string) error
}

// ConsumerDBReader is the interface for reading events from the ConsumerDB.
type ConsumerDBReader interface {
	// GetAllChainIDs gets all chain IDs that have been used in bridge events.
	GetAllChainIDs(ctx context.Context) ([]uint32, error)
	// ReadBlockNumberByChainID reads an event from the database by chainID.
	ReadBlockNumberByChainID(ctx context.Context, eventType int8, chainID uint32) (*uint64, error)
	// BridgeEventCount returns the number of bridge events.
	BridgeEventCount(ctx context.Context, chainID uint32, address *string, tokenAddress *string, directionIn bool, firstBlock uint64) (count uint64, err error)
	// GetTokenAddressesByChainID gets all token addresses that have been used in bridge events for a given chain ID.
	GetTokenAddressesByChainID(ctx context.Context, chainID uint32) ([]string, error)
	// BridgeEventsFromIdentifiers returns events given identifiers.
	BridgeEventsFromIdentifiers(ctx context.Context, chainID *uint32, address, tokenAddress, kappa, txHash *string, page int) (partialInfos []*model.PartialInfo, err error)
	// GetSwapSuccess returns if an event had a successful swap.
	GetSwapSuccess(ctx context.Context, kappa string, chainID uint32) (*bool, error)
	// GetTxHashFromKappa returns the transaction hash for a given kappa.
	GetTxHashFromKappa(ctx context.Context, kappa string) (*string, error)
	// GetKappaFromTxHash returns the kappa for a given transaction hash.
	GetKappaFromTxHash(ctx context.Context, txHash string, chainID *uint32) (*string, error)
	// GetTransactionCountForEveryAddress gets the count of transactions (origin) for each address.
	GetTransactionCountForEveryAddress(ctx context.Context, subQuery string) ([]*model.AddressRanking, error)
	// DB gets the underlying gorm db.
	DB() *gorm.DB
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
