// Package pricer contains pricing logic for RFQ relayer quotes.
package pricer

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// FeePricer is the interface for the fee pricer.
type FeePricer interface {
	// Start starts the fee pricer.
	Start(ctx context.Context)
	// GetOriginFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetOriginFee(ctx context.Context, origin, destination uint32, denomToken string, isQuote bool) (*big.Int, error)
	// GetDestinationFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string, isQuote bool) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string, isQuote bool) (*big.Int, error)
	// GetGasPrice returns the gas price for a given chainID in native units.
	GetGasPrice(ctx context.Context, chainID uint32) (*big.Int, error)
	// GetTokenPrice returns the price of a token in USD.
	GetTokenPrice(ctx context.Context, token string) (float64, error)
	// GetPricePair calculates the price of a token pair from one chain to another.
	GetPricePair(parentCtx context.Context, stepLabel string, baseTokenChain uint32, pricedTokenChain uint32, baseToken string, pricedToken string, baseValueWei big.Int) (_ *PricedValuePair, err error)
}

type feePricer struct {
	// config is the relayer config.
	config relconfig.Config
	// gasPriceCache maps chainID -> gas price
	gasPriceCache *ttlcache.Cache[uint32, *big.Int]
	// tokenPriceCache maps token name -> token price
	tokenPriceCache *ttlcache.Cache[string, float64]
	// clientFetcher is used to fetch clients.
	clientFetcher submitter.ClientFetcher
	// handler is the metrics handler.
	handler metrics.Handler
	// priceFetcher is used to fetch prices from coingecko.
	priceFetcher CoingeckoPriceFetcher
}

// PricedValuePair is a set of equivalently priced values between two disparate assets.
// From.Wei/Units/Usd and To.Wei/Units/Usd all represent the exact same "value" - but in different denominations & formats.
type TokenValue struct {
	Symbol  string     `json:"symbol"`
	SpotUsd float64    `json:"spotUsd"`
	Wei     *big.Int   `json:"wei"`
	Units   *big.Float `json:"units"`
	Usd     *big.Float `json:"usd"`
}

type PricedValuePair struct {
	BaseToken   TokenValue `json:"baseTokenValue"`
	PricedToken TokenValue `json:"pricedTokenValue"`
}

// NewFeePricer creates a new fee pricer.
func NewFeePricer(config relconfig.Config, clientFetcher submitter.ClientFetcher, priceFetcher CoingeckoPriceFetcher, handler metrics.Handler) FeePricer {
	gasPriceCache := ttlcache.New[uint32, *big.Int](
		ttlcache.WithTTL[uint32, *big.Int](time.Second*time.Duration(config.GetFeePricer().GasPriceCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[uint32, *big.Int](),
	)
	tokenPriceCache := ttlcache.New[string, float64](
		ttlcache.WithTTL[string, float64](time.Second*time.Duration(config.GetFeePricer().TokenPriceCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[string, float64](),
	)
	return &feePricer{
		config:          config,
		gasPriceCache:   gasPriceCache,
		tokenPriceCache: tokenPriceCache,
		clientFetcher:   clientFetcher,
		handler:         handler,
		priceFetcher:    priceFetcher,
	}
}

func (f *feePricer) Start(ctx context.Context) {
	// Start the TTL caches.
	go f.gasPriceCache.Start()
	go f.tokenPriceCache.Start()

	go func() {
		<-ctx.Done()
		f.gasPriceCache.Stop()
		f.tokenPriceCache.Stop()
	}()
}

func (f *feePricer) GetOriginFee(parentCtx context.Context, origin, destination uint32, denomToken string, isQuote bool) (*big.Int, error) {
	var err error
	ctx, span := f.handler.Tracer().Start(parentCtx, "getOriginFee", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(origin)),
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("is_quote", isQuote),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Calculate the origin fee
	gasEstimate, err := f.config.GetOriginGasEstimate(int(origin))
	if err != nil {
		return nil, fmt.Errorf("could not get origin gas estimate: %w", err)
	}
	fee, err := f.getFee(ctx, origin, destination, gasEstimate, denomToken, isQuote)
	if err != nil {
		return nil, fmt.Errorf("err getFee: %w", err)
	}

	// If specified, calculate and add the L1 fee
	l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(origin, true)
	if useL1Fee {
		l1Fee, err := f.getFee(ctx, l1ChainID, destination, l1GasEstimate, denomToken, isQuote)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, l1Fee)
		span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
	}
	span.SetAttributes(attribute.String("origin_fee", fee.String()))
	return fee, nil
}

func (f *feePricer) GetDestinationFee(parentCtx context.Context, _, destination uint32, denomToken string, isQuote bool) (*big.Int, error) {
	var err error
	ctx, span := f.handler.Tracer().Start(parentCtx, "getDestinationFee", trace.WithAttributes(
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("is_quote", isQuote),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Calculate the destination fee
	gasEstimate, err := f.config.GetDestGasEstimate(int(destination))
	if err != nil {
		return nil, fmt.Errorf("could not get dest gas estimate: %w", err)
	}
	fee, err := f.getFee(ctx, destination, destination, gasEstimate, denomToken, isQuote)
	if err != nil {
		return nil, err
	}

	// If specified, calculate and add the L1 fee
	l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(destination, false)
	if useL1Fee {
		l1Fee, err := f.getFee(ctx, l1ChainID, destination, l1GasEstimate, denomToken, isQuote)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, l1Fee)
		span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
	}
	span.SetAttributes(attribute.String("destination_fee", fee.String()))
	return fee, nil
}

func (f *feePricer) GetTotalFee(parentCtx context.Context, origin, destination uint32, denomToken string, isQuote bool) (_ *big.Int, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getTotalFee", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(origin)),
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("is_quote", isQuote),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	originFee, err := f.GetOriginFee(ctx, origin, destination, denomToken, isQuote)
	if err != nil {
		span.AddEvent("could not get origin fee", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))
		return nil, err
	}
	destFee, err := f.GetDestinationFee(ctx, origin, destination, denomToken, isQuote)
	if err != nil {
		span.AddEvent("could not get destination fee", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))
		return nil, err
	}
	totalFee := new(big.Int).Add(originFee, destFee)
	span.SetAttributes(
		attribute.String("origin_fee", originFee.String()),
		attribute.String("dest_fee", destFee.String()),
		attribute.String("total_fee", totalFee.String()),
	)
	return totalFee, nil
}

func (f *feePricer) GetPricePair(parentCtx context.Context, stepLabel string, baseTokenChain uint32, pricedTokenChain uint32, baseToken string, pricedToken string, baseValueWei big.Int) (_ *PricedValuePair, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "GetPricePair", trace.WithAttributes(
		// stepLabel is an arbitrary & short sting to help provide debugging context to the logs & traces for this exact invocation of GetPricePair.
		attribute.String("step_label", stepLabel),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	var (
		baseTokenSpotUsd, pricedTokenSpotUsd               float64
		baseTokenDecimals, pricedTokenDecimals             uint8
		baseTokenDecimalsFactor, pricedTokenDecimalsFactor *big.Int
	)

	// TODO: Create a common GetTokenMetadata func that can be used for base & denom here -- and elsewhere.
	// ###################   baseToken lookups:
	if common.IsHexAddress(baseToken) {
		baseToken, err = f.config.GetTokenName(baseTokenChain, baseToken)
		if err != nil {
			return nil, err
		}
	}
	baseTokenSpotUsd, err = f.GetTokenPrice(ctx, baseToken)
	if err != nil {
		return nil, err
	}
	if baseTokenSpotUsd <= 0 {
		return nil, fmt.Errorf("invalid baseTokenSpotUsd: %f", baseTokenSpotUsd)
	}

	baseTokenDecimals, err = f.config.GetTokenDecimals(baseTokenChain, baseToken)
	if err != nil {
		return nil, err
	}
	baseTokenDecimalsFactor = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(baseTokenDecimals)), nil)

	// TODO: Create a common GetTokenMetadata func that can be used for base & denom here -- and elsewhere.
	// ###################   pricedToken lookups:
	if common.IsHexAddress(pricedToken) {
		pricedToken, err = f.config.GetTokenName(pricedTokenChain, pricedToken)
		if err != nil {
			return nil, err
		}
	}
	pricedTokenSpotUsd, err = f.GetTokenPrice(ctx, pricedToken)
	if err != nil {
		return nil, err
	}

	if pricedTokenSpotUsd <= 0 {
		return nil, fmt.Errorf("invalid pricedTokenSpotUsd: %f", pricedTokenSpotUsd)
	}

	pricedTokenDecimals, err = f.config.GetTokenDecimals(pricedTokenChain, pricedToken)
	if err != nil {
		return nil, err
	}
	pricedTokenDecimalsFactor = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(pricedTokenDecimals)), nil)

	// With all token/decimal/price lookups resolved, construct a concicse pair label to aid with debug/log/trace context
	// EG: ETH.42161>USDC.8453
	//
	// In addition, for "USD" (from special USD_ constant) -- do not print the chain component since it will just be a confusing placeholder
	// EG: ETH.42161>USD
	pairLabel := fmt.Sprintf("%s>%s", strings.Replace(fmt.Sprintf("%s.%d", baseToken, baseTokenChain), "USD."+fmt.Sprint(baseTokenChain), "USD", 1), strings.Replace(fmt.Sprintf("%s.%d", pricedToken, pricedTokenChain), "USD."+fmt.Sprint(pricedTokenChain), "USD", 1))

	span.SetAttributes(attribute.String("pair_label", pairLabel))

	// if zero basevaluewei is supplied, early return a zeroed-out result.
	if baseValueWei.Cmp(big.NewInt(0)) == 0 {
		zeroResult := &PricedValuePair{
			BaseToken: TokenValue{
				Symbol:  baseToken,
				SpotUsd: baseTokenSpotUsd,
				Wei:     big.NewInt(0),
				Units:   big.NewFloat(0),
				Usd:     big.NewFloat(0),
			},
			PricedToken: TokenValue{
				Symbol:  pricedToken,
				SpotUsd: pricedTokenSpotUsd,
				Wei:     big.NewInt(0),
				Units:   big.NewFloat(0),
				Usd:     big.NewFloat(0),
			},
		}
		return zeroResult, nil
	}

	// This will be the final wei output
	var pricedValueWei big.Int

	var result PricedValuePair

	// The steps below convert a raw/wei value of our baseToken (baseValueWei) into an equivalently priced raw/wei value of pricedToken (pricedValueWei)

	// 1) convert baseValueWei into baseValue Units  --- EG: 1000000000000000000 ETHER >> 1.0 ETHER
	baseValueUnits := new(big.Float).Quo(new(big.Float).SetInt(&baseValueWei), new(big.Float).SetInt(baseTokenDecimalsFactor))

	// 2) convert baseValueUnits into baseValueUSD --- EG: 1.0 ETHER >> $2555.55 USD
	baseValueUsd := new(big.Float).Mul(baseValueUnits, new(big.Float).SetFloat64(baseTokenSpotUsd))

	pricedValueUsd := baseValueUsd // <- Intentional. USD value of baseValue is identical to pricedValue and is the common anchor point between the pairs being priced.

	var pricedValueUnits *big.Float

	// Tokens are the same & identical decimals? Then steps 3 and 4 can be bypassed. Instead, pricedValueUnits and pricedValueWei will be the same as their baseValue counterparts
	if baseToken == pricedToken && baseTokenDecimals == pricedTokenDecimals {
		pricedValueUnits = baseValueUnits
		pricedValueWei.Set(&baseValueWei)
	} else {
		// 3) convert pricedValueUsd into pricedValueUnits --- EG: $2555.55 USD to 2555.55 USDC
		pricedValueUnits = new(big.Float).Quo(pricedValueUsd, new(big.Float).SetFloat64(pricedTokenSpotUsd))

		// 4) convert pricedValueUnits into pricedValueWeiFloat, then to Int, then to final output bigint --- EG: 2555.55 USDC to 2555550000
		pricedValueWeiFloat := new(big.Float).Mul(pricedValueUnits, new(big.Float).SetInt(pricedTokenDecimalsFactor))
		pricedValueWeiInt, _ := pricedValueWeiFloat.Int(nil)
		pricedValueWei.Set(pricedValueWeiInt)
	}

	// add "pricePair" to debugOutput env var for dev/debug output
	if strings.Contains(strings.ToLower(os.Getenv("debugOutput")), "pricepair") {
		fmt.Printf("%-25s%-20s%-5s base_____wei: %s\n", stepLabel, pairLabel, baseToken, baseValueWei.String())
		fmt.Printf("%-25s%-20s%-5s base___units: %s\n", stepLabel, pairLabel, baseToken, baseValueUnits.Text('f', -1))
		fmt.Printf("%-25s%-20s%-5s base_____usd: %s\n", stepLabel, pairLabel, baseToken, baseValueUsd.Text('f', -1))
		fmt.Printf("%-25s%-20s%-5s priced___usd: %s\n", stepLabel, pairLabel, pricedToken, pricedValueUsd.Text('f', -1))
		fmt.Printf("%-25s%-20s%-5s priced_units: %s\n", stepLabel, pairLabel, pricedToken, pricedValueUnits.Text('f', -1))
		fmt.Printf("%-25s%-20s%-5s priced___wei: %s\n", stepLabel, pairLabel, pricedToken, pricedValueWei.String())
	}

	span.SetAttributes(
		attribute.String("base_token_symbol", baseToken),
		attribute.Int("base_token_decimals", int(baseTokenDecimals)),
		attribute.Int("base_token_chain", int(baseTokenChain)),
		attribute.String("base_token_wei", baseValueWei.String()),
		attribute.String("base_token_units", baseValueUnits.Text('f', -1)),
		attribute.String("base_token_usd", baseValueUsd.Text('f', -1)),

		attribute.String("priced_token_symbol", pricedToken),
		attribute.Int("priced_token_decimals", int(pricedTokenDecimals)),
		attribute.Int("priced_token_chain", int(pricedTokenChain)),
		attribute.String("priced_token_usd", pricedValueUsd.Text('f', -1)),
		attribute.String("priced_token_units", pricedValueUnits.Text('f', -1)),
		attribute.String("priced_token_wei", pricedValueWei.String()),
	)

	// all of these represent the exact same value, but in different denominations & formats.
	result = PricedValuePair{
		BaseToken: TokenValue{
			Symbol:  baseToken,
			SpotUsd: baseTokenSpotUsd,
			Wei:     new(big.Int).Set(&baseValueWei),
			Units:   new(big.Float).Set(baseValueUnits),
			Usd:     new(big.Float).Set(baseValueUsd),
		},
		PricedToken: TokenValue{
			Symbol:  pricedToken,
			SpotUsd: pricedTokenSpotUsd,
			Wei:     new(big.Int).Set(&pricedValueWei),
			Units:   new(big.Float).Set(pricedValueUnits),
			Usd:     new(big.Float).Set(pricedValueUsd),
		},
	}

	return &result, nil

}

func (f *feePricer) getFee(parentCtx context.Context, gasChain, denomChain uint32, gasEstimate int, denomToken string, isQuote bool) (_ *big.Int, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getFee", trace.WithAttributes(
		attribute.Int("gas_chain", int(gasChain)),
		attribute.Int("denom_chain", int(denomChain)),
		attribute.Int("gas_estimate", gasEstimate),
		attribute.String("denom_token", denomToken),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	gasPrice, err := f.GetGasPrice(ctx, gasChain)
	if err != nil {
		return nil, err
	}
	nativeToken, err := f.config.GetNativeToken(int(gasChain))
	if err != nil {
		return nil, err
	}

	// calculate the total Fee cost in native gas WEI
	feeNativeWei := new(big.Int).Mul(gasPrice, big.NewInt(int64(gasEstimate)))

	// price native gas WEI into the denomination token
	feeNativeWeiPriced, err := f.GetPricePair(ctx, "getFee", gasChain, denomChain, nativeToken, denomToken, *feeNativeWei)
	if err != nil {
		return nil, err
	}

	var feeDenom = new(big.Float).SetInt(feeNativeWeiPriced.PricedToken.Wei)

	var multiplier float64
	if isQuote {
		multiplier, err = f.config.GetQuoteFixedFeeMultiplier(int(gasChain))
		if err != nil {
			return nil, fmt.Errorf("could not get quote fixed fee multiplier: %w", err)
		}
	} else {
		multiplier, err = f.config.GetRelayFixedFeeMultiplier(int(gasChain))
		if err != nil {
			return nil, fmt.Errorf("could not get relay fixed fee multiplier: %w", err)
		}
	}

	// Apply the fixed fee multiplier.
	// Note that this step rounds towards zero- we may need to apply rounding here if
	// we want to be conservative and lean towards overestimating fees.
	feeDenomScaled, _ := new(big.Float).Mul(feeDenom, new(big.Float).SetFloat64(multiplier)).Int(nil)

	if feeDenomScaled == nil {
		return nil, fmt.Errorf("err getFee: nil fee return")
	}

	return feeDenomScaled, nil
}

// getGasPrice returns the gas price for a given chainID in native units.
func (f *feePricer) GetGasPrice(ctx context.Context, chainID uint32) (*big.Int, error) {
	// Attempt to fetch gas price from cache.
	gasPriceItem := f.gasPriceCache.Get(chainID)
	var gasPrice *big.Int
	if gasPriceItem == nil {
		// Fetch gas price from omnirpc.
		client, err := f.clientFetcher.GetClient(ctx, big.NewInt(int64(chainID)))
		if err != nil {
			return nil, err
		}
		gasPrice, err = client.SuggestGasPrice(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to suggest gas price on chain %d: %w", chainID, err)
		}
		if gasPrice == nil {
			return nil, fmt.Errorf("gas price is nil on chain %d", chainID)
		}
		f.gasPriceCache.Set(chainID, gasPrice, 0)
	} else {
		gasPrice = gasPriceItem.Value()
	}
	return gasPrice, nil
}

// getTokenPrice returns the price of a token in USD.
func (f *feePricer) GetTokenPrice(ctx context.Context, token string) (price float64, err error) {
	ctx, span := f.handler.Tracer().Start(ctx, "GetTokenPrice", trace.WithAttributes(
		attribute.String("token", token),
	))

	defer func() {
		span.SetAttributes(attribute.Float64("price", price))
		metrics.EndSpanWithErr(span, err)
	}()

	// Attempt to fetch gas price from cache.
	tokenPriceItem := f.tokenPriceCache.Get(token)
	//nolint:nestif
	if tokenPriceItem == nil {
		// Try to get price from coingecko.
		price, err = f.priceFetcher.GetPrice(ctx, token)

		if err == nil {
			f.tokenPriceCache.Set(token, price, 0)
			span.SetAttributes(attribute.Float64("cg_price", price))
		} else {
			span.SetAttributes(
				attribute.String("cg_error", err.Error()),
			)
			// intentionally no longer allow fallback to flat hard-configured token price
			return 0, fmt.Errorf("err price lookup %s: %v", token, err)

		}
	} else {
		price = tokenPriceItem.Value()
		span.SetAttributes(attribute.Float64("cache_price", price))
	}
	return price, nil
}
