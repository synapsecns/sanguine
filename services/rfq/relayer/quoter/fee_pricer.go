package quoter

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/core/metrics"
	omnirpcClient "github.com/synapsecns/sanguine/services/omnirpc/client"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

// FeePricer is the interface for the fee pricer.
type FeePricer interface {
	// Start starts the fee pricer.
	Start()
	// GetOriginFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetOriginFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error)
	// GetDestinationFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error)
}

type feePricer struct {
	// config is the fee pricer config.
	config relconfig.FeePricerConfig
	// tokenConfigs maps chain ID -> chain config.
	chainConfigs map[int]relconfig.ChainConfig
	// gasPriceCache maps chainID -> gas price
	gasPriceCache *ttlcache.Cache[uint32, *big.Int]
	// tokenPriceCache maps token name -> token price
	tokenPriceCache *ttlcache.Cache[string, *big.Int]
	// omnirpcClient is the omnirpc client.
	omniClient omnirpcClient.RPCClient
}

// NewFeePricer creates a new fee pricer.
func NewFeePricer(config relconfig.FeePricerConfig, chainConfigs map[int]relconfig.ChainConfig, omnirpcURL string, metricHandler metrics.Handler) FeePricer {
	omniClient := omnirpcClient.NewOmnirpcClient(omnirpcURL, metricHandler, omnirpcClient.WithCaptureReqRes())
	return &feePricer{
		config:          config,
		chainConfigs:    chainConfigs,
		gasPriceCache:   ttlcache.New[uint32, *big.Int](ttlcache.WithTTL[uint32, *big.Int](time.Second * time.Duration(config.GasPriceCacheTTL))),
		tokenPriceCache: ttlcache.New[string, *big.Int](ttlcache.WithTTL[string, *big.Int](time.Second * time.Duration(config.TokenPriceCacheTTL))),
		omniClient:      omniClient,
	}
}

func (f *feePricer) Start() {
	// Start the TTL caches.
	f.gasPriceCache.Start()
	f.tokenPriceCache.Start()
}

var nativeDecimalsFactor = new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(18)), nil)

func (f *feePricer) GetOriginFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error) {
	gasPrice, err := f.getGasPrice(ctx, origin)
	if err != nil {
		return nil, err
	}
	nativeToken, err := f.getNativeToken(origin)
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
	denomTokenDecimals, err := f.getTokenDecimals(destination, denomToken)
	if err != nil {
		return nil, err
	}

	// Compute the fee in USD terms.
	originFee := new(big.Float).SetInt(gasPrice.Mul(gasPrice, big.NewInt(int64(f.config.OriginGasEstimate))))
	originFeeUSD := new(big.Float).Mul(originFee, new(big.Float).SetFloat64(nativeTokenPrice))

	// Convert the USD value to the deonominated token.
	originFeeDenom := new(big.Float).Mul(originFeeUSD, new(big.Float).SetFloat64(denomTokenPrice))
	denomDecimalsFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(denomTokenDecimals)), nil)
	originFeeDenomDecimals, _ := new(big.Float).Mul(originFeeDenom, new(big.Float).SetInt(denomDecimalsFactor)).Int(nil)
	return originFeeDenomDecimals, nil
}

func (f *feePricer) GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error) {
	gasPrice, err := f.getGasPrice(ctx, origin)
	if err != nil {
		return nil, err
	}
	nativeToken, err := f.getNativeToken(origin)
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
	denomTokenDecimals, err := f.getTokenDecimals(destination, denomToken)
	if err != nil {
		return nil, err
	}

	// Compute the fee in USD terms.
	originFee := new(big.Float).SetInt(gasPrice.Mul(gasPrice, big.NewInt(int64(f.config.DestinationGasEstimate))))
	originFeeUSD := new(big.Float).Mul(originFee, new(big.Float).SetFloat64(nativeTokenPrice))

	// Convert the USD value to the deonominated token.
	originFeeDenom := new(big.Float).Mul(originFeeUSD, new(big.Float).SetFloat64(denomTokenPrice))
	denomDecimalsFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(denomTokenDecimals)), nil)
	originFeeDenomDecimals, _ := new(big.Float).Mul(originFeeDenom, new(big.Float).SetInt(denomDecimalsFactor)).Int(nil)
	return originFeeDenomDecimals, nil
}

func (f *feePricer) GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error) {
	originFee, err := f.GetOriginFee(ctx, origin, destination, denomToken)
	if err != nil {
		return nil, err
	}
	destFee, err := f.GetDestinationFee(ctx, origin, destination, denomToken)
	if err != nil {
		return nil, err
	}
	totalFee := new(big.Int).Add(originFee, destFee)
	return totalFee, nil
}

func (f *feePricer) getFee(origin, destination uint32, denomToken string) (*big.Int, error) {
	return nil, nil
}

// getGasPrice returns the gas price for a given chainID in native units.
func (f *feePricer) getGasPrice(ctx context.Context, chainID uint32) (*big.Int, error) {
	client, err := f.omniClient.GetChainClient(ctx, int(chainID))
	if err != nil {
		return nil, err
	}
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return header.BaseFee, nil
}

// getTokenPrice returns the price of a token in USD.
func (f *feePricer) getTokenPrice(ctx context.Context, token string) (float64, error) {
	for _, chainConfig := range f.chainConfigs {
		for tokenID, tokenConfig := range chainConfig.Tokens {
			if token == tokenID {
				return tokenConfig.PriceUSD, nil
			}

		}
	}
	return 0, fmt.Errorf("could not get price for token: %s", token)
}

func (f *feePricer) getTokenDecimals(chainID uint32, token string) (uint8, error) {
	return 0, nil
}

func (f *feePricer) getNativeToken(chainID uint32) (string, error) {
	return "", nil
}
