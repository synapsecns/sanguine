package db_test

import (
	"time"

	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
)

func (d *DBSuite) TestGetQuotesByDestChainAndToken() {
	d.RunOnAllDBs(func(testDB db.ApiDB) {
		// Arrange: Create and insert a quote
		expectedQuote := &db.QuoteModel{
			// Initialize fields like ID, DestChainID, DestTokenAddr, etc.
			ID:            1,
			DestChainID:   2,
			DestTokenAddr: "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
			DestAmount:    decimal.NewFromInt(1000),
			Price:         decimal.NewFromFloat(0.01),
			UpdatedAt:     time.Now(),
		}
		err := testDB.UpsertQuote(expectedQuote)
		d.Require().NoError(err)

		// Act: Retrieve quotes by DestChainID and DestTokenAddr
		quotes, err := testDB.GetQuotesByDestChainAndToken(expectedQuote.DestChainID, expectedQuote.DestTokenAddr)
		d.Require().NoError(err)

		// Assert: Check if the retrieved quotes match the inserted quote
		d.Len(quotes, 1)
		d.Equal(expectedQuote.ID, quotes[0].ID)
		// Continue asserting other fields
	})
}

func (d *DBSuite) TestUpsertQuote() {
	d.RunOnAllDBs(func(testDB db.ApiDB) {
		// Arrange: Create a quote
		quote := &db.QuoteModel{
			// Initialize fields like ID, DestChainID, DestTokenAddr, etc.
			ID:            1,
			DestChainID:   2,
			DestTokenAddr: "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
			DestAmount:    decimal.NewFromInt(1000),
			Price:         decimal.NewFromFloat(0.01),
			UpdatedAt:     time.Now(),
		}

		// Act & Assert: Insert new quote
		err := testDB.UpsertQuote(quote)
		d.Require().NoError(err)

		// Retrieve to verify insertion
		retrievedQuotes, err := testDB.GetQuotesByDestChainAndToken(quote.DestChainID, quote.DestTokenAddr)
		d.Require().NoError(err)
		d.Len(retrievedQuotes, 1)
		// Assert other fields if necessary

		// Act & Assert: Update the existing quote
		quote.Price = decimal.NewFromFloat(0.02)
		err = testDB.UpsertQuote(quote)
		d.Require().NoError(err)

		// Retrieve to verify update
		updatedQuotes, err := testDB.GetQuotesByDestChainAndToken(quote.DestChainID, quote.DestTokenAddr)
		d.Require().NoError(err)
		d.Len(updatedQuotes, 1)
		d.Equal(quote.Price, updatedQuotes[0].Price)
		// Assert other fields if necessary
	})
}
