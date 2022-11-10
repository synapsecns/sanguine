package node

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/backfill"
	"github.com/synapsecns/sanguine/services/explorer/config"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/consumer/client"
	fetcherpkg "github.com/synapsecns/sanguine/services/explorer/consumer/fetcher"
	"github.com/synapsecns/sanguine/services/explorer/consumer/parser"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"golang.org/x/sync/errgroup"
	"net/http"
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
func (e ExplorerBackfiller) Backfill(ctx context.Context) error {
	g, groupCtx := errgroup.WithContext(ctx)

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
		return fmt.Errorf("could not backfill explorer: %w", err)
	}

	return nil
}

// nolint gocognit,cyclop
func getChainBackfiller(consumerDB db.ConsumerDB, chainConfig config.ChainConfig, fetcher *fetcherpkg.ScribeFetcher, client bind.ContractBackend, bridgeConfigAddress common.Address, bridgeRef *bridgeconfig.BridgeConfigRef) (*backfill.ChainBackfiller, error) {
	newConfigFetcher, err := fetcherpkg.NewBridgeConfigFetcher(bridgeConfigAddress, bridgeRef)
	if err != nil || newConfigFetcher == nil {
		return nil, fmt.Errorf("could not get bridge abi: %w", err)
	}

	swapParsers := make(map[common.Address]*parser.SwapParser)

	var bridgeParser *parser.BridgeParser

	for i := range chainConfig.Contracts {
		switch chainConfig.Contracts[i].ContractType {
		case "bridge":
			bridgeParser, err = parser.NewBridgeParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), *newConfigFetcher, fetcher)
			if err != nil || bridgeParser == nil {
				return nil, fmt.Errorf("could not create bridge parser: %w", err)
			}
		case "swap":
			swapFetcher, err := fetcherpkg.NewSwapFetcher(common.HexToAddress(chainConfig.Contracts[i].Address), client)
			if err != nil || swapFetcher == nil {
				return nil, fmt.Errorf("could not create swap ScribeFetcher: %w", err)
			}

			swapParser, err := parser.NewSwapParser(consumerDB, common.HexToAddress(chainConfig.Contracts[i].Address), *swapFetcher, fetcher)
			if err != nil || swapParser == nil {
				return nil, fmt.Errorf("could not create swap parser: %w", err)
			}

			swapParsers[common.HexToAddress(chainConfig.Contracts[i].Address)] = swapParser
		}
	}

	chainBackfiller := backfill.NewChainBackfiller(consumerDB, bridgeParser, swapParsers, *fetcher, chainConfig)

	return chainBackfiller, nil
}
