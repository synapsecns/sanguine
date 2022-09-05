package gas

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/backend"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/londinium"
	"github.com/teivah/onecontext"
	"math/big"
)

var logger = log.Logger("synapse-gas")

// DefaultGasPriceEstimate estimates the default gas price.
// This is exposed as a variable so it can be modified by tests, but will be removed in a future version.
var DefaultGasPriceEstimate = big.NewInt(int64(params.GWei))

// PriceEstimator handles estimating gas prices for type 0 txes on evm based chains. The logic is entirely derived from eth's
// SuggestPrice() logic (see: https://git.io/JZFlE), but designed to be deterministic from a given block height
// (with older block gas price estimations garbage collected).
// This is needed because, unlike in cases  with a single signer gas prices need to be estimated deterministically.
// This is done by deriving the price at a given height (taken from the event log in cases where an event is being handles)
// and then following eth's estimation logic on the last n blocks. All clients arrive at the same price every time.
// The client is responsible for determining the height they wish to get the gas price at and coordinating this if
// the event log is cross chain.
type PriceEstimator interface {
	// EstimateGasPrice estimates a gas price deterministically at a given height
	EstimateGasPrice(ctx context.Context, height uint64, config gasprice.Config) (*big.Int, error)
}

type gasFeeEstimatorImpl struct {
	//nolint: containedctx
	ctx           context.Context
	oracleBackend londinium.OracleBackend
}

// NewGasPriceEstimator creates a gas estimator for v0 transactions.
func NewGasPriceEstimator(ctx context.Context, oracle backend.OracleBackendChain) (g PriceEstimator) {
	g = &gasFeeEstimatorImpl{ctx: ctx, oracleBackend: londinium.NewOracleBackendFromChain(oracle)}
	return g
}

// EstimateGasPrice deterministically determines the gas price at a given height.
func (g *gasFeeEstimatorImpl) EstimateGasPrice(ctx context.Context, height uint64, config gasprice.Config) (gasPrice *big.Int, err error) {
	// copy this to avoid mutating the default value
	defaultEstimate, _ := big.NewInt(0).SetString(DefaultGasPriceEstimate.String(), 10)

	if height < uint64(config.Blocks) {
		return defaultEstimate, nil
	}
	ctx, cancel := onecontext.Merge(g.ctx, ctx)
	defer cancel()

	oracleBackend := londinium.NewOracleBackendFromHeight(height, g.oracleBackend)
	gasOracle := londinium.NewOracle(oracleBackend, londinium.ToLondiniumConfig(config))
	price, err := gasOracle.SuggestPrice(ctx)
	if err != nil {
		logger.Warn(fmt.Errorf("could not estimate gas price: %w", err))
		return defaultEstimate, nil
	}

	return price, nil
}

var _ PriceEstimator = &gasFeeEstimatorImpl{}
