package main

import (
	"os"

	"github.com/synapsecns/sanguine/rfq/quoting-api/cmd"

	sanguineConfig "github.com/synapsecns/sanguine/core/config"
)

var (
	version = sanguineConfig.DefaultVersion
	commit  = sanguineConfig.DefaultCommit
	date    = sanguineConfig.DefaultDate
)

func main() {
	buildInfo := sanguineConfig.NewBuildInfo(version, commit, "quoter", date)
	cmd.Start(os.Args, buildInfo) 
}
