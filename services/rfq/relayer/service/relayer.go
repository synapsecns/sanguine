package service

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ipfs/go-log"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/metrics"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/listener"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/connect"
	"golang.org/x/sync/errgroup"
	"math/big"
	"time"
)

// Relayer is the core of the relayer application.
type Relayer struct {
	cfg            relconfig.Config
	metrics        metrics.Handler
	db             reldb.Service
	client         omnirpcClient.RPCClient
	chainListeners map[int]listener.ContractListener
	inventory      inventory.InventoryManager
	quoter         *quoter.QuoterManager
	submitter      submitter.TransactionSubmitter
	signer         signer.Signer
	claimCache     *ttlcache.Cache[common.Hash, bool]
}

var logger = log.Logger("relayer")

// NewRelayer creates a new relayer.
//
// The relayer is the core of the application. It is responsible for starting the listener and quoter event loops.
func NewRelayer(ctx context.Context, metricHandler metrics.Handler, cfg relconfig.Config) (*Relayer, error) {
	omniClient := omnirpcClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omnirpcClient.WithCaptureReqRes())

	// TODO: pull from config
	dbType, err := dbcommon.DBTypeFromString(cfg.Database.Type)
	if err != nil {
		return nil, fmt.Errorf("could not get db type: %w", err)
	}

	store, err := connect.Connect(ctx, dbType, cfg.Database.DSN, metricHandler)
	if err != nil {
		return nil, fmt.Errorf("could not make db: %w", err)
	}
	chainListeners := make(map[int]listener.ContractListener)

	// setup chain listeners
	for chainID, chainCFG := range cfg.Bridges {
		// TODO: consider getter for this convert step
		bridge := common.HexToAddress(chainCFG.Bridge)
		chainClient, err := omniClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		chainListener, err := listener.NewChainListener(chainClient, store, bridge, metricHandler)
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		chainListeners[chainID] = chainListener
	}

	sg, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer")
	}

	im, err := inventory.NewInventoryManager(ctx, omniClient, metricHandler, cfg, sg.Address(), store)
	if err != nil {
		return nil, fmt.Errorf("could not add imanager: %w", err)
	}

	q, err := quoter.NewQuoterManager(ctx, cfg.QuotableTokens, im, cfg.RfqAPIURL, sg)
	if err != nil {
		return nil, fmt.Errorf("could not get quoter")
	}

	sm := submitter.NewTransactionSubmitter(metricHandler, sg, omniClient, store.SubmitterDB(), &cfg.SubmitterConfig)

	cache := ttlcache.New[common.Hash, bool](ttlcache.WithTTL[common.Hash, bool](time.Second * 30))
	rel := Relayer{
		db:             store,
		client:         omniClient,
		quoter:         q,
		metrics:        metricHandler,
		claimCache:     cache,
		cfg:            cfg,
		inventory:      im,
		submitter:      sm,
		signer:         sg,
		chainListeners: chainListeners,
	}
	return &rel, nil
}

const defaultPostInterval = 1

// Start starts the relayer.
//
// This will:
// 1. Check if approvals need to be issued on chain & issue them if needed. This allows
// _pullToken to function correctly.
// 2. Start the chain parser: This will listen to and process any events on chain
// 3. Start the db selector: This will check the db for any requests that need to be processed.
// 4. Start the submitter: This will submit any transactions that need to be submitted.
func (r *Relayer) Start(ctx context.Context) error {
	err := r.inventory.ApproveAllTokens(ctx, r.submitter)
	if err != nil {
		return fmt.Errorf("could not approve all tokens: %w", err)
	}

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		err := r.startChainIndexers(ctx)
		if err != nil {
			return fmt.Errorf("could not start chain parser: %w", err)
		}
		return nil
	})

	go r.claimCache.Start()
	go func() {
		select {
		case <-ctx.Done():
			r.claimCache.Stop()
		}
	}()

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return fmt.Errorf("could not start db selector: %w", ctx.Err())
			case <-time.After(defaultPostInterval * time.Second):
				err := r.quoter.SubmitAllQuotes()
				if err != nil {
					return fmt.Errorf("could not start db selector: %w", err)
				}
			}
		}
	})

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(defaultPostInterval * time.Second):
				err := r.runDBSelector(ctx)
				if err != nil {
					return fmt.Errorf("could not start db selector: %w", err)
				}
			}
		}
	})

	g.Go(func() error {
		err := r.submitter.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not start submitter: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not start: %w", err)
	}

	return nil
}

// TODO: make this configurable.
const dbSelectorInterval = 1

func (r *Relayer) runDBSelector(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("could not run db selector: %w", ctx.Err())
		case <-time.After(dbSelectorInterval * time.Second):
			// TODO: add context w/ timeout
			// TODO: add trigger
			err := r.processDB(ctx)
			if err != nil {
				return err
			}
		}
	}
}

func (r *Relayer) processDB(ctx context.Context) error {
	requests, err := r.db.GetQuoteResultsByStatus(ctx, reldb.Seen, reldb.CommittedPending, reldb.CommittedConfirmed, reldb.RelayCompleted, reldb.ProvePosted)
	if err != nil {
		return nil
	}
	// Obviously, these are only seen.
	for _, request := range requests {
		// if deadline < now
		if request.Transaction.Deadline.Cmp(big.NewInt(time.Now().Unix())) < 0 && request.Status.Int() < reldb.RelayCompleted.Int() {
			err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionId, reldb.DeadlineExceeded)
		}

		qr, err := r.requestToHandler(ctx, request)
		if err != nil {
			return fmt.Errorf("could not get request to handler: %w", err)
		}

		err = qr.Handle(ctx, request)
		if err != nil {
			return fmt.Errorf("could not handle request: %w", err)
		}
	}
	return nil
}
