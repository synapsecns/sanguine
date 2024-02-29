package executor

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/agents/config"
	signerConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"
)

// Config is used to configure an Executor agent.
type Config struct {
	// DBConfig is the database configuration.
	DBConfig config.DBConfig `yaml:"db_config"`
	// ScribeConfig is the scribe configuration.
	ScribeConfig config.ScribeConfig `yaml:"scribe_config"`
	// Chains stores all chain information
	Chains ChainConfigs `yaml:"chains"`
	// SummitChainID is the chain ID of the chain that the summit contract is deployed on.
	SummitChainID uint32 `yaml:"summit_chain_id"`
	// SummitAddress is the address of the summit contract.
	SummitAddress string `yaml:"summit_address"`
	// InboxAddress is the address of the inbox contract.
	InboxAddress string `yaml:"inbox_address"`
	// BaseOmnirpcURL is the base url for omnirpc.
	// The format is "https://omnirpc.url". Notice the lack of "confirmations" on the URL
	// in comparison to what `Scribe` uses.
	BaseOmnirpcURL string `yaml:"base_omnirpc_url"`
	// UnbondedSigner contains the unbonded signer config for agents
	// (this is signer used to submit transactions)
	UnbondedSigner signerConfig.SignerConfig `yaml:"unbonded_signer"`
	// ExecuteInterval is the interval at which the executor agent will
	// check if messages in the database are ready to be executed.
	ExecuteInterval uint32 `yaml:"execute_interval"`
	// SetMinimumTimeInterval is the interval at which the executor agent will
	// check messages to set their minimum times from attestations.
	SetMinimumTimeInterval uint32 `yaml:"set_minimum_time_interval"`
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
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Chains.IsValid(ctx); !ok {
		return false, err
	}

	if c.BaseOmnirpcURL == "" {
		return false, fmt.Errorf("rpc url cannot be empty")
	}

	if ok, err = c.UnbondedSigner.IsValid(ctx); !ok {
		return false, fmt.Errorf("unbonded signer is invalid: %w", err)
	}

	return true, nil
}

// Encode gets the encoded config.yaml file.
func (c Config) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}

// DecodeConfig parses in a config from a file.
func DecodeConfig(filePath string) (cfg Config, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}
