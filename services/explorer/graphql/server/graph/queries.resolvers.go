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
	// If no search parameters are provided, throw an error.
	if chainID == nil && address == nil && txnHash == nil && kappa == nil {
		return nil, fmt.Errorf("must provide at least one of chainID, address, txnHash, or kappa")
	}
	// Deal with potentially nil parameters.
	if page == nil {
		page = new(int)
		*page = 1
	}
	if includePending == nil {
		includePending = new(bool)
		*includePending = true
	}
	if *includePending == false && kappa != nil {
		return nil, fmt.Errorf("cannot filter by kappa without including pending transactions")
	}
	var chainIDRef *uint32
	if chainID != nil {
		tmp := uint32(*chainID)
		chainIDRef = &tmp
	}

	var err error
	var results []*model.BridgeTransaction
	if txnHash != nil {
		// If we are given a transaction hash, we search for the bridge transaction on the origin chain, then locate
		// its counterpart on the destination chain using the kappa (the keccak256 hash of the transaction hash).
		fromInfos, err := r.DB.BridgeEventsFromIdentifiers(ctx, chainIDRef, address, tokenAddress, nil, txnHash, *page)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		results, err = r.originToDestinationBridge(ctx, address, txnHash, kappa, includePending, page, tokenAddress, fromInfos)
	} else if kappa != nil {
		// If we are given a kappa, we search for the bridge transaction on the destination chain, then locate
		// its counterpart on the origin chain using a query to find a transaction hash given a kappa.
		toInfos, err := r.DB.BridgeEventsFromIdentifiers(ctx, chainIDRef, address, tokenAddress, kappa, nil, *page)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		results, err = r.destinationToOriginBridge(ctx, address, txnHash, kappa, page, tokenAddress, toInfos)
	} else {
		// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
		// both the origin -> destination transactions that match the search parameters, and the destination -> origin
		// transactions that match the search parameters. Then we need to merge the results and remove duplicates.
		results, err = r.originOrDestinationBridge(ctx, chainIDRef, address, txnHash, kappa, includePending, page, tokenAddress)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}
	return results, nil
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
