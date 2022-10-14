package node

import (
	"context"
	"fmt"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/consumer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"golang.org/x/sync/errgroup"
	"time"
)

// Explorer is a live explorer that parses and stores all event data across all chains.
type Explorer struct {
	// consumerDB is the database to store event data in
	consumerDB db.ConsumerDB
	// explorerBackfiller is the backfiller to use to backfill the explorer
	explorerBackfiller *backfill.ExplorerBackfiller
	// config is the config for the explorer
	config config.Config
}

// NewExplorer creates a new explorer.
func NewExplorer(consumerDB db.ConsumerDB, config config.Config) (*Explorer, error) {
	// initialize the explorer backfiller
	explorerBackfiller, err := backfill.NewExplorerBackfiller(consumerDB, config)
	if err != nil {
		return nil, fmt.Errorf("could not create explorer backfiller: %w", err)
	}

	return &Explorer{
		consumerDB:         consumerDB,
		explorerBackfiller: explorerBackfiller,
		config:             config,
	}, nil
}

// Start starts the explorer. This works by starting a backfill from a predetermined block from the config, and then
// backfilling to the last block that is stored by Scribe on each chain. The last block stored by Scribe is gotten
// using the consumer's fetcher.
func (e Explorer) Start(ctx context.Context) error {
	refreshRate := e.config.RefreshRate
	if refreshRate == 0 {
		refreshRate = 1
	}
	// backfill each chain
	g, groupCtx := errgroup.WithContext(ctx)
	for i := range e.config.Chains {
		// capture the func literal
		chainConfig := e.config.Chains[i]
		g.Go(func() error {
			// backoff in case of an error
			b := &backoff.Backoff{
				Factor: 2,
				Jitter: true,
				Min:    1 * time.Second,
				Max:    30 * time.Second,
			}
			// timeout should always be 0 on the first attempt
			timeout := time.Duration(0)
			for {
				select {
				case <-groupCtx.Done():
					return fmt.Errorf("context canceled: %w", groupCtx.Err())
				case <-time.After(timeout):
					err := e.processRange(groupCtx, chainConfig.ChainID)
					if err != nil {
						timeout = b.Duration()
						continue
					}
					b.Reset()
					timeout = time.Duration(refreshRate) * time.Second
				}
			}
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("error in explorer: %w", err)
	}

	return nil
}

// nolint:gocognit
func (e Explorer) processRange(ctx context.Context, chainID uint32) error {
	// TODO add comments, take out log
	// TODO probably remove mapping from start blocks? since chainID is already decided.
	newBlock := e.config.Chains[chainID].StartBlocks[chainID]
	fmt.Println("newBlock", newBlock)

	err := e.explorerBackfiller.ChainBackfillers[chainID].Backfill(ctx)
	if err != nil {
		return fmt.Errorf("could not retrieve last confirmed block: %w", err)
	}

	lastBlockNumber, err := e.consumerDB.RetrieveLastBlock(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not retrieve last confirmed block: %w", err)
	}

	for i := lastBlockNumber; i <= newBlock; i++ {
		// check the validity of the block
		err = e.explorerBackfiller.ChainBackfillers[chainID].Backfill(ctx)
		if err != nil {
			return fmt.Errorf("could not backfill: %w", err)
		}
		// update the last confirmed block number
		err = e.consumerDB.StoreLastBlock(ctx, chainID, i)
		if err != nil {
			return fmt.Errorf("could not store last confirmed block: %w", err)
		}
	}

	return nil
}
