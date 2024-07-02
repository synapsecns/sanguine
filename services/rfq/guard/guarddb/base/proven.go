package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/services/rfq/guard/guarddb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StorePendingProven stores a quote request.
func (s Store) StorePendingProven(ctx context.Context, proven guarddb.PendingProven) error {
	model := FromPendingProven(proven)
	dbTx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: transactionIDFieldName}},
		DoUpdates: clause.AssignmentColumns([]string{transactionIDFieldName}),
	}).Create(&model)
	if dbTx.Error != nil {
		return fmt.Errorf("could not store proven: %w", dbTx.Error)
	}
	return nil
}

// GetPendingProvenByID gets a quote request by id. Should return ErrNoProvenForID if not found.
func (s Store) GetPendingProvenByID(ctx context.Context, id [32]byte) (*guarddb.PendingProven, error) {
	var modelResult PendingProvenModel
	tx := s.DB().WithContext(ctx).Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).First(&modelResult)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, guarddb.ErrNoProvenForID
	}

	if tx.Error != nil {
		return nil, fmt.Errorf("could not get proven")
	}

	qr, err := modelResult.ToPendingProven()
	if err != nil {
		return nil, err
	}
	return qr, nil
}

// UpdatePendingProvenStatus updates the status of a pending proven.
func (s Store) UpdatePendingProvenStatus(ctx context.Context, id [32]byte, status guarddb.PendingProvenStatus) error {
	tx := s.DB().WithContext(ctx).Model(&PendingProvenModel{}).
		Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).
		Update(statusFieldName, status)
	if tx.Error != nil {
		return fmt.Errorf("could not update: %w", tx.Error)
	}
	return nil
}
