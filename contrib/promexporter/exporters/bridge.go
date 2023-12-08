package exporters

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridge"
	"github.com/synapsecns/sanguine/services/explorer/contracts/bridgeconfig"
	"github.com/synapsecns/sanguine/services/explorer/contracts/swap"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"math"
	"math/big"
	"time"
)

func (e *exporter) getBridgeConfig(ctx context.Context) (*bridgeconfig.BridgeConfigRef, error) {
	client, err := e.omnirpcClient.GetConfirmationsClient(ctx, e.cfg.BridgeConfig.ChainID, 1)
	if err != nil {
		return nil, fmt.Errorf("could not get confirmations client: %w", err)
	}

	// note this will not update
	configContract, err := bridgeconfig.NewBridgeConfigRef(common.HexToAddress(e.cfg.BridgeConfig.Address), client)
	if err != nil {
		return nil, fmt.Errorf("could not get bridge config contract: %w", err)
	}
	return configContract, nil
}

// Will be a lot faster w/: https://github.com/open-telemetry/opentelemetry-go/issues/3034
// nolint: cyclop
func (e *exporter) vpriceStats(ctx context.Context, chainID int, tokenID string) error {
	meter := e.metrics.Meter(meterName)
	vpriceMetric, err := meter.Float64ObservableGauge("vpriceMetric")
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	attributes := attribute.NewSet(attribute.Int(metrics.ChainID, chainID), attribute.String("tokenID", tokenID))

	client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
	if err != nil {
		return fmt.Errorf("could not get confirmations client: %w", err)
	}

	bridgeConfig, err := e.getBridgeConfig(ctx)
	if err != nil {
		return err
	}

	token, err := bridgeConfig.GetToken(&bind.CallOpts{Context: ctx}, tokenID, big.NewInt(int64(chainID)))
	if err != nil {
		return fmt.Errorf("could not get token: %w", err)
	}

	poolConfig, err := bridgeConfig.GetPoolConfig(&bind.CallOpts{Context: ctx}, common.HexToAddress(token.TokenAddress), big.NewInt(int64(chainID)))
	if err != nil {
		return errPoolNotExist
	}

	// pool doesn't exist, no metrics to record!
	if poolConfig.PoolAddress == common.BigToAddress(big.NewInt(0)) {
		return errPoolNotExist
	}

	tokenContract, err := bridge.NewERC20(common.HexToAddress(token.TokenAddress), client)
	if err != nil {
		return fmt.Errorf("could not get tokenID contract: %w", err)
	}

	decimals, err := tokenContract.Decimals(&bind.CallOpts{Context: ctx})
	if err != nil {
		return fmt.Errorf("could not get decimals: %w", err)
	}

	iswap, err := swap.NewISwap(poolConfig.PoolAddress, client)
	if err != nil {
		return fmt.Errorf("could not get iswap contract: %w", err)
	}

	if _, err := meter.RegisterCallback(func(parentCtx context.Context, o metric.Observer) (err error) {
		ctx, span := e.metrics.Tracer().Start(parentCtx, "vprice_stats", trace.WithAttributes(
			attribute.Int(metrics.ChainID, chainID), attribute.String("tokenID", tokenID),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		realvPrice, err := iswap.GetVirtualPrice(&bind.CallOpts{Context: ctx})
		if err != nil {
			return fmt.Errorf("could not get virtual price: %w", err)
		}

		// Assuming `vpriceMetric` is of type *big.Int and `decimals` is an int

		// Convert vpriceMetric to *big.Float
		bigVPrice := new(big.Float).SetInt(realvPrice)

		// Calculate the divisor for decimals
		divisor := new(big.Float).SetFloat64(math.Pow10(int(decimals)))

		// Divide bigVPrice by the divisor to account for decimals
		realVPrice := new(big.Float).Quo(bigVPrice, divisor)

		// Convert the final value to float64
		floatVPrice, _ := realVPrice.Float64()

		// Use floatVPrice as required
		o.ObserveFloat64(vpriceMetric, floatVPrice, metric.WithAttributeSet(attributes))

		return nil
	}, vpriceMetric); err != nil {
		return fmt.Errorf("registering callback on instruments: %w", err)
	}

	return nil
}

var errPoolNotExist = errors.New("pool does not exist")
