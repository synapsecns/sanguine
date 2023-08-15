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

// Config is used to configure the explorer's data consumption.
type Config struct {
	// HTTPPort is the http port for the api
	HTTPPort uint16
	// DBAddress is the address of the database
	DBAddress string
	// HydrateCache is whether or not to hydrate the cache
	HydrateCache bool
	// ScribeURL is the URL of the Scribe server.
	ScribeURL string `yaml:"scribe_url"`
	// RPCURL is the URL of the RPC server.
	RPCURL string `yaml:"rpc_url"`
	// BridgeConfigAddress is the address of BridgeConfig contract.
	BridgeConfigAddress string `yaml:"bridge_config_address"`
	// BridgeConfigChainID is the ChainID of BridgeConfig contract.
	BridgeConfigChainID uint32 `yaml:"bridge_config_chain_id"`
	// SwapTopicHash is the hash of the swap topic.
	SwapTopicHash string `yaml:"swap_topic_hash"`
	// Chains stores the chain configurations.
	Chains map[uint32]ChainConfig `yaml:"chains"`
}

type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	/// GetLogsRange is the max number of blocks to request in a single getLogs request.
	GetLogsRange uint64 `yaml:"get_logs_range"`
	// GetLogsBatchAmount is the number of getLogs requests to include in a single batch request.
	GetLogsBatchAmount uint64 `yaml:"get_logs_batch_amount"`
	// BlockTime is the block time of the chain.
	BlockTime uint64 `yaml:"block_time"`
	// Swaps are the addresses of the swaps on the chain for parsing token address logs.
	Swaps []string `yaml:"swaps"`
	// Chains stores the chain configurations.
	Contracts ContractsConfig `yaml:"contracts"`
}

type ContractsConfig struct {
	// CCTP is the address of the cctp contract
	CCTP string `yaml:"cctp"`
	// Bridge is the URL of the Scribe server.
	Bridge string `yaml:"bridge"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if c.ScribeURL == "" || c.RPCURL == "" || c.BridgeConfigAddress == "" || c.BridgeConfigChainID == 0 || c.DBAddress == "" {
		return false, fmt.Errorf("A required global config field is empty")
	}
	for _, chain := range c.Chains {
		ok, err = chain.IsValid(ctx)
		if !ok {
			return false, err
		}
		ok, err = chain.Contracts.IsValid(ctx)
		if !ok {
			return false, err
		}
	}
	return true, nil
}

func (c *ChainConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.ChainID == 0 {
		return false, fmt.Errorf("chain ID cannot be 0")
	}
	return true, nil
}

func (c ContractsConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if c.CCTP == "" && c.Bridge == "" {
		return false, fmt.Errorf("one contract must be specified on each contract config")
	}
	return true, nil
}

// EncodeServerConfig gets the encoded config.yaml file.
func (c Config) EncodeServerConfig() ([]byte, error) {
	output, err := yaml.Marshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}

// DecodeServerConfig parses in a config from a file.
func DecodeServerConfig(filePath string) (cfg Config, err error) {
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
