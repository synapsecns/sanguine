package live

import (
	"context"
	"fmt"
	"time"

	"github.com/jpillora/backoff"
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
//
//nolint:cyclop
func (s Scribe) Start(ctx context.Context) error {
	refreshRate := s.config.RefreshRate
	if refreshRate == 0 {
		refreshRate = 1
	}
	// backfill each chain
	g, groupCtx := errgroup.WithContext(ctx)
	for _, chainConfig := range s.config.Chains {
		// capture func literal
		chainConfig := chainConfig
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
					return groupCtx.Err()
				case <-time.After(timeout):
					err := s.processRange(groupCtx, chainConfig.ChainID)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not get current block number: %v", err)
						continue
					}
					b.Reset()
					timeout = time.Duration(refreshRate) * time.Second
				}
			}
		})
	}

	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}

	return nil
}

func (s Scribe) processRange(ctx context.Context, chainID uint32) error {
	newBlock, err := s.clients[chainID].BlockNumber(ctx)
	if err != nil {
		return fmt.Errorf("could not get current block number: %w", err)
	}

	err = s.scribeBackfiller.ChainBackfillers[chainID].Backfill(ctx, newBlock)
	if err != nil {
		return fmt.Errorf("could not backfill: %w", err)
	}
	return nil
}
