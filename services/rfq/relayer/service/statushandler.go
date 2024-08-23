package service

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/mapmutex"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/limiter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// TODO: everything in this file should be moved to it's own module, at least as an interface

// QuoteRequestHandler is the helper for a quote request.
// lowercase fields are private, uppercase are public.
// the plan is to move this out of relayer which is when this distinction will matter.
type QuoteRequestHandler struct {
	// Origin is the origin chain.
	Origin chain.Chain
	// Dest is the destination chain.
	Dest chain.Chain
	// db is the database.
	db reldb.Service
	// Inventory is the inventory.
	Inventory inventory.Manager
	// Quoter is the quoter.
	Quoter quoter.Quoter
	// handlers is the map of handlers.
	handlers map[reldb.QuoteRequestStatus]Handler
	// claimCache is the cache of claims used for figuring out when we should retry the claim method.
	claimCache *ttlcache.Cache[common.Hash, bool]
	// RelayerAddress is the relayer RelayerAddress
	RelayerAddress common.Address
	// metrics is the metrics handler.
	metrics metrics.Handler
	// apiClient is used to get acks before submitting a relay transaction.
	apiClient client.AuthenticatedClient
	// mutexMiddlewareFunc is used to wrap the handler in a mutex middleware.
	// this should only be done if Handling, not forwarding.
	mutexMiddlewareFunc func(func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error) func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error
	// handlerMtx is the mutex for relaying.
	handlerMtx mapmutex.StringMapMutex
	// limiter is the rate limiter.
	limiter limiter.Limiter
	// tokenNames is the map of addresses to token names
	tokenNames map[string]relconfig.TokenConfig
	// balanceMtx is the mutex for balances.
	balanceMtx mapmutex.StringMapMutex
}

func getBalanceMtxKey(chainID uint32, token common.Address) string {
	return fmt.Sprintf("%d-%s", chainID, token.Hex())
}

// Handler is the handler for a quote request.
type Handler func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error

func (r *Relayer) requestToHandler(ctx context.Context, req reldb.QuoteRequest) (*QuoteRequestHandler, error) {
	origin, err := r.chainIDToChain(ctx, req.Transaction.OriginChainId)
	if err != nil {
		return nil, fmt.Errorf("could not get origin chain: %w", err)
	}

	dest, err := r.chainIDToChain(ctx, req.Transaction.DestChainId)
	if err != nil {
		return nil, fmt.Errorf("could not get dest chain: %w", err)
	}

	originTokens, err := r.cfg.GetTokens(req.Transaction.OriginChainId)
	if err != nil {
		return nil, fmt.Errorf("could not get tokens: %w", err)
	}

	qr := &QuoteRequestHandler{
		Origin:              *origin,
		Dest:                *dest,
		db:                  r.db,
		Inventory:           r.inventory,
		Quoter:              r.quoter,
		handlers:            make(map[reldb.QuoteRequestStatus]Handler),
		metrics:             r.metrics,
		RelayerAddress:      r.signer.Address(),
		claimCache:          r.claimCache,
		apiClient:           r.apiClient,
		mutexMiddlewareFunc: r.mutexMiddleware,
		handlerMtx:          r.handlerMtx,
		// TODO: this should be configurable
		limiter: limiter.NewRateLimiter(
			r.cfg,
			r.chainListeners[int(req.Transaction.OriginChainId)],
			r.quoter,
			r.metrics,
			originTokens,
		),
		tokenNames: originTokens,
		balanceMtx: mapmutex.NewStringMapMutex(),
	}

	// wrap in deadline middleware since the relay has not yet happened
	qr.handlers[reldb.Seen] = r.deadlineMiddleware(r.gasMiddleware(qr.handleSeen))
	qr.handlers[reldb.CommittedPending] = r.deadlineMiddleware(r.gasMiddleware(qr.handleCommitPending))
	qr.handlers[reldb.CommittedConfirmed] = r.deadlineMiddleware(r.gasMiddleware(qr.handleCommitConfirmed))

	// no-op edge case, but we still want to check the deadline
	qr.handlers[reldb.RelayStarted] = r.deadlineMiddleware(func(_ context.Context, _ trace.Span, _ reldb.QuoteRequest) error { return nil })

	// no more need for deadline middleware now, we already relayed
	qr.handlers[reldb.RelayCompleted] = qr.handleRelayCompleted
	qr.handlers[reldb.ProvePosted] = qr.handleProofPosted

	// error handlers only
	qr.handlers[reldb.NotEnoughInventory] = r.deadlineMiddleware(qr.handleNotEnoughInventory)

	return qr, nil
}

func (r *Relayer) mutexMiddleware(next func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error) func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error {
	return func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) (err error) {
		unlocker, ok := r.handlerMtx.TryLock(hexutil.Encode(req.TransactionID[:]))
		if !ok {
			span.SetAttributes(
				attribute.Bool("locked", true),
				attribute.StringSlice("current_locks", r.handlerMtx.Keys()),
			)
			return nil
		}
		defer unlocker.Unlock()

		// make sure the status has not changed since we last saw it
		dbReq, err := r.db.GetQuoteRequestByID(ctx, req.TransactionID)
		if err != nil {
			return fmt.Errorf("could not get request: %w", err)
		}
		if dbReq.Status != req.Status {
			span.SetAttributes(
				attribute.Bool("status_changed", true),
				attribute.String("db_status", dbReq.Status.String()),
				attribute.String("handler_status", req.Status.String()),
			)
			return nil
		}

		return next(ctx, span, req)
	}
}

func (r *Relayer) deadlineMiddleware(next func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error) func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error {
	return func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error {
		// apply deadline buffer
		buffer, err := r.cfg.GetDeadlineBuffer(int(req.Transaction.DestChainId))
		if err != nil {
			return fmt.Errorf("could not get deadline buffer: %w", err)
		}
		almostNow := time.Now().Add(-buffer)

		// if deadline < now, we don't even have to bother calling the underlying function
		if req.Transaction.Deadline.Cmp(big.NewInt(almostNow.Unix())) < 0 {
			err := r.db.UpdateQuoteRequestStatus(ctx, req.TransactionID, reldb.DeadlineExceeded, &req.Status)
			if err != nil {
				return fmt.Errorf("could not update request status: %w", err)
			}
			return nil
		}

		return next(ctx, span, req)
	}
}

// gasMiddleware checks that we have sufficient gas to process a request on origin and destination.
func (r *Relayer) gasMiddleware(next func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error) func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error {
	return func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) (err error) {
		var sufficientGasOrigin, sufficientGasDest bool

		defer func() {
			span.SetAttributes(
				attribute.Bool("sufficient_gas_origin", sufficientGasOrigin),
				attribute.Bool("sufficient_gas_dest", sufficientGasDest),
			)
		}()

		sufficientGasOrigin, err = r.inventory.HasSufficientGas(ctx, int(req.Transaction.OriginChainId), nil)
		if err != nil {
			return fmt.Errorf("could not check gas on origin: %w", err)
		}

		// on destination, we need to check transactor.Value as well if we are dealing with ETH
		// However, all requests with statuses CommittedPending, CommittedConfirmed and RelayStarted are considered
		// in-flight and their respective amounts are already deducted from the inventory: see Manager.GetCommittableBalances().
		// Therefore, we only need to check the gas value for requests with all the other statuses.
		isInFlight := req.Status == reldb.CommittedPending || req.Status == reldb.CommittedConfirmed || req.Status == reldb.RelayStarted
		var destGasValue *big.Int
		if req.Transaction.DestToken == util.EthAddress && !isInFlight {
			destGasValue = req.Transaction.DestAmount
			span.SetAttributes(attribute.String("dest_gas_value", destGasValue.String()))
		}
		sufficientGasDest, err = r.inventory.HasSufficientGas(ctx, int(req.Transaction.DestChainId), destGasValue)
		if err != nil {
			return fmt.Errorf("could not check gas on dest: %w", err)
		}

		if !sufficientGasOrigin || !sufficientGasDest {
			return nil
		}

		return next(ctx, span, req)
	}
}

func (r *Relayer) chainIDToChain(ctx context.Context, chainID uint32) (*chain.Chain, error) {
	id := int(chainID)

	chainClient, err := r.client.GetChainClient(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("could not get origin client: %w", err)
	}

	chain, err := chain.NewChain(ctx, r.cfg, chainClient, r.chainListeners[id], r.submitter)
	if err != nil {
		return nil, fmt.Errorf("could not create chain: %w", err)
	}
	return chain, nil
}

// shouldCheckClaim checks if we should check the claim method.
// if so it checks the claim method and updates the cache.
func (q *QuoteRequestHandler) shouldCheckClaim(request reldb.QuoteRequest) bool {
	// we use claim cache to make sure we don't hit the rpc to check to often
	if q.claimCache.Has(request.TransactionID) {
		return false
	}

	q.claimCache.Set(request.TransactionID, true, 30*time.Second)
	return true
}

// Handle handles a quote request.
// Note: this will panic if no method is available. This is done on purpose.
func (q *QuoteRequestHandler) Handle(ctx context.Context, request reldb.QuoteRequest) (err error) {
	ctx, span := q.metrics.Tracer().Start(ctx, fmt.Sprintf("handle-%s", request.Status.String()), trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(request.TransactionID[:])),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// we're handling and not forwarding, so we need to wrap the handler in a mutex middleware
	handler := q.mutexMiddlewareFunc(q.handlers[request.Status])
	return handler(ctx, span, request)
}

// Forward forwards a quote request.
// this ignores the mutex middleware.
func (q *QuoteRequestHandler) Forward(ctx context.Context, request reldb.QuoteRequest) (err error) {
	txID := hexutil.Encode(request.TransactionID[:])
	ctx, span := q.metrics.Tracer().Start(ctx, fmt.Sprintf("forward-%s", request.Status.String()), trace.WithAttributes(
		attribute.String("transaction_id", txID),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// sanity check to make sure that the lock is already acquired for this tx
	_, ok := q.handlerMtx.TryLock(txID)
	if ok {
		panic(fmt.Sprintf("attempted forward while lock was not acquired for tx: %s", txID))
	}

	return q.handlers[request.Status](ctx, span, request)
}
