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
	"go.opentelemetry.io/otel/trace"
)

func (e *exporter) stuckHeroCountStats(parentCtx context.Context, owner common.Address, chainName string) (err error) {
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

	e.otelRecorder.RecordStuckHeroCount(int64(len(stuckHeroes.Heroes)), chainName)

	return nil
}
