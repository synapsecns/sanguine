// Package main: entry point for the cli.
package main

import (
	"github.com/ipfs/go-log"
	"github.com/synapsecns/sanguine/core/config"
	"github.com/synapsecns/sanguine/core/metrics/pyroscope"
	"os"

	"github.com/synapsecns/sanguine/services/explorer/cmd"
)

var (
	version = config.DefaultVersion
	commit  = config.DefaultCommit
	date    = config.DefaultDate
)

var logger = log.Logger("main-logger")

func main() {
	buildInfo := config.NewBuildInfo(version, commit, "explorer", date)
	err := pyroscope.Monitor(buildInfo)
	if err != nil {
		logger.Warnf("could not start pyroscope: %v", err)
	}
	cmd.Start(os.Args, buildInfo)
}
