package config

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v3"
)

// RPCType is the type of the swap event.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=RPCType
type RPCType uint8

const (
	// Stable is the rpc type for RPCs without restrictions such as hosted nodes.
	Stable RPCType = iota
	// Auxiliary is the rpc type for RPCs with restrictive rate limiting or other restrictions (public).
	Auxiliary
)

func StringToRPCType(s string) RPCType {
	switch s {
	case "stable":
		return Stable
	case "auxiliary":
		return Auxiliary
	default:
		return Auxiliary // default to Auxiliary
	}
}

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
	// RPCs is a list of rpcs to use
	RPCs []RPCConfig `json:"rpcs" yaml:"rpcs"`
	// Checks is how many rpcs must return the same result for it to be used. This does not apply to height/status based methods
	Checks uint16 `yaml:"confirmations,omitempty"`
}

// RPCConfig is the config for an RPC.
type RPCConfig struct {
	// RPC is the endpoint
	RPC string `json:"rpc" yaml:"rpc"`
	// RPCType is the type of rpc
	RPCType string `json:"rpc_type" yaml:"rpc_type"`
}

func FlattenRPCs(RPCs []RPCConfig) []string {
	rpcs := make([]string, len(RPCs))
	for i, rpc := range RPCs {
		rpcs[i] = rpc.RPC
	}
	return rpcs
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
