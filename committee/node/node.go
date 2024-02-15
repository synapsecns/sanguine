package node

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/committee/config"
	"github.com/synapsecns/sanguine/committee/contracts/synapsemodule"
	"github.com/synapsecns/sanguine/committee/db"
	"github.com/synapsecns/sanguine/committee/db/connect"
	"github.com/synapsecns/sanguine/committee/listener"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

type Node struct {
	client         omnirpcClient.RPCClient
	metrics        metrics.Handler
	db             db.Service
	submitter      submitter.TransactionSubmitter
	signer         signer.Signer
	cfg            config.Config
	chainListeners map[int]listener.ContractListener
}

var logger = log.Logger("node")

// NewNode creates a new node.
func NewNode(ctx context.Context, handler metrics.Handler, cfg config.Config) (*Node, error) {
	omniClient := omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}

	var decoderContract *synapsemodule.SynapseModuleRef

	// Note: since decoder is a pointer, this will ony deref if chain length is empty & this is called.
	store, err := connect.Connect(ctx, dbType, cfg.Database.DSN, handler, func(ctx context.Context, data []byte) (synapsemodule.InterchainInterchainTransaction, error) {
		return decoderContract.DecodeInterchainTransaction(&bind.CallOpts{Context: ctx}, data)
	})
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	chainListeners := make(map[int]listener.ContractListener)

	for chainID, chainCFG := range cfg.Chains {
		synapseModule := common.HexToAddress(chainCFG.SynapseModuleAddress)
		chainClient, err := omniClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, store, synapseModule, handler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		chainListeners[chainID] = chainListener

		decoderContract, err = synapsemodule.NewSynapseModuleRef(synapseModule, chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not get synapse module ref: %w", err)
		}
	}

	sg, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer: %w", err)
	}

	sm := submitter.NewTransactionSubmitter(handler, sg, omniClient, store.SubmitterDB(), &cfg.SubmitterConfig)

	return &Node{
		client:    omniClient,
		metrics:   handler,
		cfg:       cfg,
		db:        store,
		submitter: sm,
		signer:    sg,
	}, nil
}

func (n *Node) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return n.submitter.Start(ctx)
	})

	g.Go(func() error {
		return n.startChainIndexers(ctx)
	})

	return nil

}

// startChainIndexers starts the chain indexers for each bridge.
// these listen on the chain using the chain listeners and then handle the events.
func (n *Node) startChainIndexers(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	// TODO: good chance we wanna prepare these chain listeners up front and then listen later.
	for chainID := range n.cfg.Chains {
		chainID := chainID // capture func literal

		g.Go(func() error {
			err := n.runChainIndexer(ctx, chainID)
			if err != nil {
				return fmt.Errorf("could not runChainIndexer chain indexer for chain %d: %w", chainID, err)
			}
			return nil
		})
	}
	return nil
}

// runChainIndexer runs the chain indexer for a given chain.
// any events that an action exists for are indexed.
// nolint: cyclop
func (n *Node) runChainIndexer(parentCtx context.Context, chainID int) (err error) {
	chainListener := n.chainListeners[chainID]

	parser, err := fastbridge.NewParser(chainListener.Address())
	if err != nil {
		return fmt.Errorf("could not parse: %w", err)
	}

	err = chainListener.Listen(parentCtx, func(parentCtx context.Context, log types.Log) (err error) {
		et, parsedEvent, ok := parser.ParseEvent(log)
		// handle unknown event
		if !ok {
			if len(log.Topics) != 0 {
				logger.Warnf("unknown event %s", log.Topics[0])
			}
			return nil
		}

		ctx, span := n.metrics.Tracer().Start(parentCtx, fmt.Sprintf("handleLog-%s", et), trace.WithAttributes(
			attribute.String(metrics.TxHash, log.TxHash.String()),
			attribute.Int(metrics.Origin, chainID),
			attribute.String(metrics.Contract, log.Address.String()),
			attribute.String("block_hash", log.BlockHash.String()),
			attribute.Int64("block_number", int64(log.BlockNumber)),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		switch event := parsedEvent.(type) {
		case *synapsemodule.SynapseModuleModuleMessageSent:
			err = n.handleMessageSent(ctx, event)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("listener failed: %w", err)
	}
	return nil
}

func (n *Node) handleMessageSent(ctx context.Context, event *synapsemodule.SynapseModuleModuleMessageSent) error {
	err := n.db.StoreInterchainTransactionReceived(ctx, *event)
	if err != nil {
		return fmt.Errorf("could not store interchain transaction: %w", err)
	}

	return nil
}
