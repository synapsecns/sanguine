package config

import (
	"context"
)

// Config is used to configure a Scribe instance and information about chains and contracts.
type Config struct {
	// Chains stores all chain information
	Chains ChainConfigs `toml:"Chains"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Chains.IsValid(ctx); !ok {
		return false, err
	}

	return true, nil
}
