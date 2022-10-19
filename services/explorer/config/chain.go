package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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
	// StartBlocks is a mapping from chain ID -> start block for backfilling.
	StartBlock uint64 `yaml:"start_block"`
	// SynapseBridgeAddress is the address of the SynapseBridge.sol contract
	SynapseBridgeAddress string `yaml:"synapse_bridge_address"`
	// SwapFlashLoanAddresses are the addresses of the SwapFlashLoan.sol contracts
	SwapFlashLoanAddresses []string `yaml:"swap_flash_loan_address"`
	// StartFromLastBlockStored is a flag to start from the last block(s) stored in the database for each chain.
	StartFromLastBlockStored bool `yaml:"start_from_last_block_stored"`
	// MaxGoroutines is the maximum number of goroutines that can be spawned.
	MaxGoroutines int64 `yaml:"max_goroutines"`
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
	if c.MaxGoroutines == 0 {
		return false, fmt.Errorf("must have more than 0 goroutines per chain")
	}
	if c.SynapseBridgeAddress == "" {
		return false, fmt.Errorf("field Address: %w", ErrRequiredField)
	}
	if len(c.SynapseBridgeAddress) != (common.AddressLength*2)+2 {
		return false, fmt.Errorf("field Address: %w", ErrAddressLength)
	}
	for i := range c.SwapFlashLoanAddresses {
		if c.SwapFlashLoanAddresses[i] == "" {
			return false, fmt.Errorf("field Address: %w", ErrRequiredField)
		}
		if len(c.SwapFlashLoanAddresses[i]) != (common.AddressLength*2)+2 {
			return false, fmt.Errorf("address not correct length: %w", ErrAddressLength)
		}
	}
	return true, nil
}
