package localmetrics

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/jaegertracing/jaeger/thrift-gen/jaeger"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// generateJaegerThriftTrace creates a minimal Jaeger trace in Thrift format
func generateJaegerThriftTrace() []byte {
	// Create a timestamp in microseconds
	timestamp := time.Now().UnixNano() / 1000

	// Create a test span
	span := &jaeger.Span{
		TraceIdLow:    rand.Int63(),
		TraceIdHigh:   rand.Int63(),
		SpanId:        rand.Int63(),
		OperationName: "test-operation",
		StartTime:     timestamp,
		Duration:      1000, // 1ms
		Tags: []*jaeger.Tag{
			{
				Key:   "test.key",
				VType: jaeger.TagType_STRING,
				VStr:  stringPtr("test-value"),
			},
		},
	}

	// Create a batch with the span
	batch := &jaeger.Batch{
		Process: &jaeger.Process{
			ServiceName: "test-service",
			Tags: []*jaeger.Tag{
				{
					Key:   "service.version",
					VType: jaeger.TagType_STRING,
					VStr:  stringPtr("1.0.0"),
				},
			},
		},
		Spans: []*jaeger.Span{span},
	}

	// Serialize the batch to Thrift
	transport := thrift.NewTMemoryBufferLen(1024)
	protocol := thrift.NewTBinaryProtocolTransport(transport)

	if err := batch.Write(context.Background(), protocol); err != nil {
		log.Printf("Failed to serialize Jaeger batch: %v", err)
		return nil
	}

	return transport.Bytes()
}

// waitForContainerHealth waits for the container to be healthy
func (j *testJaeger) waitForContainerHealth(resource *dockertest.Resource) error {
	startTime := time.Now()
	j.tb.Log("Starting container health check...")

	// Initial warmup period - increased to allow for full initialization
	j.tb.Log("Waiting for initial container warmup...")
	time.Sleep(time.Second * 15)
	// Check container status and get initial logs
	container := resource.Container
	if container == nil {
		return fmt.Errorf("container reference is nil")
	}

	// Get initial container logs immediately after startup
	var buf bytes.Buffer
	err := j.pool.Client.Logs(docker.LogsOptions{
		Container:    container.ID,
		OutputStream: &buf,
		Follow:       false,
		Stdout:       true,
		Stderr:       true,
		Timestamps:   true,
		Since:        0,
	})
	if err != nil {
		j.tb.Logf("Warning: Failed to get initial container logs: %v", err)
	} else if buf.Len() > 0 {
		j.tb.Logf("Initial container logs:\n%s", buf.String())
	}

	// Use pool's retry mechanism with timeout and detailed logging
	if err := j.pool.Retry(func() error {
		// Add delay between retries
		time.Sleep(time.Second * 5)

		j.tb.Logf("Container status: %s (running for %v)", container.State.Status, time.Since(startTime))
		if container.State.Status != "running" {
			return fmt.Errorf("container is not running, status: %s", container.State.Status)
		}

		// Try to connect to all endpoints with increased timeouts
		collectorEndpoint := fmt.Sprintf("http://127.0.0.1:%s/api/traces", resource.GetPort("14268/tcp"))
		queryEndpoint := fmt.Sprintf("http://127.0.0.1:%s", resource.GetPort("16686/tcp"))
		healthEndpoint := fmt.Sprintf("http://127.0.0.1:%s/health", resource.GetPort("14269/tcp"))
		otlpGrpcEndpoint := fmt.Sprintf("http://127.0.0.1:%s", resource.GetPort("4317/tcp"))
		otlpHttpEndpoint := fmt.Sprintf("http://127.0.0.1:%s", resource.GetPort("4318/tcp"))

		j.tb.Logf("Checking endpoints - collector: %s, query: %s, health: %s, otlp-grpc: %s, otlp-http: %s",
			collectorEndpoint, queryEndpoint, healthEndpoint, otlpGrpcEndpoint, otlpHttpEndpoint)

		// Check health endpoint first with increased retries and longer intervals
		healthReady := false
		maxRetries := 15 // Increased from 10
		for i := 0; i < maxRetries; i++ {
			healthReady = isEndpointReady(healthEndpoint)
			if healthReady {
				j.tb.Log("Health endpoint is ready")
				break
			}
			j.tb.Logf("Health check attempt %d/%d failed, waiting before retry...", i+1, maxRetries)

			// Get container logs after each failed attempt
			if resource.Container != nil {
				var buf bytes.Buffer
				err := j.pool.Client.Logs(docker.LogsOptions{
					Container:    resource.Container.ID,
					OutputStream: &buf,
					Follow:       false,
					Stdout:       true,
					Stderr:       true,
				})
				if err == nil && buf.Len() > 0 {
					j.tb.Logf("Container logs after failed attempt %d:\n%s", i+1, buf.String())
				}
			}

			time.Sleep(time.Second * 5) // Increased from 3
		}
		if !healthReady {
			j.tb.Log("Health endpoint not ready")
			return fmt.Errorf("health endpoint not ready (waited %v)", time.Since(startTime))
		}

		// Now check collector endpoint with increased retries
		collectorReady := false
		for i := 0; i < maxRetries; i++ {
			collectorReady = isEndpointReady(collectorEndpoint)
			if collectorReady {
				j.tb.Log("Collector endpoint is ready")
				break
			}
			j.tb.Logf("Collector check attempt %d/%d failed, waiting before retry...", i+1, maxRetries)
			time.Sleep(time.Second * 5) // Increased from 3
		}
		if !collectorReady {
			j.tb.Log("Collector endpoint not ready")
			return fmt.Errorf("collector endpoint not ready (waited %v)", time.Since(startTime))
		}

		// Quick check all endpoints with detailed logging
		queryReady := isEndpointReady(queryEndpoint)
		if !queryReady {
			j.tb.Log("Query endpoint not ready")
		}

		// Skip OTLP gRPC endpoint check as it's not an HTTP endpoint
		j.tb.Log("Skipping OTLP gRPC endpoint check (not HTTP)")

		otlpHttpReady := isEndpointReady(otlpHttpEndpoint)
		if !otlpHttpReady {
			j.tb.Log("OTLP HTTP endpoint not ready")
		}

		if !queryReady || !otlpHttpReady {
			// Get container logs on failure using Docker API
			if resource.Container != nil {
				var buf bytes.Buffer
				err := j.pool.Client.Logs(docker.LogsOptions{
					Container:    resource.Container.ID,
					OutputStream: &buf,
					Follow:       false,
					Stdout:       true,
					Stderr:       true,
				})
				if err == nil && buf.Len() > 0 {
					j.tb.Logf("Container logs: %s", buf.String())
				} else if err != nil {
					j.tb.Logf("Failed to get container logs: %v", err)
				}
			}
			return fmt.Errorf("endpoints not ready - query: %v, otlp-http: %v (waited %v)",
				queryReady, otlpHttpReady, time.Since(startTime))
		}

		j.tb.Logf("Container health check passed after %v", time.Since(startTime))
		return nil
	}); err != nil {
		return fmt.Errorf("container health check failed: %v", err)
	}

	return nil
}

// isEndpointReady performs a quick check if an endpoint is responding
func isEndpointReady(endpoint string) bool {
	client := &http.Client{
		Timeout: time.Second * 10, // Increased timeout for endpoint checks
	}

	// For the collector endpoint, we need to send a POST request with Jaeger format
	if strings.Contains(endpoint, "/api/traces") {
		payload := generateJaegerThriftTrace()
		if payload == nil {
			log.Printf("Failed to generate Jaeger trace payload")
			return false
		}

		req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
		if err != nil {
			log.Printf("Failed to create request: %v", err)
			return false
		}

		// Set the correct content type for Thrift binary encoding
		req.Header.Set("Content-Type", "application/vnd.apache.thrift.binary")
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Failed to send request: %v", err)
			return false
		}
		defer resp.Body.Close()

		// Read response body for debugging
		body, err := io.ReadAll(resp.Body)
		if err == nil {
			log.Printf("Collector endpoint response: status=%d, body=%s", resp.StatusCode, string(body))
		}

		// Jaeger accepts both 200 and 202 status codes for traces
		return resp.StatusCode == http.StatusOK || resp.StatusCode == http.StatusAccepted
	}

	// For health endpoints, require 200 OK
	if strings.Contains(endpoint, "/health") {
		resp, err := client.Get(endpoint)
		if err != nil {
			log.Printf("Health check failed for %s: %v", endpoint, err)
			return false
		}
		defer resp.Body.Close()

		// Read and log response body for debugging
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Failed to read response body: %v", err)
			return false
		}
		log.Printf("Health check response from %s: status=%d, body=%s", endpoint, resp.StatusCode, string(body))

		return resp.StatusCode == http.StatusOK
	}

	// For other endpoints, use GET request
	resp, err := client.Get(endpoint)
	if err != nil {
		log.Printf("Endpoint check failed for %s: %v", endpoint, err)
		return false
	}
	defer resp.Body.Close()

	// Read and log response for debugging
	body, err := io.ReadAll(resp.Body)
	if err == nil {
		log.Printf("Endpoint response from %s: status=%d, body=%s", endpoint, resp.StatusCode, string(body))
	}

	// For other endpoints, any response below 500 is considered ready
	return resp.StatusCode < 500
}
