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
	// RetrieveLogs retrieves logs that match a tx hash and chain id
	RetrieveLogs(ctx context.Context, txHash common.Hash, chainID uint32) (logs []*types.Log, err error)
	// RetrieveAllLogs retrieves all logs in the database
	RetrieveAllLogs(ctx context.Context) (logs []*types.Log, err error)
	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, receipt types.Receipt, chainID uint32) error
	// RetrieveReceipt retrieves a receipt by tx hash and chain id
	RetrieveReceipt(ctx context.Context, txHash common.Hash, chainID uint32) (receipt types.Receipt, err error)
	// RetrieveAllReceipts retrieves all receipts in the database
	RetrieveAllReceipts(ctx context.Context) (receipts []types.Receipt, err error)
	// StoreEthTx stores a processed transaction
	StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32) error
	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error
	// RetrieveLastIndexed retrieves the last indexed for a contract address
	RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32) (uint64, error)
}
