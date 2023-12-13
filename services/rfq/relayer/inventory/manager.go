package inventory

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// fuck the implementation. What we actaully want to be able to do here is
type InventoryManager interface {
	// GetCommittableBalance gets the total balance available for quotes
	// this does not include on-chain balances commited in previous quotes that may be
	// refunded in the event of a revert.
	GetCommittableBalance(ctx context.Context, chainID int, token common.Address)
	// GetCommitableBalances gets the total balances commitable for all tracked tokens.
	GetCommitableBalances(ctx context.Context) map[int]map[common.Address]*big.Int
}

type inventoryManagerImpl struct {
	// map chainID->address->
	tokens map[int]map[common.Address]*big.Int
}

// NewInventoryManager creates a list of tokens we should use.
func NewInventoryManager(ctx context.Context, tokens map[int][]common.Address) {

}

// Ultimately this should produce a list of all balances and remove the
// quoted amounts from the database
