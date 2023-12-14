package config

type Config struct {
	// ChainID: address
	Tokens map[int][]string `yaml:"tokens"`
	// ChainID: bridge
	Bridges map[int]string
}
