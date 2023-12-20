package pricer

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/ethergo/submitter"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
	"golang.org/x/sync/errgroup"
)

// FeePricer is the interface for the fee pricer.
type FeePricer interface {
	// Start starts the fee pricer.
	Start(ctx context.Context)
	// GetOriginFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetOriginFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error)
	// GetDestinationFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error)
	// GetGasPrice returns the gas price for a given chainID in native units.
	GetGasPrice(ctx context.Context, chainID uint32) (*big.Int, error)
	// GetTokenID returns the token ID for a given chainID and token address.
	// TODO: this can probably move into relconfig.Config, but since Quoter does not have a config,
	// we expose this functionality here.
	GetTokenID(chainID uint32, addr string) (string, error)
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
	// clientFetcher is used to fetch clients.
	clientFetcher submitter.ClientFetcher
}

// NewFeePricer creates a new fee pricer.
func NewFeePricer(config relconfig.FeePricerConfig, chainConfigs map[int]relconfig.ChainConfig, clientFetcher submitter.ClientFetcher) FeePricer {
	gasPriceCache := ttlcache.New[uint32, *big.Int](
		ttlcache.WithTTL[uint32, *big.Int](time.Second*time.Duration(config.GasPriceCacheTTLSeconds)),
		ttlcache.WithDisableTouchOnHit[uint32, *big.Int](),
	)
	return &feePricer{
		config:          config,
		chainConfigs:    chainConfigs,
		gasPriceCache:   gasPriceCache,
		tokenPriceCache: ttlcache.New[string, *big.Int](ttlcache.WithTTL[string, *big.Int](time.Second * time.Duration(config.TokenPriceCacheTTLSeconds))),
		clientFetcher:   clientFetcher,
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

func (f *feePricer) GetOriginFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error) {
	return f.getFee(ctx, origin, destination, f.config.OriginGasEstimate, denomToken)
}

func (f *feePricer) GetDestinationFee(ctx context.Context, origin, destination uint32, denomToken string) (*big.Int, error) {
	return f.getFee(ctx, destination, destination, f.config.DestinationGasEstimate, denomToken)
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

func (f *feePricer) getFee(ctx context.Context, gasChain, denomChain uint32, gasEstimate int, denomToken string) (*big.Int, error) {
	gasPrice, err := f.GetGasPrice(ctx, gasChain)
	if err != nil {
		return nil, err
	}
	nativeToken, err := f.getNativeToken(gasChain)
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
	denomTokenDecimals, err := f.getTokenDecimals(denomChain, denomToken)
	if err != nil {
		return nil, err
	}
	denomDecimalsFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(denomTokenDecimals)), nil)

	// Compute the fee in ETH terms.
	feeWei := new(big.Float).Mul(new(big.Float).SetInt(gasPrice), new(big.Float).SetFloat64(float64(gasEstimate)))
	feeEth := new(big.Float).Quo(feeWei, new(big.Float).SetInt(nativeDecimalsFactor))
	feeUSD := new(big.Float).Mul(feeEth, new(big.Float).SetFloat64(nativeTokenPrice))
	feeUSDC := new(big.Float).Mul(feeUSD, new(big.Float).SetFloat64(denomTokenPrice))
	// Note that this rounds towards zero- we may need to apply rounding here if
	// we want to be conservative and lean towards overestimating fees.
	feeUSDCDecimals, _ := new(big.Float).Mul(feeUSDC, new(big.Float).SetInt(denomDecimalsFactor)).Int(nil)
	return feeUSDCDecimals, nil
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

func (f *feePricer) GetTokenID(chain uint32, addr string) (string, error) {
	chainConfig, ok := f.chainConfigs[int(chain)]
	if !ok {
		return "", fmt.Errorf("no chain config for chain %d", chain)
	}
	for tokenID, tokenConfig := range chainConfig.Tokens {
		// TODO: probably a better way to do this.
		if strings.ToLower(tokenConfig.Address) == strings.ToLower(addr) {
			return tokenID, nil
		}
	}
	return "", fmt.Errorf("no tokenID found for chain %d and address %s", chain, addr)
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
	chainConfig, ok := f.chainConfigs[int(chainID)]
	if !ok {
		return 0, fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	for tokenID, tokenConfig := range chainConfig.Tokens {
		if token == tokenID {
			return tokenConfig.Decimals, nil
		}
	}
	return 0, fmt.Errorf("could not get token decimals for chain %d and token %s", chainID, token)
}

func (f *feePricer) getNativeToken(chainID uint32) (string, error) {
	chainConfig, ok := f.chainConfigs[int(chainID)]
	if !ok {
		return "", fmt.Errorf("could not get chain config for chainID: %d", chainID)
	}
	if len(chainConfig.NativeToken) == 0 {
		return "", fmt.Errorf("chain config for chainID %d does not have a native token", chainID)
	}
	return chainConfig.NativeToken, nil
}
