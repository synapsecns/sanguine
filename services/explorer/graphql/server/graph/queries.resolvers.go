package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strings"
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
		results, err = r.GetBridgeTxsFromOrigin(ctx, chainID, address, txnHash, includePending, page, tokenAddress, false)
		if err != nil {
			return nil, err
		}
	case kappa != nil:
		// If we are given a kappa, we search for the bridge transaction on the destination chain, then locate
		// its counterpart on the origin chain using a query to find a transaction hash given a kappa.
		results, err = r.GetBridgeTxsFromDestination(ctx, chainID, address, txnHash, kappa, page, tokenAddress)
		if err != nil {
			return nil, err
		}

	default:
		// If we have either just a chain ID or an address, or both a chain ID and an address, we need to search for
		// both the origin -> destination transactions that match the search parameters, and the destination -> origin
		// transactions that match the search parameters. Then we need to merge the results and remove duplicates.
		fromResults, err := r.GetBridgeTxsFromOrigin(ctx, chainID, address, txnHash, includePending, page, tokenAddress, false)
		if err != nil {
			return nil, err
		}
		toResults, err := r.GetBridgeTxsFromDestination(ctx, chainID, address, txnHash, kappa, page, tokenAddress)
		if err != nil {
			return nil, err
		}
		fmt.Println("LENS", len(fromResults), len(toResults))
		results = r.mergeBridgeTransactions(fromResults, toResults)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
	}

	return results, nil
}

// LatestBridgeTransactions is the resolver for the latestBridgeTransactions field.
func (r *queryResolver) LatestBridgeTransactions(ctx context.Context, includePending bool, page int) ([]*model.BridgeTransaction, error) {
	// For each chain ID, get the latest bridge transaction.
	//return r.GetBridgeTxsFromOrigin(ctx, nil, nil, nil, includePending, page, nil, true)

	//its getting the pending kappa and not the chaina one (shoudl only be getting the latest)
	var results []*model.BridgeTransaction
	//mapChains := make(map[uint32]string)
	fromBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generatePartialInfoQueryByChain(100))
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}
	fromBridgeEventsMap := make(map[string][]sql.BridgeEvent)
	for _, fromBridgeEvent := range fromBridgeEvents {
		key := fmt.Sprintf("%d", fromBridgeEvent.ChainID)
		if fromBridgeEventsMap[key] == nil {
			fromBridgeEventsMap[key] = append(fromBridgeEventsMap[key], fromBridgeEvent)
		}
	}

	var toKappaChainArr []string
	for _, bridgeEvent := range fromBridgeEvents {
		if bridgeEvent.DestinationChainID != nil {
			toKappaChainArr = append(toKappaChainArr, fmt.Sprintf("(%d,'%s')", bridgeEvent.DestinationChainID, bridgeEvent.DestinationKappa))
		}
	}
	toKappaChainStr := "(" + strings.Join(toKappaChainArr, ",") + ")"
	toBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generateToKappaPartialInfoQuery(toKappaChainStr, nil, nil, nil, nil, nil, page, true))
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}
	toBridgeEventsMap := make(map[string][]sql.BridgeEvent)
	for _, toBridgeEvent := range toBridgeEvents {
		key := fmt.Sprintf("%d", toBridgeEvent.ChainID)
		toBridgeEventsMap[key] = append(toBridgeEventsMap[key], toBridgeEvent)
	}

	// Check for pending

	var fromBridgeEventsCleaned []sql.BridgeEvent
	var toBridgeEventsCleaned []sql.BridgeEvent
	toBridgeEventsMapCleaned := make(map[string][]sql.BridgeEvent)
	if includePending {
		for key := range fromBridgeEventsMap {
			fromBridgeEventsCleaned = append(fromBridgeEventsCleaned, fromBridgeEventsMap[key][:1]...)
		}
		for key := range toBridgeEventsMap {
			toBridgeEventsMapCleaned[key] = toBridgeEventsMap[key][:1]
			toBridgeEventsCleaned = append(toBridgeEventsCleaned, toBridgeEventsMapCleaned[key]...)
		}
	} else {
		for key := range toBridgeEventsMap {
			for _, fromBridgeEvent := range fromBridgeEvents {
				for _, toBridgeEvent := range toBridgeEventsMap[key] {

					if fromBridgeEvent.DestinationKappa == toBridgeEvent.Kappa.String {
						fromBridgeEventsCleaned = append(fromBridgeEventsCleaned, fromBridgeEvent)
						toBridgeEventsCleaned = append(toBridgeEventsCleaned, toBridgeEvent)
						toBridgeEventsMapCleaned[key] = []sql.BridgeEvent{toBridgeEvent}
						goto NEXTCHAIN
					}
				}
			}
		NEXTCHAIN:
		}
	}

	fromInfos, err := GetPartialInfoFromBridgeEvent(fromBridgeEventsCleaned)
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}
	toInfos, err := GetToPartialInfoFromBridgeEvent(toBridgeEventsCleaned)
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}

	for i, fromBridgeTx := range fromBridgeEventsCleaned {
		// If we are not including pending transactions, and the transaction is pending, skip it.
		toBridgeEvent := toBridgeEventsMapCleaned[fromBridgeTx.DestinationChainID.String()]
		if fromBridgeTx.DestinationKappa == toBridgeEvent[0].Kappa.String {
			var swapSuccess bool
			if toBridgeEvent[0].SwapSuccess.Uint64() == 1 {
				swapSuccess = true
			}
			if err != nil {
				return nil, fmt.Errorf("failed to get swap success: %w", err)
			}

			pending := false
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfos[i],
				ToInfo:      toInfos[fromBridgeTx.DestinationChainID.String()],
				Kappa:       &toBridgeEvent[0].Kappa.String,
				Pending:     &pending,
				SwapSuccess: &swapSuccess,
			})
		} else {
			if includePending {
				kappa := fromBridgeTx.DestinationKappa
				results = append(results, &model.BridgeTransaction{
					FromInfo:    fromInfos[i],
					ToInfo:      nil,
					Kappa:       &kappa,
					Pending:     &includePending,
					SwapSuccess: nil,
				})
			}
		}
	}

	if err != nil {
		return nil, fmt.Errorf("failed to get bridge transaction: %w", err)
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
			chainIDFilter = generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
			blockNumberFilter = fmt.Sprintf("AND %s >= %d", sql.InsertTimeFieldName, targetTime)
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
	finalSQL := fmt.Sprintf("\nSELECT %s(toUInt256(%s)) FROM %s %s", operation, sql.AmountUSDFieldName, subQuery, additionalFilters)
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

	var results []*model.TransactionCountResult

	for i := range chainIDs {
		count, err := r.DB.GetUint64(ctx, generateBridgeEventCountQuery(chainIDs[i], address, nil, directionIn, &targetTime))
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
			`SELECT DISTINCT %s FROM bridge_events WHERE %s = %d OR %s = %d AND %s`,
			sql.TokenFieldName, sql.ChainIDFieldName, chain, sql.DestinationChainIDFieldName, chain, deDupInQuery,
		)
		tokenAddresses, err := r.DB.GetStringArray(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("failed to get token addresses: %w", err)
		}

		chainIDsToTokenAddresses[chain] = tokenAddresses
	}

	directionIn := r.getDirectionIn(direction)
	targetTime := r.getTargetTime(hours)

	// Get the number of transactions for each token address, for each chain ID.
	var results []*model.TokenCountResult

	for chain, tokenAddresses := range chainIDsToTokenAddresses {
		for i := range tokenAddresses {
			count, err := r.DB.GetUint64(ctx, generateBridgeEventCountQuery(chain, address, &tokenAddresses[i], directionIn, &targetTime))
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

	query := fmt.Sprintf(`SELECT %s AS address, COUNT(DISTINCT %s) AS count FROM %s GROUP BY address ORDER BY count DESC`, sql.TokenFieldName, sql.TxHashFieldName, subQuery)
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
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, FROM_UNIXTIME(%s, %s) AS date FROM bridge_events %s AND %s GROUP BY date ORDER BY total DESC", sql.AmountUSDFieldName, sql.TimeStampFieldName, "'%d/%m/%Y'", filter, deDupInQuery)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	case model.HistoricalResultTypeAddresses:
		subQuery = fmt.Sprintf("SELECT uniqExact(%s) AS total, FROM_UNIXTIME( %s, %s) AS date FROM  bridge_events %s AND %s GROUP BY date ORDER BY total DESC", sql.SenderFieldName, sql.TimeStampFieldName, "'%d/%m/%Y'", filter, deDupInQuery)
		query = fmt.Sprintf("SELECT uniqExact(%s) FROM bridge_events %s AND %s", sql.SenderFieldName, filter, deDupInQuery)
	case model.HistoricalResultTypeTransactions:
		subQuery = fmt.Sprintf("SELECT uniqExact(%s) AS total, FROM_UNIXTIME(%s, %s) AS date FROM  bridge_events %s AND %s GROUP BY date ORDER BY total DESC", sql.TxHashFieldName, sql.TimeStampFieldName, "'%d/%m/%Y'", filter, deDupInQuery)
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

func keyGen(chainID string, kappa string) string {
	return fmt.Sprintf("%s-%s", chainID, kappa)
}
