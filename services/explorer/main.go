// Package main: entry point for the cli.
package main

import (
	"github.com/synapsecns/sanguine/core/config"
	"os"

	"github.com/synapsecns/sanguine/services/explorer/cmd"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "explorer", date)
	cmd.Start(os.Args, buildInfo)
}
