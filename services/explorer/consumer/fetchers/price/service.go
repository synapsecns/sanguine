package price

import (
	"context"
	"encoding/json"
	"fmt"
	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/jpillora/backoff"
	"net/http"
	"strconv"

	"time"
)

// IPriceFetcher provides price data about tokens using either a cache or defillama
// cache keys sare always ${coin gecko id}_${timestamp}.
type IPriceFetcher interface {
	// GetPriceData attempts to get price data from the cache otherwise it requests defillama
	GetPriceData(context.Context, int, string) *float64
}

const cacheSize = 5000

type priceFetcherImpl struct {
	// tokenCache is the cache of the token prices
	tokenPriceCache *lru.TwoQueueCache[string, float64]
}

// NewPriceFetcher creates a new token price data service.
func NewPriceFetcher() (IPriceFetcher, error) {
	cache, err := lru.New2Q[string, float64](cacheSize)
	if err != nil {
		return nil, fmt.Errorf("could not create token price data cache: %w", err)
	}

	return &priceFetcherImpl{
		tokenPriceCache: cache,
	}, nil
}

func (t *priceFetcherImpl) GetPriceData(parentCtx context.Context, timestamp int, coinGeckoID string) *float64 {
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Minute)
	defer cancel()
	key := fmt.Sprintf("%s_%d", coinGeckoID, timestamp)

	// truncate key to save requests
	key = key[:len(key)-3]

	if data, ok := t.tokenPriceCache.Get(key); ok {
		return &data
	}
	tokenPrice := getDefiLlamaData(ctx, timestamp, coinGeckoID)
	if tokenPrice == nil {
		return nil
	}
	t.tokenPriceCache.Add(key, *tokenPrice)

	return tokenPrice
}

// tokenMetadataMaxRetry is the maximum number of times to retry requesting token metadata
// from the DeFi llama API.
const tokenMetadataMaxRetry = 20
const tokenMetadataMaxRetrySecondary = 1

// GetDefiLlamaData does a get request to defi llama for the symbol and price for a token.
//
//nolint:cyclop,gocognit
func getDefiLlamaData(ctx context.Context, timestamp int, coinGeckoID string) *float64 {
	zero := float64(0)

	if coinGeckoID == "NO_TOKEN" || coinGeckoID == "NO_PRICE" {
		// if there is no data on the token, the amount returned will be 1:1 (price will be same as the amount of token
		// and the token  symbol will say "no symbol"
		return &zero
	}
	client := http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			ResponseHeaderTimeout: 10 * time.Second,
		},
	}
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    10 * time.Millisecond,
		Max:    1 * time.Second,
	}
	timeout := time.Duration(0)
	retries := 0
	retriesSecondary := 0
RETRY:
	select {
	case <-ctx.Done():
		logger.Errorf("context canceled %s, %v", coinGeckoID, ctx.Err())

		return nil
	case <-time.After(timeout):
		if retriesSecondary > tokenMetadataMaxRetrySecondary {
			retriesSecondary = 0
			retries++
		}
		if retries >= tokenMetadataMaxRetry {
			logger.Errorf("Max retries reached, could not get token metadata for %s", coinGeckoID)
			return nil
		}
		var granularityStr string
		if retries > 1 {
			granularityStr = fmt.Sprintf("?searchWidth=%dh", 15*(retries-2)+24)
		}

		url := fmt.Sprintf("https://coins.llama.fi/prices/historical/%d/coingecko:%s%s", timestamp, coinGeckoID, granularityStr)
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			timeout = b.Duration()
			logger.Errorf("error getting response from defi llama %v", err)
			retries++
			goto RETRY
		}
		resRaw, err := client.Do(req)

		if err != nil {
			timeout = b.Duration()
			logger.Errorf("error getting response from defi llama %v", err)
			retries++
			goto RETRY
		}

		res := make(map[string]map[string]map[string]interface{})
		err = json.NewDecoder(resRaw.Body).Decode(&res)
		if err != nil {
			if retriesSecondary > retriesSecondary-2 && retries > 18 {
				logger.Errorf("Failed decoding defillama data after retries %d, retries secondary: %d  id %s: timestamp: %d error %v %s", retries, retriesSecondary, coinGeckoID, timestamp, err, url)
			}
			timeout = b.Duration()
			retriesSecondary++
			goto RETRY
		}

		priceRaw := res["coins"][fmt.Sprintf("coingecko:%s", coinGeckoID)]["price"]
		if priceRaw == nil {
			if retries >= 6 {
				logger.Errorf("error getting price from defi llama: retries: %d %s %d %v %s", retries, coinGeckoID, timestamp, res, url)
			}
			timeout = b.Duration()
			retries++

			goto RETRY
		}
		priceStr := fmt.Sprintf("%.10f", priceRaw)
		priceFloat, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			timeout = b.Duration()
			logger.Errorf("error unwrapping price from defi llama %v", err)
			retries++
			goto RETRY
		}

		price := &priceFloat

		if resRaw.Body.Close() != nil {
			logger.Error("Error closing http connection.")
		}
		return price
	}
}
