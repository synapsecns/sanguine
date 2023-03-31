package localserver

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/processlog"
	"github.com/synapsecns/sanguine/core/retry"
	"os"
	"time"
)

// StartJaegerServer starts a new jaeger instance.
func (j *testJaeger) StartJaegerServer(ctx context.Context) *uiResource {
	if core.HasEnv(internal.JAEGER_ENDPOINT) && !core.HasEnv(internal.JAEGER_UI_ENDPOINT) {
		j.tb.Fatalf("%s is set but %s is not, please remove %s or set %s", internal.JAEGER_ENDPOINT, internal.JAEGER_UI_ENDPOINT, internal.JAEGER_ENDPOINT, internal.JAEGER_UI_ENDPOINT)
		return nil
	}

	if core.HasEnv(internal.JAEGER_ENDPOINT) {
		return &uiResource{
			Resource: nil,
			uiURL:    os.Getenv(internal.JAEGER_UI_ENDPOINT),
		}
	}

	runOptions := &dockertest.RunOptions{
		Repository:   "jaegertracing/all-in-one",
		Tag:          "latest",
		ExposedPorts: []string{"14268", "16686"},
		Labels: map[string]string{
			appLabel:   "jaeger",
			runIDLabel: j.runID,
		},
	}
	resource, err := j.pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(j.tb, err)

	j.tb.Setenv(internal.JAEGER_ENDPOINT, fmt.Sprintf("http://localhost:%s/api/traces", resource.GetPort("14268/tcp")))
	j.tb.Setenv(internal.JAEGER_UI_ENDPOINT, fmt.Sprintf("http://localhost:%s", resource.GetPort("16686/tcp")))

	if !debugLocal {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		assert.Nil(j.tb, err)
	}

	logResourceChan := make(chan *uiResource, 1)

	go func() {
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithPool(j.pool), dockerutil.WithProcessLogOptions(processlog.WithLogDir(j.logDir), processlog.WithLogFileName("jaeger")), dockerutil.WithFollow(true),
			dockerutil.WithResource(resource), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
				select {
				case <-ctx.Done():
					return
				case logResourceChan <- &uiResource{
					Resource: resource,
					uiURL:    os.Getenv(internal.JAEGER_UI_ENDPOINT),
				}:
					return
				}
			}))
	}()

	// make sure client is alive
	err = retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JAEGER_ENDPOINT)), retry.WithMax(time.Millisecond*10), retry.WithMax(time.Minute))
	if err != nil {
		return nil
	}

	select {
	case <-ctx.Done():
		return nil
	case logResource := <-logResourceChan:
		return logResource
	}
}
