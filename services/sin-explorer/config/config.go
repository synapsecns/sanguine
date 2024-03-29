package config

// Config represents the configuration for the relayer.
type Config struct {
	// Chains is a map of chainID -> chain config.
	Chains map[int]ChainConfig `yaml:"chains"`
	// Omnirpc is the omnirpc configuration.
	OmnirpcURL string `yaml:"omnirpc_url"`
	// Database is the database configuration
	Database DatabaseConfig `yaml:"database"`
}

// ChainConfig represents the configuration for a chain.
type ChainConfig struct {
	// InterchainClientAddress is the interchain client address.
	InterchainClientAddress string `yaml:"interchain_client_address"`
	// StartBlock is the start block.
	StartBlock int64 `yaml:"start_block"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}
