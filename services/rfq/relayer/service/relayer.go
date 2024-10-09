package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ipfs/go-log"
	"github.com/jellydator/ttlcache/v3"
	"github.com/puzpuzpuz/xsync/v2"
	"github.com/synapsecns/sanguine/core/dbcommon"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/listener"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/cctp-relayer/attestation"
	cctpSql "github.com/synapsecns/sanguine/services/cctp-relayer/db/sql"
	"github.com/synapsecns/sanguine/services/cctp-relayer/relayer"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/guard/guardconfig"
	serviceGuard "github.com/synapsecns/sanguine/services/rfq/guard/service"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relapi"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb/connect"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/semaphore"
)

const maxConcurrentRequests = 15

// Relayer is the core of the relayer application.
type Relayer struct {
	cfg            relconfig.Config
	metrics        metrics.Handler
	db             reldb.Service
	client         omniClient.RPCClient
	chainListeners map[int]listener.ContractListener
	apiServer      *relapi.RelayerAPIServer
	apiClient      rfqAPIClient.AuthenticatedClient
	inventory      inventory.Manager
	quoter         quoter.Quoter
	submitter      submitter.TransactionSubmitter
	signer         signer.Signer
	claimCache     *ttlcache.Cache[common.Hash, bool]
	decimalsCache  *xsync.MapOf[string, *uint8]
	// semaphore is used to limit the number of concurrent requests
	semaphore *semaphore.Weighted
	// handlerMtx is used to synchronize handling of relay requests, keyed on transaction ID
	handlerMtx mapmutex.StringMapMutex
	// balanceMtx is used to synchronize balance requests, keyed on a chainID and tokenAddress pair
	balanceMtx          mapmutex.StringMapMutex
	multicallDispatcher MulticallDispatcher
	otelRecorder        iOtelRecorder
}

var logger = log.Logger("relayer")

// NewRelayer creates a new relayer.
//
// The relayer is the core of the application. It is responsible for starting the listener and quoter event loops.
func NewRelayer(ctx context.Context, metricHandler metrics.Handler, cfg relconfig.Config) (*Relayer, error) {
	omniClient := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omniClient.WithCaptureReqRes())

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
	for chainID := range cfg.GetChains() {
		rfqAddr, err := cfg.GetRFQAddress(chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get rfq address: %w", err)
		}
		chainClient, err := omniClient.GetChainClient(ctx, chainID)
		if err != nil {
			return nil, fmt.Errorf("could not get chain client: %w", err)
		}

		contract, err := fastbridge.NewFastBridgeRef(rfqAddr, chainClient)
		if err != nil {
			return nil, fmt.Errorf("could not create fast bridge contract at address %s: %w", contract.Address(), err)
		}
		startBlock, err := contract.DeployBlock(&bind.CallOpts{Context: ctx})
		if err != nil {
			return nil, fmt.Errorf("could not get deploy block at address %s: %w", contract.Address(), err)
		}
		chainListener, err := listener.NewChainListener(chainClient, store, rfqAddr, uint64(startBlock.Int64()), metricHandler, listener.WithName("relayer"))
		if err != nil {
			return nil, fmt.Errorf("could not get chain listener: %w", err)
		}
		chainListeners[chainID] = chainListener
	}

	sg, err := signerConfig.SignerFromConfig(ctx, cfg.Signer)
	if err != nil {
		return nil, fmt.Errorf("could not get signer: %w", err)
	}
	fmt.Printf("loaded signer with address: %s\n", sg.Address().String())

	sm := submitter.NewTransactionSubmitter(metricHandler, sg, omniClient, store.SubmitterDB(), &cfg.SubmitterConfig)

	im, err := inventory.NewInventoryManager(ctx, omniClient, metricHandler, cfg, sg.Address(), sm, store)
	if err != nil {
		return nil, fmt.Errorf("could not add imanager: %w", err)
	}

	priceFetcher := pricer.NewCoingeckoPriceFetcher(cfg.GetHTTPTimeout())
	fp := pricer.NewFeePricer(cfg, omniClient, priceFetcher, metricHandler)

	apiClient, err := rfqAPIClient.NewAuthenticatedClient(metricHandler, cfg.GetRFQAPIURL(), sg)
	if err != nil {
		return nil, fmt.Errorf("error creating RFQ API client: %w", err)
	}

	q, err := quoter.NewQuoterManager(cfg, metricHandler, im, sg, fp, apiClient)
	if err != nil {
		return nil, fmt.Errorf("could not get quoter")
	}

	apiServer, err := relapi.NewRelayerAPI(ctx, cfg, metricHandler, omniClient, store, sm)
	if err != nil {
		return nil, fmt.Errorf("could not get api server: %w", err)
	}

	otelRecorder, err := newOtelRecorder(metricHandler, store, sg)
	if err != nil {
		return nil, fmt.Errorf("could not get otel recorder: %w", err)
	}

	cache := ttlcache.New[common.Hash, bool](ttlcache.WithTTL[common.Hash, bool](time.Second * 30))
	rel := Relayer{
		db:             store,
		client:         omniClient,
		quoter:         q,
		metrics:        metricHandler,
		claimCache:     cache,
		decimalsCache:  xsync.NewMapOf[*uint8](),
		cfg:            cfg,
		inventory:      im,
		submitter:      sm,
		signer:         sg,
		chainListeners: chainListeners,
		apiServer:      apiServer,
		apiClient:      apiClient,
		semaphore:      semaphore.NewWeighted(maxConcurrentRequests),
		handlerMtx:     mapmutex.NewStringMapMutex(),
		balanceMtx:     mapmutex.NewStringMapMutex(),
		otelRecorder:   otelRecorder,
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
// nolint: cyclop
func (r *Relayer) Start(ctx context.Context) (err error) {
	err = r.inventory.ApproveAllTokens(ctx)
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
		<-ctx.Done()
		r.claimCache.Stop()
	}()

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return fmt.Errorf("could not start db selector: %w", ctx.Err())
			case <-time.After(defaultPostInterval * time.Second):
				err := r.quoter.SubmitAllQuotes(ctx)
				if err != nil {
					return fmt.Errorf("could not start db selector: %w", err)
				}
			}
		}
	})

	if r.cfg.SupportActiveQuoting {
		g.Go(func() error {
			err = r.quoter.SubscribeActiveRFQ(ctx)
			if err != nil {
				return fmt.Errorf("could not subscribe to active RFQ: %w", err)
			}
			return nil
		})
	}

	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(defaultPostInterval * time.Second):
				err := r.runDBSelector(ctx, false, reldb.Seen, reldb.CommittedPending, reldb.CommittedConfirmed, reldb.NotEnoughInventory)
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
				err := r.runDBSelector(ctx, true, reldb.RelayStarted, reldb.RelayCompleted, reldb.ProvePosted)
				if err != nil {
					return fmt.Errorf("could not start db selector: %w", err)
				}
			}
		}
	})

	g.Go(func() error {
		if !r.submitter.Started() {
			err := r.submitter.Start(ctx)
			if err != nil && !errors.Is(err, submitter.ErrSubmitterAlreadyStarted) {
				return fmt.Errorf("could not start submitter: %w", err)
			}
			return nil
		}
		return nil
	})

	g.Go(func() error {
		err := r.apiServer.Run(ctx)
		if err != nil {
			return fmt.Errorf("could not start api server: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		err := r.inventory.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not start inventory manager: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		err = r.multicallDispatcher.Start(ctx)
		if err != nil {
			return fmt.Errorf("could not start multicall dispatcher: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		err = r.startCCTPRelayer(ctx)
		if err != nil {
			return fmt.Errorf("could not start cctp relayer: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		err = r.startGuard(ctx)
		if err != nil {
			return fmt.Errorf("could not start guard: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		err = r.recordMetrics(ctx)
		if err != nil {
			return fmt.Errorf("could not record metrics: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return fmt.Errorf("could not start: %w", err)
	}

	return nil
}

func (r *Relayer) runDBSelector(ctx context.Context, serial bool, matchStatuses ...reldb.QuoteRequestStatus) error {
	interval := r.cfg.GetDBSelectorInterval()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("could not run db selector: %w", ctx.Err())
		case <-time.After(interval):
			// TODO: add context w/ timeout
			// TODO: add trigger
			// TODO: should not fail on error
			err := r.processDB(ctx, serial, matchStatuses...)
			if err != nil {
				return err
			}
		}
	}
}

// startCCTPRelayer starts the CCTP relayer, if a config is specified.
func (r *Relayer) startCCTPRelayer(ctx context.Context) (err error) {
	// only start the CCTP relayer if the config is specified
	cctpCfg := r.cfg.CCTPRelayerConfig
	if cctpCfg == nil {
		return nil
	}

	// build the CCTP relayer
	dbType, err := dbcommon.DBTypeFromString(r.cfg.Database.Type)
	if err != nil {
		return fmt.Errorf("could not get db type: %w", err)
	}
	store, err := cctpSql.Connect(ctx, dbType, r.cfg.Database.DSN, r.metrics)
	if err != nil {
		return fmt.Errorf("could not connect to database: %w", err)
	}
	omnirpcClient := omniClient.NewOmnirpcClient(cctpCfg.BaseOmnirpcURL, r.metrics, omniClient.WithCaptureReqRes())
	attAPI := attestation.NewCircleAPI(cctpCfg.CircleAPIURl)
	cctpRelayer, err := relayer.NewCCTPRelayer(ctx, *cctpCfg, store, omnirpcClient, r.metrics, attAPI, relayer.WithSubmitter(r.submitter))
	if err != nil {
		return fmt.Errorf("could not create cctp relayer: %w", err)
	}

	// run the cctp relayer
	err = cctpRelayer.Run(ctx)
	if err != nil {
		return fmt.Errorf("could not run cctp relayer: %w", err)
	}

	return nil
}

// startGuard starts the guard, if specified.
func (r *Relayer) startGuard(ctx context.Context) (err error) {
	if !r.cfg.UseEmbeddedGuard {
		return nil
	}

	guardCfg := guardconfig.NewGuardConfigFromRelayer(r.cfg)
	guard, err := serviceGuard.NewGuard(ctx, r.metrics, guardCfg, r.submitter)
	if err != nil {
		return fmt.Errorf("could not create guard: %w", err)
	}

	err = guard.Start(ctx)
	if err != nil {
		return fmt.Errorf("could not start guard: %w", err)
	}

	return nil
}

const defaultMetricsInterval = 10

func (r *Relayer) recordMetrics(ctx context.Context) (err error) {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("could not record metrics: %w", ctx.Err())
		case <-time.After(defaultMetricsInterval * time.Second):
			statusCounts, err := r.db.GetStatusCounts(ctx, reldb.Seen, reldb.NotEnoughInventory, reldb.CommittedPending, reldb.CommittedConfirmed, reldb.RelayStarted, reldb.RelayCompleted, reldb.ProvePosting, reldb.ProvePosted, reldb.ClaimPending)
			if err != nil {
				return fmt.Errorf("could not get status counts: %w", err)
			}
			for status, count := range statusCounts {
				r.otelRecorder.RecordStatusCount(status, count)
			}
		}
	}
}

func (r *Relayer) processDB(ctx context.Context, serial bool, matchStatuses ...reldb.QuoteRequestStatus) (err error) {
	ctx, span := r.metrics.Tracer().Start(ctx, "processDB", trace.WithAttributes(
		attribute.Bool("serial", serial),
	))
	defer func() {
		r.recordDBStats(ctx, span)
		metrics.EndSpanWithErr(span, err)
	}()

	requests, err := r.db.GetQuoteResultsByStatus(ctx, matchStatuses...)
	if err != nil {
		return fmt.Errorf("could not get quote results: %w", err)
	}

	wg := sync.WaitGroup{}
	// Obviously, these are only seen.
	for _, req := range requests {
		//nolint: nestif
		if serial {
			// process in serial
			err = r.processRequest(ctx, req)
			if err != nil {
				return fmt.Errorf("could not process request: %w", err)
			}
		} else {
			// process in parallel (new goroutine)
			request := req // capture func literal
			err = r.semaphore.Acquire(ctx, 1)
			if err != nil {
				return fmt.Errorf("could not acquire semaphore: %w", err)
			}
			wg.Add(1)
			go func() {
				defer r.semaphore.Release(1)
				defer wg.Done()
				err = r.processRequest(ctx, request)
				if err != nil {
					logger.Errorf("could not process request: %v", err)
				}
			}()
		}
	}

	// no-op if serial is specified
	wg.Wait()
	return nil
}

func (r *Relayer) processRequest(parentCtx context.Context, request reldb.QuoteRequest) (err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "processRequest", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(request.TransactionID[:])),
		attribute.String("status", request.Status.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// if deadline < now
	if request.Transaction.Deadline.Cmp(big.NewInt(time.Now().Unix())) < 0 && request.Status.Int() < reldb.RelayCompleted.Int() {
		err = r.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.DeadlineExceeded, &request.Status)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
	}

	qr, err := r.requestToHandler(ctx, request)
	if err != nil {
		return fmt.Errorf("could not get request to handler: %w", err)
	}

	err = qr.Handle(ctx, request)
	if err != nil {
		return fmt.Errorf("could not handle request: %w", err)
	}
	return nil
}

func (r *Relayer) recordDBStats(ctx context.Context, span trace.Span) {
	sqlStats, sqlErr := r.db.GetDBStats(ctx)
	if sqlErr != nil {
		span.SetAttributes(attribute.String("sql_error", sqlErr.Error()))
		return
	}
	if sqlStats != nil {
		span.SetAttributes(attribute.Int64("sql_open_conns", int64(sqlStats.OpenConnections)))
		span.SetAttributes(attribute.Int64("sql_in_use_conns", int64(sqlStats.InUse)))
		span.SetAttributes(attribute.Int64("sql_idle_conns", int64(sqlStats.Idle)))
		span.SetAttributes(attribute.Int64("sql_wait_count", sqlStats.WaitCount))
		span.SetAttributes(attribute.String("sql_wait_duration", sqlStats.WaitDuration.String()))
	}
}
