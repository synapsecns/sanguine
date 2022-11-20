package client

import (
	"context"
	"fmt"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/services/scribe/api"
)

// ScribeClient is a gRPC client to a Scribe.
type ScribeClient struct {
	// HTTPPort is the port the Scribe is listening on for HTTP requests.
	HTTPPort uint16
	// Database is the database type.
	Database string
	// Path is the path to the database or db connection.
	Path string
	// GRPCPort is the port the Scribe is listening on for gRPC requests.
	GRPCPort uint16
}

// EmbeddedScribe is a ScribeClient that is used locally.
type EmbeddedScribe struct {
	// ScribeClient is the ScribeClient.
	ScribeClient
}

// NewEmbeddedScribe creates a new EmbeddedScribe.
func NewEmbeddedScribe(database, path string) EmbeddedScribe {
	return EmbeddedScribe{
		ScribeClient: ScribeClient{
			HTTPPort: uint16(freeport.GetPort()),
			Database: database,
			Path:     path,
			GRPCPort: uint16(freeport.GetPort()),
		},
	}
}

// Start starts the EmbeddedScribe.
func (e EmbeddedScribe) Start(ctx context.Context) error {
	apiConfig := api.Config{
		HTTPPort: e.HTTPPort,
		Database: e.Database,
		Path:     e.Path,
		GRPCPort: e.GRPCPort,
	}
	err := api.Start(ctx, apiConfig)
	if err != nil {
		return fmt.Errorf("could not start api: %w", err)
	}

	return nil
}

// RemoteScribe is a ScribeClient that is used remotely.
type RemoteScribe struct {
	// ScribeClient is the ScribeClient.
	ScribeClient
}

// NewRemoteScribe creates a new RemoteScribe.
func NewRemoteScribe(httpPort uint16, database, path string, grpcPort uint16) RemoteScribe {
	return RemoteScribe{
		ScribeClient: ScribeClient{
			HTTPPort: httpPort,
			Database: database,
			Path:     path,
			GRPCPort: grpcPort,
		},
	}
}
