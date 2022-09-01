package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/synapsecns/sanguine/services/scribe/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/scribe/server/graph/resolver"
)

func (r *queryResolver) Logs(ctx context.Context, address *string) ([]*model.Log, error) {
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

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
