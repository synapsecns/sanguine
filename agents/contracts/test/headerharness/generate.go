// Package headerharness is for the HeaderHarness
package headerharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/HeaderHarness.t.sol --pkg headerharness --sol-version 0.8.17 --filename headerharness
// line after go:generate cannot be left blank
