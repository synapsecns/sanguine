// Package pricer contains pricing logic for RFQ relayer quotes.
package pricer

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
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
	GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string, isQuote bool, tx *fastbridgev2.IFastBridgeV2BridgeTransactionV2) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string, isQuote bool, tx *fastbridgev2.IFastBridgeV2BridgeTransactionV2) (*big.Int, error)
	// GetGasPrice returns the gas price for a given chainID in native units.
	GetGasPrice(ctx context.Context, chainID uint32) (*big.Int, error)
	// GetTokenPrice returns the price of a token in USD.
	GetTokenPrice(ctx context.Context, token string) (float64, error)
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
	// relayerAddress is the address of the relayer.
	relayerAddress common.Address
}

// NewFeePricer creates a new fee pricer.
func NewFeePricer(config relconfig.Config, clientFetcher submitter.ClientFetcher, priceFetcher CoingeckoPriceFetcher, handler metrics.Handler, relayerAddress common.Address) FeePricer {
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
		relayerAddress:  relayerAddress,
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

var nativeDecimalsFactor = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18)), nil)

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
		return nil, err
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

func (f *feePricer) GetDestinationFee(parentCtx context.Context, _, destination uint32, denomToken string, isQuote bool, tx *fastbridgev2.IFastBridgeV2BridgeTransactionV2) (*big.Int, error) {
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
	span.SetAttributes(attribute.String("raw_fee", fee.String()))

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

	// If specified, calculate and add the call fee, as well as the call value which will be paid by the relayer
	if tx != nil {
		if tx.CallParams != nil {
			client, err := f.clientFetcher.GetClient(ctx, big.NewInt(int64(destination)))
			if err != nil {
				return nil, fmt.Errorf("could not get client: %w", err)
			}
			parsedABI, err := abi.JSON(strings.NewReader(fastbridgev2.IFastBridgeRecipientMetaData.ABI))
			if err != nil {
				return nil, fmt.Errorf("could not parse ABI: %w", err)
			}
			methodName := "fastBridgeTransferReceived"
			encodedData, err := parsedABI.Pack(methodName, tx.DestToken, tx.DestAmount, tx.CallParams)
			if err != nil {
				return nil, fmt.Errorf("could not encode function call: %w", err)
			}

			callMsg := ethereum.CallMsg{
				From:  f.relayerAddress,
				To:    &tx.DestRecipient,
				Value: tx.CallValue,
				Data:  encodedData,
			}
			gasEstimate, err := client.EstimateGas(ctx, callMsg)
			if err != nil {
				return nil, fmt.Errorf("could not estimate gas: %w", err)
			}
			callFee, err := f.getFee(ctx, destination, destination, int(gasEstimate), denomToken, isQuote)
			if err != nil {
				return nil, err
			}
			fee = new(big.Int).Add(fee, callFee)
			span.SetAttributes(attribute.String("call_fee", callFee.String()))
		}

		if tx.CallValue != nil {
			callValueFloat := new(big.Float).SetInt(tx.CallValue)
			valueDenom, err := f.getDenomFee(ctx, destination, destination, denomToken, callValueFloat)
			if err != nil {
				return nil, err
			}
			valueScaled, err := f.getFeeWithMultiplier(ctx, destination, isQuote, valueDenom)
			if err != nil {
				return nil, err
			}
			fee = new(big.Int).Add(fee, valueScaled)
			span.SetAttributes(attribute.String("value_scaled", valueScaled.String()))
		}
	}

	span.SetAttributes(attribute.String("destination_fee", fee.String()))
	return fee, nil
}

func (f *feePricer) GetTotalFee(parentCtx context.Context, origin, destination uint32, denomToken string, isQuote bool, tx *fastbridgev2.IFastBridgeV2BridgeTransactionV2) (_ *big.Int, err error) {
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
	destFee, err := f.GetDestinationFee(ctx, origin, destination, denomToken, isQuote, tx)
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
	feeWei := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), new(big.Float).SetFloat64(float64(gasEstimate)))

	feeDenom, err := f.getDenomFee(ctx, gasChain, denomChain, denomToken, feeWei)
	if err != nil {
		return nil, err
	}

	feeScaled, err := f.getFeeWithMultiplier(ctx, gasChain, isQuote, feeDenom)
	if err != nil {
		return nil, err
	}

	span.SetAttributes(
		attribute.String("gas_price", gasPrice.String()),
		attribute.String("fee_wei", feeWei.String()),
		attribute.String("fee_denom", feeDenom.String()),
		attribute.String("fee_scaled", feeScaled.String()),
	)
	return feeScaled, nil
}

func (f *feePricer) getDenomFee(ctx context.Context, gasChain, denomChain uint32, denomToken string, feeWei *big.Float) (*big.Float, error) {
	span := trace.SpanFromContext(ctx)

	nativeToken, err := f.config.GetNativeToken(int(gasChain))
	if err != nil {
		return nil, err
	}
	nativeTokenPrice, err := f.GetTokenPrice(ctx, nativeToken)
	if err != nil {
		return nil, err
	}
	denomTokenPrice, err := f.GetTokenPrice(ctx, denomToken)
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
	span.SetAttributes(
		attribute.Float64("native_token_price", nativeTokenPrice),
		attribute.Float64("denom_token_price", denomTokenPrice),
		attribute.Int("denom_token_decimals", int(denomTokenDecimals)),
	)

	return feeDenom, nil
}

func (f *feePricer) getFeeWithMultiplier(ctx context.Context, gasChain uint32, isQuote bool, feeDenom *big.Float) (feeScaled *big.Int, err error) {
	span := trace.SpanFromContext(ctx)

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
	span.SetAttributes(
		attribute.Float64("multiplier", multiplier),
	)

	// Apply the fixed fee multiplier.
	// Note that this step rounds towards zero- we may need to apply rounding here if
	// we want to be conservative and lean towards overestimating fees.
	feeScaled, _ = new(big.Float).Mul(feeDenom, new(big.Float).SetFloat64(multiplier)).Int(nil)

	return feeScaled, nil
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
	// Attempt to fetch gas price from cache.
	tokenPriceItem := f.tokenPriceCache.Get(token)
	//nolint:nestif
	if tokenPriceItem == nil {
		// Try to get price from coingecko.
		price, err = f.priceFetcher.GetPrice(ctx, token)
		if err == nil {
			f.tokenPriceCache.Set(token, price, 0)
		} else {
			// Fallback to configured token price.
			price, err = f.getTokenPriceFromConfig(token)
			if err != nil {
				return 0, err
			}
		}
	} else {
		price = tokenPriceItem.Value()
	}
	return price, nil
}

func (f *feePricer) getTokenPriceFromConfig(token string) (float64, error) {
	for _, chainConfig := range f.config.GetChains() {
		for tokenName, tokenConfig := range chainConfig.Tokens {
			if token == tokenName {
				return tokenConfig.PriceUSD, nil
			}
		}
	}
	return 0, fmt.Errorf("could not get price for token: %s", token)
}
