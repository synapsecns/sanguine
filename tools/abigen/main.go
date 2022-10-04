// Package main provides a command line tool for generating ABIs.
package main

import (
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/tools/abigen/cmd"
	"os"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "abigen", date)
	cmd.Run(os.Args, buildInfo)
}
