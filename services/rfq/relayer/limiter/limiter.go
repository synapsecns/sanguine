// Package limiter provides a rate limiting mechanism for RFQs to protect against reorgs.
package limiter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpc "github.com/synapsecns/sanguine/services/omnirpc/client"
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
	listener      LatestBlockFetcher
	metrics       metrics.Handler
	quoter        quoter.Quoter
	cfg           relconfig.Config
	tokenNames    map[string]relconfig.TokenConfig
	omnirpcClient omnirpc.RPCClient
}

// NewRateLimiter creates a new Limiter.
// TODO: implement the sliding window: queue up requests and process them in order if cumulative volume is above limit.
func NewRateLimiter(
	cfg relconfig.Config,
	l LatestBlockFetcher,
	q quoter.Quoter,
	metricHandler metrics.Handler,
	tokens map[string]relconfig.TokenConfig,
	omnirpcClient omnirpc.RPCClient,
) Limiter {
	return &limiterImpl{
		listener:      l,
		metrics:       metricHandler,
		quoter:        q,
		cfg:           cfg,
		tokenNames:    tokens,
		omnirpcClient: omnirpcClient,
	}
}

// IsAllowed returns true if the request is allowed, false otherwise.
func (l *limiterImpl) IsAllowed(ctx context.Context, request *reldb.QuoteRequest) (_ bool, err error) {
	ctx, span := l.metrics.Tracer().Start(
		ctx, "limiter.IsAllowed", trace.WithAttributes(util.QuoteRequestToAttributes(*request)...),
	)

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// if enough confirmations, allow because reorgs are rare at this point
	hasEnoughConfirmations, err := l.hasEnoughConfirmations(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not check confirmations: %w", err)
	}
	if hasEnoughConfirmations {
		return true, nil
	}

	// if not enough confirmations, check volume
	withinSize, err := l.withinSizeLimit(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not check volume limit: %w", err)
	}

	span.SetAttributes(
		attribute.Bool("has_enough_confirmations", hasEnoughConfirmations),
		attribute.Bool("within_size_limit", withinSize),
	)

	receiptFieldsMatch, err := l.checkReceipt(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not check receipt: %w", err)
	}

	return withinSize && receiptFieldsMatch, nil
}

func (l *limiterImpl) hasEnoughConfirmations(ctx context.Context, request *reldb.QuoteRequest) (_ bool, err error) {
	_, span := l.metrics.Tracer().Start(ctx, "limiter.hasEnoughConfirmations")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	currentBlockNumber := l.listener.LatestBlock()
	if err != nil {
		return false, fmt.Errorf("could not get block number: %w", err)
	}

	requiredConfirmations, err := l.cfg.GetFinalityConfirmations(int(request.Transaction.OriginChainId))
	if err != nil {
		return false, fmt.Errorf("could not get required confirmations from config: %w", err)
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

func (l *limiterImpl) withinSizeLimit(ctx context.Context, request *reldb.QuoteRequest) (_ bool, err error) {
	ctx, span := l.metrics.Tracer().Start(ctx, "limiter.withinSizeLimit")
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

	withinSizeLimit := tokenPrice.Cmp(volumeLimit) < 0
	span.SetAttributes(
		attribute.String("volume_limit", volumeLimit.String()),
		attribute.String("token_price", tokenPrice.String()),
		attribute.Bool("within_size_limit", withinSizeLimit),
	)

	return withinSizeLimit, nil
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

func (l *limiterImpl) checkReceipt(ctx context.Context, request *reldb.QuoteRequest) (bool, error) {
	confirmationsClient, err := l.omnirpcClient.GetConfirmationsClient(
		ctx,
		int(request.Transaction.OriginChainId),
		int(l.cfg.GetRPCConfirmations()),
	)
	if err != nil {
		return false, fmt.Errorf("could not get confirmations client: %w", err)
	}

	// make sure receipt exists and has the correct fields in case of a reorg
	receipt, err := confirmationsClient.TransactionReceipt(ctx, request.OriginTxHash)
	if err != nil {
		return false, fmt.Errorf("could not check for receipt: %w", err)
	}

	// not sure if this is needed.
	if receipt.Logs[0] == nil {
		return false, fmt.Errorf("no logs in receipt")
	}
	log := receipt.Logs[0]

	// nonce check
	if log.Topics[1] != request.TransactionID {
		return false, fmt.Errorf("incorrect transactionID got %s expected %s", log.Topics[1].String(), hexutil.Encode((request.TransactionID[:])))
	}

	parser, err := fastbridge.NewParser(common.HexToAddress(""))
	if err != nil {
		return false, fmt.Errorf("could not create parser: %w", err)
	}

	_, parsedEvent, ok := parser.ParseEvent(*log)
	if !ok {
		return false, fmt.Errorf("could not parse event")
	}

	switch event := parsedEvent.(type) {
	case *fastbridge.FastBridgeBridgeRequested:
		return rfqFieldsMatch(request, event), nil
	default:
		return false, fmt.Errorf("failed to decode event: unknown event")
	}
}

func rfqFieldsMatch(request *reldb.QuoteRequest, event *fastbridge.FastBridgeBridgeRequested) bool {
	return request.TransactionID == event.TransactionId &&
		request.Sender.String() == event.Sender.String() &&
		request.Transaction.OriginAmount.String() == event.OriginAmount.String() &&
		request.Transaction.DestAmount.String() == event.DestAmount.String() &&
		request.Transaction.OriginToken.String() == event.OriginToken.String()
}
