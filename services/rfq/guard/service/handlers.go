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
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

var maxRPCRetryTime = 15 * time.Second

func (g *Guard) handleBridgeRequestedLog(parentCtx context.Context, req *fastbridgev2.FastBridgeV2BridgeRequested, chainID int) (err error) {
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

//nolint:gosec
func (g *Guard) handleProofProvidedLog(parentCtx context.Context, event *fastbridgev2.FastBridgeV2BridgeProofProvided, chainID int) (err error) {
	ctx, span := g.metrics.Tracer().Start(parentCtx, "handleProofProvidedLog-guard", trace.WithAttributes(
		attribute.Int(metrics.Origin, chainID),
		attribute.String("transaction_id", hexutil.Encode(event.TransactionId[:])),
		attribute.String("tx_hash", hexutil.Encode(event.TransactionHash[:])),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	proven := guarddb.PendingProven{
		Origin:            uint32(chainID),
		RelayerAddress:    event.Relayer,
		FastBridgeAddress: event.Raw.Address,
		TransactionID:     event.TransactionId,
		TxHash:            event.TransactionHash,
		Status:            guarddb.ProveCalled,
		BlockNumber:       event.Raw.BlockNumber,
	}
	err = g.db.StorePendingProven(ctx, proven)
	if err != nil {
		return fmt.Errorf("could not store pending proven: %w", err)
	}

	return nil
}

func (g *Guard) handleProofDisputedLog(parentCtx context.Context, event *fastbridgev2.FastBridgeV2BridgeProofDisputed) (err error) {
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

	// first, verify that the prove tx is finalized
	finalized, err := g.isFinalized(ctx, int(proven.Origin), proven.BlockNumber)
	if err != nil {
		return fmt.Errorf("could not check if tx is finalized: %w", err)
	}
	span.SetAttributes(attribute.Bool("finalized", finalized))
	if !finalized {
		return nil
	}

	// get the corresponding bridge request
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
		if g.isV2Address(int(bridgeRequest.Transaction.OriginChainId), proven.FastBridgeAddress) {
			err = g.disputeV2(ctx, proven, bridgeRequest)
		} else {
			err = g.disputeV1(ctx, proven, bridgeRequest)
		}
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

func (g *Guard) disputeV1(ctx context.Context, proven *guarddb.PendingProven, bridgeRequest *guarddb.BridgeRequest) error {
	contract, ok := g.contractsV1[int(bridgeRequest.Transaction.OriginChainId)]
	if !ok {
		return errors.New("could not get contract")
	}
	_, err := g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(bridgeRequest.Transaction.OriginChainId)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = contract.Dispute(transactor, proven.TransactionID)
		if err != nil {
			return nil, fmt.Errorf("could not dispute: %w", err)
		}

		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not dispute: %w", err)
	}

	return nil
}

func (g *Guard) disputeV2(ctx context.Context, proven *guarddb.PendingProven, bridgeRequest *guarddb.BridgeRequest) error {
	contract, ok := g.contractsV2[int(bridgeRequest.Transaction.OriginChainId)]
	if !ok {
		return errors.New("could not get contract")
	}
	_, err := g.txSubmitter.SubmitTransaction(ctx, big.NewInt(int64(bridgeRequest.Transaction.OriginChainId)), func(transactor *bind.TransactOpts) (tx *types.Transaction, err error) {
		tx, err = contract.Dispute(transactor, proven.TransactionID)
		if err != nil {
			return nil, fmt.Errorf("could not dispute: %w", err)
		}

		return tx, nil
	})
	if err != nil {
		return fmt.Errorf("could not dispute: %w", err)
	}

	return nil
}

//nolint:cyclop
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

	var rfqContractAddr string

	if g.isV2Address(int(bridgeRequest.Transaction.OriginChainId), proven.FastBridgeAddress) {
		rfqContractAddr, err = g.cfg.GetRFQAddressV2(int(bridgeRequest.Transaction.DestChainId))
		if err != nil {
			return false, fmt.Errorf("could not get rfq address v2: %w", err)
		}
	} else {
		v1addr, err := g.cfg.GetRFQAddressV1(int(bridgeRequest.Transaction.DestChainId))
		if err != nil {
			return false, fmt.Errorf("could not get rfq address v1: %w", err)
		}
		if v1addr == nil {
			return false, fmt.Errorf("rfq address v1 is nil")
		}
		rfqContractAddr = *v1addr
	}

	var valid bool
	valid, err = g.isProveValidParse(ctx, proven, bridgeRequest, receipt, rfqContractAddr)

	if err != nil {
		return false, fmt.Errorf("could not parse proof for validity: %w", err)
	}

	return valid, nil
}

func (g *Guard) isProveValidParse(ctx context.Context, proven *guarddb.PendingProven, bridgeRequest *guarddb.BridgeRequest, receipt *types.Receipt, rfqContractAddr string) (bool, error) {
	span := trace.SpanFromContext(ctx)

	parser, err := fastbridgev2.NewParser(common.HexToAddress(rfqContractAddr))
	if err != nil {
		return false, fmt.Errorf("could not get parser: %w", err)
	}

	for _, log := range receipt.Logs {
		_, parsedEvent, ok := parser.ParseEvent(*log)
		if !ok {
			continue
		}

		if log.Address != common.HexToAddress(rfqContractAddr) {
			span.AddEvent(fmt.Sprintf("log address %s does not match rfq address %s", log.Address.Hex(), rfqContractAddr))
			continue
		}

		event, ok := parsedEvent.(*fastbridgev2.FastBridgeV2BridgeRelayed)
		if !ok {
			span.AddEvent("event is not a BridgeRelayed event")
			continue
		}

		if event.Relayer != proven.RelayerAddress {
			span.AddEvent(fmt.Sprintf("relayer address %s does not match prover address %s", event.Relayer.Hex(), proven.RelayerAddress.Hex()))
			continue
		}

		details := relayDetails{
			TransactionID: event.TransactionId,
			OriginAmount:  event.OriginAmount,
			DestAmount:    event.DestAmount,
			OriginChainID: event.OriginChainId,
			To:            event.To,
			OriginToken:   event.OriginToken,
			DestToken:     event.DestToken,
		}

		return relayMatchesBridgeRequest(details, bridgeRequest), nil
	}

	return false, nil
}

type relayDetails struct {
	TransactionID [32]byte
	OriginAmount  *big.Int
	DestAmount    *big.Int
	OriginChainID uint32
	To            common.Address
	OriginToken   common.Address
	DestToken     common.Address
}

func relayMatchesBridgeRequest(details relayDetails, bridgeRequest *guarddb.BridgeRequest) bool {
	// TODO: is this exhaustive?
	if details.TransactionID != bridgeRequest.TransactionID {
		return false
	}
	if details.OriginAmount.Cmp(bridgeRequest.Transaction.OriginAmount) != 0 {
		return false
	}
	if details.DestAmount.Cmp(bridgeRequest.Transaction.DestAmount) != 0 {
		return false
	}
	if details.OriginChainID != bridgeRequest.Transaction.OriginChainId {
		return false
	}
	if details.To != bridgeRequest.Transaction.DestRecipient {
		return false
	}
	if details.OriginToken != bridgeRequest.Transaction.OriginToken {
		return false
	}
	if details.DestToken != bridgeRequest.Transaction.DestToken {
		return false
	}
	return true
}

// isFinalized checks if a transaction is finalized versus the configured confirmations threshold.
func (g *Guard) isFinalized(ctx context.Context, chainID int, txBlockNumber uint64) (bool, error) {
	span := trace.SpanFromContext(ctx)

	client, err := g.client.GetChainClient(ctx, chainID)
	if err != nil {
		return false, fmt.Errorf("could not get chain client: %w", err)
	}

	currentBlockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("could not get block number: %w", err)
	}

	chainCfg, ok := g.cfg.Chains[chainID]
	if !ok {
		return false, fmt.Errorf("could not get chain config for chain %d", chainID)
	}
	threshBlockNumber := txBlockNumber + chainCfg.Confirmations

	//nolint:gosec
	span.SetAttributes(
		attribute.Int("chain_id", chainID),
		attribute.Int("current_block_number", int(currentBlockNumber)),
		attribute.Int("tx_block_number", int(txBlockNumber)),
		attribute.Int("confirmations", int(chainCfg.Confirmations)),
	)

	return currentBlockNumber >= threshBlockNumber, nil
}
