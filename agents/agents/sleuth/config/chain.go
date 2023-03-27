package config

// ChainConfig defines the config for a specific chain.
type ChainConfig struct {
	// TempRPC is the temporary RPC endpoint for the chain.
	TempRPC string `yaml:"temp_rpc"`
	// ChainID is the ID of the chain.
	ChainID uint32 `yaml:"chain_id"`
	// OriginAddress is the address of the origin contract.
	OriginAddress string `yaml:"origin_address"`
	// DestinationAddress is the address of the destination contract.
	DestinationAddress string `yaml:"destination_address"`
}

// ChainConfigs contains an array of ChainConfigs.
type ChainConfigs []ChainConfig
