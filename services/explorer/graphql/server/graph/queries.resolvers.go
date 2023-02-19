package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	resolvers "github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/resolver"
	"golang.org/x/sync/errgroup"
)

// BridgeTransactions is the resolver for the bridgeTransactions2 field.
func (r *queryResolver) BridgeTransactions(ctx context.Context, chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txnHash *string, kappa *string, pending *bool, useMv *bool, page *int, tokenAddressFrom []*string, tokenAddressTo []*string) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction
	if useMv != nil && *useMv {
		var mvResults []*model.BridgeTransaction
		var err error
		mvResults, err = r.GetBridgeTxs(ctx, chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, maxAmountUsd, minAmountUsd, startTime, endTime, txnHash, tokenAddressTo, tokenAddressFrom, kappa, pending, page)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
		}
		sort.Sort(SortBridgeTxType(mvResults))
		return mvResults, nil
	}

	var fromResults []*model.BridgeTransaction
	var toResults []*model.BridgeTransaction

	var wg sync.WaitGroup
	var err error
	wg.Add(1)
	go func() {
		defer wg.Done()
		fromResults, err = r.GetBridgeTxsFromOrigin(ctx, useMv, chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, maxAmountUsd, minAmountUsd, startTime, endTime, txnHash, tokenAddressTo, tokenAddressFrom, kappa, pending, page, false)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		toResults, err = r.GetBridgeTxsFromDestination(ctx, useMv, chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, minAmountUsd, maxAmountUsd, startTime, endTime, txnHash, kappa, tokenAddressFrom, tokenAddressTo, page, pending)
	}()

	wg.Wait()
	if err != nil {
		return nil, err
	}
	// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
	// both the origin -> destination transactions that match the search parameters, and the destination -> origin
	// transactions that match the search parameters. Then we need to merge the results and remove duplicates.

	results = r.mergeBridgeTransactions(fromResults, toResults)
	//}
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
		return nil, fmt.Errorf("could not get message bus transactions %w", err)
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
//
// nolint:gocognit
func (r *queryResolver) AmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, platform *model.Platform, chainID *int, address *string, tokenAddress *string, useCache *bool, useMv *bool) (*model.ValueResult, error) {
	if useCache != nil && *useCache {
		res, err := r.getValueResultFromCache(fmt.Sprintf("amountStatistic, %s, %s, %s, %s, %s, %s", typeArg.String(), platform.String(), duration.String(), keyGenHandleNilInt(chainID), keyGenHandleNilString(address), keyGenHandleNilString(tokenAddress)))
		if err == nil {
			return res, nil
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
		hours := 2190
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationPast6Months:
		hours := 4380
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationPastYear:
		hours := 8760
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationAllTime:
		timestampSpecifier = ""
	}
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.SenderFieldName, &firstFilter, "")
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")

	compositeFilters := fmt.Sprintf(
		`%s%s%s`,
		timestampSpecifier, addressSpecifier, chainIDSpecifier,
	)
	var finalSQL *string
	switch *platform {
	case model.PlatformBridge:
		finalSQL, err = GenerateAmountStatisticBridgeSQL(typeArg, address, chainID, tokenAddress)
		if err != nil {
			return nil, err
		}
	case model.PlatformSwap:
		finalSQL, err = GenerateAmountStatisticSwapSQL(typeArg, compositeFilters, tokenAddress)
		if err != nil {
			return nil, err
		}
	case model.PlatformMessageBus:
		if tokenAddress != nil {
			return nil, fmt.Errorf("cannot filter by token on message bus events")
		}

		finalSQL, err = GenerateAmountStatisticMessageBusSQL(typeArg, compositeFilters)

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

		bridgeFinalSQL, err = GenerateAmountStatisticBridgeSQL(typeArg, address, chainID, tokenAddress)
		if err != nil {
			return nil, err
		}

		swapFinalSQL, err = GenerateAmountStatisticSwapSQL(typeArg, compositeFilters, tokenAddress)
		if err != nil {
			return nil, err
		}

		g, groupCtx := errgroup.WithContext(ctx)

		if tokenAddress == nil && typeArg != model.StatisticTypeTotalVolumeUsd && typeArg != model.StatisticTypeMedianVolumeUsd && typeArg != model.StatisticTypeMeanVolumeUsd {
			messageBusFinalSQL, err = GenerateAmountStatisticMessageBusSQL(typeArg, compositeFilters)
			if err != nil {
				return nil, err
			}
			g.Go(func() error {
				messageBusSum, err = r.DB.GetFloat64(groupCtx, *messageBusFinalSQL)

				if err != nil {
					return fmt.Errorf("failed to get dateResults: %w", err)
				}
				return nil
			})
		}
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

		err = g.Wait()
		if err != nil {
			return nil, fmt.Errorf("error getting data from all platforms, %w", err)
		}
		value := fmt.Sprintf("%f", bridgeSum+swapSum+messageBusSum)
		output := model.ValueResult{
			Value: &value,
		}
		err = r.Cache.CacheResponse(fmt.Sprintf("amountStatistic, %s, %s, %s, %s, %s, %s", typeArg.String(), platform.String(), duration.String(), keyGenHandleNilInt(chainID), keyGenHandleNilString(address), keyGenHandleNilString(tokenAddress)), &output)
		if err != nil {
			return nil, fmt.Errorf("error caching results, %w", err)
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
	err = r.Cache.CacheResponse(fmt.Sprintf("amountStatistic, %s, %s, %s, %s, %s, %s", typeArg.String(), platform.String(), duration.String(), keyGenHandleNilInt(chainID), keyGenHandleNilString(address), keyGenHandleNilString(tokenAddress)), &output)
	if err != nil {
		return nil, fmt.Errorf("error storing cache data, %w", err)
	}
	return &output, nil
}

// DailyStatisticsByChain is the resolver for the dailyStatisticsByChain field.
func (r *queryResolver) DailyStatisticsByChain(ctx context.Context, chainID *int, typeArg *model.DailyStatisticType, platform *model.Platform, duration *model.Duration, useCache *bool, useMv *bool) ([]*model.DateResultByChain, error) {
	if useCache != nil && *useCache {
		res, err := r.getDateResultByChainFromCache(fmt.Sprintf("dailyStatisticsByChain, %s, %s, %s, %s", keyGenHandleNilInt(chainID), typeArg.String(), duration.String(), platform.String()))
		if err == nil {
			return res, nil
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
		return nil, fmt.Errorf("could not get daily data by chain: %w", err)
	}
	err = r.Cache.CacheResponse(fmt.Sprintf("dailyStatisticsByChain, %s, %s, %s, %s", keyGenHandleNilInt(chainID), typeArg.String(), duration.String(), platform.String()), res)
	if err != nil {
		return nil, fmt.Errorf("error cahcing response, %w", err)
	}
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
	case model.DurationPast3Months:
		hours := 2190
		targetTime := r.getTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, &firstFilter, "")
	case model.DurationPast6Months:
		hours := 4380
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
