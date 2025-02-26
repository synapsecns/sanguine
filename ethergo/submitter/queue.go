package submitter

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
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
// TODO: add a way to process a confirmation queue.
func (t *txSubmitterImpl) processQueue(parentCtx context.Context) (err error) {
	// TODO: this might be too short of a deadline depending on the number of pendingTxes in the queue
	deadlineCtx, cancel := context.WithTimeout(parentCtx, time.Second*60)
	defer cancel()

	ctx, span := t.metrics.Tracer().Start(deadlineCtx, "submitter.ProcessQueue")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: parallelize resubmission by chainid, maybe w/ a locker per chain
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		err := t.processConfirmedQueue(ctx)
		if err != nil {
			span.AddEvent("processConfirmedQueue error", trace.WithAttributes(
				attribute.String("error", err.Error())))
		}
	}()

	pendingChainIDs, err := t.db.GetChainIDsByStatus(ctx, t.signer.Address(), db.Stored, db.Pending, db.FailedSubmit, db.Submitted)
	if err != nil {
		return fmt.Errorf("could not get pendingChainIDs: %w", err)
	}

	t.distinctChainIDMux.RLock()
	noOpChainIDs := outersection(pendingChainIDs, t.distinctChainIDs)
	t.distinctChainIDMux.RUnlock()

	pendingChainIDs64 := make([]int64, len(pendingChainIDs))
	for i, chainID := range pendingChainIDs {
		pendingChainIDs64[i] = chainID.Int64()
	}
	span.SetAttributes(attribute.Int64Slice("pending_chain_ids", pendingChainIDs64))

	wg.Add(len(pendingChainIDs))

	for _, chainID := range pendingChainIDs {
		go func(chainID *big.Int) {
			defer wg.Done()

			// get all the pendingTxes in the queue
			pendingTxes, err := t.db.GetTXS(ctx, t.signer.Address(), chainID, db.WithStatuses(db.Stored, db.Pending, db.FailedSubmit, db.Submitted))
			if err != nil {
				span.AddEvent("could not get pendingTxes", trace.WithAttributes(
					attribute.String("error", err.Error()), attribute.Int64("chainID", chainID.Int64()),
				))
				return
			}

			err = t.chainPendingQueue(ctx, chainID, pendingTxes)
			if err != nil {
				span.AddEvent("chainPendingQueue error", trace.WithAttributes(
					attribute.String("error", err.Error()), attribute.Int64("chainID", chainID.Int64())))
			}
		}(chainID)
	}

	for _, chainID := range noOpChainIDs {
		t.otelRecorder.RecordOldestPendingTx(uint32(chainID.Int64()), 0)
		t.otelRecorder.RecordNumPendingTxes(uint32(chainID.Int64()), 0)

		if !t.otelRecorder.HasNonceForChain(uint32(chainID.Int64())) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				evmClient, err := t.fetcher.GetClient(ctx, chainID)
				if err != nil {
					logger.Warn("could not get client", "error", err)
					return
				}
				nonce, err := evmClient.NonceAt(ctx, t.signer.Address(), nil)
				if err != nil {
					logger.Warn("could not get nonce", "error", err)
					return
				}
				t.otelRecorder.RecordNonceForChain(uint32(chainID.Int64()), nonce)
			}()
		}

		if !t.otelRecorder.HasGasBalanceForChain(uint32(chainID.Int64())) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				evmClient, err := t.fetcher.GetClient(ctx, chainID)
				if err != nil {
					logger.Warn("could not get client", "error", err)
					return
				}
				balance, err := evmClient.BalanceAt(ctx, t.signer.Address(), nil)
				if err != nil {
					logger.Warn("could not get balance", "error", err)
					return
				}
				t.otelRecorder.RecordGasBalanceForChain(uint32(chainID.Int64()), balance)
			}()
		}
	}

	wg.Wait()

	return nil
}

const maxTxesPerChain = 100

func (t *txSubmitterImpl) processConfirmedQueue(parentCtx context.Context) (err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.processConfirmedQueue")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	txs, err := t.db.GetAllTXAttemptByStatus(ctx, t.signer.Address(), nil, db.WithMaxResults(1000), db.WithStatuses(db.ReplacedOrConfirmed))
	if err != nil {
		return fmt.Errorf("could not get txs: %w", err)
	}

	sortedTXsByChainID := sortTxesByChainID(txs, maxTxesPerChain)

	t.distinctChainIDMux.RLock()
	noOpChainIDs := outersection(mapToBigIntSlice(sortedTXsByChainID), t.distinctChainIDs)
	t.distinctChainIDMux.RUnlock()

	var wg sync.WaitGroup
	wg.Add(len(sortedTXsByChainID))

	for chainID := range sortedTXsByChainID {
		go func(chainID uint64) {
			defer wg.Done()
			err := t.chainConfirmQueue(ctx, new(big.Int).SetUint64(chainID), sortedTXsByChainID[chainID])
			if err != nil {
				span.AddEvent("chainPendingQueue error", trace.WithAttributes(
					attribute.String("error", err.Error()), attribute.Int64("chainID", int64(chainID))))
			}
		}(chainID)
	}

	for _, chainID := range noOpChainIDs {
		t.otelRecorder.RecordConfirmedQueue(uint32(chainID.Int64()), 0)
	}

	wg.Wait()
	return nil
}

func (t *txSubmitterImpl) chainConfirmQueue(parentCtx context.Context, chainID *big.Int, txes []db.TX) (err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.chainConfirmQueue")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	t.otelRecorder.RecordConfirmedQueue(uint32(chainID.Int64()), len(txes))

	// chainClient is the client for the chain we're working on
	chainClient, err := t.fetcher.GetClient(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get client: %w", err)
	}

	nonceMap := groupTxesByNonce(txes)
	for nonce := range nonceMap {
		err = t.checkAndSetConfirmation(ctx, chainClient, nonceMap[nonce])
		if err != nil {
			return fmt.Errorf("could not check and set confirmation: %w", err)
		}
	}
	return nil
}

// checkAndSetConfirmation checks if the tx is confirmed and sets the status accordingly.
// note: assumes all txes have the same nonce.
func (t *txSubmitterImpl) checkAndSetConfirmation(ctx context.Context, chainClient client.EVM, txes []db.TX) error {
	// nothing do to
	if len(txes) == 0 {
		return nil
	}

	// we're going to take every tx for this nonce and get a receipt for it
	// because there can only be one transaction per nonce, as soon as we know which one has a receipt, we can assume all the
	// others are replaced.
	//
	// There are a few constraints on the logic below as it's currently implemented. Namely that the number of txes
	// can't be bigger than batch size.
	//
	// the other constraint is we treat all errors as "tx not found" errors. This is fine because we only store txes in cases
	//
	calls := make([]w3types.Caller, len(txes))
	receipts := make([]types.Receipt, len(txes))
	for i := range calls {
		calls[i] = eth.TxReceipt(txes[i].Hash()).Returns(&receipts[i])
	}

	err := chainClient.BatchWithContext(ctx, calls...)
	foundSuccessfulTX := false
	if err != nil {
		// there's no way around this type inference
		//nolint: errorlint
		callErr, ok := err.(w3.CallErrors)
		if !ok {
			//nolint: errorlint
			return fmt.Errorf("unexpected error type: %T", err)
		}

		for i := range callErr {
			if callErr[i] != nil {
				txes[i].Status = db.Replaced
			} else {
				foundSuccessfulTX = true
				txes[i].Status = db.Confirmed
			}
		}
	} else if receipts[0].TxHash == txes[0].Hash() {
		// there must be only one tx, so we can just check the first one
		// TODO: handle the case where there is more than one
		txes[0].Status = db.Confirmed
		foundSuccessfulTX = true
	}

	if foundSuccessfulTX {
		err = t.db.PutTXS(ctx, txes...)
		if err != nil {
			return fmt.Errorf("could not put txes: %w", err)
		}
	}

	return nil
}
