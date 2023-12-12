package service

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/rfq/rfq-relayer/bindings"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/config"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/db"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/listener"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/queue"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/service/quote"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	relayerTypes "github.com/synapsecns/sanguine/rfq/rfq-relayer/utils"

	"math/big"
	"strings"

	"github.com/synapsecns/sanguine/core/metrics"
	EVMClient "github.com/synapsecns/sanguine/ethergo/client"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
)

const (
	// MaxEventChanSize is the maximum size of the event channel.
	MaxEventChanSize = 1000
	// MaxSeenChanSize is the maximum size of the seen channel.
	MaxSeenChanSize = 1000
)

// IRelayer is the interface for the relayer.
type IRelayer interface {
	Start(ctx context.Context) error
}

type relayerImpl struct {
	db           db.DB
	config       *config.Config
	eventChan    chan relayerTypes.WrappedLog
	seenChan     chan relayerTypes.WrappedLog
	txSubmitter  submitter.TransactionSubmitter
	evmClients   map[uint32]EVMClient.EVM
	contracts    map[uint32]*bindings.FastBridge
	chainConfigs map[uint32]*listener.ChainListenerConfig
	claimQueue   *queue.Queue
	quoter       quote.IQuoter
}

// NewRelayer inits all necessary components for the relayer.
func NewRelayer(ctx context.Context, cfg *config.Config, db db.DB, handler metrics.Handler) (IRelayer, error) {
	// Init OmniRPC client
	omniRPCClient := omniClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omniClient.WithCaptureReqRes())

	// Init Submitter and Signer
	signer, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not make the signer: %w", err)
	}
	txSubmitter := submitter.NewTransactionSubmitter(handler, signer, omniRPCClient, db.SubmitterDB(), &cfg.SubmitterConfig)

	// Iterate over each chain configuration and init clients, listener configs, contract bindings, (etc.) for each
	listenerConfigs := make(map[uint32]*listener.ChainListenerConfig)
	evmClients := make(map[uint32]EVMClient.EVM)
	fastBridgeContracts := make(map[uint32]*bindings.FastBridge)
	for _, chainCfg := range cfg.Chains {
		// Get the EVM client for the current chain
		client, chainErr := omniRPCClient.GetClient(ctx, big.NewInt(int64(chainCfg.ChainID)))
		if chainErr != nil {
			return nil, fmt.Errorf("error creating client for chain ID %d: %w", chainCfg.ChainID, chainErr)
		}
		evmClients[chainCfg.ChainID] = client

		// Get the FastBridge contract instance for the current chain for relay(), prove(), claim() execution
		fastBridgeRef, chainErr := bindings.NewFastBridge(common.HexToAddress(chainCfg.FastBridgeAddress), client)
		if chainErr != nil {
			return nil, fmt.Errorf("error creating FastBridge Contract for chain ID %d: %w", chainCfg.ChainID, chainErr)
		}
		fastBridgeContracts[chainCfg.ChainID] = fastBridgeRef
		parsedABI, chainErr := abi.JSON(strings.NewReader(bindings.FastBridgeMetaData.ABI))
		if chainErr != nil {
			return nil, fmt.Errorf("could not parse ABI: %w", chainErr)
		}

		// Create listener config for current chain.
		listenerConfigs[chainCfg.ChainID] = &listener.ChainListenerConfig{
			ChainID:         chainCfg.ChainID,
			StartBlock:      chainCfg.FastBridgeBlockDeployed,
			BridgeAddress:   common.HexToAddress(chainCfg.FastBridgeAddress),
			Client:          client,
			PollInterval:    chainCfg.PollInterval,
			MaxGetLogsRange: chainCfg.MaxGetLogsRange,
			Confirmations:   chainCfg.Confirmations,
			ABI:             parsedABI,
		}
	}

	// Create the quoter (balance management, unconfirmed bridge events, quoter API connection.
	quoter, err := quote.NewQuoter(ctx, evmClients, cfg.Assets, common.HexToAddress(cfg.RelayerAddress), cfg.RFQURL)
	if err != nil {
		//return nil, fmt.Errorf("could not create quoter: %w", err)
	}
	// Init queue to handle claim() execution
	claimQueue, err := queue.NewQueue(ctx, cfg.MaxQueueSize, cfg.Deadline, db)
	if err != nil {
		//return nil, fmt.Errorf("could not create queue: %w", err)
	}

	// Listener Channels
	eventChan := make(chan relayerTypes.WrappedLog, MaxEventChanSize)
	seenChan := make(chan relayerTypes.WrappedLog, MaxEventChanSize)

	return &relayerImpl{
		eventChan:   eventChan,
		seenChan:    seenChan,
		db:          db,
		txSubmitter: txSubmitter,
		evmClients:  evmClients,
		contracts:   fastBridgeContracts,
		claimQueue:  claimQueue,
		quoter:      quoter,
	}, nil
}

// Start starts the relayer.
func (r *relayerImpl) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("global context canceled %w", ctx.Err())
		default:
			g, gCtx := errgroup.WithContext(ctx)

			// Listen to events on all chains
			g.Go(func() error {
				return r.RunListeners(gCtx)
			})

			// Process (confirmed) events from listener
			g.Go(func() error {
				return r.HandleConfirmedEvents(gCtx)
			})

			// Process (unconfirmed/seen) events from listener
			g.Go(func() error {
				return r.HandleUnconfirmedEvents(gCtx)
			})

			// Check queue for events ready to claim
			g.Go(func() error {
				return r.HandleClaimEvents(gCtx)
			})
			err := g.Wait()
			if err != nil {
				logger.Errorf("could not run relayer, retrying. Error: %v", err)
				continue
			}
		}
	}
}
