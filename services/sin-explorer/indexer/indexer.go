package explorer

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/listener"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/sin-explorer/config"
	"github.com/synapsecns/sanguine/services/sin-explorer/db"
	"github.com/synapsecns/sanguine/sin-executor/contracts/interchainclient"
	"math/big"
)

type indexer struct {
	// handler is the metrics handler
	handler metrics.Handler
	// db is the database
	db db.Service
	// cfg is the configuration
	cfg config.Config
	// omnirpcClient is the omnirpc client
	omnirpcClient omnirpcClient.RPCClient
}

// Indexer is the interface for the indexer.
type Indexer interface {
}

// NewIndexer creates a new indexer.
func NewIndexer(db db.Service, cfg config.Config, handler metrics.Handler) Indexer {
	omniClient := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes(), omnirpcClient.WithDefaultConfirmations(1))

	return &indexer{
		handler:       handler,
		cfg:           cfg,
		db:            db,
		omnirpcClient: omniClient,
	}
}

func (i *indexer) Start(ctx context.Context) error {
	for chainID, chain := range i.cfg.Chains {
		chainClient, err := i.omnirpcClient.GetClient(ctx, new(big.Int).SetInt64(int64(chainID)))
		if err != nil {
			return fmt.Errorf("failed to get client for chain %d: %w", chainID, err)
		}

		blockNumber, err := chainClient.BlockNumber(ctx)
		if err != nil {
			return fmt.Errorf("failed to get block number for chain %d: %w", chainID, err)
		}

		contractListener, err := listener.NewChainListener(chainClient, i.db, []common.Address{common.HexToAddress(chain.InterchainClientAddress)}, blockNumber, i.handler)
		if err != nil {
			return fmt.Errorf("failed to create chain listener for chain %d: %w", chainID, err)
		}

		// TODO: go func this out
		err = contractListener.Listen(ctx, i.parseLog())
		if err != nil {
			return fmt.Errorf("failed to listen for chain %d: %w", chainID, err)
		}
	}
	return nil
}

func (i *indexer) parseLog() listener.HandleLog {
	return func(ctx context.Context, log types.Log) error {
		// Parse the log
		filterer, err := interchainclient.NewParser(log.Address)
		if err != nil {
			return fmt.Errorf("failed to create interchain client: %w", err)
		}

		_, parsedEvent, ok := filterer.ParseEvent(log)
		if !ok {
			if len(log.Topics) == 0 {
				i.handler.ExperimentalLogger().Warnf(ctx, "unknown event %s", log.Topics[0])
			}
			return nil
		}

		switch event := parsedEvent.(type) {
		case *interchainclient.InterchainClientV1InterchainTransactionSent:
			err = i.db.PutInterchainTransactionSent(ctx, event)
			if err != nil {
				return fmt.Errorf("failed to put interchain transaction sent: %w", err)
			}
		case *interchainclient.InterchainClientV1InterchainTransactionReceived:
			err = i.db.PutInterchainTransactionReceived(ctx, event)
			if err != nil {
				return fmt.Errorf("failed to put interchain transaction received: %w", err)
			}
		}

		return nil
	}
}
