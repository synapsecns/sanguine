package backfill

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/ethergo/backends/simulated"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
	"math/big"
)

// ScribeBackfiller is a backfiller that aggregates all backfilling from ChainBackfillers.
type ScribeBackfiller struct {
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// clients is a mapping of chain IDs -> clients
	clients map[uint32][]ScribeBackend
	// ChainBackfillers is a mapping of chain IDs -> chain backfillers
	ChainBackfillers map[uint32]*ChainBackfiller
	// config is the config for the backfiller
	config config.Config
}

// NewScribeBackfiller creates a new backfiller for the scribe.
func NewScribeBackfiller(eventDB db.EventDB, clientsMap map[uint32][]ScribeBackend, config config.Config) (*ScribeBackfiller, error) {
	// initialize the list of chain backfillers
	chainBackfillers := map[uint32]*ChainBackfiller{}
	// initialize each chain backfiller
	for _, chainConfig := range config.Chains {
		chainBackfiller, err := NewChainBackfiller(chainConfig.ChainID, eventDB, clientsMap[chainConfig.ChainID], chainConfig)
		if err != nil {
			return nil, fmt.Errorf("could not create chain backfiller: %w", err)
		}
		chainBackfillers[chainConfig.ChainID] = chainBackfiller
	}

	return &ScribeBackfiller{
		eventDB:          eventDB,
		clients:          clientsMap,
		ChainBackfillers: chainBackfillers,
		config:           config,
	}, nil
}

// Backfill iterates over each chain backfiller and calls Backfill concurrently on each one.
func (s ScribeBackfiller) Backfill(ctx context.Context) error {
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)

	// iterate over each chain backfiller
	for _, chainBackfiller := range s.ChainBackfillers {
		// capture func literal
		chainBackfiller := chainBackfiller
		// call Backfill concurrently
		g.Go(func() error {
			err := chainBackfiller.Backfill(groupCtx, false)
			if err != nil {
				return fmt.Errorf("could not backfill chain: %w", err)
			}
			return nil
		})
	}
	// wait for all of the backfillers to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	return nil
}

// ScribeBackend is the set of functions that the scribe needs from a client.
type ScribeBackend interface {
	// ChainID gets the chain id from the rpc server
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockByNumber retrieves a block from the database by number, caching it
	// (associated with its hash) if found.
	BlockByNumber(ctx context.Context, number *big.Int) (*types.Block, error)
	// TransactionByHash checks the pool of pending transactions in addition to the
	// blockchain. The isPending return value indicates whether the transaction has been
	// mined yet. Note that the transaction may not be part of the canonical chain even if
	// it's not pending.
	TransactionByHash(ctx context.Context, txHash common.Hash) (tx *types.Transaction, isPending bool, err error)
	// TransactionReceipt returns the receipt of a mined transaction. Note that the
	// transaction may not be included in the current canonical chain even if a receipt
	// exists.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
	// BlockNumber gets the latest block number
	BlockNumber(ctx context.Context) (uint64, error)
	// FilterLogs executes a log filter operation, blocking during execution and
	// returning all the results in one batch.
	//
	// TODO(karalabe): Deprecate when the subscription one can return past data too.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	// HeaderByNumber returns the block header with the given block number.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
}

var _ ScribeBackend = simulated.Backend{}
