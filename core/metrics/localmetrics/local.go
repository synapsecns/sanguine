package localmetrics

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/Flaque/filet"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"github.com/synapsecns/sanguine/core/metrics/internal"
	"github.com/synapsecns/sanguine/core/retry"
	"io"
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

	// create the pool with retry and increased timeout
	var err error
	err = retry.WithBackoff(ctx, func(ctx context.Context) error {
		tj.pool, err = dockertest.NewPool("")
		if err != nil {
			tb.Logf("Failed to create Docker pool: %v", err)
			return err
		}

		// Configure pool timeout
		tj.pool.MaxWait = time.Second * 60

		// Ensure Docker daemon is responsive
		if err := tj.pool.Client.Ping(); err != nil {
			tb.Logf("Docker daemon not responsive: %v", err)
			return err
		}
		return nil
	},
		retry.WithMax(time.Second*15),
		retry.WithMaxAttempts(5))
	if err != nil {
		tb.Fatal(err)
	}

	tj.logDir = filet.TmpDir(tb, "")
	tb.Logf("Created log directory at: %s", tj.logDir)

	// Clean up existing resources thoroughly
	if err := tj.purgeAllResources(); err != nil {
		tb.Logf("Warning: Failed to clean up existing resources: %v", err)
	}

	// Wait for cleanup to complete
	time.Sleep(time.Second * 3)

	// Start containers with improved retry logic
	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		if attempt > 0 {
			tb.Logf("Retry attempt %d/%d for container startup", attempt+1, maxRetries)
			// Add exponential backoff between retries
			time.Sleep(time.Duration(attempt*5) * time.Second)

			// Thorough cleanup between attempts
			tj.purgeResources()
			if tj.network != nil {
				_ = tj.network.Close()
				tj.network = nil
			}
			// Wait for resources to be fully released
			time.Sleep(time.Second * 5)
		}

		// Create network first if needed
		if tj.cfg.requiresNetwork {
			networks := tj.getNetworks()
			if len(networks) == 0 {
				// Create network with retry
				err := retry.WithBackoff(ctx, func(ctx context.Context) error {
					networks = tj.getNetworks()
					if len(networks) == 0 {
						network, err := tj.pool.CreateNetwork(fmt.Sprintf("test-network-%s", tj.runID))
						if err != nil {
							tb.Logf("Failed to create network: %v", err)
							return err
						}
						tj.network = network
					}
					return nil
				},
					retry.WithMax(time.Second*5),
					retry.WithMaxAttempts(3))
				if err != nil {
					tb.Logf("Failed to create network after retries: %v", err)
					continue
				}
			}
			// Wait for network to stabilize
			time.Sleep(time.Second * 3)
		}

		// Start Jaeger first
		tj.jaegerResource = tj.StartJaegerServer(ctx)
		if tj.jaegerResource == nil || tj.jaegerResource.Resource == nil {
			tb.Logf("Failed to start Jaeger container")
			continue
		}

		// Wait for Jaeger to be fully ready with increased timeout
		if err := retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerEndpoint)),
			retry.WithMax(time.Second*10),
			retry.WithMaxAttempts(20)); err != nil {
			tb.Logf("Jaeger health check failed: %v", err)
			continue
		}

		// Start Pyroscope if enabled
		if tj.cfg.enablePyroscope {
			// Wait for network stability
			time.Sleep(time.Second * 5)

			tj.pyroscopeResource = tj.StartPyroscopeServer(ctx)
			if tj.pyroscopeResource == nil || tj.pyroscopeResource.Resource == nil {
				tb.Logf("Failed to start Pyroscope container")
				continue
			}

			// Wait for Pyroscope with increased timeout
			if err := retry.WithBackoff(ctx, checkURL(os.Getenv(internal.PyroscopeEndpoint)),
				retry.WithMax(time.Second*10),
				retry.WithMaxAttempts(20)); err != nil {
				tb.Logf("Pyroscope health check failed: %v", err)
				continue
			}
		}

		// Start UI component if needed
		if tj.cfg.enablePyroscopeJaeger {
			// Wait for network stability
			time.Sleep(time.Second * 2)

			tj.jaegerPyroscopeUIResource = tj.StartJaegerPyroscopeUI(ctx)
			if tj.jaegerPyroscopeUIResource == nil || tj.jaegerPyroscopeUIResource.Resource == nil {
				tb.Logf("Failed to start Jaeger Pyroscope UI container")
				continue
			}

			// Wait for UI with increased timeout
			if err := retry.WithBackoff(ctx, checkURL(os.Getenv(internal.JaegerUIEndpoint)),
				retry.WithMax(time.Second*5),
				retry.WithMaxAttempts(15)); err != nil {
				tb.Logf("Jaeger UI health check failed: %v", err)
				continue
			}
		}

		// All containers started successfully
		break
	}

	// Verify final state based on configuration
	if tj.jaegerResource == nil {
		tb.Fatal("Failed to start Jaeger container after maximum retries")
	}
	if tj.cfg.enablePyroscope && tj.pyroscopeResource == nil {
		tb.Fatal("Failed to start Pyroscope container after maximum retries")
	}
	if tj.cfg.enablePyroscopeJaeger && tj.jaegerPyroscopeUIResource == nil {
		tb.Fatal("Failed to start Jaeger Pyroscope UI container after maximum retries")
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

	// Try to use existing jaeger-test-net network first
	networks, err := j.pool.Client.ListNetworks()
	if err != nil {
		j.tb.Logf("Error listing networks: %v", err)
	} else {
		for _, network := range networks {
			if network.Name == "jaeger-test-net" {
				// Found existing network, wrap it in dockertest.Network
				networkPtr := &network
				return []*dockertest.Network{{Network: networkPtr}}
			}
		}
	}

	// Clean up existing test networks with retry
	err = retry.WithBackoff(context.Background(), func(ctx context.Context) error {
		networks, err := j.pool.Client.ListNetworks()
		if err != nil {
			j.tb.Logf("Error listing networks: %v", err)
			return err
		}

		for _, network := range networks {
			if network.Labels[runIDLabel] == j.runID {
				// List and disconnect containers
				containers, err := j.pool.Client.ListContainers(docker.ListContainersOptions{
					All:     true,
					Filters: map[string][]string{"network": {network.ID}},
				})
				if err != nil {
					j.tb.Logf("Error listing containers for network %s: %v", network.ID, err)
					continue
				}

				for _, container := range containers {
					err := j.pool.Client.DisconnectNetwork(network.ID, docker.NetworkConnectionOptions{
						Container: container.ID,
						Force:     true,
					})
					if err != nil {
						j.tb.Logf("Error disconnecting container %s from network %s: %v", container.ID, network.ID, err)
					}
				}

				// Wait for disconnections to complete
				time.Sleep(time.Second)

				if err := j.pool.Client.RemoveNetwork(network.ID); err != nil {
					j.tb.Logf("Error removing network %s: %v", network.ID, err)
					return err
				}
			}
		}
		return nil
	},
		retry.WithMax(time.Second*2),
		retry.WithMaxAttempts(3))

	if err != nil {
		j.tb.Logf("Warning: Failed to clean up existing networks: %v", err)
	}

	// Use existing jaeger-test-net network if available
	if networkExists("jaeger-test-net") {
		network, err := j.pool.Client.NetworkInfo("jaeger-test-net")
		if err == nil {
			j.tb.Log("Using existing jaeger-test-net network")
			return []*dockertest.Network{{Network: network}}
		}
		j.tb.Logf("Error getting jaeger-test-net info: %v", err)
	}

	// Create new network with retry and specific subnet
	err = retry.WithBackoff(context.Background(), func(ctx context.Context) error {
		var createErr error
		j.network, createErr = j.pool.CreateNetwork(j.runID, func(config *docker.CreateNetworkOptions) {
			config.Driver = "bridge"
			config.IPAM = &docker.IPAMOptions{
				Config: []docker.IPAMConfig{{
					Subnet:  "172.20.0.0/16",
					Gateway: "172.20.0.1",
				}},
			}
			config.Labels = map[string]string{
				runIDLabel: j.runID,
			}
			config.CheckDuplicate = true
		})
		if createErr != nil {
			j.tb.Logf("Failed to create network, retrying: %v", createErr)
			return createErr
		}

		// Verify network exists
		_, err := j.pool.Client.NetworkInfo(j.network.Network.ID)
		if err != nil {
			j.tb.Logf("Network verification failed: %v", err)
			return err
		}
		return nil
	},
		retry.WithMax(time.Second*2),
		retry.WithMaxAttempts(3))

	if err != nil {
		j.tb.Logf("Failed to create network after retries: %v", err)
		j.tb.Fatal(err)
	}

	j.tb.Logf("Successfully created network: %s", j.network.Network.ID)
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
	var cleanupMutex sync.Mutex
	resources := j.getDockerizedResources()

	// Clean up ports first to ensure they're released
	if err := j.cleanupPorts(); err != nil {
		j.tb.Logf("Warning: Failed to clean up ports: %v", err)
	}

	// First, stop all containers sequentially to avoid race conditions
	for _, resource := range resources {
		if resource == nil || resource.Container == nil {
			continue
		}
		cleanupMutex.Lock()
		// Stop container with timeout and proper error handling
		if err := j.pool.Client.StopContainer(resource.Container.ID, uint(5)); err != nil {
			if !strings.Contains(err.Error(), "No such container") {
				j.tb.Logf("Warning: Failed to stop container %s: %v", resource.Container.ID, err)
			}
		}
		cleanupMutex.Unlock()
		// Wait for container to stop
		time.Sleep(time.Second)
	}

	// Then remove containers and their volumes sequentially
	for _, resource := range resources {
		if resource == nil || resource.Container == nil {
			continue
		}
		cleanupMutex.Lock()
		opts := docker.RemoveContainerOptions{
			ID:            resource.Container.ID,
			Force:         true,
			RemoveVolumes: true,
		}
		if err := j.pool.Client.RemoveContainer(opts); err != nil {
			if !strings.Contains(err.Error(), "No such container") &&
				!strings.Contains(err.Error(), "removal of container") {
				j.tb.Logf("Warning: Failed to remove container %s: %v", resource.Container.ID, err)
			}
		}
		cleanupMutex.Unlock()
		// Wait for removal to complete
		time.Sleep(time.Second)
	}

	// Finally, clean up networks with proper synchronization
	if networks, err := j.pool.Client.ListNetworks(); err == nil {
		for _, network := range networks {
			if network.Labels[runIDLabel] == j.runID {
				cleanupMutex.Lock()
				// Ensure all containers are disconnected from network
				containers, _ := j.pool.Client.ListContainers(docker.ListContainersOptions{
					All:     true,
					Filters: map[string][]string{"network": {network.ID}},
				})
				for _, container := range containers {
					if err := j.pool.Client.DisconnectNetwork(network.ID, docker.NetworkConnectionOptions{
						Container: container.ID,
						Force:     true,
					}); err != nil {
						j.tb.Logf("Warning: Failed to disconnect container %s from network %s: %v", container.ID, network.ID, err)
					}
				}
				// Wait for disconnections to complete
				time.Sleep(time.Second * 2)
				if err := j.pool.Client.RemoveNetwork(network.ID); err != nil {
					if !strings.Contains(err.Error(), "not found") {
						j.tb.Logf("Warning: Failed to remove network %s: %v", network.ID, err)
					}
				}
				cleanupMutex.Unlock()
			}
		}
	}

	// Final port cleanup and verification
	if err := j.cleanupPorts(); err != nil {
		j.tb.Logf("Warning: Failed in final port cleanup: %v", err)
	}
	// Wait for all cleanup operations to complete
	time.Sleep(time.Second * 3)
}

// uiResource is a wrapper around dockertest.Resource that logs the container logs to a file.
type uiResource struct {
	// Resource is the underlying dockertest resource.
	// this is not guaranteed to be set.
	*dockertest.Resource
	uiURL string
}

// checkURL is a helper function that checks if a url is alive.
// it accepts 405 Method Not Allowed for trace endpoints.
func checkURL(url string) retry.RetryableFunc {
	return func(ctx context.Context) error {
		if url == "" {
			return fmt.Errorf("empty URL provided")
		}

		client := &http.Client{
			Timeout: time.Second * 5,
			Transport: &http.Transport{
				DisableKeepAlives: true,
				MaxIdleConns:      1,
				IdleConnTimeout:   time.Second,
				TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			},
		}

		// Determine HTTP method based on endpoint
		method := http.MethodGet
		if strings.Contains(url, "/api/traces") {
			method = http.MethodPost
		}

		req, err := http.NewRequestWithContext(ctx, method, url, nil)
		if err != nil {
			return fmt.Errorf("failed to create request for %s: %v", url, err)
		}

		resp, err := client.Do(req)
		if err != nil {
			if os.IsTimeout(err) {
				return fmt.Errorf("timeout connecting to %s: %v", url, err)
			}
			if strings.Contains(err.Error(), "connection refused") {
				return fmt.Errorf("connection refused to %s - container may not be ready: %v", url, err)
			}
			return fmt.Errorf("failed to connect to %s: %v", url, err)
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			return fmt.Errorf("server error at %s: status=%d", url, resp.StatusCode)
		}

		// For trace endpoints, accept 405 Method Not Allowed as valid
		if resp.StatusCode >= 400 {
			if strings.Contains(url, "/api/traces") && resp.StatusCode == http.StatusMethodNotAllowed {
				return nil
			}
			return fmt.Errorf("client error at %s: status=%d", url, resp.StatusCode)
		}

		// Read a small portion of the body to verify the response
		_, err = io.ReadAll(io.LimitReader(resp.Body, 1024))
		if err != nil {
			return fmt.Errorf("failed to read response body from %s: %v", url, err)
		}

		return nil
	}
}
