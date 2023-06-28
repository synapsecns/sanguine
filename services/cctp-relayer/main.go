package main

import (
	"github.com/synapsecns/sanguine/services/cctp-relayer/cmd"
	"github.com/synapsecns/sanguine/services/cctp-relayer/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
