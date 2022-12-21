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
	// RequiredConfirmations is the number of confirmations required for a block to be finalized.
	RequiredConfirmations uint32 `yaml:"required_confirmations"`
	// Contracts stores all the contract information for the chain.
	Contracts ContractConfigs `yaml:"contracts"`
	// BlockTimeChunkCount is the number of chunks (goroutines) to process at a time while backfilling blocktimes.
	BlockTimeChunkCount uint64 `yaml:"block_time_chunk_count"`
	// BlockTimeChunkSize is the number of blocks to process per chunk (goroutine) while backfilling blocktimes.
	BlockTimeChunkSize uint64 `yaml:"block_time_chunk_size"`
	// ContractSubChunkSize is the number of blocks to request for in each get logs request in the batch request.
	ContractSubChunkSize int `yaml:"contract_sub_chunk_size"`
	// ContractChunkSize is the number of blocks to process per chunk while backfilling contracts.
	ContractChunkSize int `yaml:"contract_chunk_size"`
	// StoreConcurrency is the number of goroutines to use when storing data.
	StoreConcurrency int `yaml:"store_concurrency"`
	// storeConcurrencyThreshold is the max number of block from head in which concurrent store is allowed.
	StoreConcurrencyThreshold uint64 `yaml:"store_concurrency_threshold"`
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
