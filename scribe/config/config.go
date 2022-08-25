package config

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/BurntSushi/toml"
	tomlCommon "github.com/synapsecns/sanguine/core/toml"
)

// Config is used to configure a Scribe instance and information about chains and contracts.
type Config struct {
	// Chains stores all chain information
	Chains ChainConfigs `toml:"Chains"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *Config) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.Chains.IsValid(ctx); !ok {
		return false, err
	}

	return true, nil
}

// Encode gets the encoded config.toml file.
func (c *Config) Encode() (string, error) {
	marshalledConfig := struct {
		Chains ChainConfigs `toml:"Chains"`
	}{
		Chains: c.Chains,
	}
	file, err := tomlCommon.Encode(marshalledConfig)
	if err != nil {
		return "", fmt.Errorf("could not encode file: %w", err)
	}
	// currently, there's a bug in the parser that requires maps to be on the same level as the parent.
	// TODO: fix
	splitFile := strings.Split(file, "\n")
	var newLines []string
	for _, line := range splitFile {
		// get rid of double spacing on maps
		indentLen := len(tomlCommon.Indent) * 2
		newLines = append(newLines, strings.ReplaceAll(line, getStringOfLength(indentLen), getStringOfLength(indentLen-2)))
	}
	return strings.Join(newLines, "\n"), nil
}

// getStringOfLength generates a blank string of length.
func getStringOfLength(length int) (res string) {
	for i := 0; i < length; i++ {
		res += " "
	}
	return res
}

// DecodeConfig parses in a config from a file.
func DecodeConfig(filepath string) (cfg *Config, err error) {
	cfg = &Config{}
	_, err = toml.DecodeFile(filepath, cfg)
	if err != nil {
		return nil, fmt.Errorf("could not parse config at path %s into %s: %w", filepath, reflect.TypeOf(cfg), err)
	}
	return cfg, nil
}
