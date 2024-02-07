// Package main is the entry point for the omnirpc service.
package main

import (
	"github.com/synapsecns/sanguine/services/omnirpc/cmd"
	"github.com/synapsecns/sanguine/services/omnirpc/metadata"
	_ "go.uber.org/automaxprocs"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
