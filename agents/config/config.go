package config

import (
	"context"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
)

// Config is used for configuring the application. It stores the configurations defined in each module.
type Config struct {
	// Domains stores all domains
	Domains DomainConfigs `yaml:"domains"`
	// Signer contains the signer config for agents
	Signer config.SignerConfig `yaml:"signer"`
	// DbConfig is the database config
	Database DBConfig `yaml:"database"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Domains.IsValid(ctx); !ok {
		return false, err
	}

	return true, nil
}
