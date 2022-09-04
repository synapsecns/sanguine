package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/services/scribe/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/scribe/server/graph/resolver"
)

func (r *queryResolver) Logs(ctx context.Context, contractAddress *string, chainID int, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int) ([]*model.Log, error) {
	logsFilter := r.buildLogFilter(contractAddress, blockNumber, txHash, txIndex, blockHash, index)
	logsFilter.ChainID = uint32(chainID)
	logs, err := r.DB.RetrieveLogsWithFilter(ctx, logsFilter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}
	return r.logsToModelLogs(logs, logsFilter.ChainID), nil
}

func (r *queryResolver) LogsRange(ctx context.Context, contractAddress *string, chainID int, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int, startBlock int, endBlock int) ([]*model.Log, error) {
	logsFilter := r.buildLogFilter(contractAddress, blockNumber, txHash, txIndex, blockHash, index)
	logsFilter.ChainID = uint32(chainID)
	logs, err := r.DB.RetrieveLogsInRange(ctx, logsFilter, uint64(startBlock), uint64(endBlock))
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}
	return r.logsToModelLogs(logs, logsFilter.ChainID), nil
}

func (r *queryResolver) LogsFromAddress(ctx context.Context, contractAddress string, chainID int) ([]*model.Log, error) {
	logs, err := r.DB.RetrieveLogsByContractAddress(ctx, common.HexToAddress(contractAddress), uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}
	return r.logsToModelLogs(logs, uint32(chainID)), nil
}

func (r *queryResolver) Receipts(ctx context.Context, chainID int, txHash *string, contractAddress *string, blockHash *string, blockNumber *int, txIndex *int) ([]*model.Receipt, error) {
	receiptsFilter := r.buildReceiptFilter(txHash, contractAddress, blockHash, blockNumber, txIndex)
	receiptsFilter.ChainID = uint32(chainID)
	receipts, err := r.DB.RetrieveReceiptsWithFilter(ctx, receiptsFilter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	return r.receiptsToModelReceipts(receipts, receiptsFilter.ChainID), nil
}

func (r *queryResolver) ReceiptsRange(ctx context.Context, chainID int, txHash *string, contractAddress *string, blockHash *string, blockNumber *int, txIndex *int, startBlock int, endBlock int) ([]*model.Receipt, error) {
	receiptsFilter := r.buildReceiptFilter(txHash, contractAddress, blockHash, blockNumber, txIndex)
	receiptsFilter.ChainID = uint32(chainID)
	receipts, err := r.DB.RetrieveReceiptsInRange(ctx, receiptsFilter, uint64(startBlock), uint64(endBlock))
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	return r.receiptsToModelReceipts(receipts, receiptsFilter.ChainID), nil
}

func (r *queryResolver) LogsFromTxHash(ctx context.Context, transactionHash string, chainID int) ([]*model.Log, error) {
	logs, err := r.DB.RetrieveLogsByTxHash(ctx, common.HexToHash(transactionHash), uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}
	return r.logsToModelLogs(logs, uint32(chainID)), nil
}

func (r *queryResolver) ReceiptsFromAddress(ctx context.Context, contractAddress string, chainID int) ([]*model.Receipt, error) {
	receipts, err := r.DB.RetrieveReceiptsByContractAddress(ctx, common.HexToAddress(contractAddress), uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	return r.receiptsToModelReceipts(receipts, uint32(chainID)), nil
}

func (r *queryResolver) ReceiptFromTxHash(ctx context.Context, txHash string, chainID int) (*model.Receipt, error) {
	receipt, err := r.DB.RetrieveReceiptByTxHash(ctx, common.HexToHash(txHash), uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipt: %w", err)
	}
	return r.receiptToModelReceipt(receipt, uint32(chainID)), nil
}

func (r *queryResolver) Transactions(ctx context.Context, txHash *string, chainID int, blockNumber *int) ([]*model.Transaction, error) {
	transactionsFilter := r.buildEthTxFilter(txHash, blockNumber)
	transactionsFilter.ChainID = uint32(chainID)
	transactions, err := r.DB.RetrieveEthTxsWithFilter(ctx, transactionsFilter)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}
	return r.ethTxsToModelTransactions(transactions, transactionsFilter.ChainID), nil
}

func (r *queryResolver) TransactionsRange(ctx context.Context, txHash *string, chainID int, blockNumber *int, startBlock int, endBlock int) ([]*model.Transaction, error) {
	transactionsFilter := r.buildEthTxFilter(txHash, blockNumber)
	transactionsFilter.ChainID = uint32(chainID)
	transactions, err := r.DB.RetrieveEthTxsInRange(ctx, transactionsFilter, uint64(startBlock), uint64(endBlock))
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}
	return r.ethTxsToModelTransactions(transactions, transactionsFilter.ChainID), nil
}

func (r *queryResolver) TransactionFromTxHash(ctx context.Context, txHash string, chainID int) (*model.Transaction, error) {
	tx, err := r.DB.RetrieveEthTxByTxHash(ctx, txHash, uint32(chainID))
	if err != nil {
		return nil, fmt.Errorf("error retrieving transaction: %w", err)
	}
	return r.ethTxToModelTransaction(tx, uint32(chainID)), nil
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
