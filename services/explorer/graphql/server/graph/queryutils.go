package graph

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math/big"
	"strconv"
	"strings"
	"time"
)

const sortingKeys = "event_index, block_number, event_type, tx_hash, chain_id, contract_address"
const maxBlockNumberSortingKeys = "event_index, event_type, tx_hash, chain_id, contract_address"

const deDupInQuery = "(" + sortingKeys + ", insert_time) IN (SELECT " + sortingKeys + ", max(insert_time) as insert_time FROM bridge_events GROUP BY " + sortingKeys + ")"
const deDupInQueryLatest = "(" + maxBlockNumberSortingKeys + ", block_number, insert_time) IN (SELECT " + maxBlockNumberSortingKeys + ", max(block_number) as block_number, max(insert_time) as insert_time FROM bridge_events GROUP BY " + maxBlockNumberSortingKeys + ")"

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
		key := keyGen(fmt.Sprintf("%d", originTx.FromInfo.ChainID), *originTx.Kappa)
		uniqueBridgeTransactions[key] = originTx
	}

	for _, destinationTx := range destination {
		key := keyGen(fmt.Sprintf("%d", destinationTx.FromInfo.ChainID), *destinationTx.Kappa)
		uniqueBridgeTransactions[key] = destinationTx
	}

	for _, v := range uniqueBridgeTransactions {
		results = append(results, v)
	}

	return results
}

// generateAddressSpecifierSQL generates a where function with an string.
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

// generateSingleSpecifierI32SQL generates a where function with an uint32.
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

// generateTimestampSpecifierSQL generates a where function with an uint64.
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

// generateDestinationChainIDSpecifierSQL generates a where function with a string.
func generateDestinationChainIDSpecifierSQL(field string, firstFilter *bool, tablePrefix string, destination bool) string {
	if destination {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s == 0", tablePrefix, field)
		}

		return fmt.Sprintf(" AND %s%s  == 0", tablePrefix, field)
	}
	if *firstFilter {
		*firstFilter = false

		return fmt.Sprintf(" WHERE %s%s > 0", tablePrefix, field)
	}
	return fmt.Sprintf(" AND %s%s  > 0", tablePrefix, field)
}

// generatePartialInfoQuery returns the query for making the PartialInfo query.
func generatePartialInfoQuery(chainID *int, address, tokenAddress, kappa, txHash *string, page int, latest bool, destination bool) string {
	firstFilter := true

	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t1.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "t1.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "t1.")
	kappaSpecifier := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "t1.")
	destinationChainIDSpecifier := generateDestinationChainIDSpecifierSQL(sql.DestinationChainIDFieldName, &firstFilter, "t1.", destination)
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "t1.")
	pageSpecifier := fmt.Sprintf(" ORDER BY %s DESC, %s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.EventIndexFieldName, sql.PageSize, (page-1)*sql.PageSize)

	compositeIdentifiers := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + destinationChainIDSpecifier + txHashSpecifier + pageSpecifier

	selectParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s, max(%s)`,
		sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName,
		sql.SenderFieldName, sql.TxHashFieldName, sql.InsertTimeFieldName,
	)
	groupByParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s`,
		sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName,
	)
	joinOnParameters := fmt.Sprintf(
		`t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s
		AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = insert_max_time`,
		sql.TxHashFieldName, sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName,
		sql.ChainIDFieldName, sql.EventTypeFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.TokenFieldName, sql.AmountFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.EventIndexFieldName,
		sql.DestinationKappaFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName, sql.SenderFieldName, sql.InsertTimeFieldName,
	)
	deDup := deDupInQuery
	if latest {
		deDup = deDupInQueryLatest
	}
	query := fmt.Sprintf(
		`
		SELECT t1.* FROM bridge_events t1
    	JOIN (
    	SELECT %s AS insert_max_time
    	FROM bridge_events WHERE %s GROUP BY %s) t2
    	    ON (%s) %s `,
		selectParameters, deDup, groupByParameters, joinOnParameters, compositeIdentifiers)

	return query
}

// generateBridgeEventCountQuery creates the query for bridge event count.
func generateBridgeEventCountQuery(chainID *int, address *string, tokenAddress *string, directionIn bool, timestamp *uint64, isTokenCount bool) string {
	chainField := sql.ChainIDFieldName

	if directionIn {
		chainField = sql.DestinationChainIDFieldName
	}

	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, chainField, &firstFilter, "")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.RecipientFieldName, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	timestampSpecifier := generateTimestampSpecifierSQL(timestamp, sql.TimeStampFieldName, &firstFilter, "")
	var query string
	if isTokenCount {
		query = fmt.Sprintf(`SELECT %s, %s AS TokenAddress, COUNT(DISTINCT (%s)) AS Count FROM bridge_events %s%s%s%s GROUP BY %s, %s ORDER BY Count Desc`,
			sql.ChainIDFieldName, sql.TokenFieldName, sql.TxHashFieldName, chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, timestampSpecifier, sql.TokenFieldName, sql.ChainIDFieldName)
	} else {
		query = fmt.Sprintf(`SELECT %s, COUNT(DISTINCT (%s)) AS Count FROM bridge_events %s%s%s%s GROUP BY %s ORDER BY Count Desc`,
			sql.ChainIDFieldName, sql.TxHashFieldName, chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, timestampSpecifier, sql.ChainIDFieldName)
	}
	return query
}

func (r *queryResolver) generateSubQuery(ctx context.Context, targetTime uint64, colOne string, colTwo string) (string, error) {
	subQuery := "("
	chainIDs, err := r.DB.GetAllChainIDs(ctx)
	if err != nil {
		return subQuery, fmt.Errorf("failed to get chain IDs: %w", err)
	}

	for i, chain := range chainIDs {
		sqlString := fmt.Sprintf("\nSELECT %s, %s, amount_usd FROM bridge_events WHERE %s = %d AND  %s >= %d AND %s", colOne, colTwo, sql.ChainIDFieldName, chain, sql.TimeStampFieldName, targetTime, deDupInQuery)

		if i != len(chainIDs)-1 {
			sqlString += " UNION ALL"
		}

		subQuery += sqlString
	}

	return subQuery + ")", nil
}

// GetPartialInfoFromBridgeEventSingle returns the partial info from bridge event.
func GetPartialInfoFromBridgeEventSingle(res sql.BridgeEvent) (*model.PartialInfo, error) {
	var partialInfos *model.PartialInfo
	chainIDInt := int(res.ChainID)
	blockNumberInt := int(res.BlockNumber)

	var recipient string

	switch {
	case res.Recipient.Valid:
		recipient = res.Recipient.String
	case res.RecipientBytes.Valid:
		recipient = res.RecipientBytes.String
	default:
		return nil, fmt.Errorf("recipient is not valid")
	}

	var tokenSymbol string

	if res.TokenSymbol.Valid && res.TokenSymbol.String != "" {
		tokenSymbol = res.TokenSymbol.String
	} else {
		return nil, fmt.Errorf("token symbol is not valid")
	}

	value := res.Amount.String()

	var formattedValue *float64

	if res.TokenDecimal != nil {
		formattedValue = getAdjustedValue(res.Amount, *res.TokenDecimal)
	} else {
		return nil, fmt.Errorf("token decimal is not valid")
	}

	var timeStamp int

	if res.TimeStamp != nil {
		timeStamp = int(*res.TimeStamp)
	} else {
		return nil, fmt.Errorf("time stamp is not valid")
	}

	partialInfos = &model.PartialInfo{
		ChainID:        &chainIDInt,
		Address:        &recipient,
		TxnHash:        &res.TxHash,
		Value:          &value,
		FormattedValue: formattedValue,
		USDValue:       res.AmountUSD,
		TokenAddress:   &res.Token,
		TokenSymbol:    &tokenSymbol,
		BlockNumber:    &blockNumberInt,
		Time:           &timeStamp,
	}

	return partialInfos, nil
}

// generateBridgeEventsWithKappaQuery gets bridge events with a (chainid, kappa) string.
func generateBridgeEventsWithKappaQuery(kappaChainStr string, chainID *int, address, tokenAddress, kappa, txHash *string, page int, latest bool, to bool) string {
	firstFilter := false
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t1.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "t1.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "t1.")
	kappaSpecifier := generateSingleSpecifierStringSQL(kappa, sql.KappaFieldName, &firstFilter, "t1.")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "t1.")

	pageSpecifier := fmt.Sprintf(" ORDER BY %s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.PageSize, (page-1)*sql.PageSize)
	var kappaChainSpecifier string
	if to {
		kappaChainSpecifier = fmt.Sprintf("WHERE (%s,%s) IN %s", sql.ChainIDFieldName, sql.KappaFieldName, kappaChainStr)
	} else {
		kappaChainSpecifier = fmt.Sprintf("WHERE (%s,%s) IN %s", sql.DestinationChainIDFieldName, sql.DestinationKappaFieldName, kappaChainStr)
	}
	compositeIdentifiers := kappaChainSpecifier + chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + pageSpecifier
	selectParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s, max(%s)`,
		sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName,
		sql.SenderFieldName, sql.TxHashFieldName, sql.InsertTimeFieldName,
	)
	groupByParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s`,
		sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName,
	)
	joinOnParameters := fmt.Sprintf(
		`t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s
		AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = insert_max_time`,
		sql.TxHashFieldName, sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName,
		sql.ChainIDFieldName, sql.EventTypeFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.TokenFieldName, sql.AmountFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.EventIndexFieldName,
		sql.DestinationKappaFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName, sql.SenderFieldName, sql.InsertTimeFieldName,
	)
	deDup := deDupInQuery
	if latest {
		deDup = deDupInQueryLatest
	}
	query := fmt.Sprintf(
		`
		SELECT t1.* FROM bridge_events t1
    	JOIN (
    	SELECT %s AS insert_max_time
    	FROM bridge_events WHERE %s  GROUP BY %s) t2
    	    ON (%s) %s `,
		selectParameters, deDup, groupByParameters, joinOnParameters, compositeIdentifiers)

	return query
}

// generatePartialInfoQueryByChain returns the query for making the PartialInfo query.
func generatePartialInfoQueryByChain(limitSize int) string {
	pageSpecifier := fmt.Sprintf(" ORDER BY (%s,%s) DESC LIMIT %d BY %s", sql.BlockNumberFieldName, sql.EventIndexFieldName, limitSize, sql.ChainIDFieldName)
	compositeIdentifiers := pageSpecifier
	selectParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s, max(%s)`,
		sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName,
		sql.SenderFieldName, sql.TxHashFieldName, sql.InsertTimeFieldName,
	)
	groupByParameters := fmt.Sprintf(
		`%s,%s,%s,%s,%s,%s,%s,%s,%s,%s`,
		sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName,
	)
	joinOnParameters := fmt.Sprintf(
		`t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s
		AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = t2.%s AND t1.%s = insert_max_time`,
		sql.TxHashFieldName, sql.TxHashFieldName, sql.ContractAddressFieldName, sql.ContractAddressFieldName, sql.ChainIDFieldName,
		sql.ChainIDFieldName, sql.EventTypeFieldName, sql.EventTypeFieldName, sql.BlockNumberFieldName, sql.BlockNumberFieldName,
		sql.TokenFieldName, sql.TokenFieldName, sql.AmountFieldName, sql.AmountFieldName, sql.EventIndexFieldName, sql.EventIndexFieldName,
		sql.DestinationKappaFieldName, sql.DestinationKappaFieldName, sql.SenderFieldName, sql.SenderFieldName, sql.InsertTimeFieldName,
	)
	query := fmt.Sprintf(
		`
		SELECT t1.* FROM bridge_events t1
    	JOIN (
    	SELECT %s AS insert_max_time
    	FROM bridge_events WHERE %s GROUP BY %s) t2
    	    ON (%s) %s`,
		selectParameters, deDupInQuery, groupByParameters, joinOnParameters, compositeIdentifiers)

	return query
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromOrigin(ctx context.Context, chainID *int, address *string, txnHash *string, includePending bool, page int, tokenAddress *string, latest bool) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	var fromBridgeEvents []sql.BridgeEvent
	if latest {
		fromBridgeEvents, err = r.DB.GetBridgeEvents(ctx, generatePartialInfoQueryByChain(100))
	} else {
		fromBridgeEvents, err = r.DB.GetBridgeEvents(ctx, generatePartialInfoQuery(chainID, address, tokenAddress, nil, txnHash, page, false, false))
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get origin bridge events from identifiers: %w", err)
	}
	if len(fromBridgeEvents) == 0 {
		return nil, nil
	}

	var toKappaChainArr []string
	var fromBridgeEventsKappaStatusMap = make(map[string]bool)
	var chainCheck = make(map[string]bool)

	// Go through every bridge event and generate a kappa query to find destination/to bridge events and fill a pending
	// map struct
	for _, bridgeEvent := range fromBridgeEvents {
		if bridgeEvent.DestinationChainID != nil {
			key := keyGen(bridgeEvent.DestinationChainID.String(), bridgeEvent.DestinationKappa)
			toKappaChainArr = append(toKappaChainArr, fmt.Sprintf("(%d,'%s')", bridgeEvent.DestinationChainID, bridgeEvent.DestinationKappa))
			fromBridgeEventsKappaStatusMap[key] = false
			chainCheck[fmt.Sprintf("%d", bridgeEvent.DestinationChainID)] = false
		}
	}
	var toKappaChainStr string
	if len(toKappaChainArr) > 1 {
		toKappaChainStr = "(" + strings.Join(toKappaChainArr, ",") + ")" // (1,'0x123'),(2,'0x456')
	} else {
		toKappaChainStr = strings.Join(toKappaChainArr, ",") // (1,'0x123'),(2,'0x456')
	}

	// Get all destination/to bridge events that match the kappa
	toBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generateBridgeEventsWithKappaQuery(toKappaChainStr, nil, nil, nil, nil, nil, page, true, true))
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}

	// Create a map of destination/to bridge events with the chain and kappa as the key
	toBridgeEventsMap := make(map[string]sql.BridgeEvent)

	// Define pending status from the origin/from bridge events.
	for _, toBridgeEvent := range toBridgeEvents {
		if toBridgeEvent.Kappa.Valid {
			key := keyGen(fmt.Sprintf("%d", toBridgeEvent.ChainID), toBridgeEvent.Kappa.String)
			fromBridgeEventsKappaStatusMap[key] = true
			toBridgeEventsMap[key] = toBridgeEvent
		}
	}

	// Iterate through all bridge events and return all partials
	for i := range fromBridgeEvents {
		fromBridgeEvent := fromBridgeEvents[i]
		if latest && chainCheck[fmt.Sprintf("%d", fromBridgeEvent.ChainID)] {
			continue
		}
		key := keyGen(fromBridgeEvent.DestinationChainID.String(), fromBridgeEvent.DestinationKappa)

		// Generate partial info
		fromInfo, err := GetPartialInfoFromBridgeEventSingle(fromBridgeEvent)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}

		// nolint:nestif
		if fromBridgeEventsKappaStatusMap[key] {
			toBridgeEvent := toBridgeEventsMap[key]

			var swapSuccess bool
			if toBridgeEvent.SwapSuccess.Uint64() == 1 {
				swapSuccess = true
			}

			pending := false

			toInfo, err := GetPartialInfoFromBridgeEventSingle(toBridgeEvent)
			if err != nil {
				return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
			}

			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      toInfo,
				Kappa:       &fromBridgeEvent.DestinationKappa,
				Pending:     &pending,
				SwapSuccess: &swapSuccess,
			})
		} else if includePending {
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      nil,
				Kappa:       &fromBridgeEvent.DestinationKappa,
				Pending:     &includePending,
				SwapSuccess: nil,
			})
		}
		if latest {
			chainCheck[fmt.Sprintf("%d", fromBridgeEvent.ChainID)] = true
		}
	}
	return results, nil
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromDestination(ctx context.Context, chainID *int, address *string, txnHash *string, kappa *string, page int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	toBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generatePartialInfoQuery(chainID, address, tokenAddress, kappa, txnHash, page, false, true))
	if err != nil {
		return nil, fmt.Errorf("failed to get destinationbridge events from identifiers: %w", err)
	}
	if len(toBridgeEvents) == 0 {
		return nil, nil
	}
	var fromKappaChainArr []string

	// Go through every bridge event and generate a kappa query to find destination/to bridge events and fill a pending
	// map struct
	for _, bridgeEvent := range toBridgeEvents {
		if bridgeEvent.Kappa.Valid {
			fromKappaChainArr = append(fromKappaChainArr, fmt.Sprintf("(%d,'%s')", bridgeEvent.ChainID, bridgeEvent.Kappa.String))
		}
	}

	var fromKappaChainStr string
	if len(fromKappaChainArr) > 1 {
		fromKappaChainStr = "(" + strings.Join(fromKappaChainArr, ",") + ")" // (1,'0x123'),(2,'0x456')
	} else {
		fromKappaChainStr = strings.Join(fromKappaChainArr, ",") // (1,'0x123'),(2,'0x456')
	}

	// Get all destination/to bridge events that match the kappa
	fromBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generateBridgeEventsWithKappaQuery(fromKappaChainStr, chainID, address, tokenAddress, nil, nil, page, true, false))
	if err != nil {
		return nil, fmt.Errorf("failed to get origin bridge events from identifiers: %w", err)
	}

	// Create a map of destination/to bridge events with the chain and kappa as the key
	toBridgeEventsMap := make(map[string]sql.BridgeEvent)

	// Define pending status from the origin/from bridge events.
	for _, toBridgeEvent := range toBridgeEvents {
		if toBridgeEvent.Kappa.Valid {
			key := keyGen(fmt.Sprintf("%d", toBridgeEvent.ChainID), toBridgeEvent.Kappa.String)
			toBridgeEventsMap[key] = toBridgeEvent
		}
	}

	// Iterate through all bridge events and return all partials
	for i := range fromBridgeEvents {
		fromBridgeEvent := fromBridgeEvents[i]
		key := keyGen(fromBridgeEvent.DestinationChainID.String(), fromBridgeEvent.DestinationKappa)

		// Generate partial info
		fromInfo, err := GetPartialInfoFromBridgeEventSingle(fromBridgeEvent)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}

		// If not pending, return a nonpending parital, otherwise, a pending partial.

		// Get a "to" bridge event
		toBridgeEvent := toBridgeEventsMap[key]
		if toBridgeEvent.TxHash == "" {
			continue
		}

		var swapSuccess bool
		if toBridgeEvent.SwapSuccess.Uint64() == 1 {
			swapSuccess = true
		}

		pending := false

		toInfo, err := GetPartialInfoFromBridgeEventSingle(toBridgeEvent)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}
		results = append(results, &model.BridgeTransaction{
			FromInfo:    fromInfo,
			ToInfo:      toInfo,
			Kappa:       &fromBridgeEvent.DestinationKappa,
			Pending:     &pending,
			SwapSuccess: &swapSuccess,
		})
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
