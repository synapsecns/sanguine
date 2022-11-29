package config

import (
	"context"
	"fmt"
	"github.com/richardwilkes/toolbox/collection"
)

// ChainConfig defines the config for a specific chain.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// OriginAddress is the address of the origin contract.
	OriginAddress string `yaml:"origin_address"`
	// AttestationCollectorAddress is the address of the attestation collector contract.
	AttestationCollectorAddress string `yaml:"attestation_collector_address"`
}

// ChainConfigs contains an array of ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
func (c ChainConfigs) IsValid(ctx context.Context) (ok bool, err error) {
	intSet := collection.Set[uint32]{}

	for _, cfg := range c {
		if intSet.Contains(cfg.ChainID) {
			return false, fmt.Errorf("chain id %d appears twice: %s", cfg.ChainID, "duplicate chain id")
		}
		intSet.Add(cfg.ChainID)
	}

	return true, nil
}

// IsValid validates the chain config.
func (c ChainConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.ChainID == 0 {
		return false, fmt.Errorf("%s: chain ID cannot be 0", "invalid chain id")
	}
	return true, nil
}
