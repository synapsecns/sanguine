package rest

import (
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// FilterQuoteAge exports filterQuoteAge for testing.
func FilterQuoteAge(cfg config.Config, dbQuotes []*db.Quote) []*db.Quote {
	return filterQuoteAge(cfg, dbQuotes)
}

// GetPassiveQuote exports getPassiveQuote for testing.
func GetPassiveQuote(cfg config.Config, quotes []*db.Quote, request *model.PutRFQRequest) (*model.QuoteData, error) {
	return getPassiveQuote(cfg, quotes, request)
}
