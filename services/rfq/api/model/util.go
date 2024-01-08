package model

import (
	"time"

	"github.com/synapsecns/sanguine/services/rfq/api/db"
)

// QuoteResponseFromDbQuote converts a db.Quote to a GetQuoteResponse.
func QuoteResponseFromDbQuote(dbQuote *db.Quote) *GetQuoteResponse {
	return &GetQuoteResponse{
		OriginChainID:           int(dbQuote.OriginChainID),
		OriginTokenAddr:         dbQuote.OriginTokenAddr,
		DestChainID:             int(dbQuote.DestChainID),
		DestTokenAddr:           dbQuote.DestTokenAddr,
		DestAmount:              dbQuote.DestAmount.String(),
		MaxOriginAmount:         dbQuote.MaxOriginAmount.String(),
		FixedFee:                dbQuote.FixedFee.String(),
		RelayerAddr:             dbQuote.RelayerAddr,
		OriginFastBridgeAddress: dbQuote.OriginFastBridgeAddress,
		DestFastBridgeAddress:   dbQuote.DestFastBridgeAddress,
		UpdatedAt:               dbQuote.UpdatedAt.Format(time.RFC3339),
	}
}
