// Package indexerconfig is the config loader for the indexer
package indexerconfig

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/richardwilkes/toolbox/collection"
	"github.com/synapsecns/sanguine/services/sinner/config"

	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
)

// ContractType is the type of contract specified by the config and used for selecting the correct parser.
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=ContractType -linecomment
type ContractType int

const (
	// OriginType is the ContractType for the origin contract.
	OriginType ContractType = iota // origin
	// ExecutionHubType is the ContractType for the execution hub contract.
	ExecutionHubType // execution_hub
	// UnknownType is the ContractType for an unknown contract.
	UnknownType // unknown
)

// ContractTypeFromString converts a string (intended to be from parsed config) into the ContractType type.
func ContractTypeFromString(s string) (ContractType, error) {
	switch s {
	case OriginType.String():
		return OriginType, nil
	case ExecutionHubType.String():
		return ExecutionHubType, nil
	default:
		return UnknownType, fmt.Errorf("unknown contract type: %s", s)
	}
}

// Config is used to configure the sinner's data consumption.
type Config struct {
	// DefaultRefreshRate is the default rate at which data is refreshed in seconds.
	DefaultRefreshRate int `yaml:"default_refresh_rate"`
	// ScribeURL is the URL of the Scribe graphql server.
	ScribeURL string `yaml:"scribe_url"`
	// Chains stores the chain configurations.
	Chains []ChainConfig `yaml:"chains"`
	// DBPath is the path to the database.
	DBPath string `yaml:"db_path"`
	// DBType is the flag signifying the type of database (mysql, sqlite, etc).
	DBType string `yaml:"db_type"`
	// SkipMigrations skips db migrations.
	SkipMigrations bool `yaml:"skip_migrations"`
}

// ChainConfig is the configuration for a chain.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// Contracts are the contracts.
	Contracts []ContractConfig `yaml:"contracts"`
}

// ContractConfig is the configuration for a contract.
type ContractConfig struct {
	// ContractType is the type of contract.
	ContractType string `yaml:"contract_type"`
	// Addresses are the addresses of the contracts
	Address string `yaml:"address"`
	// StartBlock is where to start indexing this address from.
	StartBlock uint64 `yaml:"start_block"`
	// EndBlock is where the end the indexing. This will only backfill a range and will not livefill.
	EndBlock uint64 `yaml:"end_block"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each submodule.
func (c *Config) IsValid() error {
	if c.ScribeURL == "" {
		return fmt.Errorf("scribe_url, %w", config.ErrRequiredGlobalField)
	}
	if len(c.Chains) == 0 {
		return fmt.Errorf("no chains specified for indexing")
	}
	for _, chain := range c.Chains {
		err := chain.IsValid()
		if err != nil {
			return fmt.Errorf("chain with ID %d is invalid: %w", chain.ChainID, err)
		}
	}

	return nil
}

// IsValid validates the chain config.
func (c ChainConfig) IsValid() error {
	intSet := collection.Set[string]{}
	for _, contract := range c.Contracts {
		err := contract.IsValid()
		if err != nil {
			return fmt.Errorf("contract with address %s is invalid: %w", contract.Address, err)
		}
		if intSet.Contains(contract.Address) {
			return fmt.Errorf("address %s appears twice", contract.Address)
		}
		intSet.Add(contract.Address)
	}

	return nil
}

// IsValid validates the chain config.
func (c ContractConfig) IsValid() error {
	_, err := ContractTypeFromString(c.ContractType)
	if err != nil {
		return fmt.Errorf("contract_type %s invalid for address %s", c.ContractType, c.Address)
	}

	switch {
	case c.StartBlock == 0:
		return fmt.Errorf("start_block, %w", config.ErrRequiredContractField)
	case c.Address == "":
		return fmt.Errorf("address, %w", config.ErrRequiredContractField)
	}
	return nil
}

// DecodeConfig parses in a config from a file.
func DecodeConfig(filePath string) (cfg Config, err error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return Config{}, fmt.Errorf("config file does not exist: %w", err)
	}
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
