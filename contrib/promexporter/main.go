package main

import (
	"github.com/synapsecns/sanguine/contrib/promexporter/cmd"
	"github.com/synapsecns/sanguine/core/config"
	"os"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "promexporter", date)
	cmd.Start(os.Args, buildInfo)
}
