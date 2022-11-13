// Package testmessagebus generates test abis for synapse-contracts/.../MessageBusUpgradeable.sol
package testmessagebus

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../../contracts/testcontracts/TestMessageBusUpgradeable.sol --pkg testmessagebus --sol-version 0.8.0 --filename testmessagebus
// ignore this line: go:generate cannot be the last line of a file
