package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/scribe/db"

	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/resolver"
)

// Logs is the resolver for the logs field.
func (r *queryResolver) Logs(ctx context.Context, contractAddress *string, chainID int, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int, confirmed *bool, page int) ([]*model.Log, error) {
	logsFilter := db.BuildLogFilter(contractAddress, blockNumber, txHash, txIndex, blockHash, index, confirmed)
	logsFilter.ChainID = uint32(chainID)
	logs, err := r.DB.RetrieveLogsWithFilter(ctx, logsFilter, page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}
	return r.logsToModelLogs(logs, logsFilter.ChainID), nil
}

// LogsRange is the resolver for the logsRange field.
func (r *queryResolver) LogsRange(ctx context.Context, contractAddress *string, chainID int, blockNumber *int, txHash *string, txIndex *int, blockHash *string, index *int, confirmed *bool, startBlock int, endBlock int, page int) ([]*model.Log, error) {
	logsFilter := db.BuildLogFilter(contractAddress, blockNumber, txHash, txIndex, blockHash, index, confirmed)
	logsFilter.ChainID = uint32(chainID)
	logs, err := r.DB.RetrieveLogsInRange(ctx, logsFilter, uint64(startBlock), uint64(endBlock), page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}
	return r.logsToModelLogs(logs, logsFilter.ChainID), nil
}

// Receipts is the resolver for the receipts field.
func (r *queryResolver) Receipts(ctx context.Context, chainID int, txHash *string, contractAddress *string, blockHash *string, blockNumber *int, txIndex *int, confirmed *bool, page int) ([]*model.Receipt, error) {
	receiptsFilter := db.BuildReceiptFilter(txHash, contractAddress, blockHash, blockNumber, txIndex, confirmed)
	receiptsFilter.ChainID = uint32(chainID)
	receipts, err := r.DB.RetrieveReceiptsWithFilter(ctx, receiptsFilter, page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	return r.receiptsToModelReceipts(receipts, receiptsFilter.ChainID), nil
}

// ReceiptsRange is the resolver for the receiptsRange field.
func (r *queryResolver) ReceiptsRange(ctx context.Context, chainID int, txHash *string, contractAddress *string, blockHash *string, blockNumber *int, txIndex *int, confirmed *bool, startBlock int, endBlock int, page int) ([]*model.Receipt, error) {
	receiptsFilter := db.BuildReceiptFilter(txHash, contractAddress, blockHash, blockNumber, txIndex, confirmed)
	receiptsFilter.ChainID = uint32(chainID)
	receipts, err := r.DB.RetrieveReceiptsInRange(ctx, receiptsFilter, uint64(startBlock), uint64(endBlock), page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	return r.receiptsToModelReceipts(receipts, receiptsFilter.ChainID), nil
}

// Transactions is the resolver for the transactions field.
func (r *queryResolver) Transactions(ctx context.Context, txHash *string, chainID int, blockNumber *int, blockHash *string, confirmed *bool, page int) ([]*model.Transaction, error) {
	transactionsFilter := db.BuildEthTxFilter(txHash, blockNumber, blockHash, confirmed)
	transactionsFilter.ChainID = uint32(chainID)
	transactions, err := r.DB.RetrieveEthTxsWithFilter(ctx, transactionsFilter, page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}
	return r.ethTxsToModelTransactions(transactions, transactionsFilter.ChainID), nil
}

// TransactionsRange is the resolver for the transactionsRange field.
func (r *queryResolver) TransactionsRange(ctx context.Context, txHash *string, chainID int, blockNumber *int, blockHash *string, confirmed *bool, startBlock int, endBlock int, page int) ([]*model.Transaction, error) {
	transactionsFilter := db.BuildEthTxFilter(txHash, blockNumber, blockHash, confirmed)
	transactionsFilter.ChainID = uint32(chainID)
	transactions, err := r.DB.RetrieveEthTxsInRange(ctx, transactionsFilter, uint64(startBlock), uint64(endBlock), page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}
	return r.ethTxsToModelTransactions(transactions, transactionsFilter.ChainID), nil
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
