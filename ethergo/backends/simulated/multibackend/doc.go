// Package multibackend contains a copy of https://github.com/ethereum/go-ethereum/blob/master/ethclient/simulated/backend.go
// that allows use with multiple chains by exporting new methods.
// Note: As of go-ethereum v1.13.9, the simulated backend has been moved from accounts/abi/bind/backends to ethclient/simulated,
// but a backward compatibility wrapper remains in the old location, which is what we're using here.
package multibackend
