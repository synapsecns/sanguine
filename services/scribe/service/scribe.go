package service

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	otelMetrics "go.opentelemetry.io/otel/metric"

	"golang.org/x/sync/errgroup"
)

// Scribe is a live scribe that logs all event data.
type Scribe struct {
	// eventDB is the database to store event data in.
	eventDB db.EventDB
	// clients is a mapping of chain IDs -> clients.
	clients map[uint32][]backend.ScribeBackend
	// chainIndexers are the indexers for the scribe.
	chainIndexers map[uint32]*ChainIndexer
	// config is the config for the scribe.
	config config.Config
	// handler is the metrics handler for the scribe.
	handler metrics.Handler
	// reorgMeters holds a otel counter meter for reorgs for each chain
	reorgMeters map[uint32]otelMetrics.Int64Counter
}

// NewScribe creates a new scribe.
func NewScribe(eventDB db.EventDB, clients map[uint32][]backend.ScribeBackend, config config.Config, handler metrics.Handler) (*Scribe, error) {
	chainIndexers := make(map[uint32]*ChainIndexer)
	for i := range config.Chains {
		chainConfig := config.Chains[i]
		chainIndexer, err := NewChainIndexer(eventDB, clients[chainConfig.ChainID], chainConfig, handler)
		if err != nil {
			return nil, fmt.Errorf("could not create chain indexer: %w", err)
		}
		chainIndexers[chainConfig.ChainID] = chainIndexer
	}

	return &Scribe{
		eventDB:       eventDB,
		clients:       clients,
		chainIndexers: chainIndexers,
		config:        config,
		handler:       handler,
		reorgMeters:   make(map[uint32]otelMetrics.Int64Counter),
	}, nil
}

// Start starts the scribe. A chain indexer is spun up for each chain, and a indexer is spun up for
// each contract on that chain. There is an indexer for livefillingall contracts and indexer for livefilling at the tip as well.
//
//nolint:cyclop
func (s Scribe) Start(ctx context.Context) error {
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range s.config.Chains {
		chainConfig := s.config.Chains[i]
		chainID := chainConfig.ChainID

		// Livefill the chains
		g.Go(func() error {
			err := s.chainIndexers[chainID].Index(groupCtx, nil)
			if err != nil {
				return fmt.Errorf("could not index: %w", err)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("livefill failed: %w", err)
	}

	return nil
}
