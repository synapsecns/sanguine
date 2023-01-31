package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
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

// Encode gets the encoded config.yaml file.
func (c GuardConfig) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}

// DecodeGuardConfig parses in a config from a file.
func DecodeGuardConfig(filePath string) (cfg GuardConfig, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return GuardConfig{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return GuardConfig{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}
