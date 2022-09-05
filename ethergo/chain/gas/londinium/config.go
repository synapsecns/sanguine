package londinium

import "github.com/ethereum/go-ethereum/eth/gasprice"

// ToLondiniumConfig converts a post-london gasprice.Config to a legacy confiy.
func ToLondiniumConfig(newConfig gasprice.Config) Config {
	return Config{
		Blocks:     newConfig.Blocks,
		Percentile: newConfig.Percentile,
		Default:    newConfig.Default,
		MaxPrice:   newConfig.MaxPrice,
	}
}
