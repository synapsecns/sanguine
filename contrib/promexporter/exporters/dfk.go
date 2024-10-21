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

	var totalStuckHeroes int64
	var skip int64
	var limit int64 = 100

	for {
		stuckHeroes, err := dfkClient.StuckHeroes(ctx, core.PtrTo(skip), core.PtrTo(limit), core.PtrTo(owner.String()))
		if err != nil {
			return fmt.Errorf("could not get stuck hero count: %w", err)
		}
		totalStuckHeroes += int64(len(stuckHeroes.Heroes))
		if len(stuckHeroes.Heroes) < int(limit) {
			break
		}
		skip += limit
	}

	e.otelRecorder.RecordStuckHeroCount(totalStuckHeroes, chainName)

	return nil
}
