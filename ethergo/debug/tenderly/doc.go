// Package tenderly handles interacting with tenderly for easier debugging. This currently exports transactions, but does
// not properly handle source maps. For this to work we're going to have to correctly handle combined json outputs here:
// https://ethereum.stackexchange.com/a/26216 We can't export this functionality because it doesn't exist in go.
// tenderly piggybacks off the underlying provider. This can be accomplished by overriding the abigen to include
// the combined-json source map. While we're at it we should be able to copy the solidity source code into traceutil
// and generate direct stack traces in the console for geth/ganache based tests. This is stubbed in backends/base right now
package tenderly
