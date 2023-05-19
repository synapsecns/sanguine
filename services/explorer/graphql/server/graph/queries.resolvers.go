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
	results = r.mergeBridgeTransactions(fromResults, toResults)

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
	targetTime := GetTargetTime(hours)
	results, err := r.DB.GetTxCounts(ctx, generateBridgeEventCountQuery(chainID, address, nil, directionIn, &targetTime, false))
	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}
	return results, nil
}

// CountByTokenAddress is the resolver for the countByTokenAddress field.
func (r *queryResolver) CountByTokenAddress(ctx context.Context, chainID *int, address *string, direction *model.Direction, hours *int) ([]*model.TokenCountResult, error) {
	directionIn := r.getDirectionIn(direction)
	targetTime := GetTargetTime(hours)
	results, err := r.DB.GetTokenCounts(ctx, generateBridgeEventCountQuery(chainID, address, nil, directionIn, &targetTime, true))

	if err != nil {
		return nil, fmt.Errorf("failed to get count by chain ID: %w", err)
	}

	return results, nil
}

// AddressRanking is the resolver for the addressRanking field.
func (r *queryResolver) AddressRanking(ctx context.Context, hours *int) ([]*model.AddressRanking, error) {
	targetTime := GetTargetTime(hours)
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
func (r *queryResolver) AmountStatistic(ctx context.Context, typeArg model.StatisticType, duration *model.Duration, platform *model.Platform, chainID *int, address *string, tokenAddress *string, useCache *bool, useMv *bool) (*model.ValueResult, error) {
	if useCache != nil && *useCache {
		res, err := r.getValueResultFromCache(fmt.Sprintf("amountStatistic, %s, %s, %s, %s, %s, %s", typeArg.String(), platform.String(), duration.String(), keyGenHandleNilInt(chainID), keyGenHandleNilString(address), keyGenHandleNilString(tokenAddress)))
		if err == nil {
			return res, nil
		}
	}

	var err error
	firstFilter := true
	timestampSpecifier := GetDurationFilter(duration, &firstFilter, "")
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
		var value *string
		value, err = r.getAmountStatisticsAll(ctx, typeArg, chainID, address, tokenAddress, compositeFilters)
		if err != nil {
			return nil, fmt.Errorf("could not calculate value across all platforms, %w", err)
		}
		output := model.ValueResult{
			Value: value,
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

	if useMv != nil && *useMv {
		mvResult, err := r.getDateResultByChainMv(ctx, chainID, typeArg, platform, duration)
		if err != nil {
			return nil, fmt.Errorf("error getting mv data, %w", err)
		}
		return mvResult, nil
	}

	var err error
	firstFilter := true
	timestampSpecifier := GetDurationFilter(duration, &firstFilter, "")
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
	firstFilter := true
	timestampSpecifier := GetDurationFilter(duration, &firstFilter, "")

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

// AddressData is the resolver for the addressData field.
func (r *queryResolver) AddressData(ctx context.Context, address string) (*model.AddressData, error) {
	bridgeQuery := fmt.Sprintf("SELECT toFloat64(sumKahan(famount_usd)) AS volumeTotal, toFloat64(sumKahan(tfee_amount_usd)) AS feeTotal, toInt64(uniq(fchain_id, ftx_hash)) AS txTotal FROM (SELECT * FROM mv_bridge_events where fsender = '%s' LIMIT 1 BY fchain_id,fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", address)
	swapQuery := fmt.Sprintf("SELECT toFloat64(sumKahan(multiIf(event_type = 0, amount_usd[sold_id], event_type = 1, arraySum(mapValues(amount_usd)), event_type = 9, arraySum(mapValues(amount_usd)), event_type = 10, amount_usd[sold_id],0))) AS volumeTotal, toFloat64(sumKahan(arraySum(mapValues(fee_usd)))) AS feeTotal,  toInt64(uniq(chain_id, tx_hash)) AS txTotal FROM (SELECT * FROM swap_events where sender = '%s' LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash)", address)
	rankingQuery := fmt.Sprintf("select rowNumber from (select sender, row_number() over (order by sumTotal desc ) as rowNumber from (select fsender as sender, sumKahan(famount_usd) as sumTotal from (SELECT * FROM mv_bridge_events where fsender != '' LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) where fsender != '' group by fsender)) where sender = '%s'", address)
	firstTx := fmt.Sprintf("SELECT min(ftimestamp) AS earliestTime FROM (SELECT * FROM mv_bridge_events where fsender = '%s' LIMIT 1 BY fchain_id,fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", address)
	dailyDataQuery := fmt.Sprintf("SELECT coalesce(toString(date), toString(s.date)) AS date, toFloat64(coalesce(sumTotal, 0)) + toFloat64(coalesce(s.sumTotal, 0)) as count FROM (SELECT * FROM (SELECT %s, uniq(fchain_id, ftx_hash) AS sumTotal FROM (SELECT * FROM mv_bridge_events where fsender = '%s' LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) group by date order by date) b FULL OUTER JOIN (SELECT %s, uniq(chain_id, tx_hash) AS sumTotal FROM (SELECT * FROM swap_events WHERE sender = '%s' LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date) s ON b.date = s.date) SETTINGS join_use_nulls=1", toDateSelectMv, address, toDateSelect, address)
	chainRankingQuery := fmt.Sprintf("SELECT row_number() over (order by VolumeUsd desc ) as Rank, tchain_id as ChainID, sumKahan(tamount_usd) AS VolumeUsd FROM (SELECT * FROM mv_bridge_events where fsender = '%s' LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) where ChainID > 0 group by ChainID", address)
	var bridgeVolume float64
	var bridgeFees float64
	var bridgeTxs int
	var swapVolume float64
	var swapFees float64
	var swapTxs int
	var rank int
	var earliestTxTimestamp int
	var addressChainRanking []*model.AddressChainRanking

	var addressDailyData []*model.AddressDailyCount

	g, groupCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		bridgeVolume, bridgeFees, bridgeTxs, err = r.DB.GetAddressData(groupCtx, bridgeQuery)
		if err != nil {
			return fmt.Errorf("failed to get bridge data for address %s: %w", address, err)
		}
		return nil
	})
	g.Go(func() error {
		var err error
		swapVolume, swapFees, swapTxs, err = r.DB.GetAddressData(groupCtx, swapQuery)
		if err != nil {
			return fmt.Errorf("failed to get swap data for address %s: %w", address, err)
		}
		return nil
	})
	g.Go(func() error {
		res, err := r.DB.GetUint64(groupCtx, rankingQuery)
		if err != nil {
			return fmt.Errorf("failed to get ranking for address %s: %w", address, err)
		}
		rank = int(res)
		return nil
	})
	g.Go(func() error {
		res, err := r.DB.GetUint64(groupCtx, firstTx)
		if err != nil {
			return fmt.Errorf("failed to get first timestamp for address %s: %w", address, err)
		}
		earliestTxTimestamp = int(res)
		return nil
	})
	g.Go(func() error {
		var err error
		addressDailyData, err = r.DB.GetAddressDailyData(groupCtx, dailyDataQuery)
		if err != nil {
			return fmt.Errorf("failed to get first daily data for address %s: %w", address, err)
		}
		return nil
	})

	g.Go(func() error {
		var err error
		addressChainRanking, err = r.DB.GetAddressChainRanking(groupCtx, chainRankingQuery)
		if err != nil {
			return fmt.Errorf("failed to get ranking for address %s: %w", address, err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, fmt.Errorf("could not get address data, %w", err)
	}
	res := &model.AddressData{
		BridgeVolume: &bridgeVolume,
		BridgeFees:   &bridgeFees,
		BridgeTxs:    &bridgeTxs,
		SwapVolume:   &swapVolume,
		SwapFees:     &swapFees,
		SwapTxs:      &swapTxs,
		Rank:         &rank,
		EarliestTx:   &earliestTxTimestamp,
		ChainRanking: addressChainRanking,
		DailyData:    addressDailyData,
	}
	return res, nil
}

// Leaderboard is the resolver for the leaderboard field.
func (r *queryResolver) Leaderboard(ctx context.Context, duration *model.Duration, chainID *int, useMv *bool, page *int) ([]*model.Leaderboard, error) {
	if !*useMv {
		return nil, fmt.Errorf("the leaderboard query does not support non-mv based queries")
	}
	firstFilter := false
	timestampSpecifier := GetDurationFilter(duration, &firstFilter, "f")
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "f")
	pageValue := sql.PageSize
	pageOffset := (*page - 1) * sql.PageSize
	filters := timestampSpecifier + chainIDSpecifier
	leaderboardQuery := fmt.Sprintf("select row_number() over (order by VolumeUsd desc ) as Rank, * from (select fsender as Address, toFloat64(sumKahan(famount_usd)) as VolumeUsd,toFloat64(avg(famount_usd)) as AvgVolumeUsd, count(DISTINCT ftx_hash) as Txs,toFloat64(sumKahan(tfee_amount_usd)) as Fees from (SELECT * FROM mv_bridge_events where fsender != '' LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) where fsender != '' %s group by fsender) LIMIT %d OFFSET %d", filters, pageValue, pageOffset)
	leaderboardRes, err := r.DB.GetLeaderboard(ctx, leaderboardQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get leaderboard %w", err)
	}

	return leaderboardRes, nil
}

// Query returns resolvers.QueryResolver implementation.
func (r *Resolver) Query() resolvers.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
