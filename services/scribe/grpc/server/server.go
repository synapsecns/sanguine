package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/synapsecns/sanguine/services/scribe/db"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"google.golang.org/grpc"
	"time"
)

// SetupGRPCServer sets up the grpc server.
func SetupGRPCServer(ctx context.Context, engine *gin.Engine) (*grpc.Server, error) {
	s := grpc.NewServer()
	sImpl := server{}

	mux := runtime.NewServeMux()
	pbscribe.RegisterScribeServiceServer(s, &sImpl)
	err := pbscribe.RegisterScribeServiceHandlerServer(ctx, mux, &sImpl)
	if err != nil {
		return nil, fmt.Errorf("could not register server")
	}

	engine.NoRoute(func(c *gin.Context) {
		c.Status(200)
		gin.WrapF(mux.ServeHTTP)(c)
	})

	return s, nil
}

type server struct {
	db db.EventDB
	pbscribe.UnimplementedScribeServiceServer
}

func (s *server) FilterLogs(ctx context.Context, req *pbscribe.FilterLogsRequest) (*pbscribe.FilterLogsResponse, error) {
	req.Filter.ToNative()
	panic("")
}

func (s *server) Check(context.Context, *pbscribe.HealthCheckRequest) (*pbscribe.HealthCheckResponse, error) {
	return &pbscribe.HealthCheckResponse{Status: pbscribe.HealthCheckResponse_SERVING}, nil
}

func (s *server) Watch(a *pbscribe.HealthCheckRequest, res pbscribe.ScribeService_WatchServer) error {
	for {
		select {
		case <-res.Context().Done():
			return res.Context().Err()
		case <-time.After(time.Second):
			err := res.Send(&pbscribe.HealthCheckResponse{Status: pbscribe.HealthCheckResponse_SERVING})
			if err != nil {
				return fmt.Errorf("could not check response: %w", err)
			}
		}
	}
}
