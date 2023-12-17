package base

import (
	"context"
	"fmt"
	"gorm.io/gorm/clause"

	"github.com/synapsecns/sanguine/services/rfq/api/db"
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
