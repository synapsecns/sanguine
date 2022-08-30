package main

import (
	"github.com/synapsecns/sanguine/services/scribe/cmd"
	"os"
)

func main() {
	cmd.Start(os.Args)
}
