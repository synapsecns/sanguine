// Package config provides the configuration for the Committee devnet.
package config

import "github.com/ethereum/go-ethereum/common"

// Config is the config for the Synapse module.
type ProvisionerConfig struct {
	// Chains is a map of chain IDs to deployed Synapse Modules.
	Chains map[int]string `yaml:"chains"`
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
}
