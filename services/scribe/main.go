// Package main provides the main entry point for the scribe service.
package main

import (
	"github.com/synapsecns/sanguine/core/config"
	"os"

	"github.com/synapsecns/sanguine/services/scribe/cmd"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "scribe", date)

	cmd.Start(os.Args, buildInfo)
}
