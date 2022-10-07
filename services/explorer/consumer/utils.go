package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"math"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// OpenYaml opens yaml file with coin gecko ID mapping and returns it.
func OpenYaml() (map[string]string, error) {
	pwd, _ := os.Getwd()
	// nolint:gosec
	input, err := os.ReadFile(pwd + filepath.Clean("/tokenIDToCoinGeckoID.yaml"))
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
	coinGeckoIDs, err := OpenYaml()
	if err != nil {
		fmt.Println("Error while retrieving CoinGecko ids from yaml:", err)
		return nil, nil
	}
	coinGeckoID := coinGeckoIDs[*tokenID]
	return GetDefiLlamaData(ctx, timestamp, &coinGeckoID)
}

// TODO implement this for swaps
// func GetTokenMetadataWithSymbol(blockNumber uint64, tokenID *string) (*float64, *string)

// GetDefiLlamaData does a get request to defi llama for the symbol and price for a token.
func GetDefiLlamaData(ctx context.Context, timestamp int, coinGeckoID *string) (*float64, *string) {
	client := http.Client{
		Timeout: 2 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://coins.llama.fi/prices/historical/%d/coingecko:%s", timestamp, *coinGeckoID), nil) // OK
	if err != nil {
		return nil, nil
	}
	resRaw, err := client.Do(req)
	if err != nil {
		return nil, nil
	}
	res := make(map[string]map[string]map[string]interface{})
	err = json.NewDecoder(resRaw.Body).Decode(&res)
	if err != nil {
		return nil, nil
	}

	var price *float64
	var symbol *string
	price, symbol = nil, nil
	if priceRes, ok := res["coins"][fmt.Sprintf("coingecko:%s", *coinGeckoID)]["price"].(float64); ok {
		price = &priceRes
	}
	if stringRes, ok := res["coins"][fmt.Sprintf("coingecko:%s", *coinGeckoID)]["symbol"].(string); ok {
		symbol = &stringRes
	}
	if resRaw.Body.Close() != nil {
		fmt.Println("Failed while closing connection")
	}
	return price, symbol
}

// GetAmountUSD computes the USD value of a token amount.
func GetAmountUSD(amount *big.Int, decimals uint8, price *float64) *float64 {
	trueAmount := float64(amount.Int64()) * math.Pow(10.0, float64(decimals)) * *price
	return &trueAmount
}
