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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	ethergoUtil "github.com/synapsecns/sanguine/ethergo/util"
	"github.com/synapsecns/sanguine/services/rfq/contracts/fastbridgev2"
	"github.com/synapsecns/sanguine/services/rfq/relayer/chain"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"github.com/synapsecns/sanguine/services/rfq/relayer/reldb"
	rfqUtil "github.com/synapsecns/sanguine/services/rfq/util"
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
	GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string, isQuote bool, quoteRequest *reldb.QuoteRequest) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string, isQuote bool, quoteRequest *reldb.QuoteRequest) (*big.Int, error)
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
	fee, err := f.getFee(ctx, origin, destination, gasEstimate, denomToken, isQuote, nil, true)
	if err != nil {
		return nil, err
	}

	// If specified, calculate and add the L1 fee
	l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(origin, true)
	if useL1Fee {
		l1Fee, err := f.getFee(ctx, l1ChainID, destination, l1GasEstimate, denomToken, isQuote, nil, true)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, l1Fee)
		span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
	}
	span.SetAttributes(attribute.String("origin_fee", fee.String()))
	return fee, nil
}

//nolint:gosec
func (f *feePricer) GetDestinationFee(parentCtx context.Context, _, destination uint32, denomToken string, isQuote bool, quoteRequest *reldb.QuoteRequest) (*big.Int, error) {

	var err error
	ctx, span := f.handler.Tracer().Start(parentCtx, "getDestinationFee", trace.WithAttributes(
		attribute.Int(metrics.Destination, int(destination)),
		attribute.String("denom_token", denomToken),
		attribute.Bool("is_quote", isQuote),
	))
	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	fee := big.NewInt(0)

	// Calculate the static L2 fee if it won't be incorporated by directly estimating the relay() call
	// in addZapFees().
	if quoteRequest == nil || len(quoteRequest.Transaction.ZapData) == 0 {
		gasEstimate, err := f.config.GetDestGasEstimate(int(destination))
		if err != nil {
			return nil, fmt.Errorf("could not get dest gas estimate: %w", err)
		}
		fee, err = f.getFee(ctx, destination, destination, gasEstimate, denomToken, isQuote, nil, true)
		if err != nil {
			return nil, err
		}
	}
	span.SetAttributes(attribute.String("raw_fee", fee.String()))

	// If specified, calculate and add the call fee, as well as the call value which will be paid by the relayer
	if quoteRequest != nil {
		fee, err = f.addZapFees(ctx, destination, denomToken, quoteRequest, fee)
		if err != nil {
			return nil, err
		}
	}

	chainStack, err := f.config.GetChainStack(int(destination))
	if err != nil {
		return nil, fmt.Errorf("could not get chain stack: %w", err)
	}

	// if op stack, then we must call the OP GasOracle contract to obtain the L1 gas fee.
	// otherwise, apply static L1 fee config values if they are set.
	var gasEstimateL1Wei *big.Int
	if chainStack == "evm_op_stack" && quoteRequest != nil {

		var err error
		gasEstimateL1Wei, err = f.getOpStackL1GasWeiEst(ctx, destination, quoteRequest)
		if err != nil {
			return nil, fmt.Errorf("could not get OpStack L1 wei: %w", err)
		}

		// even though we recieve wei from oracle, we need to normalize this into the denomToken via getFee -- without a multiplier
		gasEstimateL1Denom, err := f.getFee(ctx, destination, destination, 0, denomToken, isQuote, gasEstimateL1Wei, false)
		if err != nil {
			return nil, fmt.Errorf("could not normalize gasEstimateL1Wei: %w", err)
		}

		// Debug Zap Fee Summary
		// l2Fee := fee
		// fee = new(big.Int).Add(fee, gasEstimateL1Denom)
		// if len(quoteRequest.Transaction.ZapData) > 0 {
		// 	fmt.Printf("L2: %s %s + L1: (%s WEI, %s %s) = Total Fee: %s %s\n", l2Fee, denomToken, gasEstimateL1Wei, gasEstimateL1Denom, denomToken, fee, denomToken)
		// }

		span.SetAttributes(attribute.String("gasEstimateL1Wei", gasEstimateL1Wei.String()))
		span.SetAttributes(attribute.String("gasEstimateL1Denom", gasEstimateL1Denom.String()))
	} else {
		// If specified, calculate and add the L1 fee
		l1ChainID, l1GasEstimate, useL1Fee := f.config.GetL1FeeParams(destination, false)
		if useL1Fee {
			l1Fee, err := f.getFee(ctx, l1ChainID, destination, l1GasEstimate, denomToken, isQuote, nil, true)
			if err != nil {
				return nil, err
			}
			fee = new(big.Int).Add(fee, l1Fee)
			span.SetAttributes(attribute.String("l1_fee", l1Fee.String()))
		}
	}

	span.SetAttributes(attribute.String("destination_fee", fee.String()))
	return fee, nil
}

// addZapFees incorporates the cost of the call and the call value into the fee.
// Note that to be conservative, we always use the QuoteFixedFeeMultiplier over the RelayFixedFeeMultiplier.
//
//nolint:cyclop,gosec
func (f *feePricer) addZapFees(ctx context.Context, destination uint32, denomToken string, quoteRequest *reldb.QuoteRequest, fee *big.Int) (*big.Int, error) {
	span := trace.SpanFromContext(ctx)

	if quoteRequest != nil && len(quoteRequest.Transaction.ZapData) != 0 {
		gasEstimate, err := f.getZapGasEstimate(ctx, destination, quoteRequest)
		if err != nil {
			return nil, err
		}

		callFee, err := f.getFee(ctx, destination, destination, int(gasEstimate), denomToken, true, nil, false)
		if err != nil {
			return nil, err
		}
		fee = new(big.Int).Add(fee, callFee)
		span.SetAttributes(attribute.String("call_fee", callFee.String()))
	}

	if quoteRequest != nil && quoteRequest.Transaction.ZapNative != nil && quoteRequest.Transaction.ZapNative.Cmp(big.NewInt(0)) > 0 && quoteRequest.Transaction.ZapNative.Sign() > 0 {
		callValueFloat := new(big.Float).SetInt(quoteRequest.Transaction.ZapNative)
		valueDenom, err := f.getDenomFee(ctx, destination, destination, denomToken, callValueFloat)
		if err != nil {
			return nil, err
		}
		// note: amount is intentionally not scaled with any multipliers
		valueDenomInt, _ := valueDenom.Int(nil)
		fee = new(big.Int).Add(fee, valueDenomInt)
		span.SetAttributes(attribute.String("value_denom", valueDenom.String()))
	}

	return fee, nil
}

// cache so that we don't have to parse the ABI every time.
var fastBridgeV2ABI *abi.ABI

const methodName = "relayV2"

func (f *feePricer) getZapGasEstimate(ctx context.Context, destination uint32, quoteRequest *reldb.QuoteRequest) (gasEstimate uint64, err error) {

	span := trace.SpanFromContext(ctx)
	span.AddEvent("getZapGasEstimate", trace.WithAttributes(
		attribute.String("transaction_id", hexutil.Encode(quoteRequest.TransactionID[:])),
		attribute.Int("destination", int(destination)),
	))

	defer func() {
		if err != nil {
			span.AddEvent("Error in getZapGasEstimate", trace.WithAttributes(
				attribute.String("error", err.Error()),
			))
		} else {
			span.AddEvent("Completed getZapGasEstimate", trace.WithAttributes(
				attribute.String("gasEstimateUnits", fmt.Sprint(gasEstimate)),
			))
		}
	}()

	client, err := f.clientFetcher.GetClient(ctx, big.NewInt(int64(destination)))
	if err != nil {
		return 0, fmt.Errorf("could not get client: %w", err)
	}

	if fastBridgeV2ABI == nil {
		parsedABI, err := abi.JSON(strings.NewReader(fastbridgev2.IFastBridgeV2MetaData.ABI))
		if err != nil {
			return 0, fmt.Errorf("could not parse ABI: %w", err)
		}
		fastBridgeV2ABI = &parsedABI
	}

	rawRequest, err := chain.EncodeBridgeTx(quoteRequest.Transaction)
	if err != nil {
		return 0, fmt.Errorf("could not encode quote data: %w", err)
	}

	encodedData, err := fastBridgeV2ABI.Pack(methodName, rawRequest, f.relayerAddress)
	if err != nil {
		return 0, fmt.Errorf("could not encode function call: %w", err)
	}

	rfqAddr, err := f.config.GetRFQAddress(int(destination))
	if err != nil {
		return 0, fmt.Errorf("could not get RFQ address: %w", err)
	}

	callMsg := ethereum.CallMsg{
		From: f.relayerAddress,
		To:   &rfqAddr,
		Data: encodedData,
	}
	// Tx.value needs to match `DestAmount` for native gas token, or `ZapNative` for ERC20s.
	if rfqUtil.IsGasToken(quoteRequest.Transaction.DestToken) {
		callMsg.Value = quoteRequest.Transaction.DestAmount
	} else {
		callMsg.Value = quoteRequest.Transaction.ZapNative
	}

	// note: this gas limit is intentionally not modified/boosted beyond the estimate, since this is for anticipated pricing
	gasEstimate, err = client.EstimateGas(ctx, callMsg)
	if err != nil {
		errMsg := ethergoUtil.FormatError(err)
		span.RecordError(err)
		return 0, fmt.Errorf("could not estimate gas: %s", errMsg)
	}

	return gasEstimate, nil
}

// This bespoke OP-Stack functionality is necessary to estimate total gas fees with any accuracy on relevant chains.
// A typical eth_estimateGas call (performed above) will obtain the "L2" or "Execution" component of the gas fee
// whereas *this* step will obtain the "L1" or "Storage" component of the gas fee by sending the RLP encoded tx to their Gas oracle & calling getL1Fee
// https://docs.optimism.io/stack/smart-contracts#gaspriceoracle
// Note that the L2 fee is denominated in gas units and must be calculated into native gas WEI via the gas price
// but the L1 fee output from getOpStackL1GasWeiEst is already calculated in this way & denominated in native gas WEI.
func (f *feePricer) getOpStackL1GasWeiEst(ctx context.Context, destination uint32, quoteRequest *reldb.QuoteRequest) (l1GasWei *big.Int, err error) {

	if quoteRequest == nil {
		return big.NewInt(0), nil
	}

	transactionIDStr := hexutil.Encode(quoteRequest.TransactionID[:])
	span := trace.SpanFromContext(ctx)
	span.AddEvent("getOpStackL1GasWeiEst", trace.WithAttributes(
		attribute.String("transaction_id", transactionIDStr),
		attribute.Int("destination", int(destination)),
	))

	defer func() {
		if err != nil {
			span.AddEvent("Error in getOpStackL1GasWeiEst", trace.WithAttributes(
				attribute.String("error", err.Error()),
			))
		} else {
			span.AddEvent("Completed getOpStackL1GasWeiEst", trace.WithAttributes(
				attribute.String("l1GasWei", l1GasWei.String()),
			))
		}
	}()

	client, err := f.clientFetcher.GetClient(ctx, big.NewInt(int64(destination)))
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could not get client: %w", err)
	}

	if fastBridgeV2ABI == nil {
		parsedABI, err := abi.JSON(strings.NewReader(fastbridgev2.IFastBridgeV2MetaData.ABI))
		if err != nil {
			return big.NewInt(0), fmt.Errorf("could not parse ABI: %w", err)
		}
		fastBridgeV2ABI = &parsedABI
	}

	rawRequest, err := chain.EncodeBridgeTx(quoteRequest.Transaction)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could not encode quote data: %w", err)
	}

	encodedData, err := fastBridgeV2ABI.Pack(methodName, rawRequest, f.relayerAddress)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could not encode function call: %w", err)
	}

	rfqAddr, err := f.config.GetRFQAddress(int(destination))
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could not get RFQ address: %w", err)
	}

	// value needs to match `DestAmount` for native gas token, or `ZapNative` for ERC20s.
	var txValue *big.Int
	if rfqUtil.IsGasToken(quoteRequest.Transaction.DestToken) {
		txValue = quoteRequest.Transaction.DestAmount
	} else {
		txValue = quoteRequest.Transaction.ZapNative
	}

	// construct txn to encode
	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID: big.NewInt(int64(destination)),
		Nonce:   0,
		To:      &rfqAddr,
		Value:   txValue,
		// gas value here is arbitrary -- for enhanced accuracy, could supply the actual expected L2/Execution gas units,
		// but it is unlikely to alter the estimated L1 gas cost significantly
		Gas:       300000,
		GasFeeCap: big.NewInt(0),
		GasTipCap: big.NewInt(0),
		Data:      encodedData,
	})

	// RLP encode the transaction
	var rlpEncodedTx []byte
	rlpEncodedTx, err = rlp.EncodeToBytes(tx)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could not RLP encode tx: %w", err)
	}

	// fmt.Println("encodedData: ", hexutil.Encode(encodedData))
	// fmt.Println("rlpEncodedTx: ", hexutil.Encode(rlpEncodedTx))

	// getl1Fee function of Optimism Gas Oracle
	// ex: https://optimistic.etherscan.io/address/0x420000000000000000000000000000000000000F
	// ( identical address etc on all OP stack chains )
	var gasOracleABIString = `[{"inputs":[{"internalType":"bytes","name":"transaction","type":"bytes"}],"name":"getL1Fee","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`

	gasOracleABI, err := abi.JSON(strings.NewReader(gasOracleABIString))
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could not parse gasOracle ABI: %w", err)
	}

	getL1FeeCall, err := gasOracleABI.Pack("getL1Fee", rlpEncodedTx)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could not pack method call: %w", err)
	}

	oracleAddress := common.HexToAddress("0x420000000000000000000000000000000000000F")

	oracleCallMsg := ethereum.CallMsg{
		To:   &oracleAddress,
		Data: getL1FeeCall,
	}

	result, err := client.CallContract(ctx, oracleCallMsg, nil)
	if err != nil {
		return big.NewInt(0), fmt.Errorf("could call OP stack gas oracle: %w", err)
	}

	l1Gas := new(big.Int).SetBytes(result)

	return l1Gas, nil
}

func (f *feePricer) GetTotalFee(parentCtx context.Context, origin, destination uint32, denomToken string, isQuote bool, quoteRequest *reldb.QuoteRequest) (_ *big.Int, err error) {
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
	destFee, err := f.GetDestinationFee(ctx, origin, destination, denomToken, isQuote, quoteRequest)
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

func (f *feePricer) getFee(parentCtx context.Context, gasChain, denomChain uint32, gasEstimateUnits int, denomToken string, isQuote bool, gasEstimateWei *big.Int, applyFixedFeeMult bool) (_ *big.Int, err error) {

	if gasEstimateWei == nil {
		gasEstimateWei = big.NewInt(0)
	}

	ctx, span := f.handler.Tracer().Start(parentCtx, "getFee", trace.WithAttributes(
		attribute.Int("gas_chain", int(gasChain)),
		attribute.Int("denom_chain", int(denomChain)),
		attribute.Int("gasEstimateUnits", gasEstimateUnits),
		attribute.Int("gasEstimateWei", gasEstimateUnits),
		attribute.String("denom_token", denomToken),
	))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	gasPrice, err := f.GetGasPrice(ctx, gasChain)
	if err != nil {
		return nil, err
	}
	feeWei := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), new(big.Float).SetFloat64(float64(gasEstimateUnits)))

	// if gasEstimateWei was provided, these are "pre-calculated" from gas units & do not require price multiplier.
	// add them onto any wei from gasUnits that were calculated above.
	// note: This was built for gas est functions whose output is already in Wei - rather than units - such as OP Stack gas oracle
	feeWei = new(big.Float).Add(feeWei, new(big.Float).SetInt(gasEstimateWei))

	feeDenom, err := f.getDenomFee(ctx, gasChain, denomChain, denomToken, feeWei)
	if err != nil {
		return nil, err
	}

	// conditionally apply Fixed-Fee multiplier.
	// Non-Fixed fees (eg: sims / oracle calls) are typically accurate as-is, will not benefit from multiplier, & thus dont apply it.
	feeScaled := new(big.Int)
	if applyFixedFeeMult {
		var err error
		feeScaled, err = f.getFeeWithMultiplier(ctx, gasChain, isQuote, feeDenom)
		if err != nil {
			return nil, err
		}
	} else {
		feeScaled, _ = feeDenom.Int(nil)
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
