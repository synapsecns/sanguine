package localmetrics

import (
	"context"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
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
	if err != nil {
		tb.Logf("Failed to create Docker pool: %v", err)
		tb.Fatal(err)
	}

	tj.logDir = filet.TmpDir(tb, "")
	tb.Logf("Created log directory at: %s", tj.logDir)

	// Ensure Docker daemon is responsive
	if err := tj.pool.Client.Ping(); err != nil {
		tb.Logf("Docker daemon not responsive: %v", err)
		tb.Fatal(err)
	}

	// Clean up any existing containers with our labels before starting
	containers, err := tj.pool.Client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		tb.Logf("Failed to list containers: %v", err)
	} else {
		for _, container := range containers {
			if container.Labels[runIDLabel] == tj.runID || container.Labels[appLabel] == "jaeger" ||
			   container.Labels[appLabel] == "pyroscope" || container.Labels[appLabel] == "jaeger-ui" {
				tb.Logf("Removing existing container: %s", container.ID)
				err := tj.pool.Client.RemoveContainer(docker.RemoveContainerOptions{
					ID:            container.ID,
					Force:         true,
					RemoveVolumes: true,
				})
				if err != nil {
					tb.Logf("Error removing container %s: %v", container.ID, err)
				}
			}
		}
	}

	// Start containers with retry logic
	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			tb.Logf("Retry attempt %d/%d after Docker cleanup", attempt+1, maxRetries)
			// Exponential backoff
			time.Sleep(time.Duration(attempt*2) * time.Second)

			// Force cleanup between attempts
			tj.purgeResources()
			if tj.network != nil {
				_ = tj.network.Close()
				tj.network = nil
			}
		}

		// Start containers sequentially with health checks
		tj.jaegerResource = tj.StartJaegerServer(ctx)
		if tj.jaegerResource == nil || tj.jaegerResource.Resource == nil {
			tb.Logf("Failed to start Jaeger container, retrying...")
			continue
		}

		// Verify Jaeger is healthy before continuing with more lenient parameters
		if err := retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerEndpoint)),
			retry.WithMax(time.Second*2), retry.WithMaxAttempts(30)); err != nil {
			tb.Logf("Jaeger health check failed: %v", err)
			continue
		}

		tj.pyroscopeResource = tj.StartPyroscopeServer(ctx)
		if tj.pyroscopeResource == nil || tj.pyroscopeResource.Resource == nil {
			tb.Logf("Failed to start Pyroscope container, retrying...")
			continue
		}

		// Verify Pyroscope is healthy before continuing with more lenient parameters
		if err := retry.WithBackoff(ctx, checkURL(os.Getenv(internal.PyroscopeEndpoint)),
			retry.WithMax(time.Second*2), retry.WithMaxAttempts(30)); err != nil {
			tb.Logf("Pyroscope health check failed: %v", err)
			continue
		}

		tj.jaegerPyroscopeUIResource = tj.StartJaegerPyroscopeUI(ctx)
		if tj.jaegerPyroscopeUIResource == nil || tj.jaegerPyroscopeUIResource.Resource == nil {
			tb.Logf("Failed to start Jaeger Pyroscope UI container, retrying...")
			continue
		}

		// All containers started successfully
		break
	}

	// Check if all containers are running
	if tj.jaegerResource == nil || tj.pyroscopeResource == nil || tj.jaegerPyroscopeUIResource == nil {
		tb.Fatal("Failed to start all required containers after maximum retries")
	}

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

// getNetworks gets the networks to be associated with each container.
func (j *testJaeger) getNetworks() []*dockertest.Network {
	// no need to hit the mutex if no network is required
	if !j.cfg.requiresNetwork {
		return []*dockertest.Network{}
	}

	j.networkMux.Lock()
	defer j.networkMux.Unlock()

	// Clean up any existing network with the same runID
	if networks, err := j.pool.Client.ListNetworks(); err == nil {
		for _, network := range networks {
			if network.Labels[runIDLabel] == j.runID {
				// List containers connected to this network
				containers, err := j.pool.Client.ListContainers(docker.ListContainersOptions{
					All:     true,
					Filters: map[string][]string{"network": {network.ID}},
				})
				if err != nil {
					j.tb.Logf("Error listing containers for network %s: %v", network.ID, err)
					continue
				}

				// Disconnect containers from network before removal
				for _, container := range containers {
					err := j.pool.Client.DisconnectNetwork(network.ID, docker.NetworkConnectionOptions{
						Container: container.ID,
						Force:     true,
					})
					if err != nil {
						j.tb.Logf("Error disconnecting container %s from network %s: %v", container.ID, network.ID, err)
					}
				}

				// Now try to remove the network
				if err := j.pool.Client.RemoveNetwork(network.ID); err != nil {
					j.tb.Logf("Error removing network %s: %v", network.ID, err)
				}
			}
		}
	}

	var err error
	j.network, err = j.pool.CreateNetwork(j.runID, func(config *docker.CreateNetworkOptions) {
		config.Driver = "bridge"
		config.Labels = map[string]string{
			runIDLabel: j.runID,
		}
	})

	if err != nil {
		j.tb.Logf("Failed to create network: %v", err)
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
	allResources := []*uiResource{j.jaegerResource, j.pyroscopeResource, j.jaegerPyroscopeUIResource}
	for _, resource := range allResources {
		if resource == nil || resource.Resource == nil {
			j.tb.Logf("Resource is nil: %v", resource)
			continue
		}
		dockerizedResources = append(dockerizedResources, resource.Resource)
	}
	return dockerizedResources
}

// purgeResources purges the resources from the pool.
func (j *testJaeger) purgeResources() {
	var wg sync.WaitGroup
	resources := j.getDockerizedResources()

	// First, stop all containers to release network resources
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		wg.Add(1)
		go func(r *dockertest.Resource) {
			defer wg.Done()
			if err := j.pool.Client.StopContainer(r.Container.ID, 1); err != nil {
				j.tb.Logf("Error stopping container %s: %v", r.Container.ID, err)
			}
		}(resource)
	}
	wg.Wait()

	// Then remove containers and their volumes
	for _, resource := range resources {
		if resource == nil {
			continue
		}
		wg.Add(1)
		go func(r *dockertest.Resource) {
			defer wg.Done()
			if err := j.pool.Purge(r); err != nil {
				j.tb.Logf("Error purging resource: %v", err)
			}
		}(resource)
	}
	wg.Wait()

	// Finally, clean up any networks associated with our runID
	if networks, err := j.pool.Client.ListNetworks(); err == nil {
		for _, network := range networks {
			if network.Labels[runIDLabel] == j.runID {
				// Add a small delay before network removal
				time.Sleep(time.Second * 2)
				if err := j.pool.Client.RemoveNetwork(network.ID); err != nil {
					j.tb.Logf("Error removing network %s: %v", network.ID, err)
				}
			}
		}
	}
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
