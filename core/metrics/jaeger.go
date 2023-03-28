package metrics

import (
	"context"
	"fmt"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/dockerutil"
	"github.com/synapsecns/sanguine/core/processlog"
	"go.opentelemetry.io/otel/exporters/jaeger"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"os"
	"runtime"
	"sync"
	"testing"
	"time"
)

type jaegerHandler struct {
	*baseHandler
	buildInfo config.BuildInfo
	exporter  *jaeger.Exporter
}

// NewJaegerHandler creates a new jaeger handler for handling jaeger traces.
// the JAEGER_ENDPOINT environment variable must be set for this to work.
func NewJaegerHandler(buildInfo config.BuildInfo) Handler {
	return &jaegerHandler{
		buildInfo: buildInfo,
	}
}

const jaegerEnv = "JAEGER_ENDPOINT"

func (j *jaegerHandler) Start(ctx context.Context) (err error) {
	endpoint := os.Getenv(jaegerEnv)
	if endpoint == "" {
		return fmt.Errorf("could not get jaeger endpoint from env")
	}
	j.exporter, err = jaeger.New(
		jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(endpoint)),
	)
	if err != nil {
		return fmt.Errorf("could not create jaeger exporter: %w", err)
	}

	j.baseHandler = newBaseHandler(j.buildInfo, tracesdk.WithSyncer(j.exporter), tracesdk.WithSampler(tracesdk.AlwaysSample()))
	err = j.baseHandler.Start(ctx)
	if err != nil {
		return fmt.Errorf("could not start base handler: %w", err)
	}

	return nil
}

var keepAliveOnFailure = time.Minute * 10

var testMux sync.Mutex

// debugjaeger.
const debugJaeger = true

// SetupTestJaeger creates a new test jaeger instance. If the test fails, the instance is kept alive for 5 minutes.
// we also allow a GLOBAL_jaeger env var to be set to a jaeger url to send all traces to in order to avoid having to boot for long running tests.
func SetupTestJaeger(tb testing.TB) {
	tb.Helper()
	// make sure we don't setup two
	testMux.Lock()
	defer testMux.Unlock()

	ctx, cancel := context.WithCancel(context.Background())

	if os.Getenv(jaegerEnv) != "" {
		cancel()
		return
	}

	pool, err := dockertest.NewPool("")
	assert.Nil(tb, err)

	runOptions := &dockertest.RunOptions{
		Repository:   "jaegertracing/all-in-one",
		Tag:          "latest",
		ExposedPorts: []string{"14268", "16686"},
	}
	resource, err := pool.RunWithOptions(runOptions, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{Name: "no"}
	})
	assert.Nil(tb, err)

	tb.Setenv(jaegerEnv, fmt.Sprintf("http://localhost:%s/api/traces", resource.GetPort("14268/tcp")))

	if !debugJaeger {
		err = resource.Expire(uint(keepAliveOnFailure.Seconds()))
		assert.Nil(tb, err)
	}

	go func() {
		_ = dockerutil.TailContainerLogs(dockerutil.WithContext(ctx), dockerutil.WithPool(pool), dockerutil.WithResource(resource), dockerutil.WithFollow(true), dockerutil.WithCallback(func(ctx context.Context, metadata processlog.LogMetadata) {
			logger.Warnf(
				"serving jaeger instance on http://localhost:%s. Container logs will be saved to %s %s", resource.GetPort("16686/tcp"), metadata.LogDir(),
				fmt.Sprintf("if you want to persist this session, set debugjaeger to true in %s (currently %t), then set the JAEGER_ENDPOINT to %s", getCurrentFile(), debugJaeger, os.Getenv(jaegerEnv)),
			)
		}))
	}()

	tb.Cleanup(func() {
		defer cancel()
		// TODO: move me
		if tb.Failed() {
			logger.Warnf("Test failed, will temporarily continue serving jaeger instance on http://localhost:%s", resource.GetPort("16686/tcp"))
		} else if !debugJaeger {
			// TODO: uncomment me
			_ = pool.Purge(resource)
		}
	})
}

// TODO: clean me up with runtime.caller(2).
func getCurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	return file
}
