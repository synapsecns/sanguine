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

// Config is used to configure an Executor agent.
type Config struct {
	// Chains stores all chain information
	Chains ChainConfigs `yaml:"chains"`
	// AttestationCollectorChainID is the chain ID of the attestation collector.
	AttestationCollectorChainID uint32 `yaml:"attestation_collector_chain_id"`
	// AttestationCollectorAddress is the address of the Attestation Collector on SYN Chain.
	AttestationCollectorAddress string `yaml:"attestation_collector_address"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Chains.IsValid(ctx); !ok {
		return false, err
	}

	if c.AttestationCollectorChainID == 0 {
		return false, fmt.Errorf("attestation collector chain ID is not set")
	}

	if c.AttestationCollectorAddress == "" {
		return false, fmt.Errorf("attestation collector address is not set")
	}

	return true, nil
}

// Encode gets the encoded config.yaml file.
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
