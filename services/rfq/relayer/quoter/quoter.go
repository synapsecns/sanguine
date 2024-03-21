// Package quoter submits quotes to the RFQ API for which assets the relayer is willing to relay.
package quoter

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/synapsecns/sanguine/contrib/screener-api/client"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/pricer"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	"golang.org/x/exp/slices"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/ethergo/signer/signer"
	rfqAPIClient "github.com/synapsecns/sanguine/services/rfq/api/client"
	"github.com/synapsecns/sanguine/services/rfq/api/model"
	"github.com/synapsecns/sanguine/services/rfq/relayer/inventory"
)

var logger = log.Logger("quoter")

// Quoter submits quotes to the RFQ API.
type Quoter interface {
	// SubmitAllQuotes submits all quotes to the RFQ API.
	SubmitAllQuotes(ctx context.Context) (err error)
	// ShouldProcess determines if a quote should be processed.
	// We do this by either saving all quotes in-memory, and refreshing via GetSelfQuotes() through the API
	// The first comparison is does bridge transaction OriginChainID+TokenAddr match with a quote + DestChainID+DestTokenAddr, then we look to see if we have enough amount to relay it + if the price fits our bounds (based on that the Relayer is relaying the destination token for the origin)
	// validateQuote(BridgeEvent)
	ShouldProcess(ctx context.Context, quote reldb.QuoteRequest) (bool, error)
	// IsProfitable determines if a quote is profitable, i.e. we will not lose money on it, net of fees.
	IsProfitable(ctx context.Context, quote reldb.QuoteRequest) (bool, error)
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
	// simpleScreener is used to screen addresses.
	screener client.ScreenerClient
	// relayPaused is set when the RFQ API is found to be offline, which
	// lets the quoter indicate that quotes should not be relayed.
	relayPaused bool
}

// NewQuoterManager creates a new QuoterManager.
func NewQuoterManager(config relconfig.Config, metricsHandler metrics.Handler, inventoryManager inventory.Manager, relayerSigner signer.Signer, feePricer pricer.FeePricer) (Quoter, error) {
	apiClient, err := rfqAPIClient.NewAuthenticatedClient(metricsHandler, config.GetRfqAPIURL(), relayerSigner)
	if err != nil {
		return nil, fmt.Errorf("error creating RFQ API client: %w", err)
	}

	qt := make(map[string][]string)

	// fix any casing issues.
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

	return &Manager{
		config:           config,
		inventoryManager: inventoryManager,
		rfqClient:        apiClient,
		quotableTokens:   qt,
		relayerSigner:    relayerSigner,
		metricsHandler:   metricsHandler,
		feePricer:        feePricer,
		screener:         ss,
	}, nil
}

const screenerRuleset = "rfq"

// ShouldProcess determines if a quote should be processed.
func (m *Manager) ShouldProcess(parentCtx context.Context, quote reldb.QuoteRequest) (res bool, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "shouldProcess", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(quote.TransactionID[:])),
	))

	defer func() {
		span.AddEvent("result", trace.WithAttributes(attribute.Bool("result", res)))
		metrics.EndSpanWithErr(span, err)
	}()

	if m.relayPaused {
		span.AddEvent("relayPaused is set due to RFQ API being offline")
		return false, nil
	}

	if m.screener != nil {
		blocked, err := m.screener.ScreenAddress(ctx, screenerRuleset, quote.Transaction.OriginSender.String())
		if err != nil {
			span.RecordError(fmt.Errorf("error screening address: %w", err))
			return false, fmt.Errorf("error screening address: %w", err)
		}
		if blocked {
			span.AddEvent(fmt.Sprintf("address %s blocked", quote.Transaction.OriginSender))
			return false, nil
		}

		blocked, err = m.screener.ScreenAddress(ctx, screenerRuleset, quote.Transaction.DestRecipient.String())
		if err != nil {
			span.RecordError(fmt.Errorf("error screening address: %w", err))
			return false, fmt.Errorf("error screening address: %w", err)
		}
		if blocked {
			span.AddEvent(fmt.Sprintf("address %s blocked", quote.Transaction.DestRecipient))
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

	inv, err := m.inventoryManager.GetCommittableBalances(ctx)
	if err != nil {
		return fmt.Errorf("error getting committable balances: %w", err)
	}
	return m.prepareAndSubmitQuotes(ctx, inv)
}

// Prepares and submits quotes based on inventory.
func (m *Manager) prepareAndSubmitQuotes(ctx context.Context, inv map[int]map[common.Address]*big.Int) (err error) {
	ctx, span := m.metricsHandler.Tracer().Start(ctx, "prepareAndSubmitQuotes")
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var allQuotes []model.PutQuoteRequest

	// First, generate all quotes
	for chainID, balances := range inv {
		for address, balance := range balances {
			quotes, err := m.generateQuotes(ctx, chainID, address, balance)
			if err != nil {
				return err
			}
			allQuotes = append(allQuotes, quotes...)
		}
	}

	// Now, submit all the generated quotes
	for _, quote := range allQuotes {
		if err := m.submitQuote(quote); err != nil {
			span.AddEvent("error submitting quote; setting relayPaused to true", trace.WithAttributes(
				attribute.String("error", err.Error()),
				attribute.Int(metrics.Origin, quote.OriginChainID),
				attribute.Int(metrics.Destination, quote.DestChainID),
				attribute.String("origin_token_addr", quote.OriginTokenAddr),
				attribute.String("dest_token_addr", quote.DestTokenAddr),
				attribute.String("max_origin_amount", quote.MaxOriginAmount),
				attribute.String("dest_amount", quote.DestAmount),
			))
			m.relayPaused = true

			// Suppress error so that we can continue submitting quotes
			return nil
		}
	}

	// We successfully submitted all quotes, so we can set relayPaused to false
	m.relayPaused = false

	return nil
}

// generateQuotes TODO: THIS LOOP IS BROKEN
// Essentially, if we know a destination chain token balance, then we just need to find which tokens are bridgeable to it.
// We can do this by looking at the quotableTokens map, and finding the key that matches the destination chain token.
// Generates quotes for a given chain ID, address, and balance.
func (m *Manager) generateQuotes(parentCtx context.Context, chainID int, address common.Address, balance *big.Int) (quotes []model.PutQuoteRequest, err error) {
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
	quotes = []model.PutQuoteRequest{}
	for keyTokenID, itemTokenIDs := range m.quotableTokens {
		for _, tokenID := range itemTokenIDs {
			//nolint:nestif
			if tokenID == destTokenID {
				quote, quoteErr := m.generateQuote(ctx, keyTokenID, chainID, address, balance, destRFQAddr)
				if quoteErr != nil {
					// continue generating quotes even if one fails
					span.AddEvent("error generating quote", trace.WithAttributes(
						attribute.String("key_token_id", keyTokenID),
					))
					continue
				}
				quotes = append(quotes, *quote)
			}
		}
	}
	return quotes, nil
}

func (m *Manager) generateQuote(ctx context.Context, keyTokenID string, chainID int, address common.Address, balance *big.Int, destRFQAddr string) (quote *model.PutQuoteRequest, err error) {
	// Parse token info
	originStr := strings.Split(keyTokenID, "-")[0]
	origin, err := strconv.Atoi(originStr)
	if err != nil {
		logger.Error("Error converting origin chainID", "error", err)
		return nil, fmt.Errorf("error converting origin chainID: %w", err)
	}
	originTokenAddr := common.HexToAddress(strings.Split(keyTokenID, "-")[1])

	// Calculate the quote amount for this route
	quoteAmount, err := m.getQuoteAmount(ctx, origin, chainID, address, balance)
	// don't quote if gas exceeds quote
	if errors.Is(err, errMinGasExceedsQuoteAmount) {
		quoteAmount = big.NewInt(0)
	} else if err != nil {
		logger.Error("Error getting quote amount", "error", err)
		return nil, err
	}

	// Calculate the fee for this route
	destToken, err := m.config.GetTokenName(uint32(chainID), address.Hex())
	if err != nil {
		logger.Error("Error getting dest token ID", "error", err)
		return nil, fmt.Errorf("error getting dest token ID: %w", err)
	}
	fee, err := m.feePricer.GetTotalFee(ctx, uint32(origin), uint32(chainID), destToken, true)
	if err != nil {
		logger.Error("Error getting total fee", "error", err)
		return nil, fmt.Errorf("error getting total fee: %w", err)
	}
	originRFQAddr, err := m.config.GetRFQAddress(origin)
	if err != nil {
		logger.Error("Error getting RFQ address", "error", err)
		return nil, fmt.Errorf("error getting RFQ address: %w", err)
	}

	// Build the quote
	destAmount, err := m.getDestAmount(ctx, quoteAmount, chainID)
	if err != nil {
		logger.Error("Error getting dest amount", "error", err)
		return nil, fmt.Errorf("error getting dest amount: %w", err)
	}
	quote = &model.PutQuoteRequest{
		OriginChainID:           origin,
		OriginTokenAddr:         originTokenAddr.Hex(),
		DestChainID:             chainID,
		DestTokenAddr:           address.Hex(),
		DestAmount:              destAmount.String(),
		MaxOriginAmount:         quoteAmount.String(),
		FixedFee:                fee.String(),
		OriginFastBridgeAddress: originRFQAddr,
		DestFastBridgeAddress:   destRFQAddr,
	}
	return quote, nil
}

// getQuoteAmount calculates the quote amount for a given route.
func (m *Manager) getQuoteAmount(parentCtx context.Context, origin, dest int, address common.Address, balance *big.Int) (quoteAmount *big.Int, err error) {
	ctx, span := m.metricsHandler.Tracer().Start(parentCtx, "getQuoteAmount", trace.WithAttributes(
		attribute.String(metrics.Origin, strconv.Itoa(origin)),
		attribute.String(metrics.Destination, strconv.Itoa(dest)),
		attribute.String("address", address.String()),
		attribute.String("balance", balance.String()),
	))

	defer func() {
		span.SetAttributes(attribute.String("quote_amount", quoteAmount.String()))
		metrics.EndSpanWithErr(span, err)
	}()

	// First, check if we have enough gas to complete the a bridge for this route
	// If not, set the quote amount to zero to make sure a stale quote won't be used
	// TODO: handle in-flight gas; for now we can set a high min_gas_token
	sufficentGas, err := m.inventoryManager.HasSufficientGas(ctx, origin, dest)
	if err != nil {
		return nil, fmt.Errorf("error checking sufficient gas: %w", err)
	}
	if !sufficentGas {
		return big.NewInt(0), nil
	}

	// Apply the quotePct
	quotePct, err := m.config.GetQuotePct(dest)
	if err != nil {
		return nil, fmt.Errorf("error getting quote pct: %w", err)
	}
	balanceFlt := new(big.Float).SetInt(balance)
	quoteAmount, _ = new(big.Float).Mul(balanceFlt, new(big.Float).SetFloat64(quotePct/100)).Int(nil)

	// Clip the quoteAmount by the minQuoteAmount
	minQuoteAmount := m.config.GetMinQuoteAmount(dest, address)
	if quoteAmount.Cmp(minQuoteAmount) < 0 {
		span.AddEvent("quote amount less than min quote amount", trace.WithAttributes(
			attribute.String("quote_amount", quoteAmount.String()),
			attribute.String("min_quote_amount", minQuoteAmount.String()),
		))
		quoteAmount = minQuoteAmount
	}

	// Finally, clip the quoteAmount by the balance
	if quoteAmount.Cmp(balance) > 0 {
		span.AddEvent("quote amount greater than balance", trace.WithAttributes(
			attribute.String("quote_amount", quoteAmount.String()),
			attribute.String("balance", balance.String()),
		))
		quoteAmount = balance
	}

	// Deduct gas cost from the quote amount, if necessary
	if chain.IsGasToken(address) {
		// Deduct the minimum gas token balance from the quote amount
		var minGasToken *big.Int
		minGasToken, err = m.config.GetMinGasToken(dest)
		if err != nil {
			return nil, fmt.Errorf("error getting min gas token: %w", err)
		}
		quoteAmount = new(big.Int).Sub(quoteAmount, minGasToken)
		if quoteAmount.Cmp(big.NewInt(0)) < 0 {
			err = errMinGasExceedsQuoteAmount
			span.AddEvent(err.Error(), trace.WithAttributes(
				attribute.String("quote_amount", quoteAmount.String()),
				attribute.String("min_gas_token", minGasToken.String()),
			))
			return nil, err
		}
	}
	return quoteAmount, nil
}

var errMinGasExceedsQuoteAmount = errors.New("min gas token exceeds quote amount")

func (m *Manager) getDestAmount(parentCtx context.Context, quoteAmount *big.Int, chainID int) (*big.Int, error) {
	_, span := m.metricsHandler.Tracer().Start(parentCtx, "getDestAmount", trace.WithAttributes(
		attribute.String("quote_amount", quoteAmount.String()),
	))
	defer func() {
		metrics.EndSpan(span)
	}()

	quoteOffsetBps, err := m.config.GetQuoteOffsetBps(chainID)
	if err != nil {
		return nil, fmt.Errorf("error getting quote offset bps: %w", err)
	}
	quoteOffsetFraction := new(big.Float).Quo(new(big.Float).SetInt64(int64(quoteOffsetBps)), new(big.Float).SetInt64(10000))
	quoteOffsetFactor := new(big.Float).Sub(new(big.Float).SetInt64(1), quoteOffsetFraction)
	destAmount, _ := new(big.Float).Mul(new(big.Float).SetInt(quoteAmount), quoteOffsetFactor).Int(nil)

	span.SetAttributes(
		attribute.Float64("quote_offset_bps", quoteOffsetBps),
		attribute.String("quote_offset_fraction", quoteOffsetFraction.String()),
		attribute.String("quote_offset_factor", quoteOffsetFactor.String()),
		attribute.String("dest_amount", destAmount.String()),
	)
	return destAmount, nil
}

// Submits a single quote.
func (m *Manager) submitQuote(quote model.PutQuoteRequest) error {
	err := m.rfqClient.PutQuote(&quote)
	if err != nil {
		return fmt.Errorf("error submitting quote: %w", err)
	}
	return nil
}
