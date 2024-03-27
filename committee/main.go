// Package main contains the entrypoint for the committee service.
package main

import (
	"github.com/synapsecns/sanguine/committee/cmd"
	"github.com/synapsecns/sanguine/committee/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
