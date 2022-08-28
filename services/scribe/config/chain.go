package config

import (
	"context"
	"fmt"

	"github.com/richardwilkes/toolbox/collection"
)

// ChainConfig defines the config for a specific chain.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `toml:"ChainID"`
	// RPCUrl is the URL of the chain's RPC server.
	RPCUrl string `toml:"RPCUrl"`
	// ConfirmationThreshold is the number of blocks to wait before indexing events.
	ConfirmationThreshold uint32 `toml:"ConfirmationThreshold"`
	// Contracts stores all the contract information for the chain.
	Contracts ContractConfigs `toml:"Contracts"`
}

// ChainConfigs contains an array of ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
// It also calls IsValid on each individual ContractConfig.
func (c ChainConfigs) IsValid(ctx context.Context) (ok bool, err error) {
	intSet := collection.NewUint32Set()

	for _, cfg := range c {
		if intSet.Contains(cfg.ChainID) {
			return false, fmt.Errorf("chain id %d appears twice: %w", cfg.ChainID, ErrDuplicateChainID)
		}

		ok, err = cfg.IsValid(ctx)
		if !ok {
			return false, err
		}

		intSet.Add(cfg.ChainID)
	}

	return true, nil
}

// IsValid validates the chain config.
func (c ChainConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.ChainID == 0 {
		return false, fmt.Errorf("%w: chain ID cannot be 0", ErrInvalidChainID)
	}
	if c.RPCUrl == "" {
		return false, fmt.Errorf("field RPCUrl: %w", ErrRequiredField)
	}
	if ok, err = c.Contracts.IsValid(ctx); !ok {
		return false, err
	}

	return true, nil
}
