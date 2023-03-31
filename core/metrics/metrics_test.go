package metrics_test

import (
	"context"
	"github.com/brianvoe/gofakeit/v6"
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

	localmetrics.SetupTestJaeger(ctx, t)

	for _, handler := range metrics.AllHandlerTypes {
		handler := handler // capture func literal
		assert.NotPanics(t, func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			t.Setenv(metrics.HandlerEnv, handler.String())
			// used for testing
			t.Setenv("NEW_RELIC_LICENSE_KEY", gofakeit.Sentence(40)[0:40])

			_, err := metrics.NewFromEnv(ctx, config.NewBuildInfo(config.DefaultVersion, config.DefaultCommit, config.AppName, config.DefaultDate))
			Nil(t, err)
		})
	}
}
