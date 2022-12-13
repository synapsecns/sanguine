package backfill

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"golang.org/x/sync/errgroup"
)

// ScribeBackfiller is a backfiller that aggregates all backfilling from ChainBackfillers.
type ScribeBackfiller struct {
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// clients is a mapping of chain IDs -> clients.
	clients map[uint32][]ScribeBackend
	// ChainBackfillers is a mapping of chain IDs -> chain backfillers.
	ChainBackfillers map[uint32]*ChainBackfiller
	// config is the config for the backfiller.
	config config.Config
}

// NewScribeBackfiller creates a new backfiller for the scribe.
func NewScribeBackfiller(eventDB db.EventDB, clientsMap map[uint32][]ScribeBackend, config config.Config) (*ScribeBackfiller, error) {
	chainBackfillers := map[uint32]*ChainBackfiller{}

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
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range s.ChainBackfillers {
		chainBackfiller := s.ChainBackfillers[i]
		g.Go(func() error {
			LogEvent(InfoLevel, "Scribe backfilling chain", LogData{"cid": chainBackfiller.chainID})
			err := chainBackfiller.Backfill(groupCtx, nil)
			if err != nil {
				return fmt.Errorf("could not backfill chain: %w", err)
			}

			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	return nil
}
