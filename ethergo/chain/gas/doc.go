// Package gas handles generating deterministic gas price estimates.
// because our implementation requires each client to arrive at the same.
// digest for each transaction to sign, and gas is a factor in generating that transaction hash.
// we can't just call SuggestPrice(). This module handles deterministic gas price and will (TODO!) in the future
// handle gas limit determinations.
package gas
