package exporters

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/lmittmann/w3/module/eth"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const gasBalance = "gas_balance"
const nonce = "nonce"

// note: this kind of check should be deprecated in favor of submitter metrics once everything has been moved over.
func (e *exporter) submitterStats(address common.Address, chainID int, name string) error {
	meter := e.metrics.Meter(fmt.Sprintf("%s_%d", meterName, chainID))

	balanceGauge, err := meter.Float64ObservableGauge(gasBalance)
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	nonceGauge, err := meter.Int64ObservableGauge(nonce)
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	attributes := attribute.NewSet(attribute.Int(metrics.ChainID, chainID), attribute.String(metrics.EOAAddress, address.String()), attribute.String("name", name))

	if _, err := meter.RegisterCallback(
		func(parentCtx context.Context, o metric.Observer) (err error) {
			ctx, span := e.metrics.Tracer().Start(parentCtx, "submitter_stats", trace.WithAttributes(
				attribute.Int(metrics.ChainID, chainID),
				attribute.String(metrics.EOAAddress, address.String()),
			))

			defer func() {
				metrics.EndSpanWithErr(span, err)
			}()

			client, err := e.omnirpcClient.GetConfirmationsClient(ctx, chainID, 1)
			if err != nil {
				return fmt.Errorf("could not get confirmations client: %w", err)
			}

			var nonce uint64
			var balance big.Int

			err = client.BatchWithContext(ctx,
				eth.Nonce(address, nil).Returns(&nonce),
				eth.Balance(address, nil).Returns(&balance),
			)

			if err != nil {
				return fmt.Errorf("could not get balance: %w", err)
			}

			ethBalance := new(big.Float).Quo(new(big.Float).SetInt(&balance), new(big.Float).SetInt64(params.Ether))
			truncEthBalance, _ := ethBalance.Float64()

			o.ObserveFloat64(balanceGauge, truncEthBalance, metric.WithAttributeSet(attributes))
			o.ObserveInt64(nonceGauge, int64(nonce), metric.WithAttributeSet(attributes))

			return nil
		}, balanceGauge, nonceGauge); err != nil {
		return fmt.Errorf("registering callback on instruments: %w", err)
	}

	return nil
}
