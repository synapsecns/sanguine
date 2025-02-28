package londinium

import (
	"math/big"

	"github.com/ethereum/go-ethereum/eth/gasprice"
)

// Config is a config for the Londinium gas price estimator
type Config struct {
	Blocks     int
	Percentile int
	Default    *big.Int
	MaxPrice   *big.Int
}

// ToLondiniumConfig converts a post-london gasprice.Config to a legacy config.
func ToLondiniumConfig(newConfig gasprice.Config) OracleConfig {
	return OracleConfig{
		Blocks:     newConfig.Blocks,
		Percentile: newConfig.Percentile,
		// In v1.14.8, Default field has been removed from gasprice.Config
		Default:  big.NewInt(0), // Set a default value or get it from elsewhere
		MaxPrice: newConfig.MaxPrice,
	}
}
