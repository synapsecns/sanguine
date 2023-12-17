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

// StoreQuoteRequest stores a quote request.
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

// GetQuoteRequestByID gets a quote request by id. Should return ErrNoQuoteForID if not found.
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

// GetQuoteResultsByStatus gets quote results by status.
func (s Store) GetQuoteResultsByStatus(ctx context.Context, matchStatuses ...reldb.QuoteRequestStatus) (res []reldb.QuoteRequest, _ error) {
	var quoteResults []RequestForQuote

	inArgs := make([]int, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
	}

	// TODO: consider pagination
	tx := s.DB().WithContext(ctx).Model(&RequestForQuote{}).Where(fmt.Sprintf("%s IN ?", statusFieldName), inArgs).Find(&quoteResults)
	if tx.Error != nil {
		return []reldb.QuoteRequest{}, fmt.Errorf("could not get db results: %w", tx.Error)
	}

	for _, result := range quoteResults {
		marshaled, err := result.ToQuoteRequest()
		if err != nil {
			return []reldb.QuoteRequest{}, fmt.Errorf("could not get quotes")
		}
		res = append(res, *marshaled)
	}
	return res, nil
}

// UpdateQuoteRequestStatus todo: db test.
func (s Store) UpdateQuoteRequestStatus(ctx context.Context, id [32]byte, status reldb.QuoteRequestStatus) error {
	tx := s.DB().WithContext(ctx).Model(&RequestForQuote{}).
		Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).
		Update(statusFieldName, status)
	if tx.Error != nil {
		return fmt.Errorf("could not update: %w", tx.Error)
	}
	return nil
}
