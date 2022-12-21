package config

import (
	"context"
	"fmt"
)

// GuardConfig is used for configuring the guard.
type GuardConfig struct {
	// Domains stores all domains
	Domains DomainConfigs `toml:"Domains"`
	// Signer contains the signer config for agents
	Signer SignerConfig `toml:"Signer"`
	// DbConfig is the database config
	Database DBConfig `toml:"Database"`
	// RefreshIntervalInSeconds is how long to wait before refreshing the Notary state
	RefreshIntervalInSeconds int64 `toml:"RefreshIntervalInSeconds"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *GuardConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Domains.IsValid(ctx); !ok {
		return false, err
	}

	for len(c.Domains) == 0 {
		return false, fmt.Errorf("guard has no domains configured: %w", ErrInvalidDomainID)
	}

	return true, nil
}
