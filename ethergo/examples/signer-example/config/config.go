// Package config provides configuration for the example signer.
package config

import "github.com/synapsecns/sanguine/ethergo/signer/config"

// ExampleConfig is a configuration for the example signer.
type ExampleConfig struct {
	config.SignerConfig `yaml:"signer_config"`
}
