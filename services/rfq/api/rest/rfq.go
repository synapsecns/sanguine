package rest

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

func getBestQuote(a, b *model.QuoteData) *model.QuoteData {
	if a == nil && b == nil {
		return nil
	}
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	aAmount, _ := new(big.Int).SetString(*a.DestAmount, 10)
	bAmount, _ := new(big.Int).SetString(*b.DestAmount, 10)
	if aAmount.Cmp(bAmount) > 0 {
		return a
	}
	return b
}

const collectionTimeout = 1 * time.Minute

func (r *QuoterAPIServer) handleActiveRFQ(ctx context.Context, request *model.PutUserQuoteRequest, requestID string) (quote *model.QuoteData) {
	expireCtx, _ := context.WithTimeout(ctx, time.Duration(request.Data.ExpirationWindow)*time.Millisecond)
	collectionCtx, _ := context.WithTimeout(ctx, time.Duration(request.Data.ExpirationWindow)*time.Millisecond+collectionTimeout)

	// publish the quote request to all connected clients
	relayerReq := model.NewRelayerWsQuoteRequest(request.Data, requestID)
	r.wsClients.Range(func(key string, client WsClient) bool {
		client.SendQuoteRequest(expireCtx, relayerReq)
		return true
	})

	err := r.db.UpdateActiveQuoteRequestStatus(ctx, requestID, db.Pending)
	if err != nil {
		logger.Errorf("Error updating active quote request status: %v", err)
	}

	// collect responses from all clients until expiration window closes
	wg := sync.WaitGroup{}
	respMux := sync.Mutex{}
	responses := map[string]*model.RelayerWsQuoteResponse{}
	r.wsClients.Range(func(key string, client WsClient) bool {
		wg.Add(1)
		go func(client WsClient) {
			defer wg.Done()
			resp, err := client.ReceiveQuoteResponse(collectionCtx)
			if err != nil {
				logger.Errorf("Error receiving quote response: %v", err)
				return
			}
			respMux.Lock()
			responses[key] = resp
			respMux.Unlock()

			// record the response
			respStatus := db.Considered
			if expireCtx.Err() != nil {
				respStatus = db.PastExpiration
			}
			err = r.db.InsertActiveQuoteResponse(ctx, resp, respStatus)
			if err != nil {
				logger.Errorf("Error inserting active quote response: %v", err)
			}
		}(client)
		return true
	})

	// wait for all responses to be received, or expiration
	select {
	case <-expireCtx.Done():
		// request expired before all responses were received
	case <-func() chan struct{} {
		ch := make(chan struct{})
		go func() {
			wg.Wait()
			close(ch)
		}()
		return ch
	}():
		// all responses received
	}

	// construct the response
	// at this point, all responses should have been validated
	var bestQuoteID string
	for _, resp := range responses {
		quote = getBestQuote(quote, &resp.Data)
		bestQuoteID = resp.QuoteID
	}

	if quote == nil {
		err = r.db.UpdateActiveQuoteRequestStatus(ctx, requestID, db.Expired)
		if err != nil {
			logger.Errorf("Error updating active quote request status: %v", err)
		}
	} else {
		err = r.db.UpdateActiveQuoteRequestStatus(ctx, requestID, db.Fulfilled)
		if err != nil {
			logger.Errorf("Error updating active quote request status: %v", err)
		}
		err = r.db.UpdateActiveQuoteResponseStatus(ctx, bestQuoteID, db.Returned)
		if err != nil {
			logger.Errorf("Error updating active quote response status: %v", err)
		}
	}

	return quote
}

func (r *QuoterAPIServer) handlePassiveRFQ(ctx context.Context, request *model.PutUserQuoteRequest) (*model.QuoteData, error) {
	quotes, err := r.db.GetQuotesByOriginAndDestination(ctx, uint64(request.Data.OriginChainID), request.Data.OriginTokenAddr, uint64(request.Data.DestChainID), request.Data.DestTokenAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get quotes: %w", err)
	}

	originAmount, ok := new(big.Int).SetString(request.Data.OriginAmount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid origin amount")
	}

	var bestQuote *model.QuoteData
	for _, quote := range quotes {
		quoteOriginAmount, ok := new(big.Int).SetString(quote.MaxOriginAmount.String(), 10)
		if !ok {
			continue
		}
		if quoteOriginAmount.Cmp(originAmount) < 0 {
			continue
		}
		quotePrice := new(big.Float).Quo(
			new(big.Float).SetInt(quote.DestAmount.BigInt()),
			new(big.Float).SetInt(quote.MaxOriginAmount.BigInt()),
		)

		rawDestAmount := new(big.Float).Mul(
			new(big.Float).SetInt(originAmount),
			quotePrice,
		)

		rawDestAmountInt, _ := rawDestAmount.Int(nil)
		destAmount := new(big.Int).Sub(rawDestAmountInt, quote.FixedFee.BigInt()).String()
		quoteData := &model.QuoteData{
			OriginChainID:   int(quote.OriginChainID),
			DestChainID:     int(quote.DestChainID),
			OriginTokenAddr: quote.OriginTokenAddr,
			DestTokenAddr:   quote.DestTokenAddr,
			OriginAmount:    quote.MaxOriginAmount.String(),
			DestAmount:      &destAmount,
			RelayerAddress:  &quote.RelayerAddr,
		}
		bestQuote = getBestQuote(bestQuote, quoteData)
	}

	return bestQuote, nil
}
