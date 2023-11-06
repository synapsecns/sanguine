package service

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/sinner/logger"
	sinnerTypes "github.com/synapsecns/sanguine/services/sinner/types"
	"net/http"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	indexerConfig "github.com/synapsecns/sanguine/services/sinner/config/indexer"
	"github.com/synapsecns/sanguine/services/sinner/contracts/destination"
	"github.com/synapsecns/sanguine/services/sinner/contracts/origin"
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
	refreshRate := time.Duration(config.DefaultRefreshRate) * time.Second
	// Initialize each chain backfiller.
	for _, chainConfig := range config.Chains {
		chainIndexer, err := getChainIndexer(eventDB, chainConfig.ChainID, fetcher, chainConfig, refreshRate)
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
	errChan := make(chan error, len(e.config.Chains)) // Buffered to prevent goroutine blockage

	// Listen for errors
	go func() {
		for err := range errChan {
			logger.ReportSinnerError(fmt.Errorf("could not livefill explorer: %w", err), 0, logger.SinnerIndexingFailure)
		}
	}()

	for i := range e.config.Chains {
		chainConfig := e.config.Chains[i]
		chainIndexer := e.indexers[chainConfig.ChainID]

		wg.Add(1)
		go func(chainCfg *indexerConfig.ChainConfig, indexer *ChainIndexer) {
			defer wg.Done()

			for {
				chainCtx, cancelChainCtx := context.WithCancel(ctx)

				select {
				case <-ctx.Done(): // global context canceled
					errChan <- fmt.Errorf("global context canceled")
					cancelChainCtx() // cancel the local context before returning
					return
				default:
					err := chainIndexer.Index(chainCtx)
					cancelChainCtx() // cancel the local context immediately after its use

					if err != nil {
						errChan <- fmt.Errorf("error indexing chain %d: %w", chainCfg.ChainID, err)
						continue // continue trying
					}
					return
				}
			}
		}(&chainConfig, chainIndexer)
	}

	wg.Wait()
	close(errChan)

	return nil
}

// nolint gocognit,cyclop
func getChainIndexer(eventDB db.EventDB, chainID uint32, fetcher fetcherpkg.ScribeFetcher, chainConfig indexerConfig.ChainConfig, refreshRate time.Duration) (*ChainIndexer, error) {
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

	chainIndexer := NewChainIndexer(eventDB, parsers, fetcher, chainConfig, refreshRate)

	return chainIndexer, nil
}
