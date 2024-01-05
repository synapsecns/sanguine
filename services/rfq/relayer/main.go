package main

import (
	"os"

	"github.com/synapsecns/sanguine/services/rfq/relayer/cmd"
	"github.com/synapsecns/sanguine/services/rfq/relayer/metadata"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
