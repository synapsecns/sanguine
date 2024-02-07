package node

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/backfill"
	indexerConfig "github.com/synapsecns/sanguine/services/explorer/config/indexer"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/consumer/client"
	fetcherpkg "github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/fetcher/tokenprice"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/static"
	"golang.org/x/sync/errgroup"
)

// ExplorerBackfiller is a backfiller that aggregates all backfilling from ChainBackfillers.
type ExplorerBackfiller struct {
	// consumerDB is the database to store consumer data in.
	consumerDB db.ConsumerDB
	// clients is a mapping of chain IDs -> clients.
	clients map[uint32]bind.ContractBackend
	// ChainBackfillers is a mapping of chain IDs -> chain backfillers.
	ChainBackfillers map[uint32]*backfill.ChainBackfiller
	// config is the config for the backfiller.
	config indexerConfig.Config
}

// NewExplorerBackfiller creates a new backfiller for the explorer.
//
// nolint:gocognit
func NewExplorerBackfiller(consumerDB db.ConsumerDB, config indexerConfig.Config, clients map[uint32]bind.ContractBackend, handler metrics.Handler) (*ExplorerBackfiller, error) {
	chainBackfillers := make(map[uint32]*backfill.ChainBackfiller)
	httpClient := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
	fetcher := fetcherpkg.NewFetcher(gqlClient.NewClient(&httpClient, config.ScribeURL), handler)
	bridgeConfigRef, err := bridgeconfig.NewBridgeConfigRef(common.HexToAddress(config.BridgeConfigAddress), clients[config.BridgeConfigChainID])
	if err != nil || bridgeConfigRef == nil {
		return nil, fmt.Errorf("could not create bridge config ScribeFetcher: %w", err)
	}
	priceDataService, err := tokenprice.NewPriceDataService()
	if err != nil {
		return nil, fmt.Errorf("could not create price data service: %w", err)
	}
	newConfigFetcher, err := fetcherpkg.NewBridgeConfigFetcher(common.HexToAddress(config.BridgeConfigAddress), bridgeConfigRef)
	if err != nil || newConfigFetcher == nil {
		return nil, fmt.Errorf("could not get bridge abi: %w", err)
	}
	tokenSymbolToIDs, err := parser.ParseYaml(static.GetTokenSymbolToTokenIDConfig())
	if err != nil {
		return nil, fmt.Errorf("could not open yaml file: %w", err)
	}
	tokenDataService, err := tokendata.NewTokenDataService(newConfigFetcher, tokenSymbolToIDs)
	if err != nil {
		return nil, fmt.Errorf("could not create token data service: %w", err)
	}

	// Initialize each chain backfiller.
	for _, chainConfig := range config.Chains {
		chainBackfiller, err := getChainBackfiller(consumerDB, chainConfig, fetcher, clients[chainConfig.ChainID], tokenDataService, priceDataService)
		if err != nil {
			return nil, fmt.Errorf("could not get chain backfiller: %w", err)
		}

		chainBackfillers[chainConfig.ChainID] = chainBackfiller
	}

	return &ExplorerBackfiller{
		consumerDB:       consumerDB,
		clients:          clients,
		ChainBackfillers: chainBackfillers,
		config:           config,
	}, nil
}

// Backfill iterates over each chain backfiller and calls Backfill concurrently on each one.
//
// nolint:cyclop
func (e ExplorerBackfiller) Backfill(ctx context.Context, livefill bool) error {
	refreshRate := e.config.DefaultRefreshRate

	if refreshRate == 0 {
		refreshRate = 1
	}

	g, groupCtx := errgroup.WithContext(ctx)

	for i := range e.config.Chains {
		chainConfig := e.config.Chains[i]
		chainBackfiller := e.ChainBackfillers[chainConfig.ChainID]
		g.Go(func() error {
			err := chainBackfiller.Backfill(groupCtx, livefill, refreshRate)
			if err != nil {
				return fmt.Errorf("could not backfill chain %d: %w", chainConfig.ChainID, err)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		logger.Errorf("backfill completed: %v", err)

		return fmt.Errorf("could not livefill explorer: %w", err)
	}
	logger.Errorf("backfill completed with no errors")

	return nil
}

// nolint gocognit,cyclop
func getChainBackfiller(consumerDB db.ConsumerDB, chainConfig indexerConfig.ChainConfig, fetcher fetcherpkg.ScribeFetcher, client bind.ContractBackend, tokenDataService tokendata.Service, priceDataService tokenprice.Service) (*backfill.ChainBackfiller, error) {
	var err error
	var bridgeParser *parser.BridgeParser
	var messageBusParser *parser.MessageBusParser
	var cctpParser *parser.CCTPParser
	var swapService fetcherpkg.SwapService
	var cctpService fetcherpkg.CCTPService

	swapParsers := make(map[common.Address]*parser.SwapParser)

	for i := range chainConfig.Contracts {
		switch chainConfig.Contracts[i].ContractType {
		case "bridge":
			bridgeParser, err = parser.NewBridgeParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), tokenDataService, fetcher, priceDataService, false)
			if err != nil || bridgeParser == nil {
				return nil, fmt.Errorf("could not create bridge parser: %w", err)
			}
		case "swap":
			swapService, err = fetcherpkg.NewSwapFetcher(common.HexToAddress(chainConfig.Contracts[i].Address), client, false)
			if err != nil || swapService == nil {
				return nil, fmt.Errorf("could not create swapService: %w", err)
			}
			swapParser, err := parser.NewSwapParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), false, fetcher, swapService, tokenDataService, priceDataService)
			if err != nil || swapParser == nil {
				return nil, fmt.Errorf("could not create swap parser: %w", err)
			}

			swapParsers[common.HexToAddress(chainConfig.Contracts[i].Address)] = swapParser
		case "metaswap":
			if swapService == nil {
				swapService, err := fetcherpkg.NewSwapFetcher(common.HexToAddress(chainConfig.Contracts[i].Address), client, true)
				if err != nil || swapService == nil {
					return nil, fmt.Errorf("could not create swapService: %w", err)
				}
			}
			swapParser, err := parser.NewSwapParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), true, fetcher, swapService, tokenDataService, priceDataService)
			if err != nil || swapParser == nil {
				return nil, fmt.Errorf("could not create swap parser: %w", err)
			}

			swapParsers[common.HexToAddress(chainConfig.Contracts[i].Address)] = swapParser
		case "messagebus":
			messageBusParser, err = parser.NewMessageBusParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), fetcher, priceDataService)
			if err != nil || messageBusParser == nil {
				return nil, fmt.Errorf("could not create message bus parser: %w", err)
			}
		case "cctp":
			cctpService, err = fetcherpkg.NewCCTPFetcher(common.HexToAddress(chainConfig.Contracts[i].Address), client)
			if err != nil || cctpService == nil {
				return nil, fmt.Errorf("could not create cctpService: %w", err)
			}
			cctpParser, err = parser.NewCCTPParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), fetcher, cctpService, tokenDataService, priceDataService, false)
			if err != nil || cctpParser == nil {
				return nil, fmt.Errorf("could not create message bus parser: %w", err)
			}
		}
	}

	// TODO Add the cctp parser
	chainBackfiller := backfill.NewChainBackfiller(consumerDB, bridgeParser, swapParsers, messageBusParser, cctpParser, fetcher, chainConfig)

	return chainBackfiller, nil
}
