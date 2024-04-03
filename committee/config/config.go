// Package config provides the configuration for the Synapse module.
package config

import (
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
)

// Config is the config for the Synapse module.
type Config struct {
	// Chains is a map of chain IDs to chain configs.
	Chains map[int]string `yaml:"chains"`
	// OmnirpcURL is the URL of the Omni RPC.
	OmnirpcURL string `yaml:"omnirpc_url"`
	// Database is the database config.
	Database DatabaseConfig `yaml:"database"`
	// Signer is the signer config.
	Signer config.SignerConfig `yaml:"signer"`
	// Submitter is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
	// ShouldRelay is whether the node should relay.
	ShouldRelay bool `yaml:"should_relay"`
	// BootstrapPeers is the list of bootstrap peers.
	BootstrapPeers []string `yaml:"bootstrap_peers"`
	// P2PPort is the port for the p2p server.
	P2PPort int `yaml:"p2p_port"`
	// UsePeerID is wether or not to use the secp256k1 key for peer identification. This is a beta feature and currently can lead to rate limits/high bills on kms so should be turned off for now.
	UsePeerID bool `yaml:"use_peer_id"`
}

// DatabaseConfig represents the configuration for the database.
type DatabaseConfig struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"` // Data Source Name
}
