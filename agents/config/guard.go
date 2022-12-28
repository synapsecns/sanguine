package config

import (
	"context"
)

// GuardConfig is used for configuring the guard.
type GuardConfig struct {
	// OriginDomains stores all origin domains
	OriginDomains DomainConfigs `toml:"OriginDomains"`
	// AttestationDomain stores the attestaion domain
	AttestationDomain DomainConfig `toml:"AttestationDomain"`
	// DestinationDomains stores all destination domains
	DestinationDomains DomainConfigs `toml:"DestinationDomains"`
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
func (c *GuardConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.OriginDomains.IsValid(ctx); !ok {
		return false, err
	}
	if ok, err = c.DestinationDomains.IsValid(ctx); !ok {
		return false, err
	}

	return true, nil
}
