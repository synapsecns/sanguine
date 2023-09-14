// Package receiptharness generates abi data for contract ReceiptHarness.t.sol
package receiptharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/ReceiptHarness.t.sol --pkg receiptharness --sol-version 0.8.17 --filename receiptharness
// line after go:generate cannot be left blank
