// Package main provides the entry point for the executor.
package main

import (
	"github.com/synapsecns/sanguine/agents/agents/executor/cmd"
	"github.com/synapsecns/sanguine/core/config"
	"os"
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
