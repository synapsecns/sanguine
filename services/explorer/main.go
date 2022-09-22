// Package main: entry point for the cli.
package main

import (
	"os"

	"github.com/synapsecns/sanguine/services/explorer/cmd"
)

func main() {
	cmd.Start(os.Args)
}
