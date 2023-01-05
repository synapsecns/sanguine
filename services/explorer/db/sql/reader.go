package sql

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
)

/*╔══════════════════════════════════════════════════════════════════════╗*\
▏*║                        Generic Read Functions                        ║*▕
\*╚══════════════════════════════════════════════════════════════════════╝*/

// GetUint64 gets a uint64 from a given query.
func (s *Store) GetUint64(ctx context.Context, query string) (uint64, error) {
	var res int64

	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return uint64(res), nil
}

// GetFloat64 gets a float64 from a given query.
func (s *Store) GetFloat64(ctx context.Context, query string) (float64, error) {
	var res float64
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return 0, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetStringArray returns a string array for a given query.
func (s *Store) GetStringArray(ctx context.Context, query string) ([]string, error) {
	var res []string

	dbTx := s.db.WithContext(ctx).Raw(query + " SETTINGS readonly=1").Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetBridgeEvent returns a bridge event.
func (s *Store) GetBridgeEvent(ctx context.Context, query string) (*BridgeEvent, error) {
	var res BridgeEvent

	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return &res, nil
}

// GetBridgeEvents returns bridge events.
func (s *Store) GetBridgeEvents(ctx context.Context, query string) ([]BridgeEvent, error) {
	var res []BridgeEvent
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetAllBridgeEvents returns bridge events.
func (s *Store) GetAllBridgeEvents(ctx context.Context, query string) ([]HybridBridgeEvent, error) {
	var res []HybridBridgeEvent
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	return res, nil
}

// GetTxCounts returns Tx counts.
func (s *Store) GetTxCounts(ctx context.Context, query string) ([]*model.TransactionCountResult, error) {
	var res []*model.TransactionCountResult
	dbTx := s.db.WithContext(ctx).Raw(query + " SETTINGS readonly=1").Find(&res)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetTokenCounts returns Tx counts.
func (s *Store) GetTokenCounts(ctx context.Context, query string) ([]*model.TokenCountResult, error) {
	var res []*model.TokenCountResult
	dbTx := s.db.WithContext(ctx).Raw(query + " SETTINGS readonly=1").Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetDateResults returns the dya by day data.
func (s *Store) GetDateResults(ctx context.Context, query string) ([]*model.DateResult, error) {
	var res []*model.DateResult
	dbTx := s.db.WithContext(ctx).Raw(query + " SETTINGS readonly=1").Scan(&res)

	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get date results: %w", dbTx.Error)
	}

	return res, nil
}

// GetAddressRanking gets AddressRanking for a given query.
func (s *Store) GetAddressRanking(ctx context.Context, query string) ([]*model.AddressRanking, error) {
	var res []*model.AddressRanking

	dbTx := s.db.WithContext(ctx).Raw(query + " SETTINGS readonly=1").Scan(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	if len(res) == 0 {
		return nil, nil
	}

	return res, nil
}

/*╔══════════════════════════════════════════════════════════════════════╗*\
▏*║                       Specific Read Functions                        ║*▕
\*╚══════════════════════════════════════════════════════════════════════╝*/

// PartialInfosFromIdentifiers returns events given identifiers. If order is true, the events are ordered by block number.
//
//nolint:cyclop
func (s *Store) PartialInfosFromIdentifiers(ctx context.Context, query string) ([]*model.PartialInfo, error) {
	var res []BridgeEvent
	var partialInfos []*model.PartialInfo

	dbTx := s.db.WithContext(ctx).Raw(query + " SETTINGS readonly=1").Find(&res)
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

// PartialInfosFromIdentifiersByChain returns events given identifiers. If order is true, the events are ordered by block number.
//
//nolint:cyclop
func (s *Store) PartialInfosFromIdentifiersByChain(ctx context.Context, query string) (map[int]*model.PartialInfo, error) {
	var res []BridgeEvent
	output := make(map[int]*model.PartialInfo)
	dbTx := s.db.WithContext(ctx).Raw(query + " SETTINGS readonly=1").Find(&res)
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

		output[chainIDInt] = &model.PartialInfo{
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

	return output, nil
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
