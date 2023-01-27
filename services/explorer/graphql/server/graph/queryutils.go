package graph

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
)

// nolint:unparam
func generateDeDepQuery(filter string, page *int, offset *int) string {
	if page != nil || offset != nil {
		return fmt.Sprintf("SELECT * FROM bridge_events %s ORDER BY timestamp DESC, block_number DESC, event_index DESC, insert_time DESC LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash LIMIT %d OFFSET %d", filter, *page, *offset)
	}

	return fmt.Sprintf("SELECT * FROM bridge_events %s ORDER BY timestamp DESC, block_number DESC, event_index DESC, insert_time DESC LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash", filter)
}

func generateDeDepQueryCTE(filter string, page *int, offset *int, in bool) string {
	minTimestamp := " (SELECT min(timestamp) - 86400 FROM baseQuery) AS minTimestamp"
	if in {
		minTimestamp = " (SELECT min(timestamp) FROM baseQuery) AS minTimestamp"
	}
	if page != nil || offset != nil {
		return fmt.Sprintf("WITH baseQuery AS (SELECT * FROM bridge_events %s ORDER BY timestamp DESC, block_number DESC, event_index DESC, insert_time DESC LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash LIMIT %d OFFSET %d), %s, %s", filter, *page, *offset, minTimestamp, swapDeDup)
	}
	return fmt.Sprintf("WITH baseQuery AS (SELECT * FROM bridge_events %s ORDER BY timestamp DESC, block_number DESC, event_index DESC, insert_time DESC LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash), %s, %s", filter, minTimestamp, swapDeDup)
}

func (r *queryResolver) getDirectionIn(direction *model.Direction) bool {
	var directionIn bool

	if direction != nil {
		directionIn = *direction == model.DirectionIn
	} else {
		directionIn = true
	}

	return directionIn
}

func (r *queryResolver) getTargetTime(hours *int) uint64 {
	var targetTime uint64

	if hours == nil {
		targetTime = uint64(time.Now().Add(-time.Hour * 24).Unix())
	} else {
		targetTime = uint64(time.Now().Add(-time.Hour * time.Duration(*hours)).Unix())
	}

	return targetTime
}

func (r *queryResolver) mergeBridgeTransactions(origin []*model.BridgeTransaction, destination []*model.BridgeTransaction) []*model.BridgeTransaction {
	var results []*model.BridgeTransaction
	uniqueBridgeTransactions := make(map[string]*model.BridgeTransaction)

	for _, originTx := range origin {
		key := keyGen(fmt.Sprintf("%d", *originTx.FromInfo.ChainID), *originTx.Kappa)
		uniqueBridgeTransactions[key] = originTx
	}

	for _, destinationTx := range destination {
		key := keyGen(fmt.Sprintf("%d", *destinationTx.FromInfo.ChainID), *destinationTx.Kappa)
		uniqueBridgeTransactions[key] = destinationTx
	}

	for _, v := range uniqueBridgeTransactions {
		results = append(results, v)
	}

	return results
}

// generateAddressSpecifierSQL generates a where function with an string.
//
// nolint:unparam
func generateAddressSpecifierSQL(address *string, firstFilter *bool, tablePrefix string) string {
	if address != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE (%s%s = '%s' OR  %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
		}

		return fmt.Sprintf(" AND (%s%s = '%s' OR %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	}

	return ""
}

// generateEqualitySpecifierSQL generates a where function with an equality.
//
// nolint:unparam
func generateEqualitySpecifierSQL(value *int, field string, firstFilter *bool, tablePrefix string, greaterThan bool) string {
	operator := "<"
	if greaterThan {
		operator = ">"
	}
	if value != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s %s %d", tablePrefix, field, operator, *value)
		}

		return fmt.Sprintf(" AND %s%s %s %d", tablePrefix, field, operator, *value)
	}

	return ""
}

// generateDirectionSpecifierSQL generates a where function with a string.
//
// nolint:unparam
func generateDirectionSpecifierSQL(in bool, firstFilter *bool, tablePrefix string) string {
	if in {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s > 0", tablePrefix, sql.DestinationChainIDFieldName)
		}

		return fmt.Sprintf(" AND %s%s > 0", tablePrefix, sql.DestinationChainIDFieldName)
	}
	if *firstFilter {
		*firstFilter = false

		return fmt.Sprintf(" WHERE %s%s == 0", tablePrefix, sql.DestinationChainIDFieldName)
	}

	return fmt.Sprintf(" AND %s%s == 0", tablePrefix, sql.DestinationChainIDFieldName)
}

// generateSingleSpecifierI32SQL generates a where function with an uint32.
//
// nolint:unparam
func generateSingleSpecifierI32SQL(value *int, field string, firstFilter *bool, tablePrefix string) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = %d", tablePrefix, field, *value)
		}

		return fmt.Sprintf(" AND %s%s = %d", tablePrefix, field, *value)
	}

	return ""
}

// generateSingleSpecifierI32ArrSQL generates a where function with an uint32.
//
// nolint:unparam
func generateSingleSpecifierI32ArrSQL(values []*int, field string, firstFilter *bool, tablePrefix string) string {
	if len(values) == 0 {
		return ""
	}
	var final string
	if *firstFilter {
		*firstFilter = false
		final += " WHERE ("
	}

	for i := range values {
		final += fmt.Sprintf(" %s%s = %d", tablePrefix, field, *values[i])
		if i < len(values)-1 {
			final += " OR "
		}
	}

	return final + ")"
}

// GenerateSingleSpecifierStringSQL generates a where function with a string.
//
// nolint:unparam
func generateSingleSpecifierStringArrSQL(values []*string, field string, firstFilter *bool, tablePrefix string) string {

	if len(values) == 0 {
		return ""
	}
	var final string
	if *firstFilter {
		*firstFilter = false
		final += " WHERE ("
	} else {
		final += " AND ("
	}

	for i := range values {
		final += fmt.Sprintf(" %s%s = '%s'", tablePrefix, field, *values[i])
		if i < len(values)-1 {
			final += " OR "
		}
	}

	return final + ")"

}

// generateTimestampSpecifierSQL generates a where function with an uint64.
//
// nolint:unparam
func generateTimestampSpecifierSQL(value *uint64, field string, firstFilter *bool, tablePrefix string) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s >= %d", tablePrefix, field, *value)
		}

		return fmt.Sprintf(" AND %s%s >= %d", tablePrefix, field, *value)
	}

	return ""
}

// GenerateSingleSpecifierStringSQL generates a where function with a string.
//
// nolint:unparam
func generateSingleSpecifierStringSQL(value *string, field string, firstFilter *bool, tablePrefix string) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, field, *value)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, field, *value)
	}

	return ""
}

// generateKappaSpecifierSQL generates a where function with a string.
func generateKappaSpecifierSQL(value *string, field string, firstFilter *bool, tablePrefix string) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, field, *value)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, field, *value)
	}

	return ""
}

//// generateDestinationChainIDSpecifierSQL generates a where function with a string.
// func generateDestinationChainIDSpecifierSQL(field string, firstFilter *bool, tablePrefix string, destination bool) string {
//	if destination {
//		if *firstFilter {
//			*firstFilter = false
//
//			return fmt.Sprintf(" WHERE %s%s == 0", tablePrefix, field)
//		}
//
//		return fmt.Sprintf(" AND %s%s  == 0", tablePrefix, field)
//	}
//	if *firstFilter {
//		*firstFilter = false
//
//		return fmt.Sprintf(" WHERE %s%s > 0", tablePrefix, field)
//	}
//	return fmt.Sprintf(" AND %s%s  > 0", tablePrefix, field)
//}

// generateBridgeEventCountQuery creates the query for bridge event count.
func generateBridgeEventCountQuery(chainID *int, address *string, tokenAddress *string, directionIn bool, timestamp *uint64, isTokenCount bool) string {
	chainField := sql.ChainIDFieldName

	firstFilter := true
	directionSpecifier := generateDirectionSpecifierSQL(directionIn, &firstFilter, "")
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, chainField, &firstFilter, "")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.RecipientFieldName, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	timestampSpecifier := generateTimestampSpecifierSQL(timestamp, sql.TimeStampFieldName, &firstFilter, "")

	compositeFilters := fmt.Sprintf(
		`%s%s%s%s%s`,
		directionSpecifier, chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, timestampSpecifier,
	)
	var query string
	if isTokenCount {
		query = fmt.Sprintf(`%s SELECT %s, %s AS TokenAddress, COUNT(DISTINCT (%s)) AS Count FROM (SELECT %s FROM %s %s) GROUP BY %s, %s ORDER BY Count Desc`,
			generateDeDepQueryCTE(compositeFilters, nil, nil, true), sql.ChainIDFieldName, sql.TokenFieldName, sql.TxHashFieldName, singleSideCol, "baseQuery", singleSideJoinsCTE, sql.TokenFieldName, sql.ChainIDFieldName)
	} else {
		query = fmt.Sprintf(`%s SELECT %s, COUNT(DISTINCT (%s)) AS Count FROM (SELECT %s FROM %s %s) GROUP BY %s ORDER BY Count Desc`,
			generateDeDepQueryCTE(compositeFilters, nil, nil, true), sql.ChainIDFieldName, sql.TxHashFieldName, singleSideCol, "baseQuery", singleSideJoinsCTE, sql.ChainIDFieldName)
	}
	return query
}

// GetPartialInfoFromBridgeEventHybrid returns the partial info from bridge event.
//
// nolint:cyclop
func GetPartialInfoFromBridgeEventHybrid(bridgeEvent sql.HybridBridgeEvent, includePending bool) (*model.BridgeTransaction, error) {
	if includePending && bridgeEvent.TTxHash != "" {
		// nolint:nilnil
		return nil, nil
	}
	var bridgeTx model.BridgeTransaction
	fromChainID := int(bridgeEvent.FChainID)
	fromBlockNumber := int(bridgeEvent.FBlockNumber)
	fromValue := bridgeEvent.FAmount.String()
	var fromTimestamp int
	var fromFormattedValue *float64
	if bridgeEvent.FTokenDecimal != nil {
		fromFormattedValue = getAdjustedValue(bridgeEvent.FAmount, *bridgeEvent.FTokenDecimal)
	} else {
		return nil, fmt.Errorf("token decimal is not valid")
	}
	if bridgeEvent.FTimeStamp != nil {
		fromTimestamp = int(*bridgeEvent.FTimeStamp)
	} else {
		return nil, fmt.Errorf("timestamp is not valid")
	}

	fromInfos := &model.PartialInfo{
		ChainID:        &fromChainID,
		Address:        &bridgeEvent.FRecipient.String,
		TxnHash:        &bridgeEvent.FTxHash,
		Value:          &fromValue,
		FormattedValue: fromFormattedValue,
		USDValue:       bridgeEvent.FAmountUSD,
		TokenAddress:   &bridgeEvent.FToken,
		TokenSymbol:    &bridgeEvent.FTokenSymbol.String,
		BlockNumber:    &fromBlockNumber,
		Time:           &fromTimestamp,
	}

	// If not pending, return a destination partial, otherwise toInfos will be null.
	var pending bool
	var toInfos *model.PartialInfo
	// nolint:nestif
	if bridgeEvent.TTxHash != "" {
		toChainID := int(bridgeEvent.TChainID)
		toBlockNumber := int(bridgeEvent.TBlockNumber)
		toValue := bridgeEvent.TAmount.String()
		var toTimestamp int
		var toFormattedValue *float64

		if bridgeEvent.TTokenDecimal != nil {
			toFormattedValue = getAdjustedValue(bridgeEvent.TAmount, *bridgeEvent.TTokenDecimal)
		} else {
			return nil, fmt.Errorf("token decimal is not valid")
		}
		if bridgeEvent.TTimeStamp != nil {
			toTimestamp = int(*bridgeEvent.TTimeStamp)
		} else {
			return nil, fmt.Errorf("timestamp is not valid")
		}
		toInfos = &model.PartialInfo{
			ChainID:        &toChainID,
			Address:        &bridgeEvent.TRecipient.String,
			TxnHash:        &bridgeEvent.TTxHash,
			Value:          &toValue,
			FormattedValue: toFormattedValue,
			USDValue:       bridgeEvent.TAmountUSD,
			TokenAddress:   &bridgeEvent.TToken,
			TokenSymbol:    &bridgeEvent.TTokenSymbol.String,
			BlockNumber:    &toBlockNumber,
			Time:           &toTimestamp,
		}
	} else {
		toInfos = nil
		pending = true
	}

	var swapSuccess bool
	if bridgeEvent.TSwapSuccess.Uint64() == 1 {
		swapSuccess = true
	}
	if !includePending && pending {
		// nolint:nilnil
		return nil, nil
	}
	kappa := bridgeEvent.FDestinationKappa
	if kappa == "" {
		kappa = bridgeEvent.TKappa.String
	}
	bridgeTx = model.BridgeTransaction{
		FromInfo:    fromInfos,
		ToInfo:      toInfos,
		Kappa:       &kappa,
		Pending:     &pending,
		SwapSuccess: &swapSuccess,
	}
	return &bridgeTx, nil
}

// generatePartialInfoQuery returns the query for making the PartialInfo query.
//
// nolint:dupl
func generateAllBridgeEventsQueryFromOrigin(chainID *int, address, tokenAddress, txHash *string, page int, in bool) string {
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionSpecifier := generateDirectionSpecifierSQL(in, &firstFilter, "")
	compositeFilters := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + txHashSpecifier + directionSpecifier
	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	finalQuery := fmt.Sprintf("%s SELECT %s FROM %s %s", generateDeDepQueryCTE(compositeFilters, &pageValue, &pageOffset, true), originToDestCol, "baseQuery", originToDestJoins)

	return finalQuery
}

func generateAllBridgeEventsQueryFromDestination(chainID *int, address, tokenAddress, kappa, txHash *string, page int, in bool) string {
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	kappaSpecifier := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionSpecifier := generateDirectionSpecifierSQL(in, &firstFilter, "")
	compositeFilters := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + directionSpecifier
	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	finalQuery := fmt.Sprintf("%s SELECT %s FROM %s %s", generateDeDepQueryCTE(compositeFilters, &pageValue, &pageOffset, false), destToOriginCol, "baseQuery", destToOriginJoins)

	return finalQuery
}

func generateAllBridgeEventsQueryFromDestination2(chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, tokenAddress []*string, kappa *string, txHash *string, page *int, in bool) string {
	firstFilter := true

	chainIDSpecifier := generateSingleSpecifierI32ArrSQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringArrSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")

	minTimeSpecfier := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "", true)
	maxTimeSpecfier := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "", false)

	minAmountSpecfier := generateEqualitySpecifierSQL(minAmount, sql.AmountUSDFieldName, &firstFilter, "", true)
	maxAmountSpecfier := generateEqualitySpecifierSQL(maxAmount, sql.AmountUSDFieldName, &firstFilter, "", false)

	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "")
	kappaSpecifier := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionSpecifier := generateDirectionSpecifierSQL(in, &firstFilter, "")
	compositeFilters := chainIDSpecifier + tokenAddressSpecifier + minTimeSpecfier + maxTimeSpecfier + minAmountSpecfier + maxAmountSpecfier + addressSpecifier + kappaSpecifier + txHashSpecifier + directionSpecifier
	pageValue := sql.PageSize
	pageOffset := (*page - 1) * sql.PageSize
	finalQuery := fmt.Sprintf("%s SELECT %s FROM %s %s", generateDeDepQueryCTE(compositeFilters, &pageValue, &pageOffset, false), destToOriginCol, "baseQuery", destToOriginJoins)
	//fmt.Println(compositeFilters)
	return finalQuery
}

// generateAllBridgeEventsQueryFromOrigin2 gets all the filters for query from origin.
//
// nolint:dupl
func generateAllBridgeEventsQueryFromOrigin2(chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, tokenAddress []*string, txHash *string, page *int, in bool) string {
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32ArrSQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringArrSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")

	minTimeSpecfier := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "", true)
	maxTimeSpecfier := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "", false)

	minAmountSpecfier := generateEqualitySpecifierSQL(minAmount, sql.AmountUSDFieldName, &firstFilter, "", true)
	maxAmountSpecfier := generateEqualitySpecifierSQL(maxAmount, sql.AmountUSDFieldName, &firstFilter, "", false)

	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionSpecifier := generateDirectionSpecifierSQL(in, &firstFilter, "")
	compositeFilters := chainIDSpecifier + tokenAddressSpecifier + minTimeSpecfier + maxTimeSpecfier + minAmountSpecfier + maxAmountSpecfier + addressSpecifier + txHashSpecifier + directionSpecifier
	pageValue := sql.PageSize
	pageOffset := (*page - 1) * sql.PageSize
	finalQuery := fmt.Sprintf("%s SELECT %s FROM %s %s", generateDeDepQueryCTE(compositeFilters, &pageValue, &pageOffset, true), originToDestCol, "baseQuery", originToDestJoins)
	//fmt.Println(compositeFilters)
	return finalQuery
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromDestination(ctx context.Context, chainID *int, address *string, txHash *string, kappa *string, page int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generateAllBridgeEventsQueryFromDestination(chainID, address, tokenAddress, kappa, txHash, page, false))

	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}
	if len(allBridgeEvents) == 0 {
		return nil, nil
	}

	// Iterate through all bridge events and return all partials
	for i := range allBridgeEvents {
		bridgeTx, err := GetPartialInfoFromBridgeEventHybrid(allBridgeEvents[i], false)
		if err != nil {
			return nil, fmt.Errorf("failed to get partial info from bridge event: %w", err)
		}
		if bridgeTx != nil {
			results = append(results, bridgeTx)
		}
	}
	return results, nil
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromDestination2(ctx context.Context, chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, txHash *string, kappa *string, page *int, tokenAddress []*string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generateAllBridgeEventsQueryFromDestination2(chainID, address, maxAmount, minAmount, startTime, endTime, tokenAddress, kappa, txHash, page, false))

	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}
	if len(allBridgeEvents) == 0 {
		return nil, nil
	}

	// Iterate through all bridge events and return all partials
	for i := range allBridgeEvents {
		bridgeTx, err := GetPartialInfoFromBridgeEventHybrid(allBridgeEvents[i], false)
		if err != nil {
			return nil, fmt.Errorf("failed to get partial info from bridge event: %w", err)
		}
		if bridgeTx != nil {
			results = append(results, bridgeTx)
		}
	}
	return results, nil
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromOrigin(ctx context.Context, chainID *int, address *string, txHash *string, page int, tokenAddress *string, pending bool, latest bool) ([]*model.BridgeTransaction, error) {
	var err error
	var chainMap = make(map[uint32]bool)
	var results []*model.BridgeTransaction
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generateAllBridgeEventsQueryFromOrigin(chainID, address, tokenAddress, txHash, page, true))

	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}
	if len(allBridgeEvents) == 0 {
		return nil, nil
	}

	// Iterate through all bridge events and return all partials
	for i := range allBridgeEvents {
		if latest && chainMap[allBridgeEvents[i].FChainID] {
			continue
		}

		bridgeTx, err := GetPartialInfoFromBridgeEventHybrid(allBridgeEvents[i], pending)
		if err != nil {
			return nil, fmt.Errorf("failed to get partial info from bridge event: %w", err)
		}
		if bridgeTx != nil {
			results = append(results, bridgeTx)
			chainMap[allBridgeEvents[i].FChainID] = true
		}
	}
	return results, nil
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromOrigin2(ctx context.Context, chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, txHash *string, page *int, tokenAddress []*string, pending bool, latest bool) ([]*model.BridgeTransaction, error) {
	var err error
	var chainMap = make(map[uint32]bool)
	var results []*model.BridgeTransaction
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generateAllBridgeEventsQueryFromOrigin2(chainID, address, maxAmount, minAmount, startTime, endTime, tokenAddress, txHash, page, true))

	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}
	if len(allBridgeEvents) == 0 {
		return nil, nil
	}

	// Iterate through all bridge events and return all partials
	for i := range allBridgeEvents {
		if latest && chainMap[allBridgeEvents[i].FChainID] {
			continue
		}

		bridgeTx, err := GetPartialInfoFromBridgeEventHybrid(allBridgeEvents[i], pending)
		if err != nil {
			return nil, fmt.Errorf("failed to get partial info from bridge event: %w", err)
		}
		if bridgeTx != nil {
			results = append(results, bridgeTx)
			chainMap[allBridgeEvents[i].FChainID] = true
		}
	}
	return results, nil
}

// getAdjustedValue gets the adjusted value.
func getAdjustedValue(amount *big.Int, decimals uint8) *float64 {
	decimalMultiplier := new(big.Float).SetInt(big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil))
	adjustedAmount := new(big.Float).Quo(new(big.Float).SetInt(amount), decimalMultiplier)
	trueAmountStr := adjustedAmount.SetMode(big.AwayFromZero).Text('f', 4)
	priceFloat, err := strconv.ParseFloat(trueAmountStr, 64)
	if err != nil {
		return nil
	}
	return &priceFloat
}
func keyGen(chainID string, kappa string) string {
	return fmt.Sprintf("%s-%s", chainID, kappa)
}

func GenerateDailyStatisticBridgeSQL(typeArg *model.DailyStatisticType, compositeFilters string) (*string, *string, error) {
	var subQuery string
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, FROM_UNIXTIME(%s, %s) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.AmountUSDFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", singleSideCol, "baseQuery", singleSideJoinsCTE)
		query = fmt.Sprintf("%s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
	case model.DailyStatisticTypeAddresses:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, FROM_UNIXTIME(%s, %s) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", singleSideCol, "baseQuery", singleSideJoinsCTE)
		query = fmt.Sprintf("%s SELECT toFloat64(uniq(%s, %s )) AS total FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), sql.ChainIDFieldName, sql.SenderFieldName, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.DailyStatisticTypeTransactions:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s)) AS total, FROM_UNIXTIME(%s, %s) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", singleSideCol, "baseQuery", singleSideJoinsCTE)
		query = fmt.Sprintf(" %s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
	case model.DailyStatisticTypeFee:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s)) AS total, FROM_UNIXTIME(%s, %s) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%m/%d/%Y'", singleSideCol, "baseQuery", singleSideJoinsCTE)
		query = fmt.Sprintf(" %s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)

	default:
		return nil, nil, fmt.Errorf("invalid type argument")
	}
	return &subQuery, &query, nil
}

func GenerateDailyStatisticSwapSQL(typeArg *model.DailyStatisticType, compositeFilters string) (*string, *string, error) {
	var subQuery string
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", swapVolumeSelect, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
		query = fmt.Sprintf("%s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
	case model.DailyStatisticTypeAddresses:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	case model.DailyStatisticTypeTransactions:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	case model.DailyStatisticTypeFee:
		subQuery = fmt.Sprintf("SELECT sumKahan(arraySum(mapValues(%s))) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.AdminFeeUSDFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)

	default:
		return nil, nil, fmt.Errorf("invalid type argument")
	}
	return &subQuery, &query, nil
}

func GenerateDailyStatisticMessageBusSQL(typeArg *model.DailyStatisticType, compositeFilters string) (*string, *string, error) {
	var subQuery string
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		return nil, nil, nil
	case model.DailyStatisticTypeAddresses:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseMessageBus, compositeFilters)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	case model.DailyStatisticTypeTransactions:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseMessageBus, compositeFilters)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	case model.DailyStatisticTypeFee:
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.FeeUSDFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseMessageBus, compositeFilters)
		query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
	default:
		return nil, nil, fmt.Errorf("invalid type argument")
	}
	return &subQuery, &query, nil
}

func GenerateAmountStatisticBridgeSQL(typeArg model.StatisticType, compositeFilters string) (*string, error) {
	var operation string
	var finalSQL string
	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		operation = fmt.Sprintf("AVG(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("%s SELECT %s FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), operation, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.StatisticTypeMedianVolumeUsd:
		operation = fmt.Sprintf("median(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("%s SELECT %s FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), operation, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.StatisticTypeTotalVolumeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("%s SELECT %s FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), operation, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, generateDeDepQuery(compositeFilters, nil, nil))
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, generateDeDepQuery(compositeFilters, nil, nil))
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseBridge)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseBridge)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseBridge)
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	return &finalSQL, nil
}
func GenerateAmountStatisticSwapSQL(typeArg model.StatisticType, compositeFilters string) (*string, error) {
	var operation string
	var finalSQL string
	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		operation = fmt.Sprintf("AVG(%s)", swapVolumeSelect)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	case model.StatisticTypeMedianVolumeUsd:
		operation = fmt.Sprintf("median(%s)", swapVolumeSelect)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	case model.StatisticTypeTotalVolumeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", swapVolumeSelect)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseSwap)
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	return &finalSQL, nil
}

func GenerateAmountStatisticMessageBusSQL(typeArg model.StatisticType, compositeFilters string) (*string, error) {
	var operation string
	var finalSQL string
	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		return nil, nil
	case model.StatisticTypeMedianVolumeUsd:
		return nil, nil
	case model.StatisticTypeTotalVolumeUsd:
		return nil, nil
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseMessageBus)
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseMessageBus)
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseMessageBus)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseMessageBus)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, baseMessageBus)
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	return &finalSQL, nil
}

type SortBridgeTxType []*model.BridgeTransaction

func (s SortBridgeTxType) Len() int           { return len(s) }
func (s SortBridgeTxType) Less(i, j int) bool { return *s[i].FromInfo.Time > *s[j].FromInfo.Time }
func (s SortBridgeTxType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
