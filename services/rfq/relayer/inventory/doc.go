// Package inventory contains the inventory manager.
// the inventory manager is one of the core pieces of the program
// which manages the currently available inventory.
//
// "All" (committed & uncommitted inventory) is considered to be the current balance on all chains
// (so, you if you iterate through all erc-20's and call balanceOf()) that is the "All" inventory
// "Commitable" inventory is "All" inventory net of current commitments. Commitments occur when a user bridge
// request comes in and the relayer internally commits to filling it. This prevents overcommiting erc-20 balances.
package inventory
