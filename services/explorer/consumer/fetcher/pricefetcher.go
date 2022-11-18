package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jpillora/backoff"
	"log"
	"net/http"
	"time"
)

// tokenMetadataMaxRetry is the maximum number of times to retry requesting token metadata
// from the defi llama API.
var tokenMetadataMaxRetry = 10

// GetDefiLlamaData does a get request to defi llama for the symbol and price for a token.
//
//nolint:cyclop
func GetDefiLlamaData(ctx context.Context, timestamp int, coinGeckoID string) (*float64, *string) {
	if coinGeckoID == "NO_TOKEN" || coinGeckoID == "NO_PRICE" {
		// if there is no data on the token, the amount returned will be 1:1 (price will be same as the amount of token
		// and the token  symbol will say "no symbol"
		zero := float64(0)
		noSymbol := "NO_SYMBOL"

		return &zero, &noSymbol
	}
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    30 * time.Second,
	}
	timeout := time.Duration(0)
	retries := 0
RETRY:
	select {
	case <-ctx.Done():

		return nil, nil
	case <-time.After(timeout):
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://coins.llama.fi/prices/historical/%d/coingecko:%s", timestamp, coinGeckoID), nil)
		if err != nil {
			if retries >= tokenMetadataMaxRetry {
				logger.Warnf("Max retries reached, could not get token metadata for %s: %v", coinGeckoID, err)
				return nil, nil
			}
			timeout = b.Duration()
			logger.Errorf("error getting response from defi llama %v", err)
			retries++
			goto RETRY
		}
		resRaw, err := client.Do(req)

		if err != nil {
			if retries >= tokenMetadataMaxRetry {
				logger.Warnf("Max retries reached, could not get token metadata for %s: %v", coinGeckoID, err)
				return nil, nil
			}

			timeout = b.Duration()
			logger.Errorf("error getting response from defi llama %v", err)
			retries++
			goto RETRY
		}

		res := make(map[string]map[string]map[string]interface{})
		err = json.NewDecoder(resRaw.Body).Decode(&res)

		if err != nil {
			logger.Warnf("failed decoding defillama data %s: %v", coinGeckoID, err)

			return nil, nil
		}

		var price *float64
		var symbol *string
		if priceRes, ok := res["coins"][fmt.Sprintf("coingecko:%s", coinGeckoID)]["price"].(float64); ok {
			price = &priceRes
		}

		if stringRes, ok := res["coins"][fmt.Sprintf("coingecko:%s", coinGeckoID)]["symbol"].(string); ok {
			symbol = &stringRes
		}

		if resRaw.Body.Close() != nil {
			log.Printf("Error closing http connection.")
		}

		if price == nil || symbol == nil {
			if retries >= 1 {
				logger.Errorf("error getting price or symbol from defi llama, skipping: retries: %d %s %d", retries, coinGeckoID, timestamp)
				zero := float64(0)
				return &zero, nil
			}
			timeout = b.Duration()
			logger.Errorf("error getting price or symbol from defi llama: retries: %d %s %d", retries, coinGeckoID, timestamp)
			retries++

			goto RETRY
		}

		return price, symbol
	}
}
