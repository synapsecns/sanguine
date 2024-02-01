package pricer

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/sync/errgroup"
)

// FeePricer is the interface for the fee pricer.
type FeePricer interface {
	// Start starts the fee pricer.
	Start(ctx context.Context)
	// GetOriginFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetOriginFee(ctx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error)
	// GetDestinationFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error)
	// GetGasPrice returns the gas price for a given chainID in native units.
	GetGasPrice(ctx context.Context, chainID uint32) (*big.Int, error)
}

type feePricer struct {
	// config is the relayer config.
	config relconfig.Config
	// gasPriceCache maps chainID -> gas price
	gasPriceCache *ttlcache.Cache[uint32, *big.Int]
	// tokenPriceCache maps token name -> token price
	tokenPriceCache *ttlcache.Cache[string, *big.Int]
	// clientFetcher is used to fetch clients.
	clientFetcher submitter.ClientFetcher
	// handler is the metrics handler.
	handler metrics.Handler
}

// NewFeePricer creates a new fee pricer.
func NewFeePricer(config relconfig.Config, clientFetcher submitter.ClientFetcher, handler metrics.Handler) FeePricer {
	gasPriceCache := ttlcache.New[uint32, *big.Int](
		ttlcache.WithTTL[uint32, *big.Int](time.Second*time.Duration(config.GetFeePricer().GasPriceCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[uint32, *big.Int](),
	)
	return &feePricer{
		config:          config,
		gasPriceCache:   gasPriceCache,
		tokenPriceCache: ttlcache.New[string, *big.Int](ttlcache.WithTTL[string, *big.Int](time.Second * time.Duration(config.GetFeePricer().TokenPriceCacheTTLSeconds))),
		clientFetcher:   clientFetcher,
		handler:         handler,
	}
}

func (f *feePricer) Start(ctx context.Context) {
	g, _ := errgroup.WithContext(ctx)

	// Start the TTL caches.
	g.Go(func() error {
		f.gasPriceCache.Start()
		return nil
	})
	g.Go(func() error {
		f.tokenPriceCache.Start()
		return nil
	})
}

var nativeDecimalsFactor = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18)), nil)

func (f *feePricer) GetOriginFee(parentCtx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error) {
	var err error
	ctx, span := f.handler.Tracer().Start(parentCtx, "getOriginFee", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(origin)),
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("use_multiplier", useMultiplier),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Calculate the origin fee
	fee, err := f.getFee(ctx, origin, destination, f.config.GetOriginGasEstimate(origin), denomToken, useMultiplier)
	if err != nil {
		return nil, err
	}

	// If specified, calculate and add the L1 fee
	l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(origin, true)
	if useL1Fee {
		l1Fee, err := f.getFee(ctx, l1ChainID, destination, l1GasEstimate, denomToken, useMultiplier)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, l1Fee)
		span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
	}
	span.SetAttributes(attribute.String("origin_fee", fee.String()))
	return fee, nil
}

func (f *feePricer) GetDestinationFee(parentCtx context.Context, _, destination uint32, denomToken string, useMultiplier bool) (*big.Int, error) {
	var err error
	ctx, span := f.handler.Tracer().Start(parentCtx, "getDestinationFee", trace.WithAttributes(
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("use_multiplier", useMultiplier),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	// Calculate the destination fee
	fee, err := f.getFee(ctx, destination, destination, f.config.GetDestGasEstimate(destination), denomToken, useMultiplier)
	if err != nil {
		return nil, err
	}

	// If specified, calculate and add the L1 fee
	l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(destination, false)
	if useL1Fee {
		l1Fee, err := f.getFee(ctx, l1ChainID, destination, l1GasEstimate, denomToken, useMultiplier)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, l1Fee)
		span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
	}
	span.SetAttributes(attribute.String("destination_fee", fee.String()))
	return fee, nil
}

func (f *feePricer) GetTotalFee(parentCtx context.Context, origin, destination uint32, denomToken string, useMultiplier bool) (_ *big.Int, err error) {
	ctx, span := f.handler.Tracer().Start(parentCtx, "getTotalFee", trace.WithAttributes(
		attribute.Int(metrics.Origin, int(origin)),
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("use_multiplier", useMultiplier),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	originFee, err := f.GetOriginFee(ctx, origin, destination, denomToken, useMultiplier)
	if err != nil {
		span.AddEvent("could not get origin fee", trace.WithAttributes(
			attribute.String("error", err.Error()),
		))
		return nil, err
	}
	destFee, err := f.GetDestinationFee(ctx, origin, destination, denomToken, useMultiplier)
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

func (f *feePricer) getFee(parentCtx context.Context, gasChain, denomChain uint32, gasEstimate int, denomToken string, useMultiplier bool) (_ *big.Int, err error) {
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
	nativeToken, err := f.config.GetNativeToken(gasChain)
	if err != nil {
		return nil, err
	}
	nativeTokenPrice, err := f.getTokenPrice(ctx, nativeToken)
	if err != nil {
		return nil, err
	}
	denomTokenPrice, err := f.getTokenPrice(ctx, denomToken)
	if err != nil {
		return nil, err
	}
	denomTokenDecimals, err := f.config.GetTokenDecimals(denomChain, denomToken)
	if err != nil {
		return nil, err
	}
	denomDecimalsFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(denomTokenDecimals)), nil)

	// Compute the fee.
	var feeDenom *big.Float
	feeWei := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), new(big.Float).SetFloat64(float64(gasEstimate)))
	if denomToken == nativeToken {
		// Denomination token is native token, so no need for unit conversion.
		feeDenom = feeWei
	} else {
		// Convert the fee from ETH to denomToken terms.
		feeEth := new(big.Float).Quo(feeWei, new(big.Float).SetInt(nativeDecimalsFactor))
		feeUSD := new(big.Float).Mul(feeEth, new(big.Float).SetFloat64(nativeTokenPrice))
		feeUSDC := new(big.Float).Mul(feeUSD, new(big.Float).SetFloat64(denomTokenPrice))
		feeDenom = new(big.Float).Mul(feeUSDC, new(big.Float).SetInt(denomDecimalsFactor))
		span.SetAttributes(
			attribute.String("fee_wei", feeWei.String()),
			attribute.String("fee_eth", feeEth.String()),
			attribute.String("fee_usd", feeUSD.String()),
			attribute.String("fee_usdc", feeUSDC.String()),
		)
	}

	multiplier := f.config.GetFixedFeeMultiplier()
	if !useMultiplier {
		multiplier = 1
	}

	// Apply the fixed fee multiplier.
	// Note that this step rounds towards zero- we may need to apply rounding here if
	// we want to be conservative and lean towards overestimating fees.
	feeUSDCDecimalsScaled, _ := new(big.Float).Mul(feeDenom, new(big.Float).SetFloat64(multiplier)).Int(nil)
	span.SetAttributes(
		attribute.String("gas_price", gasPrice.String()),
		attribute.Float64("native_token_price", nativeTokenPrice),
		attribute.Float64("denom_token_price", denomTokenPrice),
		attribute.Int("denom_token_decimals", int(denomTokenDecimals)),
		attribute.String("fee_wei", feeWei.String()),
		attribute.String("fee_denom", feeDenom.String()),
		attribute.String("fee_usdc_decimals_scaled", feeUSDCDecimalsScaled.String()),
	)
	return feeUSDCDecimalsScaled, nil
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
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}
		gasPrice = header.BaseFee
		f.gasPriceCache.Set(chainID, gasPrice, 0)
	} else {
		gasPrice = gasPriceItem.Value()
	}
	return gasPrice, nil
}

// getTokenPrice returns the price of a token in USD.
func (f *feePricer) getTokenPrice(ctx context.Context, token string) (float64, error) {
	for _, chainConfig := range f.config.GetChains() {
		for tokenName, tokenConfig := range chainConfig.Tokens {
			if token == tokenName {
				return tokenConfig.PriceUSD, nil
			}
		}
	}
	return 0, fmt.Errorf("could not get price for token: %s", token)
}
