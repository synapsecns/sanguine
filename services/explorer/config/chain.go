package config

import (
	"context"
	"fmt"
	"github.com/richardwilkes/toolbox/collection"
)

// ChainConfig is the configuration for a chain.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// RPCURL is the ID of the chain.
	RPCURL string `yaml:"chain_id"`
	// FetchBlockIncrement is the number of blocks to fetch at a time.
	FetchBlockIncrement uint32 `yaml:"fetch_block_increment"`
	// StartBlocks is a mapping from chain ID -> start block for backfilling.
	StartBlock uint64 `yaml:"start_blocks"`
	// SynapseBridgeAddress is the address of the SynapseBridge.sol contract
	SynapseBridgeAddress string `yaml:"synapse_bridge_address"`
	// SwapFlashLoanAddresses are the addresses of the SwapFlashLoan.sol contracts
	SwapFlashLoanAddresses []string `yaml:"swap_flash_loan_address"`
}

// ChainConfigs contains an array fo ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
func (c ChainConfigs) IsValid(ctx context.Context) (ok bool, err error) {
	intSet := collection.Set[uint32]{}
	for _, cfg := range c {
		if intSet.Contains(cfg.ChainID) {
			return false, fmt.Errorf("chain id %d appears twice", cfg.ChainID)
		}
		intSet.Add(cfg.ChainID)
	}
	return true, nil
}

// IsValid validates the chain config.
func (c ChainConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.ChainID == 0 {
		return false, fmt.Errorf("chain ID cannot be 0")
	}
	return true, nil
}
