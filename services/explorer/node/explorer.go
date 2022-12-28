package node

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jpillora/backoff"
	"github.com/synapsecns/sanguine/services/explorer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/config"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/consumer/client"
	fetcherpkg "github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser/tokendata"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"github.com/synapsecns/sanguine/services/explorer/static"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
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
	config config.Config
}

// NewExplorerBackfiller creates a new backfiller for the explorer.
//
// nolint:gocognit
func NewExplorerBackfiller(consumerDB db.ConsumerDB, config config.Config, clients map[uint32]bind.ContractBackend) (*ExplorerBackfiller, error) {
	chainBackfillers := make(map[uint32]*backfill.ChainBackfiller)
	fetcher := fetcherpkg.NewFetcher(gqlClient.NewClient(http.DefaultClient, config.ScribeURL))
	bridgeConfigRef, err := bridgeconfig.NewBridgeConfigRef(common.HexToAddress(config.BridgeConfigAddress), clients[config.BridgeConfigChainID])
	if err != nil || bridgeConfigRef == nil {
		return nil, fmt.Errorf("could not create bridge config ScribeFetcher: %w", err)
	}

	// Initialize each chain backfiller.
	for _, chainConfig := range config.Chains {
		chainBackfiller, err := getChainBackfiller(consumerDB, chainConfig, fetcher, clients[chainConfig.ChainID], common.HexToAddress(config.BridgeConfigAddress), bridgeConfigRef)
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
	refreshRate := e.config.RefreshRate

	if refreshRate == 0 {
		refreshRate = 1
	}

	g, groupCtx := errgroup.WithContext(ctx)
	if !livefill {
		for i := range e.config.Chains {
			chainConfig := e.config.Chains[i]
			chainBackfiller := e.ChainBackfillers[chainConfig.ChainID]
			g.Go(func() error {
				err := chainBackfiller.Backfill(groupCtx)
				if err != nil {
					return fmt.Errorf("could not backfill chain %d: %w", chainConfig.ChainID, err)
				}
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			logger.Errorf("livefill compelted: %v", err)

			return fmt.Errorf("could not backfill explorer: %w", err)
		}
		logger.Errorf("livefill compelted")

		return nil
	}

	for i := range e.config.Chains {
		chainConfig := e.config.Chains[i]
		chainBackfiller := e.ChainBackfillers[chainConfig.ChainID]

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
				case <-groupCtx.Done():
					logger.Errorf("livefill of chain %d failed: %v", chainConfig.ChainID, groupCtx.Err())

					return fmt.Errorf("livefill of chain %d failed: %w", chainConfig.ChainID, groupCtx.Err())
				case <-time.After(timeout):
					err := chainBackfiller.Backfill(groupCtx)
					if err != nil {
						timeout = b.Duration()
						logger.Warnf("could not livefill chain, retrying %d: %v", chainConfig.ChainID, err)

						continue
					}

					b.Reset()
					timeout = time.Duration(refreshRate) * time.Second
					logger.Errorf("processed range for chain %d, continuing to livefill", chainConfig.ChainID)
				}
			}
		})
	}
	err := g.Wait()
	logger.Errorf("livefill compelted: %v", err)
	return fmt.Errorf("livefill compelted: %w", err)
}

// nolint gocognit,cyclop
func getChainBackfiller(consumerDB db.ConsumerDB, chainConfig config.ChainConfig, fetcher *fetcherpkg.ScribeFetcher, client bind.ContractBackend, bridgeConfigAddress common.Address, bridgeRef *bridgeconfig.BridgeConfigRef) (*backfill.ChainBackfiller, error) {
	newConfigFetcher, err := fetcherpkg.NewBridgeConfigFetcher(bridgeConfigAddress, bridgeRef)
	if err != nil || newConfigFetcher == nil {
		return nil, fmt.Errorf("could not get bridge abi: %w", err)
	}

	swapParsers := make(map[common.Address]*parser.SwapParser)

	var bridgeParser *parser.BridgeParser
	var messageBusParser *parser.MessageBusParser

	tokenSymbolToIDs, err := parser.ParseYaml(static.GetTokenSymbolToTokenIDConfig())
	if err != nil {
		return nil, fmt.Errorf("could not open yaml file: %w", err)
	}
	tokenDataService, err := tokendata.NewTokenDataService(newConfigFetcher, tokenSymbolToIDs)
	if err != nil {
		return nil, fmt.Errorf("could not create token data service: %w", err)
	}

	for i := range chainConfig.Contracts {
		switch chainConfig.Contracts[i].ContractType {
		case "bridge":
			bridgeParser, err = parser.NewBridgeParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), tokenDataService, fetcher)
			if err != nil || bridgeParser == nil {
				return nil, fmt.Errorf("could not create bridge parser: %w", err)
			}
		case "swap":
			swapService, err := fetcherpkg.NewSwapFetcher(common.HexToAddress(chainConfig.Contracts[i].Address), client)
			if err != nil || swapService == nil {
				return nil, fmt.Errorf("could not create swapService: %w", err)
			}
			swapParser, err := parser.NewSwapParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), fetcher, &swapService, tokenDataService)
			if err != nil || swapParser == nil {
				return nil, fmt.Errorf("could not create swap parser: %w", err)
			}

			swapParsers[common.HexToAddress(chainConfig.Contracts[i].Address)] = swapParser
		case "messagebus":
			messageBusParser, err = parser.NewMessageBusParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), fetcher)
			if err != nil || messageBusParser == nil {
				return nil, fmt.Errorf("could not create message bus parser: %w", err)
			}
		}
	}

	chainBackfiller := backfill.NewChainBackfiller(consumerDB, bridgeParser, swapParsers, messageBusParser, *fetcher, chainConfig)

	return chainBackfiller, nil
}
