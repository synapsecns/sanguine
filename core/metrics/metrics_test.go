package metrics_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	. "github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics"
	"testing"
)

func TestSetupFromEnv(t *testing.T) {
	for _, handler := range metrics.AllHandlerTypes {
		handler := handler // capture func literal
		assert.NotPanics(t, func() {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			t.Setenv(metrics.HandlerEnv, handler.String())

			err := metrics.NewFromEnv(ctx, config.NewBuildInfo(config.DefaultVersion, config.DefaultCommit, config.AppName, config.DefaultDate))
			Nil(t, err)
		})
	}
}
