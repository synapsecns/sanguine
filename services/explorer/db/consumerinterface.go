package db

import (
	"context"
	"github.com/benbjohnson/immutable"

	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"gorm.io/gorm"
)

// TODO simplify these interfaces so that there is one generic read function using an interface to accept a pointer from
// the caller so that the query's results are scanned into that pointer.

// ConsumerDBWriter is the interface for writing to the ConsumerDB.
type ConsumerDBWriter interface {
	// StoreEvent stores an event.
	StoreEvent(ctx context.Context, event interface{}) error
	// StoreEvents stores a list of events.
	StoreEvents(ctx context.Context, events []interface{}) error
	// StoreLastBlock stores the last block number that has been indexed for a given chain.
	StoreLastBlock(ctx context.Context, chainID uint32, blockNumber uint64, contractAddress string) error
	// StoreTokenIndex stores the token index data.
	StoreTokenIndex(ctx context.Context, chainID uint32, tokenIndex uint8, tokenAddress string, contractAddress string) error
	// StoreSwapFee stores the swap fee data.
	StoreSwapFee(ctx context.Context, chainID uint32, timestamp uint64, contractAddress string, fee uint64, feeType string) error
	// UNSAFE_DB gets the underlying gorm db. This is for testing only and not intended for use in production.
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
	// GetString gets a string out of the database
	GetString(ctx context.Context, query string) (string, error)
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
	// GetMVBridgeEvent returns a bridge event from the mv Table.
	GetMVBridgeEvent(ctx context.Context, query string) (*sql.HybridBridgeEvent, error)
	// GetAllBridgeEvents returns a bridge event.
	GetAllBridgeEvents(ctx context.Context, query string) ([]sql.HybridBridgeEvent, error)
	// GetAllMessageBusEvents returns a bridge event.
	GetAllMessageBusEvents(ctx context.Context, query string) ([]sql.HybridMessageBusEvent, error)
	// GetDateResults gets day by day data for a given query.
	GetDateResults(ctx context.Context, query string) ([]*model.DateResult, error)
	// GetAddressRanking gets AddressRanking for a given query.
	GetAddressRanking(ctx context.Context, query string) ([]*model.AddressRanking, error)
	// GetDailyTotals gets the daily stats for each date broken down by chain
	GetDailyTotals(ctx context.Context, query string) ([]*model.DateResultByChain, error)
	// GetRankedChainsByVolume gets the volume for each chain
	GetRankedChainsByVolume(ctx context.Context, query string) ([]*model.VolumeByChainID, error)
	// GetLastStoredBlock gets the last stored block for a given chain.
	GetLastStoredBlock(ctx context.Context, chainID uint32, contractAddress string) (uint64, error)
	// GetAddressData gets data for an address
	GetAddressData(ctx context.Context, query string) (float64, float64, int, error)
	// GetAddressDailyData gets daily data for a given address
	GetAddressDailyData(ctx context.Context, query string) ([]*model.AddressDailyCount, error)
	// GetAddressChainRanking gets ranking of an address's  chain activity
	GetAddressChainRanking(ctx context.Context, query string) ([]*model.AddressChainRanking, error)
	// GetLeaderboard gets the bridge leaderboard.
	GetLeaderboard(ctx context.Context, query string) ([]*model.Leaderboard, error)
	// GetPendingByChain gets the pending txs by chain.
	GetPendingByChain(ctx context.Context) (res *immutable.Map[int, int], err error)
	// GetBlockHeights gets the block heights for a given chain and contract type.
	GetBlockHeights(ctx context.Context, query string, contractTypeMap map[string]model.ContractType) ([]*model.BlockHeight, error)
}

// ConsumerDB is the interface for the ConsumerDB.
//
//go:generate go run github.com/vektra/mockery/v2 --name=ConsumerDB --output=mocks --case=underscore
type ConsumerDB interface {
	ConsumerDBWriter
	ConsumerDBReader
}
