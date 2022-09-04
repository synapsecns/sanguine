package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EventDBWriter interface {
	// StoreLog stores a log
	StoreLog(ctx context.Context, log types.Log, chainID uint32) error
	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, receipt types.Receipt, chainID uint32) error
	// StoreEthTx stores a processed transaction
	StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32, blockNumber uint64) error
	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error
}

type EventDBReader interface {
	// RetrieveLogsWithFilter retrieves all logs that match a filter.
	RetrieveLogsWithFilter(ctx context.Context, logFilter LogFilter) (logs []*types.Log, err error)
	// RetrieveLogsInRange retrieves all logs that match an inputted filter, and are within a range.
	RetrieveLogsInRange(ctx context.Context, logFilter LogFilter, startBlock, endBlock uint64) (logs []*types.Log, err error)
	// RetrieveLogsByTxHash retrieves logs that match a tx hash and chain id
	RetrieveLogsByTxHash(ctx context.Context, txHash common.Hash, chainID uint32) (logs []*types.Log, err error)
	// RetrieveLogsByContractAddress retrieves all logs that match a contract address and chain id.
	RetrieveLogsByContractAddress(ctx context.Context, contractAddress common.Address, chainID uint32) (logs []*types.Log, err error)
	// UnsafeRetrieveAllLogs retrieves all logs in the database. When `specific` is true, you can specify
	// a chainID and contract address to specifically search for. This is only used for testing.
	UnsafeRetrieveAllLogs(ctx context.Context, specific bool, chainID uint32, address common.Address) (logs []*types.Log, err error)
	// RetrieveReceiptsInRange retrieves receipts in a range.
	RetrieveReceiptsInRange(ctx context.Context, receiptFilter ReceiptFilter, startBlock, endBlock uint64) (receipts []types.Receipt, err error)
	// RetrieveReceiptsWithFilter retrieves receipts with a filter.
	RetrieveReceiptsWithFilter(ctx context.Context, receiptFilter ReceiptFilter) (receipts []types.Receipt, err error)
	// RetrieveReceiptByTxHash retrieves a receipt by tx hash and chain id
	RetrieveReceiptByTxHash(ctx context.Context, txHash common.Hash, chainID uint32) (receipt types.Receipt, err error)
	// RetrieveReceiptsByContractAddress retrieves all receipts with a given contract address and chain id.
	RetrieveReceiptsByContractAddress(ctx context.Context, contractAddress common.Address, chainID uint32) (receipts []types.Receipt, err error)
	// UnsafeRetrieveAllReceipts retrieves all receipts in the database. When `specific` is true, you can specify
	// a chainID to specifically search for. This is only used for testing.
	UnsafeRetrieveAllReceipts(ctx context.Context, specific bool, chainID uint32) (receipts []*types.Receipt, err error)
	// RetrieveEthTxsWithFilter retrieves eth transactions with a filter.
	RetrieveEthTxsWithFilter(ctx context.Context, ethTxFilter EthTxFilter) ([]types.Transaction, error)
	// RetrieveEthTxsInRange retrieves eth transactions in a range.
	RetrieveEthTxsInRange(ctx context.Context, ethTxFilter EthTxFilter, startBlock, endBlock uint64) ([]types.Transaction, error)
	// RetrieveEthTxByTxHash retrieves a processed transaction by tx hash and chain id.
	RetrieveEthTxByTxHash(ctx context.Context, txHash string, chainID uint32) (types.Transaction, error)
	// RetrieveLastIndexed retrieves the last indexed for a contract address
	RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32) (uint64, error)
}

// EventDB stores events.
//
//go:generate go run github.com/vektra/mockery/v2 --name EventDB --output ./mocks --case=underscore
type EventDB interface {
	EventDBWriter
	EventDBReader
}
