package pricer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/core"

	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/instrumentation/httpcapture"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// exported constant used to apply special pricing behavior for USD as opposed to any token asset
// eg: If we are pricing an ETH value into US Dollars (not USDC) then we can supply USD_ as the "Price" asset
//
// TODO: find a better home for global constants like this
const USD_ = "USD"

// PriceFetcher is an interface for fetching prices from external sources.
//
//go:generate go run github.com/vektra/mockery/v2 --name PriceFetcher --output ./mocks --case=underscore
type PriceFetcher interface {
	GetPrice(ctx context.Context, token string) (float64, error)
}

// PriceFetcherImpl is an implementation of PriceFetcher.
type PriceFetcherImpl struct {
	handler   metrics.Handler
	client    *http.Client
	relConfig *relconfig.Config
}

var (
	priceCache = ttlcache.New(
		ttlcache.WithDisableTouchOnHit[string, float64](),
	)

	// Add a separate RWMutex for tokenIDLookup
	tokenIDMu sync.RWMutex

	inFlightRequests = make(map[string]*sync.Cond)

	inFlightRequestsMu sync.Mutex
)

// NewPriceFetcher creates a new instance of PriceFetcherImpl.
func NewPriceFetcher(handler metrics.Handler, timeout time.Duration, relConfig relconfig.Config) *PriceFetcherImpl {
	client := &http.Client{
		Timeout: timeout,
	}
	handler.ConfigureHTTPClient(client)
	client.Transport = httpcapture.NewCaptureTransport(client.Transport, handler)

	return &PriceFetcherImpl{
		handler:   handler,
		client:    client,
		relConfig: &relConfig,
	}
}

type PriceSourceDetail struct {
	Source        string
	SourceTokenID string
}

type TokenPriceConfig struct {
	// PrimaryPrice is the "actionable" amount upon which our relayer will actually quote & execute trades.
	// Thereby, it should be trusted, reliable, and real-time.
	PrimaryPrice PriceSourceDetail

	// VerificationPrice is used purely as a safety consensus to compare against the PrimaryPrice.
	// It is *not* used as a backup/fallback of any kind. Quotes/executions will never use verification prices.
	// It should still be trusted and reliable, but does not need to be as accurate or as real-time as PrimaryPrice.
	VerificationPrice PriceSourceDetail

	DeviationTolerancePct float64 // How much can primary and verification disagree with eachother before we stop returning a price?

	PriceCacheTTL int
}

// TODO: this would be more appropriate on config YAML
var tokenConfigs = map[string]TokenPriceConfig{
	"USDC": {
		PrimaryPrice: PriceSourceDetail{
			Source:        "Pyth",
			SourceTokenID: "0xeaa020c61cc479712813461ce153894a96a6c00b21ed0cfc2798d1f9a9e9c94a",
		},
		VerificationPrice: PriceSourceDetail{
			Source:        "USD",
			SourceTokenID: "n/a",
		},
		DeviationTolerancePct: 2,
		PriceCacheTTL:         30,
	},
	"ETH": {
		PrimaryPrice: PriceSourceDetail{
			Source:        "Pyth",
			SourceTokenID: "0xff61491a931112ddf1bd8147cd1b641375f79f5825126d665480874634fd0ace",
		},
		VerificationPrice: PriceSourceDetail{
			Source:        "CoinGecko",
			SourceTokenID: "ethereum",
		},
		DeviationTolerancePct: 5,
		PriceCacheTTL:         5,
	},
	"BERA": {
		PrimaryPrice: PriceSourceDetail{
			Source:        "Pyth",
			SourceTokenID: "0x962088abcfdbdb6e30db2e340c8cf887d9efb311b1f2f17b155a63dbb6d40265",
		},
		VerificationPrice: PriceSourceDetail{
			Source:        "CoinGecko",
			SourceTokenID: "berachain-bera",
		},
		DeviationTolerancePct: 5,
		PriceCacheTTL:         5,
	},
	"HYPE": {
		PrimaryPrice: PriceSourceDetail{
			Source:        "KuCoin",
			SourceTokenID: "HYPE",
		},
		VerificationPrice: PriceSourceDetail{
			Source:        "CoinGecko",
			SourceTokenID: "hyperliquid",
		},
		DeviationTolerancePct: 5,
		PriceCacheTTL:         5,
	},
}

// Store original token ID map values for reset capability
var originalTokenConfig map[string]TokenPriceConfig

// init function to store the original map values
func init() {
	originalTokenConfig = make(map[string]TokenPriceConfig)
	for k, v := range tokenConfigs {
		originalTokenConfig[k] = v
	}
}

// UnsafeGetTokenConfigMap returns a copy of the tokenConfigs map.
// This method is only meant to be used in test environments and will return an error if used in production.
func UnsafeGetTokenConfigMap() (map[string]TokenPriceConfig, error) {
	// Only allow in test environments
	if os.Getenv("GO_ENVIRONMENT") != "test" && os.Getenv("GO_ENV") != "test" && os.Getenv("ENVIRONMENT") != "test" {
		return nil, fmt.Errorf("UnsafeGetTokenConfigMap can only be called in test environments")
	}

	tokenIDMu.RLock()
	defer tokenIDMu.RUnlock()

	// Make a copy of the map
	result := make(map[string]TokenPriceConfig)
	for k, v := range tokenConfigs {
		result[k] = v
	}

	return result, nil
}

// UnsafeUpdateTokenConfigMap updates the tokenConfigs map with the provided entries.
// This method is only meant to be used in test environments and will return an error if used in production.
func UnsafeUpdateTokenConfigMap(entries map[string]TokenPriceConfig) error {
	// Only allow updates in test environments
	if !(os.Getenv("GO_ENVIRONMENT") != "test" || os.Getenv("GO_ENV") != "test" || os.Getenv("ENVIRONMENT") != "test" || core.HasEnv("CI")) {
		panic("UnsafeUpdateTokenConfigMap can only be called in test environments")
	}

	tokenIDMu.Lock()
	defer tokenIDMu.Unlock()

	// Update the map with new entries
	for token, config := range entries {
		tokenConfigs[token] = config
	}

	return nil
}

// UnsafeResetTokenConfigMap resets the tokenConfigs map to its original state.
// This method is only meant to be used in test environments and will return an error if used in production.
func UnsafeResetTokenConfigMap() error {
	// Only allow resets in test environments
	if os.Getenv("GO_ENVIRONMENT") != "test" && os.Getenv("GO_ENV") != "test" && os.Getenv("ENVIRONMENT") != "test" {
		return fmt.Errorf("UnsafeResetTokenConfigMap can only be called in test environments")
	}

	tokenIDMu.Lock()
	defer tokenIDMu.Unlock()

	// Clear current map
	for k := range tokenConfigs {
		delete(tokenConfigs, k)
	}

	// Restore original values
	for k, v := range originalTokenConfig {
		tokenConfigs[k] = v
	}

	return nil
}

// fetchPriceExternal fetches a token price from external sources, validates it, and caches it
func (c *PriceFetcherImpl) fetchPriceExternal(ctx context.Context, token string, tokenConfig TokenPriceConfig) (price float64, err error) {

	price, err = c.getExternalPrice(ctx, tokenConfig.PrimaryPrice)
	if err != nil {
		return 0, fmt.Errorf("err getExternalPrice pri: %w", err)
	}

	var verificationPrice float64

	if tokenConfig.VerificationPrice.Source == "None" {
		// If the verification source is "None", this effectively means we have no verification consensus protection configured for this asset.
		// Duplicate PrimaryPrice & move along.
		verificationPrice = price
	} else {
		// Otherwise, use the verification source
		verificationPrice, err = c.getExternalPrice(ctx, tokenConfig.VerificationPrice)
		if err != nil {
			return 0, fmt.Errorf("err getExternalPrice sec: %w", err)
		}
	}

	// confirm reasonable consensus btwn pri and sec price sources
	percentageDiff := math.Abs(verificationPrice-price) / price * 100

	if percentageDiff > tokenConfig.DeviationTolerancePct {
		// Remove any existing cache of this price due to the consensus failure
		priceCache.Delete(token)

		return 0, fmt.Errorf("err price consensus: pri = $%f, sec = $%f, %f%% diff", price, verificationPrice, percentageDiff)
	}

	// add "fetchprice" to debugOutput env var for dev/debug output
	if strings.Contains(strings.ToLower(os.Getenv("debugOutput")), "fetchprice") {
		fmt.Printf("%s, Token: %6s, Cached:  No, PrimaryPrice: $%15f, VerificationPrice: $%15f, Diff: %f%%\n", time.Now().Format(time.RFC3339), token, price, verificationPrice, percentageDiff)
	}

	// cache this price so we can minimize external calls
	priceCache.Set(token, price, time.Duration(tokenConfig.PriceCacheTTL)*time.Second)

	return price, nil
}

// GetPrice fetches the USD price of a token from external sources (or a cache, if available)
func (c *PriceFetcherImpl) GetPrice(ctx context.Context, token string) (price float64, err error) {
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

	tokenIDMu.RLock()
	tokenConfig, exists := tokenConfigs[token]
	tokenIDMu.RUnlock()

	if !exists {
		return 0, fmt.Errorf("err tokenConfig not found: %s", token)
	}

	// Check price cache first before making any new external calls
	cachedItem := priceCache.Get(token)
	if cachedItem != nil {
		cached := cachedItem.Value()
		span.SetAttributes(attribute.Bool("used_cached_price", true))

		// add "fetchprice" to debugOutput env var for dev/debug output
		if strings.Contains(strings.ToLower(os.Getenv("debugOutput")), "fetchprice") {
			fmt.Printf("%s, Token: %6s, Cached: Yes, PrimaryPrice: $%15f\n", time.Now().Format(time.RFC3339), token, cached)
		}

		return cached, nil
	}

	inFlightRequestsMu.Lock()
	cond, inFlight := inFlightRequests[token]

	// if we are not already in the process of fetching this token's price via some separate call, then initiate a fresh fetch
	if !inFlight {
		cond = sync.NewCond(&inFlightRequestsMu)
		inFlightRequests[token] = cond

		inFlightRequestsMu.Unlock()

		span.SetAttributes(attribute.Bool("used_cached_false", true))
		price, err := c.fetchPriceExternal(ctx, token, tokenConfig)

		if err != nil {
			return 0, fmt.Errorf("err fetchPriceExternal: %w", err)
		}

		inFlightRequestsMu.Lock()
		delete(inFlightRequests, token)
		inFlightRequestsMu.Unlock()

		cond.Broadcast()

		return price, nil
	}
	// Otherwise, wait until the in-flight price fetch is completed & then we can share the same result via cache
	cond.Wait()
	inFlightRequestsMu.Unlock()

	// recurse, which should now pull from the cache
	return c.GetPrice(ctx, token)
}

// getExternalPrice gets the price of a token from an external source
func (c *PriceFetcherImpl) getExternalPrice(ctx context.Context, priceSource PriceSourceDetail) (float64, error) {
	var price float64
	var err error

	switch priceSource.Source {
	case "CoinGecko":
		price, err = c.getPriceCoinGecko(ctx, priceSource.SourceTokenID)
		if err != nil {
			return 0, fmt.Errorf("getPriceCoinGecko: %w", err)
		}
		return price, nil

	case "KuCoin":
		price, err = c.getPriceKucoin(ctx, priceSource.SourceTokenID)
		if err != nil {
			return 0, fmt.Errorf("getPriceKucoin: %w", err)
		}
		return price, nil

	case "Pyth":
		price, err = c.getPricePyth(ctx, priceSource.SourceTokenID)
		if err != nil {
			return 0, fmt.Errorf("getPricePyth: %w", err)
		}
		return price, nil

	case "USD": // this can be used to force an asset's primary or verification price to be 1-to-1 with the US dollar
		return 1.0, nil

	default:
		return 0, fmt.Errorf("err unknown source: %s", priceSource.Source)
	}
}

// getPricePyth gets the price of a token from Pyth
func (c *PriceFetcherImpl) getPricePyth(ctx context.Context, pythTokenId string) (float64, error) {
	var price float64

	url := fmt.Sprintf("https://hermes.pyth.network/v2/updates/price/latest?ids[]=%s&parsed=true", pythTokenId)

	// fetch price from Pyth
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("err building request: %w", err)
	}
	r, err := c.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("err sending request: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("err response status: %v", r.Status)
	}
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			fmt.Printf("Warning: getPricePyth r.Body.Close() fail: %v\n", closeErr)
		}
	}()

	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, fmt.Errorf("err response unreadable: %w", err)
	}

	// parse the price from the response
	var resp struct {
		Parsed []struct {
			Price struct {
				Price       string `json:"price"`
				Expo        int    `json:"expo"`
				PublishTime int64  `json:"publish_time"`
			} `json:"price"`
		} `json:"parsed"`
	}
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return 0, fmt.Errorf("err unmarshal fail: %w", err)
	}

	if len(resp.Parsed) == 0 {
		return 0, fmt.Errorf("err no price data found")
	}

	// do not consider price valid if it is stale by more than X minutes (ie: apparent Pyth failure)
	currentTime := time.Now().Unix()
	if currentTime-resp.Parsed[0].Price.PublishTime > 5*60 {
		return 0, fmt.Errorf("err stale price (%d)", resp.Parsed[0].Price.PublishTime)
	}

	// Convert the price to normal units using the expo value from the API response
	rawPrice, err := strconv.ParseFloat(resp.Parsed[0].Price.Price, 64)
	if err != nil {
		return 0, fmt.Errorf("err parsing price: %w", err)
	}
	expo := resp.Parsed[0].Price.Expo
	price = rawPrice * math.Pow10(expo)

	return price, nil
}

// getPriceKucoin fetches the price of a token from Kucoin.
func (c *PriceFetcherImpl) getPriceKucoin(ctx context.Context, kucoinTokenId string) (float64, error) {
	var price float64

	url := fmt.Sprintf("https://api.kucoin.com/api/v1/prices?base=USD&currencies=%s", kucoinTokenId)

	// fetch price from kucoin
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("err building request: %w", err)
	}
	r, err := c.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("err sending request: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("err response status: %v", r.Status)
	}
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			fmt.Printf("Warning: getPriceKucoin r.Body.Close() fail: %v\n", closeErr)
		}
	}()

	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, fmt.Errorf("err response unreadable: %w", err)
	}

	// parse the price
	var resp struct {
		Code string            `json:"code"`
		Data map[string]string `json:"data"`
	}
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return 0, fmt.Errorf("err unmarshal fail: %w", err)
	}

	// 200000  =  Success
	if resp.Code != "200000" {
		return 0, fmt.Errorf("err response code: %s", resp.Code)
	}

	// Extract the price for the given token ID
	if val, ok := resp.Data[kucoinTokenId]; ok {
		price, err = strconv.ParseFloat(val, 64)
		if err != nil {
			return 0, fmt.Errorf("err parsing price: %w", err)
		}
	} else {
		return 0, fmt.Errorf("err no price data found for token: %s", kucoinTokenId)
	}

	return price, nil
}

// getPriceCoinGecko fetches the price of a token from CoinGecko.
func (c *PriceFetcherImpl) getPriceCoinGecko(ctx context.Context, coingeckoTokenId string) (float64, error) {
	var price float64

	baseURL := c.relConfig.CoinGeckoAPIURL
	if baseURL == "" {
		if c.relConfig.CoinGeckoAPIKey != "" {
			baseURL = "https://pro-api.coingecko.com/api/v3"
		} else {
			baseURL = "https://api.coingecko.com/api/v3"
		}
	}

	url := fmt.Sprintf("%s/simple/price?ids=%s&vs_currencies=USD", baseURL, coingeckoTokenId)

	// attach pro API key if supplied
	if c.relConfig.CoinGeckoAPIKey != "" && c.relConfig.CoinGeckoAPIURL == "" {
		url += fmt.Sprintf("&x_cg_pro_api_key=%s", c.relConfig.CoinGeckoAPIKey)
	}

	// fetch price from coingecko
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return 0, fmt.Errorf("err building request: %w", err)
	}
	r, err := c.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("err sending request: %w", err)
	}
	if r.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("err response status: %v", r.Status)
	}
	defer func() {
		closeErr := r.Body.Close()
		if closeErr != nil {
			fmt.Printf("Warning: getPriceCoinGecko r.Body.Close() fail: %v\n", closeErr)
		}
	}()

	respBytes, err := io.ReadAll(r.Body)
	if err != nil {
		return 0, fmt.Errorf("err response unreadable: %w", err)
	}

	// parse the price
	var resp map[string]map[string]float64
	err = json.Unmarshal(respBytes, &resp)
	if err != nil {
		return 0, fmt.Errorf("err unmarshal fail: %w", err)
	}
	price, ok := resp[coingeckoTokenId]["usd"]

	if !ok {
		return 0, fmt.Errorf("err usd price not found: %v", resp)
	}

	return price, nil
}
