// Package main spins up a explorer api and introspects the graphql api
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ipfs/go-log"
	"github.com/phayes/freeport"
	util2 "github.com/synapsecns/sanguine/contrib/promexporter/internal/gql/util"
	"github.com/synapsecns/sanguine/core/ginhelper"
	"github.com/synapsecns/sanguine/core/metrics"
	baseServer "github.com/synapsecns/sanguine/core/server"
	serverConfig "github.com/synapsecns/sanguine/services/explorer/config/server"
	gqlServer "github.com/synapsecns/sanguine/services/explorer/graphql/server"
	"github.com/synapsecns/sanguine/services/explorer/metadata"
	"os"
	"time"
)

var logger = log.Logger("explorer-api")

func main() {
	// ********************
	// GraphQL Server Start:
	// ********************

	// here we're going to create a temporary graphql server in the explorer we can extract the schema from there
	// rather than a live endpoint

	// create the context for the temporary server
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// prepare the server
	router := ginhelper.New(logger)

	nullHandler, err := metrics.NewByType(ctx, metadata.BuildInfo(), metrics.Null)
	if err != nil {
		panic(fmt.Errorf("error creating null handler, %w", err))
	}
	gqlServer.EnableGraphql(router, nil, nil, nil, nil, nil, nil, nil, serverConfig.Config{}, nullHandler)

	tmpPort, err := freeport.GetFreePort()
	if err != nil {
		panic(fmt.Errorf("could not get port: %w", err))
	}

	// start the server
	go func() {
		connection := baseServer.Server{}
		err := connection.ListenAndServe(ctx, fmt.Sprintf(":%d", tmpPort), router)
		if err != nil && !errors.Is(err, context.Canceled) {
			logger.Errorf("could not start gqlServer: %v", err)
		}
	}()

	err = util2.WaitForStart(ctx, tmpPort)
	if err != nil {
		panic(err)
	}

	// ********************
	// GQLGenc  Generation:
	// ********************

	const configURL = "contrib/promexporter/internal/gql/explorer/.gqlgenc.yaml"
	endpointURL := fmt.Sprintf("http://localhost:%d%s", tmpPort, gqlServer.GraphqlEndpoint)

	err = util2.GenerateGQLFromLocalServer(ctx, configURL, endpointURL)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		//nolint: gocritic
		os.Exit(3)
	}
}
