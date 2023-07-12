package executor

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
	// DestinationAddress is the address of the destination contract.
	DestinationAddress string `yaml:"destination_address"`
	// LightInboxAddress is the address of the light inbox contract.
	LightInboxAddress string `yaml:"light_inbox_address"`
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

	if c.OriginAddress == "" {
		return false, fmt.Errorf("field OriginAddress is required")
	}

	if c.DestinationAddress == "" {
		return false, fmt.Errorf("field DestinationAddress is required")
	}

	// TODO (joe): Change from 10 to SYN chain constant, this will be wrong in prod
	if c.ChainID == 10 {
		if c.LightInboxAddress == "" {
			return false, fmt.Errorf("field LightInboxAddress is required")
		}
	}

	return true, nil
}
