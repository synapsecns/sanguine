package gasdataharness

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/GasDataHarness.t.sol --pkg gasdataharness --sol-version 0.8.17 --filename gasdataharness
// line after go:generate cannot be left blank
