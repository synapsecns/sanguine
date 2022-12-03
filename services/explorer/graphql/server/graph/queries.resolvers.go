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

	var err error
	var results []*model.BridgeTransaction

	switch {
	case kappa != nil:
		// If we are given a kappa, we search for the bridge transaction on the destination chain, then locate
		// its counterpart on the origin chain using a query to find a transaction hash given a kappa.
		results, err = r.GetBridgeTxsFromDestination(ctx, chainID, address, txnHash, kappa, *page, tokenAddress)
		if err != nil {
			return nil, err
		}
	default:
		// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
		// both the origin -> destination transactions that match the search parameters, and the destination -> origin
		// transactions that match the search parameters. Then we need to merge the results and remove duplicates.
		fromResults, err := r.GetBridgeTxsFromOrigin(ctx, chainID, address, txnHash, *includePending, *page, tokenAddress, false)
		if err != nil {
			return nil, err
		}
		toResults, err := r.GetBridgeTxsFromDestination(ctx, chainID, address, txnHash, kappa, *page, tokenAddress)
		if err != nil {
			return nil, err
		}
		results = r.mergeBridgeTransactions(fromResults, toResults)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}
	return results, nil
}

// LatestBridgeTransactions is the resolver for the latestBridgeTransactions field.
func (r *queryResolver) LatestBridgeTransactions(ctx context.Context, includePending *bool, page *int) ([]*model.BridgeTransaction, error) {
	// For each chain ID, get the latest bridge transaction.
	var results []*model.BridgeTransaction
	var err error
	results, err = r.GetBridgeTxsFromOrigin(ctx, nil, nil, nil, *includePending, *page, nil, true)

	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}

	return results, nil
}

// BridgeAmountStatistic is the resolver for the bridgeAmountStatistic field.
func (r *queryResolver) BridgeAmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, chainID *int, address *string, tokenAddress *string) (*model.ValueResult, error) {
	var err error
	var timestampSpecifier string
	firstFilter := true

	if *duration == model.DurationPastDay {
		hours := 24
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	}

	var operation string

	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		operation = fmt.Sprintf("AVG(%s)", sql.AmountUSDFieldName)
	case model.StatisticTypeMedianVolumeUsd:
		operation = fmt.Sprintf("median(%s)", sql.AmountUSDFieldName)
	case model.StatisticTypeTotalVolumeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", sql.AmountUSDFieldName)
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("COUNT(DISTINCT %s)", sql.TxHashFieldName)
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("COUNT(DISTINCT %s)", sql.SenderFieldName)
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}

	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.SenderFieldName, &firstFilter, "")
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")

	additionalFilters := fmt.Sprintf(
		`%s%s%s%s`,
		timestampSpecifier, tokenAddressSpecifier, addressSpecifier, chainIDSpecifier,
	)
	finalSQL := fmt.Sprintf("\nSELECT %s FROM bridge_events %s", operation, additionalFilters)
	res, err := r.DB.GetFloat64(ctx, finalSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}

	value := fmt.Sprintf("%f", res)
	fmt.Println(value)
	output := model.ValueResult{
		Value: &value,
	}

	return &output, nil
}

// CountByChainID is the resolver for the countByChainId field.
func (r *queryResolver) CountByChainID(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int) ([]*model.TransactionCountResult, error) {
	directionIn := r.getDirectionIn(direction)
	targetTime := r.getTargetTime(hours)
	results, err := r.DB.GetTxCounts(ctx, generateBridgeEventCountQuery(chainID, address, nil, directionIn, &targetTime, false))
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}
	return results, nil
}

// CountByTokenAddress is the resolver for the countByTokenAddress field.
func (r *queryResolver) CountByTokenAddress(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int) ([]*model.TokenCountResult, error) {
	directionIn := r.getDirectionIn(direction)
	targetTime := r.getTargetTime(hours)
	results, err := r.DB.GetTokenCounts(ctx, generateBridgeEventCountQuery(chainID, address, nil, directionIn, &targetTime, true))
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}

	return results, nil
}

// AddressRanking is the resolver for the addressRanking field.
func (r *queryResolver) AddressRanking(ctx context.Context, hours *int) ([]*model.AddressRanking, error) {
	targetTime := r.getTargetTime(hours)
	firstFilter := true
	timeStampSpecifier := generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	query := fmt.Sprintf(`SELECT %s AS address, COUNT(DISTINCT %s) AS count FROM bridge_events %s GROUP BY address ORDER BY count DESC`, sql.TokenFieldName, sql.TxHashFieldName, timeStampSpecifier)
	res, err := r.DB.GetAddressRanking(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}

	return res, nil
}

// HistoricalStatistics is the resolver for the historicalStatistics field.
func (r *queryResolver) HistoricalStatistics(ctx context.Context, chainID *int, typeArg *model.HistoricalResultType, days *int) (*model.HistoricalResult, error) {
	var subQuery string
	var query string

	startTime := uint64(time.Now().Unix() - int64(*days*86400))
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	timeStampSpecifier := generateTimestampSpecifierSQL(&startTime, sql.TimeStampFieldName, &firstFilter, "")
	filter := fmt.Sprintf("%s%s", chainIDSpecifier, timeStampSpecifier)

	// Handle the different logic needed for each query type.
	switch *typeArg {
	case model.HistoricalResultTypeBridgevolume:
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, FROM_UNIXTIME(%s, %s) AS date FROM bridge_events %s AND %s GROUP BY date ORDER BY date ASC", sql.AmountUSDFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", filter, deDupInQuery)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	case model.HistoricalResultTypeAddresses:
		subQuery = fmt.Sprintf("SELECT uniqExact(%s) AS total, FROM_UNIXTIME( %s, %s) AS date FROM  bridge_events %s AND %s GROUP BY date ORDER BY date ASC", sql.SenderFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", filter, deDupInQuery)
		query = fmt.Sprintf("SELECT uniqExact(%s) FROM bridge_events %s AND %s", sql.SenderFieldName, filter, deDupInQuery)
	case model.HistoricalResultTypeTransactions:
		subQuery = fmt.Sprintf("SELECT uniqExact(%s) AS total, FROM_UNIXTIME(%s, %s) AS date FROM  bridge_events %s AND %s GROUP BY date ORDER BY date ASC", sql.TxHashFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", filter, deDupInQuery)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
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
