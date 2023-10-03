// Package config contains the config for the prom exporter.
package config

import (
	"fmt"
	"github.com/creasty/defaults"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config contains the config for the prometheues exporter.
type Config struct {
	// Port is the port of the config
	Port int `yaml:"port"`
	// DFKApiUrl is the url of the DFK API
	DFKUrl string `yaml:"dfk_url" default:"https://defi-kingdoms-community-api-gateway-co06z8vi.uc.gateway.dev/graphql"`
	// DFKPending is the list of pending heroes
	DFKPending []DFKPending `yaml:"dfk_pending"`
	// SubmitterChecks is the list of gas checks
	SubmitterChecks []SubmitterChecks `yaml:"gas_checks"`
	OmnirpcURL      string            `yaml:"omnirpc_url" default:"https://rpc.omnirpc.io"`
}

// DFKPending contains the config for the DFK pending metric.
type DFKPending struct {
	// Owner is the owner of the pending heroes
	Owner string `yaml:"owner"`
	// ChainName is the name of the chain
	ChainName string `yaml:"chain_name"`
}

// SubmitterChecks contains the config for the gas checks.
type SubmitterChecks struct {
	// ChainID is the chain id
	ChainIDs []int `yaml:"chain_ids"`
	// Address is the address of the contract
	Address string `yaml:"address"`
	// Name of the address entity
	Name string `yaml:"name"`
}

// DecodeConfig decodes the config from the given file path.
func DecodeConfig(filePath string) (_ Config, err error) {
	cfg := &Config{}
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}

	// set some defaults
	if err := defaults.Set(cfg); err != nil {
		panic(err)
	}

	// set some other defaults that can't be set w/ default
	cfg.DFKPending = []DFKPending{
		{
			Owner:     "0x739B1666c2956f601f095298132773074c3E184b",
			ChainName: "dfk",
		},
		{
			Owner:     "0xEE258eF5F4338B37E9BA9dE6a56382AdB32056E2",
			ChainName: "klatyn",
		},
	}

	// note: when you want to add bridges, you can use the router to look up bridge addresses
	// and get the gas limit from there
	cfg.SubmitterChecks = []SubmitterChecks{
		{
			Address: "0x230a1ac45690b9ae1176389434610b9526d2f21b",
			ChainIDs: []int{
				1, 10, 25, 56, 137, 250, 288, 1088, 1284, 1285, 2000, 7700, 8217, 8453, 42161, 43114, 53935, 1313161554, 1666600000,
			},
			Name: "validators",
		},
		{
			Address: "0xaa920f7b9039e556d2442113f1fd339e4927dd9a",
			ChainIDs: []int{
				53935, 8217, 1666600000,
			},
			Name: "messenger",
		},
		{
			Address: "0x0a1e1d0eb6a1cef79e46f0e2d35b7bf2e958a26a",
			Name:    "cctp",
			ChainIDs: []int{
				1, 42161, 43114, 10,
			},
		},
	}

	err = yaml.Unmarshal(input, cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return *cfg, nil
}
