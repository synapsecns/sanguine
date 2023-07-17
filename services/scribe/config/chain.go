package config

import (
	"context"
	"fmt"

	"github.com/richardwilkes/toolbox/collection"
)

// TODO add tests for this config type

// ConfirmationConfig holds config data for reorg protection.
type ConfirmationConfig struct {
	// RequiredConfirmations is the number of confirmations required for a block to be finalized.
	RequiredConfirmations uint32 `yaml:"required_confirmations"`
	// ConfirmationThreshold is the number of blocks to wait until doing a reorg check.
	ConfirmationThreshold uint64 `yaml:"confirmation_threshold"`
	// ConfirmationMinWait is the amount of time in seconds to wait before checking confirmations
	ConfirmationRefreshRate int `yaml:"confirmation_min_wait"`
}

// ChainConfig defines the config for a specific chain.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// Contracts stores all the contract information for the chain.
	Contracts ContractConfigs `yaml:"contracts"`
	// GetLogsRange is the max number of blocks to request in a single getLogs request.
	GetLogsRange uint64 `yaml:"get_logs_range"`
	// GetLogsBatchAmount is the number of getLogs requests to include in a single batch request.
	GetLogsBatchAmount uint64 `yaml:"get_logs_batch_amount"`
	// StoreConcurrency is the number of goroutines to use when storing data.
	StoreConcurrency int `yaml:"store_concurrency"`
	// ConcurrencyThreshold is the max number of block from head in which concurrent operations (store, getlogs) is allowed.
	ConcurrencyThreshold uint64 `yaml:"concurrency_threshold"`
	// GetBlockBatchSize is the amount of blocks to get at a time when doing confirmations.
	GetBlockBatchAmount int `yaml:"get_block_batch_amount"`
	// ConfirmationConfig holds config data for reorg protection.
	ConfirmationConfig ConfirmationConfig `yaml:"confirmation_config"`
}

// ChainConfigs contains an array of ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
// It also calls IsValid on each individual ContractConfig.
func (c ChainConfigs) IsValid(ctx context.Context) (ok bool, err error) {
	intSet := collection.Set[uint32]{}

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
	if ok, err = c.Contracts.IsValid(ctx); !ok {
		return false, err
	}

	return true, nil
}
