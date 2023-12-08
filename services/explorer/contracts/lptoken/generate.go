// Package lptoken Go interface for synapse-contracts/.../LPToken_flat.sol
package lptoken

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../contracts/LPToken_flat.sol --pkg lptoken --sol-version 0.7.0 --filename lptoken

// ignore this line: go:generate cannot be the last line of a file
