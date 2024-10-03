// Package quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/synapsecns/sanguine/contrib/screener-api/client"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"github.com/synapsecns/sanguine/services/rfq/util"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/api/rest"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
)

var logger = log.Logger("quoter")

// Quoter submits quotes to the RFQ API.
//
//go:generate go run github.com/vektra/mockery/v2 --name Quoter --output ./mocks --case=underscore
type Quoter interface {
	// SubmitAllQuotes submits all quotes to the RFQ API.
	SubmitAllQuotes(ctx context.Context) (err error)
	// SubscribeActiveRFQ subscribes to the RFQ websocket API.
	SubscribeActiveRFQ(ctx context.Context) (err error)
	// ShouldProcess determines if a quote should be processed.
	// We do this by either saving all quotes in-memory, and refreshing via GetSelfQuotes() through the API
	// The first comparison is does bridge transaction OriginChainID+TokenAddr match with a quote + DestChainID+DestTokenAddr, then we look to see if we have enough amount to relay it + if the price fits our bounds (based on that the Relayer is relaying the destination token for the origin)
	// validateQuote(BridgeEvent)
	ShouldProcess(ctx context.Context, quote reldb.QuoteRequest) (bool, error)
	// IsProfitable determines if a quote is profitable, i.e. we will not lose money on it, net of fees.
	IsProfitable(ctx context.Context, quote reldb.QuoteRequest) (bool, error)
	// GetPrice gets the price of a token.
	GetPrice(ctx context.Context, tokenName string) (float64, error)
}

// Manager submits quotes to the RFQ API.
// TODO: should be unexported.
type Manager struct {
	// config is the relayer's config.
	config relconfig.Config
	// inventoryManager is used to get the relayer's inventory.
	inventoryManager inventory.Manager
	// rfqClient is used to communicate with the RFQ API.
	rfqClient rfqAPIClient.AuthenticatedClient
	// relayerSigner is the signer used by the relayer to interact on chain
	relayerSigner signer.Signer
	// feePricer is used to price fees.
	feePricer pricer.FeePricer
	// metricsHandler handles traces, etc
	metricsHandler metrics.Handler
	// quotableTokens is a map of token -> list of quotable tokens.
	// should be removed in config overhaul
	quotableTokens map[string][]string
	// screener is used to screen addresses.
	screener client.ScreenerClient
	// relayPaused is set when the RFQ API is found to be offline, which
	// lets the quoter indicate that quotes should not be relayed.
	relayPaused atomic.Bool
	// meter is the meter used by this package.
	meter metric.Meter
	// quoteAmountGauge stores a histogram of quote amounts.
	quoteAmountGauge metric.Float64ObservableGauge
	// currentQuotes is used for recording quote metrics.
	currentQuotes []model.PutRelayerQuoteRequest
}

// NewQuoterManager creates a new QuoterManager.
func NewQuoterManager(config relconfig.Config, metricsHandler metrics.Handler, inventoryManager inventory.Manager, relayerSigner signer.Signer, feePricer pricer.FeePricer, apiClient rfqAPIClient.AuthenticatedClient) (Quoter, error) {
	qt := make(map[string][]string)

	// fix any casing issues.
	var err error
	for tokenID, destTokenIDs := range config.QuotableTokens {
		processedDestTokens := make([]string, len(destTokenIDs))
		for i := range destTokenIDs {
			processedDestTokens[i], err = relconfig.SanitizeTokenID(destTokenIDs[i])
			if err != nil {
				return nil, fmt.Errorf("error sanitizing dest token ID: %w", err)
			}
		}
		sanitizedID, err := relconfig.SanitizeTokenID(tokenID)
		if err != nil {
			return nil, fmt.Errorf("error sanitizing token ID: %w", err)
		}
		qt[sanitizedID] = processedDestTokens
	}

	var ss client.ScreenerClient
	if config.ScreenerAPIUrl != "" {
		ss, err = client.NewClient(metricsHandler, config.ScreenerAPIUrl)
		if err != nil {
			return nil, fmt.Errorf("error creating screener client: %w", err)
		}
	}

	m := &Manager{
		config:           config,
		inventoryManager: inventoryManager,
		rfqClient:        apiClient,
		quotableTokens:   qt,
		relayerSigner:    relayerSigner,
		metricsHandler:   metricsHandler,
		feePricer:        feePricer,
		screener:         ss,
		meter:            metricsHandler.Meter(meterName),
		currentQuotes:    []model.PutRelayerQuoteRequest{},
	}

	m.quoteAmountGauge, err = m.meter.Float64ObservableGauge("quote_amount")
	if err != nil {
		return nil, fmt.Errorf("error creating quote amount gauge: %w", err)
	}

	_, err = m.meter.RegisterCallback(m.recordQuoteAmounts, m.quoteAmountGauge)
	if err != nil {
		return nil, fmt.Errorf("could not register callback: %w", err)
	}

	return m, nil
}

// ShouldProcess determines if a quote should be processed.
func (m *Manager) ShouldProcess(parentCtx context.Context, quote reldb.QuoteRequest) (res bool, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "shouldProcess", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(quote.TransactionID[:])),
	))

	defer func() {
		span.AddEvent("result", trace.WithAttributes(attribute.Bool("result", res)))
		metrics.EndSpanWithErr(span, err)
	}()

	if m.relayPaused.Load() {
		span.AddEvent("relayPaused is set due to RFQ API being offline")
		return false, nil
	}

	if m.screener != nil {
		// screen sender and recipient in parallel
		g, gctx := errgroup.WithContext(ctx)
		var senderBlocked, recipientBlocked bool
		g.Go(func() error {
			senderBlocked, err = m.screener.ScreenAddress(gctx, quote.Transaction.OriginSender.String())
			if err != nil {
				span.RecordError(fmt.Errorf("error screening address: %w", err))
				return fmt.Errorf("error screening address: %w", err)
			}
			return nil
		})
		g.Go(func() error {
			recipientBlocked, err = m.screener.ScreenAddress(gctx, quote.Transaction.DestRecipient.String())
			if err != nil {
				span.RecordError(fmt.Errorf("error screening address: %w", err))
				return fmt.Errorf("error screening address: %w", err)
			}
			return nil
		})

		err = g.Wait()
		if err != nil {
			return false, fmt.Errorf("error screening addresses: %w", err)
		}
		if senderBlocked || recipientBlocked {
			span.SetAttributes(
				attribute.Bool("sender_blocked", senderBlocked),
				attribute.Bool("recipient_blocked", recipientBlocked),
			)
			return false, nil
		}
	}

	// allowed pairs for this origin token on the destination
	destPairs := m.quotableTokens[quote.GetOriginIDPair()]
	if !(slices.Contains(destPairs, quote.GetDestIDPair())) {
		span.AddEvent(fmt.Sprintf("%s not in %s or %s not found", quote.GetDestIDPair(), strings.Join(destPairs, ", "), quote.GetOriginIDPair()))
		return false, nil
	}

	// handle decimals.
	// this will never get hit if we're operating correctly.
	if quote.OriginTokenDecimals != quote.DestTokenDecimals {
		span.AddEvent("Pairing tokens with two different decimals is disabled as a safety feature right now.")
		return false, nil
	}

	// all checks have passed
	return true, nil
}

// IsProfitable determines if a quote is profitable, i.e. we will not lose money on it, net of fees.
func (m *Manager) IsProfitable(parentCtx context.Context, quote reldb.QuoteRequest) (isProfitable bool, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "IsProfitable")

	defer func() {
		span.AddEvent("result", trace.WithAttributes(attribute.Bool("result", isProfitable)))
		metrics.EndSpanWithErr(span, err)
	}()

	destTokenID, err := m.config.GetTokenName(quote.Transaction.DestChainId, quote.Transaction.DestToken.String())
	if err != nil {
		return false, fmt.Errorf("error getting dest token ID: %w", err)
	}
	fee, err := m.feePricer.GetTotalFee(ctx, quote.Transaction.OriginChainId, quote.Transaction.DestChainId, destTokenID, false)
	if err != nil {
		return false, fmt.Errorf("error getting total fee: %w", err)
	}

	cost := new(big.Int).Add(quote.Transaction.DestAmount, fee)

	span.AddEvent("fee", trace.WithAttributes(attribute.String("fee", fee.String())))
	span.AddEvent("cost", trace.WithAttributes(attribute.String("cost", cost.String())))
	span.AddEvent("dest_amount", trace.WithAttributes(attribute.String("dest_amount", quote.Transaction.DestAmount.String())))
	span.AddEvent("origin_amount", trace.WithAttributes(attribute.String("origin_amount", quote.Transaction.OriginAmount.String())))

	// NOTE: this logic assumes that the origin and destination tokens have the same price.
	return quote.Transaction.OriginAmount.Cmp(cost) >= 0, nil
}

// SubmitAllQuotes submits all quotes to the RFQ API.
func (m *Manager) SubmitAllQuotes(ctx context.Context) (err error) {
	ctx, span := m.metricsHandler.Tracer().Start(ctx, "SubmitAllQuotes")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	inv, err := m.inventoryManager.GetCommittableBalances(ctx, inventory.SkipDBCache())
	if err != nil {
		return fmt.Errorf("error getting committable balances: %w", err)
	}

	return m.prepareAndSubmitQuotes(ctx, inv)
}

// SubscribeActiveRFQ subscribes to the RFQ websocket API.
// This function is blocking and will run until the context is canceled.
func (m *Manager) SubscribeActiveRFQ(ctx context.Context) (err error) {
	ctx, span := m.metricsHandler.Tracer().Start(ctx, "SubscribeActiveRFQ")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	chainIDs := []int{}
	for chainID := range m.config.Chains {
		chainIDs = append(chainIDs, chainID)
	}
	req := model.SubscribeActiveRFQRequest{
		ChainIDs: chainIDs,
	}
	span.SetAttributes(attribute.IntSlice("chain_ids", chainIDs))

	reqChan := make(chan *model.ActiveRFQMessage)
	respChan, err := m.rfqClient.SubscribeActiveQuotes(ctx, &req, reqChan)
	if err != nil {
		return fmt.Errorf("error subscribing to active quotes: %w", err)
	}
	span.AddEvent("subscribed to active quotes")
	for {
		select {
		case <-ctx.Done():
			return nil
		case msg, ok := <-respChan:
			if !ok {
				return errors.New("ws channel closed")
			}
			if msg == nil {
				continue
			}
			resp, err := m.generateActiveRFQ(ctx, msg)
			if err != nil {
				return fmt.Errorf("error generating active RFQ message: %w", err)
			}
			reqChan <- resp
		}
	}
}

// getActiveRFQ handles an active RFQ message.
//
//nolint:nilnil
func (m *Manager) generateActiveRFQ(ctx context.Context, msg *model.ActiveRFQMessage) (resp *model.ActiveRFQMessage, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(ctx, "generateActiveRFQ", trace.WithAttributes(
		attribute.String("op", msg.Op),
		attribute.String("content", string(msg.Content)),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	if msg.Op != rest.RequestQuoteOp {
		span.AddEvent("not a request quote op")
		return nil, nil
	}

	inv, err := m.inventoryManager.GetCommittableBalances(ctx, inventory.SkipDBCache())
	if err != nil {
		return nil, fmt.Errorf("error getting committable balances: %w", err)
	}

	var rfqRequest model.WsRFQRequest
	err = json.Unmarshal(msg.Content, &rfqRequest)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling quote data: %w", err)
	}
	span.SetAttributes(attribute.String("request_id", rfqRequest.RequestID))

	quoteInput := QuoteInput{
		OriginChainID:   rfqRequest.Data.OriginChainID,
		DestChainID:     rfqRequest.Data.DestChainID,
		OriginTokenAddr: common.HexToAddress(rfqRequest.Data.OriginTokenAddr),
		DestTokenAddr:   common.HexToAddress(rfqRequest.Data.DestTokenAddr),
		OriginBalance:   inv[rfqRequest.Data.OriginChainID][common.HexToAddress(rfqRequest.Data.OriginTokenAddr)],
		DestBalance:     inv[rfqRequest.Data.DestChainID][common.HexToAddress(rfqRequest.Data.DestTokenAddr)],
	}

	rawQuote, err := m.generateQuote(ctx, quoteInput)
	if err != nil {
		return nil, fmt.Errorf("error generating quote: %w", err)
	}
	span.SetAttributes(attribute.String("dest_amount", rawQuote.DestAmount))

	rfqResp := model.WsRFQResponse{
		RequestID:  rfqRequest.RequestID,
		DestAmount: rawQuote.DestAmount,
	}
	span.SetAttributes(attribute.String("dest_amount", rawQuote.DestAmount))
	respBytes, err := json.Marshal(rfqResp)
	if err != nil {
		return nil, fmt.Errorf("error serializing response: %w", err)
	}
	resp = &model.ActiveRFQMessage{
		Op:      rest.SendQuoteOp,
		Content: respBytes,
	}
	span.AddEvent("generated response")

	return resp, nil
}

// GetPrice gets the price of a token.
func (m *Manager) GetPrice(parentCtx context.Context, tokenName string) (_ float64, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "GetPrice")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	price, err := m.feePricer.GetTokenPrice(ctx, tokenName)
	if err != nil {
		return 0, fmt.Errorf("error getting price: %w", err)
	}

	return price, nil
}

// Prepares and submits quotes based on inventory.
func (m *Manager) prepareAndSubmitQuotes(ctx context.Context, inv map[int]map[common.Address]*big.Int) (err error) {
	ctx, span := m.metricsHandler.Tracer().Start(ctx, "prepareAndSubmitQuotes")
	defer func() {
		span.SetAttributes(attribute.Bool("relay_paused", m.relayPaused.Load()))
		metrics.EndSpanWithErr(span, err)
	}()

	var allQuotes []model.PutRelayerQuoteRequest

	// First, generate all quotes
	g, gctx := errgroup.WithContext(ctx)
	mtx := sync.Mutex{}
	for cid, balances := range inv {
		chainid := cid // capture loop variable
		for a, b := range balances {
			address := a
			balance := b
			g.Go(func() error {
				quotes, err := m.generateQuotes(gctx, chainid, address, balance, inv)
				if err != nil {
					return fmt.Errorf("error generating quotes: %w", err)
				}
				mtx.Lock()
				allQuotes = append(allQuotes, quotes...)
				mtx.Unlock()
				return nil
			})
		}
	}
	err = g.Wait()
	if err != nil {
		return fmt.Errorf("error generating quotes: %w", err)
	}

	span.SetAttributes(attribute.Int("num_quotes", len(allQuotes)))

	// Now, submit all the generated quotes
	if m.config.SubmitSingleQuotes {
		for _, quote := range allQuotes {
			if err := m.submitQuote(ctx, quote); err != nil {
				span.AddEvent("error submitting quote; setting relayPaused to true", trace.WithAttributes(
					attribute.String("error", err.Error()),
					attribute.Int(metrics.Origin, quote.OriginChainID),
					attribute.Int(metrics.Destination, quote.DestChainID),
					attribute.String("origin_token_addr", quote.OriginTokenAddr),
					attribute.String("dest_token_addr", quote.DestTokenAddr),
					attribute.String("max_origin_amount", quote.MaxOriginAmount),
					attribute.String("dest_amount", quote.DestAmount),
				))
				m.relayPaused.Store(true)

				// Suppress error so that we can continue submitting quotes
				return nil
			}
		}
	} else {
		err = m.submitBulkQuotes(ctx, allQuotes)
		if err != nil {
			span.AddEvent("error submitting bulk quotes; setting relayPaused to true", trace.WithAttributes(
				attribute.String("error", err.Error())))
			m.relayPaused.Store(true)
			return fmt.Errorf("error submitting bulk quotes: %w", err)
		}
	}

	// We successfully submitted all quotes, so we can set relayPaused to false
	m.relayPaused.Store(false)

	return nil
}

const meterName = "github.com/synapsecns/sanguine/services/rfq/relayer/quoter"

// Essentially, if we know a destination chain token balance, then we just need to find which tokens are bridgeable to it.
// We can do this by looking at the quotableTokens map, and finding the key that matches the destination chain token.
// Generates quotes for a given chain ID, address, and balance.
func (m *Manager) generateQuotes(parentCtx context.Context, chainID int, address common.Address, balance *big.Int, inv map[int]map[common.Address]*big.Int) (quotes []model.PutRelayerQuoteRequest, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "generateQuotes", trace.WithAttributes(
		attribute.Int(metrics.Origin, chainID),
		attribute.String("address", address.String()),
		attribute.String("balance", balance.String()),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	destRFQAddr, err := m.config.GetRFQAddress(chainID)
	if err != nil {
		return nil, fmt.Errorf("error getting destination RFQ address: %w", err)
	}
	destTokenID := fmt.Sprintf("%d-%s", chainID, address.Hex())

	// generate quotes in parallel
	g, gctx := errgroup.WithContext(ctx)
	quoteMtx := &sync.Mutex{}
	quotes = []model.PutRelayerQuoteRequest{}
	for k, itemTokenIDs := range m.quotableTokens {
		for _, tokenID := range itemTokenIDs {
			//nolint:nestif
			if tokenID == destTokenID {
				keyTokenID := k // Parse token info
				originStr := strings.Split(keyTokenID, "-")[0]
				origin, tokenErr := strconv.Atoi(originStr)
				if err != nil {
					span.AddEvent("error converting origin chainID", trace.WithAttributes(
						attribute.String("key_token_id", keyTokenID),
						attribute.String("error", tokenErr.Error()),
					))
					continue
				}
				originTokenAddr := common.HexToAddress(strings.Split(keyTokenID, "-")[1])

				var originBalance *big.Int
				originTokens, ok := inv[origin]
				if ok {
					originBalance = originTokens[originTokenAddr]
				}

				g.Go(func() error {
					input := QuoteInput{
						OriginChainID:   origin,
						DestChainID:     chainID,
						OriginTokenAddr: originTokenAddr,
						DestTokenAddr:   address,
						OriginBalance:   originBalance,
						DestBalance:     balance,
						DestRFQAddr:     destRFQAddr.Hex(),
					}

					quote, quoteErr := m.generateQuote(gctx, input)
					if quoteErr != nil {
						// continue generating quotes even if one fails
						span.AddEvent("error generating quote", trace.WithAttributes(
							attribute.String("key_token_id", keyTokenID),
							attribute.String("error", quoteErr.Error()),
						))
						return nil
					}
					quoteMtx.Lock()
					defer quoteMtx.Unlock()
					quotes = append(quotes, *quote)
					return nil
				})
			}
		}
	}
	err = g.Wait()
	if err != nil {
		return nil, fmt.Errorf("error generating quotes: %w", err)
	}

	m.currentQuotes = quotes
	return quotes, nil
}

// QuoteInput is a wrapper struct for input arguments to generateQuote.
type QuoteInput struct {
	OriginChainID   int
	DestChainID     int
	OriginTokenAddr common.Address
	DestTokenAddr   common.Address
	OriginBalance   *big.Int
	DestBalance     *big.Int
	DestRFQAddr     string
}

func (m *Manager) generateQuote(ctx context.Context, input QuoteInput) (quote *model.PutRelayerQuoteRequest, err error) {
	// Calculate the quote amount for this route
	originAmount, err := m.getOriginAmount(ctx, input)
	// don't quote if gas exceeds quote
	if errors.Is(err, errMinGasExceedsQuoteAmount) {
		originAmount = big.NewInt(0)
	} else if err != nil {
		logger.Error("Error getting quote amount", "error", err)
		return nil, err
	}

	// Calculate the fee for this route
	destToken, err := m.config.GetTokenName(uint32(input.DestChainID), input.DestTokenAddr.Hex())
	if err != nil {
		logger.Error("Error getting dest token ID", "error", err)
		return nil, fmt.Errorf("error getting dest token ID: %w", err)
	}
	fee, err := m.feePricer.GetTotalFee(ctx, uint32(input.OriginChainID), uint32(input.DestChainID), destToken, true)
	if err != nil {
		logger.Error("Error getting total fee", "error", err)
		return nil, fmt.Errorf("error getting total fee: %w", err)
	}
	originRFQAddr, err := m.config.GetRFQAddress(input.OriginChainID)
	if err != nil {
		logger.Error("Error getting RFQ address", "error", err)
		return nil, fmt.Errorf("error getting RFQ address: %w", err)
	}

	// Build the quote
	destAmount, err := m.getDestAmount(ctx, originAmount, input.DestChainID, destToken)
	if err != nil {
		logger.Error("Error getting dest amount", "error", err)
		return nil, fmt.Errorf("error getting dest amount: %w", err)
	}
	quote = &model.PutRelayerQuoteRequest{
		OriginChainID:           input.OriginChainID,
		OriginTokenAddr:         input.OriginTokenAddr.Hex(),
		DestChainID:             input.DestChainID,
		DestTokenAddr:           input.DestTokenAddr.Hex(),
		DestAmount:              destAmount.String(),
		MaxOriginAmount:         originAmount.String(),
		FixedFee:                fee.String(),
		OriginFastBridgeAddress: originRFQAddr.Hex(),
		DestFastBridgeAddress:   input.DestRFQAddr,
	}
	return quote, nil
}

// recordQuoteAmounts records the latest quotes from the relayer.
func (m *Manager) recordQuoteAmounts(_ context.Context, observer metric.Observer) (err error) {
	if m.meter == nil || m.quoteAmountGauge == nil || m.currentQuotes == nil {
		return nil
	}

	for _, quote := range m.currentQuotes {
		originMetadata, err := m.inventoryManager.GetTokenMetadata(quote.OriginChainID, common.HexToAddress(quote.OriginTokenAddr))
		if err != nil {
			return fmt.Errorf("error getting origin token metadata: %w", err)
		}
		destMetadata, err := m.inventoryManager.GetTokenMetadata(quote.DestChainID, common.HexToAddress(quote.DestTokenAddr))
		if err != nil {
			return fmt.Errorf("error getting dest token metadata: %w", err)
		}

		destAmount, err := strconv.ParseFloat(quote.DestAmount, 64)
		if err != nil {
			return fmt.Errorf("error parsing dest amount: %w", err)
		}
		opts := metric.WithAttributes(
			attribute.Int(metrics.Origin, quote.OriginChainID),
			attribute.Int(metrics.Destination, quote.DestChainID),
			attribute.String("origin_token_name", originMetadata.Name),
			attribute.String("dest_token_name", destMetadata.Name),
			attribute.String("max_origin_amount", quote.MaxOriginAmount),
			attribute.String("fixed_fee", quote.FixedFee),
			attribute.String("relayer", m.relayerSigner.Address().Hex()),
		)
		observer.ObserveFloat64(m.quoteAmountGauge, destAmount, opts)
	}

	return nil
}

// getOriginAmount calculates the origin quote amount for a given route.
//
//nolint:cyclop
func (m *Manager) getOriginAmount(parentCtx context.Context, input QuoteInput) (quoteAmount *big.Int, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "getOriginAmount", trace.WithAttributes(
		attribute.Int(metrics.Origin, input.OriginChainID),
		attribute.Int(metrics.Destination, input.DestChainID),
		attribute.String("dest_address", input.DestTokenAddr.String()),
		attribute.String("origin_address", input.OriginTokenAddr.String()),
		attribute.String("origin_balance", input.OriginBalance.String()),
		attribute.String("dest_balance", input.DestBalance.String()),
	))

	defer func() {
		span.SetAttributes(attribute.String("quote_amount", quoteAmount.String()))
		metrics.EndSpanWithErr(span, err)
	}()

	// First, check if we have enough gas to complete the a bridge for this route
	// If not, set the quote amount to zero to make sure a stale quote won't be used
	// TODO: handle in-flight gas; for now we can set a high min_gas_token
	sufficentGasOrigin, err := m.inventoryManager.HasSufficientGas(ctx, input.OriginChainID, nil)
	if err != nil {
		return nil, fmt.Errorf("error checking sufficient gas: %w", err)
	}
	sufficentGasDest, err := m.inventoryManager.HasSufficientGas(ctx, input.DestChainID, nil)
	if err != nil {
		return nil, fmt.Errorf("error checking sufficient gas: %w", err)
	}
	span.SetAttributes(
		attribute.Bool("sufficient_gas_origin", sufficentGasOrigin),
		attribute.Bool("sufficient_gas_dest", sufficentGasDest),
	)
	if !sufficentGasOrigin || !sufficentGasDest {
		return big.NewInt(0), nil
	}

	// Apply the quotePct
	quotePct, err := m.config.GetQuotePct(input.DestChainID)
	if err != nil {
		return nil, fmt.Errorf("error getting quote pct: %w", err)
	}
	balanceFlt := new(big.Float).SetInt(input.DestBalance)
	quoteAmount, _ = new(big.Float).Mul(balanceFlt, new(big.Float).SetFloat64(quotePct/100)).Int(nil)

	// Apply the quoteOffset to origin token.
	tokenName, err := m.config.GetTokenName(uint32(input.DestChainID), input.DestTokenAddr.Hex())
	if err != nil {
		return nil, fmt.Errorf("error getting token name: %w", err)
	}
	quoteOffsetBps, err := m.config.GetQuoteOffsetBps(input.OriginChainID, tokenName, true)
	if err != nil {
		return nil, fmt.Errorf("error getting quote offset bps: %w", err)
	}
	quoteAmount = m.applyOffset(ctx, quoteOffsetBps, quoteAmount)

	// Clip the quoteAmount by the minQuoteAmount
	minQuoteAmount := m.config.GetMinQuoteAmount(input.DestChainID, input.DestTokenAddr)
	if quoteAmount.Cmp(minQuoteAmount) < 0 {
		span.AddEvent("quote amount less than min quote amount", trace.WithAttributes(
			attribute.String("quote_amount", quoteAmount.String()),
			attribute.String("min_quote_amount", minQuoteAmount.String()),
		))
		quoteAmount = minQuoteAmount
	}

	// Clip the quoteAmount by the max origin balance
	maxBalance := m.config.GetMaxBalance(input.OriginChainID, input.OriginTokenAddr)
	if maxBalance != nil && input.OriginBalance != nil {
		quotableBalance := new(big.Int).Sub(maxBalance, input.OriginBalance)
		if quotableBalance.Cmp(big.NewInt(0)) <= 0 {
			span.AddEvent("non-positive quotable balance", trace.WithAttributes(
				attribute.String("quotable_balance", quotableBalance.String()),
				attribute.String("max_balance", maxBalance.String()),
				attribute.String("origin_balance", input.OriginBalance.String()),
			))
			quoteAmount = big.NewInt(0)
		} else if quoteAmount.Cmp(quotableBalance) > 0 {
			span.AddEvent("quote amount greater than quotable balance", trace.WithAttributes(
				attribute.String("quote_amount", quoteAmount.String()),
				attribute.String("quotable_balance", quotableBalance.String()),
				attribute.String("max_balance", maxBalance.String()),
				attribute.String("origin_balance", input.OriginBalance.String()),
			))
			quoteAmount = quotableBalance
		}
	}

	// Finally, clip the quoteAmount by the dest balance
	if quoteAmount.Cmp(input.DestBalance) > 0 {
		span.AddEvent("quote amount greater than destination balance", trace.WithAttributes(
			attribute.String("quote_amount", quoteAmount.String()),
			attribute.String("balance", input.DestBalance.String()),
		))
		quoteAmount = input.DestBalance
	}

	// Deduct gas cost from the quote amount, if necessary
	quoteAmount, err = m.deductGasCost(ctx, quoteAmount, input.DestTokenAddr, input.DestChainID)
	if err != nil {
		return nil, fmt.Errorf("error deducting gas cost: %w", err)
	}

	return quoteAmount, nil
}

// deductGasCost deducts the gas cost from the quote amount, if necessary.
func (m *Manager) deductGasCost(parentCtx context.Context, quoteAmount *big.Int, address common.Address, dest int) (quoteAmountAdj *big.Int, err error) {
	if !util.IsGasToken(address) {
		return quoteAmount, nil
	}

	_, span := m.metricsHandler.Tracer().Start(parentCtx, "deductGasCost", trace.WithAttributes(
		attribute.String("quote_amount", quoteAmount.String()),
	))
	defer func() {
		span.SetAttributes(attribute.String("quote_amount", quoteAmount.String()))
		metrics.EndSpanWithErr(span, err)
	}()

	// Deduct the minimum gas token balance from the quote amount
	var minGasToken *big.Int
	minGasToken, err = m.config.GetMinGasToken(dest)
	if err != nil {
		return nil, fmt.Errorf("error getting min gas token: %w", err)
	}
	quoteAmountAdj = new(big.Int).Sub(quoteAmount, minGasToken)
	if quoteAmountAdj.Cmp(big.NewInt(0)) < 0 {
		err = errMinGasExceedsQuoteAmount
		span.AddEvent(err.Error(), trace.WithAttributes(
			attribute.String("quote_amount_adj", quoteAmountAdj.String()),
			attribute.String("min_gas_token", minGasToken.String()),
		))
		return nil, err
	}
	return quoteAmountAdj, nil
}

var errMinGasExceedsQuoteAmount = errors.New("min gas token exceeds quote amount")

func (m *Manager) getDestAmount(parentCtx context.Context, originAmount *big.Int, chainID int, tokenName string) (*big.Int, error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "getDestAmount", trace.WithAttributes(
		attribute.String("quote_amount", originAmount.String()),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	quoteOffsetBps, err := m.config.GetQuoteOffsetBps(chainID, tokenName, false)
	if err != nil {
		return nil, fmt.Errorf("error getting quote offset bps: %w", err)
	}
	quoteWidthBps, err := m.config.GetQuoteWidthBps(chainID)
	if err != nil {
		return nil, fmt.Errorf("error getting quote width bps: %w", err)
	}
	totalOffsetBps := quoteOffsetBps + quoteWidthBps
	destAmount := m.applyOffset(ctx, totalOffsetBps, originAmount)

	span.SetAttributes(
		attribute.Float64("quote_offset_bps", quoteOffsetBps),
		attribute.Float64("quote_width_bps", quoteWidthBps),
		attribute.String("dest_amount", destAmount.String()),
	)
	return destAmount, nil
}

// applyOffset applies an offset (in bps) to a target.
func (m *Manager) applyOffset(parentCtx context.Context, offsetBps float64, target *big.Int) (result *big.Int) {
	_, span := m.metricsHandler.Tracer().Start(parentCtx, "applyOffset", trace.WithAttributes(
		attribute.Float64("offset_bps", offsetBps),
		attribute.String("target", target.String()),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	offsetFraction := new(big.Float).Quo(new(big.Float).SetInt64(int64(offsetBps)), new(big.Float).SetInt64(10000))
	offsetFactor := new(big.Float).Sub(new(big.Float).SetInt64(1), offsetFraction)
	result, _ = new(big.Float).Mul(new(big.Float).SetInt(target), offsetFactor).Int(nil)
	return result
}

// Submits a single quote.
func (m *Manager) submitQuote(ctx context.Context, quote model.PutRelayerQuoteRequest) error {
	quoteCtx, quoteCancel := context.WithTimeout(ctx, m.config.GetQuoteSubmissionTimeout())
	defer quoteCancel()

	err := m.rfqClient.PutQuote(quoteCtx, &quote)
	if err != nil {
		return fmt.Errorf("error submitting quote: %w", err)
	}
	return nil
}

// Submits multiple quotes.
func (m *Manager) submitBulkQuotes(ctx context.Context, quotes []model.PutRelayerQuoteRequest) error {
	quoteCtx, quoteCancel := context.WithTimeout(ctx, m.config.GetQuoteSubmissionTimeout())
	defer quoteCancel()

	req := model.PutBulkQuotesRequest{
		Quotes: quotes,
	}
	err := m.rfqClient.PutBulkQuotes(quoteCtx, &req)
	if err != nil {
		return fmt.Errorf("error submitting bulk quotes: %w", err)
	}
	return nil
}
