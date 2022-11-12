// Package messaging Go interface for synapse-contracts/.../MessageBusUpgradeable.sol
package messagebus

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../contracts/MessageBusUpgradeableV1_flat.sol --pkg messaging --sol-version 0.8.0 --filename messaging

// ignore this line: go:generate cannot be the last line of a file
