package main

import (
	"github.com/synapsecns/sanguine/services/rfq/relayer/cmd"
	"github.com/synapsecns/sanguine/services/rfq/relayer/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
