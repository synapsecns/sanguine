package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
)

// AgentConfig is used for configuring the guard.
type AgentConfig struct {
	// DBConfig is the database configuration.
	DBConfig DBConfig `yaml:"db_config"`
	// ScribeConfig is the scribe configuration.
	ScribeConfig ScribeConfig `yaml:"scribe_config"`
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
	UnbondedSigner config.SignerConfig `yaml:"unbonded_signer"`
	// BondedSigner contains the bonded signer config for agents
	BondedSigner config.SignerConfig `yaml:"bonded_signer"`
	// OwnerSigner contains the owner signer config for agents (optional).
	OwnerSigner config.SignerConfig `yaml:"owner_signer"`
	// RefreshIntervalSeconds is the refresh interval in seconds
	RefreshIntervalSeconds uint32 `yaml:"refresh_interval_seconds,omitempty"`
	// BaseOmnirpcURL is the base url for omnirpc.
	// The format is "https://omnirpc.url". Notice the lack of "confirmations" on the URL
	// in comparison to what `Scribe` uses.
	BaseOmnirpcURL string `yaml:"base_omnirpc_url"`
	// DBPrefix is the prefix for the tables in the database. This is only to be used with mysql.
	DBPrefix string `yaml:"db_prefix"`
	// SubmitterConfig is the config for the submitter.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
	// MaxRetrySeconds is the maximum number of seconds to retry an RPC call (not a transaction).
	MaxRetrySeconds uint32 `yaml:"max_retry_seconds"`
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

	if a.BaseOmnirpcURL == "" {
		return false, fmt.Errorf("rpc url cannot be empty")
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
