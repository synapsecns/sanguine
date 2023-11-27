// Package erc20 Go interface for synapse-contracts/.../ERC20_flat.sol
package erc20

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ../contracts/ERC20_flat.sol --pkg erc20 --sol-version 0.8.13 --filename erc20

// ignore this line: go:generate cannot be the last line of a file
