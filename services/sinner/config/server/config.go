// Package serverconfig is the config loader for the server
package serverconfig

import (
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/config"
)

// Config is used to configure the explorer server.
type Config struct {
	// HTTPPort is the http port for the api
	HTTPPort uint16 `yaml:"http_port"`
	// DBPath is the address of the database
	DBPath string `yaml:"db_path"`
	// DBFlag is the address of the database
	DBFlag string `yaml:"db_flag"`
	// SkipMigrations skips the database migrations.
	SkipMigrations bool `yaml:"skip_migrations"`
	// HydrateCache is a flag for enabling cache hydration.
	HydrateCache bool `yaml:"hydrate_cache"`
}

// IsValid makes sure the config is valid.
func (c *Config) IsValid() error {
	switch {
	case c.DBPath == "":
		return fmt.Errorf("db_address, %w", config.ErrRequiredGlobalField)
	}

	return nil
}
