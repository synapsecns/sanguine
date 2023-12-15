package client_test

import (
	"fmt"

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
	fmt.Println("err", err)
	c.Require().NoError(err)

	quotes, err := c.client.GetAllQuotes()
	c.Require().NoError(err)
	fmt.Println(quotes)

	// found := false
	//for _, q := range quotes {
	//	if q.MaxOriginAmount == 123 {
	//		found = true
	//		break
	//	}
	//}
	//c.Assert().True(found, "Newly added quote not found")
}

func (c *ClientSuite) TestGetSpecificQuote() {
	putData := client.APIQuotePutRequest{
		//ID:              123,
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
	fmt.Println(quotes)
	//
	// found := false
	//for _, q := range quotes {
	//	if q.ID == 123 {
	//		found = true
	//		break
	//	}
	//}
	//c.Assert().True(found, "Newly added quote not found")
}

func (c *ClientSuite) TestGetQuoteByRelayerAddress() {
	putData := client.APIQuotePutRequest{
		//ID:              123,
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
	fmt.Println(quotes)

	found := false
	for _, q := range quotes {
		if q.RelayerAddr == relayerAddr {
			found = true
			break
		}
	}
	c.Assert().True(found, "Quote for given relayer address not found")
}
