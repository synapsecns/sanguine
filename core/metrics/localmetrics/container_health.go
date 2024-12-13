package localmetrics

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/jaegertracing/jaeger/thrift-gen/jaeger"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

// generateJaegerThriftTrace creates a minimal Jaeger trace in Thrift format
func generateJaegerThriftTrace() []byte {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create a timestamp in microseconds
	timestamp := time.Now().UnixNano() / 1000

	// Create trace and span IDs
	traceIDLow := rand.Uint64()
	traceIDHigh := rand.Uint64()
	spanID := rand.Uint64()

	// Create process info first
	process := &jaeger.Process{
		ServiceName: "jaeger-health-check",
		Tags: []*jaeger.Tag{
			{
				Key:   "jaeger.version",
				VType: jaeger.TagType_STRING,
				VStr:  stringPtr("Go-2.30.0"),
			},
			{
				Key:   "hostname",
				VType: jaeger.TagType_STRING,
				VStr:  stringPtr("localhost"),
			},
		},
	}

	// Create a test span
	span := &jaeger.Span{
		TraceIdLow:    int64(traceIDLow),
		TraceIdHigh:   int64(traceIDHigh),
		SpanId:        int64(spanID),
		ParentSpanId:  0, // Root span
		OperationName: "health-check",
		References:    []*jaeger.SpanRef{},
		Flags:         1, // Sampled
		StartTime:     timestamp,
		Duration:      1000000, // 1 second in microseconds
		Tags: []*jaeger.Tag{
			{
				Key:   "span.kind",
				VType: jaeger.TagType_STRING,
				VStr:  stringPtr("client"),
			},
			{
				Key:   "sampler.type",
				VType: jaeger.TagType_STRING,
				VStr:  stringPtr("const"),
			},
			{
				Key:   "sampler.param",
				VType: jaeger.TagType_BOOL,
				VBool: boolPtr(true),
			},
		},
		Logs: []*jaeger.Log{},
	}

	// Create batch with explicit initialization
	batch := &jaeger.Batch{
		Process: process,
		Spans:   []*jaeger.Span{span},
	}

	// Use TMemoryBuffer for serialization
	memBuffer := thrift.NewTMemoryBufferLen(4096)
	protocol := thrift.NewTBinaryProtocolConf(memBuffer, &thrift.TConfiguration{
		MaxFrameSize:   16384000,
		MaxMessageSize: 16384000,
	})

	// Write the batch using explicit struct writing
	ctx := context.Background()
	if err := protocol.WriteStructBegin(ctx, "Batch"); err != nil {
		log.Printf("Failed to write struct begin: %v", err)
		return nil
	}

	// Write Process field (field 1)
	if err := protocol.WriteFieldBegin(ctx, "process", thrift.STRUCT, 1); err != nil {
		log.Printf("Failed to write process field begin: %v", err)
		return nil
	}
	if err := process.Write(ctx, protocol); err != nil {
		log.Printf("Failed to write process: %v", err)
		return nil
	}
	if err := protocol.WriteFieldEnd(ctx); err != nil {
		log.Printf("Failed to write process field end: %v", err)
		return nil
	}

	// Write Spans field (field 2)
	if err := protocol.WriteFieldBegin(ctx, "spans", thrift.LIST, 2); err != nil {
		log.Printf("Failed to write spans field begin: %v", err)
		return nil
	}
	if err := protocol.WriteListBegin(ctx, thrift.STRUCT, len(batch.Spans)); err != nil {
		log.Printf("Failed to write spans list begin: %v", err)
		return nil
	}
	for _, s := range batch.Spans {
		if err := s.Write(ctx, protocol); err != nil {
			log.Printf("Failed to write span: %v", err)
			return nil
		}
	}
	if err := protocol.WriteListEnd(ctx); err != nil {
		log.Printf("Failed to write spans list end: %v", err)
		return nil
	}
	if err := protocol.WriteFieldEnd(ctx); err != nil {
		log.Printf("Failed to write spans field end: %v", err)
		return nil
	}

	// Write struct end
	if err := protocol.WriteFieldStop(ctx); err != nil {
		log.Printf("Failed to write field stop: %v", err)
		return nil
	}
	if err := protocol.WriteStructEnd(ctx); err != nil {
		log.Printf("Failed to write struct end: %v", err)
		return nil
	}

	// Ensure protocol is flushed
	if err := protocol.Flush(ctx); err != nil {
		log.Printf("Failed to flush protocol: %v", err)
		return nil
	}

	// Get the serialized payload
	payload := memBuffer.Bytes()
	payloadSize := len(payload)

	// Create the final message with size header
	msg := make([]byte, 4+payloadSize)
	binary.BigEndian.PutUint32(msg[0:4], uint32(payloadSize))
	copy(msg[4:], payload)

	log.Printf("Generated trace batch of size: %d bytes (frame: 4, payload: %d)", len(msg), payloadSize)
	return msg
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

		// Quick check additional endpoints with detailed logging
		queryReady := isEndpointReady(queryEndpoint)
		if !queryReady {
			j.tb.Log("Query endpoint not ready (optional)")
		}

		// Skip OTLP gRPC endpoint check as it's not an HTTP endpoint
		j.tb.Log("Skipping OTLP gRPC endpoint check (not HTTP)")

		otlpHttpReady := isEndpointReady(otlpHttpEndpoint)
		if !otlpHttpReady {
			j.tb.Log("OTLP HTTP endpoint not ready (optional)")
		}

		// Log status of optional endpoints but don't fail if they're not ready
		if !queryReady || !otlpHttpReady {
			j.tb.Logf("Optional endpoints status - query: %v, otlp-http: %v (after %v)",
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
		Timeout: time.Second * 10,
	}

	// Determine endpoint type and prepare request
	var req *http.Request
	var err error

	// Check if this is a collector endpoint
	isCollector := strings.Contains(endpoint, "/api/traces")
	isHealth := strings.Contains(endpoint, "/health")

	if isCollector {
		// Generate and send a test trace
		traceData := generateJaegerThriftTrace()
		if traceData == nil {
			log.Printf("Failed to generate trace data for collector endpoint: %s", endpoint)
			return false
		}

		req, err = http.NewRequest(http.MethodPost, endpoint, bytes.NewReader(traceData))
		if err != nil {
			log.Printf("Failed to create request for collector endpoint %s: %v", endpoint, err)
			return false
		}

		// Set required headers for Jaeger collector
		req.Header.Set("Content-Type", "application/vnd.apache.thrift.binary")
		req.Header.Set("User-Agent", "jaeger-go/2.30.0")
		log.Printf("Sending trace to collector %s with headers: %v", endpoint, req.Header)
		log.Printf("Trace payload size: %d bytes", len(traceData))
	} else {
		// For other endpoints, just do a GET request
		req, err = http.NewRequest(http.MethodGet, endpoint, nil)
		if err != nil {
			log.Printf("Failed to create request for endpoint %s: %v", endpoint, err)
			return false
		}
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to connect to endpoint %s: %v", endpoint, err)
		return false
	}
	defer resp.Body.Close()

	// Read response body for debugging
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Failed to read response body from %s: %v", endpoint, err)
	}

	// Log response details
	log.Printf("Response from %s - Status: %d, Headers: %v", endpoint, resp.StatusCode, resp.Header)
	if len(body) > 0 {
		log.Printf("Response body from %s: %s", endpoint, string(body))
	}

	// For collector endpoint, accept 202 Accepted
	if isCollector {
		if resp.StatusCode != http.StatusAccepted {
			log.Printf("Collector endpoint %s returned unexpected status %d with headers: %v",
				endpoint, resp.StatusCode, resp.Header)
			return false
		}
		log.Printf("Successfully submitted trace to collector %s", endpoint)
		return true
	}

	// For health endpoint, require 200 OK
	if isHealth {
		if resp.StatusCode != http.StatusOK {
			log.Printf("Health endpoint %s returned unexpected status %d: %s",
				endpoint, resp.StatusCode, string(body))
			return false
		}
		return true
	}

	// For other endpoints (UI, etc.), accept any 2xx status code
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Printf("Endpoint %s returned non-2xx status %d: %s",
			endpoint, resp.StatusCode, string(body))
		return false
	}

	return true
}
