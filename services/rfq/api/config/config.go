// Package config implements the configuration parser for the RFQ Quoter.
package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
)

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}

// Config is the configuration for the RFQ Quoter.
type Config struct {
	Database   DatabaseConfig `yaml:"database"`
	OmniRPCURL string         `yaml:"omnirpc_url"`
	// bridges is a map of chainid->address
	Bridges map[uint32]string `yaml:"bridges"`
	Port    string            `yaml:"port"`
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
