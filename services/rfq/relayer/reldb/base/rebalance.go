package base

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"gorm.io/gorm"
)

// StoreRebalance stores a rebalance action.
func (s Store) StoreRebalance(ctx context.Context, rebalance reldb.Rebalance) error {
	reb := FromRebalance(rebalance)
	dbTx := s.DB().WithContext(ctx).Create(&reb)
	if dbTx.Error != nil {
		return fmt.Errorf("could not store rebalance: %w", dbTx.Error)
	}
	return nil
}

// UpdateRebalanceStatus updates the rebalance status.
func (s Store) UpdateRebalanceStatus(ctx context.Context, id [32]byte, origin *uint64, status reldb.RebalanceStatus) error {
	tx := s.DB().WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("could not start transaction: %w", tx.Error)
	}

	// prepare the update transaction
	var result *gorm.DB
	if origin != nil {
		result = tx.Model(&Rebalance{}).
			Where(fmt.Sprintf("%s = ?", "origin"), *origin).
			Where(fmt.Sprintf("%s = ?", statusFieldName), reldb.RebalanceInitiated.Int()).
			Updates(map[string]interface{}{
				rebalanceIDFieldName: hexutil.Encode(id[:]),
				statusFieldName:      status,
			})
	} else {
		result = tx.Model(&Rebalance{}).
			Where(fmt.Sprintf("%s = ?", rebalanceIDFieldName), hexutil.Encode(id[:])).
			Update(statusFieldName, status)
	}

	// commit the transaction if only one row is affected
	if result.Error != nil {
		tx.Rollback()
		return fmt.Errorf("could not update rebalance status: %w", result.Error)
	}
	if result.RowsAffected != 1 {
		tx.Rollback()
		return fmt.Errorf("expected 1 row to be affected, got %d", result.RowsAffected)
	}
	err := tx.Commit().Error
	if err != nil {
		return fmt.Errorf("could not commit transaction: %w", err)
	}
	return nil
}

// HasPendingRebalance checks if there is a pending rebalance for the given chain ids.
func (s Store) HasPendingRebalance(ctx context.Context, chainIDs ...uint64) (bool, error) {
	var rebalances []Rebalance

	matchStatuses := []reldb.RebalanceStatus{reldb.RebalanceInitiated, reldb.RebalancePending}
	inArgs := make([]int, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
	}

	// TODO: can be made more efficient by doing below check inside sql query
	tx := s.DB().WithContext(ctx).Model(&Rebalance{}).Where(fmt.Sprintf("%s IN ?", statusFieldName), inArgs).Find(&rebalances)
	if tx.Error != nil {
		return false, fmt.Errorf("could not get db results: %w", tx.Error)
	}

	// Check if any pending rebalances involve the given chain ids
	for _, result := range rebalances {
		for _, chainID := range chainIDs {
			if result.Origin == chainID || result.Destination == chainID {
				return true, nil
			}
		}
	}
	return false, nil
}
