// Package relconfig contains the config yaml object for the relayer.
package relconfig

import (
	"fmt"
	"os"
	"strings"

	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"

	"path/filepath"
)

type IConfig interface {
	GetChains() map[int]ChainConfig
	GetOmniRPCURL() string
	GetRfqAPIURL() string
	GetDatabase() DatabaseConfig
	GetQuotableTokens(token string) ([]string, error)
	GetSigner() config.SignerConfig
	GetFeePricer() FeePricerConfig
	GetTokenID(chain int, addr string) (string, error)
	GetNativeToken(chainID uint32) (string, error)
	GetTokenDecimals(chainID uint32, token string) (uint8, error)
	GetTokens(chainID uint32) (map[string]TokenConfig, error)
	GetTokenName(chain uint32, addr string) (string, error)
}

// Config represents the configuration for the relayer.
// TODO: validation function.
type Config struct {
	Chains          map[int]ChainConfig    `yaml:"bridges"`
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

func (c Config) GetChains() map[int]ChainConfig {
	return c.Chains
}

func (c Config) GetOmniRPCURL() string {
	return c.OmniRPCURL
}

func (c Config) GetRfqAPIURL() string {
	return c.RfqAPIURL
}

func (c Config) GetDatabase() DatabaseConfig {
	return c.Database
}

func (c Config) GetSigner() config.SignerConfig {
	return c.Signer
}

func (c Config) GetFeePricer() FeePricerConfig {
	return c.FeePricer
}

func (c Config) GetTokenID(chain int, addr string) (string, error) {
	chainConfig, ok := c.Chains[int(chain)]
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

func (c Config) GetQuotableTokens(token string) ([]string, error) {
	tokens, ok := c.QuotableTokens[token]
	if !ok {
		return nil, fmt.Errorf("no quotable tokens for token %s", token)
	}
	return tokens, nil
}

func (c Config) GetNativeToken(chainID uint32) (string, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return "", fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	if len(chainConfig.NativeToken) == 0 {
		return "", fmt.Errorf("chain config for chainID %d does not have a native token", chainID)
	}
	return chainConfig.NativeToken, nil
}

func (c Config) GetTokenDecimals(chainID uint32, token string) (uint8, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return 0, fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	for tokenName, tokenConfig := range chainConfig.Tokens {
		if token == tokenName {
			return tokenConfig.Decimals, nil
		}
	}
	return 0, fmt.Errorf("could not get token decimals for chain %d and token %s", chainID, token)
}

func (c Config) GetTokens(chainID uint32) (map[string]TokenConfig, error) {
	chainConfig, ok := c.Chains[int(chainID)]
	if !ok {
		return nil, fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	return chainConfig.Tokens, nil
}

func (c Config) GetTokenName(chain uint32, addr string) (string, error) {
	chainConfig, ok := c.Chains[int(chain)]
	if !ok {
		return "", fmt.Errorf("no chain config for chain %d", chain)
	}
	for tokenName, tokenConfig := range chainConfig.Tokens {
		// TODO: probably a better way to do this.
		if strings.ToLower(tokenConfig.Address) == strings.ToLower(addr) {
			return tokenName, nil
		}
	}
	return "", fmt.Errorf("no tokenName found for chain %d and address %s", chain, addr)
}
