package backfill

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/synapse-node/pkg/evm/client"
	"golang.org/x/sync/errgroup"
)

// ScribeBackfiller is a backfiller that aggregates all backfilling from ChainBackfillers.
type ScribeBackfiller struct {
	// eventDB is the database to store event data in
	eventDB db.EventDB
	// clients are a mapping of chain IDs -> clients
	clients map[uint32]client.EVMClient
	// chainBackfillers is the list of chain backfillers
	chainBackfillers []*ChainBackfiller
	// config is the config for the backfiller
	config config.Config
}

// NewScribeBackfiller creates a new backfiller for the scribe.
func NewScribeBackfiller(eventDB db.EventDB, clients []client.EVMClient, config config.Config) (*ScribeBackfiller, error) {
	// set up the clients mapping
	clientsMap := make(map[uint32]client.EVMClient)
	for _, client := range clients {
		chainID, err := client.ChainID(context.Background())
		if err != nil {
			return nil, fmt.Errorf("could not get chain ID: %w", err)
		}
		clientsMap[uint32(chainID.Uint64())] = client
	}
	// initialize the list of chain backfillers
	chainBackfillers := []*ChainBackfiller{}
	// initialize each chain backfiller
	for _, chainConfig := range config.Chains {
		chainBackfiller, err := NewChainBackfiller(chainConfig.ChainID, eventDB, clientsMap[chainConfig.ChainID], chainConfig)
		if err != nil {
			return nil, fmt.Errorf("could not create chain backfiller: %w", err)
		}
		chainBackfillers = append(chainBackfillers, chainBackfiller)
	}

	return &ScribeBackfiller{
		eventDB:          eventDB,
		clients:          clientsMap,
		chainBackfillers: chainBackfillers,
		config:           config,
	}, nil
}

// Backfill iterates over each chain backfiller and calls Backfill concurrently on each one.
func (s ScribeBackfiller) Backfill(ctx context.Context) error {
	// initialize the errgroup
	g, ctx := errgroup.WithContext(ctx)

	// iterate over each chain backfiller
	for _, chainBackfiller := range s.chainBackfillers {
		// capture func literal
		chainBackfiller := chainBackfiller
		// get the end height for the backfill
		currentBlock, err := s.clients[chainBackfiller.chainID].BlockNumber(ctx)
		if err != nil {
			return fmt.Errorf("could not get current block number: %w", err)
		}
		confirmationThreshold := s.config.Chains[chainBackfiller.chainID].ConfirmationThreshold
		if uint32(currentBlock) < confirmationThreshold {
			return fmt.Errorf("current block number %d is less than confirmation threshold %d", currentBlock, s.config.Chains[chainBackfiller.chainID].ConfirmationThreshold)
		}
		// call Backfill concurrently
		g.Go(func() error {
			err := chainBackfiller.Backfill(ctx, currentBlock-uint64(confirmationThreshold))
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
