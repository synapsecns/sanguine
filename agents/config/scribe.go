package config

import (
	"context"
	"fmt"

	scribeConfig "github.com/synapsecns/sanguine/services/scribe/config"
)

// ScribeConfig is used to configure a scribe for the Executor.
type ScribeConfig struct {
	// Type is the type of scribe. This can be either "embedded" or "remote".
	Type string `yaml:"type"`

	// EmbeddedDBConfig is the database configuration for an embedded scribe.
	EmbeddedDBConfig scribeConfig.DBConfig `yaml:"embedded_db_config,omitempty"`
	// EmbeddedScribeConfig is the config for the embedded scribe.
	EmbeddedScribeConfig scribeConfig.Config `yaml:"embedded_scribe_config,omitempty"`

	// Port is the port to listen on for the remote scribe.
	Port uint `yaml:"port,omitempty"`
	// URL is the URL to connect to for the remote scribe.
	URL string `yaml:"url,omitempty"`
}

// IsValid makes sure the config is valid. This is done by calling IsValid() on each
// submodule. If any method returns an error that is returned here and the entirety
// of IsValid returns false. Any warnings are logged by the submodules respective loggers.
func (c *ScribeConfig) IsValid(ctx context.Context) (ok bool, err error) {
	if ok, err = c.EmbeddedScribeConfig.IsValid(); !ok {
		return false, fmt.Errorf("embedded scribe config is invalid: %w", err)
	}

	return true, nil
}
