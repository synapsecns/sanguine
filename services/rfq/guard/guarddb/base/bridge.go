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

// StoreBridgeRequest stores a quote request.
func (s Store) StoreBridgeRequest(ctx context.Context, request guarddb.BridgeRequest) error {
	model := FromBridgeRequest(request)
	dbTx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: transactionIDFieldName}},
		DoUpdates: clause.AssignmentColumns([]string{transactionIDFieldName}),
	}).Create(&model)
	if dbTx.Error != nil {
		return fmt.Errorf("could not store request: %w", dbTx.Error)
	}
	return nil
}

// GetBridgeRequestByID gets a quote request by id. Should return ErrNoBridgeRequestForID if not found.
func (s Store) GetBridgeRequestByID(ctx context.Context, id [32]byte) (*guarddb.BridgeRequest, error) {
	var modelResult BridgeRequestModel
	tx := s.DB().WithContext(ctx).Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).First(&modelResult)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, guarddb.ErrNoBridgeRequestForID
	}

	if tx.Error != nil {
		return nil, fmt.Errorf("could not get request")
	}

	qr, err := modelResult.ToBridgeRequest()
	if err != nil {
		return nil, err
	}
	return qr, nil
}
