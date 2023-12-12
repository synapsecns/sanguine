package main

import (
	"os"

	"github.com/synapsecns/sanguine/rfq/rfq-relayer/cmd"
	"github.com/synapsecns/sanguine/rfq/rfq-relayer/metadata"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
