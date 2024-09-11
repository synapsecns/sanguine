package metrics

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials"
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

	primaryExporter, err := makeOTLPExporter(ctx, "")
	if err != nil {
		return fmt.Errorf("could not create default client: %w", err)
	}
	exporters = append(exporters, primaryExporter)

	// Loop to create additional exporters
	for i := 1; ; i++ {
		envSuffix := fmt.Sprintf("_%d", i)
		// if this is empty we can assume no config exists at all.
		endpointEnv := otelEndpointEnv + envSuffix

		// no more transports to add.
		if !core.HasEnv(endpointEnv) {
			break
		}

		exporter, err := makeOTLPExporter(ctx, envSuffix)
		if err != nil {
			return fmt.Errorf("could not create exporter %d: %v", i, err)
		}

		exporters = append(exporters, exporter)
	}

	// create the multi-exporter with all the exporters
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
	otelInsecureEvn  = "OTEL_EXPORTER_OTLP_INSECURE_MODE"
	otelHeadersEnv   = "OTEL_EXPORTER_OTLP_HEADERS"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=otlpTransportType -linecomment
type otlpTransportType uint8

const (
	otlpTransportHTTP otlpTransportType = iota + 1 // http
	otlpTransportGRPC                              // grpc
)

// getEnvSuffix returns the value of an environment variable with a suffix.
func getEnvSuffix(env, suffix, defaultRet string) string {
	newEnv := env + suffix
	return core.GetEnv(newEnv, defaultRet)
}

// makeOTLPTrace creates a new OTLP client based on the transport type and url.
func makeOTLPExporter(ctx context.Context, envSuffix string) (*otlptrace.Exporter, error) {
	transport := transportFromString(getEnvSuffix(otelTransportEnv, envSuffix, otlpTransportHTTP.String()))
	url := getEnvSuffix(otelEndpointEnv, envSuffix, "")
	insecure := getEnvSuffix(otelInsecureEvn, envSuffix, "false")
	headers := getEnvSuffix(otelHeadersEnv, envSuffix, "")

	if url == "" {
		return nil, fmt.Errorf("could not create exporter: url is empty")
	}

	oteltraceClient, err := buildClientFromTransport(
		transport,
		WithURL(url),
		// defaults to true
		WithInsecure(insecure == "false"),
		WithHeaders(headers),
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
func buildClientFromTransport(transport otlpTransportType, options ...Option) (otlptrace.Client, error) {
	to := transportOptions{}
	for _, option := range options {
		if err := option(&to); err != nil {
			return nil, fmt.Errorf("could not apply option: %w", err)
		}
	}

	// TODO: make sure url is validated

	switch transport {
	case otlpTransportHTTP:
		return otlptracehttp.NewClient(to.httpOptions...), nil
	case otlpTransportGRPC:
		return otlptracegrpc.NewClient(to.grpcOptions...), nil
	default:
		return nil, fmt.Errorf("unknown transport type: %s", transport.String())
	}
}

type transportOptions struct {
	// httpOptions are the options for the http transport.
	httpOptions []otlptracehttp.Option
	// grpcOptions are the options for the grpc transport.
	grpcOptions []otlptracegrpc.Option
}

// Option Each option appends the correspond option for both http and grpc options.
// only one will be used in creating the actual client.
type Option func(*transportOptions) error

func WithURL(url string) Option {
	return func(o *transportOptions) error {
		o.httpOptions = append(o.httpOptions, otlptracehttp.WithEndpointURL(url))
		o.grpcOptions = append(o.grpcOptions, otlptracegrpc.WithEndpointURL(url))

		return nil
	}
}

func WithInsecure(isInsecure bool) Option {
	return func(o *transportOptions) error {
		if isInsecure {
			o.httpOptions = append(o.httpOptions, otlptracehttp.WithInsecure())
			o.grpcOptions = append(o.grpcOptions, otlptracegrpc.WithInsecure())
		} else {
			tlsCreds := credentials.NewClientTLSFromCert(nil, "")
			// note: you do not need to specify the tls creds for http, this happens automatically when https:// is used as the protocol scheme.
			o.grpcOptions = append(o.grpcOptions, otlptracegrpc.WithTLSCredentials(tlsCreds))
		}

		return nil
	}
}

func WithHeaders(headers string) Option {
	return func(o *transportOptions) error {
		realHeaders := headersToMap(headers)
		o.httpOptions = append(o.httpOptions, otlptracehttp.WithHeaders(realHeaders))
		o.grpcOptions = append(o.grpcOptions, otlptracegrpc.WithHeaders(realHeaders))
		return nil
	}
}

func headersToMap(input string) map[string]string {
	// Initialize the map
	result := make(map[string]string)

	// Split the input string by comma to get key=value pairs
	pairs := strings.Split(input, ",")

	// Iterate over each pair
	for _, pair := range pairs {
		// Split each pair by '=' to get the key and value
		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 2 {
			key := kv[0]
			value := kv[1]
			// Add the key and value to the map
			result[key] = value
		}
	}

	return result
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
