// Package serverconfig is the config loader for the server
package serverconfig

import (
	"fmt"
	"github.com/richardwilkes/toolbox/collection"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"os"
	"path/filepath"

	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
)

// Config is used to configure the explorer server.
type Config struct {
	// HTTPPort is the http port for the api
	HTTPPort uint16 `yaml:"http_port"`
	// DBAddress is the address of the database
	DBAddress string `yaml:"db_address"`
	// HydrateCache is a flag for enabling cache hydration.
	HydrateCache bool `yaml:"hydrate_cache"`
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

// ChainConfig is the config for each chain in the server config.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	/// GetLogsRange is the max number of blocks to request in a single getLogs request.
	GetLogsRange uint64 `yaml:"get_logs_range"`
	// GetLogsBatchAmount is the number of getLogs requests to include in a single batch request.
	GetLogsBatchAmount uint64 `yaml:"get_logs_batch_amount"`
	// BlockTime is the block time of the chain.
	BlockTime uint64 `yaml:"avg_block_time"`
	// Swaps are the addresses of the swaps on the chain for parsing token address logs.
	Swaps []string `yaml:"swaps"`
	// Chains stores the chain configurations.
	Contracts ContractsConfig `yaml:"contracts"`
}

// ContractsConfig is config for each contract in the server config.
type ContractsConfig struct {
	// CCTP is the address of the cctp contract
	CCTP string `yaml:"cctp"`
	// Bridge is the URL of the Scribe server.
	Bridge string `yaml:"bridge"`
}

// IsValid makes sure the config is valid.
func (c *Config) IsValid() error {
	switch {
	case c.ScribeURL == "":
		return fmt.Errorf("scribe_url, %w", config.ErrRequiredGlobalField)
	case c.RPCURL == "":
		return fmt.Errorf("rpc_url, %w", config.ErrRequiredGlobalField)
	case c.BridgeConfigAddress == "":
		return fmt.Errorf("bridge_config_address, %w", config.ErrRequiredGlobalField)
	case c.BridgeConfigChainID == 0:
		return fmt.Errorf("bridge_config_chain_id, %w", config.ErrRequiredGlobalField)
	case c.DBAddress == "":
		return fmt.Errorf("db_address, %w", config.ErrRequiredGlobalField)
	}
	intSet := collection.Set[uint32]{}
	for _, chain := range c.Chains {
		err := chain.IsValid()
		if err != nil {
			return err
		}
		if intSet.Contains(chain.ChainID) {
			return fmt.Errorf("chain id %d appears twice in the server", chain.ChainID)
		}
		intSet.Add(chain.ChainID)
	}

	return nil
}

// IsValid checks if the entered ChainConfig is valid.
func (c *ChainConfig) IsValid() error {
	switch {
	case c.ChainID == 0:
		return fmt.Errorf("chain_id cannot be 0")
	case c.GetLogsRange == 0:
		return fmt.Errorf("get_logs_range, %w", config.ErrRequiredChainField)
	case c.GetLogsBatchAmount == 0:
		return fmt.Errorf("get_logs_range, %w", config.ErrRequiredChainField)
	case c.BlockTime == 0:
		return fmt.Errorf("block_time, %w", config.ErrRequiredChainField)
	}
	err := c.Contracts.IsValid()
	if err != nil {
		return err
	}
	return nil
}

// IsValid checks if the entered ContractsConfig is valid.
func (c ContractsConfig) IsValid() error {
	if c.CCTP == "" && c.Bridge == "" {
		return fmt.Errorf("one contract must be specified on each contract config")
	}
	return nil
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

	err = cfg.IsValid()
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
