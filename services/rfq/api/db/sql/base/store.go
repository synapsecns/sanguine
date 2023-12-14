package base

import (
	"errors"

	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"gorm.io/gorm"
)

func (s *Store) GetQuotesByDestChainAndToken(destChainId uint64, destTokenAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote

	result := s.db.Where("dest_chain_id = ? AND dest_token = ?", destChainId, destTokenAddr).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

func (s *Store) GetQuotesByOriginAndDestination(originChainId uint64, originTokenAddr string, destChainId uint64, destTokenAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote

	result := s.db.Where("origin_chain_id = ? AND origin_token = ? AND dest_chain_id = ? AND dest_token = ?", originChainId, originTokenAddr, destChainId, destTokenAddr).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

func (s *Store) GetQuotesByRelayerAddress(relayerAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote

	result := s.db.Where("relayer_address = ?", relayerAddr).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

// This function retrieves all quotes from the database.
func (s *Store) GetAllQuotes() ([]*db.Quote, error) {
	var quotes []*db.Quote
	result := s.db.Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

// UpsertQuote inserts a new quote into the database or updates an existing one.
func (s *Store) UpsertQuote(quote *db.Quote) error {
	var existingQuote db.Quote
	result := s.db.First(&existingQuote, quote.ID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Create new record if not found
		return s.db.Create(quote).Error
	} else if result.Error != nil {
		return result.Error
	}

	// Update existing record
	return s.db.Model(&existingQuote).Updates(quote).Error
}
