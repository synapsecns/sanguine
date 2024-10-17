// Package fastbridgemockv2 is a mock fast bridge contract for testing fast bridge interactions.
package fastbridgemockv2

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../../../../../packages/contracts-rfq/flattened/FastBridgeMock.sol --pkg fastbridgemockv2 --sol-version 0.8.20 --filename fastbridgemockv2 --evm-version istanbul
