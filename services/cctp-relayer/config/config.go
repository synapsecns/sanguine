package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/ImVexed/fasturl"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	ethConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	"github.com/synapsecns/sanguine/services/cctp-relayer/types"
	"gopkg.in/yaml.v2"
)

// Config is used to configure an Executor agent.
type Config struct {
	// Port is the RelayerAPIServer port
	Port uint16 `yaml:"port"`
	// Host is the RelayerAPIServer host
	Host string `yaml:"host"`
	// CircleAPIURl is the URL for the Circle API
	CircleAPIURl string `yaml:"circle_api_url"`
	// CCTPType is the method for executing CCTP transactions.
	CCTPType string `yaml:"cctp_type"`
	// Chains stores all chain information
	Chains ChainConfigs `yaml:"chains"`
	// BaseOmnirpcURL is the base url for omnirpc.
	// The format is "https://omnirpc.url/". Notice the lack of "confirmations" on the URL
	// in comparison to what `Scribe` uses.
	BaseOmnirpcURL string `yaml:"base_omnirpc_url"`
	// Signer contains the unbonded signer config for agents
	// (this is signer used to submit transactions)
	Signer ethConfig.SignerConfig `yaml:"unbonded_signer"`
	// RetryInterval is the interval for attestation request retries
	RetryIntervalMS int `yaml:"retry_interval_ms"`
	// HTTPBackoffMaxElapsedTime is the max elapsed time for attestation request retries
	HTTPBackoffMaxElapsedTimeMs int `yaml:"http_backoff_max_elapsed_time_ms"`
	// SubmitterConfig is the config for the transaction submitter
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
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

	if _, err := fasturl.ParseURL(c.BaseOmnirpcURL); err != nil {
		return false, fmt.Errorf("rpc url is invalid: %w", err)
	}

	if _, err := c.Chains.IsValid(ctx); err != nil {
		return false, fmt.Errorf(fmt.Errorf("could not validate chains: %w", err).Error())
	}

	if ok, err = c.Signer.IsValid(ctx); !ok {
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

const defaultCCTPType = types.SynapseMessageType

// GetCCTPType returns the CCTP method.
func (c Config) GetCCTPType() (types.MessageType, error) {
	switch c.CCTPType {
	case "synapse":
		return types.SynapseMessageType, nil
	case "circle":
		return types.CircleMessageType, nil
	default:
		if len(c.CCTPType) == 0 {
			return defaultCCTPType, nil
		}
		return 0, fmt.Errorf("invalid cctp method: %s", c.CCTPType)
	}
}
