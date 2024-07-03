package service

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/ethergo/listener"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb/connect"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

var logger = log.Logger("guard")

// Guard monitors calls to prove() and verifies them.
type Guard struct {
	cfg            relconfig.Config
	metrics        metrics.Handler
	db             guarddb.Service
	client         omniClient.RPCClient
	chainListeners map[int]listener.ContractListener
	contracts      map[int]*fastbridge.FastBridgeRef
	txSubmitter    submitter.TransactionSubmitter
}

// NewGuard creates a new Guard.
func NewGuard(ctx context.Context, metricHandler metrics.Handler, cfg relconfig.Config) (*Guard, error) {
	omniClient := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omniClient.WithCaptureReqRes())
	chainListeners := make(map[int]listener.ContractListener)

	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}
	store, err := connect.Connect(ctx, dbType, cfg.Database.DSN, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}

	// setup chain listeners
	contracts := make(map[int]*fastbridge.FastBridgeRef)
	for chainID := range cfg.GetChains() {
		rfqAddr, err := cfg.GetRFQAddress(chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get rfq address: %w", err)
		}
		chainClient, err := omniClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		contract, err := fastbridge.NewFastBridgeRef(common.HexToAddress(rfqAddr), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create fast bridge contract: %w", err)
		}
		startBlock, err := contract.DeployBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("could not get deploy block: %w", err)
		}
		chainListener, err := listener.NewChainListener(chainClient, store, common.HexToAddress(rfqAddr), uint64(startBlock.Int64()), metricHandler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		chainListeners[chainID] = chainListener

		// setup FastBridge contract on this chain
		addr, err := cfg.GetRFQAddress(chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get rfq address: %w", err)
		}
		contracts[chainID], err = fastbridge.NewFastBridgeRef(common.HexToAddress(addr), chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create bridge contract: %w", err)
		}
	}

	sg, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer: %w", err)
	}
	fmt.Printf("loaded signer with address: %s\n", sg.Address().String())

	txSubmitter := submitter.NewTransactionSubmitter(metricHandler, sg, omniClient, store.SubmitterDB(), &cfg.SubmitterConfig)

	return &Guard{
		cfg:            cfg,
		metrics:        metricHandler,
		db:             store,
		client:         omniClient,
		chainListeners: chainListeners,
		contracts:      contracts,
		txSubmitter:    txSubmitter,
	}, nil
}

const defaultDBInterval = 5

// Start starts the guard.
func (g *Guard) Start(ctx context.Context) (err error) {
	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error {
		err := g.startChainIndexers(ctx)
		if err != nil {
			return fmt.Errorf("could not start chain indexers: %w", err)
		}
		return nil
	})
	group.Go(func() error {
		err = g.runDBSelector(ctx)
		if err != nil {
			return fmt.Errorf("could not start db selector: %w", err)
		}
		return nil
	})

	err = group.Wait()
	if err != nil {
		return fmt.Errorf("could not wait for group: %w", err)
	}

	return nil
}

func (g *Guard) runDBSelector(ctx context.Context) (err error) {
	interval := g.cfg.GetDBSelectorInterval()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("could not run db selector: %w", ctx.Err())
		case <-time.After(interval):
			err := g.processDB(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func (g *Guard) startChainIndexers(ctx context.Context) error {
	group, ctx := errgroup.WithContext(ctx)

	for chainID := range g.cfg.GetChains() {
		chainID := chainID // capture func literal

		group.Go(func() error {
			err := g.runChainIndexer(ctx, chainID)
			if err != nil {
				return fmt.Errorf("could not runChainIndexer chain indexer for chain %d: %w", chainID, err)
			}
			return nil
		})
	}
	return nil
}

func (g Guard) runChainIndexer(ctx context.Context, chainID int) (err error) {
	chainListener := g.chainListeners[chainID]

	parser, err := fastbridge.NewParser(chainListener.Address())
	if err != nil {
		return fmt.Errorf("could not parse: %w", err)
	}

	err = chainListener.Listen(ctx, func(parentCtx context.Context, log types.Log) (err error) {
		et, parsedEvent, ok := parser.ParseEvent(log)
		// handle unknown event
		if !ok {
			if len(log.Topics) != 0 {
				logger.Warnf("unknown event %s", log.Topics[0])
			}
			return nil
		}

		ctx, span := g.metrics.Tracer().Start(parentCtx, fmt.Sprintf("handleLog-%s", et), trace.WithAttributes(
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
		case *fastbridge.FastBridgeBridgeRequested:
			err = g.handleBridgeRequestedLog(ctx, event, chainID)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		case *fastbridge.FastBridgeBridgeProofProvided:
			err = g.handleProofProvidedLog(ctx, event, chainID)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		case *fastbridge.FastBridgeBridgeProofDisputed:
			err = g.handleProofDisputedLog(ctx, event)
			if err != nil {
				return fmt.Errorf("could not handle request: %w", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("listener failed: %w", err)
	}
	return nil
}

var maxRPCRetryTime = 15 * time.Second

func (g *Guard) handleBridgeRequestedLog(ctx context.Context, req *fastbridge.FastBridgeBridgeRequested, chainID int) (err error) {
	originClient, err := g.client.GetChainClient(ctx, int(chainID))
	if err != nil {
		return fmt.Errorf("could not get correct omnirpc client: %w", err)
	}

	fastBridge, err := fastbridge.NewFastBridgeRef(req.Raw.Address, originClient)
	if err != nil {
		return fmt.Errorf("could not get correct fast bridge: %w", err)
	}

	var bridgeTx fastbridge.IFastBridgeBridgeTransaction
	call := func(ctx context.Context) error {
		bridgeTx, err = fastBridge.GetBridgeTransaction(&bind.CallOpts{Context: ctx}, req.Request)
		if err != nil {
			return fmt.Errorf("could not get bridge transaction: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, call, retry.WithMaxTotalTime(maxRPCRetryTime))
	if err != nil {
		return fmt.Errorf("could not make call: %w", err)
	}

	dbReq := guarddb.BridgeRequest{
		RawRequest:    req.Request,
		TransactionID: req.TransactionId,
		Transaction:   bridgeTx,
	}
	err = g.db.StoreBridgeRequest(ctx, dbReq)
	if err != nil {
		return fmt.Errorf("could not get db: %w", err)
	}
	return nil
}

func (g *Guard) handleProofProvidedLog(ctx context.Context, event *fastbridge.FastBridgeBridgeProofProvided, chainID int) (err error) {
	proven := guarddb.PendingProven{
		Origin:        uint32(chainID),
		TransactionID: event.TransactionId,
		TxHash:        event.TransactionHash,
		Status:        guarddb.ProveCalled,
	}
	err = g.db.StorePendingProven(ctx, proven)
	if err != nil {
		return fmt.Errorf("could not store pending proven: %w", err)
	}

	return nil
}

func (g *Guard) handleProofDisputedLog(ctx context.Context, event *fastbridge.FastBridgeBridgeProofDisputed) (err error) {
	err = g.db.UpdatePendingProvenStatus(ctx, event.TransactionId, guarddb.Disputed)
	if err != nil {
		return fmt.Errorf("could not update pending proven status: %w", err)
	}

	return nil
}

func (g *Guard) processDB(ctx context.Context) (err error) {
	provens, err := g.db.GetPendingProvensByStatus(ctx, guarddb.ProveCalled)
	for _, proven := range provens {
		err := g.handleProveCalled(ctx, proven)
		if err != nil {
			return fmt.Errorf("could not handle prove called: %w", err)
		}
	}

	return nil
}

func (g *Guard) handleProveCalled(ctx context.Context, proven *guarddb.PendingProven) (err error) {
	// first, get the corresponding bridge request
	bridgeRequest, err := g.db.GetBridgeRequestByID(ctx, proven.TransactionID)
	if err != nil {
		return fmt.Errorf("could not get bridge request: %w", err)
	}

	valid, err := g.isProveValid(ctx, proven, bridgeRequest)
	if err != nil {
		return fmt.Errorf("could not check prove validity: %w", err)
	}

	if valid {
		// mark as validated
		err = g.db.UpdatePendingProvenStatus(ctx, proven.TransactionID, guarddb.Validated)
		if err != nil {
			return fmt.Errorf("could not update pending proven status: %w", err)
		}
	} else {
		// trigger dispute
		contract, ok := g.contracts[int(bridgeRequest.Transaction.DestChainId)]
		if !ok {
			return fmt.Errorf("could not get contract for chain: %d", bridgeRequest.Transaction.DestChainId)
		}
		_, err := g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(proven.Origin)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
			tx, err = contract.Dispute(&bind.TransactOpts{Context: ctx}, proven.TransactionID)
			if err != nil {
				return nil, fmt.Errorf("could not dispute: %w", err)
			}

			return tx, nil
		})
		if err != nil {
			return fmt.Errorf("could not dispute: %w", err)
		}

		// mark as dispute pending
		err = g.db.UpdatePendingProvenStatus(ctx, proven.TransactionID, guarddb.DisputePending)
		if err != nil {
			return fmt.Errorf("could not update pending proven status: %w", err)
		}
	}

	return nil
}

func (g *Guard) isProveValid(ctx context.Context, proven *guarddb.PendingProven, bridgeRequest *guarddb.BridgeRequest) (bool, error) {
	// get the receipt for this tx on dest chain
	chainClient, err := g.client.GetChainClient(ctx, int(bridgeRequest.Transaction.DestChainId))
	if err != nil {
		return false, fmt.Errorf("could not get chain client: %w", err)
	}
	receipt, err := chainClient.TransactionReceipt(ctx, proven.TxHash)
	if err != nil {
		return false, fmt.Errorf("could not get receipt: %w", err)
	}
	addr, err := g.cfg.GetRFQAddress(int(bridgeRequest.Transaction.DestChainId))
	if err != nil {
		return false, fmt.Errorf("could not get rfq address: %w", err)
	}
	parser, err := fastbridge.NewParser(common.HexToAddress(addr))
	if err != nil {
		return false, fmt.Errorf("could not get parser: %w", err)
	}

	for _, log := range receipt.Logs {
		_, parsedEvent, ok := parser.ParseEvent(*log)
		if !ok {
			continue
		}

		switch event := parsedEvent.(type) {
		case *fastbridge.FastBridgeBridgeRelayed:
			return relayMatchesBridgeRequest(event, bridgeRequest), nil
		}
	}

	return false, nil
}

func relayMatchesBridgeRequest(event *fastbridge.FastBridgeBridgeRelayed, bridgeRequest *guarddb.BridgeRequest) bool {
	//TODO: is this exhaustive?
	if event.TransactionId != bridgeRequest.TransactionID {
		return false
	}
	if event.OriginAmount.Cmp(bridgeRequest.Transaction.OriginAmount) != 0 {
		return false
	}
	if event.DestAmount.Cmp(bridgeRequest.Transaction.DestAmount) != 0 {
		return false
	}
	if event.OriginChainId != bridgeRequest.Transaction.OriginChainId {
		return false
	}
	if event.To != bridgeRequest.Transaction.DestRecipient {
		return false
	}
	if event.OriginToken != bridgeRequest.Transaction.OriginToken {
		return false
	}
	if event.DestToken != bridgeRequest.Transaction.DestToken {
		return false
	}
	return true
}
