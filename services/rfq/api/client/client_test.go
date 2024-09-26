package client_test

import (
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// TODO: @aurelius tese tests make a lot less sesnes w/ a composite index

func (c *ClientSuite) TestPutAndGetQuote() {
	req := model.PutQuoteRequest{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100",
		MaxOriginAmount: "200",
		FixedFee:        "10",
	}

	err := c.client.PutQuote(c.GetTestContext(), &req)
	c.Require().NoError(err)

	quotes, err := c.client.GetAllQuotes(c.GetTestContext())
	c.Require().NoError(err)

	expectedResp := model.GetQuoteResponse{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100",
		MaxOriginAmount: "200",
		FixedFee:        "10",
		RelayerAddr:     c.testWallet.Address().String(),
		UpdatedAt:       quotes[0].UpdatedAt,
	}
	c.Len(quotes, 1)
	c.Equal(expectedResp, *quotes[0])
}

func (c *ClientSuite) TestPutAndGetBulkQuotes() {
	req := model.PutBulkQuotesRequest{
		Quotes: []model.PutQuoteRequest{
			{
				OriginChainID:   1,
				OriginTokenAddr: "0xOriginTokenAddr",
				DestChainID:     42161,
				DestTokenAddr:   "0xDestTokenAddr",
				DestAmount:      "100",
				MaxOriginAmount: "200",
				FixedFee:        "10",
			},
			{
				OriginChainID:   42161,
				OriginTokenAddr: "0xOriginTokenAddr",
				DestChainID:     1,
				DestTokenAddr:   "0xDestTokenAddr",
				DestAmount:      "100",
				MaxOriginAmount: "200",
				FixedFee:        "10",
			},
		},
	}

	err := c.client.PutBulkQuotes(c.GetTestContext(), &req)
	c.Require().NoError(err)

	quotes, err := c.client.GetAllQuotes(c.GetTestContext())
	c.Require().NoError(err)

	expectedResp := []model.GetQuoteResponse{
		{
			OriginChainID:   1,
			OriginTokenAddr: "0xOriginTokenAddr",
			DestChainID:     42161,
			DestTokenAddr:   "0xDestTokenAddr",
			DestAmount:      "100",
			MaxOriginAmount: "200",
			FixedFee:        "10",
			RelayerAddr:     c.testWallet.Address().String(),
			UpdatedAt:       quotes[0].UpdatedAt,
		},
		{
			OriginChainID:   42161,
			OriginTokenAddr: "0xOriginTokenAddr",
			DestChainID:     1,
			DestTokenAddr:   "0xDestTokenAddr",
			DestAmount:      "100",
			MaxOriginAmount: "200",
			FixedFee:        "10",
			RelayerAddr:     c.testWallet.Address().String(),
			UpdatedAt:       quotes[0].UpdatedAt,
		},
	}
	c.Len(quotes, 2)
	c.Equal(expectedResp[0], *quotes[0])
	c.Equal(expectedResp[1], *quotes[1])
}

func (c *ClientSuite) TestGetSpecificQuote() {
	req := model.PutQuoteRequest{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100",
		MaxOriginAmount: "200",
		FixedFee:        "10",
	}

	err := c.client.PutQuote(c.GetTestContext(), &req)
	c.Require().NoError(err)

	quotes, err := c.client.GetSpecificQuote(c.GetTestContext(), &model.GetQuoteSpecificRequest{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
	})
	c.Require().NoError(err)

	c.Len(quotes, 1)
	expectedResp := model.GetQuoteResponse{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100",
		MaxOriginAmount: "200",
		FixedFee:        "10",
		RelayerAddr:     c.testWallet.Address().String(),
		UpdatedAt:       quotes[0].UpdatedAt,
	}
	c.Equal(expectedResp, *quotes[0])
}

func (c *ClientSuite) TestGetQuoteByRelayerAddress() {
	req := model.PutQuoteRequest{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100",
		MaxOriginAmount: "200",
		FixedFee:        "10",
	}

	err := c.client.PutQuote(c.GetTestContext(), &req)
	c.Require().NoError(err)

	relayerAddr := c.testWallet.Address().Hex()
	quotes, err := c.client.GetQuoteByRelayerAddress(c.GetTestContext(), relayerAddr)
	c.Require().NoError(err)

	c.Len(quotes, 1)
	expectedResp := model.GetQuoteResponse{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100",
		MaxOriginAmount: "200",
		FixedFee:        "10",
		RelayerAddr:     c.testWallet.Address().String(),
		UpdatedAt:       quotes[0].UpdatedAt,
	}
	c.Equal(expectedResp, *quotes[0])
}
