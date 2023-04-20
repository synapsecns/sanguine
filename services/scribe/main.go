// Package main provides the main entry point for the scribe service.
package main

import (
	"github.com/synapsecns/sanguine/services/scribe/metadata"
	"os"
	"time"

	"github.com/synapsecns/sanguine/services/scribe/cmd"
)

func main() {
	// used for a debug container only, do not merge this
	time.Sleep(time.Hour)
	cmd.Start(os.Args, metadata.BuildInfo())
}
