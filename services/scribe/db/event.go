package db

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// EventDBWriter is an interface for writing events to a database.
//
//nolint:interfacebloat
type EventDBWriter interface {
	// StoreLog stores a log
	StoreLog(ctx context.Context, log types.Log, chainID uint32) error
	// ConfirmLog confirms a log.
	ConfirmLog(ctx context.Context, blockHash common.Hash, chainID uint32) error
	// ConfirmLogsInRange confirms logs in a range.
	ConfirmLogsInRange(ctx context.Context, startBlock, endBlock uint64, chainID uint32) error
	// DeleteLogs deletes logs with a given block hash.
	DeleteLogs(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, receipt types.Receipt, chainID uint32) error
	// ConfirmReceipt confirms a receipt.
	ConfirmReceipt(ctx context.Context, blockHash common.Hash, chainID uint32) error
	// ConfirmReceiptsInRange confirms receipts in a range.
	ConfirmReceiptsInRange(ctx context.Context, startBlock, endBlock uint64, chainID uint32) error
	// DeleteReceipts deletes receipts with a given block hash.
	DeleteReceipts(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreEthTx stores a processed transaction
	StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32, blockHash common.Hash, blockNumber uint64) error // ConfirmEthTx confirms an eth tx.
	ConfirmEthTx(ctx context.Context, blockHash common.Hash, chainID uint32) error
	// ConfirmEthTxsInRange confirms eth txs in a range.
	ConfirmEthTxsInRange(ctx context.Context, startBlock, endBlock uint64, chainID uint32) error
	// DeleteEthTxs deletes eth txs with a given block hash.
	DeleteEthTxs(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error

	// StoreLastConfirmedBlock stores the last block number that has been confirmed.
	// It updates the value if there is a previous last block confirmed value, and creates a new
	// entry if there is no previous value.
	StoreLastConfirmedBlock(ctx context.Context, chainID uint32, blockNumber uint64) error
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

	// RetrieveLastConfirmedBlock retrieves the last block number that has been confirmed.
	RetrieveLastConfirmedBlock(ctx context.Context, chainID uint32) (uint64, error)
}

// EventDB stores events.
//
//go:generate go run github.com/vektra/mockery/v2 --name EventDB --output ./mocks --case=underscore
type EventDB interface {
	EventDBWriter
	EventDBReader
}
