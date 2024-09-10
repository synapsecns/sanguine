package metrics

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

type otlpHandler struct {
	*baseHandler
	buildInfo config.BuildInfo
}

// NewOTLPMetricsHandler creates a new newrelic metrics handler.
func NewOTLPMetricsHandler(buildInfo config.BuildInfo) Handler {
	return &otlpHandler{
		buildInfo:   buildInfo,
		baseHandler: newBaseHandler(buildInfo),
	}
}

func (n *otlpHandler) Start(ctx context.Context) (err error) {
	var exporters []tracesdk.SpanExporter

	primaryExporter, err := makeOTLPExporter(ctx, otelTransportEnv, otelEndpointEnv)
	if err != nil {
		return fmt.Errorf("could not create default client: %w", err)
	}
	exporters = append(exporters, primaryExporter)

	// Loop to create additional exporters
	for i := 1; ; i++ {
		envSuffix := fmt.Sprintf("%d", i)
		transportEnv := otelTransportEnv + envSuffix
		endpointEnv := otelEndpointEnv + envSuffix

		// no more transports to add.
		if !core.HasEnv(endpointEnv) {
			break
		}

		exporter, err := makeOTLPExporter(ctx, transportEnv, endpointEnv)
		if err != nil {
			if err != nil {
				return fmt.Errorf("could not create exporter %d: %v", i, err)
			}
		}

		exporters = append(exporters, exporter)
	}

	// create the multi-exporter with optional secondary exporter
	multiExporter := NewMultiExporter(exporters...)

	n.baseHandler = newBaseHandler(
		n.buildInfo,
		tracesdk.WithBatcher(
			multiExporter,
			tracesdk.WithMaxQueueSize(defaultMaxQueueSize),
			tracesdk.WithMaxExportBatchSize(defaultMaxExportBatch),
		),
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
	)

	// start the new parent
	err = n.baseHandler.Start(ctx)
	if err != nil {
		return fmt.Errorf("could not start base handler: %w", err)
	}

	go func() {
		handleShutdown(ctx, n.baseHandler.unwrappedTP)
	}()

	return nil
}

func (n *otlpHandler) Type() HandlerType {
	return OTLP
}

// wait for the context to be canceled.
// then flush the traces and shutdown the exporter.
func handleShutdown(ctx context.Context, provider *tracesdk.TracerProvider) {
	<-ctx.Done()

	const spanWaitTime = time.Millisecond
	const shutdownAllowance = time.Second * 10

	// allow only 10 seconds for graceful shutdown.
	// we use without cancel to copy the parents values while making sure our derived context is not canceled.
	shutdownCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), shutdownAllowance)
	defer cancel()

	// don't shutdown immediately, wait for a bit to allow the last spans to be sent. This is in process and should be aymptotic to instant.
	time.Sleep(spanWaitTime)

	err := provider.ForceFlush(shutdownCtx)
	if err != nil {
		logger.Warnf("could not flush traces: %v", err)
	}
	err = provider.Shutdown(shutdownCtx)
	if err != nil {
		logger.Warnf("could not shutdown traces: %v", err)
	}
}

const (
	otelEndpointEnv  = "OTEL_EXPORTER_OTLP_ENDPOINT"
	otelTransportEnv = "OTEL_EXPORTER_OTLP_TRANSPORT"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=otlpTransportType -linecomment
type otlpTransportType uint8

const (
	otlpTransportHTTP otlpTransportType = iota + 1 // http
	otlpTransportGRPC                              // grpc
)

// makeOTLPTrace creates a new OTLP client based on the transport type and url.
func makeOTLPExporter(ctx context.Context, transportEnv, urlEnv string) (*otlptrace.Exporter, error) {
	transport := transportFromString(core.GetEnv(transportEnv, otlpTransportHTTP.String()))
	url := os.Getenv(urlEnv)

	oteltraceClient, err := buildClientFromTransport(
		transport,
		url,
	)
	if err != nil {
		return nil, fmt.Errorf("could not create client from transport: %w", err)
	}

	exporter, err := otlptrace.New(ctx, oteltraceClient)
	if err != nil {
		return nil, fmt.Errorf("ocould not create client: %w", err)
	}
	return exporter, nil
}

// buildClientFromTransport creates a new OTLP client based on the transport type and url.
func buildClientFromTransport(transport otlpTransportType, url string) (otlptrace.Client, error) {
	if url == "" {
		return nil, fmt.Errorf("no url specified")
	}

	switch transport {
	case otlpTransportHTTP:
		return otlptracehttp.NewClient(otlptracehttp.WithEndpointURL(url)), nil
	case otlpTransportGRPC:
		return otlptracegrpc.NewClient(otlptracegrpc.WithEndpointURL(url)), nil
	default:
		return nil, fmt.Errorf("unknown transport type: %s", transport.String())
	}
}

// transportFromString converts a string to a transport type.
// Defaults to http if the string is not recognized.
func transportFromString(transport string) otlpTransportType {
	switch strings.ToLower(transport) {
	case otlpTransportHTTP.String():
		return otlpTransportHTTP
	case otlpTransportGRPC.String():
		return otlpTransportGRPC
	}
	// will be unknown since we use iota+1
	// (see uber's go stye guide for details)
	return otlpTransportType(0)
}

const (
	defaultMaxQueueSize   = 1000000
	defaultMaxExportBatch = 2000
)
