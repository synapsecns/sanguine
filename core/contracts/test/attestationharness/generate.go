package attestationharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts/flattened/AttestationHarness.sol --pkg attestationharness --sol-version 0.8.13 --filename attestationharness
// line after go:generate cannot be left blank
