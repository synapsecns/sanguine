package config

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/richardwilkes/toolbox/collection"
)

// ChainConfig defines the config for a specific chain.
type ChainConfig struct { // ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// CCTPAddress is the address of the CCTP contract.
	// This will correspond to SynapseCCTP or TokenMessenger depending on the CCTPType.
	CCTPAddress string `yaml:"cctp_address"`
	// CCTP start block is the block at which the chain listener will listen for CCTP events.
	CCTPStartBlock uint64 `yaml:"cctp_start_block"`
}

// GetCCTPAddress returns the CCTP address.
func (c ChainConfig) GetCCTPAddress() common.Address {
	return common.HexToAddress(c.CCTPAddress)
}

// ChainConfigs contains an array of ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
func (c ChainConfigs) IsValid(_ context.Context) (ok bool, err error) {
	intSet := collection.Set[uint32]{}

	for _, cfg := range c {
		if intSet.Contains(cfg.ChainID) {
			return false, fmt.Errorf("chain id %d appears twice: %s", cfg.ChainID, "duplicate chain id")
		}
		intSet.Add(cfg.ChainID)

		if len(cfg.CCTPAddress) > 0 && !common.IsHexAddress(cfg.CCTPAddress) {
			return false, fmt.Errorf("invalid address %s: %s", cfg.CCTPAddress, "invalid address")
		}
	}

	return true, nil
}

// IsValid validates the chain config.
func (c ChainConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.ChainID == 0 {
		return false, fmt.Errorf("%s: chain ID cannot be 0", "invalid chain id")
	}

	if c.CCTPAddress == "" {
		return false, fmt.Errorf("a CCTP address is required")
	}

	return true, nil
}
