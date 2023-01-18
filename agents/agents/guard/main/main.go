// Package main provides the main entry point for the guard.
package main

import (
	"os"

	"github.com/synapsecns/sanguine/core/config"

	"github.com/synapsecns/sanguine/agents/agents/guard/cmd"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "guard", date)

	cmd.Start(os.Args, buildInfo)
}
