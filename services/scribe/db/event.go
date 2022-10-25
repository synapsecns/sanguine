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
	// ConfirmLogsForBlockHash confirms logs for a given block hash.
	ConfirmLogsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error
	// ConfirmLogsInRange confirms logs in a range.
	ConfirmLogsInRange(ctx context.Context, startBlock, endBlock uint64, chainID uint32) error
	// DeleteLogsForBlockHash deletes logs with a given block hash.
	DeleteLogsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, receipt types.Receipt, chainID uint32) error
	// ConfirmReceiptsForBlockHash confirms receipts for a given block hash.
	ConfirmReceiptsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error
	// ConfirmReceiptsInRange confirms receipts in a range.
	ConfirmReceiptsInRange(ctx context.Context, startBlock, endBlock uint64, chainID uint32) error
	// DeleteReceiptsForBlockHash deletes receipts with a given block hash.
	DeleteReceiptsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreEthTx stores a processed transaction
	StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32, blockHash common.Hash, blockNumber uint64, transactionIndex uint64) error
	// ConfirmEthTxsForBlockHash confirms eth txs for a given block hash.
	ConfirmEthTxsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error
	// ConfirmEthTxsInRange confirms eth txs in a range.
	ConfirmEthTxsInRange(ctx context.Context, startBlock, endBlock uint64, chainID uint32) error
	// DeleteEthTxsForBlockHash deletes eth txs with a given block hash.
	DeleteEthTxsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64) error

	// StoreLastConfirmedBlock stores the last block number that has been confirmed.
	// It updates the value if there is a previous last block confirmed value, and creates a new
	// entry if there is no previous value.
	StoreLastConfirmedBlock(ctx context.Context, chainID uint32, blockNumber uint64) error

	// StoreLastBlockTime stores the last block time stored for a chain.
	// It updates the value if there is a previous last indexed value, and creates a new
	// entry if there is no previous value.
	StoreLastBlockTime(ctx context.Context, chainID uint32, blockNumber uint64) error
	// StoreBlockTime stores a block time for a chain.
	StoreBlockTime(ctx context.Context, chainID uint32, blockNumber, timestamp uint64) error
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

	// RetrieveLastBlockTime retrieves the last block time stored for a chain.
	RetrieveLastBlockTime(ctx context.Context, chainID uint32) (uint64, error)
	// RetrieveBlockTime retrieves a block time for a chain and block number.
	RetrieveBlockTime(ctx context.Context, chainID uint32, blockNumber uint64) (uint64, error)
	// RetrieveLastBlockStored retrieves the last block number that has a stored block time.
	RetrieveLastBlockStored(ctx context.Context, chainID uint32) (uint64, error)
	// RetrieveFirstBlockStored retrieves the first block number that has a stored block time.
	RetrieveFirstBlockStored(ctx context.Context, chainID uint32) (uint64, error)
	// RetrieveLogCountForContract retrieves the number of logs for a contract.
	RetrieveLogCountForContract(ctx context.Context, contractAddress common.Address, chainID uint32) (int64, error)
}

// EventDB stores events.
//
//go:generate go run github.com/vektra/mockery/v2 --name EventDB --output ./mocks --case=underscore
type EventDB interface {
	EventDBWriter
	EventDBReader
}
