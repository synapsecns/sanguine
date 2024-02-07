// Package requestharness generates abi data for contract RequestHarness.t.sol
package requestharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/RequestHarness.t.sol --pkg requestharness --sol-version 0.8.17 --filename requestharness
// line after go:generate cannot be left blank
