package metrics

import (
	"context"
	"fmt"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/synapsecns/sanguine/core/config"
	"os"
	"sync"
)

type newRelicHandler struct {
	app       *newrelic.Application
	startMux  sync.Mutex
	buildInfo config.BuildInfo
}

func NewRelicMetricsHandler(buildInfo config.BuildInfo) Handler {
	return &newRelicHandler{
		buildInfo: buildInfo,
	}
}

func (n *newRelicHandler) Start(_ context.Context) (err error) {
	n.startMux.Lock()
	defer n.startMux.Unlock()
	if n.app == nil {
		n.app, err = newrelic.NewApplication(
			newrelic.ConfigAppName(n.buildInfo.Name()),
			newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
			newrelic.ConfigAppLogForwardingEnabled(true),
			newrelic.ConfigCodeLevelMetricsEnabled(true),
			func(c *newrelic.Config) {
				c.Labels = map[string]string{
					"version": n.buildInfo.Version(),
					"commit":  n.buildInfo.Commit(),
				}
			},
			// optional overrides
			newrelic.ConfigFromEnvironment(),
		)
		if err != nil {
			return fmt.Errorf("could not create new relic application: %w", err)
		}
	}

	return nil
}
