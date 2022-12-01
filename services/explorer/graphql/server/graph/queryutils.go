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

func (r *queryResolver) getChainIDs(ctx context.Context, chainID *int) ([]int, error) {
	var chainIDs []int

	// If the chain ID is not specified, get all chain IDs.
	if chainID == nil {
		chainIDsInt, err := r.DB.GetAllChainIDs(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to get all chain IDs: %w", err)
		}

		chainIDs = append(chainIDs, chainIDsInt...)
	} else {
		chainIDs = append(chainIDs, *chainID)
	}

	return chainIDs, nil
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
	uniqueBridgeTransactions := make(map[*model.BridgeTransaction]bool)

	for _, originTx := range origin {
		uniqueBridgeTransactions[originTx] = true
	}

	for _, destinationTx := range destination {
		uniqueBridgeTransactions[destinationTx] = true
	}

	for tx := range uniqueBridgeTransactions {
		results = append(results, tx)
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

// GenerateSingleSpecifierStringSQL generates a where function with a string.
func generateKappaSpecifierStringSQL(value *string, field string, firstFilter *bool, tablePrefix string, destination bool) string {
	if value != nil {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE %s%s = '%s'", tablePrefix, field, *value)
		}

		return fmt.Sprintf(" AND %s%s = '%s'", tablePrefix, field, *value)
	}
	if destination {
		if *firstFilter {
			*firstFilter = false

			return fmt.Sprintf(" WHERE isNotNull(%s%s) AND %s%s != ''", tablePrefix, field, tablePrefix, field)
		}

		return fmt.Sprintf(" AND isNotNull(%s%s) AND %s%s != ''", tablePrefix, field, tablePrefix, field)
	}
	return ""
}

// generatePartialInfoQuery returns the query for making the PartialInfo query.
func generatePartialInfoQuery(chainID *int, address, tokenAddress, kappa, txHash *string, page int, latest bool, destination bool) string {
	firstFilter := true

	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t1.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "t1.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "t1.")
	kappaSpecifier := generateKappaSpecifierStringSQL(kappa, sql.KappaFieldName, &firstFilter, "t1.", destination)
	//if destination && kappaSpecifier == "" {
	//	if firstFilter {
	//		kappaSpecifier = fmt.Sprintf("WHERE %s != NULL", sql.KappaFieldName)
	//		firstFilter = false
	//	} else {
	//		kappaSpecifier = fmt.Sprintf("AND %s != NULL", sql.KappaFieldName)
	//	}
	//
	//}
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "t1.")
	pageSpecifier := fmt.Sprintf(" ORDER BY %s, %s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.EventIndexFieldName, sql.PageSize, (page-1)*sql.PageSize)

	compositeIdentifiers := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + pageSpecifier

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
func generateBridgeEventCountQuery(chainID *int, address *string, tokenAddress *string, directionIn bool, timestamp *uint64) string {
	chainField := sql.ChainIDFieldName

	if directionIn {
		chainField = sql.DestinationChainIDFieldName
	}

	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, chainField, &firstFilter, "")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.RecipientFieldName, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	timestampSpecifier := generateTimestampSpecifierSQL(timestamp, sql.TimeStampFieldName, &firstFilter, "")
	query := fmt.Sprintf(`SELECT %s, COUNT(DISTINCT (%s)) AS Count FROM bridge_events %s%s%s%s GROUP BY %s`,
		sql.ChainIDFieldName, sql.TxHashFieldName, chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, timestampSpecifier, sql.ChainIDFieldName)

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

func GetPartialInfoFromBridgeEvent(res []sql.BridgeEvent) ([]*model.PartialInfo, error) {
	var partialInfos []*model.PartialInfo
	for i := range res {
		chainIDInt := int(res[i].ChainID)
		blockNumberInt := int(res[i].BlockNumber)

		var recipient string

		switch {
		case res[i].Recipient.Valid:
			recipient = res[i].Recipient.String
		case res[i].RecipientBytes.Valid:
			recipient = res[i].RecipientBytes.String
		default:
			return nil, fmt.Errorf("recipient is not valid")
		}

		var tokenSymbol string

		if res[i].TokenSymbol.Valid && res[i].TokenSymbol.String != "" {
			tokenSymbol = res[i].TokenSymbol.String
		} else {
			return nil, fmt.Errorf("token symbol is not valid")
		}

		value := res[i].Amount.String()

		var formattedValue *float64

		if res[i].TokenDecimal != nil {
			formattedValue = getAdjustedValue(res[i].Amount, *res[i].TokenDecimal)
		} else {
			return nil, fmt.Errorf("token decimal is not valid")
		}

		var timeStamp int

		if res[i].TimeStamp != nil {
			timeStamp = int(*res[i].TimeStamp)
		} else {
			return nil, fmt.Errorf("time stamp is not valid")
		}

		partialInfos = append(partialInfos, &model.PartialInfo{
			ChainID:        &chainIDInt,
			Address:        &recipient,
			TxnHash:        &res[i].TxHash,
			Value:          &value,
			FormattedValue: formattedValue,
			USDValue:       res[i].AmountUSD,
			TokenAddress:   &res[i].Token,
			TokenSymbol:    &tokenSymbol,
			BlockNumber:    &blockNumberInt,
			Time:           &timeStamp,
		})
	}

	return partialInfos, nil
}

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

// generateToPartialInfoQuery returns the query for making the PartialInfo query.
func generateToKappaPartialInfoQuery(toKappaChainStr string, chainID *int, address, tokenAddress, kappa, txHash *string, page int, latest bool) string {
	firstFilter := false
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t1.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "t1.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "t1.")
	kappaSpecifier := generateSingleSpecifierStringSQL(kappa, sql.KappaFieldName, &firstFilter, "t1.")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "t1.")

	pageSpecifier := fmt.Sprintf(" ORDER BY %s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.PageSize, (page-1)*sql.PageSize)
	compositeIdentifiers := fmt.Sprintf("WHERE (%s,%s) IN %s", sql.ChainIDFieldName, sql.KappaFieldName, toKappaChainStr) + chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + pageSpecifier
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
func GetToPartialInfoFromBridgeEvent(res []sql.BridgeEvent) (map[string]*model.PartialInfo, error) {
	partialInfos := make(map[string]*model.PartialInfo)
	for i := range res {
		chainIDInt := int(res[i].ChainID)
		blockNumberInt := int(res[i].BlockNumber)

		var recipient string

		switch {
		case res[i].Recipient.Valid:
			recipient = res[i].Recipient.String
		case res[i].RecipientBytes.Valid:
			recipient = res[i].RecipientBytes.String
		default:
			return nil, fmt.Errorf("recipient is not valid")
		}

		var tokenSymbol string

		if res[i].TokenSymbol.Valid && res[i].TokenSymbol.String != "" {
			tokenSymbol = res[i].TokenSymbol.String
		} else {
			return nil, fmt.Errorf("token symbol is not valid")
		}

		value := res[i].Amount.String()

		var formattedValue *float64

		if res[i].TokenDecimal != nil {
			formattedValue = getAdjustedValue(res[i].Amount, *res[i].TokenDecimal)
		} else {
			return nil, fmt.Errorf("token decimal is not valid")
		}

		var timeStamp int

		if res[i].TimeStamp != nil {
			timeStamp = int(*res[i].TimeStamp)
		} else {
			return nil, fmt.Errorf("time stamp is not valid")
		}
		key := fmt.Sprintf("%d", res[i].ChainID)
		partialInfos[key] = &model.PartialInfo{
			ChainID:        &chainIDInt,
			Address:        &recipient,
			TxnHash:        &res[i].TxHash,
			Value:          &value,
			FormattedValue: formattedValue,
			USDValue:       res[i].AmountUSD,
			TokenAddress:   &res[i].Token,
			TokenSymbol:    &tokenSymbol,
			BlockNumber:    &blockNumberInt,
			Time:           &timeStamp,
		}
	}

	return partialInfos, nil
}

// generateFromKappaPartialInfoQuery returns the query for making the PartialInfo query.
func generateFromKappaPartialInfoQuery(fromKappaChainStr string, chainID *int, address, tokenAddress, kappa, txHash *string, page int, latest bool) string {
	firstFilter := false
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t1.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "t1.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "t1.")
	kappaSpecifier := generateSingleSpecifierStringSQL(kappa, sql.KappaFieldName, &firstFilter, "t1.")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "t1.")

	pageSpecifier := fmt.Sprintf(" ORDER BY %s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.PageSize, (page-1)*sql.PageSize)
	compositeIdentifiers := fmt.Sprintf("WHERE (%s,%s) IN %s", sql.DestinationChainIDFieldName, sql.DestinationKappaFieldName, fromKappaChainStr) + chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + kappaSpecifier + txHashSpecifier + pageSpecifier
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
	toBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generateToKappaPartialInfoQuery(toKappaChainStr, nil, nil, nil, nil, nil, page, true))
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
		if latest && chainCheck[fmt.Sprintf("%d", fromBridgeEvent.DestinationChainID)] {
			continue
		}
		key := keyGen(fromBridgeEvent.DestinationChainID.String(), fromBridgeEvent.DestinationKappa)

		// Generate partial info
		fromInfo, err := GetPartialInfoFromBridgeEventSingle(fromBridgeEvent)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}

		// If not pending, return a non pending parital, otherwise, a pending partial.
		if fromBridgeEventsKappaStatusMap[key] {
			// Get a "to" bridge event
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
				chainCheck[fmt.Sprintf("%d", fromBridgeEvent.DestinationChainID)] = true
			}
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      toInfo,
				Kappa:       &toBridgeEvent.Kappa.String,
				Pending:     &pending,
				SwapSuccess: &swapSuccess,
			})
		} else {
			if includePending {
				if latest {
					chainCheck[fmt.Sprintf("%d", fromBridgeEvent.DestinationChainID)] = true
				}
				results = append(results, &model.BridgeTransaction{
					FromInfo:    fromInfo,
					ToInfo:      nil,
					Kappa:       &fromBridgeEvent.Kappa.String,
					Pending:     &includePending,
					SwapSuccess: nil,
				})
			}
		}
	}
	return results, nil
}

func (r *queryResolver) GetBridgeTxsFromDestination(ctx context.Context, chainID *int, address *string, txnHash *string, kappa *string, page int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	var err error
	var results []*model.BridgeTransaction
	// Get all bridge transactions
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
	fromBridgeEvents, err := r.DB.GetBridgeEvents(ctx, generateFromKappaPartialInfoQuery(fromKappaChainStr, nil, nil, nil, nil, nil, page, true))
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

		// If not pending, return a non pending parital, otherwise, a pending partial.

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
			Kappa:       &toBridgeEvent.Kappa.String,
			Pending:     &pending,
			SwapSuccess: &swapSuccess,
		})
	}
	return results, nil
}

// getAdjustedValue gets the adjusted value
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
