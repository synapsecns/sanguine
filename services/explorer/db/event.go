package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EventDBWriter is an interface for writing events to a database.
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

// EventDBReader is an interface for reading events from a database.
//
//nolint:interfacebloat
type EventDBReader interface {
	// RetrieveLogsWithFilter retrieves all logs that match a filter given a page.
	RetrieveLogsWithFilter(ctx context.Context, logFilter LogFilter, page int) (logs []*types.Log, err error)
	// RetrieveLogsInRange retrieves all logs that match an inputted filter and are within a range given a page.
	RetrieveLogsInRange(ctx context.Context, logFilter LogFilter, startBlock, endBlock uint64, page int) (logs []*types.Log, err error)
	// RetrieveReceiptsWithFilter retrieves receipts with a filter given a page.
	RetrieveReceiptsWithFilter(ctx context.Context, receiptFilter ReceiptFilter, page int) (receipts []types.Receipt, err error)
	// RetrieveReceiptsInRange retrieves receipts that match an inputted filter and are within a range given a page.
	RetrieveReceiptsInRange(ctx context.Context, receiptFilter ReceiptFilter, startBlock, endBlock uint64, page int) (receipts []types.Receipt, err error)
	// RetrieveEthTxsWithFilter retrieves eth transactions with a filter given a page.
	RetrieveEthTxsWithFilter(ctx context.Context, ethTxFilter EthTxFilter, page int) ([]types.Transaction, error)
	// RetrieveEthTxsInRange retrieves eth transactions that match an inputted filter and are within a range given a page.
	RetrieveEthTxsInRange(ctx context.Context, ethTxFilter EthTxFilter, startBlock, endBlock uint64, page int) ([]types.Transaction, error)
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
