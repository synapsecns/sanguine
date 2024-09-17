package rest_test

import (
	"context"
	"fmt"
	"math/big"

	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

func (c *ServerSuite) TestHandleActiveRFQ() {
	// Start the API server
	c.startQuoterAPIServer()

	url := fmt.Sprintf("http://localhost:%d", c.port)
	wsURL := fmt.Sprintf("ws://localhost:%d", c.wsPort)

	// Create a relayer client
	relayerSigner := localsigner.NewSigner(c.testWallet.PrivateKey())
	relayerClient, err := client.NewAuthenticatedClient(metrics.Get(), url, &wsURL, relayerSigner)
	c.Require().NoError(err)

	// Create a user client
	userWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	userSigner := localsigner.NewSigner(userWallet.PrivateKey())
	userClient, err := client.NewAuthenticatedClient(metrics.Get(), url, nil, userSigner)
	c.Require().NoError(err)

	// Create channels for active quote requests and responses
	reqChan := make(chan *model.ActiveRFQMessage)
	req := &model.SubscribeActiveRFQRequest{
		ChainIDs: []int{c.originChainID, c.destChainID},
	}
	respChan, err := relayerClient.SubscribeActiveQuotes(c.GetTestContext(), req, reqChan)
	c.Require().NoError(err)

	// Create a goroutine to handle incoming quote requests
	userRequestAmount := big.NewInt(1_000_000)
	originAmount := userRequestAmount.String()
	destAmount := new(big.Int).Sub(userRequestAmount, big.NewInt(1000)).String()
	respCtx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()
	go func() {
		for {
			select {
			case <-respCtx.Done():
				return
			case msg := <-respChan:
				if msg.Op == "request_quote" {
					quoteReq, ok := msg.Content.(*model.RelayerWsQuoteRequest)
					if !ok {
						c.Error(fmt.Errorf("msg.Content is not a model.RelayerWsQuoteRequest"))
						continue
					}
					quoteResp := &model.RelayerWsQuoteResponse{
						Data: model.QuoteData{
							OriginChainID:   quoteReq.Data.OriginChainID,
							OriginTokenAddr: quoteReq.Data.OriginTokenAddr,
							DestChainID:     quoteReq.Data.DestChainID,
							DestTokenAddr:   quoteReq.Data.DestTokenAddr,
							DestAmount:      &destAmount,
							OriginAmount:    originAmount,
						},
					}
					reqChan <- &model.ActiveRFQMessage{
						Op:      "send_quote",
						Content: quoteResp,
					}
				}
			}
		}
	}()

	// Prepare a user quote request
	userQuoteReq := &model.PutUserQuoteRequest{
		Data: model.QuoteData{
			OriginChainID:    1,
			OriginTokenAddr:  "0x1111111111111111111111111111111111111111",
			DestChainID:      2,
			DestTokenAddr:    "0x2222222222222222222222222222222222222222",
			OriginAmount:     userRequestAmount.String(),
			ExpirationWindow: 50_000,
		},
		QuoteTypes: []string{"active"},
	}

	// Submit the user quote request
	userQuoteResp, err := userClient.PutUserQuoteRequest(c.GetTestContext(), userQuoteReq)
	c.Require().NoError(err)

	// Assert the response
	c.Assert().True(userQuoteResp.Success)
	c.Assert().Equal("active", userQuoteResp.QuoteType)
	c.Assert().Equal(destAmount, userQuoteResp.Data.DestAmount)
	c.Assert().Equal(originAmount, userQuoteResp.Data.OriginAmount)
}
