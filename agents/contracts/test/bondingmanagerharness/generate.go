// Package bondingmanagerharness generates abi data for contract BondingManagerHarness.t.sol
package bondingmanagerharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/BondingManagerHarness.t.sol --pkg bondingmanagerharness --sol-version 0.8.17 --filename bondingmanagerharness
// line after go:generate cannot be left blank
