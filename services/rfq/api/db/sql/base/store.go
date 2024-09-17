package base

import (
	"context"
	"fmt"

	"gorm.io/gorm/clause"

	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// GetQuotesByDestChainAndToken gets quotes from the database by destination chain and token.
func (s *Store) GetQuotesByDestChainAndToken(ctx context.Context, destChainID uint64, destTokenAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote

	result := s.db.WithContext(ctx).Where("dest_chain_id = ? AND dest_token = ?", destChainID, destTokenAddr).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

// GetQuotesByOriginAndDestination gets quotes from the database by origin and destination.
func (s *Store) GetQuotesByOriginAndDestination(ctx context.Context, originChainID uint64, originTokenAddr string, destChainID uint64, destTokenAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote

	result := s.db.WithContext(ctx).Where("origin_chain_id = ? AND origin_token = ? AND dest_chain_id = ? AND dest_token = ?", originChainID, originTokenAddr, destChainID, destTokenAddr).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

// GetQuotesByRelayerAddress gets quotes from the database by relayer address.
func (s *Store) GetQuotesByRelayerAddress(ctx context.Context, relayerAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote

	result := s.db.WithContext(ctx).Where("relayer_address = ?", relayerAddr).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

// GetAllQuotes retrieves all quotes from the database.
func (s *Store) GetAllQuotes(ctx context.Context) ([]*db.Quote, error) {
	var quotes []*db.Quote
	result := s.db.WithContext(ctx).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

// UpsertQuote inserts a new quote into the database or updates an existing one.
func (s *Store) UpsertQuote(ctx context.Context, quote *db.Quote) error {
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(quote)

	if dbTx.Error != nil {
		return fmt.Errorf("could not update quote: %w", dbTx.Error)
	}
	return nil
}

// UpsertQuotes inserts multiple quotes into the database or updates existing ones.
func (s *Store) UpsertQuotes(ctx context.Context, quotes []*db.Quote) error {
	dbTx := s.DB().WithContext(ctx).
		Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(quotes)

	if dbTx.Error != nil {
		return fmt.Errorf("could not update quotes: %w", dbTx.Error)
	}
	return nil
}

// InsertActiveQuoteRequest inserts an active quote request into the database.
func (s *Store) InsertActiveQuoteRequest(ctx context.Context, req *model.PutUserQuoteRequest, requestID string) error {
	dbReq, err := db.FromUserRequest(req, requestID)
	if err != nil {
		return fmt.Errorf("could not convert user request to database request: %w", err)
	}
	result := s.db.WithContext(ctx).Create(dbReq)
	if result.Error != nil {
		return fmt.Errorf("could not insert active quote request: %w", result.Error)
	}
	return nil
}

// UpdateActiveQuoteRequestStatus updates the status of an active quote request in the database.
func (s *Store) UpdateActiveQuoteRequestStatus(ctx context.Context, requestID string, status db.ActiveQuoteRequestStatus) error {
	result := s.db.WithContext(ctx).
		Model(&db.ActiveQuoteRequest{}).
		Where("request_id = ?", requestID).
		Update("status", status)
	if result.Error != nil {
		return fmt.Errorf("could not update active quote request status: %w", result.Error)
	}
	return nil
}

// InsertActiveQuoteResponse inserts an active quote response into the database.
func (s *Store) InsertActiveQuoteResponse(ctx context.Context, resp *model.RelayerWsQuoteResponse, status db.ActiveQuoteResponseStatus) error {
	dbReq, err := db.FromRelayerResponse(resp, status)
	if err != nil {
		return fmt.Errorf("could not convert relayer response to database response: %w", err)
	}
	result := s.db.WithContext(ctx).Create(dbReq)
	if result.Error != nil {
		return fmt.Errorf("could not insert active quote response: %w", result.Error)
	}
	return nil
}

// UpdateActiveQuoteResponseStatus updates the status of an active quote response in the database.
func (s *Store) UpdateActiveQuoteResponseStatus(ctx context.Context, quoteID string, status db.ActiveQuoteResponseStatus) error {
	result := s.db.WithContext(ctx).
		Model(&db.ActiveQuoteResponse{}).
		Where("quote_id = ?", quoteID).
		Update("status", status)
	if result.Error != nil {
		return fmt.Errorf("could not update active quote response status: %w", result.Error)
	}
	return nil
}

// GetActiveQuoteRequests gets active quote requests from the database.
func (s *Store) GetActiveQuoteRequests(ctx context.Context, matchStatuses ...db.ActiveQuoteRequestStatus) ([]*db.ActiveQuoteRequest, error) {
	var requests []*db.ActiveQuoteRequest

	query := s.db.WithContext(ctx).Model(&db.ActiveQuoteRequest{})
	if len(matchStatuses) > 0 {
		query = query.Where("status IN ?", matchStatuses)
	}
	result := query.Find(&requests)
	if result.Error != nil {
		return nil, result.Error
	}
	return requests, nil
}
