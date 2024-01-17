package service

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
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
	// RelayerAdress is the relayer RelayerAdress
	RelayerAdress common.Address
	// metrics is the metrics handler.
	metrics metrics.Handler
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

	qr := &QuoteRequestHandler{
		Origin:        *origin,
		Dest:          *dest,
		db:            r.db,
		Inventory:     r.inventory,
		Quoter:        r.quoter,
		handlers:      make(map[reldb.QuoteRequestStatus]Handler),
		metrics:       r.metrics,
		RelayerAdress: r.signer.Address(),
		claimCache:    r.claimCache,
	}

	qr.handlers[reldb.Seen] = r.deadlineMiddleware(qr.handleSeen)
	qr.handlers[reldb.CommittedPending] = r.deadlineMiddleware(qr.handleCommitPending)
	qr.handlers[reldb.CommittedConfirmed] = r.deadlineMiddleware(qr.handleCommitConfirmed)
	// no more need for deadline middleware now, we already relayed.
	qr.handlers[reldb.RelayCompleted] = qr.handleRelayCompleted
	qr.handlers[reldb.ProvePosted] = qr.handleProofPosted
	// TODO: we probably want a claim complete state once we've seen that event on chain

	// error handlers only
	qr.handlers[reldb.NotEnoughInventory] = r.deadlineMiddleware(qr.handleNotEnoughInventory)

	return qr, nil
}

func (r *Relayer) deadlineMiddleware(next func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error) func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error {
	return func(ctx context.Context, span trace.Span, req reldb.QuoteRequest) error {
		// apply deadline buffer
		almostNow := time.Now().Add(-r.cfg.GetDeadlineBuffer(int(req.Transaction.DestChainId)))

		// if deadline < now, we don't even have to bother calling the underlying function
		if req.Transaction.Deadline.Cmp(big.NewInt(almostNow.Unix())) < 0 {
			err := r.db.UpdateQuoteRequestStatus(ctx, req.TransactionID, reldb.DeadlineExceeded)
			if err != nil {
				return fmt.Errorf("could not update request status: %w", err)
			}
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

	//nolint: wrapcheck
	return chain.NewChain(ctx, chainClient, common.HexToAddress(r.cfg.GetChains()[id].Bridge), r.chainListeners[id], r.submitter)
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

	return q.handlers[request.Status](ctx, span, request)
}
