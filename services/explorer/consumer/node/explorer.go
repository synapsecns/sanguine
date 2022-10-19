package node

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/backfill"
	gqlClient "github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// ExplorerBackfiller is a backfiller that aggregates all backfilling from ChainBackfillers.
type ExplorerBackfiller struct {
	// consumerDB is the database to store consumer data in
	consumerDB db.ConsumerDB
	// clients is a mapping of chain IDs -> clients
	clients map[uint32]bind.ContractBackend
	// ChainBackfillers is a mapping of chain IDs -> chain backfillers
	ChainBackfillers map[uint32]*backfill.ChainBackfiller
	// config is the config for the backfiller
	config config.Config
}

// NewExplorerBackfiller creates a new backfiller for the explorer.
//
// nolint:gocognit
func NewExplorerBackfiller(consumerDB db.ConsumerDB, config config.Config, clients map[uint32]bind.ContractBackend) (*ExplorerBackfiller, error) {
	// initialize the list of chain backfillers
	chainBackfillers := make(map[uint32]*backfill.ChainBackfiller)

	// create the consumer Fetcher
	fetcher := consumer.NewFetcher(gqlClient.NewClient(http.DefaultClient, config.ScribeURL))

	// create the bridge config Fetcher
	bridgeConfigRef, err := bridgeconfig.NewBridgeConfigRef(common.HexToAddress(config.BridgeConfigAddress), clients[config.BridgeConfigChainID])

	if err != nil || bridgeConfigRef == nil {
		return nil, fmt.Errorf("could not create bridge config Fetcher: %w", err)
	}
	// initialize each chain backfiller
	for _, chainConfig := range config.Chains {
		// create the chain backfiller
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
	// initialize the errgroup
	g, groupCtx := errgroup.WithContext(ctx)

	// iterate over each chain backfiller
	for i := range e.config.Chains {
		// capture func literal
		chainConfig := e.config.Chains[i]

		chainBackfiller := e.ChainBackfillers[chainConfig.ChainID]
		// call Backfill concurrently
		g.Go(func() error {
			err := chainBackfiller.Backfill(groupCtx)
			if err != nil {
				return fmt.Errorf("could not backfill chain %d: %w", chainConfig.ChainID, err)
			}
			return nil
		})
	}
	// wait for all backfills to finish
	if err := g.Wait(); err != nil {
		return fmt.Errorf("could not backfill explorer: %w", err)
	}

	return nil
}

// nolint gocognit,cyclop
func getChainBackfiller(consumerDB db.ConsumerDB, chainConfig config.ChainConfig, fetcher *consumer.Fetcher, client bind.ContractBackend, bridgeConfigAddress common.Address, bridgeRef *bridgeconfig.BridgeConfigRef) (*backfill.ChainBackfiller, error) {

	newConfigFetcher, err := consumer.NewBridgeConfigFetcher(bridgeConfigAddress, bridgeRef)
	if err != nil || newConfigFetcher == nil {
		return nil, fmt.Errorf("could not get bridge abi: %w", err)
	}

	// create the bridge parser
	bridgeParser, err := consumer.NewBridgeParser(consumerDB, common.HexToAddress(chainConfig.SynapseBridgeAddress), *newConfigFetcher, fetcher)
	if err != nil || bridgeParser == nil {
		return nil, fmt.Errorf("could not create bridge parser: %w", err)
	}
	// create the swap parsers
	swapParsers := make(map[common.Address]*consumer.SwapParser)
	for _, swapAddress := range chainConfig.SwapFlashLoanAddresses {
		// create the swap Fetcher
		swapFetcher, err := consumer.NewSwapFetcher(common.HexToAddress(swapAddress), client)
		if err != nil || swapFetcher == nil {
			return nil, fmt.Errorf("could not create swap Fetcher: %w", err)
		}
		swapParser, err := consumer.NewSwapParser(consumerDB, common.HexToAddress(swapAddress), *swapFetcher, fetcher)
		if err != nil || swapParser == nil {
			return nil, fmt.Errorf("could not create swap parser: %w", err)
		}
		swapParsers[common.HexToAddress(swapAddress)] = swapParser
	}
	chainBackfiller := backfill.NewChainBackfiller(consumerDB, bridgeParser, swapParsers, *fetcher, chainConfig)

	return chainBackfiller, nil
}
