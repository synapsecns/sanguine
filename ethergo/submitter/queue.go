package submitter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
	"time"
)

// runSelector runs the selector start loop.
func (t *txSubmitterImpl) runSelector(parentCtx context.Context, i int) (shouldExit bool, err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.Start", trace.WithAttributes(attribute.Int("i", i)))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	select {
	case <-ctx.Done():
		return true, fmt.Errorf("context done: %w", ctx.Err())
	case <-time.After(t.GetRetryInterval()):
		err = t.processQueue(ctx)
	case <-t.retryNow:
		err = t.processQueue(ctx)
	}
	return false, err
}

// processQueue processes the queue of transactions.
func (t *txSubmitterImpl) processQueue(parentCtx context.Context) (err error) {
	// TODO: this might be too short of a deadline depending on the number of transactions in the queue
	deadlineCtx, cancel := context.WithDeadline(parentCtx, time.Now().Add(time.Second*60))
	defer cancel()

	ctx, span := t.metrics.Tracer().Start(deadlineCtx, "submitter.ProcessQueue")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// get all the transactions in the queue
	transactions, err := t.db.GetTXS(ctx, t.signer.Address(), nil, db.Pending, db.Replaced)
	if err != nil {
		return fmt.Errorf("could not get transactions: %w", err)
	}

	// fetch txes into a map by chainid.
	sortedTXes := sortTxes(transactions)

	// TODO: parallelize resubmission by chainid, maybe w/ a locker per chain
	g, ctx := errgroup.WithContext(ctx)
	_ = g
	for chainID := range sortedTXes {
		_ = chainID
	}

	return nil
}

func (t *txSubmitterImpl) chainQueue(parentCtx context.Context, chainID uint64, txes []*types.Transaction) (err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.ChainQueue")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
	_ = ctx

	return nil
}
