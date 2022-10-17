package config

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config is used to configure the explorer's data consumption.
type Config struct {
	// Chains stores the chain configurations.
	Chains ChainConfigs `yaml:"chains"`
	// RefreshRate is the rate at which the explorer will refresh the last block height in seconds.
	RefreshRate uint `yaml:"refresh_rate"`
	// ScribeURL is the URL of the Scribe server.
	ScribeURL string `yaml:"scribe_url"`
	// BridgeConfigAddress is the address of BridgeConfig contract.
	BridgeConfigAddress string `yaml:"bridge_config_abi"`
	// BridgeConfigChainID is the ChainID of BridgeConfig contract.
	BridgeConfigChainID uint32 `yaml:"bridge_config_chain_id"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Chains.IsValid(ctx); !ok {
		return false, err
	}
	return true, nil
}

// Encode gets the encoded config.toml file.
func (c Config) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}

// DecodeConfig parses in a config from a file.
func DecodeConfig(filePath string) (cfg Config, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}
