package service

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/retry"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var maxRPCRetryTime = 15 * time.Second

func (g *Guard) handleBridgeRequestedLog(parentCtx context.Context, req *fastbridge.FastBridgeBridgeRequested, chainID int) (err error) {
	ctx, span := g.metrics.Tracer().Start(parentCtx, "handleBridgeRequestedLog-guard", trace.WithAttributes(
		attribute.Int(metrics.Origin, chainID),
		attribute.String("transaction_id", hexutil.Encode(req.TransactionId[:])),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	originClient, err := g.client.GetChainClient(ctx, chainID)
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

func (g *Guard) handleProofProvidedLog(parentCtx context.Context, event *fastbridge.FastBridgeBridgeProofProvided, chainID int) (err error) {
	ctx, span := g.metrics.Tracer().Start(parentCtx, "handleProofProvidedLog-guard", trace.WithAttributes(
		attribute.Int(metrics.Origin, chainID),
		attribute.String("transaction_id", hexutil.Encode(event.TransactionId[:])),
		attribute.String("tx_hash", hexutil.Encode(event.TransactionHash[:])),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

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

func (g *Guard) handleProofDisputedLog(parentCtx context.Context, event *fastbridge.FastBridgeBridgeProofDisputed) (err error) {
	ctx, span := g.metrics.Tracer().Start(parentCtx, "handleProofDisputedLog-guard", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(event.TransactionId[:])),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	err = g.db.UpdatePendingProvenStatus(ctx, event.TransactionId, guarddb.Disputed)
	if err != nil {
		return fmt.Errorf("could not update pending proven status: %w", err)
	}

	return nil
}

func (g *Guard) handleProveCalled(parentCtx context.Context, proven *guarddb.PendingProven) (err error) {
	ctx, span := g.metrics.Tracer().Start(parentCtx, "handleProveCalled", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(proven.TransactionID[:])),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// first, get the corresponding bridge request
	bridgeRequest, err := g.db.GetBridgeRequestByID(ctx, proven.TransactionID)
	if err != nil {
		return fmt.Errorf("could not get bridge request for txid %s: %w", hexutil.Encode(proven.TransactionID[:]), err)
	}

	valid, err := g.isProveValid(ctx, proven, bridgeRequest)
	if err != nil {
		return fmt.Errorf("could not check prove validity: %w", err)
	}
	span.SetAttributes(attribute.Bool("valid", valid))

	//nolint:nestif
	if valid {
		// mark as validated
		err = g.db.UpdatePendingProvenStatus(ctx, proven.TransactionID, guarddb.Validated)
		if err != nil {
			return fmt.Errorf("could not update pending proven status: %w", err)
		}
	} else {
		// trigger dispute
		contract, ok := g.contracts[int(bridgeRequest.Transaction.OriginChainId)]
		if !ok {
			return fmt.Errorf("could not get contract for chain: %d", bridgeRequest.Transaction.OriginChainId)
		}
		_, err = g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(bridgeRequest.Transaction.OriginChainId)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
			tx, err = contract.Dispute(transactor, proven.TransactionID)
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
	if errors.Is(err, ethereum.NotFound) {
		// if tx hash does not exist, we want to consider the proof invalid
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("could not get receipt: %w", err)
	}
	rfqAddr, err := g.cfg.GetRFQAddress(int(bridgeRequest.Transaction.DestChainId))
	if err != nil {
		return false, fmt.Errorf("could not get rfq address: %w", err)
	}
	parser, err := fastbridge.NewParser(common.HexToAddress(rfqAddr))
	if err != nil {
		return false, fmt.Errorf("could not get parser: %w", err)
	}

	for _, log := range receipt.Logs {
		_, parsedEvent, ok := parser.ParseEvent(*log)
		if !ok {
			continue
		}

		if log.Address != common.HexToAddress(rfqAddr) {
			continue
		}

		event, ok := parsedEvent.(*fastbridge.FastBridgeBridgeRelayed)
		if ok {
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
