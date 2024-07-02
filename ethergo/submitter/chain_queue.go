package submitter

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/params"

	"github.com/synapsecns/sanguine/ethergo/util"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/lmittmann/w3/module/eth"
	"github.com/lmittmann/w3/w3types"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/ethergo/submitter/db"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// chainQueue is a single use queue for a single chain.
type chainQueue struct {
	*txSubmitterImpl
	// client is the client for this chain
	client client.EVM
	// client is the nonce used for this chain
	nonce uint64
	// txsHaveConfirmed is true if any of the txes have confirmed
	txsHaveConfirmed bool
	// chainID is the chainID for this queue
	chainID *big.Int
	// reprocessQueue is a list of transactions that should be resubmitted
	reprocessQueue []db.TX
	// reprocessQueueMux is a mutex for reprocessQueue
	reprocessQueueMux sync.Mutex
}

func (c *chainQueue) chainIDInt() int {
	return int(c.chainID.Int64())
}

func (t *txSubmitterImpl) chainPendingQueue(parentCtx context.Context, chainID *big.Int, txes []db.TX) (err error) {
	ctx, span := t.metrics.Tracer().Start(parentCtx, "submitter.ChainQueue", trace.WithAttributes(
		attribute.String("chain_id", chainID.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// chainClient is the client for the chain we're working on
	chainClient, err := t.fetcher.GetClient(ctx, chainID)
	if err != nil {
		return fmt.Errorf("could not get client: %w", err)
	}

	currentNonce, err := chainClient.NonceAt(ctx, t.signer.Address(), nil)
	if err != nil {
		return fmt.Errorf("could not get nonce: %w", err)
	}
	span.SetAttributes(attribute.Int("nonce", int(currentNonce)))

	// record metrics for txes.
	t.otelRecorder.RecordNonceForChain(uint32(chainID.Int64()), currentNonce)
	t.otelRecorder.RecordNumPendingTxes(uint32(chainID.Int64()), calculatePendingTxes(txes, currentNonce))
	t.otelRecorder.RecordOldestPendingTx(uint32(chainID.Int64()), time.Since(fetchOldestPendingTx(txes, currentNonce)))

	wg := &sync.WaitGroup{}

	cq := chainQueue{
		txSubmitterImpl: t,
		chainID:         core.CopyBigInt(chainID),
		nonce:           currentNonce,
		client:          chainClient,
	}

	// so now, we have a list of transactions that need to be resubmitted.
	// these are already ordered by nonce and we should have at most one per nonce, so we can just iterate through them.
	// we need to figure out which ones are still valid, and which ones need to be bumped.
	// once this is done, we'll be ready to submit them.
	// we're going to handle this by updating txes in place.
	// note that if we do not have sufficient gas on this chain, no txes will be bumped.
	gasBalance, err := chainClient.BalanceAt(ctx, t.signer.Address(), nil)
	if err != nil {
		return fmt.Errorf("could not get gas balance: %w", err)
	}
	t.otelRecorder.RecordGasBalanceForChain(uint32(chainID.Int64()), core.CopyBigInt(gasBalance))
	span.SetAttributes(attribute.String("gas_balance", gasBalance.String()))

	for i := range txes {
		tx := txes[i]

		if tx.Nonce() < currentNonce {
			cq.txsHaveConfirmed = true
			continue
		}

		if tx.Cost().Cmp(gasBalance) > 0 {
			span.SetAttributes(attribute.Bool("out_of_gas", true))
			span.AddEvent("tx out of gas", trace.WithAttributes(txToAttributes(tx.Transaction, tx.UUID)...))
			break
		}

		// bump tx in new goroutine
		wg.Add(1)
		go func() {
			defer wg.Done()
			bumpErr := cq.bumpTX(ctx, tx)
			if bumpErr != nil {
				logger.Errorf("could not bump tx: %v", bumpErr)
			}
		}()
	}

	// update old tx statuses in parallel
	wg.Add(1)
	go func() {
		defer wg.Done()
		updateErr := cq.updateOldTxStatuses(ctx)
		if updateErr != nil {
			logger.Errorf("could not update old tx statuses: %v", updateErr)
		}
	}()
	wg.Wait()

	sort.Slice(cq.reprocessQueue, func(i, j int) bool {
		return cq.reprocessQueue[i].Nonce() < cq.reprocessQueue[j].Nonce()
	})

	calls := make([]w3types.Caller, len(cq.reprocessQueue))
	txHashes := make([]common.Hash, len(cq.reprocessQueue))
	for i, tx := range cq.reprocessQueue {
		calls[i] = eth.SendTx(tx.Transaction).Returns(&txHashes[i])
	}

	cq.storeAndSubmit(ctx, calls, span)

	return nil
}

func toFloat(wei *big.Int) float64 {
	// Convert wei to float64
	weiFloat := new(big.Float).SetInt(wei)
	weiAsFloat64, _ := weiFloat.Float64()

	// Perform the division to convert wei to ether
	return weiAsFloat64 / params.Ether
}

// storeAndSubmit stores the txes in the database and submits them to the chain.
func (c *chainQueue) storeAndSubmit(ctx context.Context, calls []w3types.Caller, span trace.Span) {
	var wg sync.WaitGroup
	wg.Add(2)

	storeCtx, cancelStore := context.WithCancel(ctx)

	go func() {
		defer wg.Done()
		err := c.db.PutTXS(storeCtx, c.reprocessQueue...)
		if err != nil {
			span.AddEvent("could not store txes", trace.WithAttributes(attribute.String("error", err.Error())))
		}
	}()

	go func() {
		defer wg.Done()
		err := c.client.BatchWithContext(ctx, calls...)
		cancelStore()
		for i := range c.reprocessQueue {
			if err != nil {
				c.reprocessQueue[i].Status = db.FailedSubmit
			} else {
				c.reprocessQueue[i].Status = db.Submitted
			}
		}

		err = c.db.PutTXS(ctx, c.reprocessQueue...)
		if err != nil {
			span.AddEvent("could not store txes", trace.WithAttributes(attribute.String("error", err.Error())))
		}
	}()
	wg.Wait()
}

// nolint: cyclop
func (c *chainQueue) bumpTX(parentCtx context.Context, ogTx db.TX) (err error) {
	if !c.isBumpIntervalElapsed(ogTx) {
		c.addToReprocessQueue(ogTx)
		return nil
	}
	// copy the transaction, switching the type if we need to.
	// this is required if the config changes to use legacy transactions on a tx that is already bumped.
	tx, err := util.CopyTX(ogTx.Transaction, util.WithTxType(c.txTypeForChain(c.chainID)))
	if err != nil {
		return fmt.Errorf("could not copy tx: %w", err)
	}

	ctx, span := c.metrics.Tracer().Start(parentCtx, "chainPendingQueue.bumpTX", trace.WithAttributes(attribute.Stringer(metrics.TxHash, tx.Hash())))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	newGasEstimate, err := c.getGasEstimate(ctx, c.client, c.chainIDInt(), tx)
	if err != nil {
		return fmt.Errorf("could not get gas estimate: %w", err)
	}

	transactor, err := c.signer.GetTransactor(ctx, c.chainID)
	if err != nil {
		return fmt.Errorf("could not get transactor: %w", err)
	}

	transactor.NoSend = true
	transactor.Nonce = new(big.Int).SetUint64(tx.Nonce())
	transactor.GasLimit = newGasEstimate

	err = c.setGasPrice(ctx, c.client, transactor, c.chainID, ogTx.Transaction)
	if err != nil {
		return fmt.Errorf("could not set gas price: %w", err)
	}

	switch tx.Type() {
	case types.LegacyTxType:
		tx = types.NewTx(&types.LegacyTx{
			Nonce:    tx.Nonce(),
			GasPrice: transactor.GasPrice,
			Gas:      transactor.GasLimit,
			To:       tx.To(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		})
	case types.DynamicFeeTxType:
		tx = types.NewTx(&types.DynamicFeeTx{
			ChainID:   tx.ChainId(),
			Nonce:     tx.Nonce(),
			GasTipCap: core.CopyBigInt(transactor.GasTipCap),
			GasFeeCap: core.CopyBigInt(transactor.GasFeeCap),
			Gas:       transactor.GasLimit,
			To:        tx.To(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		})
	default:
		return fmt.Errorf("unknown tx type: %v", ogTx.Type())
	}

	tx, err = transactor.Signer(transactor.From, tx)
	if err != nil {
		return fmt.Errorf("could not sign tx: %w", err)
	}

	span.AddEvent("add to reprocess queue")
	span.SetAttributes(txToAttributes(tx, ogTx.UUID)...)

	c.addToReprocessQueue(db.TX{
		UUID:        ogTx.UUID,
		Transaction: tx,
		Status:      db.Stored,
	})

	return nil
}

// addToReprocessQueue adds a tx to the reprocess queue.
func (c *chainQueue) addToReprocessQueue(tx db.TX) {
	c.reprocessQueueMux.Lock()
	defer c.reprocessQueueMux.Unlock()

	c.reprocessQueue = append(c.reprocessQueue, tx)
}

// TODO: test this mehtod.
func (c *chainQueue) isBumpIntervalElapsed(tx db.TX) bool {
	bumpInterval := c.config.GetBumpInterval(c.chainIDInt())
	elapsedSeconds := time.Since(tx.CreationTime().Add(bumpInterval)).Seconds()

	return elapsedSeconds >= 0
}

// updateOldTxStatuses updates the status of txes that are before the current nonce
// this will only run if we have txes that have confirmed.
func (c *chainQueue) updateOldTxStatuses(parentCtx context.Context) (err error) {
	// nothing to do
	if !c.txsHaveConfirmed {
		return
	}

	ctx, span := c.metrics.Tracer().Start(parentCtx, "chainPendingQueue.updateOldTxStatuses")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	err = c.db.MarkAllBeforeNonceReplacedOrConfirmed(ctx, c.signer.Address(), c.chainID, c.nonce)
	if err != nil {
		return fmt.Errorf("could not mark txes: %w", err)
	}
	return nil
}
