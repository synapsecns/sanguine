package pricer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/core"

	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation/httpcapture"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// exported constant used to apply special pricing behavior for USD as opposed to any token asset
// eg: If we are pricing an ETH value into US Dollars (not USDC) then we can supply USD_ as the "Price" asset
const USD_ = "USD"

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

	// Add a separate RWMutex for coingeckoIDLookup
	coingeckoMu sync.RWMutex
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

// Store original coingecko map values for reset capability
var originalCoingeckoIDLookup map[string]string

// init function to store the original map values
func init() {
	originalCoingeckoIDLookup = make(map[string]string)
	for k, v := range coingeckoIDLookup {
		originalCoingeckoIDLookup[k] = v
	}
}

// UnsafeGetCoingeckoIDMap returns a copy of the coingeckoIDLookup map.
// This method is only meant to be used in test environments and will return an error if used in production.
func UnsafeGetCoingeckoIDMap() (map[string]string, error) {
	// Only allow in test environments
	if os.Getenv("GO_ENVIRONMENT") != "test" && os.Getenv("GO_ENV") != "test" && os.Getenv("ENVIRONMENT") != "test" {
		return nil, fmt.Errorf("UnsafeGetCoingeckoIDMap can only be called in test environments")
	}

	coingeckoMu.RLock()
	defer coingeckoMu.RUnlock()

	// Make a copy of the map
	result := make(map[string]string)
	for k, v := range coingeckoIDLookup {
		result[k] = v
	}

	return result, nil
}

// UnsafeUpdateCoingeckoMap updates the coingeckoIDLookup map with the provided entries.
// This method is only meant to be used in test environments and will return an error if used in production.
func UnsafeUpdateCoingeckoMap(entries map[string]string) error {
	// Only allow updates in test environments
	if !(os.Getenv("GO_ENVIRONMENT") != "test" || os.Getenv("GO_ENV") != "test" || os.Getenv("ENVIRONMENT") != "test" || core.HasEnv("CI")) {
		panic("UnsafeUpdateCoingeckoMap can only be called in test environments")
	}

	coingeckoMu.Lock()
	defer coingeckoMu.Unlock()

	// Update the map with new entries
	for token, coinGeckoID := range entries {
		coingeckoIDLookup[token] = coinGeckoID
	}

	return nil
}

// UnsafeResetCoingeckoMap resets the coingeckoIDLookup map to its original state.
// This method is only meant to be used in test environments and will return an error if used in production.
func UnsafeResetCoingeckoMap() error {
	// Only allow resets in test environments
	if os.Getenv("GO_ENVIRONMENT") != "test" && os.Getenv("GO_ENV") != "test" && os.Getenv("ENVIRONMENT") != "test" {
		return fmt.Errorf("UnsafeResetCoingeckoMap can only be called in test environments")
	}

	coingeckoMu.Lock()
	defer coingeckoMu.Unlock()

	// Clear current map
	for k := range coingeckoIDLookup {
		delete(coingeckoIDLookup, k)
	}

	// Restore original values
	for k, v := range originalCoingeckoIDLookup {
		coingeckoIDLookup[k] = v
	}

	return nil
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

	// USD_ is a special identifier we can use to price assets directly to USD while also following roughly the same logic as any other asset.
	if token == USD_ {
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

	coingeckoMu.RLock()
	coingeckoID, ok := coingeckoIDLookup[token]
	coingeckoMu.RUnlock()

	if !ok {
		return price, fmt.Errorf("could not get coingecko id for token: %s", token)
	}

	apiKey := c.relConfig.CoinGeckoApiKey
	var baseURL string

	// Use the custom CoinGecko API URL if configured, otherwise use the default
	if c.relConfig.CoinGeckoAPIURL != "" {
		baseURL = c.relConfig.CoinGeckoAPIURL
	} else if apiKey != "" {
		baseURL = "https://pro-api.coingecko.com/api/v3"
	} else {
		baseURL = "https://api.coingecko.com/api/v3"
	}

	var url string
	if apiKey != "" && c.relConfig.CoinGeckoAPIURL == "" { // Only add API key if using official endpoint
		url = fmt.Sprintf("%s/simple/price?ids=%s&vs_currencies=USD&x_cg_pro_api_key=%s", baseURL, coingeckoID, apiKey)
	} else {
		url = fmt.Sprintf("%s/simple/price?ids=%s&vs_currencies=USD", baseURL, coingeckoID)
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
