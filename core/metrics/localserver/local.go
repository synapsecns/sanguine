package localserver

import (
	"context"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/stretchr/testify/assert"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/retry"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"testing"
	"time"
)

var keepAliveOnFailure = time.Minute * 10

var testMux sync.Mutex

// debugLocal should be turned on to persist containers.
var debugLocal = true

const appLabel = "app"
const runIDLabel = "runID"

type testJaeger struct {
	tb                        testing.TB
	pool                      *dockertest.Pool
	logDir                    string
	runID                     string
	pyroscopeResource         *uiResource
	jaegerResource            *uiResource
	jaegerPyroscopeUIResource *uiResource
	networkMux                sync.Mutex
	// this should not be used directly, use getNetwork
	network *dockertest.Network
}

// StartServer starts a local jaeger server for testing.
func StartServer(parentCtx context.Context, tb testing.TB) {
	startServer(parentCtx, tb)
}

// startServer starts a local jaeger server for testing.
// this is a separate function so we can export testJaeger for testing.
func startServer(parentCtx context.Context, tb testing.TB) *testJaeger {
	tb.Helper()
	// create the test jaegar instance
	tj := testJaeger{
		tb:    tb,
		runID: gofakeit.UUID(),
	}

	tb.Helper()
	// make sure we don't setup two
	testMux.Lock()
	defer testMux.Unlock()

	// if we have a global jaegerResource env var, don't setup a local one
	ctx, cancel := context.WithCancel(parentCtx)

	// create the pool
	var err error
	tj.pool, err = dockertest.NewPool("")
	assert.Nil(tb, err)

	tj.logDir = filet.TmpDir(tb, "")

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		tj.jaegerResource = tj.StartJaegerServer(ctx)
	}()

	go func() {
		defer wg.Done()
		tj.pyroscopeResource = tj.StartPyroscopeServer(ctx)
	}()
	go func() {
		defer wg.Done()
		tj.jaegerPyroscopeUIResource = tj.StartJaegerPyroscopeUI(ctx)
	}()
	wg.Wait()

	logger.Warnf(tj.buildLogMessage(true))

	tb.Cleanup(func() {
		defer cancel()
		// TODO: move me
		if tb.Failed() {
			logger.Warn("Test failed, will temporarily continue serving \n" + tj.buildLogMessage(false))
		} else if !debugLocal {
			tj.purgeResources()
			if tj.network != nil {
				_ = tj.network.Close()
			}
		}
	})

	return &tj
}

// createDockerResources creates the docker resources.
// this must be called by each container.
func (j *testJaeger) getNetwork() *dockertest.Network {
	j.networkMux.Lock()
	defer j.networkMux.Unlock()

	if j.network != nil {
		return j.network
	}

	var err error
	j.network, err = j.pool.CreateNetwork(j.runID, func(config *docker.CreateNetworkOptions) {
		config.Driver = "bridge"
		config.Labels = map[string]string{
			runIDLabel: j.runID,
		}
	})

	if err != nil {
		j.tb.Fatal(err)
	}

	return j.network
}

// buildLogMessage builds a log message for the test jaeger instance.
func (j *testJaeger) buildLogMessage(includeAuxiliary bool) string {
	var messages []string
	messages = append(messages, fmt.Sprintf("jaeger ui: %s", os.Getenv(internal.JaegerUiEndpoint)))
	messages = append(messages, fmt.Sprintf("pyroscope ui: %s", j.pyroscopeResource.uiURL))

	var bootMessages []string
	if len(j.getDockerizedResources()) > 0 {
		bootMessages = append(bootMessages, fmt.Sprintf("Container logs will be saved to %s", j.logDir))
	}

	if j.jaegerResource != nil {
		bootMessages = append(bootMessages, fmt.Sprintf("if you want to persist this session, set debugLocal to true in %s (currently %t), then set the JAEGER_ENDPOINT to %s", getCurrentFile(), debugLocal, os.Getenv(internal.JaegerEndpoint)))
	}

	if len(bootMessages) > 0 && includeAuxiliary {
		messages = append(messages, strings.Join(bootMessages, " "))
	}

	return strings.Join(messages, "\n")
}

// getDockerizedResources gets all resources that have been dockerized by this process.
func (j *testJaeger) getDockerizedResources() (dockerizedResources []*dockertest.Resource) {
	allResources := []*uiResource{j.jaegerResource, j.pyroscopeResource}
	for _, resource := range allResources {
		if resource.Resource == nil {
			return nil
		}
		dockerizedResources = append(dockerizedResources, resource.Resource)
	}
	return dockerizedResources
}

// purgeResources purges the resources from the pool.
func (j *testJaeger) purgeResources() {
	var wg sync.WaitGroup
	resources := j.getDockerizedResources()

	wg.Add(len(resources))
	for _, resource := range resources {
		go func(resource *dockertest.Resource) {
			defer wg.Done()
			_ = j.pool.Purge(resource)
		}(resource)
	}
	wg.Wait()
}

// uiResource is a wrapper around dockertest.Resource that logs the container logs to a file.
type uiResource struct {
	// Resource is the underlying dockertest resource.
	// this is not guaranteed to be set.
	*dockertest.Resource
	uiURL string
}

// checkURL is a helper function that checks if a url is alive.
// it does not check the status code.
func checkURL(url string) retry.RetryableFunc {
	return func(ctx context.Context) error {
		client := http.DefaultClient
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		if err != nil {
			return fmt.Errorf("could not create request: %w", err)
		}

		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("could not get response: %w", err)
		}

		if resp != nil {
			_ = resp.Body.Close()
		}

		return nil
	}
}

// TODO: clean me up with runtime.caller(2).
func getCurrentFile() string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		return "unknown"
	}
	return file
}
