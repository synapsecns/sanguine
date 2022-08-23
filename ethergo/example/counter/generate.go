package counter

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ./counter.sol --pkg counter --sol-version 0.8.4 --filename counter
// there needs to be a trailing line after go:generate or generation will fail
