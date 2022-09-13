package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/synapsecns/sanguine/services/scribe/db"
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/types/v1"
	"google.golang.org/grpc"
)

// SetupGRPCServer sets up the grpc server.
func SetupGRPCServer(ctx context.Context, engine *gin.Engine) (*grpc.Server, error) {
	s := grpc.NewServer()
	sImpl := server{}

	mux := runtime.NewServeMux()
	pbscribe.RegisterLogServiceServer(s, &sImpl)
	err := pbscribe.RegisterLogServiceHandlerServer(ctx, mux, &sImpl)
	if err != nil {
		return nil, fmt.Errorf("could not register server")
	}

	engine.NoRoute(func(c *gin.Context) {
		mux.ServeHTTP(c.Writer, c.Request)
	})

	return s, nil
}

type server struct {
	db db.EventDB
	pbscribe.UnimplementedLogServiceServer
}

func (s *server) FilterLogs(ctx context.Context, req *pbscribe.FilterLogsRequest) (*pbscribe.FilterLogsResponse, error) {
	req.Filter.ToNative()
	panic("")
}
