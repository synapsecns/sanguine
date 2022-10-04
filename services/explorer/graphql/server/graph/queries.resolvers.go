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
	chainIDs, err := r.getChainIDs(ctx, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain IDs: %w", err)
	}
	directionIn := r.getDirectionIn(direction)
	targetTime := r.getTargetTime(hours)

	// get the number of transactions for each chain ID
	var results []*model.TransactionCountResult
	for _, chain := range chainIDs {
		startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, chain, 0, targetTime)
		if err != nil {
			return nil, fmt.Errorf("failed to get start block number: %w", err)
		}
		count, err := r.DB.BridgeEventCount(ctx, chain, address, nil, directionIn, startBlock)
		if err != nil {
			return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
		}
		chainInt := int(chain)
		countInt := int(count)
		results = append(results, &model.TransactionCountResult{
			ChainID: &chainInt,
			Count:   &countInt,
		})
	}

	return results, nil
}

// CountByTokenAddress is the resolver for the countByTokenAddress field.
func (r *queryResolver) CountByTokenAddress(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int) ([]*model.TokenCountResult, error) {
	chainIDs, err := r.getChainIDs(ctx, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain IDs: %w", err)
	}
	chainIDsToTokenAddresses := make(map[uint32][]string)
	for _, chain := range chainIDs {
		tokenAddresses, err := r.DB.GetTokenAddressesByChainID(ctx, chain)
		if err != nil {
			return nil, fmt.Errorf("failed to get token addresses: %w", err)
		}
		chainIDsToTokenAddresses[chain] = tokenAddresses
	}
	directionIn := r.getDirectionIn(direction)
	targetTime := r.getTargetTime(hours)

	// get the number of transactions for each token address, for each chain ID
	var results []*model.TokenCountResult
	for chain, tokenAddresses := range chainIDsToTokenAddresses {
		startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, chain, 0, targetTime)
		if err != nil {
			return nil, fmt.Errorf("failed to get start block number: %w", err)
		}
		for _, tokenAddress := range tokenAddresses {
			count, err := r.DB.BridgeEventCount(ctx, chain, address, &tokenAddress, directionIn, startBlock)
			if err != nil {
				return nil, fmt.Errorf("failed to get count by token address: %w", err)
			}
			chainInt := int(chain)
			countInt := int(count)
			results = append(results, &model.TokenCountResult{
				ChainID:      &chainInt,
				TokenAddress: &tokenAddress,
				Count:        &countInt,
			})
		}
	}

	return results, nil
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
