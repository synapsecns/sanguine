// Package stipconfig contains the configuration structures and logic for the STIP relayer service.
package stipconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"
)

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}

// FeeRebate represents the fee and rebate values.
type FeeRebate struct {
	Fee       float64 `yaml:"fee"`        // Fee is the cost that will be charged.
	RebateBps float64 `yaml:"rebate_bps"` // RebateBps is the amount that will be returned, in units of basis points.
}

// TokenFeeRebate is a map where the key is a string representing a token,
// and the value is a FeeRebate struct representing the fee and rebate for that token.
type TokenFeeRebate map[string]FeeRebate

// ModuleFeeRebate is a map where the key is a string representing a module,
// and the value is a TokenFeeRebate map representing the fee and rebate for each token in that module.
type ModuleFeeRebate map[string]TokenFeeRebate

// FeesAndRebates is a map where the key is an integer representing a specific category or group,
// and the value is a ModuleFeeRebate map representing the fee and rebate for each module in that category or group.
type FeesAndRebates map[int]ModuleFeeRebate

// Config holds the configuration for the STIP relayer service.
type Config struct {
	Signer config.SignerConfig `yaml:"signer"`
	// Submitter is the submitter config.
	SubmitterConfig  submitterConfig.Config `yaml:"submitter_config"`
	ArbAddress       string                 `yaml:"arb_address"`
	ArbChainID       uint64                 `yaml:"arb_chain_id"`
	StartDate        time.Time              `yaml:"start_date"`
	Database         DatabaseConfig         `yaml:"database"`
	OmniRPCURL       string                 `yaml:"omnirpc_url"`
	FeesAndRebates   FeesAndRebates         `yaml:"fees_and_rebates"`
	DuneInterval     time.Duration          `yaml:"dune_interval"`
	RebateInterval   time.Duration          `yaml:"rebate_interval"`
	StipAPIPort      string                 `yaml:"stip_api_port"`
	ARBMaxTransfer   int64                  `yaml:"ARB_max_transfer"`
	ArbCapPerAddress int64                  `yaml:"arb_cap_per_address"`
}

const defaultArbCapPerAddress = 2000

// GetArbCapPerAddress returns the configured arb cap per address, in human-readable units.
func (c Config) GetArbCapPerAddress() int64 {
	if c.ArbCapPerAddress == 0 {
		return defaultArbCapPerAddress
	}
	return c.ArbCapPerAddress
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
