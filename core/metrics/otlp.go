package metrics

import (
	"context"
	"fmt"
	"github.com/synapsecns/sanguine/core"
	"github.com/synapsecns/sanguine/core/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"os"
	"strings"
	"time"
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
	var client otlptrace.Client
	transport := transportFromString(core.GetEnv(otlpTransportEnv, otlpTransportGRPC.String()))
	switch transport {
	case otlpTransportHTTP:
		client = otlptracehttp.NewClient()
	case otlpTransportGRPC:
		client = otlptracegrpc.NewClient()
	default:
		return fmt.Errorf("unknown transport type: %s", os.Getenv(otlpTransportEnv))
	}

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		return fmt.Errorf("failed to create otlp exporter: %w", err)
	}

	n.baseHandler = newBaseHandler(n.buildInfo, tracesdk.WithBatcher(exporter, tracesdk.WithMaxQueueSize(1000000), tracesdk.WithMaxExportBatchSize(2000)), tracesdk.WithSampler(tracesdk.AlwaysSample()))

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
	otlpTransportEnv = "OTEL_EXPORTER_OTLP_TRANSPORT"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=otlpTransportType -linecomment
type otlpTransportType uint8

const (
	otlpTransportHTTP otlpTransportType = iota + 1 // http
	otlpTransportGRPC                              // grpc
)

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
