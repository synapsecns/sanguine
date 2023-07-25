package metrics_test

import (
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics"
	"testing"
)

func TestNoopExporter(t *testing.T) {
	exporter := metrics.NewNoOpExporter()
	temporality := exporter.Temporality(1)
	NotNil(t, temporality)
	aggregation := exporter.Aggregation(1)
	NotNil(t, aggregation)
	export := exporter.Export(nil, nil)
	Nil(t, export)
	flush := exporter.ForceFlush(nil)
	Nil(t, flush)
	shutdown := exporter.Shutdown(nil)
	Nil(t, shutdown)
}
