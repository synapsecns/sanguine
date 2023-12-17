package client_test

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"strconv"

	"github.com/synapsecns/sanguine/services/rfq/api/client"
)

// TODO: @aurelius tese tests make a lot less sesnes w/ a composite index

func (c *ClientSuite) TestPutAndGetQuote() {
	putData := client.APIQuotePutRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100.0",
		Price:           "50.0",
		MaxOriginAmount: "200.0",
	}

	err := c.client.PutQuote(&putData)
	c.Require().NoError(err)

	quotes, err := c.client.GetAllQuotes()
	c.Require().NoError(err)

	c.assertEqual(putData, quotes[0])
}

func (c *ClientSuite) assertEqual(q1 client.APIQuotePutRequest, quote *db.Quote) {
	c.Assert().Equal(q1.OriginChainID, strconv.Itoa(int(quote.OriginChainID)))
	c.Assert().Equal(q1.OriginTokenAddr, quote.OriginTokenAddr)
	c.Assert().Equal(q1.DestChainID, strconv.Itoa(int(quote.DestChainID)))
	c.Assert().Equal(q1.DestTokenAddr, quote.DestTokenAddr)
	c.assertStringFloatEqual(q1.DestAmount, quote.DestAmount)
	c.assertStringFloatEqual(q1.Price, quote.Price)
	c.assertStringFloatEqual(q1.MaxOriginAmount, quote.MaxOriginAmount)
}

func (c *ClientSuite) assertStringFloatEqual(s1 string, f2 decimal.Decimal) {
	qDA, exact := f2.Float64()
	c.True(exact)

	q1DA, err := strconv.ParseFloat(s1, 64)
	c.Require().NoError(err)

	c.Assert().Equal(qDA, q1DA)
}

func (c *ClientSuite) TestGetSpecificQuote() {
	putData := client.APIQuotePutRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100.0",
		Price:           "50.0",
		MaxOriginAmount: "200.0",
	}

	err := c.client.PutQuote(&putData)
	fmt.Println("err", err)
	c.Require().NoError(err)

	quotes, err := c.client.GetSpecificQuote(&client.APIQuoteSpecificGetRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
	})
	c.Require().NoError(err)
	c.assertEqual(putData, quotes[0])
}

func (c *ClientSuite) TestGetQuoteByRelayerAddress() {
	putData := client.APIQuotePutRequest{
		OriginChainID:   "1",
		OriginTokenAddr: "0xOriginTokenAddr",
		DestChainID:     "42161",
		DestTokenAddr:   "0xDestTokenAddr",
		DestAmount:      "100.0",
		Price:           "50.0",
		MaxOriginAmount: "200.0",
	}

	err := c.client.PutQuote(&putData)
	fmt.Println("err", err)
	c.Require().NoError(err)

	relayerAddr := c.testWallet.Address().Hex()

	quotes, err := c.client.GetQuoteByRelayerAddress(relayerAddr)
	c.Require().NoError(err)
	c.assertEqual(putData, quotes[0])

	found := false
	for _, q := range quotes {
		if q.RelayerAddr == relayerAddr {
			found = true
			break
		}
	}
	c.Assert().True(found, "Quote for given relayer address not found")
}
