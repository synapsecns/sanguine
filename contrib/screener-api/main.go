// Package main provides the entry point for the screener-api.
package main

import (
	"github.com/synapsecns/sanguine/contrib/screener-api/cmd"
	"github.com/synapsecns/sanguine/contrib/screener-api/metadata"
	"os"
)

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
