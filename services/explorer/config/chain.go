package config

import (
	"fmt"
	"github.com/richardwilkes/toolbox/collection"
)

// ChainConfig is the configuration for a chain.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// RPCURL is the RPC of the chain.
	RPCURL string `yaml:"rpc_url"`
	// FetchBlockIncrement is the number of blocks to fetch at a time.
	FetchBlockIncrement uint64 `yaml:"fetch_block_increment"`
	// MaxGoroutines is the maximum number of goroutines that can be spawned.
	MaxGoroutines int `yaml:"max_goroutines"`
	// Contracts are the contracts.
	Contracts ContractConfigs `yaml:"contracts"`
}

// ChainConfigs contains an array fo ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
func (c ChainConfigs) IsValid() (ok bool, err error) {
	intSet := collection.Set[uint32]{}
	for _, cfg := range c {
		if intSet.Contains(cfg.ChainID) {
			return false, fmt.Errorf("chain id %d appears twice", cfg.ChainID)
		}
		ok, err = cfg.IsValid()
		if !ok {
			return false, err
		}
		intSet.Add(cfg.ChainID)
	}
	return true, nil
}

// IsValid validates the chain config.
func (c ChainConfig) IsValid() (ok bool, err error) {
	if c.ChainID == 0 {
		return false, fmt.Errorf("chain ID cannot be 0")
	}
	if c.FetchBlockIncrement == 0 {
		return false, fmt.Errorf("field FetchBlockIncrement: %w", ErrRequiredField)
	}
	if c.MaxGoroutines == 0 {
		return false, fmt.Errorf("must have more than 0 goroutines per chain")
	}
	ok, err = c.Contracts.IsValid()
	if !ok {
		return false, err
	}
	return true, nil
}
