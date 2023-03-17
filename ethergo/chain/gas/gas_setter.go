package gas

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/params"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/ethergo/chain/client"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/backend"
	"github.com/synapsecns/sanguine/ethergo/chain/gas/london"
	"math/big"
)

// feeCapBumpPercentage is how much the fee cap should bump (expressed as a percentage).
const feeCapBumpPercentage = 7

// Setter handles setting gas prices/fees in both pre (v0) and post (v2) eip-1559 transactions.
// this allows the setter to work on both ethereum and third party rpcs regardless of eip-1559 support.
type Setter interface {
	// SetGasFeeByBlock gets a deterministic gas price based on the block
	SetGasFeeByBlock(ctx context.Context, transactor *bind.TransactOpts, gasBlock uint64, maxPrice *big.Int) error
	// SetGasFee sets a gas fee non-deterministically. This is useful for speeding up tests
	SetGasFee(ctx context.Context, transactor *bind.TransactOpts, londonBlock uint64, maxPrice *big.Int) error
}

// GetConfig ges the gas price config.
func GetConfig() gasprice.Config {
	config := ethconfig.FullNodeGPO
	config.MaxPrice = big.NewInt(params.GWei * 750)
	config.IgnorePrice = core.CopyBigInt(ethconfig.FullNodeGPO.IgnorePrice)
	config.Percentile = 90
	config.Default = big.NewInt(params.GWei)
	return config
}

type gasSetterImpl struct {
	//nolint: containedctx
	ctx            context.Context
	oracleBackend  backend.OracleBackendChain
	priceEstimator PriceEstimator
}

// NewGasSetter creates a new gas setter for v0 & v2 transactions depending on the chain settings.
func NewGasSetter(ctx context.Context, oracle backend.OracleBackendChain) Setter {
	return &gasSetterImpl{
		ctx:            ctx,
		oracleBackend:  oracle,
		priceEstimator: NewGasPriceEstimator(ctx, oracle),
	}
}

func (g gasSetterImpl) SetGasFeeByBlock(ctx context.Context, transactor *bind.TransactOpts, gasBlock uint64, maxPrice *big.Int) (err error) {
	oracleConfig := GetConfig()
	oracleConfig.MaxPrice = maxPrice

	// if london is activated, we only care about the max tip cap. Max fee remains the same for every tx
	//nolint: nestif
	if client.UsesLondon(g.oracleBackend.ChainConfig(), gasBlock) {
		gasHeader, err := g.oracleBackend.HeaderByNumber(ctx, new(big.Int).SetUint64(gasBlock))
		if err != nil {
			return fmt.Errorf("could not get gas block %d on chain %d: %w", gasBlock, g.oracleBackend.ChainConfig().ChainID, err)
		}

		oracle := london.NewFeeOracle(g.oracleBackend, gasBlock, oracleConfig)

		tipCap, err := oracle.SuggestTipCap(ctx)
		if err != nil {
			return fmt.Errorf("could not get tip cap: %w", err)
		}

		transactor.GasFeeCap = maxPrice
		transactor.GasTipCap = BumpByPercent(tipCap, feeCapBumpPercentage)

		if transactor.GasTipCap.Cmp(gasHeader.BaseFee) > 0 {
			transactor.GasTipCap = new(big.Int).Sub(transactor.GasFeeCap, gasHeader.BaseFee)
		}
	} else {
		transactor.GasPrice, err = g.priceEstimator.EstimateGasPrice(ctx, gasBlock, oracleConfig)
		if err != nil {
			return fmt.Errorf("could not get gas price: %w", err)
		}
	}
	return nil
}

func (g gasSetterImpl) SetGasFee(ctx context.Context, transactor *bind.TransactOpts, londonBlock uint64, maxPrice *big.Int) (err error) {
	//nolint: nestif
	if client.UsesLondon(g.oracleBackend.ChainConfig(), londonBlock) {
		gasHeader, err := g.oracleBackend.HeaderByNumber(ctx, new(big.Int).SetUint64(londonBlock))
		if err != nil {
			return fmt.Errorf("could not get gas block %d on chain %d: %w", londonBlock, g.oracleBackend.ChainConfig().ChainID, err)
		}

		tipCap, err := g.oracleBackend.SuggestGasTipCap(ctx)
		if err != nil {
			return fmt.Errorf("could not calculate tip cap: %w", err)
		}

		// add the tip cap to the gas price and bump
		transactor.GasFeeCap = maxPrice
		transactor.GasTipCap = BumpByPercent(tipCap, feeCapBumpPercentage)

		if transactor.GasTipCap.Cmp(gasHeader.BaseFee) > 0 {
			transactor.GasTipCap = new(big.Int).Sub(transactor.GasFeeCap, gasHeader.BaseFee)
		}
	} else {
		gasPrice, err := g.oracleBackend.SuggestGasPrice(ctx)
		if err != nil {
			return fmt.Errorf("could not get gas price: %w", err)
		}

		transactor.GasPrice = min(BumpByPercent(gasPrice, feeCapBumpPercentage), maxPrice)
	}

	return nil
}
