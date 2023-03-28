package internal

//go:generate go run github.com/synapsecns/sanguine/tools/abigen generate --sol ./testsignature.sol --pkg internal --sol-version 0.8.4 --filename testsignature
// there needs to be a trailing line after go:generate or generation will fail
