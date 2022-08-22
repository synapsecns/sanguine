package config

import (
	"context"
	"fmt"

	"github.com/richardwilkes/toolbox/collection"
)

// ContractConfig defines the config for a specific contract.
type ContractConfig struct {
	// Address is the address of the contract.
	Address string `toml:"Address"`
	// StartBlock is the block number to start indexing events from.
	StartBlock uint64 `toml:"StartBlock"`
}

// ContractConfigs contains a map of name->contract config.
type ContractConfigs map[string]ContractConfig

// IsValid validates the contract configs by asserting no two contracts appear twice.
// It also calls IsValid on each individual ContractConfig.
func (c ContractConfigs) IsValid(ctx context.Context) (ok bool, err error) {
	intSet := collection.NewStringSet()

	for _, cfg := range c {
		if intSet.Contains(cfg.Address) {
			return false, fmt.Errorf("duplicate contract address %s was found: %w", cfg.Address, ErrDuplicateAddress)
		}

		ok, err = cfg.IsValid(ctx)
		if !ok {
			return false, err
		}

		intSet.Add(cfg.Address)
	}

	return true, nil
}

// IsValid validates the contract config.
func (c ContractConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.Address == "" {
		return false, fmt.Errorf("field Address: %w", ErrRequiredField)
	}
	return true, nil
}
