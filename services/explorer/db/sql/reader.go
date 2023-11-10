package sql

import (
	"context"
	"fmt"
	"github.com/benbjohnson/immutable"
	"github.com/synapsecns/sanguine/core"
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

// GetString gets a string from a given query.
func (s *Store) GetString(ctx context.Context, query string) (string, error) {
	var res string
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return "", fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
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

// GetMVBridgeEvent gets a bridge event from the materialized view table.
func (s *Store) GetMVBridgeEvent(ctx context.Context, query string) (*HybridBridgeEvent, error) {
	var res HybridBridgeEvent

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

// GetDailyTotals returns bridge events.
func (s *Store) GetDailyTotals(ctx context.Context, query string) ([]*model.DateResultByChain, error) {
	var res []*model.DateResultByChain

	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}
	return res, nil
}

// GetAllMessageBusEvents returns message bus events.
func (s *Store) GetAllMessageBusEvents(ctx context.Context, query string) ([]HybridMessageBusEvent, error) {
	var res []HybridMessageBusEvent
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read message bus event: %w", dbTx.Error)
	}
	return res, nil
}

// GetTxCounts returns Tx counts.
func (s *Store) GetTxCounts(ctx context.Context, query string) ([]*model.TransactionCountResult, error) {
	var res []*model.TransactionCountResult
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetTokenCounts returns Tx counts.
func (s *Store) GetTokenCounts(ctx context.Context, query string) ([]*model.TokenCountResult, error) {
	var res []*model.TokenCountResult
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to read bridge event: %w", dbTx.Error)
	}

	return res, nil
}

// GetRankedChainsByVolume gets ranked chains by volume.
func (s *Store) GetRankedChainsByVolume(ctx context.Context, query string) ([]*model.VolumeByChainID, error) {
	var res []*model.VolumeByChainID
	dbTx := s.db.WithContext(ctx).Raw(query).Find(&res)
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

// GetAddressData returns the address data.
func (s *Store) GetAddressData(ctx context.Context, query string) (float64, float64, int, error) {
	type addressData struct {
		VolumeTotal float64 `gorm:"column:volumeTotal"`
		FeeTotal    float64 `gorm:"column:feeTotal"`
		TxTotal     int     `gorm:"column:txTotal"`
	}
	var res addressData
	// var test map[string]interface{}
	dbTx := s.db.WithContext(ctx).Raw(query).Scan(&res)
	if dbTx.Error != nil {
		return 0, 0, 0, fmt.Errorf("failed to get address data: %w", dbTx.Error)
	}

	return res.VolumeTotal, res.FeeTotal, res.TxTotal, nil
}

// GetAddressChainRanking ranks chains by volume for a given address.
func (s *Store) GetAddressChainRanking(ctx context.Context, query string) ([]*model.AddressChainRanking, error) {
	var res []*model.AddressChainRanking
	dbTx := s.db.WithContext(ctx).Raw(query).Scan(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get address chain ranking: %w", dbTx.Error)
	}

	return res, nil
}

// GetAddressDailyData gets daily data (number of txs_ for a given address.
func (s *Store) GetAddressDailyData(ctx context.Context, query string) ([]*model.AddressDailyCount, error) {
	var res []*model.AddressDailyCount
	dbTx := s.db.WithContext(ctx).Raw(query).Scan(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get address daily data: %w", dbTx.Error)
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

// GetLastStoredBlock returns the last stored block for a given chainID and contract.
func (s *Store) GetLastStoredBlock(ctx context.Context, chainID uint32, contract string) (uint64, error) {
	query := fmt.Sprintf("SELECT %s FROM last_blocks WHERE %s = %d AND %s = '%s' ORDER BY %s DESC LIMIT 1", BlockNumberFieldName, ChainIDFieldName, chainID, ContractAddressFieldName, contract, BlockNumberFieldName)
	lastBlock, err := s.GetUint64(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed to get last block: %w", err)
	}

	return lastBlock, nil
}

// GetLeaderboard gets the bridge leaderboard.
func (s *Store) GetLeaderboard(ctx context.Context, query string) ([]*model.Leaderboard, error) {
	var res []*model.Leaderboard
	dbTx := s.db.WithContext(ctx).Raw(query).Scan(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get leaderboard: %w", dbTx.Error)
	}
	if len(res) == 0 {
		return nil, nil
	}

	return res, nil
}

// GetPendingByChain gets the bridge leaderboard by chain.
// returns chainid, count
// TODO: test this.
func (s *Store) GetPendingByChain(ctx context.Context) (res *immutable.Map[int, int], err error) {
	const query = `SELECT
		toInt64(destination_chain_id) as destination_chain_id,
		toInt64(COUNTDistinct(destination_kappa)) AS distinct_count
	FROM bridge_events
	WHERE destination_kappa NOT IN (
		SELECT kappa
		FROM bridge_events
		WHERE kappa != ''
	)
	GROUP BY destination_chain_id`

	type PendingByChain struct {
		DestinationChainID *int64 `gorm:"column:destination_chain_id"`
		DistinctCount      *int64 `gorm:"column:distinct_count"`
	}

	var pending []*PendingByChain

	dbTx := s.db.WithContext(ctx).Raw(query).Scan(&pending)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get pending by chain: %w", dbTx.Error)
	}

	builder := immutable.NewMapBuilder[int, int](nil)

	for _, kvPair := range pending {
		builder.Set(int(*kvPair.DestinationChainID), int(*kvPair.DistinctCount))
	}

	return builder.Map(), nil
}

func (s *Store) GetBlockHeights(ctx context.Context, query string, contractTypeMap map[string]model.ContractType) ([]*model.BlockHeight, error) {
	var res []*LastBlock
	dbTx := s.db.WithContext(ctx).Raw(query).Scan(&res)
	if dbTx.Error != nil {
		return nil, fmt.Errorf("failed to get block heights: %w", dbTx.Error)
	}
	if len(res) == 0 {
		return nil, nil
	}

	var formatted []*model.BlockHeight
	for _, block := range res {
		chainID := int(block.ChainID)
		blockNumber := int(block.BlockNumber)
		formatted = append(formatted, &model.BlockHeight{
			ChainID:     &chainID,
			Type:        core.PtrTo(contractTypeMap[block.ContractAddress]),
			BlockNumber: &blockNumber,
		})
	}

	return formatted, nil
}
