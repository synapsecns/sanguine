// Package main provides the main entry point for the sinner service.
package main

import (
	"github.com/synapsecns/sanguine/services/sinner/metadata"
	"os"

	"github.com/synapsecns/sanguine/services/sinner/cmd"
)

func main() {
	buildInfo := metadata.BuildInfo()
	cmd.Start(os.Args, buildInfo)
}
