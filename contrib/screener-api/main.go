// Package main provides the entry point for the screener-api.
package main

import (
	"os"

	"github.com/synapsecns/sanguine/contrib/screener-api/cmd"
	"github.com/synapsecns/sanguine/contrib/screener-api/metadata"
)

//go:generate go run github.com/swaggo/swag/cmd/swag init

func main() {
	cmd.Start(os.Args, metadata.BuildInfo())
}
