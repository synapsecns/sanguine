// parses the input config.yaml file
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

// Config holds the config for the explorer.
type Config struct {
	// SynapseBridgeAddress is the address of the SynapseBridge.sol contract
	SynapseBridgeAddress string `yaml:"synapse_bridge_address"`
	// BridgeConfigV3Address is the address of the BridgeConfigV3.sol contract
	BridgeConfigV3Address string `yaml:"bridge_config_v3_address"`
	// SwapFlashLoanAddress is the address of the SwapFlashLoan.sol contract
	SwapFlashLoanAddress string `yaml:"swap_flash_loan_address"`
}

// TODO: add more checks
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if c.SynapseBridgeAddress != c.BridgeConfigV3Address && c.SynapseBridgeAddress != c.SwapFlashLoanAddress {
		return true, nil
	}
	return false, err
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
