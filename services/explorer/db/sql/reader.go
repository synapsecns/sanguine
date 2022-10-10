package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"gorm.io/gorm"
	"math"
	"strconv"
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
			return nil, fmt.Errorf("failed to read event: %w", dbTx.Error)
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
func (s *Store) GetKappaFromTxHash(ctx context.Context, query string) (*string, error) {
	var res BridgeEvent

	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
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
func (s *Store) GetTokenAddressesByChainID(ctx context.Context, query string) ([]string, error) {
	var res []string
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
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
func (s *Store) BridgeEventCount(ctx context.Context, query string) (count uint64, err error) {
	var res int64
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
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
	if *typeArg == model.HistoricalResultTypeAddresses {
		dbTxFinal = s.db.WithContext(ctx).Raw(fmt.Sprintf("SELECT uniqExact(%s) FROM bridge_events %s", SenderFieldName, filter)).Scan(&sum)
	} else {
		// TODO pass table from previous query to prevent redoing this query.
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

// PartialInfosFromIdentifiers returns events given identifiers. If order is true, the events are ordered by block number.
func (s *Store) PartialInfosFromIdentifiers(ctx context.Context, query string) (partialInfos []*model.PartialInfo, err error) {
	var res []BridgeEvent

	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

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
