// Package main is the entry point for the omnirpc service.
package main

import (
	"github.com/synapsecns/sanguine/services/omnirpc/cmd"
	"os"
)

func main() {
	cmd.Start(os.Args)
}
