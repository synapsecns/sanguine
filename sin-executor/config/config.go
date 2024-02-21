package config

import (
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
)

// Config is the config for the Synapse module.
type Config struct {
	// Chains is a map of chain IDs to chain configs.
	Chains map[int]ChainConfig `yaml:"chains"`
	// OmnirpcURL is the URL of the Omni RPC.
	OmnirpcURL string
	// Database is the database config.
	Database DatabaseConfig `yaml:"database"`
	// Signer is the signer config.
	Signer config.SignerConfig `yaml:"signer"`
	// Submitter is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
}

// ChainConfig is the config for a chain.
type ChainConfig struct {
	// SynapseClientAddress is the address of the Synapse module.
	SynapseClientAddress string `yaml:"synapse_client"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}
