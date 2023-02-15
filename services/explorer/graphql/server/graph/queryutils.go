package graph

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/contracts/user"
	"golang.org/x/sync/errgroup"

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
	minTimestamp := " (SELECT min(timestamp) - 86400 FROM baseQuery) AS minTimestamp, (SELECT count(*) FROM baseQuery) AS rowCount"
	if in {
		minTimestamp = " (SELECT min(timestamp) FROM baseQuery) AS minTimestamp, (SELECT count(*) FROM baseQuery) AS rowCount"
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
	var fromTimeStampFormatted string
	if bridgeEvent.FTokenDecimal != nil {
		fromFormattedValue = getAdjustedValue(bridgeEvent.FAmount, *bridgeEvent.FTokenDecimal)
	} else {
		return nil, fmt.Errorf("token decimal is not valid")
	}
	if bridgeEvent.FTimeStamp != nil {
		fromTimestamp = int(*bridgeEvent.FTimeStamp)
		fromTimeStampFormatted = time.Unix(int64(*bridgeEvent.FTimeStamp), 0).String()
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
		FormattedTime:  &fromTimeStampFormatted,
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
		var toTimeStampFormatted string
		if bridgeEvent.TTokenDecimal != nil {
			toFormattedValue = getAdjustedValue(bridgeEvent.TAmount, *bridgeEvent.TTokenDecimal)
		} else {
			return nil, fmt.Errorf("token decimal is not valid")
		}
		if bridgeEvent.TTimeStamp != nil {
			toTimestamp = int(*bridgeEvent.TTimeStamp)
			toTimeStampFormatted = time.Unix(int64(*bridgeEvent.TTimeStamp), 0).String()
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
			FormattedTime:  &toTimeStampFormatted,
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

func generateMessageBusQuery(chainID []*int, address *string, startTime *int, endTime *int, messageID *string, pending bool, reverted bool, txHash *string, page int) string {
	firstFilter := true

	chainIDSpecifier := generateSingleSpecifierI32ArrSQL(chainID, sql.ChainIDFieldName, &firstFilter, "")

	minTimeSpecfier := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "", true)
	maxTimeSpecfier := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "", false)

	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "")
	messageIDSpecifier := generateSingleSpecifierStringSQL(messageID, "message_id", &firstFilter, "")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	operation := " = ''"
	if !pending {
		operation = " != ''"
	}
	pendingSpecifier := fmt.Sprintf(" WHERE t.message_id %s", operation)
	compositeFilters := chainIDSpecifier + minTimeSpecfier + maxTimeSpecfier + addressSpecifier + messageIDSpecifier + txHashSpecifier
	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize

	cte := fmt.Sprintf("WITH baseQuery AS (SELECT * FROM message_bus_events %s ORDER BY timestamp DESC, block_number DESC, event_index DESC, insert_time DESC LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash), (SELECT min(timestamp) FROM baseQuery) AS minTimestamp", compositeFilters)

	finalQuery := fmt.Sprintf("%s SELECT * FROM (SELECT * FROM (SELECT * FROM %s WHERE %s = 1 ) f LEFT JOIN (SELECT * FROM (%s) WHERE %s = 0) t ON f.%s = t.%s %s)  LIMIT %d OFFSET %d", cte, "baseQuery", sql.EventTypeFieldName, baseMessageBus, sql.EventTypeFieldName, "message_id", "message_id", pendingSpecifier, pageValue, pageOffset)

	if reverted {
		finalQuery = fmt.Sprintf("%s SELECT * FROM  (SELECT * FROM (select * from (%s) WHERE %s = 1) f RIGHT OUTER JOIN (Select r.reverted_reason AS reverted_reason, j.reverted_reason AS rrr, * FROM (select * from %s WHERE event_type = 0 and status = 'Fail') j LEFT JOIN (select reverted_reason, tx_hash from (%s) WHERE %s = 2) r on j.tx_hash = r.tx_hash) t ON f.%s = t.%s)  LIMIT %d OFFSET %d", cte, baseMessageBus, sql.EventTypeFieldName, "baseQuery", baseMessageBus, sql.EventTypeFieldName, "message_id", "message_id", pageValue, pageOffset)
	}
	return finalQuery
}
func generateAllBridgeEventsQueryFromDestination2(chainIDTo []*int, chainIDFrom []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, tokenAddressFrom []*string, tokenAddressTo []*string, kappa *string, txHash *string, page int, in bool) string {
	firstFilter := true
	chainIDToFilter := generateSingleSpecifierI32ArrSQL(chainIDTo, sql.ChainIDFieldName, &firstFilter, "")
	minTimeFilter := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "", true)
	maxTimeFilter := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "", false)
	addressToFilter := generateAddressSpecifierSQL(addressTo, &firstFilter, "")
	kappaFilter := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "")
	txHashFilter := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionFilter := generateDirectionSpecifierSQL(in, &firstFilter, "")

	toFilters := chainIDToFilter + minTimeFilter + maxTimeFilter + addressToFilter + kappaFilter + txHashFilter + directionFilter

	firstFilter = false
	chainIDFromFilter := generateSingleSpecifierI32ArrSQL(chainIDFrom, sql.ChainIDFieldName, &firstFilter, "")
	addressFromFilter := generateAddressSpecifierSQL(addressFrom, &firstFilter, "")

	fromFilters := chainIDFromFilter + addressFromFilter

	// TODO change these feuld names
	firstFilter = true
	minAmountFilter := generateEqualitySpecifierSQL(minAmount, "tamount", &firstFilter, "", true)
	minAmountFilterUsd := generateEqualitySpecifierSQL(minAmountUsd, "tamount_usd", &firstFilter, "", true)
	maxAmountFilter := generateEqualitySpecifierSQL(maxAmount, "famount", &firstFilter, "", false)
	maxAmountFilterUsd := generateEqualitySpecifierSQL(maxAmountUsd, "famount_usd", &firstFilter, "", false)
	tokenAddressToFilter := generateSingleSpecifierStringArrSQL(tokenAddressTo, "ttoken", &firstFilter, "")
	tokenAddressFromFilter := generateSingleSpecifierStringArrSQL(tokenAddressFrom, "ftoken", &firstFilter, "")
	postJoinFilters := minAmountFilter + minAmountFilterUsd + maxAmountFilter + maxAmountFilterUsd + tokenAddressToFilter + tokenAddressFromFilter

	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	if postJoinFilters == "" {
		return fmt.Sprintf("%s SELECT %s FROM %s %s %s %s", generateDeDepQueryCTE(toFilters, &pageValue, &pageOffset, false), destToOriginCol, "baseQuery", destToOriginJoinsPt1, fromFilters, destToOriginJoinsPt2)
	}
	return fmt.Sprintf("%s SELECT * FROM (SELECT %s FROM %s %s %s %s) %s LIMIT %d OFFSET %d", generateDeDepQueryCTE(toFilters, nil, nil, false), destToOriginCol, "baseQuery", destToOriginJoinsPt1, fromFilters, destToOriginJoinsPt2, postJoinFilters, pageValue, pageOffset)
}

func generateAllBridgeEventsQueryFromDestination(chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, tokenAddress []*string, kappa *string, txHash *string, page int, in bool) string {
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
	pageOffset := (page - 1) * sql.PageSize
	finalQuery := fmt.Sprintf("%s SELECT %s FROM %s %s", generateDeDepQueryCTE(compositeFilters, &pageValue, &pageOffset, false), destToOriginCol, "baseQuery", destToOriginJoins)
	return finalQuery
}

// generateAllBridgeEventsQueryFromOrigin gets all the filters for query from origin.
//
// nolint:dupl
func generateAllBridgeEventsQueryFromOrigin(chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, tokenAddress []*string, txHash *string, page int, in bool) string {
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
	pageOffset := (page - 1) * sql.PageSize
	finalQuery := fmt.Sprintf("%s SELECT %s FROM %s %s", generateDeDepQueryCTE(compositeFilters, &pageValue, &pageOffset, true), originToDestCol, "baseQuery", originToDestJoins)
	return finalQuery
}

func generateAllBridgeEventsQueryFromOrigin2(chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, tokenAddressFrom []*string, tokenAddressTo []*string, txHash *string, page int, pending bool, in bool) string {
	firstFilter := true
	chainIDFromFilter := generateSingleSpecifierI32ArrSQL(chainIDFrom, sql.ChainIDFieldName, &firstFilter, "")
	minTimeFilter := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "", true)
	maxTimeFilter := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "", false)
	addressFromFilter := generateAddressSpecifierSQL(addressFrom, &firstFilter, "")
	txHashFilter := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionFilter := generateDirectionSpecifierSQL(in, &firstFilter, "")

	fromFilters := chainIDFromFilter + minTimeFilter + maxTimeFilter + addressFromFilter + txHashFilter + directionFilter

	firstFilter = false
	chainIDToFilter := generateSingleSpecifierI32ArrSQL(chainIDTo, sql.ChainIDFieldName, &firstFilter, "")
	addressToFilter := generateAddressSpecifierSQL(addressTo, &firstFilter, "")

	toFilters := chainIDToFilter + addressToFilter

	firstFilter = false
	minAmountFilter := generateEqualitySpecifierSQL(minAmount, "tamount", &firstFilter, "", true)
	minAmountFilterUsd := generateEqualitySpecifierSQL(minAmountUsd, "tamount_usd", &firstFilter, "", true)
	maxAmountFilter := generateEqualitySpecifierSQL(maxAmount, "famount", &firstFilter, "", false)
	maxAmountFilterUsd := generateEqualitySpecifierSQL(maxAmountUsd, "famount_usd", &firstFilter, "", false)
	tokenAddressToFilter := generateSingleSpecifierStringArrSQL(tokenAddressTo, "ttoken", &firstFilter, "")
	tokenAddressFromFilter := generateSingleSpecifierStringArrSQL(tokenAddressFrom, "ftoken", &firstFilter, "")

	operation := " = ''"
	if !pending {
		operation = " != ''"
	}
	pendingFilter := fmt.Sprintf(" WHERE t%s %s", sql.KappaFieldName, operation)
	postJoinFilters := minAmountFilter + minAmountFilterUsd + maxAmountFilter + maxAmountFilterUsd + tokenAddressToFilter + tokenAddressFromFilter

	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	if !pending && postJoinFilters == "" {
		return fmt.Sprintf("%s SELECT %s FROM %s %s %s %s", generateDeDepQueryCTE(fromFilters, &pageValue, &pageOffset, false), originToDestCol, "baseQuery", originToDestJoinsPt1, toFilters, originToDestJoinsPt2)
	}
	return fmt.Sprintf("%s SELECT * FROM (SELECT %s FROM %s %s %s %s) %s LIMIT %d OFFSET %d", generateDeDepQueryCTE(fromFilters, nil, nil, false), originToDestCol, "baseQuery", originToDestJoinsPt1, toFilters, originToDestJoinsPt2, pendingFilter+postJoinFilters, pageValue, pageOffset)
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromDestination(ctx context.Context, chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, txHash *string, kappa *string, page *int, tokenAddress []*string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generateAllBridgeEventsQueryFromDestination(chainID, address, maxAmount, minAmount, startTime, endTime, tokenAddress, kappa, txHash, *page, false))

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

func (r *queryResolver) GetBridgeTxsFromDestination2(ctx context.Context, chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txHash *string, kappa *string, page *int, tokenAddressFrom []*string, tokenAddressTo []*string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	query := generateAllBridgeEventsQueryFromDestination2(chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, minAmountUsd, maxAmountUsd, startTime, endTime, tokenAddressFrom, tokenAddressTo, kappa, txHash, *page, false)

	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, query)

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

//
//// nolint:gocognit,cyclop
// func (r *queryResolver) GetBridgeTxsFromOrigin(ctx context.Context, chainID []*int, address *string, maxAmount *int, minAmount *int, startTime *int, endTime *int, txHash *string, page *int, tokenAddress []*string, pending bool, latest bool) ([]*model.BridgeTransaction, error) {
//	var err error
//	var chainMap = make(map[uint32]bool)
//	var results []*model.BridgeTransaction
//	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generateAllBridgeEventsQueryFromOrigin2(chainID, address, maxAmount, minAmount, startTime, endTime, tokenAddress, txHash, *page, true))
//
//	if err != nil {
//		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
//	}
//	if len(allBridgeEvents) == 0 {
//		return nil, nil
//	}
//
//	// Iterate through all bridge events and return all partials
//	for i := range allBridgeEvents {
//		if latest && chainMap[allBridgeEvents[i].FChainID] {
//			continue
//		}
//
//		bridgeTx, err := GetPartialInfoFromBridgeEventHybrid(allBridgeEvents[i], pending)
//		if err != nil {
//			return nil, fmt.Errorf("failed to get partial info from bridge event: %w", err)
//		}
//		if bridgeTx != nil {
//			results = append(results, bridgeTx)
//			chainMap[allBridgeEvents[i].FChainID] = true
//		}
//	}
//	return results, nil
//}

func (r *queryResolver) GetBridgeTxsFromOrigin2(ctx context.Context, chainIDTo []*int, chainIDFrom []*int, addressTo *string, addressFrom *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txHash *string, page *int, tokenAddressTo []*string, tokenAddressFrom []*string, pending bool, latest bool) ([]*model.BridgeTransaction, error) {
	var err error
	var chainMap = make(map[uint32]bool)
	var results []*model.BridgeTransaction
	query := generateAllBridgeEventsQueryFromOrigin2(chainIDTo, chainIDFrom, addressTo, addressFrom, maxAmount, minAmount, maxAmountUsd, minAmountUsd, startTime, endTime, tokenAddressTo, tokenAddressFrom, txHash, *page, pending, true)
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, query)

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

// GetPartialInfoFromMessageBusEventHybrid returns the partial info from message bus event.
//
// nolint:cyclop
func GetPartialInfoFromMessageBusEventHybrid(ctx context.Context, messageBusEvent sql.HybridMessageBusEvent, pending bool) (*model.MessageBusTransaction, error) {
	var messageBusTx model.MessageBusTransaction
	fromChainID := int(messageBusEvent.FChainID)
	fromDestinationChainID := int(messageBusEvent.FDestinationChainID.Uint64())
	fromBlockNumber := int(messageBusEvent.FBlockNumber)
	fromTimeStamp := int(*messageBusEvent.FTimeStamp)
	fromTimeStampFormatted := time.Unix(int64(*messageBusEvent.FTimeStamp), 0).String()

	toChainID := int(messageBusEvent.TChainID)
	toBlockNumber := int(messageBusEvent.TBlockNumber)
	toTimeStamp := int(*messageBusEvent.TTimeStamp)
	toTimeStampFormatted := time.Unix(int64(*messageBusEvent.TTimeStamp), 0).String()

	fromInfos := &model.PartialMessageBusInfo{
		ChainID:            &fromChainID,
		DestinationChainID: &fromDestinationChainID,
		ContractAddress:    &messageBusEvent.FContractAddress,
		TxnHash:            &messageBusEvent.FTxHash,
		Message:            &messageBusEvent.FMessage.String,
		BlockNumber:        &fromBlockNumber,
		Time:               &fromTimeStamp,
		FormattedTime:      &fromTimeStampFormatted,
		RevertedReason:     nil,
	}

	toInfos := &model.PartialMessageBusInfo{
		ChainID:            &toChainID,
		DestinationChainID: nil,
		ContractAddress:    &messageBusEvent.TContractAddress,
		TxnHash:            &messageBusEvent.TTxHash,
		Message:            &messageBusEvent.TMessage.String,
		BlockNumber:        &toBlockNumber,
		Time:               &toTimeStamp,
		FormattedTime:      &toTimeStampFormatted,
		RevertedReason:     &messageBusEvent.TRevertedReason.String,
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		fromInfos.MessageType = user.Decode(ctx, messageBusEvent.FMessage.String)
	}()

	go func() {
		defer wg.Done()
		toInfos.MessageType = user.Decode(ctx, messageBusEvent.TMessage.String)
	}()

	wg.Wait()

	messageBusTx = model.MessageBusTransaction{
		FromInfo:  fromInfos,
		ToInfo:    toInfos,
		MessageID: &messageBusEvent.FMessageID.String,
		Pending:   &pending,
	}
	return &messageBusTx, nil
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetMessageBusTxs(ctx context.Context, chainID []*int, address *string, startTime *int, endTime *int, txHash *string, messageID *string, pending bool, reverted bool, page *int) ([]*model.MessageBusTransaction, error) {
	var err error
	allMessageBusEvents, err := r.DB.GetAllMessageBusEvents(ctx, generateMessageBusQuery(chainID, address, startTime, endTime, messageID, pending, reverted, txHash, *page))
	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}

	if len(allMessageBusEvents) == 0 {
		return nil, nil
	}

	results := make([]*model.MessageBusTransaction, len(allMessageBusEvents))
	var sliceMux sync.Mutex
	g, ctx := errgroup.WithContext(ctx)
	// Iterate through all bridge events and return all partials
	for i := range allMessageBusEvents {
		i := i // capture func literal
		g.Go(func() error {
			messageBusTx, err := GetPartialInfoFromMessageBusEventHybrid(ctx, allMessageBusEvents[i], pending)
			if err != nil {
				return fmt.Errorf("failed to get partial info from bridge event: %w", err)
			}
			if messageBusTx != nil {
				sliceMux.Lock()
				results[i] = messageBusTx
				sliceMux.Unlock()
			}

			return nil
		})
	}
	err = g.Wait()

	if err != nil {
		return nil, fmt.Errorf("could not get partial info from message bus event: %w", err)
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

func GenerateDailyStatisticBridgeSQL(typeArg *model.DailyStatisticType, compositeFilters string, firstFilter *bool) (*string, *string, error) {
	var subQuery string
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		compositeFilters += directionSpecifier
		subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.AmountUSDFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
		query = fmt.Sprintf("%s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
	case model.DailyStatisticTypeAddresses:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
		query = fmt.Sprintf("%s SELECT toFloat64(uniq(%s, %s )) AS total FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), sql.ChainIDFieldName, sql.SenderFieldName, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.DailyStatisticTypeTransactions:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s)) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
		query = fmt.Sprintf(" %s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
	case model.DailyStatisticTypeFee:
		subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s)) AS total, FtoDate(ROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
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
		subQuery = fmt.Sprintf("SELECT sumKahan(arraySum(mapValues(%s))) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.FeeUSDFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
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

func GenerateAmountStatisticBridgeSQL(typeArg model.StatisticType, compositeFilters string, firstFilter *bool) (*string, error) {
	var operation string
	var finalSQL string
	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		compositeFilters += directionSpecifier
		operation = fmt.Sprintf("AVG(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("%s SELECT %s FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), operation, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.StatisticTypeMedianVolumeUsd:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		compositeFilters += directionSpecifier
		operation = fmt.Sprintf("median(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("%s SELECT %s FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), operation, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.StatisticTypeTotalVolumeUsd:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		compositeFilters += directionSpecifier
		operation = fmt.Sprintf("sumKahan(%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("%s SELECT %s FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), operation, singleSideCol, "baseQuery", singleSideJoinsCTE)
	case model.StatisticTypeCountTransactions:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		compositeFilters += directionSpecifier
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, generateDeDepQuery(compositeFilters, nil, nil))
	case model.StatisticTypeCountAddresses:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		compositeFilters += directionSpecifier
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s)", operation, generateDeDepQuery(compositeFilters, nil, nil))
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseBridge, compositeFilters)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseBridge, compositeFilters)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseBridge, compositeFilters)

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
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
	case model.StatisticTypeMedianVolumeUsd:
		operation = fmt.Sprintf("median(%s)", swapVolumeSelect)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
	case model.StatisticTypeTotalVolumeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", swapVolumeSelect)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseSwap, compositeFilters)
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
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseMessageBus, compositeFilters)
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(%s, source_address) AS res", sql.ChainIDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseMessageBus, compositeFilters)
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseMessageBus, compositeFilters)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseMessageBus, compositeFilters)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", sql.FeeUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s) %s", operation, baseMessageBus, compositeFilters)
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	return &finalSQL, nil
}

//	func GenerateDailyStatisticBridgeSQL(typeArg *model.DailyStatisticType, compositeFilters string, firstFilter *bool) (*string, *string, error) {
//		var subQuery string
//		var query string
//		switch *typeArg {
//		case model.DailyStatisticTypeVolume:
//			directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
//			compositeFilters += directionSpecifier
//			subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.AmountUSDFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
//			query = fmt.Sprintf("%s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
//		case model.DailyStatisticTypeAddresses:
//			subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
//			query = fmt.Sprintf("%s SELECT toFloat64(uniq(%s, %s )) AS total FROM (SELECT %s FROM %s %s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), sql.ChainIDFieldName, sql.SenderFieldName, singleSideCol, "baseQuery", singleSideJoinsCTE)
//		case model.DailyStatisticTypeTransactions:
//			subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s)) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
//			query = fmt.Sprintf(" %s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
//		case model.DailyStatisticTypeFee:
//			subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s)) AS total, FtoDate(ROM_UNIXTIME(%s, %s)) AS date FROM (SELECT %s FROM %s %s) GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", singleSideCol, "baseQuery", singleSideJoinsCTE)
//			query = fmt.Sprintf(" %s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
//
//		default:
//			return nil, nil, fmt.Errorf("invalid type argument")
//		}
//		return &subQuery, &query, nil
//	}
//
//	func GenerateDailyStatisticSwapSQL(typeArg *model.DailyStatisticType, compositeFilters string) (*string, *string, error) {
//		var subQuery string
//		var query string
//		switch *typeArg {
//		case model.DailyStatisticTypeVolume:
//			subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", swapVolumeSelect, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
//			query = fmt.Sprintf("%s SELECT sumKahan(total) FROM (%s)", generateDeDepQueryCTE(compositeFilters, nil, nil, true), subQuery)
//		case model.DailyStatisticTypeAddresses:
//			subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
//			query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
//		case model.DailyStatisticTypeTransactions:
//			subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
//			query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
//		case model.DailyStatisticTypeFee:
//			subQuery = fmt.Sprintf("SELECT sumKahan(arraySum(mapValues(%s))) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.FeeUSDFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseSwap, compositeFilters)
//			query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
//
//		default:
//			return nil, nil, fmt.Errorf("invalid type argument")
//		}
//		return &subQuery, &query, nil
//	}
//
//	func GenerateDailyStatisticMessageBusSQL(typeArg *model.DailyStatisticType, compositeFilters string) (*string, *string, error) {
//		var subQuery string
//		var query string
//		switch *typeArg {
//		case model.DailyStatisticTypeVolume:
//			return nil, nil, nil
//		case model.DailyStatisticTypeAddresses:
//			subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.SenderFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseMessageBus, compositeFilters)
//			query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
//		case model.DailyStatisticTypeTransactions:
//			subQuery = fmt.Sprintf("SELECT toFloat64(uniq(%s, %s )) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.ChainIDFieldName, sql.TxHashFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseMessageBus, compositeFilters)
//			query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
//		case model.DailyStatisticTypeFee:
//			subQuery = fmt.Sprintf("SELECT sumKahan(%s) AS total, toDate(FROM_UNIXTIME(%s, %s)) AS date FROM (%s) %s GROUP BY date ORDER BY date ASC", sql.FeeUSDFieldName, sql.TimeStampFieldName, "'%Y/%m/%d'", baseMessageBus, compositeFilters)
//			query = fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)
//		default:
//			return nil, nil, fmt.Errorf("invalid type argument")
//		}
//		return &subQuery, &query, nil
//	}
func GenerateRankedChainsByVolumeSQL(compositeFilters string, firstFilter *bool) string {
	directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
	return fmt.Sprintf("%s %s FULL OUTER JOIN (SELECT chain_id, sumKahan(multiIf(event_type = 0, amount_usd[sold_id], event_type = 1, arraySum(mapValues(amount_usd)), event_type = 9, arraySum(mapValues(amount_usd)), event_type = 10, amount_usd[sold_id], 0)) as usdTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by chain_id) s ON b.pre_fchain_id = s.chain_id ORDER BY total DESC SETTINGS join_use_nulls = 1", generateDeDepQueryCTE(compositeFilters+directionSpecifier, nil, nil, true), rankedChainsBridgeVolume, compositeFilters)
}

func GenerateDailyStatisticByChainAllSQL(typeArg *model.DailyStatisticType, compositeFilters string, firstFilter *bool) (*string, error) {
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		query = fmt.Sprintf("%s %s FULL OUTER JOIN (SELECT %s, chain_id, sumKahan(multiIf(event_type = 0, amount_usd[sold_id], event_type = 1,    arraySum(mapValues(amount_usd)), event_type = 9,    arraySum(mapValues(amount_usd)), event_type = 10, amount_usd[sold_id],    0) )     as usdTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.pre_fchain_id = s.chain_id) group by date order by date) SETTINGS join_use_nulls=1", generateDeDepQueryCTE(compositeFilters+directionSpecifier, nil, nil, true), dailyVolumeBridge, toDateSelect, compositeFilters)
	case model.DailyStatisticTypeFee:
		query = fmt.Sprintf("%s FROM ( SELECT %s, chain_id, sumKahan(fee_usd) as sumTotal FROM (SELECT * FROM bridge_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) GROUP BY date, chain_id) b  FULL OUTER JOIN ( SELECT %s, chain_id, sumKahan(arraySum(mapValues(fee_usd))) AS sumTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.chain_id = s.chain_id  FULL OUTER JOIN ( SELECT %s, chain_id, sumKahan(fee_usd) AS sumTotal FROM (SELECT * FROM message_bus_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) m ON b.date = m.date AND b.chain_id = m.chain_id) group by date order by date ) SETTINGS join_use_nulls = 1", dailyStatisticGenericSelect, toDateSelect, compositeFilters, toDateSelect, compositeFilters, toDateSelect, compositeFilters)
	case model.DailyStatisticTypeAddresses:
		query = fmt.Sprintf("%s FROM ( SELECT %s, chain_id, uniq(chain_id, sender) as sumTotal FROM (SELECT * FROM bridge_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) GROUP BY date, chain_id) b  FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, sender) AS sumTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.chain_id = s.chain_id  FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, source_address) AS sumTotal FROM (SELECT * FROM message_bus_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) m ON b.date = m.date AND b.chain_id = m.chain_id) group by date order by date ) SETTINGS join_use_nulls = 1", dailyStatisticGenericSelect, toDateSelect, compositeFilters, toDateSelect, compositeFilters, toDateSelect, compositeFilters)
	case model.DailyStatisticTypeTransactions:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		query = fmt.Sprintf("%s FROM ( SELECT %s, chain_id, uniq(chain_id, tx_hash) as sumTotal FROM (SELECT * FROM bridge_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) GROUP BY date, chain_id) b  FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, tx_hash) AS sumTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.chain_id = s.chain_id  FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, tx_hash) AS sumTotal FROM (SELECT * FROM message_bus_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) m ON b.date = m.date AND b.chain_id = m.chain_id) group by date order by date ) SETTINGS join_use_nulls = 1", dailyStatisticGenericSelect, toDateSelect, compositeFilters+directionSpecifier, toDateSelect, compositeFilters, toDateSelect, compositeFilters)
	default:
		return nil, fmt.Errorf("unsupported statistic type")
	}
	return &query, nil
}

// TODO make this more dynamic.
func GenerateDailyStatisticByChainBridgeSQL(typeArg *model.DailyStatisticType, compositeFilters string, firstFilter *bool) (*string, error) {
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		query = fmt.Sprintf("%s  %s sumKahan(amount_usd) AS sumTotal %s group by date, chain_id order by date, chain_id) group by date order by date )", generateDeDepQueryCTE(compositeFilters+directionSpecifier, nil, nil, true), dailyStatisticGenericSinglePlatform, dailyStatisticBridge)
	case model.DailyStatisticTypeFee:
		query = fmt.Sprintf("%s  %s sumKahan(fee_usd) AS sumTotal %s group by date, chain_id order by date, chain_id) group by date order by date )", generateDeDepQueryCTE(compositeFilters, nil, nil, true), dailyStatisticGenericSinglePlatform, dailyStatisticBridge)
	case model.DailyStatisticTypeAddresses:
		query = fmt.Sprintf("%s  %s uniq(chain_id, sender) AS sumTotal %s group by date, chain_id order by date, chain_id) group by date order by date )", generateDeDepQueryCTE(compositeFilters, nil, nil, true), dailyStatisticGenericSinglePlatform, dailyStatisticBridge)
	case model.DailyStatisticTypeTransactions:
		directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
		query = fmt.Sprintf("%s %s uniq(chain_id, tx_hash) AS sumTotal  %s group by date, chain_id order by date, chain_id) group by date order by date )", generateDeDepQueryCTE(compositeFilters+directionSpecifier, nil, nil, true), dailyStatisticGenericSinglePlatform, dailyStatisticBridge)
	default:
		return nil, fmt.Errorf("unsupported statistic type")
	}
	return &query, nil
}

func GenerateDailyStatisticByChainSwapSQL(typeArg *model.DailyStatisticType, compositeFilters string) (*string, error) {
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		query = fmt.Sprintf("%s sumKahan(multiIf(event_type = 0, amount_usd[sold_id], event_type = 1, arraySum(mapValues(amount_usd)), event_type = 9, arraySum(mapValues(amount_usd)), event_type = 10, amount_usd[sold_id],0)) AS sumTotal FROM (%s) %s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, baseSwap, compositeFilters)
	case model.DailyStatisticTypeFee:
		query = fmt.Sprintf("%s sumKahan(arraySum(mapValues(%s))) AS sumTotal FROM (%s) %s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, sql.FeeUSDFieldName, baseSwap, compositeFilters)
	case model.DailyStatisticTypeAddresses:
		query = fmt.Sprintf("%s uniq(%s, %s) AS sumTotal FROM (%s) %s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, sql.ChainIDFieldName, sql.SenderFieldName, baseSwap, compositeFilters)
	case model.DailyStatisticTypeTransactions:
		query = fmt.Sprintf("%s uniq(%s, %s) AS sumTotal FROM (%s) %s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, sql.ChainIDFieldName, sql.TxHashFieldName, baseSwap, compositeFilters)
	default:
		return nil, fmt.Errorf("unsupported statistic type")
	}
	return &query, nil
}

func GenerateDailyStatisticByChainMessageBusSQL(typeArg *model.DailyStatisticType, compositeFilters string) (*string, error) {
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		return nil, fmt.Errorf("cannot calculate volume for messagebus")
	case model.DailyStatisticTypeFee:
		query = fmt.Sprintf("%s sumKahan(%s) AS sumTotal FROM (%s) %s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, sql.FeeUSDFieldName, baseMessageBus, compositeFilters)
	case model.DailyStatisticTypeAddresses:
		query = fmt.Sprintf("%s uniq(%s, %s) AS sumTotal FROM (%s)%s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, sql.ChainIDFieldName, sql.SenderFieldName, baseMessageBus, compositeFilters)
	case model.DailyStatisticTypeTransactions:
		query = fmt.Sprintf("%s uniq(%s, %s) AS sumTotal FROM (%s) %s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, sql.ChainIDFieldName, sql.SenderFieldName, baseMessageBus, compositeFilters)
	default:
		return nil, fmt.Errorf("unsupported statistic type")
	}
	return &query, nil
}

type SortBridgeTxType []*model.BridgeTransaction

func (s SortBridgeTxType) Len() int           { return len(s) }
func (s SortBridgeTxType) Less(i, j int) bool { return *s[i].FromInfo.Time > *s[j].FromInfo.Time }
func (s SortBridgeTxType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type SortMessageBusTxType []*model.MessageBusTransaction

func (s SortMessageBusTxType) Len() int           { return len(s) }
func (s SortMessageBusTxType) Less(i, j int) bool { return *s[i].FromInfo.Time > *s[j].FromInfo.Time }
func (s SortMessageBusTxType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func keyGenHandleNilInt(item *int) string {
	if item != nil {
		return fmt.Sprintf("%d", *item)
	}
	return ""
}

func keyGenHandleNilString(item *string) string {
	if item != nil {
		return fmt.Sprintf("%s", *item)
	}
	return ""
}
