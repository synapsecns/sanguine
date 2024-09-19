package rest

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
)

const collectionTimeout = 1 * time.Minute

func (r *QuoterAPIServer) handleActiveRFQ(ctx context.Context, request *model.PutUserQuoteRequest, requestID string) (quote *model.QuoteData) {
	// publish the quote request to all connected clients
	relayerReq := model.NewRelayerWsQuoteRequest(request.Data, requestID)
	r.wsClients.Range(func(key string, client WsClient) bool {
		client.SendQuoteRequest(ctx, relayerReq)
		return true
	})
	err := r.db.UpdateActiveQuoteRequestStatus(ctx, requestID, db.Pending)
	if err != nil {
		logger.Errorf("Error updating active quote request status: %v", err)
	}

	// collect the responses and determine the best quote
	responses := r.collectRelayerResponses(ctx, request)
	var quoteID string
	var isUpdated bool
	for _, resp := range responses {
		quote, isUpdated = getBestQuote(quote, resp.Data)
		if isUpdated {
			quoteID = resp.QuoteID
		}
	}
	err = r.recordActiveQuote(ctx, quote, requestID, quoteID)
	if err != nil {
		logger.Errorf("Error recording active quote: %v", err)
	}

	return quote
}

func (r *QuoterAPIServer) collectRelayerResponses(ctx context.Context, request *model.PutUserQuoteRequest) (responses map[string]*model.RelayerWsQuoteResponse) {
	expireCtx, expireCancel := context.WithTimeout(ctx, time.Duration(request.Data.ExpirationWindow)*time.Millisecond)
	defer expireCancel()

	// don't cancel the collection context so that late responses can be collected in background
	collectionCtx, _ := context.WithTimeout(ctx, time.Duration(request.Data.ExpirationWindow)*time.Millisecond+collectionTimeout)

	wg := sync.WaitGroup{}
	respMux := sync.Mutex{}
	responses = map[string]*model.RelayerWsQuoteResponse{}
	r.wsClients.Range(func(relayerAddr string, client WsClient) bool {
		wg.Add(1)
		go func(client WsClient) {
			defer wg.Done()
			resp, err := client.ReceiveQuoteResponse(collectionCtx)
			if err != nil {
				logger.Errorf("Error receiving quote response: %v", err)
				return
			}

			// validate the response
			respStatus := getQuoteResponseStatus(expireCtx, resp, relayerAddr)
			if respStatus == db.Considered {
				respMux.Lock()
				responses[relayerAddr] = resp
				respMux.Unlock()
			}

			// record the response
			err = r.db.InsertActiveQuoteResponse(collectionCtx, resp, respStatus)
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

	return responses
}

func getBestQuote(a, b *model.QuoteData) (*model.QuoteData, bool) {
	if a == nil && b == nil {
		return nil, false
	}
	if a == nil {
		return b, true
	}
	if b == nil {
		return a, false
	}
	aAmount, _ := new(big.Int).SetString(*a.DestAmount, 10)
	bAmount, _ := new(big.Int).SetString(*b.DestAmount, 10)
	if aAmount.Cmp(bAmount) > 0 {
		return a, false
	}
	return b, true
}

func getQuoteResponseStatus(ctx context.Context, resp *model.RelayerWsQuoteResponse, relayerAddr string) db.ActiveQuoteResponseStatus {
	respStatus := db.Considered
	err := validateRelayerQuoteResponse(relayerAddr, resp)
	if err != nil {
		respStatus = db.Malformed
		logger.Errorf("Error validating quote response: %v", err)
	} else if ctx.Err() != nil {
		respStatus = db.PastExpiration
	}
	return respStatus
}

func validateRelayerQuoteResponse(relayerAddr string, resp *model.RelayerWsQuoteResponse) error {
	if resp.Data.RelayerAddress == nil {
		return fmt.Errorf("relayer address is nil")
	}
	// TODO: compute quote ID from request
	resp.QuoteID = uuid.New().String()
	resp.Data.RelayerAddress = &relayerAddr
	return nil
}

func (r *QuoterAPIServer) recordActiveQuote(ctx context.Context, quote *model.QuoteData, requestID, quoteID string) (err error) {
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
		err = r.db.UpdateActiveQuoteResponseStatus(ctx, quoteID, db.Returned)
		if err != nil {
			return fmt.Errorf("error updating active quote response status: %w", err)
		}
	}
	return nil
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
		bestQuote, _ = getBestQuote(bestQuote, quoteData)
	}

	return bestQuote, nil
}
