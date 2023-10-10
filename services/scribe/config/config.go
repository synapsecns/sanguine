package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/davecgh/go-spew/spew"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
)

// Config is used to configure a Scribe instance and information about chains and contracts.
type Config struct {
	// Chains stores all chain information
	Chains ChainConfigs `yaml:"chains"`
	// RPCURL is the url of the omnirpc.
	RPCURL string `yaml:"rpc_url"`
	// Verbose is used to enable verbose logging.
	Verbose bool `yaml:"verbose"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid() (ok bool, err error) {
	if ok, err = c.Chains.IsValid(); !ok {
		return false, err
	}
	if c.RPCURL == "" {
		return false, fmt.Errorf("%w: rpc url cannot be empty", ErrRequiredField)
	}
	return true, nil
}

// Encode gets the encoded config.yaml file.
func (c Config) Encode() ([]byte, error) {
	output, err := yaml.Marshal(&c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(spew.Sdump(c), 20), err)
	}
	return output, nil
}

// DecodeConfig parses in a config from a file.
func DecodeConfig(filePath string) (cfg Config, err error) {
	input, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return Config{}, fmt.Errorf("failed to read file: %w", err)
	}
	err = yaml.Unmarshal(input, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("could not unmarshall config %s: %w", ellipsis.Shorten(string(input), 30), err)
	}
	return cfg, nil
}
