package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v3"
)

// Config holds the config for the chain.
type Config struct {
	// chainID -> chainConfig
	Chains map[uint32]ChainConfig `yaml:"chains"`
	// Port is the port
	Port uint16 `yaml:"port,omitempty"`
	// RefreshInterval is the refresh interval of rpc latency
	// expressed in seconds
	RefreshInterval int `yaml:"refresh_interval,omitempty"`
	// ClientType is the client type to use
	ClientType string `yaml:"client_type,omitempty"`
}

// ChainConfig is the config for a single chain.
type ChainConfig struct {
	// RPCS is a list of rpcs to use
	RPCs []string `yaml:"rpcs"`
	// Checks is how many rpcs must return the same result for it to be used. This does not apply to height/status based methods
	Checks uint16 `yaml:"confirmations,omitempty"`
}

// UnmarshallConfig unmarshalls a config.
func UnmarshallConfig(input []byte) (cfg Config, err error) {
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}

// Marshall a config to yaml.
func (c Config) Marshall() ([]byte, error) {
	output, err := yaml.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}
