package inventory

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
)

type InventoryManager interface{}

type inventoryManagerImpl struct {
	tokens map[int][]common.Address
}

// NewInventoryManager creates a list of tokens we should use.
func NewInventoryManager(ctx context.Context, tokens map[int][]common.Address) {

}
