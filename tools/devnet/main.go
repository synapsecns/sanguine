package main

import (
	"flag"
	"os"

	"gopkg.in/yaml.v2"
)

type ChainConfig struct {
	Name    string `yaml:"name"`
	ChainID int    `yaml:"chain-id"`
	Port    string `yaml:"port"`
}

type DockerComposeConfig struct {
	Services map[string]ChainConfig `yaml:"services"`
}

func main() {
	var dockerPath string
	flag.StringVar(&dockerPath, "d", "", "path docker compose file")
	flag.Parse()
	if len(dockerPath) == 0 {
		panic("expected docker path to be set (use -d flag)")
	}

	// Read the Docker Compose YAML file
	data, err := os.ReadFile(dockerPath)
	if err != nil {
		panic(err)
	}

	// Parse the YAML data into a DockerComposeConfig struct
	var dockerComposeConfig DockerComposeConfig
	err = yaml.Unmarshal(data, &dockerComposeConfig)
	if err != nil {
		panic(err)
	}
}
