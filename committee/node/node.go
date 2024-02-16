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
	"github.com/synapsecns/sanguine/committee/p2p"
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
	"math/big"
	"sync"
	"time"
)

type Node struct {
	client              omnirpcClient.RPCClient
	metrics             metrics.Handler
	db                  db.Service
	submitter           submitter.TransactionSubmitter
	signer              signer.Signer
	cfg                 config.Config
	chainListeners      map[int]listener.ContractListener
	interchainContracts map[int]*synapsemodule.SynapseModuleRef
	peerManager         p2p.LibP2PManager
}

var logger = log.Logger("node")

// NewNode creates a new node.
func NewNode(ctx context.Context, handler metrics.Handler, cfg config.Config) (*Node, error) {
	node := &Node{
		metrics: handler,
		cfg:     cfg,
	}

	node.client = omnirpcClient.NewOmnirpcClient(cfg.OmnirpcURL, handler, omnirpcClient.WithCaptureReqRes())

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}

	var decoderContract *synapsemodule.SynapseModuleRef

	// Note: since decoder is a pointer, this will ony deref if chain length is empty & this is called.
	node.db, err = connect.Connect(ctx, dbType, cfg.Database.DSN, handler, func(ctx context.Context, data []byte) (synapsemodule.InterchainInterchainTransaction, error) {
		return decoderContract.DecodeInterchainTransaction(&bind.CallOpts{Context: ctx}, data)
	})
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	node.chainListeners = make(map[int]listener.ContractListener)
	node.interchainContracts = make(map[int]*synapsemodule.SynapseModuleRef)

	for chainID, chainCFG := range cfg.Chains {
		synapseModule := common.HexToAddress(chainCFG.SynapseModuleAddress)
		chainClient, err := node.client.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, node.db, synapseModule, handler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		node.chainListeners[chainID] = chainListener

		node.interchainContracts[chainID], err = synapsemodule.NewSynapseModuleRef(synapseModule, chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not get synapse module ref: %w", err)
		}

		// just use the last chain as the decoder contract
		decoderContract = node.interchainContracts[chainID]
	}

	node.signer, err = signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer: %w", err)
	}

	node.submitter = submitter.NewTransactionSubmitter(handler, node.signer, node.client, node.db.SubmitterDB(), &cfg.SubmitterConfig)
	//node.peerManager, err = p2p.NewLibP2PManager(ctx, node.signer)
	//if err != nil {
	//	return nil, fmt.Errorf("could not get peer manager: %w", err)
	//}

	return node, nil
}

const defaultDBInterval = 3

func (n *Node) createPeerManager(parentCtx context.Context) (err error) {
	ctx, span := n.metrics.Tracer().Start(parentCtx, "createPeerManager")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var mux sync.Mutex
	interchainValidators := make(map[int][]common.Address)

	g, gCtx := errgroup.WithContext(ctx)
	// fo reach chain
	for chainID, contract := range n.interchainContracts {
		chainID := chainID
		contract := contract
		g.Go(func() error {
			var verifiers []common.Address

			i := 0
			for {
				// query all validators
				validator, err := contract.Verifiers(&bind.CallOpts{Context: gCtx}, big.NewInt(int64(i)))
				if err != nil {
					return fmt.Errorf("could not get validator: %w", err)
				}
				// until we hit one that doesn't exist and then we're done
				if validator.String() == (common.Address{}).String() {
					break
				}
				i++
				verifiers = append(verifiers, validator)
			}
			mux.Lock()
			defer mux.Unlock()
			interchainValidators[chainID] = verifiers
			return nil
		})
	}
	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not get validators: %w", err)
	}

	n.peerManager, err = p2p.NewLibP2PManager(ctx, n.signer, n.db)
	if err != nil {
		return fmt.Errorf("could not get peer manager: %w", err)
	}

	return nil
}

func (n *Node) Start(ctx context.Context) error {
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		return n.submitter.Start(ctx)
	})

	g.Go(func() error {
		return n.startChainIndexers(ctx)
	})

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(defaultDBInterval * time.Second):
				err := n.runDBSelector(ctx)
				if err != nil {
					logger.Errorf("could not cleanup: %w", err)
				}
			}
		}
	})

	return nil
}

func (n *Node) runDBSelector(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(defaultDBInterval * time.Second):
			dbItems, err := n.db.GetQuoteResultsByStatus(ctx, db.Seen)
			if err != nil {
				return fmt.Errorf("could not cleanup: %w", err)
			}

			// I should sign and broadcast now.
			for _, request := range dbItems {
				switch request.Status {
				case db.Seen:
					err := n.signAndBroadcast(ctx, request)
					if err != nil {
						logger.Errorf("could not sign and broadcast: %w", err)
					}

				}
			}

			_ = dbItems
		}
	}
}

func (n *Node) signAndBroadcast(ctx context.Context, request db.SignRequest) error {
	// first try to sign
	signedTx, err := n.signer.SignMessage(ctx, request.TransactionID[:], false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	// we'll add this to the struct now, but we're not going to save until it's broadcasted.
	request.Signature[n.signer.Address()] = signer.Encode(signedTx)

	// broadcast the transaction.
	err = n.peerManager.PutSignature(ctx, request.OriginChainID, int(request.Nonce), request.Signature[n.signer.Address()])
	if err != nil {
		return fmt.Errorf("could not broadcast: %w", err)
	}

	// save the transaction.

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
