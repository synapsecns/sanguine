package config

// GasEstimationMethod is the method to use for gas estimation.
type GasEstimationMethod int

const (
	// GethGasEstimation is the default gas estimation method.
	GethGasEstimation GasEstimationMethod = iota + 1
	// ArbitrumGasEstimation is the gas estimation method for Arbitrum.
	ArbitrumGasEstimation
)

func (g GasEstimationMethod) String() string {
	switch g {
	case GethGasEstimation:
		return "geth"
	case ArbitrumGasEstimation:
		return "arbitrum"
	}
	return ""
}

const arbitrumChainID = 42161
const arbitrumSepoliaChainID = 421614

// GetGasEstimationMethod returns the gas estimation method to use for the chain.
func (c *Config) GetGasEstimationMethod(chainID int) GasEstimationMethod {
	if chainID == arbitrumChainID || chainID == arbitrumSepoliaChainID {
		return ArbitrumGasEstimation
	}
	return GethGasEstimation
}
