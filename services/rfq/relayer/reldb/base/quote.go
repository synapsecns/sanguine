package base

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// StoreQuoteRequest stores a quote request
func (s Store) StoreQuoteRequest(ctx context.Context, request reldb.QuoteRequest) error {
	rq := FromQuoteRequest(request)

	dbTx := s.DB().WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: transactionIDFieldName}},
		DoUpdates: clause.AssignmentColumns([]string{transactionIDFieldName}),
	}).Create(&rq)
	if dbTx.Error != nil {
		return fmt.Errorf("could not store quote: %w", dbTx.Error)
	}
	return nil
}

func (s Store) GetQuoteRequestByID(ctx context.Context, id [32]byte) (*reldb.QuoteRequest, error) {
	var modelResult RequestForQuote
	tx := s.DB().WithContext(ctx).Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).First(&modelResult)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, reldb.ErrNoQuoteForID
	}

	if tx.Error != nil {
		return nil, fmt.Errorf("could not get quote id")
	}

	qr, err := modelResult.ToQuoteRequest()
	if err != nil {
		return nil, err
	}
	return qr, nil
}
