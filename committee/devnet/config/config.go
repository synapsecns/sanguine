// Package config provides the configuration for the Committee devnet.
package config

import "github.com/ethereum/go-ethereum/common"

// Config is the config for the Synapse module.
type Config struct {
	// Chains is a map of chain IDs to deployed Synapse Modules.
	Chains map[int]string `yaml:"chains"`
	// ValidatorAddresses is a list of addresses of the validators
	// we want to add to the Synapse Module.
	ValidatorAddresses []common.Address `yaml:"validator_addresses"`
}
