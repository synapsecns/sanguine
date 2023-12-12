// Package config implements the configuration parser for the RFQ Quoter.
package config

import (
	"fmt"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config is the configuration for the RFQ Quoter.
type Config struct {
	// AuthExpiryDelta is the delta in seconds to add to the current time to get the expiry time for the auth token.
	// TODO: consider only exporting via getter to prevent accidental mutation
	AuthExpiryDelta int64
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
