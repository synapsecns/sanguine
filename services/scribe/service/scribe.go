package service

import (
	"context"
	"fmt"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/scribe/backend"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/logger"
	otelMetrics "go.opentelemetry.io/otel/metric"
	"time"

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
// each contract on that chain. There is an indexer for livefilling all contracts and indexer for livefilling at the tip as well.
//
//nolint:cyclop
func (s Scribe) Start(ctx context.Context) error {
	g, groupCtx := errgroup.WithContext(ctx)
	b := backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    10 * time.Second,
	}
	retryRate := time.Second * 0
	for i := range s.config.Chains {
		chainConfig := s.config.Chains[i]
		chainID := chainConfig.ChainID

		// Run chain indexer for each chain
		g.Go(func() error {
			// Each chain gets its own context so it can retry on its own if there is a fatal error.
			// If the global scribe context fails, all chains will fail.
			chainCtx, cancelChain := context.WithCancel(ctx)
			defer cancelChain()
			for {
				select {
				case <-groupCtx.Done(): // Global context cancel, destroy all chain indexers.
					cancelChain() // redundant, but clean.
					return fmt.Errorf("global scribe context cancel %w", groupCtx.Err())
				case <-chainCtx.Done(): // Chain level context cancel, retry and recreate context.
					logger.ReportScribeError(fmt.Errorf("chain level scribe context cancel, %w", chainCtx.Err()), chainID, logger.ContextCancelled)
					chainCtx, cancelChain = context.WithCancel(ctx)
					retryRate = b.Duration()
					continue
				case <-time.After(retryRate):
					err := s.chainIndexers[chainID].Index(groupCtx)
					if err != nil {
						logger.ReportScribeError(fmt.Errorf("error running chain indexer %w", err), chainID, logger.FatalScribeError)
						retryRate = b.Duration()
						continue
					}
					cancelChain()
					return nil // This shouldn't really ever be hit
				}
			}
		})
	}
	if err := g.Wait(); err != nil {
		return fmt.Errorf("scribe failed: %w", err)
	}

	return nil
}
