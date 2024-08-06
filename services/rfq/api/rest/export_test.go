package rest

import (
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
)

// FilterQuoteAge exports filterQuoteAge for testing.
func FilterQuoteAge(cfg config.Config, dbQuotes []*db.Quote) []*db.Quote {
	return filterQuoteAge(cfg, dbQuotes)
}
