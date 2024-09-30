// Package limiter provides a rate limiting mechanism for RFQs to protect against reorgs.
package limiter

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridge"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// LatestBlockFetcher is the interface for fetching the latest block number.
type LatestBlockFetcher interface {
	// LatestBlock returns the latest block number.
	LatestBlock() uint64
}

// Limiter is the interface for rate limiting RFQs.
type Limiter interface {
	// IsAllowed returns true if the request is allowed, false otherwise.
	IsAllowed(ctx context.Context, request *reldb.QuoteRequest) (bool, error)
}

type limiterImpl struct {
	listener   LatestBlockFetcher
	metrics    metrics.Handler
	quoter     quoter.Quoter
	cfg        relconfig.Config
	tokenNames map[string]relconfig.TokenConfig
	evmClient  client.EVM
}

// NewRateLimiter creates a new Limiter.
func NewRateLimiter(
	cfg relconfig.Config,
	l LatestBlockFetcher,
	q quoter.Quoter,
	metricHandler metrics.Handler,
	tokens map[string]relconfig.TokenConfig,
	evmClient client.EVM,
) Limiter {
	return &limiterImpl{
		listener:   l,
		metrics:    metricHandler,
		quoter:     q,
		cfg:        cfg,
		tokenNames: tokens,
		evmClient:  evmClient,
	}
}

// IsAllowed returns true if the request is allowed, false otherwise.
func (l *limiterImpl) IsAllowed(ctx context.Context, request *reldb.QuoteRequest) (_ bool, err error) {
	ctx, span := l.metrics.Tracer().Start(
		ctx, "limiter.IsAllowed", trace.WithAttributes(util.QuoteRequestToAttributes(request)...),
	)

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()
	// if not enough confirmations, check volume. if under limit, wait 0 confirmations.
	underVolumeLimit, err := l.isUnderVolumeLimit(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not check volume limit: %w", err)
	}
	if underVolumeLimit {
		return true, nil
	}

	// if enough confirmations, allow because reorgs are rare at this point.
	hasEnoughConfirmations, err := l.hasEnoughConfirmations(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not check confirmations: %w", err)
	}
	if hasEnoughConfirmations {
		// we need to check if the receipt exists, parse the events from it, check for possible reverts,
		// and has the correct fields in case of a reorg. then, and only then, we can be sure a reorg will not
		// revert this.
		receiptFieldsMatch, err := l.checkReceipt(ctx, request)
		if err != nil {
			return false, fmt.Errorf("could not check receipt: %w", err)
		}
		return receiptFieldsMatch, nil
	}

	span.SetAttributes(
		attribute.Bool("has_enough_confirmations", hasEnoughConfirmations),
		attribute.Bool("within_size_limit", underVolumeLimit),
	)

	return false, nil
}

// hasEnoughConfirmations returns true if the request has enough confirmations, false otherwise.
func (l *limiterImpl) hasEnoughConfirmations(ctx context.Context, request *reldb.QuoteRequest) (_ bool, err error) {
	_, span := l.metrics.Tracer().Start(ctx, "limiter.hasEnoughConfirmations")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	currentBlockNumber := l.listener.LatestBlock()
	if err != nil {
		return false, fmt.Errorf("could not get block number: %w", err)
	}

	requiredConfirmations, err := l.getNumberOfConfirmationsToWait(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not get number of confirmations to wait: %w", err)
	}

	actualConfirmations := currentBlockNumber - request.BlockNumber
	hasEnoughConfirmations := actualConfirmations >= requiredConfirmations

	//nolint: gosec
	span.SetAttributes(
		attribute.Int64("current_block_number", int64(currentBlockNumber)),
		attribute.Int64("required_confirmations", int64(requiredConfirmations)),
		attribute.Int64("actual_confirmations", int64(actualConfirmations)),
		attribute.Bool("has_enough_confirmations", hasEnoughConfirmations),
	)

	return hasEnoughConfirmations, nil
}

// getNumberOfConfirmationsToWait returns the number of confirmations to wait for the request.
// confirmations = requestVolume / volumeLimitForChain, e.g. we wait 1 confirmation per every `volumeLimitForChain` USD.
func (l *limiterImpl) getNumberOfConfirmationsToWait(ctx context.Context, request *reldb.QuoteRequest) (_ uint64, err error) {
	ctx, span := l.metrics.Tracer().Start(ctx, "limiter.getNumberOfConfirmationsToWait")

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	tokenVolume, err := l.getRequestVolumeOfToken(ctx, request)
	if err != nil {
		return 0, fmt.Errorf("could not get token volume: %w", err)
	}

	volumeLimitForChain := l.cfg.GetVolumeLimit(int(request.Transaction.OriginChainId), request.Transaction.OriginToken)

	return uint64(new(big.Int).Div(
		tokenVolume,
		volumeLimitForChain,
	).Int64()), nil
}

// isUnderVolumeLimit returns true if the request is under the volume limit, false otherwise.
func (l *limiterImpl) isUnderVolumeLimit(ctx context.Context, request *reldb.QuoteRequest) (_ bool, err error) {
	ctx, span := l.metrics.Tracer().Start(ctx, "limiter.underVolumeLimitLimit")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	volumeLimit := l.cfg.GetVolumeLimit(int(request.Transaction.OriginChainId), request.Transaction.OriginToken)
	// There is no limit.
	if volumeLimit.Cmp(big.NewInt(-1)) == 0 {
		return true, nil
	}

	tokenPrice, err := l.getRequestVolumeOfToken(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not get USD amount of token: %w", err)
	}

	underVolumeLimit := tokenPrice.Cmp(volumeLimit) < 0
	span.SetAttributes(
		attribute.String("volume_limit", volumeLimit.String()),
		attribute.String("token_price", tokenPrice.String()),
		attribute.Bool("within_size_limit", underVolumeLimit),
	)

	return underVolumeLimit, nil
}

// getRequestVolumeOfToken returns the volume of the token in USD. This value is NOT human readable.
// We first get the price of the token in human readable units, then we multiply it by the OriginAmount.
func (l *limiterImpl) getRequestVolumeOfToken(
	ctx context.Context,
	request *reldb.QuoteRequest,
) (_ *big.Int, err error) {
	ctx, span := l.metrics.Tracer().Start(ctx, "limiter.getRequestVolumeOfToken")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var tokenName string
	for tn, tokenConfig := range l.tokenNames {
		if common.HexToAddress(tokenConfig.Address).Hex() == request.Transaction.OriginToken.Hex() {
			tokenName = tn
			break
		}
	}

	// Get the human readable price of the token.
	price, err := l.quoter.GetPrice(ctx, tokenName)
	if err != nil {
		return nil, fmt.Errorf("could not get price: %w", err)
	}
	priceFlt := new(big.Float).SetFloat64(price)
	// OriginAmount is NOT human readable. E.g. 3 USDC is 3000000 (3 * 10^6).
	originAmountFlt := new(big.Float).SetInt(request.Transaction.OriginAmount)

	product, _ := new(big.Float).Mul(priceFlt, originAmountFlt).Int(nil)

	span.SetAttributes(
		attribute.String("token_name", tokenName),
		attribute.Float64("price", price),
		attribute.String("origin_amount", request.Transaction.OriginAmount.String()),
		attribute.String("product", product.String()),
	)

	return product, nil
}

// checkReceipt checks if the receipt exists and has the correct fields in the case a reorg happened.
func (l *limiterImpl) checkReceipt(ctx context.Context, request *reldb.QuoteRequest) (_ bool, err error) {
	// Make sure receipt exists and has the correct fields in case of a reorg.
	// Note: https://community.infura.io/t/does-eth-gettransactionreceipt-respond-to-re-orged-transactions/7765
	// "You will get a tx receipt back but, as you note, there is a small chance of a reorg.
	// In a re-org the shorter side chain (one block usually, two occasionally) will have all its tx’s reverted and
	// placed back in the mempool. Calling eth_getTransactionReceipt at this point will return null, until the tx is
	// added to a new block and validated."
	ctx, span := l.metrics.Tracer().Start(ctx, "limiter.checkReceipt")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	receipt, err := l.evmClient.TransactionReceipt(ctx, request.OriginTxHash)
	if err != nil {
		if errors.Is(err, ethereum.NotFound) {
			return false, nil
		}
		return false, fmt.Errorf("could not get transaction receipt: %w", err)
	}

	rfqAddr, err := l.cfg.GetRFQAddress(int(request.Transaction.OriginChainId))
	if err != nil {
		return false, fmt.Errorf("could not get RFQ address: %w", err)
	}

	parser, err := fastbridge.NewParser(rfqAddr)
	if err != nil {
		return false, fmt.Errorf("could not create parser: %w", err)
	}

	for _, log := range receipt.Logs {
		_, parsedEvent, ok := parser.ParseEvent(*log)
		if !ok {
			continue
		}

		event, ok := parsedEvent.(*fastbridge.FastBridgeBridgeRequested)
		if ok {
			return rfqFieldsMatch(request, event) && !log.Removed, nil
		}
	}
	span.SetAttributes(
		attribute.String("receipt txHash", receipt.TxHash.Hex()),
		attribute.String("receipt log address", receipt.ContractAddress.String()),
		attribute.String("receipt data", hex.EncodeToString(receipt.Logs[0].Data)),
	)

	return false, nil
}

// TODO: is this exhaustive?
func rfqFieldsMatch(request *reldb.QuoteRequest, event *fastbridge.FastBridgeBridgeRequested) bool {
	transactionIDMatch := request.TransactionID == event.TransactionId
	senderMatch := request.Sender.String() == event.Sender.String()
	originAmountMatch := request.Transaction.OriginAmount.String() == event.OriginAmount.String()
	destAmountMatch := request.Transaction.DestAmount.String() == event.DestAmount.String()
	originTokenMatch := request.Transaction.OriginToken.String() == event.OriginToken.String()

	return transactionIDMatch && senderMatch && originAmountMatch && destAmountMatch && originTokenMatch
}
