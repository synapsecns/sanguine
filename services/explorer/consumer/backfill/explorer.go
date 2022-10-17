package backfill

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/explorer/consumer"
	"github.com/synapsecns/sanguine/services/explorer/consumer/client"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"github.com/synapsecns/sanguine/services/explorer/db"
	"golang.org/x/sync/errgroup"
	"net/http"
)

// ExplorerBackfiller is a backfiller that aggregates all backfilling from ChainBackfillers.
type ExplorerBackfiller struct {
	// consumerDB is the database to store consumer data in
	consumerDB db.ConsumerDB
	// ChainBackfillers is a mapping of chain IDs -> chain backfillers
	ChainBackfillers map[uint32]*ChainBackfiller
	// config is the config for the backfiller
	config config.Config
}

// NewExplorerBackfiller creates a new backfiller for the explorer.
//
// nolint:gocognit
func NewExplorerBackfiller(ctx context.Context, consumerDB db.ConsumerDB, config config.Config) (*ExplorerBackfiller, error) {
	// initialize the list of chain backfillers
	chainBackfillers := make(map[uint32]*ChainBackfiller)
	// initialize each chain backfiller
	for _, chainConfig := range config.Chains {
		bridgeConfigV3ABI, err := bridgeconfig.BridgeConfigV3MetaData.GetAbi()
		if err != nil || bridgeConfigV3ABI == nil {
			return nil, fmt.Errorf("could not get bridge config v3 abi: %w", err)
		}
		// create the chain backfiller
		chainBackfiller, err := getChainBackfiller(ctx, consumerDB, chainConfig, config.ScribeURL)
		if err != nil {
			return nil, fmt.Errorf("could not get chain backfiller: %w", err)
		}

		chainBackfillers[chainConfig.ChainID] = chainBackfiller
	}
	return &ExplorerBackfiller{
		consumerDB:       consumerDB,
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

// nolint:gocognit,cyclop
func getChainBackfiller(ctx context.Context, consumerDB db.ConsumerDB, chainConfig config.ChainConfig, baseURL string) (*ChainBackfiller, error) {
	// get the ABI for each contract
	bridgeConfigABI, err := bridgeconfig.BridgeConfigV3MetaData.GetAbi()
	if err != nil || bridgeConfigABI == nil {
		return nil, fmt.Errorf("could not get bridge config v3 abi: %w", err)
	}
	swapABI, err := swap.SwapFlashLoanMetaData.GetAbi()
	if err != nil || swapABI == nil {
		return nil, fmt.Errorf("could not get swap flash loan abi: %w", err)
	}
	bridgeABI, err := bridge.SynapseBridgeMetaData.GetAbi()
	if err != nil || bridgeABI == nil {
		return nil, fmt.Errorf("could not get bridge abi: %w", err)
	}

	// create the client
	ethClient, err := ethclient.DialContext(ctx, chainConfig.RPCURL)
	// create the consumer Fetcher
	fetcher := consumer.NewFetcher(client.NewClient(http.DefaultClient, baseURL))
	// create the bridge config Fetcher
	bridgeConfigFetcher, err := consumer.NewBridgeConfigFetcher(common.HexToAddress(chainConfig.BridgeConfigV3Address), ethClient)
	if err != nil || bridgeConfigFetcher == nil {
		return nil, fmt.Errorf("could not create bridge config Fetcher: %w", err)
	}
	// create the bridge parser
	bridgeParser, err := consumer.NewBridgeParser(consumerDB, common.HexToAddress(chainConfig.SynapseBridgeAddress), *bridgeConfigFetcher, fetcher)
	if err != nil || bridgeParser == nil {
		return nil, fmt.Errorf("could not create bridge parser: %w", err)
	}
	// create the swap parsers
	swapParsers := make(map[common.Address]*consumer.SwapParser)
	for _, swapAddress := range chainConfig.SwapFlashLoanAddresses {
		// create the swap Fetcher
		swapFetcher, err := consumer.NewSwapFetcher(common.HexToAddress(swapAddress), ethClient)
		if err != nil || swapFetcher == nil {
			return nil, fmt.Errorf("could not create swap Fetcher: %w", err)
		}
		swapParser, err := consumer.NewSwapParser(consumerDB, common.HexToAddress(swapAddress), *swapFetcher, fetcher)
		if err != nil || swapParser == nil {
			return nil, fmt.Errorf("could not create swap parser: %w", err)
		}
		swapParsers[common.HexToAddress(swapAddress)] = swapParser
	}
	chainBackfiller := NewChainBackfiller(consumerDB, bridgeParser, swapParsers, *fetcher, chainConfig)

	return chainBackfiller, nil
}
