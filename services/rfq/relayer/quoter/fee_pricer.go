package quoter

import (
	"math/big"
	"time"

	"github.com/jellydator/ttlcache/v3"
	"github.com/synapsecns/sanguine/services/rfq/relayer/relconfig"
)

// FeePricer is the interface for the fee pricer.
type FeePricer interface {
	// Start starts the fee pricer.
	Start()
	// GetOriginFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetOriginFee(chainID uint32, denomToken string) (*big.Int, error)
	// GetDestinationFee returns the total fee for a given chainID and gas limit, denominated in a given token.
	GetDestinationFee(chainID uint32, denomToken string) (*big.Int, error)
	// GetTotalFee returns the total fee for a given origin and destination chainID, denominated in a given token.
	GetTotalFee(origin, destination uint32, denomToken string) (*big.Int, error)
}

type feePricer struct {
	config relconfig.FeePricerConfig
	// gasPriceCache maps chainID -> gas price
	gasPriceCache *ttlcache.Cache[uint32, *big.Int]
	// tokenPriceCache maps token name -> token price
	tokenPriceCache *ttlcache.Cache[string, *big.Int]
}

// NewFeePricer creates a new fee pricer.
func NewFeePricer(config relconfig.FeePricerConfig) FeePricer {
	return &feePricer{
		config:          config,
		gasPriceCache:   ttlcache.New[uint32, *big.Int](ttlcache.WithTTL[uint32, *big.Int](time.Second * time.Duration(config.GasPriceCacheTTL))),
		tokenPriceCache: ttlcache.New[string, *big.Int](ttlcache.WithTTL[string, *big.Int](time.Second * time.Duration(config.TokenPriceCacheTTL))),
	}
}

func (f *feePricer) Start() {
	f.gasPriceCache.Start()
	f.tokenPriceCache.Start()
}

func (f *feePricer) GetOriginFee(chainID uint32, denomToken string) (*big.Int, error) {
	return nil, nil
}

func (f *feePricer) GetDestinationFee(chainID uint32, denomToken string) (*big.Int, error) {
	return nil, nil
}

func (f *feePricer) GetTotalFee(origin, destination uint32, denomToken string) (*big.Int, error) {
	return nil, nil
}
