package relconfig

// TODO: validation function
type Config struct {
	// ChainID: address
	Tokens map[int][]string `yaml:"tokens"`
	// ChainID: bridge
	Bridges    map[int]string `yaml:"bridges"`
	OmnirpcURL string         `yaml:"omnirpc_url"`
}
