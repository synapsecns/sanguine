package base

import (
	"context"
	"errors"
	"fmt"

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

// UpdateRebalance updates the rebalance status.
//
//nolint:cyclop
func (s Store) UpdateRebalance(ctx context.Context, rebalance reldb.Rebalance, updateID bool) error {
	tx := s.DB().WithContext(ctx).Begin()
	if tx.Error != nil {
		return fmt.Errorf("could not start transaction: %w", tx.Error)
	}
	rebalanceModel := FromRebalance(rebalance)

	// prepare the updates
	updates := map[string]interface{}{
		statusFieldName: rebalance.Status,
	}
	if rebalanceModel.OriginTxHash.Valid {
		updates[originTxHashFieldName] = rebalanceModel.OriginTxHash
	}
	if rebalanceModel.DestTxHash.Valid {
		updates[destTxHashFieldName] = rebalanceModel.DestTxHash
	}

	// prepare the update transaction
	var result *gorm.DB
	if updateID {
		if rebalance.Origin == 0 {
			tx.Rollback()
			return fmt.Errorf("origin chain id is required if id is not set")
		}
		updates[rebalanceIDFieldName] = rebalanceModel.RebalanceID
		result = tx.Model(&Rebalance{}).
			Where(fmt.Sprintf("%s = ?", "origin"), rebalance.Origin).
			Where(fmt.Sprintf("%s = ?", statusFieldName), reldb.RebalanceInitiated.Int()).
			Updates(updates)
	} else {
		result = tx.Model(&Rebalance{}).
			Where(fmt.Sprintf("%s = ?", rebalanceIDFieldName), rebalanceModel.RebalanceID).
			Updates(updates)
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

// GetRebalanceByID gets a rebalance by id. Should return ErrNoRebalanceForID if not found.
func (s Store) GetRebalanceByID(ctx context.Context, rebalanceID string) (*reldb.Rebalance, error) {
	var modelResult Rebalance
	tx := s.DB().WithContext(ctx).Where(fmt.Sprintf("%s = ?", rebalanceIDFieldName), rebalanceID).First(&modelResult)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, reldb.ErrNoRebalanceForID
	}

	if tx.Error != nil {
		return nil, fmt.Errorf("could not get quote")
	}

	rebalance, err := modelResult.ToRebalance()
	if err != nil {
		return nil, fmt.Errorf("could not get rebalance: %w", err)
	}

	return rebalance, nil
}
