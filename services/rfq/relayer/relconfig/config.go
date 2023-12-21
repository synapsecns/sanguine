// Package relconfig contains the config yaml object for the relayer.
package relconfig

import (
	"fmt"
	"os"

	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"

	"path/filepath"
)

// Config represents the configuration for the relayer.
// TODO: validation function.
type Config struct {
	// ChainID: address
	// TODO(aurelius): move under ChainConfig
	// TODO: this can actually be replaced by quotable tokens.
	Tokens map[int][]string `yaml:"tokens"`
	// ChainID: bridge
	Bridges         map[int]ChainConfig    `yaml:"bridges"`
	OmniRPCURL      string                 `yaml:"omnirpc_url"`
	RfqAPIURL       string                 `yaml:"rfq_url"`
	Database        DatabaseConfig         `yaml:"database"`
	QuotableTokens  map[string][]string    `yaml:"quotable_tokens"`
	Signer          config.SignerConfig    `yaml:"signer"`
	SubmitterConfig submitterConfig.Config `yaml:"submitter"`
	FeePricer       FeePricerConfig        `yaml:"fee_pricer"`
}

func (c Config) GetTokenID(chain int, addr string) (string, error) {
	chainConfig, ok := c.Bridges[int(chain)]
	if !ok {
		return "", fmt.Errorf("no chain config for chain %d", chain)
	}
	for tokenID, tokenConfig := range chainConfig.Tokens {
		if tokenConfig.Address == addr {
			return tokenID, nil
		}
	}
	return "", fmt.Errorf("no tokenID found for chain %d and address %s", chain, addr)
}

// ChainConfig represents the configuration for a chain.
type ChainConfig struct {
	// Bridge is the bridge confirmation count.
	Bridge string `yaml:"address"`
	// Confirmations is the number of required confirmations
	Confirmations uint64 `yaml:"confirmations"`
	// Tokens is a map of token ID -> token config.
	Tokens map[string]TokenConfig `yaml:"tokens"`
	// NativeToken is the native token of the chain (pays gas).
	NativeToken string `yaml:"native_token"`
}

// TokenConfig represents the configuration for a token.
type TokenConfig struct {
	// Address is the token address.
	Address string `yaml:"address"`
	// Decimals is the token decimals.
	Decimals uint8 `yaml:"decimals"`
	// For now, specify the USD price of the token in the config.
	PriceUSD float64 `yaml:"price_usd"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}

// FeePricerConfig represents the configuration for the fee pricer.
type FeePricerConfig struct {
	// OriginGasEstimate is the gas required to execute prove + claim transactions on origin chain.
	OriginGasEstimate int `yaml:"origin_gas_estimate"`
	// DestinationGasEstimate is the gas required to execute relay transaction on destination chain.
	DestinationGasEstimate int `yaml:"destination_gas_estimate"`
	// GasPriceCacheTTLSeconds is the TTL for the gas price cache.
	GasPriceCacheTTLSeconds int `yaml:"gas_price_cache_ttl"`
	// TokenPriceCacheTTLSeconds is the TTL for the token price cache.
	TokenPriceCacheTTLSeconds int `yaml:"token_price_cache_ttl"`
}

// LoadConfig loads the config from the given path.
func LoadConfig(path string) (config Config, err error) {
	input, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &config)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return config, nil
}
