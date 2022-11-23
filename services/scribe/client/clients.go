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
	// GRPCPort is the port the Scribe is listening on for gRPC requests.
	GRPCPort uint16
	// URL is the URL for the connection.
	URL string
}

// EmbeddedScribe is a ScribeClient that is used locally.
type EmbeddedScribe struct {
	// ScribeClient is the ScribeClient.
	ScribeClient
	// database is the database type.
	database string
	// path is the path to the database or db connection.
	path string
}

// NewEmbeddedScribe creates a new EmbeddedScribe.
func NewEmbeddedScribe(database, path string) EmbeddedScribe {
	return EmbeddedScribe{
		ScribeClient: ScribeClient{
			HTTPPort: uint16(freeport.GetPort()),
			GRPCPort: uint16(freeport.GetPort()),
			URL:      "localhost",
		},
		database: database,
		path:     path,
	}
}

// OverrideURL overrides the URL for the RemoteScribe.
func (r RemoteScribe) OverrideURL(url string) {
	r.URL = url
}

// Start starts the EmbeddedScribe.
func (e EmbeddedScribe) Start(ctx context.Context) error {
	apiConfig := api.Config{
		HTTPPort: e.HTTPPort,
		Database: e.database,
		Path:     e.path,
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
func NewRemoteScribe(httpPort uint16, grpcPort uint16, url string) RemoteScribe {
	return RemoteScribe{
		ScribeClient: ScribeClient{
			HTTPPort: httpPort,
			GRPCPort: grpcPort,
			URL:      url,
		},
	}
}
