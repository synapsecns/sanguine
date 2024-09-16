// Package config implements the configuration parser for the RFQ Quoter.
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

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
	Bridges         map[uint32]string `yaml:"bridges"`
	Port            string            `yaml:"port"`
	RelayAckTimeout time.Duration     `yaml:"relay_ack_timeout"`
	MaxQuoteAge     time.Duration     `yaml:"max_quote_age"`
	WebsocketPort   *string           `yaml:"websocket_port"`
}

const defaultRelayAckTimeout = 30 * time.Second

// GetRelayAckTimeout returns the relay ack timeout.
func (c Config) GetRelayAckTimeout() time.Duration {
	if c.RelayAckTimeout == 0 {
		return defaultRelayAckTimeout
	}
	return c.RelayAckTimeout
}

const defaultMaxQuoteAge = 1 * time.Hour

// GetMaxQuoteAge returns the max quote age.
func (c Config) GetMaxQuoteAge() time.Duration {
	if c.MaxQuoteAge == 0 {
		return defaultMaxQuoteAge
	}
	return c.MaxQuoteAge
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
