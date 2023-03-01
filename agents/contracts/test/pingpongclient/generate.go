// Package pingpongclient generates abi data for contract PingPongClient.t.sol
package pingpongclient

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol  ../../../../packages/contracts-core/flattened/PingPongClient.sol --pkg pingpongclient --sol-version 0.8.17 --filename pingpongclient
// line after go:generate cannot be left blank
