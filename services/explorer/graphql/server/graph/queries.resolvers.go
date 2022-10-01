package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/resolver"
)

// BridgeTransactions is the resolver for the bridgeTransactions field.
func (r *queryResolver) BridgeTransactions(ctx context.Context, chainID *int, address *string, txnHash *string, kappa *string, includePending *bool, page *int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	panic(fmt.Errorf("not implemented: BridgeTransactions - bridgeTransactions"))
}

// LatestBridgeTransactions is the resolver for the latestBridgeTransactions field.
func (r *queryResolver) LatestBridgeTransactions(ctx context.Context, includePending *bool, page *int) ([]*model.BridgeTransaction, error) {
	panic(fmt.Errorf("not implemented: LatestBridgeTransactions - latestBridgeTransactions"))
}

// BridgeAmountStatistic is the resolver for the bridgeAmountStatistic field.
func (r *queryResolver) BridgeAmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, chainID *int, address *string, tokenAddress *string) (*model.ValueResult, error) {
	panic(fmt.Errorf("not implemented: BridgeAmountStatistic - bridgeAmountStatistic"))
}

// CountByChainID is the resolver for the countByChainId field.
func (r *queryResolver) CountByChainID(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int) ([]*model.TransactionCountResult, error) {
	// panic(fmt.Errorf("not implemented: CountByChainID - countByChainId"))
	cId := 1
	cnt := 2
	m := &model.TransactionCountResult{
		ChainID: &cId,
		Count:   &cnt,
	}
	return []*model.TransactionCountResult{
		m,
	}, nil
}

// CountByTokenAddress is the resolver for the countByTokenAddress field.
func (r *queryResolver) CountByTokenAddress(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int) ([]*model.TokenCountResult, error) {
	panic(fmt.Errorf("not implemented: CountByTokenAddress - countByTokenAddress"))
}

// AddressRanking is the resolver for the addressRanking field.
func (r *queryResolver) AddressRanking(ctx context.Context, hours *int) ([]*model.AddressRanking, error) {
	panic(fmt.Errorf("not implemented: AddressRanking - addressRanking"))
}

// GetCSV is the resolver for the getCsv field.
func (r *queryResolver) GetCSV(ctx context.Context, address string) (*model.CSVData, error) {
	panic(fmt.Errorf("not implemented: GetCSV - getCsv"))
}

// HistoricalStatistics is the resolver for the historicalStatistics field.
func (r *queryResolver) HistoricalStatistics(ctx context.Context, chainID *int, typeArg *model.HistoricalResultType, days *int) (*model.HistoricalResult, error) {
	panic(fmt.Errorf("not implemented: HistoricalStatistics - historicalStatistics"))
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
