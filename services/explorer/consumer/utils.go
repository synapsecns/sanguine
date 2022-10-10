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
	"strings"
	"time"
)

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
		fmt.Println("Error while retrieving CoinGecko ids from yaml:", err)
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
		fmt.Println("Error while retrieving CoinGecko ids from yaml:", err)
		return nil, nil
	}
	coinGeckoID := coinGeckoIDs[strings.ToLower(*tokenSymbol)]
	return GetDefiLlamaData(ctx, timestamp, &coinGeckoID)
}

// GetDefiLlamaData does a get request to defi llama for the symbol and price for a token.
func GetDefiLlamaData(ctx context.Context, timestamp int, coinGeckoID *string) (*float64, *string) {
	if *coinGeckoID == "NO_TOKEN" || *coinGeckoID == "NO_PRICE" {
		// if there is no data on the token, the amount returned will be 1:1 (price will be same as the amount of token
		// and the token  symbol will say "no symbol"
		one := float64(1)
		noSymbol := "NO_SYMBOL"
		return &one, &noSymbol
	}

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
