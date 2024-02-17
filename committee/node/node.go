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
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"math/big"
	"strings"
	"sync"
	"time"
)

type Node struct {
	client               omnirpcClient.RPCClient
	metrics              metrics.Handler
	db                   db.Service
	submitter            submitter.TransactionSubmitter
	signer               signer.Signer
	cfg                  config.Config
	chainListeners       map[int]listener.ContractListener
	interchainContracts  map[int]*synapsemodule.SynapseModuleRef
	peerManager          p2p.LibP2PManager
	interchainValidators map[int][]common.Address
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

	node.db, err = connect.Connect(ctx, dbType, cfg.Database.DSN, handler)
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
	}

	node.signer, err = signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer: %w", err)
	}

	node.submitter = submitter.NewTransactionSubmitter(handler, node.signer, node.client, node.db.SubmitterDB(), &cfg.SubmitterConfig)

	// this can't be done in the constructor because we need to wait for the peer manager to be created.
	err = node.createPeerManager(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not create peer manager: %w", err)
	}

	return node, nil
}

func (n *Node) IPFSAddress() (addresses []string) {
	for _, addr := range n.peerManager.Host().Addrs() {
		addresses = append(addresses, fmt.Sprintf("%s/p2p/%s", addr, n.peerManager.Host().ID()))
	}
	return addresses
}

func (n *Node) Address() common.Address {
	return n.signer.Address()
}

const defaultDBInterval = 3

func (n *Node) createPeerManager(ctx context.Context) (err error) {
	n.peerManager, err = p2p.NewLibP2PManager(ctx, n.signer, n.db)
	if err != nil {
		return fmt.Errorf("could not get peer manager: %w", err)
	}

	return nil
}

func (n *Node) setValidators(parentCtx context.Context) (err error) {
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
				if err != nil && strings.Contains(err.Error(), "execution reverted") {
					break
				}
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

	n.interchainValidators = interchainValidators
	return nil
}

func (n *Node) Start(parentContext context.Context) error {
	g, ctx := errgroup.WithContext(parentContext)

	err := n.setValidators(ctx)
	if err != nil {
		return fmt.Errorf("could not set validators: %w", err)
	}

	err = n.startP2P(ctx)
	if err != nil {
		return fmt.Errorf("could not start p2p: %w", err)
	}

	g.Go(func() error {
		// nolint: errcheck
		return n.submitter.Start(ctx)
	})

	g.Go(func() error {
		// nolint: errcheck
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

	// TODO: call g.wait, return error on error

	return nil
}

func (n *Node) startP2P(ctx context.Context) error {
	err := n.peerManager.Start(ctx, n.cfg.BootstarpPeers)
	if err != nil {
		return fmt.Errorf("could not start peer manager: %w", err)
	}

	uniqueValidators := map[common.Address]struct{}{}
	for _, validators := range n.interchainValidators {
		for _, validator := range validators {
			uniqueValidators[validator] = struct{}{}
		}
	}

	var allValidators []common.Address
	for validator := range uniqueValidators {
		allValidators = append(allValidators, validator)
	}

	err = n.peerManager.AddValidators(ctx, allValidators...)
	if err != nil {
		return fmt.Errorf("could not add validators: %w", err)
	}

	return nil
}

func (n *Node) runDBSelector(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-time.After(defaultDBInterval * time.Second):
			statusQueries := []db.SynapseRequestStatus{db.Seen}
			if n.cfg.ShouldRelay {
				statusQueries = append(statusQueries, db.Signed)
			}

			dbItems, err := n.db.GetQuoteResultsByStatus(ctx, statusQueries...)
			if err != nil {
				return fmt.Errorf("could not cleanup: %w", err)
			}

			for _, request := range dbItems {
				switch request.Status {
				case db.Seen:
					err := n.signAndBroadcast(ctx, request)
					if err != nil {
						logger.Errorf("could not sign and broadcast: %v", err)
					}

					fmt.Printf("original signed entry hash: %s\n", request.SignedEntryHash.String())
				case db.Signed:
					err := n.submit(ctx, request)
					if err != nil {
						logger.Errorf("could not submit: %w", err)
					}
				default:
					panic("unhandled default case")
				}
			}
		}
	}
}

func (n *Node) submit(ctx context.Context, request db.SignRequest) error {
	contract := n.interchainContracts[int(request.DestChainId.Int64())]
	threshold, err := contract.RequiredThreshold(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get threshold: %w", err)
	}

	var signatures [][]byte
	for _, validator := range n.interchainValidators[int(request.DestChainId.Int64())] {
		signature, err := n.peerManager.GetSignature(ctx, validator, int(request.InterchainEntry.SrcChainId.Int64()), int(request.InterchainEntry.WriterNonce.Uint64()))
		if err != nil {
			logger.Errorf("could not get signature for peer %s message: %w", validator, err)
		}
		signatures = append(signatures, signature)
	}

	if len(signatures) < int(threshold.Uint64()) {
		return fmt.Errorf("not enough signatures")
	}

	nonce, err := n.submitter.SubmitTransaction(ctx, request.DestChainId, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		// should be request.InterchainEntry
		return contract.VerifyEntry(transactor, request.InterchainEntry, signatures)
	})

	go func() {
		for {
			time.Sleep(time.Second * 5)

			yo, err := n.submitter.GetSubmissionStatus(ctx, request.DestChainId, nonce)
			if err != nil {
				logger.Errorf("could not get submission status: %w", err)
			}

			fmt.Println("fuck")
			fmt.Printf("tx hash: %s\n", yo.TxHash().String())
			fmt.Printf("new signed hash: %s \n", request.SignedEntryHash.String())
			fmt.Println("you")
		}
	}()

	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}

	err = n.db.UpdateSignRequestStatus(ctx, request.SignedEntryHash, db.Broadcast)
	if err != nil {
		return fmt.Errorf("could not update status: %w", err)
	}

	return nil
}

func (n *Node) signAndBroadcast(ctx context.Context, request db.SignRequest) error {
	// first try to sign
	signedTx, err := n.signer.SignMessage(ctx, request.SignedEntryHash[:], false)
	if err != nil {
		return fmt.Errorf("could not sign message: %w", err)
	}

	// broadcast the transaction.
	tweakedSig := signer.NewSignature(new(big.Int).Add(big.NewInt(27), signedTx.V()), signedTx.R(), signedTx.S())

	err = n.peerManager.PutSignature(ctx, int(request.InterchainEntry.SrcChainId.Int64()), int(request.InterchainEntry.WriterNonce.Uint64()), signer.Encode(tweakedSig))
	if err != nil {
		return fmt.Errorf("could not broadcast: %w", err)
	}

	// save the transaction.
	err = n.db.UpdateSignRequestStatus(ctx, request.SignedEntryHash, db.Signed)
	if err != nil {
		return fmt.Errorf("could not update status: %w", err)
	}

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

	parser, err := synapsemodule.NewParser(chainListener.Address())
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
		case *synapsemodule.SynapseModuleVerificationRequested:
			err = n.handleMessageSent(ctx, event)
		case *synapsemodule.SynapseModuleEntryVerified:
			err = n.db.UpdateSignRequestStatus(ctx, event.SignableEntryHash, db.Completed)
		}
		// stop the world.
		if err != nil {
			return fmt.Errorf("could not handle event: %w", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("listener failed: %w", err)
	}
	return nil
}

func (n *Node) handleMessageSent(ctx context.Context, event *synapsemodule.SynapseModuleVerificationRequested) error {
	err := n.db.StoreInterchainTransactionReceived(ctx, *event)
	if err != nil {
		return fmt.Errorf("could not store interchain transaction: %w", err)
	}

	return nil
}
