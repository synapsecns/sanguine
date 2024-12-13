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

	// Create process info with all required fields explicitly set
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
			{
				Key:   "client-uuid",
				VType: jaeger.TagType_STRING,
				VStr:  stringPtr("test-client"),
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

	// Create batch
	batch := &jaeger.Batch{
		Process: process,
		Spans:   []*jaeger.Span{span},
	}

	// Log batch details for debugging
	log.Printf("Batch details - Process: %+v", batch.Process)
	log.Printf("Batch details - Process ServiceName: %s", batch.Process.ServiceName)
	log.Printf("Batch details - Process Tags count: %d", len(batch.Process.Tags))
	log.Printf("Batch details - Spans count: %d", len(batch.Spans))

	// Use TMemoryBuffer for serialization with TBinaryProtocol
	memBuffer := thrift.NewTMemoryBufferLen(4096)
	protocol := thrift.NewTBinaryProtocolConf(memBuffer, &thrift.TConfiguration{
		MaxFrameSize:   16384000,
		MaxMessageSize: 16384000,
	})

	// Write Thrift message envelope for emitBatch
	ctx := context.Background()
	if err := protocol.WriteMessageBegin(ctx, "emitBatch", thrift.CALL, 1); err != nil {
		log.Printf("Failed to write message begin: %v", err)
		return nil
	}

	// Write emitBatch_args struct begin
	if err := protocol.WriteStructBegin(ctx, "emitBatch_args"); err != nil {
		log.Printf("Failed to write args struct begin: %v", err)
		return nil
	}

	// Write batch field (field 1)
	if err := protocol.WriteFieldBegin(ctx, "batch", thrift.STRUCT, 1); err != nil {
		log.Printf("Failed to write batch field begin: %v", err)
		return nil
	}

	// Write batch struct begin
	if err := protocol.WriteStructBegin(ctx, "Batch"); err != nil {
		log.Printf("Failed to write batch struct begin: %v", err)
		return nil
	}

	// Write Process field (field 1)
	if err := protocol.WriteFieldBegin(ctx, "process", thrift.STRUCT, 1); err != nil {
		log.Printf("Failed to write process field begin: %v", err)
		return nil
	}

	// Write Process struct
	if err := protocol.WriteStructBegin(ctx, "Process"); err != nil {
		log.Printf("Failed to write process struct begin: %v", err)
		return nil
	}

	// Write ServiceName field (field 1)
	if err := protocol.WriteFieldBegin(ctx, "serviceName", thrift.STRING, 1); err != nil {
		log.Printf("Failed to write serviceName field begin: %v", err)
		return nil
	}

	if err := protocol.WriteString(ctx, batch.Process.ServiceName); err != nil {
		log.Printf("Failed to write serviceName: %v", err)
		return nil
	}

	if err := protocol.WriteFieldEnd(ctx); err != nil {
		log.Printf("Failed to write serviceName field end: %v", err)
		return nil
	}

	// Write Tags field (field 2)
	if err := protocol.WriteFieldBegin(ctx, "tags", thrift.LIST, 2); err != nil {
		log.Printf("Failed to write tags field begin: %v", err)
		return nil
	}

	// Write tags list
	if err := protocol.WriteListBegin(ctx, thrift.STRUCT, len(batch.Process.Tags)); err != nil {
		log.Printf("Failed to write tags list begin: %v", err)
		return nil
	}

	for _, tag := range batch.Process.Tags {
		// Write Tag struct begin
		if err := protocol.WriteStructBegin(ctx, "Tag"); err != nil {
			log.Printf("Failed to write tag struct begin: %v", err)
			return nil
		}

		// Write key field (field 1)
		if err := protocol.WriteFieldBegin(ctx, "key", thrift.STRING, 1); err != nil {
			log.Printf("Failed to write tag key field begin: %v", err)
			return nil
		}
		if err := protocol.WriteString(ctx, tag.Key); err != nil {
			log.Printf("Failed to write tag key: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write tag key field end: %v", err)
			return nil
		}

		// Write vType field (field 2)
		if err := protocol.WriteFieldBegin(ctx, "vType", thrift.I32, 2); err != nil {
			log.Printf("Failed to write tag vType field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI32(ctx, int32(tag.VType)); err != nil {
			log.Printf("Failed to write tag vType: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write tag vType field end: %v", err)
			return nil
		}

		// Write value field based on vType
		switch tag.VType {
		case jaeger.TagType_STRING:
			if tag.VStr != nil {
				if err := protocol.WriteFieldBegin(ctx, "vStr", thrift.STRING, 3); err != nil {
					log.Printf("Failed to write tag vStr field begin: %v", err)
					return nil
				}
				if err := protocol.WriteString(ctx, *tag.VStr); err != nil {
					log.Printf("Failed to write tag vStr: %v", err)
					return nil
				}
				if err := protocol.WriteFieldEnd(ctx); err != nil {
					log.Printf("Failed to write tag vStr field end: %v", err)
					return nil
				}
			}
		case jaeger.TagType_BOOL:
			if tag.VBool != nil {
				if err := protocol.WriteFieldBegin(ctx, "vBool", thrift.BOOL, 5); err != nil {
					log.Printf("Failed to write tag vBool field begin: %v", err)
					return nil
				}
				if err := protocol.WriteBool(ctx, *tag.VBool); err != nil {
					log.Printf("Failed to write tag vBool: %v", err)
					return nil
				}
				if err := protocol.WriteFieldEnd(ctx); err != nil {
					log.Printf("Failed to write tag vBool field end: %v", err)
					return nil
				}
			}
		}

		// End Tag struct
		if err := protocol.WriteFieldStop(ctx); err != nil {
			log.Printf("Failed to write tag field stop: %v", err)
			return nil
		}
		if err := protocol.WriteStructEnd(ctx); err != nil {
			log.Printf("Failed to write tag struct end: %v", err)
			return nil
		}
	}

	if err := protocol.WriteListEnd(ctx); err != nil {
		log.Printf("Failed to write tags list end: %v", err)
		return nil
	}

	if err := protocol.WriteFieldEnd(ctx); err != nil {
		log.Printf("Failed to write tags field end: %v", err)
		return nil
	}

	// End Process struct
	if err := protocol.WriteFieldStop(ctx); err != nil {
		log.Printf("Failed to write process field stop: %v", err)
		return nil
	}

	if err := protocol.WriteStructEnd(ctx); err != nil {
		log.Printf("Failed to write process struct end: %v", err)
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

	// Write spans list
	if err := protocol.WriteListBegin(ctx, thrift.STRUCT, len(batch.Spans)); err != nil {
		log.Printf("Failed to write spans list begin: %v", err)
		return nil
	}

	for _, span := range batch.Spans {
		// Write Span struct begin
		if err := protocol.WriteStructBegin(ctx, "Span"); err != nil {
			log.Printf("Failed to write span struct begin: %v", err)
			return nil
		}

		// Write traceIdLow (field 1)
		if err := protocol.WriteFieldBegin(ctx, "traceIdLow", thrift.I64, 1); err != nil {
			log.Printf("Failed to write traceIdLow field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI64(ctx, span.TraceIdLow); err != nil {
			log.Printf("Failed to write traceIdLow: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write traceIdLow field end: %v", err)
			return nil
		}

		// Write traceIdHigh (field 2)
		if err := protocol.WriteFieldBegin(ctx, "traceIdHigh", thrift.I64, 2); err != nil {
			log.Printf("Failed to write traceIdHigh field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI64(ctx, span.TraceIdHigh); err != nil {
			log.Printf("Failed to write traceIdHigh: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write traceIdHigh field end: %v", err)
			return nil
		}

		// Write spanId (field 3)
		if err := protocol.WriteFieldBegin(ctx, "spanId", thrift.I64, 3); err != nil {
			log.Printf("Failed to write spanId field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI64(ctx, span.SpanId); err != nil {
			log.Printf("Failed to write spanId: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write spanId field end: %v", err)
			return nil
		}

		// Write parentSpanId (field 4)
		if err := protocol.WriteFieldBegin(ctx, "parentSpanId", thrift.I64, 4); err != nil {
			log.Printf("Failed to write parentSpanId field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI64(ctx, span.ParentSpanId); err != nil {
			log.Printf("Failed to write parentSpanId: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write parentSpanId field end: %v", err)
			return nil
		}

		// Write operationName (field 5)
		if err := protocol.WriteFieldBegin(ctx, "operationName", thrift.STRING, 5); err != nil {
			log.Printf("Failed to write operationName field begin: %v", err)
			return nil
		}
		if err := protocol.WriteString(ctx, span.OperationName); err != nil {
			log.Printf("Failed to write operationName: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write operationName field end: %v", err)
			return nil
		}

		// Write flags (field 7)
		if err := protocol.WriteFieldBegin(ctx, "flags", thrift.I32, 7); err != nil {
			log.Printf("Failed to write flags field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI32(ctx, span.Flags); err != nil {
			log.Printf("Failed to write flags: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write flags field end: %v", err)
			return nil
		}

		// Write startTime (field 8)
		if err := protocol.WriteFieldBegin(ctx, "startTime", thrift.I64, 8); err != nil {
			log.Printf("Failed to write startTime field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI64(ctx, span.StartTime); err != nil {
			log.Printf("Failed to write startTime: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write startTime field end: %v", err)
			return nil
		}

		// Write duration (field 9)
		if err := protocol.WriteFieldBegin(ctx, "duration", thrift.I64, 9); err != nil {
			log.Printf("Failed to write duration field begin: %v", err)
			return nil
		}
		if err := protocol.WriteI64(ctx, span.Duration); err != nil {
			log.Printf("Failed to write duration: %v", err)
			return nil
		}
		if err := protocol.WriteFieldEnd(ctx); err != nil {
			log.Printf("Failed to write duration field end: %v", err)
			return nil
		}

		// Write optional tags (field 10)
		if len(span.Tags) > 0 {
			if err := protocol.WriteFieldBegin(ctx, "tags", thrift.LIST, 10); err != nil {
				log.Printf("Failed to write span tags field begin: %v", err)
				return nil
			}
			if err := protocol.WriteListBegin(ctx, thrift.STRUCT, len(span.Tags)); err != nil {
				log.Printf("Failed to write span tags list begin: %v", err)
				return nil
			}
			for _, tag := range span.Tags {
				// Write Tag struct begin
				if err := protocol.WriteStructBegin(ctx, "Tag"); err != nil {
					log.Printf("Failed to write tag struct begin: %v", err)
					return nil
				}

				// Write key field (field 1)
				if err := protocol.WriteFieldBegin(ctx, "key", thrift.STRING, 1); err != nil {
					log.Printf("Failed to write tag key field begin: %v", err)
					return nil
				}
				if err := protocol.WriteString(ctx, tag.Key); err != nil {
					log.Printf("Failed to write tag key: %v", err)
					return nil
				}
				if err := protocol.WriteFieldEnd(ctx); err != nil {
					log.Printf("Failed to write tag key field end: %v", err)
					return nil
				}

				// Write vType field (field 2)
				if err := protocol.WriteFieldBegin(ctx, "vType", thrift.I32, 2); err != nil {
					log.Printf("Failed to write tag vType field begin: %v", err)
					return nil
				}
				if err := protocol.WriteI32(ctx, int32(tag.VType)); err != nil {
					log.Printf("Failed to write tag vType: %v", err)
					return nil
				}
				if err := protocol.WriteFieldEnd(ctx); err != nil {
					log.Printf("Failed to write tag vType field end: %v", err)
					return nil
				}

				// Write value field based on vType
				switch tag.VType {
				case jaeger.TagType_STRING:
					if tag.VStr != nil {
						if err := protocol.WriteFieldBegin(ctx, "vStr", thrift.STRING, 3); err != nil {
							log.Printf("Failed to write tag vStr field begin: %v", err)
							return nil
						}
						if err := protocol.WriteString(ctx, *tag.VStr); err != nil {
							log.Printf("Failed to write tag vStr: %v", err)
							return nil
						}
						if err := protocol.WriteFieldEnd(ctx); err != nil {
							log.Printf("Failed to write tag vStr field end: %v", err)
							return nil
						}
					}
				case jaeger.TagType_BOOL:
					if tag.VBool != nil {
						if err := protocol.WriteFieldBegin(ctx, "vBool", thrift.BOOL, 5); err != nil {
							log.Printf("Failed to write tag vBool field begin: %v", err)
							return nil
						}
						if err := protocol.WriteBool(ctx, *tag.VBool); err != nil {
							log.Printf("Failed to write tag vBool: %v", err)
							return nil
						}
						if err := protocol.WriteFieldEnd(ctx); err != nil {
							log.Printf("Failed to write tag vBool field end: %v", err)
							return nil
						}
					}
				}

				// End Tag struct
				if err := protocol.WriteFieldStop(ctx); err != nil {
					log.Printf("Failed to write tag field stop: %v", err)
					return nil
				}
				if err := protocol.WriteStructEnd(ctx); err != nil {
					log.Printf("Failed to write tag struct end: %v", err)
					return nil
				}
			}
			if err := protocol.WriteListEnd(ctx); err != nil {
				log.Printf("Failed to write span tags list end: %v", err)
				return nil
			}
			if err := protocol.WriteFieldEnd(ctx); err != nil {
				log.Printf("Failed to write span tags field end: %v", err)
				return nil
			}
		}

		// End Span struct
		if err := protocol.WriteFieldStop(ctx); err != nil {
			log.Printf("Failed to write span field stop: %v", err)
			return nil
		}
		if err := protocol.WriteStructEnd(ctx); err != nil {
			log.Printf("Failed to write span struct end: %v", err)
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

	// End batch struct
	if err := protocol.WriteFieldStop(ctx); err != nil {
		log.Printf("Failed to write batch field stop: %v", err)
		return nil
	}

	if err := protocol.WriteStructEnd(ctx); err != nil {
		log.Printf("Failed to write batch struct end: %v", err)
		return nil
	}

	// End emitBatch_args struct
	if err := protocol.WriteFieldStop(ctx); err != nil {
		log.Printf("Failed to write args field stop: %v", err)
		return nil
	}
	if err := protocol.WriteStructEnd(ctx); err != nil {
		log.Printf("Failed to write args struct end: %v", err)
		return nil
	}

	// End message
	if err := protocol.WriteMessageEnd(ctx); err != nil {
		log.Printf("Failed to write message end: %v", err)
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
		req.Header.Set("Content-Type", "application/x-thrift")
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
