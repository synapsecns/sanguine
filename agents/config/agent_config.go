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

// AgentConfig is used for configuring the guard.
type AgentConfig struct {
	// Domains stores all the domains
	Domains DomainConfigs `yaml:"domains"`
	// DomainID is the domain of the chain that this agent is assigned to.
	// 	For Guards, it is 0 meaning all domains.
	// 	For Notaries, it will be a specific domain greater than 0.
	DomainID uint32 `yaml:"domain_id"`
	// SummitDomainID is the domain of the chain that has the Summit contract (ie SYN chain).
	SummitDomainID uint32 `yaml:"summit_domain_id"`
	// UnbondedSigner contains the unbonded signer config for agents
	// (this is signer used to submit transactions)
	UnbondedSigner SignerConfig `yaml:"unbonded_signer"`
	// BondedSigner contains the bonded signer config for agents
	BondedSigner SignerConfig `yaml:"bonded_signer"`
	// RefreshIntervalSeconds is the refresh interval in seconds
	RefreshIntervalSeconds int `yaml:"refresh_interval_seconds,omitempty"`
	// IsTestHarness marks if this is a test harness
	IsTestHarness bool `yaml:"is_test_harness"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (a *AgentConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if int64(a.RefreshIntervalSeconds) == int64(0) {
		return false, fmt.Errorf("refresh_interval_seconds cannot be 0")
	}

	if ok, err = a.Domains.IsValid(ctx); !ok {
		return false, err
	}

	hasAssignedDomain := (a.DomainID == uint32(0))
	hasSummitDomainID := false

	for _, cfg := range a.Domains {
		if !hasAssignedDomain {
			if cfg.DomainID == a.DomainID {
				hasAssignedDomain = true
			}
		}

		if !hasSummitDomainID {
			if cfg.DomainID == a.SummitDomainID {
				hasSummitDomainID = true
				if cfg.SummitAddress == "" {
					return false, fmt.Errorf("field SummitAddress: %w", ErrRequiredField)
				}
			}
		}
	}

	return true, nil
}

// Encode gets the encoded config.yaml file.
func (a AgentConfig) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&a)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(a), 20), err)
	}
	return output, nil
}

// DecodeAgentConfig parses in a config from a file.
func DecodeAgentConfig(filePath string) (a AgentConfig, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return AgentConfig{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &a)
	if err != nil {
		return AgentConfig{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return a, nil
}
