package main

import (
	"github.com/synapsecns/sanguine/services/sin-explorer/cmd"
	"github.com/synapsecns/sanguine/services/sin-explorer/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
