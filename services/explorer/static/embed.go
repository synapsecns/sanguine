// Package static contains static config files
package static

import _ "embed" // used for embedding coingecko files

//go:embed tokenIDToCoinGeckoID.yaml
var tokenIDToCoingeckoMap []byte

//go:embed tokenSymbolToCoinGeckoID.yaml
var tokenSymbolToCoingeckoMap []byte

// GetTokenIDToCoingekoConfig returns the tokenID yaml files.
func GetTokenIDToCoingekoConfig() []byte {
	return tokenIDToCoingeckoMap
}

// GetTokenSymbolToCoingeckoConfig returns the token Symbol yaml files.
func GetTokenSymbolToCoingeckoConfig() []byte {
	return tokenSymbolToCoingeckoMap
}
