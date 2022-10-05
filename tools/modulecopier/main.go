// Package main provides a command line tool for copying modules.
package main

import (
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/tools/modulecopier/cmd"
	"os"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "modulecopier", date)

	cmd.Run(os.Args, buildInfo)
}
