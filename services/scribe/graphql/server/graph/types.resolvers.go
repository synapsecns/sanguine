package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/scribe/db"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/scribe/graphql/server/graph/resolver"
	"github.com/synapsecns/sanguine/services/scribe/graphql/server/types"
)

// Transaction is the resolver for the transaction field.
func (r *logResolver) Transaction(ctx context.Context, obj *model.Log) (*model.Transaction, error) {
	transactionFilter := db.EthTxFilter{
		ChainID: uint32(obj.ChainID),
		TxHash:  obj.TxHash,
	}

	transactions, err := r.DB.RetrieveEthTxsWithFilter(ctx, transactionFilter, obj.Page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}
	if len(transactions) == 0 {
		return nil, fmt.Errorf("no transaction found for log")
	}
	if len(transactions) > 1 {
		return nil, fmt.Errorf("multiple transactions found for log")
	}

	return r.ethTxToModelTransaction(transactions[0].Tx, uint32(obj.ChainID)), nil
}

// Receipt is the resolver for the receipt field.
func (r *logResolver) Receipt(ctx context.Context, obj *model.Log) (*model.Receipt, error) {
	receiptFilter := db.ReceiptFilter{
		ChainID: uint32(obj.ChainID),
		TxHash:  obj.TxHash,
	}

	receipts, err := r.DB.RetrieveReceiptsWithFilter(ctx, receiptFilter, obj.Page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	if len(receipts) == 0 {
		return nil, fmt.Errorf("no receipt found for log")
	}
	if len(receipts) > 1 {
		return nil, fmt.Errorf("multiple receipts found for log")
	}

	return r.receiptToModelReceipt(receipts[0], uint32(obj.ChainID)), nil
}

// JSON is the resolver for the json field.
func (r *logResolver) JSON(ctx context.Context, obj *model.Log) (types.JSON, error) {
	json, err := types.UnmarshalJSON(obj)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return json, nil
}

// Logs is the resolver for the logs field.
func (r *receiptResolver) Logs(ctx context.Context, obj *model.Receipt) ([]*model.Log, error) {
	logFilter := db.LogFilter{
		ChainID: uint32(obj.ChainID),
		TxHash:  obj.TxHash,
	}

	logs, err := r.DB.RetrieveLogsWithFilter(ctx, logFilter, obj.Page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}

	return r.logsToModelLogs(logs, uint32(obj.ChainID)), nil
}

// Transaction is the resolver for the transaction field.
func (r *receiptResolver) Transaction(ctx context.Context, obj *model.Receipt) (*model.Transaction, error) {
	transactionFilter := db.EthTxFilter{
		ChainID: uint32(obj.ChainID),
		TxHash:  obj.TxHash,
	}

	transactions, err := r.DB.RetrieveEthTxsWithFilter(ctx, transactionFilter, obj.Page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving transactions: %w", err)
	}
	if len(transactions) == 0 {
		return nil, fmt.Errorf("no transaction found for receipt")
	}
	if len(transactions) > 1 {
		return nil, fmt.Errorf("multiple transactions found for receipt")
	}

	return r.ethTxToModelTransaction(transactions[0].Tx, uint32(obj.ChainID)), nil
}

// JSON is the resolver for the json field.
func (r *receiptResolver) JSON(ctx context.Context, obj *model.Receipt) (types.JSON, error) {
	json, err := types.UnmarshalJSON(obj)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return json, nil
}

// Logs is the resolver for the logs field.
func (r *transactionResolver) Logs(ctx context.Context, obj *model.Transaction) ([]*model.Log, error) {
	logFilter := db.LogFilter{
		ChainID: uint32(obj.ChainID),
		TxHash:  obj.TxHash,
	}

	logs, err := r.DB.RetrieveLogsWithFilter(ctx, logFilter, obj.Page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving logs: %w", err)
	}

	return r.logsToModelLogs(logs, uint32(obj.ChainID)), nil
}

// Receipt is the resolver for the receipt field.
func (r *transactionResolver) Receipt(ctx context.Context, obj *model.Transaction) (*model.Receipt, error) {
	receiptFilter := db.ReceiptFilter{
		ChainID: uint32(obj.ChainID),
		TxHash:  obj.TxHash,
	}

	receipts, err := r.DB.RetrieveReceiptsWithFilter(ctx, receiptFilter, obj.Page)
	if err != nil {
		return nil, fmt.Errorf("error retrieving receipts: %w", err)
	}
	if len(receipts) == 0 {
		return nil, fmt.Errorf("no receipt found for transaction")
	}
	if len(receipts) > 1 {
		return nil, fmt.Errorf("multiple receipts found for transaction")
	}

	return r.receiptToModelReceipt(receipts[0], uint32(obj.ChainID)), nil
}

// JSON is the resolver for the json field.
func (r *transactionResolver) JSON(ctx context.Context, obj *model.Transaction) (types.JSON, error) {
	json, err := types.UnmarshalJSON(obj)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling json: %w", err)
	}

	return json, nil
}

// Log returns resolvers.LogResolver implementation.
func (r *Resolver) Log() resolvers.LogResolver { return &logResolver{r} }

// Receipt returns resolvers.ReceiptResolver implementation.
func (r *Resolver) Receipt() resolvers.ReceiptResolver { return &receiptResolver{r} }

// Transaction returns resolvers.TransactionResolver implementation.
func (r *Resolver) Transaction() resolvers.TransactionResolver { return &transactionResolver{r} }

type logResolver struct{ *Resolver }
type receiptResolver struct{ *Resolver }
type transactionResolver struct{ *Resolver }
