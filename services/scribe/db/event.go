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
	// StoreLogs stores a log
	StoreLogs(ctx context.Context, chainID uint32, log ...types.Log) error
	// StoreLogsAtHead stores a log at the tip.
	StoreLogsAtHead(ctx context.Context, chainID uint32, log ...types.Log) error
	// ConfirmLogsForBlockHash confirms logs for a given block hash.
	ConfirmLogsForBlockHash(ctx context.Context, chainID uint32, blockHash common.Hash) error
	// DeleteLogsForBlockHash deletes logs with a given block hash.
	DeleteLogsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreReceipt stores a receipt
	StoreReceipt(ctx context.Context, chainID uint32, receipt types.Receipt) error
	// StoreReceiptAtHead stores a receipt to the tip
	StoreReceiptAtHead(ctx context.Context, chainID uint32, receipt types.Receipt) error
	// DeleteReceiptsForBlockHash deletes receipts with a given block hash.
	DeleteReceiptsForBlockHash(ctx context.Context, chainID uint32, blockHash common.Hash) error

	// StoreEthTx stores a processed transaction
	StoreEthTx(ctx context.Context, tx *types.Transaction, chainID uint32, blockHash common.Hash, blockNumber uint64, transactionIndex uint64) error
	// StoreEthTxAtHead stores a processed transaction at the tip.
	StoreEthTxAtHead(ctx context.Context, tx *types.Transaction, chainID uint32, blockHash common.Hash, blockNumber uint64, transactionIndex uint64) error
	// ConfirmEthTxsForBlockHash confirms eth txs for a given block hash.
	ConfirmEthTxsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error
	// DeleteEthTxsForBlockHash deletes eth txs with a given block hash.
	DeleteEthTxsForBlockHash(ctx context.Context, blockHash common.Hash, chainID uint32) error

	// StoreLastIndexed stores the last indexed for a contract address
	StoreLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, blockNumber uint64, livefillAtHead bool) error
	// StoreLastIndexedMultiple stores the last indexed block numbers for numerous contracts.
	StoreLastIndexedMultiple(ctx context.Context, contractAddresses []common.Address, chainID uint32, blockNumber uint64) error

	// StoreLastConfirmedBlock stores the last block number that has been confirmed.
	// It updates the value if there is a previous last block confirmed value, and creates a new
	// entry if there is no previous value.
	StoreLastConfirmedBlock(ctx context.Context, chainID uint32, blockNumber uint64) error

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
	// RetrieveLogsInRangeAsc retrieves all logs that match an inputted filter and are within a range given a page - in ascending order.
	RetrieveLogsInRangeAsc(ctx context.Context, logFilter LogFilter, startBlock, endBlock uint64, page int) (logs []*types.Log, err error)

	// RetrieveReceiptsWithFilter retrieves receipts with a filter given a page.
	RetrieveReceiptsWithFilter(ctx context.Context, receiptFilter ReceiptFilter, page int) (receipts []types.Receipt, err error)
	// RetrieveReceiptsInRange retrieves receipts that match an inputted filter and are within a range given a page.
	RetrieveReceiptsInRange(ctx context.Context, receiptFilter ReceiptFilter, startBlock, endBlock uint64, page int) (receipts []types.Receipt, err error)

	// RetrieveEthTxsWithFilter retrieves eth transactions with a filter given a page.
	RetrieveEthTxsWithFilter(ctx context.Context, ethTxFilter EthTxFilter, page int) ([]TxWithBlockNumber, error)
	// RetrieveEthTxsInRange retrieves eth transactions that match an inputted filter and are within a range given a page.
	RetrieveEthTxsInRange(ctx context.Context, ethTxFilter EthTxFilter, startBlock, endBlock uint64, page int) ([]TxWithBlockNumber, error)

	// RetrieveLastIndexed retrieves the last indexed for a contract address
	RetrieveLastIndexed(ctx context.Context, contractAddress common.Address, chainID uint32, livefillAtHead bool) (uint64, error)

	// RetrieveLastIndexedMultiple retrieves the last indexed block numbers for numerous contracts.
	RetrieveLastIndexedMultiple(ctx context.Context, contractAddresses []common.Address, chainID uint32) (map[common.Address]uint64, error)
	// RetrieveLastConfirmedBlock retrieves the last block number that has been confirmed.
	RetrieveLastConfirmedBlock(ctx context.Context, chainID uint32) (uint64, error)

	// RetrieveBlockTime retrieves a block time for a chain and block number.
	RetrieveBlockTime(ctx context.Context, chainID uint32, blockNumber uint64) (uint64, error)
	// RetrieveLastBlockStored retrieves the last block number that has a stored block time.
	RetrieveLastBlockStored(ctx context.Context, chainID uint32) (uint64, error)
	// RetrieveFirstBlockStored retrieves the first block number that has a stored block time.
	RetrieveFirstBlockStored(ctx context.Context, chainID uint32) (uint64, error)
	// RetrieveLogCountForContract retrieves the number of logs for a contract.
	RetrieveLogCountForContract(ctx context.Context, contractAddress common.Address, chainID uint32) (int64, error)
	// RetrieveReceiptCountForChain retrieves the number of receipts for a chain.
	RetrieveReceiptCountForChain(ctx context.Context, chainID uint32) (int64, error)
	// RetrieveBlockTimesCountForChain retrieves the number of block times stored for a chain.
	RetrieveBlockTimesCountForChain(ctx context.Context, chainID uint32) (int64, error)
	// RetrieveReceiptsWithStaleBlockHash gets receipts that are from a reorged/stale block.
	RetrieveReceiptsWithStaleBlockHash(ctx context.Context, chainID uint32, blockHashes []string, startBlock uint64, endBlock uint64) ([]types.Receipt, error)

	// RetrieveLogsFromHeadRangeQuery gets unconfirmed logs from the head in a range.
	RetrieveLogsFromHeadRangeQuery(ctx context.Context, logFilter LogFilter, startBlock uint64, endBlock uint64, page int) (logs []*types.Log, err error)
	// RetrieveReceiptsFromHeadRangeQuery gets unconfirmed receipts from the head in a range.
	RetrieveReceiptsFromHeadRangeQuery(ctx context.Context, receiptFilter ReceiptFilter, startBlock uint64, endBlock uint64, page int) ([]types.Receipt, error)
	// RetrieveUnconfirmedEthTxsFromHeadRangeQuery retrieves all unconfirmed ethTx for a given chain ID and range.
	RetrieveUnconfirmedEthTxsFromHeadRangeQuery(ctx context.Context, receiptFilter EthTxFilter, startBlock uint64, endBlock uint64, lastIndexed uint64, page int) ([]TxWithBlockNumber, error)

	// FlushFromHeadTables flushes unconfirmed logs, receipts, and txs from the head.
	FlushFromHeadTables(ctx context.Context, time int64) error
}

// EventDB stores events.
//
//go:generate go run github.com/vektra/mockery/v2 --name EventDB --output ./mocks --case=underscore
type EventDB interface {
	EventDBWriter
	EventDBReader
}

// TxWithBlockNumber is a transaction with a block number and is used for specifically for batching data in explorer.
type TxWithBlockNumber struct {
	Tx          types.Transwaction
	BlockNumber uint64
}
