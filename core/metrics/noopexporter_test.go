package metrics_test

import (
	"context"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"testing"
)

func TestNoopExporter(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exporter := metrics.NewNoOpExporter()
	temporality := exporter.Temporality(1)
	NotNil(t, temporality)
	aggregation := exporter.Aggregation(1)
	NotNil(t, aggregation)
	export := exporter.Export(ctx, nil)
	Nil(t, export)
	flush := exporter.ForceFlush(ctx)
	Nil(t, flush)
	shutdown := exporter.Shutdown(ctx)
	Nil(t, shutdown)
}
