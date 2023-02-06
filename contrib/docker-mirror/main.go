// Package main is the entry point for the docker mirror tool.
package main

import (
	"github.com/synapsecns/sanguine/contrib/docker-mirror/cmd"
	"github.com/synapsecns/sanguine/core/config"
	_ "go.uber.org/automaxprocs"
	"os"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "docker-mirror", date)
	err := cmd.Start(os.Args, buildInfo)
	if err != nil {
		panic(err)
	}
}
