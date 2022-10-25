package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jpillora/backoff"
	"gopkg.in/yaml.v2"
	"log"
	"math"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// tokenMetadataMaxRetry is the maximum number of times to retry requesting token metadata
// from the defi llama API.
var tokenMetadataMaxRetry = 10

// OpenYaml opens yaml file with coin gecko ID mapping and returns it.
func OpenYaml(path string) (map[string]string, error) {
	// nolint:gosec
	input, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error opening yaml file %w", err)
	}
	var res map[string]string
	err = yaml.Unmarshal(input, &res)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling yaml file %w", err)
	}
	return res, nil
}

// GetTokenMetadataWithTokenID gets the token metadata (symbol, price).
func GetTokenMetadataWithTokenID(ctx context.Context, timestamp int, tokenID *string) (*float64, *string) {
	pwd, _ := os.Getwd()
	path := pwd + filepath.Clean("/tokenIDToCoinGeckoID.yaml")
	coinGeckoIDs, err := OpenYaml(path)
	if err != nil {
		return nil, nil
	}
	coinGeckoID := coinGeckoIDs[*tokenID]
	return GetDefiLlamaData(ctx, timestamp, &coinGeckoID)
}

// GetTokenMetadataWithTokenSymbol gets the token metadata (symbol, price).
func GetTokenMetadataWithTokenSymbol(ctx context.Context, timestamp int, tokenSymbol *string) (*float64, *string) {
	pwd, _ := os.Getwd()
	path := pwd + filepath.Clean("/tokenIDToCoinGeckoID.yaml")
	coinGeckoIDs, err := OpenYaml(path)
	if err != nil {
		return nil, nil
	}
	coinGeckoID := coinGeckoIDs[strings.ToLower(*tokenSymbol)]
	return GetDefiLlamaData(ctx, timestamp, &coinGeckoID)
}

// GetDefiLlamaData does a get request to defi llama for the symbol and price for a token.
//
//nolint:cyclop
func GetDefiLlamaData(ctx context.Context, timestamp int, coinGeckoID *string) (*float64, *string) {
	if *coinGeckoID == "NO_TOKEN" || *coinGeckoID == "NO_PRICE" {
		// if there is no data on the token, the amount returned will be 1:1 (price will be same as the amount of token
		// and the token  symbol will say "no symbol"
		one := float64(1)
		noSymbol := "NO_SYMBOL"
		return &one, &noSymbol
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	// backoff in the case of an error
	b := &backoff.Backoff{
		Factor: 2,
		Jitter: true,
		Min:    1 * time.Second,
		Max:    30 * time.Second,
	}
	// timeout should always be 0 on the first attempt
	timeout := time.Duration(0)
	// keep track of the number of retries
	retries := 0
RETRY:
	select {
	case <-ctx.Done():
		return nil, nil
	case <-time.After(timeout):
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://coins.llama.fi/prices/historical/%d/coingecko:%s", timestamp, *coinGeckoID), nil)
		if err != nil {
			if retries >= tokenMetadataMaxRetry {
				return nil, nil
			}
			timeout = b.Duration()
			logger.Errorf("error creating request to defi llama %v", err)
			retries++
			goto RETRY
		}
		resRaw, err := client.Do(req)
		if err != nil {
			if retries >= tokenMetadataMaxRetry {
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
			return nil, nil
		}

		var price *float64
		var symbol *string
		if priceRes, ok := res["coins"][fmt.Sprintf("coingecko:%s", *coinGeckoID)]["price"].(float64); ok {
			price = &priceRes
		}
		if stringRes, ok := res["coins"][fmt.Sprintf("coingecko:%s", *coinGeckoID)]["symbol"].(string); ok {
			symbol = &stringRes
		}
		if resRaw.Body.Close() != nil {
			log.Printf("Error closing http connection.")
		}
		return price, symbol
	}
}

// GetAmountUSD computes the USD value of a token amount.
func GetAmountUSD(amount *big.Int, decimals uint8, price *float64) *float64 {
	trueAmount := float64(amount.Int64()) * math.Pow(10.0, float64(decimals)) * *price
	return &trueAmount
}
