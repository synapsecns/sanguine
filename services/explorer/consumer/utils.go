package consumer

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"math"
	"math/big"
	"net/http"
	"os"
	"path/filepath"
)

// OpenYaml opens yaml file with coin gecko ID mapping and returns it.
func OpenYaml(path string) (map[string]string, error) {
	input, err := os.ReadFile(filepath.Clean(path))
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
func GetTokenMetadataWithTokenID(timestamp int, tokenID *string, path string) (*float64, *string) {
	coinGeckoIDs, err := OpenYaml(path)
	if err != nil {
		fmt.Println("Error while retrieving CoinGecko ids from yaml:", err)
		return nil, nil
	}
	coinGeckoID := coinGeckoIDs[*tokenID]
	return GetDefiLlamaData(timestamp, &coinGeckoID)
}

// TODO implement this for swaps
// func GetTokenMetadataWithSymbol(blockNumber uint64, tokenID *string) (*float64, *string)

// GetDefiLlamaData does a get request to defi llama for the symbol and price for a token.
func GetDefiLlamaData(timestamp int, coinGeckoID *string) (*float64, *string) {
	fmt.Println(timestamp, *coinGeckoID)
	resp, err := http.Get(fmt.Sprintf("https://coins.llama.fi/prices/historical/%d/coingecko:%s", timestamp, *coinGeckoID))
	if err != nil {
		return nil, nil
	}
	res := make(map[string]map[string]map[string]interface{})
	if json.NewDecoder(resp.Body).Decode(&res) != nil {
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
	if resp.Body.Close() != nil {
		fmt.Println("Failed while closing connection")
	}
	return price, symbol
}

// GetAmountUSD computes the USD value of a token amount.
func GetAmountUSD(amount *big.Int, decimals uint8, price *float64) *float64 {
	fmt.Println(amount, decimals, price)
	trueAmount := float64(amount.Int64()) * math.Pow(10.0, float64(decimals)) * *price
	return &trueAmount
}
