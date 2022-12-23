package graph

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
)

const columnRename = "if(ti.token_address = '', be.token, ti.token_address) AS token, if(se.tokens_bought > 0 , se.tokens_bought, be.amount) AS amount, be.token_symbol AS token_symbol, if(se.amount_usd[ti.token_index] > 0, ((toFloat64(amount)/exp10(be.token_decimal)) * se.amount_usd[ti.token_index]), be.amount_usd) AS amount_usd, be.event_type AS event_type,be.token AS token_raw, be.tx_hash AS tx_hash,be.chain_id AS chain_id,be.contract_address AS contract_address, be.token_symbol AS token_symbol,be.destination_kappa AS destination_kappa,be.sender AS sender, be.recipient AS recipient, be.recipient_bytes AS recipient_bytes,be.fee AS fee, be.kappa AS kappa, be.token_index_from AS token_index_from, be.token_index_to AS token_index_to, be.min_dy AS min_dy, be.deadline AS deadline, be.swap_success AS swap_success, be.swap_token_index AS swap_token_index, be.swap_min_amount AS swap_min_amount, be.swap_deadline AS swap_deadline, be.token_id AS token_id, be.fee_amount_usd AS fee_amount_usd, be.token_decimal AS token_decimal, be.timestamp AS timestamp,be.destination_chain_id AS destination_chain_id, be.insert_time AS insert_time"
const joinsToOrigin = "SELECT DISTINCT ON (be.chain_id, be.contract_address, be.event_type, be.block_number, be.event_index, be.tx_hash) be.*, se.*, ti.* FROM bridge_events be LEFT JOIN swap_events se ON be.tx_hash = se.tx_hash AND be.chain_id = se.chain_id LEFT JOIN (SELECT DISTINCT ON (chain_id, token_index, contract_address) * FROM token_indices) ti ON se.chain_id = ti.chain_id AND se.contract_address = ti.contract_address AND ti.token_index = se.bought_id"
const joins = "SELECT DISTINCT ON (be.chain_id, be.contract_address, be.event_type, be.block_number, be.event_index, be.tx_hash) be.*, se.*, ti.* FROM bridge_events be LEFT JOIN swap_events se ON be.tx_hash = se.tx_hash AND be.chain_id = se.chain_id LEFT JOIN (SELECT DISTINCT ON (chain_id, token_index, contract_address) * FROM token_indices) ti ON se.chain_id = ti.chain_id AND se.contract_address = ti.contract_address AND ti.token_index = se.sold_id"
const simpleDeDup = "SELECT DISTINCT ON (chain_id, contract_address, event_type, block_number, event_index, tx_hash) * FROM bridge_events ORDER BY block_number DESC, event_index DESC"

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

// generateDirectionSpecifierSQL generates a where function with a string.
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
//
// nolint:dupl
func generatePartialInfoQuery(chainID *int, address, tokenAddress, kappa, txHash *string, page int, latest bool, destination bool) string {
	firstFilter := true
	fmt.Println(latest)
	generatePartialInfoQuerySimpleNew(chainID, address, tokenAddress, kappa, txHash, page)
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "be.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "be.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "be.")
	kappaSpecifier := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "be.")
	destinationChainIDSpecifier := generateDestinationChainIDSpecifierSQL(sql.DestinationChainIDFieldName, &firstFilter, "be.", destination)
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "be.")
	pageSpecifier := fmt.Sprintf(" ORDER BY be.%s DESC, be.%s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.EventIndexFieldName, sql.PageSize, (page-1)*sql.PageSize)
	compositeIdentifiers := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + destinationChainIDSpecifier + txHashSpecifier + pageSpecifier
	query := fmt.Sprintf("SELECT %s FROM (%s %s)  ", columnRename, joins, compositeIdentifiers)
	return query
}

// TODO working on this @simon
// generatePartialInfoQuery returns the query for making the PartialInfo query.
//
// nolint:dupl
func generatePartialInfoQuerySimpleNew(chainID *int, address, tokenAddress, kappa, txHash *string, page int) string {
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "be.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "be.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "be.")
	kappaSpecifier := generateKappaSpecifierSQL(kappa, sql.KappaFieldName, &firstFilter, "be.")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "be.")
	pageSpecifier := fmt.Sprintf(" ORDER BY be.%s DESC, be.%s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.EventIndexFieldName, sql.PageSize, (page-1)*sql.PageSize)
	compositeIdentifiers := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + pageSpecifier
	fromBEquery := fmt.Sprintf("SELECT %s FROM (%s %s)  ", columnRename, joins, compositeIdentifiers)
	toBEquery := fmt.Sprintf("SELECT %s FROM (%s %s)  ", columnRename, joinsToOrigin, compositeIdentifiers)
	finalQuery := fmt.Sprintf("SELECT * FROM (%s) f LEFT JOIN (%s) t ON f.%s = t.%s AND f.%s = t.%s", fromBEquery, toBEquery, sql.DestinationChainIDFieldName, sql.ChainIDFieldName, sql.DestinationKappaFieldName, sql.KappaFieldName)
	fmt.Println("Test Query", finalQuery)

	return finalQuery
}

// generateBridgeEventCountQuery creates the query for bridge event count.
func generateBridgeEventCountQuery(chainID *int, address *string, tokenAddress *string, directionIn bool, timestamp *uint64, isTokenCount bool) string {
	chainField := sql.ChainIDFieldName

	firstFilter := true
	directionSpecifier := generateDirectionSpecifierSQL(directionIn, &firstFilter, "be.")
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, chainField, &firstFilter, "be.")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.RecipientFieldName, &firstFilter, "be.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "be.")
	timestampSpecifier := generateTimestampSpecifierSQL(timestamp, sql.TimeStampFieldName, &firstFilter, "be.")

	compositeFilters := fmt.Sprintf(
		`%s%s%s%s%s`,
		directionSpecifier, chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, timestampSpecifier,
	)
	var query string
	if isTokenCount {
		query = fmt.Sprintf(`SELECT %s, %s AS TokenAddress, COUNT(DISTINCT (%s)) AS Count FROM (SELECT %s FROM (%s %s)) GROUP BY %s, %s ORDER BY Count Desc`,
			sql.ChainIDFieldName, sql.TokenFieldName, sql.TxHashFieldName, columnRename, joins, compositeFilters, sql.TokenFieldName, sql.ChainIDFieldName)
	} else {
		query = fmt.Sprintf(`SELECT %s, COUNT(DISTINCT (%s)) AS Count FROM (SELECT %s FROM (%s %s)) GROUP BY %s ORDER BY Count Desc`,
			sql.ChainIDFieldName, sql.TxHashFieldName, columnRename, joins, compositeFilters, sql.ChainIDFieldName)
	}
	return query
}

// GetPartialInfoFromBridgeEventSingle returns the partial info from bridge event.
//
// nolint:cyclop
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
		switch res.TokenSymbol.String {
		case "nETH":
			tokenSymbol = "ETH"
		case "nUSD":
			tokenSymbol = "USDC"
		default:
			tokenSymbol = res.TokenSymbol.String
		}
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
	fmt.Println(latest)
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "be.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "be.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "be.")
	kappaSpecifier := generateSingleSpecifierStringSQL(kappa, sql.KappaFieldName, &firstFilter, "be.")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "be.")

	pageSpecifier := fmt.Sprintf(" ORDER BY be.%s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.PageSize, (page-1)*sql.PageSize)

	var kappaChainSpecifier string
	if to {
		kappaChainSpecifier = fmt.Sprintf("WHERE (be.%s,be.%s) IN %s", sql.ChainIDFieldName, sql.KappaFieldName, kappaChainStr)
	} else {
		kappaChainSpecifier = fmt.Sprintf("WHERE (be.%s,be.%s) IN %s", sql.DestinationChainIDFieldName, sql.DestinationKappaFieldName, kappaChainStr)
	}
	compositeIdentifiers := kappaChainSpecifier + chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + pageSpecifier

	query := fmt.Sprintf("SELECT %s FROM (%s %s)  ", columnRename, joinsToOrigin, compositeIdentifiers)
	return query
}

// generatePartialInfoQueryByChain returns the query for making the PartialInfo query.
func generatePartialInfoQueryByChain(limitSize int) string {
	pageSpecifier := fmt.Sprintf(" ORDER BY (be.%s,be.%s) DESC", sql.BlockNumberFieldName, sql.EventIndexFieldName)
	limitSpecifier := fmt.Sprintf("LIMIT %d BY be.%s", limitSize, sql.ChainIDFieldName)
	compositeIdentifiers := pageSpecifier
	query := fmt.Sprintf("SELECT %s FROM (%s %s) %s  ", columnRename, joins, compositeIdentifiers, limitSpecifier)
	return query
}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromOrigin(ctx context.Context, chainID *int, address *string, txnHash *string, includePending bool, page int, tokenAddress *string, latest bool) ([]*model.BridgeTransaction, error) {
	var err error
	// test, _ := r.GetBridgeTxsFromDestinationNew(ctx, chainID, address, txnHash, nil, page, tokenAddress)
	// fmt.Println(test)
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
		if bridgeEvent.DestinationChainID.Uint64() > 0 {
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
			if latest {
				chainCheck[fmt.Sprintf("%d", fromBridgeEvent.ChainID)] = true
			}
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      toInfo,
				Kappa:       &fromBridgeEvent.DestinationKappa,
				Pending:     &pending,
				SwapSuccess: &swapSuccess,
			})
		} else if includePending {
			if latest {
				chainCheck[fmt.Sprintf("%d", fromBridgeEvent.ChainID)] = true
			}
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      nil,
				Kappa:       &fromBridgeEvent.DestinationKappa,
				Pending:     &includePending,
				SwapSuccess: nil,
			})
		}
	}
	return results, nil
}

// nolint:gocognit,cyclop
// TODO finish this to speed up tx queries.
// func (r *queryResolver) GetBridgeTxsFromDestinationNew(ctx context.Context, chainID *int, address *string, txHash *string, kappa *string, page int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
//	var results []*model.BridgeTransaction
//	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generatePartialInfoQuerySimpleNew(chainID, address, tokenAddress, kappa, txHash, page))
//	if err != nil {
//		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
//	}
//
//	for i := range allBridgeEvents {
//		bridgeEvent := allBridgeEvents[i]
//		fmt.Println("FROMEVENT", bridgeEvent.FChainID, "TOEVENT", bridgeEvent.TChainID)
//
//		// key := keyGen(fromBridgeEvent.DestinationChainID.String(), fromBridgeEvent.DestinationKappa)
//		//
//		//// Generate partial info
//		//fromInfo, err := GetPartialInfoFromBridgeEventSingle(fromBridgeEvent)
//		//if err != nil {
//		//	return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
//		//}
//		//
//		//// If not pending, return a nonpending parital, otherwise, a pending partial.
//		//
//		//// Get a "to" bridge event
//		//toBridgeEvent := toBridgeEventsMap[key]
//		//if toBridgeEvent.TxHash == "" {
//		//	continue
//		//}
//		//
//		//var swapSuccess bool
//		//if toBridgeEvent.SwapSuccess.Uint64() == 1 {
//		//	swapSuccess = true
//		//}
//		//
//		//pending := false
//		//
//		//toInfo, err := GetPartialInfoFromBridgeEventSingle(toBridgeEvent)
//		//if err != nil {
//		//	return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
//		//}
//		//results = append(results, &model.BridgeTransaction{
//		//	FromInfo:    fromInfo,
//		//	ToInfo:      toInfo,
//		//	Kappa:       &fromBridgeEvent.DestinationKappa,
//		//	Pending:     &pending,
//		//	SwapSuccess: &swapSuccess,
//		//})
//	}
//	return results, nil
//}

// nolint:gocognit,cyclop
func (r *queryResolver) GetBridgeTxsFromDestination(ctx context.Context, chainID *int, address *string, txHash *string, kappa *string, page int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	toBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generatePartialInfoQuery(chainID, address, tokenAddress, kappa, txHash, page, false, true))

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
