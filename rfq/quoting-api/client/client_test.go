package client_test

import (
	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/rfq/quoting-api/models"
)

func (c *ClientSuite) TestCreateQuote() {
	originAmount, _ := decimal.NewFromString("100")
	destAmount, _ := decimal.NewFromString("200")

	// Created a dummy quote
	quote := models.Quote{
		Relayer:       c.testHandler.TestWallet().Address().String(),
		OriginChainID: 1,
		DestChainID:   chainID,
		OriginToken:   "0x1",
		DestToken:     "0x2",
		OriginAmount:  originAmount,
		DestAmount:    destAmount,
	}

	err := c.client.CreateQuote(quote)
	c.NoError(err)
	// TODO: Make sure the quote is there
}
