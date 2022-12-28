package config

// Config is the configuration for the release-monitor service.
type Config struct {
	// name->chainConfig
	Chains map[string]ChainConfig `json:"chains"`
	// GithubAPIKey is the API key for the Github API used to bypass rate limits
	GithubAPIKey string `json:"github_api_key"`
}

// ChainConfig is the config for a chain.
type ChainConfig struct {
	// ChainID is the chain id.
	ChainId uint64 `json:"chainId"`
	// URL is the github url of the release to monitor.
	URL string `json:"url"`
}
