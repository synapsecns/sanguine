package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var (
	maxTotalTime    = 15 * time.Second
	maxRPCRetryTime = 30 * time.Second
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

	unlocker, ok := r.handlerMtx.TryLock(hexutil.Encode(req.TransactionId[:]))
	if !ok {
		span.SetAttributes(attribute.Bool("locked", true))
		// already processing this request
		return nil
	}

	defer unlocker.Unlock()

	_, err = r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	// expect no results
	if !errors.Is(err, reldb.ErrNoQuoteForID) {
		// maybe a db err? if so error out & try again later
		if err != nil {
			return fmt.Errorf("could not call db: %w", err)
		}

		span.AddEvent("already known")
		// already seen this request
		return nil
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

	// TODO: you can just pull these out of inventory. If they don't exist mark as invalid.
	originDecimals, destDecimals, err := r.getDecimalsFromBridgeTx(ctx, bridgeTx)
	// can't use errors.is here
	if err != nil && strings.Contains(err.Error(), "no contract code at given address") {
		logger.Warnf("invalid token, skipping")
		return nil
	}

	if err != nil || originDecimals == nil || destDecimals == nil {
		return fmt.Errorf("could not get decimals: %w", err)
	}

	dbReq := reldb.QuoteRequest{
		BlockNumber:         req.Raw.BlockNumber,
		RawRequest:          req.Request,
		OriginTokenDecimals: *originDecimals,
		DestTokenDecimals:   *destDecimals,
		TransactionID:       req.TransactionId,
		Sender:              req.Sender,
		Transaction:         bridgeTx,
		Status:              reldb.Seen,
		OriginTxHash:        req.Raw.TxHash,
	}
	err = r.db.StoreQuoteRequest(ctx, dbReq)
	if err != nil {
		return fmt.Errorf("could not get db: %w", err)
	}

	// immediately forward the request to handleSeen
	span.AddEvent("sending to handleSeen")
	qr, err := r.requestToHandler(ctx, dbReq)
	if err != nil {
		return fmt.Errorf("could not get quote request handler: %w", err)
	}
	// Forward instead of lock since we called lock above.
	fwdErr := qr.Forward(ctx, dbReq)
	if fwdErr != nil {
		logger.Errorf("could not forward to handle seen: %w", fwdErr)
		span.AddEvent("could not forward to handle seen")
	}

	return nil
}

// handleSeen handles the seen status.
// Step 2: CommittedPending
// Possible Errors: WillNotProcess, NotEnoughInventory
//
// This is the second step in the bridge process. It is emitted when the relayer sees the request.
// We check if we have enough inventory to process the request and mark it as committed pending.
//
//nolint:cyclop
func (q *QuoteRequestHandler) handleSeen(ctx context.Context, span trace.Span, request reldb.QuoteRequest) (err error) {
	shouldProcess, err := q.Quoter.ShouldProcess(ctx, request)
	if err != nil {
		// will retry later
		return fmt.Errorf("could not determine if should process: %w", err)
	}
	if !shouldProcess {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.WillNotProcess, &request.Status)
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

	// check balance and mark it as CommitPending or NotEnoughInventory.
	// note that this will update the request.Status in-place.
	err = q.commitPendingBalance(ctx, span, &request)
	if err != nil {
		return fmt.Errorf("could not commit pending balance: %w", err)
	}

	//nolint:exhaustive
	switch request.Status {
	case reldb.CommittedPending:
		// immediately forward the request to handleCommitPending
		span.AddEvent("forwarding to handleCommitPending")
		fwdErr := q.Forward(ctx, request)
		if fwdErr != nil {
			logger.Errorf("could not forward to handle commit pending: %w", fwdErr)
			span.AddEvent(fmt.Sprintf("could not forward to handle commit pending: %s", fwdErr.Error()))
		}
	case reldb.NotEnoughInventory:
		span.AddEvent("not enough inventory; not forwarding")
	default:
		logger.Errorf("unexpected request status: %s", request.Status.String())
		span.AddEvent(fmt.Sprintf("unexpected request status: %s", request.Status.String()))
	}

	return nil
}

// commitPendingBalance locks the balance and marks the request as CommitPending.
func (q *QuoteRequestHandler) commitPendingBalance(ctx context.Context, span trace.Span, request *reldb.QuoteRequest) (err error) {
	// lock the consumed balance
	key := getBalanceMtxKey(q.Dest.ChainID, request.Transaction.DestToken)
	span.SetAttributes(attribute.String("balance_lock_key", key))
	unlocker, ok := q.balanceMtx.TryLock(key)
	if !ok {
		// balance is locked due to concurrent request, try again later
		span.SetAttributes(attribute.Bool("locked", true))
		return nil
	}
	defer unlocker.Unlock()

	// get destination committable balance
	committableBalance, err := q.Inventory.GetCommittableBalance(ctx, int(q.Dest.ChainID), request.Transaction.DestToken)
	if errors.Is(err, inventory.ErrUnsupportedChain) {
		// don't process request if chain is currently unsupported
		span.AddEvent("dropping unsupported chain")
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.WillNotProcess, &request.Status)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("could not get committable balance: %w", err)
	}

	// check if we have enough inventory to handle the request
	if committableBalance.Cmp(request.Transaction.DestAmount) < 0 {
		request.Status = reldb.NotEnoughInventory
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.NotEnoughInventory, &request.Status)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
		return nil
	}

	// get ack from API to synchronize calls with other relayers and avoid reverts
	req := model.PutAckRequest{
		TxID:        hexutil.Encode(request.TransactionID[:]),
		DestChainID: int(request.Transaction.DestChainId),
	}
	resp, err := q.apiClient.PutRelayAck(ctx, &req)
	if err != nil {
		return fmt.Errorf("could not get relay ack: %w", err)
	}
	span.SetAttributes(
		attribute.String("transaction_id", hexutil.Encode(request.TransactionID[:])),
		attribute.Bool("should_relay", resp.ShouldRelay),
		attribute.String("relayer_address", resp.RelayerAddress),
	)
	if !resp.ShouldRelay {
		span.AddEvent("not relaying due to ack")
		return nil
	}

	request.Status = reldb.CommittedPending
	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.CommittedPending, &request.Status)
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

	var bs uint8
	call := func(ctx context.Context) error {
		bs, err = q.Origin.Bridge.BridgeStatuses(&bind.CallOpts{Context: ctx}, request.TransactionID)
		if err != nil {
			return fmt.Errorf("could not get bridge status: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, call, retry.WithMaxTotalTime(maxRPCRetryTime))
	if err != nil {
		return fmt.Errorf("could not make contract call: %w", err)
	}

	span.AddEvent("status_check", trace.WithAttributes(attribute.String("chain_bridge_status", fastbridge.BridgeStatus(bs).String())))

	// sanity check to make sure it's still requested.
	if bs != fastbridge.REQUESTED.Int() {
		return nil
	}

	request.Status = reldb.CommittedConfirmed
	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.CommittedConfirmed, &request.Status)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}

	// immediately forward to handleCommitConfirmed
	span.AddEvent("forwarding to handleCommitConfirmed")
	fwdErr := q.Forward(ctx, request)
	if fwdErr != nil {
		logger.Errorf("could not forward to handle commit confirmed: %w", fwdErr)
		span.AddEvent("could not forward to handle commit confirmed")
	}

	return nil
}

// handleCommitConfirmed handles the committed confirmed status.
// Step 4: RelayStarted
//
// This is the fourth step in the bridge process. Here we submit the relay transaction to the destination chain.
// TODO: just to be safe, we should probably check if another relayer has already relayed this.
func (q *QuoteRequestHandler) handleCommitConfirmed(ctx context.Context, span trace.Span, request reldb.QuoteRequest) (err error) {
	// TODO: store the dest txhash connected to the nonce
	nonce, _, err := q.Dest.SubmitRelay(ctx, request)
	if err != nil {
		return fmt.Errorf("could not submit relay: %w", err)
	}
	span.AddEvent("relay successfully submitted")
	span.SetAttributes(attribute.Int("relay_nonce", int(nonce)))

	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.RelayStarted, &request.Status)
	if err != nil {
		return fmt.Errorf("could not update quote request status: %w", err)
	}

	err = q.db.UpdateRelayNonce(ctx, request.TransactionID, nonce)
	if err != nil {
		return fmt.Errorf("could not update relay nonce: %w", err)
	}

	return nil
}

// handleRelayStarted handles the relay started status and marks the relay as completed.
// Step 5: RelayCompleted
//
// This is the fifth step in the bridge process. Here we check if the relay has been completed on the destination chain.
// Notably, this is polled from the chain listener rather than the database since we wait for the log to show up.
func (r *Relayer) handleRelayLog(parentCtx context.Context, req *fastbridge.FastBridgeBridgeRelayed) (err error) {
	ctx, span := r.metrics.Tracer().Start(parentCtx, "handleRelayLog",
		trace.WithAttributes(attribute.String("transaction_id", hexutil.Encode(req.TransactionId[:]))),
	)
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	reqID, err := r.db.GetQuoteRequestByID(ctx, req.TransactionId)
	if err != nil {
		return fmt.Errorf("could not get quote request: %w", err)
	}
	// we might've accidentally gotten this later, if so we'll just ignore it
	// note that in the edge case where we pessimistically marked as DeadlineExceeded
	// and the relay was actually successful, we should continue the proving process
	if reqID.Status != reldb.RelayStarted && reqID.Status != reldb.DeadlineExceeded {
		logger.Warnf("got relay log for request that was not relay started (transaction id: %s, txhash: %s)", hexutil.Encode(reqID.TransactionID[:]), req.Raw.TxHash)
		return nil
	}

	// TODO: this can still get re-orged
	err = r.db.UpdateDestTxHash(ctx, req.TransactionId, req.Raw.TxHash)
	if err != nil {
		return fmt.Errorf("could not update dest tx hash: %w", err)
	}
	span.SetAttributes(attribute.String("dest_tx_hash", hexutil.Encode(req.Raw.TxHash[:])))

	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.RelayCompleted, nil)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// handleRelayCompleted handles the relay completed status and marks the claim as started.
// Step 6: ProvePosting
//
// This is the sixth step in the bridge process. Here we submit the claim transaction to the origin chain.
func (q *QuoteRequestHandler) handleRelayCompleted(ctx context.Context, span trace.Span, request reldb.QuoteRequest) (err error) {
	// fetch the block of the relay transaction to confirm that it has been finalized
	relayBlockNumber, err := q.getRelayBlockNumber(ctx, request)
	if err != nil {
		// relay tx must have gotten reorged; mark as RelayRaceLost
		span.SetAttributes(attribute.String("receipt_error", err.Error()))
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.RelayRaceLost, nil)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
		return fmt.Errorf("could not get relay block number: %w", err)
	}

	currentBlockNumber := q.Dest.LatestBlock()
	proveConfirmations, err := q.cfg.GetFinalityConfirmations(int(q.Dest.ChainID))
	if err != nil {
		return fmt.Errorf("could not get prove confirmations: %w", err)
	}

	//nolint:gosec
	span.SetAttributes(
		attribute.Int("current_block_number", int(currentBlockNumber)),
		attribute.Int("relay_block_number", int(relayBlockNumber)),
		attribute.Int("prove_confirmations", int(proveConfirmations)),
	)
	if currentBlockNumber < relayBlockNumber+proveConfirmations {
		span.AddEvent("not enough confirmations")
		return nil
	}

	// relay has been finalized, it's time to go back to the origin chain and try to prove
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

	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.ProvePosting, &request.Status)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// getRelayBlockNumber fetches the block number of the relay transaction for a given quote request.
func (q *QuoteRequestHandler) getRelayBlockNumber(ctx context.Context, request reldb.QuoteRequest) (blockNumber uint64, err error) {
	// fetch the transaction receipt for corresponding tx hash
	var receipt *types.Receipt
	err = retry.WithBackoff(ctx, func(context.Context) error {
		receipt, err = q.Dest.Client.TransactionReceipt(ctx, request.DestTxHash)
		if err != nil {
			return fmt.Errorf("could not get transaction receipt: %w", err)
		}
		return nil
	}, retry.WithMaxTotalTime(maxTotalTime))
	if err != nil {
		return blockNumber, fmt.Errorf("could not get receipt: %w", err)
	}
	parser, err := fastbridge.NewParser(q.Dest.Bridge.Address())
	if err != nil {
		return blockNumber, fmt.Errorf("could not create parser: %w", err)
	}

	// check that a Relayed event was emitted
	for _, log := range receipt.Logs {
		if log == nil {
			continue
		}
		_, parsedEvent, ok := parser.ParseEvent(*log)
		if !ok {
			continue
		}
		_, ok = parsedEvent.(*fastbridge.FastBridgeBridgeRelayed)
		if ok {
			return receipt.BlockNumber.Uint64(), nil
		}
	}

	return blockNumber, fmt.Errorf("relayed event not found for dest tx hash: %s", request.DestTxHash.Hex())
}

// handleProofProvided handles the ProofProvided event emitted by the Bridge.
// Step 7: ProvePosted
//
// This is the seventh step in the bridge process. Here we process the event that the proof was posted on chain.
func (r *Relayer) handleProofProvided(ctx context.Context, req *fastbridge.FastBridgeBridgeProofProvided) (err error) {
	if req.Relayer != r.signer.Address() {
		return nil
	}

	// TODO: this can still get re-orged
	// ALso: we should make sure the previous status  is ProvePosting
	err = r.db.UpdateQuoteRequestStatus(ctx, req.TransactionId, reldb.ProvePosted, nil)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// handleProofPosted handles the proof posted status and marks the claim as pending.
// Step 8: ClaimPending
//
// we'll wait until optimistic period is over to check if we can claim.
//
//nolint:cyclop
func (q *QuoteRequestHandler) handleProofPosted(ctx context.Context, span trace.Span, request reldb.QuoteRequest) (err error) {
	// we shouldnt' check the claim yet
	if !q.shouldCheckClaim(request) {
		return nil
	}

	// make sure relayer hasn't already proved. This is neeeded in case of an abrupt halt in event sourcing
	// note:  this assumes caller has already checked the sender is the relayer.
	var bs uint8
	call := func(ctx context.Context) error {
		bs, err = q.Origin.Bridge.BridgeStatuses(&bind.CallOpts{Context: ctx}, request.TransactionID)
		if err != nil {
			return fmt.Errorf("could not get bridge status: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, call, retry.WithMaxTotalTime(maxRPCRetryTime))
	if err != nil {
		return fmt.Errorf("could not make contract call: %w", err)
	}
	switch bs {
	case fastbridge.RelayerProved.Int():
		// no op
	case fastbridge.RelayerClaimed.Int():
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.ClaimCompleted, &request.Status)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
	default:
		if span != nil {
			span.SetAttributes(attribute.Int("claim_bridge_status", int(bs)))
			span.AddEvent("unexpected bridge status for claim")
		}
		return nil
	}

	proofs, err := q.Origin.Bridge.BridgeProofs(&bind.CallOpts{Context: ctx}, request.TransactionID)
	if err != nil {
		return fmt.Errorf("could not get bridge proofs: %w", err)
	}
	if proofs.Relayer != q.RelayerAddress {
		if span != nil {
			span.SetAttributes(attribute.String("proof_relayer", proofs.Relayer.String()))
			span.AddEvent("unexpected relayer in proof")
		}
		return fmt.Errorf("onchain proof does not match our relayer")
	}

	var canClaim bool
	claimCall := func(ctx context.Context) error {
		canClaim, err = q.Origin.Bridge.CanClaim(&bind.CallOpts{Context: ctx}, request.TransactionID, q.RelayerAddress)
		if err != nil {
			return fmt.Errorf("could not check if can claim: %w", err)
		}
		return nil
	}
	err = retry.WithBackoff(ctx, claimCall, retry.WithMaxTotalTime(maxRPCRetryTime))
	if err != nil {
		return fmt.Errorf("could not make call: %w", err)
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

	err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.ClaimPending, &request.Status)
	if err != nil {
		return fmt.Errorf("could not update request status: %w", err)
	}
	return nil
}

// Error Handlers Only from this point below.
//
// handleNotEnoughInventory handles the not enough inventory status.
func (q *QuoteRequestHandler) handleNotEnoughInventory(ctx context.Context, span trace.Span, request reldb.QuoteRequest) (err error) {
	// acquire balance lock
	key := getBalanceMtxKey(q.Dest.ChainID, request.Transaction.DestToken)
	span.SetAttributes(attribute.String("balance_lock_key", key))
	unlocker, ok := q.balanceMtx.TryLock(key)
	if !ok {
		// balance is locked due to concurrent request, try again later
		span.SetAttributes(attribute.Bool("locked", true))
		return nil
	}
	defer unlocker.Unlock()

	// commit destination balance
	committableBalance, err := q.Inventory.GetCommittableBalance(ctx, int(q.Dest.ChainID), request.Transaction.DestToken)
	if err != nil {
		return fmt.Errorf("could not get committable balance: %w", err)
	}
	if committableBalance.Cmp(request.Transaction.DestAmount) > 0 {
		err = q.db.UpdateQuoteRequestStatus(ctx, request.TransactionID, reldb.CommittedPending, &request.Status)
		if err != nil {
			return fmt.Errorf("could not update request status: %w", err)
		}
	}
	return nil
}
