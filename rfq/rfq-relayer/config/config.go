package config

import (
	"fmt"
	"github.com/jftuga/ellipsis"
	"os"
	"path/filepath"

	ethConfig "github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
	"gopkg.in/yaml.v2"
)

// ChainConfig represents the configuration for each chain.
type ChainConfig struct {
	ChainID                 uint32 `yaml:"chainId"`
	RpcUrl                  string `yaml:"rpcUrl"`
	FastBridgeAddress       string `yaml:"fastBridgeContract"`
	FastBridgeBlockDeployed uint64 `yaml:"fastBridgeBlockDeployed"`
	PollInterval            int    `yaml:"pollInterval"`
	MaxGetLogsRange         uint64 `yaml:"maxGetLogsRange"`
	Confirmations           uint64 `yaml:"confirmations"`
}

type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}

type AssetConfig struct {
	Address string `yaml:"address"`
	ChainID uint32 `yaml:"chainid"`
}

// Config represents the application's configuration structure.
type Config struct {
	Chains         map[uint32]ChainConfig `yaml:"chains"`
	Database       DatabaseConfig         `yaml:"database"`
	Assets         []AssetConfig          `yaml:"assets"`
	RelayerAddress string                 `yaml:"relayer_address"`
	// (this is signer used to submit transactions)
	Signer            ethConfig.SignerConfig `yaml:"unbonded_signer"`
	SubmitterConfig   submitterConfig.Config `yaml:"submitter_config"`
	OmnirpcURL        string                 `yaml:"omnirpc_url"`
	MaxQueueSize      int                    `yaml:"max_queue_size"`
	Deadline          int64                  `yaml:"deadline"`
	QueuePollInterval int                    `yaml:"queue_poll_interval"`
	SkipMigrations    bool                   `yaml:"skip_migrations"`
	RFQURL            string                 `yaml:"rfq_url"`
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
