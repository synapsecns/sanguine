package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"github.com/synapsecns/sanguine/services/omnirpc/rpcmap"
	"gopkg.in/yaml.v2"
	"os"
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
}

// ChainConfig is the config for a single chain.
type ChainConfig struct {
	// RPCS is a list of rpcs to use
	RPCs []string `yaml:"rpcs"`
	// Checks is how many rpcs must return the same result for it to be used. This does not apply to height/status absed methods
	Checks int `yaml:"confirmations,omitempty"`
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

// UnmarshallRPCMap unmarshalls an rpc config from an input.
func UnmarshallRPCMap(input string) (*rpcmap.RPCMap, error) {
	var rawMap map[int][]string
	err := yaml.Unmarshal([]byte(input), &rawMap)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall rpc map: %w", err)
	}
	return rpcmap.NewRPCMapFromMap(rawMap), nil
}

// UnmarshallConfigFromFile gets a config from a file.
func UnmarshallConfigFromFile(file string) (*rpcmap.RPCMap, error) {
	//nolint: gosec
	contents, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %w", err)
	}

	return UnmarshallRPCMap(string(contents))
}

// MarshallFromMap marshalls a config from an rpc map.
func MarshallFromMap(rpcMap *rpcmap.RPCMap) string {
	// errors are impossible here
	output, _ := yaml.Marshal(rpcMap.RawMap())

	return string(output)
}
