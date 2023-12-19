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

// ChainConfig represents the configuration for a chain.
type ChainConfig struct {
	// Bridge is the bridge confirmation count.
	Bridge string `yaml:"address"`
	// Confirmations is the number of required confirmations
	Confirmations uint64 `yaml:"confirmations"`
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
	// GasPriceCacheTTL is the TTL for the gas price cache.
	GasPriceCacheTTL int `yaml:"gas_price_cache_ttl"`
	// TokenPriceCacheTTL is the TTL for the token price cache.
	TokenPriceCacheTTL int `yaml:"token_price_cache_ttl"`
	// Tokens is a map of chain id -> token address -> token name.
	Tokens map[uint32]map[string]string `yaml:"tokens"`
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
