package base

import (
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"gorm.io/gorm"
)

func (s *Store) GetQuotesByDestChainAndToken(destChainId uint64, destTokenAddr string) ([]*db.Quote, error) {
	var quotes []*db.Quote
	result := s.db.Where("dest_chain_id = ? AND token = ?", destChainId, destTokenAddr).Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

func (s *Store) GetAllQuotes() ([]*db.Quote, error) {
	var quotes []*db.Quote
	result := s.db.Find(&quotes)
	if result.Error != nil {
		return nil, result.Error
	}
	return quotes, nil
}

func (s *Store) UpsertQuote(quote *db.Quote) error {
	var existingQuote db.Quote
	result := s.db.First(&existingQuote, quote.ID)

	if result.Error == gorm.ErrRecordNotFound {
		// Create new record if not found
		return s.db.Create(quote).Error
	} else if result.Error != nil {
		return result.Error
	}

	// Update existing record
	return s.db.Model(&existingQuote).Updates(quote).Error
}
