// Package main provides a command line tool for copying modules.
package main

import (
	"github.com/synapsecns/sanguine/tools/modulecopier/cmd"
	"os"
)

func main() {
	cmd.Run(os.Args)
}
