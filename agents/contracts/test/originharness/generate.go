package originharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts/flattened/OriginHarness.sol --pkg originharness --sol-version 0.8.13 --filename originharness
// line after go:generate cannot be left blank
