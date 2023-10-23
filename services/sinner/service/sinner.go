package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/sinner/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/contracts/origin"
	"github.com/synapsecns/sanguine/services/sinner/logger"

	"github.com/ethereum/go-ethereum/common"
	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	"github.com/synapsecns/sanguine/services/sinner/db"
	fetcherpkg "github.com/synapsecns/sanguine/services/sinner/fetcher"
	gqlClient "github.com/synapsecns/sanguine/services/sinner/fetcher/client"

	"golang.org/x/sync/errgroup"
)

// Sinner parses messages stored in scribe.
type Sinner struct {
	// consumerDB is the database to store consumer data in.
	consumerDB db.EventDB
	// indexers is a mapping of chain IDs -> chain indexers.
	indexers map[uint32]*ChainIndexer
	// config is the config for the indexer.
	config indexerConfig.Config
}

// Parsers holds all the parsers for a given chain.
type Parsers struct {
	// ChainID is the chain these parsers are for.
	ChainID uint32
	// OriginParser parses logs from the origin contract.
	OriginParser *origin.ParserImpl
	// DestinationParser parses logs from the execution hub contract.
	DestinationParser *destination.ParserImpl
}

// NewSinner creates a new sinner indexer service.
//
// nolint:gocognit
func NewSinner(eventDB db.EventDB, config indexerConfig.Config, handler metrics.Handler) (*Sinner, error) {
	chainIndexers := make(map[uint32]*ChainIndexer)
	httpClient := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}

	fetcher := fetcherpkg.NewFetcher(gqlClient.NewClient(&httpClient, config.ScribeURL), handler)

	// Initialize each chain backfiller.
	for _, chainConfig := range config.Chains {
		chainIndexer, err := getChainIndexer(eventDB, chainConfig.ChainID, fetcher, chainConfig)
		if err != nil {
			return nil, fmt.Errorf("could not get chain indexer: %w", err)
		}

		chainIndexers[chainConfig.ChainID] = chainIndexer
	}

	return &Sinner{
		consumerDB: eventDB,
		indexers:   chainIndexers,
		config:     config,
	}, nil
}

// Index iterates over each chain backfiller and calls index concurrently on each one.
//
// nolint:cyclop
func (e Sinner) Index(ctx context.Context) error {
	g, groupCtx := errgroup.WithContext(ctx)

	for i := range e.config.Chains {
		chainConfig := e.config.Chains[i]
		chainIndexer := e.indexers[chainConfig.ChainID]

		g.Go(func() error {
			// generate new context
			chainContext := context.Background()
			for {
				select {
				case <-groupCtx.Done(): // global context canceled
					return fmt.Errorf("global context canceled")
				case <-chainContext.Done(): // local context canceled, reset context
					chainContext = context.Background()
				default:

					err := chainIndexer.Index(chainContext)
					if err != nil {
						// return fmt.Errorf("could not index chain %d: %w", chainConfig.ChainID, err)
						continue // continue trying
					}
					return nil
				}
			}
		})
	}

	if err := g.Wait(); err != nil {
		logger.ReportSinnerError(fmt.Errorf("could not livefill explorer: %w", err), 0, logger.SinnerIndexingFailure)

		return fmt.Errorf("could not livefill explorer: %w", err)
	}
	return nil
}

// nolint gocognit,cyclop
func getChainIndexer(eventDB db.EventDB, chainID uint32, fetcher fetcherpkg.ScribeFetcher, chainConfig indexerConfig.ChainConfig) (*ChainIndexer, error) {
	parsers := Parsers{
		ChainID: chainID,
	}
	for i := range chainConfig.Contracts {
		switch chainConfig.Contracts[i].ContractType {
		case "origin":
			originParser, err := origin.NewParser(common.HexToAddress(chainConfig.Contracts[i].Address), eventDB, chainID)
			if err != nil {
				return nil, fmt.Errorf("could not create origin parser: %w", err)
			}
			parsers.OriginParser = originParser
		case "execution_hub":
			destinationParser, err := destination.NewParser(common.HexToAddress(chainConfig.Contracts[i].Address), eventDB, chainID)
			if err != nil {
				return nil, fmt.Errorf("could not create execution_hub parser: %w", err)
			}
			parsers.DestinationParser = destinationParser
		}
	}

	chainIndexer := NewChainIndexer(eventDB, parsers, fetcher, chainConfig)

	return chainIndexer, nil
}
