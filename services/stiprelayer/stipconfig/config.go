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

// Config holds the configuration for the STIP relayer service.
type Config struct {
	Signer config.SignerConfig `yaml:"signer"`
	// Submitter is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
	ArbAddress      string                 `yaml:"arb_address"`
	ArbChainID      uint64                 `yaml:"arb_chain_id"`
	StartDate       time.Time              `yaml:"start_date"`
	Database        DatabaseConfig         `yaml:"database"`
	OmniRPCURL      string                 `yaml:"omnirpc_url"`
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
