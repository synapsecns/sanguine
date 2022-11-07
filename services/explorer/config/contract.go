package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/richardwilkes/toolbox/collection"
)

// ContractConfig is the configuration for a contract.
type ContractConfig struct {
	// ContractType is the type of contract.
	ContractType string `yaml:"contract_type"`
	// Addresses are the addresses of the contracts
	Address string `yaml:"address"`
	// StartBlock is where to start backfilling this address from.
	StartBlock int64 `yaml:"start_block"`
}

// ContractConfigs contains an array fo ChainConfigs.
type ContractConfigs []ContractConfig

// IsValid validates the contract config by asserting no two contracts appear twice.
func (c ContractConfigs) IsValid() (ok bool, err error) {
	intSet := collection.Set[string]{}
	for _, cfg := range c {
		if cfg.Address == "" {
			return false, fmt.Errorf("field Address: %w", ErrRequiredField)
		}
		if len(cfg.Address) != (common.AddressLength*2)+2 {
			return false, fmt.Errorf("address not correct length: %w", ErrAddressLength)
		}
		if intSet.Contains(cfg.Address) {
			return false, fmt.Errorf("address %s appears twice", cfg.Address)
		}
		ok, err = cfg.IsValid()
		if !ok {
			return false, err
		}
		intSet.Add(cfg.Address)
	}
	return true, nil
}

// IsValid validates the chain config.
func (c ContractConfig) IsValid() (ok bool, err error) {
	if c.ContractType == "" {
		return false, fmt.Errorf("field Address: %w", ErrRequiredField)
	}
	return true, nil
}
