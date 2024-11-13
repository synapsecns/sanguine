package rest

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/api/config"
	"github.com/synapsecns/sanguine/services/rfq/api/db"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const collectionTimeout = 1 * time.Minute

func (r *QuoterAPIServer) handleActiveRFQ(ctx context.Context, request *model.PutRFQRequest, requestID string) (quote *model.QuoteData) {
	ctx, span := r.handler.Tracer().Start(ctx, "handleActiveRFQ", trace.WithAttributes(
		attribute.String("user_address", request.UserAddress),
		attribute.String("request_id", requestID),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	// publish the quote request to all connected clients
	relayerReq := model.NewWsRFQRequest(request.Data, requestID)
	r.wsClients.Range(func(relayerAddr string, client WsClient) bool {
		sendCtx, sendSpan := r.handler.Tracer().Start(ctx, "sendQuoteRequest", trace.WithAttributes(
			attribute.String("relayer_address", relayerAddr),
			attribute.String("request_id", requestID),
		))
		defer metrics.EndSpan(sendSpan)

		subscribed := r.pubSubManager.IsSubscribed(relayerAddr, request.Data.OriginChainID, request.Data.DestChainID)
		span.SetAttributes(attribute.Bool("subscribed", subscribed))
		if subscribed {
			err := client.SendQuoteRequest(sendCtx, relayerReq)
			if err != nil {
				logger.Errorf("Error sending quote request to %s: %v", relayerAddr, err)
			}
		}
		return true
	})
	err := r.db.UpdateActiveQuoteRequestStatus(ctx, requestID, nil, db.Pending)
	if err != nil {
		logger.Errorf("Error updating active quote request status: %v", err)
	}

	// collect the responses and determine the best quote
	responses := r.collectRelayerResponses(ctx, request, requestID)
	for r, resp := range responses {
		relayerAddr := r
		quote = getBestQuote(quote, getRelayerQuoteData(request, resp))
		quote.RelayerAddress = &relayerAddr
	}
	err = r.recordActiveQuote(ctx, quote, requestID)
	if err != nil {
		logger.Errorf("Error recording active quote: %v", err)
	}

	return quote
}

func (r *QuoterAPIServer) collectRelayerResponses(ctx context.Context, request *model.PutRFQRequest, requestID string) (responses map[string]*model.WsRFQResponse) {
	ctx, span := r.handler.Tracer().Start(ctx, "collectRelayerResponses", trace.WithAttributes(
		attribute.String("user_address", request.UserAddress),
		attribute.String("request_id", requestID),
	))
	defer metrics.EndSpan(span)

	expireCtx, expireCancel := context.WithTimeout(ctx, time.Duration(request.Data.ExpirationWindow)*time.Millisecond)
	defer expireCancel()

	// don't cancel the collection context so that late responses can be collected in background
	// nolint:govet
	collectionCtx, _ := context.WithTimeout(ctx, time.Duration(request.Data.ExpirationWindow)*time.Millisecond+collectionTimeout)

	wg := sync.WaitGroup{}
	respMux := sync.Mutex{}
	responses = map[string]*model.WsRFQResponse{}
	r.wsClients.Range(func(relayerAddr string, client WsClient) bool {
		wg.Add(1)
		go func(client WsClient) {
			var respStatus db.ActiveQuoteResponseStatus
			var err error
			_, clientSpan := r.handler.Tracer().Start(collectionCtx, "collectRelayerResponses", trace.WithAttributes(
				attribute.String("relayer_address", relayerAddr),
				attribute.String("request_id", requestID),
			))
			defer func() {
				clientSpan.SetAttributes(attribute.String("status", respStatus.String()))
				metrics.EndSpanWithErr(clientSpan, err)
			}()

			defer wg.Done()
			resp, err := client.ReceiveQuoteResponse(collectionCtx, requestID)
			if err != nil {
				logger.Errorf("Error receiving quote response: %v", err)
				return
			}
			clientSpan.AddEvent("received quote response", trace.WithAttributes(
				attribute.String("relayer_address", relayerAddr),
				attribute.String("request_id", requestID),
				attribute.String("dest_amount", resp.DestAmount),
			))

			// validate the response
			respStatus = getQuoteResponseStatus(expireCtx, resp)
			if respStatus == db.Considered {
				respMux.Lock()
				responses[relayerAddr] = resp
				respMux.Unlock()
			}

			// record the response
			err = r.db.InsertActiveQuoteResponse(collectionCtx, resp, relayerAddr, respStatus)
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

func getRelayerQuoteData(request *model.PutRFQRequest, resp *model.WsRFQResponse) *model.QuoteData {
	return &model.QuoteData{
		OriginChainID:     request.Data.OriginChainID,
		DestChainID:       request.Data.DestChainID,
		OriginTokenAddr:   request.Data.OriginTokenAddr,
		DestTokenAddr:     request.Data.DestTokenAddr,
		OriginAmountExact: request.Data.OriginAmountExact,
		DestAmount:        &resp.DestAmount,
		QuoteID:           &resp.QuoteID,
	}
}

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

func getQuoteResponseStatus(ctx context.Context, resp *model.WsRFQResponse) db.ActiveQuoteResponseStatus {
	respStatus := db.Considered
	err := validateRelayerQuoteResponse(resp)
	if err != nil {
		respStatus = db.Malformed
		logger.Errorf("Error validating quote response: %v", err)
	} else if ctx.Err() != nil {
		respStatus = db.PastExpiration
	}
	return respStatus
}

func validateRelayerQuoteResponse(resp *model.WsRFQResponse) error {
	_, ok := new(big.Int).SetString(resp.DestAmount, 10)
	if !ok {
		return fmt.Errorf("dest amount is invalid")
	}
	// TODO: compute quote ID from request
	resp.QuoteID = uuid.New().String()
	return nil
}

func (r *QuoterAPIServer) recordActiveQuote(ctx context.Context, quote *model.QuoteData, requestID string) (err error) {
	if quote == nil {
		err = r.db.UpdateActiveQuoteRequestStatus(ctx, requestID, nil, db.Expired)
		if err != nil {
			logger.Errorf("Error updating active quote request status: %v", err)
		}
	} else {
		err = r.db.UpdateActiveQuoteRequestStatus(ctx, requestID, quote.QuoteID, db.Closed)
		if err != nil {
			logger.Errorf("Error updating active quote request status: %v", err)
		}
		err = r.db.UpdateActiveQuoteResponseStatus(ctx, *quote.QuoteID, db.Returned)
		if err != nil {
			return fmt.Errorf("error updating active quote response status: %w", err)
		}
	}
	return nil
}

func (r *QuoterAPIServer) handlePassiveRFQ(ctx context.Context, request *model.PutRFQRequest) (*model.QuoteData, error) {
	ctx, span := r.handler.Tracer().Start(ctx, "handlePassiveRFQ", trace.WithAttributes(
		attribute.String("user_address", request.UserAddress),
	))
	defer metrics.EndSpan(span)

	quotes, err := r.db.GetQuotesByOriginAndDestination(ctx, uint64(request.Data.OriginChainID), request.Data.OriginTokenAddr, uint64(request.Data.DestChainID), request.Data.DestTokenAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to get quotes: %w", err)
	}

	quote, err := getPassiveQuote(r.cfg, quotes, request)
	if err != nil {
		return nil, fmt.Errorf("failed to get passive quote: %w", err)
	}

	return quote, nil
}

func getPassiveQuote(cfg config.Config, quotes []*db.Quote, request *model.PutRFQRequest) (*model.QuoteData, error) {
	quotes = filterQuoteAge(cfg, quotes)

	originAmount, ok := new(big.Int).SetString(request.Data.OriginAmountExact, 10)
	if !ok {
		return nil, errors.New("invalid origin amount exact")
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
		if rawDestAmountInt.Cmp(quote.FixedFee.BigInt()) < 0 {
			continue
		}
		destAmount := new(big.Int).Sub(rawDestAmountInt, quote.FixedFee.BigInt()).String()
		//nolint:gosec
		quoteData := &model.QuoteData{
			OriginChainID:     int(quote.OriginChainID),
			DestChainID:       int(quote.DestChainID),
			OriginTokenAddr:   quote.OriginTokenAddr,
			DestTokenAddr:     quote.DestTokenAddr,
			OriginAmountExact: quote.MaxOriginAmount.String(),
			DestAmount:        &destAmount,
			RelayerAddress:    &quote.RelayerAddr,
		}
		bestQuote = getBestQuote(bestQuote, quoteData)
	}

	return bestQuote, nil
}
