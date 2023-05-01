package client

import (
	"context"
	"fmt"
	"github.com/phayes/freeport"
	"github.com/synapsecns/sanguine/core/metrics"
	"github.com/synapsecns/sanguine/services/scribe/api"
)

// ScribeClient is a gRPC client to a Scribe.
type ScribeClient struct {
	// Port is the port the Scribe is listening on for HTTP requests.
	Port uint16
	// URL is the URL for the connection.
	URL string
	// metrics is the metrics metrics.
	metrics metrics.Handler
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
func NewEmbeddedScribe(database, path string, metrics metrics.Handler) EmbeddedScribe {
	return EmbeddedScribe{
		ScribeClient: ScribeClient{
			Port:    uint16(freeport.GetPort()),
			URL:     "localhost",
			metrics: metrics,
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
		Port:     e.Port,
		Database: e.database,
		Path:     e.path,
	}
	err := api.Start(ctx, apiConfig, e.ScribeClient.metrics)
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
func NewRemoteScribe(httpPort uint16, url string, metrics metrics.Handler) RemoteScribe {
	return RemoteScribe{
		ScribeClient: ScribeClient{
			Port:    httpPort,
			URL:     url,
			metrics: metrics,
		},
	}
}
