// Package bridge Go interface for synapse-contracts/.../SynapseBridge.sol
package bridge

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../contracts/SynapseBridgeV2_flat.sol --pkg bridge --sol-version 0.6.12 --filename bridge

// ignore this line: go:generate cannot be the last line of a file
