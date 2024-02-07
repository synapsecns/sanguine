package localmetrics

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
	"strings"
	"sync"
	"testing"
	"time"
)

var keepAliveOnFailure = time.Minute * 10

var testMux sync.Mutex

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
	// this is not guaranteed to be set and only required for pyroscope-jaeger
	network *dockertest.Network
	cfg     *config
}

// startServer starts a local jaeger server for testing.
// this is a separate function so we can export testJaeger for testing.
func startServer(parentCtx context.Context, tb testing.TB, options ...Option) *testJaeger {
	tb.Helper()

	// create the test jaegar instance
	tj := testJaeger{
		tb:    tb,
		runID: gofakeit.UUID(),
		cfg:   makeConfig(options),
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
		// Do not keep containers on ci.
		if tb.Failed() && os.Getenv("CI") == "" {
			logger.Warn("Test failed, will temporarily continue serving \n" + tj.buildLogMessage(false))
		} else if !tj.cfg.keepContainers {
			tj.purgeResources()
			if tj.network != nil {
				_ = tj.network.Close()
			}
		}
	})

	return &tj
}

// getNetworks gets the networks to be associaed with each container.
func (j *testJaeger) getNetworks() []*dockertest.Network {
	// no need to hit the mutex if no network is required
	if !j.cfg.requiresNetwork {
		return []*dockertest.Network{}
	}

	j.networkMux.Lock()
	defer j.networkMux.Unlock()

	if j.network != nil {
		return []*dockertest.Network{j.network}
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

	return []*dockertest.Network{j.network}
}

// buildLogMessage builds a log message for the test jaeger instance.
func (j *testJaeger) buildLogMessage(includeAuxiliary bool) string {
	var messages []string
	messages = append(messages, fmt.Sprintf("jaeger ui: %s", os.Getenv(internal.JaegerUIEndpoint)))
	if j.cfg.enablePyroscope {
		messages = append(messages, fmt.Sprintf("pyroscope ui: %s", os.Getenv(internal.PyroscopeEndpoint)))
	}

	var bootMessages []string
	if len(j.getDockerizedResources()) > 0 {
		bootMessages = append(bootMessages, fmt.Sprintf("Container logs will be saved to %s", j.logDir))
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
		if resource == nil || resource.Resource == nil {
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
			_ = j.pool.Purge(resource)
			wg.Done()
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
