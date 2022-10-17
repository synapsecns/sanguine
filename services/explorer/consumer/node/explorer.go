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

//
//// Start starts the explorer. This works by starting a backfill from a predetermined block from the config, and then
//// back filling to the last block that is stored by Scribe on each chain. The last block stored by Scribe is gotten
//// using the consumer's fetcher.
// func (e Explorer) Start(ctx context.Context) error {
//	refreshRate := e.config.RefreshRate
//	if refreshRate == 0 {
//		refreshRate = 1
//	}
//	// backfill each chain
//	g, groupCtx := errgroup.WithContext(ctx)
//	for i := range e.config.Chains {
//		// capture the func literal
//		chainConfig := e.config.Chains[i]
//		g.Go(func() error {
//			// backoff in case of an error
//			b := &backoff.Backoff{
//				Factor: 2,
//				Jitter: true,
//				Min:    1 * time.Second,
//				Max:    30 * time.Second,
//			}
//			// timeout should always be 0 on the first attempt
//			timeout := time.Duration(0)
//			for {
//				select {
//				case <-groupCtx.Done():
//					return fmt.Errorf("context canceled: %w", groupCtx.Err())
//				case <-time.After(timeout):
//					err := e.explorerBackfiller.ChainBackfillers[chainConfig.ChainID].Backfill(groupCtx)
//					if err != nil {
//						timeout = b.Duration()
//						continue
//					}
//					b.Reset()
//					timeout = time.Duration(refreshRate) * time.Second
//				}
//			}
//		})
//	}
//
//	if err := g.Wait(); err != nil {
//		return fmt.Errorf("error in explorer: %w", err)
//	}
//
//	return nil
//}
