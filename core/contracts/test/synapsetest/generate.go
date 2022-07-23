package synapsetest

// TODO: this requires a modification to build:go to pull in foundry deps
//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts/flattened/SynapseTest.sol --pkg synapsetest --sol-version 0.8.13 --filename synapsetest
// line after go:generate cannot be left blank
