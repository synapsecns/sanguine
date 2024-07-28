package base

import (
	"context"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
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
		return nil, fmt.Errorf("could not get quote")
	}

	qr, err := modelResult.ToQuoteRequest()
	if err != nil {
		return nil, err
	}
	return qr, nil
}

// GetQuoteRequestByOriginTxHash gets a quote request by tx hash. Should return ErrNoQuoteForID if not found.
func (s Store) GetQuoteRequestByOriginTxHash(ctx context.Context, txHash common.Hash) (*reldb.QuoteRequest, error) {
	var modelResult RequestForQuote
	tx := s.DB().WithContext(ctx).Where(fmt.Sprintf("%s = ?", originTxHashFieldName), txHash.String()).First(&modelResult)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, reldb.ErrNoQuoteForTxHash
	}

	if tx.Error != nil {
		return nil, fmt.Errorf("could not get quote %w", tx.Error)
	}

	qr, err := modelResult.ToQuoteRequest()
	if err != nil {
		return nil, fmt.Errorf("could not convert to quote request %w", err)
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
func (s Store) UpdateQuoteRequestStatus(ctx context.Context, id [32]byte, status reldb.QuoteRequestStatus, prevStatus *reldb.QuoteRequestStatus) error {
	if prevStatus == nil {
		req, err := s.GetQuoteRequestByID(ctx, id)
		if err != nil {
			return fmt.Errorf("could not get quote: %w", reldb.ErrNoQuoteForID)
		}
		prevStatus = &req.Status
	}
	if !isValidStateTransition(*prevStatus, status) {
		return nil
	}

	tx := s.DB().WithContext(ctx).Model(&RequestForQuote{}).
		Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).
		Update(statusFieldName, status)
	if tx.Error != nil {
		return fmt.Errorf("could not update: %w", tx.Error)
	}
	return nil
}

// UpdateDestTxHash todo: db test.
func (s Store) UpdateDestTxHash(ctx context.Context, id [32]byte, destTxHash common.Hash) error {
	tx := s.DB().WithContext(ctx).Model(&RequestForQuote{}).
		Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).
		Update(destTxHashFieldName, destTxHash.String())
	if tx.Error != nil {
		return fmt.Errorf("could not update: %w", tx.Error)
	}
	return nil
}

// UpdateRelayNonce todo: db test.
func (s Store) UpdateRelayNonce(ctx context.Context, id [32]byte, nonce uint64) error {
	tx := s.DB().WithContext(ctx).Model(&RequestForQuote{}).
		Where(fmt.Sprintf("%s = ?", transactionIDFieldName), hexutil.Encode(id[:])).
		Update(relayNonceFieldName, nonce)
	if tx.Error != nil {
		return fmt.Errorf("could not update: %w", tx.Error)
	}
	return nil
}

func isValidStateTransition(prevStatus, status reldb.QuoteRequestStatus) bool {
	if status == reldb.DeadlineExceeded || status == reldb.WillNotProcess {
		return true
	}
	return status >= prevStatus
}

type statusCount struct {
	Status int
	Count  int64
}

// GetStatusCounts gets the counts of quote requests by status.
func (s Store) GetStatusCounts(ctx context.Context, matchStatuses ...reldb.QuoteRequestStatus) (map[reldb.QuoteRequestStatus]int, error) {
	inArgs := make([]int, len(matchStatuses))
	for i := range matchStatuses {
		inArgs[i] = int(matchStatuses[i].Int())
	}

	var results []statusCount
	tx := s.DB().
		WithContext(ctx).
		Model(&RequestForQuote{}).
		Select(fmt.Sprintf("%s, COUNT(*) as count", statusFieldName)).
		Where(fmt.Sprintf("%s IN ?", statusFieldName), inArgs).
		Group(statusFieldName).
		Scan(&results)
	if tx.Error != nil {
		return nil, fmt.Errorf("could not get db results: %w", tx.Error)
	}

	statuses := make(map[reldb.QuoteRequestStatus]int)
	for _, result := range results {
		statuses[reldb.QuoteRequestStatus(result.Status)] = int(result.Count)
	}

	return statuses, nil
}
