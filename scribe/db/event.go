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
	// StoreEthTx stores a processed transaction
	StoreEthTx(ctx context.Context, tx *types.Transaction) error
	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error
	// RetrieveLastIndexed retrieves the last indexed for a contract address
	RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32) (uint64, error)
}
