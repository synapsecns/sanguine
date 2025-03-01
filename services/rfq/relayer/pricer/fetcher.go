package pricer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation/httpcapture"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// CoingeckoPriceFetcher is an interface for fetching prices from coingecko.
//
//go:generate go run github.com/vektra/mockery/v2 --name CoingeckoPriceFetcher --output ./mocks --case=underscore
type CoingeckoPriceFetcher interface {
	GetPrice(ctx context.Context, token string) (float64, error)
}

// CoingeckoPriceFetcherImpl is an implementation of CoingeckoPriceFetcher.
type CoingeckoPriceFetcherImpl struct {
	handler   metrics.Handler
	client    *http.Client
	relConfig *relconfig.Config
}

type cachedPrice struct {
	price      float64
	source     string
	timestamp  time.Time
	timeToLive time.Duration
}

var (
	globalCache = make(map[string]cachedPrice)
	cacheMu     sync.Mutex
)

// NewCoingeckoPriceFetcher creates a new instance of CoingeckoPriceFetcherImpl.
func NewCoingeckoPriceFetcher(handler metrics.Handler, timeout time.Duration, relConfig relconfig.Config) *CoingeckoPriceFetcherImpl {
	client := &http.Client{
		Timeout: timeout,
	}
	handler.ConfigureHTTPClient(client)
	client.Transport = httpcapture.NewCaptureTransport(client.Transport, handler)

	return &CoingeckoPriceFetcherImpl{
		handler:   handler,
		client:    client,
		relConfig: &relConfig,
	}
}

var coingeckoIDLookup = map[string]string{
	"USDC": "usd-coin",
	"ETH":  "ethereum",
	"BERA": "berachain-bera",
	"HYPE": "hyperliquid",
}

// GetPrice fetches the price of a token from coingecko.
func (c *CoingeckoPriceFetcherImpl) GetPrice(ctx context.Context, token string) (price float64, err error) {
	ctx, span := c.handler.Tracer().Start(ctx, "GetPrice", trace.WithAttributes(
		attribute.String("token", token),
	))

	defer func() {
		span.SetAttributes(attribute.Float64("price", price))
		metrics.EndSpanWithErr(span, err)
	}()

	// "DirectUSD" is a special identifier we can use to price assets directly to USD while also following roughly the same logic as any other asset.
	if token == "DirectUSD" {
		price = 1.0
		return price, nil
	}

	// Check global cache first before calling api
	cacheMu.Lock()
	cached, found := globalCache[token]
	cacheMu.Unlock()

	if found && time.Since(cached.timestamp) < cached.timeToLive {
		span.SetAttributes(attribute.Bool("used_cached_price", true))
		return cached.price, nil
	}

	span.SetAttributes(attribute.Bool("used_cached_price", false))
	coingeckoID, ok := coingeckoIDLookup[token]
	if !ok {
		return price, fmt.Errorf("could not get coingecko id for token: %s", token)
	}
	apiKey := c.relConfig.CoinGeckoApiKey
	var url string
	if apiKey != "" {
		url = fmt.Sprintf("https://pro-api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=USD&x_cg_pro_api_key=%s", coingeckoID, apiKey)
	} else {
		url = fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=USD", coingeckoID)
	}

	// fetch price from coingecko
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return price, fmt.Errorf("could not build request: %w", err)
	}
	r, err := c.client.Do(req)
	if err != nil {
		return price, fmt.Errorf("could not get price from coingecko: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return price, fmt.Errorf("coingecko stat code: %v", r.Status)
	}
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
		}
	}()

	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return price, fmt.Errorf("could not read response body: %w", err)
	}

	// parse the price
	var resp map[string]map[string]float64
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return price, fmt.Errorf("could not unmarshal response body: %w", err)
	}
	price, ok = resp[coingeckoID]["usd"]

	if !ok {
		return price, fmt.Errorf("could not get price from coingecko response: %v", resp)
	}

	//todo: make this actually work
	secondarySource := "None"

	if secondarySource != "None" {
		// obtain secondary price for the asset
		secondaryPrice := 5.0 //, err := c.GetSecondaryPrice(ctx, token)

		if err != nil {
			return price, fmt.Errorf("err GetSecondaryPrice fail: %w", err)
		}

		// confirm reasonable consensus btwn pri and sec price sources
		percentageDiff := (secondaryPrice - price) / price * 100
		if percentageDiff < 0 {
			percentageDiff = -percentageDiff
		}
		if percentageDiff > 5 {
			return price, fmt.Errorf("err price consensus: pri = $%f, sec = $%f, %f%% diff", price, secondaryPrice, percentageDiff)
		}
	}

	// Update global cache
	cacheMu.Lock()
	ttl := 3 * time.Second // Default TTL

	// Longer TTL for some tokens
	if token == "USDC" {
		ttl = 10 * time.Minute
	}

	globalCache[token] = cachedPrice{
		price:      price,
		source:     "coingecko",
		timestamp:  time.Now(),
		timeToLive: ttl, // Set the TTL for the cache entry
	}
	cacheMu.Unlock()

	return price, nil
}

func (c *CoingeckoPriceFetcherImpl) GetSecondaryPrice(ctx context.Context, token string) (price float64, err error) {
	return 1, nil
}
