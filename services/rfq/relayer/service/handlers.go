package service

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// handleBridgeRequestedLog handles the BridgeRequestedLog event.
// Step 1: Seen
//
// This is the first event emitted in the bridge process. It is emitted when a user calls bridge on chain.
// To process it, we decode the bridge transaction and store all the data, marking it as seen.
// This marks the event as seen.
func (r *Relayer) handleBridgeRequestedLog(parentCtx context.Context, req *fastbridge.FastBridgeBridgeRequested, chainID uint64) (err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "handleBridgeRequestedLog", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(req.TransactionId[:])),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// TODO: consider a mapmutex
	_, err = r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	// expect no results
	if !errors.Is(err, reldb.ErrNoQuoteForID) {
		// maybe a db err? if so error out & try again later
		if err != nil {
			return fmt.Errorf("could not call db: %w", err)
		}
	}

	// TODO: these should be premade
	originClient, err := r.client.GetChainClient(ctx, int(chainID))
	if err != nil {
		return fmt.Errorf("could not get correct omnirpc client: %w", err)
	}

	fastBridge, err := fastbridge.NewFastBridgeRef(req.Raw.Address, originClient)
	if err != nil {
		return fmt.Errorf("could not get correct fast bridge: %w", err)
	}

	bridgeTx, err := fastBridge.GetBridgeTransaction(&bind.CallOpts{Context: ctx}, req.Request)
	if err != nil {
		return fmt.Errorf("could not get bridge transaction: %w", err)
	}

	// TODO: you can just pull these out of inventory. If they don't exist mark as invalid.
	decimals, err := r.getDecimals(ctx, bridgeTx)
	// can't use errors.is here
	if err != nil && strings.Contains(err.Error(), "no contract code at given address") {
		logger.Warnf("invalid token, skipping")
		return nil
	}

	if err != nil {
		return fmt.Errorf("could not get decimals: %w", err)
	}

	err = r.db.StoreQuoteRequest(ctx, reldb.QuoteRequest{
		BlockNumber:         req.Raw.BlockNumber,
		RawRequest:          req.Request,
		OriginTokenDecimals: decimals.originDecimals,
		DestTokenDecimals:   decimals.destDecimals,
		TransactionID:       req.TransactionId,
		Sender:              req.Sender,
		Transaction:         bridgeTx,
		Status:              reldb.Seen,
		OriginTxHash:        req.Raw.TxHash,
	})
	if err != nil {
		return fmt.Errorf("could not get db: %w", err)
	}

	return nil
}

// handleSeen handles the seen status.
// Step 2: CommittedPending
// Possible Errors: WillNotProcess, NotEnoughInventory
//
// This is the second step in the bridge process. It is emitted when the relayer sees the request.
// We check if we have enough inventory to process the request and mark it as committed pending.
func (q *QuoteRequestHandler) handleSeen(ctx context.Context, span trace.Span, request reldb.QuoteRequest) (err error) {
	shouldProcess, err := q.Quoter.ShouldProcess(ctx, request)
	if err != nil {
		// will retry later
		return fmt.Errorf("could not determine if should process: %w", err)
	}
	if !shouldProcess {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.WillNotProcess)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
		// shouldn't process from here on out
		return nil
	}

	// check if the quote is profitable
	isProfitable, err := q.Quoter.IsProfitable(ctx, request)
	if err != nil {
		// will retry later
		return fmt.Errorf("could not determine if profitable: %w", err)
	}
	if !isProfitable {
		// will retry later since profitability is dependent on dynamic gas prices
		span.AddEvent("quote is not profitable")
		return nil
	}

	// get destination committable balancs
	committableBalance, err := q.Inventory.GetCommittableBalance(ctx, int(q.Dest.ChainID), request.Transaction.DestToken)
	if err != nil {
		return fmt.Errorf("could not get committable balance: %w", err)
	}
	// if committableBalance > destAmount
	if committableBalance.Cmp(request.Transaction.DestAmount) < 0 {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.NotEnoughInventory)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
		return nil
	}
	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.CommittedPending)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// handleCommitPending handles the committed pending status.
// Step 3: CommittedConfirmed
//
// This is the third step in the bridge process. Here we check the current chain block number against our confirmation count
// if we've passed the conf threshold, we mark the request as committed confirmed.
//
// TODO: we still need to handle the reorg state here, but for right now it just gets stuck in the queue and
// never relayed.
// Reorgs are rare enough that its questionable wether this is ever worth building or if we can just
// leave these in the queue.
func (q *QuoteRequestHandler) handleCommitPending(ctx context.Context, span trace.Span, request reldb.QuoteRequest) (err error) {
	earliestConfirmBlock := request.BlockNumber + q.Origin.Confirmations

	latestBlock := q.Origin.LatestBlock()
	shouldContinue := latestBlock >= earliestConfirmBlock
	span.AddEvent("pending_check", trace.WithAttributes(
		attribute.Int("latest_block", int(latestBlock)),
		attribute.Int("earliest_confirm_block", int(earliestConfirmBlock)),
		attribute.Bool("should_continue", shouldContinue),
	))

	if !shouldContinue {
		span.AddEvent("will_not_continue")
		// can't complete yet, do nothing
		return nil
	}

	bs, err := q.Origin.Bridge.BridgeStatuses(&bind.CallOpts{Context: ctx}, request.TransactionID)
	if err != nil {
		return fmt.Errorf("could not get bridge status: %w", err)
	}

	span.AddEvent("status_check", trace.WithAttributes(attribute.String("chain_bridge_status", fastbridge.BridgeStatus(bs).String())))

	// sanity check to make sure it's still requested.
	if bs == fastbridge.REQUESTED.Int() {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.CommittedConfirmed)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
	}
	return nil
}

// handleCommitConfirmed handles the committed confirmed status.
// Step 4: RelayStarted
//
// This is the fourth step in the bridge process. Here we submit the relay transaction to the destination chain.
// TODO: just to be safe, we should probably check if another relayer has already relayed this.
func (q *QuoteRequestHandler) handleCommitConfirmed(ctx context.Context, _ trace.Span, request reldb.QuoteRequest) (err error) {
	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.RelayStarted)
	if err != nil {
		return fmt.Errorf("could not update quote request status: %w", err)
	}

	// TODO: store the dest txhash connected to the nonce
	nonce, _, err := q.Dest.SubmitRelay(ctx, request)
	if err != nil {
		return fmt.Errorf("could not submit relay: %w", err)
	}
	_ = nonce

	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// handleRelayStarted handles the relay started status and marks the relay as completed.
// Step 5: RelayCompleted
//
// This is the fifth step in the bridge process. Here we check if the relay has been completed on the destination chain.
// Notably, this is polled from the chain listener rather than the database since we wait for the log to show up.
func (r *Relayer) handleRelayLog(ctx context.Context, req *fastbridge.FastBridgeBridgeRelayed) (err error) {
	reqID, err := r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	if err != nil {
		return fmt.Errorf("could not get quote request: %w", err)
	}
	// we might've accidentally gotten this later, if so we'll just ignore it
	if reqID.Status != reldb.RelayStarted {
		logger.Warnf("got relay log for request that was not relay started (transaction id: %s, txhash: %s)", hexutil.Encode(reqID.TransactionID[:]), req.Raw.TxHash)
		return nil
	}

	// TODO: this can still get re-orged
	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.RelayCompleted)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	err = r.db.UpdateDestTxHash(ctx, req.TransactionId, req.Raw.TxHash)
	if err != nil {
		return fmt.Errorf("could not update dest tx hash: %w", err)
	}
	return nil
}

// handleRelayCompleted handles the relay completed status and marks the claim as started.
// Step 6: ProvePosting
//
// This is the sixth step in the bridge process. Here we submit the claim transaction to the origin chain.
func (q *QuoteRequestHandler) handleRelayCompleted(ctx context.Context, _ trace.Span, request reldb.QuoteRequest) (err error) {
	// relays been completed, it's time to go back to the origin chain and try to prove
	_, err = q.Origin.SubmitTransaction(ctx, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = q.Origin.Bridge.Prove(transactor, request.RawRequest, request.DestTxHash)
		if err != nil {
			return nil, fmt.Errorf("could not relay: %w", err)
		}

		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}

	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.ProvePosting)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// handleProofProvided handles the ProofProvided event emitted by the Bridge.
// Step 7: ProvePosted
//
// This is the seventh step in the bridge process. Here we process the event that the proof was posted on chain.
func (r *Relayer) handleProofProvided(ctx context.Context, req *fastbridge.FastBridgeBridgeProofProvided) (err error) {
	// TODO: this can still get re-orged
	// ALso: we should make sure the previous status  is ProvePosting
	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.ProvePosted)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// handleProofPosted handles the proof posted status and marks the claim as pending.
// Step 8: ClaimPending
//
// we'll wait until optimistic period is over to check if we can claim.
func (q *QuoteRequestHandler) handleProofPosted(ctx context.Context, _ trace.Span, request reldb.QuoteRequest) (err error) {
	// we shouldnt' check the claim yet
	if !q.shouldCheckClaim(request) {
		return nil
	}

	// make sure relayer hasn't already proved. This is neeeded in case of an abrupt halt in event sourcing
	// note:  this assumes caller has already checked the sender is the relayer.
	bs, err := q.Origin.Bridge.BridgeStatuses(&bind.CallOpts{Context: ctx}, request.TransactionID)
	if err != nil {
		return fmt.Errorf("could not get bridge status: %w", err)
	}

	if bs == fastbridge.RelayerProved.Int() {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.ClaimPending)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
		return nil
	}

	if bs == fastbridge.RelayerClaimed.Int() {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.ClaimCompleted)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
		return nil
	}

	canClaim, err := q.Origin.Bridge.CanClaim(&bind.CallOpts{Context: ctx}, request.TransactionID, q.RelayerAddress)
	if err != nil {
		return fmt.Errorf("could not check if can claim: %w", err)
	}

	// can't claim yet. we'll check again later
	if !canClaim {
		return nil
	}
	_, err = q.Origin.SubmitTransaction(ctx, func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = q.Origin.Bridge.Claim(transactor, request.RawRequest, transactor.From)
		if err != nil {
			return nil, fmt.Errorf("could not relay: %w", err)
		}
		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not submit transaction: %w", err)
	}

	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.ClaimPending)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// Error Handlers Only from this point below.
//
// handleNotEnoughInventory handles the not enough inventory status.
func (q *QuoteRequestHandler) handleNotEnoughInventory(ctx context.Context, _ trace.Span, request reldb.QuoteRequest) (err error) {
	committableBalance, err := q.Inventory.GetCommittableBalance(ctx, int(q.Dest.ChainID), request.Transaction.DestToken)
	if err != nil {
		return fmt.Errorf("could not get committable balance: %w", err)
	}
	// if committableBalance > destAmount
	if committableBalance.Cmp(request.Transaction.DestAmount) > 0 {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.CommittedPending)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
	}
	return nil
}
