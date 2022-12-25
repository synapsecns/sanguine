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

func generateDeDepQuery(filter string, page *int, offset *int, latest bool) string {
	if page != nil || offset != nil {
		if latest {
			return fmt.Sprintf("SELECT * FROM (SELECT * FROM bridge_events %s ORDER BY block_number DESC, event_index DESC, insert_time DESC LIMIT %d BY chain_id) LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash", filter, *page/4)
		}
		return fmt.Sprintf("SELECT * FROM bridge_events %s ORDER BY block_number DESC, event_index DESC, insert_time DESC LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash LIMIT %d OFFSET %d", filter, *page, *offset)
	}

	return fmt.Sprintf("SELECT * FROM bridge_events %s ORDER BY block_number DESC, event_index DESC, insert_time DESC LIMIT 1 BY chain_id, contract_address, event_type, block_number, event_index, tx_hash", filter)
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
		query = fmt.Sprintf(`SELECT %s, %s AS TokenAddress, COUNT(DISTINCT (%s)) AS Count FROM (SELECT %s FROM (%s) %s) GROUP BY %s, %s ORDER BY Count Desc`,
			sql.ChainIDFieldName, sql.TokenFieldName, sql.TxHashFieldName, singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins, sql.TokenFieldName, sql.ChainIDFieldName)
	} else {
		query = fmt.Sprintf(`SELECT %s, COUNT(DISTINCT (%s)) AS Count FROM (SELECT %s FROM (%s) %s) GROUP BY %s ORDER BY Count Desc`,
			sql.ChainIDFieldName, sql.TxHashFieldName, singleSideCol, generateDeDepQuery(compositeFilters, nil, nil, false), singleSideJoins, sql.ChainIDFieldName)
	}
	return query
}

// GetPartialInfoFromBridgeEventHybrid returns the partial info from bridge event.
//
// nolint:cyclop
func GetPartialInfoFromBridgeEventHybrid(bridgeEvent sql.HybridBridgeEvent, includePending bool) (*model.BridgeTransaction, error) {
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
	bridgeTx = model.BridgeTransaction{
		FromInfo:    fromInfos,
		ToInfo:      toInfos,
		Kappa:       &bridgeEvent.FDestinationKappa,
		Pending:     &pending,
		SwapSuccess: &swapSuccess,
	}
	return &bridgeTx, nil
}

// TODO working on this @simon
// generatePartialInfoQuery returns the query for making the PartialInfo query.
//
// nolint:dupl
func generateAllBridgeEventsQueryFromOrigin(chainID *int, address, tokenAddress, txHash *string, page int, in bool, latest bool) string {
	firstFilter := true
	chainIDSpecifier := generateSingleSpecifierI32SQL(chainID, sql.ChainIDFieldName, &firstFilter, "")
	addressSpecifier := generateAddressSpecifierSQL(address, &firstFilter, "")
	tokenAddressSpecifier := generateSingleSpecifierStringSQL(tokenAddress, sql.TokenFieldName, &firstFilter, "")
	txHashSpecifier := generateSingleSpecifierStringSQL(txHash, sql.TxHashFieldName, &firstFilter, "")
	directionSpecifier := generateDirectionSpecifierSQL(in, &firstFilter, "")
	compositeFilters := chainIDSpecifier + addressSpecifier + tokenAddressSpecifier + txHashSpecifier + directionSpecifier
	pageValue := sql.PageSize
	pageOffset := (page - 1) * sql.PageSize
	finalQuery := fmt.Sprintf("SELECT %s FROM (%s) %s", originToDestCol, generateDeDepQuery(compositeFilters, &pageValue, &pageOffset, latest), originToDestJoins)
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
	finalQuery := fmt.Sprintf("SELECT %s FROM (%s) %s", destToOriginCol, generateDeDepQuery(compositeFilters, &pageValue, &pageOffset, false), destToOriginJoins)

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
func (r *queryResolver) GetBridgeTxsFromOrigin(ctx context.Context, chainID *int, address *string, txHash *string, page int, tokenAddress *string, pending bool, latest bool) ([]*model.BridgeTransaction, error) {
	var err error
	var chainMap = make(map[uint32]bool)
	var results []*model.BridgeTransaction
	allBridgeEvents, err := r.DB.GetAllBridgeEvents(ctx, generateAllBridgeEventsQueryFromOrigin(chainID, address, tokenAddress, txHash, page, true, latest))

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
