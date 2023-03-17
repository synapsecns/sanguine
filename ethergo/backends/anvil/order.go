package anvil

// Order is the Order of transactions to use.
// see: https://github.com/foundry-rs/foundry/blob/master/anvil/src/eth/pool/transactions.rs#L56
//
//go:generate go run golang.org/x/tools/cmd/stringer -type=Order -linecomment
type Order uint8

const (
	// Fees orders transactions by fees in the mempool.
	Fees Order = iota + 1 // fees
	// Fifo orders transactions by first in first out in the mempool ignoring fees.
	Fifo // fifo
)

var allOrders = []Order{
	Fees,
	Fifo,
}

func init() {
	if len(_Order_index)-1 != len(allOrders) {
		panic("not all orders have been added to the stringer")
	}
}
