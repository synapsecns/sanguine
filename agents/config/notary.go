package config

import (
	"context"
	"fmt"
)

// NotaryConfig is used for configuring the notary.
type NotaryConfig struct {
	// Destination is the destination domain of this notary is assigned to
	DestinationID uint32 `toml:"DestinationID"`
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
func (c *NotaryConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Domains.IsValid(ctx); !ok {
		return false, err
	}

	for _, cfg := range c.Domains {
		if cfg.DomainID == c.DestinationID {
			return false, fmt.Errorf("origin domain id %d is same as Notary's assigned destination id %d: %w", cfg.DomainID, c.DestinationID, ErrInvalidDomainID)
		}
	}

	return true, nil
}
