package rest

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

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

func (r *QuoterAPIServer) handleActiveRFQ(ctx context.Context, request *model.PutUserQuoteRequest) (quote *model.QuoteData) {
	rfqCtx, _ := context.WithTimeout(ctx, time.Duration(request.Data.ExpirationWindow)*time.Millisecond)
	fmt.Printf("started rfq ctx at %s\n", time.Now().Format("2006-01-02 15:04:05"))

	// publish the quote request to all connected clients
	relayerReq := model.NewRelayerWsQuoteRequest(request.Data)
	r.wsClients.Range(func(key string, client WsClient) bool {
		client.SendQuoteRequest(rfqCtx, relayerReq)
		fmt.Printf("sent quote request at %s\n", time.Now().Format("2006-01-02 15:04:05"))
		return true
	})

	// collect responses from all clients until expiration window closes
	wg := sync.WaitGroup{}
	respMux := sync.Mutex{}
	responses := map[string]*model.RelayerWsQuoteResponse{}
	r.wsClients.Range(func(key string, client WsClient) bool {
		wg.Add(1)
		go func(client WsClient) {
			defer wg.Done()
			resp, err := client.ReceiveQuoteResponse(rfqCtx)
			fmt.Printf("got quote response at %s\n", time.Now().Format("2006-01-02 15:04:05"))
			if err != nil {
				logger.Errorf("Error receiving quote response: %v", err)
				return
			}
			respMux.Lock()
			responses[key] = resp
			respMux.Unlock()
		}(client)
		return true
	})

	select {
	case <-rfqCtx.Done():
		// Context expired before all responses were received
	case <-func() chan struct{} {
		ch := make(chan struct{})
		go func() {
			wg.Wait()
			close(ch)
		}()
		return ch
	}():
		// All responses received
	}

	// construct the response
	// at this point, all responses should have been validated
	for _, resp := range responses {
		quote = getBestQuote(quote, &resp.Data)
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