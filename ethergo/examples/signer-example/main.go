// Package main provides the command line interface for the example signer.
package main

import (
	"github.com/synapsecns/sanguine/ethergo/examples/signer-example/cmd"
	"github.com/synapsecns/sanguine/ethergo/examples/signer-example/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
