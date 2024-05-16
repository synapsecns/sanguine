// Package main contains the entrypoint for the committee devnet.
package main

import (
	"os"

	"github.com/synapsecns/sanguine/committee/devnet/cmd"

	"github.com/synapsecns/sanguine/committee/devnet/metadata"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
