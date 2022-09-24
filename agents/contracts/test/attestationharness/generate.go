// Package attestationharness generates abi data for contract AttestationHarness.sol
package attestationharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/AttestationHarness.sol --pkg attestationharness --sol-version 0.8.13 --filename attestationharness
// line after go:generate cannot be left blank
