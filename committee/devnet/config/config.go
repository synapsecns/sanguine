// Package config provides the configuration for the Committee devnet.
package config

import (
	"github.com/ethereum/go-ethereum/common"
	dbconfig "github.com/synapsecns/sanguine/committee/config"
	"github.com/synapsecns/sanguine/ethergo/signer/config"
	submitterConfig "github.com/synapsecns/sanguine/ethergo/submitter/config"
)

// Config is the config for the Synapse module.
type ProvisionerConfig struct {
	// Chains is a map of chain IDs to deployed Synapse Modules.
	SynapseModuleDeployments map[int]string `yaml:"synapse_module_deployments"`
	// GasOracleDeployments is a map of chain IDs to deployed Gas Oracles
	GasOracleDeployments map[int]string `yaml:"gas_oracle_deployments"`
	// ValidatorAddresses is a list of addresses of the validators we want to add to the Synapse Module.
	ValidatorAddresses []common.Address `yaml:"validator_addresses"`
	// OmniRPCURL is the url for OmniRPC
	OmnirpcURL string `yaml:"omnirpc_url"`
}

type SenderConfig struct {
	// OmniRPCURL is the url for OmniRPC
	OmnirpcURL string `yaml:"omnirpc_url"`
	// OriginChainID is the chain ID of the origin chain.
	OriginChainID int `yaml:"origin_chain_id"`
	// DestinationChainID is the chain ID of the destination chain.
	DestinationChainID int `yaml:"destination_chain_id"`
	// Signer is the signer config.
	SignerConfig config.SignerConfig `yaml:"signer_config"`
	// Submitter is the submitter config.
	SubmitterConfig submitterConfig.Config `yaml:"submitter_config"`
	// Database is the database config.
	Database dbconfig.DatabaseConfig `yaml:"database"`
}
