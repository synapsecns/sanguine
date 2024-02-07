package unifiedconfig

import (
	"fmt"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"github.com/synapsecns/sanguine/services/sinner/config/indexer"
	serverConfig "github.com/synapsecns/sanguine/services/sinner/config/server"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// UnifiedConfig combines both the server and indexer config into a single type.
type UnifiedConfig struct {
	// DBPath is the address of the database
	DBPath string `yaml:"db_path"`
	// DBType is the flag signifying the type of database (mysql, sqlite, etc).
	DBType string `yaml:"db_type"`
	// SkipMigrations skips the database migrations.
	SkipMigrations bool `yaml:"skip_migrations"`
	// HTTPPort is the http port for the api.
	HTTPPort uint16 `yaml:"http_port"`
	// ScribeURL is the URL of the Scribe graphql server (indexer).
	ScribeURL string `yaml:"scribe_url"`
	// DefaultRefreshRate is the default rate at which data is refreshed in seconds (indexer).
	DefaultRefreshRate int `yaml:"default_refresh_rate"`
	// Chains stores the chain configurations (indexer).
	Chains []indexerconfig.ChainConfig `yaml:"chains"`
}

// IsValid makes sure the config is valid.
func (c *UnifiedConfig) IsValid() error {
	if c.ScribeURL == "" {
		return fmt.Errorf("scribe_url, %w", config.ErrRequiredGlobalField)
	}
	if len(c.Chains) == 0 {
		return fmt.Errorf("no chains specified for indexing")
	}
	for _, chain := range c.Chains {
		err := chain.IsValid()
		if err != nil {
			return fmt.Errorf("chain with ID %d is invalid: %w", chain.ChainID, err)
		}
	}

	return nil
}

// ServerConfig converts the config to a server config for the api.
func (c *UnifiedConfig) ServerConfig() serverConfig.Config {
	return serverConfig.Config{
		HTTPPort:       c.HTTPPort,
		DBPath:         c.DBPath,
		DBType:         c.DBType,
		SkipMigrations: c.SkipMigrations,
	}
}

// IndexerConfig converts the config to a indexer config for the indexer.
func (c *UnifiedConfig) IndexerConfig() indexerconfig.Config {
	return indexerconfig.Config{
		DefaultRefreshRate: c.DefaultRefreshRate,
		ScribeURL:          c.ScribeURL,
		Chains:             c.Chains,
		DBPath:             c.DBPath,
		DBType:             c.DBType,
		SkipMigrations:     c.SkipMigrations,
	}
}

// DecodeConfig parses in a config from a file.
func DecodeConfig(filePath string) (cfg UnifiedConfig, err error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return UnifiedConfig{}, fmt.Errorf("config file does not exist: %w", err)
	}
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return UnifiedConfig{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return UnifiedConfig{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}

	err = cfg.IsValid()
	if err != nil {
		return UnifiedConfig{}, fmt.Errorf("malformed unified config %w", err)
	}

	return cfg, nil
}
