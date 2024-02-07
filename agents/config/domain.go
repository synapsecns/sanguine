package config

import (
	"context"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/utils"
	"github.com/richardwilkes/toolbox/collection"
)

// DomainConfigs contains a map of name->domain config.
type DomainConfigs map[string]DomainConfig

// IsValid validates the domain configs by asserting no two domains appear twice
// it also calls IsValid on each individual DomainConfig.
func (d DomainConfigs) IsValid(ctx context.Context) (ok bool, err error) {
	intSet := collection.NewSet[uint32]()

	for _, cfg := range d {
		if intSet.Contains(cfg.DomainID) {
			return false, fmt.Errorf("domain id %d appears twice: %w", cfg.DomainID, ErrInvalidDomainID)
		}

		ok, err = cfg.IsValid(ctx)
		if !ok {
			return false, err
		}

		intSet.Add(cfg.DomainID)
	}

	return true, nil
}

// DomainConfig defines the config for a specific domain.
type DomainConfig struct {
	// DomainID is the domain of the chain
	DomainID uint32 `yaml:"domain_id"`
	// Type of the chain (e.g. evm)
	Type string `yaml:"type"`
	// RequiredConfirmations is the number of confirmations to way
	RequiredConfirmations uint32 `yaml:"required_confirmations"`
	// OriginAddress gets origin contract address
	OriginAddress string `yaml:"origin_address"`
	// SummitAddress contains the summit address (if present)
	SummitAddress string `yaml:"summit_address"`
	// DestinationAddress gets destination contract address
	DestinationAddress string `yaml:"destination_address"`
	// LightManagerAddress gets then light manager contract address
	LightManagerAddress string `yaml:"light_manager_address"`
	// BondingManagerAddress gets the bonding manager contract address
	BondingManagerAddress string `yaml:"bonding_manager_address"`
	// LightInboxAddress gets the light inbox contract address
	LightInboxAddress string `yaml:"light_inbox_address"`
	// InboxAddress gets the inbox contract address
	InboxAddress string `yaml:"inbox_address"`
}

// IsValid validates the domain config.
func (d DomainConfig) IsValid(_ context.Context) (ok bool, err error) {
	if !utils.ContainsString(chainTypeList, d.Type) {
		return false, fmt.Errorf("invalid chain type %s, %w", d.Type, ErrInvalidChainType)
	}

	// TODO: this might require more significant checking against a list of presets (for london/non-london etc)
	if d.DomainID == 0 {
		return false, fmt.Errorf("%w: cannot be 0", ErrInvalidDomainID)
	}

	// TODO: we should defer to chain-specific config here for verification
	if d.OriginAddress == "" {
		return false, fmt.Errorf("field OriginAddress: %w", ErrRequiredField)
	}

	if d.DestinationAddress == "" {
		return false, fmt.Errorf("field DestinationAddress: %w", ErrRequiredField)
	}

	if d.SummitAddress == "" {
		if d.LightManagerAddress == "" {
			return false, fmt.Errorf("field LightManagerAddress: %w", ErrRequiredField)
		}
	} else {
		if d.BondingManagerAddress == "" {
			return false, fmt.Errorf("field BondingManagerAddress: %w", ErrRequiredField)
		}
	}

	return true, nil
}
