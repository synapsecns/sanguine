package relconfig

import "github.com/ethereum/go-ethereum/common"

// TODO: validation function.
type Config struct {
	// ChainID: address
	// TODO(aurelius): move under ChainConfig
	Tokens map[int][]string `yaml:"tokens"`
	// ChainID: bridge
	Bridges    map[int]ChainConfig `yaml:"bridges"`
	OmnirpcURL string              `yaml:"omnirpc_url"`
	DBConfig   string
	// TODO: remove, replace w/ pkey recover
	RelayerAddress common.Address
}

type ChainConfig struct {
	// Bridge is the bridge confrimation count.
	Bridge string `yaml:"address"`
	// Confirmations is the number of required confirmations
	Confirmations uint64 `yaml:"confirmations"`
}
