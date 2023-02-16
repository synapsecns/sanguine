package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/resolver"
	"golang.org/x/sync/errgroup"
)

// BridgeTransactions is the resolver for the bridgeTransactions2 field.
func (r *queryResolver) BridgeTransactions(ctx context.Context, chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txnHash *string, kappa *string, pending *bool, page *int, tokenAddressFrom []*string, tokenAddressTo []*string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	switch {
	case kappa != nil:
		// If we are given a kappa, we search for the bridge transaction on the destination chain, then locate
		// its counterpart on the origin chain using a query to find a transaction hash given a kappa.
		results, err = r.GetBridgeTxsFromDestination(ctx, chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, minAmountUsd, maxAmountUsd, startTime, endTime, txnHash, kappa, page, tokenAddressFrom, tokenAddressTo)
		if err != nil {
			return nil, err
		}
	default:
		var fromResults []*model.BridgeTransaction
		var toResults []*model.BridgeTransaction
		var wg sync.WaitGroup
		var err error
		wg.Add(1)
		go func() {
			defer wg.Done()
			fromResults, err = r.GetBridgeTxsFromOrigin(ctx, chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, maxAmountUsd, minAmountUsd, startTime, endTime, txnHash, page, tokenAddressTo, tokenAddressFrom, *pending, false)
		}()
		if !*pending {
			wg.Add(1)
			go func() {
				defer wg.Done()
				toResults, err = r.GetBridgeTxsFromDestination(ctx, chainIDTo, chainIDFrom, addressTo, addressFrom, maxAmount, minAmount, minAmountUsd, maxAmountUsd, startTime, endTime, txnHash, kappa, page, tokenAddressFrom, tokenAddressTo)
			}()
		}
		wg.Wait()
		if err != nil {
			return nil, err
		}
		// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
		// both the origin -> destination transactions that match the search parameters, and the destination -> origin
		// transactions that match the search parameters. Then we need to merge the results and remove duplicates.

		results = r.mergeBridgeTransactions(fromResults, toResults)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}
	sort.Sort(SortBridgeTxType(results))
	return results, nil
}

// MessageBusTransactions is the resolver for the messageBusTransactions field.
func (r *queryResolver) MessageBusTransactions(ctx context.Context, chainID []*int, contractAddress *string, startTime *int, endTime *int, txnHash *string, messageID *string, pending *bool, reverted *bool, page *int) ([]*model.MessageBusTransaction, error) {
	var err error
	var results []*model.MessageBusTransaction
	results, err = r.GetMessageBusTxs(ctx, chainID, contractAddress, startTime, endTime, txnHash, messageID, *pending, *reverted, page)
	if err != nil {
		fmt.Errorf("could not get message bus transactions %v", err)
	}
	sort.Sort(SortMessageBusTxType(results))
	return results, nil
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
	query := fmt.Sprintf(`SELECT %s AS address, COUNT(DISTINCT %s) AS Count FROM (%s) GROUP BY %s ORDER BY Count Desc`, sql.SenderFieldName, sql.TxHashFieldName, generateDeDepQuery(compositeFilters, nil, nil), sql.SenderFieldName)
	res, err := r.DB.GetAddressRanking(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}

	return res, nil
}

// AmountStatistic is the resolver for the amountStatistic field.
func (r *queryResolver) AmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, platform *model.Platform, chainID *int, address *string, tokenAddress *string, useCache *bool) (*model.ValueResult, error) {
	if useCache != nil && *useCache {
		cacheResult := r.Cache.GetCache(fmt.Sprintf("amountStatistic, %s, %s, %s, %s, %s, %s", typeArg.String(), platform.String(), duration.String(), keyGenHandleNilInt(chainID), keyGenHandleNilString(address), keyGenHandleNilString(tokenAddress)))
		if cacheResult != nil {
			result := *(cacheResult.(*interface{}))
			return result.(*model.ValueResult), nil
		}
	}

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
	firstFilter = true
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.SenderFieldName, &firstFilter, "")
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")

	compositeFilters := fmt.Sprintf(
		`%s%s%s%s`,
		timestampSpecifier, tokenAddressSpecifier, addressSpecifier, chainIDSpecifier,
	)

	var finalSQL *string
	switch *platform {
	case model.PlatformBridge:
		finalSQL, err = GenerateAmountStatisticBridgeSQL(typeArg, compositeFilters, &firstFilter, tokenAddressSpecifier)
		if err != nil {
			return nil, err
		}
	case model.PlatformSwap:
		finalSQL, err = GenerateAmountStatisticSwapSQL(typeArg, compositeFilters, tokenAddressSpecifier)
		if err != nil {
			return nil, err
		}
	case model.PlatformMessageBus:

		finalSQL, err = GenerateAmountStatisticMessageBusSQL(typeArg, compositeFilters, tokenAddressSpecifier)

		if err != nil {
			return nil, err
		}
	case model.PlatformAll:
		if typeArg == model.StatisticTypeMedianVolumeUsd || typeArg == model.StatisticTypeMeanVolumeUsd || typeArg == model.StatisticTypeMedianFeeUsd || typeArg == model.StatisticTypeMeanFeeUsd {
			return nil, fmt.Errorf("cannot calculate averages or medians across all platforms")
		}
		var bridgeFinalSQL *string
		var swapFinalSQL *string
		var messageBusFinalSQL *string

		var bridgeSum float64
		var swapSum float64
		var messageBusSum float64

		bridgeFinalSQL, err = GenerateAmountStatisticBridgeSQL(typeArg, compositeFilters, &firstFilter, tokenAddressSpecifier)
		if err != nil {
			return nil, err
		}

		swapFinalSQL, err = GenerateAmountStatisticSwapSQL(typeArg, compositeFilters, tokenAddressSpecifier)
		if err != nil {
			return nil, err
		}

		messageBusFinalSQL, err = GenerateAmountStatisticMessageBusSQL(typeArg, compositeFilters, tokenAddressSpecifier)
		if err != nil {
			return nil, err
		}

		g, groupCtx := errgroup.WithContext(ctx)
		g.Go(func() error {
			bridgeSum, err = r.DB.GetFloat64(groupCtx, *bridgeFinalSQL)
			if err != nil {
				return fmt.Errorf("failed to get dateResults: %w", err)
			}
			return nil
		})
		g.Go(func() error {
			swapSum, err = r.DB.GetFloat64(groupCtx, *swapFinalSQL)
			if err != nil {
				return fmt.Errorf("failed to get dateResults: %w", err)
			}

			return nil
		})
		if typeArg != model.StatisticTypeTotalVolumeUsd {
			g.Go(func() error {
				messageBusSum, err = r.DB.GetFloat64(groupCtx, *messageBusFinalSQL)

				if err != nil {
					return fmt.Errorf("failed to get dateResults: %w", err)
				}
				return nil
			})
		}
		err = g.Wait()
		if err != nil {
			return nil, err
		}
		value := fmt.Sprintf("%f", bridgeSum+swapSum+messageBusSum)
		output := model.ValueResult{
			Value: &value,
		}

		return &output, nil
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	if finalSQL == nil {
		return nil, fmt.Errorf("invalid statistic or platform type: %s", typeArg)
	}
	res, err := r.DB.GetFloat64(ctx, *finalSQL)
	if err != nil {
		return nil, fmt.Errorf("failed to get amount data stats: %w", err)
	}

	value := fmt.Sprintf("%f", res)
	output := model.ValueResult{
		Value: &value,
	}
	r.Cache.CacheResponse(fmt.Sprintf("amountStatistic, %s, %s, %s, %s, %s, %s", typeArg.String(), platform.String(), duration.String(), keyGenHandleNilInt(chainID), keyGenHandleNilString(address), keyGenHandleNilString(tokenAddress)), &output)
	return &output, nil
}

// DailyStatistics is the resolver for the dailyStatistics field.
func (r *queryResolver) DailyStatistics(ctx context.Context, chainID *int, typeArg *model.DailyStatisticType, platform *model.Platform, days *int) (*model.DailyResult, error) {
	if *typeArg == model.DailyStatisticTypeVolume && *platform == model.PlatformMessageBus {
		return nil, fmt.Errorf("cannot calculate volume for the message bus")
	}
	var subQuery *string
	var query *string
	var err error
	startTime := uint64(time.Now().Unix() - int64(*days*86400))
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	timeStampSpecifier := generateTimestampSpecifierSQL(&startTime, sql.TimeStampFieldName, &firstFilter, "")
	compositeFilters := fmt.Sprintf("%s%s", chainIDSpecifier, timeStampSpecifier)

	switch *platform {
	case model.PlatformBridge:
		subQuery, query, err = GenerateDailyStatisticBridgeSQL(typeArg, compositeFilters, &firstFilter)
		if err != nil {
			return nil, err
		}
	case model.PlatformSwap:
		subQuery, query, err = GenerateDailyStatisticSwapSQL(typeArg, compositeFilters)
		if err != nil {
			return nil, err
		}
	case model.PlatformMessageBus:
		subQuery, query, err = GenerateDailyStatisticMessageBusSQL(typeArg, compositeFilters)
		if err != nil {
			return nil, err
		}
	case model.PlatformAll:
		bridgeSubQuery, bridgeQuery, bridgeErr := GenerateDailyStatisticBridgeSQL(typeArg, compositeFilters, &firstFilter)
		if bridgeErr != nil {
			return nil, bridgeErr
		}
		swapSubQuery, swapQuery, swapErr := GenerateDailyStatisticSwapSQL(typeArg, compositeFilters)
		if swapErr != nil {
			return nil, swapErr
		}

		messageBusSubQuery, messageBusQuery, messageBusErr := GenerateDailyStatisticMessageBusSQL(typeArg, compositeFilters)
		if messageBusErr != nil {
			return nil, messageBusErr
		}

		var bridgeSum float64
		var swapSum float64
		var messageBusSum float64
		var dailyBridgeData []*model.DateResult
		var dailySwapData []*model.DateResult
		var dailyMessageBusData []*model.DateResult
		g, groupCtx := errgroup.WithContext(ctx)

		// Get Bridge
		g.Go(func() error {
			dailyBridgeData, err = r.DB.GetDateResults(groupCtx, fmt.Sprintf("%s %s", generateDeDepQueryCTE(compositeFilters, nil, nil, true), *bridgeSubQuery))
			if err != nil {
				return fmt.Errorf("failed to get dateResults: %w", err)
			}
			return nil
		})
		g.Go(func() error {
			bridgeSum, err = r.DB.GetFloat64(groupCtx, *bridgeQuery)
			if err != nil {
				return fmt.Errorf("failed to get total sum: %w", err)
			}
			return nil
		})

		// Get Swap
		g.Go(func() error {
			dailySwapData, err = r.DB.GetDateResults(groupCtx, fmt.Sprintf("%s %s", generateDeDepQueryCTE(compositeFilters, nil, nil, true), *swapSubQuery))
			if err != nil {
				return fmt.Errorf("failed to get dateResults: %w", err)
			}
			return nil
		})
		g.Go(func() error {
			swapSum, err = r.DB.GetFloat64(groupCtx, *swapQuery)
			if err != nil {
				return fmt.Errorf("failed to get total sum: %w", err)
			}
			return nil
		})

		// if volume skip
		// Get Message Bus
		if *typeArg != model.DailyStatisticTypeVolume {
			g.Go(func() error {
				dailyMessageBusData, err = r.DB.GetDateResults(groupCtx, fmt.Sprintf("%s %s", generateDeDepQueryCTE(compositeFilters, nil, nil, true), *messageBusSubQuery))
				if err != nil {
					return fmt.Errorf("failed to get dateResults: %w", err)
				}
				return nil
			})
			g.Go(func() error {
				messageBusSum, err = r.DB.GetFloat64(groupCtx, *messageBusQuery)
				if err != nil {
					return fmt.Errorf("failed to get total sum: %w", err)
				}
				return nil
			})
		}
		err = g.Wait()

		totalDailyResults := make(map[string]float64)

		for i := range dailyBridgeData {
			key := *dailyBridgeData[i].Date
			totalDailyResults[key] = totalDailyResults[key] + *dailyBridgeData[i].Total
		}
		for i := range dailySwapData {
			key := *dailySwapData[i].Date
			totalDailyResults[key] = totalDailyResults[key] + *dailySwapData[i].Total
		}
		for i := range dailyMessageBusData {
			key := *dailyMessageBusData[i].Date
			totalDailyResults[key] = totalDailyResults[key] + *dailyMessageBusData[i].Total
		}

		var finalDailyData []*model.DateResult
		for k := range totalDailyResults {
			date := k
			value := totalDailyResults[k]
			entry := model.DateResult{&date, &value}
			finalDailyData = append(finalDailyData, &entry)
		}

		totalSum := bridgeSum + swapSum + messageBusSum
		payload := model.DailyResult{
			Total:       &totalSum,
			DateResults: finalDailyData,
		}
		return &payload, nil
	}

	var sum float64
	var dayByDayData []*model.DateResult
	g, groupCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		dayByDayData, err = r.DB.GetDateResults(groupCtx, fmt.Sprintf("%s %s", generateDeDepQueryCTE(compositeFilters, nil, nil, true), *subQuery))
		if err != nil {
			return fmt.Errorf("failed to get dateResults: %w", err)
		}
		return nil
	})

	g.Go(func() error {
		sum, err = r.DB.GetFloat64(groupCtx, *query)
		if err != nil {
			return fmt.Errorf("failed to get total sum: %w", err)
		}
		return nil
	})
	err = g.Wait()

	if err != nil {
		return nil, fmt.Errorf("could not get historical data: %w", err)
	}

	payload := model.DailyResult{
		Total:       &sum,
		DateResults: dayByDayData,
		Type:        typeArg,
	}
	return &payload, nil
}

// DailyStatisticsByChain is the resolver for the dailyStatisticsByChain field.
func (r *queryResolver) DailyStatisticsByChain(ctx context.Context, chainID *int, typeArg *model.DailyStatisticType, platform *model.Platform, duration *model.Duration, useCache *bool) ([]*model.DateResultByChain, error) {
	if useCache != nil && *useCache {
		cacheResult := r.Cache.GetCache(fmt.Sprintf("dailyStatisticsByChain, %s, %s, %s, %s", keyGenHandleNilInt(chainID), typeArg.String(), duration.String(), platform.String()))
		if cacheResult != nil {
			result := *(cacheResult.(*interface{}))
			return result.([]*model.DateResultByChain), nil
		}
	}
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
	case model.DurationPast3Months:
		hours := 2160
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationPast6Months:
		hours := 4320
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationPastYear:
		hours := 8760
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationAllTime:
		timestampSpecifier = ""
	}

	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	compositeFilters := fmt.Sprintf(
		`%s%s`,
		timestampSpecifier, chainIDSpecifier,
	)

	var res []*model.DateResultByChain
	var query *string
	g, groupCtx := errgroup.WithContext(ctx)
	switch *platform {
	case model.PlatformBridge:
		query, err = GenerateDailyStatisticByChainBridgeSQL(typeArg, compositeFilters, &firstFilter)
		if err != nil {
			return nil, err
		}
	case model.PlatformSwap:
		query, err = GenerateDailyStatisticByChainSwapSQL(typeArg, compositeFilters)
		if err != nil {
			return nil, err
		}
	case model.PlatformMessageBus:
		query, err = GenerateDailyStatisticByChainMessageBusSQL(typeArg, compositeFilters)
		if err != nil {
			return nil, err
		}
	case model.PlatformAll:
		query, err = GenerateDailyStatisticByChainAllSQL(typeArg, compositeFilters, &firstFilter)
	default:
		return nil, fmt.Errorf("unsupported platform")
	}
	g.Go(func() error {
		res, err = r.DB.GetDailyTotals(groupCtx, *query)
		if err != nil {
			return fmt.Errorf("failed to get dateResults: %w", err)
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get daily data: %w", err)
	}
	r.Cache.CacheResponse(fmt.Sprintf("dailyStatisticsByChain, %s, %s, %s, %s", keyGenHandleNilInt(chainID), typeArg.String(), duration.String(), platform.String()), res)
	return res, nil
}

// RankedChainIDsByVolume is the resolver for the rankedChainIDsByVolume field.
func (r *queryResolver) RankedChainIDsByVolume(ctx context.Context, duration *model.Duration, useCache *bool) ([]*model.VolumeByChainID, error) {
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
	case model.DurationPastYear:
		hours := 8760
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationAllTime:
		timestampSpecifier = ""
	}

	query := GenerateRankedChainsByVolumeSQL(timestampSpecifier, &firstFilter)

	var res []*model.VolumeByChainID
	g, groupCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		res, err = r.DB.GetRankedChainsByVolume(groupCtx, query)
		if err != nil {
			return fmt.Errorf("failed to get dateResults: %w", err)
		}
		return nil
	})

	err = g.Wait()

	if err != nil {
		return nil, fmt.Errorf("could not get daily data: %w", err)
	}

	return res, nil
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
