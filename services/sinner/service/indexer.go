package service

import (
	"context"
	"fmt"
	"github.com/jpillora/backoff"
	indexerconfig "github.com/synapsecns/sanguine/services/explorer/config/indexer"
	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	"github.com/synapsecns/sanguine/services/sinner/db"
	"github.com/synapsecns/sanguine/services/sinner/fetcher"
	"github.com/synapsecns/sanguine/services/sinner/logger"
	"github.com/synapsecns/sanguine/services/sinner/types"
	"golang.org/x/sync/errgroup"
	"time"
)

// ChainIndexer indexes message logs for a chain.
type ChainIndexer struct {
	// consumerDB is the database to store consumer data in.
	eventDB db.EventDB
	// parsers are the parsers for this chain.
	parsers types.Parsers
	// fetcher is the scribe fetcher.
	fetcher fetcher.ScribeFetcher
	// config is the config for the backfiller.
	config indexerConfig.ChainConfig
}

func NewChainIndexer(eventDB db.EventDB, parsers types.Parsers, fetcher fetcher.ScribeFetcher, config indexerConfig.ChainConfig) *ChainIndexer {
	chainIndexer := ChainIndexer{
		eventDB,
		parsers,
		fetcher,
		config,
	}
	return &chainIndexer
}

func (c ChainIndexer) Index(ctx context.Context) error {
	g, chainCtx := errgroup.WithContext(ctx)

	for i := range c.config.Contracts {
		contract := c.config.Contracts[i]
		contractType, err := indexerConfig.ContractTypeFromString(contract.ContractType)
		if err != nil {
			return  fmt.Errorf("could not create event parser for unknown contract type: %s", contract.ContractType)
		}

		var eventParser interface{}
		switch contractType {
		case indexerConfig.OriginType:
			eventParser = c.parsers.OriginParser
		case indexerConfig.ExecutionHubType:
			eventParser = c.parsers.DestinationParser

		}

		g.Go(func() error {
			b := &backoff.Backoff{
				Factor: 2,
				Jitter: true,
				Min:    1 * time.Second,
				Max:    3 * time.Second,
			}
			timeout := time.Duration(0)
			for {
				select {
				case <-chainCtx.Done():
					logger.ReportSinnerError(fmt.Errorf("sinner indexer %s on chain %d context cancelled: %v", contract.Address, c.config.ChainID, chainCtx.Err()), c.config.ChainID, logger.ContextCancelled)

					return fmt.Errorf("sinner indexer %s on chain %d context cancelled: %v", contract.Address, c.config.ChainID, chainCtx.Err())
				case <-time.After(timeout):
					err := c.indexLogs(chainCtx, contract)
					if err != nil {
						timeout = b.Duration()

						continue
					}
					b.Reset()
					timeout = time.Duration(1) * time.Second

				}
			}
		}
	}

}

