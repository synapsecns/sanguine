package live

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/scribe/backfill"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// Scribe is a live scribe that logs all event data.
type Scribe struct {
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// clients is a mapping of chain IDs -> clients
	clients map[uint32]backfill.ScribeBackend
	// scribeBackfiller is the backfiller for the scribe
	scribeBackfiller *backfill.ScribeBackfiller
	// config is the config for the scribe
	config config.Config
}

// NewScribe creates a new scribe.
func NewScribe(eventDB db.EventDB, clients map[uint32]backfill.ScribeBackend, config config.Config) (*Scribe, error) {
	// initialize the scribe backfiller
	scribeBackfiller, err := backfill.NewScribeBackfiller(eventDB, clients, config)
	if err != nil {
		return nil, fmt.Errorf("could not create scribe backfiller: %w", err)
	}

	return &Scribe{
		eventDB:          eventDB,
		clients:          clients,
		scribeBackfiller: scribeBackfiller,
		config:           config,
	}, nil
}

// Start starts the scribe. This works by starting a backfill and recording what the
// current block, which it will backfill to. Then, each chain will listen for new block
// heights and backfill to that height.
func (s Scribe) Start(ctx context.Context) error {
	currentBlocks := make(map[uint32]uint64)

	// backfill each chain
	g, groupCtx := errgroup.WithContext(ctx)
	for _, chainConfig := range s.config.Chains {
		// capture func literal
		chainConfig := chainConfig
		// start the chain backfiller
		currentBlock, err := s.clients[chainConfig.ChainID].BlockNumber(groupCtx)
		currentBlocks[chainConfig.ChainID] = currentBlock
		if err != nil {
			return fmt.Errorf("could not get current block number: %w", err)
		}
		g.Go(func() error {
			err = s.scribeBackfiller.ChainBackfillers[chainConfig.ChainID].Backfill(groupCtx, currentBlock)
			if err != nil {
				return fmt.Errorf("could not backfill: %w", err)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	g, groupCtx = errgroup.WithContext(ctx)
	// start listening for new blocks and backfilling
	for _, chainConfig := range s.config.Chains {
		// capture func literal
		chainConfig := chainConfig
		g.Go(func() error {
			for {
				newBlock, err := s.clients[chainConfig.ChainID].BlockNumber(groupCtx)
				if err != nil {
					return fmt.Errorf("could not get current block number: %w", err)
				}
				if newBlock > currentBlocks[chainConfig.ChainID] {
					err = s.scribeBackfiller.ChainBackfillers[chainConfig.ChainID].Backfill(groupCtx, newBlock)
					if err != nil {
						return fmt.Errorf("could not backfill: %w", err)
					}
					currentBlocks[chainConfig.ChainID] = newBlock
				}
			}
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}
	fmt.Println("never")

	return nil
}
