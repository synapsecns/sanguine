// Package main provides the main entry point for the scribe service.
package main

import (
	"github.com/synapsecns/sanguine/services/scribe/cmd"
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
