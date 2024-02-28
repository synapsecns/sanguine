// Package config provides the configuration for the Synapse module.
package config

import (
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
)

// Config is the config for the Synapse module.
type Config struct {
	// Chains is a map of chain IDs to execution service contract addresses.
	Chains map[int]ChainConfig `yaml:"chains"`
	// OmnirpcURL is the URL of the Omni RPC.
	OmnirpcURL string `yaml:"omnirpc_url"`
	// Database is the database config.
	Database DatabaseConfig `yaml:"database"`
	// Signer is the signer config.
	Signer config.SignerConfig `yaml:"signer"`
	// Submitter is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
}

type ChainConfig struct {
	// ExecutionService is the address of the execution service contract.
	ExecutionService string `yaml:"execution_service"`
	// Client is the address of the interchain client contract.
	Client string `yaml:"client"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}
