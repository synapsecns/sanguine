// Package main is the entry point for the STIP Relayer
package main

import (
	"os"

	"github.com/synapsecns/sanguine/services/stiprelayer/cmd"
	"github.com/synapsecns/sanguine/services/stiprelayer/metadata"
)

// main is the entry point for the stip relayer.
func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
