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
	"golang.org/x/sync/errgroup"
)

// BridgeTransactions is the resolver for the bridgeTransactions field.
func (r *queryResolver) BridgeTransactions(ctx context.Context, chainID *int, address *string, txnHash *string, kappa *string, includePending *bool, page *int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction

	switch {
	case kappa != nil:
		// If we are given a kappa, we search for the bridge transaction on the destination chain, then locate
		// its counterpart on the origin chain using a query to find a transaction hash given a kappa.
		results, err = r.GetBridgeTxsFromDestination(ctx, nil, address, txnHash, kappa, *page, tokenAddress)
		if err != nil {
			return nil, err
		}
	default:
		// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
		// both the origin -> destination transactions that match the search parameters, and the destination -> origin
		// transactions that match the search parameters. Then we need to merge the results and remove duplicates.
		fromResults, err := r.GetBridgeTxsFromOrigin(ctx, chainID, address, txnHash, *page, tokenAddress, *includePending, false)
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

// BridgeAmountStatistic is the resolver for the bridgeAmountStatistic field.
func (r *queryResolver) BridgeAmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, chainID *int, address *string, tokenAddress *string) (*model.ValueResult, error) {
	var err error
	var timestampSpecifier string
	firstFilter := true

	switch *duration {
	case model.DurationPastDay:
		hours := 24
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationPastMonth:
		hours := 720
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationAllTime:
		timestampSpecifier = ""
	}
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.SenderFieldName, &firstFilter, "")
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	directionSpecifier := generateDirectionSpecifierSQL(true, &firstFilter, "")

	compositeFilters := fmt.Sprintf(
		`%s%s%s%s%s`,
		timestampSpecifier, tokenAddressSpecifier, addressSpecifier, chainIDSpecifier, directionSpecifier,
	)

	var operation string
	var finalSQL string
	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		operation = fmt.Sprintf("AVG(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("\nSELECT %s FROM (SELECT %s FROM (%s) %s)", operation, singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins)
	case model.StatisticTypeMedianVolumeUsd:
		operation = fmt.Sprintf("median(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("\nSELECT %s FROM (SELECT %s FROM (%s) %s)", operation, singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins)
	case model.StatisticTypeTotalVolumeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("\nSELECT %s FROM (SELECT %s FROM (%s) %s)", operation, singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins)
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
		finalSQL = fmt.Sprintf("\nSELECT %s FROM (%s)", operation, generateDeDepQuery(compositeFilters, nil, nil, false))
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
		finalSQL = fmt.Sprintf("\nSELECT %s FROM (%s)", operation, generateDeDepQuery(compositeFilters, nil, nil, false))
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}

	res, err := r.DB.GetFloat64(ctx, finalSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to get amount data stats: %w", err)
	}

	value := fmt.Sprintf("%f", res)
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
	directionSpecifier := generateDirectionSpecifierSQL(true, &firstFilter, "")
	compositeFilters := fmt.Sprintf("%s%s", timeStampSpecifier, directionSpecifier)
	query := fmt.Sprintf(`SELECT %s AS address, COUNT(DISTINCT %s) AS Count FROM (%s) GROUP BY %s ORDER BY Count Desc`, sql.SenderFieldName, sql.TxHashFieldName, generateDeDepQuery(compositeFilters, nil, nil, false), sql.SenderFieldName)
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
	directionSpecifier := generateDirectionSpecifierSQL(true, &firstFilter, "")
	compositeFilters := fmt.Sprintf("%s%s%s", chainIDSpecifier, timeStampSpecifier, directionSpecifier)

	// Handle the different logic needed for each query type.
	switch *typeArg {
	case model.HistoricalResultTypeBridgevolume:
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, FROM_UNIXTIME(%s, %s) AS date FROM (SELECT %s FROM (%s) %s) GROUP BY date ORDER BY date ASC", sql.AmountUSDFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	case model.HistoricalResultTypeAddresses:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, FROM_UNIXTIME(%s, %s) AS date FROM (SELECT %s FROM (%s) %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins)
		query = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total FROM (SELECT %s FROM (%s) %s)", sql.ChainIDFieldName, sql.SenderFieldName, singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins)
	case model.HistoricalResultTypeTransactions:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s)) AS total, FROM_UNIXTIME(%s, %s) AS date FROM (SELECT %s FROM (%s) %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)

	default:
		return nil, fmt.Errorf("invalid type argument")
	}

	var sum float64
	var err error
	var dayByDayData []*model.DateResult
	g, groupCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		dayByDayData, err = r.DB.GetDateResults(groupCtx, subQuery)
		if err != nil {
			return fmt.Errorf("failed to get dateResults: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		sum, err = r.DB.GetFloat64(groupCtx, query)
		if err != nil {
			return fmt.Errorf("failed to get total sum: %w", err)
		}
		return nil
	})
	err = g.Wait()

	if err != nil {
		return nil, fmt.Errorf("could not get historical data: %w", err)
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) LatestBridgeTransactions(ctx context.Context, includePending *bool, page *int) ([]*model.BridgeTransaction, error) {
	// For each chain ID, get the latest bridge transaction.
	var results []*model.BridgeTransaction
	var err error
	results, err = r.GetBridgeTxsFromOrigin(ctx, nil, nil, nil, *page, nil, *includePending, true)

	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}

	return results, nil
}
