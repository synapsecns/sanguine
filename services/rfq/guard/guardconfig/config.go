// Package guardconfig contains the config yaml object for the relayer.
package guardconfig

import (
	"fmt"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"gopkg.in/yaml.v2"

	"path/filepath"
)

// Config represents the configuration for the relayer.
type Config struct {
	// Chains is a map of chainID -> chain config.
	Chains map[int]ChainConfig `yaml:"chains"`
	// OmniRPCURL is the URL of the OmniRPC server.
	OmniRPCURL string `yaml:"omnirpc_url"`
	// Database is the database config.
	Database DatabaseConfig `yaml:"database"`
	// Signer is the signer config.
	Signer config.SignerConfig `yaml:"signer"`
	// SubmitterConfig is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
	// DBSelectorInterval is the interval for the db selector.
	DBSelectorInterval time.Duration `yaml:"db_selector_interval"`
}

// ChainConfig represents the configuration for a chain.
type ChainConfig struct {
	// Bridge is the rfq bridge contract address.
	RFQAddress string `yaml:"rfq_address"`
	// Confirmations is the number of required confirmations.
	Confirmations uint64 `yaml:"confirmations"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
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
	err = config.Validate()
	if err != nil {
		return config, fmt.Errorf("error validating config: %w", err)
	}
	return config, nil
}

// Validate validates the config.
func (c Config) Validate() (err error) {
	for chainID := range c.Chains {
		addr, err := c.GetRFQAddress(chainID)
		if err != nil {
			return fmt.Errorf("could not get rfq address: %w", err)
		}
		if !common.IsHexAddress(addr) {
			return fmt.Errorf("invalid rfq address: %s", addr)
		}
	}

	return nil
}

// GetChains returns the chains config.
func (c Config) GetChains() map[int]ChainConfig {
	return c.Chains
}

// GetRFQAddress returns the RFQ address for the given chain.
func (c Config) GetRFQAddress(chainID int) (string, error) {
	chainCfg, ok := c.Chains[chainID]
	if !ok {
		return "", fmt.Errorf("chain config not found for chain %d", chainID)
	}
	return chainCfg.RFQAddress, nil
}

const defaultDBSelectorIntervalSeconds = 1

// GetDBSelectorInterval returns the interval for the DB selector.
func (c Config) GetDBSelectorInterval() time.Duration {
	interval := c.DBSelectorInterval
	if interval <= 0 {
		interval = time.Duration(defaultDBSelectorIntervalSeconds) * time.Second
	}
	return interval
}

// NewGuardConfigFromRelayer creates a new guard config from a relayer config.
func NewGuardConfigFromRelayer(relayerCfg relconfig.Config) Config {
	cfg := Config{
		Chains:             make(map[int]ChainConfig),
		OmniRPCURL:         relayerCfg.OmniRPCURL,
		Database:           DatabaseConfig(relayerCfg.Database),
		Signer:             relayerCfg.Signer,
		SubmitterConfig:    relayerCfg.SubmitterConfig,
		DBSelectorInterval: relayerCfg.DBSelectorInterval,
	}
	for chainID, chainCfg := range relayerCfg.GetChains() {
		cfg.Chains[chainID] = ChainConfig{
			RFQAddress:    chainCfg.RFQAddress,
			Confirmations: chainCfg.Confirmations,
		}
	}
	return cfg
}
