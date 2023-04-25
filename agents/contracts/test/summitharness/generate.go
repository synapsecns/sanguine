// Package summitharness generates abi data for contract SummitHarness.t.sol
package summitharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/SummitHarness.t.sol --pkg summitharness --sol-version 0.8.17 --filename summitharness --optimizer-runs 1
// line after go:generate cannot be left blank
