// Package tipsharness generates abi data for contract TipsHarness.t.sol
package tipsharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/TipsHarness.t.sol --pkg tipsharness --sol-version 0.8.17 --filename tipsharness
// line after go:generate cannot be left blank
