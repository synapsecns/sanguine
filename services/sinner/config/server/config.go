// Package serverconfig is the config loader for the server
package serverconfig

import (
	"fmt"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/services/sinner/config"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config is used to configure the explorer server.
type Config struct {
	// HTTPPort is the http port for the api
	HTTPPort uint16 `yaml:"http_port"`
	// DBPath is the address of the database
	DBPath string `yaml:"db_path"`
	// DBType is the flag signifying the type of database (mysql, sqlite, etc).
	DBType string `yaml:"db_type"`
	// SkipMigrations skips the database migrations.
	SkipMigrations bool `yaml:"skip_migrations"`
}

// IsValid makes sure the config is valid.
func (c *Config) IsValid() error {
	if c.DBPath == "" {
		return fmt.Errorf("db_address, %w", config.ErrRequiredGlobalField)
	}
	if c.HTTPPort == 0 {
		return fmt.Errorf("http_port cannot be zero, %w", config.ErrRequiredGlobalField)
	}
	return nil
}

// DecodeServerConfig parses in a config from a file.
func DecodeServerConfig(filePath string) (cfg Config, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}

	err = cfg.IsValid()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
