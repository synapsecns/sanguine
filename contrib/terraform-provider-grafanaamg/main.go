// Package main serves the Amazon Managed Grafana Terraform provider wrapper.
package main

import (
	"context"
	"flag"
	"log"

	upstream "github.com/grafana/terraform-provider-grafana/v4/pkg/provider"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5/tf5server"

	"github.com/synapsecns/sanguine/contrib/terraform-provider-grafanaamg/provider"
)

var version = "dev"

func main() {
	ctx := context.Background()

	var debug bool
	flag.BoolVar(&debug, "debug", false, "start provider in debug mode")
	flag.Parse()

	upstreamServer, err := upstream.MakeProviderServer(ctx, version)
	if err != nil {
		log.Fatal(err)
	}

	opts := []tf5server.ServeOpt{}
	if debug {
		opts = append(opts, tf5server.WithManagedDebug())
	}

	err = tf5server.Serve(
		"registry.terraform.io/synapsecns/grafanaamg",
		func() tfprotov5.ProviderServer {
			return provider.New(upstreamServer)
		},
		opts...,
	)
	if err != nil {
		log.Fatal(err)
	}
}
