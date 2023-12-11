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
	"github.com/synapsecns/sanguine/services/explorer/types/bridge"
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

// GetTargetTime converts the number of hours into a timestamp.
func GetTargetTime(hours *int) uint64 {
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
	// if address != nil {
	//	if *firstFilter {
	//		*firstFilter = false
	//
	//		return fmt.Sprintf(" WHERE (%s%s = '%s' OR  %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//	}
	//
	//	return fmt.Sprintf(" AND (%s%s = '%s' OR %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//}
	//
	// return ""
	if address != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, sql.SenderFieldName, *address)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, sql.SenderFieldName, *address)
	}

	return ""
}

// generateAddressSpecifierSQL generates a where function with an string.
//
// nolint:unparam
func generateAddressSpecifierSQLMv(address *string, firstFilter *bool, firstInLocale *bool, tablePrefix string) string {
	// if address != nil {
	//	if *firstFilter {
	//		*firstFilter = false
	//
	//		return fmt.Sprintf(" WHERE (%s%s = '%s' OR  %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//	}
	//
	//	return fmt.Sprintf(" AND (%s%s = '%s' OR %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//}
	//
	// return ""
	if address != nil {
		if *firstInLocale {
			*firstFilter = false
			*firstInLocale = false
			return fmt.Sprintf("  %s%s = '%s'", tablePrefix, sql.SenderFieldName, *address)
		}
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, sql.SenderFieldName, *address)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, sql.SenderFieldName, *address)
	}

	return ""
}

func generateRecipientSpecifierSQL(address *string, firstFilter *bool, tablePrefix string) string {
	// if address != nil {
	//	if *firstFilter {
	//		*firstFilter = false
	//
	//		return fmt.Sprintf(" WHERE (%s%s = '%s' OR  %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//	}
	//
	//	return fmt.Sprintf(" AND (%s%s = '%s' OR %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//}
	//
	// return ""
	if address != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, sql.RecipientFieldName, *address)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, sql.RecipientFieldName, *address)
	}

	return ""
}

func generateRecipientSpecifierSQLMv(address *string, firstFilter *bool, firstInLocale *bool, tablePrefix string) string {
	// if address != nil {
	//	if *firstFilter {
	//		*firstFilter = false
	//
	//		return fmt.Sprintf(" WHERE (%s%s = '%s' OR  %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//	}
	//
	//	return fmt.Sprintf(" AND (%s%s = '%s' OR %s%s = '%s')", tablePrefix, sql.RecipientFieldName, *address, tablePrefix, sql.SenderFieldName, *address)
	//}
	//
	// return ""
	if address != nil {
		if *firstInLocale {
			*firstFilter = false
			*firstInLocale = false
			return fmt.Sprintf(" %s%s = '%s'", tablePrefix, sql.RecipientFieldName, *address)
		}
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, sql.RecipientFieldName, *address)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, sql.RecipientFieldName, *address)
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

// generateCCTPSpecifierSQLMv generates a where function with event type to filter only cctp events.
func generateCCTPSpecifierSQL(onlyCctp *bool, to bool, field string, firstFilter *bool, tablePrefix string) string {
	if onlyCctp != nil && *onlyCctp {
		// From explorer/types/bridge/eventtypes.go
		eventType := 10
		if to {
			eventType = 11
		}

		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s =  %d", tablePrefix, field, eventType)
		}

		return fmt.Sprintf(" AND %s%s = %d", tablePrefix, field, eventType)
	}

	return ""
}

// generateEqualitySpecifierSQL generates a where function with an equality.
//
// nolint:unparam
func generateEqualitySpecifierSQLMv(value *int, field string, firstFilter *bool, firstInLocale *bool, tablePrefix string, greaterThan bool) string {
	operator := "<"
	if greaterThan {
		operator = ">"
	}
	if value != nil {
		if *firstInLocale {
			*firstFilter = false
			*firstInLocale = false
			return fmt.Sprintf(" %s%s %s %d", tablePrefix, field, operator, *value)
		}
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

		return fmt.Sprintf(" WHERE %s%s = 0", tablePrefix, sql.DestinationChainIDFieldName)
	}

	return fmt.Sprintf(" AND %s%s = 0", tablePrefix, sql.DestinationChainIDFieldName)
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
		final += whereString
	}

	for i := range values {
		final += fmt.Sprintf(" %s%s = %d", tablePrefix, field, *values[i])
		if i < len(values)-1 {
			final += orString
		}
	}

	return final + ")"
}

// generateSingleSpecifierI32ArrSQL generates a where function with an uint32.
//
// nolint:unparam
func generateSingleSpecifierI32ArrSQLMv(values []*int, field string, firstFilter *bool, firstInLocale *bool, tablePrefix string) string {
	if len(values) == 0 {
		return ""
	}
	var final string
	if *firstInLocale {
		*firstInLocale = false
		*firstFilter = false
		final += " ("
	} else if *firstFilter {
		*firstFilter = false
		final += whereString
	}
	for i := range values {
		final += fmt.Sprintf(" %s%s = %d", tablePrefix, field, *values[i])
		if i < len(values)-1 {
			final += orString
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
		final += whereString
	} else {
		final += " AND ("
	}

	for i := range values {
		if values[i] != nil {
			final += fmt.Sprintf(" %s%s = '%s'", tablePrefix, field, *values[i])
			if i < len(values)-1 {
				final += orString
			}
		}
	}

	return final + ")"
}

// GenerateSingleSpecifierStringSQL generates a where function with a string.
//
// nolint:unparam
func generateSingleSpecifierStringArrSQLMv(values []*string, field string, firstFilter *bool, firstInLocale *bool, tablePrefix string) string {
	if len(values) == 0 {
		return ""
	}
	var final string
	if *firstInLocale {
		*firstInLocale = false
		*firstFilter = false
		final += " ("
	} else {
		if *firstFilter {
			*firstFilter = false
			final += whereString
		} else {
			final += " AND ("
		}
	}
	for i := range values {
		if values[i] != nil {
			final += fmt.Sprintf(" %s%s = '%s'", tablePrefix, field, *values[i])
			if i < len(values)-1 {
				final += orString
			}
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

			return fmt.Sprintf(" WHERE %s%s =  '%s'", tablePrefix, field, *value)
		}

		return fmt.Sprintf(" AND %s%s =  '%s'", tablePrefix, field, *value)
	}

	return ""
}

// GenerateSingleSpecifierStringSQL generates a where function with a string.
//
// nolint:unparam
func generateSingleSpecifierStringSQLMv(value *string, field string, firstFilter *bool, firstLocale *bool, tablePrefix string) string {
	if value != nil {
		if *firstLocale {
			*firstFilter = false
			*firstLocale = false
			return fmt.Sprintf(" %s%s = '%s'", tablePrefix, field, *value)
		}
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

// generateKappaSpecifierSQL generates a where function with a string.
func generateKappaSpecifierSQLMv(value *string, field string, firstFilter *bool, firstInLocale *bool, tablePrefix string) string {
	if value != nil {
		if *firstInLocale {
			*firstFilter = false
			*firstInLocale = false
			return fmt.Sprintf(" %s%s = '%s'", tablePrefix, field, *value)
		}
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, field, *value)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, field, *value)
	}

	return ""
}

// generateCCTPSpecifierSQLMv generates a where function with event type to filter only cctp events.
func generateCCTPSpecifierSQLMv(onlyCctp *bool, to bool, field string, firstFilter *bool, firstInLocale *bool, tablePrefix string) string {
	if onlyCctp != nil && *onlyCctp {
		// From explorer/types/bridge/eventtypes.go
		eventType := 10
		if to {
			eventType = 11
		}

		if *firstInLocale {
			*firstFilter = false
			*firstInLocale = false
			return fmt.Sprintf(" %s%s = %d", tablePrefix, field, eventType)
		}
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s =  %d", tablePrefix, field, eventType)
		}

		return fmt.Sprintf(" AND %s%s = %d", tablePrefix, field, eventType)
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
func GetPartialInfoFromBridgeEventHybrid(bridgeEvent sql.HybridBridgeEvent, includePending *bool) (*model.BridgeTransaction, error) {
	if includePending != nil && *includePending && bridgeEvent.TTxHash != "" {
		// nolint:nilnil
		return nil, nil
	}
	var bridgeTx model.BridgeTransaction
	fromChainID := int(bridgeEvent.FChainID)
	fromDestinationChainID := int(bridgeEvent.FDestinationChainID.Uint64())
	fromBlockNumber := int(bridgeEvent.FBlockNumber)
	fromValue := bridgeEvent.FAmount.String()
	fromEventTypeFormatted := bridge.GetEventType(bridgeEvent.FEventType)
	fromEventType := int(bridgeEvent.FEventType)

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

	fAddress := bridgeEvent.FRecipient.String
	if bridgeEvent.FEventType == bridge.CircleRequestSentEvent.Int() {
		fAddress = bridgeEvent.FSender
	}
	fromInfos := &model.PartialInfo{
		ChainID:            &fromChainID,
		DestinationChainID: &fromDestinationChainID,
		Address:            &fAddress,
		TxnHash:            &bridgeEvent.FTxHash,
		Value:              &fromValue,
		FormattedValue:     fromFormattedValue,
		USDValue:           bridgeEvent.FAmountUSD,
		TokenAddress:       &bridgeEvent.FToken,
		TokenSymbol:        &bridgeEvent.FTokenSymbol.String,
		BlockNumber:        &fromBlockNumber,
		Time:               &fromTimestamp,
		FormattedTime:      &fromTimeStampFormatted,
		FormattedEventType: &fromEventTypeFormatted,
		EventType:          &fromEventType,
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
		toEventTypeFormatted := bridge.GetEventType(bridgeEvent.TEventType)
		toEventType := int(bridgeEvent.TEventType)

		tAddress := bridgeEvent.TRecipient.String
		if bridgeEvent.FEventType == bridge.CircleRequestFulfilledEvent.Int() {
			tAddress = bridgeEvent.TSender
		}

		toInfos = &model.PartialInfo{
			ChainID:            &toChainID,
			Address:            &tAddress,
			TxnHash:            &bridgeEvent.TTxHash,
			Value:              &toValue,
			FormattedValue:     toFormattedValue,
			USDValue:           bridgeEvent.TAmountUSD,
			TokenAddress:       &bridgeEvent.TToken,
			TokenSymbol:        &bridgeEvent.TTokenSymbol.String,
			BlockNumber:        &toBlockNumber,
			Time:               &toTimestamp,
			FormattedTime:      &toTimeStampFormatted,
			FormattedEventType: &toEventTypeFormatted,
			EventType:          &toEventType,
		}
	} else {
		toInfos = nil
		pending = true
	}

	var swapSuccess bool
	if bridgeEvent.TSwapSuccess.Uint64() == 1 {
		swapSuccess = true
	}
	if includePending != nil && !*includePending && pending {
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
func generateAllBridgeEventsQueryFromDestination(chainIDTo []*int, chainIDFrom []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, tokenAddressFrom []*string, tokenAddressTo []*string, kappa *string, txHash *string, onlyCctp *bool, page int, in bool) string {
	firstFilter := true
	chainIDToFilter := generateSingleSpecifierI32ArrSQL(chainIDTo, sql.ChainIDFieldName, &firstFilter, "")
	minTimeFilter := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "", true)
	maxTimeFilter := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "", false)
	addressToFilter := generateAddressSpecifierSQL(addressTo, &firstFilter, "")
	kappaFilter := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "")
	txHashFilter := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionFilter := generateDirectionSpecifierSQL(in, &firstFilter, "")
	cctpFilter := generateCCTPSpecifierSQL(onlyCctp, true, sql.EventTypeFieldName, &firstFilter, "")

	toFilters := chainIDToFilter + minTimeFilter + maxTimeFilter + addressToFilter + kappaFilter + txHashFilter + directionFilter + cctpFilter

	firstFilter = false
	chainIDFromFilter := generateSingleSpecifierI32ArrSQL(chainIDFrom, sql.ChainIDFieldName, &firstFilter, "")
	addressFromFilter := generateAddressSpecifierSQL(addressFrom, &firstFilter, "")

	fromFilters := chainIDFromFilter + addressFromFilter

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

func generateAllBridgeEventsQueryFromDestinationMv(chainIDTo []*int, addressTo *string, minAmount *int, minAmountUsd *int, startTime *int, endTime *int, tokenAddressTo []*string, kappa *string, txHash *string, pending *bool, page int) string {
	firstFilter := true
	chainIDToFilter := generateSingleSpecifierI32ArrSQL(chainIDTo, sql.ChainIDFieldName, &firstFilter, "t")
	minTimeFilter := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "t", true)
	maxTimeFilter := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "t", false)
	addressToFilter := generateRecipientSpecifierSQL(addressTo, &firstFilter, "t")
	kappaFilter := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "t")
	txHashFilter := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "t")
	minAmountFilter := generateEqualitySpecifierSQL(minAmount, "tamount", &firstFilter, "", true)
	minAmountFilterUsd := generateEqualitySpecifierSQL(minAmountUsd, "tamount_usd", &firstFilter, "", true)
	tokenAddressToFilter := generateSingleSpecifierStringArrSQL(tokenAddressTo, "ttoken", &firstFilter, "")

	// firstFilter = false
	// chainIDFromFilter := generateSingleSpecifierI32ArrSQL(chainIDFrom, sql.ChainIDFieldName, &firstFilter, "")
	// addressFromFilter := generateAddressSpecifierSQL(addressFrom, &firstFilter, "")
	// maxAmountFilter := generateEqualitySpecifierSQL(maxAmount, "famount", &firstFilter, "", false)
	// maxAmountFilterUsd := generateEqualitySpecifierSQL(maxAmountUsd, "famount_usd", &firstFilter, "", false)
	// tokenAddressFromFilter := generateSingleSpecifierStringArrSQL(tokenAddressFrom, "ftoken", &firstFilter, "")

	// fromFilters := chainIDFromFilter + addressFromFilter

	// firstFilter = true
	// minAmountFilter := generateEqualitySpecifierSQL(minAmount, "tamount", &firstFilter, "", true)
	// minAmountFilterUsd := generateEqualitySpecifierSQL(minAmountUsd, "tamount_usd", &firstFilter, "", true)
	pendingFilter := ""
	if pending != nil {
		prefix := " AND "
		if firstFilter {
			prefix = " WHERE "
		}
		if *pending {
			pendingFilter = prefix + "fdestination_kappa = ''"
		} else {
			pendingFilter = prefix + "fdestination_tkappa != ''"
		}
	}

	toFilters := chainIDToFilter + minTimeFilter + maxTimeFilter + addressToFilter + kappaFilter + txHashFilter + minAmountFilter + minAmountFilterUsd + tokenAddressToFilter + pendingFilter

	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize

	return fmt.Sprintf("SELECT * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash LIMIT %d OFFSET %d ", toFilters, pageValue, pageOffset)
}

// generateAllBridgeEventsQueryFromOrigin gets all the filters for query from origin.
//
// nolint:dupl
func generateAllBridgeEventsQueryFromOrigin(chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, tokenAddressFrom []*string, tokenAddressTo []*string, txHash *string, pending *bool, onlyCctp *bool, page int, in bool) string {
	firstFilter := true
	chainIDFromFilter := generateSingleSpecifierI32ArrSQL(chainIDFrom, sql.ChainIDFieldName, &firstFilter, "")
	minTimeFilter := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "", true)
	maxTimeFilter := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "", false)
	addressFromFilter := generateAddressSpecifierSQL(addressFrom, &firstFilter, "")
	txHashFilter := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionFilter := generateDirectionSpecifierSQL(in, &firstFilter, "")
	cctpFilter := generateCCTPSpecifierSQL(onlyCctp, false, sql.EventTypeFieldName, &firstFilter, "")
	fromFilters := chainIDFromFilter + minTimeFilter + maxTimeFilter + addressFromFilter + txHashFilter + directionFilter + cctpFilter

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
	if pending != nil && !*pending {
		operation = " != ''"
	}
	pendingFilter := fmt.Sprintf(" WHERE t%s %s", sql.KappaFieldName, operation)
	postJoinFilters := minAmountFilter + minAmountFilterUsd + maxAmountFilter + maxAmountFilterUsd + tokenAddressToFilter + tokenAddressFromFilter

	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	if pending != nil && !*pending && postJoinFilters == "" {
		return fmt.Sprintf("%s SELECT %s FROM %s %s %s %s", generateDeDepQueryCTE(fromFilters, &pageValue, &pageOffset, false), originToDestCol, "baseQuery", originToDestJoinsPt1, toFilters, originToDestJoinsPt2)
	}
	return fmt.Sprintf("%s SELECT * FROM (SELECT %s FROM %s %s %s %s) %s LIMIT %d OFFSET %d", generateDeDepQueryCTE(fromFilters, nil, nil, false), originToDestCol, "baseQuery", originToDestJoinsPt1, toFilters, originToDestJoinsPt2, pendingFilter+postJoinFilters, pageValue, pageOffset)
}

// generateAllBridgeEventsQueryFromOriginMv gets all the filters for query from origin.
//
// nolint:dupl
func generateAllBridgeEventsQueryFromOriginMv(chainIDFrom []*int, addressFrom *string, maxAmount *int, maxAmountUsd *int, startTime *int, endTime *int, tokenAddressFrom []*string, txHash *string, kappa *string, pending *bool, page int) string {
	firstFilter := true
	chainIDFromFilter := generateSingleSpecifierI32ArrSQL(chainIDFrom, sql.ChainIDFieldName, &firstFilter, "f")
	minTimeFilter := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "f", true)
	maxTimeFilter := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "f", false)
	addressFromFilter := generateAddressSpecifierSQL(addressFrom, &firstFilter, "f")
	txHashFilter := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "f")
	tokenAddressFromFilter := generateSingleSpecifierStringArrSQL(tokenAddressFrom, "ftoken", &firstFilter, "")
	maxAmountFilter := generateEqualitySpecifierSQL(maxAmount, "famount", &firstFilter, "", false)
	maxAmountFilterUsd := generateEqualitySpecifierSQL(maxAmountUsd, "famount_usd", &firstFilter, "", false)
	kappaFilter := generateKappaSpecifierSQL(kappa, sql.DestinationKappaFieldName, &firstFilter, "f")
	// firstFilter = false
	// chainIDToFilter := generateSingleSpecifierI32ArrSQL(chainIDTo, sql.ChainIDFieldName, &firstFilter, "")
	// addressToFilter := generateAddressSpecifierSQL(addressTo, &firstFilter, "")
	// minAmountFilter := generateEqualitySpecifierSQL(minAmount, "tamount", &firstFilter, "", true)
	// minAmountFilterUsd := generateEqualitySpecifierSQL(minAmountUsd, "tamount_usd", &firstFilter, "", true)
	// tokenAddressToFilter := generateSingleSpecifierStringArrSQL(tokenAddressTo, "ttoken", &firstFilter, "")

	pendingFilter := ""
	if pending != nil {
		prefix := " AND "
		if firstFilter {
			prefix = " WHERE "
		}
		if *pending {
			pendingFilter = prefix + "tkappa = ''"
		} else {
			pendingFilter = prefix + "tkappa != ''"
		}
	}

	fromFilters := chainIDFromFilter + minTimeFilter + maxTimeFilter + addressFromFilter + txHashFilter + tokenAddressFromFilter + maxAmountFilter + maxAmountFilterUsd + pendingFilter + kappaFilter
	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	return fmt.Sprintf("SELECT * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash LIMIT %d OFFSET %d ", fromFilters, pageValue, pageOffset)
}
func generateAllBridgeEventsQueryMv(chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, tokenAddressFrom []*string, tokenAddressTo []*string, txHash *string, kappa *string, pending *bool, onlyCctp *bool, page int) string {
	firstFilter := true
	firstInLocale := true
	chainIDFromFilter := generateSingleSpecifierI32ArrSQLMv(chainIDFrom, sql.ChainIDFieldName, &firstFilter, &firstInLocale, "f")
	addressFromFilter := generateAddressSpecifierSQLMv(addressFrom, &firstFilter, &firstInLocale, "f")
	txHashFromFilter := generateSingleSpecifierStringSQLMv(txHash, sql.TxHashFieldName, &firstFilter, &firstInLocale, "f")
	tokenAddressFromFilter := generateSingleSpecifierStringArrSQLMv(tokenAddressFrom, sql.TokenFieldName, &firstFilter, &firstInLocale, "f")
	maxAmountFilter := generateEqualitySpecifierSQLMv(maxAmount, sql.AmountFieldName, &firstFilter, &firstInLocale, "f", false)
	maxAmountFilterUsd := generateEqualitySpecifierSQLMv(maxAmountUsd, sql.AmountUSDFieldName, &firstFilter, &firstInLocale, "f", false)
	kappaFromFilter := generateKappaSpecifierSQLMv(kappa, sql.DestinationKappaFieldName, &firstFilter, &firstInLocale, "f")
	onlyCCTPFromFilter := generateCCTPSpecifierSQLMv(onlyCctp, false, sql.EventTypeFieldName, &firstFilter, &firstInLocale, "f")

	// firstFilter = false
	firstInLocale = true
	chainIDToFilter := generateSingleSpecifierI32ArrSQLMv(chainIDTo, sql.ChainIDFieldName, &firstFilter, &firstInLocale, "t")
	addressToFilter := generateRecipientSpecifierSQLMv(addressTo, &firstFilter, &firstInLocale, "t")
	txHashToFilter := generateSingleSpecifierStringSQLMv(txHash, sql.TxHashFieldName, &firstFilter, &firstInLocale, "t")
	tokenAddressToFilter := generateSingleSpecifierStringArrSQLMv(tokenAddressTo, sql.TokenFieldName, &firstFilter, &firstInLocale, "t")
	minAmountFilter := generateEqualitySpecifierSQLMv(minAmount, sql.AmountFieldName, &firstFilter, &firstInLocale, "t", true)
	minAmountFilterUsd := generateEqualitySpecifierSQLMv(minAmountUsd, sql.AmountUSDFieldName, &firstFilter, &firstInLocale, "t", true)
	kappaToFilter := generateKappaSpecifierSQLMv(kappa, sql.KappaFieldName, &firstFilter, &firstInLocale, "t")
	onlyCCTPToFilter := generateCCTPSpecifierSQLMv(onlyCctp, true, sql.EventTypeFieldName, &firstFilter, &firstInLocale, "t")

	toFilters := chainIDFromFilter + addressFromFilter + txHashFromFilter + tokenAddressFromFilter + maxAmountFilter + maxAmountFilterUsd + kappaFromFilter + onlyCCTPFromFilter
	fromFilters := chainIDToFilter + addressToFilter + txHashToFilter + tokenAddressToFilter + minAmountFilter + minAmountFilterUsd + kappaToFilter + onlyCCTPToFilter

	minTimeFilter := generateEqualitySpecifierSQL(startTime, sql.TimeStampFieldName, &firstFilter, "f", true)
	maxTimeFilter := generateEqualitySpecifierSQL(endTime, sql.TimeStampFieldName, &firstFilter, "f", false)

	var allFilters string
	switch {
	case fromFilters != "" && toFilters != "":
		allFilters = fmt.Sprintf("WHERE ((%s) OR (%s)) %s", fromFilters, toFilters, minTimeFilter+maxTimeFilter)
	case fromFilters != "" && toFilters == "":
		allFilters = fmt.Sprintf("WHERE (%s) %s", fromFilters, minTimeFilter+maxTimeFilter)
	case fromFilters == "" && toFilters != "":
		allFilters = fmt.Sprintf("WHERE (%s) %s ", toFilters, minTimeFilter+maxTimeFilter)
	default:
		allFilters = minTimeFilter + maxTimeFilter
	}

	var pendingFilter string
	if pending != nil {
		if *pending {
			pendingFilter = "WHERE tkappa = '' AND fdestination_chain_id != 121014925"
		} else {
			pendingFilter = " WHERE tkappa != ''"
		}
	}
	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	return fmt.Sprintf("SELECT * FROM(SELECT * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) %s LIMIT %d OFFSET %d SETTINGS memory_overcommit_ratio_denominator=4000, memory_usage_overcommit_max_wait_microseconds=500 ", allFilters, pendingFilter, pageValue, pageOffset)
}

// nolint:cyclop
func (r *queryResolver) GetBridgeTxsFromDestination(ctx context.Context, useMv *bool, chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txHash *string, kappa *string, tokenAddressFrom []*string, tokenAddressTo []*string, onlyCctp *bool, page *int, pending *bool) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	var query string
	if useMv != nil && *useMv {
		if chainIDTo == nil && addressTo == nil && minAmount == nil && minAmountUsd == nil && startTime == nil && endTime == nil && tokenAddressTo == nil && kappa == nil && txHash == nil {
			return nil, nil
		}
		query = generateAllBridgeEventsQueryFromDestinationMv(chainIDTo, addressTo, minAmount, minAmountUsd, startTime, endTime, tokenAddressTo, kappa, txHash, pending, *page)
	} else {
		query = generateAllBridgeEventsQueryFromDestination(chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, minAmountUsd, maxAmountUsd, startTime, endTime, tokenAddressFrom, tokenAddressTo, kappa, txHash, onlyCctp, *page, false)
	}
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}
	if len(allBridgeEvents) == 0 {
		return nil, nil
	}

	// Iterate through all bridge events and return all partials
	for i := range allBridgeEvents {
		bridgeTx, err := GetPartialInfoFromBridgeEventHybrid(allBridgeEvents[i], nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get partial info from bridge event: %w", err)
		}
		if bridgeTx != nil {
			results = append(results, bridgeTx)
		}
	}
	return results, nil
}

func (r *queryResolver) GetBridgeTxsFromOrigin(ctx context.Context, useMv *bool, chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txHash *string, tokenAddressTo []*string, tokenAddressFrom []*string, kappa *string, pending *bool, onlyCctp *bool, page *int, latest bool) ([]*model.BridgeTransaction, error) {
	var err error
	var chainMap = make(map[uint32]bool)
	var results []*model.BridgeTransaction
	query := generateAllBridgeEventsQueryFromOrigin(chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, maxAmountUsd, minAmountUsd, startTime, endTime, tokenAddressFrom, tokenAddressTo, txHash, pending, onlyCctp, *page, true)
	if useMv != nil && *useMv {
		query = generateAllBridgeEventsQueryFromOriginMv(chainIDFrom, addressFrom, maxAmount, maxAmountUsd, startTime, endTime, tokenAddressFrom, txHash, kappa, pending, *page)
	}
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

func (r *queryResolver) GetBridgeTxs(ctx context.Context, chainIDFrom []*int, chainIDTo []*int, addressFrom *string, addressTo *string, maxAmount *int, minAmount *int, maxAmountUsd *int, minAmountUsd *int, startTime *int, endTime *int, txHash *string, tokenAddressTo []*string, tokenAddressFrom []*string, kappa *string, pending *bool, onlyCctp *bool, page *int) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	query := generateAllBridgeEventsQueryMv(chainIDFrom, chainIDTo, addressFrom, addressTo, maxAmount, minAmount, maxAmountUsd, minAmountUsd, startTime, endTime, tokenAddressFrom, tokenAddressTo, txHash, kappa, pending, onlyCctp, *page)
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}
	if len(allBridgeEvents) == 0 {
		return nil, nil
	}

	// Iterate through all bridge events and return all partials
	for i := range allBridgeEvents {
		bridgeTx, err := GetPartialInfoFromBridgeEventHybrid(allBridgeEvents[i], pending)
		if err != nil {
			return nil, fmt.Errorf("failed to get partial info from bridge event: %w", err)
		}
		if bridgeTx != nil {
			results = append(results, bridgeTx)
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

// GenerateAmountStatisticBridgeSQL generate sql for the bridge platform.
func GenerateAmountStatisticBridgeSQL(typeArg model.StatisticType, address *string, chainID *int, tokenAddress *string) (*string, error) {
	var operation string
	var finalSQL string
	firstFilter2 := true
	addressFilter := generateSingleSpecifierStringSQL(address, sql.SenderFieldName, &firstFilter2, "f")
	chainIDFilter := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter2, "f")
	tokenAddressFilter := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter2, "f")
	compositeFilters := addressFilter + chainIDFilter + tokenAddressFilter
	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		operation = fmt.Sprintf("AVG(f%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s from (select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)  ", operation, compositeFilters)
	case model.StatisticTypeMedianVolumeUsd:
		operation = fmt.Sprintf("median(f%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s FROM (select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)  ", operation, compositeFilters)
	case model.StatisticTypeTotalVolumeUsd:
		operation = fmt.Sprintf("sumKahan(f%s)", sql.AmountUSDFieldName)
		finalSQL = fmt.Sprintf("SELECT %s from ( select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", operation, compositeFilters)
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("uniq(f%s, f%s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
		finalSQL = fmt.Sprintf("SELECT %s from ( select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", operation, compositeFilters)
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(f%s, f%s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
		finalSQL = fmt.Sprintf("SELECT %s from ( select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", operation, compositeFilters)
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(%s)", "tfee_amount_usd")
		finalSQL = fmt.Sprintf("SELECT %s from ( select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", operation, compositeFilters)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(%s)", "tfee_amount_usd")
		finalSQL = fmt.Sprintf("SELECT %s from ( select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", operation, compositeFilters)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", "tfee_amount_usd")
		finalSQL = fmt.Sprintf("SELECT %s from ( select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash)", operation, compositeFilters)

	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	return &finalSQL, nil
}

// GenerateAmountStatisticSwapSQL generates sql to get statistics on the swap platform.
//
// nolint:cyclop
func GenerateAmountStatisticSwapSQL(typeArg model.StatisticType, compositeFilters string, tokenAddress *string) (*string, error) {
	var operation string
	var finalSQL string

	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		operation = fmt.Sprintf("AVG(%s)", swapVolumeSelect)
	case model.StatisticTypeMedianVolumeUsd:
		operation = fmt.Sprintf("median(%s)", swapVolumeSelect)
	case model.StatisticTypeTotalVolumeUsd:
		operation = fmt.Sprintf("sumKahan(%s)", swapVolumeSelect)
	case model.StatisticTypeCountTransactions:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.TxHashFieldName)
	case model.StatisticTypeCountAddresses:
		operation = fmt.Sprintf("uniq(%s, %s) AS res", sql.ChainIDFieldName, sql.SenderFieldName)
	case model.StatisticTypeMeanFeeUsd:
		operation = fmt.Sprintf("AVG(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
	case model.StatisticTypeMedianFeeUsd:
		operation = fmt.Sprintf("median(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
	case model.StatisticTypeTotalFeeUsd:
		operation = fmt.Sprintf("sumKahan(arraySum(mapValues(%s)))", sql.FeeUSDFieldName)
	default:
		return nil, fmt.Errorf("invalid statistic type: %s", typeArg)
	}
	if tokenAddress == nil {
		finalSQL = fmt.Sprintf("SELECT %s FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash)", operation, compositeFilters)
	} else {
		firstFilter := true
		tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
		finalSQL = fmt.Sprintf("SELECT %s FROM (%s %s %s %s", operation, baseSwapWithTokenPt1, compositeFilters, baseSwapWithTokenPt2, tokenAddressSpecifier)
	}
	return &finalSQL, nil
}

// GenerateAmountStatisticMessageBusSQL generates sql for getting stats on the message bus platform.
func GenerateAmountStatisticMessageBusSQL(typeArg model.StatisticType, compositeFilters string) (*string, error) {
	var operation string
	var finalSQL string
	switch typeArg {
	case model.StatisticTypeMeanVolumeUsd:
		return nil, fmt.Errorf("cannot calculate volume data for messagebus events")
	case model.StatisticTypeMedianVolumeUsd:
		return nil, fmt.Errorf("cannot calculate volume data for messagebus events")
	case model.StatisticTypeTotalVolumeUsd:
		return nil, fmt.Errorf("cannot calculate volume data for messagebus events")
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

// GenerateRankedChainsByVolumeSQL generates sql for getting all chains ranked in order of volume.
func GenerateRankedChainsByVolumeSQL(compositeFilters string, firstFilter *bool) string {
	directionSpecifier := generateDirectionSpecifierSQL(true, firstFilter, "")
	return fmt.Sprintf("%s %s FULL OUTER JOIN (SELECT chain_id, sumKahan(multiIf(event_type = 0, amount_usd[sold_id], event_type = 1, arraySum(mapValues(amount_usd)), event_type = 9, arraySum(mapValues(amount_usd)), event_type = 10, amount_usd[sold_id], 0)) as usdTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by chain_id) s ON b.pre_fchain_id = s.chain_id ORDER BY total DESC SETTINGS join_use_nulls = 1", generateDeDepQueryCTE(compositeFilters+directionSpecifier, nil, nil, true), rankedChainsBridgeVolume, compositeFilters)
}

// GenerateDailyStatisticByChainAllSQL generates sql for getting daily stats across all chains.
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

// GenerateDailyStatisticByChainBridgeSQL generates sql for getting data for daily stats across the bridge platform.
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

// GenerateDailyStatisticByChainSwapSQL generates sql for getting daily stats across the swap platform.
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

// GenerateDailyStatisticByChainMessageBusSQL generates sql for getting daily stats across the message bus platform.
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
		query = fmt.Sprintf("%s uniq(%s, %s) AS sumTotal FROM (%s) %s group by date, chain_id) group by date order by date)", dailyStatisticGenericSinglePlatform, sql.ChainIDFieldName, sql.TxHashFieldName, baseMessageBus, compositeFilters)
	default:
		return nil, fmt.Errorf("unsupported statistic type")
	}
	return &query, nil
}

// SortBridgeTxType sorts bridge transactions by time.
type SortBridgeTxType []*model.BridgeTransaction

func (s SortBridgeTxType) Len() int           { return len(s) }
func (s SortBridgeTxType) Less(i, j int) bool { return *s[i].FromInfo.Time > *s[j].FromInfo.Time }
func (s SortBridgeTxType) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// SortMessageBusTxType sorts message bus transactions by time.
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
		return *item
	}
	return ""
}

// Gets the value result from cache.
func (r *queryResolver) getValueResultFromCache(key string) (*model.ValueResult, error) {
	cacheResult := r.Cache.GetCache(key)
	if cacheResult != nil {
		rawCache, ok := cacheResult.(*interface{})
		if !ok || rawCache == nil {
			return nil, fmt.Errorf("type assertion error when converting to *interface{}, rawCache %v", rawCache)
		}
		res, ok := (*rawCache).(*model.ValueResult)
		if !ok || res == nil {
			return nil, fmt.Errorf("type assertion error when converting to *model.ValueResult, res %v", res)
		}
		return res, nil
	}
	return nil, fmt.Errorf("could not get cached data")
}

// Gets the value result from cache.
func (r *queryResolver) getDateResultByChainFromCache(key string) ([]*model.DateResultByChain, error) {
	cacheResult := r.Cache.GetCache(key)
	if cacheResult != nil {
		rawCache, ok := cacheResult.(*interface{})
		if !ok || rawCache == nil {
			return nil, fmt.Errorf("type assertion error when converting to *interface{}, rawCache %v", rawCache)
		}
		res, ok := (*rawCache).([]*model.DateResultByChain)
		if !ok || res == nil {
			return nil, fmt.Errorf("type assertion error when converting to []*model.DateResultByChain, res %v", res)
		}
		return res, nil
	}
	return nil, fmt.Errorf("could not get cached data")
}

// GetDurationFilter creates a filter for the various time ranges for analysis.
func GetDurationFilter(duration *model.Duration, firstFilter *bool, prefix string) string {
	var timestampSpecifier string
	switch *duration {
	case model.DurationPastDay:
		hours := 24
		targetTime := GetTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, firstFilter, prefix)
	case model.DurationPastMonth:
		hours := 720
		targetTime := GetTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, firstFilter, prefix)
	case model.DurationPast3Months:
		hours := 2190
		targetTime := GetTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, firstFilter, prefix)
	case model.DurationPast6Months:
		hours := 4380
		targetTime := GetTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, firstFilter, prefix)
	case model.DurationPastYear:
		hours := 8760
		targetTime := GetTargetTime(&hours)
		timestampSpecifier = generateTimestampSpecifierSQL(&targetTime, sql.TimeStampFieldName, firstFilter, prefix)
	case model.DurationAllTime:
		timestampSpecifier = ""
	}
	return timestampSpecifier
}

// nolint:cyclop
func (r *queryResolver) getAmountStatisticsAll(ctx context.Context, typeArg model.StatisticType, chainID *int, address *string, tokenAddress *string, compositeFilters string) (*string, error) {
	if typeArg == model.StatisticTypeMedianVolumeUsd || typeArg == model.StatisticTypeMeanVolumeUsd || typeArg == model.StatisticTypeMedianFeeUsd || typeArg == model.StatisticTypeMeanFeeUsd {
		return nil, fmt.Errorf("cannot calculate averages or medians across all platforms")
	}
	var bridgeFinalSQL *string
	var swapFinalSQL *string
	var messageBusFinalSQL *string
	var err error
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
	return &value, nil
}

// nolint:cyclop
func (r *queryResolver) getDateResultByChainMv(ctx context.Context, chainID *int, typeArg *model.DailyStatisticType, platform *model.Platform, duration *model.Duration) ([]*model.DateResultByChain, error) {
	var err error
	firstFilter := true
	timestampSpecifierMv := GetDurationFilter(duration, &firstFilter, "f")
	chainIDSpecifierMv := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "f")
	compositeFiltersMv := fmt.Sprintf(
		`%s%s`,
		timestampSpecifierMv, chainIDSpecifierMv,
	)
	firstFilter = true
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
		// Change chainID filter to destination chainID as that's where fees are collected.
		if *typeArg == model.DailyStatisticTypeFee {
			chainIDSpecifierMv = generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t")
			compositeFiltersMv = fmt.Sprintf(
				`%s%s`,
				timestampSpecifierMv, chainIDSpecifierMv,
			)
		}
		query, err = GenerateDailyStatisticByChainBridgeSQLMv(typeArg, compositeFiltersMv)
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
		query, err = GenerateDailyStatisticByChainAllSQLMv(typeArg, compositeFilters, compositeFiltersMv)
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
		return nil, fmt.Errorf("error caching response, %w", err)
	}
	return res, nil
}

// GenerateDailyStatisticByChainBridgeSQLMv generates sql for getting data for daily stats across the bridge platform.
func GenerateDailyStatisticByChainBridgeSQLMv(typeArg *model.DailyStatisticType, compositeFilters string) (*string, error) {
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		query = fmt.Sprintf("%s  sumKahan(famount_usd) AS sumTotal from (select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) group by date, chain_id) group by date order by date)  ", dailyStatisticGenericSinglePlatformMv, compositeFilters)
	case model.DailyStatisticTypeFee:
		query = fmt.Sprintf("%s  sumKahan(tfee_amount_usd) AS sumTotal from (select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) group by date, chain_id) group by date order by date)  ", dailyStatisticGenericSinglePlatformMvFee, compositeFilters)
	case model.DailyStatisticTypeAddresses:
		query = fmt.Sprintf("%s  uniq(fchain_id, fsender) AS sumTotal from (select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) group by date, chain_id) group by date order by date)  ", dailyStatisticGenericSinglePlatformMv, compositeFilters)
	case model.DailyStatisticTypeTransactions:
		query = fmt.Sprintf("%s  uniq(fchain_id, ftx_hash) AS sumTotal from (select * FROM mv_bridge_events %s ORDER BY ftimestamp DESC, fblock_number DESC, fevent_index DESC, insert_time DESC LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) group by date, chain_id) group by date order by date) ", dailyStatisticGenericSinglePlatformMv, compositeFilters)
	default:
		return nil, fmt.Errorf("unsupported statistic type")
	}
	return &query, nil
}

// GenerateDailyStatisticByChainAllSQLMv generates sql for getting daily stats across all chains.
func GenerateDailyStatisticByChainAllSQLMv(typeArg *model.DailyStatisticType, compositeFilters string, compositeFiltersMv string) (*string, error) {
	var query string
	switch *typeArg {
	case model.DailyStatisticTypeVolume:
		query = fmt.Sprintf("%s %s %s FULL OUTER JOIN (SELECT %s, chain_id, sumKahan(multiIf(event_type = 0, amount_usd[sold_id], event_type = 1,    arraySum(mapValues(amount_usd)), event_type = 9,    arraySum(mapValues(amount_usd)), event_type = 10, amount_usd[sold_id],    0) )     as usdTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.chain_id = s.chain_id) group by date order by date) SETTINGS join_use_nulls=1", dailyVolumeBridgeMvPt1, compositeFiltersMv, dailyVolumeBridgeMvPt2, toDateSelect, compositeFilters)
	case model.DailyStatisticTypeFee: // destination chain fee used
		query = fmt.Sprintf("%s FROM ( SELECT %s, tchain_id AS chain_id, sumKahan(tfee_amount_usd) as sumTotal FROM (SELECT * FROM mv_bridge_events %s LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) GROUP BY date, chain_id) b FULL OUTER JOIN ( SELECT %s, chain_id, sumKahan(arraySum(mapValues(fee_usd))) AS sumTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.chain_id = s.chain_id  FULL OUTER JOIN ( SELECT %s, chain_id, sumKahan(fee_usd) AS sumTotal FROM (SELECT * FROM message_bus_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) m ON b.date = m.date AND b.chain_id = m.chain_id) group by date order by date ) SETTINGS join_use_nulls = 1", dailyStatisticGenericSelect, toDateSelectMv, compositeFiltersMv, toDateSelect, compositeFilters, toDateSelect, compositeFilters)
	case model.DailyStatisticTypeAddresses:
		query = fmt.Sprintf("%s FROM ( SELECT %s, fchain_id AS chain_id, uniq(fchain_id, fsender) as sumTotal FROM (SELECT * FROM mv_bridge_events %s LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) GROUP BY date, chain_id) b FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, sender) AS sumTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.chain_id = s.chain_id  FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, source_address) AS sumTotal FROM (SELECT * FROM message_bus_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) m ON b.date = m.date AND b.chain_id = m.chain_id) group by date order by date ) SETTINGS join_use_nulls = 1", dailyStatisticGenericSelect, toDateSelectMv, compositeFiltersMv, toDateSelect, compositeFilters, toDateSelect, compositeFilters)
	case model.DailyStatisticTypeTransactions:
		query = fmt.Sprintf("%s FROM ( SELECT %s, fchain_id AS chain_id, uniq(fchain_id, ftx_hash) as sumTotal FROM (SELECT * FROM mv_bridge_events %s LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash) GROUP BY date, chain_id) b FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, tx_hash) AS sumTotal FROM (SELECT * FROM swap_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) s ON b.date = s.date AND b.chain_id = s.chain_id  FULL OUTER JOIN ( SELECT %s, chain_id, uniq(chain_id, tx_hash) AS sumTotal FROM (SELECT * FROM message_bus_events %s LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash) group by date, chain_id ) m ON b.date = m.date AND b.chain_id = m.chain_id) group by date order by date ) SETTINGS join_use_nulls = 1", dailyStatisticGenericSelect, toDateSelectMv, compositeFiltersMv, toDateSelect, compositeFilters, toDateSelect, compositeFilters)
	default:
		return nil, fmt.Errorf("unsupported statistic type")
	}
	return &query, nil
}

// increase this to enable querying the db
const fallbackTime = time.Second * 0

// GetOriginBridgeTxBW gets an origin bridge tx.
func (r *queryResolver) GetOriginBridgeTxBW(ctx context.Context, chainID int, txnHash string, eventType model.BridgeType) (*model.BridgeWatcherTx, error) {
	txType := model.BridgeTxTypeOrigin
	query := fmt.Sprintf("SELECT * FROM mv_bridge_events WHERE fchain_id = %d AND ftx_hash = '%s' ORDER BY insert_time desc LIMIT 1 BY fchain_id, fcontract_address, fevent_type, fblock_number, fevent_index, ftx_hash", chainID, txnHash)

	bwQueryCtx, cancel := context.WithTimeout(ctx, fallbackTime)
	defer cancel()

	bridgeEventMV, err := r.DB.GetMVBridgeEvent(bwQueryCtx, query)

	if err != nil || bridgeEventMV == nil || bridgeEventMV.FChainID == 0 {
		switch eventType {
		case model.BridgeTypeBridge:
			return r.bwOriginFallback(ctx, uint32(chainID), txnHash)
		case model.BridgeTypeCctp:
			return r.bwOriginFallbackCCTP(ctx, uint32(chainID), txnHash)
		}
	}
	return bwBridgeMVToBWTxOrigin(bridgeEventMV, txType)
}

// GetDestinationBridgeTxBW returns the destination bridge transaction for the bridgewatcher.
func (r *queryResolver) GetDestinationBridgeTxBW(ctx context.Context, chainID int, address string, kappa string, timestamp int, historical bool, bridgeType model.BridgeType) (*model.BridgeWatcherTx, error) {
	var err error
	txType := model.BridgeTxTypeDestination
	bwQueryCtx, cancel := context.WithTimeout(ctx, fallbackTime)
	defer cancel()

	query := fmt.Sprintf("SELECT * FROM mv_bridge_events WHERE tchain_id = %d AND tkappa = '%s' ORDER BY insert_time desc LIMIT 1 BY tchain_id, tcontract_address, tevent_type, tblock_number, tevent_index, ttx_hash", chainID, kappa)
	bridgeEventMV, err := r.DB.GetMVBridgeEvent(bwQueryCtx, query)

	var bridgeTx model.PartialInfo
	isPending := true

	if err != nil || bridgeEventMV == nil || bridgeEventMV.TChainID == 0 {
		var txFromChain *model.BridgeWatcherTx
		txFromChain, err = r.bwDestinationFallback(ctx, uint32(chainID), address, kappa, timestamp, historical, bridgeType)
		if err != nil {
			if err.Error() == kappaDoesNotExist {
				pendingKappa := model.KappaStatusPending
				return &model.BridgeWatcherTx{
					BridgeTx:    &bridgeTx,
					Pending:     &isPending,
					Type:        &txType,
					Kappa:       &kappa,
					KappaStatus: &pendingKappa,
				}, nil
			}
			return nil, fmt.Errorf("failed to get destination bridge event from chain: %w", err)
		}
		return txFromChain, nil
	}
	return bwBridgeMVToBWTxDestination(bridgeEventMV, txType)
}

func bwBridgeToBWTx(bridgeEvent *sql.BridgeEvent, txType model.BridgeTxType) (*model.BridgeWatcherTx, error) {
	var bridgeTx model.PartialInfo
	chainID := int(bridgeEvent.ChainID)
	isPending := false
	blockNumber := int(bridgeEvent.BlockNumber)
	value := bridgeEvent.Amount.String()
	var timestamp int
	var formattedValue *float64
	var timeStampFormatted string
	if bridgeEvent.TokenDecimal != nil {
		formattedValue = getAdjustedValue(bridgeEvent.Amount, *bridgeEvent.TokenDecimal)
	} else {
		return nil, fmt.Errorf("token decimal is not valid")
	}
	if bridgeEvent.TimeStamp != nil {
		timestamp = int(*bridgeEvent.TimeStamp)
		timeStampFormatted = time.Unix(int64(*bridgeEvent.TimeStamp), 0).String()
	} else {
		return nil, fmt.Errorf("timestamp is not valid")
	}

	kappa := bridgeEvent.DestinationKappa
	destinationChainID := int(bridgeEvent.ChainID)
	if txType == model.BridgeTxTypeOrigin {
		destinationChainID = int(bridgeEvent.DestinationChainID.Uint64())
	}
	if txType == model.BridgeTxTypeDestination {
		kappa = bridgeEvent.Kappa.String
	}
	bridgeTx = model.PartialInfo{
		ChainID:            &chainID,
		DestinationChainID: &destinationChainID,
		Address:            &bridgeEvent.Recipient.String,
		TxnHash:            &bridgeEvent.TxHash,
		Value:              &value,
		FormattedValue:     formattedValue,
		USDValue:           bridgeEvent.AmountUSD,
		TokenAddress:       &bridgeEvent.Token,
		TokenSymbol:        &bridgeEvent.TokenSymbol.String,
		BlockNumber:        &blockNumber,
		Time:               &timestamp,
		FormattedTime:      &timeStampFormatted,
	}
	result := &model.BridgeWatcherTx{
		BridgeTx: &bridgeTx,
		Pending:  &isPending,
		Type:     &txType,
		Kappa:    &kappa,
	}
	return result, nil
}

func bwBridgeMVToBWTxOrigin(bridgeEvent *sql.HybridBridgeEvent, txType model.BridgeTxType) (*model.BridgeWatcherTx, error) {
	var bridgeTx model.PartialInfo
	chainID := int(bridgeEvent.FChainID)
	isPending := false
	blockNumber := int(bridgeEvent.FBlockNumber)
	value := bridgeEvent.FAmount.String()
	var timestamp int
	var formattedValue *float64
	var timeStampFormatted string
	if bridgeEvent.FTokenDecimal != nil {
		formattedValue = getAdjustedValue(bridgeEvent.FAmount, *bridgeEvent.FTokenDecimal)
	} else {
		return nil, fmt.Errorf("token decimal is not valid")
	}
	if bridgeEvent.FTimeStamp != nil {
		timestamp = int(*bridgeEvent.FTimeStamp)
		timeStampFormatted = time.Unix(int64(*bridgeEvent.FTimeStamp), 0).String()
	} else {
		return nil, fmt.Errorf("timestamp is not valid")
	}

	kappa := bridgeEvent.FDestinationKappa
	destinationChainID := int(bridgeEvent.FDestinationChainID.Uint64())
	kappaStatus := model.KappaStatusUnknown
	bridgeTx = model.PartialInfo{
		ChainID:            &chainID,
		DestinationChainID: &destinationChainID,
		Address:            &bridgeEvent.FRecipient.String,
		TxnHash:            &bridgeEvent.FTxHash,
		Value:              &value,
		FormattedValue:     formattedValue,
		USDValue:           bridgeEvent.FAmountUSD,
		TokenAddress:       &bridgeEvent.FToken,
		TokenSymbol:        &bridgeEvent.FTokenSymbol.String,
		BlockNumber:        &blockNumber,
		Time:               &timestamp,
		FormattedTime:      &timeStampFormatted,
	}
	result := &model.BridgeWatcherTx{
		BridgeTx:    &bridgeTx,
		Pending:     &isPending,
		Type:        &txType,
		Kappa:       &kappa,
		KappaStatus: &kappaStatus,
	}
	return result, nil
}

func bwBridgeMVToBWTxDestination(bridgeEvent *sql.HybridBridgeEvent, txType model.BridgeTxType) (*model.BridgeWatcherTx, error) {
	var bridgeTx model.PartialInfo
	chainID := int(bridgeEvent.TChainID)
	isPending := false
	blockNumber := int(bridgeEvent.TBlockNumber)
	value := bridgeEvent.TAmount.String()
	var timestamp int
	var formattedValue *float64
	var timeStampFormatted string
	if bridgeEvent.TTokenDecimal != nil {
		formattedValue = getAdjustedValue(bridgeEvent.TAmount, *bridgeEvent.TTokenDecimal)
	} else {
		return nil, fmt.Errorf("token decimal is not valid")
	}
	if bridgeEvent.TTimeStamp != nil {
		timestamp = int(*bridgeEvent.TTimeStamp)
		timeStampFormatted = time.Unix(int64(*bridgeEvent.TTimeStamp), 0).String()
	} else {
		return nil, fmt.Errorf("timestamp is not valid")
	}

	destinationChainID := int(bridgeEvent.TChainID)
	kappa := bridgeEvent.TKappa.String
	kappaStatus := model.KappaStatusExists
	bridgeTx = model.PartialInfo{
		ChainID:            &chainID,
		DestinationChainID: &destinationChainID,
		Address:            &bridgeEvent.TRecipient.String,
		TxnHash:            &bridgeEvent.TTxHash,
		Value:              &value,
		FormattedValue:     formattedValue,
		USDValue:           bridgeEvent.TAmountUSD,
		TokenAddress:       &bridgeEvent.TToken,
		TokenSymbol:        &bridgeEvent.TTokenSymbol.String,
		BlockNumber:        &blockNumber,
		Time:               &timestamp,
		FormattedTime:      &timeStampFormatted,
	}
	result := &model.BridgeWatcherTx{
		BridgeTx:    &bridgeTx,
		Pending:     &isPending,
		Type:        &txType,
		Kappa:       &kappa,
		KappaStatus: &kappaStatus,
	}
	return result, nil
}

func (r *queryResolver) checkIfChainIDExists(chainIDNeeded uint32, bridgeType model.BridgeType) bool {
	exists := false
	for chainID, chainConfig := range r.Config.Chains {
		if chainID == chainIDNeeded {
			switch bridgeType {
			case model.BridgeTypeBridge:
				if chainConfig.Contracts.Bridge != "" {
					exists = true
				}
			case model.BridgeTypeCctp:
				if chainConfig.Contracts.CCTP != "" {
					exists = true
				}
			}
		}
	}
	return exists
}

func (r *queryResolver) getContractAddressFromType(chainID uint32, contractType model.ContractType) (string, error) {
	if _, ok := r.Config.Chains[chainID]; !ok {
		return "", fmt.Errorf("chain ID not found")
	}
	switch contractType {
	case model.ContractTypeBridge:
		return r.Config.Chains[chainID].Contracts.Bridge, nil
	case model.ContractTypeCctp:
		return r.Config.Chains[chainID].Contracts.CCTP, nil
	default:
		return "", fmt.Errorf("contract type not supported")
	}
}
