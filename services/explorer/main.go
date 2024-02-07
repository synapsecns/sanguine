// Package main: entry point for the cli.
package main

import (
	"github.com/synapsecns/sanguine/services/explorer/metadata"
	"os"

	"github.com/synapsecns/sanguine/services/explorer/cmd"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
