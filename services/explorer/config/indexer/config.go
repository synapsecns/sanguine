// Package indexerconfig is the config loader for the indexer
package indexerconfig

import (
	"fmt"
	"github.com/richardwilkes/toolbox/collection"
	"github.com/synapsecns/sanguine/services/explorer/config"
	"os"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
)

// ContractType is the type of contract specified by the config and used for selecting the correct parser.
type ContractType int

const (
	// BridgeContractType is the ContractType for the bridge contract.
	BridgeContractType ContractType = iota
	// SwapContractType is the ContractType for the swap contract.
	SwapContractType
	// MessageBusContractType is the ContractType for the message bus contract.
	MessageBusContractType
	// MetaSwapContractType is the ContractType for the meta swap contract.
	MetaSwapContractType
	// CCTPContractType is the ContractType for the cctp contract.
	CCTPContractType
)

func (c ContractType) String() string {
	return [...]string{"bridge", "swap", "messagebus", "metaswap", "cctp"}[c]
}

// ContractTypeFromString converts a string (intended to be from parsed config) into the ContractType type.
func ContractTypeFromString(s string) (ContractType, error) {
	switch s {
	case "bridge":
		return BridgeContractType, nil
	case "swap":
		return SwapContractType, nil
	case "messagebus":
		return MessageBusContractType, nil
	case "metaswap":
		return MetaSwapContractType, nil
	case "cctp":
		return CCTPContractType, nil
	default:
		return -1, fmt.Errorf("unknown contract type: %s", s)
	}
}

// Config is used to configure the explorer's data consumption.
type Config struct {
	// DefaultRefreshRate is the default rate at which data is refreshed.
	DefaultRefreshRate int `yaml:"default_refresh_rate"`
	// ScribeURL is the URL of the Scribe server.
	ScribeURL string `yaml:"scribe_url"`
	// RPCURL is the URL of the RPC server.
	RPCURL string `yaml:"rpc_url"`
	// BridgeConfigAddress is the address of BridgeConfig contract.
	BridgeConfigAddress string `yaml:"bridge_config_address"`
	// BridgeConfigChainID is the ChainID of BridgeConfig contract.
	BridgeConfigChainID uint32 `yaml:"bridge_config_chain_id"`
	// Chains stores the chain configurations.
	Chains []ChainConfig `yaml:"chains"`
}

// ChainConfig is the configuration for a chain.
type ChainConfig struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// RPCURL is the RPC of the chain.
	RPCURL string `yaml:"rpc_url"`
	// FetchBlockIncrement is the number of blocks to fetch at a time.
	FetchBlockIncrement uint64 `yaml:"fetch_block_increment"`
	// MaxGoroutines is the maximum number of goroutines that can be spawned.
	MaxGoroutines int `yaml:"max_goroutines"`
	// Contracts are the contracts.
	Contracts []ContractConfig `yaml:"contracts"`
}

// ContractConfig is the configuration for a contract.
type ContractConfig struct {
	// ContractType is the type of contract.
	ContractType string `yaml:"contract_type"`
	// Addresses are the addresses of the contracts
	Address string `yaml:"address"`
	// StartBlock is where to start backfilling this address from.
	StartBlock int64 `yaml:"start_block"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid() error {
	switch {
	case c.BridgeConfigAddress == "":
		return fmt.Errorf("bridge_config_address, %w", config.ErrRequiredGlobalField)
	case c.ScribeURL == "":
		return fmt.Errorf("scribe_url, %w", config.ErrRequiredGlobalField)
	case c.RPCURL == "":
		return fmt.Errorf("rpc_url, %w", config.ErrRequiredGlobalField)
	case c.BridgeConfigChainID == 0:
		return fmt.Errorf("chain_id cannot be 0")
	}
	if len(c.BridgeConfigAddress) != (common.AddressLength*2)+2 {
		return fmt.Errorf("field Address: %w", config.ErrAddressLength)
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
	switch {
	case c.ChainID == 0:
		return fmt.Errorf("chain_id, %w", config.ErrRequiredGlobalField)
	case c.MaxGoroutines == 0:
		return fmt.Errorf("max_goroutines, %w", config.ErrRequiredGlobalField)
	}

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
