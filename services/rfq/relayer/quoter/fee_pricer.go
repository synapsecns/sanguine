package quoter

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
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
	config relconfig.FeePricerConfig
	// gasPriceCache maps chainID -> gas price
	gasPriceCache *ttlcache.Cache[uint32, *big.Int]
	// tokenPriceCache maps token name -> token price
	tokenPriceCache *ttlcache.Cache[string, *big.Int]
	omniClient      omnirpcClient.RPCClient
	httpClient      *http.Client
}

const coingeckoURL = "https://api.coingecko.com/api/v3/simple/price?vs_currencies=USD&ids="

// NewFeePricer creates a new fee pricer.
func NewFeePricer(config relconfig.FeePricerConfig, omnirpcURL string, metricHandler metrics.Handler) FeePricer {
	omniClient := omnirpcClient.NewOmnirpcClient(omnirpcURL, metricHandler, omnirpcClient.WithCaptureReqRes())
	httpClient := &http.Client{Timeout: time.Duration(config.RequestTimeoutMs) * time.Millisecond}
	return &feePricer{
		config:          config,
		gasPriceCache:   ttlcache.New[uint32, *big.Int](ttlcache.WithTTL[uint32, *big.Int](time.Second * time.Duration(config.GasPriceCacheTTL))),
		tokenPriceCache: ttlcache.New[string, *big.Int](ttlcache.WithTTL[string, *big.Int](time.Second * time.Duration(config.TokenPriceCacheTTL))),
		omniClient:      omniClient,
		httpClient:      httpClient,
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
	resp, err := f.httpClient.Get(fmt.Sprintf("%s%s", coingeckoURL, token))
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("bad status code: %v", resp.Status)
		return 0, err
	}
	defer resp.Body.Close()
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("could not read response: %w", err)
		return 0, err
	}
	var result map[string]map[string]float64
	err = json.Unmarshal(respBytes, &result)
	if err != nil {
		err = fmt.Errorf("could not unmarshal response: %w", err)
		return 0, err
	}
	for _, val := range result {
		for _, num := range val {
			return num, nil
		}
	}
	return 0, fmt.Errorf("could not get token price")
}

func (f *feePricer) getTokenDecimals(chainID uint32, token string) (uint8, error) {
	return 0, nil
}

func (f *feePricer) getNativeToken(chainID uint32) (string, error) {
	return "", nil
}
