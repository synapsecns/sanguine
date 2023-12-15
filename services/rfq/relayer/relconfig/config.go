package relconfig

import (
	"fmt"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"
	"os"

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
}

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
