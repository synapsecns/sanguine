// Package static contains static config files
package static

import _ "embed" // used for embedding coingecko files

//go:embed tokenIDToCoinGeckoID.yaml
var tokenIDToCoingeckoMap []byte

//go:embed tokenSymbolToCoinGeckoID.yaml
var tokenSymbolToCoingeckoMap []byte

//go:embed tokenSymbolToTokenID.yaml
var tokenSymbolToTokenIDMap []byte

// GetTokenIDToCoingekoConfig returns the tokenID yaml files.
func GetTokenIDToCoingekoConfig() []byte {
	return tokenIDToCoingeckoMap
}

// GetTokenSymbolToCoingeckoConfig returns the token Symbol yaml files.
func GetTokenSymbolToCoingeckoConfig() []byte {
	return tokenSymbolToCoingeckoMap
}

// GetTokenSymbolToTokenIDConfig returns the token Symbol to token ID yaml files.
func GetTokenSymbolToTokenIDConfig() []byte {
	return tokenSymbolToTokenIDMap
}
