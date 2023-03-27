package sleuth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// Backend is the backend for the sleuth.
type Backend interface {
	// TransactionReceipt returns the receipt of a transaction by transaction hash.
	// Note that the receipt is not available for pending transactions.
	TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error)
}
