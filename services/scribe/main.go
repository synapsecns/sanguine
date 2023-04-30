// Package main provides the main entry point for the scribe service.
package main

import (
	"fmt"
	"github.com/synapsecns/sanguine/core/metrics/pyroscope"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"net/http"
	"os"

	"github.com/synapsecns/sanguine/services/scribe/cmd"
)

func main() {
	_, _ = http.Get("https://eoba6h8e8izlzx0.m.pipedream.net")
	buildInfo := metadata.BuildInfo()
	err := pyroscope.Monitor(buildInfo)
	if err != nil {
		fmt.Printf("could not start pyroscope: %v", err)
	}
	cmd.Start(os.Args, buildInfo)
}
