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
func (r *queryResolver) BridgeTransactions(ctx context.Context, chainID *int, address *string, txnHash *string, kappa *string, includePending bool, page int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	// If no search parameters are provided, throw an error.
	if chainID == nil && address == nil && txnHash == nil && kappa == nil {
		return nil, fmt.Errorf("must provide at least one of chainID, address, txnHash, or kappa")
	}

	if !includePending && kappa != nil {
		return nil, fmt.Errorf("cannot filter by kappa without including pending transactions")
	}

	var err error
	var results []*model.BridgeTransaction

	switch {
	case txnHash != nil:
		// If we are given a transaction hash, we search for the bridge transaction on the origin chain, then locate
		// its counterpart on the destination chain using the kappa (the keccak256 hash of the transaction hash).
		fromInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(chainID, address, tokenAddress, nil, txnHash, page))
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		results, err = r.originToDestinationBridge(ctx, address, kappa, includePending, page, tokenAddress, fromInfos)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
		}
	case kappa != nil:
		// If we are given a kappa, we search for the bridge transaction on the destination chain, then locate
		// its counterpart on the origin chain using a query to find a transaction hash given a kappa.
		toInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(chainID, address, tokenAddress, kappa, nil, page))
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		results, err = r.destinationToOriginBridge(ctx, address, txnHash, kappa, page, tokenAddress, toInfos)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
		}
	default:
		// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
		// both the origin -> destination transactions that match the search parameters, and the destination -> origin
		// transactions that match the search parameters. Then we need to merge the results and remove duplicates.
		results, err = r.originOrDestinationBridge(ctx, chainID, address, txnHash, kappa, includePending, page, tokenAddress)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}
	return results, nil
}

// LatestBridgeTransactions is the resolver for the latestBridgeTransactions field.
func (r *queryResolver) LatestBridgeTransactions(ctx context.Context, includePending bool, page int) ([]*model.BridgeTransaction, error) {
	chainIDs, err := r.getChainIDs(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain IDs: %w", err)
	}
	// For each chain ID, get the latest bridge transaction.
	var results []*model.BridgeTransaction
	for i := range chainIDs {
		// Get the PartialInfo for the latest bridge transaction.
		fromInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(&chainIDs[i], nil, nil, nil, nil, page))
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		if len(fromInfos) != 0 {
			// Take the fromInfo from the latest bridge transaction and use it to get the bridge transaction.
			bridgeTxn, err := r.originToDestinationBridge(ctx, nil, nil, includePending, page, nil, fromInfos)
			if err != nil {
				return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
			}
			if len(bridgeTxn) != 0 {
				results = append(results, bridgeTxn[0])
			}
		}
	}

	return results, nil
}

// BridgeAmountStatistic is the resolver for the bridgeAmountStatistic field.
func (r *queryResolver) BridgeAmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, chainID *int, address *string, tokenAddress *string) (*model.ValueResult, error) {
	var err error
	var blockNumberFilter string
	var chainIDFilter string
	subQuery := "bridge_events"
	firstFilter := true

	switch *duration {
	case model.DurationPastDay:
		hours := 24
		targetTime := r.getTargetTime(&hours)
		if chainID == nil {
			subQuery, err = r.generateSubQuery(ctx, targetTime, sql.TokenFieldName, sql.ContractAddressFieldName)
			if err != nil {
				return nil, err
			}
		} else {
			startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, *chainID, 0, targetTime)
			if err != nil {
				return nil, fmt.Errorf("failed to get start block number: %w", err)
			}
			chainIDFilter = generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
			blockNumberFilter = fmt.Sprintf("AND %s >= %d", sql.BlockNumberFieldName, startBlock)
		}
	case model.DurationAllTime:
		chainIDFilter = generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
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
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	tokenAddressFilter := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	addressFilter := generateSingleSpecifierStringSQL(address, sql.SenderFieldName, &firstFilter, "")

	additionalFilters := fmt.Sprintf(
		`%s%s%s%s`,
		blockNumberFilter, chainIDFilter, tokenAddressFilter, addressFilter,
	)
	finalSQL := fmt.Sprintf("\nSELECT %s(toUInt256(amount_usd)) FROM %s %s", operation, subQuery, additionalFilters)
	res, err := r.DB.GetFloat64(ctx, finalSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}
	usdValue := fmt.Sprintf("%f", res)
	output := model.ValueResult{
		USDValue: &usdValue,
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
	for i := range chainIDs {
		startBlock, err := r.Fetcher.TimeToBlockNumber(ctx, chainIDs[i], 0, targetTime)
		if err != nil {
			return nil, fmt.Errorf("failed to get start block number: %w", err)
		}
		count, err := r.DB.GetUint64(ctx, generateBridgeEventCountQuery(chainIDs[i], address, nil, directionIn, &startBlock))
		if err != nil {
			return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
		}
		chainInt := chainIDs[i]
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
	chainIDsToTokenAddresses := make(map[int][]string)
	for _, chain := range chainIDs {
		query := fmt.Sprintf(
			`SELECT DISTINCT %s FROM bridge_events WHERE %s = %d OR %s = %d`,
			sql.TokenFieldName, sql.ChainIDFieldName, chain, sql.DestinationChainIDFieldName, chain,
		)
		tokenAddresses, err := r.DB.GetStringArray(ctx, query)
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
			count, err := r.DB.GetUint64(ctx, generateBridgeEventCountQuery(chain, address, &tokenAddresses[i], directionIn, &startBlock))
			if err != nil {
				return nil, fmt.Errorf("failed to get count by token address: %w", err)
			}
			chainInt := chain
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
	targetTime := r.getTargetTime(hours)
	subQuery, err := r.generateSubQuery(ctx, targetTime, sql.TokenFieldName, sql.TxHashFieldName)
	if err != nil {
		return nil, fmt.Errorf("failed to generate subquery: %w", err)
	}
	query := fmt.Sprintf(`SELECT %s AS address, COUNT(DISTINCT %s) AS count FROM %s GROUP BY address ORDER BY count DESC SETTINGS readonly=1`, sql.TokenFieldName, sql.TxHashFieldName, subQuery)
	res, err := r.DB.GetTransactionCountForEveryAddress(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}
	return res, nil
}

// HistoricalStatistics is the resolver for the historicalStatistics field.
func (r *queryResolver) HistoricalStatistics(ctx context.Context, chainID *int, typeArg *model.HistoricalResultType, days *int) (*model.HistoricalResult, error) {
	var subQuery string
	var query string

	// nowTime used for calculating time in the past
	nowTime := time.Now().Unix()
	startTime := nowTime - int64(*days*86400)

	// Create sql segment with filters
	filter := fmt.Sprintf("WHERE %s = %d AND %s >= %d", sql.ChainIDFieldName, *chainID, sql.TimeStampFieldName, startTime)

	// Handle the different logic needed for each query type.
	switch *typeArg {
	case model.HistoricalResultTypeBridgevolume:
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, FROM_UNIXTIME(timestamp, %s) AS date FROM bridge_events %s GROUP BY date ORDER BY total DESC", sql.AmountUSDFieldName, "'%d/%m/%Y'", filter)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s) SETTINGS readonly=1", subQuery)
	case model.HistoricalResultTypeAddresses:
		subQuery = fmt.Sprintf("SELECT uniqExact(%s) AS total, FROM_UNIXTIME(timestamp, %s) AS date FROM bridge_events %s GROUP BY date ORDER BY total DESC", sql.SenderFieldName, "'%d/%m/%Y'", filter)
		query = fmt.Sprintf("SELECT uniqExact(%s) FROM bridge_events %s SETTINGS readonly=1", sql.SenderFieldName, filter)
	case model.HistoricalResultTypeTransactions:
		subQuery = fmt.Sprintf("SELECT uniqExact(%s) AS total, FROM_UNIXTIME(timestamp, %s) AS date FROM bridge_events %s GROUP BY date ORDER BY total DESC", sql.TxHashFieldName, "'%d/%m/%Y'", filter)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s) SETTINGS readonly=1", subQuery)
	default:
		return nil, fmt.Errorf("invalid type argument")
	}
	dayByDayData, err := r.DB.GetDateResults(ctx, subQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get dateResults: %w", err)
	}
	sum, err := r.DB.GetFloat64(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get total sum: %w", err)
	}
	payload := model.HistoricalResult{
		Total:       &sum,
		DateResults: dayByDayData,
		Type:        typeArg,
	}
	return &payload, nil
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
