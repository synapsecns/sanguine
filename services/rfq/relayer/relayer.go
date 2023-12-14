package relayer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
)

// Relayer is the core of the relayer application
type Relayer struct {
	cfg     relconfig.Config
	metrics metrics.Handler
	db      reldb.Service
	client  omnirpcClient.RPCClient
}

var logger = log.Logger("relayer")

// NewRelayer creates a new relayer.
func NewRelayer(ctx context.Context, metricHandler metrics.Handler, cfg relconfig.Config) (Relayer, error) {
	omniClient := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, metricHandler, omnirpcClient.WithCaptureReqRes())

	// TODO: add bd
	return Relayer{
		client:  omniClient,
		metrics: metricHandler,
		cfg:     cfg,
	}, nil
}

func (r *Relayer) Start(ctx context.Context) error {
	return nil
}

func (r *Relayer) startChainParser(ctx context.Context) error {
	// TODO: good chance we wanna prepare these chain listeners up front and then listen later.
	for chainID, bridgeStr := range r.cfg.Bridges {
		// TODO: consider getter for this convert step
		bridge := common.HexToAddress(bridgeStr)
		chainClient, err := r.client.GetChainClient(ctx, chainID)
		if err != nil {
			return fmt.Errorf("could not get chain client: %w", err)
		}

		parser, err := fastbridge.NewParser(bridge)
		if err != nil {
			return fmt.Errorf("could not parse: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, r.db, bridge, r.metrics)
		if err != nil {
			return fmt.Errorf("could not get chain listener: %w", err)
		}

		err = chainListener.Listen(ctx, func(ctx context.Context, log types.Log) error {
			_, parsedEvent, ok := parser.ParseEvent(log)
			// handle unknown event
			if !ok {
				if len(log.Topics) != 0 {
					logger.Warnf("unknown event %s", log.Topics[0])
				}
				return nil
			}

			switch event := parsedEvent.(type) {
			case *fastbridge.FastBridgeBridgeRequested:
				// TODO store this if not already seen
				_ = event
			case *fastbridge.FastBridgeBridgeRelayed:
				panic("implement me")

			}

			return nil
		})
	}
	return nil
}
