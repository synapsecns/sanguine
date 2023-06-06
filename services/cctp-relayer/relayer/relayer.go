package relayer

import (
	"context"
	"fmt"

	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/scribe/config"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CCTPRelayer struct{}

func NewCCTPRelayer(ctx context.Context, cfg config.Config, handler metrics.Handler) (*CCTPRelayer, error) {
	scribeClient := cfg.GetScribeConfig()
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("%s:%d", scribeClient.URL, scribeClient.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
		grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor(otelgrpc.WithTracerProvider(handler.GetTracerProvider()))),
	)
	if err != nil {
		return nil, fmt.Errorf("could not dial grpc: %w", err)
	}

	grpcClient := pbscribe.NewScribeServiceClient(conn)

	// Ensure that gRPC is up and running.
	healthCheck, err := grpcClient.Check(ctx, &pbscribe.HealthCheckRequest{}, grpc.WaitForReady(true))
	if err != nil {
		return nil, fmt.Errorf("could not check: %w", err)
	}
	if healthCheck.Status != pbscribe.HealthCheckResponse_SERVING {
		return nil, fmt.Errorf("not serving: %s", healthCheck.Status)
	}

	return &CCTPRelayer{
		scribeClient,
		grpClient,
	}, nil
}
