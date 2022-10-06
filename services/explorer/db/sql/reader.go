package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
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
		return nil, nil
	}
	if res.SwapSuccess.Uint64() == 1 {
		trueVal := true
		return &trueVal, nil
	} else if res.SwapSuccess.Uint64() == 0 {
		falseVal := false
		return &falseVal, nil
	}
	return nil, nil
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

// PartialInfosFromIdentifiers returns events given identifiers. If order is true, the events are ordered by block number.
func (s *Store) PartialInfosFromIdentifiers(ctx context.Context, chainID *uint32, address, tokenAddress, kappa, txHash *string, page int, order bool) (partialInfos []*model.PartialInfo, err error) {
	var res []BridgeEvent
	and := ""
	var chainIDSpecifier string
	if chainID != nil {
		chainIDSpecifier = fmt.Sprintf(" %s %s = %d", and, ChainIDFieldName, *chainID)
		and = "AND"
	}
	var addressSpecifier string
	if address != nil {
		addressSpecifier = fmt.Sprintf(" %s (%s = '%s' OR %s = '%s')", and, RecipientFieldName, *address, SenderFieldName, *address)
		and = "AND"
	}
	var tokenAddressSpecifier string
	if tokenAddress != nil {
		tokenAddressSpecifier = fmt.Sprintf(" %s %s = '%s'", and, TokenFieldName, *tokenAddress)
		and = "AND"
	}
	if page < 1 {
		page = 1
	}
	//pageSpecifier := fmt.Sprintf(" LIMIT %d OFFSET %d", PageSize, (page-1)*PageSize)
	pageSpecifier := ""

	compositeIdentifiers := fmt.Sprintf(
		`WHERE%s%s%s%s`,
		chainIDSpecifier, addressSpecifier, tokenAddressSpecifier, pageSpecifier,
	)
	if kappa != nil {
		compositeIdentifiers += fmt.Sprintf(" %s %s = '%s'", and, KappaFieldName, *kappa)
	}
	if txHash != nil {
		compositeIdentifiers += fmt.Sprintf(" %s %s = '%s'", and, TxHashFieldName, *txHash)
	}
	if order {
		compositeIdentifiers += fmt.Sprintf(" ORDER BY %s DESC", BlockNumberFieldName)
	}

	fmt.Println(fmt.Sprintf(
		`SELECT * FROM bridge_events %s`,
		compositeIdentifiers))
	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
		`SELECT * FROM bridge_events %s`,
		compositeIdentifiers,
	)).Find(&res)
	fmt.Printf("res: %+v\n", res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	//if kappa != nil {
	//	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
	//		`SELECT * FROM bridge_events WHERE %s = %s AND %s`,
	//		KappaFieldName, *kappa, compositeIdentifiers,
	//	)).Find(&res)
	//	if dbTx.Error != nil {
	//		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	//	}
	//} else if txHash != nil {
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
	//} else {
	//	dbTx := s.db.WithContext(ctx).Raw(fmt.Sprintf(
	//		`SELECT * FROM bridge_events WHERE %s`,
	//		compositeIdentifiers,
	//	)).Find(&res)
	//	if dbTx.Error != nil {
	//		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	//	}
	//}

	for _, event := range res {
		chainIDInt := int(event.ChainID)
		blockNumberInt := int(event.BlockNumber)
		txHashh := event.TxHash
		var recipient string
		if event.Recipient.Valid {
			recipient = event.Recipient.String
		} else if event.RecipientBytes.Valid {
			recipient = event.RecipientBytes.String
		} else {
			recipient = ""
		}
		partialInfos = append(partialInfos, &model.PartialInfo{
			ChainID:      &chainIDInt,
			Address:      &recipient,
			TxnHash:      &txHashh,
			TokenAddress: &event.Token,
			BlockNumber:  &blockNumberInt,
		})
	}

	return partialInfos, nil
}
