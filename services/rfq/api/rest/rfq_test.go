package rest_test

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/signer/signer/localsigner"
	"github.com/synapsecns/sanguine/ethergo/signer/wallet"
	"github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

func runMockRelayer(c *ServerSuite, respCtx context.Context, relayerWallet wallet.Wallet, quoteResp *model.RelayerWsQuoteResponse, url, wsURL string) {
	// Create a relayer client
	relayerSigner := localsigner.NewSigner(relayerWallet.PrivateKey())
	relayerClient, err := client.NewAuthenticatedClient(metrics.Get(), url, &wsURL, relayerSigner)
	c.Require().NoError(err)

	// Create channels for active quote requests and responses
	reqChan := make(chan *model.ActiveRFQMessage)
	req := &model.SubscribeActiveRFQRequest{
		ChainIDs: []int{c.originChainID, c.destChainID},
	}
	respChan, err := relayerClient.SubscribeActiveQuotes(c.GetTestContext(), req, reqChan)
	c.Require().NoError(err)

	go func() {
		for {
			select {
			case <-respCtx.Done():
				return
			case msg := <-respChan:
				if msg == nil {
					continue
				}
				if msg.Op == "request_quote" {
					var quoteReq model.RelayerWsQuoteRequest
					err := json.Unmarshal(msg.Content, &quoteReq)
					if err != nil {
						c.Error(fmt.Errorf("error unmarshaling quote request: %w", err))
						continue
					}
					relayerAddr := relayerWallet.Address().Hex()
					quoteResp.RequestID = quoteReq.RequestID
					quoteResp.Data.RelayerAddress = &relayerAddr
					rawRespData, err := json.Marshal(quoteResp)
					if err != nil {
						c.Error(fmt.Errorf("error marshaling quote response: %w", err))
						continue
					}
					reqChan <- &model.ActiveRFQMessage{
						Op:      "send_quote",
						Content: json.RawMessage(rawRespData),
					}
				}
			}
		}
	}()
}

func verifyActiveQuoteRequest(c *ServerSuite, userReq *model.PutUserQuoteRequest, activeQuoteRequest *db.ActiveQuoteRequest, status db.ActiveQuoteRequestStatus) {
	c.Assert().Equal(uint64(userReq.Data.OriginChainID), activeQuoteRequest.OriginChainID)
	c.Assert().Equal(userReq.Data.OriginTokenAddr, activeQuoteRequest.OriginTokenAddr)
	c.Assert().Equal(uint64(userReq.Data.DestChainID), activeQuoteRequest.DestChainID)
	c.Assert().Equal(userReq.Data.DestTokenAddr, activeQuoteRequest.DestTokenAddr)
	c.Assert().Equal(userReq.Data.OriginAmount, activeQuoteRequest.OriginAmount.String())
	c.Assert().Equal(status, activeQuoteRequest.Status)
}

const (
	originTokenAddr = "0x1111111111111111111111111111111111111111"
	destTokenAddr   = "0x2222222222222222222222222222222222222222"
)

func (c *ServerSuite) TestActiveRFQSingleRelayer() {
	// Start the API server
	c.startQuoterAPIServer()

	url := fmt.Sprintf("http://localhost:%d", c.port)
	wsURL := fmt.Sprintf("ws://localhost:%d", c.wsPort)

	// Create a user client
	userWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	userSigner := localsigner.NewSigner(userWallet.PrivateKey())
	userClient, err := client.NewAuthenticatedClient(metrics.Get(), url, nil, userSigner)
	c.Require().NoError(err)

	// Prepare a user quote request
	userRequestAmount := big.NewInt(1_000_000)
	userQuoteReq := &model.PutUserQuoteRequest{
		Data: model.QuoteData{
			OriginChainID:    c.originChainID,
			OriginTokenAddr:  originTokenAddr,
			DestChainID:      c.destChainID,
			DestTokenAddr:    destTokenAddr,
			OriginAmount:     userRequestAmount.String(),
			ExpirationWindow: 10_000,
		},
		QuoteTypes: []string{"active"},
	}

	// Prepare the relayer quote response
	originAmount := userRequestAmount.String()
	destAmount := new(big.Int).Sub(userRequestAmount, big.NewInt(1000)).String()
	quoteResp := &model.RelayerWsQuoteResponse{
		Data: &model.QuoteData{
			OriginChainID:   c.originChainID,
			OriginTokenAddr: originTokenAddr,
			DestChainID:     c.destChainID,
			DestTokenAddr:   destTokenAddr,
			DestAmount:      &destAmount,
			OriginAmount:    originAmount,
		},
	}
	respCtx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()
	runMockRelayer(c, respCtx, c.relayerWallets[0], quoteResp, url, wsURL)

	// Submit the user quote request
	userQuoteResp, err := userClient.PutUserQuoteRequest(c.GetTestContext(), userQuoteReq)
	c.Require().NoError(err)

	// Assert the response
	c.Assert().True(userQuoteResp.Success)
	c.Assert().Equal("active", userQuoteResp.QuoteType)
	c.Assert().Equal(destAmount, *userQuoteResp.Data.DestAmount)
	c.Assert().Equal(originAmount, userQuoteResp.Data.OriginAmount)

	// Verify ActiveQuoteRequest insertion
	activeQuoteRequests, err := c.database.GetActiveQuoteRequests(c.GetTestContext())
	c.Require().NoError(err)
	verifyActiveQuoteRequest(c, userQuoteReq, activeQuoteRequests[0], db.Closed)
}

func (c *ServerSuite) TestActiveRFQExpiredRequest() {
	// Start the API server
	c.startQuoterAPIServer()

	url := fmt.Sprintf("http://localhost:%d", c.port)
	wsURL := fmt.Sprintf("ws://localhost:%d", c.wsPort)

	// Create a user client
	userWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	userSigner := localsigner.NewSigner(userWallet.PrivateKey())
	userClient, err := client.NewAuthenticatedClient(metrics.Get(), url, nil, userSigner)
	c.Require().NoError(err)

	// Prepare a user quote request
	userRequestAmount := big.NewInt(1_000_000)
	userQuoteReq := &model.PutUserQuoteRequest{
		Data: model.QuoteData{
			OriginChainID:    c.originChainID,
			OriginTokenAddr:  originTokenAddr,
			DestChainID:      c.destChainID,
			DestTokenAddr:    destTokenAddr,
			OriginAmount:     userRequestAmount.String(),
			ExpirationWindow: 0,
		},
		QuoteTypes: []string{"active"},
	}

	// Prepare the relayer quote response
	originAmount := userRequestAmount.String()
	destAmount := new(big.Int).Sub(userRequestAmount, big.NewInt(1000)).String()
	quoteResp := &model.RelayerWsQuoteResponse{
		Data: &model.QuoteData{
			OriginChainID:   c.originChainID,
			OriginTokenAddr: originTokenAddr,
			DestChainID:     c.destChainID,
			DestTokenAddr:   destTokenAddr,
			DestAmount:      &destAmount,
			OriginAmount:    originAmount,
		},
	}
	respCtx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()
	runMockRelayer(c, respCtx, c.relayerWallets[0], quoteResp, url, wsURL)

	// Submit the user quote request
	userQuoteResp, err := userClient.PutUserQuoteRequest(c.GetTestContext(), userQuoteReq)
	c.Require().NoError(err)

	// Assert the response
	c.Assert().False(userQuoteResp.Success)
	c.Assert().Equal("no quotes found", userQuoteResp.Reason)

	// Verify ActiveQuoteRequest insertion
	activeQuoteRequests, err := c.database.GetActiveQuoteRequests(c.GetTestContext())
	c.Require().NoError(err)
	verifyActiveQuoteRequest(c, userQuoteReq, activeQuoteRequests[0], db.Expired)
}

func (c *ServerSuite) TestActiveRFQMultipleRelayers() {
	// Start the API server
	c.startQuoterAPIServer()

	url := fmt.Sprintf("http://localhost:%d", c.port)
	wsURL := fmt.Sprintf("ws://localhost:%d", c.wsPort)

	// Create a user client
	userWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	userSigner := localsigner.NewSigner(userWallet.PrivateKey())
	userClient, err := client.NewAuthenticatedClient(metrics.Get(), url, nil, userSigner)
	c.Require().NoError(err)

	// Prepare a user quote request
	userRequestAmount := big.NewInt(1_000_000)
	userQuoteReq := &model.PutUserQuoteRequest{
		Data: model.QuoteData{
			OriginChainID:    c.originChainID,
			OriginTokenAddr:  originTokenAddr,
			DestChainID:      c.destChainID,
			DestTokenAddr:    destTokenAddr,
			OriginAmount:     userRequestAmount.String(),
			ExpirationWindow: 10_000,
		},
		QuoteTypes: []string{"active"},
	}

	// Prepare the relayer quote responses
	originAmount := userRequestAmount.String()
	destAmount := "999000"
	quoteResp := model.RelayerWsQuoteResponse{
		Data: &model.QuoteData{
			OriginChainID:   c.originChainID,
			OriginTokenAddr: originTokenAddr,
			DestChainID:     c.destChainID,
			DestTokenAddr:   destTokenAddr,
			DestAmount:      &destAmount,
			OriginAmount:    originAmount,
		},
	}

	// Create additional responses with worse prices
	destAmount2 := "998000"
	quoteResp2 := model.RelayerWsQuoteResponse{
		Data: &model.QuoteData{
			OriginChainID:   c.originChainID,
			OriginTokenAddr: originTokenAddr,
			DestChainID:     c.destChainID,
			DestTokenAddr:   destTokenAddr,
			DestAmount:      &destAmount2,
			OriginAmount:    originAmount,
		},
	}
	destAmount3 := "997000"
	quoteResp3 := model.RelayerWsQuoteResponse{
		Data: &model.QuoteData{
			OriginChainID:   c.originChainID,
			OriginTokenAddr: originTokenAddr,
			DestChainID:     c.destChainID,
			DestTokenAddr:   destTokenAddr,
			DestAmount:      &destAmount3,
			OriginAmount:    originAmount,
		},
	}
	respCtx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()
	runMockRelayer(c, respCtx, c.relayerWallets[0], &quoteResp, url, wsURL)
	runMockRelayer(c, respCtx, c.relayerWallets[1], &quoteResp2, url, wsURL)
	runMockRelayer(c, respCtx, c.relayerWallets[2], &quoteResp3, url, wsURL)

	// Submit the user quote request
	userQuoteResp, err := userClient.PutUserQuoteRequest(c.GetTestContext(), userQuoteReq)
	c.Require().NoError(err)

	// Assert the response
	c.Assert().True(userQuoteResp.Success)
	c.Assert().Equal("active", userQuoteResp.QuoteType)
	c.Assert().Equal(destAmount, *userQuoteResp.Data.DestAmount)
	c.Assert().Equal(originAmount, userQuoteResp.Data.OriginAmount)

	// Verify ActiveQuoteRequest insertion
	activeQuoteRequests, err := c.database.GetActiveQuoteRequests(c.GetTestContext())
	c.Require().NoError(err)
	verifyActiveQuoteRequest(c, userQuoteReq, activeQuoteRequests[0], db.Closed)
}

func (c *ServerSuite) TestActiveRFQFallbackToPassive() {
	// Start the API server
	c.startQuoterAPIServer()

	url := fmt.Sprintf("http://localhost:%d", c.port)
	wsURL := fmt.Sprintf("ws://localhost:%d", c.wsPort)

	// Create a user client
	userWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	userSigner := localsigner.NewSigner(userWallet.PrivateKey())
	userClient, err := client.NewAuthenticatedClient(metrics.Get(), url, nil, userSigner)
	c.Require().NoError(err)

	userRequestAmount := big.NewInt(1_000_000)

	// Upsert passive quotes into the database
	passiveQuotes := []db.Quote{
		{
			RelayerAddr:     c.relayerWallets[0].Address().Hex(),
			OriginChainID:   uint64(c.originChainID),
			OriginTokenAddr: originTokenAddr,
			DestChainID:     uint64(c.destChainID),
			DestTokenAddr:   destTokenAddr,
			DestAmount:      decimal.NewFromBigInt(new(big.Int).Sub(userRequestAmount, big.NewInt(1000)), 0),
			MaxOriginAmount: decimal.NewFromBigInt(userRequestAmount, 0),
			FixedFee:        decimal.NewFromInt(1000),
		},
	}

	for _, quote := range passiveQuotes {
		err := c.database.UpsertQuote(c.GetTestContext(), &quote)
		c.Require().NoError(err)
	}

	// Prepare user quote request with 0 expiration window
	userQuoteReq := &model.PutUserQuoteRequest{
		Data: model.QuoteData{
			OriginChainID:    c.originChainID,
			OriginTokenAddr:  originTokenAddr,
			DestChainID:      c.destChainID,
			DestTokenAddr:    destTokenAddr,
			OriginAmount:     userRequestAmount.String(),
			ExpirationWindow: 0,
		},
		QuoteTypes: []string{"active", "passive"},
	}

	// Prepare mock relayer response (which should be ignored due to 0 expiration window)
	destAmount := new(big.Int).Sub(userRequestAmount, big.NewInt(1000)).String()
	quoteResp := &model.RelayerWsQuoteResponse{
		Data: &model.QuoteData{
			OriginChainID:   c.originChainID,
			OriginTokenAddr: originTokenAddr,
			DestChainID:     c.destChainID,
			DestTokenAddr:   destTokenAddr,
			DestAmount:      &destAmount,
			OriginAmount:    userQuoteReq.Data.OriginAmount,
		},
	}

	respCtx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()

	// Run mock relayer even though we expect it to be ignored
	runMockRelayer(c, respCtx, c.relayerWallets[0], quoteResp, url, wsURL)

	// Submit the user quote request
	userQuoteResp, err := userClient.PutUserQuoteRequest(c.GetTestContext(), userQuoteReq)
	c.Require().NoError(err)

	// Assert the response
	c.Assert().True(userQuoteResp.Success)
	c.Assert().Equal("passive", userQuoteResp.QuoteType)
	c.Assert().Equal("998000", *userQuoteResp.Data.DestAmount) // destAmount is quote destAmount minus fixed fee
	c.Assert().Equal(userRequestAmount.String(), userQuoteResp.Data.OriginAmount)
	c.Assert().Equal(c.relayerWallets[0].Address().Hex(), *userQuoteResp.Data.RelayerAddress)
}

func (c *ServerSuite) TestActiveRFQPassiveBestQuote() {
	// Start the API server
	c.startQuoterAPIServer()

	url := fmt.Sprintf("http://localhost:%d", c.port)
	wsURL := fmt.Sprintf("ws://localhost:%d", c.wsPort)

	// Create a user client
	userWallet, err := wallet.FromRandom()
	c.Require().NoError(err)
	userSigner := localsigner.NewSigner(userWallet.PrivateKey())
	userClient, err := client.NewAuthenticatedClient(metrics.Get(), url, nil, userSigner)
	c.Require().NoError(err)

	userRequestAmount := big.NewInt(1_000_000)

	// Upsert passive quotes into the database
	passiveQuotes := []db.Quote{
		{
			RelayerAddr:     c.relayerWallets[0].Address().Hex(),
			OriginChainID:   uint64(c.originChainID),
			OriginTokenAddr: originTokenAddr,
			DestChainID:     uint64(c.destChainID),
			DestTokenAddr:   destTokenAddr,
			DestAmount:      decimal.NewFromBigInt(new(big.Int).Sub(userRequestAmount, big.NewInt(100)), 0),
			MaxOriginAmount: decimal.NewFromBigInt(userRequestAmount, 0),
			FixedFee:        decimal.NewFromInt(1000),
		},
	}

	for _, quote := range passiveQuotes {
		err := c.database.UpsertQuote(c.GetTestContext(), &quote)
		c.Require().NoError(err)
	}

	// Prepare user quote request with 0 expiration window
	userQuoteReq := &model.PutUserQuoteRequest{
		Data: model.QuoteData{
			OriginChainID:    c.originChainID,
			OriginTokenAddr:  originTokenAddr,
			DestChainID:      c.destChainID,
			DestTokenAddr:    destTokenAddr,
			OriginAmount:     userRequestAmount.String(),
			ExpirationWindow: 0,
		},
		QuoteTypes: []string{"active", "passive"},
	}

	// Prepare mock relayer response (which should be ignored due to 0 expiration window)
	destAmount := new(big.Int).Sub(userRequestAmount, big.NewInt(1000)).String()
	quoteResp := model.RelayerWsQuoteResponse{
		Data: &model.QuoteData{
			OriginChainID:   c.originChainID,
			OriginTokenAddr: originTokenAddr,
			DestChainID:     c.destChainID,
			DestTokenAddr:   destTokenAddr,
			DestAmount:      &destAmount,
			OriginAmount:    userQuoteReq.Data.OriginAmount,
		},
	}

	respCtx, cancel := context.WithCancel(c.GetTestContext())
	defer cancel()

	// Create additional responses with worse prices
	quoteResp2 := quoteResp
	destAmount2 := new(big.Int).Sub(userRequestAmount, big.NewInt(2000))
	destAmount2Str := destAmount2.String()
	quoteResp2.Data.DestAmount = &destAmount2Str
	quoteResp3 := quoteResp
	destAmount3 := new(big.Int).Sub(userRequestAmount, big.NewInt(3000))
	destAmount3Str := destAmount3.String()
	quoteResp3.Data.DestAmount = &destAmount3Str

	runMockRelayer(c, respCtx, c.relayerWallets[0], &quoteResp, url, wsURL)
	runMockRelayer(c, respCtx, c.relayerWallets[1], &quoteResp2, url, wsURL)
	runMockRelayer(c, respCtx, c.relayerWallets[2], &quoteResp3, url, wsURL)

	// Submit the user quote request
	userQuoteResp, err := userClient.PutUserQuoteRequest(c.GetTestContext(), userQuoteReq)
	c.Require().NoError(err)

	// Assert the response
	c.Assert().True(userQuoteResp.Success)
	c.Assert().Equal("passive", userQuoteResp.QuoteType)
	c.Assert().Equal("998900", *userQuoteResp.Data.DestAmount) // destAmount is quote destAmount minus fixed fee
	c.Assert().Equal(userRequestAmount.String(), userQuoteResp.Data.OriginAmount)
	c.Assert().Equal(c.relayerWallets[0].Address().Hex(), *userQuoteResp.Data.RelayerAddress)
}