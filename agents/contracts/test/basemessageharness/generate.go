// Package basemessageharness generates abi data for contract BaseMessageHarness.t.sol
package basemessageharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/BaseMessageHarness.t.sol --pkg basemessageharness --sol-version 0.8.17 --filename basemessageharness
// line after go:generate cannot be left blank
