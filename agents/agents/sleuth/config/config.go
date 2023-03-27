package config

// Config is the configuration for the Sleuth agent.
type Config struct {
	// Chains stores all chain information
	Chains ChainConfigs `yaml:"chains"`
	// SummitChainID is the chain ID of the chain that the summit contract is deployed on.
	SummitChainID uint32 `yaml:"summit_chain_id"`
	// SummitAddress is the address of the summit contract.
	SummitAddress string `yaml:"summit_address"`
	// BaseOmnirpcURL is the base url for omnirpc.
	// The format is "https://omnirpc.url/". Notice the lack of "confirmations" on the URL
	// in comparison to what `Scribe` uses.
	BaseOmnirpcURL string `yaml:"base_omnirpc_url"`
}
