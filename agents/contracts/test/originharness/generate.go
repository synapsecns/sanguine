// Package originharness generates abi data for contract OriginHarness.t.sol
package originharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/OriginHarness.t.sol --pkg originharness --sol-version 0.8.17 --filename originharness
// line after go:generate cannot be left blank
