// Package lightmanagerharness generates abi data for contract LightManagerHarness.t.sol
package lightmanagerharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/LightManagerHarness.t.sol --pkg lightmanagerharness --sol-version 0.8.17 --filename lightmanagerharness
// line after go:generate cannot be left blank
