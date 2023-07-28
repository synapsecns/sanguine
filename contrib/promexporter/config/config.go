package config

import (
	"fmt"
	"github.com/jftuga/ellipsis"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Config contains the config for the prometheues exporter.
type Config struct {
	// Port is the port of the config
	Port int `yaml:"port"`
	// DFKApiUrl is the url of the DFK API
	DFKUrl string `yaml:"dfk_url"`
}

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
