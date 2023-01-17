// Package main provides the main entry point for the executor.
package main

import (
	"github.com/synapsecns/sanguine/core/config"
	"os"

	"github.com/synapsecns/sanguine/agents/agents/executor/cmd"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "executor", date)

	cmd.Start(os.Args, buildInfo)
}
