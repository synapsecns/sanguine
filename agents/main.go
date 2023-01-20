// Package main provides the main entry point for all agents.
package main

import (
	"github.com/synapsecns/sanguine/agents/cmd"
	"github.com/synapsecns/sanguine/core/config"
	"os"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "agents", date)

	cmd.Start(os.Args, buildInfo)
}
