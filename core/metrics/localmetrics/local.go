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

	// Ensure ports are free before starting containers
	if err := tj.cleanupPorts(); err != nil {
		tb.Logf("Warning: Failed to clean up ports: %v", err)
		// Even if port cleanup fails, continue with container startup
		// as the container itself might handle port conflicts
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
			if err := tj.purgeResources(); err != nil {
				tb.Logf("Warning: Failed to clean up resources between attempts: %v", err)
			}
			if tj.network != nil && tj.pool != nil && tj.pool.Client != nil {
				if err := tj.pool.Client.RemoveNetwork(tj.network.Network.ID); err != nil {
					if !strings.Contains(err.Error(), "not found") {
						tb.Logf("Warning: Failed to remove network %s: %v", tj.network.Network.ID, err)
					}
				}
				tj.network = nil
			}
			// Wait for resources to be fully released
			time.Sleep(time.Second * 5)
		}

		// Create network first if needed
		if tj.cfg.requiresNetwork {
			// Create network with retry
			err := retry.WithBackoff(ctx, func(ctx context.Context) error {
				networkName := fmt.Sprintf("test-network-%s", tj.runID)

				// Clean up all stale test networks first
				networks, err := tj.pool.Client.ListNetworks()
				if err != nil {
					tb.Logf("Failed to list networks: %v", err)
					return err
				}

				// Clean up any existing test networks to prevent subnet conflicts
				for _, n := range networks {
					if strings.HasPrefix(n.Name, "test-network-") {
						if err := tj.pool.Client.RemoveNetwork(n.ID); err != nil {
							if !strings.Contains(err.Error(), "not found") {
								tb.Logf("Warning: Failed to remove existing network %s: %v", n.ID, err)
							}
						}
					}
				}

				// Wait for network cleanup to complete
				time.Sleep(time.Second * 3)

				// Create new network with specific subnet
				network, err := tj.pool.Client.CreateNetwork(docker.CreateNetworkOptions{
					Name: networkName,
					Labels: map[string]string{
						runIDLabel: tj.runID,
					},
					Driver: "bridge",
					IPAM: &docker.IPAMOptions{
						Driver: "default",
						Config: []docker.IPAMConfig{
							{
								Subnet:  "172.21.0.0/16",
								Gateway: "172.21.0.1",
							},
						},
					},
					CheckDuplicate: true,
					EnableIPv6:     false,
				})
				if err != nil {
					tb.Logf("Failed to create network: %v", err)
					// Try to clean up the failed network
					if cleanupErr := tj.pool.Client.RemoveNetwork(network.ID); cleanupErr != nil {
						tb.Logf("Warning: Failed to clean up failed network: %v", cleanupErr)
					}
					return err
				}

				// Verify network was created
				createdNetwork, err := tj.pool.Client.NetworkInfo(network.ID)
				if err != nil {
					tb.Logf("Failed to verify network creation: %v", err)
					return err
				}

				// Additional network validation
				if createdNetwork == nil {
					return fmt.Errorf("network created but info returned nil")
				}

				// Verify network is in expected state
				if createdNetwork.Driver != "bridge" {
					return fmt.Errorf("network driver mismatch: expected bridge, got %s", createdNetwork.Driver)
				}

				tj.network = &dockertest.Network{Network: createdNetwork}
				return nil
			},
				retry.WithMax(time.Second*15),
				retry.WithMaxAttempts(5))

			if err != nil {
				tb.Logf("Failed to create network after retries: %v", err)
				return nil // Fail fast if network creation fails
			}

			// Verify network exists and is ready
			if tj.network == nil || tj.network.Network == nil {
				tb.Logf("Network creation succeeded but network reference is nil")
				return nil // Fail fast if network validation fails
			}

			// Wait for network to stabilize
			time.Sleep(time.Second * 5)
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
			if err := tj.purgeResources(); err != nil {
				logger.Warnf("Failed to clean up resources during cleanup: %v", err)
			}
			if tj.network != nil && tj.pool != nil && tj.pool.Client != nil {
				if err := tj.pool.Client.RemoveNetwork(tj.network.Network.ID); err != nil {
					if !strings.Contains(err.Error(), "not found") {
						logger.Warnf("Failed to remove network %s: %v", tj.network.Network.ID, err)
					}
				}
				tj.network = nil
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

	// First, ensure cleanup of any existing networks
	networks, err := j.pool.Client.ListNetworks()
	if err != nil {
		j.tb.Logf("Error listing networks: %v", err)
		return []*dockertest.Network{}
	}

	// Clean up existing networks with retries
	for _, network := range networks {
		if network.Labels[runIDLabel] == j.runID || strings.HasPrefix(network.Name, "test-network") {
			// Disconnect all containers first
			containers, err := j.pool.Client.ListContainers(docker.ListContainersOptions{
				All:     true,
				Filters: map[string][]string{"network": {network.ID}},
			})
			if err != nil {
				j.tb.Logf("Warning: Failed to list containers for network %s: %v", network.ID, err)
				continue
			}

			for _, container := range containers {
				if err := j.pool.Client.DisconnectNetwork(network.ID, docker.NetworkConnectionOptions{
					Container: container.ID,
					Force:     true,
				}); err != nil {
					j.tb.Logf("Warning: Failed to disconnect container %s from network %s: %v",
						container.ID, network.ID, err)
				}
			}

			// Wait for disconnections to complete
			time.Sleep(time.Second * 2)

			if err := j.pool.Client.RemoveNetwork(network.ID); err != nil {
				j.tb.Logf("Warning: Failed to remove network %s: %v", network.ID, err)
			}
		}
	}

	// If we already have a network, return it
	if j.network != nil {
		return []*dockertest.Network{j.network}
	}

	// Create a new network with retries
	networkName := fmt.Sprintf("test-network-%s", j.runID)
	var network *docker.Network
	err = retry.WithBackoff(context.Background(), func(ctx context.Context) error {
		var createErr error
		network, createErr = j.pool.Client.CreateNetwork(docker.CreateNetworkOptions{
			Name: networkName,
			Labels: map[string]string{
				runIDLabel: j.runID,
				appLabel:   "jaeger-test",
			},
			Driver: "bridge",
			IPAM: &docker.IPAMOptions{
				Driver: "default",
				Config: []docker.IPAMConfig{
					{
						Subnet: "172.21.0.0/16",
					},
				},
			},
			CheckDuplicate: true,
		})
		if createErr != nil {
			j.tb.Logf("Network creation attempt failed: %v", createErr)
			return createErr
		}
		return nil
	},
		retry.WithMax(time.Second*30),
		retry.WithMaxAttempts(5))

	if err != nil {
		j.tb.Logf("Failed to create network after retries: %v", err)
		return []*dockertest.Network{}
	}

	j.network = &dockertest.Network{Network: network}
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

// purgeResources purges the resources from the pool and returns any errors encountered.
func (j *testJaeger) purgeResources() error {
	if j == nil {
		return fmt.Errorf("testJaeger instance is nil")
	}
	if j.pool == nil {
		return fmt.Errorf("Docker pool is nil")
	}
	if j.pool.Client == nil {
		return fmt.Errorf("Docker client is nil")
	}

	var cleanupMutex sync.Mutex
	var errs []error
	resources := j.getDockerizedResources()

	// Clean up ports first to ensure they're released
	if err := j.cleanupPorts(); err != nil {
		errs = append(errs, fmt.Errorf("port cleanup failed: %w", err))
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
				errs = append(errs, fmt.Errorf("failed to stop container %s: %w", resource.Container.ID, err))
				j.tb.Logf("Warning: Failed to stop container %s: %v", resource.Container.ID, err)
			}
		}
		cleanupMutex.Unlock()
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
				errs = append(errs, fmt.Errorf("failed to remove container %s: %w", resource.Container.ID, err))
				j.tb.Logf("Warning: Failed to remove container %s: %v", resource.Container.ID, err)
			}
		}
		cleanupMutex.Unlock()
		time.Sleep(time.Second)
	}

	// Finally, clean up networks with proper synchronization and retries
	if j.pool != nil && j.pool.Client != nil {
		networks, err := j.pool.Client.ListNetworks()
		if err != nil {
			errs = append(errs, fmt.Errorf("failed to list networks: %w", err))
			j.tb.Logf("Warning: Failed to list networks: %v", err)
		} else {
			for _, network := range networks {
				// Check network validity using ID and Name fields
				if network.ID == "" || network.Name == "" || network.Labels == nil {
					continue
				}
				if network.Labels[runIDLabel] == j.runID {
					cleanupMutex.Lock()
					// Disconnect all containers with retries
					containers, err := j.pool.Client.ListContainers(docker.ListContainersOptions{
						All:     true,
						Filters: map[string][]string{"network": {network.ID}},
					})
					if err != nil {
						errs = append(errs, fmt.Errorf("failed to list containers for network %s: %w", network.ID, err))
						j.tb.Logf("Warning: Failed to list containers for network %s: %v", network.ID, err)
						cleanupMutex.Unlock()
						continue
					}

					// Wait for network operations to stabilize
					time.Sleep(time.Second * 2)

					for _, container := range containers {
						// Retry disconnection up to 3 times with increasing backoff
						var disconnectErr error
						for attempt := 0; attempt < 3; attempt++ {
							disconnectErr = j.pool.Client.DisconnectNetwork(network.ID, docker.NetworkConnectionOptions{
								Container: container.ID,
								Force:     true,
							})
							if disconnectErr == nil || strings.Contains(disconnectErr.Error(), "No such container") {
								break
							}
							time.Sleep(time.Second * time.Duration(attempt+1))
						}
						if disconnectErr != nil && !strings.Contains(disconnectErr.Error(), "No such container") {
							errs = append(errs, fmt.Errorf("failed to disconnect container %s from network %s: %w", container.ID, network.ID, disconnectErr))
							j.tb.Logf("Warning: Failed to disconnect container %s from network %s: %v", container.ID, network.ID, disconnectErr)
						}
					}

					// Wait for disconnections to complete
					time.Sleep(time.Second * 3)

					// Retry network removal up to 5 times with increasing backoff
					var removeErr error
					for attempt := 0; attempt < 5; attempt++ {
						removeErr = j.pool.Client.RemoveNetwork(network.ID)
						if removeErr == nil || strings.Contains(removeErr.Error(), "not found") {
							removeErr = nil
							break
						}
						time.Sleep(time.Second * time.Duration(attempt+1))
					}
					if removeErr != nil {
						errs = append(errs, fmt.Errorf("failed to remove network %s: %w", network.ID, removeErr))
						j.tb.Logf("Warning: Failed to remove network %s: %v", network.ID, removeErr)
					}
					cleanupMutex.Unlock()
				}
			}
		}
	}

	// Final port cleanup
	if err := j.cleanupPorts(); err != nil {
		errs = append(errs, fmt.Errorf("final port cleanup failed: %w", err))
		j.tb.Logf("Warning: Failed in final port cleanup: %v", err)
	}

	// Wait for all cleanup operations to complete
	time.Sleep(time.Second * 3)

	if len(errs) > 0 {
		return fmt.Errorf("resource cleanup encountered errors: %v", errs)
	}
	return nil
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
