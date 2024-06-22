// Package main provides the entry point for the opbot.
package main

import (
	"os"

	"github.com/synapsecns/sanguine/contrib/opbot/cmd"
	"github.com/synapsecns/sanguine/contrib/opbot/metadata"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
