package graph

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/synapsecns/sanguine/services/explorer/db/sql"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math"
	"strconv"
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

func (r *queryResolver) originToDestinationBridge(ctx context.Context, address *string, kappa *string, includePending bool, page int, tokenAddress *string, fromInfos []*model.PartialInfo) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction

	for _, fromInfo := range fromInfos {
		txHash := common.HexToHash(*fromInfo.TxnHash)
		destinationKappa := crypto.Keccak256Hash(txHash.Bytes()).String()
		if kappa != nil {
			destinationKappa = *kappa
		}

		toInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(nil, address, tokenAddress, &destinationKappa, nil, page, true))
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}

		switch len(toInfos) {
		case 1:
			var swapSuccess bool

			toInfo := toInfos[0]
			swapBridgeEventQuery := fmt.Sprintf(
				`SELECT * FROM bridge_events WHERE %s = '%s' AND %s = %d SETTINGS readonly=1`,
				sql.KappaFieldName, destinationKappa, sql.ChainIDFieldName, *toInfo.ChainID,
			)
			swapBridgeEvent, err := r.DB.GetBridgeEvent(ctx, swapBridgeEventQuery)
			if swapBridgeEvent.SwapSuccess.Uint64() == 1 {
				swapSuccess = true
			}
			if err != nil {
				return nil, fmt.Errorf("failed to get swap success: %w", err)
			}

			pending := false
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      toInfo,
				Kappa:       &destinationKappa,
				Pending:     &pending,
				SwapSuccess: &swapSuccess,
			})
		case 0:
			if includePending {
				results = append(results, &model.BridgeTransaction{
					FromInfo:    fromInfo,
					ToInfo:      nil,
					Kappa:       &destinationKappa,
					Pending:     &includePending,
					SwapSuccess: nil,
				})
			}
		default:
			return nil, fmt.Errorf("multiple toInfos found for kappa %s", destinationKappa)
		}
	}

	return results, nil
}

// nolint:cyclop
func (r *queryResolver) destinationToOriginBridge(ctx context.Context, address *string, txnHash *string, kappa *string, page int, tokenAddress *string, toInfos []*model.PartialInfo) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction

	pending := false

	for _, toInfo := range toInfos {
		var swapSuccess bool

		// gets destination tx (to tx)
		toBridgeEventQuery := fmt.Sprintf(
			`SELECT * FROM bridge_events WHERE %s = '%s' AND %s = %d SETTINGS readonly=1`,
			sql.TxHashFieldName, *toInfo.TxnHash, sql.ChainIDFieldName, *toInfo.ChainID,
		)
		toBridgeEvent, err := r.DB.GetBridgeEvent(ctx, toBridgeEventQuery)
		if err != nil {
			return nil, fmt.Errorf("failed to get swap success: %w", err)
		}
		if toBridgeEvent.SwapSuccess.Uint64() == 1 {
			swapSuccess = true
		}

		// Gets kappa from the destination tx
		if kappa == nil {
			if toBridgeEvent.Kappa.Valid {
				kappa = &toBridgeEvent.Kappa.String
			} else {
				return nil, fmt.Errorf("failed to get kappa from bridge event: %v", toBridgeEvent)
			}
		}

		// Get origin tx
		query := fmt.Sprintf(
			`SELECT * FROM bridge_events WHERE %s = '%s' SETTINGS readonly=1`, sql.DestinationKappaFieldName, *kappa)

		originTxHash := txnHash
		if txnHash == nil {
			fromBridgeEvent, err := r.DB.GetBridgeEvent(ctx, query)
			if err != nil {
				return nil, fmt.Errorf("failed to get bridge event: %w", err)
			}

			originTxHash = &fromBridgeEvent.TxHash
		}

		fromInfos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(nil, address, tokenAddress, nil, originTxHash, page, true))
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
		}

		switch {
		case len(fromInfos) > 1:
			return nil, fmt.Errorf("multiple fromInfos found for kappa %s", *kappa)
		case len(fromInfos) == 1:
			fromInfo := fromInfos[0]
			results = append(results, &model.BridgeTransaction{
				FromInfo:    fromInfo,
				ToInfo:      toInfo,
				Kappa:       kappa,
				Pending:     &pending,
				SwapSuccess: &swapSuccess,
			})
		case len(fromInfos) == 0:
			return nil, fmt.Errorf("no fromInfo found for kappa %s", *kappa)
		}
	}

	return results, nil
}

func (r *queryResolver) originOrDestinationBridge(ctx context.Context, chainID *int, address *string, txnHash *string, kappa *string, includePending bool, page int, tokenAddress *string) ([]*model.BridgeTransaction, error) {
	var results []*model.BridgeTransaction
	var toInfos []*model.PartialInfo
	var fromInfos []*model.PartialInfo

	infos, err := r.DB.PartialInfosFromIdentifiers(ctx, generatePartialInfoQuery(chainID, address, tokenAddress, kappa, txnHash, page, false))
	if err != nil {
		return nil, fmt.Errorf("failed to get bridge events from identifiers: %w", err)
	}

	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")

	for _, info := range infos {
		txHashSpecifier := generateSingleSpecifierStringSQL(info.TxnHash, sql.TxHashFieldName, &firstFilter, "")
		query := fmt.Sprintf(`SELECT * FROM bridge_events %s%s`, chainIDSpecifier, txHashSpecifier)
		bridgeEvent, err := r.DB.GetBridgeEvent(ctx, query)
		if err != nil {
			return nil, fmt.Errorf("failed to get kappa from tx hash: %w", err)
		}

		// Check bridge event kappa.
		if !bridgeEvent.Kappa.Valid || bridgeEvent.Kappa.String == "" {
			fromInfos = append(fromInfos, info)
		} else {
			toInfos = append(toInfos, info)
		}
	}

	originResults, err := r.originToDestinationBridge(ctx, nil, nil, includePending, page, tokenAddress, fromInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to get origin -> destination bridge transactions: %w", err)
	}

	destinationResults, err := r.destinationToOriginBridge(ctx, nil, nil, nil, page, tokenAddress, toInfos)
	if err != nil {
		return nil, fmt.Errorf("failed to get destination -> origin bridge transactions: %w", err)
	}

	results = r.mergeBridgeTransactions(originResults, destinationResults)

	return results, nil
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

// generatePartialInfoQuery returns the query for making the PartialInfo query.
func generatePartialInfoQuery(chainID *int, address, tokenAddress, kappa, txHash *string, page int, latest bool) string {
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "t1.")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "t1.")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "t1.")
	kappaSpecifier := generateSingleSpecifierStringSQL(kappa, sql.KappaFieldName, &firstFilter, "t1.")
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
func generateBridgeEventCountQuery(chainID int, address *string, tokenAddress *string, directionIn bool, timestamp *uint64) string {
	chainField := sql.ChainIDFieldName

	if directionIn {
		chainField = sql.DestinationChainIDFieldName
	}

	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(&chainID, chainField, &firstFilter, "")
	addressSpecifier := generateSingleSpecifierStringSQL(address, sql.RecipientFieldName, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	timestampSpecifier := generateTimestampSpecifierSQL(timestamp, sql.TimeStampFieldName, &firstFilter, "")
	query := fmt.Sprintf(`SELECT COUNT(DISTINCT (%s, %s)) FROM bridge_events %s%s%s%s`,
		sql.TxHashFieldName, sql.EventIndexFieldName, chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, timestampSpecifier)

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
	var err error
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

		var formattedValue float64

		if res[i].TokenDecimal != nil {
			formattedValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse float: %w", err)
			}

			formattedValue /= math.Pow10(int(*res[i].TokenDecimal))
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
			FormattedValue: &formattedValue,
			USDValue:       res[i].AmountUSD,
			TokenAddress:   &res[i].Token,
			TokenSymbol:    &tokenSymbol,
			BlockNumber:    &blockNumberInt,
			Time:           &timeStamp,
		})
	}

	return partialInfos, nil
}

// generateToPartialInfoQuery returns the query for making the PartialInfo query.
func generateToPartialInfoQuery(toKappaChainStr string, page int) string {
	pageSpecifier := fmt.Sprintf(" ORDER BY %s DESC LIMIT %d OFFSET %d", sql.BlockNumberFieldName, sql.PageSize, (page-1)*sql.PageSize)
	compositeIdentifiers := fmt.Sprintf("WHERE (%s,%s) IN %s", sql.ChainIDFieldName, sql.KappaFieldName, toKappaChainStr) + pageSpecifier
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
	deDup := deDupInQueryLatest
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
	var err error
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

		var formattedValue float64

		if res[i].TokenDecimal != nil {
			formattedValue, err = strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("failed to parse float: %w", err)
			}

			formattedValue /= math.Pow10(int(*res[i].TokenDecimal))
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
			FormattedValue: &formattedValue,
			USDValue:       res[i].AmountUSD,
			TokenAddress:   &res[i].Token,
			TokenSymbol:    &tokenSymbol,
			BlockNumber:    &blockNumberInt,
			Time:           &timeStamp,
		}
	}

	return partialInfos, nil
}
