package server

import (
	pbscribe "github.com/synapsecns/sanguine/services/scribe/grpc/types/scribe"
	"google.golang.org/grpc"
)

// SetupGRPCServer sets up the grpc server.
func SetupGRPCServer() *grpc.Server {
	s := grpc.NewServer()
	sImpl := server{}

	pbscribe.RegisterLogServiceServer(s, &sImpl)

	return s
}

type server struct {
	pbscribe.UnimplementedLogServiceServer
}
