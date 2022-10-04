package sql

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/services/explorer/graphql/server/graph/model"
	"math/big"
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

// GetAllChainIDs gets all chain IDs that have been used in bridge events.
func (s *Store) GetAllChainIDs(ctx context.Context) ([]uint32, error) {
	var resOut []uint32
	dbTx := s.db.WithContext(ctx).Raw(`SELECT DISTINCT chain_id FROM bridge_events`).Find(&resOut)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	var resIn []big.Int
	dbTx = s.db.WithContext(ctx).Raw(`SELECT DISTINCT destination_chain_id FROM bridge_events`).Find(&resIn)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	uniqueChainIDs := make(map[uint32]bool)
	for _, chainID := range resOut {
		uniqueChainIDs[chainID] = true
	}
	for _, chainID := range resIn {
		uniqueChainIDs[uint32(chainID.Uint64())] = true
	}
	var res []uint32
	for chainID := range uniqueChainIDs {
		res = append(res, chainID)
	}

	return res, nil
}

// GetTokenAddressesByChainID gets all token addresses that have been used in bridge events for a given chain ID.
func (s *Store) GetTokenAddressesByChainID(ctx context.Context, chainID uint32) ([]string, error) {
	var res []string
	dbTx := s.db.WithContext(ctx).Raw(`SELECT DISTINCT %s FROM bridge_events WHERE %s = %d`, TokenFieldName, ChainIDFieldName, chainID).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	return res, nil
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
	fmt.Println("resresresresres", res)
	return uint64(res), nil
}
