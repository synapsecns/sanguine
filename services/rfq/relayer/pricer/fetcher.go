package pricer

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// CoingeckoPriceFetcher is an interface for fetching prices from coingecko.
//
//go:generate go run github.com/vektra/mockery/v2 --name CoingeckoPriceFetcher --output ./mocks --case=underscore
type CoingeckoPriceFetcher interface {
	GetPrice(ctx context.Context, token string) (float64, error)
}

// CoingeckoPriceFetcherImpl is an implementation of CoingeckoPriceFetcher.
type CoingeckoPriceFetcherImpl struct {
	client *http.Client
}

// NewCoingeckoPriceFetcher creates a new instance of CoingeckoPriceFetcherImpl.
func NewCoingeckoPriceFetcher(timeoutMs int) *CoingeckoPriceFetcherImpl {
	return &CoingeckoPriceFetcherImpl{
		client: &http.Client{
			Timeout: time.Duration(timeoutMs) * time.Millisecond,
		},
	}
}

var coingeckoIDLookup = map[string]string{
	"ETH": "ethereum",
}

// GetPrice fetches the price of a token from coingecko.
func (c *CoingeckoPriceFetcherImpl) GetPrice(ctx context.Context, token string) (price float64, err error) {
	coingeckoID, ok := coingeckoIDLookup[token]
	if !ok {
		return price, fmt.Errorf("could not get coingecko id for token: %s", token)
	}
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=USD", coingeckoID)

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
		return price, fmt.Errorf("bad status code fetching price from coingecko: %v", r.Status)
	}
	defer func() {
		err = r.Body.Close()
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
	return price, nil
}
