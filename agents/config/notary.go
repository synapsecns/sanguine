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

// NotaryConfig is used for configuring the notary.
type NotaryConfig struct {
	// OriginDomains stores all origin domains
	OriginDomains DomainConfigs `yaml:"origin_domains"`
	// SummitDomain stores the summit domain
	SummitDomain DomainConfig `yaml:"summit_domain"`
	// DestinationDomain stores  the destination domain
	DestinationDomain DomainConfig `yaml:"destination_domain"`
	// UnbondedSigner contains the unbonded signer config for agents
	// (this is signer used to submit transactions)
	UnbondedSigner SignerConfig `yaml:"unbonded_signer"`
	// BondedSigner contains the bonded signer config for agents
	BondedSigner SignerConfig `yaml:"bonded_signer"`
	// DbConfig is the database config
	Database DBConfig `yaml:"database"`
	// RefreshIntervalInSeconds is how long to wait before refreshing the Notary state
	RefreshIntervalInSeconds int64 `yaml:"refresh_interval_in_seconds"`
	// DBPrefix is the table prefix in the database
	DBPrefix string `yaml:"db_prefix"`
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

// Encode gets the encoded config.yaml file.
func (c NotaryConfig) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}

// DecodeNotaryConfig parses in a notary config from a file.
func DecodeNotaryConfig(filePath string) (cfg NotaryConfig, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return NotaryConfig{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return NotaryConfig{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}
