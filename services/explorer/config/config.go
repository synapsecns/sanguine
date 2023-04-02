package config

import (
	"context"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// TODO: these should be put into the contracts themselves and implement a custom type.
const (
	// BridgeContractType is the type of a bridge contract.
	BridgeContractType = "bridge"
	// SwapContractType is the type of the swap contract.
	SwapContractType = "swap"
	// MessageBusContractType is the type of a message bus contract.
	MessageBusContractType = "messagebus"
	// MetaSwapContractType is the type of a meta swap contract.
	MetaSwapContractType = "metaswap"
)

// Config is used to configure the explorer's data consumption.
type Config struct {
	// RefreshRate is the rate at which the explorer will refresh the last block height in seconds.
	RefreshRate int `yaml:"refresh_rate"`
	// ScribeURL is the URL of the Scribe server.
	ScribeURL string `yaml:"scribe_url"`
	// RPCURL is the URL of the RPC server.
	RPCURL string `yaml:"rpc_url"`
	// BridgeConfigAddress is the address of BridgeConfig contract.
	BridgeConfigAddress string `yaml:"bridge_config_address"`
	// BridgeConfigChainID is the ChainID of BridgeConfig contract.
	BridgeConfigChainID uint32 `yaml:"bridge_config_chain_id"`
	// Chains stores the chain configurations.
	Chains ChainConfigs `yaml:"chains"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if c.BridgeConfigAddress == "" {
		return false, fmt.Errorf("field Address: %w", ErrRequiredField)
	}
	if c.ScribeURL == "" {
		return false, fmt.Errorf("field Address: %w", ErrRequiredField)
	}
	if c.RPCURL == "" {
		return false, fmt.Errorf("field RPCURL: %w", ErrRequiredField)
	}

	if len(c.BridgeConfigAddress) != (common.AddressLength*2)+2 {
		return false, fmt.Errorf("field Address: %w", ErrAddressLength)
	}
	if c.BridgeConfigChainID == 0 {
		return false, fmt.Errorf("BridgeConfigChainID chain ID cannot be 0")
	}

	// Checks validity of each chain config.
	if ok, err = c.Chains.IsValid(); !ok {
		return false, err
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
