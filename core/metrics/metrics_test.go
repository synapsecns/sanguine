package metrics_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/core/metrics/localmetrics"
	"testing"
)

func TestSetupFromEnv(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Set up test Jaeger once for all tests
	ts := localmetrics.SetupTestJaeger(ctx, t)
	jaegerAvailable := ts != nil

	for _, handler := range metrics.AllHandlerTypes {
		handler := handler // capture func literal
		t.Run(handler.String(), func(t *testing.T) {
			if handler.String() == "Jaeger" && !jaegerAvailable {
				t.Skip("Skipping Jaeger test due to container startup failure")
				return
			}

			assert.NotPanics(t, func() {
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()

				t.Setenv(metrics.HandlerEnv, handler.String())

				// Set Jaeger environment variables only for Jaeger handler
				if handler.String() == "Jaeger" {
					t.Setenv("JAEGER_ENDPOINT", "http://localhost:14268/api/traces")
					t.Setenv("JAEGER_UI", "http://localhost:16686")
					t.Setenv("JAEGER_AGENT_HOST", "localhost")
					t.Setenv("JAEGER_AGENT_PORT", "6831")
				}

				_, err := metrics.NewFromEnv(ctx, config.NewBuildInfo(config.DefaultVersion, config.DefaultCommit, config.AppName, config.DefaultDate))
				Nil(t, err)
			})
		})
	}
}
