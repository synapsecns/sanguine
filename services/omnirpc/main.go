// Package main is the entry point for the omnirpc service.
package main

import (
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/services/omnirpc/cmd"
	_ "go.uber.org/automaxprocs"
	"os"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "omnirpc", date)
	cmd.Start(os.Args, buildInfo)
}
