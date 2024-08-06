package limiter

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	omniClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/quoter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"go.uber.org/ratelimit"
)

type Conditional func() bool

// Limiter is the interface for rate limiting RFQs.
type Limiter interface {
	// IsAllowed returns true if the request is allowed, false otherwise.
	IsAllowed(ctx context.Context, request reldb.QuoteRequest) bool
	// Take blocks until a token is available.
	Take() time.Time
}

// essentially a wrapper around golang.org/x/time/rate.Limiter but we keep interfacability
type limiterImpl struct {
	ratelimit.Limiter

	client     omniClient.RPCClient
	quoter     quoter.Quoter
	cfg        relconfig.Config
	tokenNames map[string]relconfig.TokenConfig
}

// NewRateLimiter creates a new Limiter.
func NewRateLimiter(cfg relconfig.Config, q quoter.Quoter, metricHandler metrics.Handler, req reldb.QuoteRequest) Limiter {
	omniClient := omniClient.NewOmnirpcClient(cfg.OmniRPCURL, metricHandler, omniClient.WithCaptureReqRes())

	return &limiterImpl{
		Limiter: ratelimit.New(cfg.BlockWindow),

		client:     omniClient,
		quoter:     q,
		cfg:        cfg,
		tokenNames: cfg.Chains[int(req.Transaction.OriginChainId)].Tokens,
	}
}

// IsAllowed returns true if the request is allowed, false otherwise.
func (l *limiterImpl) IsAllowed(ctx context.Context, request reldb.QuoteRequest) bool {
	return l.withinVolumeLimit(ctx, request) && l.hasEnoughConfirmations(ctx, request)
}

func (l *limiterImpl) hasEnoughConfirmations(ctx context.Context, request reldb.QuoteRequest) bool {
	cc, err := l.client.GetChainClient(ctx, int(request.Transaction.OriginChainId))
	if err != nil {
		return false
	}

	currentBlockNumber, err := cc.BlockNumber(ctx)
	if err != nil {
		return false
	}

	return currentBlockNumber-request.BlockNumber >= l.cfg.BaseChainConfig.Confirmations
}

func (l *limiterImpl) withinVolumeLimit(ctx context.Context, request reldb.QuoteRequest) bool {
	tokenPrice, err := l.getUSDAmountOfToken(ctx, l.quoter, request)
	if err != nil {
		return false
	}
	return tokenPrice <= l.cfg.VolumeLimit
}

func (l *limiterImpl) getUSDAmountOfToken(ctx context.Context, q quoter.Quoter, request reldb.QuoteRequest) (float64, error) {
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
