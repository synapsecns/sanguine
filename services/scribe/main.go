// Package main provides the main entry point for the scribe service.
package main

import (
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"os"

	"github.com/synapsecns/sanguine/services/scribe/cmd"
)

func main() {
	buildInfo := metadata.BuildInfo()
	cmd.Start(os.Args, buildInfo)
}
