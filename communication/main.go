// Package main contains the entrypoint for the committee service.
package main

import (
	"github.com/synapsecns/sanguine/communication/cmd"
	"github.com/synapsecns/sanguine/communication/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
