package exporters

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/synapsecns/sanguine/contrib/promexporter/internal/gql/dfk"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/metrics"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
)

const stuckHeroMetric = "dfk_pending_heroes"

func (e *exporter) stuckHeroCountStats(owner common.Address, chainName string) error {
	meter := e.metrics.Meter(meterName)
	attributes := attribute.NewSet(attribute.String("chain_name", chainName))

	stuckCount, err := meter.Int64ObservableGauge(stuckHeroMetric)
	if err != nil {
		return fmt.Errorf("could not create gauge: %w", err)
	}

	if _, err := meter.RegisterCallback(func(parentCtx context.Context, o metric.Observer) (err error) {
		ctx, span := e.metrics.Tracer().Start(parentCtx, "dfk_stats", trace.WithAttributes(
			attribute.String("chain_name", chainName),
		))

		defer func() {
			metrics.EndSpanWithErr(span, err)
		}()

		ctx, cancel := context.WithTimeout(ctx, time.Minute)
		defer cancel()

		dfkClient := dfk.NewClient(e.client, e.cfg.DFKUrl)

		stuckHeroes, err := dfkClient.StuckHeroes(ctx, core.PtrTo[int64](0), core.PtrTo(owner.String()))
		if err != nil {
			return fmt.Errorf("could not get stuck hero count: %w", err)
		}

		// TODO: this maxes out at 100 now. Need binary search or something.
		o.ObserveInt64(stuckCount, int64(len(stuckHeroes.Heroes)), metric.WithAttributeSet(attributes))

		return nil
	}, stuckCount); err != nil {
		return fmt.Errorf("registering callback on instruments: %w", err)
	}

	return nil
}
