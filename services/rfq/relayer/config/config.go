package config

type Config struct {
	// ChainID: address
	Tokens map[int][]string `yaml:"tokens"`
}
