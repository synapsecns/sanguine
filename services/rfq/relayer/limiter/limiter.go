package limiter

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.uber.org/ratelimit"
)

// Limiter is the interface for rate limiting RFQs.
type Limiter interface {
	// IsAllowed returns true if the request is allowed, false otherwise.
	IsAllowed(ctx context.Context, request reldb.QuoteRequest) (bool, error)
	// Take blocks until a token is available.
	Take() time.Time
}

type limiterImpl struct {
	// TODO: Possibly unneeded?
	ratelimit.Limiter

	client     client.EVM
	quoter     quoter.Quoter
	cfg        relconfig.Config
	tokenNames map[string]relconfig.TokenConfig
}

// NewRateLimiter creates a new Limiter.
// TODO: implement the sliding window: queue up requests and process them in order if cumulative volume is above limit
func NewRateLimiter(
	cfg relconfig.Config,
	q quoter.Quoter,
	client client.EVM,
	metricHandler metrics.Handler,
	tokens map[string]relconfig.TokenConfig,
) Limiter {
	return &limiterImpl{
		Limiter:    ratelimit.New(cfg.MaxRFQSize),
		client:     client,
		quoter:     q,
		cfg:        cfg,
		tokenNames: tokens,
	}
}

// IsAllowed returns true if the request is allowed, false otherwise.
func (l *limiterImpl) IsAllowed(ctx context.Context, request reldb.QuoteRequest) (bool, error) {
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

func (l *limiterImpl) hasEnoughConfirmations(ctx context.Context, request reldb.QuoteRequest) (bool, error) {
	currentBlockNumber, err := l.client.BlockNumber(ctx)
	if err != nil {
		return false, fmt.Errorf("could not get block number: %w", err)
	}
	requiredConfirmations, err := l.cfg.GetConfirmations(int(request.Transaction.OriginChainId))
	if err != nil {
		return false, fmt.Errorf("could not get required confirmations from config: %w", err)
	}

	actualConfirmations := currentBlockNumber - request.BlockNumber

	return actualConfirmations >= requiredConfirmations, nil
}

func (l *limiterImpl) withinSizeLimit(ctx context.Context, request reldb.QuoteRequest) (bool, error) {
	tokenPrice, err := l.getUSDAmountOfToken(ctx, l.quoter, request)
	if err != nil {
		return false, fmt.Errorf("could not get USD amount of token: %w", err)
	}
	return tokenPrice <= l.cfg.VolumeLimit, nil
}

func (l *limiterImpl) getUSDAmountOfToken(
	ctx context.Context,
	q quoter.Quoter,
	request reldb.QuoteRequest,
) (float64, error) {
	var tokenName string
	for tn, tokenConfig := range l.tokenNames {
		if common.HexToAddress(tokenConfig.Address).Hex() == request.Transaction.OriginToken.Hex() {
			tokenName = tn
		}
	}

	price, err := q.GetPrice(ctx, tokenName)
	if err != nil {
		return 0, fmt.Errorf("could not get price: %w", err)
	}

	return price * float64(request.Transaction.OriginAmount.Int64()), nil
}
