package client_test

import (
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

// TODO: @aurelius tese tests make a lot less sesnes w/ a composite index

func (c *ClientSuite) TestPutAndGetQuote() {
	req := model.PutQuoteRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100.0",
		MaxOriginAmount: "200.0",
		FixedFee:        "10.0",
	}

	err := c.client.PutQuote(&req)
	c.Require().NoError(err)

	quotes, err := c.client.GetAllQuotes()
	c.Require().NoError(err)

	expectedResp := model.GetQuoteResponse{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      100.0,
		MaxOriginAmount: 200.0,
		FixedFee:        10.0,
		RelayerAddr:     c.testWallet.Address().String(),
		UpdatedAt:       quotes[0].UpdatedAt,
	}
	c.Len(quotes, 1)
	c.Equal(expectedResp, *quotes[0])
}

func (c *ClientSuite) TestGetSpecificQuote() {
	putData := model.PutQuoteRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100.0",
		MaxOriginAmount: "200.0",
		FixedFee:        "10.0",
	}

	err := c.client.PutQuote(&putData)
	c.Require().NoError(err)

	quotes, err := c.client.GetSpecificQuote(&model.GetQuoteSpecificRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
	})
	c.Require().NoError(err)

	c.Len(quotes, 1)
	expectedResp := model.GetQuoteResponse{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      100.0,
		MaxOriginAmount: 200.0,
		FixedFee:        10.0,
		RelayerAddr:     c.testWallet.Address().String(),
		UpdatedAt:       quotes[0].UpdatedAt,
	}
	c.Equal(expectedResp, *quotes[0])
}

func (c *ClientSuite) TestGetQuoteByRelayerAddress() {
	putData := model.PutQuoteRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100.0",
		MaxOriginAmount: "200.0",
		FixedFee:        "10.0",
	}

	err := c.client.PutQuote(&putData)
	c.Require().NoError(err)

	relayerAddr := c.testWallet.Address().Hex()
	quotes, err := c.client.GetQuoteByRelayerAddress(relayerAddr)
	c.Require().NoError(err)

	c.Len(quotes, 1)
	expectedResp := model.GetQuoteResponse{
		OriginChainID:   1,
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     42161,
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      100.0,
		MaxOriginAmount: 200.0,
		FixedFee:        10.0,
		RelayerAddr:     c.testWallet.Address().String(),
		UpdatedAt:       quotes[0].UpdatedAt,
	}
	c.Equal(expectedResp, *quotes[0])
}
