package api

import (
	"context"
	"fmt"

	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/ginhelper"
	baseServer "github.com/synapsecns/sanguine/core/server"
	"golang.org/x/sync/errgroup"
)

var logger = log.Logger("guard-api")

// Start starts the api server.
func Start(parentCtx context.Context, metricsPort uint16) error {
	router := ginhelper.New(logger)

	g, ctx := errgroup.WithContext(parentCtx)

	g.Go(func() error {
		connection := baseServer.Server{}
		err := connection.ListenAndServe(ctx, fmt.Sprintf(":%d", metricsPort), router)
		if err != nil {
			return fmt.Errorf("could not start gqlServer: %w", err)
		}
		return nil
	})

	err := g.Wait()
	if err != nil {
		return fmt.Errorf("could not start api: %w", err)
	}

	return nil
}
