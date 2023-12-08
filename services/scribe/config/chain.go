package config

import (
	"fmt"

	"github.com/richardwilkes/toolbox/collection"
)

// TODO add tests for this config type

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
	// Confirmations is the number of blocks away from the head to livefill to.
	Confirmations uint64 `yaml:"confirmations"`
	// LivefillThreshold is the number of blocks away from the head minus confirmations to livefill to.
	LivefillThreshold uint64 `yaml:"livefill_threshold"`
	// LivefillRange is the number of blocks that the livefill indexer with request for with get logs at once.
	LivefillRange uint64 `yaml:"livefill_range"`
	// LivefillFlushInterval is how long to wait before flushing the livefill indexer db (in seconds)
	LivefillFlushInterval uint64 `yaml:"livefill_flush_interval"`
	// OmniRPCConfirmations is the OmniRPC URL.
	OmniRPCConfirmations uint `yaml:"omnirpc_confirmations"`
}

// ChainConfigs contains an array of ChainConfigs.
type ChainConfigs []ChainConfig

// IsValid validates the chain config by asserting no two chains appear twice.
// It also calls IsValid on each individual ContractConfig.
func (c ChainConfigs) IsValid() (ok bool, err error) {
	intSet := collection.Set[uint32]{}

	for _, cfg := range c {
		if intSet.Contains(cfg.ChainID) {
			return false, fmt.Errorf("chain id %d appears twice: %w", cfg.ChainID, ErrDuplicateChainID)
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
		return false, fmt.Errorf("%w: chain ID cannot be 0", ErrInvalidChainID)
	}
	if ok, err = c.Contracts.IsValid(); !ok {
		return false, err
	}

	return true, nil
}
