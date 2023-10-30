package service

import (
	"context"
	"fmt"
	sinnerTypes "github.com/synapsecns/sanguine/services/sinner/types"
	"net/http"
	"sync"
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
	OriginParser sinnerTypes.EventParser
	// DestinationParser parses logs from the execution hub contract.
	DestinationParser sinnerTypes.EventParser
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
	var wg sync.WaitGroup
	errChan := make(chan error)

	for i := range e.config.Chains {
		chainConfig := e.config.Chains[i]
		chainIndexer := e.indexers[chainConfig.ChainID]

		go func(chainCfg *indexerConfig.ChainConfig, indexer *ChainIndexer) {
			defer wg.Done()
			// generate new context
			chainContext := context.WithoutCancel(ctx)
			for {
				select {
				case <-ctx.Done(): // global context canceled
					errChan <- fmt.Errorf("global context canceled")
					return
				case <-chainContext.Done(): // local context canceled, reset context
					chainContext = context.WithoutCancel(ctx)
				default:
					err := chainIndexer.Index(chainContext)
					if err != nil {
						errChan <- fmt.Errorf(" error indexing chain %d: %w", chainConfig.ChainID, err)
						continue // continue trying
					}
					return
				}
			}
		}(&chainConfig, chainIndexer)
	}

	// Goroutine to collect errors, closes once the wait group is done.
	go func() {
		for err := range errChan {
			logger.ReportSinnerError(fmt.Errorf("could not livefill explorer: %w", err), 0, logger.SinnerIndexingFailure)
		}
	}()

	// Wait for all goroutines to finish
	wg.Wait()
	close(errChan)
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
