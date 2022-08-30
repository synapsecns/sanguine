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

// Start starts the scribe.
func (s Scribe) Start(ctx context.Context) error {
	for _, chainConfig := range s.config.Chains {
		g, ctx := errgroup.WithContext(ctx)
		// start the chain backfiller
		g.Go(func() error {
			currentBlock, err := s.clients[chainConfig.ChainID].BlockNumber(ctx)
			if err != nil {
				return fmt.Errorf("could not get current block number: %w", err)
			}
			err = s.scribeBackfiller.chainBackfillers[chainConfig.ChainID].Backfill(ctx, currentBlock)
		})

	}

	// for each chain, get the current block
	currentBlocks := make(map[uint32]uint64)
	for _, chainConfig := range s.config.Chains {
		currentBlock, err := s.clients[chainConfig.ChainID].BlockNumber(ctx)
		if err != nil {
			return fmt.Errorf("could not get current block number %w", err)
		}
		currentBlocks[chainConfig.ChainID] = currentBlock
	}

	return nil
}
