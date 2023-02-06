package config

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strconv"
)

// Config is the result of the parsed yaml file.
type Config struct {
	Cleanup      bool         `yaml:"cleanup"`
	Workers      int          `yaml:"workers"`
	Repositories []Repository `yaml:"repositories,flow"`
	Target       TargetConfig `yaml:"target"`
}

// TargetConfig contains info on where to mirror repositories to.
type TargetConfig struct {
	Registry string `yaml:"registry"`
	Prefix   string `yaml:"prefix"`
}

// Repository is a single docker hub repository to mirror.
type Repository struct {
	PrivateRegistry string            `yaml:"private_registry"`
	Name            string            `yaml:"name"`
	MatchTags       []string          `yaml:"match_tag"`
	DropTags        []string          `yaml:"ignore_tag"`
	MaxTags         int               `yaml:"max_tags"`
	MaxTagAge       *Duration         `yaml:"max_tag_age"`
	RemoteTagSource string            `yaml:"remote_tags_source"`
	RemoteTagConfig map[string]string `yaml:"remote_tags_config"`
	TargetPrefix    *string           `yaml:"target_prefix"`
	Host            string            `yaml:"host"`
}

// Setup is a function that can be called to perform any setup required for a config.
func (c *Config) Setup() error {
	if c.Target.Registry == "" {
		return errors.New("missing `target -> registry` yaml c")
	}

	if c.Workers == 0 {
		c.Workers = runtime.NumCPU()
	}

	// number of workers
	if w := os.Getenv("NUM_WORKERS"); w != "" {
		p, err := strconv.Atoi(w)
		if err != nil {
			log.Fatal(fmt.Sprintf("Could not parse NUM_WORKERS env: %s", err))
		}

		c.Workers = p
	}

	return nil
}
