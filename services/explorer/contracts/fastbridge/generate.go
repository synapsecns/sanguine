// Package fastbridge Go interface for synapse-contracts/.../SynapseRFQV2.sol
package fastbridge

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../contracts/SynapseRFQV2_flat.sol --pkg fastbridge --sol-version 0.8.20 --filename fastbridge --evm-version istanbul

// ignore this line: go:generate cannot be the last line of a file
