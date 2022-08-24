package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EventDB stores events.
type EventDB interface {
	// StoreLog stores a log
	StoreLog(ctx context.Context, log types.Log, chainID uint32) error
	// RetrieveLogsByTxHash retrieves a log by tx hash
	RetrieveLogsByTxHash(ctx context.Context, txHash common.Hash) (logs []*types.Log, err error)
	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, receipt types.Receipt) error
	// RetrieveReceiptByTxHash retrieves a receipt by tx hash
	RetrieveReceiptByTxHash(ctx context.Context, txHash common.Hash) (receipt types.Receipt, err error)
}
