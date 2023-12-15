package db_test

import (
	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
)

func (d *DBSuite) TestGetQuotesByDestChainAndToken() {
	d.RunOnAllDBs(func(testDB db.ApiDB) {
		// Arrange: Create and insert a quote
		expectedQuote := &db.Quote{
			// Initialize fields like ID, DestChainID, DestTokenAddr, etc.
			OriginChainID:   1,
			OriginTokenAddr: "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
			DestChainID:     42161,
			DestTokenAddr:   "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
			DestAmount:      decimal.NewFromInt(1000),
			Price:           decimal.NewFromFloat(0.01),
			MaxOriginAmount: decimal.NewFromInt(1000).Div(decimal.NewFromFloat(0.01)),
		}
		err := testDB.UpsertQuote(d.GetTestContext(), expectedQuote)
		d.Require().NoError(err)

		// Act: Retrieve quotes by DestChainID and DestTokenAddr
		quotes, err := testDB.GetQuotesByDestChainAndToken(d.GetTestContext(), expectedQuote.DestChainID, expectedQuote.DestTokenAddr)
		d.Require().NoError(err)

		// Assert: Check if the retrieved quotes match the inserted quote
		d.Len(quotes, 1)
		d.Equal(expectedQuote.OriginChainID, quotes[0].OriginChainID)
		d.Equal(expectedQuote.OriginTokenAddr, quotes[0].OriginTokenAddr)
		d.Equal(expectedQuote.DestChainID, quotes[0].DestChainID)
		d.Equal(expectedQuote.DestTokenAddr, quotes[0].DestTokenAddr)
		// TODO: decimal comparison
		// d.Equal(expectedQuote.DestAmount, quotes[0].DestAmount)
		// d.Equal(expectedQuote.Price, quotes[0].Price)
		// d.Equal(expectedQuote.MaxOriginAmount, quotes[0].MaxOriginAmount)
		d.NotEqual(quotes[0].UpdatedAt, nil)
		// Continue asserting other fields
	})
}

func (d *DBSuite) TestUpsertQuote() {
	d.RunOnAllDBs(func(testDB db.ApiDB) {
		// Arrange: Create a quote
		quote := &db.Quote{
			// Initialize fields like ID, DestChainID, DestTokenAddr, etc.
			OriginChainID:   1,
			OriginTokenAddr: "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
			DestChainID:     42161,
			DestTokenAddr:   "0x3f5CE5FBFe3E9af3971dD833D26bA9b5C936f0bE",
			DestAmount:      decimal.NewFromInt(1000),
			Price:           decimal.NewFromFloat(0.01),
			MaxOriginAmount: decimal.NewFromInt(1000).Div(decimal.NewFromFloat(0.01)),
		}

		// Act & Assert: Insert new quote
		err := testDB.UpsertQuote(d.GetTestContext(), quote)
		d.Require().NoError(err)

		// Retrieve to verify insertion
		retrievedQuotes, err := testDB.GetQuotesByDestChainAndToken(d.GetTestContext(), quote.DestChainID, quote.DestTokenAddr)
		d.Require().NoError(err)
		d.Len(retrievedQuotes, 1)
		// Assert other fields if necessary

		// Act & Assert: Update the existing quote
		quote.Price = decimal.NewFromFloat(0.02)
		err = testDB.UpsertQuote(d.GetTestContext(), quote)
		d.Require().NoError(err)

		// Retrieve to verify update
		updatedQuotes, err := testDB.GetQuotesByDestChainAndToken(d.GetTestContext(), quote.DestChainID, quote.DestTokenAddr)
		d.Require().NoError(err)
		d.Len(updatedQuotes, 1)
		d.Equal(quote.Price, updatedQuotes[0].Price)
		// Assert other fields if necessary
	})
}
