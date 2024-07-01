package main

import (
	"os"

	"github.com/synapsecns/sanguine/services/rfq/tools/withdrawal/cmd"
	"github.com/synapsecns/sanguine/services/rfq/tools/withdrawal/metadata"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
