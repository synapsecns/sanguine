// Package limiter provides a rate limiting mechanism for RFQs to protect against reorgs.
package limiter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type LatestBlockFetcher interface {
	LatestBlock() uint64
}

// Limiter is the interface for rate limiting RFQs.
type Limiter interface {
	// IsAllowed returns true if the request is allowed, false otherwise.
	IsAllowed(ctx context.Context, request reldb.QuoteRequest) (bool, error)
}

type limiterImpl struct {
	listener   LatestBlockFetcher
	metrics    metrics.Handler
	quoter     quoter.Quoter
	cfg        relconfig.Config
	tokenNames map[string]relconfig.TokenConfig
}

// NewRateLimiter creates a new Limiter.
// TODO: implement the sliding window: queue up requests and process them in order if cumulative volume is above limit.
func NewRateLimiter(
	cfg relconfig.Config,
	l LatestBlockFetcher,
	q quoter.Quoter,
	metricHandler metrics.Handler,
	tokens map[string]relconfig.TokenConfig,
) Limiter {
	return &limiterImpl{
		listener:   l,
		metrics:    metricHandler,
		quoter:     q,
		cfg:        cfg,
		tokenNames: tokens,
	}
}

// IsAllowed returns true if the request is allowed, false otherwise.
func (l *limiterImpl) IsAllowed(ctx context.Context, request reldb.QuoteRequest) (_ bool, err error) {
	ctx, span := l.metrics.Tracer().Start(
		ctx, "limiter.IsAllowed", trace.WithAttributes(util.QuoteRequestToAttributes(request)...),
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

	return withinSize, nil
}

func (l *limiterImpl) hasEnoughConfirmations(ctx context.Context, request reldb.QuoteRequest) (_ bool, err error) {
	_, span := l.metrics.Tracer().Start(ctx, "limiter.hasEnoughConfirmations")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	currentBlockNumber := l.listener.LatestBlock()
	if err != nil {
		return false, fmt.Errorf("could not get block number: %w", err)
	}

	requiredConfirmations, err := l.cfg.GetConfirmations(int(request.Transaction.OriginChainId))
	if err != nil {
		return false, fmt.Errorf("could not get required confirmations from config: %w", err)
	}

	actualConfirmations := currentBlockNumber - request.BlockNumber
	hasEnoughConfirmations := actualConfirmations >= requiredConfirmations

	span.SetAttributes(
		attribute.Int64("current_block_number", int64(currentBlockNumber)),
		attribute.Int64("required_confirmations", int64(requiredConfirmations)),
		attribute.Int64("actual_confirmations", int64(actualConfirmations)),
		attribute.Bool("has_enough_confirmations", hasEnoughConfirmations),
	)

	return hasEnoughConfirmations, nil
}

func (l *limiterImpl) withinSizeLimit(ctx context.Context, request reldb.QuoteRequest) (bool, error) {
	volumeLimit := l.cfg.GetVolumeLimit(int(request.Transaction.OriginChainId), request.Transaction.OriginToken)
	// There is no limit.
	if volumeLimit.Cmp(big.NewInt(-1)) == 0 {
		return true, nil
	}

	tokenPrice, err := l.getUSDAmountOfToken(ctx, request)
	if err != nil {
		return false, fmt.Errorf("could not get USD amount of token: %w", err)
	}

	return tokenPrice.Cmp(volumeLimit) < 0, nil
}

func (l *limiterImpl) getUSDAmountOfToken(
	ctx context.Context,
	request reldb.QuoteRequest,
) (_ *big.Int, err error) {
	ctx, span := l.metrics.Tracer().Start(ctx, "limiter.getUSDAmountOfToken")
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

	price, err := l.quoter.GetPrice(ctx, tokenName)
	if err != nil {
		return nil, fmt.Errorf("could not get price: %w", err)
	}
	priceFlt := new(big.Float).SetFloat64(price)
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
