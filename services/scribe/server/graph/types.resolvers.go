package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/synapsecns/sanguine/services/scribe/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/scribe/server/graph/resolver"
)

func (r *logResolver) Transaction(ctx context.Context, obj *model.Log) (*model.Transaction, error) {
	// eventdb.transactionbyhash(obj.TxHash)
	res := model.TxTypeLegacyTx
	return &model.Transaction{
		Type:      &res,
		Data:      "asdfads",
		Gas:       4,
		GasPrice:  1,
		GasTipCap: "adsfds",
		GasFeeCap: "asdfds",
		Value:     "asdfa",
		Nonce:     0,
		To:        "",
	}, nil
}

func (r *transactionResolver) Logs(ctx context.Context, obj *model.Transaction) ([]*model.Log, error) {
	// TODO: make sure address is not nil
	return []*model.Log{
		{
			ID:          "1",
			BlockHash:   "dsa",
			BlockNumber: 131,
			TxHash:      "x",
			Transaction: nil,
		},
	}, nil
}

// Log returns resolvers.LogResolver implementation.
func (r *Resolver) Log() resolvers.LogResolver { return &logResolver{r} }

// Transaction returns resolvers.TransactionResolver implementation.
func (r *Resolver) Transaction() resolvers.TransactionResolver { return &transactionResolver{r} }

type logResolver struct{ *Resolver }
type transactionResolver struct{ *Resolver }
