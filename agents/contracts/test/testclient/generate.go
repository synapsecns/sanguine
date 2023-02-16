// Package testclient generates abi data for contract TestClient.t.sol
package testclient

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/TestClient.sol --pkg testclient --sol-version 0.8.17 --filename testclient
// line after go:generate cannot be left blank
