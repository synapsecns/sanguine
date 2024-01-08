package usdt

// See package readme before modifying this go:generate statement
// TODO use new abigen. This will require significant workarounds for sol 4.
// TODO: this is currently disable since a geth upgrade: https://github.com/ethereum/go-ethereum/issues/25604
// go:generate go run github.com/ethereum/go-ethereum/cmd/abigen --sol ../../../external/tether/TetherToken.sol --pkg usdt --out usdt.abigen.go --solc=./tether.sh --alias=_totalSupply=underTotalSupply
