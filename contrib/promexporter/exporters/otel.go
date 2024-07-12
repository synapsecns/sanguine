package exporters

import (
	"context"

	"github.com/hedzr/log"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

type otelRecorder struct {
	metrics metrics.Handler
	// meter is the metrics meter.
	meter metric.Meter
	// bridge metrics
	vpriceGauge        metric.Float64ObservableGauge
	bridgeBalanceGauge metric.Float64ObservableGauge
	feeBalanceGauge    metric.Float64ObservableGauge
	totalSupplyGauge   metric.Float64ObservableGauge
	gasBalanceGauge    metric.Float64ObservableGauge

	// submitter metrics
	balanceGauge metric.Float64ObservableGauge
	nonceGauge   metric.Int64ObservableGauge

	// dfk metrics
	stuckCount metric.Int64ObservableGauge
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
	if otr.vpriceGauge, err = otr.meter.Float64ObservableGauge(
		metricName("promexporter.vpriceGauge"),
		metric.WithDescription("vprice gauge"),
		metric.WithUnit("price")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.bridgeBalanceGauge, err = otr.meter.Float64ObservableGauge(
		metricName("promexporter.bridgeBalanceGauge"),
		metric.WithDescription("bridge balance"),
		metric.WithUnit("eth")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.feeBalanceGauge, err = otr.meter.Float64ObservableGauge(
		metricName("promexporter.feeBalanceGauage"),
		metric.WithDescription("fee balance gauge"),
		metric.WithUnit("gwei")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.totalSupplyGauge, err = otr.meter.Float64ObservableGauge(
		metricName("promexporter.vpriceGauage"),
		metric.WithDescription("vprice gauge"),
		metric.WithUnit("virtual price")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.gasBalanceGauge, err = otr.meter.Float64ObservableGauge(
		metricName("promexporter.gasBalance"),
		metric.WithDescription("vprice gauge"),
		metric.WithUnit("virtual price")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.balanceGauge, err = otr.meter.Float64ObservableGauge(
		metricName("promexporter.balanceGauge"),
		metric.WithDescription("balance gauge"),
		metric.WithUnit("eth")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.nonceGauge, err = otr.meter.Int64ObservableGauge(
		metricName("promexporter.nonceGauge"),
		metric.WithDescription("nonce gauge"),
		metric.WithUnit("nonce")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	if otr.stuckCount, err = otr.meter.Int64ObservableGauge(
		metricName("promexporter.stuckCount"),
		metric.WithDescription("stuck count gauge"),
		metric.WithUnit("count")); err != nil {
		log.Warnf("failed to create gauge: %v", err)
	}

	// register callbacks
	return otr
}

func (o *otelRecorder) recordVpriceGauge(
	parentCtx context.Context,
	vPrice float64,
	chainID int,
	tokenID string,
	observer metric.Observer,
) (err error) {
	if o.metrics == nil || o.vpriceGauge == nil {
		return nil
	}

	attributes := attribute.NewSet(
		attribute.Int(metrics.ChainID, chainID),
		attribute.String("tokenID", tokenID),
	)

	_, span := o.metrics.Tracer().Start(
		parentCtx,
		"vprice_stats",
		trace.WithAttributes(
			attribute.Int(metrics.ChainID, chainID),
			attribute.String("tokenID", tokenID),
		))

	defer func() {
		metrics.EndSpanWithErr(span, err)
	}()

	observer.ObserveFloat64(o.vpriceGauge, vPrice, metric.WithAttributeSet(attributes))

	return nil
}
