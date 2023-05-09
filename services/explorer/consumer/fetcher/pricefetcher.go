package fetcher

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/jpillora/backoff"
)

// PriceService --output=mocks --case=underscore.
type PriceService interface {
	// GetDefiLlamaData historical price data
	GetDefiLlamaData(context.Context, int, string) *float64
}

// tokenMetadataMaxRetry is the maximum number of times to retry requesting token metadata
// from the defi llama API.
const tokenMetadataMaxRetry = 20
const tokenMetadataMaxRetrySecondary = 1

var endpoints = []string{
	"https://o7zsk6wjki.execute-api.us-east-1.amazonaws.com/default/ipDivider0?url=",
	"https://56tczseny8.execute-api.us-east-1.amazonaws.com/default/ipDivider1?url=",
	"https://7ormerg7r2.execute-api.us-east-1.amazonaws.com/default/ipDivider2?url=",
	"https://kam552ykx0.execute-api.us-east-1.amazonaws.com/default/ipDivider3?url=",
	"https://tliz2ftmib.execute-api.us-east-1.amazonaws.com/default/ipDivider4?url=",
	"https://tliz2ftmib.execute-api.us-east-1.amazonaws.com/default/ipDivider4?url=",
	"https://kmhpuaswf1.execute-api.us-east-1.amazonaws.com/default/ipDivider5?url=",
	"https://xyunzulw3b.execute-api.us-east-1.amazonaws.com/default/ipDivider6?url=",
	"https://7u5ixxsvrb.execute-api.us-east-1.amazonaws.com/default/ipDivider7?url=",
	"https://pgsrika6v6.execute-api.us-east-1.amazonaws.com/default/ipDivider8?url=",
	"https://vy6svlr3s6.execute-api.us-east-1.amazonaws.com/default/ipDivider9?url=",
	"https://rkdoixpf1a.execute-api.us-east-1.amazonaws.com/default/ipDivider10?url=",
	"https://uha0wvok1l.execute-api.us-east-1.amazonaws.com/default/ipDivider11?url=",
	"https://cvp8yvofjj.execute-api.us-east-1.amazonaws.com/default/ipDivider12?url=",
	"https://5gisfyd0hj.execute-api.us-east-1.amazonaws.com/default/ipDivider13?url=",
	"https://62wcvmki11.execute-api.us-east-1.amazonaws.com/default/ipDivider14?url=",
	"https://z0tes6abp9.execute-api.us-east-1.amazonaws.com/default/ipDivider15?url=",
}

//
// const coinGeckoRetryThreshold = 15
// func GetCoinGeckoPriceData(ctx context.Context, timestamp int, coinGeckoID string, retries int) *float64 {
//	client := http.Client{
//		Timeout: 5 * time.Second,
//	}
//	granularity := timestamp + (1380 + (retries-coinGeckoRetryThreshold)*100)
//	reqUrl := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/%s/market_chart/range?vs_currency=usd&from=%d&to=%d", coinGeckoID, timestamp, granularity)
//	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
//	if err != nil {
//		logger.Error("CoinGecko price request failed", err)
//	}
//	resRaw, err := client.Do(req)
//	res := make(map[string][][]float64)
//	err = json.NewDecoder(resRaw.Body).Decode(&res)
//	if len(res["prices"]) > 0 && len(res["prices"][0]) > 0 {
//		priceRaw := res["prices"][0][1]
//		priceStr := fmt.Sprintf("%.4f", priceRaw)
//		priceFloat, err := strconv.ParseFloat(priceStr, 64)
//		if err != nil {
//			return nil
//		}
//		if resRaw.Body.Close() != nil {
//			logger.Error("Error closing http connection.")
//		}
//		return &priceFloat
//	}
//	return nil
//}

// GetDefiLlamaData does a get request to defi llama for the symbol and price for a token.
//
//nolint:cyclop,gocognit
func GetDefiLlamaData(ctx context.Context, timestamp int, coinGeckoID string) *float64 {
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

		// nolint:gosec
		if rand.Intn(retries/3+1) == 1 {
			// nolint:gosec
			url = endpoints[rand.Intn(len(endpoints))] + fmt.Sprintf("'https://coins.llama.fi/prices/historical/%d/coingecko:%s%s'", timestamp, coinGeckoID, granularityStr)
		}
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
