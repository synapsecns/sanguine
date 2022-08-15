package messageharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts/flattened/MessageHarness.sol --pkg messageharness --sol-version 0.8.13 --filename messageharness
// line after go:generate cannot be left blank
