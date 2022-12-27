package config

import (
	"context"
	"fmt"
)

// NotaryConfig is used for configuring the notary.
type NotaryConfig struct {
	// OriginDomains stores all origin domains
	OriginDomains DomainConfigs `toml:"OriginDomains"`
	// AttestationDomain stores the attestaion domain
	AttestationDomain DomainConfig `toml:"AttestationDomain"`
	// DestinationDomain stores  the destination domain
	DestinationDomain DomainConfig `toml:"DestinationDomain"`
	// UnbondedSigner contains the unbonded signer config for agents
	// (this is signer used to submit transactions)
	UnbondedSigner SignerConfig `toml:"UnbondedSigner"`
	// BondedSigner contains the bonded signer config for agents
	BondedSigner SignerConfig `toml:"BondedSigner"`
	// DbConfig is the database config
	Database DBConfig `toml:"Database"`
	// RefreshIntervalInSeconds is how long to wait before refreshing the Notary state
	RefreshIntervalInSeconds int64 `toml:"RefreshIntervalInSeconds"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *NotaryConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.OriginDomains.IsValid(ctx); !ok {
		return false, err
	}

	for _, cfg := range c.OriginDomains {
		if cfg.DomainID == c.DestinationDomain.DomainID {
			return false, fmt.Errorf("origin domain id %d is same as Notary's assigned destination id %d: %w", cfg.DomainID, c.DestinationDomain.DomainID, ErrInvalidDomainID)
		}
	}

	return true, nil
}
