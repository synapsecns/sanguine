package submitter

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/synapsecns/sanguine/ethergo/submitter/config"
)

// GasEstimator is an interface for estimating gas.
type GasEstimator interface {
	EstimateGas(ctx context.Context, call ethereum.CallMsg) (gas uint64, err error)
}

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

func isArbitrumChain(chainID int) bool {
	return chainID == arbitrumChainID || chainID == arbitrumSepoliaChainID
}

// GetGasEstimationMethod returns the gas estimation method to use for the chain.
func GetGasEstimationMethod(cfg config.IConfig, chainID int) GasEstimationMethod {
	if cfg.NativeGasEstimation(chainID) && isArbitrumChain(chainID) {
		return ArbitrumGasEstimation
	}
	return GethGasEstimation
}
