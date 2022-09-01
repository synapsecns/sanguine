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
	// UnsafeRetrieveAllLogs retrieves all logs in the database. When `specific` is true, you can specify
	// a chainID and contract address to specifically search for. This is only used for testing.
	UnsafeRetrieveAllLogs(ctx context.Context, specific bool, chainID uint32, address common.Address) (logs []*types.Log, err error)
	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, receipt types.Receipt, chainID uint32) error
	// RetrieveReceipt retrieves a receipt by tx hash and chain id
	RetrieveReceipt(ctx context.Context, txHash common.Hash, chainID uint32) (receipt types.Receipt, err error)
	// UnsafeRetrieveAllReceipts retrieves all receipts in the database. When `specific` is true, you can specify
	// a chainID to specifically search for. This is only used for testing.
	UnsafeRetrieveAllReceipts(ctx context.Context, specific bool, chainID uint32) (receipts []*types.Receipt, err error)
	// StoreEthTx stores a processed transaction
	StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32, blockNumber uint64) error
	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error
	// RetrieveLastIndexed retrieves the last indexed for a contract address
	RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32) (uint64, error)
}
