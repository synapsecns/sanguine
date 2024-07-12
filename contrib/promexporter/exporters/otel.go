package exporters

import (
	"context"
	"math/big"

	"github.com/hedzr/log"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

type otelRecorder struct {
	metrics metrics.Handler
	meter   metric.Meter

	vPrice        float64
	bridgeBalance float64
	feeBalance    float64
	totalSupply   float64
	gasBalance    float64
	balance       float64
	nonce         int64
	stuckHeroes   int64

	tokenID string
	chainID int

	vpriceGauge        metric.Float64ObservableGauge
	bridgeBalanceGauge metric.Float64ObservableGauge
	feeBalanceGauge    metric.Float64ObservableGauge
	totalSupplyGauge   metric.Float64ObservableGauge
	gasBalanceGauge    metric.Float64ObservableGauge
	balanceGauge       metric.Float64ObservableGauge
	nonceGauge         metric.Int64ObservableGauge
	stuckCount         metric.Int64ObservableGauge
}

func newOtelRecorder(meterHandler metrics.Handler) otelRecorder {
	otr := otelRecorder{
		metrics: meterHandler,
		meter:   meterHandler.Meter(meterName),
	}
	// todo: make an option
	metricName := func(metricName string) string {
		return metricName
	}

	var err error
	if otr.vpriceGauge, err = otr.meter.Float64ObservableGauge(metricName("promexporter.vpriceGauge"), metric.WithDescription("vprice gauge"), metric.WithUnit("price")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.bridgeBalanceGauge, err = otr.meter.Float64ObservableGauge(metricName("promexporter.bridgeBalanceGauge"), metric.WithDescription("bridge balance"), metric.WithUnit("eth")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.feeBalanceGauge, err = otr.meter.Float64ObservableGauge(metricName("promexporter.feeBalanceGauage"), metric.WithDescription("fee balance gauge"), metric.WithUnit("gwei")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.totalSupplyGauge, err = otr.meter.Float64ObservableGauge(metricName("promexporter.vpriceGauage"), metric.WithDescription("vprice gauge"), metric.WithUnit("virtual price")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.gasBalanceGauge, err = otr.meter.Float64ObservableGauge(metricName("promexporter.gasBalance"), metric.WithDescription("vprice gauge"), metric.WithUnit("virtual price")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.balanceGauge, err = otr.meter.Float64ObservableGauge(metricName("promexporter.balanceGauge"), metric.WithDescription("balance gauge"), metric.WithUnit("eth")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.nonceGauge, err = otr.meter.Int64ObservableGauge(metricName("promexporter.nonceGauge"), metric.WithDescription("nonce gauge"), metric.WithUnit("nonce")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.stuckCount, err = otr.meter.Int64ObservableGauge(metricName("promexporter.stuckCount"), metric.WithDescription("stuck count gauge"), metric.WithUnit("count")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if _, err = otr.meter.RegisterCallback(otr.recordVpriceGauge, otr.vpriceGauge); err != nil {
		log.Warnf("failed to register callback for vprice gauge: %v", err)
	}

	if _, err = otr.meter.RegisterCallback(otr.recordTokenBalance, otr.balanceGauge); err != nil {
		log.Warnf("failed to register callback for balance gauge: %v", err)
	}

	// register callbacks
	return otr
}

// Virtual Price Metrics
func (o *otelRecorder) RecordVPrice(vPrice float64) {
	o.vPrice = vPrice
}

func (o *otelRecorder) recordVpriceGauge(
	parentCtx context.Context,
	observer metric.Observer,
) (err error) {
	if o.metrics == nil || o.vpriceGauge == nil {
		return nil
	}

	observer.ObserveFloat64(
		o.vpriceGauge,
		o.vPrice,
		metric.WithAttributes(attribute.Int(metrics.ChainID, o.chainID)),
	)

	return nil
}

// Token Balance Metrics
func (o *otelRecorder) RecordTokenBalance(
	parentCtx context.Context,
	bridgeBalance *big.Int,
	feeBalance *big.Int,
	totalSupply *big.Int,
	chainID int,
	tokenData []tokenData,
) (err error) {
	o.bridgeBalance = core.BigToDecimals(bridgeBalance, 18)
	o.feeBalance = core.BigToDecimals(feeBalance, 18)
	o.totalSupply = core.BigToDecimals(totalSupply, 18)
	return nil
}

func (o *otelRecorder) recordTokenBalance(
	parentCtx context.Context,
	observer metric.Observer,
) (err error) {

	if o.metrics == nil || o.balanceGauge == nil {
		return nil
	}

	observer.ObserveFloat64(
		o.gasBalanceGauge,
		o.gasBalance,
		metric.WithAttributes(attribute.Int(metrics.ChainID, o.chainID)),
	)

	observer.ObserveFloat64(
		o.balanceGauge,
		o.balance,
		metric.WithAttributes(
			attribute.Int(metrics.ChainID, o.chainID),
			attribute.String("tokenID", o.tokenID),
		),
	)

	observer.ObserveFloat64(
		o.feeBalanceGauge,
		o.feeBalance,
		metric.WithAttributes(
			attribute.Int(metrics.ChainID, o.chainID),
			attribute.String("tokenID", o.tokenID),
		),
	)

	observer.ObserveFloat64(
		o.totalSupplyGauge,
		o.totalSupply,
		metric.WithAttributes(
			attribute.Int(metrics.ChainID, o.chainID),
			attribute.String("tokenID", o.tokenID),
		),
	)

	return nil

}

type tokenData struct {
	metadata        TokenConfig
	contractBalance *big.Int
	totalSuppply    *big.Int
	feeBalance      *big.Int
}
