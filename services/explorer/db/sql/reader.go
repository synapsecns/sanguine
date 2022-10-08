package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"gorm.io/gorm"
)

// EventType is an enum for event types.
type EventType int8

const (
	// Bridge - SynapseBridge event.
	Bridge int8 = iota
	// Swap - SwapFlashLoan event.
	Swap
)

// ReadBlockNumberByChainID provides an easy-to-use interface to validate database
// data from a recent write event via chain id.
func (s *Store) ReadBlockNumberByChainID(ctx context.Context, eventType int8, chainID uint32) (*uint64, error) {
	// If reading a bridge event
	var blockNumber uint64
	switch eventType {
	case Bridge:
		var resp BridgeEvent
		dbTx := s.DB().WithContext(ctx).
			Find(&resp, "chain_id = ?", chainID)
		if dbTx.Error != nil {
			return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
		blockNumber = resp.BlockNumber

	// If reading a swap event
	case Swap:
		var resp SwapEvent
		dbTx := s.DB().WithContext(ctx).
			Find(&resp, "chain_id = ?", chainID)
		if dbTx.Error != nil {
			return nil, fmt.Errorf("failed to store read event: %w", dbTx.Error)
		}
		blockNumber = resp.BlockNumber
	}

	return &blockNumber, nil
}

// GetTxHashFromKappa returns the transaction hash for a given kappa.
func (s *Store) GetTxHashFromKappa(ctx context.Context, kappa string) (*string, error) {
	var res BridgeEvent
	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
		`SELECT * FROM bridge_events WHERE %s = '%s'`,
		DestinationKappaFieldName, kappa,
	)).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return &res.TxHash, nil
}

// GetKappaFromTxHash returns the kappa for a given transaction hash.
func (s *Store) GetKappaFromTxHash(ctx context.Context, txHash string, chainID *uint32) (*string, error) {
	var res BridgeEvent
	var chainIDSpecifier string
	if chainID != nil {
		chainIDSpecifier = fmt.Sprintf(" AND %s = %d", ChainIDFieldName, *chainID)
	}
	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
		`SELECT * FROM bridge_events WHERE %s = '%s'%s`,
		TxHashFieldName, txHash, chainIDSpecifier,
	)).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	// nolint:nilnil
	if !res.Kappa.Valid || res.Kappa.String == "" {
		return nil, nil
	}

	return &res.Kappa.String, nil
}

// GetSwapSuccess returns if an event had a successful swap.
func (s *Store) GetSwapSuccess(ctx context.Context, kappa string, chainID uint32) (*bool, error) {
	var res BridgeEvent
	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
		`SELECT * FROM bridge_events WHERE %s = '%s' AND %s = %d`,
		KappaFieldName, kappa, ChainIDFieldName, chainID,
	)).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	if res.SwapSuccess == nil {
		return nil, fmt.Errorf("GetSwapSuccess returned a nil: %w", dbTx.Error)
	}
	if res.SwapSuccess.Uint64() == 1 {
		trueVal := true
		return &trueVal, nil
	}
	falseVal := false
	return &falseVal, nil
}

// GetAllChainIDs gets all chain IDs that have been used in bridge events.
func (s *Store) GetAllChainIDs(ctx context.Context) ([]uint32, error) {
	var res []uint32
	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
		`SELECT DISTINCT %s FROM bridge_events UNION DISTINCT SELECT DISTINCT toUInt32(%s) FROM bridge_events`,
		ChainIDFieldName, DestinationChainIDFieldName,
	)).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetTokenAddressesByChainID gets all token addresses that have been used in bridge events for a given chain ID.
func (s *Store) GetTokenAddressesByChainID(ctx context.Context, chainID uint32) ([]string, error) {
	var res []string
	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
		`SELECT DISTINCT %s FROM bridge_events WHERE %s = %d OR %s = %d`,
		TokenFieldName, ChainIDFieldName, chainID, DestinationChainIDFieldName, chainID,
	)).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetBridgeStatistic gets the bridge statistics.
func (s *Store) GetBridgeStatistic(ctx context.Context, subQuery string) (*string, error) {
	var res float64
	dbTx := s.db.WithContext(ctx).Raw(subQuery).Find(&res)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	output := fmt.Sprintf("%f", res)
	if len(output) == 0 {
		return nil, fmt.Errorf("GetBridgeStatistic returned nil: %w", dbTx.Error)
	}
	return &output, nil
}

// BridgeEventCount returns the number of bridge events.
func (s *Store) BridgeEventCount(ctx context.Context, chainID uint32, address *string, tokenAddress *string, directionIn bool, firstBlock uint64) (count uint64, err error) {
	var res int64
	var addressSpecifier string
	if address != nil {
		addressSpecifier = fmt.Sprintf(" AND %s = '%s'", RecipientFieldName, *address)
	}
	var tokenAddressSpecifier string
	if tokenAddress != nil {
		tokenAddressSpecifier = fmt.Sprintf(" AND %s = '%s'", TokenFieldName, *tokenAddress)
	}

	if directionIn {
		dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
			`SELECT COUNT(DISTINCT (%s, %s)) FROM bridge_events WHERE %s = %d AND %s >= %d%s%s`,
			TxHashFieldName, EventIndexFieldName, DestinationChainIDFieldName, chainID, BlockNumberFieldName, firstBlock, addressSpecifier, tokenAddressSpecifier,
		)).Find(&res)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
	} else {
		dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
			`SELECT COUNT(DISTINCT %s, %s) FROM bridge_events WHERE %s = %d AND %s >= %d%s%s`,
			TxHashFieldName, EventIndexFieldName, ChainIDFieldName, chainID, BlockNumberFieldName, firstBlock, addressSpecifier, tokenAddressSpecifier,
		)).Find(&res)
		if dbTx.Error != nil {
			return 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
		}
	}

	return uint64(res), nil
}

// GetTransactionCountForEveryAddress gets the count of transactions (origin) for each address per chain id.
func (s *Store) GetTransactionCountForEveryAddress(ctx context.Context, subQuery string) ([]*model.AddressRanking, error) {
	var res []*model.AddressRanking
	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(`SELECT %s AS address, COUNT(DISTINCT %s) AS count FROM (%s) GROUP BY address ORDER BY count DESC`, TokenFieldName, TxHashFieldName, subQuery)).Scan(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	if len(res) == 0 {
		return nil, nil
	}
	return res, nil
}

// GetHistoricalData gets historical data for an address.
func (s *Store) GetHistoricalData(ctx context.Context, subQuery string, typeArg *model.HistoricalResultType, filter string) (*model.HistoricalResult, error) {
	// TODO clean up
	var res []*model.DateResult

	// Get day by day data.
	dbTx := s.db.WithContext(ctx).Raw(subQuery).Scan(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	var sum float64
	var dbTxFinal *gorm.DB

	// Get the rest of the data depending on query type.
	switch *typeArg {
	case model.HistoricalResultTypeAddresses:
		dbTxFinal = s.db.WithContext(ctx).Raw(fmt.Sprintf("SELECT uniqExact(%s) FROM bridge_events %s", SenderFieldName, filter)).Scan(&sum)
	case model.HistoricalResultTypeTransactions:
		dbTxFinal = s.db.WithContext(ctx).Raw(fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)).Scan(&sum)
	case model.HistoricalResultTypeBridgevolume: // Extra in case things change.
		dbTxFinal = s.db.WithContext(ctx).Raw(fmt.Sprintf("SELECT sumKahan(total) FROM (%s)", subQuery)).Scan(&sum)
	}
	if dbTxFinal.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	payload := model.HistoricalResult{
		Total:       &sum,
		DateResults: res,
		Type:        typeArg,
	}

	return &payload, nil
}

// GenerateAddressSpecifierSQL generates a where function with an string.
func GenerateAddressSpecifierSQL(address *string, firstFilter *bool) string {
	if address != nil {
		if *firstFilter {
			*firstFilter = false
			return fmt.Sprintf(" WHERE (%s = '%s' OR %s = '%s')", RecipientFieldName, *address, SenderFieldName, *address)
		}
		return fmt.Sprintf(" AND (%s = '%s OR %s = '%s)", RecipientFieldName, *address, SenderFieldName, *address)
	}
	return ""
}

// GenerateSingleSpecifierI32SQL generates a where function with an uint32.
func GenerateSingleSpecifierI32SQL(value *uint32, field string, firstFilter *bool) string {
	if value != nil {
		if *firstFilter {
			return fmt.Sprintf(" WHERE %s = %d", field, *value)
		}
		return fmt.Sprintf("AND %s = %d", field, *value)
	}
	return ""
}

// GenerateSingleSpecifierStringSQL generates a where function with a string.
func GenerateSingleSpecifierStringSQL(value *string, field string, firstFilter *bool) string {
	if value != nil {
		if *firstFilter {
			return fmt.Sprintf(" WHERE %s = '%s'", field, *value)
		}
		return fmt.Sprintf("AND %s = '%s'", field, *value)
	}
	return ""
}

// PartialInfosFromIdentifiers returns events given identifiers. If order is true, the events are ordered by block number.
func (s *Store) PartialInfosFromIdentifiers(ctx context.Context, chainID *uint32, address, tokenAddress, kappa, txHash *string, page int) (partialInfos []*model.PartialInfo, err error) {
	var res []BridgeEvent
	firstFilter := true
	chainIDSpecifier := GenerateSingleSpecifierI32SQL(chainID, ChainIDFieldName, &firstFilter)
	addressSpecifier := GenerateAddressSpecifierSQL(address, &firstFilter)
	tokenAddressSpecifier := GenerateSingleSpecifierStringSQL(tokenAddress, TokenFieldName, &firstFilter)
	kappaSpecifier := GenerateSingleSpecifierStringSQL(kappa, KappaFieldName, &firstFilter)
	txHashSpecifier := GenerateSingleSpecifierStringSQL(txHash, TxHashFieldName, &firstFilter)

	pageSpecifier := fmt.Sprintf(" ORDER BY %s DESC LIMIT %d OFFSET %d", BlockNumberFieldName, PageSize, (page-1)*PageSize)

	compositeIdentifiers := fmt.Sprintf(
		`%s%s%s%s%s%s`,
		chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, kappaSpecifier, txHashSpecifier, pageSpecifier,
	)

	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
		`SELECT * FROM bridge_events %s`,
		compositeIdentifiers,
	)).Find(&res)
	fmt.Printf("res: %+v\n", res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	// if kappa != nil {
	//	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
	//		`SELECT * FROM bridge_events WHERE %s = %s AND %s`,
	//		KappaFieldName, *kappa, compositeIdentifiers,
	//	)).Find(&res)
	//	if dbTx.Error != nil {
	//		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	//	}
	// } else if txHash != nil {
	//	fmt.Println(fmt.Sprintf(
	//		`SELECT * FROM bridge_events WHERE %s = '%s' AND %s`,
	//		TxHashFieldName, *txHash, compositeIdentifiers,
	//	))
	//	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
	//		`SELECT * FROM bridge_events WHERE %s = '%s' AND %s`,
	//		TxHashFieldName, *txHash, compositeIdentifiers,
	//	)).Find(&res)
	//	if dbTx.Error != nil {
	//		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	//	}
	// } else {
	//	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
	//		`SELECT * FROM bridge_events WHERE %s`,
	//		compositeIdentifiers,
	//	)).Find(&res)
	//	if dbTx.Error != nil {
	//		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	//	}
	//}

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
			recipient = ""
		}

		partialInfos = append(partialInfos, &model.PartialInfo{
			ChainID:      &chainIDInt,
			Address:      &recipient,
			TxnHash:      &res[i].TxHash,
			TokenAddress: &res[i].Token,
			BlockNumber:  &blockNumberInt,
		})
	}

	return partialInfos, nil
}
