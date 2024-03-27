// Package main is the entrypoint to the sin-executor.
package main

import (
	"github.com/synapsecns/sanguine/sin-executor/cmd"
	"github.com/synapsecns/sanguine/sin-executor/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
