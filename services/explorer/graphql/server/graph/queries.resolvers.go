package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/db/sql"
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
	if !*includePending && kappa != nil {
		return nil, fmt.Errorf("cannot filter by kappa without including pending transactions")
	}
	var chainIDRef *uint32
	if chainID != nil {
		tmp := uint32(*chainID)
		chainIDRef = &tmp
	}

	var err error
	var results []*model.BridgeTransaction

	switch {
	case txnHash != nil:
		// If we are given a transaction hash, we search for the bridge transaction on the origin chain, then locate
		// its counterpart on the destination chain using the kappa (the keccak256 hash of the transaction hash).
		fromInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, chainIDRef, address, tokenAddress, nil, txnHash, *page, false)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		results, err = r.originToDestinationBridge(ctx, address, kappa, includePending, page, tokenAddress, fromInfos, false)
		if err != nil {
			fmt.Println("originToDestinationBridge threw an error")
		}
	case kappa != nil:
		// If we are given a kappa, we search for the bridge transaction on the destination chain, then locate
		// its counterpart on the origin chain using a query to find a transaction hash given a kappa.
		toInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, chainIDRef, address, tokenAddress, kappa, nil, *page, false)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		results, err = r.destinationToOriginBridge(ctx, address, txnHash, kappa, page, tokenAddress, toInfos)
		if err != nil {
			fmt.Println("destinationToOriginBridge threw an error")
		}
	default:
		// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
		// both the origin -> destination transactions that match the search parameters, and the destination -> origin
		// transactions that match the search parameters. Then we need to merge the results and remove duplicates.
		results, err = r.originOrDestinationBridge(ctx, chainIDRef, address, txnHash, kappa, includePending, page, tokenAddress, false)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}
	return results, nil
}

// LatestBridgeTransactions is the resolver for the latestBridgeTransactions field.
func (r *queryResolver) LatestBridgeTransactions(ctx context.Context, includePending *bool, page *int) ([]*model.BridgeTransaction, error) {
	// Deal with potentially nil parameters.
	chainIDs, err := r.getChainIDs(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain IDs: %w", err)
	}
	if page == nil {
		page = new(int)
		*page = 1
	}
	if includePending == nil {
		includePending = new(bool)
		*includePending = true
	}
	// For each chain ID, get the latest bridge transaction.
	var results []*model.BridgeTransaction
	for i := range chainIDs {
		// Get the PartialInfo for the latest bridge transaction.
		fromInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, &chainIDs[i], nil, nil, nil, nil, *page, true)
		if err != nil || len(fromInfos) == 0 {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}

		// Take the fromInfo from the latest bridge transaction and use it to get the bridge transaction.
		bridgeTxn, err := r.originToDestinationBridge(ctx, nil, nil, includePending, page, nil, fromInfos, true)
		if err != nil || len(bridgeTxn) == 0 {
			return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
		}
		results = append(results, bridgeTxn[0])
	}

	return results, nil
}

// BridgeAmountStatistic is the resolver for the bridgeAmountStatistic field.
func (r *queryResolver) BridgeAmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, chainID *int, address *string, tokenAddress *string) (*model.ValueResult, error) {
	getSubQuery := func(targetTime uint64) (string, error) {
		subQuery := "("
		chainIDs, err := r.DB.GetAllChainIDs(ctx)
		if err != nil {
			return subQuery, fmt.Errorf("failed to get chain IDs: %w", err)
		}

		for i, chain := range chainIDs {
			startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, chain, 0, targetTime)
			if err != nil {
				return subQuery, fmt.Errorf("failed to get start block number: %w", err)
			}
			sqlString := fmt.Sprintf("\nSELECT %s, %s, amount_usd FROM bridge_events WHERE %s = %d AND  %s >= %d", sql.TokenFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName, chain, sql.BlockNumberFieldName, startBlock)
			if i != len(chainIDs)-1 {
				sqlString += " UNION ALL"
			} else {
				sqlString += ")"
			}
			subQuery += sqlString
		}
		return subQuery, nil
	}
	var err error
	subQuery := "bridge_events"

	firstFilter := true
	blockNumberFilter := ""
	chainIDFilter := ""
	switch *duration {
	case model.DurationPastDay:
		hours := 24
		targetTime := r.getTargetTime(&hours)
		if chainID == nil {
			subQuery, err = getSubQuery(targetTime)
			if err != nil {
				return nil, err
			}
		} else {
			startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, uint32(*chainID), 0, targetTime)
			if err != nil {
				return nil, fmt.Errorf("failed to get start block number: %w", err)
			}
			chainID32 := uint32(*chainID)
			chainIDFilter = sql.GenerateSingleSpecifierI32SQL(&chainID32, sql.ChainIDFieldName, &firstFilter)
			blockNumberFilter = fmt.Sprintf("AND %s >= %d", sql.BlockNumberFieldName, startBlock)
		}
	case model.DurationAllTime:
		chainID32 := uint32(*chainID)
		chainIDFilter = sql.GenerateSingleSpecifierI32SQL(&chainID32, sql.ChainIDFieldName, &firstFilter)
	}
	var operation string
	switch typeArg {
	case model.StatisticTypeMean:
		operation = "AVG"
	case model.StatisticTypeTotal:
		operation = "sumKahan"
	case model.StatisticTypeMedian:
		operation = "median"
	case model.StatisticTypeCount:
		operation = "COUNT"
	}
	tokenAddressFilter := sql.GenerateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter)
	// TODO double check this
	addressFilter := sql.GenerateSingleSpecifierStringSQL(address, sql.ContractAddressFieldName, &firstFilter)

	additionalFilters := fmt.Sprintf(
		`%s%s%s%s`,
		blockNumberFilter, chainIDFilter, tokenAddressFilter, addressFilter,
	)
	finalSQL := fmt.Sprintf("\nSELECT %s(toUInt256(amount_usd)) FROM %s %s", operation, subQuery, additionalFilters)
	res, err := r.DB.GetBridgeStatistic(ctx, finalSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}
	output := model.ValueResult{
		USDValue: res,
	}
	return &output, nil
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
		for i := range tokenAddresses {
			count, err := r.DB.BridgeEventCount(ctx, chain, address, &tokenAddresses[i], directionIn, startBlock)
			if err != nil {
				return nil, fmt.Errorf("failed to get count by token address: %w", err)
			}
			chainInt := int(chain)
			countInt := int(count)
			results = append(results, &model.TokenCountResult{
				ChainID:      &chainInt,
				TokenAddress: &tokenAddresses[i],
				Count:        &countInt,
			})
		}
	}

	return results, nil
}

// AddressRanking is the resolver for the addressRanking field.
func (r *queryResolver) AddressRanking(ctx context.Context, hours *int) ([]*model.AddressRanking, error) {
	chainIDs, err := r.DB.GetAllChainIDs(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain IDs: %w", err)
	}
	targetTime := r.getTargetTime(hours)

	// this is a generated sql sub query that will allow all addresses across all chains to be queried at once
	// 1. this is the proper way to query this and 2. this will allow us to leverage sql's ORDER BY
	var genSQL string
	for i, chain := range chainIDs {
		startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, chain, 0, targetTime)
		if err != nil {
			return nil, fmt.Errorf("failed to get start block number: %w", err)
		}
		sqlString := fmt.Sprintf("\nSELECT %s, %s FROM bridge_events WHERE %s = %d AND %s >= %d", sql.TokenFieldName, sql.TxHashFieldName, sql.ChainIDFieldName, chain, sql.BlockNumberFieldName, startBlock)
		if i != len(chainIDs)-1 {
			sqlString += " UNION ALL"
		}
		genSQL += sqlString
	}

	res, err := r.DB.GetTransactionCountForEveryAddress(ctx, genSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}
	return res, nil
}

// HistoricalStatistics is the resolver for the historicalStatistics field.
func (r *queryResolver) HistoricalStatistics(ctx context.Context, chainID *int, typeArg *model.HistoricalResultType, days *int) (*model.HistoricalResult, error) {
	var operation string

	// Handle the different logic needed for each query type.
	switch *typeArg {
	case model.HistoricalResultTypeBridgevolume:
		operation = "sumKahan(amount_usd)"
	case model.HistoricalResultTypeAddresses:
		operation = fmt.Sprintf("uniqExact(%s)", sql.SenderFieldName)
	case model.HistoricalResultTypeTransactions:
		operation = fmt.Sprintf("uniqExact(%s)", sql.TxHashFieldName)
	}

	// nowTime used for calculating time in the past
	nowTime := time.Now().Unix()
	startTime := nowTime - int64(*days*86400)

	// Create sql segment with filters
	filter := fmt.Sprintf("WHERE %s = %d AND timestamp >= %d", sql.ChainIDFieldName, *chainID, startTime)

	// Create query for getting day by day data
	subQuery := fmt.Sprintf("SELECT %s AS total, FROM_UNIXTIME(timestamp, %s) AS date FROM bridge_events %s GROUP BY date ORDER BY total DESC", operation, "'%d/%m/%Y'", filter)

	// get data
	res, err := r.DB.GetHistoricalData(ctx, subQuery, typeArg, filter)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}
	return res, nil
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
