package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/resolver"
	"time"
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
	var chainIDs []uint32
	// if the chain ID is not specified, get all chain IDs
	if chainID == nil {
		chainIDsInt, err := r.DB.GetAllChainIDs(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get all chain IDs: %w", err)
		}
		for _, chain := range chainIDsInt {
			chainIDs = append(chainIDs, uint32(chain))
		}
	} else {
		chainIDs = append(chainIDs, uint32(*chainID))
	}
	var directionIn bool
	if direction != nil {
		directionIn = *direction == model.DirectionIn
	} else {
		directionIn = true
	}
	var targetTime uint64
	if hours == nil {
		targetTime = uint64(time.Now().Add(-time.Hour * 24).Unix())
	} else {
		targetTime = uint64(time.Now().Add(-time.Hour * time.Duration(*hours)).Unix())
	}

	// get the number of transactions for each chain ID
	var results []*model.TransactionCountResult
	for _, chain := range chainIDs {
		startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, chain, 0, targetTime)
		if err != nil {
			return nil, fmt.Errorf("failed to get start block number: %w", err)
		}
		count, err := r.DB.BridgeCountByChainID(ctx, chain, address, directionIn, startBlock)
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
	//
	//kill, me := r.Fetcher.FetchClient.GetFirstStoredBlockNumber(ctx, *chainID)
	//if me != nil {
	//	return nil, me
	//}
	//fmt.Println("kill me", kill)
	//
	////// panic(fmt.Errorf("not implemented: CountByChainID - countByChainId"))
	////
	//cId := 1
	//cnt := 2
	//m := &model.TransactionCountResult{
	//	ChainID: &cId,
	//	Count:   &cnt,
	//}
	//return []*model.TransactionCountResult{
	//	m,
	//}, nil
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
